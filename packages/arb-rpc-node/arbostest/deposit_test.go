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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"math/big"
	"strings"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestDepositEthTx(t *testing.T) {
	depositDest := common.RandAddress()

	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	failIfError(t, err)

	deployTx := makeSimpleConstructorTx(hexutil.MustDecode(arbostestcontracts.SimpleBin), big.NewInt(0))

	// Deposit to EOA
	tx := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(1000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(1),
			DestAddress: depositDest,
			Payment:     big.NewInt(100),
			Data:        nil,
		}),
	}

	getBalance1 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        snapshot.GetBalanceData(depositDest),
	}

	// Deposit to contract that succeeds
	tx2 := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(3),
			DestAddress: connAddress1,
			Payment:     big.NewInt(200),
			Data:        makeFuncData(t, simpleABI.Methods["acceptPayment"]),
		}),
	}

	getBalance2 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(4),
		DestAddress: common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        snapshot.GetBalanceData(connAddress1),
	}

	// Deposit to contract that reverts
	tx3 := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.Transaction{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(5),
			DestAddress: connAddress1,
			Payment:     big.NewInt(50),
			Data:        makeFuncData(t, simpleABI.Methods["rejectPayment"]),
		}),
	}

	getBalance3 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(6),
		DestAddress: common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        snapshot.GetBalanceData(connAddress1),
	}

	getBalance4 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(7),
		DestAddress: common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        snapshot.GetBalanceData(sender),
	}

	tx4 := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(10000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: connAddress1,
				Payment:     big.NewInt(100),
				Data:        makeFuncData(t, simpleABI.Methods["acceptPayment"]),
			},
		}),
	}

	getBalance5 := message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(8),
		DestAddress: common.NewAddressFromEth(arbos.ARB_INFO_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        snapshot.GetBalanceData(connAddress1),
	}

	messages := []message.Message{
		message.NewSafeL2Message(deployTx),
		tx,
		message.NewSafeL2Message(getBalance1),
		tx2,
		message.NewSafeL2Message(getBalance2),
		tx3,
		message.NewSafeL2Message(getBalance3),
		message.NewSafeL2Message(getBalance4),
		tx4,
		message.NewSafeL2Message(getBalance5),
	}
	logs, _, _, _ := runAssertion(t, makeSimpleInbox(messages), 10, 0)
	results := processTxResults(t, logs)

	checkConstructorResult(t, results[0], connAddress1)
	succeededTxCheck(t, results[1])
	succeededTxCheck(t, results[2])
	succeededTxCheck(t, results[3])
	succeededTxCheck(t, results[4])
	revertedTxCheck(t, results[5])
	succeededTxCheck(t, results[6])
	succeededTxCheck(t, results[7])
	succeededTxCheck(t, results[8])
	succeededTxCheck(t, results[9])

	balance1, err := snapshot.ParseBalanceResult(results[2])
	failIfError(t, err)
	balance2, err := snapshot.ParseBalanceResult(results[4])
	failIfError(t, err)
	balance3, err := snapshot.ParseBalanceResult(results[6])
	failIfError(t, err)
	balance4, err := snapshot.ParseBalanceResult(results[7])
	failIfError(t, err)
	balance5, err := snapshot.ParseBalanceResult(results[9])
	failIfError(t, err)

	if balance1.Cmp(big.NewInt(100)) != 0 {
		t.Error("wrong balance1", balance1)
	}

	if balance2.Cmp(big.NewInt(200)) != 0 {
		t.Error("wrong balance2", balance2)
	}

	if balance3.Cmp(big.NewInt(200)) != 0 {
		t.Error("wrong balance3", balance3)
	}

	if balance4.Cmp(big.NewInt(50)) != 0 {
		t.Error("wrong balance4", balance4)
	}

	if balance5.Cmp(big.NewInt(300)) != 0 {
		t.Error("wrong balance5", balance5)
	}

	if results[1].IncomingRequest.Kind != message.EthDepositTxType {
		t.Error("wrong msg type")
	}
	if results[8].IncomingRequest.Kind != message.EthDepositTxType {
		t.Error("wrong msg type")
	}
}
