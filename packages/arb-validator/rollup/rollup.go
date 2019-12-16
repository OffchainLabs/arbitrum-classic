/*
* Copyright 2019, Offchain Labs, Inc.
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
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"log"
	"math/big"
)

type Chain struct {
	rollupAddr      common.Address
	vmParams        ChainParams
	latestConfirmed *Node
	leaves          *LeafList
	nodeFromHash    map[[32]byte]*Node
	stakerList      *StakerList
	challenges      map[common.Address]*Challenge
}

func NewChain(_rollupAddr common.Address, _machine machine.Machine, _vmParams ChainParams) *Chain {
	ret := &Chain{
		rollupAddr: _rollupAddr,
		vmParams:   _vmParams,
		leaves:     NewLeafList(),
	}
	ret.CreateInitialNode(_machine)
	return ret
}

type LeafList struct {
	arr []*Node
	idx map[[32]byte]uint
}

func NewLeafList() *LeafList {
	return &LeafList{}
}

func (ll *LeafList) IsLeaf(node *Node) bool {
	_, ok := ll.idx[node.hash]
	return ok
}

func (ll *LeafList) Add(node *Node) {
	if _, ok := ll.idx[node.hash]; ok {
		log.Fatal("tried to insert leaf twice")
	}
	ll.idx[node.hash] = uint(len(ll.arr))
	ll.arr = append(ll.arr, node)
}

func (ll *LeafList) Delete(node *Node) {
	slot, ok := ll.idx[node.hash]
	if !ok {
		log.Fatal("tried to remove nonexistent leaf")
	}
	if int(slot) < len(ll.arr)-1 {
		ll.arr[slot] = ll.arr[len(ll.arr)-1]
		ll.idx[ll.arr[slot].hash] = slot
	}
	ll.arr = ll.arr[:len(ll.arr)-1]
}

type Staker struct {
	address      common.Address
	location     *Node
	creationTime *big.Int
	challenge    *Challenge
}

type StakerList struct {
	arr []*Staker
	idx map[common.Address]uint
}

func NewStakerList() *StakerList {
	return &StakerList{}
}

func (sl *StakerList) Add(newStaker *Staker) {
	if _, ok := sl.idx[newStaker.address]; ok {
		log.Fatal("tried to insert staker twice")
	}
	sl.idx[newStaker.address] = uint(len(sl.arr))
	sl.arr = append(sl.arr, newStaker)
}

func (sl *StakerList) Delete(staker *Staker) {
	slot, ok := sl.idx[staker.address]
	if !ok {
		log.Fatal("tried to remove nonexistent staker")
	}
	if int(slot) < len(sl.arr)-1 {
		sl.arr[slot] = sl.arr[len(sl.arr)-1]
		sl.idx[sl.arr[slot].address] = slot
	}
	sl.arr = sl.arr[:len(sl.arr)-1]
}

func (sl *StakerList) Get(addr common.Address) *Staker {
	return sl.arr[sl.idx[addr]]
}

type ChainParams struct {
	stakeRequirement  *big.Int
	gracePeriod       uint32
	maxExecutionSteps uint32
	pendingInbox      *PendingInbox
}

type PendingInbox struct {
	//TODO
}

type DisputableNode struct {
	hash     [32]byte
	deadline *big.Int
}

type Node struct {
	hash            [32]byte
	disputable      *DisputableNode
	machineHash     [32]byte
	machine         machine.Machine // nil if unknown
	inboxHash       [32]byte
	inbox           value.Value // nil if unknown
	prev            *Node
	linkType        ChildType
	hasSuccessors   bool
	successorHashes [MaxChildType + 1][32]byte
}

type ChildType uint

const (
	ValidChildType            ChildType = 0
	InvalidPendingChildType   ChildType = 1
	InvalidMessagesChildType  ChildType = 2
	InvalidExecutionChildType ChildType = 3

	MinChildType        ChildType = 0
	MinInvalidChildType ChildType = 1
	MaxChildType        ChildType = 3
)

var zeroBytes32 [32]byte // deliberately zeroed

func (chain *Chain) CreateInitialNode(machine machine.Machine) {
	newNode := &Node{
		machineHash: machine.Hash(),
		machine:     machine.Clone(),
		inboxHash:   value.NewEmptyTuple().Hash(),
		inbox:       value.NewEmptyTuple(),
		linkType:    ValidChildType,
	}
	newNode.setHash()
	chain.leaves.Add(newNode)
	chain.latestConfirmed = newNode
}

func (chain *Chain) CreateNodesOnAssert(
	prevNode *Node,
	dispNode *DisputableNode,
	afterMachineHash [32]byte,
	afterMachine machine.Machine, // if known
	afterInboxHash [32]byte,
	afterInbox value.Value, // if known
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
		disputable:  dispNode,
		prev:        prevNode,
		linkType:    ValidChildType,
		machineHash: afterMachineHash,
		machine:     afterMachine,
		inboxHash:   afterInboxHash,
		inbox:       afterInbox,
	}
	newNode.setHash()
	prevNode.successorHashes[ValidChildType] = newNode.hash
	chain.leaves.Add(newNode)

	// create nodes for invalid branches
	for kind := MinInvalidChildType; kind <= MaxChildType; kind++ {
		newNode := &Node{
			disputable:  dispNode,
			prev:        prevNode,
			linkType:    kind,
			machineHash: prevNode.machineHash,
			machine:     prevNode.machine,
			inboxHash:   prevNode.inboxHash,
			inbox:       prevNode.inbox,
		}
		newNode.setHash()
		prevNode.successorHashes[kind] = newNode.hash
		chain.leaves.Add(newNode)
	}
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
		node.inboxHash,
	)
	var ret [32]byte
	copy(ret[:], retSlice)
	return ret
}

func (node *Node) removePrev() {
	oldNode := node.prev
	node.prev = nil // so garbage collector doesn't preserve prev anymore
	if oldNode != nil {
		oldNode.successorHashes[node.linkType] = zeroBytes32
		oldNode.considerRemoving()
	}
}

func (node *Node) considerRemoving() {
	for kind := MinChildType; kind <= MaxChildType; kind++ {
		if node.successorHashes[kind] != zeroBytes32 {
			return
		}
	}
	node.removePrev()
}

func (chain *Chain) ConfirmNode(nodeHash [32]byte) {
	node := chain.nodeFromHash[nodeHash]
	chain.latestConfirmed = node
	node.removePrev()
}

func (chain *Chain) PruneNode(nodeHash [32]byte) {
	node := chain.nodeFromHash[nodeHash]
	delete(chain.nodeFromHash, nodeHash)
	node.removePrev()
}

func (chain *Chain) CreateStake(stakerAddr common.Address, nodeHash [32]byte, creationTime *big.Int) {
	staker := &Staker{
		stakerAddr,
		chain.nodeFromHash[nodeHash],
		creationTime,
		nil,
	}
	chain.stakerList.Add(staker)
}

func (chain *Chain) MoveStake(stakerAddr common.Address, nodeHash [32]byte) {
	chain.stakerList.Get(stakerAddr).location = chain.nodeFromHash[nodeHash]
}

func (chain *Chain) RemoveStake(stakerAddr common.Address) {
	chain.stakerList.Delete(chain.stakerList.Get(stakerAddr))
}

func (staker *Staker) MarshalToBuf() *StakerBuf {
	challengeStr := ""
	if staker.challenge != nil {
		challengeStr = string(staker.challenge.contract.Bytes())
	}
	return &StakerBuf{
		Address:       string(staker.address.Bytes()),
		Location:      string(string(staker.location.hash[:])),
		CreationTime:  string(staker.creationTime.Bytes()),
		ChallengeAddr: challengeStr,
	}
}

func (buf *StakerBuf) Unmarshal(chain *Chain) *Staker {
	// chain.nodeFromHash and chain.challenges must have already been unmarshaled
	var locArr [32]byte
	copy(locArr[:], []byte(buf.Location))
	if buf.ChallengeAddr == "" {
		return &Staker{
			address:      common.BytesToAddress([]byte(buf.Address)),
			location:     chain.nodeFromHash[locArr],
			creationTime: new(big.Int).SetBytes([]byte(buf.CreationTime)),
			challenge:    nil,
		}
	} else {
		return &Staker{
			address:      common.BytesToAddress([]byte(buf.Address)),
			location:     chain.nodeFromHash[locArr],
			creationTime: new(big.Int).SetBytes([]byte(buf.CreationTime)),
			challenge:    chain.challenges[common.BytesToAddress([]byte(buf.ChallengeAddr))],
		}
	}
}

type Challenge struct {
	contract   common.Address
	asserter   common.Address
	challenger common.Address
	kind       ChallengeType
}

type ChallengeType uint

const (
	InvalidPendingTopChallenge ChallengeType = 1
	InvalidMessagesChallenge   ChallengeType = 2
	InvalidExecutionChallenge  ChallengeType = 3
)

func (chain *Chain) NewChallenge(contract, asserter, challenger common.Address, kind ChallengeType) *Challenge {
	ret := &Challenge{contract, asserter, challenger, kind}
	chain.challenges[contract] = ret
	chain.stakerList.Get(asserter).challenge = ret
	chain.stakerList.Get(challenger).challenge = ret
	return ret
}

func (chain *Chain) ChallengeResolved(contract, winner, loser common.Address) {
	chain.RemoveStake(loser)
	delete(chain.challenges, contract)
}

func (chal *Challenge) MarshalToBuf() *ChallengeBuf {
	return &ChallengeBuf{
		Contract:   string(chal.contract.Bytes()),
		Asserter:   string(chal.asserter.Bytes()),
		Challenger: string(chal.challenger.Bytes()),
		Kind:       uint32(chal.kind),
	}
}

func (buf *ChallengeBuf) Unmarshal() *Challenge {
	return &Challenge{
		common.BytesToAddress([]byte(buf.Contract)),
		common.BytesToAddress([]byte(buf.Asserter)),
		common.BytesToAddress([]byte(buf.Challenger)),
		ChallengeType(buf.Kind),
	}
}
