/*
* Copyright 2019-2020, Offchain Labs, Inc.
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
	"math/big"
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

var dummyAddress common.Address

func TestCreateEmptyChain(t *testing.T) {
	chain, _, err := setUpChain()
	if err != nil {
		t.Fatal(err)
	}
	if chain.nodeGraph.leaves.NumLeaves() != 1 {
		t.Fatal("unexpected leaf count")
	}
	tryMarshalUnmarshal(chain, t)
	cp := structures.NewRollupCheckpointer("dummy", 1000000, "contract.ao")
	tryMarshalUnmarshalWithCheckpointer(chain, cp, t)
}

func tryMarshalUnmarshal(chain *ChainObserver, t *testing.T) {
	ctx := structures.NewCheckpointContextImpl()
	chainBuf := chain.MarshalForCheckpoint(ctx)
	chain2 := chainBuf.UnmarshalFromCheckpoint(ctx, nil)
	if !chain.Equals(chain2) {
		t.Fail()
	}
}

func tryMarshalUnmarshalWithCheckpointer(chain *ChainObserver, cp *structures.RollupCheckpointer, t *testing.T) {
	blockHeight := big.NewInt(7337)
	ctx := structures.NewCheckpointContextImpl()
	buf, err := chain.MarshalToBytes(ctx)
	if err != nil {
		t.Fatal(err)
	}
	cp.SaveCheckpoint(blockHeight, buf, ctx)
	chain2, err := UnmarshalChainObserverFromBytes(buf, ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !chain.Equals(chain2) {
		t.Fail()
	}
}

func TestDoAssertion(t *testing.T) {
	chain, _, err := setUpChain()
	if err != nil {
		t.Fatal(err)
	}

	doAnAssertion(chain, chain.nodeGraph.latestConfirmed)
	validTip := chain.nodeGraph.latestConfirmed.GetSuccessor(chain.nodeGraph.NodeGraph, structures.ValidChildType)
	doAnAssertion(chain, validTip)
	if chain.nodeGraph.leaves.NumLeaves() != 7 {
		t.Fatal("unexpected leaf count")
	}

	tryMarshalUnmarshal(chain, t)
}

func TestChallenge(t *testing.T) {
	chain, _, err := setUpChain()
	if err != nil {
		t.Fatal(err)
	}

	doAnAssertion(chain, chain.nodeGraph.latestConfirmed)
	staker1addr := common.BytesToAddress([]byte{1})
	staker2addr := common.BytesToAddress([]byte{2})
	contractAddr := common.BytesToAddress([]byte{3})
	validTip := chain.nodeGraph.latestConfirmed.GetSuccessor(chain.nodeGraph.NodeGraph, structures.ValidChildType)
	tip2 := chain.nodeGraph.latestConfirmed.GetSuccessor(chain.nodeGraph.NodeGraph, structures.InvalidMessagesChildType)
	n1, _, childType, err := chain.nodeGraph.GetConflictAncestor(validTip, tip2)
	if err != nil {
		t.Fatal(err)
	}
	confNode := n1.prev
	if !confNode.Equals(chain.nodeGraph.latestConfirmed) {
		t.Fatal("unexpected value for conflict ancestor")
	}
	if childType != structures.InvalidMessagesChildType {
		t.Fatal("unexpected value for conflict type")
	}

	createOneStaker(chain, staker1addr, validTip.hash)
	createOneStaker(chain, staker2addr, tip2.hash)

	chain.nodeGraph.NewChallenge(contractAddr, staker1addr, staker2addr, structures.InvalidMessagesChildType)

	tryMarshalUnmarshal(chain, t)

	chain.nodeGraph.ChallengeResolved(contractAddr, staker1addr, staker2addr)

	tryMarshalUnmarshal(chain, t)
}

func doAnAssertion(chain *ChainObserver, baseNode *Node) {
	theMachine := baseNode.machine
	timeBounds := &protocol.TimeBoundsBlocks{
		Start: &protocol.TimeBlocksBuf{Val: utils.MarshalInt64ToBigIntBuf(0)},
		End:   &protocol.TimeBlocksBuf{Val: utils.MarshalInt64ToBigIntBuf(1000)},
	}
	execAssertion, numGas := theMachine.ExecuteAssertion(1, timeBounds, value.NewEmptyTuple())
	_ = execAssertion

	assertionParams := &structures.AssertionParams{
		NumSteps:             1,
		TimeBounds:           timeBounds,
		ImportedMessageCount: big.NewInt(0),
	}
	assertionStub := &protocol.ExecutionAssertionStub{
		AfterHash:        value.NewHashBuf(theMachine.Hash()),
		DidInboxInsn:     false,
		NumGas:           uint64(numGas),
		FirstMessageHash: value.NewHashBuf([32]byte{}),
		LastMessageHash:  value.NewHashBuf([32]byte{}),
		FirstLogHash:     value.NewHashBuf([32]byte{}),
		LastLogHash:      value.NewHashBuf([32]byte{}),
	}
	assertionClaim := &structures.AssertionClaim{
		AfterPendingTop:       chain.pendingInbox.GetTopHash(),
		ImportedMessagesSlice: value.NewEmptyTuple().Hash(),
		AssertionStub:         assertionStub,
	}
	disputableNode := structures.NewDisputableNode(
		assertionParams,
		assertionClaim,
		chain.pendingInbox.GetTopHash(),
		big.NewInt(0),
	)
	chain.nodeGraph.CreateNodesOnAssert(
		baseNode,
		disputableNode,
		theMachine,
		protocol.NewTimeBlocks(big.NewInt(10)),
		[32]byte{},
	)
}

func TestCreateStakers(t *testing.T) {
	chain, _, err := setUpChain()
	if err != nil {
		t.Fatal(err)
	}

	createSomeStakers(chain)
	tryMarshalUnmarshal(chain, t)
}

func setUpChain() (*ChainObserver, machine.Machine, error) {
	var dummyAddress common.Address
	theMachine, err := loader.LoadMachineFromFile("contract.ao", true, "test")
	if err != nil {
		return nil, nil, err
	}
	chain := NewChain(
		dummyAddress,
		theMachine,
		structures.ChainParams{
			StakeRequirement:        big.NewInt(1),
			GracePeriod:             structures.TimeFromSeconds(60 * 60),
			MaxExecutionSteps:       1000000,
			ArbGasSpeedLimitPerTick: 1000,
		},
		false,
		big.NewInt(10),
	)
	return chain, theMachine, nil
}

func createSomeStakers(chain *ChainObserver) {
	for i := 0; i < 5; i++ {
		createOneStaker(chain, common.BytesToAddress([]byte{byte(i)}), chain.nodeGraph.latestConfirmed.hash)
	}
}

func createOneStaker(chain *ChainObserver, stakerAddr common.Address, nodeHash [32]byte) {
	chain.CreateStake(
		ethbridge.StakeCreatedEvent{
			Staker:   stakerAddr,
			NodeHash: nodeHash,
		},
		structures.TimeFromSeconds(73),
	)
}
