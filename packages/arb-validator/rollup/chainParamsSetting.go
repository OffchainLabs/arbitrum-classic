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

package rollup

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func DefaultChainParams() structures.ChainParams {
	gracePeriodInBlocks := int64(30)
	return structures.ChainParams{
		StakeRequirement:        big.NewInt(10),
		GracePeriod:             common.TimeFromBlockNum(common.NewTimeBlocks(big.NewInt(gracePeriodInBlocks))),
		MaxExecutionSteps:       1000000000,
		MaxTimeBoundsWidth:      20,
		ArbGasSpeedLimitPerTick: 20000000,
	}
}
