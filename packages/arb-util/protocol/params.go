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
	"math/rand"
	
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type ChainParams struct {
	GracePeriod               *common.TimeBlocks
	ArbGasSpeedLimitPerSecond uint64
}

func NewRandomChainParams() ChainParams {
	return ChainParams{
		GracePeriod:               common.NewTimeBlocks(common.RandBigInt()),
		ArbGasSpeedLimitPerSecond: rand.Uint64(),
	}
}

func (cp ChainParams) WithGracePeriod(period *common.TimeBlocks) ChainParams {
	ret := cp
	ret.GracePeriod = period
	return ret
}

func (cp ChainParams) WithArbGasSpeedLimitPerSecond(limit uint64) ChainParams {
	ret := cp
	ret.ArbGasSpeedLimitPerSecond = limit
	return ret
}

func (cp ChainParams) Equals(cp2 ChainParams) bool {
	return cp.GracePeriod == cp2.GracePeriod && cp.ArbGasSpeedLimitPerSecond == cp2.ArbGasSpeedLimitPerSecond
}
