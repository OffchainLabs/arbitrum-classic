package evm

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestByteStackConversion(t *testing.T) {
	data := make([]byte, 100)
	rand.Read(data)
	bytestack, err := BytesToByteStack(data)
	if err != nil {
		t.Error(err)
	}
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
