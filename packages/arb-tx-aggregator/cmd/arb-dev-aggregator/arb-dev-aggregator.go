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
	"flag"
	"fmt"
	accounts2 "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
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
	"time"
)

var logger zerolog.Logger
var pprofMux *http.ServeMux

func init() {
	pprofMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()
}

func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// Print line number that log was created on
	logger = log.With().Caller().Str("component", "arb-dev-aggregator").Logger()

	ctx := context.Background()
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	rpcVars := utils2.AddRPCFlags(fs)

	enablePProf := fs.Bool("pprof", false, "enable profiling server")
	saveMessages := fs.String("save", "", "save messages")
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

	db := txdb.New(l1, cp, as, rollupAddress)

	if err := db.Load(ctx); err != nil {
		logger.Fatal().Err(err).Send()
	}

	if err := db.AddInitialBlock(ctx, big.NewInt(0)); err != nil {
		logger.Fatal().Err(err).Send()
	}

	signer := types.NewEIP155Signer(message.ChainAddressToID(rollupAddress))
	backend := NewBackend(db, l1, signer)

	if err := backend.addInboxMessage(ctx, initMsg, rollupAddress); err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	wallet, err := hdwallet.NewFromMnemonic(*mnemonic)
	if err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	depositSize, ok := new(big.Int).SetString("100000000000000000000", 10)
	if !ok {
		logger.Fatal().Stack().Send()
	}

	accounts := make([]accounts2.Account, 0)
	for i := 0; i < 10; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%v", i))
		account, err := wallet.Derive(path, false)
		if err != nil {
			logger.Fatal().Stack().Err(err).Send()
		}
		deposit := message.Eth{
			Dest:  common.NewAddressFromEth(account.Address),
			Value: depositSize,
		}
		if err := backend.addInboxMessage(ctx, deposit, rollupAddress); err != nil {
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		data, err := inbox.TestVectorJSON(backend.messages, nil, nil)
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

	if err := rpc.LaunchAggregatorAdvanced(
		big.NewInt(0),
		db,
		rollupAddress,
		"8547",
		"8548",
		rpcVars,
		backend,
	); err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error running LaunchAggregator")
	}
}

type l1BlockInfo struct {
	blockId   *common.BlockId
	timestamp *big.Int
}

type Backend struct {
	db         *txdb.TxDB
	l1Emulator *L1Emulator
	signer     types.Signer

	newTxFeed event.Feed

	msgCount int64
	messages []inbox.InboxMessage
}

func NewBackend(db *txdb.TxDB, l1 *L1Emulator, signer types.Signer) *Backend {
	return &Backend{
		db:         db,
		l1Emulator: l1,
		signer:     signer,
	}
}

// Return nil if no pending transaction count is available
func (b *Backend) PendingTransactionCount(context.Context, common.Address) *uint64 {
	return nil
}

func (b *Backend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	arbTx := message.NewCompressedECDSAFromEth(tx)
	sender, err := types.Sender(b.signer, tx)
	if err != nil {
		return err
	}
	arbMsg, err := message.NewL2Message(arbTx)
	if err != nil {
		return err
	}

	return b.addInboxMessage(ctx, arbMsg, common.NewAddressFromEth(sender))
}

func (b *Backend) addInboxMessage(ctx context.Context, msg message.Message, sender common.Address) error {
	block := b.l1Emulator.generateBlock()

	chainTime := inbox.ChainTime{
		BlockNum:  block.blockId.Height,
		Timestamp: block.timestamp,
	}

	inboxMessage := message.NewInboxMessage(msg, sender, big.NewInt(b.msgCount), chainTime)

	if err := b.db.AddMessages(ctx, []inbox.InboxMessage{inboxMessage}, block.blockId); err != nil {
		return err
	}

	b.messages = append(b.messages, inboxMessage)
	b.msgCount++
	return nil
}

func (b *Backend) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.newTxFeed.Subscribe(ch)
}

// Return nil if no pending snapshot is available
func (b *Backend) PendingSnapshot() *snapshot.Snapshot {
	return nil
}

type L1Emulator struct {
	l1Blocks       map[uint64]l1BlockInfo
	l1BlocksByHash map[common.Hash]l1BlockInfo
	latest         uint64
}

func NewL1Emulator() *L1Emulator {
	genesis := l1BlockInfo{
		blockId: &common.BlockId{
			Height:     common.NewTimeBlocksInt(0),
			HeaderHash: common.RandHash(),
		},
		timestamp: big.NewInt(time.Now().Unix()),
	}

	b := &L1Emulator{
		l1Blocks:       make(map[uint64]l1BlockInfo),
		l1BlocksByHash: make(map[common.Hash]l1BlockInfo),
	}
	b.addBlock(genesis)
	return b
}

func (b *L1Emulator) BlockIdForHeight(_ context.Context, height *common.TimeBlocks) (*common.BlockId, error) {
	return b.l1Blocks[height.AsInt().Uint64()].blockId, nil
}

func (b *L1Emulator) TimestampForBlockHash(_ context.Context, hash common.Hash) (*big.Int, error) {
	info, ok := b.l1BlocksByHash[hash]
	if !ok {
		return nil, errors.Errorf("no info for block with hash %v", hash)
	}
	return info.timestamp, nil
}

func (b *L1Emulator) addBlock(info l1BlockInfo) {
	b.l1Blocks[info.blockId.Height.AsInt().Uint64()] = info
	b.l1BlocksByHash[info.blockId.HeaderHash] = info
}

func (b *L1Emulator) generateBlock() l1BlockInfo {
	info := l1BlockInfo{
		blockId: &common.BlockId{
			Height:     common.NewTimeBlocksInt(int64(b.latest)),
			HeaderHash: common.RandHash(),
		},
		timestamp: big.NewInt(time.Now().Unix()),
	}
	b.addBlock(info)
	b.latest++
	return info
}
