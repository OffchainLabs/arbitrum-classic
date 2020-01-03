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
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

//go:generate protoc -I.. -I. --go_out=paths=source_relative:. rollup.proto

type Chain struct {
	rollupAddr      common.Address
	vmParams        ChainParams
	pendingInbox    *PendingInbox
	latestConfirmed *Node
	leaves          *LeafSet
	nodeFromHash    map[[32]byte]*Node
	stakers         *StakerSet
	challenges      map[common.Address]*Challenge
}

func NewChain(_rollupAddr common.Address, _machine machine.Machine, _vmParams ChainParams) *Chain {
	ret := &Chain{
		_rollupAddr,
		_vmParams,
		NewPendingInbox(),
		nil,
		NewLeafSet(),
		make(map[[32]byte]*Node),
		NewStakerSet(),
		make(map[common.Address]*Challenge),
	}
	ret.CreateInitialNode(_machine)
	return ret
}

func (chain *Chain) MarshalToBuf() *ChainBuf {
	var allNodes []*NodeBuf
	for _, v := range chain.nodeFromHash {
		allNodes = append(allNodes, v.MarshalToBuf())
	}
	var allStakers []*StakerBuf
	chain.stakers.forall(func(staker *Staker) {
		allStakers = append(allStakers, staker.MarshalToBuf())
	})
	var leafHashes []string
	chain.leaves.forall(func(node *Node) {
		leafHashes = append(leafHashes, string(node.hash[:]))
	})
	var allChallenges []*ChallengeBuf
	for _, v := range chain.challenges {
		allChallenges = append(allChallenges, v.MarshalToBuf())
	}
	return &ChainBuf{
		ContractAddress:     string(chain.rollupAddr.Bytes()),
		VmParams:            chain.vmParams.MarshalToBuf(),
		PendingInbox:        chain.pendingInbox.MarshalToBuf(),
		Nodes:               allNodes,
		LatestConfirmedHash: string(chain.latestConfirmed.hash[:]),
		LeafHashes:          leafHashes,
		Stakers:             allStakers,
		Challenges:          allChallenges,
	}
}

func (buf *ChainBuf) Unmarshal() *Chain {
	chain := &Chain{
		common.BytesToAddress([]byte(buf.ContractAddress)),
		buf.VmParams.Unmarshal(),
		buf.PendingInbox.Unmarshal(),
		nil,
		NewLeafSet(),
		make(map[[32]byte]*Node),
		NewStakerSet(),
		make(map[common.Address]*Challenge),
	}
	for _, chalBuf := range buf.Challenges {
		chal := &Challenge{
			common.BytesToAddress([]byte(chalBuf.Contract)),
			common.BytesToAddress([]byte(chalBuf.Asserter)),
			common.BytesToAddress([]byte(chalBuf.Challenger)),
		}
		chain.challenges[chal.contract] = chal
	}
	for _, nodeBuf := range buf.Nodes {
		var nodeHash [32]byte
		copy(nodeHash[:], []byte(nodeBuf.Hash))
		node := chain.nodeFromHash[nodeHash]
		var prevHash [32]byte
		copy(prevHash[:], []byte(nodeBuf.PrevHash))
		if prevHash != zeroBytes32 {
			prev := chain.nodeFromHash[prevHash]
			node.prev = prev
			prev.successorHashes[node.linkType] = nodeHash
		}
	}
	for _, leafHashStr := range buf.LeafHashes {
		var leafHash [32]byte
		copy(leafHash[:], []byte(leafHashStr))
		chain.leaves.Add(chain.nodeFromHash[leafHash])
	}
	for _, stakerBuf := range buf.Stakers {
		var locationHash [32]byte
		copy(locationHash[:], []byte(stakerBuf.Location))
		chain.stakers.Add(&Staker{
			common.BytesToAddress([]byte(stakerBuf.Address)),
			chain.nodeFromHash[locationHash],
			stakerBuf.CreationTime.Unmarshal(),
			chain.challenges[common.BytesToAddress([]byte(stakerBuf.ChallengeAddr))],
		})
	}
	var lcHash [32]byte
	copy(lcHash[:], []byte(buf.LatestConfirmedHash))
	chain.latestConfirmed = chain.nodeFromHash[lcHash]

	return chain
}

type LeafSet struct {
	idx map[[32]byte]*Node
}

func NewLeafSet() *LeafSet {
	return &LeafSet{
		make(map[[32]byte]*Node),
	}
}

func (ll *LeafSet) IsLeaf(node *Node) bool {
	_, ok := ll.idx[node.hash]
	return ok
}

func (ll *LeafSet) Add(node *Node) {
	if ll.IsLeaf(node) {
		log.Fatal("tried to insert leaf twice")
	}
	ll.idx[node.hash] = node
}

func (ll *LeafSet) Delete(node *Node) {
	delete(ll.idx, node.hash)
}

func (ll *LeafSet) forall(f func(*Node)) {
	for _, v := range ll.idx {
		f(v)
	}
}

type Staker struct {
	address      common.Address
	location     *Node
	creationTime RollupTime
	challenge    *Challenge
}

type StakerSet struct {
	idx map[common.Address]*Staker
}

func NewStakerSet() *StakerSet {
	return &StakerSet{make(map[common.Address]*Staker)}
}

func (sl *StakerSet) Add(newStaker *Staker) {
	if _, ok := sl.idx[newStaker.address]; ok {
		log.Fatal("tried to insert staker twice")
	}
	sl.idx[newStaker.address] = newStaker
}

func (sl *StakerSet) Delete(staker *Staker) {
	delete(sl.idx, staker.address)
}

func (sl *StakerSet) Get(addr common.Address) *Staker {
	return sl.idx[addr]
}

func (sl *StakerSet) forall(f func(*Staker)) {
	for _, v := range sl.idx {
		f(v)
	}
}

type ChainParams struct {
	stakeRequirement  *big.Int
	gracePeriod       RollupTime
	maxExecutionSteps uint32
}

func (params *ChainParams) MarshalToBuf() *ChainParamsBuf {
	return &ChainParamsBuf{
		StakeRequirement:  string(params.stakeRequirement.Bytes()),
		GracePeriod:       params.gracePeriod.MarshalToBuf(),
		MaxExecutionSteps: params.maxExecutionSteps,
	}
}

func (buf *ChainParamsBuf) Unmarshal() ChainParams {
	return ChainParams{
		new(big.Int).SetBytes([]byte(buf.StakeRequirement)),
		buf.GracePeriod.Unmarshal(),
		buf.MaxExecutionSteps,
	}
}

type DisputableNode struct {
	hash           [32]byte
	pendingTopHash [32]byte
	deadline       RollupTime
}

func (dn *DisputableNode) MarshalToBuf() *DisputableNodeBuf {
	return &DisputableNodeBuf{
		Hash:       string(dn.hash[:]),
		PendingTop: string(dn.pendingTopHash[:]),
		Deadline:   dn.deadline.MarshalToBuf(),
	}
}

func (buf *DisputableNodeBuf) Unmarshal() *DisputableNode {
	var hashBuf [32]byte
	copy(hashBuf[:], []byte(buf.Hash))
	var pthBuf [32]byte
	copy(pthBuf[:], []byte(buf.PendingTop))
	return &DisputableNode{
		hash:           hashBuf,
		pendingTopHash: pthBuf,
		deadline:       buf.Deadline.Unmarshal(),
	}
}

type Node struct {
	hash            [32]byte
	disputable      *DisputableNode
	machineHash     [32]byte
	machine         machine.Machine // nil if unknown
	pendingTopHash  [32]byte
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
		machineHash:    machine.Hash(),
		machine:        machine.Clone(),
		pendingTopHash: value.NewEmptyTuple().Hash(),
		linkType:       ValidChildType,
	}
	newNode.setHash()
	chain.leaves.Add(newNode)
	chain.latestConfirmed = newNode
}

func (chain *Chain) notifyNewBlockNumber(blockNum *big.Int) {
	//TODO: checkpoint, and take other appropriate actions for new block
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
		disputable:     dispNode,
		prev:           prevNode,
		linkType:       ValidChildType,
		machineHash:    afterMachineHash,
		pendingTopHash: dispNode.pendingTopHash,
		machine:        afterMachine,
	}
	newNode.setHash()
	prevNode.successorHashes[ValidChildType] = newNode.hash
	chain.leaves.Add(newNode)

	// create nodes for invalid branches
	for kind := MinInvalidChildType; kind <= MaxChildType; kind++ {
		newNode := &Node{
			disputable:     dispNode,
			prev:           prevNode,
			linkType:       kind,
			machineHash:    prevNode.machineHash,
			machine:        prevNode.machine,
			pendingTopHash: prevNode.pendingTopHash,
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

func (node *Node) MarshalToBuf() *NodeBuf {
	if node.machine != nil {
		//TODO: marshal node.machine
	}
	return &NodeBuf{
		DisputableNode: node.disputable.MarshalToBuf(),
		MachineHash:    string(node.machineHash[:]),
		PendingTopHash: string(node.pendingTopHash[:]),
		LinkType:       uint32(node.linkType),
		PrevHash:       string(node.prev.hash[:]),
	}
}

func (buf *NodeBuf) Unmarshal(chain *Chain) (*Node, [32]byte) {
	var machineHashArr [32]byte
	copy(machineHashArr[:], []byte(buf.MachineHash))
	var prevHashArr [32]byte
	copy(prevHashArr[:], []byte(buf.PrevHash))
	var pthArr [32]byte
	copy(pthArr[:], []byte(buf.PendingTopHash))
	node := &Node{
		disputable:     buf.DisputableNode.Unmarshal(),
		machineHash:    machineHashArr,
		pendingTopHash: pthArr,
		linkType:       ChildType(buf.LinkType),
	}
	//TODO: try to retrieve machine from checkpoint DB; might fail
	node.setHash()
	chain.nodeFromHash[node.hash] = node

	// can't set up prev and successorHash fields yet; return prevHashArr so caller can do this later
	return node, prevHashArr
}

func (chain *Chain) CreateStake(stakerAddr common.Address, nodeHash [32]byte, creationTime RollupTime) {
	staker := &Staker{
		stakerAddr,
		chain.nodeFromHash[nodeHash],
		creationTime,
		nil,
	}
	chain.stakers.Add(staker)
}

func (chain *Chain) MoveStake(stakerAddr common.Address, nodeHash [32]byte) {
	chain.stakers.Get(stakerAddr).location = chain.nodeFromHash[nodeHash]
}

func (chain *Chain) RemoveStake(stakerAddr common.Address) {
	chain.stakers.Delete(chain.stakers.Get(stakerAddr))
}

func (staker *Staker) MarshalToBuf() *StakerBuf {
	challengeStr := ""
	if staker.challenge != nil {
		challengeStr = string(staker.challenge.contract.Bytes())
	}
	return &StakerBuf{
		Address:       string(staker.address.Bytes()),
		Location:      string(string(staker.location.hash[:])),
		CreationTime:  staker.creationTime.MarshalToBuf(),
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
			creationTime: buf.CreationTime.Unmarshal(),
			challenge:    nil,
		}
	} else {
		return &Staker{
			address:      common.BytesToAddress([]byte(buf.Address)),
			location:     chain.nodeFromHash[locArr],
			creationTime: buf.CreationTime.Unmarshal(),
			challenge:    chain.challenges[common.BytesToAddress([]byte(buf.ChallengeAddr))],
		}
	}
}

type Challenge struct {
	contract   common.Address
	asserter   common.Address
	challenger common.Address
}

type ChallengeType uint

const (
	InvalidPendingTopChallenge ChallengeType = 1
	InvalidMessagesChallenge   ChallengeType = 2
	InvalidExecutionChallenge  ChallengeType = 3
)

func (chain *Chain) NewChallenge(contract, asserter, challenger common.Address) *Challenge {
	ret := &Challenge{contract, asserter, challenger}
	chain.challenges[contract] = ret
	chain.stakers.Get(asserter).challenge = ret
	chain.stakers.Get(challenger).challenge = ret
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
	}
}

func (buf *ChallengeBuf) Unmarshal(chain *Chain) *Challenge {
	ret := &Challenge{
		common.BytesToAddress([]byte(buf.Contract)),
		common.BytesToAddress([]byte(buf.Asserter)),
		common.BytesToAddress([]byte(buf.Challenger)),
	}
	chain.challenges[ret.contract] = ret
	return ret
}
