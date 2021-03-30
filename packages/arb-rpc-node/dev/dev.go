package dev

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	core2 "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/staker"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

var logger = log.With().Caller().Stack().Str("component", "dev").Logger()

func NewDevNode(dir string, config protocol.ChainParams) (*staker.Monitor, *Backend, *txdb.TxDB, common.Address) {
	owner := common.RandAddress()
	rollupAddress := common.RandAddress()
	initMsg := message.Init{
		ChainParams: config,
		Owner:       owner,
		ExtraConfig: nil,
	}

	monitor, err := staker.NewMonitor(dir, arbos.Path())
	if err != nil {
		logger.Fatal().Err(err).Msg("error opening monitor")
	}

	db, err := txdb.New(context.Background(), monitor.Core, monitor.Storage.GetNodeStore(), rollupAddress, 10*time.Millisecond)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	signer := types.NewEIP155Signer(message.ChainAddressToID(rollupAddress))
	l1 := NewL1Emulator()
	backend, err := NewBackend(monitor.Core, db, l1, signer)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	if _, err := backend.AddInboxMessage(initMsg, rollupAddress); err != nil {
		logger.Fatal().Err(err).Send()
	}

	return monitor, backend, db, rollupAddress
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

type Backend struct {
	sync.Mutex
	arbcore    core.ArbCore
	db         *txdb.TxDB
	l1Emulator *L1Emulator
	signer     types.Signer

	newTxFeed event.Feed
}

func NewBackend(arbcore core.ArbCore, db *txdb.TxDB, l1 *L1Emulator, signer types.Signer) (*Backend, error) {
	return &Backend{
		arbcore:    arbcore,
		db:         db,
		l1Emulator: l1,
		signer:     signer,
	}, nil
}

func (b *Backend) ExportData() ([]byte, error) {
	b.Lock()
	messageCount, err := b.arbcore.GetMessageCount()
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	messages, err := b.arbcore.GetMessages(big.NewInt(0), messageCount)
	if err != nil {
		logger.Fatal().Err(err).Send()
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
	arbMsg, err := message.NewL2Message(arbTx)
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
	if _, err := b.addInboxMessage(arbMsg, common.NewAddressFromEth(sender), block); err != nil {
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
		if _, err := b.addInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), common.Address{}, block); err != nil {
			return err
		}

		return web3.HandleCallError(res, true)
	}

	return nil
}

func (b *Backend) AddInboxMessage(msg message.Message, sender common.Address) (common.Hash, error) {
	b.Lock()
	defer b.Unlock()
	return b.addInboxMessage(msg, sender, b.l1Emulator.GenerateBlock())
}

func (b *Backend) addInboxMessage(msg message.Message, sender common.Address, block L1BlockInfo) (common.Hash, error) {
	chainTime := inbox.ChainTime{
		BlockNum:  block.blockId.Height,
		Timestamp: block.timestamp,
	}
	msgCount, err := b.arbcore.GetMessageCount()
	if err != nil {
		return common.Hash{}, err
	}
	inboxMessage := message.NewInboxMessage(msg, sender, new(big.Int).Set(msgCount), big.NewInt(0), chainTime)

	requestId := message.CalculateRequestId(b.signer.ChainID(), msgCount)
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
