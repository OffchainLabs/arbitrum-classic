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

package nitroexport

import (
	"os"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/arblog"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

var logger = arblog.Logger.With().Str("component", "nitroexport").Logger()

func ExportState(arbcore core.ArbCore, height uint64, dirname string) error {
	cursor, err := arbcore.GetExecutionCursorAtEndOfBlock(height, true)
	if err != nil {
		return err
	}
	logger.Info().Uint64("block", height).Msg("taking machine")
	machine, err := arbcore.TakeMachine(cursor)
	if err != nil {
		return err
	}
	_, statsErr := os.Stat(dirname)
	if !os.IsNotExist(statsErr) {
		return errors.Errorf("%v already exists", dirname)
	}
	err = os.MkdirAll(dirname, os.ModePerm)
	if err != nil {
		return err
	}
	return arbcore.DumpArbosState(machine, height, dirname)
}
