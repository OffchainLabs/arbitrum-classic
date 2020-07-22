/*
* Copyright 2020, Offchain Labs, Inc.
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
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func TestListToStackValue(t *testing.T) {
	vals := make([]value.Value, 0)
	for i := int64(0); i < 10; i++ {
		vals = append(vals, value.NewInt64Value(i))
	}
	stackVal := ListToStackValue(vals)

	vals2, err := StackValueToList(stackVal)
	if err != nil {
		t.Fatal(err)
	}

	if len(vals) != len(vals2) {
		t.Fatal("wrong val count")
	}

	for i, val := range vals {
		if !value.Eq(val, vals2[i]) {
			t.Fatal("val not equal")
		}
	}
}

func TestByteStackConversion(t *testing.T) {
	data := make([]byte, 100)
	rand.Read(data)
	bytestack := BytesToByteStack(data)
	data2, err := ByteStackToHex(bytestack)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data2) {
		t.Log("data:", hexutil.Encode(data))
		t.Log("data2:", hexutil.Encode(data2))
		t.Error("data changed in conversion")
	}
}

func TestStackValueToListFailures(t *testing.T) {
	intVal := value.NewInt64Value(0)
	if _, err := StackValueToList(intVal); err == nil {
		t.Error("should fail when passed non-tuple")
	}

	tup, _ := value.NewTupleFromSlice([]value.Value{intVal, intVal, intVal})
	if _, err := StackValueToList(tup); err == nil {
		t.Error("should fail when passed tuple not of size 2")
	}
}

func TestByteStackToHexFailures(t *testing.T) {
	intVal := value.NewInt64Value(0)
	if _, err := ByteStackToHex(intVal); err == nil {
		t.Error("should fail when passed non-tuple")
	}

	tup, _ := value.NewTupleFromSlice([]value.Value{intVal, intVal, intVal})
	if _, err := ByteStackToHex(tup); err == nil {
		t.Error("should fail when passed tuple not of size 2")
	}

	tup, _ = value.NewTupleFromSlice([]value.Value{value.NewEmptyTuple(), intVal})
	if _, err := ByteStackToHex(tup); err == nil {
		t.Error("should fail when first value is not an int")
	}

	tup, _ = value.NewTupleFromSlice([]value.Value{intVal, intVal})
	if _, err := ByteStackToHex(tup); err == nil {
		t.Error("should fail when second value is not a stack value")
	}

	listVal := ListToStackValue([]value.Value{value.NewEmptyTuple()})
	tup, _ = value.NewTupleFromSlice([]value.Value{intVal, listVal})
	if _, err := ByteStackToHex(tup); err == nil {
		t.Error("should fail when second value contains non ints in the stack")
	}
}
