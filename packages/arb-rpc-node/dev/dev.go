/*
* Copyright 2021, Offchain Labs, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package dev

import (
	"context"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	core2 "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var logger = log.With().Caller().Stack().Str("component", "dev").Logger()

func NewDevNode(ctx context.Context, dir string, arbosPath string, chainId *big.Int, agg common.Address, initialL1Height uint64) (*Backend, *txdb.TxDB, func(), <-chan error, error) {
	nodeCacheConfig := configuration.NodeCache{
		AllowSlowLookup: true,
		LRUSize:         1000,
		TimedExpire:     20 * time.Minute,
	}
	coreConfig := configuration.DefaultCoreSettings()

	mon, err := monitor.NewMonitor(dir, arbosPath, coreConfig)
	if err != nil {
		return nil, nil, nil, nil, errors.Wrap(err, "error opening monitor")
	}

	backendCore, err := NewBackendCore(ctx, mon.Core, chainId)
	if err != nil {
		mon.Close()
		return nil, nil, nil, nil, err
	}

	db, errChan, err := txdb.New(ctx, mon.Core, mon.Storage.GetNodeStore(), 10*time.Millisecond, &nodeCacheConfig)
	if err != nil {
		mon.Close()
		return nil, nil, nil, nil, errors.Wrap(err, "error opening txdb")
	}

	cancel := func() {
		db.Close()
		mon.Close()
	}
	signer := types.NewEIP155Signer(chainId)
	l1 := NewL1Emulator(initialL1Height)
	backend := NewBackend(ctx, backendCore, db, l1, signer, agg, big.NewInt(100000000000))

	return backend, db, cancel, errChan, nil
}

type EVM struct {
	backend   *Backend
	snapshots map[uint64]uint64 // message count to block count
}

func NewEVM(backend *Backend) *EVM {
	return &EVM{
		backend:   backend,
		snapshots: make(map[uint64]uint64),
	}
}

func (s *EVM) Snapshot() (hexutil.Uint64, error) {
	messageCount, err := s.backend.arbcore.GetMessageCount()
	if err != nil {
		return 0, err
	}
	latestHeight := s.backend.l1Emulator.LatestHeight()
	s.snapshots[messageCount.Uint64()] = latestHeight
	logger.Info().Uint64("count", messageCount.Uint64()).Msg("created snapshot")
	return hexutil.Uint64(messageCount.Uint64()), nil
}

func (s *EVM) Revert(snapId hexutil.Uint64) error {
	messageCount := uint64(snapId)
	logger.Info().Uint64("snap", messageCount).Msg("revert")
	blockCount, ok := s.snapshots[messageCount]
	if !ok {
		return errors.New("no such snapshot")
	}
	err := s.backend.Reorg(messageCount, blockCount)
	if err != nil {
		logger.Error().Err(err).Msg("can't revert")
	}
	return err
}

func (s *EVM) Mine(timestamp *hexutil.Uint64) error {
	if timestamp != nil {
		s.backend.l1Emulator.SetTime(int64(*timestamp))
	}
	_, err := s.backend.AddInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), common.Address{})
	return err
}

func (s *EVM) IncreaseTime(amount int64) (string, error) {
	s.backend.l1Emulator.IncreaseTime(amount)
	_, err := s.backend.AddInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), common.Address{})
	return strconv.FormatInt(amount, 10), err
}

type BackendCore struct {
	ctx          context.Context
	arbcore      core.ArbCore
	chainID      *big.Int
	delayedCount *big.Int
}

func NewBackendCore(ctx context.Context, arbcore core.ArbCore, chainID *big.Int) (*BackendCore, error) {
	delayedCount, err := arbcore.GetTotalDelayedMessagesSequenced()
	if err != nil {
		return nil, err
	}
	return &BackendCore{
		ctx:          ctx,
		arbcore:      arbcore,
		chainID:      chainID,
		delayedCount: delayedCount,
	}, nil
}

func (b *BackendCore) addInboxMessage(msg message.Message, sender common.Address, gasPrice *big.Int, block L1BlockInfo) (common.Hash, error) {
	chainTime := inbox.ChainTime{
		BlockNum:  block.blockId.Height,
		Timestamp: block.timestamp,
	}
	msgCount, err := b.arbcore.GetMessageCount()
	if err != nil {
		return common.Hash{}, err
	}
	inboxMessage := message.NewInboxMessage(msg, sender, new(big.Int).Set(msgCount), gasPrice, chainTime)

	requestId := message.CalculateRequestId(b.chainID, msgCount)
	var prevHash common.Hash
	if msgCount.Cmp(big.NewInt(0)) > 0 {
		prevHash, err = b.arbcore.GetInboxAcc(new(big.Int).Sub(msgCount, big.NewInt(1)))
		if err != nil {
			return common.Hash{}, err
		}
	}
	seqBatchItem := inbox.NewSequencerItem(b.delayedCount, inboxMessage, prevHash)
	nextBlockMessage := inbox.InboxMessage{
		Kind:        6,
		Sender:      common.Address{},
		InboxSeqNum: new(big.Int).Add(msgCount, big.NewInt(1)),
		GasPrice:    big.NewInt(0),
		Data:        []byte{},
		ChainTime: inbox.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		},
	}
	nextBlockBatchItem := inbox.NewSequencerItem(b.delayedCount, nextBlockMessage, seqBatchItem.Accumulator)
	err = core.DeliverMessagesAndWait(b.arbcore, msgCount, prevHash, []inbox.SequencerBatchItem{seqBatchItem, nextBlockBatchItem}, nil, nil)
	if err != nil {
		return common.Hash{}, err
	}
	for {
		if b.arbcore.MachineIdle() {
			break
		}
		select {
		case <-b.ctx.Done():
			return [32]byte{}, errors.New("dev node canceled")
		case <-time.After(time.Millisecond * 200):
		}

	}
	for {
		cursorPos, err := b.arbcore.LogsCursorPosition(big.NewInt(0))
		if err != nil {
			return common.Hash{}, err
		}
		coreLogs, err := b.arbcore.GetLogCount()
		if err != nil {
			return common.Hash{}, err
		}
		if cursorPos.Cmp(coreLogs) == 0 {
			break
		}
		select {
		case <-b.ctx.Done():
			return [32]byte{}, errors.New("dev node canceled")
		case <-time.After(time.Millisecond * 200):
		}
	}

	return requestId, nil
}

type Backend struct {
	sync.Mutex
	*BackendCore
	ctx               context.Context
	db                *txdb.TxDB
	l1Emulator        *L1Emulator
	signer            types.Signer
	currentAggregator common.Address
	chainAggregator   common.Address
	l1GasPrice        *big.Int

	newTxFeed event.Feed
}

func NewBackend(ctx context.Context, core *BackendCore, db *txdb.TxDB, l1 *L1Emulator, signer types.Signer, aggregator common.Address, l1GasPrice *big.Int) *Backend {
	return &Backend{
		BackendCore:       core,
		ctx:               ctx,
		db:                db,
		l1Emulator:        l1,
		signer:            signer,
		currentAggregator: aggregator,
		chainAggregator:   aggregator,
		l1GasPrice:        l1GasPrice,
	}
}

func (b *Backend) ExportData() ([]byte, error) {
	b.Lock()
	messageCount, err := b.arbcore.GetMessageCount()
	if err != nil {
		return nil, err
	}
	messages, err := b.arbcore.GetMessages(big.NewInt(0), messageCount)
	if err != nil {
		return nil, err
	}
	b.Unlock()
	return inbox.TestVectorJSON(messages, nil, nil)
}

func (b *Backend) Reorg(messageCount, blockCount uint64) error {
	b.Lock()
	defer b.Unlock()
	return b.reorg(messageCount, blockCount)
}

func (b *Backend) reorg(messageCount, blockCount uint64) error {
	b.l1Emulator.Reorg(blockCount)
	logger.Info().Uint64("message", messageCount).Uint64("block", blockCount).Msg("Reorged chain")
	if err := core.ReorgAndWait(b.arbcore, new(big.Int).SetUint64(messageCount)); err != nil {
		return err
	}
	return b.waitForBlockCount(blockCount)
}

func (b *Backend) waitForBlockCount(blockCount uint64) error {
	for {
		blocks, err := b.db.BlockCount()
		if err != nil {
			return err
		}
		if blocks == blockCount {
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
	return nil
}

func (b *Backend) PendingTransactionCount(_ context.Context, _ common.Address) (*uint64, error) {
	b.Lock()
	defer b.Unlock()
	return nil, nil
}

func (b *Backend) SendTransaction(_ context.Context, tx *types.Transaction) error {
	b.Lock()
	defer b.Unlock()
	arbTx := message.NewCompressedECDSAFromEth(tx)
	sender, err := types.Sender(b.signer, tx)
	if err != nil {
		return err
	}

	arbMsg, err := message.NewTransactionBatchFromMessages([]message.AbstractL2Message{arbTx})
	if err != nil {
		return err
	}

	logger.
		Info().
		Uint64("gasLimit", tx.Gas()).
		Str("gasPrice", tx.GasPrice().String()).
		Uint64("nonce", tx.Nonce()).
		Str("from", sender.Hex()).
		Str("value", tx.Value().String()).
		Hex("hash", tx.Hash().Bytes()).
		Msg("sent transaction")

	startHeight := b.l1Emulator.LatestHeight()
	startCount, err := b.arbcore.GetMessageCount()
	if err != nil {
		return err
	}

	block := b.l1Emulator.GenerateBlock()
	if _, err := b.addInboxMessage(message.NewSafeL2Message(arbMsg), b.currentAggregator, b.l1GasPrice, block); err != nil {
		return err
	}
	if err := b.waitForBlockCount(block.blockId.Height.AsInt().Uint64()); err != nil {
		return err
	}
	txHash := common.NewHashFromEth(tx.Hash())
	res, err := b.db.GetRequest(txHash)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("tx res not found")
	}

	if res.ResultCode != evm.ReturnCode {
		logger.Warn().Int("code", int(res.ResultCode)).Msg("transaction failed")
		// If transaction failed, rollback the block
		if err := b.reorg(startCount.Uint64(), startHeight); err != nil {
			return err
		}

		// Insert an empty block instead
		block := b.l1Emulator.GenerateBlock()
		if _, err := b.addInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), b.currentAggregator, b.l1GasPrice, block); err != nil {
			return err
		}

		return evm.HandleCallError(res, true)
	}

	return nil
}

func (b *Backend) Aggregator() *common.Address {
	return &b.chainAggregator
}

func (b *Backend) AddInboxMessage(msg message.Message, sender common.Address) (common.Hash, error) {
	b.Lock()
	defer b.Unlock()
	return b.addInboxMessage(msg, sender, big.NewInt(0), b.l1Emulator.GenerateBlock())
}

func (b *Backend) SubscribeNewTxsEvent(ch chan<- core2.NewTxsEvent) event.Subscription {
	b.Lock()
	defer b.Unlock()
	return b.newTxFeed.Subscribe(ch)
}

func (b *Backend) PendingSnapshot() (*snapshot.Snapshot, error) {
	b.Lock()
	defer b.Unlock()
	return nil, nil
}

func (b *Backend) Start(ctx context.Context) {
}

type L1BlockInfo struct {
	blockId   *common.BlockId
	timestamp *big.Int
}

type L1Emulator struct {
	sync.Mutex
	timeIncrease int64
	latestHeight uint64
}

func NewL1Emulator(initialHeight uint64) *L1Emulator {
	b := &L1Emulator{
		latestHeight: initialHeight,
	}
	b.addBlock()
	return b
}

func (b *L1Emulator) LatestHeight() uint64 {
	return b.latestHeight
}

func (b *L1Emulator) Reorg(height uint64) {
	b.Lock()
	defer b.Unlock()
	b.latestHeight = height
}

func (b *L1Emulator) addBlock() L1BlockInfo {
	info := L1BlockInfo{
		blockId: &common.BlockId{
			Height:     common.NewTimeBlocksInt(int64(b.latestHeight + 1)),
			HeaderHash: common.RandHash(),
		},
		timestamp: big.NewInt(time.Now().Unix() + b.timeIncrease),
	}
	b.latestHeight += 1
	return info
}

func (b *L1Emulator) GenerateBlock() L1BlockInfo {
	b.Lock()
	defer b.Unlock()
	return b.addBlock()
}

func (b *L1Emulator) SetTime(timestamp int64) {
	b.Lock()
	defer b.Unlock()
	b.timeIncrease = timestamp - time.Now().Unix()
}

func (b *L1Emulator) IncreaseTime(amount int64) {
	b.Lock()
	defer b.Unlock()
	if amount < 0 {
		amount = 0
	}
	b.timeIncrease += amount
}

func EnableFees(srv *aggregator.Server, ownerAuth *bind.TransactOpts, aggregator ethcommon.Address) error {

	client := web3.NewEthClient(srv, true)
	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, client)
	if err != nil {
		return errors.Wrap(err, "error connecting to arb owner")
	}

	tx, err := arbOwner.SetFairGasPriceSender(ownerAuth, aggregator, true)
	if err != nil {
		return errors.Wrap(err, "error calling SetFairGasPriceSender")
	}
	_, err = ethbridge.WaitForReceiptWithResultsSimple(context.Background(), client, tx.Hash())
	if err != nil {
		return errors.Wrap(err, "error getting SetFairGasPriceSender receipt")
	}
	_, err = arbOwner.SetChainParameter(ownerAuth, arbos.FeesEnabledParamId, big.NewInt(1))
	if err != nil {
		return errors.Wrap(err, "error calling SetChainParameter")
	}
	return nil
}
