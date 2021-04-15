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
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
	"testing"
)

func TestGas(t *testing.T) {
	conData := hexutil.MustDecode(arbostestcontracts.GasUsedBin)
	conData = append(conData, math.U256Bytes(big.NewInt(0))...)
	constructorTx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{0},
		Payment:     big.NewInt(0),
		Data:        conData,
	}

	noopEOACallTx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.RandAddress(),
		Payment:     big.NewInt(0),
		Data:        nil,
	}

	noopFuncCallTx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode("0x5dfc2e4a"),
	}

	storeFuncCallTx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(3),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode("0x703c2d1a"),
	}

	store2FuncCallTx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(4),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode("0x703c2d1a"),
	}

	messages := []message.Message{
		makeEthDeposit(sender, big.NewInt(10000)),
		message.NewSafeL2Message(constructorTx),
		message.NewSafeL2Message(noopEOACallTx),
		message.NewSafeL2Message(noopFuncCallTx),
		message.NewSafeL2Message(storeFuncCallTx),
		message.NewSafeL2Message(store2FuncCallTx),
	}

	logs, _, _, _ := runAssertion(t, makeSimpleInbox(messages), len(messages), 0)
	results := processTxResults(t, logs)

	allResultsSucceeded(t, results)

	checkConstructorResult(t, results[1], connAddress1)
	validGasCheck(t, results[2])
	validGasCheck(t, results[3])
	validGasCheck(t, results[4])
	validGasCheck(t, results[5])
}

func validGasCheck(t *testing.T, res *evm.TxResult) *big.Int {
	t.Log("GasUsed", res.GasUsed)
	return res.GasUsed
}
