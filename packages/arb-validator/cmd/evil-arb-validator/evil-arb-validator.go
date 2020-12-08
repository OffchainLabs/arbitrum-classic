/*
 * Copyright 2019, Offchain Labs, Inc.
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
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	golog "log"
	"math/big"
	"os"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/cmdhelper"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupmanager"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rolluptest"
)

var logger zerolog.Logger

// Launches the rollup validator with the following command line arguments:
// 1) Compiled Arbitrum bytecode file
// 2) private key file
// 3) Global EthBridge addresses json file
// 4) ethURL
func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	// Print line number that log was created on
	logger = log.With().Caller().Str("component", "evil-arb-validator").Logger()

	// Check number of args
	flag.Parse()
	switch os.Args[1] {
	case "validate":
		if err := cmdhelper.ValidateRollupChain("evil-arb-validator", createEvilManager); err != nil {
			logger.Fatal().Stack().Err(err).Msg("Error with ValidateRollupChain")
		}
	default:
	}
}

func createEvilManager(ctx context.Context, rollupAddress common.Address, client arbbridge.ArbClient, contractFile string, dbPath string) (*rollupmanager.Manager, error) {
	cp, err := rolluptest.NewEvilRollupCheckpointer(
		rollupAddress,
		dbPath,
		big.NewInt(100),
		false,
	)
	if err != nil {
		return nil, err
	}
	return rollupmanager.CreateManagerAdvanced(
		ctx,
		rollupAddress,
		true,
		client,
		cp,
		contractFile,
	)
}
