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

package structures

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
)

var tester *ethbridgetestcontracts.RollupTester

func TestMainSetup(m *testing.T) {
	client, auths := test.SimulatedBackend()
	auth := auths[0]

	_, machineTx, deployedArbRollup, err := ethbridgetestcontracts.DeployRollupTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	client.Commit()
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

	tester = deployedArbRollup
}

func TestGenerateLastMessageHash(t *testing.T) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	node := NewInitialNode(mach.Clone(), common.Hash{})

	results := make([]*evm.Result, 0, 10)
	for i := int32(0); i < 5; i++ {
		stop := evm.NewRandomResult(message.NewRandomEth(), 2)
		results = append(results, stop)
	}

	nextNode := NewRandomNodeFromValidPrev(node, results)
	assert := nextNode.Assertion()
	expectedHash := nextNode.Disputable().AssertionClaim.AssertionStub.LastMessageHash

	ethbridgeHash, _, err := tester.GenerateLastMessageHash(
		nil,
		assert.OutMsgsData,
		big.NewInt(0),
		big.NewInt(int64(len(assert.OutMsgsData))))
	if err != nil {
		t.Fatal(err)
	}

	if expectedHash != ethbridgeHash {
		t.Error(errors.New("calculated wrong last message hash"))
		fmt.Println(expectedHash)
		fmt.Println(ethbridgeHash)
	}
}

func TestCalculateLeafFromPath(t *testing.T) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	node := NewInitialNode(mach.Clone(), common.Hash{})

	results := make([]*evm.Result, 0, 10)
	for i := int32(0); i < 5; i++ {
		stop := evm.NewRandomResult(message.NewRandomEth(), 2)
		results = append(results, stop)
	}

	nextNode := NewRandomNodeFromValidPrev(node, results)
	path := GeneratePathProof(node, nextNode)

	bridgeHash, err := tester.CalculateLeafFromPath(nil, node.Hash(), common.HashSliceToRaw(path))
	if nextNode.Hash().ToEthHash() != bridgeHash {
		fmt.Println(bridgeHash)
		fmt.Println(nextNode.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}

func TestChildNodeHash(t *testing.T) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	node := NewInitialNode(mach.Clone(), common.Hash{})

	results := make([]*evm.Result, 0, 10)
	for i := int32(0); i < 7; i++ {
		stop := evm.NewRandomResult(message.NewRandomEth(), 2)
		results = append(results, stop)
	}

	nextNode := NewRandomNodeFromValidPrev(node, results)

	bridgeHash, err := tester.ChildNodeHash(
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
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	node := NewInitialNode(mach.Clone(), common.Hash{})

	results := make([]*evm.Result, 0, 10)
	for i := int32(0); i < 8; i++ {
		stop := evm.NewRandomResult(message.NewRandomEth(), 2)
		results = append(results, stop)
	}

	nextNode := NewRandomNodeFromValidPrev(node, results)
	protoState := nextNode.VMProtoData()

	bridgeHash, err := tester.ComputeProtoHashBefore(
		nil,
		protoState.MachineHash,
		protoState.InboxTop,
		protoState.InboxCount,
		protoState.MessageCount,
		protoState.LogCount,
	)

	if protoState.Hash().ToEthHash() != bridgeHash {
		fmt.Println(bridgeHash)
		fmt.Println(protoState.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}
