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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	//"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
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

func (al *AnnouncerListener) SawAssertion(arbbridge.AssertedEvent, *protocol.TimeBlocks, [32]byte) {
	log.Println("SawAssertion")
}

func (al *AnnouncerListener) ConfirmedNode(arbbridge.ConfirmedEvent) {
	log.Println("ConfirmedNode")
}

func (al *AnnouncerListener) AssertionPrepared(*preparedAssertion) {
	log.Println("AssertionPrepared")
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

func (lis *AnnouncerListener) AdvancedKnownAssertion(*protocol.ExecutionAssertion, [32]byte) {
	log.Println("AdvancedKnownAssertion")
}
