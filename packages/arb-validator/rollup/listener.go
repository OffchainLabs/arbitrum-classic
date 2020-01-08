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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arb"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"log"
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenges"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type ChainListener interface {
	StakeCreated(arbbridge.StakeCreatedEvent)
	StakeRemoved(arbbridge.StakeRefundedEvent)
	StakeMoved(arbbridge.StakeMovedEvent)
	StartedChallenge(arbbridge.ChallengeStartedEvent, *Node, *Node)
	CompletedChallenge(event arbbridge.ChallengeCompletedEvent)
	SawAssertion(arbbridge.AssertedEvent, *protocol.TimeBlocks, [32]byte)
	ConfirmedNode(arbbridge.ConfirmedEvent)

	AssertionPrepared(*preparedAssertion)
	PrunableLeafs([]pruneParams)
	MootableStakes([]recoverStakeMootedParams)
	OldStakes([]recoverStakeOldParams)

	AdvancedKnownAssertion(*protocol.ExecutionAssertion, [32]byte)
}

type StakerListener struct {
	sync.Mutex
	myAddr   common.Address
	auth     *bind.TransactOpts
	client   *ethclient.Client
	contract arbbridge.ArbRollup
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

func (staker *StakerListener) makeAssertion(ctx context.Context, opp *preparedAssertion, proof [][32]byte) error {
	staker.Lock()
	err := staker.contract.MakeAssertion(
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
	staker.Unlock()
	return err
}

func (staker *StakerListener) challengePendingTop(contractAddress common.Address, pendingInbox *structures.PendingInbox) {
	go challenges.ChallengePendingTopClaim(
		staker.auth,
		staker.client,
		contractAddress,
		pendingInbox,
	)
}

func (staker *StakerListener) challengeMessages(contractAddress common.Address, pendingInbox *structures.PendingInbox, conflictNode *Node) {
	go challenges.ChallengeMessagesClaim(
		staker.auth,
		staker.client,
		contractAddress,
		pendingInbox,
		conflictNode.vmProtoData.PendingTop,
		conflictNode.disputable.AssertionClaim.AfterPendingTop,
	)
}

func (staker *StakerListener) challengeExecution(contractAddress common.Address, mach machine.Machine, pre *protocol.Precondition) {
	go challenges.ChallengeExecutionClaim(
		staker.auth,
		staker.client,
		contractAddress,
		pre,
		mach,
	)
}

func (staker *StakerListener) defendPendingTop(contractAddress common.Address, pendingInbox *structures.PendingInbox, conflictNode *Node) {
	go challenges.DefendPendingTopClaim(
		staker.auth,
		staker.client,
		contractAddress,
		pendingInbox,
		conflictNode.disputable.AssertionClaim.AfterPendingTop,
		conflictNode.disputable.MaxPendingTop,
	)
}

func (staker *StakerListener) defendMessages(contractAddress common.Address, pendingInbox *structures.PendingInbox, conflictNode *Node) {
	go challenges.DefendMessagesClaim(
		staker.auth,
		staker.client,
		contractAddress,
		pendingInbox,
		conflictNode.vmProtoData.PendingTop,
		conflictNode.disputable.AssertionClaim.AfterPendingTop,
		conflictNode.disputable.AssertionClaim.ImportedMessagesSlice,
	)
}

func (staker *StakerListener) defendExecution(contractAddress common.Address, mach machine.Machine, pre *protocol.Precondition, numSteps uint32) {
	go challenges.DefendExecutionClaim(
		staker.auth,
		staker.client,
		contractAddress,
		pre,
		numSteps,
		mach,
	)
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

func (lis *ValidatorChainListener) AddStaker(client *ethclient.Client, auth *bind.TransactOpts) error {
	contract, err := arb.NewRollup(lis.chain.rollupAddr, client, auth)
	if err != nil {
		return err
	}
	location := lis.chain.knownValidNode
	proof1 := GeneratePathProof(lis.chain.nodeGraph.latestConfirmed, location)
	proof2 := GeneratePathProof(location, lis.chain.nodeGraph.getLeaf(location))
	go contract.PlaceStake(context.TODO(), lis.chain.nodeGraph.params.StakeRequirement, proof1, proof2)
	address := auth.From
	staker := &StakerListener{
		myAddr:   address,
		client:   client,
		contract: contract,
	}
	lis.stakers[address] = staker
	return nil
}

func (lis *ValidatorChainListener) StakeCreated(ev arbbridge.StakeCreatedEvent) {
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

func (lis *ValidatorChainListener) StakeRemoved(arbbridge.StakeRefundedEvent) {

}

func (lis *ValidatorChainListener) StakeMoved(ev arbbridge.StakeMovedEvent) {
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

func (lis *ValidatorChainListener) StartedChallenge(ev arbbridge.ChallengeStartedEvent, conflictNode *Node, challengerAncestor *Node) {
	asserter, ok := lis.stakers[ev.Asserter]
	if ok {
		switch conflictNode.linkType {
		case structures.InvalidPendingChildType:
			asserter.defendPendingTop(ev.ChallengeContract, lis.chain.pendingInbox, conflictNode)
		case structures.InvalidMessagesChildType:
			asserter.defendMessages(ev.ChallengeContract, lis.chain.pendingInbox, conflictNode)
		case structures.InvalidExecutionChildType:
			asserter.defendExecution(
				ev.ChallengeContract,
				conflictNode.machine,
				lis.chain.ExecutionPrecondition(conflictNode),
				conflictNode.disputable.AssertionParams.NumSteps,
			)
		}
	}

	challenger, ok := lis.stakers[ev.Challenger]
	if ok {
		switch conflictNode.linkType {
		case structures.InvalidPendingChildType:
			challenger.challengePendingTop(ev.ChallengeContract, lis.chain.pendingInbox)
		case structures.InvalidMessagesChildType:
			challenger.challengeMessages(ev.ChallengeContract, lis.chain.pendingInbox, conflictNode)
		case structures.InvalidExecutionChildType:
			challenger.challengeExecution(
				ev.ChallengeContract,
				conflictNode.machine,
				lis.chain.ExecutionPrecondition(conflictNode),
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

func (lis *ValidatorChainListener) SawAssertion(arbbridge.AssertedEvent, *protocol.TimeBlocks, [32]byte) {

}

func (lis *ValidatorChainListener) ConfirmedNode(arbbridge.ConfirmedEvent) {

}

func (lis *ValidatorChainListener) AssertionPrepared(prepared *preparedAssertion) {
	leaf, ok := lis.chain.nodeGraph.nodeFromHash[prepared.prevLeaf]
	if ok {
		for _, staker := range lis.stakers {
			stakerPos := lis.chain.nodeGraph.stakers.Get(staker.myAddr)
			if stakerPos != nil {
				proof := GeneratePathProof(stakerPos.location, leaf)
				if proof != nil {
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

func (lis *ValidatorChainListener) PrunableLeafs(params []pruneParams) {
	for _, staker := range lis.stakers {
		for _, prune := range params {
			go staker.contract.PruneLeaf(
				context.TODO(),
				prune.ancestor.hash,
				prune.leafProof,
				prune.ancProof,
			)
		}
		break
	}

}

func (lis *ValidatorChainListener) MootableStakes(params []recoverStakeMootedParams) {
	for _, staker := range lis.stakers {
		for _, moot := range params {
			go staker.contract.RecoverStakeMooted(
				context.TODO(),
				moot.ancestor.hash,
				moot.addr,
				moot.lcProof,
				moot.stProof,
			)
		}
		break
	}
}

func (lis *ValidatorChainListener) OldStakes(params []recoverStakeOldParams) {
	for _, staker := range lis.stakers {
		for _, old := range params {
			go staker.contract.RecoverStakeOld(
				context.TODO(),
				old.addr,
				old.proof,
			)
		}
		break
	}
}

func (lis *ValidatorChainListener) AdvancedKnownAssertion(*protocol.ExecutionAssertion, [32]byte) {}
