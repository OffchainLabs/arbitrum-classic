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

package protocol

import (
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type TimeBounds struct {
	LowerBoundBlock *common.TimeBlocks
	UpperBoundBlock *common.TimeBlocks
}

func sortedRandomBigInts() (*big.Int, *big.Int) {
	a := common.RandBigInt()
	b := common.RandBigInt()
	if a.Cmp(b) > 0 {
		a, b = b, a
	}
	return b, a
}

func NewRandomTimeBounds() *TimeBounds {
	lowerBlockBound, upperBlockBound := sortedRandomBigInts()
	return &TimeBounds{
		LowerBoundBlock: common.NewTimeBlocks(upperBlockBound),
		UpperBoundBlock: common.NewTimeBlocks(lowerBlockBound),
	}
}

func (tb *TimeBounds) MarshalToBuf() *TimeBoundsBlocksBuf {
	return &TimeBoundsBlocksBuf{
		LowerBoundBlock: tb.LowerBoundBlock.Marshal(),
		UpperBoundBlock: tb.UpperBoundBlock.Marshal(),
	}
}

func (x *TimeBoundsBlocksBuf) Unmarshal() *TimeBounds {
	return &TimeBounds{
		LowerBoundBlock: x.LowerBoundBlock.Unmarshal(),
		UpperBoundBlock: x.UpperBoundBlock.Unmarshal(),
	}
}

func (tb *TimeBounds) Clone() *TimeBounds {
	return &TimeBounds{
		LowerBoundBlock: tb.LowerBoundBlock.Clone(),
		UpperBoundBlock: tb.UpperBoundBlock.Clone(),
	}
}

func (tb *TimeBounds) AsIntArray() [2]*big.Int {
	return [2]*big.Int{tb.LowerBoundBlock.AsInt(), tb.UpperBoundBlock.AsInt()}
}

func (tb *TimeBounds) Equals(other *TimeBounds) bool {
	return tb.LowerBoundBlock.AsInt().Cmp(other.LowerBoundBlock.AsInt()) == 0 &&
		tb.UpperBoundBlock.AsInt().Cmp(other.UpperBoundBlock.AsInt()) == 0
}

func (tb *TimeBounds) IsValidTime(block *common.TimeBlocks) error {
	lowerBoundBlock := tb.LowerBoundBlock.AsInt()
	if block.AsInt().Cmp(lowerBoundBlock) < 0 {
		return errors.New("TimeBounds minimum block must less than the block")
	}
	upperBoundBlock := tb.UpperBoundBlock.AsInt()
	if block.AsInt().Cmp(upperBoundBlock) > 0 {
		return errors.New("TimeBounds maximum block must greater than the block")
	}
	return nil
}

func (tb *TimeBounds) AsValue() value.TupleValue {
	newTup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(tb.LowerBoundBlock.AsInt()),
		value.NewIntValue(tb.UpperBoundBlock.AsInt()),
	})
	return newTup
}
