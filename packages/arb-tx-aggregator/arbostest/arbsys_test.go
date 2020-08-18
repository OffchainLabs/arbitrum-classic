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
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
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
)

func generateFib(val *big.Int) ([]byte, error) {
	fib, err := abi.JSON(strings.NewReader(FibonacciABI))
	if err != nil {
		return nil, err
	}

	generateFibABI := fib.Methods["generateFib"]
	generateFibData, err := generateFibABI.Inputs.Pack(val)
	if err != nil {
		return nil, err
	}

	generateSignature, err := hexutil.Decode("0x2ddec39b")
	if err != nil {
		return nil, err
	}
	return append(generateSignature, generateFibData...), nil
}

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
	randDest := common.RandAddress()
	correctTxCount := 0

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	checkTxCount := func(target int) error {
		snap := snapshot.NewSnapshot(mach, chainTime, message.ChainAddressToID(chain), big.NewInt(9999999))
		txCount, err := snap.GetTransactionCount(addr)
		if err != nil {
			t.Fatal(err)
		}
		if txCount.Cmp(big.NewInt(int64(target))) != 0 {
			return fmt.Errorf("wrong tx count %v", txCount)
		}
		t.Log("Current tx count is", txCount)
		return nil
	}

	runMessage(t, mach, initMsg(), chain)

	if err := checkTxCount(0); err != nil {
		t.Fatal(err)
	}

	depositEth(t, mach, addr, big.NewInt(1000))

	// Deposit doesn't increase tx count
	if err := checkTxCount(correctTxCount); err != nil {
		t.Fatal(err)
	}

	tx1 := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(int64(correctTxCount)),
		DestAddress: randDest,
		Payment:     big.NewInt(300),
		Data:        []byte{},
	}

	_, err = runValidTransaction(t, mach, tx1, addr)
	if err != nil {
		t.Fatal(err)
	}
	correctTxCount++

	// Payment to EOA increases tx count
	if err := checkTxCount(correctTxCount); err != nil {
		t.Fatal(err)
	}

	tx2 := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(int64(correctTxCount) + 1),
		DestAddress: randDest,
		Payment:     big.NewInt(10),
		Data:        []byte{},
	}

	runMessage(t, mach, message.NewSafeL2Message(tx2), addr)

	// Payment to EOA with incorrect sequence number shouldn't increase tx count
	if err := checkTxCount(correctTxCount); err != nil {
		t.Fatal(err)
	}

	tx3 := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(int64(correctTxCount)),
		DestAddress: randDest,
		Payment:     big.NewInt(30000),
		Data:        []byte{},
	}

	_, err = runTransaction(t, mach, tx3, addr)
	if err != nil {
		t.Fatal(err)
	}

	// Payment to EOA with insufficient funds shouldn't increase tx count
	if err := checkTxCount(correctTxCount); err != nil {
		t.Fatal(err)
	}

	constructorData, err := hexutil.Decode(FibonacciBin)
	if err != nil {
		t.Fatal(err)
	}

	fibAddress, err := deployContract(t, mach, addr, constructorData, big.NewInt(int64(correctTxCount)))
	if err != nil {
		t.Fatal(err)
	}
	correctTxCount++

	// Contract deployment increases tx count
	if err := checkTxCount(correctTxCount); err != nil {
		t.Fatal(err)
	}

	fibData, err := generateFib(big.NewInt(20))
	if err != nil {
		t.Fatal(err)
	}

	generateTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(int64(correctTxCount)),
		DestAddress: fibAddress,
		Payment:     big.NewInt(300),
		Data:        fibData,
	}

	_, err = runValidTransaction(t, mach, generateTx, addr)
	if err != nil {
		t.Fatal(err)
	}

	correctTxCount++

	// Tx call increases tx count
	if err := checkTxCount(correctTxCount); err != nil {
		t.Fatal(err)
	}

	generateTx2 := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(int64(correctTxCount + 1)),
		DestAddress: fibAddress,
		Payment:     big.NewInt(300),
		Data:        fibData,
	}

	runMessage(t, mach, message.NewSafeL2Message(generateTx2), addr)

	// Tx call with incorrect sequence number doesn't affect the count
	if err := checkTxCount(correctTxCount); err != nil {
		t.Fatal(err)
	}

	generateTx3 := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(int64(correctTxCount)),
		DestAddress: fibAddress,
		Payment:     big.NewInt(100000),
		Data:        fibData,
	}

	res, err := runTransaction(t, mach, generateTx3, addr)
	if err != nil {
		t.Fatal(err)
	}
	if res.ResultCode != evm.InsufficientTxFundsCode {
		t.Fatal("incorrect return code", res.ResultCode)
	}

	// Tx call with insufficient balance doesn't affect the count
	if err := checkTxCount(correctTxCount); err != nil {
		t.Fatal(err)
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

	depositMsg := message.Eth{
		Dest:  addr,
		Value: big.NewInt(10000),
	}

	depositValue := big.NewInt(100)
	withdrawDest := common.RandAddress()
	tx := withdrawEthTx(t, big.NewInt(0), depositValue, withdrawDest)

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(depositMsg, addr, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx), addr, big.NewInt(2), chainTime),
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
