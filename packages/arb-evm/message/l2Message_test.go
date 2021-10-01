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
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"

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

	randomBatch, err := NewRandomTransactionBatch(10, pk, 0, common.RandBigInt())
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

func TestCompressedECDSAEncoding(t *testing.T) {
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
	v := byte(0)

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

	correctEncoded := []byte{0xff, 128, 128, 133, 232, 212, 165, 16, 0, 148, 243, 101, 124, 147, 250, 217, 103, 9, 37, 122, 103, 44, 160, 214, 230, 81, 119, 46, 3, 73, 128, 119, 22, 2, 247, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 48, 242, 130, 147, 4, 245, 114, 73, 189, 227, 45, 98, 143, 159, 80, 204, 118, 150, 152, 52, 26, 32, 58, 251, 218, 23, 113, 41, 44, 137, 37, 254, 26, 67, 153, 128, 134, 97, 76, 163, 86, 191, 84, 156, 105, 228, 138, 94, 56, 87, 196, 213, 22, 68, 22, 238, 39, 162, 182, 18, 31, 165, 43, 232, 0}

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

func TestCompressedECDSAConversion(t *testing.T) {
	chainId := common.RandBigInt()
	signer := types.NewEIP155Signer(chainId)
	tx := types.NewTransaction(rand.Uint64(), common.RandAddress().ToEthAddress(), common.RandBigInt(), rand.Uint64(), common.RandBigInt(), common.RandBytes(100))
	pk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, signer, pk)
	if err != nil {
		t.Fatal(err)
	}
	compressed := NewCompressedECDSAFromEth(signedTx)
	tx2, err := compressed.AsEthTx(chainId)
	if err != nil {
		t.Fatal(err)
	}
	tx1JSON, err := signedTx.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	tx2JSON, err := tx2.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(tx1JSON))
	t.Log(string(tx2JSON))
	sender1, err := types.Sender(signer, signedTx)
	if err != nil {
		t.Fatal(err)
	}
	sender2, err := types.Sender(signer, tx2)
	if err != nil {
		t.Fatal(err)
	}
	if sender1 != sender2 {
		t.Fatal("senders don't match")
	}
	if signedTx.Hash() != tx2.Hash() {
		t.Fatal("decoded tx incorrectly")
	}
}
