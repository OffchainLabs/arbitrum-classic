/*
* Copyright 2020, Offchain Labs, Inc.
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
	"context"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type ChainListener interface {
	StakeCreated(arbbridge.StakeCreatedEvent)
	StakeRemoved(arbbridge.StakeRefundedEvent)
	StakeMoved(arbbridge.StakeMovedEvent)
	StartedChallenge(arbbridge.ChallengeStartedEvent, *Node, *Node)
	CompletedChallenge(event arbbridge.ChallengeCompletedEvent)
	SawAssertion(arbbridge.AssertedEvent, *common.TimeBlocks, common.Hash)
	ConfirmedNode(arbbridge.ConfirmedEvent)
	PrunedLeaf(arbbridge.PrunedEvent)
	MessageDelivered(arbbridge.MessageDeliveredEvent)

	AssertionPrepared(*preparedAssertion)
	ValidNodeConfirmable(*confirmValidOpportunity)
	InvalidNodeConfirmable(*confirmInvalidOpportunity)
	PrunableLeafs([]pruneParams)
	MootableStakes([]recoverStakeMootedParams)
	OldStakes([]recoverStakeOldParams)

	AdvancedKnownValidNode(common.Hash)
	AdvancedKnownAssertion(*protocol.ExecutionAssertion, common.Hash)
}

type ValidatorChainListener struct {
	chain                  *ChainObserver
	stakers                map[common.Address]*StakingKey
	broadcastAssertions    map[common.Hash]bool
	broadcastConfirmations map[common.Hash]bool
	broadcastLeafPrunes    map[common.Hash]bool
}

func NewValidatorChainListener(
	chain *ChainObserver,
) *ValidatorChainListener {
	return &ValidatorChainListener{
		chain:                  chain,
		stakers:                make(map[common.Address]*StakingKey),
		broadcastAssertions:    make(map[common.Hash]bool),
		broadcastConfirmations: make(map[common.Hash]bool),
		broadcastLeafPrunes:    make(map[common.Hash]bool),
	}
}

func (lis *ValidatorChainListener) AddStaker(client arbbridge.ArbAuthClient) error {
	contract, err := client.NewRollup(lis.chain.rollupAddr)
	if err != nil {
		return err
	}

	address := client.Address()
	staker := &StakingKey{
		myAddr:   address,
		client:   client,
		contract: contract,
	}

	isStaked, err := contract.IsStaked(client.Address())
	if err != nil {
		return err
	}
	if !isStaked {
		log.Println("Staking", address.Hex())
		lis.chain.RLock()
		location := lis.chain.knownValidNode
		proof1 := GeneratePathProof(lis.chain.nodeGraph.latestConfirmed, location)
		proof2 := GeneratePathProof(location, lis.chain.nodeGraph.getLeaf(location))
		lis.chain.RUnlock()
		go func() {
			staker.Lock()
			contract.PlaceStake(context.TODO(), lis.chain.nodeGraph.params.StakeRequirement, proof1, proof2)
			staker.Unlock()
		}()
	} else {
		log.Println("Already staked", address.Hex())
	}

	lis.stakers[address] = staker
	return nil
}

func (lis *ValidatorChainListener) StakeCreated(ev arbbridge.StakeCreatedEvent) {
	staker, ok := lis.stakers[ev.Staker]
	if ok {
		opps := lis.chain.nodeGraph.checkChallengeOpportunityAllPairs()
		for _, opp := range opps {
			go staker.initiateChallenge(context.TODO(), opp)
		}
	} else {
		lis.challengeStakerIfPossible(context.TODO(), ev.Staker)
	}
}

func (lis *ValidatorChainListener) StakeRemoved(arbbridge.StakeRefundedEvent) {

}

func (lis *ValidatorChainListener) StakeMoved(ev arbbridge.StakeMovedEvent) {
	lis.challengeStakerIfPossible(context.TODO(), ev.Staker)
}

func (lis *ValidatorChainListener) challengeStakerIfPossible(ctx context.Context, stakerAddr common.Address) {
	_, ok := lis.stakers[stakerAddr]
	if ok {
		// Can't challenge yourself
		return
	}

	newStaker := lis.chain.nodeGraph.stakers.Get(stakerAddr)
	if newStaker == nil {
		log.Fatalf("Nonexistant staker moved %v", stakerAddr)
	}

	for myAddr, staker := range lis.stakers {
		meAsStaker := lis.chain.nodeGraph.stakers.Get(myAddr)
		if meAsStaker == nil {
			continue
		}
		opp := lis.chain.nodeGraph.checkChallengeOpportunityPair(newStaker, meAsStaker)
		if opp != nil {
			staker.initiateChallenge(ctx, opp)
			return
		}
		opp = lis.chain.nodeGraph.checkChallengeOpportunityAny(newStaker)
		if opp != nil {
			go staker.initiateChallenge(ctx, opp)
			return
		}
	}

}

func (lis *ValidatorChainListener) StartedChallenge(ev arbbridge.ChallengeStartedEvent, conflictNode *Node, challengerAncestor *Node) {
	startBlockId := ev.BlockId
	startLogIndex := ev.LogIndex - 1
	asserter, ok := lis.stakers[ev.Asserter]
	if ok {
		switch conflictNode.linkType {
		case structures.InvalidPendingChildType:
			go asserter.defendPendingTop(ev.ChallengeContract, startBlockId, startLogIndex, lis.chain.pendingInbox, conflictNode)
		case structures.InvalidMessagesChildType:
			go asserter.defendMessages(ev.ChallengeContract, startBlockId, startLogIndex, lis.chain.pendingInbox, conflictNode)
		case structures.InvalidExecutionChildType:
			go asserter.defendExecution(
				ev.ChallengeContract,
				startBlockId,
				startLogIndex,
				conflictNode.machine,
				lis.chain.executionPrecondition(conflictNode),
				conflictNode.disputable.AssertionParams.NumSteps,
			)
		}
	}

	challenger, ok := lis.stakers[ev.Challenger]
	if ok {
		switch conflictNode.linkType {
		case structures.InvalidPendingChildType:
			go challenger.challengePendingTop(ev.ChallengeContract, startBlockId, startLogIndex, lis.chain.pendingInbox)
		case structures.InvalidMessagesChildType:
			go challenger.challengeMessages(ev.ChallengeContract, startBlockId, startLogIndex, lis.chain.pendingInbox, conflictNode)
		case structures.InvalidExecutionChildType:
			go challenger.challengeExecution(
				ev.ChallengeContract,
				startBlockId,
				startLogIndex,
				conflictNode.machine,
				lis.chain.executionPrecondition(conflictNode),
			)
		}
	}
}

func (lis *ValidatorChainListener) CompletedChallenge(ev arbbridge.ChallengeCompletedEvent) {
	_, ok := lis.stakers[ev.Winner]
	if ok {
		lis.wonChallenge(ev)
	}

	_, ok = lis.stakers[ev.Loser]
	if ok {
		lis.lostChallenge(ev)
	}
	lis.challengeStakerIfPossible(context.TODO(), ev.Winner)
}

func (lis *ValidatorChainListener) lostChallenge(arbbridge.ChallengeCompletedEvent) {

}

func (lis *ValidatorChainListener) wonChallenge(arbbridge.ChallengeCompletedEvent) {

}

func (lis *ValidatorChainListener) SawAssertion(arbbridge.AssertedEvent, *common.TimeBlocks, common.Hash) {

}

func (lis *ValidatorChainListener) ConfirmedNode(arbbridge.ConfirmedEvent) {

}

func (lis *ValidatorChainListener) PrunedLeaf(arbbridge.PrunedEvent) {

}

func (lis *ValidatorChainListener) MessageDelivered(arbbridge.MessageDeliveredEvent) {

}

func (lis *ValidatorChainListener) AssertionPrepared(prepared *preparedAssertion) {
	_, alreadySent := lis.broadcastAssertions[prepared.leafHash]
	if alreadySent {
		return
	}
	leaf, ok := lis.chain.nodeGraph.nodeFromHash[prepared.leafHash]
	if ok {
		for _, staker := range lis.stakers {
			stakerPos := lis.chain.nodeGraph.stakers.Get(staker.myAddr)
			if stakerPos != nil {
				proof := GeneratePathProof(stakerPos.location, leaf)
				if proof != nil {
					lis.broadcastAssertions[prepared.leafHash] = true
					go func() {
						err := staker.makeAssertion(context.TODO(), prepared, proof)
						if err != nil {
							log.Println("Error making assertion", err)
						} else {
							log.Println("Successfully made assertion")
						}
					}()

					break
				}
			}
		}
	}
}

func (lis *ValidatorChainListener) ValidNodeConfirmable(conf *confirmValidOpportunity) {
	_, alreadySent := lis.broadcastConfirmations[conf.nodeHash]
	if alreadySent {
		return
	}
	for _, staker := range lis.stakers {
		lis.broadcastConfirmations[conf.nodeHash] = true
		go func() {
			staker.Lock()
			staker.contract.ConfirmValid(
				context.TODO(),
				conf.deadlineTicks,
				conf.messages,
				conf.logsAcc,
				conf.vmProtoStateHash,
				conf.stakerAddresses,
				conf.stakerProofs,
				conf.stakerProofOffsets,
			)
			staker.Unlock()
		}()
		break
	}
}

func (lis *ValidatorChainListener) InvalidNodeConfirmable(conf *confirmInvalidOpportunity) {
	_, alreadySent := lis.broadcastConfirmations[conf.nodeHash]
	if alreadySent {
		return
	}
	for _, staker := range lis.stakers {
		lis.broadcastConfirmations[conf.nodeHash] = true
		go func() {
			staker.Lock()
			staker.contract.ConfirmInvalid(
				context.TODO(),
				conf.deadlineTicks,
				conf.challengeNodeData,
				conf.branch,
				conf.vmProtoStateHash,
				conf.stakerAddresses,
				conf.stakerProofs,
				conf.stakerProofOffsets,
			)
			staker.Unlock()
		}()
		break
	}
}

func (lis *ValidatorChainListener) PrunableLeafs(params []pruneParams) {
	for _, staker := range lis.stakers {
		for _, prune := range params {
			_, alreadySent := lis.broadcastLeafPrunes[prune.leafHash]
			if alreadySent {
				continue
			}
			lis.broadcastLeafPrunes[prune.leafHash] = true
			pruneCopy := prune.Clone()
			go func() {
				staker.Lock()
				staker.contract.PruneLeaf(
					context.TODO(),
					pruneCopy.ancestorHash,
					pruneCopy.leafProof,
					pruneCopy.ancProof,
				)
				staker.Unlock()
			}()
		}
		break
	}
}

func (lis *ValidatorChainListener) MootableStakes(params []recoverStakeMootedParams) {
	for _, staker := range lis.stakers {
		for _, moot := range params {
			go func() {
				staker.Lock()
				staker.contract.RecoverStakeMooted(
					context.TODO(),
					moot.ancestorHash,
					moot.addr,
					moot.lcProof,
					moot.stProof,
				)
				staker.Unlock()
			}()
		}
		break
	}
}

func (lis *ValidatorChainListener) OldStakes(params []recoverStakeOldParams) {
	for _, staker := range lis.stakers {
		for _, old := range params {
			go func() {
				staker.Lock()
				staker.contract.RecoverStakeOld(
					context.TODO(),
					old.addr,
					old.proof,
				)
				staker.Unlock()
			}()
		}
		break
	}
}

func (lis *ValidatorChainListener) AdvancedKnownValidNode(common.Hash)                               {}
func (lis *ValidatorChainListener) AdvancedKnownAssertion(*protocol.ExecutionAssertion, common.Hash) {}
