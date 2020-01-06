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
	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenges"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type ChainListener interface {
	StakeCreated(ethbridge.StakeCreatedEvent)
	StakeRemoved(ethbridge.StakeRefundedEvent)
	StakeMoved(ethbridge.StakeMovedEvent)
	StartedChallenge(ethbridge.ChallengeStartedEvent, *Node, structures.ChildType)
	CompletedChallenge(event ethbridge.ChallengeCompletedEvent)
}

type ValidatorChainListener struct {
	chain  *ChainObserver
	myAddr common.Address
	ch     chan interface{}
}

func NewValidatorChainListener(chain *ChainObserver, myAddr common.Address, runLoop func(*ValidatorChainListener)) {
	ret := &ValidatorChainListener{chain, myAddr, make(chan interface{}, 1024)}
	go runLoop(ret)
}

func (lis *ValidatorChainListener) StakeCreated(ethbridge.StakeCreatedEvent) {

}

func (lis *ValidatorChainListener) StakeRemoved(ethbridge.StakeRefundedEvent) {

}

func (lis *ValidatorChainListener) StakeMoved(ev ethbridge.StakeMovedEvent) {

}

func (lis *ValidatorChainListener) StartedChallenge(ev ethbridge.ChallengeStartedEvent, conflictNode *Node, disputeType structures.ChildType) {
	if utils.AddressesEqual(lis.myAddr, ev.Asserter) {
		lis.actAsAsserter(ev, conflictNode, disputeType)
	}
	if utils.AddressesEqual(lis.myAddr, ev.Challenger) {
		lis.actAsChallenger(ev, conflictNode, disputeType)
	}
}

func (lis *ValidatorChainListener) actAsChallenger(ev ethbridge.ChallengeStartedEvent, conflictNode *Node, disputeType structures.ChildType) {
	switch disputeType {
	case structures.InvalidPendingChildType:
		go challenges.ChallengePendingTopClaim(
			nil,
			nil,
			ev.ChallengeContract,
			lis.chain.pendingInbox,
		)
	case structures.InvalidMessagesChildType:
		go challenges.ChallengeMessagesClaim(
			nil,
			nil,
			ev.ChallengeContract,
			lis.chain.pendingInbox,
			conflictNode.vmProtoData.PendingTop,
			conflictNode.disputable.AssertionClaim.AfterPendingTop,
		)
	case structures.InvalidExecutionChildType:
		go challenges.ChallengeExecutionClaim(
			nil,
			nil,
			ev.ChallengeContract,
			conflictNode.ExecutionPrecondition(),
			conflictNode.machine,
		)
	}
}

func (lis *ValidatorChainListener) actAsAsserter(ev ethbridge.ChallengeStartedEvent, conflictNode *Node, disputeType structures.ChildType) {
	switch disputeType {
	case structures.InvalidPendingChildType:
		go challenges.DefendPendingTopClaim(
			nil,
			nil,
			ev.ChallengeContract,
			lis.chain.pendingInbox,
			conflictNode.disputable.AssertionClaim.AfterPendingTop,
			conflictNode.disputable.MaxPendingTop,
		)
	case structures.InvalidMessagesChildType:
		go challenges.DefendMessagesClaim(
			nil,
			nil,
			ev.ChallengeContract,
			lis.chain.pendingInbox,
			conflictNode.vmProtoData.PendingTop,
			conflictNode.disputable.AssertionClaim.AfterPendingTop,
			conflictNode.disputable.AssertionClaim.ImportedMessagesSlice,
		)
	case structures.InvalidExecutionChildType:
		go challenges.DefendExecutionClaim(
			nil,
			nil,
			ev.ChallengeContract,
			conflictNode.ExecutionPrecondition(),
			conflictNode.disputable.AssertionParams.NumSteps,
			conflictNode.disputable.AssertionClaim.AssertionStub,
			conflictNode.machine,
		)
	}
}

func (lis *ValidatorChainListener) CompletedChallenge(ev ethbridge.ChallengeCompletedEvent) {
	if utils.AddressesEqual(lis.myAddr, ev.Winner) {
		lis.wonChallenge(ev)
	}
	if utils.AddressesEqual(lis.myAddr, ev.Loser) {
		lis.lostChallenge(ev)
	}
}

func (lis *ValidatorChainListener) lostChallenge(ethbridge.ChallengeCompletedEvent) {

}

func (lis *ValidatorChainListener) wonChallenge(ethbridge.ChallengeCompletedEvent) {

}
