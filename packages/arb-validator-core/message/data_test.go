package message

import (
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

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
