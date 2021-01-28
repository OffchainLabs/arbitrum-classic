/*
 * Copyright 2020, Offchain Labs, Inc.
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

package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum"
	accounts2 "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"io/ioutil"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

var logger zerolog.Logger
var pprofMux *http.ServeMux

func init() {
	pprofMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()

	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	gethlog.Root().SetHandler(gethlog.LvlFilterHandler(gethlog.LvlInfo, gethlog.StreamHandler(os.Stderr, gethlog.TerminalFormat(true))))

	// Print line number that log was created on
	logger = log.With().Caller().Str("component", "arb-dev-aggregator").Logger()
}

func main() {
	ctx := context.Background()
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	rpcVars := utils2.AddRPCFlags(fs)

	enablePProf := fs.Bool("pprof", false, "enable profiling server")
	saveMessages := fs.String("save", "", "save messages")
	walletcount := fs.Int("walletcount", 10, "number of wallets to fund")
	walletbalance := fs.Int64("walletbalance", 100, "amount of funds in each wallet (Eth)")
	mnemonic := fs.String(
		"mnemonic",
		"jar deny prosper gasp flush glass core corn alarm treat leg smart",
		"mnemonic to generate accounts from",
	)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error parsing arguments")
	}

	if *enablePProf {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	cp := checkpointing.NewInMemoryCheckpointer()
	if err := cp.Initialize(arbos.Path()); err != nil {
		logger.Fatal().Err(err).Send()
	}
	as := machine.NewInMemoryAggregatorStore()

	config := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(10),
		StakeToken:              common.Address{},
		GracePeriod:             common.TimeTicks{Val: big.NewInt(13000 * 2)},
		MaxExecutionSteps:       10000000000,
		ArbGasSpeedLimitPerTick: 200000,
	}
	owner := common.RandAddress()
	rollupAddress := common.RandAddress()
	initMsg := message.Init{
		ChainParams: config,
		Owner:       owner,
		ExtraConfig: nil,
	}

	l1 := NewL1Emulator()
	initialBlock := l1.GenerateBlock()
	eventCreated := arbbridge.ChainInfo{
		BlockId:  initialBlock.blockId,
		LogIndex: 0,
	}
	db, err := txdb.New(l1, cp, as, rollupAddress, eventCreated, initialBlock.timestamp)
	if err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	if _, err := db.Load(ctx); err != nil {
		logger.Fatal().Err(err).Send()
	}

	signer := types.NewEIP155Signer(message.ChainAddressToID(rollupAddress))
	backend := NewBackend(db, l1, signer)

	if _, err := backend.AddInboxMessage(ctx, initMsg, rollupAddress, backend.l1Emulator.GenerateBlock()); err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	wallet, err := hdwallet.NewFromMnemonic(*mnemonic)
	if err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	depositSize, ok := new(big.Int).SetString("1000000000000000000", 10)
	if !ok {
		logger.Fatal().Stack().Send()
	}
	depositSize = depositSize.Mul(depositSize, big.NewInt(*walletbalance))

	accounts := make([]accounts2.Account, 0)
	for i := 0; i < *walletcount; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%v", i))
		account, err := wallet.Derive(path, false)
		if err != nil {
			logger.Fatal().Stack().Err(err).Send()
		}
		deposit := message.Eth{
			Dest:  common.NewAddressFromEth(account.Address),
			Value: depositSize,
		}
		if _, err := backend.AddInboxMessage(ctx, deposit, rollupAddress, backend.l1Emulator.GenerateBlock()); err != nil {
			logger.Fatal().Stack().Err(err).Send()
		}
		accounts = append(accounts, account)
	}

	fmt.Println("Arbitrum Dev Chain")
	fmt.Println("")
	fmt.Println("Available Accounts")
	fmt.Println("==================")
	for i, account := range accounts {
		fmt.Printf("(%v) %v (100 ETH)\n", i, account.Address.Hex())
	}

	fmt.Println("\nPrivate Keys")
	fmt.Println("==================")
	for i, account := range accounts {
		privKey, err := wallet.PrivateKeyHex(account)
		if err != nil {
			logger.Fatal().Stack().Err(err).Send()
		}
		fmt.Printf("(%v) 0x%v\n", i, privKey)
	}
	fmt.Println("")

	privateKeys := make([]*ecdsa.PrivateKey, 0)
	for _, account := range accounts {
		privKey, err := wallet.PrivateKey(account)
		if err != nil {
			logger.Fatal().Stack().Err(err).Send()
		}
		privateKeys = append(privateKeys, privKey)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		backend.Lock()
		messages := backend.messages
		backend.Unlock()
		data, err := inbox.TestVectorJSON(messages, nil, nil)
		if err != nil {
			logger.Fatal().Err(err).Send()
		}
		if *saveMessages != "" {
			if err := ioutil.WriteFile(*saveMessages, data, 777); err != nil {
				logger.Fatal().Err(err).Send()
			}
		}
		os.Exit(0)
	}()

	plugins := make(map[string]interface{})
	plugins["evm"] = &EVM{backend: backend, as: as}

	if err := rpc.LaunchAggregatorAdvanced(
		big.NewInt(0),
		db,
		rollupAddress,
		"8547",
		"8548",
		rpcVars,
		backend,
		privateKeys,
		true,
		plugins,
	); err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error running LaunchAggregator")
	}
}

type EVM struct {
	backend *Backend
	as      *machine.InMemoryAggregatorStore
}

func (s *EVM) Snapshot() (hexutil.Uint64, error) {
	logCount, err := s.as.LogCount()
	if err != nil {
		return 0, err
	}
	messageCount, err := s.as.MessageCount()
	if err != nil {
		return 0, err
	}
	latestHeight := s.backend.l1Emulator.Latest().blockId.Height.AsInt().Uint64()
	logger.
		Info().
		Uint64("latest", latestHeight).
		Uint64("logcount", logCount).
		Uint64("messagecount", messageCount).
		Msg("created snapshot")
	return hexutil.Uint64(latestHeight), nil
}

func (s *EVM) Revert(ctx context.Context, snapId hexutil.Uint64) error {
	logger.Info().Uint64("snap", uint64(snapId)).Msg("revert")
	err := s.backend.Reorg(ctx, uint64(snapId))
	if err != nil {
		logger.Error().Err(err).Msg("can't revert")
	}
	return err
}

func (s *EVM) Mine(ctx context.Context, timestamp *hexutil.Uint64) error {
	if timestamp != nil {
		s.backend.l1Emulator.SetTime(int64(*timestamp))
	}
	block := s.backend.l1Emulator.GenerateBlock()
	_, err := s.backend.AddInboxMessage(ctx, message.NewSafeL2Message(message.HeartbeatMessage{}), common.Address{}, block)
	return err
}

func (s *EVM) IncreaseTime(ctx context.Context, amount int64) (string, error) {
	s.backend.l1Emulator.IncreaseTime(amount)
	block := s.backend.l1Emulator.GenerateBlock()
	_, err := s.backend.AddInboxMessage(ctx, message.NewSafeL2Message(message.HeartbeatMessage{}), common.Address{}, block)
	return strconv.FormatInt(amount, 10), err
}

type Backend struct {
	sync.Mutex
	db         *txdb.TxDB
	l1Emulator *L1Emulator
	signer     types.Signer

	newTxFeed event.Feed

	messages []inbox.InboxMessage
}

func NewBackend(db *txdb.TxDB, l1 *L1Emulator, signer types.Signer) *Backend {
	return &Backend{
		db:         db,
		l1Emulator: l1,
		signer:     signer,
	}
}

func (b *Backend) Reorg(ctx context.Context, height uint64) error {
	b.Lock()
	defer b.Unlock()
	return b.reorg(ctx, height)
}

func (b *Backend) reorg(ctx context.Context, height uint64) error {
	startHeight := b.db.LatestBlock()
	b.l1Emulator.Reorg(height)
	latestHeight := b.l1Emulator.Latest().blockId.Height.AsInt().Uint64()
	if latestHeight != height {
		panic("wrong height")
	}
	if _, err := b.db.Load(ctx); err != nil {
		return err
	}
	logger.
		Info().
		Uint64("start", startHeight.Height.AsInt().Uint64()).
		Uint64("end", b.db.LatestBlock().Height.AsInt().Uint64()).
		Uint64("height", height).
		Msg("Reorged chain")
	b.messages = b.messages[:height-1]
	return nil
}

// Return nil if no pending transaction count is available
func (b *Backend) PendingTransactionCount(_ context.Context, account common.Address) *uint64 {
	b.Lock()
	defer b.Unlock()
	return nil
}

func (b *Backend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
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
		Msg("sent transaction")
	startHeight := b.l1Emulator.Latest().blockId.Height.AsInt().Uint64()
	block := b.l1Emulator.GenerateBlock()
	results, err := b.addInboxMessage(ctx, arbMsg, common.NewAddressFromEth(sender), block)
	if err != nil {
		return err
	}
	txHash := common.NewHashFromEth(tx.Hash())
	for _, res := range results {
		// Found matching receipt
		if res.IncomingRequest.MessageID == txHash {
			if res.ResultCode == evm.RevertCode {
				// If transaction failed, rollback the block
				if err := b.reorg(ctx, startHeight); err != nil {
					return err
				}

				// Insert an empty block instead
				block := b.l1Emulator.GenerateBlock()
				if _, err := b.addInboxMessage(ctx, message.NewSafeL2Message(message.HeartbeatMessage{}), common.Address{}, block); err != nil {
					return err
				}

				return web3.HandleCallError(res, true)
			} else {
				return nil
			}
		}
	}
	return nil
}

func (b *Backend) AddInboxMessage(ctx context.Context, msg message.Message, sender common.Address, block L1BlockInfo) ([]*evm.TxResult, error) {
	b.Lock()
	defer b.Unlock()
	return b.addInboxMessage(ctx, msg, sender, block)
}

func (b *Backend) addInboxMessage(ctx context.Context, msg message.Message, sender common.Address, block L1BlockInfo) ([]*evm.TxResult, error) {
	chainTime := inbox.ChainTime{
		BlockNum:  block.blockId.Height,
		Timestamp: block.timestamp,
	}

	inboxMessage := message.NewInboxMessage(msg, sender, big.NewInt(int64(len(b.messages))), chainTime)

	b.messages = append(b.messages, inboxMessage)
	results, err := b.db.AddMessages(ctx, []inbox.InboxMessage{inboxMessage}, block.blockId)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (b *Backend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	b.Lock()
	defer b.Unlock()
	return b.newTxFeed.Subscribe(ch)
}

// Return nil if no pending snapshot is available
func (b *Backend) PendingSnapshot() *snapshot.Snapshot {
	b.Lock()
	defer b.Unlock()
	return nil
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
