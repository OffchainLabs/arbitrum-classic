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

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
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

func (f *Flat) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	s := f.Clone()
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

func (f *Flat) CloneImpl() *Flat {
	f.verifyHeight()
	ints := make([]value.IntValue, len(f.ints))
	copy(ints, f.ints)
	tuples := make([]value.TupleValue, len(f.tuples))
	copy(tuples, f.tuples)
	codePoints := make([]value.CodePointValue, len(f.codePoints))
	copy(codePoints, f.codePoints)
	hashOnly := make([]value.HashOnlyValue, len(f.hashOnly))
	copy(hashOnly, f.hashOnly)
	itemTypes := make([]byte, len(f.itemTypes))
	copy(itemTypes, f.itemTypes)
	hashes := make([][32]byte, len(f.hashes))
	copy(hashes, f.hashes)
	newS := &Flat{ints, tuples, codePoints, hashOnly, itemTypes, hashes, f.size}
	newS.verifyHeight()
	return newS
}

func (f *Flat) Clone() Stack {
	return f.CloneImpl()
}

func (f *Flat) Equal(yin Stack) (bool, string) {
	y := yin.(*Flat)
	if len(f.itemTypes) != len(y.itemTypes) {
		return false, fmt.Sprintf("Flat stack lengths are different (%v and %v)", len(f.itemTypes), len(y.itemTypes))
	}
	if !bytes.Equal(f.itemTypes, y.itemTypes) {
		return false, "Flat stacks contain values of different types"
	}

	for i := range f.ints {
		if !(value.Eq(f.ints[i], y.ints[i])) {
			return false, fmt.Sprintf("Flat stacks contain ints of different value (expected %v found %v)", f.ints[i], y.ints[i])
		}
	}
	for i := range f.tuples {
		if !(value.Eq(f.tuples[i], y.tuples[i])) {
			return false, "Flat stack tuples different"
		}
	}
	for i := range f.codePoints {
		if !(value.Eq(f.codePoints[i], y.codePoints[i])) {
			return false, "Flat stack codePoints different"
		}
	}
	for i := range f.hashOnly {
		if !(value.Eq(f.hashOnly[i], y.hashOnly[i])) {
			return false, "Flat stack hashOnly different"
		}
	}
	for i := range f.hashes {
		if f.hashes[i] != y.hashes[i] {
			return false, "Flat stack hashes different"
		}
	}

	if f.size != y.size {
		return false, fmt.Sprintf("Flat stack sizes are different (%v and %v)", f.size, y.size)
	}

	return true, ""
}

func (f *Flat) PushInt(v value.IntValue) {
	f.verifyHeight()
	f.ints = append(f.ints, v)
	f.addedValue(v.TypeCode(), v.Size())
	f.verifyHeight()
}

func (f *Flat) PushTuple(v value.TupleValue) {
	f.verifyHeight()
	f.tuples = append(f.tuples, v)
	f.addedValue(v.TypeCode(), v.Size())
	f.verifyHeight()
}

func (f *Flat) PushCodePoint(v value.CodePointValue) {
	f.verifyHeight()
	f.codePoints = append(f.codePoints, v)
	f.addedValue(v.TypeCode(), v.Size())
	f.verifyHeight()
}

func (f *Flat) PushHashOnly(b value.HashOnlyValue) {
	f.verifyHeight()
	f.hashOnly = append(f.hashOnly, b)
	f.addedValue(b.TypeCode(), b.Size())
	f.verifyHeight()
}

func (f *Flat) Push(val value.Value) {
	f.verifyHeight()
	switch v := val.(type) {
	case value.IntValue:
		f.PushInt(v)
	case value.TupleValue:
		f.PushTuple(v)
	case value.CodePointValue:
		f.PushCodePoint(v)
	case value.HashOnlyValue:
		f.PushHashOnly(v)
	default:
		panic("PushValue: unhandled case")
	}
	f.verifyHeight()
}

func (f *Flat) PopInt() (val value.IntValue, err error) {
	f.verifyHeight()
	if err := f.tryPop(value.TypeCodeInt); err != nil {
		return value.IntValue{}, err
	}
	val = f.popIntUnchecked()
	f.verifyHeight()
	return
}

func (f *Flat) PopTuple() (value.TupleValue, error) {
	f.verifyHeight()
	if err := f.tryPop(value.TypeCodeTuple); err != nil {
		return value.TupleValue{}, err
	}
	val := f.popTupleUnchecked()
	f.verifyHeight()
	return val, nil
}

func (f *Flat) PopCodePoint() (val value.CodePointValue, err error) {
	f.verifyHeight()
	if err := f.tryPop(value.TypeCodeCodePoint); err != nil {
		return value.CodePointValue{}, err
	}
	val = f.popCodePointUnchecked()
	f.verifyHeight()
	return
}

func (f *Flat) PopHashOnly() (val value.HashOnlyValue, err error) {
	f.verifyHeight()
	if err := f.tryPop(value.TypeCodeHashOnly); err != nil {
		return value.HashOnlyValue{}, err
	}
	val = f.popHashOnlyUnchecked()
	f.verifyHeight()
	return
}

func (f *Flat) Pop() (value.Value, error) {
	f.verifyHeight()
	if len(f.itemTypes) == 0 {
		return nil, EmptyError{}
	}
	valType := f.itemTypes[len(f.itemTypes)-1]
	switch valType {
	case value.TypeCodeInt:
		return f.popIntUnchecked(), nil
	case value.TypeCodeTuple:
		return f.popTupleUnchecked(), nil
	case value.TypeCodeCodePoint:
		return f.popCodePointUnchecked(), nil
	case value.TypeCodeHashOnly:
		return f.popHashOnlyUnchecked(), nil
	default:
		panic("PopValue: Unhandled type")
	}
}

func (f *Flat) IsEmpty() bool {
	return len(f.itemTypes) == 0
}

func (f *Flat) Size() int64 {
	return f.size
}

func (f *Flat) Count() int64 {
	return int64(len(f.itemTypes))
}

func (f *Flat) StateValue() value.Value {
	f.updateHashes()
	if len(f.itemTypes) == 0 {
		return value.NewHashOnlyValue(value.NewEmptyTuple().Hash(), 1)
	}
	return value.NewHashOnlyValue(f.hashes[len(f.hashes)-1], f.size)
}

func (f *Flat) ProofValue(stackInfo []byte) value.Value {
	f.updateHashes()
	c := f.CloneImpl()
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

func (f *Flat) SolidityProofValue(stackInfo []byte) (value.HashOnlyValue, []value.Value) {
	f.updateHashes()
	c := f.CloneImpl()
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

func (f *Flat) FullyExpandedValue() value.Value {
	s2 := f.Clone()
	return s2.(*Flat).FullyExpandedValueImpl()
}

func (f *Flat) FullyExpandedValueImpl() value.Value {
	if f.IsEmpty() {
		return value.NewEmptyTuple()
	}
	top, _ := f.Pop() // ignore error because we just checked for empty stack
	return value.NewTuple2(top, f.FullyExpandedValueImpl())
}

func (f *Flat) addedValueAddHash(itemHash1 [32]byte) {
	var prevItem [32]byte
	if len(f.hashes) > 0 {
		prevItem = f.hashes[len(f.hashes)-1]
	} else {
		prevItem = value.NewEmptyTuple().Hash()
	}
	val := solsha3.SoliditySHA3(
		solsha3.Uint8(value.TypeCodeTuple+2),
		value.Bytes32ArrayEncoded([][32]byte{itemHash1, prevItem}),
	)
	var ret [32]byte
	copy(ret[:], val)
	f.hashes = append(f.hashes, ret)
}

func (f *Flat) countOfType(tipe byte) int {
	switch tipe {
	case value.TypeCodeInt:
		return len(f.ints)
	case value.TypeCodeTuple:
		return len(f.tuples)
	case value.TypeCodeCodePoint:
		return len(f.codePoints)
	case value.TypeCodeHashOnly:
		return len(f.hashOnly)
	default:
		panic("PopValue: Unhandled type")
	}
}

func (f *Flat) hashOfItem(tipe byte, offset int) [32]byte {
	switch tipe {
	case value.TypeCodeInt:
		return f.ints[offset].Hash()
	case value.TypeCodeTuple:
		return f.tuples[offset].Hash()
	case value.TypeCodeCodePoint:
		return f.codePoints[offset].Hash()
	case value.TypeCodeHashOnly:
		return f.hashOnly[offset].Hash()
	default:
		panic("PopValue: Unhandled type")
	}
}

func (f *Flat) addedValue(tipe byte, size int64) {
	f.itemTypes = append(f.itemTypes, tipe)
	f.size += size + 1
	if len(f.itemTypes)-len(f.hashes) > 10 {
		f.updateHashes()
	}
}

func (f *Flat) removedValue(size int64) {
	f.size -= size + 1
	f.itemTypes = f.itemTypes[:len(f.itemTypes)-1]
	if len(f.hashes) > len(f.itemTypes) {
		f.hashes = f.hashes[:len(f.itemTypes)]
	}
}

func (f *Flat) updateHashes() {
	f.verifyHeight()
	var counts [200]int
	for i := len(f.hashes); i < len(f.itemTypes); i++ {
		counts[f.itemTypes[i]]++
	}

	for i := len(f.hashes); i < len(f.itemTypes); i++ {
		tipe := f.itemTypes[i]
		totalCount := f.countOfType(tipe)
		waitingCount := counts[tipe]
		hash := f.hashOfItem(tipe, totalCount-waitingCount)
		f.addedValueAddHash(hash)
		counts[tipe]--
	}
	f.verifyHeight()
}

func (f *Flat) tryPop(tipe byte) error {
	f.verifyHeight()
	if len(f.itemTypes) == 0 {
		return EmptyError{}
	}
	valType := f.itemTypes[len(f.itemTypes)-1]
	if valType != tipe {
		switch valType {
		case value.TypeCodeInt:
			f.popIntUnchecked()
		case value.TypeCodeTuple:
			f.popTupleUnchecked()
		case value.TypeCodeCodePoint:
			f.popCodePointUnchecked()
		case value.TypeCodeHashOnly:
			f.popHashOnlyUnchecked()
		default:
			panic("PopValue: Unhandled type")
		}
		f.verifyHeight()
		return TypeError{}
	}
	f.verifyHeight()
	return nil
}

func (f *Flat) popIntUnchecked() (lastItem value.IntValue) {
	f.verifyHeight()
	lastItem = f.ints[len(f.ints)-1]
	f.ints = f.ints[:len(f.ints)-1]
	f.removedValue(lastItem.Size())
	f.verifyHeight()
	return
}

func (f *Flat) popTupleUnchecked() (lastItem value.TupleValue) {
	f.verifyHeight()
	lastItem = f.tuples[len(f.tuples)-1]
	f.tuples = f.tuples[:len(f.tuples)-1]
	f.removedValue(lastItem.Size())
	f.verifyHeight()
	return
}

func (f *Flat) popCodePointUnchecked() (lastItem value.CodePointValue) {
	f.verifyHeight()
	lastItem = f.codePoints[len(f.codePoints)-1]
	f.codePoints = f.codePoints[:len(f.codePoints)-1]
	f.removedValue(lastItem.Size())
	f.verifyHeight()
	return
}

func (f *Flat) popHashOnlyUnchecked() (lastItem value.HashOnlyValue) {
	f.verifyHeight()
	lastItem = f.hashOnly[len(f.hashOnly)-1]
	f.hashOnly = f.hashOnly[:len(f.hashOnly)-1]
	f.removedValue(lastItem.Size())
	f.verifyHeight()
	return
}

func (f *Flat) verifyHeight() {
	if len(f.ints)+len(f.tuples)+len(f.hashOnly)+len(f.codePoints) != len(f.itemTypes) {
		panic("Bad stack height")
	}
}

func (f *Flat) duplicate() {
	f.verifyHeight()
	tipe := f.itemTypes[len(f.itemTypes)-1]
	switch tipe {
	case value.TypeCodeInt:
		f.PushInt(f.ints[len(f.ints)-1])
	case value.TypeCodeTuple:
		f.PushTuple(f.tuples[len(f.tuples)-1])
	case value.TypeCodeCodePoint:
		f.PushCodePoint(f.codePoints[len(f.codePoints)-1])
	case value.TypeCodeHashOnly:
		f.PushHashOnly(f.hashOnly[len(f.hashOnly)-1])
	default:
		panic("PopValue: Unhandled type")
	}
	f.verifyHeight()
}
