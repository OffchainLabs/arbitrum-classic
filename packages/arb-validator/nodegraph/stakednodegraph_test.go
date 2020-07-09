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

package nodegraph

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"math/big"
	"testing"
	"time"
)

func getStakeData(stakerAddress common.Address, node *structures.Node) (arbbridge.StakeCreatedEvent, Staker) {
	chainInfo := arbbridge.ChainInfo{
		BlockId: &common.BlockId{
			Height:     common.NewTimeBlocks(big.NewInt(73)),
			HeaderHash: common.Hash{},
		},
		LogIndex: 0,
		TxHash:   [32]byte{},
	}
	stakeEvent := arbbridge.StakeCreatedEvent{
		ChainInfo: chainInfo,
		Staker:    stakerAddress,
		NodeHash:  node.Hash(),
	}

	expectedStaker := Staker{
		stakeEvent.Staker,
		node,
		common.TicksFromBlockNum(stakeEvent.BlockId.Height),
		common.Address{},
	}

	return stakeEvent, expectedStaker
}

func TestAddStake(t *testing.T) {
	mach, _, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	expectedNode := structures.NewInitialNode(mach, txHash)

	stakerAddress := common.Address{1}
	stakeEvent, expectedStaker := getStakeData(stakerAddress, expectedNode)
	stakedNodeGraph.CreateStake(stakeEvent)
	staker := stakedNodeGraph.Stakers().Get(stakerAddress)

	if !expectedStaker.Equals(staker) {
		t.Fatal("incorrect staker")
	}

	if stakedNodeGraph.Stakers().GetSize() != 1 {
		t.Fatal("incorrect staker count")
	}

	if !staker.Location().Equals(expectedNode) {
		t.Fatal("incorrect staker location")
	}

	dispNode, execAssert := getDispNode(expectedNode)
	err, newNode, _ := createAndAssertNodes(
		stakedNodeGraph,
		expectedNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
		common.Hash{})
	if err != nil {
		t.Fatal("error making new node")
	}

	stakerAddress2 := common.Address{2}
	stakeEvent2, expectedStaker2 := getStakeData(stakerAddress2, newNode)
	stakedNodeGraph.CreateStake(stakeEvent2)
	staker2 := stakedNodeGraph.Stakers().Get(stakerAddress2)

	if !expectedStaker2.Equals(staker2) {
		t.Fatal("incorrect staker")
	}

	if stakedNodeGraph.Stakers().GetSize() != 2 {
		t.Fatal("incorrect staker count")
	}

	if !staker2.Location().Equals(newNode) {
		t.Fatal("incorrect staker location")
	}
}

func TestMoveStake(t *testing.T) {
	mach, _, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	expectedNode := structures.NewInitialNode(mach, txHash)

	stakerAddress := common.Address{1}
	stakeEvent, expectedStaker := getStakeData(stakerAddress, expectedNode)
	stakedNodeGraph.CreateStake(stakeEvent)
	staker := stakedNodeGraph.Stakers().Get(stakerAddress)

	if !expectedStaker.Equals(staker) {
		t.Fatal("incorrect staker")
	}

	if stakedNodeGraph.Stakers().GetSize() != 1 {
		t.Fatal("incorrect staker count")
	}

	if !staker.Location().Equals(expectedNode) {
		t.Fatal("incorrect staker location")
	}

	dispNode, execAssert := getDispNode(expectedNode)
	err, newNode, _ := createAndAssertNodes(
		stakedNodeGraph,
		expectedNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
		common.Hash{})
	if err != nil {
		t.Fatal("error making new node")
	}

	stakedNodeGraph.MoveStake(stakerAddress, newNode.Hash())
	staker2 := stakedNodeGraph.Stakers().Get(stakerAddress)

	expectedStaker2 := Staker{
		stakerAddress,
		newNode,
		common.TicksFromBlockNum(stakeEvent.BlockId.Height),
		common.Address{},
	}

	if !expectedStaker2.Equals(staker2) {
		t.Fatal("incorrect staker")
	}

	if stakedNodeGraph.Stakers().GetSize() != 1 {
		t.Fatal("incorrect staker count")
	}

	if !staker2.Location().Equals(newNode) {
		t.Fatal("incorrect staker location")
	}
}

func TestRemoveStake(t *testing.T) {
	mach, _, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	expectedNode := structures.NewInitialNode(mach, txHash)

	stakerAddress := common.Address{1}
	stakeEvent, expectedStaker := getStakeData(stakerAddress, expectedNode)
	stakedNodeGraph.CreateStake(stakeEvent)
	staker := stakedNodeGraph.Stakers().Get(stakerAddress)

	if !expectedStaker.Equals(staker) {
		t.Fatal("incorrect staker")
	}

	if stakedNodeGraph.Stakers().GetSize() != 1 {
		t.Fatal("incorrect staker count")
	}

	if !staker.Location().Equals(expectedNode) {
		t.Fatal("incorrect staker location")
	}

	dispNode, execAssert := getDispNode(expectedNode)
	err, newNode, _ := createAndAssertNodes(
		stakedNodeGraph,
		expectedNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
		common.Hash{})
	if err != nil {
		t.Fatal("error making new node")
	}

	stakerAddress2 := common.Address{2}
	stakeEvent2, expectedStaker2 := getStakeData(stakerAddress2, newNode)
	stakedNodeGraph.CreateStake(stakeEvent2)
	staker2 := stakedNodeGraph.Stakers().Get(stakerAddress2)

	if !expectedStaker2.Equals(staker2) {
		t.Fatal("incorrect staker")
	}

	if stakedNodeGraph.Stakers().GetSize() != 2 {
		t.Fatal("incorrect staker count")
	}

	if !staker2.Location().Equals(newNode) {
		t.Fatal("incorrect staker location")
	}

	stakedNodeGraph.RemoveStake(stakerAddress2)
	staker2 = stakedNodeGraph.Stakers().Get(stakerAddress2)
	if staker2 != nil {
		t.Fatal("incorrect removal")
	}

	if stakedNodeGraph.Stakers().GetSize() != 1 {
		t.Fatal("incorrect staker count")
	}

	stakedNodeGraph.RemoveStake(stakerAddress)
	staker = stakedNodeGraph.Stakers().Get(stakerAddress)
	if staker != nil {
		t.Fatal("incorrect removal")
	}

	if stakedNodeGraph.Stakers().GetSize() != 0 {
		t.Fatal("incorrect staker count")
	}
}

func TestNodeGraphChallenges(t *testing.T) {
	mach, _, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	initialNode := structures.NewInitialNode(mach, txHash)

	stakerAddress := common.Address{1}
	stakerAddress2 := common.Address{2}
	challengeContract := common.Address{3}
	stakeEvent, expectedStaker := getStakeData(stakerAddress, initialNode)
	stakeEvent2, _ := getStakeData(stakerAddress2, initialNode)
	stakedNodeGraph.CreateStake(stakeEvent)
	stakedNodeGraph.CreateStake(stakeEvent2)
	if stakedNodeGraph.Stakers().GetSize() != 2 {
		t.Fatal("incorrect staker count")
	}

	challenge := NewChallenge(
		&common.BlockId{
			Height:     common.NewTimeBlocks(big.NewInt(73)),
			HeaderHash: common.Hash{},
		},
		1,
		stakerAddress,
		stakerAddress2,
		challengeContract,
		initialNode)

	stakedNodeGraph.NewChallenge(challenge)
	if stakedNodeGraph.Challenges.GetSize() != 1 {
		t.Fatal("challenges count incorrect")
	}

	stakedNodeGraph.ChallengeResolved(challenge.contract, stakerAddress, stakerAddress2)
	if stakedNodeGraph.Challenges.GetSize() != 0 {
		t.Fatal("challenges count incorrect")
	}

	staker := stakedNodeGraph.Stakers().Get(stakerAddress)

	if !expectedStaker.Equals(staker) {
		t.Fatal("incorrect staker")
	}

	if stakedNodeGraph.Stakers().GetSize() != 1 {
		t.Fatal("incorrect staker count")
	}

	if !staker.Location().Equals(initialNode) {
		t.Fatal("incorrect staker location")
	}

	staker2 := stakedNodeGraph.Stakers().Get(stakerAddress2)
	if staker2 != nil {
		t.Fatal("incorrect challenge resoliution")
	}
}

func TestHasReferenceWithStakers(t *testing.T) {
	mach, _, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	initialNode := structures.NewInitialNode(mach, txHash)
	if !stakedNodeGraph.HasReference(initialNode) {
		t.Fatal("reference error")
	}
	stakedNodeGraph.DeleteLeaf(initialNode)
	if stakedNodeGraph.HasReference(initialNode) {
		t.Fatal("reference error")
	}

	stakerAddress := common.Address{1}
	stakeEvent, _ := getStakeData(stakerAddress, initialNode)
	stakedNodeGraph.CreateStake(stakeEvent)

	node := stakedNodeGraph.NodeFromHash(initialNode.Hash())
	if !stakedNodeGraph.HasReference(node) {
		t.Fatal("reference error")
	}
}

func getDispNode(baseNode *structures.Node) (*valprotocol.DisputableNode, *protocol.ExecutionAssertion) {
	theMachine := baseNode.Machine()
	timeBounds := &protocol.TimeBounds{
		LowerBoundBlock: common.NewTimeBlocks(big.NewInt(0)),
		UpperBoundBlock: common.NewTimeBlocks(big.NewInt(1000)),
	}
	execAssertion, numSteps := theMachine.ExecuteAssertion(1, value.NewEmptyTuple(), time.Hour)
	_ = execAssertion

	assertionParams := &valprotocol.AssertionParams{
		NumSteps:             numSteps,
		TimeBounds:           timeBounds,
		ImportedMessageCount: big.NewInt(0),
	}
	assertionStub := valprotocol.NewExecutionAssertionStubFromAssertion(execAssertion)
	assertionClaim := &valprotocol.AssertionClaim{
		AfterInboxTop:         common.Hash{},
		ImportedMessagesSlice: value.NewEmptyTuple().Hash(),
		AssertionStub:         assertionStub,
	}
	return valprotocol.NewDisputableNode(
		assertionParams,
		assertionClaim,
		common.Hash{},
		big.NewInt(0),
	), execAssertion
}

func createAndAssertNodes(
	nodeGraph *StakedNodeGraph,
	baseNode *structures.Node,
	dispNode *valprotocol.DisputableNode,
	execAssertion *protocol.ExecutionAssertion,
	currentTime *common.TimeBlocks,
	assertionTxHash common.Hash) (error, *structures.Node, []*structures.Node) {

	nodes := nodeGraph.CreateNodesOnAssert(
		baseNode,
		dispNode,
		currentTime,
		assertionTxHash,
	)

	nextValid := nodeGraph.GetSuccessor(baseNode, valprotocol.ValidChildType)

	if err := nextValid.UpdateValidOpinion(baseNode.Machine(), execAssertion); err != nil {
		return err, nil, nil
	}
	return nil, nextValid, nodes
}

func getStakedNodeGraph(t *testing.T) (machine.Machine, valprotocol.ChainParams, common.Hash, *StakedNodeGraph) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	vmParams := valprotocol.ChainParams{
		StakeRequirement:        big.NewInt(1),
		GracePeriod:             common.TicksFromSeconds(60 * 60),
		MaxExecutionSteps:       1000000,
		MaxBlockBoundsWidth:     20,
		ArbGasSpeedLimitPerTick: 1000,
	}
	txHash := common.Hash{}
	return mach, vmParams, txHash, NewStakedNodeGraph(mach, vmParams, txHash)
}
