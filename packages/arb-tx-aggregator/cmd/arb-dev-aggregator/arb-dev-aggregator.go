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
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

	// Print line number that log was created on
	logger = log.With().Caller().Str("component", "arb-dev-aggregator").Logger()

	ctx := context.Background()
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	rpcVars := utils2.AddRPCFlags(fs)

	enablePProf := fs.Bool("pprof", false, "enable profiling server")

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		//messages, err := globalEth.GetDeliveredEvents(ctx, nil, nil)
		//if err != nil {
		//	logger.Fatal().Err(err).Send()
		//}
		//inboxMessages := make([]inbox.InboxMessage, 0)
		//for _, msg := range messages {
		//	inboxMessages = append(inboxMessages, msg.Message)
		//}
		//data, err := inbox.TestVectorJSON(inboxMessages, nil, nil)
		//if err != nil {
		//	logger.Fatal().Err(err).Send()
		//}
		//
		//log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		//log.Info().Msg(string(data))
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
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(b.msgCount),
		Timestamp: big.NewInt(time.Now().Unix()),
	}

	inboxMessage := message.NewInboxMessage(msg, sender, big.NewInt(b.msgCount), chainTime)

	block := b.l1Emulator.generateBlock()

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
