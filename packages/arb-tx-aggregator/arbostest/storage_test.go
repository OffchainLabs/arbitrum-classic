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
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"log"
	"math/big"
	"testing"
)

func TestGetStorageAt(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	chain := common.RandAddress()
	sender := common.RandAddress()

	connAddr := common.HexToAddress("0x2aad3e8302f74e0818b7bcd10c2c050526707755")

	constructorTx := makeConstructorTx(
		hexutil.MustDecode(arbostestcontracts.StorageBin),
		big.NewInt(0),
	)

	getStorageAtTx := message.ContractTransaction{
		MaxGas:      big.NewInt(100000000000),
		GasPriceBid: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        snapshot.GetStorageAtData(connAddr, big.NewInt(1)),
	}

	failGetStorageAtTx := message.ContractTransaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		DestAddress: connAddr,
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode("0x188f9139"),
	}

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(constructorTx), sender, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(getStorageAtTx), common.Address{}, big.NewInt(2), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(failGetStorageAtTx), sender, big.NewInt(3), chainTime),
	}

	assertion, _ := mach.ExecuteAssertion(10000000000, inboxMessages, 0)
	testCase, err := inbox.TestVectorJSON(inboxMessages, assertion.ParseLogs(), assertion.ParseOutMessages())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(testCase))
	logs := assertion.ParseLogs()
	sends := assertion.ParseOutMessages()

	if len(logs) != 3 {
		log.Println("unxpected log count", len(logs))
	}

	if len(sends) != 0 {
		log.Println("unxpected send count", len(sends))
	}

	constructorRes, err := evm.NewTxResultFromValue(logs[0])
	if err != nil {
		t.Fatal(err)
	}
	if constructorRes.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected constructor result", constructorRes.ResultCode)
	}
	connAddrCalc, err := getConstructorResult(constructorRes)
	if err != nil {
		t.Fatal(err)
	}
	if connAddrCalc != connAddr {
		t.Fatal("constructed address doesn't match")
	}
	getStorageAtRes, err := evm.NewTxResultFromValue(logs[1])
	if err != nil {
		t.Fatal(err)
	}
	if getStorageAtRes.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected get storage at result", getStorageAtRes.ResultCode)
	}
	storageVal := new(big.Int).SetBytes(getStorageAtRes.ReturnData)
	if storageVal.Cmp(big.NewInt(12345)) != 0 {
		t.Fatal("expected storage to be 12345 but got", storageVal)
	}
	log.Println(hexutil.Encode(getStorageAtRes.ReturnData))
	failGetStorageAtRes, err := evm.NewTxResultFromValue(logs[2])
	if err != nil {
		t.Fatal(err)
	}
	if failGetStorageAtRes.ResultCode != evm.RevertCode {
		t.Fatal("unexpected fail get storage at result", failGetStorageAtRes.ResultCode)
	}
}
