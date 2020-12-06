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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"strings"
	"testing"
)

func TestTransfer(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}
	chain := common.HexToAddress("0x037c4d7bbb0407d1e2c64981855ad8681d0d86d1")
	sender := common.HexToAddress("0xe91e00167939cb6694d2c422acd208a007293948")
	transfer1Address := common.HexToAddress("0x2aad3e8302f74e0818b7bcd10c2c050526707755")
	transfer2Address := common.HexToAddress("0x016cb751543d1cca5dd02976ac8dbdc0ecaacafd")

	constructorData := hexutil.MustDecode(arbostestcontracts.TransferBin)

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	constructorTx1 := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{0},
		Payment:     big.NewInt(100),
		Data:        constructorData,
	}

	constructorTx2 := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.Address{0},
		Payment:     big.NewInt(100),
		Data:        constructorData,
	}

	transferABI, err := abi.JSON(strings.NewReader(arbostestcontracts.TransferABI))
	if err != nil {
		t.Fatal(err)
	}

	sendABI := transferABI.Methods["send2"]

	sendData, err := sendABI.Inputs.Pack(transfer2Address)
	if err != nil {
		t.Fatal(err)

	}

	connCallTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: transfer1Address,
		Payment:     big.NewInt(0),
		Data:        append(hexutil.MustDecode("0x3386b1a2"), sendData...),
	}

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(message.Eth{Dest: sender, Value: big.NewInt(10000)}, chain, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(constructorTx1), sender, big.NewInt(2), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(constructorTx2), sender, big.NewInt(3), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(connCallTx), sender, big.NewInt(4), chainTime),
	}

	assertion, _ := mach.ExecuteAssertion(10000000000, inboxMessages, 0)
	logs := assertion.ParseLogs()
	sends := assertion.ParseOutMessages()
	testCase, err := inbox.TestVectorJSON(inboxMessages, logs, sends)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(testCase))

	results := processTxResults(t, assertion.ParseLogs())
	if len(results) != 3 {
		t.Fatal("unxpected log count", len(results))
	}

	if len(sends) != 0 {
		t.Fatal("unxpected send count", len(sends))
	}

	allResultsSucceeded(t, results)

	checkConstructorResult(t, results[0], transfer1Address)
	checkConstructorResult(t, results[1], transfer2Address)

	res := results[2]
	t.Log("GasUsed", res.GasUsed)
	t.Log("GasLimit", connCallTx.MaxGas)

	snap := snapshot.NewSnapshot(mach, chainTime, message.ChainAddressToID(chain), big.NewInt(4))
	transfer1Balance, err := snap.GetBalance(transfer1Address)
	if err != nil {
		t.Fatal(err)
	}
	transfer2Balance, err := snap.GetBalance(transfer2Address)
	if err != nil {
		t.Fatal(err)
	}
	senderBalance, err := snap.GetBalance(sender)
	if err != nil {
		t.Fatal(err)
	}

	if transfer1Balance.Cmp(big.NewInt(101)) != 0 {
		t.Error("unexpected transfer conn1 balance", transfer1Balance)
	}

	if transfer2Balance.Cmp(big.NewInt(99)) != 0 {
		t.Error("unexpected transfer conn2 balance", transfer2Balance)
	}

	if senderBalance.Cmp(big.NewInt(9800)) != 0 {
		t.Error("unexpected sender balance", senderBalance)
	}
}
