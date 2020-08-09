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
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

func TestTransactionCount(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
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

	txCount := getTransactionCountCall(t, mach, addr)
	if txCount.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("wrong tx count", txCount)
	}

	constructorData, err := hexutil.Decode(FibonacciBin)
	if err != nil {
		t.Fatal(err)
	}

	fibAddress, err := deployContract(t, mach, addr, constructorData, big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fibAddress.Hex())

	depositEth(t, mach, addr, big.NewInt(1000))

	txCount = getTransactionCountCall(t, mach, addr)
	if txCount.Cmp(big.NewInt(1)) != 0 {
		t.Fatal("wrong tx count", txCount)
	}

	fib, err := abi.JSON(strings.NewReader(FibonacciABI))
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
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: fibAddress,
		Payment:     big.NewInt(300),
		Data:        append(generateSignature, generateFibData...),
	}

	_, err = runTransaction(t, mach, generateTx, addr)
	if err != nil {
		t.Fatal(err)
	}

	txCount = getTransactionCountCall(t, mach, addr)
	if txCount.Cmp(big.NewInt(2)) != 0 {
		t.Fatal("wrong tx count", txCount)
	}
}

func TestWithdrawEth(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	addr := common.RandAddress()
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
	depositMsg := message.Eth{
		Dest:  addr,
		Value: big.NewInt(10000),
	}

	depositValue := big.NewInt(100)
	withdrawDest := common.RandAddress()
	tx := withdrawEthTx(t, big.NewInt(0), depositValue, withdrawDest)

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg, chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(depositMsg, addr, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewL2Message(tx), addr, big.NewInt(2), chainTime),
	}

	assertion, _ := mach.ExecuteAssertion(10000000000, inboxMessages, 0)
	testCase, err := inbox.TestVectorJSON(inboxMessages, assertion.ParseLogs(), assertion.ParseOutMessages())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(testCase))
	logs := assertion.ParseLogs()

	if len(logs) != 1 {
		t.Fatal("unexpected log count", len(logs))
	}

	res, err := evm.NewTxResultFromValue(logs[0])
	if err != nil {
		t.Fatal(err)
	}
	if res.ResultCode != evm.ReturnCode {
		t.Fatal("incorrect tx response", res.ResultCode)
	}

	sends := assertion.ParseOutMessages()
	if len(sends) != 1 {
		t.Fatal("unexpected send count")
	}

	outMsg, err := message.NewOutMessageFromValue(sends[0])
	if err != nil {
		t.Fatal(err)
	}

	if outMsg.Kind != message.EthType {
		t.Fatal("outgoing message had wrong type", outMsg.Kind)
	}

	if outMsg.Sender != addr {
		t.Fatal("wrong withdraw sender")
	}

	outEthMsg := message.NewEthFromData(outMsg.Data)

	if outEthMsg.Value.Cmp(depositValue) != 0 {
		t.Fatal("wrong withdraw value", outEthMsg.Value)
	}

	if outEthMsg.Dest != withdrawDest {
		t.Fatal("wrong withdraw destination", outEthMsg.Dest)
	}
}
