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
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestTrace(t *testing.T) {
	constructorData := hexutil.MustDecode(arbostestcontracts.SimpleBin)
	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	failIfError(t, err)

	tx1 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{0},
		Payment:     big.NewInt(0),
		Data:        constructorData,
	}

	tx2 := message.Transaction{
		MaxGas:      big.NewInt(100000),
		GasPriceBid: big.NewInt(1),
		SequenceNum: big.NewInt(1),
		DestAddress: connAddress1,
		Payment:     big.NewInt(200),
		Data:        makeFuncData(t, simpleABI.Methods["trace"], big.NewInt(42356)),
	}

	messages := []message.Message{
		makeEthDeposit(sender, big.NewInt(10000000)),
		message.NewSafeL2Message(tx1),
		message.NewSafeL2Message(tx2),
	}
	inboxMessages := makeSimpleInbox(t, messages)

	results, debugPrintsLists, _ := runTxAssertionWithCount(t, inboxMessages, len(messages))

	allResultsSucceeded(t, results)

	checkConstructorResult(t, results[1], connAddress1)

	for i, debugPrints := range debugPrintsLists[1:] {
		trace, err := evm.GetTraceFromLogLines(debugPrints)
		test.FailIfError(t, err)
		t.Log("Trace", i, "of length", len(trace.Items))
		topFrame, err := trace.FrameTree()
		test.FailIfError(t, err)
		fmt.Println("frame", topFrame)
		depthCount := 0
		for i, item := range trace.Items {
			if _, ok := item.(*evm.CallTrace); ok {
				depthCount++
			}
			if _, ok := item.(*evm.ReturnTrace); ok {
				if depthCount == 0 {
					t.Fatal("can only return from inside call")
				}
				depthCount--
			}
			if _, ok := item.(*evm.CreateTrace); ok {
				if _, ok := trace.Items[i+1].(*evm.CallTrace); !ok {
					t.Fatal("call must come after create")
				}
			}
			if _, ok := item.(*evm.Create2Trace); ok {
				if _, ok := trace.Items[i+1].(*evm.CallTrace); !ok {
					t.Fatal("call must come after create2")
				}
			}
		}
		if depthCount != 0 {
			t.Fatal("must end at depth 0")
		}
	}

}
