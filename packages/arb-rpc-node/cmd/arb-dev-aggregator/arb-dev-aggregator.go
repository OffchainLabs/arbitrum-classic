/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/pkg/errors"
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

	"github.com/ethereum/go-ethereum"
	accounts2 "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-rpc-node/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
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

	tmpDir, err := ioutil.TempDir(".", "arbitrum")
	if err != nil {
		logger.Fatal().Err(err).Msg("error generating temporary directory")
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}()
	storage, err := cmachine.NewArbStorage(tmpDir)
	if err != nil {
		logger.Fatal().Err(err).Msg("error opening ArbStorage")
	}
	defer storage.CloseArbStorage()

	err = storage.Initialize(arbos.Path())
	if err != nil {
		logger.Fatal().Err(err).Msg("error initializing ArbStorage")
	}

	arbCore := storage.GetArbCore()
	started := arbCore.StartThread()
	if !started {
		logger.Fatal().Msg("failed to start thread")
	}

	config := protocol.ChainParams{
		StakeRequirement:        big.NewInt(10),
		StakeToken:              common.Address{},
		GracePeriod:             common.NewTimeBlocksInt(3),
		MaxExecutionSteps:       10000000000,
		ArbGasSpeedLimitPerTick: 20000000000,
	}
	owner := common.RandAddress()
	rollupAddress := common.RandAddress()
	initMsg := message.Init{
		ChainParams: config,
		Owner:       owner,
		ExtraConfig: nil,
	}

	as := storage.GetAggregatorStore()

	db, err := txdb.New(arbCore, as, rollupAddress)
	if err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	signer := types.NewEIP155Signer(message.ChainAddressToID(rollupAddress))
	l1 := NewL1Emulator()
	backend, err := NewBackend(arbCore, db, l1, signer)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	if err := backend.AddInboxMessage(initMsg, rollupAddress, backend.l1Emulator.GenerateBlock()); err != nil {
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
		deposit := message.EthDepositTx{
			L2Message: message.NewSafeL2Message(message.ContractTransaction{
				BasicTx: message.BasicTx{
					MaxGas:      big.NewInt(1000000),
					GasPriceBid: big.NewInt(0),
					DestAddress: common.NewAddressFromEth(account.Address),
					Payment:     depositSize,
					Data:        nil,
				},
			}),
		}
		if err := backend.AddInboxMessage(deposit, rollupAddress, backend.l1Emulator.GenerateBlock()); err != nil {
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
	as      machine.AggregatorStore
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
	block := s.backend.l1Emulator.GenerateBlock()
	return s.backend.AddInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), common.Address{}, block)
}

func (s *EVM) IncreaseTime(amount int64) (string, error) {
	s.backend.l1Emulator.IncreaseTime(amount)
	block := s.backend.l1Emulator.GenerateBlock()
	err := s.backend.AddInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), common.Address{}, block)
	return strconv.FormatInt(amount, 10), err
}

type Backend struct {
	sync.Mutex
	arbcore    core.ArbCore
	db         *txdb.TxDB
	logReader  *core.LogReader
	l1Emulator *L1Emulator
	signer     types.Signer

	newTxFeed event.Feed

	messages []inbox.InboxMessage
}

func NewBackend(arbcore core.ArbCore, db *txdb.TxDB, l1 *L1Emulator, signer types.Signer) (*Backend, error) {
	logReader := core.NewLogReader(db, arbcore, big.NewInt(0), big.NewInt(10))
	errChan := logReader.Start(context.Background())
	go func() {
		err := <-errChan
		log.Fatal().Err(err).Msg("error reading logs")
	}()
	return &Backend{
		arbcore:    arbcore,
		db:         db,
		logReader:  logReader,
		l1Emulator: l1,
		signer:     signer,
	}, nil
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
	latestHeight := b.l1Emulator.Latest().blockId.Height.AsInt().Uint64()
	if latestHeight != height {
		panic("wrong height")
	}
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
	b.messages = b.messages[:height-1]
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
	if err := b.addInboxMessage(arbMsg, common.NewAddressFromEth(sender), block); err != nil {
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
	if res.ResultCode == evm.RevertCode {
		// If transaction failed, rollback the block
		if err := b.reorg(startHeight); err != nil {
			return err
		}

		// Insert an empty block instead
		block := b.l1Emulator.GenerateBlock()
		if err := b.addInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), common.Address{}, block); err != nil {
			return err
		}

		return web3.HandleCallError(res, true)
	}

	return nil
}

func (b *Backend) AddInboxMessage(msg message.Message, sender common.Address, block L1BlockInfo) error {
	b.Lock()
	defer b.Unlock()
	return b.addInboxMessage(msg, sender, block)
}

func (b *Backend) addInboxMessage(msg message.Message, sender common.Address, block L1BlockInfo) error {
	chainTime := inbox.ChainTime{
		BlockNum:  block.blockId.Height,
		Timestamp: block.timestamp,
	}

	inboxMessage := message.NewInboxMessage(msg, sender, big.NewInt(int64(len(b.messages))), big.NewInt(0), chainTime)

	b.messages = append(b.messages, inboxMessage)
	msgCount, err := b.arbcore.GetMessageCount()
	if err != nil {
		return err
	}
	var prevHash common.Hash
	if msgCount.Cmp(big.NewInt(0)) > 0 {
		prevHash, err = b.arbcore.GetInboxAcc(msgCount.Sub(msgCount, big.NewInt(1)))
		if err != nil {
			return err
		}
	}
	successful, err := core.DeliverMessagesAndWait(b.arbcore, []inbox.InboxMessage{inboxMessage}, prevHash, true)
	if err != nil {
		return err
	}
	if !successful {
		return errors.New("failed to deliver message")
	}
	for {
		if b.arbcore.MachineIdle() {
			break
		}
		<-time.After(time.Millisecond * 1000)
	}
	for {
		txdbLogs, err := b.db.CurrentLogCount()
		if err != nil {
			return err
		}
		coreLogs, err := b.arbcore.GetLogCount()
		if err != nil {
			return err
		}
		if txdbLogs.Cmp(coreLogs) == 0 {
			break
		}
		<-time.After(time.Millisecond * 200)
	}

	return nil
}

func (b *Backend) SubscribeNewTxsEvent(ch chan<- ethcore.NewTxsEvent) event.Subscription {
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
