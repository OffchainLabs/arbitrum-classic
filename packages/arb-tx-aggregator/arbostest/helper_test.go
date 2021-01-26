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
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbosmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
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
		Owner:       owner,
		ExtraConfig: []byte{},
	}
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

func runAssertion(t *testing.T, inboxMessages []inbox.InboxMessage, logCount int, sendCount int) ([]value.Value, [][]byte, machine.Machine, *protocol.ExecutionAssertion) {
	t.Helper()
	cmach, err := cmachine.New(arbos.Path())
	failIfError(t, err)
	mach := arbosmachine.New(cmach)
	assertion, _, _ := mach.ExecuteAssertion(10000000000, true, inboxMessages, true)
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

	return logs, sends, mach, assertion
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

func makeFuncData(t *testing.T, methodABI abi.Method, args ...interface{}) []byte {
	t.Helper()
	methodData, err := methodABI.Inputs.Pack(args...)
	failIfError(t, err)
	return append(methodABI.ID, methodData...)
}
