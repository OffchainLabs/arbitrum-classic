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

package ethbridgemachine

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	goarbitrum "github.com/offchainlabs/arbitrum/packages/arb-provider-go"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"log"
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

func runMessage(t *testing.T, mach machine.Machine, msg message.Message, sender common.Address) []*evm.Result {
	chainTime := message.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	inbox := structures.NewVMInbox()
	inbox.DeliverMessage(message.NewInboxMessage(msg, sender, big.NewInt(0), chainTime))
	assertion, steps := mach.ExecuteAssertion(1000000000, inbox.AsValue(), 0)
	t.Log("Ran assertion for", steps, "steps and had", assertion.LogsCount, "logs")
	if mach.CurrentStatus() != machine.Extensive {
		t.Fatal("machine should still be working")
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

func runTransaction(t *testing.T, mach machine.Machine, msg message.Message, sender common.Address) (*evm.Result, error) {
	results := runMessage(t, mach, msg, sender)
	if len(results) != 1 {
		return nil, fmt.Errorf("unexpected log count %v", len(results))
	}
	result := results[0]
	if result.ResultCode != evm.ReturnCode {
		return nil, fmt.Errorf("transaction failed unexpectedly %v", result)
	}
	return result, nil
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

	getBalance := message.Call{
		MaxGas:      big.NewInt(0),
		GasPriceBid: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(goarbitrum.ARB_INFO_ADDRESS),
		Data:        append(getBalanceSignature, getBalanceData...),
	}
	balanceResult, err := runTransaction(t, mach, message.L2Message{Msg: getBalance}, common.Address{})
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

func deployFib(t *testing.T, mach machine.Machine, sender common.Address) (common.Address, error) {
	constructorData, err := hexutil.Decode(FibonacciBin)
	if err != nil {
		return common.Address{}, err
	}

	constructorTx := message.Transaction{
		MaxGas:      big.NewInt(0),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        constructorData,
	}

	constructorResult, err := runTransaction(t, mach, message.L2Message{Msg: constructorTx}, sender)
	if err != nil {
		return common.Address{}, err
	}
	if len(constructorResult.ReturnData) != 32 {
		return common.Address{}, errors.New("unexpected constructor result length")
	}
	var fibAddress common.Address
	copy(fibAddress[:], constructorResult.ReturnData[12:])
	return fibAddress, nil
}

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

	addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))
	chain := common.RandAddress()

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
	results := runMessage(t, mach, initMsg, chain)
	log.Println(results)

	fibAddress, err := deployFib(t, mach, addr)
	if err != nil {
		t.Fatal(err)
	}
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
		SequenceNum: big.NewInt(1),
		DestAddress: fibAddress,
		Payment:     big.NewInt(300),
		Data:        append(generateSignature, generateFibData...),
	}

	generateResult, err := runTransaction(t, mach, message.L2Message{Msg: generateTx}, addr)
	if err != nil {
		t.Fatal(err)
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

	getFibResult, err := runTransaction(t, mach, message.L2Message{Msg: getFibTx}, addr)
	if err != nil {
		t.Fatal(err)
	}
	if hexutil.Encode(getFibResult.ReturnData) != "0x0000000000000000000000000000000000000000000000000000000000000008" {
		t.Fatal("getFib had incorrect result")
	}
}

func TestDeposit(t *testing.T) {
	mach, err := loader.LoadMachineFromFile(gotest.TestMachinePath(), false, "cpp")
	if err != nil {
		t.Fatal(err)
	}
	inbox := structures.NewVMInbox()
	mach.ExecuteAssertion(1000000000, inbox.AsValue(), 0)

	pk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))

	msg := message.Eth{
		Dest:  addr,
		Value: big.NewInt(1000),
	}

	depositResults := runMessage(t, mach, msg, addr)
	if len(depositResults) != 0 {
		t.Fatal("deposit should not have had a result")
	}

	if getBalanceCall(t, mach, addr).Cmp(msg.Value) != 0 {
		t.Fatal("incorrect balance")
	}
}

func TestBatch(t *testing.T) {
	chain := common.RandAddress()
	chainId := new(big.Int).SetBytes(chain[14:])

	mach, err := loader.LoadMachineFromFile(gotest.TestMachinePath(), false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

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
	results := runMessage(t, mach, initMsg, chain)
	log.Println(results)

	dest, err := deployFib(t, mach, common.RandAddress())
	if err != nil {
		t.Fatal(err)
	}
	batchSize := 20
	txes := make([]message.BatchTx, 0, batchSize)
	senders := make([]common.Address, 0, batchSize)
	for i := 0; i < batchSize; i++ {
		pk, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		addr := common.NewAddressFromEth(crypto.PubkeyToAddress(pk.PublicKey))
		senders = append(senders, addr)
		depositResults := runMessage(
			t,
			mach,
			message.Eth{
				Dest:  addr,
				Value: big.NewInt(1000),
			},
			addr,
		)
		if len(depositResults) != 0 {
			t.Fatal("deposit should not have had a result")
		}

		tx := types.NewTransaction(0, dest.ToEthAddress(), big.NewInt(0), 0, big.NewInt(0), []byte{})
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), pk)
		if err != nil {
			t.Fatal(err)
		}
		txes = append(txes, message.NewBatchTxFromSignedEthTx(signedTx))
	}

	msg := message.TransactionBatch{Transactions: txes}
	results = runMessage(t, mach, message.L2Message{Msg: msg}, common.RandAddress())
	if len(results) != len(txes) {
		t.Fatal("incorrect result count")
	}
	for i, result := range results {
		if result.L1Message.Sender != senders[i] {
			t.Fatal("message had incorrect sender", result.L1Message.Sender, senders[i])
		}
	}
}
