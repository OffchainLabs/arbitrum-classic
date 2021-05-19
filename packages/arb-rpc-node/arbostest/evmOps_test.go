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
	"bytes"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func TestEVMOps(t *testing.T) {
	opcodesABI, err := abi.JSON(strings.NewReader(arbostestcontracts.OpCodesABI))
	failIfError(t, err)

	tx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode(arbostestcontracts.OpCodesBin),
	}

	tx2 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode(arbostestcontracts.OpCodesBin),
	}

	tx3 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, opcodesABI.Methods["getBlockHash"]),
	}

	tx4 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(3),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, opcodesABI.Methods["getOrigin"]),
	}

	tx5 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(4),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, opcodesABI.Methods["getNestedOrigin"], connAddress2),
	}

	tx6 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(5),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, opcodesABI.Methods["getSender"]),
	}

	tx7 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(6),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, opcodesABI.Methods["getNestedSend"], connAddress2),
	}

	messages := []message.Message{
		message.NewSafeL2Message(tx),
		message.NewSafeL2Message(tx2),
		message.NewSafeL2Message(tx3),
		message.NewSafeL2Message(tx4),
		message.NewSafeL2Message(tx5),
		message.NewSafeL2Message(tx6),
		message.NewSafeL2Message(tx7),
	}

	results, _ := runSimpleTxAssertion(t, messages)
	allResultsSucceeded(t, results)
	checkConstructorResult(t, results[0], connAddress1)
	checkConstructorResult(t, results[1], connAddress2)

	if !bytes.Equal(results[2].ReturnData, common.Hash{}.Bytes()) {
		t.Error("Unexpected block hash result")
	}

	var correctOrigin common.Hash
	copy(correctOrigin[12:], sender.Bytes())

	if !bytes.Equal(results[3].ReturnData, correctOrigin.Bytes()) {
		t.Error("Unexpected origin")
	}

	if !bytes.Equal(results[4].ReturnData, correctOrigin.Bytes()) {
		t.Error("Unexpected origin")
	}

	if !bytes.Equal(results[5].ReturnData, correctOrigin.Bytes()) {
		t.Error("Unexpected sender")
	}

	var correctSender common.Hash
	copy(correctSender[12:], connAddress1.Bytes())

	if !bytes.Equal(results[6].ReturnData, correctSender.Bytes()) {
		t.Error("Unexpected sender")
	}
}
