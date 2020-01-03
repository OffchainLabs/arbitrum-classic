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

package rollup

import "math/big"

type RollupTime struct {
	val *big.Int
}

var _timeConversionFactor *big.Int

func init() {
	_timeConversionFactor = big.NewInt(13000)
}

func RollupTimeFromBlockNum(blockNum *big.Int) RollupTime {
	return RollupTime{new(big.Int).Mul(_timeConversionFactor, blockNum)}
}

func (rt RollupTime) Add(rt2 RollupTime) RollupTime {
	return RollupTime{new(big.Int).Add(rt.val, rt2.val)}
}

func (rt RollupTime) Cmp(rt2 RollupTime) int {
	return rt.val.Cmp(rt2.val)
}

func (rt RollupTime) MarshalToBuf() *RollupTimeBuf {
	return &RollupTimeBuf{
		Val: marshalBigInt(rt.val),
	}
}

func (rtb *RollupTimeBuf) Unmarshal() RollupTime {
	return RollupTime{unmarshalBigInt(rtb.Val)}
}
