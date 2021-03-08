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
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbosmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func initMsg() message.Init {
	return message.Init{
		ChainParams: protocol.ChainParams{
			StakeRequirement:          big.NewInt(0),
			StakeToken:                common.Address{},
			GracePeriod:               common.NewTimeBlocks(big.NewInt(3)),
			MaxExecutionSteps:         0,
			ArbGasSpeedLimitPerSecond: 1000000000,
		},
		Owner:       owner,
		ExtraConfig: []byte{},
	}
}

func withdrawEthTx(sequenceNum *big.Int, amount *big.Int, dest common.Address) message.Transaction {
	return message.Transaction{
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
		Payment:     amount,
		Data:        arbos.WithdrawEthData(dest),
	}
}

func makeSimpleConstructorTx(code []byte, sequenceNum *big.Int) message.Transaction {
	return makeConstructorTx(code, sequenceNum, big.NewInt(0))
}

func makeConstructorTx(code []byte, sequenceNum *big.Int, payment *big.Int) message.Transaction {
	return message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.Address{},
		Payment:     payment,
		Data:        code,
	}
}

func checkConstructorResult(t *testing.T, res *evm.TxResult, correctAddress common.Address) {
	t.Helper()
	succeededTxCheck(t, res)

	if len(res.ReturnData) != 32 {
		t.Fatal("unexpected constructor result length")
	}
	var connAddrCalc common.Address
	copy(connAddrCalc[:], res.ReturnData[12:])
	if connAddrCalc != correctAddress {
		t.Fatal("constructed address doesn't match:", connAddrCalc, "instead of", correctAddress)
	}
}

func processResults(t *testing.T, logs []value.Value) []evm.Result {
	t.Helper()
	results := make([]evm.Result, 0, len(logs))
	for _, avmLog := range logs {
		res, err := evm.NewResultFromValue(avmLog)
		failIfError(t, err)
		results = append(results, res)
	}
	return results
}

func processTxResults(t *testing.T, logs []value.Value) []*evm.TxResult {
	t.Helper()
	results := processResults(t, logs)
	txResults := make([]*evm.TxResult, 0, len(results))
	for _, res := range results {
		txRes, ok := res.(*evm.TxResult)
		if !ok {
			t.Fatalf("expected result to be tx result but got %T", res)
		}
		txResults = append(txResults, txRes)
	}
	return txResults
}

func txResultCheck(t *testing.T, res *evm.TxResult, correct evm.ResultType) {
	t.Helper()
	if res.ResultCode != correct {
		t.Log("result", res)
		t.Log("data", hexutil.Encode(res.ReturnData))
		t.Fatal("unexpected result", res.ResultCode)
	}
}

func revertedTxCheck(t *testing.T, res *evm.TxResult) {
	t.Helper()
	txResultCheck(t, res, evm.RevertCode)
}

func succeededTxCheck(t *testing.T, res *evm.TxResult) {
	t.Helper()
	txResultCheck(t, res, evm.ReturnCode)
}

func allResultsSucceeded(t *testing.T, results []*evm.TxResult) {
	t.Helper()
	for i, res := range results {
		t.Log("Checking result", i)
		succeededTxCheck(t, res)
	}
}

func failIfError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func runAssertion(t *testing.T, inboxMessages []inbox.InboxMessage, logCount int, sendCount int) ([]value.Value, [][]byte, *snapshot.Snapshot, *protocol.ExecutionAssertion) {
	t.Helper()
	cmach, err := cmachine.New(arbos.Path())
	failIfError(t, err)
	mach := arbosmachine.New(cmach)

	assertion, _, _ := mach.ExecuteAssertion(10000000000, false, inboxMessages, false)
	testCase, err := inbox.TestVectorJSON(inboxMessages, assertion.Logs, assertion.Sends)
	failIfError(t, err)
	t.Log(string(testCase))

	if len(assertion.Logs) != logCount {
		t.Fatal("unexpected log count ", len(assertion.Logs), "instead of", logCount)
	}

	if len(assertion.Sends) != sendCount {
		t.Fatal("unxpected send count ", len(assertion.Sends), "instead of", sendCount)
	}

	var snap *snapshot.Snapshot
	if len(inboxMessages) > 0 {
		lastMessage := inboxMessages[len(inboxMessages)-1]
		seq := new(big.Int).Add(lastMessage.InboxSeqNum, big.NewInt(1))
		msg := message.NewInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), sender, seq, big.NewInt(0), lastMessage.ChainTime)
		mach.ExecuteAssertionAdvanced(10000000000, false, []inbox.InboxMessage{msg}, true, nil, true, common.Hash{}, common.Hash{})
		snap = snapshot.NewSnapshot(mach.Clone(), lastMessage.ChainTime, message.ChainAddressToID(chain), seq)
	}
	return assertion.Logs, assertion.Sends, snap, assertion
}

func makeSimpleInbox(messages []message.Message) []inbox.InboxMessage {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	inboxMessages := make([]inbox.InboxMessage, 0)
	inboxMessages = append(inboxMessages, message.NewInboxMessage(initMsg(), chain, big.NewInt(0), big.NewInt(0), chainTime))
	for i, msg := range messages {
		inboxMessages = append(inboxMessages, message.NewInboxMessage(msg, sender, big.NewInt(int64(1+i)), big.NewInt(0), chainTime))
	}
	return inboxMessages
}

func checkBalance(t *testing.T, snap *snapshot.Snapshot, account common.Address, balance *big.Int) {
	t.Helper()
	bal, err := snap.GetBalance(account)
	failIfError(t, err)
	if bal.Cmp(balance) != 0 {
		t.Error("unexpected balance", bal, "for account", account)
	}
}

func makeFuncData(t *testing.T, methodABI abi.Method, args ...interface{}) []byte {
	t.Helper()
	methodData, err := methodABI.Inputs.Pack(args...)
	failIfError(t, err)
	return append(methodABI.ID, methodData...)
}
