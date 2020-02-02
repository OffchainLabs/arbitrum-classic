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
	"log"
	"math/big"
	"sort"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. --go_out=paths=source_relative:. *.proto"

var zeroBytes32 common.Hash // deliberately zeroed

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

func (sng *StakedNodeGraph) MarshalForCheckpoint(ctx structures.CheckpointContext) *StakedNodeGraphBuf {
	var allStakers []*StakerBuf
	sng.stakers.forall(func(staker *Staker) {
		allStakers = append(allStakers, staker.MarshalToBuf())
	})
	return &StakedNodeGraphBuf{
		NodeGraph: sng.NodeGraph.MarshalForCheckpoint(ctx),
		Stakers:   allStakers,
	}
}

func (m *StakedNodeGraphBuf) UnmarshalFromCheckpoint(ctx structures.RestoreContext) *StakedNodeGraph {
	chain := &StakedNodeGraph{
		NodeGraph: m.NodeGraph.UnmarshalFromCheckpoint(ctx),
		stakers:   NewStakerSet(),
	}
	for _, stakerBuf := range m.Stakers {
		chain.stakers.Add(stakerBuf.Unmarshal(chain))
	}
	return chain
}

func (sng *StakedNodeGraph) DebugString(prefix string) string {
	subPrefix := prefix + "  "
	return "\n" + prefix + "nodes:\n" + sng.NodeGraph.DebugString(sng.stakers, subPrefix) + sng.stakers.DebugString(prefix)
}

func (sng *StakedNodeGraph) Equals(s2 *StakedNodeGraph) bool {
	return sng.NodeGraph.Equals(s2.NodeGraph) &&
		sng.stakers.Equals(s2.stakers)
}

func (sng *StakedNodeGraph) CreateStake(ev arbbridge.StakeCreatedEvent) {
	node, ok := sng.nodeFromHash[ev.NodeHash]
	if !ok {
		log.Println("Bad location", ev.NodeHash)
		panic("Tried to create stake on bad node")
	}
	sng.stakers.Add(&Staker{
		ev.Staker,
		node,
		common.TimeFromBlockNum(ev.BlockID.Height),
		common.Address{},
	})
}

func (sng *StakedNodeGraph) MoveStake(stakerAddr common.Address, nodeHash common.Hash) {
	staker := sng.stakers.Get(stakerAddr)
	if staker == nil {
		log.Fatalf("Moved nonexistant staker %v to node %v", stakerAddr, nodeHash)
	}
	staker.location.numStakers--
	// no need to consider pruning staker.location, because a successor of it is getting a stake
	newLocation, ok := sng.nodeFromHash[nodeHash]
	if !ok {
		log.Fatalf("Moved staker %v to nonexistant node %v", stakerAddr, nodeHash)
	}
	staker.location = newLocation
	staker.location.numStakers++
}

func (sng *StakedNodeGraph) RemoveStake(stakerAddr common.Address) {
	staker := sng.stakers.Get(stakerAddr)
	staker.location.numStakers--
	sng.considerPruningNode(staker.location)
	sng.stakers.Delete(staker)
}

func (sng *StakedNodeGraph) NewChallenge(contract, asserter, challenger common.Address, kind structures.ChildType) {
	sng.stakers.Get(asserter).challenge = contract
	sng.stakers.Get(challenger).challenge = contract
}

func (sng *StakedNodeGraph) ChallengeResolved(contract, winner, loser common.Address) {
	sng.stakers.Get(winner).challenge = common.Address{}
	sng.RemoveStake(loser)
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

func (sng *StakedNodeGraph) generateNodePruneInfo() []pruneParams {
	prunesToDo := make([]pruneParams, 0)
	sng.leaves.forall(func(leaf *Node) {
		if leaf != sng.latestConfirmed {
			leafAncestor, _, _, err := GetConflictAncestor(leaf, sng.latestConfirmed)
			if err == nil {
				noStakersOnLeaf := true
				sng.stakers.forall(func(s *Staker) {
					if s.location.Equals(leaf) {
						noStakersOnLeaf = false
					}
				})
				if noStakersOnLeaf {
					prunesToDo = append(prunesToDo, pruneParams{
						leafHash:     leaf.hash,
						ancestorHash: leafAncestor.prev.hash,
						leafProof:    GeneratePathProof(leafAncestor.prev, leaf),
						ancProof:     GeneratePathProof(leafAncestor.prev, sng.latestConfirmed),
					})
				}
			}
		}
	})
	return prunesToDo
}

type confirmValidOpportunity struct {
	nodeHash           common.Hash
	deadlineTicks      common.TimeTicks
	messages           []value.Value
	logsAcc            common.Hash
	vmProtoStateHash   common.Hash
	stakerAddresses    []common.Address
	stakerProofs       []common.Hash
	stakerProofOffsets []*big.Int
}

type confirmInvalidOpportunity struct {
	nodeHash           common.Hash
	deadlineTicks      common.TimeTicks
	challengeNodeData  common.Hash
	branch             structures.ChildType
	vmProtoStateHash   common.Hash
	stakerAddresses    []common.Address
	stakerProofs       []common.Hash
	stakerProofOffsets []*big.Int
}

func (sng *StakedNodeGraph) generateNextConfProof(
	currentTime common.TimeTicks,
) (*confirmValidOpportunity, *confirmInvalidOpportunity) {
	stakerAddrs := make([]common.Address, 0)
	sng.stakers.forall(func(st *Staker) {
		stakerAddrs = append(stakerAddrs, st.address)
	})
	sort.Sort(SortableAddressList(stakerAddrs))

	for _, successor := range sng.latestConfirmed.successorHashes {
		if successor == zeroBytes32 {
			continue
		}
		node := sng.nodeFromHash[successor]
		proof, offsets := sng.generateAlignedStakersProof(
			node,
			currentTime,
			stakerAddrs,
		)

		if proof != nil {
			if node.linkType == structures.ValidChildType {
				if node.assertion == nil {
					return nil, nil
				}
				return &confirmValidOpportunity{
					nodeHash:           node.hash,
					deadlineTicks:      common.TimeTicks{Val: new(big.Int).Set(node.deadline.Val)},
					messages:           node.assertion.OutMsgs,
					logsAcc:            node.disputable.AssertionClaim.AssertionStub.LastLogHash,
					vmProtoStateHash:   node.vmProtoData.Hash(),
					stakerAddresses:    stakerAddrs,
					stakerProofs:       proof,
					stakerProofOffsets: offsets,
				}, nil
			} else {
				return nil, &confirmInvalidOpportunity{
					nodeHash:           node.hash,
					deadlineTicks:      common.TimeTicks{Val: new(big.Int).Set(node.deadline.Val)},
					challengeNodeData:  node.nodeDataHash,
					branch:             node.linkType,
					vmProtoStateHash:   node.vmProtoData.Hash(),
					stakerAddresses:    stakerAddrs,
					stakerProofs:       proof,
					stakerProofOffsets: offsets,
				}
			}
		}
	}

	return nil, nil
}

func (sng *StakedNodeGraph) generateAlignedStakersProof(
	confirmingNode *Node,
	currentTime common.TimeTicks,
	stakerAddrs []common.Address,
) ([]common.Hash, []*big.Int) {
	proof := make([]common.Hash, 0)
	offsets := make([]*big.Int, 0)
	deadline := confirmingNode.deadline
	if currentTime.Cmp(deadline) < 0 {
		return nil, nil
	}
	offsets = append(offsets, big.NewInt(0))
	for _, sa := range stakerAddrs {
		staker := sng.stakers.Get(sa)
		if staker.creationTime.Cmp(deadline) >= 0 {
			offsets = append(offsets, new(big.Int).SetUint64(uint64(len(proof))))
			continue
		}
		subProof := GeneratePathProof(confirmingNode, staker.location)
		if subProof == nil {
			return nil, nil
		}
		proof = append(proof, subProof...)
		offsets = append(offsets, new(big.Int).SetUint64(uint64(len(proof))))
	}
	if len(offsets) == 1 {
		return nil, nil
	}
	return proof, offsets
}

func (sng *StakedNodeGraph) generateStakerPruneInfo() ([]recoverStakeMootedParams, []recoverStakeOldParams) {
	mootedToDo := make([]recoverStakeMootedParams, 0)
	oldToDo := make([]recoverStakeOldParams, 0)
	sng.stakers.forall(func(staker *Staker) {
		stakerAncestor, _, _, err := GetConflictAncestor(staker.location, sng.latestConfirmed)
		if err == nil {
			mootedToDo = append(mootedToDo, recoverStakeMootedParams{
				addr:         staker.address,
				ancestorHash: stakerAncestor.prev.hash,
				lcProof:      GeneratePathProof(stakerAncestor.prev, sng.latestConfirmed),
				stProof:      GeneratePathProof(stakerAncestor.prev, staker.location),
			})
		} else if staker.location.depth < sng.latestConfirmed.depth {
			oldToDo = append(oldToDo, recoverStakeOldParams{
				addr:  staker.address,
				proof: GeneratePathProof(staker.location, sng.latestConfirmed),
			})
		}
	})
	return mootedToDo, oldToDo
}

type challengeOpportunity struct {
	asserter              common.Address
	challenger            common.Address
	prevNodeHash          common.Hash
	deadlineTicks         common.TimeTicks
	asserterNodeType      structures.ChildType
	challengerNodeType    structures.ChildType
	asserterVMProtoHash   common.Hash
	challengerVMProtoHash common.Hash
	asserterProof         []common.Hash
	challengerProof       []common.Hash
	asserterNodeHash      common.Hash
	challengerDataHash    common.Hash
	challengerPeriodTicks common.TimeTicks
}

func (sng *StakedNodeGraph) checkChallengeOpportunityPair(staker1, staker2 *Staker) *challengeOpportunity {
	if !staker1.challenge.IsZero() || !staker2.challenge.IsZero() {
		return nil
	}
	staker1Ancestor, staker2Ancestor, _, err := GetConflictAncestor(staker1.location, staker2.location)
	if err != nil {
		return nil
	}
	linkType1 := staker1Ancestor.linkType
	linkType2 := staker2Ancestor.linkType

	var asserterStaker *Staker
	var asserterAncestor *Node
	var challengerStaker *Staker
	var challengerAncestor *Node
	if linkType1 < linkType2 {
		asserterStaker = staker2
		asserterAncestor = staker2Ancestor
		challengerStaker = staker1
		challengerAncestor = staker1Ancestor
	} else {
		asserterStaker = staker1
		asserterAncestor = staker1Ancestor
		challengerStaker = staker2
		challengerAncestor = staker2Ancestor
	}

	challengerDataHash, challengerPeriodTicks := challengerAncestor.ChallengeNodeData(sng.params)

	return &challengeOpportunity{
		asserter:              asserterStaker.address,
		challenger:            challengerStaker.address,
		prevNodeHash:          asserterAncestor.prev.hash,
		deadlineTicks:         asserterAncestor.deadline,
		asserterNodeType:      asserterAncestor.linkType,
		challengerNodeType:    challengerAncestor.linkType,
		asserterVMProtoHash:   asserterAncestor.vmProtoData.Hash(),
		challengerVMProtoHash: challengerAncestor.vmProtoData.Hash(),
		asserterProof:         GeneratePathProof(asserterAncestor, asserterStaker.location),
		challengerProof:       GeneratePathProof(challengerAncestor, challengerStaker.location),
		asserterNodeHash:      asserterAncestor.nodeDataHash,
		challengerDataHash:    challengerDataHash,
		challengerPeriodTicks: challengerPeriodTicks,
	}
}

func (sng *StakedNodeGraph) checkChallengeOpportunityAny(staker *Staker) *challengeOpportunity {
	if !staker.challenge.IsZero() {
		return nil
	}
	var ret *challengeOpportunity
	sng.stakers.forall(func(staker2 *Staker) {
		if !staker2.Equals(staker) {
			opp := sng.checkChallengeOpportunityPair(staker, staker2)
			if opp != nil {
				ret = opp
				return
			}
		}
	})
	return ret
}
