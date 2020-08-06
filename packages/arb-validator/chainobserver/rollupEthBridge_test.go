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

package chainobserver

import (
	"context"
	"errors"
	"log"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

var tester *ethbridgetestcontracts.RollupTester

func TestMainSetup(m *testing.T) {
	client, pks := test.SimulatedBackend()
	auth := bind.NewKeyedTransactor(pks[0])

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

var contractPath = arbos.Path()

var dummyRollupAddress = common.Address{1}
var dummyAddress common.Address

func setUpChain(rollupAddress common.Address, checkpointType string, contractPath string) (*ChainObserver, error) {
	var checkpointer checkpointing.RollupCheckpointer
	switch checkpointType {
	case "dummy":
		checkpointer = NewDummyCheckpointer()
	case "fresh_rocksdb":
		var err error
		checkpointer, err = checkpointing.NewIndexedCheckpointer(rollupAddress, "", big.NewInt(1000000), true)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid checkpoint type")
	}
	if err := checkpointer.Initialize(contractPath); err != nil {
		return nil, err
	}
	chain, err := newChain(
		dummyAddress,
		checkpointer,
		valprotocol.ChainParams{
			StakeRequirement:        big.NewInt(1),
			GracePeriod:             common.TicksFromSeconds(60 * 60),
			MaxExecutionSteps:       1000000,
			ArbGasSpeedLimitPerTick: 1000,
		},
		&common.BlockId{
			Height:     common.NewTimeBlocks(big.NewInt(10)),
			HeaderHash: common.Hash{},
		},
		common.Hash{},
	)
	if err != nil {
		return nil, err
	}
	chain.Inbox = &structures.Inbox{MessageStack: structures.NewRandomMessageStack(100)}
	chain.Start(context.Background())
	return chain, nil
}

func TestComputePrevLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	prepared, err := chain.prepareAssertion(chain.latestBlockId)
	if err != nil {
		t.Fatal(err)
	}

	prevData, err := tester.ComputePrevLeaf(
		nil,
		prepared.GetAssertionParams(),
		prepared.GetAssertionParams2(),
		uint32(prepared.Prev.LinkType()),
		prepared.Params.NumSteps,
		prepared.AssertionStub.NumGas,
		prepared.Assertion.OutMsgsCount,
		prepared.Assertion.LogsCount,
	)
	if err != nil {
		t.Fatal(err)
	}

	if prepared.Prev.Hash().ToEthHash() != prevData.PrevLeaf {
		t.Error(prevData.PrevLeaf)
	}
}

func randomAssertion(t *testing.T, ms *structures.MessageStack, prevNode *structures.Node) (*protocol.ExecutionAssertion, *valprotocol.ExecutionAssertionStub) {
	logs := make([]value.Value, 0, 5)
	sends := make([]value.Value, 0)
	sends = append(sends, ethTransfer(common.Address{}, big.NewInt(75)))

	beforeInboxHash := prevNode.VMProtoData().InboxTop
	messages, err := ms.GetMessages(beforeInboxHash, 5)
	if err != nil {
		t.Fatal(err)
	}

	assertion := protocol.NewExecutionAssertionFromValues(
		prevNode.VMProtoData().MachineHash,
		common.RandHash(),
		rand.Uint64(),
		uint64(len(messages)),
		sends,
		logs,
	)
	return assertion, structures.NewExecutionAssertionStubFromWholeAssertion(assertion, beforeInboxHash, ms)
}

func TestGenerateInvalidInboxLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	prevNode := chain.NodeGraph.LatestConfirmed()
	assertion, assertionStub := randomAssertion(t, chain.Inbox.MessageStack, prevNode)
	newNode := structures.NewRandomInvalidNodeFromValidPrev(prevNode, assertionStub, assertion, valprotocol.InvalidInboxTopChildType, chain.GetChainParams())

	prepared, err := chain.prepareAssertion(chain.latestBlockId)
	if err != nil {
		t.Fatal(err)
	}
	prepared.Assertion = assertion
	prepared.AssertionStub = assertionStub

	prevData, err := tester.ComputePrevLeaf(
		nil,
		prepared.GetAssertionParams(),
		prepared.GetAssertionParams2(),
		uint32(prepared.Prev.LinkType()),
		prepared.Params.NumSteps,
		prepared.AssertionStub.NumGas,
		prepared.Assertion.OutMsgsCount,
		prepared.Assertion.LogsCount,
	)
	if err != nil {
		t.Fatal(err)
	}

	if newNode.PrevHash().ToEthHash() != prevData.PrevLeaf {
		t.Error("invalid prev leaf", hexutil.Encode(prevData.PrevLeaf[:]), "instead of", newNode.PrevHash())
	}

	invalidInboxHash, err := tester.ChildNodeHash(
		nil,
		newNode.PrevHash(),
		newNode.Deadline().Val,
		newNode.NodeDataHash(),
		new(big.Int).SetUint64(uint64(valprotocol.InvalidInboxTopChildType)),
		newNode.VMProtoData().Hash())

	if newNode.Hash().ToEthHash() != invalidInboxHash {
		t.Log(invalidInboxHash)
		t.Log(newNode.Hash())
		t.Error("incorrect child node hash")
	}
}

func TestGenerateInvalidExecutionLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	prevNode := chain.NodeGraph.LatestConfirmed()
	assertion, assertionStub := randomAssertion(t, chain.Inbox.MessageStack, prevNode)
	newNode := structures.NewRandomInvalidNodeFromValidPrev(prevNode, assertionStub, assertion, valprotocol.InvalidExecutionChildType, chain.GetChainParams())

	prepared, err := chain.prepareAssertion(chain.latestBlockId)
	if err != nil {
		t.Fatal(err)
	}
	prepared.Assertion = assertion
	prepared.AssertionStub = assertionStub

	prevData, err := tester.ComputePrevLeaf(
		nil,
		prepared.GetAssertionParams(),
		prepared.GetAssertionParams2(),
		uint32(prepared.Prev.LinkType()),
		prepared.Params.NumSteps,
		prepared.AssertionStub.NumGas,
		prepared.Assertion.OutMsgsCount,
		prepared.Assertion.LogsCount,
	)
	if err != nil {
		t.Fatal(err)
	}

	if newNode.PrevHash().ToEthHash() != prevData.PrevLeaf {
		t.Error("incorrect prev leaf hash", hexutil.Encode(prevData.PrevLeaf[:]))
	}

	invalidExecutionHash, err := tester.ChildNodeHash(
		nil,
		newNode.PrevHash(),
		newNode.Deadline().Val,
		newNode.NodeDataHash(),
		new(big.Int).SetUint64(uint64(valprotocol.InvalidExecutionChildType)),
		newNode.VMProtoData().Hash())

	if newNode.Hash().ToEthHash() != invalidExecutionHash {
		t.Log(invalidExecutionHash)
		t.Log(newNode.Hash())
		t.Error("invalid prev leaf")
	}
}
