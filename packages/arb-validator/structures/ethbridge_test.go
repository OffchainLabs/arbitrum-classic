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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"log"
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
)

var tester *ethbridgetestcontracts.RollupTester

func TestMainSetup(m *testing.T) {
	backend, pks := test.SimulatedBackend()
	ctx := context.Background()
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := bind.NewKeyedTransactor(pks[0])
	authClient, err := ethbridge.NewEthAuthClient(ctx, client, auth)
	if err != nil {
		log.Fatal(err)
	}

	rollupAddr, machineTx, err := authClient.MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		return ethbridgetestcontracts.DeployRollupTester(auth, client)
	})

	client.Commit()

	_, err = ethbridge.WaitForReceiptWithResults(
		ctx,
		client,
		auth.From,
		machineTx,
		"deployedMachineTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	tester, err = ethbridgetestcontracts.NewRollupTester(rollupAddr, client)
}

func TestGenerateLastMessageHash(t *testing.T) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	node := NewInitialNode(mach.Clone())
	ms := NewRandomMessageStack(20)
	nextNode := NewRandomNodeFromValidPrev(node, ms, 5)
	assert := nextNode.Assertion()
	expectedHash := nextNode.Disputable().Assertion.LastMessageHash

	ethbridgeHash, _, err := tester.GenerateLastMessageHash(
		nil,
		assert.OutMsgsData,
		big.NewInt(0),
		new(big.Int).SetUint64(assert.OutMsgsCount))
	if err != nil {
		t.Fatal(err)
	}

	if expectedHash != ethbridgeHash {
		t.Error(errors.New("calculated wrong last l2message hash"))
		fmt.Println(expectedHash)
		fmt.Println(ethbridgeHash)
	}
}

func TestCalculateLeafFromPath(t *testing.T) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	node := NewInitialNode(mach.Clone())
	ms := NewRandomMessageStack(20)
	nextNode := NewRandomNodeFromValidPrev(node, ms, 5)
	path := GeneratePathProof(node, nextNode)

	bridgeHash, err := tester.CalculateLeafFromPath(nil, node.Hash(), common.HashSliceToRaw(path))
	if err != nil {
		t.Fatal(err)
	}
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

	node := NewInitialNode(mach.Clone())
	ms := NewRandomMessageStack(20)
	nextNode := NewRandomNodeFromValidPrev(node, ms, 5)

	bridgeHash, err := tester.ChildNodeHash(
		nil,
		nextNode.PrevHash(),
		nextNode.Deadline().Val,
		nextNode.NodeDataHash(),
		new(big.Int).SetUint64(uint64(nextNode.LinkType())),
		nextNode.VMProtoData().Hash())
	if err != nil {
		t.Fatal(err)
	}

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

	node := NewInitialNode(mach.Clone())
	ms := NewRandomMessageStack(20)
	nextNode := NewRandomNodeFromValidPrev(node, ms, 5)
	protoState := nextNode.VMProtoData()

	bridgeHash, err := tester.ComputeProtoHashBefore(
		nil,
		protoState.MachineHash,
		protoState.InboxTop,
		protoState.InboxCount,
		protoState.MessageCount,
		protoState.LogCount,
	)
	if err != nil {
		t.Fatal(err)
	}

	if protoState.Hash().ToEthHash() != bridgeHash {
		fmt.Println(bridgeHash)
		fmt.Println(protoState.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}
