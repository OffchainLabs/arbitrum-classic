/*
* Copyright 2021, Offchain Labs, Inc.
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
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestMinimumGas(t *testing.T) {
	tx1 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode(arbostestcontracts.FailedSendBin),
	}
	tx2 := message.Transaction{
		MaxGas:      big.NewInt(1000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        nil,
	}

	messages := []message.Message{
		message.NewSafeL2Message(tx1),
		message.NewSafeL2Message(tx2),
	}
	logs, _, _, _ := runAssertion(t, makeSimpleInbox(messages), len(messages), 0)
	results := processTxResults(t, logs)
	incoming := extractIncomingMessages(t, results)
	l2Messages := filterL2Messages(t, incoming)
	if len(l2Messages) != 2 {
		t.Fatal("expected 2 l2 messages")
	}
	checkConstructorResult(t, results[0], connAddress1)
	tx2Ret, ok := l2Messages[1].(message.Transaction)
	if !ok {
		t.Fatal("expected 2nd receipt to be from transaction message")
	}
	if arbosVersion < 7 {
		if !tx2Ret.Equals(tx1) {
			t.Error("expected to hit incorrect behavior in old version")
		}
		txResultCheck(t, results[1], evm.RevertCode)
	} else {
		if !tx2Ret.Equals(tx2) {
			t.Error("2nd receipt's incoming message doesn't match input")
		}
		txResultCheck(t, results[1], evm.MinArbGasForContractTx)
	}
}
