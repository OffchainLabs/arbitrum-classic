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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	golog "log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	utils2 "github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/utils"
	_ "net/http/pprof"
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

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Print line number that log was created on
	logger = log.With().Caller().Str("component", "arb-tx-aggregator").Logger()

	ctx := context.Background()
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	walletArgs := utils.AddWalletFlags(fs)
	rpcVars := utils2.AddRPCFlags(fs)
	keepPendingState := fs.Bool("pending", false, "enable pending state tracking")

	maxBatchTime := fs.Int64(
		"maxBatchTime",
		10,
		"maxBatchTime=NumSeconds",
	)

	forwardTxURL := fs.String("forward-url", "", "url of another aggregator to send transactions through")

	enablePProf := fs.Bool("pprof", false, "enable profiling server")

	//go http.ListenAndServe("localhost:6060", nil)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error parsing arguments")
	}

	if fs.NArg() != 3 {
		logger.Fatal().Msgf(
			"usage: arb-tx-aggregator [--maxBatchTime=NumSeconds] %s %s",
			utils.WalletArgsString,
			utils.RollupArgsString,
		)
	}

	if *enablePProf {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	rollupArgs := utils.ParseRollupCommand(fs, 0)

	ethclint, err := ethutils.NewRPCEthClient(rollupArgs.EthURL)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error running NewRPcEthClient")
	}

	logger.Info().Hex("chainaddress", rollupArgs.Address.Bytes()).Hex("chainid", message.ChainAddressToID(rollupArgs.Address).Bytes()).Msg("Launching aggregator")

	var batcherMode rpc.BatcherMode
	if *forwardTxURL != "" {
		logger.Info().Str("forwardTxURL", *forwardTxURL).Msg("Aggregator starting in forwarder mode")
		batcherMode = rpc.ForwarderBatcherMode{NodeURL: *forwardTxURL}
	} else {
		auth, err := utils.GetKeystore(rollupArgs.ValidatorFolder, walletArgs, fs)
		if err != nil {
			logger.Fatal().Stack().Err(err).Msg("Error running GetKeystore")
		}

		logger.Info().Hex("from", auth.From.Bytes()).Msg("Aggregator submitting batches")

		if err := arbbridge.WaitForBalance(
			ctx,
			ethbridge.NewEthClient(ethclint),
			common.Address{},
			common.NewAddressFromEth(auth.From),
		); err != nil {
			logger.Fatal().Stack().Err(err).Msg("error waiting for balance")
		}

		if *keepPendingState {
			batcherMode = rpc.StatefulBatcherMode{Auth: auth}
		} else {
			batcherMode = rpc.StatelessBatcherMode{Auth: auth}
		}
	}

	contractFile := filepath.Join(rollupArgs.ValidatorFolder, "contract.mexe")
	dbPath := filepath.Join(rollupArgs.ValidatorFolder, "checkpoint_db")

	if err := rpc.LaunchAggregator(
		ctx,
		ethclint,
		rollupArgs.Address,
		contractFile,
		dbPath,
		"8547",
		"8548",
		rpcVars,
		time.Duration(*maxBatchTime)*time.Second,
		batcherMode,
	); err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error running LaunchAggregator")
	}
}
