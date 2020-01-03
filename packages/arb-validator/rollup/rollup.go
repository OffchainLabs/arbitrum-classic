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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. --go_out=paths=source_relative:. *.proto"

type ChainObserver struct {
	rollupAddr       common.Address
	vmParams         ChainParams
	pendingInbox     *PendingInbox
	latestConfirmed  *Node
	leaves           *LeafSet
	nodeFromHash     map[[32]byte]*Node
	oldestNode       *Node
	stakers          *StakerSet
	challenges       map[common.Address]*Challenge
	listenForAddress common.Address
	listener         ChainEventListener
}

func NewChain(
	_rollupAddr common.Address,
	_machine machine.Machine,
	_vmParams ChainParams,
	_listenForAddress common.Address,
	_listener ChainEventListener,
) *ChainObserver {
	ret := &ChainObserver{
		_rollupAddr,
		_vmParams,
		NewPendingInbox(),
		nil,
		NewLeafSet(),
		make(map[[32]byte]*Node),
		nil,
		NewStakerSet(),
		make(map[common.Address]*Challenge),
		_listenForAddress,
		_listener,
	}
	ret.CreateInitialNode(_machine)
	return ret
}

func (chain *ChainObserver) MarshalToBuf() *ChainObserverBuf {
	var allNodes []*NodeBuf
	for _, v := range chain.nodeFromHash {
		allNodes = append(allNodes, v.MarshalToBuf())
	}
	var allStakers []*StakerBuf
	chain.stakers.forall(func(staker *Staker) {
		allStakers = append(allStakers, staker.MarshalToBuf())
	})
	var leafHashes [][32]byte
	chain.leaves.forall(func(node *Node) {
		leafHashes = append(leafHashes, node.hash)
	})
	var allChallenges []*ChallengeBuf
	for _, v := range chain.challenges {
		allChallenges = append(allChallenges, v.MarshalToBuf())
	}
	return &ChainObserverBuf{
		ContractAddress:     chain.rollupAddr.Bytes(),
		VmParams:            chain.vmParams.MarshalToBuf(),
		PendingInbox:        chain.pendingInbox.MarshalToBuf(),
		Nodes:               allNodes,
		OldestNodeHash:      marshalHash(chain.oldestNode.hash),
		LatestConfirmedHash: marshalHash(chain.latestConfirmed.hash),
		LeafHashes:          marshalSliceOfHashes(leafHashes),
		Stakers:             allStakers,
		Challenges:          allChallenges,
	}
}

func (buf *ChainObserverBuf) Unmarshal(_listenForAddress common.Address, _listener ChainEventListener) *ChainObserver {
	chain := &ChainObserver{
		common.BytesToAddress([]byte(buf.ContractAddress)),
		buf.VmParams.Unmarshal(),
		buf.PendingInbox.Unmarshal(),
		nil,
		NewLeafSet(),
		make(map[[32]byte]*Node),
		nil,
		NewStakerSet(),
		make(map[common.Address]*Challenge),
		_listenForAddress,
		_listener,
	}
	for _, chalBuf := range buf.Challenges {
		chal := &Challenge{
			common.BytesToAddress([]byte(chalBuf.Contract)),
			common.BytesToAddress([]byte(chalBuf.Asserter)),
			common.BytesToAddress([]byte(chalBuf.Challenger)),
			ChallengeType(chalBuf.Kind),
		}
		chain.challenges[chal.contract] = chal
	}
	for _, nodeBuf := range buf.Nodes {
		nodeHash := unmarshalHash(nodeBuf.Hash)
		node := chain.nodeFromHash[nodeHash]
		prevHash := unmarshalHash(nodeBuf.PrevHash)
		if prevHash != zeroBytes32 {
			prev := chain.nodeFromHash[prevHash]
			node.prev = prev
			prev.successorHashes[node.linkType] = nodeHash
		}
	}
	chain.oldestNode = chain.nodeFromHash[unmarshalHash(buf.OldestNodeHash)]
	for _, leafHashStr := range buf.LeafHashes {
		leafHash := unmarshalHash(leafHashStr)
		chain.leaves.Add(chain.nodeFromHash[leafHash])
	}
	for _, stakerBuf := range buf.Stakers {
		locationHash := unmarshalHash(stakerBuf.Location)
		newStaker := &Staker{
			common.BytesToAddress(stakerBuf.Address),
			chain.nodeFromHash[locationHash],
			stakerBuf.CreationTime.Unmarshal(),
			chain.challenges[common.BytesToAddress(stakerBuf.ChallengeAddr)],
		}
		newStaker.location.numStakers++
		chain.stakers.Add(newStaker)
	}
	lcHash := unmarshalHash(buf.LatestConfirmedHash)
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
		StakeRequirement:  marshalBigInt(params.stakeRequirement),
		GracePeriod:       params.gracePeriod.MarshalToBuf(),
		MaxExecutionSteps: params.maxExecutionSteps,
	}
}

func (buf *ChainParamsBuf) Unmarshal() ChainParams {
	return ChainParams{
		unmarshalBigInt(buf.StakeRequirement),
		buf.GracePeriod.Unmarshal(),
		buf.MaxExecutionSteps,
	}
}

type DisputableNode struct {
	prevNodeHash          [32]byte
	timeBounds            [2]RollupTime
	afterPendingTop       [32]byte
	importedMessagesSlice [32]byte
	importedMessageCount  *big.Int
	assertionStub         *protocol.AssertionStub
	hash                  [32]byte
}

func (dn *DisputableNode) MarshalToBuf() *DisputableNodeBuf {
	return &DisputableNodeBuf{
		PrevNodeHash:          marshalHash(dn.prevNodeHash),
		TimeLowerBound:        dn.timeBounds[0].MarshalToBuf(),
		TimeUpperBound:        dn.timeBounds[1].MarshalToBuf(),
		AfterPendingTop:       marshalHash(dn.afterPendingTop),
		ImportedMessagesSlice: marshalHash(dn.importedMessagesSlice),
		ImportedMessageCount:  marshalBigInt(dn.importedMessageCount),
		AssertionStub:         dn.assertionStub,
	}
}

func (buf *DisputableNodeBuf) Unmarshal() *DisputableNode {
	ret := &DisputableNode{
		prevNodeHash:          unmarshalHash(buf.PrevNodeHash),
		timeBounds:            [2]RollupTime{buf.TimeLowerBound.Unmarshal(), buf.TimeUpperBound.Unmarshal()},
		afterPendingTop:       unmarshalHash(buf.AfterPendingTop),
		importedMessagesSlice: unmarshalHash(buf.ImportedMessagesSlice),
		importedMessageCount:  unmarshalBigInt(buf.ImportedMessageCount),
		assertionStub:         buf.AssertionStub,
	}
	ret.hash = ret._hash()
	return ret
}

func (dn *DisputableNode) _hash() [32]byte {
	var ret [32]byte
	retSlice := solsha3.SoliditySHA3(
		solsha3.Bytes32(unmarshalHash(dn.assertionStub.AfterHash)),
		solsha3.Bool(dn.assertionStub.DidInboxInsn),
		solsha3.Uint32(dn.assertionStub.NumSteps),
		solsha3.Uint64(dn.assertionStub.NumGas),
		solsha3.Bytes32(unmarshalHash(dn.assertionStub.FirstMessageHash)),
		solsha3.Bytes32(unmarshalHash(dn.assertionStub.LastMessageHash)),
		solsha3.Bytes32(unmarshalHash(dn.assertionStub.FirstLogHash)),
		solsha3.Bytes32(unmarshalHash(dn.assertionStub.LastLogHash)),
	)
	copy(ret[:], retSlice)
	return ret
}

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

func (chain *ChainObserver) CreateInitialNode(machine machine.Machine) {
	newNode := &Node{
		depth:          0,
		machineHash:    machine.Hash(),
		machine:        machine.Clone(),
		pendingTopHash: value.NewEmptyTuple().Hash(),
		linkType:       ValidChildType,
		numStakers:     0,
	}
	newNode.setHash()
	chain.leaves.Add(newNode)
	chain.latestConfirmed = newNode
}

func (chain *ChainObserver) GeneratePathProof(from, to *Node) [][32]byte {
	// returns nil if no proof exists
	if to == nil {
		return nil
	}
	if from == to {
		return [][32]byte{}
	} else {
		sub := chain.GeneratePathProof(from, to.prev)
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

func (chain *ChainObserver) GenerateConflictProof(from, to1, to2 *Node) ([][32]byte, [][32]byte) {
	// returns nil, nil if no proof exists
	proof1 := chain.GeneratePathProof(from, to1)
	proof2 := chain.GeneratePathProof(from, to2)
	if proof1 == nil || proof2 == nil || len(proof1) == 0 || len(proof2) == 0 || proof1[0] == proof2[0] {
		return nil, nil
	} else {
		return proof1, proof2
	}
}

func (chain *ChainObserver) CommonAncestor(n1, n2 *Node) *Node {
	for n1.depth > n2.depth {
		n1 = n1.prev
	}
	for n2.depth > n1.depth {
		n2 = n2.prev
	}
	for n1 != n2 {
		n1 = n1.prev
		n2 = n2.prev
	}
	return n1
}

func (chain *ChainObserver) CommonAncestorIfConflict(n1, n2 *Node) *Node { // return common ancestor, or nil if not conflicting
	ret := chain.CommonAncestor(n1, n2)
	if ret == n1 || ret == n2 {
		return nil
	} else {
		return ret
	}
}

func (chain *ChainObserver) notifyNewBlockNumber(blockNum *big.Int) {
	//TODO: checkpoint, and take other appropriate actions for new block
}

func (chain *ChainObserver) notifyAssert(
	prevLeafHash [32]byte,
	timeBounds [2]RollupTime,
	afterPendingTop [32]byte,
	importedMessagesSlice [32]byte,
	importedMessageCount *big.Int,
	assertionStub *protocol.AssertionStub,
) {
	disputableNode := &DisputableNode{
		prevNodeHash:          prevLeafHash,
		timeBounds:            timeBounds,
		afterPendingTop:       afterPendingTop,
		importedMessagesSlice: importedMessagesSlice,
		importedMessageCount:  importedMessageCount,
		assertionStub:         assertionStub,
	}
	disputableNode.hash = disputableNode._hash()
	chain.CreateNodesOnAssert(chain.nodeFromHash[prevLeafHash], disputableNode, unmarshalHash(disputableNode.assertionStub.AfterHash), nil)
}

func (chain *ChainObserver) CreateNodesOnAssert(
	prevNode *Node,
	dispNode *DisputableNode,
	afterMachineHash [32]byte,
	afterMachine machine.Machine,
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
		depth:          1 + prevNode.depth,
		disputable:     dispNode,
		prev:           prevNode,
		linkType:       ValidChildType,
		machineHash:    afterMachineHash,
		pendingTopHash: dispNode.afterPendingTop,
		machine:        afterMachine,
		numStakers:     0,
	}
	newNode.setHash()
	prevNode.successorHashes[ValidChildType] = newNode.hash
	chain.leaves.Add(newNode)

	// create nodes for invalid branches
	for kind := MinInvalidChildType; kind <= MaxChildType; kind++ {
		newNode := &Node{
			depth:          1 + prevNode.depth,
			disputable:     dispNode,
			prev:           prevNode,
			linkType:       kind,
			machineHash:    prevNode.machineHash,
			machine:        prevNode.machine,
			pendingTopHash: prevNode.pendingTopHash,
			numStakers:     0,
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

func (chain *ChainObserver) pruneNode(node *Node) {
	oldNode := node.prev
	node.prev = nil // so garbage collector doesn't preserve prev anymore
	if oldNode != nil {
		oldNode.successorHashes[node.linkType] = zeroBytes32
		chain.considerPruningNode(oldNode)
	}
	delete(chain.nodeFromHash, node.hash)
}

func (chain *ChainObserver) considerPruningNode(node *Node) {
	if node.numStakers > 0 {
		return
	}
	for kind := MinChildType; kind <= MaxChildType; kind++ {
		if node.successorHashes[kind] != zeroBytes32 {
			return
		}
	}
	chain.pruneNode(node)
}

func (chain *ChainObserver) ConfirmNode(nodeHash [32]byte) {
	node := chain.nodeFromHash[nodeHash]
	chain.latestConfirmed = node
	chain.considerPruningNode(node.prev)
	for chain.oldestNode != chain.latestConfirmed {
		if chain.oldestNode.numStakers > 0 {
			break
		}
		numSuccessors := 0
		for kind := MinChildType; kind <= MaxChildType; kind++ {
			if node.successorHashes[kind] != zeroBytes32 {
				numSuccessors++
			}
		}
		if numSuccessors > 1 {
			break
		}
		newOldestNode := chain.nodeFromHash[chain.oldestNode]
		chain.pruneNode(chain.oldestNode)
		chain.oldestNode = newOldestNode
	}
}

func (chain *ChainObserver) PruneNodeByHash(nodeHash [32]byte) {
	chain.pruneNode(chain.nodeFromHash[nodeHash])
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

func (chain *ChainObserver) CreateStake(stakerAddr common.Address, nodeHash [32]byte, creationTime RollupTime) {
	staker := &Staker{
		stakerAddr,
		chain.nodeFromHash[nodeHash],
		creationTime,
		nil,
	}
	staker.location.numStakers++
	chain.stakers.Add(staker)
}

func (chain *ChainObserver) MoveStake(stakerAddr common.Address, nodeHash [32]byte) {
	staker := chain.stakers.Get(stakerAddr)
	staker.location.numStakers--
	// no need to consider pruning staker.location, because a successor of it is getting a stake
	staker.location = chain.nodeFromHash[nodeHash]
	staker.location.numStakers++
}

func (chain *ChainObserver) RemoveStake(stakerAddr common.Address) {
	staker := chain.stakers.Get(stakerAddr)
	staker.location.numStakers--
	chain.considerPruningNode(staker.location)
	chain.stakers.Delete(staker)
}

func (staker *Staker) MarshalToBuf() *StakerBuf {
	if staker.challenge == nil {
		return &StakerBuf{
			Address:      staker.address.Bytes(),
			Location:     marshalHash(staker.location.hash),
			CreationTime: staker.creationTime.MarshalToBuf(),
			InChallenge:  false,
		}
	} else {
		return &StakerBuf{
			Address:       staker.address.Bytes(),
			Location:      marshalHash(staker.location.hash),
			CreationTime:  staker.creationTime.MarshalToBuf(),
			InChallenge:   true,
			ChallengeAddr: staker.challenge.contract.Bytes(),
		}
	}
}

func (buf *StakerBuf) Unmarshal(chain *ChainObserver) *Staker {
	// chain.nodeFromHash and chain.challenges must have already been unmarshaled
	locArr := unmarshalHash(buf.Location)
	if buf.InChallenge {
		return &Staker{
			address:      common.BytesToAddress([]byte(buf.Address)),
			location:     chain.nodeFromHash[locArr],
			creationTime: buf.CreationTime.Unmarshal(),
			challenge:    chain.challenges[common.BytesToAddress(buf.ChallengeAddr)],
		}
	} else {
		return &Staker{
			address:      common.BytesToAddress([]byte(buf.Address)),
			location:     chain.nodeFromHash[locArr],
			creationTime: buf.CreationTime.Unmarshal(),
			challenge:    nil,
		}
	}
}

type Challenge struct {
	contract   common.Address
	asserter   common.Address
	challenger common.Address
	kind       ChallengeType
}

type ChallengeType uint32

const (
	InvalidPendingTopChallenge ChallengeType = 0
	InvalidMessagesChallenge   ChallengeType = 1
	InvalidExecutionChallenge  ChallengeType = 2
)

func (chain *ChainObserver) NewChallenge(contract, asserter, challenger common.Address, kind ChallengeType) *Challenge {
	ret := &Challenge{contract, asserter, challenger, kind}
	chain.challenges[contract] = ret
	chain.stakers.Get(asserter).challenge = ret
	chain.stakers.Get(challenger).challenge = ret
	return ret
}

func (chain *ChainObserver) ChallengeResolved(contract, winner, loser common.Address) {
	chain.RemoveStake(loser)
	delete(chain.challenges, contract)
}

func (chal *Challenge) MarshalToBuf() *ChallengeBuf {
	return &ChallengeBuf{
		Contract:   chal.contract.Bytes(),
		Asserter:   chal.asserter.Bytes(),
		Challenger: chal.challenger.Bytes(),
	}
}

func (buf *ChallengeBuf) Unmarshal(chain *ChainObserver) *Challenge {
	ret := &Challenge{
		common.BytesToAddress(buf.Contract),
		common.BytesToAddress(buf.Asserter),
		common.BytesToAddress(buf.Challenger),
		ChallengeType(buf.Kind),
	}
	chain.challenges[ret.contract] = ret
	return ret
}
