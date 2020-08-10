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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"math/big"
	"testing"
)

func getStakeData(stakerAddress common.Address, node *structures.Node) (arbbridge.StakeCreatedEvent, Staker) {
	chainInfo := arbbridge.ChainInfo{
		BlockId: &common.BlockId{
			Height:     common.NewTimeBlocks(big.NewInt(73)),
			HeaderHash: common.Hash{},
		},
		LogIndex: 0,
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

func graphWithOneStaker(t *testing.T) (*StakedNodeGraph, *structures.Node, common.Address, arbbridge.StakeCreatedEvent) {
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
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

	return stakedNodeGraph, expectedNode, stakerAddress, stakeEvent
}

func assertAndCreateNodes(
	t *testing.T,
	prevNode *structures.Node,
	stakedNodeGraph *StakedNodeGraph,
) (*structures.Node, []*structures.Node) {
	dispNode, execAssert := getDisputableNode(prevNode)
	err, nextValid, nodes := createNodesOnAssert(
		stakedNodeGraph.NodeGraph,
		prevNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if err != nil {
		t.Fatal("error making new node")
	}

	return nextValid, nodes
}

func stake1stAndVerify(
	t *testing.T,
	stakedNodeGraph *StakedNodeGraph,
	node *structures.Node,
	address common.Address) {
	stakeEvent2, expectedStaker2 := getStakeData(address, node)
	stakedNodeGraph.CreateStake(stakeEvent2)
	staker2 := stakedNodeGraph.Stakers().Get(address)

	if !expectedStaker2.Equals(staker2) {
		t.Fatal("incorrect staker")
	}

	if stakedNodeGraph.Stakers().GetSize() != 1 {
		t.Fatal("incorrect staker count")
	}

	if !staker2.Location().Equals(node) {
		t.Fatal("incorrect staker location")
	}
}

func stake2ndAndVerify(
	t *testing.T,
	stakedNodeGraph *StakedNodeGraph,
	node *structures.Node,
	address common.Address) {
	stakeEvent2, expectedStaker2 := getStakeData(address, node)
	stakedNodeGraph.CreateStake(stakeEvent2)
	staker2 := stakedNodeGraph.Stakers().Get(address)

	if !expectedStaker2.Equals(staker2) {
		t.Fatal("incorrect staker")
	}

	if stakedNodeGraph.Stakers().GetSize() != 2 {
		t.Fatal("incorrect staker count")
	}

	if !staker2.Location().Equals(node) {
		t.Fatal("incorrect staker location")
	}
}

func TestAddStake(t *testing.T) {
	stakedNodeGraph, expectedNode, _, _ := graphWithOneStaker(t)
	newNode, _ := assertAndCreateNodes(t, expectedNode, stakedNodeGraph)
	stake2ndAndVerify(t, stakedNodeGraph, newNode, common.Address{2})
}

func TestStakerPruneInfoInitial(t *testing.T) {
	stakedNodeGraph, _, _, _ := graphWithOneStaker(t)
	mootedParams, oldParams := stakedNodeGraph.GenerateStakerPruneInfo()
	if len(mootedParams) != 0 {
		t.Fatal("incorrect results")
	}
	if len(oldParams) != 0 {
		t.Fatal("incorrect results")
	}
}

func TestNodePruneInfoInitial(t *testing.T) {
	stakedNodeGraph, _, _, _ := graphWithOneStaker(t)
	params := stakedNodeGraph.GenerateNodePruneInfo(stakedNodeGraph.Stakers())
	if len(params) != 0 {
		t.Fatal("incorrect results")
	}
}

func TestStakerPruneInfoBase(t *testing.T) {
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	initialNode := structures.NewInitialNode(mach, txHash)
	_, _ = assertAndCreateNodes(t, initialNode, stakedNodeGraph)

	stakerAddress := common.Address{1}
	stake1stAndVerify(t, stakedNodeGraph, initialNode, stakerAddress)

	mootedParams, oldParams := stakedNodeGraph.GenerateStakerPruneInfo()
	if len(mootedParams) != 0 {
		t.Fatal("incorrect results")
	}
	if len(oldParams) != 0 {
		t.Fatal("incorrect results")
	}
}

func TestNodePruneInfoBase(t *testing.T) {
	stakedNodeGraph, initialNode, _, _ := graphWithOneStaker(t)
	_, _ = assertAndCreateNodes(t, initialNode, stakedNodeGraph)

	params := stakedNodeGraph.GenerateNodePruneInfo(stakedNodeGraph.Stakers())
	if len(params) != 0 {
		log.Println("params ", params)
		t.Fatal("incorrect results")
	}
}

func stakeAllNodes(stakedNodeGraph *StakedNodeGraph, nodes []*structures.Node, startAddress int) {
	for index, node := range nodes {
		var addrr = byte(startAddress + index)
		stakerAddress := common.Address{addrr}
		stakeEvent, _ := getStakeData(stakerAddress, node)
		stakedNodeGraph.CreateStake(stakeEvent)
	}
}

func TestStakerPruneInfo(t *testing.T) {
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	initialNode := structures.NewInitialNode(mach, txHash)

	_, nodes := assertAndCreateNodes(t, initialNode, stakedNodeGraph)
	stakeAllNodes(stakedNodeGraph, nodes, 2)

	mootedParams, oldParams := stakedNodeGraph.GenerateStakerPruneInfo()
	if len(mootedParams) != 3 {
		t.Fatal("incorrect results, mootedParams, ", mootedParams)
	}
	if len(oldParams) != 0 {
		t.Fatal("incorrect results, oldParams, ", oldParams)
	}
}

func TestNodePruneInfo(t *testing.T) {
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	initialNode := structures.NewInitialNode(mach, txHash)

	_, nodes := assertAndCreateNodes(t, initialNode, stakedNodeGraph)
	stakeAllNodes(stakedNodeGraph, nodes, 2)

	params := stakedNodeGraph.GenerateNodePruneInfo(stakedNodeGraph.Stakers())
	if len(params) != 0 {
		log.Println("params ", params)
		t.Fatal("incorrect results")
	}
}

func TestNodePruneInfo2(t *testing.T) {
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	initialNode := structures.NewInitialNode(mach, txHash)

	nextValid, nodes := assertAndCreateNodes(t, initialNode, stakedNodeGraph)
	stakeAllNodes(stakedNodeGraph, nodes, 2)
	stakedNodeGraph.UpdateLatestConfirmed(nextValid)

	params := stakedNodeGraph.GenerateNodePruneInfo(stakedNodeGraph.Stakers())
	if len(params) != 0 {
		log.Println("params ", params)
		t.Fatal("incorrect results")
	}
}

func TestStakerPruneInfo2(t *testing.T) {
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	initialNode := structures.NewInitialNode(mach, txHash)

	nextValid, nodes := assertAndCreateNodes(t, initialNode, stakedNodeGraph)
	stakeAllNodes(stakedNodeGraph, nodes, 2)
	stakedNodeGraph.UpdateLatestConfirmed(nextValid)

	mootedParams, oldParams := stakedNodeGraph.GenerateStakerPruneInfo()
	if len(mootedParams) != 2 {
		t.Fatal("incorrect results, mootedParams, ", mootedParams)
	}
	if len(oldParams) != 0 {
		t.Fatal("incorrect results, oldParams, ", oldParams)
	}
}

func TestNodePruneInfo3(t *testing.T) {
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	initialNode := structures.NewInitialNode(mach, txHash)

	nextValid, nodes := assertAndCreateNodes(t, initialNode, stakedNodeGraph)
	nextValid2, _ := assertAndCreateNodes(t, nextValid, stakedNodeGraph)

	stakeAllNodes(stakedNodeGraph, nodes, 2)
	stakedNodeGraph.UpdateLatestConfirmed(nextValid2)

	params := stakedNodeGraph.GenerateNodePruneInfo(stakedNodeGraph.Stakers())
	if len(params) != 2 {
		t.Fatal("incorrect results", len(params))
	}
}

func TestStakerPruneInfo3(t *testing.T) {
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	initialNode := structures.NewInitialNode(mach, txHash)
	nextValid, _ := assertAndCreateNodes(t, initialNode, stakedNodeGraph)
	nextValid2, nodes2 := assertAndCreateNodes(t, nextValid, stakedNodeGraph)

	stakeAllNodes(stakedNodeGraph, nodes2, 2)
	stakedNodeGraph.UpdateLatestConfirmed(nextValid2)

	mootedParams, oldParams := stakedNodeGraph.GenerateStakerPruneInfo()
	if len(mootedParams) != 2 {
		t.Error("incorrect results, mootedParams, ", len(mootedParams))
	}
	if len(oldParams) != 0 {
		t.Error("incorrect results, oldParams, ", len(oldParams))
	}
}

func TestStakerPruneInfo4(t *testing.T) {
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
	initialNode := structures.NewInitialNode(mach, txHash)
	nextValid, nodes := assertAndCreateNodes(t, initialNode, stakedNodeGraph)
	nextValid2, _ := assertAndCreateNodes(t, nextValid, stakedNodeGraph)
	stakeAllNodes(stakedNodeGraph, nodes, 2)
	stakedNodeGraph.UpdateLatestConfirmed(nextValid2)

	mootedParams, oldParams := stakedNodeGraph.GenerateStakerPruneInfo()
	if len(mootedParams) != 2 {
		t.Fatal("incorrect results, mootedParams, ", mootedParams)
	}
	if len(oldParams) != 1 {
		t.Fatal("incorrect results, oldParams, ", oldParams)
	}
}

func TestMoveStake(t *testing.T) {
	stakedNodeGraph, expectedNode, stakerAddress, stakeEvent := graphWithOneStaker(t)
	newNode, _ := assertAndCreateNodes(t, expectedNode, stakedNodeGraph)

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
	stakedNodeGraph, expectedNode, stakerAddress, _ := graphWithOneStaker(t)
	newNode, _ := assertAndCreateNodes(t, expectedNode, stakedNodeGraph)

	stakerAddress2 := common.Address{2}
	stake2ndAndVerify(t, stakedNodeGraph, newNode, stakerAddress2)
	staker2 := stakedNodeGraph.Stakers().Get(stakerAddress2)

	stakedNodeGraph.RemoveStake(stakerAddress2)
	staker2 = stakedNodeGraph.Stakers().Get(stakerAddress2)
	if staker2 != nil {
		t.Fatal("incorrect removal")
	}

	if stakedNodeGraph.Stakers().GetSize() != 1 {
		t.Fatal("incorrect staker count")
	}

	stakedNodeGraph.RemoveStake(stakerAddress)
	staker := stakedNodeGraph.Stakers().Get(stakerAddress)
	if staker != nil {
		t.Fatal("incorrect removal")
	}

	if stakedNodeGraph.Stakers().GetSize() != 0 {
		t.Fatal("incorrect staker count")
	}
}

func TestNodeGraphChallenges(t *testing.T) {
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
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
	mach, txHash, stakedNodeGraph := getStakedNodeGraph(t)
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

func getStakedNodeGraph(t *testing.T) (machine.Machine, common.Hash, *StakedNodeGraph) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	return mach, common.Hash{}, NewStakedNodeGraph(mach, vmParams, common.Hash{})
}
