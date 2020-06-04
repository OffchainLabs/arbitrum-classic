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

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type NodeGraph struct {
	latestConfirmed *Node
	leaves          *LeafSet
	nodeFromHash    map[common.Hash]*Node
	oldestNode      *Node
	params          valprotocol.ChainParams
}

func NewNodeGraph(machine machine.Machine, params valprotocol.ChainParams) *NodeGraph {
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

func (chain *NodeGraph) MarshalForCheckpoint(ctx *checkpointing.CheckpointContext) *NodeGraphBuf {
	var allNodes []*NodeBuf
	for _, n := range chain.nodeFromHash {
		allNodes = append(allNodes, n.MarshalForCheckpoint(ctx))
	}
	var leafHashes []common.Hash
	chain.leaves.forall(func(node *Node) {
		leafHashes = append(leafHashes, node.hash)
	})
	return &NodeGraphBuf{
		Nodes:               allNodes,
		OldestNodeHash:      chain.oldestNode.hash.MarshalToBuf(),
		LatestConfirmedHash: chain.latestConfirmed.hash.MarshalToBuf(),
		LeafHashes:          common.MarshalSliceOfHashes(leafHashes),
		Params:              chain.params.MarshalToBuf(),
	}
}

func (buf *NodeGraphBuf) UnmarshalFromCheckpoint(ctx checkpointing.RestoreContext) *NodeGraph {
	chain := &NodeGraph{
		latestConfirmed: nil,
		leaves:          NewLeafSet(),
		nodeFromHash:    make(map[common.Hash]*Node),
		oldestNode:      nil,
		params:          buf.Params.Unmarshal(),
	}

	// unmarshal nodes; their prev/successors will not be set up yet
	for _, nodeBuf := range buf.Nodes {
		node := nodeBuf.UnmarshalFromCheckpoint(ctx)
		chain.nodeFromHash[node.hash] = node
	}
	// now set up prevs and successors for all nodes
	for _, node := range chain.nodeFromHash {
		if !node.isInitial() {
			prev, ok := chain.nodeFromHash[node.prevHash]
			if !ok {
				log.Fatalf("Prev node %v not found for node %v while unmarshalling graph\n", node.prevHash, node.hash)
			}
			node.prev = prev
			prev.successorHashes[node.linkType] = node.hash
		}
	}

	chain.oldestNode = chain.nodeFromHash[buf.OldestNodeHash.Unmarshal()]
	for _, leafHashStr := range buf.LeafHashes {
		leafHash := leafHashStr.Unmarshal()
		node := chain.nodeFromHash[leafHash]
		if node == nil {
			log.Fatal("unexpected nil node")
		}
		chain.leaves.Add(node)
	}

	lcHash := buf.LatestConfirmedHash.Unmarshal()
	chain.latestConfirmed = chain.nodeFromHash[lcHash]

	return chain
}

func (ng *NodeGraph) DebugString(stakers *StakerSet, prefix string) string {
	return ng.DebugStringForNodeRecursive(ng.oldestNode, stakers, prefix)
}

func (ng *NodeGraph) DebugStringForNodeRecursive(node *Node, stakers *StakerSet, prefix string) string {
	ret := prefix + strconv.Itoa(int(node.linkType)) + ":" + node.hash.ShortString()
	if ng.leaves.IsLeaf(node) {
		ret = ret + " leaf"
	}
	if node == ng.latestConfirmed {
		ret = ret + " latestConfirmed"
	}
	stakers.forall(func(s *Staker) {
		if s.location.Equals(node) {
			ret = ret + " stake:" + s.address.ShortString()
		}
	})
	ret = ret + "\n"
	subPrefix := prefix + "  "
	for i := valprotocol.MinChildType; i <= valprotocol.MaxChildType; i++ {
		succi := node.successorHashes[i]
		if !succi.Equals(common.Hash{}) {
			ret = ret + ng.DebugStringForNodeRecursive(ng.nodeFromHash[succi], stakers, subPrefix)
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

func (chain *NodeGraph) pruneNode(node *Node) {
	oldNode := node.prev
	node.prev = nil // so garbage collector doesn't preserve prev anymore
	if oldNode != nil {
		oldNode.successorHashes[node.linkType] = zeroBytes32
		chain.considerPruningNode(oldNode)
	}
	delete(chain.nodeFromHash, node.hash)
}

func (chain *NodeGraph) pruneOldestNode(oldest *Node) {
	for i := valprotocol.MinChildType; i <= valprotocol.MaxChildType; i++ {
		succHash := oldest.successorHashes[i]
		if !succHash.Equals(common.Hash{}) {
			chain.nodeFromHash[succHash].prev = nil
		}
	}
	delete(chain.nodeFromHash, oldest.hash)
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
	dispNode *valprotocol.DisputableNode,
	currentTime *common.TimeBlocks,
	assertionTxHash common.Hash,
) {
	if !chain.leaves.IsLeaf(prevNode) {
		log.Fatal("can't assert on non-leaf node")
	}
	chain.leaves.Delete(prevNode)

	// create nodes for invalid branches
	for kind := valprotocol.ChildType(0); kind <= valprotocol.MaxInvalidChildType; kind++ {
		newNode := NewNodeFromInvalidPrev(prevNode, dispNode, kind, chain.params, currentTime, assertionTxHash)
		chain.nodeFromHash[newNode.hash] = newNode
		chain.leaves.Add(newNode)
	}

	newNode := NewNodeFromValidPrev(prevNode, dispNode, chain.params, currentTime, assertionTxHash)
	chain.nodeFromHash[newNode.hash] = newNode
	chain.leaves.Add(newNode)
}

func (chain *NodeGraph) PruneNodeByHash(nodeHash common.Hash) {
	node := chain.nodeFromHash[nodeHash]
	chain.pruneNode(node)
}

func (chain *NodeGraph) CommonAncestor(n1, n2 *Node) *Node {
	n1, _, _, _ = chain.GetConflictAncestor(n1, n2)
	return n1.prev
}

func (chain *NodeGraph) GetConflictAncestor(n1, n2 *Node) (*Node, *Node, valprotocol.ChildType, error) {
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

func (chain *NodeGraph) GetSuccessor(node *Node, kind valprotocol.ChildType) *Node {
	return chain.nodeFromHash[node.successorHashes[kind]]
}
