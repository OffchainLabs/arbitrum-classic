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

package protocol

import (
	"errors"
	"math/big"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func NewTimeBounds(startTime, endTime uint64) *TimeBounds {
	return &TimeBounds{StartTime: startTime, EndTime: endTime}
}

func (tb *TimeBounds) Equals(other TimeBounds) bool {
	return tb.StartTime == other.StartTime && tb.EndTime == other.EndTime
}

func (tb *TimeBounds) IsValidTime(time uint64) error {
	if time < tb.StartTime {
		return errors.New("TimeBounds minimum time must less than the time")
	}
	if time > tb.EndTime {
		return errors.New("TimeBounds maximum time must greater than the time")
	}
	return nil
}

func (tb *TimeBounds) AsValue() value.Value {
	newTup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(new(big.Int).SetUint64(tb.StartTime)),
		value.NewIntValue(new(big.Int).SetUint64(tb.EndTime)),
	})
	return newTup
}

type Precondition struct {
	BeforeHash  [32]byte
	TimeBounds  *TimeBounds
	BeforeInbox value.HashOnlyValue
}

func NewPrecondition(beforeHash [32]byte, timeBounds *TimeBounds, beforeInbox value.Value) *Precondition {
	return &Precondition{beforeHash, timeBounds, value.NewHashOnlyValueFromValue(beforeInbox)}
}

func (pre *Precondition) Clone() *Precondition {
	return NewPrecondition(pre.BeforeHash, pre.TimeBounds, pre.BeforeInbox.Clone())
}

func (pre *Precondition) Equals(b *Precondition) bool {
	if pre.BeforeHash != b.BeforeHash {
		return false
	}
	if pre.TimeBounds != b.TimeBounds {
		return false
	}
	if !value.Eq(pre.BeforeInbox, b.BeforeInbox) {
		return false
	}
	return true
}

func (pre *Precondition) Hash() [32]byte {
	var ret [32]byte
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(pre.BeforeHash),
		solsha3.Uint64(pre.TimeBounds.StartTime),
		solsha3.Uint64(pre.TimeBounds.EndTime),
		solsha3.Bytes32(pre.BeforeInbox.Hash()),
	))
	return ret
}
