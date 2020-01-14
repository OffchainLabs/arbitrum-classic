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
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenges"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

type StakingKey struct {
	sync.Mutex
	myAddr   common.Address
	client   arbbridge.ArbAuthClient
	contract arbbridge.ArbRollup
}

func (staker *StakingKey) initiateChallenge(ctx context.Context, opp *challengeOpportunity) {
	staker.Lock()
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
		opp.asserterNodeHash,
		opp.challengerDataHash,
		opp.challengerPeriodTicks,
	)
	staker.Unlock()
}

func (staker *StakingKey) makeAssertion(ctx context.Context, opp *preparedAssertion, proof []common.Hash) error {
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

func (staker *StakingKey) challengePendingTop(contractAddress common.Address, pendingInbox *structures.PendingInbox) {
	staker.Lock()
	challenges.ChallengePendingTopClaim(
		staker.client,
		contractAddress,
		pendingInbox.MessageStack,
	)
	staker.Unlock()
}

func (staker *StakingKey) challengeMessages(contractAddress common.Address, pendingInbox *structures.PendingInbox, conflictNode *Node) {
	staker.Lock()
	challenges.ChallengeMessagesClaim(
		staker.client,
		contractAddress,
		pendingInbox.MessageStack,
		conflictNode.vmProtoData.PendingTop,
		conflictNode.disputable.AssertionClaim.AfterPendingTop,
	)
	staker.Unlock()
}

func (staker *StakingKey) challengeExecution(contractAddress common.Address, mach machine.Machine, pre *valprotocol.Precondition) {
	staker.Lock()
	challenges.ChallengeExecutionClaim(
		staker.client,
		contractAddress,
		pre,
		mach,
		false,
	)
	staker.Unlock()
}

func (staker *StakingKey) defendPendingTop(contractAddress common.Address, pendingInbox *structures.PendingInbox, conflictNode *Node) {
	staker.Lock()
	challenges.DefendPendingTopClaim(
		staker.client,
		contractAddress,
		pendingInbox.MessageStack,
		conflictNode.disputable.AssertionClaim.AfterPendingTop,
		conflictNode.disputable.MaxPendingTop,
		100,
	)
	staker.Unlock()
}

func (staker *StakingKey) defendMessages(contractAddress common.Address, pendingInbox *structures.PendingInbox, conflictNode *Node) {
	staker.Lock()
	challenges.DefendMessagesClaim(
		staker.client,
		contractAddress,
		pendingInbox.MessageStack,
		conflictNode.vmProtoData.PendingTop,
		conflictNode.disputable.AssertionClaim.AfterPendingTop,
		conflictNode.disputable.AssertionClaim.ImportedMessagesSlice,
		100,
	)
	staker.Unlock()
}

func (staker *StakingKey) defendExecution(contractAddress common.Address, mach machine.Machine, pre *valprotocol.Precondition, numSteps uint32) {
	staker.Lock()
	challenges.DefendExecutionClaim(
		staker.client,
		contractAddress,
		pre,
		mach,
		numSteps,
		50,
	)
	staker.Unlock()
}
