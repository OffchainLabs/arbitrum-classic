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
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestFailedNestedSend(t *testing.T) {
	dest := common.RandAddress()

	tx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode(arbostestcontracts.FailedSendBin),
	}

	failedSend, err := abi.JSON(strings.NewReader(arbostestcontracts.FailedSendABI))
	failIfError(t, err)
	sendTx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: connAddress1,
		Payment:     big.NewInt(300),
		Data:        makeFuncData(t, failedSend.Methods["send"], dest),
	}

	messages := []message.Message{
		makeEthDeposit(sender, big.NewInt(1000)),
		message.NewSafeL2Message(tx),
		message.NewSafeL2Message(sendTx),
	}

	logs, _, _, _ := runSimpleAssertion(t, messages)
	results := processTxResults(t, logs)
	checkConstructorResult(t, results[1], connAddress1)
	revertedTxCheck(t, results[2])
}

func TestRevertedNestedCall(t *testing.T) {
	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	failIfError(t, err)

	tx1 := makeSimpleConstructorTx(hexutil.MustDecode(arbostestcontracts.SimpleBin), big.NewInt(0))
	tx2 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, simpleABI.Methods["nestedCall"], big.NewInt(0)),
	}
	tx3 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, simpleABI.Methods["nestedCall"], big.NewInt(10)),
	}
	messages := []message.Message{
		message.NewSafeL2Message(tx1),
		message.NewSafeL2Message(tx2),
		message.NewSafeL2Message(tx3),
	}
	logs, _, _, _ := runSimpleAssertion(t, messages)
	results := processTxResults(t, logs)
	checkConstructorResult(t, results[0], connAddress1)
	succeededTxCheck(t, results[1])
	succeededTxCheck(t, results[2])
}
