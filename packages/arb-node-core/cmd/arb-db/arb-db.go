/*
 * Copyright 2022, Offchain Labs, Inc.
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
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/pkg/errors"
	golog "log"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

var logger zerolog.Logger

func main() {
	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Print line number that log was created on
	logger = log.With().Caller().Stack().Str("component", "arb-db").Logger()

	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running arb-db")
	}
}

func startup() error {
	config, err := configuration.ParseDBTool()
	if err != nil || len(config.Persistent.Chain) == 0 {
		fmt.Printf("\n")
		fmt.Printf("Sample usage: %s --persistent.chain='.arbitrum/mainnet' --core.database.metadata\n", os.Args[0])
		fmt.Printf("              %s --persistent.chain='.arbitrum/mainnet' --core.database.make-validator\n", os.Args[0])
		fmt.Printf("              %s --persistent.chain='.arbitrum/mainnet' --core.database.prune-on-startup\n", os.Args[0])
		if err != nil && !strings.Contains(err.Error(), "help requested") {
			fmt.Printf("%s\n", err.Error())
		}

		return nil
	}

	if config.Core.Database.MakeValidator {
		// Exit immediately after converting database
		return cmdhelp.NodeToValidator(config)
	}

	// Make sure arbcore does not continue to run
	config.Core.Database.ExitAfter = true

	var databasePath string
	databasePath = config.GetNodeDatabasePath()
	if !configuration.DatabaseInDirectory(databasePath) {
		databasePath = config.GetValidatorDatabasePath()
		if !configuration.DatabaseInDirectory(databasePath) {
			// The db directory does not exist, try just the supplied chain directory
			databasePath = config.Persistent.Chain
			if !configuration.DatabaseInDirectory(databasePath) {
				return errors.New("unable to access database in " + databasePath)
			}
		}
	}

	println("Using database: ", databasePath)

	if config.Core.Database.Metadata {
		if err = cmdhelp.PrintDatabaseMetadata(databasePath, &config.Core); err != nil {
			return errors.New("issue printing database " + databasePath)
		}
	}

	storage, err := cmachine.NewArbStorage(databasePath, &config.Core)
	if err != nil {
		return err
	}

	// Any important errors already printed
	_ = storage.ApplyConfig()

	return nil
}
