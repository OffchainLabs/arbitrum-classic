/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

type TimeBlocks big.Int

var _durationPerBlock time.Duration

func init() {
	_durationPerBlock = time.Duration(2) * time.Second
}

func SetDurationPerBlock(d time.Duration) {
	_durationPerBlock = d
}

func NewTimeBlocks(val *big.Int) *TimeBlocks {
	return (*TimeBlocks)(val)
}

func NewTimeBlocksInt(val int64) *TimeBlocks {
	return NewTimeBlocks(big.NewInt(val))
}

func (tb *TimeBlocks) Clone() *TimeBlocks {
	return NewTimeBlocks(new(big.Int).Set(tb.AsInt()))
}

func (tb *TimeBlocks) AsInt() *big.Int {
	return (*big.Int)(tb)
}

func BlocksFromSeconds(seconds int64) *TimeBlocks {
	return (*TimeBlocks)(big.NewInt(int64(time.Duration(seconds) * time.Second / AverageDurationPerBlock)))
}

func (tb *TimeBlocks) Duration() time.Duration {
	return TimeFromBlockNum(tb).Duration()
}

func (tb *TimeBlocks) Cmp(tb2 *TimeBlocks) int {
	return (*big.Int)(tb).Cmp((*big.Int)(tb2))
}

func (tb *TimeBlocks) Marshal() *TimeBlocksBuf {
	return &TimeBlocksBuf{Val: MarshalBigInt(tb.AsInt())}
}

func (tb *TimeBlocksBuf) Unmarshal() *TimeBlocks {
	return (*TimeBlocks)(tb.Val.Unmarshal())
}
