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
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"sort"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. --go_out=paths=source_relative:. *.proto"

type ChallengeType uint32

const (
	InvalidPendingTopChallenge ChallengeType = 0
	InvalidMessagesChallenge   ChallengeType = 1
	InvalidExecutionChallenge  ChallengeType = 2
)

var zeroBytes32 [32]byte // deliberately zeroed

type Challenge struct {
	contract   common.Address
	asserter   common.Address
	challenger common.Address
	kind       ChallengeType
}

func (chal *Challenge) MarshalToBuf() *ChallengeBuf {
	return &ChallengeBuf{
		Contract:   chal.contract.Bytes(),
		Asserter:   chal.asserter.Bytes(),
		Challenger: chal.challenger.Bytes(),
	}
}

func (m *ChallengeBuf) Unmarshal() *Challenge {
	return &Challenge{
		contract:   common.BytesToAddress(m.Contract),
		asserter:   common.BytesToAddress(m.Asserter),
		challenger: common.BytesToAddress(m.Challenger),
		kind:       ChallengeType(m.Kind),
	}
}

type StakedNodeGraph struct {
	*NodeGraph
	stakers    *StakerSet
	challenges map[common.Address]*Challenge
}

func NewStakedNodeGraph(machine machine.Machine) *StakedNodeGraph {
	ret := &StakedNodeGraph{
		NodeGraph:  NewNodeGraph(machine),
		stakers:    NewStakerSet(),
		challenges: make(map[common.Address]*Challenge),
	}
	ret.startCleanupThread(nil)
	return ret
}

func (chain *StakedNodeGraph) MarshalToBuf() *StakedNodeGraphBuf {
	var allStakers []*StakerBuf
	chain.stakers.forall(func(staker *Staker) {
		allStakers = append(allStakers, staker.MarshalToBuf())
	})
	var allChallenges []*ChallengeBuf
	for _, v := range chain.challenges {
		allChallenges = append(allChallenges, v.MarshalToBuf())
	}
	return &StakedNodeGraphBuf{
		NodeGraph:  chain.NodeGraph.MarshalToBuf(),
		Stakers:    allStakers,
		Challenges: allChallenges,
	}
}

func (m *StakedNodeGraphBuf) Unmarshal() *StakedNodeGraph {
	chain := &StakedNodeGraph{
		NodeGraph:  m.NodeGraph.Unmarshal(),
		stakers:    NewStakerSet(),
		challenges: make(map[common.Address]*Challenge),
	}
	for _, chalBuf := range m.Challenges {
		chal := chalBuf.Unmarshal()
		chain.challenges[chal.contract] = chal
	}
	for _, stakerBuf := range m.Stakers {
		chain.stakers.Add(stakerBuf.Unmarshal(chain))
	}
	chain.startCleanupThread(nil)

	return chain
}

func (chain *StakedNodeGraph) CreateStake(stakerAddr common.Address, nodeHash [32]byte, creationTime RollupTime) {
	chain.stakers.Add(&Staker{
		stakerAddr,
		chain.nodeFromHash[nodeHash],
		creationTime,
		nil,
	})
}

func (chain *StakedNodeGraph) MoveStake(stakerAddr common.Address, nodeHash [32]byte) {
	staker := chain.stakers.Get(stakerAddr)
	staker.location.numStakers--
	// no need to consider pruning staker.location, because a successor of it is getting a stake
	staker.location = chain.nodeFromHash[nodeHash]
	staker.location.numStakers++
}

func (chain *StakedNodeGraph) RemoveStake(stakerAddr common.Address) {
	staker := chain.stakers.Get(stakerAddr)
	staker.location.numStakers--
	chain.considerPruningNode(staker.location)
	chain.stakers.Delete(staker)
}

func (chain *StakedNodeGraph) NewChallenge(contract, asserter, challenger common.Address, kind ChallengeType) *Challenge {
	ret := &Challenge{contract, asserter, challenger, kind}
	chain.challenges[contract] = ret
	chain.stakers.Get(asserter).challenge = ret
	chain.stakers.Get(challenger).challenge = ret
	return ret
}

func (chain *StakedNodeGraph) ChallengeResolved(contract, winner, loser common.Address) {
	chain.RemoveStake(loser)
	delete(chain.challenges, contract)
}

type SortableAddressList []common.Address

func (sa SortableAddressList) Len() int {
	return len(sa)
}

func (sa SortableAddressList) Less(i, j int) bool {
	return bytes.Compare(sa[i][:], sa[j][:]) < 0
}

func (sa SortableAddressList) Swap(i, j int) {
	sa[i], sa[j] = sa[j], sa[i]
}

func (sng *StakedNodeGraph) generateAlignedStakersProof(
	confirmingNode *Node,
	deadline RollupTime,
) (stakerAddrs []common.Address, proof [][32]byte, offsets []uint64) {
	sng.stakers.forall(func(st *Staker) {
		stakerAddrs = append(stakerAddrs, st.address)
	})
	sort.Sort(SortableAddressList(stakerAddrs))

	for i, sa := range stakerAddrs {
		staker := sng.stakers.Get(sa)
		if staker.creationTime.Cmp(deadline) >= 0 {
			offsets[i] = uint64(len(proof))
			subProof := GeneratePathProof(confirmingNode, staker.location)
			proof = append(proof, subProof...)
		}
	}
	return
}
