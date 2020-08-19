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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestCrossContract(t *testing.T) {
	distributionsConstructorData, err := hexutil.Decode(DistributionsV0Bin)
	if err != nil {
		t.Fatal(err)
	}

	pointsConstructorData, err := hexutil.Decode(SubredditPointsV0Bin)
	if err != nil {
		t.Fatal(err)
	}

	distABI, err := abi.JSON(strings.NewReader(DistributionsV0ABI))
	if err != nil {
		t.Fatal(err)
	}

	instantiateContractTestABI := distABI.Methods["instantiateContractTest"]
	instantiateContractSignature, err := hexutil.Decode("0x1c2a2551")
	if err != nil {
		t.Fatal(err)
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	addr := common.Address{1, 2, 3, 4, 5}

	distributionsAddress := common.HexToAddress("0xba59937520bd4c1067bac24fb774b981b4b8c115")
	pointsAddress := common.HexToAddress("0x93fe8c8771c698af5a59a9a049ed02f2c71fefc4")

	instantiateContractData, err := instantiateContractTestABI.Inputs.Pack(pointsAddress)
	if err != nil {
		t.Fatal(err)
	}

	inboxMessages := make([]inbox.InboxMessage, 0)
	inboxMessages = append(inboxMessages, message.NewInboxMessage(initMsg(), addr, big.NewInt(0), chainTime))
	inboxMessages = append(inboxMessages, message.NewInboxMessage(
		message.NewSafeL2Message(makeConstructorTx(distributionsConstructorData, big.NewInt(0))),
		addr,
		big.NewInt(1),
		chainTime,
	))

	inboxMessages = append(inboxMessages, message.NewInboxMessage(
		message.NewSafeL2Message(makeConstructorTx(pointsConstructorData, big.NewInt(1))),
		addr,
		big.NewInt(2),
		chainTime,
	))

	inboxMessages = append(inboxMessages, message.NewInboxMessage(
		message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(1000000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(2),
			DestAddress: distributionsAddress,
			Payment:     big.NewInt(0),
			Data:        append(instantiateContractSignature, instantiateContractData...),
		}),
		addr,
		big.NewInt(3),
		chainTime,
	))

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	assertion, _ := mach.ExecuteAssertion(1000000000, inboxMessages, 0)
	data, err := inbox.TestVectorJSON(inboxMessages, assertion.ParseLogs(), assertion.ParseOutMessages())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))

	logs := assertion.ParseLogs()
	log.Println("Assertion had", len(logs), "logs")

	for _, logVal := range assertion.ParseLogs() {
		res, err := evm.NewResultFromValue(logVal)
		if err != nil {
			t.Fatal(err)
		}

		txRes, ok := res.(*evm.TxResult)
		if !ok {
			t.Fatal("incorrect res type", res)
		}

		if txRes.ResultCode != evm.ReturnCode {
			t.Error("tx failed", txRes.ResultCode)
		}
		log.Println("ReturnData", hexutil.Encode(txRes.ReturnData))
		if txRes.IncomingRequest.Kind == message.L2Type {
			l2, err := message.L2Message{Data: txRes.IncomingRequest.Data}.AbstractMessage()
			if err != nil {
				t.Fatal(err)
			}
			log.Println(l2)
		}

	}
}
