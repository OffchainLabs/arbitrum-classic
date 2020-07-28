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
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"
	"strings"
	"testing"
)

func runMessage(t *testing.T, mach machine.Machine, msg message.Message, sender common.Address) []*evm.Result {
	chainTime := message.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	inbox := value.NewEmptyTuple()
	inbox = value.NewTuple2(inbox, message.NewInboxMessage(msg, sender, big.NewInt(0), chainTime).AsValue())
	assertion, steps := mach.ExecuteAssertion(1000000000, inbox, 0)
	//data, err := value.TestVectorJSON(inbox, assertion.ParseLogs(), assertion.ParseOutMessages())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(string(data))
	t.Log("Ran assertion for", steps, "steps and had", assertion.LogsCount, "logs")
	if mach.CurrentStatus() != machine.Extensive {
		t.Fatal("machine should still be working")
	}
	blockReason := mach.IsBlocked(false)
	if blockReason == nil {
		t.Fatal("machine not blocked")
	}

	if _, ok := blockReason.(machine.InboxBlocked); !ok {
		t.Fatal("Machine blocked for weird reason", blockReason)
	}
	results := make([]*evm.Result, 0)
	for _, avmLog := range assertion.ParseLogs() {
		result, err := evm.NewResultFromValue(avmLog)
		if err != nil {
			t.Fatal(err)
		}
		results = append(results, result)
	}
	return results
}

func runTransaction(t *testing.T, mach machine.Machine, msg l2message.AbstractL2Message, sender common.Address) (*evm.Result, error) {
	results := runMessage(t, mach, message.L2Message{Data: l2message.L2MessageAsData(msg)}, sender)
	if len(results) != 1 {
		return nil, fmt.Errorf("unexpected log count %v", len(results))
	}
	result := results[0]
	if result.ResultCode != evm.ReturnCode {
		return nil, fmt.Errorf("transaction failed unexpectedly %v", result)
	}
	return result, nil
}

func getTransactionCountCall(t *testing.T, mach machine.Machine, address common.Address) *big.Int {
	arbsys, err := abi.JSON(strings.NewReader(arboscontracts.ArbSysABI))
	if err != nil {
		t.Fatal(err)
	}

	txabi := arbsys.Methods["getTransactionCount"]
	txData, err := txabi.Inputs.Pack(address)
	if err != nil {
		t.Fatal(err)
	}

	funcSig, err := hexutil.Decode("0x23ca0cd2")
	if err != nil {
		t.Fatal(err)
	}

	call := l2message.Call{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
		Data:        append(funcSig, txData...),
	}
	funcResult, err := runTransaction(t, mach.Clone(), call, common.Address{})
	if err != nil {
		t.Fatal(err)
	}
	vals, err := txabi.Outputs.UnpackValues(funcResult.ReturnData)
	if len(vals) != 1 {
		t.Fatal("unexpected tx result")
	}
	val, ok := vals[0].(*big.Int)
	if !ok {
		t.Fatal("unexpected tx result")
	}
	return val
}

func getBalanceCall(t *testing.T, mach machine.Machine, address common.Address) *big.Int {
	info, err := abi.JSON(strings.NewReader(arboscontracts.ArbInfoABI))
	if err != nil {
		t.Fatal(err)
	}

	getBalanceABI := info.Methods["getBalance"]
	getBalanceData, err := getBalanceABI.Inputs.Pack(address)
	if err != nil {
		t.Fatal(err)
	}

	getBalanceSignature, err := hexutil.Decode("0xf8b2cb4f")
	if err != nil {
		t.Fatal(err)
	}

	getBalance := l2message.Call{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS),
		Data:        append(getBalanceSignature, getBalanceData...),
	}
	balanceResult, err := runTransaction(t, mach.Clone(), getBalance, common.Address{})
	if err != nil {
		t.Fatal(err)
	}
	vals, err := getBalanceABI.Outputs.UnpackValues(balanceResult.ReturnData)
	if len(vals) != 1 {
		t.Fatal("unexpected tx result")
	}
	val, ok := vals[0].(*big.Int)
	if !ok {
		t.Fatal("unexpected tx result")
	}
	return val
}

func makeConstructorTx(code []byte, sequenceNum *big.Int) l2message.Transaction {
	return l2message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        code,
	}
}

func deployContract(t *testing.T, mach machine.Machine, sender common.Address, code []byte, sequenceNum *big.Int) (common.Address, error) {
	constructorTx := makeConstructorTx(code, sequenceNum)
	constructorResult, err := runTransaction(t, mach, constructorTx, sender)
	if err != nil {
		return common.Address{}, err
	}
	if len(constructorResult.ReturnData) != 32 {
		return common.Address{}, errors.New("unexpected constructor result length")
	}
	var contractAddress common.Address
	copy(contractAddress[:], constructorResult.ReturnData[12:])
	return contractAddress, nil
}

func depositEth(t *testing.T, mach machine.Machine, dest common.Address, amount *big.Int) {
	msg := message.Eth{
		Dest:  dest,
		Value: amount,
	}

	depositResults := runMessage(t, mach, msg, dest)
	if len(depositResults) != 0 {
		t.Fatal("deposit should not have had a result")
	}
}
