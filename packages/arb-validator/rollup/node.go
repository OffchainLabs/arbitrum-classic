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

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type ChildType uint

const (
	InvalidPendingChildType   ChildType = 0
	InvalidMessagesChildType  ChildType = 1
	InvalidExecutionChildType ChildType = 2
	ValidChildType            ChildType = 3

	MinChildType        ChildType = 0
	MaxInvalidChildType ChildType = 2
	MaxChildType        ChildType = 3
)

type VMProtoData struct {
	machineHash  [32]byte
	inboxHash    [32]byte
	pendingTop   [32]byte
	pendingCount *big.Int
}

func (d *VMProtoData) Hash() [32]byte {
	var ret [32]byte
	copy(ret[:], solsha3.SoliditySHA3(
		solsha3.Bytes32(d.machineHash),
		solsha3.Bytes32(d.inboxHash),
		solsha3.Bytes32(d.pendingTop),
		solsha3.Uint256(d.pendingCount),
	))
	return ret
}

func (node *VMProtoData) MarshalToBuf() *VMProtoDataBuf {
	return &VMProtoDataBuf{
		MachineHash:  marshalHash(node.machineHash),
		InboxHash:    marshalHash(node.inboxHash),
		PendingTop:   marshalHash(node.pendingTop),
		PendingCount: marshalBigInt(node.pendingCount),
	}
}

func (buf *VMProtoDataBuf) Unmarshal() *VMProtoData {
	return &VMProtoData{
		machineHash:  unmarshalHash(buf.MachineHash),
		inboxHash:    unmarshalHash(buf.InboxHash),
		pendingTop:   unmarshalHash(buf.PendingTop),
		pendingCount: unmarshalBigInt(buf.PendingCount),
	}
}

type Node struct {
	depth       uint64
	hash        [32]byte
	disputable  *DisputableNode
	vmProtoData *VMProtoData
	machine     machine.Machine // nil if unknown

	prev            *Node
	linkType        ChildType
	hasSuccessors   bool
	successorHashes [MaxChildType + 1][32]byte
	numStakers      uint64
}

func (node1 *Node) Equals(node2 *Node) bool {
	return node1.hash == node2.hash
}

func (node *Node) setHash() {
	var prevHashArr [32]byte
	if node.prev != nil {
		prevHashArr = node.prev.hash
	}
	innerHash := solsha3.SoliditySHA3(
		solsha3.Bytes32(node.disputable.hash),
		solsha3.Int256(node.linkType),
		solsha3.Bytes32(node.vmProtoData.Hash()),
	)
	hashSlice := solsha3.SoliditySHA3(
		solsha3.Bytes32(prevHashArr),
		solsha3.Bytes32(innerHash),
	)
	copy(node.hash[:], hashSlice)
}

func (node *Node) MarshalToBuf() *NodeBuf {
	if node.machine != nil {
		//TODO: marshal node.machine
	}
	return &NodeBuf{
		Depth:          node.depth,
		DisputableNode: node.disputable.MarshalToBuf(),
		VmProtoData:    node.vmProtoData.MarshalToBuf(),
		LinkType:       uint32(node.linkType),
		PrevHash:       marshalHash(node.prev.hash),
	}
}

func (buf *NodeBuf) Unmarshal(chain *ChainObserver) (*Node, [32]byte) {
	prevHashArr := unmarshalHash(buf.PrevHash)
	node := &Node{
		depth:       buf.Depth,
		disputable:  buf.DisputableNode.Unmarshal(),
		vmProtoData: buf.VmProtoData.Unmarshal(),
		linkType:    ChildType(buf.LinkType),
		numStakers:  0,
	}
	//TODO: try to retrieve machine from checkpoint DB; might fail
	node.setHash()
	chain.nodeFromHash[node.hash] = node

	// can't set up prev and successorHash fields yet; return prevHashArr so caller can do this later
	return node, prevHashArr
}

func GeneratePathProof(from, to *Node) [][32]byte {
	// returns nil if no proof exists
	if to == nil {
		return nil
	}
	if from == to {
		return [][32]byte{}
	}
	sub := GeneratePathProof(from, to.prev)
	if sub == nil {
		return nil
	}
	var inner32 [32]byte
	innerHash := solsha3.SoliditySHA3(
		solsha3.Bytes32(to.disputable.hash),
		solsha3.Int256(to.linkType),
		solsha3.Bytes32(to.vmProtoData.Hash()),
	)
	copy(inner32[:], innerHash)
	return append(sub, inner32)
}

func GenerateConflictProof(from, to1, to2 *Node) ([][32]byte, [][32]byte) {
	// returns nil, nil if no proof exists
	proof1 := GeneratePathProof(from, to1)
	proof2 := GeneratePathProof(from, to2)
	if proof1 == nil || proof2 == nil || len(proof1) == 0 || len(proof2) == 0 || proof1[0] == proof2[0] {
		return nil, nil
	} else {
		return proof1, proof2
	}
}

func (n *Node) EqualsFull(n2 *Node) bool {
	return n.Equals(n2) &&
		n.depth == n2.depth &&
		n.machineHash == n2.machineHash &&
		n.linkType == n2.linkType &&
		n.hasSuccessors == n2.hasSuccessors &&
		n.successorHashes == n2.successorHashes &&
		n.numStakers == n2.numStakers
}
