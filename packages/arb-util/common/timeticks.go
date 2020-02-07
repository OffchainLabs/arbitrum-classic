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

package common

import (
	"math/big"
	"time"
)

const TicksPerBlock = int64(1000)

type TimeTicks struct {
	Val *big.Int
}

func TicksFromBlockNum(blockNum *TimeBlocks) TimeTicks {
	return TimeTicks{new(big.Int).Mul(big.NewInt(TicksPerBlock), blockNum.AsInt())}
}

func TicksFromSeconds(seconds int64) TimeTicks {
	return TimeTicks{big.NewInt(int64(time.Duration(seconds*TicksPerBlock) * time.Second / _durationPerBlock))}
}

func (rt TimeTicks) Clone() TimeTicks {
	return TimeTicks{Val: new(big.Int).Set(rt.Val)}
}

func (rt TimeTicks) Add(rt2 TimeTicks) TimeTicks {
	return TimeTicks{new(big.Int).Add(rt.Val, rt2.Val)}
}

func (rt TimeTicks) Cmp(rt2 TimeTicks) int {
	return rt.Val.Cmp(rt2.Val)
}

func (rt TimeTicks) Duration() time.Duration {
	return time.Duration(rt.Val.Int64()) * _durationPerBlock / time.Duration(TicksPerBlock)
}

func (rt TimeTicks) MarshalToBuf() *TimeTicksBuf {
	return &TimeTicksBuf{
		Val: MarshalBigInt(rt.Val),
	}
}

func (rtb *TimeTicksBuf) Unmarshal() TimeTicks {
	return TimeTicks{rtb.Val.Unmarshal()}
}

func (rt TimeTicks) Equals(rt2 TimeTicks) bool {
	return rt.Val.Cmp(rt2.Val) == 0
}

func (rt TimeTicks) String() string {
	return rt.Val.String()
}
