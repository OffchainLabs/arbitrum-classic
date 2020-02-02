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

	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

const BytesPerInt = 32

var hashOfOne common.Hash
var hashOfZero common.Hash
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

func NewIntValueFromReader(rd io.Reader) (IntValue, error) {
	var data common.Hash
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
	switch val := val.(type) {
	case HashOnlyValue:
		return iv.Hash() == val.hash
	case IntValue:
		return iv.val.Cmp(val.val) == 0
	default:
		return false
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

func (iv IntValue) hashImpl() common.Hash {
	return hashing.SoliditySHA3(
		hashing.Uint256(iv.BigInt()),
	)
}

func (iv IntValue) ToBytes() [32]byte {
	var data [32]byte
	copy(data[:], math.PaddedBigBytes(math.U256(new(big.Int).Set(iv.val)), 32))
	return data
}

func (iv IntValue) Hash() common.Hash {
	switch {
	case iv.val.Cmp(big.NewInt(0)) == 0:
		return hashOfZero
	case iv.val.Cmp(big.NewInt(1)) == 0:
		return hashOfOne
	default:
		return iv.hashImpl()
	}
}

func (iv IntValue) Marshal(w io.Writer) error {
	bytesVal := iv.ToBytes()
	_, err := w.Write(bytesVal[:])
	return err
}

func (iv IntValue) MarshalForProof(wr io.Writer) error {
	return iv.Marshal(wr)
}
