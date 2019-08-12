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
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

const BytesPerInt = 32

var hashOfOne [32]byte
var hashOfZero [32]byte
var IntegerZero IntValue

func init() {
	hashOfOne = NewInt64Value(1).hashImpl()
	hashOfZero = NewInt64Value(0).hashImpl()
	IntegerZero = NewInt64Value(0)
}

type IntValue struct {
	val *big.Int
}

func NewIntValue(x *big.Int) IntValue {
	return IntValue{x}
}

func NewInt64Value(x int64) IntValue {
	return IntValue{big.NewInt(x)}
}

func NewBooleanValue(val bool) IntValue {
	if val {
		return NewInt64Value(1)
	}
	return NewInt64Value(0)
}

func NewIntValueFromString(str string) Value {
	val, valid := math.ParseBig256(str)
	if !valid {
		return NewEmptyTuple()
	}
	return NewIntValue(val)
}

func NewIntValueFromReader(rd io.Reader) (IntValue, error) {
	var data [32]byte
	_, err := rd.Read(data[:])
	if err != nil {
		return IntValue{}, err
	}
	ret := new(big.Int).SetBytes(data[:])
	return NewIntValue(ret), err
}

func (iv IntValue) TypeCode() uint8 {
	return TypeCodeInt
}

func (iv IntValue) InternalTypeCode() uint8 {
	return TypeCodeInt
}

func (iv IntValue) Clone() Value {
	return IntValue{new(big.Int).Set(iv.val)}
}

func (iv IntValue) CloneShallow() Value {
	return IntValue{iv.val}
}

func (iv IntValue) Equal(val Value) bool {
	if val.TypeCode() == TypeCodeHashOnly {
		return iv.Hash() == val.Hash()
	} else if val.TypeCode() != TypeCodeInt {
		return false
	} else {
		return iv.val.Cmp(val.(IntValue).val) == 0
	}
}

func (iv IntValue) Size() int64 {
	return 1
}

func (iv IntValue) BigInt() *big.Int {
	return new(big.Int).Set(iv.val)
}

func (iv IntValue) String() string {
	return iv.val.String()
}

func (iv IntValue) hashImpl() [32]byte {
	hashVal := solsha3.SoliditySHA3(
		solsha3.Uint256(iv.BigInt()),
	)
	ret := [32]byte{}
	copy(ret[:], hashVal)
	return ret
}

func (iv IntValue) ToBytes() [32]byte {
	var data [32]byte
	copy(data[:], math.PaddedBigBytes(math.U256(new(big.Int).Set(iv.val)), 32))
	return data
}

func (iv IntValue) Hash() [32]byte {
	if iv.val.Cmp(big.NewInt(0)) == 0 {
		return hashOfZero
	} else if iv.val.Cmp(big.NewInt(1)) == 0 {
		return hashOfOne
	} else {
		return iv.hashImpl()
	}
}

func (iv IntValue) Marshal(w io.Writer) error {
	_, err := w.Write(math.PaddedBigBytes(math.U256(new(big.Int).Set(iv.val)), 32))
	return err
}
