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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"sort"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

//go:generate bash -c "protoc -I$(go list -f '{{ .Dir }}' -m github.com/offchainlabs/arbitrum/packages/arb-util) -I. --go_out=paths=source_relative:. *.proto"

var zeroBytes32 common.Hash // deliberately zeroed

const (
	MaxAssertionSize = 120
)

type StakedNodeGraph struct {
	*NodeGraph
	stakers    *StakerSet
	challenges *ChallengeSet
}

func NewStakedNodeGraph(machine machine.Machine, params valprotocol.ChainParams) *StakedNodeGraph {
	return &StakedNodeGraph{
		NodeGraph:  NewNodeGraph(machine, params),
		stakers:    NewStakerSet(),
		challenges: NewChallengeSet(),
	}
}

func (sng *StakedNodeGraph) MarshalForCheckpoint(ctx *ckptcontext.CheckpointContext) *StakedNodeGraphBuf {
	var allStakers []*StakerBuf
	sng.stakers.forall(func(staker *Staker) {
		allStakers = append(allStakers, staker.MarshalToBuf())
	})
	var allChallenges []*ChallengeBuf
	sng.challenges.forall(func(c *Challenge) {
		allChallenges = append(allChallenges, c.MarshalToBuf())
	})
	return &StakedNodeGraphBuf{
		NodeGraph:  sng.NodeGraph.MarshalForCheckpoint(ctx),
		Stakers:    allStakers,
		Challenges: allChallenges,
	}
}

func (x *StakedNodeGraphBuf) UnmarshalFromCheckpoint(ctx ckptcontext.RestoreContext) *StakedNodeGraph {
	chain := &StakedNodeGraph{
		NodeGraph:  x.NodeGraph.UnmarshalFromCheckpoint(ctx),
		stakers:    NewStakerSet(),
		challenges: NewChallengeSet(),
	}
	for _, stakerBuf := range x.Stakers {
		chain.stakers.Add(stakerBuf.Unmarshal(chain.NodeGraph))
	}
	for _, challengeBuf := range x.Challenges {
		chain.challenges.Add(challengeBuf.Unmarshal(chain.NodeGraph))
	}
	return chain
}

func (sng *StakedNodeGraph) DebugString(prefix string) string {
	subPrefix := prefix + "  "
	return "\n" + prefix + "nodes:\n" + sng.NodeGraph.DebugString(sng.stakers, subPrefix) + sng.stakers.DebugString(prefix)
}

func (sng *StakedNodeGraph) Equals(s2 *StakedNodeGraph) bool {
	return sng.NodeGraph.Equals(s2.NodeGraph) &&
		sng.stakers.Equals(s2.stakers) &&
		sng.challenges.Equals(s2.challenges)
}

func (sng *StakedNodeGraph) CreateStake(ev arbbridge.StakeCreatedEvent) {
	nd, ok := sng.nodeFromHash[ev.NodeHash]
	if !ok {
		log.Println("Bad location", ev.NodeHash)
		panic("Tried to create stake on bad node")
	}
	sng.stakers.Add(&Staker{
		ev.Staker,
		nd,
		common.TicksFromBlockNum(ev.BlockId.Height),
		common.Address{},
	})
}

func (sng *StakedNodeGraph) MoveStake(stakerAddr common.Address, nodeHash common.Hash) {
	staker := sng.stakers.Get(stakerAddr)
	if staker == nil {
		log.Fatalf("Moved nonexistant staker %v to node %v", stakerAddr, nodeHash)
	}
	staker.location.RemoveStaker()
	// no need to consider pruning staker.location, because a successor of it is getting a stake
	newLocation, ok := sng.nodeFromHash[nodeHash]
	if !ok {
		log.Fatalf("Moved staker %v to nonexistant node %v", stakerAddr, nodeHash)
	}
	staker.location = newLocation
	staker.location.AddStaker()
}

func (sng *StakedNodeGraph) RemoveStake(stakerAddr common.Address) {
	staker := sng.stakers.Get(stakerAddr)
	staker.location.RemoveStaker()
	sng.considerPruningNode(staker.location)
	sng.stakers.Delete(staker)
}

func (sng *StakedNodeGraph) NewChallenge(chal *Challenge) {
	sng.stakers.Get(chal.asserter).challenge = chal.contract
	sng.stakers.Get(chal.challenger).challenge = chal.contract
	sng.challenges.Add(chal)
}

func (sng *StakedNodeGraph) ChallengeResolved(contract, winner, loser common.Address) {
	sng.stakers.Get(winner).challenge = common.Address{}
	sng.RemoveStake(loser)
	sng.challenges.Delete(contract)
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

func (sng *StakedNodeGraph) generateNodePruneInfo(stakers *StakerSet) []valprotocol.PruneParams {
	var prunesToDo []valprotocol.PruneParams
	sng.leaves.forall(func(leaf *structures.Node) {
		if leaf != sng.latestConfirmed {
			leafAncestor, _, err := structures.GetConflictAncestor(leaf, sng.latestConfirmed)
			if err == nil {
				noStakersOnLeaf := true
				sng.stakers.forall(func(s *Staker) {
					if s.location.Equals(leaf) {
						noStakersOnLeaf = false
					}
				})
				if noStakersOnLeaf {
					prunesToDo = append(prunesToDo, valprotocol.PruneParams{
						LeafHash:     leaf.Hash(),
						AncestorHash: leafAncestor.Prev().Hash(),
						LeafProof:    structures.GeneratePathProof(leafAncestor.Prev(), leaf),
						AncProof:     structures.GeneratePathProof(leafAncestor.Prev(), sng.latestConfirmed),
					})
				}
			}
		}
	})
	return prunesToDo
}

func (sng *StakedNodeGraph) generateNextConfProof(
	currentTime common.TimeTicks,
) *valprotocol.ConfirmOpportunity {
	stakerAddrs := make([]common.Address, 0)
	sng.stakers.forall(func(st *Staker) {
		stakerAddrs = append(stakerAddrs, st.address)
	})
	sort.Sort(SortableAddressList(stakerAddrs))

	nodeOps := make([]valprotocol.ConfirmNodeOpportunity, 0)
	conf := sng.latestConfirmed
	keepChecking := true
	confNodes := make([]*structures.Node, 0)
	for ; keepChecking; keepChecking = false {
		for _, successor := range conf.SuccessorHashes() {
			if successor == zeroBytes32 {
				continue
			}
			nd := sng.nodeFromHash[successor]

			confirmable := sng.isConfirmableNode(
				nd,
				currentTime,
				stakerAddrs,
			)
			if confirmable {
				var confOpp valprotocol.ConfirmNodeOpportunity
				if nd.LinkType() == valprotocol.ValidChildType {
					// We need to know the contents of the actual assertion to confirm it
					// We've only seen the hash accumulator of the messages before whereas this requires the full values
					assertion := nd.Assertion()
					if assertion == nil {
						break
					}
					confOpp = valprotocol.ConfirmValidOpportunity{
						DeadlineTicks:    nd.Deadline(),
						Messages:         assertion.OutMsgs,
						LogsAcc:          nd.Disputable().AssertionClaim.AssertionStub.LastLogHash,
						VMProtoStateHash: nd.VMProtoData().Hash(),
					}

				} else {
					confOpp = valprotocol.ConfirmInvalidOpportunity{
						DeadlineTicks:     nd.Deadline(),
						ChallengeNodeData: nd.NodeDataHash(),
						Branch:            nd.LinkType(),
						VMProtoStateHash:  nd.VMProtoData().Hash(),
					}
				}
				confNodes = append(confNodes, nd)
				nodeOps = append(nodeOps, confOpp)
				conf = nd
				keepChecking = true
				break
			}
		}
	}

	if len(nodeOps) == 0 {
		return nil
	}

	nodeLimit := len(nodeOps)
	for nodeLimit > 0 {
		totalSize := 0
		for _, opp := range nodeOps[:nodeLimit] {
			totalSize += opp.ProofSize()
		}
		proofs := sng.generateAlignedStakersProofs(
			confNodes[nodeLimit-1],
			currentTime,
			stakerAddrs,
		)
		for _, proof := range proofs {
			totalSize += len(proof)
		}
		if totalSize < MaxAssertionSize || nodeLimit == 1 {
			return &valprotocol.ConfirmOpportunity{
				Nodes:                  nodeOps[:nodeLimit],
				CurrentLatestConfirmed: sng.latestConfirmed.Hash(),
				StakerAddresses:        stakerAddrs,
				StakerProofs:           proofs,
			}
		}
	}
	panic("Unreachable code")
}

func (sng *StakedNodeGraph) generateAlignedStakersProofs(
	confirmingNode *structures.Node,
	currentTime common.TimeTicks,
	stakerAddrs []common.Address,
) [][]common.Hash {
	proofs := make([][]common.Hash, 0)
	deadline := confirmingNode.Deadline()
	if currentTime.Cmp(deadline) < 0 {
		return nil
	}
	for _, sa := range stakerAddrs {
		staker := sng.stakers.Get(sa)
		if staker.creationTime.Cmp(deadline) >= 0 {
			continue
		}
		subProof := structures.GeneratePathProof(confirmingNode, staker.location)
		if subProof == nil {
			return nil
		}
		proofs = append(proofs, subProof)
	}
	return proofs
}

func (sng *StakedNodeGraph) isConfirmableNode(
	confirmingNode *structures.Node,
	currentTime common.TimeTicks,
	stakerAddrs []common.Address,
) bool {
	deadline := confirmingNode.Deadline()
	stakeCount := 0
	if currentTime.Cmp(deadline) < 0 {
		return false
	}
	for _, sa := range stakerAddrs {
		staker := sng.stakers.Get(sa)
		if staker.creationTime.Cmp(deadline) >= 0 {
			continue
		}
		subProof := structures.GeneratePathProof(confirmingNode, staker.location)
		if subProof == nil {
			return false
		}
		stakeCount++
	}
	return stakeCount > 0
}

func (sng *StakedNodeGraph) generateStakerPruneInfo() ([]RecoverStakeMootedParams, []RecoverStakeOldParams) {
	var mootedToDo []RecoverStakeMootedParams
	var oldToDo []RecoverStakeOldParams
	sng.stakers.forall(func(staker *Staker) {
		stakerAncestor, _, _, err := GetConflictAncestor(staker.location, sng.latestConfirmed)
		if err == nil {
			prev := stakerAncestor.Prev()
			mootedToDo = append(mootedToDo, RecoverStakeMootedParams{
				addr:         staker.address,
				ancestorHash: prev.Hash(),
				lcProof:      structures.GeneratePathProof(prev, sng.latestConfirmed),
				stProof:      structures.GeneratePathProof(prev, staker.location),
			})
		} else if staker.location.Depth() < sng.latestConfirmed.Depth() {
			oldToDo = append(oldToDo, RecoverStakeOldParams{
				addr:  staker.address,
				proof: structures.GeneratePathProof(staker.location, sng.latestConfirmed),
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
	asserterNodeType      valprotocol.ChildType
	challengerNodeType    valprotocol.ChildType
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
	staker1Ancestor, staker2Ancestor, err := structures.GetConflictAncestor(staker1.location, staker2.location)
	if err != nil {
		return nil
	}
	linkType1 := staker1Ancestor.LinkType()
	linkType2 := staker2Ancestor.LinkType()

	var asserterStaker *Staker
	var asserterAncestor *structures.Node
	var challengerStaker *Staker
	var challengerAncestor *structures.Node
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
		prevNodeHash:          asserterAncestor.Prev().Hash(),
		deadlineTicks:         asserterAncestor.Deadline(),
		asserterNodeType:      asserterAncestor.LinkType(),
		challengerNodeType:    challengerAncestor.LinkType(),
		asserterVMProtoHash:   asserterAncestor.VMProtoData().Hash(),
		challengerVMProtoHash: challengerAncestor.VMProtoData().Hash(),
		asserterProof:         structures.GeneratePathProof(asserterAncestor, asserterStaker.location),
		challengerProof:       structures.GeneratePathProof(challengerAncestor, challengerStaker.location),
		asserterNodeHash:      asserterAncestor.NodeDataHash(),
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

func (sng *StakedNodeGraph) checkChallengeOpportunityAllPairs() []*challengeOpportunity {
	var ret []*challengeOpportunity
	var stakers []*Staker
	sng.stakers.forall(func(s *Staker) {
		stakers = append(stakers, s)
	})
	for i, s1 := range stakers {
		for j := i + 1; j < len(stakers); j++ {
			opp := sng.checkChallengeOpportunityPair(s1, stakers[j])
			if opp != nil {
				ret = append(ret, opp)
				break
			}
		}
	}
	return ret
}
