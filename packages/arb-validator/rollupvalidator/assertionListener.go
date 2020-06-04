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

package rollupvalidator

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type FinalizedAssertion struct {
	Assertion     *protocol.ExecutionAssertion // Disputable assertion
	OnChainTxHash common.Hash                  // Disputable assertion on-chain Tx hash
	NodeHash      common.Hash
}

type AssertionListener struct {
	AdvancedNodeChan chan *structures.Node
}

func NewAssertionListener(advancedNodeChan chan *structures.Node) *AssertionListener {
	return &AssertionListener{
		AdvancedNodeChan: advancedNodeChan,
	}
}

func (al *AssertionListener) RestartingFromLatestValid(context.Context, *rollup.ChainObserver) {
}

func (al *AssertionListener) StakeCreated(context.Context, *rollup.ChainObserver, arbbridge.StakeCreatedEvent) {
}
func (al *AssertionListener) StakeRemoved(context.Context, *rollup.ChainObserver, arbbridge.StakeRefundedEvent) {
}
func (al *AssertionListener) StakeMoved(context.Context, *rollup.ChainObserver, arbbridge.StakeMovedEvent) {
}
func (al *AssertionListener) StartedChallenge(context.Context, *rollup.ChainObserver, *rollup.Challenge) {
}
func (al *AssertionListener) ResumedChallenge(context.Context, *rollup.ChainObserver, *rollup.Challenge) {

}
func (al *AssertionListener) CompletedChallenge(context.Context, *rollup.ChainObserver, arbbridge.ChallengeCompletedEvent) {
}
func (al *AssertionListener) SawAssertion(context.Context, *rollup.ChainObserver, arbbridge.AssertedEvent) {
}
func (al *AssertionListener) ConfirmedNode(context.Context, *rollup.ChainObserver, arbbridge.ConfirmedEvent) {
}
func (al *AssertionListener) PrunedLeaf(context.Context, *rollup.ChainObserver, arbbridge.PrunedEvent) {
}
func (al *AssertionListener) MessageDelivered(context.Context, *rollup.ChainObserver, arbbridge.MessageDeliveredEvent) {
}

func (al *AssertionListener) AssertionPrepared(context.Context, *rollup.ChainObserver, *rollup.PreparedAssertion) {
}
func (al *AssertionListener) ConfirmableNodes(context.Context, *rollup.ChainObserver, *valprotocol.ConfirmOpportunity) {
}
func (al *AssertionListener) PrunableLeafs(context.Context, *rollup.ChainObserver, []valprotocol.PruneParams) {
}
func (al *AssertionListener) MootableStakes(context.Context, *rollup.ChainObserver, []rollup.RecoverStakeMootedParams) {
}
func (al *AssertionListener) OldStakes(context.Context, *rollup.ChainObserver, []rollup.RecoverStakeOldParams) {
}

func (al *AssertionListener) AdvancedCalculatedValidNode(context.Context, *rollup.ChainObserver, common.Hash) {
}
func (al *AssertionListener) AdvancedKnownNode(ctx context.Context, chain *rollup.ChainObserver, node *structures.Node) {
	al.AdvancedNodeChan <- node
}
