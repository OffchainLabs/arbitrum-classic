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
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func TestCreate2(t *testing.T) {
	backend, auths := test.SimulatedBackend(t)
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := auths[0]
	factoryConnAddress, _, cf, err := arbostestcontracts.DeployCloneFactory(auth, client)
	failIfError(t, err)

	simpleConnAddress, _, _, err := arbostestcontracts.DeploySimple(auth, client)
	failIfError(t, err)

	backend.Commit()

	tx, err := cf.Create2Clone(auth, simpleConnAddress, big.NewInt(0))
	failIfError(t, err)
	backend.Commit()

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	failIfError(t, err)

	ethEv, err := cf.ParseCreatedClone(*receipt.Logs[0])
	failIfError(t, err)

	cloneConnAddress := ethEv.Clone

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	factoryConstructorTx := makeSimpleConstructorTx(hexutil.MustDecode(arbostestcontracts.CloneFactoryBin), big.NewInt(0))
	simpleConstructorTx := makeSimpleConstructorTx(hexutil.MustDecode(arbostestcontracts.SimpleBin), big.NewInt(1))

	factoryABI, err := abi.JSON(strings.NewReader(arbostestcontracts.CloneFactoryABI))
	failIfError(t, err)
	create2Tx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: common.NewAddressFromEth(factoryConnAddress),
		Payment:     big.NewInt(0),
		Data:        makeFuncData(t, factoryABI.Methods["create2Clone"], simpleConnAddress, big.NewInt(0)),
	}

	simpleABI, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	failIfError(t, err)
	existsABI := simpleABI.Methods["exists"]
	existsCloneTx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(3),
		DestAddress: common.NewAddressFromEth(cloneConnAddress),
		Payment:     big.NewInt(0),
		Data:        existsABI.ID,
	}

	sender := common.NewAddressFromEth(auth.From)
	inboxMessages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(t, nil), chain, big.NewInt(0), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(factoryConstructorTx), sender, big.NewInt(1), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(simpleConstructorTx), sender, big.NewInt(2), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(create2Tx), sender, big.NewInt(3), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(existsCloneTx), sender, big.NewInt(4), big.NewInt(0), chainTime),
	}

	results, snap := runTxAssertion(t, inboxMessages)
	allResultsSucceeded(t, results)

	checkConstructorResult(t, results[0], common.NewAddressFromEth(factoryConnAddress))
	checkConstructorResult(t, results[1], common.NewAddressFromEth(simpleConnAddress))

	create2Res := results[2]
	if len(create2Res.EVMLogs) != 1 {
		t.Fatal("wrong EVM log count")
	}
	ev, err := cf.ParseCreatedClone(*create2Res.ToEthReceipt(common.Hash{}).Logs[0])
	failIfError(t, err)
	t.Log("ArbOS clone address:", ev.Clone.Hex())
	if ev.Clone != cloneConnAddress {
		t.Fatal("incorrect clone address")
	}

	existsCloneRes := results[3]
	existsCloneOutputs, err := existsABI.Outputs.UnpackValues(existsCloneRes.ReturnData)
	failIfError(t, err)
	if len(existsCloneOutputs) != 1 {
		t.Fatal("wrong output count")
	}
	if existsCloneOutputs[0].(*big.Int).Cmp(big.NewInt(10)) != 0 {
		t.Fatal("wrong exists clone output")
	}

	cloneCode, err := snap.GetCode(common.NewAddressFromEth(cloneConnAddress))
	failIfError(t, err)
	if len(cloneCode) != 45 {
		t.Fatal("wrong clone code length")
	}
}
