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
	"bytes"
	"errors"
	"math/big"

	"github.com/golang/protobuf/proto"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type TimeBlocks big.Int

func NewTimeBlocks(val *big.Int) *TimeBlocks {
	return (*TimeBlocks)(val)
}

func (tb *TimeBlocks) AsInt() *big.Int {
	return (*big.Int)(tb)
}

func (tb *TimeBlocks) Marshal() *TimeBlocksBuf {
	return &TimeBlocksBuf{Val: utils.MarshalBigInt(tb.AsInt())}
}

func (tb *TimeBlocksBuf) Unmarshal() *TimeBlocks {
	return (*TimeBlocks)(utils.UnmarshalBigInt(tb.Val))
}

func NewTimeBoundsBlocks(startTimeBlocks, endTimeBlocks *TimeBlocks) *TimeBoundsBlocks {
	return &TimeBoundsBlocks{
		Start: startTimeBlocks.Marshal(),
		End:   endTimeBlocks.Marshal(),
	}
}

func (tb *TimeBoundsBlocks) Clone() *TimeBoundsBlocks {
	return proto.Clone(tb).(*TimeBoundsBlocks)
}

func (tb *TimeBoundsBlocks) AsIntArray() [2]*big.Int {
	return [2]*big.Int{utils.UnmarshalBigInt(tb.Start.Val), utils.UnmarshalBigInt(tb.End.Val)}
}

func (tb *TimeBoundsBlocks) Equals(other *TimeBoundsBlocks) bool {
	return tb.Start.Unmarshal().AsInt().Cmp(other.Start.Unmarshal().AsInt()) == 0 &&
		tb.End.Unmarshal().AsInt().Cmp(other.End.Unmarshal().AsInt()) == 0
}

func (tb *TimeBoundsBlocks) IsValidTime(time *TimeBlocks) error {
	startTime := tb.Start.Unmarshal().AsInt()
	if time.AsInt().Cmp(startTime) < 0 {
		return errors.New("TimeBounds minimum time must less than the time")
	}
	endTime := tb.End.Unmarshal().AsInt()
	if time.AsInt().Cmp(endTime) > 0 {
		return errors.New("TimeBounds maximum time must greater than the time")
	}
	return nil
}

func (tb *TimeBoundsBlocks) AsValue() value.Value {
	newTup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(tb.Start.Unmarshal().AsInt()),
		value.NewIntValue(tb.End.Unmarshal().AsInt()),
	})
	return newTup
}

func NewPrecondition(beforeHash [32]byte, timeBounds *TimeBoundsBlocks, beforeInbox value.Value) *Precondition {
	return &Precondition{BeforeHash: value.NewHashBuf(beforeHash), TimeBounds: timeBounds, BeforeInbox: value.NewHashBuf(beforeInbox.Hash())}
}

func (pre *Precondition) BeforeHashValue() [32]byte {
	var ret [32]byte
	copy(ret[:], pre.BeforeHash.Value)
	return ret
}

func (pre *Precondition) BeforeInboxValue() [32]byte {
	var ret [32]byte
	copy(ret[:], pre.BeforeInbox.Value)
	return ret
}

func (pre *Precondition) Equals(b *Precondition) bool {
	if !bytes.Equal(pre.BeforeHash.Value, b.BeforeHash.Value) {
		return false
	}
	if pre.TimeBounds != b.TimeBounds {
		return false
	}
	if !bytes.Equal(pre.BeforeInbox.Value, b.BeforeInbox.Value) {
		return false
	}
	return true
}

func (pre *Precondition) Hash() [32]byte {
	var ret [32]byte
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(pre.BeforeHash.Value),
		solsha3.Uint128(utils.UnmarshalBigInt(pre.TimeBounds.Start.Val)),
		solsha3.Uint128(utils.UnmarshalBigInt(pre.TimeBounds.End.Val)),
		solsha3.Bytes32(pre.BeforeInbox.Value),
	))
	return ret
}
