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
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"strings"
	"testing"
)

func TestTransfer(t *testing.T) {

	constructorData := hexutil.MustDecode(arbostestcontracts.TransferBin)

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	constructorTx1 := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{0},
		Payment:     big.NewInt(100),
		Data:        constructorData,
	}

	constructorTx2 := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: common.Address{0},
		Payment:     big.NewInt(100),
		Data:        constructorData,
	}

	transferABI, err := abi.JSON(strings.NewReader(arbostestcontracts.TransferABI))
	failIfError(t, err)
	sendABI := transferABI.Methods["send2"]
	sendData, err := sendABI.Inputs.Pack(connAddress2)
	failIfError(t, err)
	connCallTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        append(sendABI.ID, sendData...),
	}

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(message.Eth{Dest: sender, Value: big.NewInt(10000)}, chain, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(constructorTx1), sender, big.NewInt(2), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(constructorTx2), sender, big.NewInt(3), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(connCallTx), sender, big.NewInt(4), chainTime),
	}

	logs, _, mach := runAssertion(t, inboxMessages, 3, 0)
	results := processTxResults(t, logs)

	allResultsSucceeded(t, results)

	checkConstructorResult(t, results[0], connAddress1)
	checkConstructorResult(t, results[1], connAddress2)

	res := results[2]
	t.Log("GasUsed", res.GasUsed)
	t.Log("GasLimit", connCallTx.MaxGas)

	snap := snapshot.NewSnapshot(mach, chainTime, message.ChainAddressToID(chain), big.NewInt(4))
	transfer1Balance, err := snap.GetBalance(connAddress1)
	failIfError(t, err)
	transfer2Balance, err := snap.GetBalance(connAddress2)
	failIfError(t, err)
	senderBalance, err := snap.GetBalance(sender)
	failIfError(t, err)

	if transfer1Balance.Cmp(big.NewInt(101)) != 0 {
		t.Error("unexpected transfer conn1 balance", transfer1Balance)
	}

	if transfer2Balance.Cmp(big.NewInt(99)) != 0 {
		t.Error("unexpected transfer conn2 balance", transfer2Balance)
	}

	if senderBalance.Cmp(big.NewInt(9800)) != 0 {
		t.Error("unexpected sender balance", senderBalance)
	}
}
