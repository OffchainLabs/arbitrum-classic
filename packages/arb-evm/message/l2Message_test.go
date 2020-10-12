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
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

const AddressSize = 32
const TransactionHeaderSize = 32*4 + AddressSize

func TestL2MessageSerialization(t *testing.T) {
	pk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	tx := NewRandomTransaction()
	txData := tx.AsDataSafe()
	if len(txData) != TransactionHeaderSize+len(tx.Data) {
		t.Error("serialized tx has incorrect size")
	}

	randomBatch, err := NewRandomTransactionBatch(10, common.RandAddress(), pk, 0)
	if err != nil {
		t.Fatal(err)
	}
	l2Messages := []AbstractL2Message{
		tx,
		NewRandomContractTransaction(),
		NewRandomCall(),
		randomBatch,
	}

	for _, msg := range l2Messages {
		t.Run(fmt.Sprintf("%T", msg), func(t *testing.T) {
			l2Message, err := NewL2Message(msg)
			if err != nil {
				t.Fatal(err)
			}
			decoded, err := l2Message.AbstractMessage()
			if err != nil {
				t.Fatal(err)
			}
			data, err := decoded.AsData()
			if err != nil {
				t.Fatal(err)
			}
			if bytes.Equal(data, l2Message.AsData()) {
				t.Fatal("decoded l2 l2message not equal")
			}
		})
	}
}

func TestTransactionHash(t *testing.T) {
	txData, err := hexutil.Decode("0x000000000000000000000000000000000000000000000000000000174876e800000000000000000000000000000000000000000000000000000000000074aa31000000000000000000000000000000000000000000000000000000000052e4d40000000000000000000000007f6999eb9d18a44784045d87f3c67cf22746e99500000000000000000000000000000000000000000000000000000000051b5aa4af5a25367951baa2ff6cd471c483f15fb90badb37c5821b6d95526a41a9504680b4e7c8b763a1b1d49d4955c8486216325253fec738dd7a9e28bf921119c160f0702448615bbda08313f6a8eb668d20bf5059875921e668a5bdf2c7fc4844592d2572bcd")
	if err != nil {
		t.Fatal(err)
	}
	sender := common.HexToAddress("0xe91e00167939cb6694d2c422acd208a007293948")
	chain := common.HexToAddress("0x037c4d7bbb0407d1e2c64981855ad8681d0d86d1")
	targetHash := common.HexToHash("0x00532596242ba0ded0a8a17d8897344282fa1b29de676aa41aad6f737898e4a2")

	if newTransactionFromData(txData).MessageID(sender, chain) != targetHash {
		t.Error("incorrect hash")
	}
}

func TestCompressedECDSAFormat(t *testing.T) {
	calldata := []byte{119, 22, 2, 247, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	gasLimit, correct := new(big.Int).SetString("e8d4a51000", 16)
	if !correct {
		t.Fatal("bad gas limit")
	}

	compressedTx := CompressedTx{
		SequenceNum: big.NewInt(0),
		GasPrice:    big.NewInt(0),
		GasLimit:    gasLimit,
		To:          CompressedAddressFull{common.HexToAddress("0xf3657c93fad96709257a672ca0d6e651772e0349")},
		Payment:     big.NewInt(0),
		Calldata:    calldata,
	}

	r, _ := new(big.Int).SetString("22139494912332618468746784620225298078926562928862475257085431569185247929854", 10)
	s, _ := new(big.Int).SetString("11879572248721183921017568834333234971060281339844537723742821302023917743080", 10)
	v, _ := new(big.Int).SetString("2258", 10)

	tx := CompressedECDSATransaction{
		CompressedTx: compressedTx,
		V:            v,
		R:            r,
		S:            s,
	}

	encoded, err := tx.AsData()
	if err != nil {
		t.Fatal(err)
	}

	correctEncoded := []byte{128, 128, 133, 232, 212, 165, 16, 0, 148, 243, 101, 124, 147, 250, 217, 103, 9, 37, 122, 103, 44, 160, 214, 230, 81, 119, 46, 3, 73, 128, 119, 22, 2, 247, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 48, 242, 130, 147, 4, 245, 114, 73, 189, 227, 45, 98, 143, 159, 80, 204, 118, 150, 152, 52, 26, 32, 58, 251, 218, 23, 113, 41, 44, 137, 37, 254, 26, 67, 153, 128, 134, 97, 76, 163, 86, 191, 84, 156, 105, 228, 138, 94, 56, 87, 196, 213, 22, 68, 22, 238, 39, 162, 182, 18, 31, 165, 43, 232, 210}

	if !bytes.Equal(encoded, correctEncoded) {
		t.Error("incorrect encoded output")
	}

	l2, err := NewL2Message(tx)
	if err != nil {
		t.Fatal(err)
	}

	abs, err := l2.AbstractMessage()
	if err != nil {
		t.Fatal(err)
	}

	recovered, err := abs.AsData()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(hexutil.Encode(correctEncoded))
	t.Log(hexutil.Encode(recovered))
	if !bytes.Equal(correctEncoded, recovered) {
		t.Error("failed to unmarshall correctly")
	}
}
