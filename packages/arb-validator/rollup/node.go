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
	"fmt"
	"log"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type Node struct {
	prev        *Node
	deadline    common.TimeTicks
	disputable  *structures.DisputableNode
	linkType    structures.ChildType
	vmProtoData *structures.VMProtoData

	machine         machine.Machine              // nil if unknown
	assertion       *protocol.ExecutionAssertion // nil if not valid node or unknown
	depth           uint64
	nodeDataHash    common.Hash
	innerHash       common.Hash
	hash            common.Hash
	assertionTxHash common.Hash

	successorHashes [structures.MaxChildType + 1]common.Hash
	numStakers      uint64
}

func (n *Node) String() string {
	return fmt.Sprintf("Node(type: %v, disputable: %v, deadline: %v, protodata: %v)", n.linkType, n.disputable, n.deadline.Val, n.vmProtoData)
}

func NewInitialNode(mach machine.Machine) *Node {
	ret := &Node{
		prev:       nil,
		deadline:   common.TimeTicks{big.NewInt(0)},
		disputable: nil,
		linkType:   0,
		vmProtoData: structures.NewVMProtoData(
			mach.Hash(),
			value.NewEmptyTuple().Hash(),
			big.NewInt(0),
		),
		machine: mach,
		depth:   0,
	}
	ret.setHash(common.Hash{})
	return ret
}

func newInitialNode_hashOnly(machHash common.Hash) *Node {
	ret := &Node{
		prev:       nil,
		deadline:   common.TimeTicks{big.NewInt(0)},
		disputable: nil,
		linkType:   0,
		vmProtoData: structures.NewVMProtoData(
			machHash,
			value.NewEmptyTuple().Hash(),
			big.NewInt(0),
		),
		depth: 0,
	}
	ret.setHash(common.Hash{})
	return ret
}

func MakeInitialNodeBuf(machineHash common.Hash) *NodeBuf {
	initNode := newInitialNode_hashOnly(machineHash)
	nodeBuf := initNode.MarshalForCheckpoint(nil)
	nodeBuf.MachineHash = machineHash.MarshalToBuf()
	return nodeBuf
}

func NewNodeFromValidPrev(
	prev *Node,
	disputable *structures.DisputableNode,
	params structures.ChainParams,
	currentTime *common.TimeBlocks,
	assertionTxHash common.Hash,
) *Node {
	return NewNodeFromPrev(
		prev,
		disputable,
		structures.ValidChildType,
		params,
		currentTime,
		disputable.ValidAfterVMProtoData(prev.vmProtoData),
		assertionTxHash,
	)
}

func NewNodeFromInvalidPrev(
	prev *Node,
	disputable *structures.DisputableNode,
	kind structures.ChildType,
	params structures.ChainParams,
	currentTime *common.TimeBlocks,
	assertionTxHash common.Hash,
) *Node {
	return NewNodeFromPrev(
		prev,
		disputable,
		kind,
		params,
		currentTime,
		prev.vmProtoData,
		assertionTxHash,
	)
}

func NewNodeFromPrev(
	prev *Node,
	disputable *structures.DisputableNode,
	kind structures.ChildType,
	params structures.ChainParams,
	currentTime *common.TimeBlocks,
	vmProtoData *structures.VMProtoData,
	assertionTxHash common.Hash,
) *Node {
	checkTime := disputable.CheckTime(params)
	deadlineTicks := common.TimeFromBlockNum(currentTime).Add(params.GracePeriod)
	if deadlineTicks.Cmp(prev.deadline) >= 0 {
		deadlineTicks = deadlineTicks.Add(checkTime)
	} else {
		deadlineTicks = prev.deadline.Add(checkTime)
	}

	ret := &Node{
		prev:            prev,
		deadline:        deadlineTicks,
		disputable:      disputable,
		linkType:        kind,
		vmProtoData:     vmProtoData,
		depth:           prev.depth + 1,
		assertionTxHash: assertionTxHash,
	}
	ret.setHash(ret.NodeDataHash(params))
	prev.successorHashes[kind] = ret.hash
	return ret
}

func (node *Node) Equals(node2 *Node) bool {
	return node.hash == node2.hash
}

func (node *Node) PrevHash() common.Hash {
	if node.prev != nil {
		return node.prev.hash
	} else {
		return common.Hash{}
	}
}

func (node *Node) GetSuccessor(chain *NodeGraph, kind structures.ChildType) *Node {
	return chain.nodeFromHash[node.successorHashes[kind]]
}

func (node *Node) ExecutionPreconditionHash() common.Hash {
	vmProtoData := node.prev.vmProtoData
	pre := &valprotocol.Precondition{
		BeforeHash:  vmProtoData.MachineHash,
		TimeBounds:  node.disputable.AssertionParams.TimeBounds,
		BeforeInbox: value.NewHashOnlyValue(node.disputable.AssertionClaim.ImportedMessagesSlice, 0),
	}
	return pre.Hash()
}

func (node *Node) NodeDataHash(params structures.ChainParams) common.Hash {
	if node.disputable == nil {
		return common.Hash{}
	}
	if node.linkType == structures.ValidChildType {
		return hashing.SoliditySHA3(
			hashing.Bytes32(node.disputable.AssertionClaim.AssertionStub.LastMessageHash),
			hashing.Bytes32(node.disputable.AssertionClaim.AssertionStub.LastLogHash),
		)
	} else {
		challengeDataHash, challengePeriodTicks := node.ChallengeNodeData(params)
		return hashing.SoliditySHA3(
			hashing.Bytes32(challengeDataHash),
			hashing.TimeTicks(challengePeriodTicks),
		)
	}
}

func (node *Node) ChallengeNodeData(params structures.ChainParams) (common.Hash, common.TimeTicks) {
	vmProtoData := node.prev.vmProtoData
	switch node.linkType {
	case structures.InvalidPendingChildType:
		pendingLeft := new(big.Int).Add(vmProtoData.PendingCount, node.disputable.AssertionParams.ImportedMessageCount)
		pendingLeft = pendingLeft.Sub(node.disputable.MaxPendingCount, pendingLeft)
		ret := structures.PendingTopChallengeDataHash(
			node.disputable.AssertionClaim.AfterPendingTop,
			node.disputable.MaxPendingTop,
			pendingLeft,
		)
		challengePeriod := params.GracePeriod.Add(common.TimeFromBlockNum(common.NewTimeBlocks(big.NewInt(1))))
		return ret, challengePeriod
	case structures.InvalidMessagesChildType:
		ret := structures.MessageChallengeDataHash(
			vmProtoData.PendingTop,
			node.disputable.AssertionClaim.AfterPendingTop,
			value.NewEmptyTuple().Hash(),
			node.disputable.AssertionClaim.ImportedMessagesSlice,
			node.disputable.AssertionParams.ImportedMessageCount,
		)
		challengePeriod := params.GracePeriod.Add(common.TimeFromBlockNum(common.NewTimeBlocks(big.NewInt(1))))
		return ret, challengePeriod
	case structures.InvalidExecutionChildType:
		ret := structures.ExecutionDataHash(
			node.disputable.AssertionParams.NumSteps,
			node.ExecutionPreconditionHash(),
			node.disputable.AssertionClaim.AssertionStub.Hash(),
		)
		challengePeriod := params.GracePeriod.Add(node.disputable.CheckTime(params))
		return ret, challengePeriod
	default:
		log.Fatal("Unhandled challenge type", node.linkType)
		return common.Hash{}, common.TimeTicks{}
	}
}

func (node *Node) setHash(nodeDataHash common.Hash) {
	var prevHashArr common.Hash
	if node.prev != nil {
		prevHashArr = node.prev.hash
	}
	innerHash := hashing.SoliditySHA3(
		hashing.Bytes32(node.vmProtoData.Hash()),
		hashing.TimeTicks(node.deadline),
		hashing.Bytes32(nodeDataHash),
		hashing.Uint256(new(big.Int).SetUint64(uint64(node.linkType))),
	)
	hash := hashing.SoliditySHA3(
		hashing.Bytes32(prevHashArr),
		hashing.Bytes32(innerHash),
	)
	node.nodeDataHash = nodeDataHash
	node.innerHash = innerHash
	node.hash = hash
}

func (node *Node) MarshalForCheckpoint(ctx structures.CheckpointContext) *NodeBuf {
	var machineHash *common.HashBuf
	if node.machine != nil {
		ctx.AddMachine(node.machine)
		machineHash = node.machine.Hash().MarshalToBuf()
	}
	var prevHashBuf *common.HashBuf
	if node.prev != nil {
		prevHashBuf = node.prev.hash.MarshalToBuf()
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
		NodeDataHash:   node.nodeDataHash.MarshalToBuf(),
		InnerHash:      node.innerHash.MarshalToBuf(),
		Hash:           node.hash.MarshalToBuf(),
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
		depth:        m.Depth,
		nodeDataHash: m.NodeDataHash.Unmarshal(),
		innerHash:    m.InnerHash.Unmarshal(),
		hash:         m.Hash.Unmarshal(),
	}

	if m.MachineHash != nil {
		node.machine = ctx.GetMachine(m.MachineHash.Unmarshal())
	}

	chain.nodeFromHash[node.hash] = node

	// can't set up prev and successorHash fields yet; caller must do this later
	return node
}

func GeneratePathProof(from, to *Node) []common.Hash {
	// returns nil if no proof exists
	if to == nil {
		return nil
	}
	if from == to {
		return []common.Hash{}
	}
	sub := GeneratePathProof(from, to.prev)
	if sub == nil {
		return nil
	}
	return append(sub, to.innerHash)
}

func GenerateConflictProof(from, to1, to2 *Node) ([]common.Hash, []common.Hash) {
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

	if n1.linkType == n2.linkType {
		return n1, n2, errors.New("no conflict")
	}
	return n1, n2, nil
}
