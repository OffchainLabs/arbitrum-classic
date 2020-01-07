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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type ExecutionNodeData struct {
}

type Node struct {
	prev        *Node
	deadline    structures.TimeTicks
	disputable  *structures.DisputableNode
	linkType    structures.ChildType
	vmProtoData *structures.VMProtoData

	machine      machine.Machine // nil if unknown
	depth        uint64
	nodeDataHash [32]byte
	innerHash    [32]byte
	hash         [32]byte

	successorHashes [structures.MaxChildType + 1][32]byte
	numStakers      uint64
}

func NewInitialNode(
	machine machine.Machine,
) *Node {
	ret := &Node{
		prev:       nil,
		deadline:   structures.TimeTicks{big.NewInt(0)},
		disputable: nil,
		linkType:   0,
		vmProtoData: structures.NewVMProtoData(
			machine.Hash(),
			value.NewEmptyTuple().Hash(),
			big.NewInt(0),
		),
		machine: machine,
		depth:   0,
	}
	ret.setHash([32]byte{})
	return ret
}

func NewNodeFromValidPrev(
	prev *Node,
	disputable *structures.DisputableNode,
	machine machine.Machine,
	params structures.ChainParams,
	currentTime *protocol.TimeBlocks,
) *Node {
	return NewNodeFromPrev(
		prev,
		disputable,
		structures.ValidChildType,
		params,
		currentTime,
		disputable.ValidAfterVMProtoData(prev.vmProtoData),
		machine,
	)
}

func NewNodeFromInvalidPrev(
	prev *Node,
	disputable *structures.DisputableNode,
	kind structures.ChildType,
	params structures.ChainParams,
	currentTime *protocol.TimeBlocks,
) *Node {
	return NewNodeFromPrev(
		prev,
		disputable,
		kind,
		params,
		currentTime,
		prev.vmProtoData,
		prev.machine,
	)
}

func NewNodeFromPrev(
	prev *Node,
	disputable *structures.DisputableNode,
	kind structures.ChildType,
	params structures.ChainParams,
	currentTime *protocol.TimeBlocks,
	vmProtoData *structures.VMProtoData,
	machine machine.Machine,
) *Node {
	checkTime := disputable.CheckTime(params)
	deadlineTicks := structures.TimeFromBlockNum(currentTime).Add(params.GracePeriod)
	if deadlineTicks.Cmp(prev.deadline) >= 0 {
		deadlineTicks = deadlineTicks.Add(checkTime)
	} else {
		deadlineTicks = prev.deadline.Add(checkTime)
	}

	ret := &Node{
		prev:        prev,
		deadline:    deadlineTicks,
		disputable:  disputable,
		linkType:    kind,
		vmProtoData: vmProtoData,
		machine:     machine,
		depth:       prev.depth + 1,
	}
	ret.setHash(ret.NodeDataHash(params))
	prev.successorHashes[kind] = ret.hash
	return ret
}

func (node *Node) Equals(node2 *Node) bool {
	return node.hash == node2.hash
}

func (node *Node) PrevHash() [32]byte {
	if node.prev != nil {
		return node.prev.hash
	} else {
		return [32]byte{}
	}
}

func (node *Node) GetSuccessor(chain *NodeGraph, kind structures.ChildType) *Node {
	return chain.nodeFromHash[node.successorHashes[kind]]
}

func (node *Node) ExecutionPreconditionHash() [32]byte {
	vmProtoData := node.prev.vmProtoData
	beforeInbox := protocol.AddMessagesHashToInboxHash(value.NewEmptyTuple().Hash(), node.disputable.AssertionClaim.ImportedMessagesSlice)
	pre := &protocol.Precondition{
		BeforeHash:  vmProtoData.MachineHash,
		TimeBounds:  node.disputable.AssertionParams.TimeBounds,
		BeforeInbox: value.NewHashOnlyValue(beforeInbox, 0),
	}
	return pre.Hash()
}

func (node *Node) NodeDataHash(
	params structures.ChainParams,
) [32]byte {
	ret := [32]byte{}
	if node.disputable == nil {
		return ret
	}
	if node.linkType == structures.ValidChildType {
		copy(ret[:], solsha3.SoliditySHA3(
			solsha3.Bytes32(node.disputable.AssertionClaim.AssertionStub.LastMessageHashValue()),
			solsha3.Bytes32(node.disputable.AssertionClaim.AssertionStub.LastLogHashValue()),
		))
	} else {
		challengeDataHash, challengePeriodTicks := node.ChallengeNodeData(params)
		copy(ret[:], solsha3.SoliditySHA3(
			solsha3.Bytes32(challengeDataHash),
			solsha3.Uint256(challengePeriodTicks.Val),
		))
	}
	return ret
}

func (node *Node) ChallengeNodeData(
	params structures.ChainParams,
) ([32]byte, structures.TimeTicks) {
	ret := [32]byte{}
	vmProtoData := node.prev.vmProtoData

	switch node.linkType {
	case structures.InvalidPendingChildType:
		pendingLeft := new(big.Int).Add(vmProtoData.PendingCount, node.disputable.AssertionParams.ImportedMessageCount)
		pendingLeft = pendingLeft.Sub(node.disputable.MaxPendingCount, pendingLeft)
		copy(ret[:], solsha3.SoliditySHA3(
			solsha3.Bytes32(node.disputable.AssertionClaim.AfterPendingTop),
			solsha3.Bytes32(node.disputable.MaxPendingTop),
			solsha3.Uint256(pendingLeft),
		))
		challengePeriod := params.GracePeriod.Add(structures.TimeFromBlockNum(protocol.NewTimeBlocks(big.NewInt(1))))
		return ret, challengePeriod
	case structures.InvalidMessagesChildType:
		copy(ret[:], solsha3.SoliditySHA3(
			solsha3.Bytes32(vmProtoData.PendingTop),
			solsha3.Bytes32(node.disputable.AssertionClaim.AfterPendingTop),
			solsha3.Bytes32([32]byte{}),
			solsha3.Bytes32(node.disputable.AssertionClaim.ImportedMessagesSlice),
			solsha3.Uint256(node.disputable.AssertionParams.ImportedMessageCount),
		))
		challengePeriod := params.GracePeriod.Add(structures.TimeFromBlockNum(protocol.NewTimeBlocks(big.NewInt(1))))
		return ret, challengePeriod
	case structures.InvalidExecutionChildType:
		copy(ret[:], solsha3.SoliditySHA3(
			solsha3.Uint32(node.disputable.AssertionParams.NumSteps),
			solsha3.Bytes32(node.ExecutionPreconditionHash()),
			solsha3.Bytes32(node.disputable.AssertionClaim.AssertionStub.Hash()),
		))
		challengePeriod := params.GracePeriod.Add(node.disputable.CheckTime(params))
		return ret, challengePeriod
	default:
		panic("Unhandled challenge type")
	}
}

func (node *Node) setHash(
	nodeDataHash [32]byte,
) {
	var prevHashArr [32]byte
	if node.prev != nil {
		prevHashArr = node.prev.hash
	}
	innerHash := solsha3.SoliditySHA3(
		solsha3.Bytes32(node.vmProtoData.Hash()),
		solsha3.Int256(node.deadline),
		solsha3.Bytes32(nodeDataHash),
		solsha3.Int256(node.linkType),
	)
	hashSlice := solsha3.SoliditySHA3(
		solsha3.Bytes32(prevHashArr),
		solsha3.Bytes32(innerHash),
	)
	node.nodeDataHash = nodeDataHash
	copy(node.innerHash[:], innerHash)
	copy(node.hash[:], hashSlice)
}

func (node *Node) MarshalForCheckpoint(ctx structures.CheckpointContext) *NodeBuf {
	var machineHash *value.HashBuf
	if node.machine != nil {
		ctx.AddMachine(node.machine)
		machineHash = utils.MarshalHash(node.machine.Hash())
	}
	var prevHashBuf *value.HashBuf
	if node.prev != nil {
		prevHashBuf = utils.MarshalHash(node.prev.hash)
	}
	var disputableNodeBuf *structures.DisputableNodeBuf
	if node.disputable != nil {
		disputableNodeBuf = node.disputable.MarshalToBuf()
	}
	return &NodeBuf{
		PrevHash:       prevHashBuf,
		Deadline:       node.deadline.MarshalToBuf(),
		DisputableNode: disputableNodeBuf,
		LinkType:       uint32(node.linkType),
		VmProtoData:    node.vmProtoData.MarshalToBuf(),
		MachineHash:    machineHash,
		Depth:          node.depth,
		NodeDataHash:   utils.MarshalHash(node.nodeDataHash),
		InnerHash:      utils.MarshalHash(node.innerHash),
		Hash:           utils.MarshalHash(node.hash),
	}
}

func (m *NodeBuf) UnmarshalFromCheckpoint(ctx structures.RestoreContext, chain *NodeGraph) *Node {
	var disputableNode *structures.DisputableNode
	if m.DisputableNode != nil {
		disputableNode = m.DisputableNode.Unmarshal()
	}
	node := &Node{
		prev:         nil,
		deadline:     m.Deadline.Unmarshal(),
		disputable:   disputableNode,
		linkType:     structures.ChildType(m.LinkType),
		vmProtoData:  m.VmProtoData.Unmarshal(),
		machine:      nil,
		depth:        m.Depth,
		nodeDataHash: utils.UnmarshalHash(m.NodeDataHash),
		innerHash:    utils.UnmarshalHash(m.InnerHash),
		hash:         utils.UnmarshalHash(m.Hash),
	}

	if m.MachineHash != nil {
		node.machine = ctx.GetMachine(utils.UnmarshalHash(m.MachineHash))
	}

	chain.nodeFromHash[node.hash] = node

	// can't set up prev and successorHash fields yet; caller must do this later
	return node
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
	return append(sub, to.innerHash)
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

func (node *Node) EqualsFull(n2 *Node) bool {
	return node.Equals(n2) &&
		node.depth == n2.depth &&
		node.vmProtoData.Equals(n2.vmProtoData) &&
		node.linkType == n2.linkType &&
		node.successorHashes == n2.successorHashes &&
		node.numStakers == n2.numStakers
}

func CommonAncestor(n1, n2 *Node) *Node {
	n1, _, _ = GetConflictAncestor(n1, n2)
	return n1.prev
}

func GetConflictAncestor(n1, n2 *Node) (*Node, *Node, error) {
	n1Orig := n1
	n2Orig := n2
	for n1.depth > n2.depth {
		n1 = n1.prev
	}
	for n2.depth > n1.depth {
		n2 = n2.prev
	}

	// Now n1 and n2 are at the same height so we can start looking for a challenge

	for n1.prev != n2.prev {
		n1 = n1.prev
		n2 = n2.prev
	}

	if n1 == n1Orig || n1 == n2Orig {
		return n1, n2, errors.New("no conflict")
	}
	return n1, n2, nil
}
