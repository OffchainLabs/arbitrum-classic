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

package value

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
)

const MaxTupleSize = 8

type TupleValue struct {
	contentsArr [MaxTupleSize]Value
	itemCount   int8
	size        int64
}

func NewEmptyTuple() *TupleValue {
	return &TupleValue{[MaxTupleSize]Value{}, 0, 1}
}

func NewTupleOfSizeWithContents(contents [MaxTupleSize]Value, size int8) (*TupleValue, error) {
	if !IsValidTupleSizeI64(int64(size)) {
		return nil, errors.New("requested empty tuple size is too big")
	}
	ret := &TupleValue{contents, size, 0}
	ret.size = ret.internalSize()
	return ret, nil
}

func NewTupleFromSlice(slice []Value) (*TupleValue, error) {
	if !IsValidTupleSizeI64(int64(len(slice))) {
		return nil, errors.New("requested tuple size is too big")
	}
	var contents [MaxTupleSize]Value
	for i, v := range slice {
		contents[i] = v
	}
	return NewTupleOfSizeWithContents(contents, int8(len(slice)))
}

func NewTuple2(value1 Value, value2 Value) *TupleValue {
	ret := &TupleValue{[MaxTupleSize]Value{value1, value2}, 2, 0}
	ret.size = ret.internalSize()
	return ret
}

func NewSizedTupleFromReader(rd io.Reader, size byte) (*TupleValue, error) {
	var contentsArr [MaxTupleSize]Value
	sz := int8(size)
	for i := 0; i < int(sz); i++ {
		boxedVal, err := UnmarshalValue(rd)
		if err != nil {
			return nil, err
		}
		contentsArr[i] = boxedVal
	}
	return NewTupleOfSizeWithContents(contentsArr, sz)
}

func IsValidTupleSizeI64(size int64) bool {
	return size >= 0 && size <= MaxTupleSize
}

func (tv *TupleValue) Contents() []Value {
	return tv.contentsArr[:tv.itemCount]
}

func (tv *TupleValue) Len() int64 {
	return int64(tv.itemCount)
}

func (tv *TupleValue) GetByInt64(idx int64) (Value, error) {
	if idx < 0 || idx >= tv.Len() {
		return nil, errors.New("tuple index out of bounds")
	}
	return tv.contentsArr[idx], nil
}

func (tv *TupleValue) TypeCode() uint8 {
	return TypeCodeTuple + byte(tv.itemCount)
}

func (tv *TupleValue) Equal(val Value) bool {
	tup, ok := val.(*TupleValue)
	if !ok {
		return false
	}
	if tup.Len() != tv.Len() {
		return false
	}
	for i, val := range tv.Contents() {
		if !Eq(val, tup.contentsArr[i]) {
			return false
		}
	}
	return true
}

func (tv *TupleValue) internalSize() int64 {
	ret := int64(1)
	for _, bv := range tv.Contents() {
		ret = ret + bv.Size()
	}
	return ret
}

func (tv *TupleValue) Size() int64 {
	return tv.size
}

func (tv *TupleValue) String() string {
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
