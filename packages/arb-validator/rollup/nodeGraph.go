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
	"errors"
	"log"
	"math/big"
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type NodeGraph struct {
	*sync.RWMutex
	latestConfirmed *Node
	leaves          *LeafSet
	nodeFromHash    map[[32]byte]*Node
	oldestNode      *Node
}

func NewNodeGraph(machine machine.Machine) *NodeGraph {
	ret := &NodeGraph{
		latestConfirmed: nil,
		leaves:          NewLeafSet(),
		nodeFromHash:    make(map[[32]byte]*Node),
		oldestNode:      nil,
	}
	ret.CreateInitialNode(machine)
	return ret
}

func (chain *NodeGraph) MarshalToBuf() *NodeGraphBuf {
	var allNodes []*NodeBuf
	for _, v := range chain.nodeFromHash {
		allNodes = append(allNodes, v.MarshalToBuf())
	}
	var leafHashes [][32]byte
	chain.leaves.forall(func(node *Node) {
		leafHashes = append(leafHashes, node.hash)
	})
	return &NodeGraphBuf{
		Nodes:               allNodes,
		OldestNodeHash:      marshalHash(chain.oldestNode.hash),
		LatestConfirmedHash: marshalHash(chain.latestConfirmed.hash),
		LeafHashes:          marshalSliceOfHashes(leafHashes),
	}
}

func (buf *NodeGraphBuf) Unmarshal() *NodeGraph {
	chain := &NodeGraph{
		latestConfirmed: nil,
		leaves:          NewLeafSet(),
		nodeFromHash:    make(map[[32]byte]*Node),
		oldestNode:      nil,
	}

	for _, nodeBuf := range buf.Nodes {
		nodeHash := unmarshalHash(nodeBuf.Hash)
		node := chain.nodeFromHash[nodeHash]
		prevHash := unmarshalHash(nodeBuf.PrevHash)
		if prevHash != zeroBytes32 {
			prev := chain.nodeFromHash[prevHash]
			node.prev = prev
			prev.successorHashes[node.linkType] = nodeHash
		}
	}
	chain.oldestNode = chain.nodeFromHash[unmarshalHash(buf.OldestNodeHash)]
	for _, leafHashStr := range buf.LeafHashes {
		leafHash := unmarshalHash(leafHashStr)
		chain.leaves.Add(chain.nodeFromHash[leafHash])
	}

	lcHash := unmarshalHash(buf.LatestConfirmedHash)
	chain.latestConfirmed = chain.nodeFromHash[lcHash]

	return chain
}

func (chain *NodeGraph) CreateInitialNode(machine machine.Machine) {
	newNode := &Node{
		depth:          0,
		machineHash:    machine.Hash(),
		machine:        machine.Clone(),
		pendingTopHash: value.NewEmptyTuple().Hash(),
		linkType:       ValidChildType,
		numStakers:     0,
	}
	newNode.setHash()
	chain.leaves.Add(newNode)
	chain.latestConfirmed = newNode
}

func (chain *NodeGraph) pruneNode(node *Node) {
	oldNode := node.prev
	node.prev = nil // so garbage collector doesn't preserve prev anymore
	if oldNode != nil {
		oldNode.successorHashes[node.linkType] = zeroBytes32
		chain.considerPruningNode(oldNode)
	}
	delete(chain.nodeFromHash, node.hash)
}

func (chain *NodeGraph) considerPruningNode(node *Node) {
	if node.numStakers > 0 {
		return
	}
	for kind := MinChildType; kind <= MaxChildType; kind++ {
		if node.successorHashes[kind] != zeroBytes32 {
			return
		}
	}
	chain.pruneNode(node)
}

func (chain *NodeGraph) CreateNodesOnAssert(
	prevNode *Node,
	dispNode *DisputableNode,
	afterMachineHash [32]byte,
	afterMachine machine.Machine,
) {
	if !chain.leaves.IsLeaf(prevNode) {
		log.Fatal("can't assert on non-leaf node")
	}
	chain.leaves.Delete(prevNode)
	prevNode.hasSuccessors = true

	// create node for valid branch
	if afterMachine != nil {
		afterMachine = afterMachine.Clone()
	}
	newNode := &Node{
		depth:          1 + prevNode.depth,
		disputable:     dispNode,
		prev:           prevNode,
		linkType:       ValidChildType,
		machineHash:    afterMachineHash,
		pendingTopHash: dispNode.afterPendingTop,
		machine:        afterMachine,
		numStakers:     0,
	}
	newNode.setHash()
	prevNode.successorHashes[ValidChildType] = newNode.hash
	chain.leaves.Add(newNode)

	// create nodes for invalid branches
	for kind := ChildType(0); kind <= MaxInvalidChildType; kind++ {
		newNode := &Node{
			depth:          1 + prevNode.depth,
			disputable:     dispNode,
			prev:           prevNode,
			linkType:       kind,
			machineHash:    prevNode.machineHash,
			machine:        prevNode.machine,
			pendingTopHash: prevNode.pendingTopHash,
			numStakers:     0,
		}
		newNode.setHash()
		prevNode.successorHashes[kind] = newNode.hash
		chain.leaves.Add(newNode)
	}
}

func (chain *NodeGraph) notifyAssert(
	prevLeafHash [32]byte,
	timeBounds [2]RollupTime,
	afterPendingTop [32]byte,
	importedMessagesSlice [32]byte,
	importedMessageCount *big.Int,
	assertionStub *protocol.AssertionStub,
) {
	disputableNode := &DisputableNode{
		prevNodeHash:          prevLeafHash,
		timeBounds:            timeBounds,
		afterPendingTop:       afterPendingTop,
		importedMessagesSlice: importedMessagesSlice,
		importedMessageCount:  importedMessageCount,
		assertionStub:         assertionStub,
	}
	disputableNode.hash = disputableNode._hash()
	chain.CreateNodesOnAssert(chain.nodeFromHash[prevLeafHash], disputableNode, unmarshalHash(disputableNode.assertionStub.AfterHash), nil)
}

func (chain *NodeGraph) ConfirmNode(nodeHash [32]byte) {
	node := chain.nodeFromHash[nodeHash]
	chain.latestConfirmed = node
	chain.considerPruningNode(node.prev)
	for chain.oldestNode != chain.latestConfirmed {
		if chain.oldestNode.numStakers > 0 {
			return
		}
		var successor *Node
		for kind := MinChildType; kind <= MaxChildType; kind++ {
			if node.successorHashes[kind] != zeroBytes32 {
				if successor != nil {
					return
				}
				successor = chain.nodeFromHash[node.successorHashes[kind]]
			}
		}
		chain.pruneNode(chain.oldestNode)
		chain.oldestNode = successor
	}
}

func (chain *NodeGraph) PruneNodeByHash(nodeHash [32]byte) {
	chain.pruneNode(chain.nodeFromHash[nodeHash])
}

func (chain *NodeGraph) CommonAncestor(n1, n2 *Node) *Node {
	n1, _, _ = chain.GetConflictAncestor(n1, n2)
	return n1
}

func (chain *NodeGraph) GetConflictAncestor(n1, n2 *Node) (*Node, ChildType, error) {
	n1Orig := n1
	n2Orig := n2
	prevN1 := n1
	prevN2 := n1
	for n1.depth > n2.depth {
		prevN1 = n1
		n1 = n1.prev
	}
	for n2.depth > n1.depth {
		prevN2 = n2
		n2 = n2.prev
	}

	for n1 != n2 {
		prevN1 = n1
		prevN2 = n2
		n1 = n1.prev
		n2 = n2.prev
	}

	if n1 == n1Orig || n1 == n2Orig {
		return n1, 0, errors.New("no conflict")
	}
	linkType := prevN1.linkType
	if prevN2.linkType < linkType {
		linkType = prevN2.linkType
	}

	return n1, linkType, nil
}
