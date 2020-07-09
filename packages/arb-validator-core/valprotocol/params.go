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

package valprotocol

import (
	"math/big"
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ChainParams struct {
	StakeRequirement        *big.Int         // in Wei
	GracePeriod             common.TimeTicks // in Ticks
	MaxExecutionSteps       uint64
	MaxBlockBoundsWidth     uint64 // in blocks
	ArbGasSpeedLimitPerTick uint64 // in ArbGas per tick
}

func NewRandomChainParams() ChainParams {
	return ChainParams{
		StakeRequirement:        common.RandBigInt(),
		GracePeriod:             common.TimeTicks{Val: common.RandBigInt()},
		MaxExecutionSteps:       rand.Uint64(),
		MaxBlockBoundsWidth:     rand.Uint64(),
		ArbGasSpeedLimitPerTick: rand.Uint64(),
	}
}

func (cp ChainParams) WithStakeRequirement(amountInWei *big.Int) ChainParams {
	ret := cp
	ret.StakeRequirement = amountInWei
	return ret
}

func (cp ChainParams) WithGracePeriod(period common.TimeTicks) ChainParams {
	ret := cp
	ret.GracePeriod = period
	return ret
}

func (cp ChainParams) WithGracePeriodBlocks(period common.TimeBlocks) ChainParams {
	return cp.WithGracePeriod(common.TicksFromBlockNum(&period))
}

func (cp ChainParams) WithMaxExecutionSteps(steps uint64) ChainParams {
	ret := cp
	ret.MaxExecutionSteps = steps
	return ret
}

func (cp ChainParams) WithMaxBlocksBoundsWidth(width uint64) ChainParams {
	ret := cp
	ret.MaxBlockBoundsWidth = width
	return ret
}

func (cp ChainParams) WithArbGasSpeedLimitPerTick(limit uint64) ChainParams {
	ret := cp
	ret.ArbGasSpeedLimitPerTick = limit
	return ret
}

func (params ChainParams) MarshalToBuf() *ChainParamsBuf {
	return &ChainParamsBuf{
		StakeRequirement:        common.MarshalBigInt(params.StakeRequirement),
		GracePeriod:             params.GracePeriod.MarshalToBuf(),
		MaxExecutionSteps:       params.MaxExecutionSteps,
		MaxBlockBoundsWidth:     params.MaxBlockBoundsWidth,
		ArbGasSpeedLimitPerTick: params.ArbGasSpeedLimitPerTick,
	}
}

func (m *ChainParamsBuf) Unmarshal() ChainParams {
	return ChainParams{
		StakeRequirement:        m.StakeRequirement.Unmarshal(),
		GracePeriod:             m.GracePeriod.Unmarshal(),
		MaxExecutionSteps:       m.MaxExecutionSteps,
		MaxBlockBoundsWidth:     m.MaxBlockBoundsWidth,
		ArbGasSpeedLimitPerTick: m.ArbGasSpeedLimitPerTick,
	}
}

func (cp ChainParams) Equals(cp2 ChainParams) bool {
	return cp.StakeRequirement.Cmp(cp2.StakeRequirement) == 0 &&
		cp.GracePeriod == cp2.GracePeriod &&
		cp.MaxExecutionSteps == cp2.MaxExecutionSteps &&
		cp.MaxBlockBoundsWidth == cp2.MaxBlockBoundsWidth &&
		cp.ArbGasSpeedLimitPerTick == cp2.ArbGasSpeedLimitPerTick
}
