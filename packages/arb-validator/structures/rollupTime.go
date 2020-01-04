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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/utils"
)

type RollupTime struct {
	Val *big.Int
}

var _timeConversionFactor *big.Int

func init() {
	_timeConversionFactor = big.NewInt(13000)
}

func RollupTimeFromBlockNum(blockNum *big.Int) RollupTime {
	return RollupTime{new(big.Int).Mul(_timeConversionFactor, blockNum)}
}

func (rt RollupTime) Add(rt2 RollupTime) RollupTime {
	return RollupTime{new(big.Int).Add(rt.Val, rt2.Val)}
}

func (rt RollupTime) Cmp(rt2 RollupTime) int {
	return rt.Val.Cmp(rt2.Val)
}

func (rt RollupTime) MarshalToBuf() *RollupTimeBuf {
	return &RollupTimeBuf{
		Val: utils.MarshalBigInt(rt.Val),
	}
}

func (rtb *RollupTimeBuf) Unmarshal() RollupTime {
	return RollupTime{utils.UnmarshalBigInt(rtb.Val)}
}

func (rt RollupTime) Equals(rt2 RollupTime) bool {
	return rt.Val.Cmp(rt2.Val) == 0
}
