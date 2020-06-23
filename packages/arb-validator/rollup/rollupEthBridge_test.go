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

package rollup

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/rolluptester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"math/big"
	"testing"
)

func TestCalculateLeafFromPath(t *testing.T) {
	rolluptester := getTester(t)

	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	node := structures.NewInitialNode(mach.Clone(), common.Hash{})

	results := make([]evm.Result, 0, 10)
	for i := int32(0); i < 5; i++ {
		stop := evm.NewRandomStop(message.NewRandomEth(), 2)
		results = append(results, stop)
	}

	nextNode := structures.NewRandomNodeFromValidPrev(node, results)
	path := structures.GeneratePathProof(node, nextNode)

	bridgeHash, err := rolluptester.CalculateLeafFromPath(nil, node.Hash(), common.HashSliceToRaw(path))
	if nextNode.Hash().ToEthHash() != bridgeHash {
		fmt.Println(bridgeHash)
		fmt.Println(nextNode.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}

func TestChildNodeHash(t *testing.T) {
	rolluptester := getTester(t)

	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	node := structures.NewInitialNode(mach.Clone(), common.Hash{})

	results := make([]evm.Result, 0, 10)
	for i := int32(0); i < 7; i++ {
		stop := evm.NewRandomStop(message.NewRandomEth(), 2)
		results = append(results, stop)
	}

	nextNode := structures.NewRandomNodeFromValidPrev(node, results)

	bridgeHash, err := rolluptester.ChildNodeHash(
		nil,
		nextNode.PrevHash(),
		nextNode.Deadline().Val,
		nextNode.NodeDataHash(),
		new(big.Int).SetUint64(uint64(nextNode.LinkType())),
		nextNode.VMProtoData().Hash())

	if nextNode.Hash().ToEthHash() != bridgeHash {
		fmt.Println(bridgeHash)
		fmt.Println(nextNode.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}

func TestProtoStateHash(t *testing.T) {
	rolluptester := getTester(t)

	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	node := structures.NewInitialNode(mach.Clone(), common.Hash{})

	results := make([]evm.Result, 0, 10)
	for i := int32(0); i < 8; i++ {
		stop := evm.NewRandomStop(message.NewRandomEth(), 2)
		results = append(results, stop)
	}

	nextNode := structures.NewRandomNodeFromValidPrev(node, results)
	protoState := nextNode.VMProtoData()

	bridgeHash, err := rolluptester.ComputeProtoHashBefore(
		nil,
		protoState.MachineHash,
		protoState.InboxTop,
		protoState.InboxCount)

	if protoState.Hash().ToEthHash() != bridgeHash {
		fmt.Println(bridgeHash)
		fmt.Println(protoState.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}

var dummyRollupAddress = common.Address{1}

func TestComputePrevLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	assertion := chain.prepareAssertion()
	rolluptester := getTester(t)

	bridgeHash, _, err := rolluptester.ComputePrevLeaf(
		nil,
		assertion.getAssertionParams(),
		assertion.beforeState.InboxCount,
		assertion.prev.Deadline().Val,
		uint32(assertion.prev.LinkType()),
		assertion.params.NumSteps,
		assertion.params.TimeBounds.AsIntArray(),
		assertion.params.ImportedMessageCount,
		assertion.claim.AssertionStub.DidInboxInsn,
		assertion.claim.AssertionStub.NumGas)
	if err != nil {
		t.Fatal(err)
	}

	if assertion.prev.Hash().ToEthHash() != bridgeHash {
		t.Error(bridgeHash)
	}
}

//func TestGenerateInvalidInboxTopLeaf(t *testing.T) {
//	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	assertion := chain.prepareAssertion()
//	rolluptester := getTester(t)
//}

var privHex = "27e926925fb5903ee038c894d9880f74d3dd6518e23ab5e5651de93327c7dffa"

func getTester(m *testing.T) *rolluptester.RollupTester {
	auth, err := test.SetupAuth(privHex)
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(test.GetEthUrl())
	if err != nil {
		log.Fatal(err)
	}

	_, machineTx, deployedArbRollup, err := rolluptester.DeployRollupTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		machineTx,
		"deployedMachineTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	return deployedArbRollup
}
