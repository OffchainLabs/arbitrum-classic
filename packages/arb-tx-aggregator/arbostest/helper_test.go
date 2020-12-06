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
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/pkg/errors"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

func initMsg() message.Init {
	return message.Init{
		ChainParams: valprotocol.ChainParams{
			StakeRequirement:        big.NewInt(0),
			StakeToken:              common.Address{},
			GracePeriod:             common.TimeTicks{Val: big.NewInt(0)},
			MaxExecutionSteps:       0,
			ArbGasSpeedLimitPerTick: 0,
		},
		Owner:       common.Address{},
		ExtraConfig: []byte{},
	}
}

func runMessage(t *testing.T, mach machine.Machine, msg message.Message, sender common.Address) ([]*evm.TxResult, []message.OutMessage) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	assertion, steps := mach.ExecuteAssertion(
		1000000000,
		[]inbox.InboxMessage{message.NewInboxMessage(msg, sender, big.NewInt(0), chainTime)},
		0,
	)
	//data, err := value.TestVectorJSON(inbox, assertion.ParseLogs(), assertion.ParseOutMessages())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(string(data))
	t.Log("Ran assertion for", steps, "steps and had", assertion.LogsCount, "logs and", assertion.OutMsgsCount, "messages")
	if mach.CurrentStatus() != machine.Extensive {
		t.Fatal("machine should still be working")
	}
	blockReason := mach.IsBlocked(false)
	if blockReason == nil {
		t.Fatal("machine not blocked")
	}

	if _, ok := blockReason.(machine.InboxBlocked); !ok {
		t.Fatal("Machine blocked for weird reason", blockReason)
	}

	results := processTxResults(t, assertion.ParseLogs())
	sends := make([]message.OutMessage, 0)
	for _, send := range assertion.ParseOutMessages() {
		msg, err := message.NewOutMessageFromValue(send)
		failIfError(t, err)
		sends = append(sends, msg)
	}
	return results, sends
}

func runValidTransaction(t *testing.T, mach machine.Machine, msg message.AbstractL2Message, sender common.Address) *evm.TxResult {
	t.Helper()
	result := runTransaction(t, mach, msg, sender)
	succeededTxCheck(t, result)
	return result
}

func runTransaction(t *testing.T, mach machine.Machine, msg message.AbstractL2Message, sender common.Address) *evm.TxResult {
	t.Helper()
	l2, err := message.NewL2Message(msg)
	failIfError(t, err)
	results, sends := runMessage(t, mach, l2, sender)
	if len(results) != 1 {
		t.Fatalf("unexpected log count %v", len(results))
	}
	if len(sends) != 0 {
		t.Fatalf("unexpected send count %v", len(sends))
	}
	return results[0]
}

func withdrawEthTx(sequenceNum *big.Int, amount *big.Int, dest common.Address) message.Transaction {
	return message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
		Payment:     amount,
		Data:        snapshot.WithdrawEthData(dest),
	}
}

func withdrawERC20Tx(sequenceNum *big.Int, amount *big.Int, dest common.Address) message.Transaction {
	return message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        snapshot.WithdrawERC20Data(dest, amount),
	}
}

func withdrawERC721Tx(sequenceNum *big.Int, id *big.Int, dest common.Address) message.Transaction {
	return message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
		Payment:     big.NewInt(0),
		Data:        snapshot.WithdrawERC721Data(dest, id),
	}
}

func makeConstructorTx(code []byte, sequenceNum *big.Int, payment *big.Int) message.Transaction {
	if payment == nil {
		payment = big.NewInt(0)
	}
	return message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: sequenceNum,
		DestAddress: common.Address{},
		Payment:     payment,
		Data:        code,
	}
}

func deployContract(t *testing.T, mach machine.Machine, sender common.Address, code []byte, sequenceNum *big.Int, payment *big.Int) (common.Address, error) {
	constructorTx := makeConstructorTx(code, sequenceNum, payment)
	constructorResult := runValidTransaction(t, mach, constructorTx, sender)
	return getConstructorResult(constructorResult)
}

func getConstructorResult(constructorResult *evm.TxResult) (common.Address, error) {
	if len(constructorResult.ReturnData) != 32 {
		return common.Address{}, errors.New("unexpected constructor result length")
	}
	var contractAddress common.Address
	copy(contractAddress[:], constructorResult.ReturnData[12:])
	return contractAddress, nil
}

func checkConstructorResult(t *testing.T, res *evm.TxResult, correctAddress common.Address) {
	t.Helper()
	succeededTxCheck(t, res)
	connAddrCalc, err := getConstructorResult(res)
	failIfError(t, err)
	if connAddrCalc != correctAddress {
		t.Fatal("constructed address doesn't match:", connAddrCalc, "instead of", correctAddress)
	}
}

func depositEth(t *testing.T, mach machine.Machine, dest common.Address, amount *big.Int) {
	msg := message.Eth{
		Dest:  dest,
		Value: amount,
	}

	depositResults, sendResults := runMessage(t, mach, msg, dest)
	if len(depositResults) != 0 {
		t.Fatal("deposit should not have had a result")
	}
	if len(sendResults) != 0 {
		t.Fatal("deposit should not trigger sends")
	}
}

func processTxResults(t *testing.T, logs []value.Value) []*evm.TxResult {
	t.Helper()
	results := make([]*evm.TxResult, 0, len(logs))
	for _, avmLog := range logs {
		res, err := evm.NewTxResultFromValue(avmLog)
		failIfError(t, err)
		results = append(results, res)
	}
	return results
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
	for _, res := range results {
		succeededTxCheck(t, res)
	}
}

func failIfError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func runAssertion(t *testing.T, inboxMessages []inbox.InboxMessage, logCount int, sendCount int) ([]value.Value, []value.Value, machine.Machine) {
	t.Helper()
	mach, err := cmachine.New(arbos.Path())
	failIfError(t, err)
	assertion, _ := mach.ExecuteAssertion(10000000000, inboxMessages, 0)
	logs := assertion.ParseLogs()
	sends := assertion.ParseOutMessages()
	testCase, err := inbox.TestVectorJSON(inboxMessages, logs, sends)
	failIfError(t, err)
	t.Log(string(testCase))

	if len(logs) != logCount {
		t.Fatal("unexpected log count ", len(logs))
	}

	if len(sends) != sendCount {
		t.Fatal("unxpected send count ", len(sends))
	}

	return logs, sends, mach
}

func makeSimpleInbox(messages []message.Message) []inbox.InboxMessage {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	inboxMessages := make([]inbox.InboxMessage, 0)
	inboxMessages = append(inboxMessages, message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime))
	for i, msg := range messages {
		inboxMessages = append(inboxMessages, message.NewInboxMessage(msg, sender, big.NewInt(int64(1+i)), chainTime))
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
