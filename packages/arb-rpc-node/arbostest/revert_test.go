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

package arbostest

import (
	"bytes"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
)

func TestRevert(t *testing.T) {
	simpleConstructorTx := makeSimpleConstructorTx(hexutil.MustDecode(arbostestcontracts.SimpleBin), big.NewInt(0))

	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	if err != nil {
		t.Fatal(err)
	}

	revertsTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        simpleABI.Methods["reverts"].ID,
	}

	messages := []message.Message{
		message.NewSafeL2Message(simpleConstructorTx),
		message.NewSafeL2Message(revertsTx),
	}

	logs, _, _, _ := runAssertion(t, makeSimpleInbox(messages), 2, 0)
	results := processTxResults(t, logs)

	checkConstructorResult(t, results[0], connAddress1)
	revertedTxCheck(t, results[1])

	correctResult := []byte("this is a test")
	if !bytes.Contains(results[1].ReturnData, correctResult) {
		t.Error("incorrect return data", hexutil.Encode(results[1].ReturnData))
	}
}
