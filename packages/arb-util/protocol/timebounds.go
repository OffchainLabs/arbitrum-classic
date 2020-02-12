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

type TimeBoundsBlocks struct {
	Start *common.TimeBlocks
	End   *common.TimeBlocks
}

func (a *TimeBoundsBlocks) MarshalToBuf() *TimeBoundsBlocksBuf {
	return &TimeBoundsBlocksBuf{
		Start: a.Start.Marshal(),
		End:   a.End.Marshal(),
	}
}

func (a *TimeBoundsBlocksBuf) Unmarshal() *TimeBoundsBlocks {
	return &TimeBoundsBlocks{
		Start: a.Start.Unmarshal(),
		End:   a.End.Unmarshal(),
	}
}

func (tb *TimeBoundsBlocks) Clone() *TimeBoundsBlocks {
	return &TimeBoundsBlocks{
		Start: tb.Start.Clone(),
		End:   tb.End.Clone(),
	}
}

func (tb *TimeBoundsBlocks) AsIntArray() [2]*big.Int {
	return [2]*big.Int{tb.Start.AsInt(), tb.End.AsInt()}
}

func (tb *TimeBoundsBlocks) Equals(other *TimeBoundsBlocks) bool {
	return tb.Start.AsInt().Cmp(other.Start.AsInt()) == 0 &&
		tb.End.AsInt().Cmp(other.End.AsInt()) == 0
}

func (tb *TimeBoundsBlocks) IsValidTime(time *common.TimeBlocks) error {
	startTime := tb.Start.AsInt()
	if time.AsInt().Cmp(startTime) < 0 {
		return errors.New("TimeBounds minimum time must less than the time")
	}
	endTime := tb.End.AsInt()
	if time.AsInt().Cmp(endTime) > 0 {
		return errors.New("TimeBounds maximum time must greater than the time")
	}
	return nil
}

func (tb *TimeBoundsBlocks) AsValue() value.TupleValue {
	newTup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(tb.Start.AsInt()),
		value.NewIntValue(tb.End.AsInt()),
	})
	return newTup
}
