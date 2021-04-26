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
	"encoding/hex"
	"encoding/json"
	"flag"
	"io/ioutil"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/nodehealth"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/staker"
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

	ctx, cancelFunc := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cancelFunc()
	}()

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
		logger.Fatal().Msg(usageStr)
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

	if err := flagSet.Parse(os.Args[6:]); err != nil {
		logger.Fatal().Err(err).Msg("failed parsing command line arguments")
	}
	if err := cmdhelp.ParseLogFlags(gethLogLevel, arbLogLevel); err != nil {
		logger.Fatal().Err(err).Send()
	}

	if *enablePProf {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
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
		logger.Fatal().Err(err).Msg("Error creating Ethereum RPC client")
	}
	var l1ChainId *big.Int
	for {
		l1ChainId, err = client.ChainID(context.Background())
		if err == nil {
			break
		}
		logger.Warn().Err(err).Msg("Error getting chain ID")
		time.Sleep(time.Second * 5)
	}

	logger.Debug().Str("chainid", l1ChainId.String()).Msg("connected to l1 chain")

	rollupAddr := ethcommon.HexToAddress(os.Args[3])
	validatorUtilsAddr := ethcommon.HexToAddress(os.Args[4])
	auth, err := cmdhelp.GetKeystore(folder, walletFlags, flagSet, l1ChainId)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error loading wallet keystore")
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
		logger.Fatal().Msg("Unsupported strategy specified. Currently supported: MakeNodes, StakeLatest")
	}

	chainState := ChainState{}
	chainStatePath := path.Join(folder, "chainState.json")
	chainStateFile, err := os.Open(chainStatePath)
	if err != nil {
		if !os.IsNotExist(err) {
			logger.Fatal().Err(err).Msg("Failed to open chainState.json")
		}
	} else {
		chainStateData, err := ioutil.ReadAll(chainStateFile)
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to read chain state")
		}
		err = json.Unmarshal(chainStateData, &chainState)
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to unmarshal chain state")
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
			logger.Fatal().Err(err).Msg("Failed to marshal chain state")
		}
		if err := ioutil.WriteFile(chainStatePath, newChainStateData, 0644); err != nil {
			logger.Fatal().Err(err).Msg("Failed to write chain state config")
		}
	} else {
		validatorAddress = ethcommon.HexToAddress(chainState.ValidatorWallet)
	}

	storage, err := cmachine.NewArbStorage(path.Join(folder, "arbStorage"))
	if err != nil {
		logger.Fatal().Err(err).Msg("Error creating ArbStorage")
	}
	defer func() {
		storage.CloseArbStorage()
	}()
	dbPath := path.Join(folder, "arbStorage")
	arbosPath := path.Join(folder, "arbos.mexe")
	mon, err := monitor.NewMonitor(dbPath, arbosPath)
	if err != nil {
		logger.Fatal().Err(err).Msg("error opening mon")
	}
	defer mon.Close()

	valAuth, err := ethbridge.NewTransactAuth(ctx, client, auth)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error creating connecting to chain")
	}
	val, err := ethbridge.NewValidator(validatorAddress, rollupAddr, client, valAuth)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error creating validator wallet")
	}

	stakerManager, _, err := staker.NewStaker(ctx, mon.Core, client, val, common.NewAddressFromEth(validatorUtilsAddr), strategy)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error setting up staker")
	}

	chainMachineHash, err := stakerManager.GetInitialMachineHash(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error checking initial chain state")
	}
	initialExecutionCursor, err := mon.Core.GetExecutionCursor(big.NewInt(0))
	if err != nil {
		logger.Fatal().Err(err).Msg("Error loading initial ArbCore machine")
	}
	initialMachineHash, err := initialExecutionCursor.MachineHash()
	if err != nil {
		logger.Fatal().Err(err).Msg("Error getting initial machine hash")
	}
	if initialMachineHash != chainMachineHash {
		logger.Fatal().Str("chain", hex.EncodeToString(chainMachineHash[:])).Str("arbCore", hex.EncodeToString(initialMachineHash[:])).Msg("Initial machine hash loaded from arbos.mexe doesn't match chain's initial machine hash")
	}

	reader, err := mon.StartInboxReader(ctx, client, common.NewAddressFromEth(rollupAddr), healthChan)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to create inbox reader")
	}
	defer reader.Stop()

	logger.Info().Int("strategy", int(strategy)).Msg("Initialized validator")
	<-stakerManager.RunInBackground(ctx)
}
