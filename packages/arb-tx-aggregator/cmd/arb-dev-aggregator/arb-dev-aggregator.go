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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/machineobserver"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
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

	//go http.ListenAndServe("localhost:6060", nil)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error parsing arguments")
	}

	if fs.NArg() != 2 {
		logger.Fatal().Msgf(
			"usage: arb-dev-aggregator [--maxBatchTime=NumSeconds] %s %s",
			utils.WalletArgsString,
			utils.RollupArgsString,
		)
	}

	ethURL := fs.Arg(0)
	privKeyHash := fs.Arg(1)

	if *enablePProf {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	ethclint, err := ethutils.NewRPCEthClient(ethURL)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error running NewRPcEthClient")
	}

	privKey, err := crypto.HexToECDSA(privKeyHash)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	auth := bind.NewKeyedTransactor(privKey)

	arbClient := ethbridge.NewEthClientAdvanced(ethclint, time.Millisecond*100)

	ethAuthClient, err := ethbridge.NewEthAuthClientAdvanced(ctx, arbClient, auth)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	inboxAddress, err := ethbridge.DeployGlobalInbox(ctx, ethAuthClient)
	rollupAddress := common.NewAddressFromEth(auth.From)

	logger.Info().Hex("from", auth.From.Bytes()).Msg("Aggregator submitting batches")

	curBalance, err := ethclint.BalanceAt(ctx, auth.From, nil)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}
	if curBalance.Cmp(big.NewInt(0)) <= 0 {
		log.Fatal().Msg("insufficient balance")
	}

	cp := checkpointing.NewInMemoryCheckpointer()
	if err := cp.Initialize(arbos.Path()); err != nil {
		logger.Fatal().Err(err).Send()
	}
	as := machine.NewInMemoryAggregatorStore()

	db := txdb.New(ethAuthClient, cp, as, rollupAddress)

	globalEth, err := ethAuthClient.NewGlobalInbox(common.NewAddressFromEth(inboxAddress), rollupAddress)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	startId, err := ethAuthClient.BlockIdForHeight(ctx, nil)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	startTime, err := ethAuthClient.TimestampForBlockHash(ctx, startId.HeaderHash)
	if err != nil {
		logger.Fatal().Err(err).Send()
	}

	config := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(10),
		StakeToken:              common.Address{},
		GracePeriod:             common.TimeTicks{Val: big.NewInt(13000 * 2)},
		MaxExecutionSteps:       10000000000,
		ArbGasSpeedLimitPerTick: 200000,
	}
	initMsg := message.Init{
		ChainParams: config,
		Owner:       common.NewAddressFromEth(auth.From),
		ExtraConfig: nil,
	}
	if err := globalEth.SendInitializationMessage(ctx, initMsg.AsData()); err != nil {
		logger.Fatal().Err(err).Send()
	}

	if err := machineobserver.ExecuteObserverAdvanced(
		ctx,
		ethAuthClient,
		big.NewInt(100000),
		db,
		globalEth,
		arbbridge.ChainInfo{
			BlockId:  startId,
			LogIndex: 0,
		},
		startTime,
	); err != nil {
		logger.Fatal().Err(err).Send()
	}

	batch := &Batcher{globalInbox: globalEth}

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
		startId.Height.AsInt(),
		db,
		rollupAddress,
		"8547",
		"8548",
		rpcVars,
		batch,
	); err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error running LaunchAggregator")
	}
}

type Batcher struct {
	globalInbox arbbridge.GlobalInboxSender
	newTxFeed   event.Feed

	messages []inbox.InboxMessage
}

// Return nil if no pending transaction count is available
func (b *Batcher) PendingTransactionCount(context.Context, common.Address) *uint64 {
	return nil
}

func (b *Batcher) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	arbTx := message.NewCompressedECDSAFromEth(tx)
	arbMsg, err := message.NewL2Message(arbTx)
	if err != nil {
		return err
	}
	delivered, err := b.globalInbox.SendL2Message(ctx, arbMsg.Data)
	if err != nil {
		return err
	}
	b.messages = append(b.messages, delivered.Message)
	return nil
}

func (b *Batcher) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.newTxFeed.Subscribe(ch)
}

// Return nil if no pending snapshot is available
func (b *Batcher) PendingSnapshot() *snapshot.Snapshot {
	return nil
}
