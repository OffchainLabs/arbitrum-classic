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

type Node struct {
	depth           uint64
	hash            [32]byte
	disputable      *DisputableNode
	machineHash     [32]byte
	machine         machine.Machine // nil if unknown
	pendingTopHash  [32]byte
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
		solsha3.Bytes32(node.protoStateHash()),
	)
	hashSlice := solsha3.SoliditySHA3(
		solsha3.Bytes32(prevHashArr),
		solsha3.Bytes32(innerHash),
	)
	copy(node.hash[:], hashSlice)
}

func (node *Node) protoStateHash() [32]byte {
	retSlice := solsha3.SoliditySHA3(
		node.machineHash,
	)
	var ret [32]byte
	copy(ret[:], retSlice)
	return ret
}

func (node *Node) MarshalToBuf() *NodeBuf {
	if node.machine != nil {
		//TODO: marshal node.machine
	}
	return &NodeBuf{
		Depth:          node.depth,
		DisputableNode: node.disputable.MarshalToBuf(),
		MachineHash:    marshalHash(node.machineHash),
		PendingTopHash: marshalHash(node.pendingTopHash),
		LinkType:       uint32(node.linkType),
		PrevHash:       marshalHash(node.prev.hash),
	}
}

func (buf *NodeBuf) Unmarshal(chain *ChainObserver) (*Node, [32]byte) {
	machineHashArr := unmarshalHash(buf.MachineHash)
	prevHashArr := unmarshalHash(buf.PrevHash)
	pthArr := unmarshalHash(buf.PendingTopHash)
	node := &Node{
		depth:          buf.Depth,
		disputable:     buf.DisputableNode.Unmarshal(),
		machineHash:    machineHashArr,
		pendingTopHash: pthArr,
		linkType:       ChildType(buf.LinkType),
		numStakers:     0,
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
	} else {
		sub := GeneratePathProof(from, to.prev)
		if sub == nil {
			return nil
		}
		var inner32 [32]byte
		innerHash := solsha3.SoliditySHA3(
			solsha3.Bytes32(to.disputable.hash),
			solsha3.Int256(to.linkType),
			solsha3.Bytes32(to.protoStateHash()),
		)
		copy(inner32[:], innerHash)
		return append(sub, inner32)
	}
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
