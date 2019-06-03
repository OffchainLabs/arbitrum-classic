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

package stack

import (
	"bytes"
	"fmt"
	"github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-avm/value"
)

type Flat struct {
	ints       []value.IntValue
	tuples     []value.TupleValue
	codePoints []value.CodePointValue
	hashOnly   []value.HashOnlyValue
	itemTypes  []byte
	hashes     [][32]byte
	size       int64
}

func (m *Flat) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	s := m.Clone()
	for !s.IsEmpty() {
		val, _ := s.Pop()
		buf.WriteString(fmt.Sprintf("%v", val))
		if !s.IsEmpty() {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("]")
	return buf.String()
}

func NewEmptyFlat() *Flat {
	s := &Flat{nil, nil, nil, nil, nil, nil, 1}
	s.verifyHeight()
	return s
}

func FlatFromTupleChain(val value.Value) *Flat {
	contents := val.(value.TupleValue).Contents()
	if len(contents) == 0 {
		return NewEmptyFlat()
	}
	fs := FlatFromTupleChain(contents[1])
	fs.Push(contents[0])
	fs.verifyHeight()
	return fs
}

func (s *Flat) CloneImpl() *Flat {
	s.verifyHeight()
	ints := make([]value.IntValue, len(s.ints))
	copy(ints, s.ints)
	tuples := make([]value.TupleValue, len(s.tuples))
	copy(tuples, s.tuples)
	codePoints := make([]value.CodePointValue, len(s.codePoints))
	copy(codePoints, s.codePoints)
	hashOnly := make([]value.HashOnlyValue, len(s.hashOnly))
	copy(hashOnly, s.hashOnly)
	itemTypes := make([]byte, len(s.itemTypes))
	copy(itemTypes, s.itemTypes)
	hashes := make([][32]byte, len(s.hashes))
	copy(hashes, s.hashes)
	newS := &Flat{ints, tuples, codePoints, hashOnly, itemTypes, hashes, s.size}
	newS.verifyHeight()
	return newS
}

func (s *Flat) Clone() Stack {
	return s.CloneImpl()
}

func (s *Flat) Equal(yin Stack) (bool, string) {
	y := yin.(*Flat)
	if len(s.itemTypes) != len(y.itemTypes) {
		return false, fmt.Sprintf("Flat stack lengths are different (%v and %v)", len(s.itemTypes), len(y.itemTypes))
	}
	if !bytes.Equal(s.itemTypes, y.itemTypes) {
		return false, "Flat stacks contain values of different types"
	}

	for i := range s.ints {
		if !(value.Eq(s.ints[i], y.ints[i])) {
			return false, fmt.Sprintf("Flat stacks contain ints of different value (expected %v found %v)", s.ints[i], y.ints[i])
		}
	}
	for i := range s.tuples {
		if !(value.Eq(s.tuples[i], y.tuples[i])) {
			return false, "Flat stack tuples different"
		}
	}
	for i := range s.codePoints {
		if !(value.Eq(s.codePoints[i], y.codePoints[i])) {
			return false, "Flat stack codePoints different"
		}
	}
	for i := range s.hashOnly {
		if !(value.Eq(s.hashOnly[i], y.hashOnly[i])) {
			return false, "Flat stack hashOnly different"
		}
	}
	for i := range s.hashes {
		if s.hashes[i] != y.hashes[i] {
			return false, "Flat stack hashes different"
		}
	}

	if s.size != y.size {
		return false, fmt.Sprintf("Flat stack sizes are different (%v and %v)", s.size, y.size)
	}

	return true, ""
}

func (s *Flat) PushInt(v value.IntValue) {
	s.verifyHeight()
	s.ints = append(s.ints, v)
	s.addedValue(v.TypeCode(), v.Size())
	s.verifyHeight()
}

func (s *Flat) PushTuple(v value.TupleValue) {
	s.verifyHeight()
	s.tuples = append(s.tuples, v)
	s.addedValue(v.TypeCode(), v.Size())
	s.verifyHeight()
}

func (s *Flat) PushCodePoint(v value.CodePointValue) {
	s.verifyHeight()
	s.codePoints = append(s.codePoints, v)
	s.addedValue(v.TypeCode(), v.Size())
	s.verifyHeight()
}

func (s *Flat) PushHashOnly(b value.HashOnlyValue) {
	s.verifyHeight()
	s.hashOnly = append(s.hashOnly, b)
	s.addedValue(b.TypeCode(), b.Size())
	s.verifyHeight()
}

func (s *Flat) Push(val value.Value) {
	s.verifyHeight()
	switch v := val.(type) {
	case value.IntValue:
		s.PushInt(v)
	case value.TupleValue:
		s.PushTuple(v)
	case value.CodePointValue:
		s.PushCodePoint(v)
	case value.HashOnlyValue:
		s.PushHashOnly(v)
	default:
		panic("PushValue: unhandled case")
	}
	s.verifyHeight()
}

func (s *Flat) PopInt() (val value.IntValue, err error) {
	s.verifyHeight()
	if err := s.tryPop(value.TypeCodeInt); err != nil {
		return value.IntValue{}, err
	}
	val = s.popIntUnchecked()
	s.verifyHeight()
	return
}

func (s *Flat) PopTuple() (value.TupleValue, error) {
	s.verifyHeight()
	if err := s.tryPop(value.TypeCodeTuple); err != nil {
		return value.TupleValue{}, err
	}
	val := s.popTupleUnchecked()
	s.verifyHeight()
	return val, nil
}

func (s *Flat) PopCodePoint() (val value.CodePointValue, err error) {
	s.verifyHeight()
	if err := s.tryPop(value.TypeCodeCodePoint); err != nil {
		return value.CodePointValue{}, err
	}
	val = s.popCodePointUnchecked()
	s.verifyHeight()
	return
}

func (s *Flat) PopHashOnly() (val value.HashOnlyValue, err error) {
	s.verifyHeight()
	if err := s.tryPop(value.TypeCodeHashOnly); err != nil {
		return value.HashOnlyValue{}, err
	}
	val = s.popHashOnlyUnchecked()
	s.verifyHeight()
	return
}

func (s *Flat) Pop() (lastItem value.Value, err error) {
	s.verifyHeight()
	if len(s.itemTypes) == 0 {
		return nil, EmptyError{}
	}
	valType := s.itemTypes[len(s.itemTypes)-1]
	switch valType {
	case value.TypeCodeInt:
		lastItem = s.popIntUnchecked()
	case value.TypeCodeTuple:
		lastItem = s.popTupleUnchecked()
	case value.TypeCodeCodePoint:
		lastItem = s.popCodePointUnchecked()
	case value.TypeCodeHashOnly:
		lastItem = s.popHashOnlyUnchecked()
	default:
		panic("PopValue: Unhandled type")
	}
	s.verifyHeight()
	return
}

func (s *Flat) IsEmpty() bool {
	return len(s.itemTypes) == 0
}

func (s *Flat) Size() int64 {
	return s.size
}

func (s *Flat) Count() int64 {
	return int64(len(s.itemTypes))
}

func (s *Flat) StateValue() value.Value {
	s.updateHashes()
	if len(s.itemTypes) > 0 {
		return value.NewHashOnlyValue(s.hashes[len(s.hashes)-1], s.size)
	} else {
		return value.NewHashOnlyValue(value.NewEmptyTuple().Hash(), 1)
	}
}

func (s *Flat) ProofValue(stackInfo []byte) value.Value {
	s.updateHashes()
	c := s.CloneImpl()
	vals := make([]value.Value, 0, len(stackInfo))
	for range stackInfo {
		val, _ := c.Pop()
		vals = append(vals, val)
	}
	stack := NewTuple(c.StateValue())
	for i := len(stackInfo) - 1; i >= 0; i-- {
		if stackInfo[i] == 1 {
			stack.Push(vals[i].CloneShallow())
		} else {
			stack.Push(value.NewHashOnlyValueFromValue(vals[i]))
		}
	}
	return stack.stack
}

func (s *Flat) SolidityProofValue(stackInfo []byte) (value.HashOnlyValue, []value.Value) {
	s.updateHashes()
	c := s.CloneImpl()
	vals := make([]value.Value, 0, len(stackInfo))
	for i := range stackInfo {
		val, _ := c.Pop()
		if stackInfo[i] == 1 {
			vals = append(vals, val.CloneShallow())
		} else {
			vals = append(vals, value.NewHashOnlyValueFromValue(val))
		}
	}
	return value.NewHashOnlyValueFromValue(c.StateValue()), vals
}

func (s *Flat) FullyExpandedValue() value.Value {
	s2 := s.Clone()
	return s2.(*Flat).FullyExpandedValueImpl()
}

func (s *Flat) FullyExpandedValueImpl() value.Value {
	if s.IsEmpty() {
		return value.NewEmptyTuple()
	} else {
		top, _ := s.Pop() // ignore error because we just checked for empty stack
		return value.NewTuple2(top, s.FullyExpandedValueImpl())
	}
}

func (s *Flat) addedValueAddHash(itemHash1 [32]byte) {
	var prevItem [32]byte
	if len(s.hashes) > 0 {
		prevItem = s.hashes[len(s.hashes)-1]
	} else {
		prevItem = value.NewEmptyTuple().Hash()
	}
	val := solsha3.SoliditySHA3(
		solsha3.Uint8(value.TypeCodeTuple+2),
		value.Bytes32ArrayEncoded([][32]byte{itemHash1, prevItem}),
	)
	var ret [32]byte
	copy(ret[:], val)
	s.hashes = append(s.hashes, ret)
}

func (s *Flat) countOfType(tipe byte) int {
	switch tipe {
	case value.TypeCodeInt:
		return len(s.ints)
	case value.TypeCodeTuple:
		return len(s.tuples)
	case value.TypeCodeCodePoint:
		return len(s.codePoints)
	case value.TypeCodeHashOnly:
		return len(s.hashOnly)
	default:
		panic("PopValue: Unhandled type")
	}
}

func (s *Flat) hashOfItem(tipe byte, offset int) [32]byte {
	switch tipe {
	case value.TypeCodeInt:
		return s.ints[offset].Hash()
	case value.TypeCodeTuple:
		return s.tuples[offset].Hash()
	case value.TypeCodeCodePoint:
		return s.codePoints[offset].Hash()
	case value.TypeCodeHashOnly:
		return s.hashOnly[offset].Hash()
	default:
		panic("PopValue: Unhandled type")
	}
}

func (s *Flat) addedValue(tipe byte, size int64) {
	s.itemTypes = append(s.itemTypes, tipe)
	s.size += size + 1
	if len(s.itemTypes)-len(s.hashes) > 10 {
		s.updateHashes()
	}
}

func (s *Flat) removedValue(size int64) {
	s.size -= size + 1
	s.itemTypes = s.itemTypes[:len(s.itemTypes)-1]
	if len(s.hashes) > len(s.itemTypes) {
		s.hashes = s.hashes[:len(s.itemTypes)]
	}
}

func (s *Flat) updateHashes() {
	s.verifyHeight()
	var counts [200]int
	for i := len(s.hashes); i < len(s.itemTypes); i++ {
		counts[s.itemTypes[i]]++
	}

	for i := len(s.hashes); i < len(s.itemTypes); i++ {
		tipe := s.itemTypes[i]
		totalCount := s.countOfType(tipe)
		waitingCount := counts[tipe]
		hash := s.hashOfItem(tipe, totalCount-waitingCount)
		s.addedValueAddHash(hash)
		counts[tipe]--
	}
	s.verifyHeight()
}

func (s *Flat) tryPop(tipe byte) error {
	s.verifyHeight()
	if len(s.itemTypes) == 0 {
		return EmptyError{}
	}
	valType := s.itemTypes[len(s.itemTypes)-1]
	if valType != tipe {
		switch valType {
		case value.TypeCodeInt:
			s.popIntUnchecked()
		case value.TypeCodeTuple:
			s.popTupleUnchecked()
		case value.TypeCodeCodePoint:
			s.popCodePointUnchecked()
		case value.TypeCodeHashOnly:
			s.popHashOnlyUnchecked()
		default:
			panic("PopValue: Unhandled type")
		}
		s.verifyHeight()
		return TypeError{}
	}
	s.verifyHeight()
	return nil
}

func (s *Flat) popIntUnchecked() (lastItem value.IntValue) {
	s.verifyHeight()
	lastItem = s.ints[len(s.ints)-1]
	s.ints = s.ints[:len(s.ints)-1]
	s.removedValue(lastItem.Size())
	s.verifyHeight()
	return
}

func (s *Flat) popTupleUnchecked() (lastItem value.TupleValue) {
	s.verifyHeight()
	lastItem = s.tuples[len(s.tuples)-1]
	s.tuples = s.tuples[:len(s.tuples)-1]
	s.removedValue(lastItem.Size())
	s.verifyHeight()
	return
}

func (s *Flat) popCodePointUnchecked() (lastItem value.CodePointValue) {
	s.verifyHeight()
	lastItem = s.codePoints[len(s.codePoints)-1]
	s.codePoints = s.codePoints[:len(s.codePoints)-1]
	s.removedValue(lastItem.Size())
	s.verifyHeight()
	return
}

func (s *Flat) popHashOnlyUnchecked() (lastItem value.HashOnlyValue) {
	s.verifyHeight()
	lastItem = s.hashOnly[len(s.hashOnly)-1]
	s.hashOnly = s.hashOnly[:len(s.hashOnly)-1]
	s.removedValue(lastItem.Size())
	s.verifyHeight()
	return
}

func (s *Flat) verifyHeight() {
	if len(s.ints)+len(s.tuples)+len(s.hashOnly)+len(s.codePoints) != len(s.itemTypes) {
		panic("Bad stack height")
	}
}

func (s *Flat) duplicate() {
	s.verifyHeight()
	tipe := s.itemTypes[len(s.itemTypes)-1]
	switch tipe {
	case value.TypeCodeInt:
		s.PushInt(s.ints[len(s.ints)-1])
	case value.TypeCodeTuple:
		s.PushTuple(s.tuples[len(s.tuples)-1])
	case value.TypeCodeCodePoint:
		s.PushCodePoint(s.codePoints[len(s.codePoints)-1])
	case value.TypeCodeHashOnly:
		s.PushHashOnly(s.hashOnly[len(s.hashOnly)-1])
	default:
		panic("PopValue: Unhandled type")
	}
	s.verifyHeight()
}
