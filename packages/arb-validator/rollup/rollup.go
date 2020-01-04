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
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

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
		&PendingInbox{buf.PendingInbox.Unmarshal()},
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
	n1, _, _ = chain.GetConflictAncestor(n1, n2)
	return n1
}

func (chain *ChainObserver) GetConflictAncestor(n1, n2 *Node) (*Node, ChildType, error) {
	n1Orig := n1
	n2Orig := n2
	prevN1 := n1
	prevN2 := n1
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
		return n1, 0, errors.New("no conflict")
	}
	linkType := prevN1.linkType
	if prevN2.linkType < linkType {
		linkType = prevN2.linkType
	}

	return n1, linkType, nil
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
			return
		}
		var successor *Node
		for kind := MinChildType; kind <= MaxChildType; kind++ {
			if node.successorHashes[kind] != zeroBytes32 {
				if successor != nil {
					return
				}
				successor = chain.nodeFromHash[node.successorHashes[kind]]
			}
		}
		chain.pruneNode(chain.oldestNode)
		chain.oldestNode = successor
	}
}

func (chain *ChainObserver) PruneNodeByHash(nodeHash [32]byte) {
	chain.pruneNode(chain.nodeFromHash[nodeHash])
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
