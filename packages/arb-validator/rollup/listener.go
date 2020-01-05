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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenges"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type ChainEventListener interface {
	StakeCreated(ethbridge.StakeCreatedEvent)
	StakeRemoved(ethbridge.StakeRefundedEvent)
	StakeMoved(ethbridge.StakeMovedEvent)
	StartedChallenge(ethbridge.ChallengeStartedEvent, *Node, structures.ChildType)
	Challenged(ethbridge.ChallengeStartedEvent, *Node, structures.ChildType)
	LostChallenge(ethbridge.ChallengeCompletedEvent)
	WonChallenge(ethbridge.ChallengeCompletedEvent)
}

type ChanCEListener struct {
	chain *ChainObserver
	ch    chan interface{}
}

func NewChanCEListener(chain *ChainObserver, runLoop func(*ChanCEListener)) {
	ret := &ChanCEListener{chain, make(chan interface{}, 1024)}
	go runLoop(ret)
}

func (lis *ChanCEListener) StakeCreated(ethbridge.StakeCreatedEvent) {

}

func (lis *ChanCEListener) StakeRemoved(ethbridge.StakeRefundedEvent) {

}

func (lis *ChanCEListener) StakeMoved(ev ethbridge.StakeMovedEvent) {

}

func (lis *ChanCEListener) StartedChallenge(ev ethbridge.ChallengeStartedEvent, conflictNode *Node, disputeType structures.ChildType) {
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

func (lis *ChanCEListener) Challenged(ev ethbridge.ChallengeStartedEvent, conflictNode *Node, disputeType structures.ChildType) {
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

func (lis *ChanCEListener) LostChallenge(ethbridge.ChallengeCompletedEvent) {

}

func (lis *ChanCEListener) WonChallenge(ethbridge.ChallengeCompletedEvent) {

}
