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
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
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
		MaxGas:      big.NewInt(1000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, simpleABI.Methods["trace"], big.NewInt(42356)),
	}

	messages := []message.Message{
		makeEthDeposit(sender, big.NewInt(10000)),
		message.NewSafeL2Message(tx1),
		message.NewSafeL2Message(tx2),
	}
	inboxMessages := makeSimpleInbox(t, messages)

	results, _, _ := runTxAssertionWithCount(t, inboxMessages, len(messages))

	allResultsSucceeded(t, results)

	checkConstructorResult(t, results[1], connAddress1)
}

func TestConTrace(t *testing.T) {
	client, keys := test.SimulatedBackend(t)
	auth, err := bind.NewKeyedTransactorWithChainID(keys[0], big.NewInt(1337))
	test.FailIfError(t, err)
	_, _, simple, err := arbostestcontracts.DeploySimple(auth, client)
	test.FailIfError(t, err)
	client.Commit()

	_, err = simple.Trace(auth, big.NewInt(42356))
	test.FailIfError(t, err)
}
