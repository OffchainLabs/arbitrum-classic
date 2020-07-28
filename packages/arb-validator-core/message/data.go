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

package message

import (
	"bytes"
	"errors"
	errors2 "github.com/pkg/errors"
	"math"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func toEth(val *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(val.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
}

func extractUInt256(data []byte) (*big.Int, []byte) {
	val := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	return val, data
}

func extractAddress(data []byte) (common.Address, []byte) {
	data = data[12:] // Skip first 12 bytes of 32 byte address data
	var addr common.Address
	copy(addr[:], data[:])
	data = data[20:]
	return addr, data
}

func addressData(addr common.Address) []byte {
	ret := make([]byte, 0, 32)
	ret = append(ret, make([]byte, 12)...)
	ret = append(ret, addr[:]...)
	return ret
}

func bytesToValues(val []byte) []value.Value {
	var ints []value.Value
	for i := 0; i < len(val); i += 32 {
		remaining := len(val) - i
		if remaining < 32 {
			data := [32]byte{}
			copy(data[:], val[i:])
			ints = append(
				ints,
				value.NewIntValue(new(big.Int).SetBytes(data[:])),
			)
		} else {
			ints = append(
				ints,
				value.NewIntValue(new(big.Int).SetBytes(val[i:i+32])),
			)
		}
	}
	return ints
}

var errInt = errors.New("expected int value")
var errTupleSize2 = errors.New("expected 2-tuple value")

func StackValueToList(val value.Value) ([]value.Value, error) {
	values := make([]value.Value, 0)
	for !val.Equal(value.NewEmptyTuple()) {
		tupVal, ok := val.(value.TupleValue)
		if !ok {
			return nil, errors2.Wrap(errTupleSize2, val.String())
		}
		if tupVal.Len() != 2 {
			return nil, errors2.Wrap(errTupleSize2, val.String())
		}

		member, _ := tupVal.GetByInt64(0)
		val, _ = tupVal.GetByInt64(1)

		values = append(values, member)
	}

	for i := len(values)/2 - 1; i >= 0; i-- {
		opp := len(values) - 1 - i
		values[i], values[opp] = values[opp], values[i]
	}
	return values, nil
}

func ListToStackValue(vals []value.Value) value.TupleValue {
	ret := value.NewEmptyTuple()
	for _, val := range vals {
		ret = value.NewTuple2(val, ret)
	}
	return ret
}

func ByteStackToHex(val value.Value) ([]byte, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return nil, errors2.Wrap(errTupleSize2, val.String())
	}
	if tup.Len() != 2 {
		return nil, errors2.Wrap(errTupleSize2, val.String())
	}
	lengthVal, _ := tup.GetByInt64(0)
	lengthIntVal, ok := lengthVal.(value.IntValue)
	if !ok {
		return nil, errInt
	}
	intLength := lengthIntVal.BigInt().Uint64()

	stackVal, _ := tup.GetByInt64(1)

	byteChunks := make([][32]byte, 0)
	vals, err := StackValueToList(stackVal)
	if err != nil {
		return nil, err
	}

	for _, val := range vals {
		intVal, ok := val.(value.IntValue)
		if !ok {
			return nil, errInt
		}
		byteChunks = append(byteChunks, intVal.ToBytes())
	}

	var buf bytes.Buffer
	for _, chunk := range byteChunks {
		buf.Write(chunk[:])
	}
	return buf.Bytes()[:intLength], nil
}

func BytesToByteStack(val []byte) value.TupleValue {
	chunks := bytesToValues(val)
	ret := ListToStackValue(chunks)
	return value.NewTuple2(value.NewInt64Value(int64(len(val))), ret)
}
