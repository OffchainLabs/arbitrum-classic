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
	"sort"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. --go_out=paths=source_relative:. *.proto"

var zeroBytes32 [32]byte // deliberately zeroed

type StakedNodeGraph struct {
	*NodeGraph
	stakers *StakerSet
}

func NewStakedNodeGraph(machine machine.Machine, params structures.ChainParams) *StakedNodeGraph {
	return &StakedNodeGraph{
		NodeGraph: NewNodeGraph(machine, params),
		stakers:   NewStakerSet(),
	}
}

func (chain *StakedNodeGraph) MarshalToBuf() *StakedNodeGraphBuf {
	var allStakers []*StakerBuf
	chain.stakers.forall(func(staker *Staker) {
		allStakers = append(allStakers, staker.MarshalToBuf())
	})
	return &StakedNodeGraphBuf{
		NodeGraph: chain.NodeGraph.MarshalToBuf(),
		Stakers:   allStakers,
	}
}

func (m *StakedNodeGraphBuf) Unmarshal() *StakedNodeGraph {
	chain := &StakedNodeGraph{
		NodeGraph: m.NodeGraph.Unmarshal(),
		stakers:   NewStakerSet(),
	}
	for _, stakerBuf := range m.Stakers {
		chain.stakers.Add(stakerBuf.Unmarshal(chain))
	}
	return chain
}

func (s *StakedNodeGraph) Equals(s2 *StakedNodeGraph) bool {
	return s.NodeGraph.Equals(s2.NodeGraph) &&
		s.stakers.Equals(s2.stakers)
}

func (chain *StakedNodeGraph) CreateStake(ev ethbridge.StakeCreatedEvent, currentTime structures.TimeTicks) {
	chain.stakers.Add(&Staker{
		ev.Staker,
		chain.nodeFromHash[ev.NodeHash],
		currentTime,
		common.Address{},
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

func (chain *StakedNodeGraph) NewChallenge(contract, asserter, challenger common.Address, kind structures.ChildType) {
	chain.stakers.Get(asserter).challenge = contract
	chain.stakers.Get(challenger).challenge = contract
}

func (chain *StakedNodeGraph) ChallengeResolved(contract, winner, loser common.Address) {
	chain.stakers.Get(winner).challenge = common.Address{}
	chain.RemoveStake(loser)
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
	deadline structures.TimeTicks,
) (stakerAddrs []common.Address, proof [][32]byte, offsets []uint64) {
	sng.stakers.forall(func(st *Staker) {
		stakerAddrs = append(stakerAddrs, st.address)
	})
	sort.Sort(SortableAddressList(stakerAddrs))

	for _, sa := range stakerAddrs {
		staker := sng.stakers.Get(sa)
		if staker.creationTime.Cmp(deadline) >= 0 {
			offsets = append(offsets, uint64(len(proof)))
			subProof := GeneratePathProof(confirmingNode, staker.location)
			proof = append(proof, subProof...)
		}
	}
	return
}

func (chain *StakedNodeGraph) generateStakerPruneInfo() ([]recoverStakeMootedParams, []recoverStakeOldParams) {
	mootedToDo := []recoverStakeMootedParams{}
	oldToDo := []recoverStakeOldParams{}
	chain.stakers.forall(func(staker *Staker) {
		ancestor, _, err := chain.GetConflictAncestor(staker.location, chain.latestConfirmed)
		if err == nil {
			mootedToDo = append(mootedToDo, recoverStakeMootedParams{
				addr:     staker.address,
				ancestor: ancestor,
				lcProof:  GeneratePathProof(ancestor, chain.latestConfirmed),
				stProof:  GeneratePathProof(ancestor, staker.location),
			})
		} else if staker.location.depth < chain.latestConfirmed.depth {
			oldToDo = append(oldToDo, recoverStakeOldParams{
				addr:  staker.address,
				proof: GeneratePathProof(staker.location, chain.latestConfirmed),
			})
		}
	})
	return mootedToDo, oldToDo
}
