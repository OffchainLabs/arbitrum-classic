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
)

type TimeBlocks big.Int

func NewTimeBlocks(val *big.Int) *TimeBlocks {
	return (*TimeBlocks)(val)
}

func (tb *TimeBlocks) Clone() *TimeBlocks {
	return NewTimeBlocks(new(big.Int).Set(tb.AsInt()))
}

func (tb *TimeBlocks) AsInt() *big.Int {
	return (*big.Int)(tb)
}

func (tb *TimeBlocks) Marshal() *TimeBlocksBuf {
	return &TimeBlocksBuf{Val: MarshalBigInt(tb.AsInt())}
}

func (tb *TimeBlocksBuf) Unmarshal() *TimeBlocks {
	return (*TimeBlocks)(tb.Val.Unmarshal())
}
