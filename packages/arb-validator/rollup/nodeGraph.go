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

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type NodeGraph struct {
	latestConfirmed *Node
	leaves          *LeafSet
	nodeFromHash    map[[32]byte]*Node
	oldestNode      *Node
	params          structures.ChainParams
}

func NewNodeGraph(machine machine.Machine, params structures.ChainParams) *NodeGraph {
	newNode := NewInitialNode(machine)
	nodeFromHash := make(map[[32]byte]*Node)
	nodeFromHash[newNode.hash] = newNode
	leaves := NewLeafSet()
	leaves.Add(newNode)
	return &NodeGraph{
		latestConfirmed: newNode,
		leaves:          leaves,
		nodeFromHash:    nodeFromHash,
		oldestNode:      newNode,
		params:          params,
	}
}

func (chain *NodeGraph) MarshalForCheckpoint(ctx structures.CheckpointContext) *NodeGraphBuf {
	var allNodes []*NodeBuf
	for _, n := range chain.nodeFromHash {
		allNodes = append(allNodes, n.MarshalForCheckpoint(ctx))
	}
	var leafHashes [][32]byte
	chain.leaves.forall(func(node *Node) {
		leafHashes = append(leafHashes, node.hash)
	})
	return &NodeGraphBuf{
		Nodes:               allNodes,
		OldestNodeHash:      utils.MarshalHash(chain.oldestNode.hash),
		LatestConfirmedHash: utils.MarshalHash(chain.latestConfirmed.hash),
		LeafHashes:          utils.MarshalSliceOfHashes(leafHashes),
		Params:              chain.params.MarshalToBuf(),
	}
}

func (buf *NodeGraphBuf) UnmarshalFromCheckpoint(ctx structures.RestoreContext) *NodeGraph {
	chain := &NodeGraph{
		latestConfirmed: nil,
		leaves:          NewLeafSet(),
		nodeFromHash:    make(map[[32]byte]*Node),
		oldestNode:      nil,
		params:          buf.Params.Unmarshal(),
	}

	// unmarshal nodes; their prev/successors will not be set up yet
	for _, nodeBuf := range buf.Nodes {
		_ = nodeBuf.UnmarshalFromCheckpoint(ctx, chain)
	}
	// now set up prevs and successors for all nodes
	for _, nodeBuf := range buf.Nodes {
		nodeHash := utils.UnmarshalHash(nodeBuf.Hash)
		node := chain.nodeFromHash[nodeHash]
		if nodeBuf.PrevHash != nil {
			prevHash := utils.UnmarshalHash(nodeBuf.PrevHash)
			prev := chain.nodeFromHash[prevHash]
			node.prev = prev
			prev.successorHashes[node.linkType] = nodeHash
		}
	}

	chain.oldestNode = chain.nodeFromHash[utils.UnmarshalHash(buf.OldestNodeHash)]
	for _, leafHashStr := range buf.LeafHashes {
		leafHash := utils.UnmarshalHash(leafHashStr)
		node := chain.nodeFromHash[leafHash]
		if node == nil {
			log.Fatal("unexpected nil node")
		}
		chain.leaves.Add(node)
	}

	lcHash := utils.UnmarshalHash(buf.LatestConfirmedHash)
	chain.latestConfirmed = chain.nodeFromHash[lcHash]

	return chain
}

func (ng *NodeGraph) Equals(ng2 *NodeGraph) bool {
	if !ng.latestConfirmed.Equals(ng2.latestConfirmed) ||
		!ng.oldestNode.Equals(ng2.oldestNode) ||
		!ng.leaves.Equals(ng2.leaves) ||
		len(ng.nodeFromHash) != len(ng2.nodeFromHash) ||
		!ng.params.Equals(ng.params) {
		return false
	}
	for h, n := range ng.nodeFromHash {
		if ng2.nodeFromHash[h] == nil || !n.Equals(ng2.nodeFromHash[h]) {
			return false
		}
	}
	return true
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

func (chain *NodeGraph) HasReference(node *Node) bool {
	if node.numStakers > 0 || chain.leaves.IsLeaf(node) {
		return true
	}
	for _, nodeHash := range node.successorHashes {
		if nodeHash != zeroBytes32 {
			return true
		}
	}
	return false
}

func (chain *NodeGraph) considerPruningNode(node *Node) {
	if !chain.HasReference(node) {
		chain.pruneNode(node)
	}
}

func (chain *NodeGraph) getLeaf(node *Node) *Node {
	for _, successor := range node.successorHashes {
		if successor != zeroBytes32 {
			return chain.getLeaf(chain.nodeFromHash[successor])
		}
	}
	return node
}

func (chain *NodeGraph) CreateNodesOnAssert(
	prevNode *Node,
	dispNode *structures.DisputableNode,
	afterMachine machine.Machine,
	currentTime *protocol.TimeBlocks,
	assertionTxHash [32]byte,
) {
	if !chain.leaves.IsLeaf(prevNode) {
		log.Fatal("can't assert on non-leaf node")
	}
	chain.leaves.Delete(prevNode)

	// create nodes for invalid branches
	for kind := structures.ChildType(0); kind <= structures.MaxInvalidChildType; kind++ {
		newNode := NewNodeFromInvalidPrev(prevNode, dispNode, kind, chain.params, currentTime, assertionTxHash)
		chain.nodeFromHash[newNode.hash] = newNode
		chain.leaves.Add(newNode)
	}

	// create node for valid branch
	if afterMachine != nil {
		afterMachine = afterMachine.Clone()
	}

	newNode := NewNodeFromValidPrev(prevNode, dispNode, afterMachine, chain.params, currentTime, assertionTxHash)
	chain.nodeFromHash[newNode.hash] = newNode
	chain.leaves.Add(newNode)
}

func (chain *NodeGraph) PruneNodeByHash(nodeHash [32]byte) {
	node := chain.nodeFromHash[nodeHash]
	chain.pruneNode(node)
}

func (chain *NodeGraph) CommonAncestor(n1, n2 *Node) *Node {
	n1, _, _, _ = chain.GetConflictAncestor(n1, n2)
	return n1.prev
}

func (chain *NodeGraph) generateNodePruneInfo() []pruneParams {
	prunesToDo := []pruneParams{}
	chain.leaves.forall(func(leaf *Node) {
		if leaf != chain.latestConfirmed {
			leafAncestor, _, err := GetConflictAncestor(leaf, chain.latestConfirmed)
			if err == nil {
				prunesToDo = append(prunesToDo, pruneParams{
					leafHash:     leaf.hash,
					ancestorHash: leafAncestor.prev.hash,
					leafProof:    GeneratePathProof(leafAncestor.prev, leaf),
					ancProof:     GeneratePathProof(leafAncestor.prev, chain.latestConfirmed),
				})
			}
		}
	})
	return prunesToDo
}

func (chain *NodeGraph) GetConflictAncestor(n1, n2 *Node) (*Node, *Node, structures.ChildType, error) {
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
		return prevN1, prevN2, 0, errors.New("no conflict")
	}
	linkType := prevN1.linkType
	if prevN2.linkType < linkType {
		linkType = prevN2.linkType
	}

	return prevN1, prevN2, linkType, nil
}
