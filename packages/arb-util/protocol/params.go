/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

package protocol

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
	"math/rand"
)

// If StakeToken is 0, stake requirement is ETH measured in Wei, otherwise stake requirement is units of stakeStoken
type ChainParams struct {
	StakeRequirement        *big.Int
	StakeToken              common.Address
	GracePeriod             *common.TimeBlocks
	MaxExecutionSteps       uint64
	ArbGasSpeedLimitPerTick uint64 // in ArbGas per tick
}

func NewRandomChainParams() ChainParams {
	return ChainParams{
		StakeRequirement:        common.RandBigInt(),
		StakeToken:              common.RandAddress(),
		GracePeriod:             common.NewTimeBlocks(common.RandBigInt()),
		MaxExecutionSteps:       rand.Uint64(),
		ArbGasSpeedLimitPerTick: rand.Uint64(),
	}
}

func (cp ChainParams) WithStakeRequirement(amount *big.Int) ChainParams {
	ret := cp
	ret.StakeRequirement = amount
	return ret
}

func (cp ChainParams) WithStakeToken(address common.Address) ChainParams {
	ret := cp
	ret.StakeToken = address
	return ret
}

func (cp ChainParams) WithGracePeriod(period *common.TimeBlocks) ChainParams {
	ret := cp
	ret.GracePeriod = period
	return ret
}

func (cp ChainParams) WithMaxExecutionSteps(steps uint64) ChainParams {
	ret := cp
	ret.MaxExecutionSteps = steps
	return ret
}

func (cp ChainParams) WithArbGasSpeedLimitPerTick(limit uint64) ChainParams {
	ret := cp
	ret.ArbGasSpeedLimitPerTick = limit
	return ret
}

func (cp ChainParams) Equals(cp2 ChainParams) bool {
	return cp.StakeRequirement.Cmp(cp2.StakeRequirement) == 0 &&
		cp.StakeToken == cp2.StakeToken &&
		cp.GracePeriod == cp2.GracePeriod &&
		cp.MaxExecutionSteps == cp2.MaxExecutionSteps &&
		cp.ArbGasSpeedLimitPerTick == cp2.ArbGasSpeedLimitPerTick
}
