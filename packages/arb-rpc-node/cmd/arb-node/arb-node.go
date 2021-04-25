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
	golog "log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/staker"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/nodehealth"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-rpc-node/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var logger zerolog.Logger

var pprofMux *http.ServeMux

const largeChannelBuffer = 200

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
	logger = log.With().Caller().Stack().Str("component", "arb-node").Logger()

	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running node")
	}
}

func startup() error {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	walletArgs := cmdhelp.AddWalletFlags(fs)
	rpcVars := utils2.AddRPCFlags(fs)
	keepPendingState := fs.Bool("pending", false, "enable pending state tracking")
	waitToCatchUp := fs.Bool("wait-to-catch-up", false, "wait to catch up to the chain before opening the RPC")

	//Healthcheck Config
	disablePrimaryCheck := fs.Bool("disable-primary-check", false, "disable checking the health of the primary")
	disableOpenEthereumCheck := fs.Bool("disable-openethereum-check", false, "disable checking the health of the OpenEthereum node")
	healthcheckMetrics := fs.Bool("metrics", false, "enable prometheus endpoint")
	healthcheckRPC := fs.String("healthcheck-rpc", "", "address to bind the healthcheck RPC to")

	maxBatchTime := fs.Int64(
		"maxBatchTime",
		10,
		"maxBatchTime=NumSeconds",
	)
	inboxAddressStr := fs.String("inbox", "", "address of the inbox contract")
	forwardTxURL := fs.String("forward-url", "", "url of another node to send transactions through")

	enablePProf := fs.Bool("pprof", false, "enable profiling server")
	gethLogLevel, arbLogLevel := cmdhelp.AddLogFlags(fs)

	//go http.ListenAndServe("localhost:6060", nil)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "error parsing arguments")
	}

	if fs.NArg() != 3 {
		fmt.Printf("usage: arb-node [--maxBatchTime=NumSeconds] %s %s", cmdhelp.WalletArgsString, utils.RollupArgsString)
		return errors.New("invalid arguments")
	}

	if err := cmdhelp.ParseLogFlags(gethLogLevel, arbLogLevel); err != nil {
		return err
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
		return errors.Wrap(err, "error running NewRPcEthClient")
	}

	l1ChainId, err := ethclint.ChainID(context.Background())
	if err != nil {
		return errors.Wrap(err, "error getting chain ID")
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
			return errors.Wrap(err, "error running GetKeystore")
		}

		if *inboxAddressStr == "" {
			return errors.New("must submit inbox address via --inbox if not running in forwarder mode")
		}
		inboxAddress := common.HexToAddress(*inboxAddressStr)

		logger.Info().Hex("from", auth.From.Bytes()).Msg("Arbitrum node submitting batches")

		if err := ethbridge.WaitForBalance(
			context.Background(),
			ethclint,
			common.Address{},
			common.NewAddressFromEth(auth.From),
		); err != nil {
			return errors.Wrap(err, "error waiting for balance")
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
		return errors.Wrap(err, "error opening monitor")
	}
	defer monitor.Close()

	db, err := txdb.New(context.Background(), monitor.Core, monitor.Storage.GetNodeStore(), rollupArgs.Address, 100*time.Millisecond)
	if err != nil {
		return errors.Wrap(err, "error opening txdb")
	}
	defer db.Close()

	ctx := context.Background()
	healthChan := make(chan nodehealth.Log, largeChannelBuffer)
	go func() {
		err := nodehealth.StartNodeHealthCheck(ctx, healthChan)
		if err != nil {
			log.Error().Err(err).Msg("healthcheck server failed")
		}
	}()

	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckMetrics", ValBool: *healthcheckMetrics}
	healthChan <- nodehealth.Log{Config: true, Var: "disablePrimaryCheck", ValBool: *disablePrimaryCheck}
	healthChan <- nodehealth.Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: *disableOpenEthereumCheck}
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckRPC", ValStr: *healthcheckRPC}

	if *forwardTxURL != "" {
		healthChan <- nodehealth.Log{Config: true, Var: "primaryHealthcheckRPC", ValStr: *forwardTxURL}
	}
	healthChan <- nodehealth.Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: rollupArgs.EthURL}
	nodehealth.Init(healthChan)

	var inboxReader *staker.InboxReader
	for {
		inboxReader, err = monitor.StartInboxReader(context.Background(), rollupArgs.EthURL, rollupArgs.Address, healthChan)
		if err == nil {
			break
		}
		logger.Warn().Err(err).
			Str("url", rollupArgs.EthURL).
			Str("rollup", rollupArgs.Address.Hex()).
			Msg("failed to start inbox reader, waiting and retrying")
		time.Sleep(time.Second * 5)
	}

	if *waitToCatchUp {
		inboxReader.WaitToCatchUp()
	}

	return rpc.LaunchNode(
		ctx,
		ethclint,
		rollupArgs.Address,
		db,
		"8547",
		"8548",
		rpcVars,
		time.Duration(*maxBatchTime)*time.Second,
		batcherMode,
	)
}
