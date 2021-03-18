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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
	"strings"
	"testing"
)

func TestTransfer(t *testing.T) {
	constructorData := hexutil.MustDecode(arbostestcontracts.TransferBin)

	constructorTx1 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{0},
		Payment:     big.NewInt(100),
		Data:        constructorData,
	}

	constructorTx2 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.Address{0},
		Payment:     big.NewInt(100),
		Data:        constructorData,
	}

	transferABI, err := abi.JSON(strings.NewReader(arbostestcontracts.TransferABI))
	failIfError(t, err)
	connCallTx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, transferABI.Methods["send2"], connAddress2),
	}

	inboxMessages := makeSimpleInbox([]message.Message{
		message.Eth{Dest: sender, Value: big.NewInt(10000)},
		message.NewSafeL2Message(constructorTx1),
		message.NewSafeL2Message(constructorTx2),
		message.NewSafeL2Message(connCallTx),
	})

	logs, _, snap, _ := runAssertion(t, inboxMessages, 3, 0)
	results := processTxResults(t, logs)

	allResultsSucceeded(t, results)

	checkConstructorResult(t, results[0], connAddress1)
	checkConstructorResult(t, results[1], connAddress2)

	res := results[2]
	t.Log("GasUsed", res.GasUsed)
	t.Log("GasLimit", connCallTx.MaxGas)

	checkBalance(t, snap, connAddress1, big.NewInt(101))
	checkBalance(t, snap, connAddress2, big.NewInt(99))
	checkBalance(t, snap, sender, big.NewInt(9800))
}
