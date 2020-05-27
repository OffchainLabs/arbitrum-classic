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

package utils

import (
	"flag"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type RollupArgs struct {
	ValidatorFolder string
	EthURL          string
	Address         common.Address
}

func ParseRollupCommand(fs *flag.FlagSet, startIndex int) RollupArgs {
	validatorFolder := fs.Arg(startIndex)
	ethURL := fs.Arg(startIndex + 1)
	addressString := fs.Arg(startIndex + 2)

	address := common.HexToAddress(addressString)

	return RollupArgs{
		ValidatorFolder: validatorFolder,
		EthURL:          ethURL,
		Address:         address,
	}
}

const RollupArgsString = "<validator_folder> <ethURL> <rollup_address>"
