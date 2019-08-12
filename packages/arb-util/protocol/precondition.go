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
	"encoding/binary"
	"errors"
	"io"
	"math/big"

	"github.com/golang/protobuf/proto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type TimeBounds [2]uint64

func NewTimeBounds(startTime, endTime uint64) TimeBounds {
	return TimeBounds{startTime, endTime}
}

func NewTimeBoundsFromReader(rd io.Reader) (TimeBounds, error) {
	var tb TimeBounds
	if err := binary.Read(rd, binary.LittleEndian, &tb[0]); err != nil {
		return tb, err
	}
	if err := binary.Read(rd, binary.LittleEndian, &tb[1]); err != nil {
		return tb, err
	}
	return tb, nil
}

func (tb TimeBounds) Equals(other TimeBounds) bool {
	return tb == other
}

func (tb TimeBounds) IsValidTime(time uint64) error {
	if time < tb[0] {
		return errors.New("TimeBounds minimum time must less than the time")
	}
	if time > tb[1] {
		return errors.New("TimeBounds maximum time must greater than the time")
	}
	return nil
}

func (tb TimeBounds) Marshal(wr io.Writer) error {
	if err := binary.Write(wr, binary.LittleEndian, &tb[0]); err != nil {
		return err
	}
	if err := binary.Write(wr, binary.LittleEndian, &tb[1]); err != nil {
		return err
	}
	return nil
}

func (tb TimeBounds) AsValue() value.Value {
	newTup, _ := value.NewTupleFromSlice([]value.Value{
		value.NewIntValue(new(big.Int).SetUint64(tb[0])),
		value.NewIntValue(new(big.Int).SetUint64(tb[1])),
	})
	return newTup
}

type Precondition struct {
	BeforeHash    [32]byte
	TimeBounds    TimeBounds
	BeforeBalance *BalanceTracker
	BeforeInbox   value.HashOnlyValue
}

func NewPrecondition(beforeHash [32]byte, timeBounds TimeBounds, beforeBalance *BalanceTracker, beforeInbox value.Value) *Precondition {
	return &Precondition{beforeHash, timeBounds, beforeBalance, value.NewHashOnlyValueFromValue(beforeInbox)}
}

func (pre *Precondition) Clone() *Precondition {
	return NewPrecondition(pre.BeforeHash, pre.TimeBounds, pre.BeforeBalance.Clone(), pre.BeforeInbox.Clone())
}

func NewPreconditionFromReader(rd io.Reader) (*Precondition, error) {
	length := uint64(0)
	err := binary.Read(rd, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 0, length)
	_, err = io.ReadFull(rd, buf)
	if err != nil {
		return nil, err
	}
	pre := &PreconditionBuf{}
	err = proto.Unmarshal(buf, pre)
	if err != nil {
		return nil, err
	}
	return NewPreconditionFromBuf(pre), nil
}

func (pre *Precondition) Marshal(wr io.Writer) error {
	preData, err := proto.Marshal(NewPreconditionBuf(pre))
	if err != nil {
		return err
	}
	length := uint64(len(preData))
	err = binary.Write(wr, binary.LittleEndian, &length)
	if err != nil {
		return err
	}
	_, err = wr.Write(preData)
	return err
}

func (pre *Precondition) Equals(b *Precondition) bool {
	if pre.BeforeHash != b.BeforeHash {
		return false
	}
	if pre.TimeBounds != b.TimeBounds {
		return false
	}
	if !pre.BeforeBalance.Equals(b.BeforeBalance) {
		return false
	}
	if !value.Eq(pre.BeforeInbox, b.BeforeInbox) {
		return false
	}
	return true
}

func (pre *Precondition) Hash() [32]byte {
	tokenTypes := make([][21]byte, 0, len(pre.BeforeBalance.TokenTypes))
	for _, tokType := range pre.BeforeBalance.TokenTypes {
		tokenTypes = append(tokenTypes, tokType)
	}
	var ret [32]byte
	hashVal := solsha3.SoliditySHA3(
		solsha3.Bytes32(pre.BeforeHash),
		solsha3.Uint64(pre.TimeBounds[0]),
		solsha3.Uint64(pre.TimeBounds[1]),
		solsha3.Bytes32(pre.BeforeInbox.Hash()),
		TokenTypeArrayEncoded(tokenTypes),
		solsha3.Uint256Array(pre.BeforeBalance.TokenAmounts),
	)
	copy(ret[:], hashVal)
	return ret
}
