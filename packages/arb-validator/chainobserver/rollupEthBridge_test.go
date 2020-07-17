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
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"

	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"math/big"
	"testing"
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

	bridgeHash, _, err := tester.ComputePrevLeaf(
		nil,
		prepared.GetAssertionParams(),
		prepared.GetAssertionParams2(),
		uint32(prepared.Prev.LinkType()),
		prepared.Params.NumSteps,
		prepared.Claim.AssertionStub.DidInboxInsn,
		prepared.Claim.AssertionStub.NumGas,
		prepared.Assertion.OutMsgsCount,
		prepared.Assertion.LogsCount,
	)
	if err != nil {
		t.Fatal(err)
	}

	if prepared.Prev.Hash().ToEthHash() != bridgeHash {
		t.Error(bridgeHash)
	}
}

func randomAssertion() *protocol.ExecutionAssertion {
	results := make([]*evm.Result, 0, 5)
	messages := make([]value.Value, 0)
	messages = append(messages, message.NewInboxMessage(
		message.Eth{
			Dest:  common.Address{},
			Value: big.NewInt(75),
		},
		common.NewAddressFromEth(auth.From),
		big.NewInt(0),
		message.NewRandomChainTime(),
	).AsValue())
	for i := int32(0); i < 5; i++ {
		stop := evm.NewRandomResult(message.NewRandomEth(), 2)
		results = append(results, stop)
		messages = append(messages, message.NewRandomInboxMessage(message.NewRandomEth()).AsValue())
	}

	return evm.NewRandomEVMAssertion(results, messages)
}

func TestGenerateInvalidMsgLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	prevNode := chain.NodeGraph.LatestConfirmed()
	assertion := randomAssertion()

	newNode := structures.NewRandomInvalidNodeFromValidPrev(prevNode, assertion, valprotocol.InvalidMessagesChildType, chain.GetChainParams())

	prepared, err := chain.prepareAssertion(chain.latestBlockId)
	if err != nil {
		t.Fatal(err)
	}
	prepared.Assertion = assertion
	prepared.Claim.AssertionStub = valprotocol.NewExecutionAssertionStubFromAssertion(assertion)

	bridgeHash, _, err := tester.ComputePrevLeaf(
		nil,
		prepared.GetAssertionParams(),
		prepared.GetAssertionParams2(),
		uint32(prepared.Prev.LinkType()),
		prepared.Params.NumSteps,
		prepared.Claim.AssertionStub.DidInboxInsn,
		prepared.Claim.AssertionStub.NumGas,
		prepared.Assertion.OutMsgsCount,
		prepared.Assertion.LogsCount,
	)
	if err != nil {
		t.Fatal(err)
	}

	if newNode.PrevHash().ToEthHash() != bridgeHash {
		t.Error(bridgeHash)
	}

	invalidMsgHash, err := tester.ChildNodeHash(
		nil,
		newNode.PrevHash(),
		newNode.Deadline().Val,
		newNode.NodeDataHash(),
		new(big.Int).SetUint64(uint64(valprotocol.InvalidMessagesChildType)),
		newNode.VMProtoData().Hash())

	if newNode.Hash().ToEthHash() != invalidMsgHash {
		fmt.Println(bridgeHash)
		fmt.Println(newNode.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}

func TestGenerateInvalidInboxLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	prevNode := chain.NodeGraph.LatestConfirmed()
	assertion := randomAssertion()
	newNode := structures.NewRandomInvalidNodeFromValidPrev(prevNode, assertion, valprotocol.InvalidInboxTopChildType, chain.GetChainParams())

	prepared, err := chain.prepareAssertion(chain.latestBlockId)
	if err != nil {
		t.Fatal(err)
	}
	prepared.Assertion = assertion
	prepared.Claim.AssertionStub = valprotocol.NewExecutionAssertionStubFromAssertion(assertion)

	bridgeHash, _, err := tester.ComputePrevLeaf(
		nil,
		prepared.GetAssertionParams(),
		prepared.GetAssertionParams2(),
		uint32(prepared.Prev.LinkType()),
		prepared.Params.NumSteps,
		prepared.Claim.AssertionStub.DidInboxInsn,
		prepared.Claim.AssertionStub.NumGas,
		prepared.Assertion.OutMsgsCount,
		prepared.Assertion.LogsCount,
	)
	if err != nil {
		t.Fatal(err)
	}

	if newNode.PrevHash().ToEthHash() != bridgeHash {
		t.Error(bridgeHash)
	}

	invalidInboxHash, err := tester.ChildNodeHash(
		nil,
		newNode.PrevHash(),
		newNode.Deadline().Val,
		newNode.NodeDataHash(),
		new(big.Int).SetUint64(uint64(valprotocol.InvalidInboxTopChildType)),
		newNode.VMProtoData().Hash())

	if newNode.Hash().ToEthHash() != invalidInboxHash {
		fmt.Println(bridgeHash)
		fmt.Println(newNode.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}

func TestGenerateInvalidExecutionLeaf(t *testing.T) {
	chain, err := setUpChain(dummyRollupAddress, "dummy", contractPath)
	if err != nil {
		t.Fatal(err)
	}

	prevNode := chain.NodeGraph.LatestConfirmed()
	assertion := randomAssertion()
	newNode := structures.NewRandomInvalidNodeFromValidPrev(prevNode, assertion, valprotocol.InvalidExecutionChildType, chain.GetChainParams())

	prepared, err := chain.prepareAssertion(chain.latestBlockId)
	if err != nil {
		t.Fatal(err)
	}
	prepared.Assertion = assertion
	prepared.Claim.AssertionStub = valprotocol.NewExecutionAssertionStubFromAssertion(assertion)

	bridgeHash, _, err := tester.ComputePrevLeaf(
		nil,
		prepared.GetAssertionParams(),
		prepared.GetAssertionParams2(),
		uint32(prepared.Prev.LinkType()),
		prepared.Params.NumSteps,
		prepared.Claim.AssertionStub.DidInboxInsn,
		prepared.Claim.AssertionStub.NumGas,
		prepared.Assertion.OutMsgsCount,
		prepared.Assertion.LogsCount,
	)
	if err != nil {
		t.Fatal(err)
	}

	if newNode.PrevHash().ToEthHash() != bridgeHash {
		t.Error(bridgeHash)
	}

	invalidExecutionHash, err := tester.ChildNodeHash(
		nil,
		newNode.PrevHash(),
		newNode.Deadline().Val,
		newNode.NodeDataHash(),
		new(big.Int).SetUint64(uint64(valprotocol.InvalidExecutionChildType)),
		newNode.VMProtoData().Hash())

	if newNode.Hash().ToEthHash() != invalidExecutionHash {
		fmt.Println(bridgeHash)
		fmt.Println(newNode.Hash().ToEthHash())
		t.Error(bridgeHash)
	}
}
