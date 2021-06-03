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
	"flag"
	"fmt"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/pkg/errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/metrics"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/nodehealth"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcastclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
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
	defer logger.Log().Msg("Cleanly shutting down node")
	ctx, cancelFunc, cancelChan := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	fs := flag.NewFlagSet("", flag.ContinueOnError)
	walletArgs := cmdhelp.AddWalletFlags(fs)
	keepPendingState := fs.Bool("pending", false, "enable pending state tracking")
	sequencerMode := fs.Bool("sequencer", false, "act as sequencer")
	waitToCatchUp := fs.Bool("wait-to-catch-up", false, "wait to catch up to the chain before opening the RPC")
	delayedMessagesTargetDelay := fs.Int64("delayed-messages-target-delay", 12, "delay before sequencing delayed messages")
	createBatchBlockInterval := fs.Int64("create-batch-block-interval", 1, "block interval at which to create new batches")
	gasPriceUrl := fs.String("gas-price-url", "", "gas price rpc url (etherscan compatible)")

	chainId64 := fs.Uint64("chainid", 42161, "chain id of the arbitrum chain")

	//Healthcheck Config
	disablePrimaryCheck := fs.Bool("disable-primary-check", false, "disable checking the health of the primary")
	disableOpenEthereumCheck := fs.Bool("disable-openethereum-check", false, "disable checking the health of the OpenEthereum node")
	healthcheckMetrics := fs.Bool("metrics", false, "enable prometheus endpoint")
	healthcheckRPC := fs.String("healthcheck-rpc", "", "address to bind the healthcheck RPC to")
	metricsPrefix := fs.String("metrics-prefix", "", "prepend the specified prefix to the exported metrics names")

	maxBatchTime := fs.Int64(
		"maxBatchTime",
		10,
		"maxBatchTime=NumSeconds",
	)
	inboxAddressStr := fs.String("inbox", "", "address of the inbox contract")
	forwardTxURL := fs.String("forward-url", "", "url of another node to send transactions through")
	feedURL := fs.String("feed-url", "", "URL of sequencer feed source")
	rpcAddr := fs.String("rpc.addr", "0.0.0.0", "RPC address")
	rpcPort := fs.String("rpc.port", "8547", "RPC port")
	wsAddr := fs.String("ws.addr", "0.0.0.0", "websocket address")
	wsPort := fs.String("ws.port", "8548", "websocket port")
	feedOutputAddr := fs.String("feedoutput.addr", "0.0.0.0", "address to bind the relay feed output to")
	feedOutputPort := fs.String("feedoutput.port", "9642", "port to bind the relay feed output to")
	feedOutputPingInterval := fs.Duration("feedoutput.ping", 5*time.Second, "number of seconds for ping interval")
	feedOutputTimeout := fs.Duration("feedoutput.timeout", 15*time.Second, "number of seconds for timeout")
	enablePProf := fs.Bool("pprof", false, "enable profiling server")
	gethLogLevel, arbLogLevel := cmdhelp.AddLogFlags(fs)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "error parsing arguments")
	}

	if fs.NArg() != 4 || (*sequencerMode && *feedURL != "") {
		fmt.Printf("\n")
		fmt.Printf("usage       sequencer: arb-node --sequencer [optional arguments] %s %s\n", cmdhelp.WalletArgsString, utils.RollupArgsString)
		fmt.Printf("   or aggregator node: arb-node --feed-url=<feed address> --inbox=<inbox address> [optional arguments] %s %s\n", cmdhelp.WalletArgsString, utils.RollupArgsString)
		fmt.Printf("   or            node: arb-node --feed-url=<feed address> --forward-url=<sequencer RPC> [optional arguments] %s %s\n\n", cmdhelp.WalletArgsString, utils.RollupArgsString)
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

	l1ChainId, err := ethclint.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "error getting chain ID")
	}
	logger.Debug().Str("chainid", l1ChainId.String()).Msg("connected to l1 chain")

	l2ChainId := new(big.Int).SetUint64(*chainId64)
	logger.Info().Hex("chainaddress", rollupArgs.Address.Bytes()).Hex("chainid", l2ChainId.Bytes()).Msg("Launching arbitrum node")

	contractFile := filepath.Join(rollupArgs.ValidatorFolder, "arbos.mexe")
	dbPath := filepath.Join(rollupArgs.ValidatorFolder, "checkpoint_db")

	mon, err := monitor.NewMonitor(dbPath, contractFile)
	if err != nil {
		return errors.Wrap(err, "error opening monitor")
	}
	defer mon.Close()

	metricsConfig := metrics.NewMetricsConfig(metricsPrefix)
	healthChan := make(chan nodehealth.Log, largeChannelBuffer)
	go func() {
		err := nodehealth.StartNodeHealthCheck(ctx, healthChan, metricsConfig.Registry, metricsConfig.Registerer)
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

	var sequencerFeed chan broadcaster.BroadcastFeedMessage
	if !*sequencerMode {
		if *feedURL == "" {
			logger.Warn().Msg("Missing --feed-url so not subscribing to feed")
		} else {
			broadcastClient := broadcastclient.NewBroadcastClient(*feedURL, nil, 20*time.Second)
			for {
				sequencerFeed, err = broadcastClient.Connect(ctx)
				if err == nil {
					break
				}
				logger.Warn().Err(err).
					Msg("failed connect to sequencer broadcast, waiting and retrying")

				select {
				case <-ctx.Done():
					return errors.New("ctx cancelled broadcast client connect")
				case <-time.After(5 * time.Second):
				}
			}
		}
	}
	var inboxReader *monitor.InboxReader
	for {
		inboxReader, err = mon.StartInboxReader(ctx, ethclint, rollupArgs.Address, rollupArgs.BridgeUtilsAddress, healthChan, sequencerFeed)
		if err == nil {
			break
		}
		logger.Warn().Err(err).
			Str("url", rollupArgs.EthURL).
			Str("rollup", rollupArgs.Address.Hex()).
			Str("bridgeUtils", rollupArgs.BridgeUtilsAddress.Hex()).
			Msg("failed to start inbox reader, waiting and retrying")

		select {
		case <-ctx.Done():
			return errors.New("ctx cancelled StartInboxReader retry loop")
		case <-time.After(5 * time.Second):
		}
	}

	var broadcasterSettings broadcaster.Settings
	var dataSigner func([]byte) ([]byte, error)
	var batcherMode rpc.BatcherMode
	if *forwardTxURL != "" {
		logger.Info().Str("forwardTxURL", *forwardTxURL).Msg("Arbitrum node starting in forwarder mode")
		batcherMode = rpc.ForwarderBatcherMode{NodeURL: *forwardTxURL}
	} else {
		var auth *bind.TransactOpts
		auth, dataSigner, err = cmdhelp.GetKeystore(rollupArgs.ValidatorFolder, walletArgs, fs, l1ChainId)
		if err != nil {
			return errors.Wrap(err, "error running GetKeystore")
		}

		var inboxAddress common.Address
		if !*sequencerMode {
			if *inboxAddressStr == "" {
				return errors.New("must submit inbox address via --inbox if not running in forwarder or sequencer mode")
			}
			inboxAddress = common.HexToAddress(*inboxAddressStr)
		}

		logger.Info().Hex("from", auth.From.Bytes()).Msg("Arbitrum node submitting batches")

		if err := ethbridge.WaitForBalance(
			ctx,
			ethclint,
			common.Address{},
			common.NewAddressFromEth(auth.From),
		); err != nil {
			return errors.Wrap(err, "error waiting for balance")
		}

		if *sequencerMode {
			batcherMode = rpc.SequencerBatcherMode{
				Auth:                       auth,
				Core:                       mon.Core,
				InboxReader:                inboxReader,
				DelayedMessagesTargetDelay: big.NewInt(*delayedMessagesTargetDelay),
				CreateBatchBlockInterval:   big.NewInt(*createBatchBlockInterval),
			}

			broadcasterSettings = broadcaster.Settings{
				Addr:                    *feedOutputAddr + ":" + *feedOutputPort,
				Workers:                 128,
				Queue:                   1,
				IoReadWriteTimeout:      2 * time.Second,
				ClientPingInterval:      *feedOutputPingInterval,
				ClientNoResponseTimeout: *feedOutputTimeout,
			}
		} else if *keepPendingState {
			batcherMode = rpc.StatefulBatcherMode{Auth: auth, InboxAddress: inboxAddress}
		} else {
			batcherMode = rpc.StatelessBatcherMode{Auth: auth, InboxAddress: inboxAddress}
		}
	}

	nodeStore := mon.Storage.GetNodeStore()
	metrics.RegisterNodeStoreMetrics(nodeStore, metricsConfig)
	db, txDBErrChan, err := txdb.New(ctx, mon.Core, nodeStore, 100*time.Millisecond)
	if err != nil {
		return errors.Wrap(err, "error opening txdb")
	}
	defer db.Close()

	if *waitToCatchUp {
		inboxReader.WaitToCatchUp(ctx)
	}

	var batch batcher.TransactionBatcher
	for {
		batch, err = rpc.SetupBatcher(
			ctx,
			ethclint,
			rollupArgs.Address,
			l2ChainId,
			db,
			time.Duration(*maxBatchTime)*time.Second,
			batcherMode,
			dataSigner,
			broadcasterSettings,
			*gasPriceUrl,
		)
		if err == nil {
			break
		}
		logger.Warn().Err(err).Msg("failed to setup batcher, waiting and retrying")

		select {
		case <-ctx.Done():
			return errors.New("ctx cancelled setup batcher")
		case <-time.After(5 * time.Second):
		}
	}

	metricsConfig.RegisterSystemMetrics()
	metricsConfig.RegisterStaticMetrics()

	srv := aggregator.NewServer(batch, rollupArgs.Address, l2ChainId, db)
	web3Server, err := web3.GenerateWeb3Server(srv, nil, false, nil, metricsConfig)
	if err != nil {
		return err
	}
	errChan := make(chan error, 1)
	go func() {
		err := rpc.LaunchPublicServer(ctx, web3Server, *rpcAddr, *rpcPort, *wsAddr, *wsPort)
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-txDBErrChan:
		return err
	case err := <-errChan:
		return err
	case <-cancelChan:
		return nil
	}
}
