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

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenges"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type ChainListener interface {
	StakeCreated(ethbridge.StakeCreatedEvent)
	StakeRemoved(ethbridge.StakeRefundedEvent)
	StakeMoved(ethbridge.StakeMovedEvent)
	StartedChallenge(ethbridge.ChallengeStartedEvent, *Node, *Node)
	CompletedChallenge(event ethbridge.ChallengeCompletedEvent)

	AssertionPrepared(*preparedAssertion)
}

type StakerListener struct {
	myAddr   common.Address
	auth     *bind.TransactOpts
	client   *ethclient.Client
	contract *ethbridge.ArbRollup
}

func (staker *StakerListener) initiateChallenge(ctx context.Context, opp *challengeOpportunity) {
	go func() { // we're holding a lock on the chain, so launch the challenge asynchronously
		staker.contract.StartChallenge(
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
			opp.asserterDataHash,
			opp.asserterPeriodTicks,
			opp.challengerNodeHash,
		)
	}()
}

func (staker *StakerListener) makeAssertion(ctx context.Context, opp *preparedAssertion, proof [][32]byte) {
	go func() { // we're holding a lock on the chain, so launch the challenge asynchronously
		staker.contract.MakeAssertion(
			ctx,
			opp.prevPrevLeafHash,
			opp.prevDataHash,
			opp.prevDeadline,
			opp.prevChildType,
			opp.beforeState,
			opp.params,
			opp.claim,
			proof,
		)
	}()
}

func (staker *StakerListener) actAsChallenger(pendingInbox *structures.PendingInbox, ev ethbridge.ChallengeStartedEvent, conflictNode *Node) {
	switch conflictNode.linkType {
	case structures.InvalidPendingChildType:
		go challenges.ChallengePendingTopClaim(
			staker.auth,
			staker.client,
			ev.ChallengeContract,
			pendingInbox,
		)
	case structures.InvalidMessagesChildType:
		go challenges.ChallengeMessagesClaim(
			staker.auth,
			staker.client,
			ev.ChallengeContract,
			pendingInbox,
			conflictNode.vmProtoData.PendingTop,
			conflictNode.disputable.AssertionClaim.AfterPendingTop,
		)
	case structures.InvalidExecutionChildType:
		go challenges.ChallengeExecutionClaim(
			staker.auth,
			staker.client,
			ev.ChallengeContract,
			conflictNode.ExecutionPrecondition(),
			conflictNode.machine,
		)
	}
}

func (staker *StakerListener) actAsAsserter(pendingInbox *structures.PendingInbox, ev ethbridge.ChallengeStartedEvent, conflictNode *Node) {
	switch conflictNode.linkType {
	case structures.InvalidPendingChildType:
		go challenges.DefendPendingTopClaim(
			staker.auth,
			staker.client,
			ev.ChallengeContract,
			pendingInbox,
			conflictNode.disputable.AssertionClaim.AfterPendingTop,
			conflictNode.disputable.MaxPendingTop,
		)
	case structures.InvalidMessagesChildType:
		go challenges.DefendMessagesClaim(
			staker.auth,
			staker.client,
			ev.ChallengeContract,
			pendingInbox,
			conflictNode.vmProtoData.PendingTop,
			conflictNode.disputable.AssertionClaim.AfterPendingTop,
			conflictNode.disputable.AssertionClaim.ImportedMessagesSlice,
		)
	case structures.InvalidExecutionChildType:
		go challenges.DefendExecutionClaim(
			staker.auth,
			staker.client,
			ev.ChallengeContract,
			conflictNode.ExecutionPrecondition(),
			conflictNode.disputable.AssertionParams.NumSteps,
			conflictNode.disputable.AssertionClaim.AssertionStub,
			conflictNode.machine,
		)
	}
}

type ValidatorChainListener struct {
	chain   *ChainObserver
	stakers map[common.Address]*StakerListener
}

func NewValidatorChainListener(
	chain *ChainObserver,
) *ValidatorChainListener {
	return &ValidatorChainListener{chain, make(map[common.Address]*StakerListener)}
}

func (lis *ValidatorChainListener) AddStaker(address common.Address, client *ethclient.Client, auth *bind.TransactOpts) error {
	contract, err := ethbridge.NewRollup(lis.chain.rollupAddr, client, auth)
	if err != nil {
		return err
	}
	lis.stakers[address] = &StakerListener{
		myAddr:   address,
		client:   client,
		contract: contract,
	}
	return nil
}

func (lis *ValidatorChainListener) StakeCreated(ev ethbridge.StakeCreatedEvent) {
	staker, ok := lis.stakers[ev.Staker]
	if ok {
		opps := lis.chain.nodeGraph.checkChallengeOpportunityAllPairs()
		for _, opp := range opps {
			staker.initiateChallenge(context.TODO(), opp)
		}
	} else {
		lis.challengeStakerIfPossible(context.TODO(), ev.Staker)
	}
}

func (lis *ValidatorChainListener) StakeRemoved(ethbridge.StakeRefundedEvent) {

}

func (lis *ValidatorChainListener) StakeMoved(ev ethbridge.StakeMovedEvent) {
	lis.challengeStakerIfPossible(context.TODO(), ev.Staker)
}

func (lis *ValidatorChainListener) challengeStakerIfPossible(ctx context.Context, stakerAddr common.Address) {
	_, ok := lis.stakers[stakerAddr]
	if !ok {
		newStaker := lis.chain.nodeGraph.stakers.Get(stakerAddr)
		for myAddr, staker := range lis.stakers {
			meAsStaker := lis.chain.nodeGraph.stakers.Get(myAddr)
			if meAsStaker != nil {
				opp := lis.chain.nodeGraph.checkChallengeOpportunityPair(newStaker, meAsStaker)
				if opp != nil {
					staker.initiateChallenge(ctx, opp)
					return
				}
			}
			opp := lis.chain.nodeGraph.checkChallengeOpportunityAny(newStaker)
			if opp != nil {
				staker.initiateChallenge(ctx, opp)
				return
			}
		}
	}
}

func (lis *ValidatorChainListener) StartedChallenge(ev ethbridge.ChallengeStartedEvent, asserterAncestor *Node, challengerAncestor *Node) {
	asserter, ok := lis.stakers[ev.Asserter]
	if ok {
		asserter.actAsAsserter(lis.chain.pendingInbox, ev, asserterAncestor)
	}

	challenger, ok := lis.stakers[ev.Challenger]
	if ok {
		challenger.actAsChallenger(lis.chain.pendingInbox, ev, asserterAncestor)
	}
}

func (lis *ValidatorChainListener) CompletedChallenge(ev ethbridge.ChallengeCompletedEvent) {
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

func (lis *ValidatorChainListener) lostChallenge(ethbridge.ChallengeCompletedEvent) {

}

func (lis *ValidatorChainListener) wonChallenge(ethbridge.ChallengeCompletedEvent) {

}

func (lis *ValidatorChainListener) AssertionPrepared(prepared *preparedAssertion) {
	leaf, ok := lis.chain.nodeGraph.nodeFromHash[prepared.prevLeaf]
	if ok {
		for _, staker := range lis.stakers {
			stakerPos := lis.chain.nodeGraph.stakers.Get(staker.myAddr)
			proof := GeneratePathProof(stakerPos.location, leaf)
			if proof != nil {
				staker.makeAssertion(context.TODO(), prepared, proof)
			}
		}
	}
}
