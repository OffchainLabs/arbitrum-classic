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

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestWithdrawEth(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	addr := common.RandAddress()
	chain := common.RandAddress()

	depositMsg := message.Eth{
		Dest:  addr,
		Value: big.NewInt(100),
	}

	withdrawValue := big.NewInt(100)
	withdrawDest := common.RandAddress()
	tx := withdrawEthTx(big.NewInt(0), withdrawValue, withdrawDest)

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(depositMsg, addr, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx), addr, big.NewInt(2), chainTime),
	}

	// Last parameter returned is number of steps executed
	assertion, _ := mach.ExecuteAssertion(10000000000, inboxMessages, 0)
	testCase, err := inbox.TestVectorJSON(inboxMessages, assertion.ParseLogs(), assertion.ParseOutMessages())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(testCase))
	logs := assertion.ParseLogs()

	if len(logs) != 1 {
		t.Fatal("unexpected log count", len(logs))
	}

	res, err := evm.NewTxResultFromValue(logs[0])
	if err != nil {
		t.Fatal(err)
	}
	if res.ResultCode != evm.ReturnCode {
		t.Fatal("incorrect tx response", res.ResultCode)
	}

	t.Log(res)

	sends := assertion.ParseOutMessages()
	if len(sends) != 1 {
		t.Fatal("unexpected send count")
	}

	outMsg, err := message.NewOutMessageFromValue(sends[0])
	if err != nil {
		t.Fatal(err)
	}

	if outMsg.Kind != message.EthType {
		t.Fatal("outgoing message had wrong type", outMsg.Kind)
	}

	if outMsg.Sender != addr {
		t.Fatal("wrong withdraw sender")
	}

	outEthMsg := message.NewEthFromData(outMsg.Data)

	if outEthMsg.Value.Cmp(withdrawValue) != 0 {
		t.Fatal("wrong withdraw value", outEthMsg.Value)
	}

	if outEthMsg.Dest != withdrawDest {
		t.Fatal("wrong withdraw destination", outEthMsg.Dest)
	}
}
