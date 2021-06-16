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
	"fmt"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
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
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcastclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
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

	config, wallet, err := configuration.Parse()
	if err != nil || len(config.Database.Path) == 0 || len(config.L1.URL) == 0 ||
		len(config.Rollup.Address) == 0 || len(config.Bridge.Utils.Address) == 0 ||
		(config.Sequencer && config.Feed.URL != "") {
		fmt.Printf("\n")
		fmt.Printf("usage                  arb-node --conf=<filename> \n")
		fmt.Printf("   or       sequencer: arb-node --sequencer --database.path=<path> --l1.url=<url> --rollup.address=<address> --bridgeutils.address=<address> [optional arguments] %s\n", cmdhelp.WalletArgsString)
		fmt.Printf("   or aggregator node: arb-node --feed.url=<feed websocket> --inbox=<inbox address> --database.path=<path> --l1.url=<url> --rollup.address=<address> --bridgeutils.address=<address> [optional arguments] %s\n", cmdhelp.WalletArgsString)
		fmt.Printf("   or            node: arb-node --feed.url=<feed websocket> --forward.url=<sequencer RPC> --database.path=<path> --l1.url=<url> --rollup.address=<address> --bridgeutils.address=<address> [optional arguments] %s\n\n", cmdhelp.WalletArgsString)
		fmt.Printf("   or            node: arb-node --l1.url=<url> --database.path=<path> --mainnet.arb1 \n")
		fmt.Printf("   or            node: arb-node --l1.url=<url> --database.path=<path> --testnet.rinkeby \n")
		return errors.New("invalid arguments")
	}

	if err := cmdhelp.ParseLogFlags(&config.RPC.LogLevel, &config.LogLevel); err != nil {
		return err
	}

	if config.Pprof.Enabled {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	ethclint, err := ethutils.NewRPCEthClient(config.L1.URL)
	if err != nil {
		return errors.Wrap(err, "error running NewRPcEthClient")
	}

	l1ChainId, err := ethclint.ChainID(ctx)
	if err != nil {
		return errors.Wrap(err, "error getting chain ID")
	}
	logger.Debug().Str("chainid", l1ChainId.String()).Msg("connected to l1 chain")

	l2ChainId := new(big.Int).SetUint64(config.ChainID)
	rollupAddress := common.HexToAddress(config.Rollup.Address)
	logger.Info().Hex("chainaddress", rollupAddress.Bytes()).Hex("chainid", l2ChainId.Bytes()).Msg("Launching arbitrum node")

	contractFile := filepath.Join(config.Database.Path, "arbos.mexe")
	dbPath := filepath.Join(config.Database.Path, "checkpoint_db")

	mon, err := monitor.NewMonitor(dbPath, contractFile)
	if err != nil {
		return errors.Wrap(err, "error opening monitor")
	}
	defer mon.Close()

	metricsConfig := metrics.NewMetricsConfig(&config.Healthcheck.Metrics.Prefix)
	healthChan := make(chan nodehealth.Log, largeChannelBuffer)
	go func() {
		err := nodehealth.StartNodeHealthCheck(ctx, healthChan, metricsConfig.Registry, metricsConfig.Registerer)
		if err != nil {
			log.Error().Err(err).Msg("healthcheck server failed")
		}
	}()

	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckMetrics", ValBool: config.Healthcheck.Metrics.Enabled}
	healthChan <- nodehealth.Log{Config: true, Var: "disablePrimaryCheck", ValBool: !config.Healthcheck.Sequencer.Enabled}
	healthChan <- nodehealth.Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: !config.Healthcheck.L1Node.Enabled}
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckRPC", ValStr: config.Healthcheck.Addr + ":" + config.Healthcheck.Port}

	if config.Forward.URL != "" {
		healthChan <- nodehealth.Log{Config: true, Var: "primaryHealthcheckRPC", ValStr: config.Forward.URL}
	}
	healthChan <- nodehealth.Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: config.L1.URL}
	nodehealth.Init(healthChan)

	var sequencerFeed chan broadcaster.BroadcastFeedMessage
	if !config.Sequencer {
		if config.Feed.URL == "" {
			logger.Warn().Msg("Missing --feed.url so not subscribing to feed")
		} else {
			broadcastClient := broadcastclient.NewBroadcastClient(config.Feed.URL, nil, 20*time.Second)
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
		inboxReader, err = mon.StartInboxReader(ctx, ethclint, common.HexToAddress(config.Rollup.Address), common.HexToAddress(config.Bridge.Utils.Address), healthChan, sequencerFeed)
		if err == nil {
			break
		}
		logger.Warn().Err(err).
			Str("url", config.L1.URL).
			Str("rollup", config.Rollup.Address).
			Str("bridgeUtils", config.Bridge.Utils.Address).
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
	if config.Forward.URL != "" {
		logger.Info().Str("forwardTxURL", config.Forward.URL).Msg("Arbitrum node starting in forwarder mode")
		batcherMode = rpc.ForwarderBatcherMode{NodeURL: config.Forward.URL}
	} else {
		var auth *bind.TransactOpts
		auth, dataSigner, err = cmdhelp.GetKeystore(config.Database.Path, wallet, l1ChainId)
		if err != nil {
			return errors.Wrap(err, "error running GetKeystore")
		}

		var inboxAddress common.Address
		if !config.Sequencer {
			if config.Inbox.Address == "" {
				return errors.New("must submit inbox address via --inbox if not running in forwarder or sequencer mode")
			}
			inboxAddress = common.HexToAddress(config.Inbox.Address)
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

		if config.Sequencer {
			batcherMode = rpc.SequencerBatcherMode{
				Auth:                       auth,
				Core:                       mon.Core,
				InboxReader:                inboxReader,
				DelayedMessagesTargetDelay: big.NewInt(config.DelayedMessagesTargetDelay),
				CreateBatchBlockInterval:   big.NewInt(config.CreateBatchBlockInterval),
			}

			ping, err := time.ParseDuration(config.Feed.Output.Ping)
			if err != nil {
				logger.Fatal().Err(err).Msg("error parsing feedoutput ping duration")
			}
			timeout, err := time.ParseDuration(config.Feed.Output.Timeout)
			if err != nil {
				logger.Fatal().Err(err).Msg("error parsing feedoutput timeout")
			}
			broadcasterSettings = broadcaster.Settings{
				Addr:                    config.Feed.Output.Addr + ":" + config.Feed.Output.Port,
				Workers:                 128,
				Queue:                   1,
				IoReadWriteTimeout:      2 * time.Second,
				ClientPingInterval:      ping,
				ClientNoResponseTimeout: timeout,
			}
		} else if config.Pending {
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

	if config.WaitToCatchUp {
		inboxReader.WaitToCatchUp(ctx)
	}

	var batch batcher.TransactionBatcher
	for {
		batch, err = rpc.SetupBatcher(
			ctx,
			ethclint,
			rollupAddress,
			l2ChainId,
			db,
			time.Duration(config.MaxBatchTime)*time.Second,
			batcherMode,
			dataSigner,
			broadcasterSettings,
			config.GasPriceUrl,
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

	srv := aggregator.NewServer(batch, rollupAddress, l2ChainId, db)
	web3Server, err := web3.GenerateWeb3Server(srv, nil, false, nil, metricsConfig)
	if err != nil {
		return err
	}
	errChan := make(chan error, 1)
	go func() {
		err := rpc.LaunchPublicServer(ctx, web3Server, config.RPC.Addr, config.RPC.Port, config.WS.Addr, config.WS.Port)
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
