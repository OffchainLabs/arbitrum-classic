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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenges"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type ChainListener interface {
	StakeCreated(*ChainObserver, arbbridge.StakeCreatedEvent)
	StakeRemoved(*ChainObserver, arbbridge.StakeRefundedEvent)
	StakeMoved(*ChainObserver, arbbridge.StakeMovedEvent)
	StartedChallenge(*ChainObserver, arbbridge.ChallengeStartedEvent, *Node, *Node)
	CompletedChallenge(*ChainObserver, arbbridge.ChallengeCompletedEvent)
	SawAssertion(*ChainObserver, arbbridge.AssertedEvent, *common.TimeBlocks, common.Hash)
	ConfirmedNode(*ChainObserver, arbbridge.ConfirmedEvent)
	PrunedLeaf(*ChainObserver, arbbridge.PrunedEvent)
	MessageDelivered(*ChainObserver, arbbridge.MessageDeliveredEvent)

	AssertionPrepared(*ChainObserver, *preparedAssertion)
	ValidNodeConfirmable(*ChainObserver, *confirmValidOpportunity)
	InvalidNodeConfirmable(*ChainObserver, *confirmInvalidOpportunity)
	PrunableLeafs(*ChainObserver, []pruneParams)
	MootableStakes(*ChainObserver, []recoverStakeMootedParams)
	OldStakes(*ChainObserver, []recoverStakeOldParams)

	AdvancedKnownValidNode(*ChainObserver, common.Hash)
	AdvancedKnownAssertion(*ChainObserver, *protocol.ExecutionAssertion, common.Hash)
}

type StakingKey struct {
	client   arbbridge.ArbAuthClient
	contract arbbridge.ArbRollup
}

type ValidatorChainListener struct {
	actor                  arbbridge.ArbRollup
	rollupAddress          common.Address
	stakingKeys            map[common.Address]*StakingKey
	broadcastAssertions    map[common.Hash]*structures.AssertionParams
	broadcastConfirmations map[common.Hash]bool
	broadcastLeafPrunes    map[common.Hash]bool
}

func NewValidatorChainListener(rollupAddress common.Address, actor arbbridge.ArbRollup) *ValidatorChainListener {
	return &ValidatorChainListener{
		actor:                  actor,
		rollupAddress:          rollupAddress,
		stakingKeys:            make(map[common.Address]*StakingKey),
		broadcastAssertions:    make(map[common.Hash]*structures.AssertionParams),
		broadcastConfirmations: make(map[common.Hash]bool),
		broadcastLeafPrunes:    make(map[common.Hash]bool),
	}
}

func stakeLatestValid(ctx context.Context, chain *ChainObserver, stakingKey *StakingKey) error {
	location := chain.knownValidNode
	proof1 := GeneratePathProof(chain.nodeGraph.latestConfirmed, location)
	proof2 := GeneratePathProof(location, chain.nodeGraph.getLeaf(location))
	stakeAmount := chain.nodeGraph.params.StakeRequirement

	log.Println("Placing stake for", stakingKey.client.Address())
	return stakingKey.contract.PlaceStake(ctx, stakeAmount, proof1, proof2)
}

func (lis *ValidatorChainListener) AddStaker(client arbbridge.ArbAuthClient) error {
	contract, err := client.NewRollup(lis.rollupAddress)
	if err != nil {
		return err
	}

	address := client.Address()
	lis.stakingKeys[address] = &StakingKey{
		client:   client,
		contract: contract,
	}
	return nil
}

func makeAssertion(ctx context.Context, rollup arbbridge.ArbRollup, prepared *preparedAssertion, proof []common.Hash) {
	err := rollup.MakeAssertion(
		context.TODO(),
		prepared.prevPrevLeafHash,
		prepared.prevDataHash,
		prepared.prevDeadline,
		prepared.prevChildType,
		prepared.beforeState,
		prepared.params,
		prepared.claim,
		proof,
	)
	if err != nil {
		log.Println("Error making assertion", err)
	} else {
		log.Println("Successfully made assertion")
	}
}

func (lis *ValidatorChainListener) AssertionPrepared(chain *ChainObserver, prepared *preparedAssertion) {
	// Anyone confirm a node
	// No need to have your own stake
	prevParams, alreadySent := lis.broadcastAssertions[prepared.leafHash]
	if alreadySent && prevParams.Equals(prepared.params) {
		return
	}

	leaf, ok := chain.nodeGraph.nodeFromHash[prepared.leafHash]
	if !ok {
		log.Println("Prepared assertion on top of invalid node")
		return
	}

	for stakingAddress, stakingKey := range lis.stakingKeys {
		stakerPos := chain.nodeGraph.stakers.Get(stakingAddress)
		if stakerPos == nil {
			// stakingKey is not staked
			continue
		}
		proof := GeneratePathProof(stakerPos.location, leaf)
		if proof == nil {
			// staker can't move to new asertion
			continue
		}
		lis.broadcastAssertions[prepared.leafHash] = prepared.params
		go makeAssertion(context.TODO(), stakingKey.contract, prepared, proof)
		return
	}

	for stakingAddress, stakingKey := range lis.stakingKeys {
		stakerPos := chain.nodeGraph.stakers.Get(stakingAddress)
		if stakerPos != nil {
			// stakingKey is already
			continue
		}
		// Put down new stake so that we can assert next time
		go stakeLatestValid(context.TODO(), chain, stakingKey)
		return
	}
}

func (lis *ValidatorChainListener) initiateChallenge(ctx context.Context, opp *challengeOpportunity) error {
	return lis.actor.StartChallenge(
		ctx,
		opp.asserter,
		opp.challenger,
		opp.prevNodeHash,
		opp.deadlineTicks.Val,
		opp.asserterNodeType,
		opp.challengerNodeType,
		opp.asserterVMProtoHash,
		opp.challengerVMProtoHash,
		opp.asserterProof,
		opp.challengerProof,
		opp.asserterNodeHash,
		opp.challengerDataHash,
		opp.challengerPeriodTicks,
	)
}

func (lis *ValidatorChainListener) StakeCreated(chain *ChainObserver, ev arbbridge.StakeCreatedEvent) {
	_, ok := lis.stakingKeys[ev.Staker]
	if ok {
		staker := chain.nodeGraph.stakers.Get(ev.Staker)
		if staker == nil {
			panic("Stake created but address is not in graph")
		}
		opp := chain.nodeGraph.checkChallengeOpportunityAny(staker)
		if opp != nil {
			go lis.initiateChallenge(context.TODO(), opp)
		}
	} else {
		lis.challengeStakerIfPossible(context.TODO(), chain, ev.Staker)
	}
}

func (lis *ValidatorChainListener) StakeMoved(chain *ChainObserver, ev arbbridge.StakeMovedEvent) {
	lis.challengeStakerIfPossible(context.TODO(), chain, ev.Staker)
}

func (lis *ValidatorChainListener) challengeStakerIfPossible(ctx context.Context, chain *ChainObserver, stakerAddr common.Address) {
	_, ok := lis.stakingKeys[stakerAddr]
	if ok {
		// Don't challenge yourself
		return
	}

	newStaker := chain.nodeGraph.stakers.Get(stakerAddr)
	if newStaker == nil {
		log.Fatalf("Nonexistant staker moved %v", stakerAddr)
	}

	// Search for an already staked staking key
	for myAddr, _ := range lis.stakingKeys {
		meAsStaker := chain.nodeGraph.stakers.Get(myAddr)
		if meAsStaker == nil {
			continue
		}
		opp := chain.nodeGraph.checkChallengeOpportunityPair(newStaker, meAsStaker)
		if opp != nil {
			go lis.initiateChallenge(context.TODO(), opp)
			return
		}
	}
	opp := chain.nodeGraph.checkChallengeOpportunityAny(newStaker)
	if opp != nil {
		go lis.initiateChallenge(context.TODO(), opp)
		return
	}
}

// All functions below are either only called if you have a stake down, or don't require a stake

func (lis *ValidatorChainListener) StartedChallenge(chain *ChainObserver, ev arbbridge.ChallengeStartedEvent, conflictNode *Node, challengerAncestor *Node) {
	// Must already be staked to be challenged
	startBlockId := ev.BlockId
	startLogIndex := ev.LogIndex - 1
	asserterKey, ok := lis.stakingKeys[ev.Asserter]
	if ok {
		switch conflictNode.linkType {
		case structures.InvalidPendingChildType:
			go challenges.DefendPendingTopClaim(
				asserterKey.client,
				ev.ChallengeContract,
				startBlockId,
				startLogIndex,
				chain.pendingInbox.MessageStack,
				conflictNode.disputable.AssertionClaim.AfterPendingTop,
				conflictNode.disputable.MaxPendingTop,
				100,
			)
		case structures.InvalidMessagesChildType:
			go challenges.DefendMessagesClaim(
				asserterKey.client,
				ev.ChallengeContract,
				startBlockId,
				startLogIndex,
				chain.pendingInbox.MessageStack,
				conflictNode.vmProtoData.PendingTop,
				conflictNode.disputable.AssertionClaim.AfterPendingTop,
				conflictNode.disputable.AssertionClaim.ImportedMessagesSlice,
				100,
			)
		case structures.InvalidExecutionChildType:
			go challenges.DefendExecutionClaim(
				asserterKey.client,
				ev.ChallengeContract,
				startBlockId,
				startLogIndex,
				chain.executionPrecondition(conflictNode),
				conflictNode.machine,
				conflictNode.disputable.AssertionParams.NumSteps,
				50,
			)
		}
	}

	challenger, ok := lis.stakingKeys[ev.Challenger]
	if ok {
		switch conflictNode.linkType {
		case structures.InvalidPendingChildType:
			go challenges.ChallengePendingTopClaim(
				challenger.client,
				ev.ChallengeContract,
				startBlockId,
				startLogIndex,
				chain.pendingInbox.MessageStack,
			)
		case structures.InvalidMessagesChildType:
			go challenges.ChallengeMessagesClaim(
				challenger.client,
				ev.ChallengeContract,
				startBlockId,
				startLogIndex,
				chain.pendingInbox.MessageStack,
				conflictNode.vmProtoData.PendingTop,
				conflictNode.disputable.AssertionClaim.AfterPendingTop,
			)
		case structures.InvalidExecutionChildType:
			go challenges.ChallengeExecutionClaim(
				challenger.client,
				ev.ChallengeContract,
				startBlockId,
				startLogIndex,
				chain.executionPrecondition(conflictNode),
				conflictNode.machine,
				false,
			)
		}
	}
}

func (lis *ValidatorChainListener) CompletedChallenge(chain *ChainObserver, ev arbbridge.ChallengeCompletedEvent) {
	// Must be staked to have challenge completed
	_, ok := lis.stakingKeys[ev.Winner]
	if ok {
		lis.wonChallenge(ev)
	}

	_, ok = lis.stakingKeys[ev.Loser]
	if ok {
		lis.lostChallenge(ev)
	}
	lis.challengeStakerIfPossible(context.TODO(), chain, ev.Winner)
}

func (lis *ValidatorChainListener) ValidNodeConfirmable(observer *ChainObserver, conf *confirmValidOpportunity) {
	// Anyone confirm a node
	// No need to have your own stake
	_, alreadySent := lis.broadcastConfirmations[conf.nodeHash]
	if alreadySent {
		return
	}
	lis.broadcastConfirmations[conf.nodeHash] = true
	go func() {
		lis.actor.ConfirmValid(
			context.TODO(),
			conf.deadlineTicks,
			conf.messages,
			conf.logsAcc,
			conf.vmProtoStateHash,
			conf.stakerAddresses,
			conf.stakerProofs,
			conf.stakerProofOffsets,
		)
	}()
}

func (lis *ValidatorChainListener) InvalidNodeConfirmable(observer *ChainObserver, conf *confirmInvalidOpportunity) {
	// Anyone confirm a node
	// No need to have your own stake
	_, alreadySent := lis.broadcastConfirmations[conf.nodeHash]
	if alreadySent {
		return
	}
	lis.broadcastConfirmations[conf.nodeHash] = true
	go func() {
		lis.actor.ConfirmInvalid(
			context.TODO(),
			conf.deadlineTicks,
			conf.challengeNodeData,
			conf.branch,
			conf.vmProtoStateHash,
			conf.stakerAddresses,
			conf.stakerProofs,
			conf.stakerProofOffsets,
		)
	}()
}

func (lis *ValidatorChainListener) PrunableLeafs(observer *ChainObserver, params []pruneParams) {
	// Anyone can prune a leaf
	for _, prune := range params {
		_, alreadySent := lis.broadcastLeafPrunes[prune.leafHash]
		if alreadySent {
			continue
		}
		lis.broadcastLeafPrunes[prune.leafHash] = true
		pruneCopy := prune.Clone()
		go func() {
			lis.actor.PruneLeaf(
				context.TODO(),
				pruneCopy.ancestorHash,
				pruneCopy.leafProof,
				pruneCopy.ancProof,
			)
		}()
	}
}

func (lis *ValidatorChainListener) MootableStakes(observer *ChainObserver, params []recoverStakeMootedParams) {
	// Anyone can moot any stake
	for _, moot := range params {
		go func() {
			lis.actor.RecoverStakeMooted(
				context.TODO(),
				moot.ancestorHash,
				moot.addr,
				moot.lcProof,
				moot.stProof,
			)
		}()
	}
}

func (lis *ValidatorChainListener) OldStakes(observer *ChainObserver, params []recoverStakeOldParams) {
	// Anyone can remove an old stake
	for _, old := range params {
		go func() {
			lis.actor.RecoverStakeOld(
				context.TODO(),
				old.addr,
				old.proof,
			)
		}()
	}
}

func (lis *ValidatorChainListener) StakeRemoved(*ChainObserver, arbbridge.StakeRefundedEvent) {}
func (lis *ValidatorChainListener) lostChallenge(arbbridge.ChallengeCompletedEvent)           {}
func (lis *ValidatorChainListener) wonChallenge(arbbridge.ChallengeCompletedEvent)            {}
func (lis *ValidatorChainListener) SawAssertion(*ChainObserver, arbbridge.AssertedEvent, *common.TimeBlocks, common.Hash) {
}
func (lis *ValidatorChainListener) ConfirmedNode(*ChainObserver, arbbridge.ConfirmedEvent)           {}
func (lis *ValidatorChainListener) PrunedLeaf(*ChainObserver, arbbridge.PrunedEvent)                 {}
func (lis *ValidatorChainListener) MessageDelivered(*ChainObserver, arbbridge.MessageDeliveredEvent) {}

func (lis *ValidatorChainListener) AdvancedKnownValidNode(*ChainObserver, common.Hash) {}
func (lis *ValidatorChainListener) AdvancedKnownAssertion(*ChainObserver, *protocol.ExecutionAssertion, common.Hash) {
}
