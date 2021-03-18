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
	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/staker"
	golog "log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-rpc-node/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
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

	gethlog.Root().SetHandler(gethlog.LvlFilterHandler(gethlog.LvlInfo, gethlog.StreamHandler(os.Stderr, gethlog.TerminalFormat(true))))

	// Print line number that log was created on
	logger = log.With().Caller().Str("component", "arb-node").Logger()

	ctx := context.Background()
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	walletArgs := cmdhelp.AddWalletFlags(fs)
	rpcVars := utils2.AddRPCFlags(fs)
	keepPendingState := fs.Bool("pending", false, "enable pending state tracking")

	maxBatchTime := fs.Int64(
		"maxBatchTime",
		10,
		"maxBatchTime=NumSeconds",
	)
	inboxAddressStr := fs.String("inbox", "", "address of the inbox contract")
	forwardTxURL := fs.String("forward-url", "", "url of another node to send transactions through")

	enablePProf := fs.Bool("pprof", false, "enable profiling server")

	//go http.ListenAndServe("localhost:6060", nil)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error parsing arguments")
	}

	if fs.NArg() != 3 {
		logger.Fatal().Msgf(
			"usage: arb-node [--maxBatchTime=NumSeconds] %s %s",
			cmdhelp.WalletArgsString,
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

	l1ChainId, err := ethclint.ChainID(context.Background())
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error getting chain ID")
	}
	logger.Debug().Str("chainid", l1ChainId.String()).Msg("connected to l1 chain")

	logger.Info().Hex("chainaddress", rollupArgs.Address.Bytes()).Hex("chainid", message.ChainAddressToID(rollupArgs.Address).Bytes()).Msg("Launching arbitrum node")

	var batcherMode rpc.BatcherMode
	if *forwardTxURL != "" {
		logger.Info().Str("forwardTxURL", *forwardTxURL).Msg("Arbitrum node starting in forwarder mode")
		batcherMode = rpc.ForwarderBatcherMode{NodeURL: *forwardTxURL}
	} else {
		auth, err := cmdhelp.GetKeystore(rollupArgs.ValidatorFolder, walletArgs, fs, l1ChainId)
		if err != nil {
			logger.Fatal().Stack().Err(err).Msg("Error running GetKeystore")
		}

		if *inboxAddressStr == "" {
			logger.Fatal().Msg("must submit inbox addres via --inbox if not running in forwarder mode")
		}
		inboxAddress := common.HexToAddress(*inboxAddressStr)

		logger.Info().Hex("from", auth.From.Bytes()).Msg("Arbitrum node submitting batches")

		if err := ethbridge.WaitForBalance(
			ctx,
			ethclint,
			common.Address{},
			common.NewAddressFromEth(auth.From),
		); err != nil {
			logger.Fatal().Stack().Err(err).Msg("error waiting for balance")
		}

		if *keepPendingState {
			batcherMode = rpc.StatefulBatcherMode{Auth: auth, InboxAddress: inboxAddress}
		} else {
			batcherMode = rpc.StatelessBatcherMode{Auth: auth, InboxAddress: inboxAddress}
		}
	}

	contractFile := filepath.Join(rollupArgs.ValidatorFolder, "arbos.mexe")
	dbPath := filepath.Join(rollupArgs.ValidatorFolder, "checkpoint_db")

	monitor, err := staker.NewMonitor(dbPath, contractFile)
	if err != nil {
		logger.Fatal().Err(err).Msg("error opening monitor")
	}
	defer monitor.Close()

	db, err := txdb.New(context.Background(), monitor.Core, monitor.Storage.GetNodeStore(), rollupArgs.Address, 100*time.Millisecond)
	if err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	_, err = monitor.StartInboxReader(context.Background(), rollupArgs.EthURL, rollupArgs.Address)
	if err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	if err := rpc.LaunchNode(
		ctx,
		ethclint,
		rollupArgs.Address,
		db,
		"8547",
		"8548",
		rpcVars,
		time.Duration(*maxBatchTime)*time.Second,
		batcherMode,
	); err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error running LaunchNode")
	}
}
