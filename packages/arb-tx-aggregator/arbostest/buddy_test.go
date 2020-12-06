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
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestBuddyContract(t *testing.T) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	addr := common.Address{1, 2, 3, 4, 5}

	l1contract := common.RandAddress()

	buddyConstructor := message.BuddyDeployment{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode(arbostestcontracts.SimpleBin),
	}

	l2Tx := message.Transaction{
		MaxGas:      big.NewInt(100000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: l1contract,
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode("0x267c4ae4"),
	}

	messages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), addr, big.NewInt(0), chainTime),
		message.NewInboxMessage(buddyConstructor, l1contract, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(l2Tx), common.RandAddress(), big.NewInt(2), chainTime),
	}

	logs, sends, _ := runAssertion(t, messages)

	results := processTxResults(t, logs)
	if len(results) != 2 {
		t.Fatal("unexpected log count", len(results))
	}

	if len(sends) != 1 {
		t.Fatal("unexpected send count", len(sends))
	}
	allResultsSucceeded(t, results)
	for i, res := range results {
		if i == 0 {
			if len(res.ReturnData) != 32 {
				t.Fatal("Unexpected return data length")
			}
			if !bytes.Equal(res.ReturnData[12:], l1contract[:]) {
				t.Log("Returned address", hexutil.Encode(res.ReturnData))
				t.Log("l1 address", l1contract)
				t.Error("constructor returned incorrect address")
			}
		} else {
			t.Log("ReturnData", hexutil.Encode(res.ReturnData))
			if len(res.ReturnData) == 0 {
				t.Error("expected return data")
			}
		}
	}

	for _, sendVal := range sends {
		msg, err := message.NewOutMessageFromValue(sendVal)
		failIfError(t, err)
		if msg.Sender != l1contract {
			t.Error("Buddy contract created at wrong address")
		}
	}
}
