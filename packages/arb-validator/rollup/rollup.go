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

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. --go_out=paths=source_relative:. *.proto"

type ChainObserver struct {
	*NodeGraph
	rollupAddr       common.Address
	vmParams         ChainParams
	pendingInbox     *PendingInbox
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
		NodeGraph:        NewNodeGraph(_machine),
		rollupAddr:       _rollupAddr,
		vmParams:         _vmParams,
		pendingInbox:     NewPendingInbox(),
		stakers:          NewStakerSet(),
		challenges:       make(map[common.Address]*Challenge),
		listenForAddress: _listenForAddress,
		listener:         _listener,
	}
	return ret
}

func (chain *ChainObserver) MarshalToBuf() *ChainObserverBuf {
	var allStakers []*StakerBuf
	chain.stakers.forall(func(staker *Staker) {
		allStakers = append(allStakers, staker.MarshalToBuf())
	})
	var allChallenges []*ChallengeBuf
	for _, v := range chain.challenges {
		allChallenges = append(allChallenges, v.MarshalToBuf())
	}
	return &ChainObserverBuf{
		ContractAddress: chain.rollupAddr.Bytes(),
		VmParams:        chain.vmParams.MarshalToBuf(),
		PendingInbox:    chain.pendingInbox.MarshalToBuf(),
		NodeGraph:       chain.NodeGraph.MarshalToBuf(),
		Stakers:         allStakers,
		Challenges:      allChallenges,
	}
}

func (m *ChainObserverBuf) Unmarshal(_listenForAddress common.Address, _listener ChainEventListener) *ChainObserver {
	chain := &ChainObserver{
		NodeGraph:        m.NodeGraph.Unmarshal(),
		rollupAddr:       common.BytesToAddress(m.ContractAddress),
		vmParams:         m.VmParams.Unmarshal(),
		pendingInbox:     &PendingInbox{m.PendingInbox.Unmarshal()},
		stakers:          NewStakerSet(),
		challenges:       make(map[common.Address]*Challenge),
		listenForAddress: _listenForAddress,
		listener:         _listener,
	}
	for _, chalBuf := range m.Challenges {
		chal := &Challenge{
			contract:   common.BytesToAddress(chalBuf.Contract),
			asserter:   common.BytesToAddress(chalBuf.Asserter),
			challenger: common.BytesToAddress(chalBuf.Challenger),
			kind:       ChallengeType(chalBuf.Kind),
		}
		chain.challenges[chal.contract] = chal
	}
	for _, stakerBuf := range m.Stakers {
		locationHash := unmarshalHash(stakerBuf.Location)
		newStaker := &Staker{
			address:      common.BytesToAddress(stakerBuf.Address),
			location:     chain.nodeFromHash[locationHash],
			creationTime: stakerBuf.CreationTime.Unmarshal(),
			challenge:    chain.challenges[common.BytesToAddress(stakerBuf.ChallengeAddr)],
		}
		newStaker.location.numStakers++
		chain.stakers.Add(newStaker)
	}

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

func (m *ChainParamsBuf) Unmarshal() ChainParams {
	return ChainParams{
		stakeRequirement:  unmarshalBigInt(m.StakeRequirement),
		gracePeriod:       m.GracePeriod.Unmarshal(),
		maxExecutionSteps: m.MaxExecutionSteps,
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

func (m *ChallengeBuf) Unmarshal(chain *ChainObserver) *Challenge {
	ret := &Challenge{
		contract:   common.BytesToAddress(m.Contract),
		asserter:   common.BytesToAddress(m.Asserter),
		challenger: common.BytesToAddress(m.Challenger),
		kind:       ChallengeType(m.Kind),
	}
	chain.challenges[ret.contract] = ret
	return ret
}
