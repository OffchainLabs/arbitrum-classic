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
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
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

	laterChainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(1),
		Timestamp: big.NewInt(1),
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
	noOpTx := message.Transaction{
		MaxGas:      big.NewInt(100000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.RandAddress(),
		Payment:     big.NewInt(0),
		Data:        nil,
	}
	contractCreation2 := makeSimpleConstructorTx(fibCode, big.NewInt(1))

	messages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), big.NewInt(0), chainTime),
		message.NewInboxMessage(buddyConstructor, connAddress1, big.NewInt(1), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(contractCreation), sender, big.NewInt(2), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(noOpTx), sender, big.NewInt(3), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(contractCreation2), sender, big.NewInt(4), big.NewInt(0), chainTime),
		message.NewInboxMessage(buddyConstructor, connAddress2, big.NewInt(5), big.NewInt(0), laterChainTime),
	}

	logs, _, snap, _ := runAssertion(t, messages, 8, 1)
	results := processResults(t, logs)

	buddyConRes, ok := results[0].(*evm.TxResult)
	if !ok {
		t.Fatal("expected tx res")
	}

	contractConRes, ok := results[1].(*evm.TxResult)
	if !ok {
		t.Fatal("expected tx res")
	}

	noOpRes, ok := results[2].(*evm.TxResult)
	if !ok {
		t.Fatal("expected tx res")
	}

	contractCon2Res, ok := results[3].(*evm.TxResult)
	if !ok {
		t.Fatal("expected tx res")
	}

	buddySendRes, ok := results[4].(*evm.SendResult)
	if !ok {
		t.Fatal("expected send res")
	}

	buddyCon2Res, ok := results[7].(*evm.TxResult)
	if !ok {
		t.Fatal("expected tx res")
	}

	checkConstructorResult(t, buddyConRes, connAddress1)

	if buddySendRes.BatchNumber.Cmp(big.NewInt(0)) != 0 {
		t.Error("unexpected batch num")
	}

	if buddySendRes.BatchIndex.Cmp(big.NewInt(0)) != 0 {
		t.Error("unexpected batch index")
	}

	buddyRes, err := evm.NewBuddyResultFromData(buddySendRes.Data)
	failIfError(t, err)
	if buddyRes.Address != connAddress1 {
		t.Error("Buddy contract created at wrong address")
	}

	txResultCheck(t, contractConRes, evm.ContractAlreadyExists)
	succeededTxCheck(t, noOpRes)
	checkConstructorResult(t, contractCon2Res, connAddress2)
	txResultCheck(t, buddyCon2Res, evm.ContractAlreadyExists)

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
