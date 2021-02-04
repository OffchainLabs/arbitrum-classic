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

package inbox

import (
	"github.com/pkg/errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

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

var errTupleSize2 = errors.New("expected 2-tuple value")

func ByteArrayToBytes(val value.Value) ([]byte, error) {
	tupVal, ok := val.(*value.TupleValue)
	if !ok || tupVal.Len() != 2 {
		return nil, errors.New("expected byte array to be 2 tuple")
	}
	sizeVal, _ := tupVal.GetByInt64(0)
	contents, _ := tupVal.GetByInt64(1)

	sizeInt, ok := sizeVal.(value.IntValue)
	if !ok {
		return nil, errors.New("size must be an int")
	}
	contentsBuffer, ok := contents.(*value.Buffer)
	if !ok {
		return nil, errors.New("contents must be an buffer")
	}

	size := sizeInt.BigInt().Uint64()
	if uint64(len(contentsBuffer.Data())) > size {
		return nil, errors.Errorf("bytearray buffer too small, size=%v, length=%v", size, len(contentsBuffer.Data()))
	}
	data := make([]byte, size)
	copy(data[:], contentsBuffer.Data())
	return data, nil
}

func StackValueToList(val value.Value) ([]value.Value, error) {
	tupVal, ok := val.(*value.TupleValue)
	if !ok {
		return nil, errors.Wrap(errTupleSize2, val.String())
	}
	values := make([]value.Value, 0)
	for tupVal.Len() != 0 {
		if tupVal.Len() != 2 {
			return nil, errors.Wrap(errTupleSize2, val.String())
		}

		// Tuple size already verified above, so error can be ignored
		member, _ := tupVal.GetByInt64(0)
		val, _ := tupVal.GetByInt64(1)

		tupVal, ok = val.(*value.TupleValue)
		if !ok {
			return nil, errors.Wrap(errTupleSize2, val.String())
		}

		values = append(values, member)
	}

	for i := len(values)/2 - 1; i >= 0; i-- {
		opp := len(values) - 1 - i
		values[i], values[opp] = values[opp], values[i]
	}
	return values, nil
}

func ListToStackValue(vals []value.Value) *value.TupleValue {
	ret := value.NewEmptyTuple()
	for _, val := range vals {
		ret = value.NewTuple2(val, ret)
	}
	return ret
}

func BytesToByteStack(val []byte) *value.TupleValue {
	chunks := bytesToValues(val)
	ret := ListToStackValue(chunks)
	return value.NewTuple2(value.NewInt64Value(int64(len(val))), ret)
}
