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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"testing"
)

func TestGas(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}
	chain := common.HexToAddress("0x037c4d7bbb0407d1e2c64981855ad8681d0d86d1")
	sender := common.HexToAddress("0xe91e00167939cb6694d2c422acd208a007293948")
	connAddress := common.HexToAddress("0x2aad3e8302f74e0818b7bcd10c2c050526707755")

	constructorData := hexutil.MustDecode(arbostestcontracts.GasUsedBin)

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	constructorTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{0},
		Payment:     big.NewInt(0),
		Data:        constructorData,
	}

	noopEOACallTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.RandAddress(),
		Payment:     big.NewInt(0),
		Data:        nil,
	}

	noopFuncCallTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: connAddress,
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode("0x5dfc2e4a"),
	}

	storeFuncCallTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(3),
		DestAddress: connAddress,
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode("0x703c2d1a"),
	}

	store2FuncCallTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(4),
		DestAddress: connAddress,
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode("0x703c2d1a"),
	}

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(message.Eth{Dest: sender, Value: big.NewInt(10000)}, chain, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(constructorTx), sender, big.NewInt(2), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(noopEOACallTx), sender, big.NewInt(3), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(noopFuncCallTx), sender, big.NewInt(4), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(storeFuncCallTx), sender, big.NewInt(5), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(store2FuncCallTx), sender, big.NewInt(6), chainTime),
	}

	assertion, _ := mach.ExecuteAssertion(10000000000, inboxMessages, 0)
	logs := assertion.ParseLogs()
	sends := assertion.ParseOutMessages()
	testCase, err := inbox.TestVectorJSON(inboxMessages, logs, sends)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(testCase))

	results := processTxResults(t, logs)
	if len(results) != len(inboxMessages)-2 {
		t.Fatal("unxpected log count", len(results))
	}

	if len(sends) != 0 {
		t.Fatal("unxpected send count", len(sends))
	}

	checkConstructorResult(t, results[0], connAddress)
	validGasCheck(t, results[1])
	validGasCheck(t, results[2])
	validGasCheck(t, results[3])
	validGasCheck(t, results[4])
}

func validGasCheck(t *testing.T, res *evm.TxResult) *big.Int {
	t.Log("GasUsed", res.GasUsed)
	if res.ResultCode != evm.ReturnCode {
		t.Log("result", res)
		t.Error("unexpected result", res.ResultCode)
	}
	return res.GasUsed
}
