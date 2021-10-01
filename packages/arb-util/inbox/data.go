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
	"math/big"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

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
		return nil, errors.New("byte array size must be an int")
	}
	contentsBuffer, ok := contents.(*value.Buffer)
	if !ok {
		return nil, errors.New("contents must be an buffer")
	}

	return BufAndLengthToBytes(sizeInt.BigInt(), contentsBuffer)
}

func BufAndLengthToBytes(sizeInt *big.Int, contents *value.Buffer) ([]byte, error) {
	size := sizeInt.Uint64()
	if uint64(len(contents.Data())) > size {
		return nil, errors.Errorf("buffer too small, size=%v, length=%v", size, len(contents.Data()))
	}
	data := make([]byte, size)
	copy(data[:], contents.Data())
	return data, nil
}

func BufOffsetAndLengthToBytes(sizeInt, offsetInt *big.Int, contents *value.Buffer) []byte {
	size := sizeInt.Uint64()
	offset := offsetInt.Uint64()
	data := make([]byte, size)
	if offset > uint64(len(contents.Data())) {
		return data
	}
	max := offset + size
	if max > uint64(len(contents.Data())) {
		max = uint64(len(contents.Data()))
	}
	copy(data[:], contents.Data()[offset:max])
	return data
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
