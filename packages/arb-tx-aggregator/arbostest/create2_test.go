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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"log"
	"math/big"
	"strings"
	"testing"
)

func TestCreate2(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	chain := common.RandAddress()
	sender := common.HexToAddress("0x8c988ec54f112dd35666e19e7b0904bb12df1b6c")

	factoryConnAddress := common.HexToAddress("0x7cc1af94bfb4676c4facfc6a56430ec35c45b8b0")
	simpleConnAddress := common.HexToAddress("0x21a02de2bb91c1b23eede23c4a8d11fda2c57ad8")
	cloneConnAddress := common.HexToAddress("0xb51F90cC1f31Da7191a325EEFC5Fe2534AAF5F7f")

	factoryConstructorTx := makeConstructorTx(
		hexutil.MustDecode(arbostestcontracts.CloneFactoryBin),
		big.NewInt(0),
	)

	simpleConstructorTx := makeConstructorTx(
		hexutil.MustDecode(arbostestcontracts.SimpleBin),
		big.NewInt(1),
	)

	factoryABI, err := abi.JSON(strings.NewReader(arbostestcontracts.CloneFactoryABI))
	if err != nil {
		t.Fatal(factoryABI)
	}

	create2ABI := factoryABI.Methods["create2Clone"]
	create2Data, err := create2ABI.Inputs.Pack(simpleConnAddress, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	if err != nil {
		t.Fatal(factoryABI)
	}

	existsABI := simpleABI.Methods["exists"]

	create2Tx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: factoryConnAddress,
		Payment:     big.NewInt(0),
		Data:        append(hexutil.MustDecode("0xc91091c3"), create2Data...),
	}

	existsCloneTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(3),
		DestAddress: cloneConnAddress,
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode("0x267c4ae4"),
	}

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(factoryConstructorTx), sender, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(simpleConstructorTx), sender, big.NewInt(2), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(create2Tx), sender, big.NewInt(3), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(existsCloneTx), sender, big.NewInt(4), chainTime),
	}

	assertion, _ := mach.ExecuteAssertion(10000000000, inboxMessages, 0)
	testCase, err := inbox.TestVectorJSON(inboxMessages, assertion.ParseLogs(), assertion.ParseOutMessages())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(testCase))
	logs := assertion.ParseLogs()
	sends := assertion.ParseOutMessages()

	if len(logs) != 4 {
		log.Println("unxpected log count", len(logs))
	}

	if len(sends) != 0 {
		log.Println("unxpected send count", len(sends))
	}

	factoryConstructorRes, err := evm.NewTxResultFromValue(logs[0])
	if err != nil {
		t.Fatal(err)
	}
	if factoryConstructorRes.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected constructor result", factoryConstructorRes.ResultCode)
	}
	factoryConnAddrCalc, err := getConstructorResult(factoryConstructorRes)
	if err != nil {
		t.Fatal(err)
	}
	if factoryConnAddrCalc != factoryConnAddress {
		t.Fatal("constructed address doesn't match:", factoryConnAddrCalc, "instead of", factoryConnAddress)
	}
	simpleConstructorRes, err := evm.NewTxResultFromValue(logs[1])
	if err != nil {
		t.Fatal(err)
	}
	if simpleConstructorRes.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected constructor result", simpleConstructorRes.ResultCode)
	}
	simpleConnAddrCalc, err := getConstructorResult(simpleConstructorRes)
	if err != nil {
		t.Fatal(err)
	}
	if simpleConnAddrCalc != simpleConnAddress {
		t.Fatal("constructed address doesn't match:", simpleConnAddrCalc, "instead of", simpleConnAddress)
	}

	create2Res, err := evm.NewTxResultFromValue(logs[2])
	if err != nil {
		t.Fatal(err)
	}
	if create2Res.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected create2 result", create2Res.ResultCode)
	}
	create2Outputs, err := create2ABI.Outputs.UnpackValues(create2Res.ReturnData)
	if err != nil {
		t.Fatal(err)
	}
	if len(create2Outputs) != 1 {
		t.Fatal("wrong output count")
	}
	createdAddress := common.NewAddressFromEth(create2Outputs[0].(ethcommon.Address))
	if createdAddress != cloneConnAddress {
		t.Fatal("incorrect clone address")
	}

	existsCloneRes, err := evm.NewTxResultFromValue(logs[3])
	if err != nil {
		t.Fatal(err)
	}
	if existsCloneRes.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected exists clone result", existsCloneRes.ResultCode)
	}

	existsCloneOutputs, err := existsABI.Outputs.UnpackValues(existsCloneRes.ReturnData)
	if err != nil {
		t.Fatal(err)
	}
	if len(existsCloneOutputs) != 1 {
		t.Fatal("wrong output count")
	}
	if existsCloneOutputs[0].(*big.Int).Cmp(big.NewInt(10)) != 0 {
		t.Fatal("wrong exists clone output")
	}
	snap := snapshot.NewSnapshot(mach, chainTime, message.ChainAddressToID(chain), big.NewInt(4))
	cloneCode, err := snap.GetCode(cloneConnAddress)
	if err != nil {
		t.Fatal(err)
	}
	if len(cloneCode) != 45 {
		t.Fatal("wrong clone code length")
	}
}
