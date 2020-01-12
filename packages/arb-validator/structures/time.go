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

type TimeTicks struct {
	Val *big.Int
}

var _timeConversionFactor *big.Int
var _timeTicksPerSecond *big.Int

func init() {
	_timeTicksPerSecond = big.NewInt(1000)
	_timeConversionFactor = new(big.Int).Mul(big.NewInt(13), _timeTicksPerSecond)
}

func TimeFromBlockNum(blockNum *common.TimeBlocks) TimeTicks {
	return TimeTicks{new(big.Int).Mul(_timeConversionFactor, blockNum.AsInt())}
}

func TimeFromSeconds(seconds int64) TimeTicks {
	return TimeTicks{new(big.Int).Mul(_timeTicksPerSecond, big.NewInt(seconds))}
}

func (rt TimeTicks) Add(rt2 TimeTicks) TimeTicks {
	return TimeTicks{new(big.Int).Add(rt.Val, rt2.Val)}
}

func (rt TimeTicks) Cmp(rt2 TimeTicks) int {
	return rt.Val.Cmp(rt2.Val)
}

func (rt TimeTicks) MarshalToBuf() *TimeTicksBuf {
	return &TimeTicksBuf{
		Val: common.MarshalBigInt(rt.Val),
	}
}

func (rtb *TimeTicksBuf) Unmarshal() TimeTicks {
	return TimeTicks{rtb.Val.Unmarshal()}
}

func (rt TimeTicks) Equals(rt2 TimeTicks) bool {
	return rt.Val.Cmp(rt2.Val) == 0
}
