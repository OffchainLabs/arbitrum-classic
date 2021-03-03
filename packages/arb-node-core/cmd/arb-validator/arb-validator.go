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
	"io/ioutil"
	golog "log"
	"net/http"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/staker"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var logger zerolog.Logger
var pprofMux *http.ServeMux

func init() {
	pprofMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()
}

type ContractAddresses struct {
	validatorWallet string
	rollup          string
	validatorUtils  string
}

type Config struct {
	arbStorageDir     string
	contractAddresses ContractAddresses
	privateKey        string
	rpc               string
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
		logger.Fatal().Msg("Usage: arb-validator [folder]")
	}
	folder := os.Args[1]
	configPath := path.Join(folder, "config.json")
	configFile, err := os.Open(configPath)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Failed to open config.json in specified path")
	}
	configData, err := ioutil.ReadAll(configFile)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Failed to read config")
	}
	config := Config{
		arbStorageDir: path.Join(folder, "storage"),
	}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Failed to unmarshal config")
	}

	validatorAddress := ethcommon.HexToAddress(config.contractAddresses.validatorWallet)
	rollupAddr := ethcommon.HexToAddress(config.contractAddresses.rollup)
	validatorUtilsAddr := ethcommon.HexToAddress(config.contractAddresses.validatorUtils)
	auth := bind.NewKeyedTransactor(nil) // TODO
	client, err := ethutils.NewRPCEthClient(config.rpc)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error creating Ethereum RPC client")
	}

	strategy := staker.MakeNodesStrategy

	ctx := context.Background()

	storage, err := cmachine.NewArbStorage(config.arbStorageDir)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error creating ArbStorage")
	}
	defer func() {
		storage.CloseArbStorage()
	}()

	err = storage.Initialize(arbos.Path())
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

	<-stakerManager.RunInBackground(ctx, logger)
}
