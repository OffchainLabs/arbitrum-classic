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

package structures

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

var zeroBytes32 common.Hash // deliberately zeroed

type Node struct {
	prevHash    common.Hash
	prev        *Node // Node with hash prevHash if non-nil
	deadline    common.TimeTicks
	disputable  *valprotocol.DisputableNode
	linkType    valprotocol.ChildType
	vmProtoData *valprotocol.VMProtoData

	machine      machine.Machine              // nil if unknown
	assertion    *protocol.ExecutionAssertion // nil if not valid node or unknown
	depth        uint64
	nodeDataHash common.Hash
	innerHash    common.Hash
	hash         common.Hash

	successorHashes [valprotocol.MaxChildType + 1]common.Hash
	numStakers      uint64
}

func (node *Node) String() string {
	return fmt.Sprintf("Node(type: %v, disputable: %v, deadline: %v, protodata: %v)", node.linkType, node.disputable, node.deadline.Val, node.vmProtoData)
}

func NewInitialNode(mach machine.Machine) *Node {
	ret := &Node{
		prevHash:   common.Hash{},
		prev:       nil,
		deadline:   common.TimeTicks{Val: big.NewInt(0)},
		disputable: nil,
		linkType:   0,
		vmProtoData: valprotocol.NewVMProtoData(
			mach.Hash(),
			common.Hash{},
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(0),
		),
		machine: mach,
		depth:   0,
	}
	ret.setHash(common.Hash{})
	return ret
}

func NewValidNodeFromPrev(
	prev *Node,
	disputable *valprotocol.DisputableNode,
	params valprotocol.ChainParams,
	currentTime *common.TimeBlocks,
) *Node {
	return NewNodeFromPrev(
		prev,
		disputable,
		valprotocol.ValidChildType,
		params,
		currentTime,
		disputable.ValidAfterVMProtoData(prev.vmProtoData),
	)
}

func NewRandomNodeFromValidPrev(prev *Node, inboxStack *MessageStack, messageCount uint64) *Node {
	assertion := protocol.NewExecutionAssertionFromValues(
		common.RandHash(),
		common.RandHash(),
		rand.Uint64(),
		messageCount,
		[]value.Value{value.NewInt64Value(0), value.NewInt64Value(2)},
		[]value.Value{value.NewInt64Value(1), value.NewInt64Value(2)},
	)

	disputableNode := valprotocol.NewRandomDisputableNode(
		NewExecutionAssertionStubFromWholeAssertion(assertion, prev.VMProtoData().InboxTop, inboxStack),
	)
	nextNode := NewValidNodeFromPrev(
		prev,
		disputableNode,
		valprotocol.NewRandomChainParams(),
		common.NewTimeBlocks(common.RandBigInt()),
	)

	_ = nextNode.UpdateValidOpinion(nil, assertion)
	return nextNode
}

func NewRandomInvalidNodeFromValidPrev(
	prev *Node,
	stub *valprotocol.ExecutionAssertionStub,
	kind valprotocol.ChildType,
	params valprotocol.ChainParams,
) *Node {
	disputableNode := valprotocol.NewRandomDisputableNode(stub)

	nextNode := NewInvalidNodeFromPrev(
		prev,
		disputableNode,
		kind,
		params,
		common.NewTimeBlocks(common.RandBigInt()),
	)

	_ = nextNode.UpdateInvalidOpinion()
	return nextNode
}

func NewInvalidNodeFromPrev(
	prev *Node,
	disputable *valprotocol.DisputableNode,
	kind valprotocol.ChildType,
	params valprotocol.ChainParams,
	currentTime *common.TimeBlocks,
) *Node {
	return NewNodeFromPrev(
		prev,
		disputable,
		kind,
		params,
		currentTime,
		prev.vmProtoData,
	)
}

func NewNodeFromPrev(
	prev *Node,
	disputable *valprotocol.DisputableNode,
	kind valprotocol.ChildType,
	params valprotocol.ChainParams,
	currentTime *common.TimeBlocks,
	vmProtoData *valprotocol.VMProtoData,
) *Node {
	deadlineTicks := valprotocol.CalculateNodeDeadline(
		disputable.Assertion,
		params,
		prev.deadline,
		common.TicksFromBlockNum(currentTime),
	)
	ret := &Node{
		prevHash:    prev.hash,
		prev:        prev,
		deadline:    deadlineTicks,
		disputable:  disputable,
		linkType:    kind,
		vmProtoData: vmProtoData,
		depth:       prev.depth + 1,
	}
	ret.setHash(ret.calculateNodeDataHash(params))
	return ret
}

func (node *Node) LinkSuccessor(successor *Node) error {
	if successor.prevHash != node.hash {
		return errors.New("node is not successor")
	}
	node.successorHashes[successor.linkType] = successor.hash
	return nil
}

func (node *Node) Hash() common.Hash {
	return node.hash
}

func (node *Node) LinkType() valprotocol.ChildType {
	return node.linkType
}

func (node *Node) PrevHash() common.Hash {
	return node.prevHash
}

func (node *Node) Prev() *Node {
	return node.prev
}

func (node *Node) ClearPrev() {
	node.prev = nil
	node.prevHash = zeroBytes32
}

func (node *Node) UnlinkPrev() bool {
	hasPrev := node.prev != nil
	if hasPrev {
		node.prev.successorHashes[node.LinkType()] = zeroBytes32
		node.ClearPrev()
	}
	return hasPrev
}

func (node *Node) Deadline() common.TimeTicks {
	return node.deadline
}

func (node *Node) Disputable() *valprotocol.DisputableNode {
	return node.disputable
}

func (node *Node) VMProtoData() *valprotocol.VMProtoData {
	return node.vmProtoData
}

func (node *Node) Machine() machine.Machine {
	return node.machine
}

func (node *Node) Assertion() *protocol.ExecutionAssertion {
	return node.assertion
}

func (node *Node) UpdateValidOpinion(machine machine.Machine, assertion *protocol.ExecutionAssertion) error {
	if node.linkType != valprotocol.ValidChildType {
		return errors.New("node is invalid")
	}
	node.machine = machine
	node.assertion = assertion
	return nil
}

func (node *Node) UpdateInvalidOpinion() error {
	if node.linkType == valprotocol.ValidChildType {
		return errors.New("node is valid")
	}
	node.machine = node.prev.machine.Clone()
	return nil
}

func (node *Node) Depth() uint64 {
	return node.depth
}

func (node *Node) NodeDataHash() common.Hash {
	return node.nodeDataHash
}

func (node *Node) SuccessorHashes() [valprotocol.MaxChildType + 1]common.Hash {
	return node.successorHashes
}

func (node *Node) NumStakers() uint64 {
	return node.numStakers
}

func (node *Node) AddStaker() {
	node.numStakers++
}

func (node *Node) RemoveStaker() {
	node.numStakers--
}

func (node *Node) HasAncestor() bool {
	emptyHash := common.Hash{}
	return node.prevHash != emptyHash
}

func (node *Node) Equals(node2 *Node) bool {
	return node.hash == node2.hash
}

func (node *Node) calculateNodeDataHash(params valprotocol.ChainParams) common.Hash {
	if node.disputable == nil {
		return common.Hash{}
	}
	if node.linkType == valprotocol.ValidChildType {
		return hashing.SoliditySHA3(
			hashing.Uint256(node.prev.VMProtoData().MessageCount),
			hashing.Bytes32(node.disputable.Assertion.LastMessageHash),
			hashing.Bytes32(node.disputable.Assertion.LastLogHash),
		)
	} else {
		challengeDataHash, challengePeriodTicks := node.ChallengeNodeData(params)
		return hashing.SoliditySHA3(
			hashing.Bytes32(challengeDataHash),
			hashing.TimeTicks(challengePeriodTicks),
		)
	}
}

func (node *Node) ChallengeNodeData(params valprotocol.ChainParams) (common.Hash, common.TimeTicks) {
	vmProtoData := node.prev.vmProtoData
	switch node.linkType {
	case valprotocol.InvalidInboxTopChildType:
		inboxLeft := new(big.Int).Add(vmProtoData.InboxCount, node.disputable.AssertionParams.ImportedMessageCount)
		inboxLeft = inboxLeft.Sub(node.disputable.MaxInboxCount, inboxLeft)
		ret := valprotocol.InboxTopChallengeDataHash(
			node.disputable.Assertion.AfterInboxHash,
			node.disputable.MaxInboxTop,
			inboxLeft,
		)
		challengePeriod := params.GracePeriod.Add(common.TicksFromBlockNum(common.NewTimeBlocks(big.NewInt(1))))
		return ret, challengePeriod
	case valprotocol.InvalidExecutionChildType:
		ret := valprotocol.ExecutionDataHash(
			node.disputable.AssertionParams.NumSteps,
			node.disputable.Assertion,
		)
		challengePeriod := params.GracePeriod.Add(node.disputable.Assertion.CheckTime(params))
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

func Link(nd *Node, prev *Node) error {
	if nd.prevHash != prev.hash {
		return errors.New("node is not parent")
	}
	nd.prev = prev
	prev.successorHashes[nd.linkType] = nd.hash
	return nil
}

func (node *Node) MarshalForCheckpoint(ctx *ckptcontext.CheckpointContext, includeMachine bool) *NodeBuf {
	var machineHash *common.HashBuf
	if includeMachine && node.machine != nil {
		ctx.AddMachine(node.machine)
		machineHash = node.machine.Hash().MarshalToBuf()
	}
	var disputableNodeBuf *valprotocol.DisputableNodeBuf
	if node.disputable != nil {
		disputableNodeBuf = node.disputable.MarshalToBuf()
	}

	return &NodeBuf{
		PrevHash:       node.prevHash.MarshalToBuf(),
		Deadline:       node.deadline.MarshalToBuf(),
		DisputableNode: disputableNodeBuf,
		LinkType:       uint32(node.linkType),
		VmProtoData:    node.vmProtoData.MarshalToBuf(),
		MachineHash:    machineHash,
		Assertion:      node.assertion,
		Depth:          node.depth,
		NodeDataHash:   node.nodeDataHash.MarshalToBuf(),
		InnerHash:      node.innerHash.MarshalToBuf(),
		Hash:           node.hash.MarshalToBuf(),
	}
}

func (x *NodeBuf) UnmarshalFromCheckpoint(ctx ckptcontext.RestoreContext) (*Node, error) {
	var disputableNode *valprotocol.DisputableNode
	if x.DisputableNode != nil {
		disputableNode = x.DisputableNode.Unmarshal()
	}

	node := &Node{
		prevHash:     x.PrevHash.Unmarshal(),
		prev:         nil,
		deadline:     x.Deadline.Unmarshal(),
		disputable:   disputableNode,
		linkType:     valprotocol.ChildType(x.LinkType),
		vmProtoData:  x.VmProtoData.Unmarshal(),
		machine:      nil,
		assertion:    x.Assertion,
		depth:        x.Depth,
		nodeDataHash: x.NodeDataHash.Unmarshal(),
		innerHash:    x.InnerHash.Unmarshal(),
		hash:         x.Hash.Unmarshal(),
		numStakers:   0,
	}

	valueCache, err := cmachine.NewValueCache()
	if err != nil {
		return nil, err
	}
	defer cmachine.DestroyValueCache(valueCache)

	if x.MachineHash != nil {
		node.machine = ctx.GetMachine(x.MachineHash.Unmarshal(), valueCache)
	}

	// can't set up prev and successorHash fields yet; caller must do this later
	return node, nil
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

func (node *Node) EqualsFull(n2 *Node) bool {
	return node.Equals(n2) &&
		node.depth == n2.depth &&
		node.vmProtoData.Equals(n2.vmProtoData) &&
		node.linkType == n2.linkType &&
		node.successorHashes == n2.successorHashes &&
		node.numStakers == n2.numStakers
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
