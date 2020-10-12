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
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"log"
	"math/big"
	"strings"
	"testing"
)

func TestCreate2(t *testing.T) {
	client, pks := test.SimulatedBackend()
	auth := bind.NewKeyedTransactor(pks[0])

	factoryConnAddress, _, cf, err := arbostestcontracts.DeployCloneFactory(auth, client)
	if err != nil {
		t.Fatal(err)
	}

	simpleConnAddress, _, _, err := arbostestcontracts.DeploySimple(auth, client)
	if err != nil {
		t.Fatal(err)
	}
	client.Commit()

	tx, err := cf.Create2Clone(auth, simpleConnAddress, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	client.Commit()

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		t.Fatal(err)
	}

	ethEv, err := cf.ParseCreatedClone(*receipt.Logs[0])
	if err != nil {
		t.Fatal(err)
	}

	cloneConnAddress := ethEv.Clone

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	chain := common.RandAddress()
	sender := common.NewAddressFromEth(auth.From)

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
		DestAddress: common.NewAddressFromEth(factoryConnAddress),
		Payment:     big.NewInt(0),
		Data:        append(hexutil.MustDecode("0xc91091c3"), create2Data...),
	}

	existsCloneTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(3),
		DestAddress: common.NewAddressFromEth(cloneConnAddress),
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
	//testCase, err := inbox.TestVectorJSON(inboxMessages, assertion.ParseLogs(), assertion.ParseOutMessages())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(string(testCase))
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
	if factoryConnAddrCalc.ToEthAddress() != factoryConnAddress {
		t.Fatal("constructed address doesn't match:", factoryConnAddrCalc, "instead of", factoryConnAddress.Hex())
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
	if simpleConnAddrCalc.ToEthAddress() != simpleConnAddress {
		t.Fatal("constructed address doesn't match:", simpleConnAddrCalc, "instead of", simpleConnAddress.Hex())
	}

	create2Res, err := evm.NewTxResultFromValue(logs[2])
	if err != nil {
		t.Fatal(err)
	}
	if create2Res.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected create2 result", create2Res.ResultCode)
	}
	if len(create2Res.EVMLogs) != 1 {
		t.Fatal("wrong EVM log count")
	}

	ev, err := cf.ParseCreatedClone(*create2Res.ToEthReceipt(common.Hash{}).Logs[0])
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ArbOS clone address:", ev.Clone.Hex())

	if ev.Clone != cloneConnAddress {
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
	cloneCode, err := snap.GetCode(common.NewAddressFromEth(cloneConnAddress))
	if err != nil {
		t.Fatal(err)
	}
	if len(cloneCode) != 45 {
		t.Fatal("wrong clone code length")
	}
}
