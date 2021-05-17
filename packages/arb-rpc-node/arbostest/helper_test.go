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

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

const printArbOSLog = false

func initMsg(t *testing.T, options []message.ChainConfigOption) message.Init {
	params := protocol.ChainParams{
		StakeRequirement:          big.NewInt(0),
		StakeToken:                common.Address{},
		GracePeriod:               common.NewTimeBlocks(big.NewInt(3)),
		MaxExecutionSteps:         0,
		ArbGasSpeedLimitPerSecond: 1000000000,
	}
	init, err := message.NewInitMessage(params, owner, options)
	test.FailIfError(t, err)
	return init
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

func makeEthDeposit(dest common.Address, amount *big.Int) message.EthDepositTx {
	return message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: dest,
				Payment:     amount,
				Data:        nil,
			},
		}),
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
		if res, ok := res.(*evm.BlockInfo); ok {
			if res.GasLimit().Cmp(big.NewInt(100000000000)) > 0 {
				t.Error("block gas limit too high", res.GasLimit())
			}
		}
		results = append(results, res)
	}
	return results
}

func processDebugPrints(t *testing.T, debugPrints []value.Value) []evm.EVMLogLine {
	t.Helper()
	results := make([]evm.EVMLogLine, 0, len(debugPrints))
	for _, debugPrint := range debugPrints {
		res, err := evm.NewLogLineFromValue(debugPrint)
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

func extractTxResults(t *testing.T, results []evm.Result) []*evm.TxResult {
	t.Helper()
	txResults := make([]*evm.TxResult, 0, len(results))
	for _, res := range results {
		txRes, ok := res.(*evm.TxResult)
		if !ok {
			continue
		}
		txResults = append(txResults, txRes)
	}
	return txResults
}

func txResultCheck(t *testing.T, res *evm.TxResult, correct evm.ResultType) {
	t.Helper()
	if res.ResultCode != correct {
		t.Log("result", res)
		nested, err := message.NestedMessage(res.IncomingRequest.Data, res.IncomingRequest.Kind)
		if err != nil {
			t.Log("Invalid nested", err)
		} else {
			t.Log("Nested:", nested)
		}
		t.Log("data", hexutil.Encode(res.ReturnData))
		t.Fatal("unexpected result", res.ResultCode, "instead of", correct)
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

func extractIncomingMessages(t *testing.T, results []*evm.TxResult) []message.Message {
	t.Helper()
	var messages []message.Message
	for _, res := range results {
		incoming, err := message.NestedMessage(res.IncomingRequest.Data, res.IncomingRequest.Kind)
		test.FailIfError(t, err)
		messages = append(messages, incoming)
	}
	return messages
}

func filterL2Messages(t *testing.T, messages []message.Message) []message.AbstractL2Message {
	var l2Messages []message.AbstractL2Message
	for _, msg := range messages {
		nested, ok := msg.(message.L2Message)
		if !ok {
			continue
		}
		abs, err := nested.AbstractMessage()
		test.FailIfError(t, err)
		l2Messages = append(l2Messages, abs)
	}
	return l2Messages
}

func failIfError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func runSimpleTxAssertion(t *testing.T, messages []message.Message) ([]*evm.TxResult, *snapshot.Snapshot) {
	t.Helper()
	return runTxAssertion(t, makeSimpleInbox(t, messages))
}

func runTxAssertion(t *testing.T, messages []inbox.InboxMessage) ([]*evm.TxResult, *snapshot.Snapshot) {
	t.Helper()
	if len(messages) == 0 {
		t.Fatal("must have at least one message")
	}
	results, _, snap := runTxAssertionWithCount(t, messages, len(messages)-1)
	return results, snap
}

func runTxAssertionWithCount(t *testing.T, messages []inbox.InboxMessage, logCount int) ([]*evm.TxResult, [][]evm.EVMLogLine, *snapshot.Snapshot) {
	t.Helper()
	results, sends, debugPrints, snap := runBasicAssertion(t, messages)
	if len(sends) != 0 {
		t.Fatal("expected no sends", len(sends))
	}
	txResults := extractTxResults(t, results)
	if len(txResults) != logCount {
		t.Fatal("unexpected log count ", len(txResults), "instead of", logCount)
	}
	return txResults, debugPrints, snap
}

func runAssertion(t *testing.T, inboxMessages []inbox.InboxMessage, logCount int, sendCount int) ([]evm.Result, [][]byte, [][]evm.EVMLogLine, *snapshot.Snapshot) {
	t.Helper()
	results, sends, debugPrints, snap := runBasicAssertion(t, inboxMessages)
	if logCount != math.MaxInt32 && len(results) != logCount+1 {
		t.Fatal("unexpected log count ", len(results), "instead of", logCount+1)
	}

	if len(sends) != sendCount {
		t.Fatal("unxpected send count ", len(sends), "instead of", sendCount)
	}
	return results, sends, debugPrints, snap
}

func runBasicAssertion(t *testing.T, inboxMessages []inbox.InboxMessage) ([]evm.Result, [][]byte, [][]evm.EVMLogLine, *snapshot.Snapshot) {
	t.Helper()
	if inboxMessages[0].Kind != message.InitType {
		t.Fatal("inbox must start with init message")
	}
	mach, err := cmachine.New(*arbosfile)
	failIfError(t, err)

	var logs []value.Value
	var sends [][]byte
	var debugPrints [][]evm.EVMLogLine
	assertion, _, _, err := mach.ExecuteAssertion(10000000000, false, nil)
	failIfError(t, err)
	logs = append(logs, assertion.Logs...)
	sends = append(sends, assertion.Sends...)
	for i, msg := range inboxMessages {
		t.Log("Message", i)
		assertion, dPrints, _, err := mach.ExecuteAssertion(10000000000, false, []inbox.InboxMessage{msg})
		failIfError(t, err)
		parsedDebugPrints := processDebugPrints(t, dPrints)
		for _, d := range parsedDebugPrints {
			t.Log("debugprint", d)
		}
		logs = append(logs, assertion.Logs...)
		sends = append(sends, assertion.Sends...)

		debugPrints = append(debugPrints, parsedDebugPrints)

		if len(assertion.Logs) != 1 {
			continue
		}
		res, err := evm.NewTxResultFromValue(assertion.Logs[0])
		if err != nil {
			continue
		}
		avmGasFactor := big.NewInt(100)
		avmGas := new(big.Int).Mul(res.FeeStats.UnitsUsed.L2Computation, avmGasFactor)
		uncountedComputation := new(big.Int).Sub(new(big.Int).SetUint64(assertion.NumGas), avmGas)
		chargeRatio := new(big.Rat).SetFrac(avmGas, new(big.Int).SetUint64(assertion.NumGas))
		// Note: These ratio's were set based on measurements to prevent any regressions
		// If in the future arbos tries to provide a stronger bound on unmetered computation, this can be adjusted
		if arbosVersion >= 8 && chargeRatio.Cmp(big.NewRat(7, 10)) < 0 && uncountedComputation.Cmp(big.NewInt(300000)) > 0 {
			t.Errorf("didn't charge enough for tx %v=%v (%v uncharged)", chargeRatio, chargeRatio.FloatString(2), uncountedComputation)
		}
	}

	var snap *snapshot.Snapshot
	if len(inboxMessages) > 0 {
		lastMessage := inboxMessages[len(inboxMessages)-1]
		seq := new(big.Int).Add(lastMessage.InboxSeqNum, big.NewInt(1))
		msg := message.NewInboxMessage(
			message.EndBlockMessage{},
			common.Address{},
			seq,
			big.NewInt(0),
			inbox.ChainTime{
				BlockNum:  common.NewTimeBlocksInt(0),
				Timestamp: big.NewInt(0),
			},
		)
		_, _, _, err = mach.ExecuteAssertionAdvanced(10000000000, false, []inbox.InboxMessage{msg}, nil, true, common.Hash{}, common.Hash{})
		test.FailIfError(t, err)
		snap, err = snapshot.NewSnapshot(mach.Clone(), lastMessage.ChainTime, message.ChainAddressToID(chain), seq)
		test.FailIfError(t, err)
	}
	if printArbOSLog {
		testCase, err := inbox.TestVectorJSON(inboxMessages, logs, sends)
		failIfError(t, err)
		t.Log(string(testCase))
	}
	return processResults(t, logs), sends, debugPrints, snap
}

type InboxBuilder struct {
	Messages []inbox.InboxMessage
}

func (ib *InboxBuilder) AddMessage(msg message.Message, sender common.Address, gasPrice *big.Int, time inbox.ChainTime) {
	newMsg := message.NewInboxMessage(msg, sender, big.NewInt(int64(len(ib.Messages))), gasPrice, time)
	ib.Messages = append(ib.Messages, newMsg)
}

func makeSimpleInbox(t *testing.T, messages []message.Message) []inbox.InboxMessage {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	ib := &InboxBuilder{}
	ib.AddMessage(initMsg(t, nil), chain, big.NewInt(0), chainTime)
	for _, msg := range messages {
		ib.AddMessage(msg, sender, big.NewInt(0), chainTime)
	}
	return ib.Messages
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
