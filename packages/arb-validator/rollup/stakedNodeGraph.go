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
	"sort"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"

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

func (chain *StakedNodeGraph) MarshalForCheckpoint(ctx checkpointing.CheckpointContext) *StakedNodeGraphBuf {
	var allStakers []*StakerBuf
	chain.stakers.forall(func(staker *Staker) {
		allStakers = append(allStakers, staker.MarshalToBuf())
	})
	var allChallenges []*ChallengeBuf
	chain.challenges.forall(func(c *Challenge) {
		allChallenges = append(allChallenges, c.MarshalToBuf())
	})
	return &StakedNodeGraphBuf{
		NodeGraph:  chain.NodeGraph.MarshalForCheckpoint(ctx),
		Stakers:    allStakers,
		Challenges: allChallenges,
	}
}

func (m *StakedNodeGraphBuf) UnmarshalFromCheckpoint(ctx checkpointing.RestoreContext) *StakedNodeGraph {
	chain := &StakedNodeGraph{
		NodeGraph:  m.NodeGraph.UnmarshalFromCheckpoint(ctx),
		stakers:    NewStakerSet(),
		challenges: NewChallengeSet(),
	}
	for _, stakerBuf := range m.Stakers {
		chain.stakers.Add(stakerBuf.Unmarshal(chain.NodeGraph))
	}
	for _, challengeBuf := range m.Challenges {
		chain.challenges.Add(challengeBuf.Unmarshal(chain.NodeGraph))
	}
	return chain
}

func (m *StakedNodeGraph) DebugString(prefix string) string {
	subPrefix := prefix + "  "
	return "\n" + prefix + "nodes:\n" + m.NodeGraph.DebugString(m.stakers, subPrefix) + m.stakers.DebugString(prefix)
}

func (s *StakedNodeGraph) Equals(s2 *StakedNodeGraph) bool {
	return s.NodeGraph.Equals(s2.NodeGraph) &&
		s.stakers.Equals(s2.stakers) &&
		s.challenges.Equals(s2.challenges)
}

func (chain *StakedNodeGraph) CreateStake(ev arbbridge.StakeCreatedEvent) {
	node, ok := chain.nodeFromHash[ev.NodeHash]
	if !ok {
		log.Println("Bad location", ev.NodeHash)
		panic("Tried to create stake on bad node")
	}
	chain.stakers.Add(&Staker{
		ev.Staker,
		node,
		common.TicksFromBlockNum(ev.BlockId.Height),
		common.Address{},
	})
}

func (chain *StakedNodeGraph) MoveStake(stakerAddr common.Address, nodeHash common.Hash) {
	staker := chain.stakers.Get(stakerAddr)
	if staker == nil {
		log.Fatalf("Moved nonexistant staker %v to node %v", stakerAddr, nodeHash)
	}
	staker.location.numStakers--
	// no need to consider pruning staker.location, because a successor of it is getting a stake
	newLocation, ok := chain.nodeFromHash[nodeHash]
	if !ok {
		log.Fatalf("Moved staker %v to nonexistant node %v", stakerAddr, nodeHash)
	}
	staker.location = newLocation
	staker.location.numStakers++
}

func (chain *StakedNodeGraph) RemoveStake(stakerAddr common.Address) {
	staker := chain.stakers.Get(stakerAddr)
	staker.location.numStakers--
	chain.considerPruningNode(staker.location)
	chain.stakers.Delete(staker)
}

func (chain *StakedNodeGraph) NewChallenge(chal *Challenge) {
	chain.stakers.Get(chal.asserter).challenge = chal.contract
	chain.stakers.Get(chal.challenger).challenge = chal.contract
	chain.challenges.Add(chal)
}

func (chain *StakedNodeGraph) ChallengeResolved(contract, winner, loser common.Address) {
	chain.stakers.Get(winner).challenge = common.Address{}
	chain.RemoveStake(loser)
	chain.challenges.Delete(contract)
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

func (chain *StakedNodeGraph) generateNodePruneInfo(stakers *StakerSet) []valprotocol.PruneParams {
	prunesToDo := []valprotocol.PruneParams{}
	chain.leaves.forall(func(leaf *Node) {
		if leaf != chain.latestConfirmed {
			leafAncestor, _, err := GetConflictAncestor(leaf, chain.latestConfirmed)
			if err == nil {
				noStakersOnLeaf := true
				chain.stakers.forall(func(s *Staker) {
					if s.location.Equals(leaf) {
						noStakersOnLeaf = false
					}
				})
				if noStakersOnLeaf {
					prunesToDo = append(prunesToDo, valprotocol.PruneParams{
						LeafHash:     leaf.hash,
						AncestorHash: leafAncestor.prev.hash,
						LeafProof:    GeneratePathProof(leafAncestor.prev, leaf),
						AncProof:     GeneratePathProof(leafAncestor.prev, chain.latestConfirmed),
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
	confNodes := make([]*Node, 0)
	for ; keepChecking; keepChecking = false {
		for _, successor := range conf.successorHashes {
			if successor == zeroBytes32 {
				continue
			}
			node := sng.nodeFromHash[successor]

			confirmable := sng.isConfirmableNode(
				node,
				currentTime,
				stakerAddrs,
			)
			if confirmable {
				var confOpp valprotocol.ConfirmNodeOpportunity
				if node.linkType == valprotocol.ValidChildType {
					// We need to know the contents of the actual assertion to confirm it
					// We've only seen the hash accumulator of the messages before whereas this requires the full values
					if node.assertion == nil {
						break
					}
					confOpp = valprotocol.ConfirmValidOpportunity{
						DeadlineTicks:    node.deadline,
						Messages:         node.assertion.OutMsgs,
						LogsAcc:          node.disputable.AssertionClaim.AssertionStub.LastLogHash,
						VMProtoStateHash: node.vmProtoData.Hash(),
					}

				} else {
					confOpp = valprotocol.ConfirmInvalidOpportunity{
						DeadlineTicks:     node.deadline,
						ChallengeNodeData: node.nodeDataHash,
						Branch:            node.linkType,
						VMProtoStateHash:  node.vmProtoData.Hash(),
					}
				}
				confNodes = append(confNodes, node)
				nodeOps = append(nodeOps, confOpp)
				conf = node
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
				CurrentLatestConfirmed: sng.latestConfirmed.hash,
				StakerAddresses:        stakerAddrs,
				StakerProofs:           proofs,
			}
		}
	}
	panic("Unreachable code")
}

func (sng *StakedNodeGraph) generateAlignedStakersProofs(
	confirmingNode *Node,
	currentTime common.TimeTicks,
	stakerAddrs []common.Address,
) [][]common.Hash {
	proofs := make([][]common.Hash, 0)
	deadline := confirmingNode.deadline
	if currentTime.Cmp(deadline) < 0 {
		return nil
	}
	for _, sa := range stakerAddrs {
		staker := sng.stakers.Get(sa)
		if staker.creationTime.Cmp(deadline) >= 0 {
			continue
		}
		subProof := GeneratePathProof(confirmingNode, staker.location)
		if subProof == nil {
			return nil
		}
		proofs = append(proofs, subProof)
	}
	return proofs
}

func (sng *StakedNodeGraph) isConfirmableNode(
	confirmingNode *Node,
	currentTime common.TimeTicks,
	stakerAddrs []common.Address,
) bool {
	deadline := confirmingNode.deadline
	stakeCount := 0
	if currentTime.Cmp(deadline) < 0 {
		return false
	}
	for _, sa := range stakerAddrs {
		staker := sng.stakers.Get(sa)
		if staker.creationTime.Cmp(deadline) >= 0 {
			continue
		}
		subProof := GeneratePathProof(confirmingNode, staker.location)
		if subProof == nil {
			return false
		}
		stakeCount++
	}
	return stakeCount > 0
}

func (chain *StakedNodeGraph) generateStakerPruneInfo() ([]recoverStakeMootedParams, []recoverStakeOldParams) {
	mootedToDo := []recoverStakeMootedParams{}
	oldToDo := []recoverStakeOldParams{}
	chain.stakers.forall(func(staker *Staker) {
		stakerAncestor, _, _, err := chain.GetConflictAncestor(staker.location, chain.latestConfirmed)
		if err == nil {
			mootedToDo = append(mootedToDo, recoverStakeMootedParams{
				addr:         staker.address,
				ancestorHash: stakerAncestor.prev.hash,
				lcProof:      GeneratePathProof(stakerAncestor.prev, chain.latestConfirmed),
				stProof:      GeneratePathProof(stakerAncestor.prev, staker.location),
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

func (chain *StakedNodeGraph) checkChallengeOpportunityPair(staker1, staker2 *Staker) *challengeOpportunity {
	if !staker1.challenge.IsZero() || !staker2.challenge.IsZero() {
		return nil
	}
	staker1Ancestor, staker2Ancestor, err := GetConflictAncestor(staker1.location, staker2.location)
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

	challengerDataHash, challengerPeriodTicks := challengerAncestor.ChallengeNodeData(chain.params)

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

func (chain *StakedNodeGraph) checkChallengeOpportunityAny(staker *Staker) *challengeOpportunity {
	if !staker.challenge.IsZero() {
		return nil
	}
	var ret *challengeOpportunity
	chain.stakers.forall(func(staker2 *Staker) {
		if !staker2.Equals(staker) {
			opp := chain.checkChallengeOpportunityPair(staker, staker2)
			if opp != nil {
				ret = opp
				return
			}
		}
	})
	return ret
}

func (chain *StakedNodeGraph) checkChallengeOpportunityAllPairs() []*challengeOpportunity {
	ret := []*challengeOpportunity{}
	stakers := []*Staker{}
	chain.stakers.forall(func(s *Staker) {
		stakers = append(stakers, s)
	})
	for i, s1 := range stakers {
		for j := i + 1; j < len(stakers); j++ {
			opp := chain.checkChallengeOpportunityPair(s1, stakers[j])
			if opp != nil {
				ret = append(ret, opp)
				break
			}
		}
	}
	return ret
}
