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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type FinalizedAssertion struct {
	Assertion     *protocol.ExecutionAssertion // Disputable assertion
	OnChainTxHash common.Hash                  // Disputable assertion on-chain Tx hash
}

type AssertionListener struct {
	CompletedAssertionChan chan FinalizedAssertion
}

func (al *AssertionListener) StakeCreated(context.Context, *ChainObserver, arbbridge.StakeCreatedEvent) {
}
func (al *AssertionListener) StakeRemoved(context.Context, *ChainObserver, arbbridge.StakeRefundedEvent) {
}
func (al *AssertionListener) StakeMoved(context.Context, *ChainObserver, arbbridge.StakeMovedEvent) {}
func (al *AssertionListener) StartedChallenge(context.Context, *ChainObserver, arbbridge.ChallengeStartedEvent, *Node, *Node) {
}
func (al *AssertionListener) CompletedChallenge(context.Context, *ChainObserver, arbbridge.ChallengeCompletedEvent) {
}
func (al *AssertionListener) SawAssertion(context.Context, *ChainObserver, arbbridge.AssertedEvent) {
}
func (al *AssertionListener) ConfirmedNode(context.Context, *ChainObserver, arbbridge.ConfirmedEvent) {
}
func (al *AssertionListener) PrunedLeaf(context.Context, *ChainObserver, arbbridge.PrunedEvent) {}
func (al *AssertionListener) MessageDelivered(context.Context, *ChainObserver, arbbridge.MessageDeliveredEvent) {
}

func (al *AssertionListener) AssertionPrepared(context.Context, *ChainObserver, *preparedAssertion) {}
func (al *AssertionListener) NodesConfirmable(context.Context, *ChainObserver, *structures.ConfirmOpportunity) {
}
func (al *AssertionListener) PrunableLeafs(context.Context, *ChainObserver, []pruneParams) {}
func (al *AssertionListener) MootableStakes(context.Context, *ChainObserver, []recoverStakeMootedParams) {
}
func (al *AssertionListener) OldStakes(context.Context, *ChainObserver, []recoverStakeOldParams) {}

func (al *AssertionListener) AdvancedCalculatedValidNode(context.Context, *ChainObserver, common.Hash) {
}
func (al *AssertionListener) AdvancedKnownAssertion(ctx context.Context, chain *ChainObserver, assertion *protocol.ExecutionAssertion, txHash common.Hash) {
	al.CompletedAssertionChan <- FinalizedAssertion{
		Assertion:     assertion,
		OnChainTxHash: txHash,
	}
}
