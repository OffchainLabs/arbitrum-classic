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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ChainParams struct {
	StakeRequirement        *big.Int
	GracePeriod             common.TimeTicks
	MaxExecutionSteps       uint32
	ArbGasSpeedLimitPerTick uint64
}

func (params ChainParams) MarshalToBuf() *ChainParamsBuf {
	return &ChainParamsBuf{
		StakeRequirement:  common.MarshalBigInt(params.StakeRequirement),
		GracePeriod:       params.GracePeriod.MarshalToBuf(),
		MaxExecutionSteps: params.MaxExecutionSteps,
	}
}

func (m *ChainParamsBuf) Unmarshal() ChainParams {
	return ChainParams{
		StakeRequirement:  m.StakeRequirement.Unmarshal(),
		GracePeriod:       m.GracePeriod.Unmarshal(),
		MaxExecutionSteps: m.MaxExecutionSteps,
	}
}

func (cp ChainParams) Equals(cp2 ChainParams) bool {
	return cp.StakeRequirement.Cmp(cp2.StakeRequirement) == 0 &&
		cp.GracePeriod == cp2.GracePeriod &&
		cp.MaxExecutionSteps == cp2.MaxExecutionSteps
}
