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
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func testWithdrawal(t *testing.T, depositMsg message.Message, withdrawalTx message.Transaction, withdrawalSender common.Address, logCount int) ([]*evm.TxResult, []message.OutMessage) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(depositMsg, common.RandAddress(), big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(withdrawalTx), withdrawalSender, big.NewInt(1), chainTime),
	}

	logs, sends, _, _ := runAssertion(t, inboxMessages, logCount, 1)
	results := processTxResults(t, logs)

	allResultsSucceeded(t, results)

	parsedSends := make([]message.OutMessage, 0, len(sends))
	for _, avmSend := range sends {
		outMsg, err := message.NewOutMessageFromValue(avmSend)
		failIfError(t, err)
		parsedSends = append(parsedSends, outMsg)
	}

	return results, parsedSends
}

func TestWithdrawEth(t *testing.T) {
	addr := common.RandAddress()

	depositMsg := message.Eth{
		Dest:  addr,
		Value: big.NewInt(100),
	}

	withdrawValue := big.NewInt(100)
	withdrawDest := common.RandAddress()
	tx := withdrawEthTx(big.NewInt(0), withdrawValue, withdrawDest)

	results, sends := testWithdrawal(t, depositMsg, tx, addr, 1)

	if len(results) != 1 {
		t.Fatal("unexpected log count", len(results))
	}

	res := results[0]

	if len(res.EVMLogs) != 1 {
		t.Fatal("unexpected log count")
	}
	ev, err := snapshot.ParseEthWithdrawalEvent(res.EVMLogs[0])
	failIfError(t, err)
	if ev.Amount.Cmp(withdrawValue) != 0 {
		t.Error("unexpected withdrawal value")
	}
	if ev.DestAddr != withdrawDest.ToEthAddress() {
		t.Error("unexpected dest address")
	}

	if len(sends) != 1 {
		t.Fatal("unexpected send count")
	}

	outMsg := sends[0]

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

func TestWithdrawERC20(t *testing.T) {
	addr := common.RandAddress()
	token := common.RandAddress()

	depositMsg := message.ERC20{
		Token: token,
		Dest:  addr,
		Value: big.NewInt(100),
	}

	withdrawValue := big.NewInt(100)
	withdrawDest := common.RandAddress()
	tx := withdrawERC20Tx(big.NewInt(1), withdrawValue, withdrawDest)

	results, sends := testWithdrawal(t, depositMsg, tx, token, 2)

	depositRes := results[0]
	t.Log("Deposit gas cost", depositRes.GasUsed)

	withdrawRes := results[1]

	if len(withdrawRes.EVMLogs) != 1 {
		t.Fatal("unexpected log count", len(withdrawRes.EVMLogs))
	}
	ev, err := snapshot.ParseERC20WithdrawalEvent(withdrawRes.EVMLogs[0])
	failIfError(t, err)
	if ev.Amount.Cmp(withdrawValue) != 0 {
		t.Error("unexpected withdrawal value")
	}
	if ev.DestAddr != withdrawDest.ToEthAddress() {
		t.Error("unexpected dest address")
	}
	if ev.TokenAddr != token.ToEthAddress() {
		t.Error("unexpected token address", ev.TokenAddr.Hex())
	}

	if len(sends) != 1 {
		t.Fatal("unexpected send count")
	}

	outMsg := sends[0]
	if outMsg.Kind != message.ERC20Type {
		t.Fatal("outgoing message had wrong type", outMsg.Kind)
	}

	// TODO: Update ArbOS to use actual sender
	emptyAddress := common.Address{}
	if outMsg.Sender != emptyAddress {
		t.Error("wrong withdraw sender")
	}

	outEthMsg := message.NewERC20FromData(outMsg.Data)

	if outEthMsg.Value.Cmp(withdrawValue) != 0 {
		t.Fatal("wrong withdraw value", outEthMsg.Value)
	}

	if outEthMsg.Dest != withdrawDest {
		t.Fatal("wrong withdraw destination", outEthMsg.Dest)
	}
}

func TestWithdrawERC721(t *testing.T) {
	addr := common.RandAddress()
	token := common.RandAddress()

	depositMsg := message.ERC721{
		Token: token,
		Dest:  addr,
		ID:    big.NewInt(100),
	}

	withdrawDest := common.RandAddress()
	tx := withdrawERC721Tx(big.NewInt(1), depositMsg.ID, withdrawDest)

	results, sends := testWithdrawal(t, depositMsg, tx, token, 2)

	if len(results) != 2 {
		t.Fatal("unexpected log count", len(results))
	}

	depositRes := results[0]
	t.Log("Deposit gas cost", depositRes.GasUsed)

	withdrawRes := results[1]

	if len(withdrawRes.EVMLogs) != 1 {
		t.Fatal("unexpected log count", len(withdrawRes.EVMLogs))
	}
	ev, err := snapshot.ParseERC721WithdrawalEvent(withdrawRes.EVMLogs[0])
	failIfError(t, err)
	if ev.Id.Cmp(depositMsg.ID) != 0 {
		t.Error("unexpected withdrawal id")
	}
	if ev.DestAddr != withdrawDest.ToEthAddress() {
		t.Error("unexpected dest address")
	}
	if ev.TokenAddr != token.ToEthAddress() {
		t.Error("unexpected token address", ev.TokenAddr.Hex())
	}

	if len(sends) != 1 {
		t.Fatal("unexpected send count")
	}

	outMsg := sends[0]
	if outMsg.Kind != message.ERC721Type {
		t.Fatal("outgoing message had wrong type", outMsg.Kind)
	}
	// TODO: Update ArbOS to use actual sender
	emptyAddress := common.Address{}
	if outMsg.Sender != emptyAddress {
		t.Error("wrong withdraw sender")
	}

	outERC721Msg := message.NewERC721FromData(outMsg.Data)

	if outERC721Msg.ID.Cmp(depositMsg.ID) != 0 {
		t.Fatal("wrong withdraw value", outERC721Msg.ID)
	}

	if outERC721Msg.Dest != withdrawDest {
		t.Fatal("wrong withdraw destination", outERC721Msg.Dest)
	}
}
