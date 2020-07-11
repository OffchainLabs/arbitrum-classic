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

package proofmachine

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/gotest"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func TestFib(t *testing.T) {
	mach, err := loader.LoadMachineFromFile(gotest.TestMachinePath(), false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	fib, err := abi.JSON(strings.NewReader(FibonacciABI))
	if err != nil {
		t.Fatal(err)
	}

	pk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	runMessage := func(msg message.AbstractL2Message) *evm.Result {
		addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))

		chainTime := message.ChainTime{
			BlockNum:  common.NewTimeBlocksInt(0),
			Timestamp: big.NewInt(0),
		}

		inbox := structures.NewVMInbox()
		inbox.DeliverMessage(message.NewInboxMessage(message.L2Message{Msg: msg}, addr, big.NewInt(0), chainTime))
		assertion, _ := mach.ExecuteAssertion(1000000000, inbox.AsValue(), 0)

		if mach.CurrentStatus() != machine.Extensive {
			t.Fatal("machine should still be working")
		}
		logs := assertion.ParseLogs()
		if len(logs) != 1 {
			t.Fatal("unexpected log count")
		}

		res, err := evm.NewResultFromValue(logs[0])
		if err != nil {
			t.Fatal(err)
		}

		return res
	}

	constructorData, err := hexutil.Decode(FibonacciBin)
	if err != nil {
		t.Fatal(err)
	}

	constructorTx := message.Transaction{
		MaxGas:      big.NewInt(0),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        constructorData,
	}

	constructorResult := runMessage(constructorTx)
	if constructorResult.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected result", constructorResult.ResultCode)
	}
	if len(constructorResult.ReturnData) != 32 {
		t.Fatal("unexpected constructor result length")
	}
	var fibAddress common.Address
	copy(fibAddress[:], constructorResult.ReturnData[12:])

	generateFibABI := fib.Methods["generateFib"]
	generateFibData, err := generateFibABI.Inputs.Pack(big.NewInt(20))
	if err != nil {
		t.Fatal(err)
	}

	generateSignature, err := hexutil.Decode("0x2ddec39b")
	if err != nil {
		t.Fatal(err)
	}

	generateTx := message.Transaction{
		MaxGas:      big.NewInt(0),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: fibAddress,
		Payment:     big.NewInt(0),
		Data:        append(generateSignature, generateFibData...),
	}

	generateResult := runMessage(generateTx)
	if generateResult.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected result", generateResult.ResultCode)
	}
	if len(generateResult.EVMLogs) != 1 {
		t.Fatal("incorrect log count")
	}
	evmLog := generateResult.EVMLogs[0]
	if evmLog.Address != fibAddress {
		t.Fatal("log came from incorrect address")
	}
	if evmLog.Topics[0].ToEthHash() != fib.Events["TestEvent"].ID {
		t.Fatal("incorrect log topic")
	}
	if hexutil.Encode(evmLog.Data) != "0x0000000000000000000000000000000000000000000000000000000000000014" {
		t.Fatal("incorrect log data")
	}

	getFibABI := fib.Methods["getFib"]
	getFibData, err := getFibABI.Inputs.Pack(big.NewInt(5))
	if err != nil {
		t.Fatal(err)
	}

	getFibSignature, err := hexutil.Decode("0x90a3e3de")
	if err != nil {
		t.Fatal(err)
	}

	getFibTx := message.Call{
		MaxGas:      big.NewInt(0),
		GasPriceBid: big.NewInt(0),
		DestAddress: fibAddress,
		Data:        append(getFibSignature, getFibData...),
	}

	getFibResult := runMessage(getFibTx)
	if getFibResult.ResultCode != evm.ReturnCode {
		t.Fatal("unexpected result", getFibResult.ResultCode)
	}
	if hexutil.Encode(getFibResult.ReturnData) != "0x0000000000000000000000000000000000000000000000000000000000000008" {
		t.Fatal("getFib had incorrect result")
	}
}
