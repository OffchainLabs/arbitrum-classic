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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/nodegraph"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type AnnouncerListener struct {
	Prefix string
}

func (al *AnnouncerListener) AddedToChain(context.Context, []*structures.Node) {
	log.Println("AddedToChain")
}

func (al *AnnouncerListener) RestartingFromLatestValid(context.Context, *structures.Node) {
	log.Println("RestartingFromLatestValid")
}

func (al *AnnouncerListener) StakeCreated(ctx context.Context, ng *nodegraph.StakedNodeGraph, ev arbbridge.StakeCreatedEvent) {
	log.Printf("%v Staker %v created at %v\n", al.Prefix, ev.Staker, ev.NodeHash)
}

func (al *AnnouncerListener) StakeRemoved(ctx context.Context, ev arbbridge.StakeRefundedEvent) {
	log.Printf("%v Staker %v removed\n", al.Prefix, ev.Staker)
}

func (al *AnnouncerListener) StakeMoved(ctx context.Context, ng *nodegraph.StakedNodeGraph, ev arbbridge.StakeMovedEvent) {
	log.Printf("%v Staker %v moved to location: %v\n", al.Prefix, ev.Staker, ev.Location)
}

func (al *AnnouncerListener) StartedChallenge(
	context.Context,
	*structures.MessageStack,
	*nodegraph.Challenge) {
	log.Println(al.Prefix, "StartedChallenge")
}

func (al *AnnouncerListener) ResumedChallenge(
	context.Context,
	*structures.MessageStack,
	*nodegraph.Challenge) {
	log.Println(al.Prefix, "ResumedChallenge")
}

func (al *AnnouncerListener) CompletedChallenge(
	ctx context.Context,
	ng *nodegraph.StakedNodeGraph,
	event arbbridge.ChallengeCompletedEvent,
) {
	log.Println(al.Prefix, "CompletedChallenge")
}

func (al *AnnouncerListener) SawAssertion(ctx context.Context, ev arbbridge.AssertedEvent) {
	log.Println(al.Prefix, "SawAssertion on leaf", ev.PrevLeafHash)
	log.Println(al.Prefix, "Params:", ev.Disputable.AssertionParams)
	log.Println(al.Prefix, "Assertion:", ev.Disputable.Assertion)
}

func (al *AnnouncerListener) ConfirmedNode(ctx context.Context, ev arbbridge.ConfirmedEvent) {
	log.Println(al.Prefix, "ConfirmedNode", ev.NodeHash)
}

func (al *AnnouncerListener) PrunedLeaf(ctx context.Context, ev arbbridge.PrunedEvent) {
	log.Println(al.Prefix, "PrunedLeaf", ev.Leaf)
}

func (al *AnnouncerListener) MessageDelivered(_ context.Context, ev arbbridge.MessageDeliveredEvent) {
	log.Println(al.Prefix, "MessageDelivered", ev.Message)
}

func (al *AnnouncerListener) AssertionPrepared(
	context.Context,
	valprotocol.ChainParams,
	*nodegraph.StakedNodeGraph,
	*structures.Node,
	*common.BlockId,
	*PreparedAssertion,
) {
	log.Println(al.Prefix, "AssertionPrepared")
}
func (al *AnnouncerListener) ConfirmableNodes(context.Context, *valprotocol.ConfirmOpportunity) {
	log.Println(al.Prefix, "ConfirmableNodes")
}
func (al *AnnouncerListener) PrunableLeafs(context.Context, []valprotocol.PruneParams) {
	log.Println(al.Prefix, "PrunableLeafs")
}
func (al *AnnouncerListener) MootableStakes(context.Context, []nodegraph.RecoverStakeMootedParams) {
	log.Println(al.Prefix, "MootableStakes")
}
func (al *AnnouncerListener) OldStakes(context.Context, []nodegraph.RecoverStakeOldParams) {
	log.Println(al.Prefix, "OldStakes")
}

func (al *AnnouncerListener) AdvancedKnownNode(context.Context, *nodegraph.StakedNodeGraph, *structures.Node) {
	log.Println(al.Prefix, "AdvancedKnownNode")
}
