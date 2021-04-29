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
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/nodehealth"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/staker"
	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcastclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var logger zerolog.Logger
var pprofMux *http.ServeMux

func init() {
	pprofMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()
}

type ChainState struct {
	ValidatorWallet string `json:"validatorWallet"`
}

func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Print line number that log was created on
	logger = log.With().Caller().Stack().Str("component", "arb-validator").Logger()

	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running validator")
	}
}

func startup() error {
	defer logger.Log().Msg("Cleanly shutting down validator")
	ctx, cancelFunc, cancelChan := cmdhelp.CreateLaunchContext()
	defer cancelFunc()

	const largeChannelBuffer = 200
	healthChan := make(chan nodehealth.Log, largeChannelBuffer)

	go func() {
		err := nodehealth.StartNodeHealthCheck(ctx, healthChan)
		if err != nil {
			log.Error().Err(err).Msg("healthcheck server failed")
		}
	}()

	if len(os.Args) < 2 {
		usageStr := "Usage: arb-validator [folder] [RPC URL] [rollup address] [validator utils address] [strategy] " + cmdhelp.WalletArgsString
		fmt.Println(usageStr)
		return errors.New("invalid arguments")
	}
	flagSet := flag.NewFlagSet("validator", flag.ExitOnError)
	walletFlags := cmdhelp.AddWalletFlags(flagSet)
	enablePProf := flagSet.Bool("pprof", false, "enable profiling server")
	gethLogLevel, arbLogLevel := cmdhelp.AddLogFlags(flagSet)

	//Healthcheck Config
	disablePrimaryCheck := flagSet.Bool("disable-primary-check", true, "disable checking the health of the primary")
	disableOpenEthereumCheck := flagSet.Bool("disable-openethereum-check", false, "disable checking the health of the OpenEthereum node")
	healthcheckMetrics := flagSet.Bool("metrics", false, "enable prometheus endpoint")
	healthcheckRPC := flagSet.String("healthcheck-rpc", "", "address to bind the healthcheck RPC to")
	sequencerWebsocketURL := flagSet.String("sequencer-websocket-url", "", "websocket address of sequencer feed")

	if err := flagSet.Parse(os.Args[6:]); err != nil {
		return errors.Wrap(err, "failed parsing command line arguments")
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

	broadcastClient := broadcastclient.NewBroadcastClient(*sequencerWebsocketURL, nil)
	sequencerFeed, err := broadcastClient.Connect()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to start broadcastclient")
	}

	folder := os.Args[1]
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckMetrics", ValBool: *healthcheckMetrics}
	healthChan <- nodehealth.Log{Config: true, Var: "disablePrimaryCheck", ValBool: *disablePrimaryCheck}
	healthChan <- nodehealth.Log{Config: true, Var: "disableOpenEthereumCheck", ValBool: *disableOpenEthereumCheck}
	healthChan <- nodehealth.Log{Config: true, Var: "healthcheckRPC", ValStr: *healthcheckRPC}
	healthChan <- nodehealth.Log{Config: true, Var: "openethereumHealthcheckRPC", ValStr: os.Args[2]}
	nodehealth.Init(healthChan)

	client, err := ethutils.NewRPCEthClient(os.Args[2])
	if err != nil {
		return errors.Wrap(err, "error creating Ethereum RPC client")
	}
	var l1ChainId *big.Int
	for {
		l1ChainId, err = client.ChainID(ctx)
		if err == nil {
			break
		}
		logger.Warn().Err(err).Msg("Error getting chain ID")
		time.Sleep(time.Second * 5)
	}

	logger.Debug().Str("chainid", l1ChainId.String()).Msg("connected to l1 chain")

	rollupAddr := ethcommon.HexToAddress(os.Args[3])
	validatorUtilsAddr := ethcommon.HexToAddress(os.Args[4])
	auth, _, err := cmdhelp.GetKeystore(folder, walletFlags, flagSet, l1ChainId)
	if err != nil {
		return errors.Wrap(err, "error loading wallet keystore")
	}
	logger.Info().Str("address", auth.From.String()).Msg("Loaded wallet")

	strategyString := os.Args[5]
	var strategy staker.Strategy
	if strategyString == "MakeNodes" {
		strategy = staker.MakeNodesStrategy
	} else if strategyString == "StakeLatest" {
		strategy = staker.StakeLatestStrategy
	} else if strategyString == "Defensive" {
		strategy = staker.DefensiveStrategy
	} else {
		return errors.New("unsupported strategy specified. Currently supported: MakeNodes, StakeLatest")
	}

	chainState := ChainState{}
	chainStatePath := path.Join(folder, "chainState.json")
	chainStateFile, err := os.Open(chainStatePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return errors.Wrap(err, "failed to open chainState.json")
		}
	} else {
		chainStateData, err := ioutil.ReadAll(chainStateFile)
		if err != nil {
			return errors.Wrap(err, "failed to read chain state")
		}
		err = json.Unmarshal(chainStateData, &chainState)
		if err != nil {
			return errors.Wrap(err, "failed to unmarshal chain state")
		}
	}

	validatorAddress := ethcommon.Address{}
	if chainState.ValidatorWallet == "" {
		for {
			validatorAddress, _, _, err = ethbridgecontracts.DeployValidator(auth, client)
			if err == nil {
				break
			}
			logger.Warn().Err(err).
				Str("sender", auth.From.Hex()).
				Msg("Failed to deploy validator wallet")
			time.Sleep(time.Second * 5)
		}
		chainState.ValidatorWallet = validatorAddress.String()

		newChainStateData, err := json.Marshal(chainState)
		if err != nil {
			return errors.Wrap(err, "failed to marshal chain state")
		}
		if err := ioutil.WriteFile(chainStatePath, newChainStateData, 0644); err != nil {
			return errors.Wrap(err, "failed to write chain state config")
		}
	} else {
		validatorAddress = ethcommon.HexToAddress(chainState.ValidatorWallet)
	}

	dbPath := path.Join(folder, "arbStorage")
	arbosPath := path.Join(folder, "arbos.mexe")
	mon, err := monitor.NewMonitor(dbPath, arbosPath)
	if err != nil {
		return errors.Wrap(err, "error opening monitor")
	}
	defer mon.Close()

	valAuth, err := ethbridge.NewTransactAuth(ctx, client, auth)
	if err != nil {
		return errors.Wrap(err, "error creating connecting to chain")
	}
	val, err := ethbridge.NewValidator(validatorAddress, rollupAddr, client, valAuth)
	if err != nil {
		return errors.Wrap(err, "error creating validator wallet")
	}

	stakerManager, _, err := staker.NewStaker(ctx, mon.Core, client, val, common.NewAddressFromEth(validatorUtilsAddr), strategy)
	if err != nil {
		return errors.Wrap(err, "error setting up staker")
	}

	chainMachineHash, err := stakerManager.GetInitialMachineHash(ctx)
	if err != nil {
		return errors.Wrap(err, "error checking initial chain state")
	}
	initialExecutionCursor, err := mon.Core.GetExecutionCursor(big.NewInt(0))
	if err != nil {
		return errors.Wrap(err, "error loading initial ArbCore machine")
	}
	initialMachineHash, err := initialExecutionCursor.MachineHash()
	if err != nil {
		return errors.Wrap(err, "error getting initial machine hash")
	}
	if initialMachineHash != chainMachineHash {
		return errors.Errorf("Initial machine hash loaded from arbos.mexe doesn't match chain's initial machine hash: chain %v, arbCore %v", hexutil.Encode(chainMachineHash[:]), initialMachineHash)
	}

	_, err = mon.StartInboxReader(ctx, client, common.NewAddressFromEth(rollupAddr), healthChan, sequencerFeed)
	if err != nil {
		return errors.Wrap(err, "failed to create inbox reader")
	}

	logger.Info().Int("strategy", int(strategy)).Msg("Initialized validator")
	select {
	case <-cancelChan:
		return nil
	case <-stakerManager.RunInBackground(ctx):
		return nil
	}
}
