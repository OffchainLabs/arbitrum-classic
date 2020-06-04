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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"strconv"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type NodeGraph struct {
	latestConfirmed *structures.Node
	leaves          *LeafSet
	nodeFromHash    map[common.Hash]*structures.Node
	oldestNode      *structures.Node
	params          valprotocol.ChainParams
}

func NewNodeGraph(machine machine.Machine, params valprotocol.ChainParams) *NodeGraph {
	newNode := structures.NewInitialNode(machine)
	nodeFromHash := make(map[common.Hash]*structures.Node)
	nodeFromHash[newNode.Hash()] = newNode
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

func (ng *NodeGraph) MarshalForCheckpoint(ctx *ckptcontext.CheckpointContext) *NodeGraphBuf {
	var allNodes []*structures.NodeBuf
	for _, n := range ng.nodeFromHash {
		allNodes = append(allNodes, n.MarshalForCheckpoint(ctx, true))
	}
	var leafHashes []common.Hash
	ng.leaves.forall(func(node *structures.Node) {
		leafHashes = append(leafHashes, node.Hash())
	})
	return &NodeGraphBuf{
		Nodes:               allNodes,
		OldestNodeHash:      ng.oldestNode.Hash().MarshalToBuf(),
		LatestConfirmedHash: ng.latestConfirmed.Hash().MarshalToBuf(),
		LeafHashes:          common.MarshalSliceOfHashes(leafHashes),
		Params:              ng.params.MarshalToBuf(),
	}
}

func (x *NodeGraphBuf) UnmarshalFromCheckpoint(ctx ckptcontext.RestoreContext) *NodeGraph {
	chain := &NodeGraph{
		latestConfirmed: nil,
		leaves:          NewLeafSet(),
		nodeFromHash:    make(map[common.Hash]*structures.Node),
		oldestNode:      nil,
		params:          x.Params.Unmarshal(),
	}

	// unmarshal nodes; their prev/successors will not be set up yet
	for _, nodeBuf := range x.Nodes {
		nd := nodeBuf.UnmarshalFromCheckpoint(ctx)
		chain.nodeFromHash[nd.Hash()] = nd
	}
	// now set up prevs and successors for all nodes
	for _, nd := range chain.nodeFromHash {
		if !nd.IsInitial() {
			prev, ok := chain.nodeFromHash[nd.PrevHash()]
			if !ok {
				log.Fatalf("Prev node %v not found for node %v while unmarshalling graph\n", nd.PrevHash(), nd.Hash())
			}
			if err := structures.Link(nd, prev); err != nil {
				// This can only fail if prev is not actually the prev of nd
				log.Fatal(err)
			}
		}
	}

	chain.oldestNode = chain.nodeFromHash[x.OldestNodeHash.Unmarshal()]
	for _, leafHashStr := range x.LeafHashes {
		leafHash := leafHashStr.Unmarshal()
		nd := chain.nodeFromHash[leafHash]
		if nd == nil {
			log.Fatal("unexpected nil node")
		}
		chain.leaves.Add(nd)
	}

	lcHash := x.LatestConfirmedHash.Unmarshal()
	chain.latestConfirmed = chain.nodeFromHash[lcHash]

	return chain
}

func (ng *NodeGraph) DebugString(stakers *StakerSet, prefix string) string {
	return ng.DebugStringForNodeRecursive(ng.oldestNode, stakers, prefix)
}

func (ng *NodeGraph) DebugStringForNodeRecursive(node *structures.Node, stakers *StakerSet, prefix string) string {
	ret := prefix + strconv.Itoa(int(node.LinkType())) + ":" + node.Hash().ShortString()
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
		succi := node.SuccessorHashes()[i]
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

func (ng *NodeGraph) pruneNode(node *structures.Node) {
	oldNode := node.Prev()
	if node.UnlinkPrev() {
		ng.considerPruningNode(oldNode)
	}
	delete(ng.nodeFromHash, node.Hash())
}

func (ng *NodeGraph) pruneOldestNode(oldest *structures.Node) {
	for i := valprotocol.MinChildType; i <= valprotocol.MaxChildType; i++ {
		succHash := oldest.SuccessorHashes()[i]
		if !succHash.Equals(common.Hash{}) {
			ng.nodeFromHash[succHash].ClearPrev()
		}
	}
	delete(ng.nodeFromHash, oldest.Hash())
}

func (ng *NodeGraph) HasReference(node *structures.Node) bool {
	if node.NumStakers() > 0 || ng.leaves.IsLeaf(node) {
		return true
	}
	for _, nodeHash := range node.SuccessorHashes() {
		if nodeHash != zeroBytes32 {
			return true
		}
	}
	return false
}

func (ng *NodeGraph) considerPruningNode(node *structures.Node) {
	if !ng.HasReference(node) {
		ng.pruneNode(node)
	}
}

func (ng *NodeGraph) getLeaf(node *structures.Node) *structures.Node {
	for _, successor := range node.SuccessorHashes() {
		if successor != zeroBytes32 {
			return ng.getLeaf(ng.nodeFromHash[successor])
		}
	}
	return node
}

func (ng *NodeGraph) CreateNodesOnAssert(
	prevNode *structures.Node,
	dispNode *valprotocol.DisputableNode,
	currentTime *common.TimeBlocks,
	assertionTxHash common.Hash,
) {
	if !ng.leaves.IsLeaf(prevNode) {
		log.Fatal("can't assert on non-leaf node")
	}
	ng.leaves.Delete(prevNode)

	// create nodes for invalid branches
	for kind := valprotocol.ChildType(0); kind <= valprotocol.MaxInvalidChildType; kind++ {
		newNode := structures.NewNodeFromInvalidPrev(prevNode, dispNode, kind, ng.params, currentTime, assertionTxHash)
		ng.nodeFromHash[newNode.Hash()] = newNode
		ng.leaves.Add(newNode)
	}

	newNode := structures.NewNodeFromValidPrev(prevNode, dispNode, ng.params, currentTime, assertionTxHash)
	ng.nodeFromHash[newNode.Hash()] = newNode
	ng.leaves.Add(newNode)
}

func (ng *NodeGraph) PruneNodeByHash(nodeHash common.Hash) {
	nd := ng.nodeFromHash[nodeHash]
	ng.pruneNode(nd)
}

func GetConflictAncestor(n1, n2 *structures.Node) (*structures.Node, *structures.Node, valprotocol.ChildType, error) {
	n1Orig := n1
	n2Orig := n2
	prevN1 := n1
	prevN2 := n2
	for n1.Depth() > n2.Depth() {
		prevN1 = n1
		n1 = n1.Prev()
	}
	for n2.Depth() > n1.Depth() {
		prevN2 = n2
		n2 = n2.Prev()
	}

	for n1 != n2 {
		prevN1 = n1
		prevN2 = n2
		n1 = n1.Prev()
		n2 = n2.Prev()
	}

	if n1 == n1Orig || n1 == n2Orig {
		return prevN1, prevN2, 0, errors.New("no conflict")
	}
	linkType := prevN1.LinkType()
	if prevN2.LinkType() < linkType {
		linkType = prevN2.LinkType()
	}

	return prevN1, prevN2, linkType, nil
}

func (ng *NodeGraph) GetSuccessor(node *structures.Node, kind valprotocol.ChildType) *structures.Node {
	return ng.nodeFromHash[node.SuccessorHashes()[kind]]
}
