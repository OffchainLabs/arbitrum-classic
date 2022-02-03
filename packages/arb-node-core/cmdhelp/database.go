/*
 * Copyright 2021, Offchain Labs, Inc.
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

package cmdhelp

import (
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/pkg/errors"
	"os"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

func NodeToValidator(config *configuration.Config) error {
	// Convert node database to validator database
	if !configuration.DatabaseInDirectory(config.GetNodeDatabasePath()) {
		return errors.New("node database doesn't exist in " + config.GetNodeDatabasePath())
	}

	if configuration.DatabaseInDirectory(config.GetValidatorDatabasePath()) {
		return errors.New("validator database already exists in " + config.GetValidatorDatabasePath())
	}

	err := os.Rename(config.GetNodeDatabasePath(), config.GetValidatorDatabasePath())
	if err != nil {
		return errors.Wrap(err, "unable to move node database to validator database")
	}

	storage, err := cmachine.NewArbStorage(config.GetValidatorDatabasePath(), &config.Core)
	if err != nil {
		return err
	}

	err = storage.CleanupValidator()
	if err != nil {
		return err
	}

	logger.Info().Str("path", config.Persistent.Chain).Msg("Database has been converted to validator database")

	return nil
}
