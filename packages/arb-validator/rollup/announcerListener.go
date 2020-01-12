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
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type AnnouncerListener struct{}

func (al *AnnouncerListener) StakeCreated(arbbridge.StakeCreatedEvent) {
	log.Println("StakeCreated")
}
func (al *AnnouncerListener) StakeRemoved(arbbridge.StakeRefundedEvent) {
	log.Println("StakeRemoved")
}
func (al *AnnouncerListener) StakeMoved(ev arbbridge.StakeMovedEvent) {
	log.Printf("StakeMoved(staker: %v, location: %v)\n", hexutil.Encode(ev.Staker[:]), hexutil.Encode(ev.Location[:]))
}
func (al *AnnouncerListener) StartedChallenge(arbbridge.ChallengeStartedEvent, *Node, *Node) {
	log.Println("StartedChallenge")
}
func (al *AnnouncerListener) CompletedChallenge(event arbbridge.ChallengeCompletedEvent) {
	log.Println("CompletedChallenge")
}

func (al *AnnouncerListener) SawAssertion(ev arbbridge.AssertedEvent, time *protocol.TimeBlocks, txHash [32]byte) {
	log.Println("SawAssertion")
	log.Println("Params:", ev.Params)
	log.Println("Claim:", ev.Claim)
}

func (al *AnnouncerListener) ConfirmedNode(ev arbbridge.ConfirmedEvent) {
	log.Println("ConfirmedNode", hexutil.Encode(ev.NodeHash[:]))
}

func (al *AnnouncerListener) PrunedLeaf(ev arbbridge.PrunedEvent) {
	log.Println("PrunedLeaf", hexutil.Encode(ev.Leaf[:]))
}

func (al *AnnouncerListener) AssertionPrepared(*preparedAssertion) {
	log.Println("AssertionPrepared")
}
func (al *AnnouncerListener) ValidNodeConfirmable(*confirmValidOpportunity) {
	log.Println("ValidNodeConfirmable")
}
func (al *AnnouncerListener) InvalidNodeConfirmable(*confirmInvalidOpportunity) {
	log.Println("InvalidNodeConfirmable")
}
func (al *AnnouncerListener) PrunableLeafs([]pruneParams) {
	log.Println("PrunableLeafs")
}
func (al *AnnouncerListener) MootableStakes([]recoverStakeMootedParams) {
	log.Println("MootableStakes")
}
func (al *AnnouncerListener) OldStakes([]recoverStakeOldParams) {
	log.Println("OldStakes")
}

func (al *AnnouncerListener) AdvancedKnownValidNode(nodeHash [32]byte) {
	log.Println("AdvancedKnownValidNode", hexutil.Encode(nodeHash[:]))
}

func (lis *AnnouncerListener) AdvancedKnownAssertion(*protocol.ExecutionAssertion, [32]byte) {
	log.Println("AdvancedKnownAssertion")
}
