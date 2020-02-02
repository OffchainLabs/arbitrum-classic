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
	"strconv"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type NodeGraph struct {
	latestConfirmed *Node
	leaves          *LeafSet
	nodeFromHash    map[common.Hash]*Node
	oldestNode      *Node
	params          structures.ChainParams
}

func NewNodeGraph(machine machine.Machine, params structures.ChainParams) *NodeGraph {
	newNode := NewInitialNode(machine)
	nodeFromHash := make(map[common.Hash]*Node)
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

func (ng *NodeGraph) MarshalForCheckpoint(ctx structures.CheckpointContext) *NodeGraphBuf {
	allNodes := make([]*NodeBuf, 0, len(ng.nodeFromHash))
	for _, n := range ng.nodeFromHash {
		allNodes = append(allNodes, n.MarshalForCheckpoint(ctx))
	}
	var leafHashes []common.Hash
	ng.leaves.forall(func(node *Node) {
		leafHashes = append(leafHashes, node.hash)
	})
	return &NodeGraphBuf{
		Nodes:               allNodes,
		OldestNodeHash:      ng.oldestNode.hash.MarshalToBuf(),
		LatestConfirmedHash: ng.latestConfirmed.hash.MarshalToBuf(),
		LeafHashes:          common.MarshalSliceOfHashes(leafHashes),
		Params:              ng.params.MarshalToBuf(),
	}
}

func (m *NodeGraphBuf) UnmarshalFromCheckpoint(ctx structures.RestoreContext) *NodeGraph {
	chain := &NodeGraph{
		latestConfirmed: nil,
		leaves:          NewLeafSet(),
		nodeFromHash:    make(map[common.Hash]*Node),
		oldestNode:      nil,
		params:          m.Params.Unmarshal(),
	}

	// unmarshal nodes; their prev/successors will not be set up yet
	for _, nodeBuf := range m.Nodes {
		node := nodeBuf.UnmarshalFromCheckpoint(ctx, chain)
		chain.nodeFromHash[node.hash] = node
	}
	// now set up prevs and successors for all nodes
	for _, nodeBuf := range m.Nodes {
		nodeHash := nodeBuf.Hash.Unmarshal()
		node := chain.nodeFromHash[nodeHash]
		if nodeBuf.PrevHash != nil {
			prevHash := nodeBuf.PrevHash.Unmarshal()
			prev, ok := chain.nodeFromHash[prevHash]
			if !ok {
				log.Fatalf("Prev node %v not found for node %v while unmarshalling graph\n", prevHash, nodeHash)
			}
			node.prev = prev
			prev.successorHashes[node.linkType] = nodeHash
		}
	}

	chain.oldestNode = chain.nodeFromHash[m.OldestNodeHash.Unmarshal()]
	for _, leafHashStr := range m.LeafHashes {
		leafHash := leafHashStr.Unmarshal()
		node := chain.nodeFromHash[leafHash]
		if node == nil {
			log.Fatal("unexpected nil node")
		}
		chain.leaves.Add(node)
	}

	lcHash := m.LatestConfirmedHash.Unmarshal()
	chain.latestConfirmed = chain.nodeFromHash[lcHash]

	return chain
}

func (ng *NodeGraph) DebugString(stakers *StakerSet, prefix string) string {
	return ng.DebugStringForNodeRecursive(ng.oldestNode, stakers, prefix)
}

func (ng *NodeGraph) DebugStringForNodeRecursive(node *Node, stakers *StakerSet, prefix string) string {
	ret := prefix + strconv.Itoa(int(node.linkType)) + ":" + node.hash.ShortString()
	if ng.leaves.IsLeaf(node) {
		ret += " leaf"
	}
	if node == ng.latestConfirmed {
		ret += " latestConfirmed"
	}
	stakers.forall(func(s *Staker) {
		if s.location.Equals(node) {
			ret += " stake:" + s.address.ShortString()
		}
	})
	ret += "\n"
	subPrefix := prefix + "  "
	for i := structures.MinChildType; i <= structures.MaxChildType; i++ {
		succi := node.successorHashes[i]
		if !succi.Equals(common.Hash{}) {
			ret += ng.DebugStringForNodeRecursive(ng.nodeFromHash[succi], stakers, subPrefix)
		}
	}
	return ret
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

func (ng *NodeGraph) pruneNode(node *Node) {
	oldNode := node.prev
	node.prev = nil // so garbage collector doesn't preserve prev anymore
	if oldNode != nil {
		oldNode.successorHashes[node.linkType] = zeroBytes32
		ng.considerPruningNode(oldNode)
	}
	delete(ng.nodeFromHash, node.hash)
}

func (ng *NodeGraph) pruneOldestNode(oldest *Node) {
	for i := structures.MinChildType; i <= structures.MaxChildType; i++ {
		succHash := oldest.successorHashes[i]
		if !succHash.Equals(common.Hash{}) {
			ng.nodeFromHash[succHash].prev = nil
		}
	}
	delete(ng.nodeFromHash, oldest.hash)
}

func (ng *NodeGraph) HasReference(node *Node) bool {
	if node.numStakers > 0 || ng.leaves.IsLeaf(node) {
		return true
	}
	for _, nodeHash := range node.successorHashes {
		if nodeHash != zeroBytes32 {
			return true
		}
	}
	return false
}

func (ng *NodeGraph) considerPruningNode(node *Node) {
	if !ng.HasReference(node) {
		ng.pruneNode(node)
	}
}

func (ng *NodeGraph) getLeaf(node *Node) *Node {
	for _, successor := range node.successorHashes {
		if successor != zeroBytes32 {
			return ng.getLeaf(ng.nodeFromHash[successor])
		}
	}
	return node
}

func (ng *NodeGraph) createNodesOnAssert(
	prevNode *Node,
	dispNode *structures.DisputableNode,
	currentTime *common.TimeBlocks,
	assertionTxHash common.Hash,
) error {
	if !ng.leaves.IsLeaf(prevNode) {
		return errors.New("can't assert on non-leaf node")
	}
	ng.leaves.Delete(prevNode)

	// create nodes for invalid branches
	for kind := structures.ChildType(0); kind <= structures.MaxInvalidChildType; kind++ {
		newNode := NewNodeFromInvalidPrev(prevNode, dispNode, kind, ng.params, currentTime, assertionTxHash)
		ng.nodeFromHash[newNode.hash] = newNode
		ng.leaves.Add(newNode)
	}

	newNode := NewNodeFromValidPrev(prevNode, dispNode, ng.params, currentTime, assertionTxHash)
	ng.nodeFromHash[newNode.hash] = newNode
	ng.leaves.Add(newNode)
	return nil
}

func (ng *NodeGraph) PruneNodeByHash(nodeHash common.Hash) {
	node := ng.nodeFromHash[nodeHash]
	ng.pruneNode(node)
}

func GetConflictAncestor(n1, n2 *Node) (*Node, *Node, structures.ChildType, error) {
	n1Orig := n1
	n2Orig := n2
	prevN1 := n1
	prevN2 := n2
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
