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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

// TestBuddyContract verifies that buddy contract deployment works and that
// regular contract deployment and buddy deployment interact correctly
func TestBuddyContract(t *testing.T) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	simpleCode := hexutil.MustDecode(arbostestcontracts.SimpleBin)
	buddyConstructor := message.BuddyDeployment{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		Payment:     big.NewInt(0),
		Data:        simpleCode,
	}

	fibCode := hexutil.MustDecode(arbostestcontracts.FibonacciBin)
	contractCreation := makeSimpleConstructorTx(fibCode, big.NewInt(0))
	contractCreation2 := makeSimpleConstructorTx(fibCode, big.NewInt(1))

	messages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(buddyConstructor, connAddress1, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(contractCreation), sender, big.NewInt(2), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(contractCreation2), sender, big.NewInt(3), chainTime),
		message.NewInboxMessage(buddyConstructor, connAddress2, big.NewInt(4), chainTime),
	}

	logs, sends, mach := runAssertion(t, messages, 4, 1)
	results := processTxResults(t, logs)

	checkConstructorResult(t, results[0], connAddress1)
	txResultCheck(t, results[1], evm.ContractAlreadyExists)
	checkConstructorResult(t, results[2], connAddress2)
	txResultCheck(t, results[3], evm.ContractAlreadyExists)

	msg, err := message.NewOutMessageFromValue(sends[0])
	failIfError(t, err)
	if msg.Sender != connAddress1 {
		t.Error("Buddy contract created at wrong address")
	}

	snap := snapshot.NewSnapshot(mach.Clone(), inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}, message.ChainAddressToID(chain), big.NewInt(5))

	conn1Code, err := snap.GetCode(connAddress1)
	failIfError(t, err)
	if !bytes.Contains(simpleCode, conn1Code) {
		t.Error("wrong code for first contract")
	}

	conn2Code, err := snap.GetCode(connAddress2)
	failIfError(t, err)
	if !bytes.Contains(fibCode, conn2Code) {
		t.Error("wrong code for second contract")
	}
}
