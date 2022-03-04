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
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
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

const (
	failLimit            = 6
	checkFrequency       = time.Second * 30
	blockCheckCountDelay = 5
)

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
	logger = arblog.Logger.With().Str("component", "arb-node").Logger()

	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running node")
	}
}

func printSampleUsage() {
	fmt.Printf("\n")
	fmt.Printf("Sample usage:                  arb-node --conf=<filename> \n")
	fmt.Printf("          or:  forwarder node: arb-node --l1.url=<L1 RPC> [optional arguments]\n\n")
	fmt.Printf("          or: aggregator node: arb-node --l1.url=<L1 RPC> --node.type=aggregator [optional arguments] %s\n", cmdhelp.WalletArgsString)
	fmt.Printf("          or:       sequencer: arb-node --l1.url=<L1 RPC> --node.type=sequencer [optional arguments] %s\n", cmdhelp.WalletArgsString)
}

func startup() error {
	ctx, cancelFunc, cancelChan := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	config, walletConfig, l1Client, l1ChainId, err := configuration.ParseNode(ctx)
	if err != nil || len(config.Persistent.GlobalConfig) == 0 || len(config.L1.URL) == 0 ||
		len(config.Rollup.Address) == 0 || len(config.BridgeUtilsAddress) == 0 ||
		((config.Node.Type != "sequencer") && len(config.Node.Sequencer.Lockout.Redis) != 0) ||
		((len(config.Node.Sequencer.Lockout.Redis) == 0) != (len(config.Node.Sequencer.Lockout.SelfRPCURL) == 0)) {
		printSampleUsage()
		if err != nil && !strings.Contains(err.Error(), "help requested") {
			fmt.Printf("%s\n", err.Error())
		}

		return nil
	}

	if config.Core.Database.Metadata {
		return cmdhelp.PrintDatabaseMetadata(config.GetNodeDatabasePath(), &config.Core)
	}

	if config.Core.Database.MakeValidator {
		// Exit immediately after converting database
		return cmdhelp.NodeToValidator(config)
	}

	badConfig := false
	if config.BridgeUtilsAddress == "" {
		badConfig = true
		fmt.Println("Missing --bridge-utils-address")
	}
	if config.Persistent.Chain == "" {
		badConfig = true
		fmt.Println("Missing --persistent.chain")
	}
	if config.Rollup.Address == "" {
		badConfig = true
		fmt.Println("Missing --rollup.address")
	}
	if config.Node.ChainID == 0 {
		badConfig = true
		fmt.Println("Missing --node.chain-id")
	}
	if config.Rollup.Machine.Filename == "" {
		badConfig = true
		fmt.Println("Missing --rollup.machine.filename")
	}

	var rpcMode web3.RpcMode
	if config.Node.Type == "forwarder" {
		if config.Node.Forwarder.Target == "" {
			badConfig = true
			fmt.Println("Forwarder node needs --node.forwarder.target")
		}

		if config.Node.Forwarder.RpcMode == "full" {
			rpcMode = web3.NormalMode
		} else if config.Node.Forwarder.RpcMode == "non-mutating" {
			rpcMode = web3.NonMutatingMode
		} else if config.Node.Forwarder.RpcMode == "forwarding-only" {
			rpcMode = web3.ForwardingOnlyMode
		} else {
			badConfig = true
			fmt.Printf("Unrecognized RPC mode %s", config.Node.Forwarder.RpcMode)
		}
	} else if config.Node.Type == "aggregator" {
		if config.Node.Aggregator.InboxAddress == "" {
			badConfig = true
			fmt.Println("Aggregator node needs --node.aggregator.inbox-address")
		}
	} else if config.Node.Type == "sequencer" {
		// Sequencer always waits
		config.WaitToCatchUp = true
	} else {
		badConfig = true
		fmt.Printf("Unrecognized node type %s", config.Node.Type)
	}

	if badConfig {
		return nil
	}

	if config.Node.Sequencer.Dangerous != (configuration.SequencerDangerous{}) {
		logger.
			Error().
			Interface("dangerousSequencerConfig", config.Node.Sequencer.Dangerous).
			Msg("sequencer starting up with dangerous options enabled!")
	}

	defer logger.Log().Msg("Cleanly shutting down node")

	if err := cmdhelp.ParseLogFlags(&config.Log.RPC, &config.Log.Core); err != nil {
		return err
	}

	if config.PProfEnable {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	l2ChainId := new(big.Int).SetUint64(config.Node.ChainID)
	rollupAddress := common.HexToAddress(config.Rollup.Address)
	logger.Info().
		Hex("chainaddress", rollupAddress.Bytes()).
		Hex("chainid", l2ChainId.Bytes()).
		Str("type", config.Node.Type).
		Int64("fromBlock", config.Rollup.FromBlock).
		Msg("Launching arbitrum node")

	mon, err := monitor.NewInitializedMonitor(config.GetNodeDatabasePath(), config.Rollup.Machine.Filename, &config.Core)
	if err != nil {
		return errors.Wrap(err, "error opening monitor")
	}
	defer mon.Close()

	metricsConfig := metrics.NewMetricsConfig(config.MetricsServer, &config.Healthcheck.MetricsPrefix)
	healthChan := make(chan nodehealth.Log, largeChannelBuffer)
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckMetrics", ValBool: config.Healthcheck.Metrics}
	healthChan <- nodehealth.Log{Config: true, Var: "disablePrimaryCheck", ValBool: !config.Healthcheck.Sequencer}
	healthChan <- nodehealth.Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: !config.Healthcheck.L1Node}
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckRPC", ValStr: config.Healthcheck.Addr + ":" + config.Healthcheck.Port}

	if config.Node.Type == "forwarder" {
		healthChan <- nodehealth.Log{Config: true, Var: "primaryHealthcheckRPC", ValStr: config.Node.Forwarder.Target}
	}
	healthChan <- nodehealth.Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: config.L1.URL}
	nodehealth.Init(healthChan)

	go func() {
		err := nodehealth.StartNodeHealthCheck(ctx, healthChan, metricsConfig.Registry)
		if err != nil {
			log.Error().Err(err).Msg("healthcheck server failed")
		}
	}()

	var sequencerFeed chan broadcaster.BroadcastFeedMessage
	if len(config.Feed.Input.URLs) == 0 {
		logger.Warn().Msg("Missing --feed.input.url so not subscribing to feed")
	} else {
		sequencerFeed = make(chan broadcaster.BroadcastFeedMessage, 4096)
		for _, url := range config.Feed.Input.URLs {
			broadcastClient := broadcastclient.NewBroadcastClient(url, nil, config.Feed.Input.Timeout)
			broadcastClient.ConnectInBackground(ctx, sequencerFeed)
		}
	}

	var inboxReader *monitor.InboxReader
	for {
		inboxReader, err = mon.StartInboxReader(
			ctx,
			l1Client,
			common.HexToAddress(config.Rollup.Address),
			config.Rollup.FromBlock,
			common.HexToAddress(config.BridgeUtilsAddress),
			healthChan,
			sequencerFeed,
			config.Node.InboxReader,
		)
		if err == nil {
			break
		}
		logger.Warn().Err(err).
			Str("url", config.L1.URL).
			Str("rollup", config.Rollup.Address).
			Str("bridgeUtils", config.BridgeUtilsAddress).
			Int64("fromBlock", config.Rollup.FromBlock).
			Msg("failed to start inbox reader, waiting and retrying")

		select {
		case <-ctx.Done():
			return errors.New("ctx cancelled StartInboxReader retry loop")
		case <-time.After(5 * time.Second):
		}
	}

	var dataSigner func([]byte) ([]byte, error)
	var batcherMode rpc.BatcherMode
	if config.Node.Type == "forwarder" {
		logger.Info().Str("forwardTxURL", config.Node.Forwarder.Target).Msg("Arbitrum node starting in forwarder mode")
		batcherMode = rpc.ForwarderBatcherMode{Config: config.Node.Forwarder}
	} else {
		var auth *bind.TransactOpts
		auth, dataSigner, err = cmdhelp.GetKeystore(config, walletConfig, l1ChainId, true)
		if err != nil {
			return errors.Wrap(err, "error running GetKeystore")
		}

		if config.Node.Sequencer.Dangerous.DisableBatchPosting {
			logger.Info().Hex("from", auth.From.Bytes()).Msg("Arbitrum node with disabled batch posting")
		} else {
			logger.Info().Hex("from", auth.From.Bytes()).Msg("Arbitrum node submitting batches")
		}

		if err := ethbridge.WaitForBalance(
			ctx,
			l1Client,
			common.Address{},
			common.NewAddressFromEth(auth.From),
		); err != nil {
			return errors.Wrap(err, "error waiting for balance")
		}

		if config.Node.Type == "sequencer" {
			batcherMode = rpc.SequencerBatcherMode{
				Auth:        auth,
				Core:        mon.Core,
				InboxReader: inboxReader,
			}
		} else {
			inboxAddress := common.HexToAddress(config.Node.Aggregator.InboxAddress)
			if config.Node.Aggregator.Stateful {
				batcherMode = rpc.StatefulBatcherMode{Auth: auth, InboxAddress: inboxAddress}
			} else {
				batcherMode = rpc.StatelessBatcherMode{Auth: auth, InboxAddress: inboxAddress}
			}
		}
	}

	nodeStore := mon.Storage.GetNodeStore()
	metricsConfig.RegisterNodeStoreMetrics(nodeStore)
	metricsConfig.RegisterArbCoreMetrics(mon.Core)
	db, txDBErrChan, err := txdb.New(ctx, mon.Core, nodeStore, &config.Node)
	if err != nil {
		return errors.Wrap(err, "error opening txdb")
	}
	defer db.Close()

	if config.WaitToCatchUp {
		inboxReader.WaitToCatchUp(ctx)
	}

	var batch batcher.TransactionBatcher
	errChan := make(chan error, 1)
	for {
		batch, err = rpc.SetupBatcher(
			ctx,
			l1Client,
			rollupAddress,
			l2ChainId,
			db,
			time.Duration(config.Node.Aggregator.MaxBatchTime)*time.Second,
			batcherMode,
			dataSigner,
			config,
			walletConfig,
		)
		lockoutConf := config.Node.Sequencer.Lockout
		if err == nil {
			seqBatcher, ok := batch.(*batcher.SequencerBatcher)
			if lockoutConf.Redis != "" {
				// Setup the lockout. This will take care of the initial delayed sequence.
				batch, err = rpc.SetupLockout(ctx, seqBatcher, mon.Core, inboxReader, lockoutConf, errChan)
			} else if ok {
				// Ensure we sequence delayed messages before opening the RPC.
				err = seqBatcher.SequenceDelayedMessages(ctx, false)
			}
		}
		if err == nil {
			go batch.Start(ctx)
			break
		}
		logger.Warn().Err(err).Msg("failed to setup batcher, waiting and retrying")

		select {
		case <-ctx.Done():
			return errors.New("ctx cancelled setup batcher")
		case <-time.After(5 * time.Second):
		}
	}

	var web3InboxReaderRef *monitor.InboxReader
	if config.Node.RPC.EnableL1Calls {
		web3InboxReaderRef = inboxReader
	}

	srv := aggregator.NewServer(batch, l2ChainId, db)
	serverConfig := web3.ServerConfig{
		Mode:          rpcMode,
		MaxCallAVMGas: config.Node.RPC.MaxCallGas * 100, // Multiply by 100 for arb gas to avm gas conversion
		DevopsStubs:   config.Node.RPC.EnableDevopsStubs,
	}
	web3Server, err := web3.GenerateWeb3Server(srv, nil, serverConfig, nil, web3InboxReaderRef)
	if err != nil {
		return err
	}
	go func() {
		err := rpc.LaunchPublicServer(ctx, web3Server, config.Node.RPC, config.Node.WS)
		if err != nil {
			errChan <- err
		}
	}()

	if config.Node.Type == "forwarder" && config.Node.Forwarder.Target != "" {
		go func() {
			clnt, err := ethclient.DialContext(ctx, config.Node.Forwarder.Target)
			if err != nil {
				log.Warn().Err(err).Msg("failed to connect to forward target")
				clnt = nil
			}
			failCount := 0
			for {
				valid, err := checkBlockHash(ctx, clnt, db)
				if err != nil {
					log.Warn().Err(err).Msg("failed to lookup blockhash for consistency check")
					clnt, err = ethclient.DialContext(ctx, config.Node.Forwarder.Target)
					if err != nil {
						log.Warn().Err(err).Msg("failed to connect to forward target")
						clnt = nil
					}
				} else {
					if !valid {
						failCount++
					} else {
						failCount = 0
					}
				}
				if failCount >= failLimit {
					log.Error().Msg("exiting due to repeated block hash mismatches")
					cancelFunc()
					return
				}
				select {
				case <-ctx.Done():
					return
				case <-time.After(checkFrequency):
				}
			}
		}()

	}

	select {
	case err := <-txDBErrChan:
		return err
	case err := <-errChan:
		return err
	case <-cancelChan:
		return nil
	}
}

func checkBlockHash(ctx context.Context, clnt *ethclient.Client, db *txdb.TxDB) (bool, error) {
	if clnt == nil {
		return false, errors.New("need a client to check block hash")
	}
	blockCount, err := db.BlockCount()
	if err != nil {
		return false, err
	}
	if blockCount < blockCheckCountDelay {
		return true, nil
	}
	// Use a small block delay here in case the upstream node isn't full caught up
	block, err := db.GetBlock(blockCount - blockCheckCountDelay)
	if err != nil {
		return false, err
	}
	remoteHeader, err := clnt.HeaderByNumber(ctx, block.Header.Number)
	if err != nil {
		return false, err
	}
	if remoteHeader.Hash() == block.Header.Hash() {
		return true, nil
	}
	logger.Warn().
		Str("remote", remoteHeader.Hash().Hex()).
		Str("local", block.Header.Hash().Hex()).
		Msg("mismatched block header")
	return false, nil
}
