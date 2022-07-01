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
	"flag"

	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/rs/zerolog"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

func AddLogFlags(fs *flag.FlagSet) (*string, *string) {
	gethLogLevel := fs.String("rpcloglevel", "info", "log level for rpc")
	arbLogLevel := fs.String("arbloglevel", "info", "log level for general arb node logging")
	return gethLogLevel, arbLogLevel
}

func ParseLogFlags(gethLogLevel, arbLogLevel *string, coreHandler gethlog.Handler) error {
	gethLevel, err := gethlog.LvlFromString(*gethLogLevel)
	if err != nil {
		return err
	}
	arbLevel, err := zerolog.ParseLevel(*arbLogLevel)
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(arbLevel)
	h := gethlog.LvlFilterHandler(gethLevel, coreHandler)
	h = gethlog.CallerFuncHandler(h)
	gethlog.Root().SetHandler(h)
	return nil
}

func PrintDatabaseMetadata(path string, coreConfig *configuration.Core) error {
	storage, err := cmachine.NewArbStorage(path, coreConfig)
	if err != nil {
		return err
	}
	defer storage.CloseArbStorage()

	storage.PrintDatabaseMetadata()

	return nil
}
