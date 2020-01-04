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

package structures

import (
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
)

type ChainParams struct {
	stakeRequirement        *big.Int
	gracePeriod             TimeTicks
	maxExecutionSteps       uint32
	arbGasSpeedLimitPerTick uint64
}

func (params ChainParams) MarshalToBuf() *ChainParamsBuf {
	return &ChainParamsBuf{
		StakeRequirement:  utils.MarshalBigInt(params.stakeRequirement),
		GracePeriod:       params.gracePeriod.MarshalToBuf(),
		MaxExecutionSteps: params.maxExecutionSteps,
	}
}

func (m *ChainParamsBuf) Unmarshal() ChainParams {
	return ChainParams{
		stakeRequirement:  utils.UnmarshalBigInt(m.StakeRequirement),
		gracePeriod:       m.GracePeriod.Unmarshal(),
		maxExecutionSteps: m.MaxExecutionSteps,
	}
}

func (cp ChainParams) Equals(cp2 ChainParams) bool {
	return cp.stakeRequirement.Cmp(cp2.stakeRequirement) == 0 &&
		cp.gracePeriod == cp2.gracePeriod &&
		cp.maxExecutionSteps == cp2.maxExecutionSteps
}
