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
	"github.com/rs/zerolog"
)

type AnnouncerListener struct {
	logger zerolog.Logger
}

func NewAnnouncerListener(address common.Address) *AnnouncerListener {
	return &AnnouncerListener{logger: logger.With().Hex("validator", address.Bytes()).Logger()}
}

func (al *AnnouncerListener) AddedToChain(context.Context, []*structures.Node) {
	logger.Info().Msg("AddedToChain")
}

func (al *AnnouncerListener) RestartingFromLatestValid(context.Context, *structures.Node) {
	logger.Info().Msg("RestartingFromLatestValid")
}

func (al *AnnouncerListener) StakeCreated(ctx context.Context, ng *nodegraph.StakedNodeGraph, ev arbbridge.StakeCreatedEvent) {
	logger.Info().
		Hex("staker", ev.Staker.Bytes()).
		Hex("node", ev.NodeHash.Bytes()).
		Msg("StakeCreated")
}

func (al *AnnouncerListener) StakeRemoved(ctx context.Context, ev arbbridge.StakeRefundedEvent) {
	logger.Info().
		Hex("staker", ev.Staker.Bytes()).
		Msg("StakeRemoved")
}

func (al *AnnouncerListener) StakeMoved(ctx context.Context, ng *nodegraph.StakedNodeGraph, ev arbbridge.StakeMovedEvent) {
	logger.Info().
		Hex("staker", ev.Staker.Bytes()).
		Hex("location", ev.Location.Bytes()).
		Msg("StakeMoved")
}

func (al *AnnouncerListener) StartedChallenge(
	context.Context,
	*structures.MessageStack,
	*nodegraph.Challenge) {
	logger.Info().
		Msg("StartedChallenge")
}

func (al *AnnouncerListener) ResumedChallenge(
	context.Context,
	*structures.MessageStack,
	*nodegraph.Challenge) {
	logger.Info().
		Msg("ResumedChallenge")
}

func (al *AnnouncerListener) CompletedChallenge(
	ctx context.Context,
	ng *nodegraph.StakedNodeGraph,
	event arbbridge.ChallengeCompletedEvent,
) {
	logger.Info().
		Msg("CompletedChallenge")
}

func (al *AnnouncerListener) SawAssertion(ctx context.Context, ev arbbridge.AssertedEvent) {
	logger.Info().
		Hex("leaf", ev.PrevLeafHash.Bytes()).
		Str("leaf", ev.AssertionParams.String()).
		Msg("SawAssertion")
}

func (al *AnnouncerListener) ConfirmedNode(ctx context.Context, ev arbbridge.ConfirmedEvent) {
	logger.Info().
		Hex("node", ev.NodeHash.Bytes()).
		Msg("ConfirmedNode")
}

func (al *AnnouncerListener) PrunedLeaf(ctx context.Context, ev arbbridge.PrunedEvent) {
	logger.Info().
		Hex("leaf", ev.Leaf.Bytes()).
		Msg("PrunedLeaf")
}

func (al *AnnouncerListener) MessageDelivered(_ context.Context, ev arbbridge.MessageDeliveredEvent) {
	/*
		logger.Info().
			Str("message", ev.Message.String()).
			Msg("MessageDelivered")
	*/
}

func (al *AnnouncerListener) AssertionPrepared(
	context.Context,
	valprotocol.ChainParams,
	*nodegraph.StakedNodeGraph,
	*structures.Node,
	*PreparedAssertion,
) {
	logger.Info().
		Msg("AssertionPrepared")
}
func (al *AnnouncerListener) ConfirmableNodes(context.Context, *valprotocol.ConfirmOpportunity) {
	logger.Info().
		Msg("ConfirmableNodes")
}
func (al *AnnouncerListener) PrunableLeafs(context.Context, []valprotocol.PruneParams) {
	logger.Info().
		Msg("PrunableLeafs")
}
func (al *AnnouncerListener) MootableStakes(context.Context, []nodegraph.RecoverStakeMootedParams) {
	logger.Info().
		Msg("MootableStakes")
}
func (al *AnnouncerListener) OldStakes(context.Context, []nodegraph.RecoverStakeOldParams) {
	logger.Info().
		Msg("OldStakes")
}

func (al *AnnouncerListener) AdvancedKnownNode(context.Context, *nodegraph.StakedNodeGraph, *structures.Node) {
	logger.Info().
		Msg("AdvancedKnownNode")
}
