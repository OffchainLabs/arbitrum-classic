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

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	core2 "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/staker"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

var logger = log.With().Caller().Stack().Str("component", "dev").Logger()

func NewDevNode(ctx context.Context, dir string, arbosPath string, params protocol.ChainParams, owner common.Address, config []message.ChainConfigOption) (*Backend, *txdb.TxDB, common.Address, func(), <-chan error, error) {
	initMsg, err := message.NewInitMessage(params, owner, config)
	if err != nil {
		return nil, nil, [20]byte{}, nil, nil, err
	}

	rollupAddress := common.RandAddress()
	signer := types.NewEIP155Signer(message.ChainAddressToID(rollupAddress))

	aggregator := common.RandAddress()
	for i := range config {
		opt := config[len(config)-1-i]
		if aggConfig, ok := opt.(message.DefaultAggConfig); ok {
			aggregator = aggConfig.Aggregator
			break
		}
	}

	monitor, err := staker.NewMonitor(dir, arbosPath)
	if err != nil {
		return nil, nil, [20]byte{}, nil, nil, errors.Wrap(err, "error opening monitor")
	}

	l1 := NewL1Emulator()
	backendCore := NewBackendCore(monitor.Core, signer.ChainID())

	db, errChan, err := txdb.New(ctx, monitor.Core, monitor.Storage.GetNodeStore(), rollupAddress, 10*time.Millisecond)
	if err != nil {
		monitor.Close()
		return nil, nil, [20]byte{}, nil, nil, errors.Wrap(err, "error opening txdb")
	}

	if _, err := backendCore.addInboxMessage(initMsg, rollupAddress, big.NewInt(0), l1.GenerateBlock()); err != nil {
		monitor.Close()
		return nil, nil, [20]byte{}, nil, nil, errors.Wrap(err, "error adding init message to inbox")
	}

	cancel := func() {
		db.Close()
		monitor.Close()
	}

	backend := NewBackend(backendCore, db, l1, signer, aggregator, big.NewInt(100000000000))

	return backend, db, rollupAddress, cancel, errChan, nil
}

type EVM struct {
	backend *Backend
}

func NewEVM(backend *Backend) *EVM {
	return &EVM{backend: backend}
}

func (s *EVM) Snapshot() (hexutil.Uint64, error) {
	logCount, err := s.backend.arbcore.GetLogCount()
	if err != nil {
		return 0, err
	}
	sendCount, err := s.backend.arbcore.GetSendCount()
	if err != nil {
		return 0, err
	}
	latestHeight := s.backend.l1Emulator.Latest().blockId.Height.AsInt().Uint64()
	logger.
		Info().
		Uint64("latest", latestHeight).
		Uint64("logcount", logCount.Uint64()).
		Uint64("sendCount", sendCount.Uint64()).
		Msg("created snapshot")
	return hexutil.Uint64(latestHeight), nil
}

func (s *EVM) Revert(snapId hexutil.Uint64) error {
	logger.Info().Uint64("snap", uint64(snapId)).Msg("revert")
	err := s.backend.Reorg(uint64(snapId))
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
	arbcore core.ArbCore
	chainID *big.Int
}

func NewBackendCore(arbcore core.ArbCore, chainID *big.Int) *BackendCore {
	return &BackendCore{
		arbcore: arbcore,
		chainID: chainID,
	}
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
		prevHash, err = b.arbcore.GetInboxAcc(msgCount.Sub(msgCount, big.NewInt(1)))
		if err != nil {
			return common.Hash{}, err
		}
	}
	successful, err := core.DeliverMessagesAndWait(b.arbcore, []inbox.InboxMessage{inboxMessage}, prevHash, true)
	if err != nil {
		return common.Hash{}, err
	}
	if !successful {
		return common.Hash{}, errors.New("failed to deliver message")
	}
	for {
		if b.arbcore.MachineIdle() {
			break
		}
		<-time.After(time.Millisecond * 1000)
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
		<-time.After(time.Millisecond * 200)
	}

	return requestId, nil
}

type Backend struct {
	sync.Mutex
	*BackendCore
	db         *txdb.TxDB
	l1Emulator *L1Emulator
	signer     types.Signer
	aggregator common.Address
	l1GasPrice *big.Int

	newTxFeed event.Feed
}

func NewBackend(core *BackendCore, db *txdb.TxDB, l1 *L1Emulator, signer types.Signer, aggregator common.Address, l1GasPrice *big.Int) *Backend {
	return &Backend{
		BackendCore: core,
		db:          db,
		l1Emulator:  l1,
		signer:      signer,
		aggregator:  aggregator,
		l1GasPrice:  l1GasPrice,
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

func (b *Backend) Reorg(height uint64) error {
	b.Lock()
	defer b.Unlock()
	return b.reorg(height)
}

func (b *Backend) reorg(height uint64) error {
	startHeight, err := b.db.BlockCount()
	if err != nil {
		return err
	}
	b.l1Emulator.Reorg(height)
	newLatest, err := b.db.BlockCount()
	if err != nil {
		return err
	}
	logger.
		Info().
		Uint64("start", startHeight).
		Uint64("end", newLatest).
		Uint64("height", height).
		Msg("Reorged chain")
	if err := core.ReorgAndWait(b.arbcore, new(big.Int).SetUint64(height)); err != nil {
		return err
	}
	afterCount, err := b.arbcore.GetMessageCount()
	if err != nil {
		return err
	}
	if afterCount.Uint64() != height {
		panic("wrong after count")
	}
	return nil
}

// Return nil if no pending transaction count is available
func (b *Backend) PendingTransactionCount(_ context.Context, _ common.Address) *uint64 {
	b.Lock()
	defer b.Unlock()
	return nil
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
	startHeight := b.l1Emulator.Latest().blockId.Height.AsInt().Uint64()
	block := b.l1Emulator.GenerateBlock()
	if _, err := b.addInboxMessage(message.NewSafeL2Message(arbMsg), b.aggregator, b.l1GasPrice, block); err != nil {
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
		if err := b.reorg(startHeight); err != nil {
			return err
		}

		// Insert an empty block instead
		block := b.l1Emulator.GenerateBlock()
		if _, err := b.addInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), b.aggregator, b.l1GasPrice, block); err != nil {
			return err
		}

		return web3.HandleCallError(res, true)
	}

	return nil
}

func (b *Backend) Aggregator() *common.Address {
	return &b.aggregator
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

// Return nil if no pending snapshot is available
func (b *Backend) PendingSnapshot() (*snapshot.Snapshot, error) {
	b.Lock()
	defer b.Unlock()
	return nil, nil
}

type L1BlockInfo struct {
	blockId   *common.BlockId
	timestamp *big.Int
}

type L1Emulator struct {
	sync.Mutex
	l1Blocks       []L1BlockInfo
	l1BlocksByHash map[common.Hash]L1BlockInfo
	timeIncrease   int64
}

func NewL1Emulator() *L1Emulator {
	b := &L1Emulator{
		l1BlocksByHash: make(map[common.Hash]L1BlockInfo),
	}
	b.addBlock()
	return b
}

func (b *L1Emulator) Latest() L1BlockInfo {
	return b.l1Blocks[uint64(len(b.l1Blocks))-1]
}

func (b *L1Emulator) Reorg(height uint64) {
	b.Lock()
	defer b.Unlock()
	for i := uint64(len(b.l1Blocks)) - 1; i > height; i-- {
		delete(b.l1BlocksByHash, b.l1Blocks[i].blockId.HeaderHash)
	}
	b.l1Blocks = b.l1Blocks[:height+1]

	latestHeight := b.Latest().blockId.Height.AsInt().Uint64()
	if latestHeight != height {
		panic("wrong height")
	}
}

func (b *L1Emulator) BlockIdForHeight(_ context.Context, height *common.TimeBlocks) (*common.BlockId, error) {
	b.Lock()
	defer b.Unlock()
	h := height.AsInt().Uint64()
	if h >= uint64(len(b.l1Blocks)) {
		return nil, ethereum.NotFound
	}
	return b.l1Blocks[h].blockId, nil
}

func (b *L1Emulator) TimestampForBlockHash(_ context.Context, hash common.Hash) (*big.Int, error) {
	b.Lock()
	defer b.Unlock()
	info, ok := b.l1BlocksByHash[hash]
	if !ok {
		return nil, ethereum.NotFound
	}
	return info.timestamp, nil
}

func (b *L1Emulator) addBlock() L1BlockInfo {
	info := L1BlockInfo{
		blockId: &common.BlockId{
			Height:     common.NewTimeBlocksInt(int64(len(b.l1Blocks))),
			HeaderHash: common.RandHash(),
		},
		timestamp: big.NewInt(time.Now().Unix() + b.timeIncrease),
	}
	b.l1Blocks = append(b.l1Blocks, info)
	b.l1BlocksByHash[info.blockId.HeaderHash] = info
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
