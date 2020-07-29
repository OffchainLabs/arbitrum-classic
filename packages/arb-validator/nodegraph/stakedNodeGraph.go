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

package nodegraph

import (
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"sort"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

var ZeroBytes32 common.Hash // deliberately zeroed

type RecoverStakeOldParams struct {
	Addr  common.Address
	Proof []common.Hash
}

type RecoverStakeMootedParams struct {
	Addr         common.Address
	AncestorHash common.Hash
	LcProof      []common.Hash
	StProof      []common.Hash
}

const (
	MaxAssertionSize = 120
)

type StakedNodeGraph struct {
	*NodeGraph
	stakers    *StakerSet
	Challenges *ChallengeSet
}

func (sg *StakedNodeGraph) Stakers() *StakerSet {
	return sg.stakers
}

func NewStakedNodeGraph(machine machine.Machine, params valprotocol.ChainParams, creationTxHash common.Hash) *StakedNodeGraph {
	return &StakedNodeGraph{
		NodeGraph:  NewNodeGraph(machine, params, creationTxHash),
		stakers:    NewStakerSet(),
		Challenges: NewChallengeSet(),
	}
}

func (sng *StakedNodeGraph) MarshalForCheckpoint(ctx *ckptcontext.CheckpointContext) *StakedNodeGraphBuf {
	var allStakers []*StakerBuf
	sng.stakers.forall(func(staker *Staker) {
		allStakers = append(allStakers, staker.MarshalToBuf())
	})
	var allChallenges []*ChallengeBuf
	sng.Challenges.Forall(func(c *Challenge) {
		allChallenges = append(allChallenges, c.MarshalToBuf())
	})
	return &StakedNodeGraphBuf{
		NodeGraph:  sng.NodeGraph.MarshalForCheckpoint(ctx),
		Stakers:    allStakers,
		Challenges: allChallenges,
	}
}

func (x *StakedNodeGraphBuf) UnmarshalFromCheckpoint(ctx ckptcontext.RestoreContext) (*StakedNodeGraph, error) {
	nodeGraph, err := x.NodeGraph.UnmarshalFromCheckpoint(ctx)
	if err != nil {
		return nil, err
	}
	chain := &StakedNodeGraph{
		NodeGraph:  nodeGraph,
		stakers:    NewStakerSet(),
		Challenges: NewChallengeSet(),
	}
	for _, stakerBuf := range x.Stakers {
		chain.stakers.Add(stakerBuf.Unmarshal(chain.NodeGraph))
	}
	for _, challengeBuf := range x.Challenges {
		chain.Challenges.Add(challengeBuf.Unmarshal(chain.NodeGraph))
	}
	return chain, nil
}

func (sng *StakedNodeGraph) DebugString(prefix string, labels map[*structures.Node][]string) string {
	subPrefix := prefix + "  "
	return "\n" + prefix + "nodes:\n" + sng.NodeGraph.DebugString(sng.stakers, subPrefix, labels) + sng.stakers.DebugString(prefix)
}

func (sng *StakedNodeGraph) Equals(s2 *StakedNodeGraph) bool {
	return sng.NodeGraph.Equals(s2.NodeGraph) &&
		sng.stakers.Equals(s2.stakers) &&
		sng.Challenges.Equals(s2.Challenges)
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
	sng.ConsiderPruningNode(staker.location)
	sng.stakers.Delete(staker)
}

func (sng *StakedNodeGraph) NewChallenge(chal *Challenge) {
	sng.stakers.Get(chal.asserter).challenge = chal.contract
	sng.stakers.Get(chal.challenger).challenge = chal.contract
	sng.Challenges.Add(chal)
}

func (sng *StakedNodeGraph) ChallengeResolved(contract, winner, loser common.Address) {
	sng.stakers.Get(winner).challenge = common.Address{}
	sng.RemoveStake(loser)
	sng.Challenges.Delete(contract)
}

func (sng *StakedNodeGraph) GenerateNodePruneInfo(stakers *StakerSet) []valprotocol.PruneParams {
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
					prunesToDo = append(prunesToDo, newPruneParams(leaf, leafAncestor, sng.latestConfirmed))
				}
			}
		}
	})
	return prunesToDo
}

func (sng *StakedNodeGraph) GenerateNextConfProof(
	currentTime common.TimeTicks,
) (*valprotocol.ConfirmOpportunity, []*structures.Node) {
	stakerAddrs := make([]common.Address, 0)
	sng.stakers.forall(func(st *Staker) {
		stakerAddrs = append(stakerAddrs, st.address)
	})
	sort.Sort(SortableAddressList(stakerAddrs))

	nodeOps := make([]valprotocol.ConfirmNodeOpportunity, 0)
	currentConfirmedNode := sng.latestConfirmed
	keepChecking := true
	confNodes := make([]*structures.Node, 0)
	for ; keepChecking; keepChecking = false {
		for _, successor := range currentConfirmedNode.SuccessorHashes() {
			if successor == ZeroBytes32 {
				continue
			}
			currentNode := sng.nodeFromHash[successor]

			confirmable := sng.isConfirmableNode(
				currentNode,
				currentTime,
				stakerAddrs,
			)
			if confirmable {
				confOpp := confirmNodeOpp(currentNode)

				if confOpp != nil {
					confNodes = append(confNodes, currentNode)
					nodeOps = append(nodeOps, confOpp)
					currentConfirmedNode = currentNode
					keepChecking = true
				}
				break
			}
		}
	}

	if len(nodeOps) == 0 {
		return nil, nil
	} else {
		return sng.makeConfirmOpp(
			nodeOps,
			confNodes,
			currentTime,
			stakerAddrs)
	}
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
			proofs = append(proofs, []common.Hash{})
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

func (sng *StakedNodeGraph) GenerateStakerPruneInfo() ([]RecoverStakeMootedParams, []RecoverStakeOldParams) {
	var mootedToDo []RecoverStakeMootedParams
	var oldToDo []RecoverStakeOldParams
	sng.stakers.forall(func(staker *Staker) {
		stakerAncestor, _, _, err := GetConflictAncestor(staker.location, sng.latestConfirmed)
		if err == nil {
			prev := stakerAncestor.Prev()
			mootedToDo = append(mootedToDo, RecoverStakeMootedParams{
				Addr:         staker.address,
				AncestorHash: stakerAncestor.PrevHash(),
				LcProof:      structures.GeneratePathProof(prev, sng.latestConfirmed),
				StProof:      structures.GeneratePathProof(prev, staker.location),
			})
		} else if staker.location.Depth() < sng.latestConfirmed.Depth() {
			oldToDo = append(oldToDo, RecoverStakeOldParams{
				Addr:  staker.address,
				Proof: structures.GeneratePathProof(staker.location, sng.latestConfirmed),
			})
		}
	})
	return mootedToDo, oldToDo
}

func (sng *StakedNodeGraph) CheckChallengeOpportunityPair(staker1, staker2 *Staker) *ChallengeOpportunity {
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

	return &ChallengeOpportunity{
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

func (sng *StakedNodeGraph) CheckChallengeOpportunityAny(staker *Staker) *ChallengeOpportunity {
	if !staker.challenge.IsZero() {
		return nil
	}
	var ret *ChallengeOpportunity
	sng.stakers.forall(func(staker2 *Staker) {
		if !staker2.Equals(staker) {
			opp := sng.CheckChallengeOpportunityPair(staker, staker2)
			if opp != nil {
				ret = opp
				return
			}
		}
	})
	return ret
}

func (sng *StakedNodeGraph) checkChallengeOpportunityAllPairs() []*ChallengeOpportunity {
	var ret []*ChallengeOpportunity
	var stakers []*Staker
	sng.stakers.forall(func(s *Staker) {
		stakers = append(stakers, s)
	})
	for i, s1 := range stakers {
		for j := i + 1; j < len(stakers); j++ {
			opp := sng.CheckChallengeOpportunityPair(s1, stakers[j])
			if opp != nil {
				ret = append(ret, opp)
				break
			}
		}
	}
	return ret
}

func (sng *StakedNodeGraph) makeConfirmOpp(
	nodeOps []valprotocol.ConfirmNodeOpportunity,
	confNodes []*structures.Node,
	currentTime common.TimeTicks,
	stakerAddrs []common.Address,
) (*valprotocol.ConfirmOpportunity, []*structures.Node) {
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
			}, confNodes
		}
	}
	panic("Unreachable code")
}
