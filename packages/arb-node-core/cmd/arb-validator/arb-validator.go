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
	"encoding/json"
	"flag"
	"io/ioutil"
	golog "log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	ethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
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
	logger = log.With().Caller().Str("component", "arb-validator").Logger()

	if len(os.Args) < 2 {
		usageStr := "Usage: arb-validator [folder] [RPC URL] [rollup address] [validator utils address] [strategy] " + cmdhelp.WalletArgsString
		logger.Fatal().Msg(usageStr)
	}
	flagSet := flag.NewFlagSet("validator", flag.ExitOnError)
	walletFlags := cmdhelp.AddWalletFlags(flagSet)
	enablePProf := flagSet.Bool("pprof", false, "enable profiling server")
	if err := flagSet.Parse(os.Args[6:]); err != nil {
		logger.Fatal().Err(err).Msg("failed parsing command line arguments")
	}

	if *enablePProf {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	folder := os.Args[1]

	rollupAddr := ethcommon.HexToAddress(os.Args[3])
	validatorUtilsAddr := ethcommon.HexToAddress(os.Args[4])
	auth, err := cmdhelp.GetKeystore(folder, walletFlags, flagSet)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error loading wallet keystore")
	}
	logger.Info().Str("address", auth.From.String()).Msg("Loaded wallet")
	client, err := ethutils.NewRPCEthClient(os.Args[2])
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error creating Ethereum RPC client")
	}
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
			logger.Fatal().Stack().Err(err).Msg("Failed to open chainState.json")
		}
	} else {
		chainStateData, err := ioutil.ReadAll(chainStateFile)
		if err != nil {
			logger.Fatal().Stack().Err(err).Msg("Failed to read chain state")
		}
		err = json.Unmarshal(chainStateData, &chainState)
		if err != nil {
			logger.Fatal().Stack().Err(err).Msg("Failed to unmarshal chain state")
		}
	}

	validatorAddress := ethcommon.Address{}
	if chainState.ValidatorWallet == "" {
		validatorAddress, _, _, err = ethbridgecontracts.DeployValidator(auth, client)
		if err != nil {
			logger.Fatal().Stack().Err(err).Msg("Failed to deploy validator wallet")
		}
		chainState.ValidatorWallet = validatorAddress.String()

		newChainStateData, err := json.Marshal(chainState)
		if err != nil {
			logger.Fatal().Stack().Err(err).Msg("Failed to marshal chain state")
		}
		if err := ioutil.WriteFile(chainStatePath, newChainStateData, 0644); err != nil {
			logger.Fatal().Stack().Err(err).Msg("Failed to write chain state config")
		}
	} else {
		validatorAddress = ethcommon.HexToAddress(chainState.ValidatorWallet)
	}

	ctx := context.Background()

	storage, err := cmachine.NewArbStorage(path.Join(folder, "arbStorage"))
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error creating ArbStorage")
	}
	defer func() {
		storage.CloseArbStorage()
	}()

	arbosPath := path.Join(folder, "arbos.mexe")
	err = storage.Initialize(arbosPath)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error initializing ArbStorage")
	}

	arbCore := storage.GetArbCore()
	started := arbCore.StartThread()
	if !started {
		logger.Fatal().Msg("Error starting ArbCore thread")
	}

	val, err := ethbridge.NewValidator(validatorAddress, rollupAddr, client, ethbridge.NewTransactAuth(auth))
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error creating validator wallet")
	}

	stakerManager, bridge, err := staker.NewStaker(ctx, arbCore, client, val, common.NewAddressFromEth(validatorUtilsAddr), strategy)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error setting up staker")
	}

	reader, err := staker.NewInboxReader(ctx, bridge, arbCore)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Failed to create inbox reader")
	}
	reader.Start(ctx)

	<-stakerManager.RunInBackground(ctx)
}
