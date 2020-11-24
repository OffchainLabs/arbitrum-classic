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
	"log"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestRevert(t *testing.T) {
	mach, err := cmachine.New(arbos.Path())
	if err != nil {
		t.Fatal(err)
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	chain := common.RandAddress()
	sender := common.HexToAddress("0x8c988ec54f112dd35666e19e7b0904bb12df1b6c")
	connAddr := common.HexToAddress("0x7cc1af94bfb4676c4facfc6a56430ec35c45b8b0")

	simpleConstructorTx := makeConstructorTx(hexutil.MustDecode(arbostestcontracts.SimpleBin), big.NewInt(0), nil)

	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	if err != nil {
		t.Fatal(err)
	}

	revertsTx := message.Transaction{
		MaxGas:      big.NewInt(1000000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: connAddr,
		Payment:     big.NewInt(0),
		Data:        simpleABI.Methods["reverts"].ID,
	}

	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(simpleConstructorTx), sender, big.NewInt(1), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(revertsTx), sender, big.NewInt(2), chainTime),
	}

	// Last parameter returned is number of tests executed
	assertion, _ := mach.ExecuteAssertion(10000000000, inboxMessages, 0)
	//testCase, err := inbox.TestVectorJSON(inboxMessages, assertion.ParseLogs(), assertion.ParseOutMessages())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//t.Log(string(testCase))
	logs := assertion.ParseLogs()
	sends := assertion.ParseOutMessages()

	results := make([]*evm.TxResult, 0, len(logs))
	for _, avmLog := range logs {
		res, err := evm.NewTxResultFromValue(avmLog)
		if err != nil {
			t.Fatal(err)
		}
		results = append(results, res)
	}

	if len(results) != 2 {
		log.Println("unxpected log count", len(results))
	}

	if len(sends) != 0 {
		log.Println("unxpected send count", len(sends))
	}

	checkConstructorResult(t, logs[0], connAddr)
	revertedTxCheck(t, results[1])

	correctResult := []byte("this is a test")
	if !bytes.Contains(results[1].ReturnData, correctResult) {
		t.Error("incorrect return data", hexutil.Encode(results[1].ReturnData))
	}
}
