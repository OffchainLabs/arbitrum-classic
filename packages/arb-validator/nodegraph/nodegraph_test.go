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
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"math/big"
	"testing"
)

var contractPath = arbos.Path()

func TestInitial(t *testing.T) {
	mach, nodeGraph := getStakedNodeGraph(t)
	expectedNode := structures.NewInitialNode(mach)

	if nodeGraph.Stakers().GetSize() != 0 {
		t.Fatal("initial stakers incorrect")
	}
	if nodeGraph.Challenges.GetSize() != 0 {
		t.Fatal("initial challenges incorrect")
	}
	if !expectedNode.Equals(nodeGraph.LatestConfirmed()) {
		t.Fatal("initial node incorrect")
	}
	if !expectedNode.Equals(nodeGraph.OldestNode()) {
		t.Fatal("initial node incorrect")
	}
	if nodeGraph.NodeFromHash(expectedNode.Hash()).Hash() != expectedNode.Hash() {
		t.Fatal("initial setup incorrect")
	}
}

func TestCreateNodes(t *testing.T) {
	mach, stakedNodeGraph := getNodeGraph(t)
	initialNode := structures.NewInitialNode(mach)
	dispNode, execAssert := getDisputableNode(initialNode)
	err, nextValid, nodes := createNodesOnAssert(
		stakedNodeGraph,
		initialNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if err != nil {
		t.Fatal("error making new node")
	}

	verifyNewNodes(t, dispNode, initialNode, nodes, stakedNodeGraph)
	if stakedNodeGraph.Leaves().IsLeaf(initialNode) {
		t.Fatal("error updating graph")
	}

	dispNode2, execAssert := getDisputableNode(nextValid)
	err, _, nodes2 := createNodesOnAssert(
		stakedNodeGraph,
		nextValid,
		dispNode2,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if err != nil {
		t.Fatal("error making new node")
	}

	verifyNewNodes(t, dispNode2, nextValid, nodes2, stakedNodeGraph)
	if stakedNodeGraph.Leaves().IsLeaf(nextValid) {
		t.Fatal("error updating graph")
	}
}

func TestGetLeaves(t *testing.T) {
	mach, stakedNodeGraph := getNodeGraph(t)
	initialNode := structures.NewInitialNode(mach)
	leaves := stakedNodeGraph.Leaves()

	if leaves.NumLeaves() != 1 || !leaves.IsLeaf(initialNode) {
		t.Fatal("incorrect leaves")
	}

	dispNode, execAssert := getDisputableNode(initialNode)
	err, nextValid, nodes := createNodesOnAssert(
		stakedNodeGraph,
		initialNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if err != nil {
		t.Fatal("error making new node")
	}

	if stakedNodeGraph.Leaves().IsLeaf(initialNode) {
		t.Fatal("error updating graph")
	}
	if stakedNodeGraph.Leaves().NumLeaves() != len(nodes) {
		t.Fatal("incorrect leaves")
	}

	for _, node := range nodes {
		if !stakedNodeGraph.Leaves().IsLeaf(node) {
			t.Fatal("incorrect leaves")
		}
	}

	dispNode2, execAssert := getDisputableNode(nextValid)
	err, _, nodes2 := createNodesOnAssert(
		stakedNodeGraph,
		nextValid,
		dispNode2,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if err != nil {
		t.Fatal("error making new node")
	}

	if stakedNodeGraph.Leaves().IsLeaf(nextValid) {
		t.Fatal("error updating graph")
	}
	for _, node := range nodes2 {
		if !stakedNodeGraph.Leaves().IsLeaf(node) {
			t.Fatal("incorrect leaves")
		}
	}

	if stakedNodeGraph.Leaves().NumLeaves() != (len(nodes) - 1 + len(nodes2)) {
		t.Fatal("incorrect leaves")
	}
}

func TestPruneInitialNodes(t *testing.T) {
	mach, stakedNodeGraph := getNodeGraph(t)
	initialNode := structures.NewInitialNode(mach)
	stakedNodeGraph.PruneNodeByHash(initialNode.Hash())
	getNode := stakedNodeGraph.NodeFromHash(initialNode.Hash())
	if getNode != nil {
		t.Fatal("error pruning")
	}
}

func TestPrunePrevNodes(t *testing.T) {
	mach, stakedNodeGraph := getNodeGraph(t)
	initialNode := structures.NewInitialNode(mach)
	dispNode, execAssert := getDisputableNode(initialNode)
	// Last value returned is not an error type
	err, nextValid, _ := createNodesOnAssert(
		stakedNodeGraph,
		initialNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if err != nil {
		t.Fatal("error making new node")
	}

	stakedNodeGraph.PruneNodeByHash(initialNode.Hash())
	getNode := stakedNodeGraph.NodeFromHash(initialNode.Hash())
	if getNode != nil {
		t.Fatal("error pruning")
	}

	getNode2 := stakedNodeGraph.NodeFromHash(nextValid.Hash())
	if getNode2 == nil {
		t.Fatal("error pruning")
	}
}

func TestPrunePrevNode(t *testing.T) {
	mach, stakedNodeGraph := getNodeGraph(t)
	initialNode := structures.NewInitialNode(mach)
	dispNode, execAssert := getDisputableNode(initialNode)
	err, _, nodes := createNodesOnAssert(
		stakedNodeGraph,
		initialNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if err != nil {
		t.Fatal("error making new node")
	}

	stakedNodeGraph.PruneNodeByHash(initialNode.Hash())
	getNode := stakedNodeGraph.NodeFromHash(initialNode.Hash())
	if getNode != nil {
		t.Fatal("error pruning")
	}

	for _, node := range nodes {
		getNode2 := stakedNodeGraph.NodeFromHash(node.Hash())
		if getNode2 == nil {
			t.Fatal("error pruning")
		}
	}
}

func TestGetLeaf(t *testing.T) {
	mach, stakedNodeGraph := getNodeGraph(t)
	initialNode := structures.NewInitialNode(mach)
	if !stakedNodeGraph.Leaves().IsLeaf(initialNode) {
		t.Fatal("error getting leaf")
	}
	dispNode, execAssert := getDisputableNode(initialNode)
	err, _, nodes := createNodesOnAssert(
		stakedNodeGraph,
		initialNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if err != nil {
		t.Fatal("error making new node")
	}

	if stakedNodeGraph.Leaves().IsLeaf(initialNode) {
		t.Fatal("error getting leaf")
	}
	for _, node := range nodes {
		if !stakedNodeGraph.Leaves().IsLeaf(node) {
			t.Fatal("error getting leaf")
		}
	}
}

func TestHasReference(t *testing.T) {
	mach, stakedNodeGraph := getNodeGraph(t)
	initialNode := structures.NewInitialNode(mach)
	if !stakedNodeGraph.Leaves().IsLeaf(initialNode) {
		t.Fatal("error getting leaf")
	}
	if !stakedNodeGraph.HasReference(initialNode) {
		t.Fatal("reference error")
	}
	stakedNodeGraph.DeleteLeaf(initialNode)
	if stakedNodeGraph.HasReference(initialNode) {
		t.Fatal("reference error")
	}
}

func TestHasReferenceWithSuccessors(t *testing.T) {
	mach, stakedNodeGraph := getNodeGraph(t)
	initialNode := structures.NewInitialNode(mach)

	dispNode, execAssert := getDisputableNode(initialNode)
	// Last value returned is not an error type
	_, _, _ = createNodesOnAssert(
		stakedNodeGraph,
		initialNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if !stakedNodeGraph.HasReference(initialNode) {
		t.Fatal("reference error")
	}

	stakedNodeGraph.DeleteLeaf(initialNode)
	if !stakedNodeGraph.HasReference(initialNode) {
		t.Fatal("reference error")
	}
}

func TestPruneNewNode(t *testing.T) {
	mach, stakedNodeGraph := getNodeGraph(t)
	initialNode := structures.NewInitialNode(mach)
	dispNode, execAssert := getDisputableNode(initialNode)
	// Last value returned is not an error type
	err, nextValid, _ := createNodesOnAssert(
		stakedNodeGraph,
		initialNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if err != nil {
		t.Fatal("error making new node")
	}

	stakedNodeGraph.PruneNodeByHash(nextValid.Hash())
	getNode := stakedNodeGraph.NodeFromHash(nextValid.Hash())
	if getNode != nil {
		t.Fatal("error pruning")
	}
}

func TestPruneAllNodes(t *testing.T) {
	mach, nodeGraph := getNodeGraph(t)
	initialNode := structures.NewInitialNode(mach)
	dispNode, execAssert := getDisputableNode(initialNode)
	err, _, nodes := createNodesOnAssert(
		nodeGraph,
		initialNode,
		dispNode,
		execAssert,
		common.NewTimeBlocks(big.NewInt(10)),
	)
	if err != nil {
		t.Fatal("error making new node")
	}

	for _, node := range nodes {
		nodeGraph.PruneNodeByHash(node.Hash())
		getNode := nodeGraph.NodeFromHash(node.Hash())
		if getNode != nil {
			t.Fatal("error pruning")
		}
	}

	getNode2 := nodeGraph.NodeFromHash(initialNode.Hash())
	if getNode2 != nil {
		t.Fatal("error pruning")
	}
}

func verifyNewNodes(
	t *testing.T,
	dispNode *valprotocol.DisputableNode,
	baseNode *structures.Node,
	nodes []*structures.Node,
	nodeGraph *NodeGraph) {
	for index, node := range nodes {
		var expectedNode *structures.Node
		if valprotocol.ChildType(index) <= valprotocol.MaxInvalidChildType {
			expectedNode = structures.NewInvalidNodeFromPrev(
				baseNode,
				dispNode,
				valprotocol.ChildType(index),
				nodeGraph.params,
				common.NewTimeBlocks(big.NewInt(10)),
			)
		} else {
			expectedNode = structures.NewValidNodeFromPrev(
				baseNode,
				dispNode,
				nodeGraph.params,
				common.NewTimeBlocks(big.NewInt(10)),
			)
		}

		getNode := nodeGraph.NodeFromHash(expectedNode.Hash())
		if !expectedNode.Equals(node) || getNode == nil || !node.Equals(getNode) {
			t.Fatal("incorrect node")
		}
	}

	for index, node := range nodes {
		expectedNode := nodeGraph.GetSuccessor(baseNode, valprotocol.ChildType(index))
		if !expectedNode.Equals(node) {
			t.Fatal("incorrect node")
		}
	}
}

func createNodesOnAssert(
	nodeGraph *NodeGraph,
	baseNode *structures.Node,
	dispNode *valprotocol.DisputableNode,
	execAssertion *protocol.ExecutionAssertion,
	currentTime *common.TimeBlocks,
) (error, *structures.Node, []*structures.Node) {

	nodes := nodeGraph.CreateNodesOnAssert(
		baseNode,
		dispNode,
		currentTime,
	)

	nextValid := nodeGraph.GetSuccessor(baseNode, valprotocol.ValidChildType)

	if err := nextValid.UpdateValidOpinion(baseNode.Machine(), execAssertion); err != nil {
		return err, nil, nil
	}
	return nil, nextValid, nodes
}

func getDisputableNode(baseNode *structures.Node) (*valprotocol.DisputableNode, *protocol.ExecutionAssertion) {
	theMachine := baseNode.Machine()
	execAssertion, _, numSteps := theMachine.ExecuteAssertion(1, true, nil, true)

	assertionParams := &valprotocol.AssertionParams{
		NumSteps:             numSteps,
		ImportedMessageCount: big.NewInt(0),
	}
	assertionStub := structures.NewExecutionAssertionStubFromWholeAssertion(execAssertion, baseNode.VMProtoData().InboxTop, nil)
	return valprotocol.NewDisputableNode(
		assertionParams,
		assertionStub,
		common.Hash{},
		big.NewInt(0),
	), execAssertion
}

var vmParams = valprotocol.ChainParams{
	StakeRequirement:        big.NewInt(1),
	GracePeriod:             common.TicksFromSeconds(60 * 60),
	MaxExecutionSteps:       1000000,
	ArbGasSpeedLimitPerTick: 1000,
}

func getNodeGraph(t *testing.T) (machine.Machine, *NodeGraph) {
	mach, err := loader.LoadMachineFromFile(contractPath, false, "cpp")
	if err != nil {
		t.Fatal(err)
	}

	return mach, NewNodeGraph(mach, vmParams)
}
