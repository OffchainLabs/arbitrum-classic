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

package evm

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func getBigTuple(val value.Value, index uint64) (value.Value, error) {
	tup, ok := val.(value.TupleValue)
	if !ok {
		return nil, errors.New("not in bigtuple format")
	}
	if tup.Len() == 0 {
		return value.NewInt64Value(0), nil
	} else if index == 0 {
		return tup.GetByInt64(7)
	} else {
		subTup, err := tup.GetByInt64(int64(index) % 7)
		if err != nil {
			return nil, err
		}
		return getBigTuple(subTup, index/7)
	}
}

func setBigTuple(tupleVal value.Value, index uint64, val value.Value) (value.Value, error) {
	tup, ok := tupleVal.(value.TupleValue)
	if !ok {
		return nil, errors.New("not in bigtuple format")
	}
	if tup.Len() == 0 {
		tup, _ = value.NewRepeatedTuple(value.NewEmptyTuple(), 8)
	}
	if index == 0 {
		tup, _ = tup.SetByInt64(7, val)
		return tup, nil
	}

	subTup, err := tup.GetByInt64(int64(index) % 7)
	if err != nil {
		return nil, err
	}
	newSubTup, err := setBigTuple(subTup, index/7, val)
	if err != nil {
		return nil, err
	}
	return tup.SetByInt64(int64(index)%7, newSubTup)
}

func getByteArray(val value.Value, index uint64) (*big.Int, error) {
	item := index / 32
	extraBytes := (index % 32) * 8
	if extraBytes == 0 {
		subVal, err := getBigTuple(val, item)
		if err != nil {
			return nil, err
		}
		intVal, ok := subVal.(value.IntValue)
		if !ok {
			return nil, errors.New("bytearray expected int value")
		}
		return intVal.BigInt(), nil
	}

	first, err := getBigTuple(val, item)
	if err != nil {
		return nil, err
	}
	firstInt, ok := first.(value.IntValue)
	if !ok {
		return nil, errors.New("bytearray expected int value")
	}
	firstBig := firstInt.BigInt()

	second, err := getBigTuple(val, item+1)
	if err != nil {
		return nil, err
	}
	secondInt, ok := second.(value.IntValue)
	if !ok {
		return nil, errors.New("bytearray expected int value")
	}
	secondBig := secondInt.BigInt()

	firstBig = math.U256(firstBig.Lsh(firstBig, uint(extraBytes)))
	secondBig = math.U256(secondBig.Rsh(secondBig, uint(256-extraBytes)))
	return firstBig.Or(firstBig, secondBig), nil
}

func SizedByteArrayToHex(val value.Value) ([]byte, error) {
	var buf bytes.Buffer
	tup, ok := val.(value.TupleValue)
	if !ok {
		return nil, errors.New("sized bytearray expected tuple value")
	}
	lengthVal, err := tup.GetByInt64(1)
	if err != nil {
		return nil, err
	}
	lengthIntVal, ok := lengthVal.(value.IntValue)
	if !ok {
		return nil, errors.New("sized bytearray expected int value")
	}
	intLength := lengthIntVal.BigInt().Uint64()

	byteArrayVal, err := tup.GetByInt64(0)
	if err != nil {
		return nil, err
	}

	for i := uint64(0); i < intLength; i += 32 {
		val, err := getByteArray(byteArrayVal, i)
		if err != nil {
			return nil, err
		}
		valBytes := value.NewIntValue(val).ToBytes()
		buf.Write(valBytes[:])
	}
	return buf.Bytes()[:intLength], nil
}

func IntsToBigTuple(ints []*big.Int) (value.Value, error) {
	var tuple value.Value = value.NewEmptyTuple()
	var err error
	for i, val := range ints {
		tuple, err = setBigTuple(tuple, uint64(i), value.NewIntValue(val))
		if err != nil {
			return nil, err
		}
	}
	return tuple, nil
}

func BytesToByteArray(val []byte) (value.Value, error) {
	var ints []*big.Int

	for i := 0; i < len(val); i += 32 {
		remaining := len(val) - i
		if remaining < 32 {
			data := [32]byte{}
			copy(data[:], val[i:])
			ints = append(ints, new(big.Int).SetBytes(data[:]))
		} else {
			ints = append(ints, new(big.Int).SetBytes(val[i:i+32]))
		}
	}
	return IntsToBigTuple(ints)
}

func BytesToSizedByteArray(val []byte) (value.Value, error) {
	arr, err := BytesToByteArray(val)
	if err != nil {
		return nil, err
	}
	return value.NewTuple2(arr, value.NewInt64Value(int64(len(val)))), nil
}

func StackValueToList(val value.Value) ([]value.Value, error) {
	values := make([]value.Value, 0)
	for !val.Equal(value.NewEmptyTuple()) {
		tupVal, ok := val.(value.TupleValue)
		if !ok {
			return nil, errors.New("value was not in stack format")
		}
		member, err := tupVal.GetByInt64(1)
		if err != nil {
			return nil, err
		}
		values = append(values, member)
		val, err = tupVal.GetByInt64(0)
		if err != nil {
			return nil, err
		}
	}
	return values, nil
}
