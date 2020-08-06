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

package chainlistener

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/nodegraph"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type NoopListener struct{}

func (nl *NoopListener) AddedToChain(context.Context, []*structures.Node) {
}

func (nl *NoopListener) RestartingFromLatestValid(context.Context, *structures.Node) {
}

func (NoopListener) StakeCreated(context.Context, *nodegraph.StakedNodeGraph, arbbridge.StakeCreatedEvent) {
}

func (NoopListener) StakeRemoved(context.Context, arbbridge.StakeRefundedEvent) {
}
func (NoopListener) StakeMoved(context.Context, *nodegraph.StakedNodeGraph, arbbridge.StakeMovedEvent) {
}

func (NoopListener) StartedChallenge(
	context.Context,
	*structures.MessageStack,
	*nodegraph.Challenge) {
}
func (NoopListener) ResumedChallenge(
	context.Context,
	*structures.MessageStack,
	*nodegraph.Challenge) {

}
func (NoopListener) CompletedChallenge(context.Context, *nodegraph.StakedNodeGraph, arbbridge.ChallengeCompletedEvent) {
}
func (NoopListener) SawAssertion(context.Context, arbbridge.AssertedEvent) {
}
func (NoopListener) ConfirmedNode(context.Context, arbbridge.ConfirmedEvent) {
}
func (NoopListener) PrunedLeaf(context.Context, arbbridge.PrunedEvent) {
}
func (NoopListener) MessageDelivered(context.Context, arbbridge.MessageDeliveredEvent) {
}

func (NoopListener) AssertionPrepared(
	context.Context,
	valprotocol.ChainParams,
	*nodegraph.StakedNodeGraph,
	*structures.Node,
	*common.BlockId,
	*PreparedAssertion) {
}
func (NoopListener) ConfirmableNodes(context.Context, *valprotocol.ConfirmOpportunity) {
}
func (NoopListener) PrunableLeafs(context.Context, []valprotocol.PruneParams) {
}
func (NoopListener) MootableStakes(context.Context, []nodegraph.RecoverStakeMootedParams) {
}
func (NoopListener) OldStakes(context.Context, []nodegraph.RecoverStakeOldParams) {
}

func (NoopListener) AdvancedKnownNode(context.Context, *nodegraph.StakedNodeGraph, *structures.Node) {
}
