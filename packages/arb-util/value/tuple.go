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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

const MaxTupleSize = 8

var hashOfNone common.Hash

func init() {
	hashOfNone = hashing.SoliditySHA3(hashing.Uint8(TypeCodeTuple))
}

type TupleValue struct {
	contentsArr     [MaxTupleSize]Value
	itemCount       int8
	cachedHash      common.Hash
	size            int64
	deferredHashing bool
}

func NewEmptyTuple() TupleValue {
	return TupleValue{[MaxTupleSize]Value{}, 0, hashOfNone, 1, false}
}

func NewTupleOfSizeWithContents(contents [MaxTupleSize]Value, size int8) (TupleValue, error) {
	if !IsValidTupleSizeI64(int64(size)) {
		return TupleValue{}, errors.New("requested empty tuple size is too big")
	}
	ret := TupleValue{contents, size, common.Hash{}, 0, true}
	ret.size = ret.internalSize()
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
	ret := TupleValue{[MaxTupleSize]Value{value1, value2}, 2, common.Hash{}, 0, true}
	ret.size = ret.internalSize()
	return ret
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

func (tv TupleValue) MarshalForProof(wr io.Writer) error {
	return tv.Marshal(wr)
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
	return TupleValue{newContents, tv.itemCount, tv.cachedHash, tv.size, tv.deferredHashing}
}

func (tv TupleValue) CloneShallow() Value {
	var newContents [MaxTupleSize]Value
	for i, b := range tv.Contents() {
		if b.TypeCode() == TypeCodeInt {
			newContents[i] = b
		} else {
			newContents[i] = NewHashOnlyValueFromValue(b)
		}
	}
	return TupleValue{newContents, tv.itemCount, tv.cachedHash, tv.size, tv.deferredHashing}
}

func (tv TupleValue) Equal(val Value) bool {
	switch val := val.(type) {
	case HashOnlyValue:
		return tv.Hash() == val.Hash()
	case TupleValue:
		return tv.Hash() == val.Hash()
	default:
		return false
	}
}

func (tv TupleValue) internalSize() int64 {
	ret := int64(1)
	for _, bv := range tv.Contents() {
		ret += bv.Size()
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

func (tv TupleValue) internalHash() common.Hash {
	hashes := make([]common.Hash, 0, tv.itemCount)
	for _, v := range tv.Contents() {
		hashes = append(hashes, v.Hash())
	}

	return hashing.SoliditySHA3(
		hashing.Uint8(tv.InternalTypeCode()),
		hashing.Bytes32ArrayEncoded(hashes),
	)
}

func (tv TupleValue) Hash() common.Hash {
	if tv.deferredHashing {
		tv.cachedHash = tv.internalHash()
		tv.deferredHashing = false
	}
	return tv.cachedHash
}
