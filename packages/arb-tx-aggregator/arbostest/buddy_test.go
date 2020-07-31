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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"log"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

func TestBuddyContract(t *testing.T) {
	arbERC20Data, err := hexutil.Decode(ArbERC20Bin)
	if err != nil {
		t.Fatal(err)
	}

	//erc20ABI, err := abi.JSON(strings.NewReader(ArbERC20ABI))
	//if err != nil {
	//	t.Fatal(err)
	//}

	//getNameABI := erc20ABI.Methods["name"]
	//getNameSignature, err := hexutil.Decode("0x06fdde03")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//generateTx := l2message.Transaction{
	//	MaxGas:      big.NewInt(1000000000),
	//	GasPriceBid: big.NewInt(0),
	//	SequenceNum: big.NewInt(1),
	//	DestAddress: fibAddress,
	//	Payment:     big.NewInt(300),
	//	Data:        append(generateSignature, generateFibData...),
	//}

	chainTime := message.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	addr := common.Address{1, 2, 3, 4, 5}

	//initializeBuddyContractData, err := initializeBuddyContractABI.Inputs.Pack(chainAddress, inboxAddress)
	//if err != nil {
	//	t.Fatal(err)
	//}

	//instantiateContractData, err = initializeBuddyContractABI.Inputs.Pack(inboxAddress)
	//if err != nil {
	//	t.Fatal(err)
	//}

	inbox := value.NewEmptyTuple()

	initMsg := message.Init{
		ChainParams: valprotocol.ChainParams{
			StakeRequirement:        big.NewInt(0),
			GracePeriod:             common.TimeTicks{Val: big.NewInt(0)},
			MaxExecutionSteps:       0,
			ArbGasSpeedLimitPerTick: 0,
		},
		Owner:       common.Address{},
		ExtraConfig: []byte{},
	}
	inbox = value.NewTuple2(inbox, message.NewInboxMessage(initMsg, addr, big.NewInt(0), chainTime).AsValue())

	l1contract := common.RandAddress()

	buddyConstructor := l2message.ContractTransaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        arbERC20Data,
	}
	buddyMsg := message.NewInboxMessage(
		message.BuddyDeployment{Data: l2message.L2MessageAsData(buddyConstructor)},
		l1contract,
		big.NewInt(1),
		chainTime,
	)

	//buddyContractAddress := common.HexToAddress("0x4ee09d87c0112181f1aa950e259a3e2d3bbd7e49")

	inbox = value.NewTuple2(inbox, buddyMsg.AsValue())

	//inbox = value.NewTuple2(inbox, message.NewInboxMessage(
	//	message.L2Message{Data: l2message.L2MessageAsData(makeConstructorTx(pointsConstructorData, big.NewInt(1)))},
	//	addr,
	//	big.NewInt(2),
	//	chainTime,
	//).AsValue())
	//
	//inbox = value.NewTuple2(inbox, message.NewInboxMessage(
	//	message.L2Message{Data: l2message.L2MessageAsData(l2message.Transaction{
	//		MaxGas:      big.NewInt(1000000000),
	//		GasPriceBid: big.NewInt(0),
	//		SequenceNum: big.NewInt(2),
	//		DestAddress: distributionsAddress,
	//		Payment:     big.NewInt(0),
	//		Data:        append(instantiateContractSignature, initializeBuddyContractData...),
	//	})},
	//	addr,
	//	big.NewInt(3),
	//	chainTime,
	//).AsValue())

	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	assertion, _ := mach.ExecuteAssertion(1000000000, inbox, 0)
	data, err := value.TestVectorJSON(inbox, assertion.ParseLogs(), assertion.ParseOutMessages())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))

	logs := assertion.ParseLogs()
	if len(logs) != 1 {
		t.Fatal("unexpected log count", len(logs))
	}

	sends := assertion.ParseOutMessages()
	if len(sends) != 1 {
		t.Fatal("unexpected send count", len(sends))
	}

	t.Log("send", sends[0])

	for _, logVal := range assertion.ParseLogs() {
		res, err := evm.NewResultFromValue(logVal)
		if err != nil {
			t.Fatal(err)
		}
		if res.ResultCode != evm.ReturnCode {
			t.Error("tx failed", res.ResultCode)
		}
		log.Println("ReturnData", hexutil.Encode(res.ReturnData))
		//if res.L1Message.Kind == message.L2Type {
		//	l2, err := l2message.NewL2MessageFromData(res.L1Message.Data)
		//	if err != nil {
		//		t.Fatal(err)
		//	}
		//}
	}
}
