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
	txData := tx.AsData()
	if len(txData) != TransactionHeaderSize+len(tx.Data) {
		t.Error("serialized tx has incorrect size")
	}

	l2Messages := []AbstractL2Message{
		tx,
		NewRandomContractTransaction(),
		NewRandomCall(),
		NewRandomTransactionBatch(10, common.RandAddress(), pk),
	}

	for _, msg := range l2Messages {
		t.Run(fmt.Sprintf("%T", msg), func(t *testing.T) {
			l2Message := NewL2Message(msg)
			decoded, err := l2Message.AbstractMessage()
			if err != nil {
				t.Fatal(err)
			}
			if bytes.Equal(decoded.AsData(), l2Message.AsData()) {
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
