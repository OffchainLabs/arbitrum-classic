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

package value

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

const MaxTupleSize = 8

var hashOfNone [32]byte

func init() {
	hashOfNoneVal := solsha3.SoliditySHA3(solsha3.Uint8(TypeCodeTuple))
	copy(hashOfNone[:], hashOfNoneVal)
}

type TupleValue struct {
	contentsArr [MaxTupleSize]Value
	itemCount   int8
	cachedHash  [32]byte
	size        int64
}

func NewEmptyTuple() TupleValue {
	return TupleValue{[MaxTupleSize]Value{}, 0, hashOfNone, 1}
}

func NewTupleOfSizeWithContents(contents [MaxTupleSize]Value, size int8) (TupleValue, error) {
	if !IsValidTupleSizeI64(int64(size)) {
		return TupleValue{}, errors.New("requested empty tuple size is too big")
	}
	ret := TupleValue{contents, size, [32]byte{}, 0}
	ret.size = ret.internalSize()
	ret.cachedHash = ret.internalHash()
	return ret, nil
}

func NewRepeatedTuple(value Value, size int64) (TupleValue, error) {
	if !IsValidTupleSize(big.NewInt(size)) {
		return TupleValue{}, errors.New("requested tuple size is too big")
	}
	ret := TupleValue{[MaxTupleSize]Value{}, int8(size), [32]byte{}, 0}
	for i := int64(0); i < size; i++ {
		ret.contentsArr[i] = value
	}
	ret.size = ret.internalSize()
	ret.cachedHash = ret.internalHash()
	return ret, nil
}

func NewTupleFromSlice(slice []Value) (TupleValue, error) {
	if !IsValidTupleSizeI64(int64(len(slice))) {
		return TupleValue{}, errors.New("requested tuple size is too big")
	}
	var contents [MaxTupleSize]Value
	for i, v := range slice {
		contents[i] = v
	}
	return NewTupleOfSizeWithContents(contents, int8(len(slice)))
}

func NewTuple2(value1 Value, value2 Value) TupleValue {
	ret := TupleValue{[MaxTupleSize]Value{value1, value2}, 2, [32]byte{}, 0}
	ret.size = ret.internalSize()
	ret.cachedHash = ret.internalHash()
	return ret
}

func (tv TupleValue) init2(value1 Value, value2 Value) {
	tv.contentsArr[0] = value1
	tv.contentsArr[1] = value2
	tv.itemCount = 2
	tv.size = tv.internalSize()
	tv.cachedHash = tv.internalHash()
}

func NewSizedTupleFromReader(rd io.Reader, size byte) (TupleValue, error) {
	var contentsArr [MaxTupleSize]Value
	sz := int8(size)
	for i := 0; i < int(sz); i++ {
		boxedVal, err := UnmarshalValue(rd)
		if err != nil {
			return TupleValue{}, err
		}
		contentsArr[i] = boxedVal
	}
	return NewTupleOfSizeWithContents(contentsArr, sz)
}

func (tv TupleValue) Marshal(wr io.Writer) error {
	for _, v := range tv.Contents() {
		if err := MarshalValue(v, wr); err != nil {
			return err
		}
	}
	return nil
}

func IsValidTupleSizeI64(size int64) bool {
	return size >= 0 && size <= MaxTupleSize
}

func IsValidTupleSize(size *big.Int) bool {
	return size.Cmp(big.NewInt(0)) >= 0 && size.Cmp(big.NewInt(MaxTupleSize)) <= 0
}

func (tv TupleValue) Contents() []Value {
	return tv.contentsArr[:tv.itemCount]
}

func (tv TupleValue) Len() int64 {
	return int64(tv.itemCount)
}

func (tv TupleValue) IsValidIndex(idx IntValue) bool {
	return idx.val.Cmp(big.NewInt(0)) >= 0 && idx.val.Cmp(big.NewInt(tv.Len())) < 0
}

func (tv TupleValue) Get(idx IntValue) (Value, error) {
	return tv.GetByInt64(idx.val.Int64())
}

func (tv TupleValue) GetByInt64(idx int64) (Value, error) {
	if idx < 0 || idx >= tv.Len() {
		return nil, errors.New("tuple index out of bounds")
	}
	return tv.contentsArr[idx], nil
}

func (tv TupleValue) Set(idx IntValue, val Value) (TupleValue, error) {
	return tv.SetByInt64(idx.val.Int64(), val)
}

func (tv TupleValue) SetByInt64(idx int64, val Value) (TupleValue, error) {
	if idx < 0 || idx >= tv.Len() {
		return TupleValue{}, errors.New("tuple index out of bounds")
	}
	var contents [MaxTupleSize]Value
	for i, v := range tv.Contents() {
		contents[i] = v
	}
	contents[idx] = val
	return NewTupleOfSizeWithContents(contents, tv.itemCount)
}

func (tv TupleValue) TypeCode() uint8 {
	return TypeCodeTuple
}

func (tv TupleValue) InternalTypeCode() uint8 {
	return TypeCodeTuple + byte(tv.itemCount)
}

func (tv TupleValue) Clone() Value {
	var newContents [MaxTupleSize]Value
	for i, b := range tv.Contents() {
		newContents[i] = b.Clone()
	}
	return TupleValue{newContents, tv.itemCount, tv.cachedHash, tv.size}
}

func (tv TupleValue) CloneShallow() Value {
	var newContents [MaxTupleSize]Value
	for i, b := range tv.Contents() {
		newContents[i] = NewHashOnlyValueFromValue(b)
	}
	return TupleValue{newContents, tv.itemCount, tv.cachedHash, tv.size}
}

func (tv TupleValue) Equal(val Value) bool {
	if val.TypeCode() == TypeCodeHashOnly {
		return tv.Hash() == val.Hash()
	} else if val.TypeCode() != TypeCodeTuple {
		return false
	} else {
		return tv.cachedHash == val.(TupleValue).cachedHash
	}
}

func (tv TupleValue) internalSize() int64 {
	ret := int64(1)
	for _, bv := range tv.Contents() {
		ret = ret + bv.Size()
	}
	return ret
}

func (tv TupleValue) Size() int64 {
	return tv.size
}

func (tv TupleValue) String() string {
	var buf bytes.Buffer
	buf.WriteString("Tuple(")
	for i, v := range tv.Contents() {
		buf.WriteString(fmt.Sprintf("%v", v))
		if int64(i) != tv.Len()-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString(")")
	return buf.String()
}

func Bytes32ArrayEncoded(input [][32]byte) []byte {
	var values []byte
	for _, val := range input {
		values = append(values, common.RightPadBytes(val[:], 32)...)
	}
	return values
}

func (tv TupleValue) internalHash() [32]byte {
	hashes := make([][32]byte, 0, tv.itemCount)
	for _, v := range tv.Contents() {
		hashes = append(hashes, v.Hash())
	}

	hashVal := solsha3.SoliditySHA3(
		solsha3.Uint8(tv.InternalTypeCode()),
		Bytes32ArrayEncoded(hashes),
	)
	ret := [32]byte{}
	copy(ret[:], hashVal)
	return ret
}

func (tv TupleValue) Hash() [32]byte {
	return tv.cachedHash
}
