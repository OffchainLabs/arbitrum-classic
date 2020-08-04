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

const (
	PruneSizeLimit = 120
)

type ChainListener interface {
	// This function is called when a ChainListener is added to a ChainObserver
	AddedToChain(context.Context, []*structures.Node)

	// This function is called every time ChainObserver starts running. This
	// includes both the initial run, and after a reorg. The third parameter
	// is the current calculated valid node
	RestartingFromLatestValid(context.Context, *structures.Node)
	StakeCreated(context.Context, *nodegraph.StakedNodeGraph, arbbridge.StakeCreatedEvent)
	StakeRemoved(context.Context, arbbridge.StakeRefundedEvent)
	StakeMoved(context.Context, *nodegraph.StakedNodeGraph, arbbridge.StakeMovedEvent)
	StartedChallenge(
		context.Context,
		*structures.MessageStack,
		*nodegraph.Challenge)
	ResumedChallenge(
		context.Context,
		*structures.MessageStack,
		*nodegraph.Challenge)
	CompletedChallenge(context.Context, *nodegraph.StakedNodeGraph, arbbridge.ChallengeCompletedEvent)
	SawAssertion(context.Context, arbbridge.AssertedEvent)
	ConfirmedNode(context.Context, arbbridge.ConfirmedEvent)
	PrunedLeaf(context.Context, arbbridge.PrunedEvent)
	MessageDelivered(context.Context, arbbridge.MessageDeliveredEvent)
	AssertionPrepared(
		context.Context,
		valprotocol.ChainParams,
		*nodegraph.StakedNodeGraph,
		*structures.Node,
		*common.BlockId,
		*PreparedAssertion)
	ConfirmableNodes(context.Context, *valprotocol.ConfirmOpportunity)
	PrunableLeafs(context.Context, []valprotocol.PruneParams)
	MootableStakes(context.Context, []nodegraph.RecoverStakeMootedParams)
	OldStakes(context.Context, []nodegraph.RecoverStakeOldParams)

	AdvancedKnownNode(context.Context, *nodegraph.StakedNodeGraph, *structures.Node)
}
