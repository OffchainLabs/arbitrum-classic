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
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type AnnouncerListener struct {
	Prefix string
}

func (al *AnnouncerListener) StakeCreated(ctx context.Context, observer *ChainObserver, ev arbbridge.StakeCreatedEvent) {
	log.Printf("%v Staker %v created at %v\n", al.Prefix, ev.Staker, ev.NodeHash)
}

func (al *AnnouncerListener) StakeRemoved(ctx context.Context, observer *ChainObserver, ev arbbridge.StakeRefundedEvent) {
	log.Printf("%v Staker %v removed\n", al.Prefix, ev.Staker)
}

func (al *AnnouncerListener) StakeMoved(ctx context.Context, observer *ChainObserver, ev arbbridge.StakeMovedEvent) {
	log.Printf("%v Staker %v moved to location: %v\n", al.Prefix, ev.Staker, ev.Location)
}

func (al *AnnouncerListener) StartedChallenge(context.Context, *ChainObserver, arbbridge.ChallengeStartedEvent, *Node, *Node) {
	log.Println(al.Prefix, "StartedChallenge")
}

func (al *AnnouncerListener) CompletedChallenge(ctx context.Context, observer *ChainObserver, event arbbridge.ChallengeCompletedEvent) {
	log.Println(al.Prefix, "CompletedChallenge")
}

func (al *AnnouncerListener) SawAssertion(ctx context.Context, observer *ChainObserver, ev arbbridge.AssertedEvent) {
	log.Println(al.Prefix, "SawAssertion on leaf", ev.PrevLeafHash)
	log.Println(al.Prefix, "Params:", ev.Params)
	log.Println(al.Prefix, "Claim:", ev.Claim)
}

func (al *AnnouncerListener) ConfirmedNode(ctx context.Context, observer *ChainObserver, ev arbbridge.ConfirmedEvent) {
	log.Println(al.Prefix, "ConfirmedNode", ev.NodeHash)
}

func (al *AnnouncerListener) PrunedLeaf(ctx context.Context, observer *ChainObserver, ev arbbridge.PrunedEvent) {
	log.Println(al.Prefix, "PrunedLeaf", ev.Leaf)
}

func (al *AnnouncerListener) MessageDelivered(context.Context, *ChainObserver, arbbridge.MessageDeliveredEvent) {
	log.Println(al.Prefix, "MessageDelivered")
}

func (al *AnnouncerListener) AssertionPrepared(context.Context, *ChainObserver, *preparedAssertion) {
	log.Println(al.Prefix, "AssertionPrepared")
}
func (al *AnnouncerListener) NodesConfirmable(context.Context, *ChainObserver, *valprotocol.ConfirmOpportunity) {
	log.Println(al.Prefix, "NodesConfirmable")
}
func (al *AnnouncerListener) PrunableLeafs(context.Context, *ChainObserver, []pruneParams) {
	log.Println(al.Prefix, "PrunableLeafs")
}
func (al *AnnouncerListener) MootableStakes(context.Context, *ChainObserver, []recoverStakeMootedParams) {
	log.Println(al.Prefix, "MootableStakes")
}
func (al *AnnouncerListener) OldStakes(context.Context, *ChainObserver, []recoverStakeOldParams) {
	log.Println(al.Prefix, "OldStakes")
}

func (al *AnnouncerListener) AdvancedCalculatedValidNode(ctx context.Context, observer *ChainObserver, nodeHash common.Hash) {
	log.Println(al.Prefix, "AdvancedCalculatedValidNode", nodeHash)
}

func (al *AnnouncerListener) AdvancedKnownAssertion(context.Context, *ChainObserver, *protocol.ExecutionAssertion, common.Hash) {
	log.Println(al.Prefix, "AdvancedKnownAssertion")
}
