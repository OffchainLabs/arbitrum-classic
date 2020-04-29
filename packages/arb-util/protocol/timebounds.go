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
	StartBlock *common.TimeBlocks
	EndBlock   *common.TimeBlocks
	StartTime  *big.Int
	EndTime    *big.Int
}

func (a *TimeBounds) MarshalToBuf() *TimeBoundsBlocksBuf {
	return &TimeBoundsBlocksBuf{
		StartBlock: a.StartBlock.Marshal(),
		EndBlock:   a.EndBlock.Marshal(),
		StartTime:  common.MarshalBigInt(a.StartTime),
		EndTime:    common.MarshalBigInt(a.EndTime),
	}
}

func (a *TimeBoundsBlocksBuf) Unmarshal() *TimeBounds {
	return &TimeBounds{
		StartBlock: a.StartBlock.Unmarshal(),
		EndBlock:   a.EndBlock.Unmarshal(),
		StartTime:  a.StartTime.Unmarshal(),
		EndTime:    a.EndTime.Unmarshal(),
	}
}

func (tb *TimeBounds) Clone() *TimeBounds {
	return &TimeBounds{
		StartBlock: tb.StartBlock.Clone(),
		EndBlock:   tb.EndBlock.Clone(),
		StartTime:  new(big.Int).Set(tb.StartTime),
		EndTime:    new(big.Int).Set(tb.EndTime),
	}
}

func (tb *TimeBounds) AsIntArray() [4]*big.Int {
	return [4]*big.Int{tb.StartBlock.AsInt(), tb.EndBlock.AsInt(), tb.StartTime, tb.EndTime}
}

func (tb *TimeBounds) Equals(other *TimeBounds) bool {
	return tb.StartBlock.AsInt().Cmp(other.StartBlock.AsInt()) == 0 &&
		tb.EndBlock.AsInt().Cmp(other.EndBlock.AsInt()) == 0 &&
		tb.StartTime.Cmp(other.StartTime) == 0 &&
		tb.EndTime.Cmp(other.EndTime) == 0
}

func (tb *TimeBounds) IsValidTime(block *common.TimeBlocks, timestamp *big.Int) error {
	startTime := tb.StartBlock.AsInt()
	if block.AsInt().Cmp(startTime) < 0 {
		return errors.New("TimeBounds minimum block must less than the block")
	}
	endTime := tb.EndBlock.AsInt()
	if block.AsInt().Cmp(endTime) > 0 {
		return errors.New("TimeBounds maximum block must greater than the block")
	}

	if timestamp.Cmp(tb.StartTime) < 0 {
		return errors.New("TimeBounds minimum timestamp must less than the timestamp")
	}
	if timestamp.Cmp(tb.EndTime) > 0 {
		return errors.New("TimeBounds maximum timestamp must greater than the timestamp")
	}

	return nil
}

func (tb *TimeBounds) AsValue() value.TupleValue {
	newTup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(tb.StartBlock.AsInt()),
		value.NewIntValue(tb.EndBlock.AsInt()),
		value.NewIntValue(tb.StartTime),
		value.NewIntValue(tb.EndTime),
	})
	return newTup
}
