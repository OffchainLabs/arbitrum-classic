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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type FinalizedAssertion struct {
	Assertion     *protocol.ExecutionAssertion // Disputable assertion
	OnChainTxHash [32]byte                     // Disputable assertion on-chain Tx hash
}

type AssertionListener struct {
	CompletedAssertionChan chan FinalizedAssertion
}

func (al *AssertionListener) StakeCreated(arbbridge.StakeCreatedEvent)                             {}
func (al *AssertionListener) StakeRemoved(arbbridge.StakeRefundedEvent)                            {}
func (al *AssertionListener) StakeMoved(arbbridge.StakeMovedEvent)                                 {}
func (al *AssertionListener) StartedChallenge(arbbridge.ChallengeStartedEvent, *Node, *Node)       {}
func (al *AssertionListener) CompletedChallenge(event arbbridge.ChallengeCompletedEvent)           {}
func (al *AssertionListener) SawAssertion(arbbridge.AssertedEvent, *protocol.TimeBlocks, [32]byte) {}
func (al *AssertionListener) ConfirmedNode(arbbridge.ConfirmedEvent)                               {}
func (al *AssertionListener) PrunedLeaf(arbbridge.PrunedEvent)                                     {}

func (al *AssertionListener) AssertionPrepared(*preparedAssertion)              {}
func (al *AssertionListener) ValidNodeConfirmable(*confirmValidOpportunity)     {}
func (al *AssertionListener) InvalidNodeConfirmable(*confirmInvalidOpportunity) {}
func (al *AssertionListener) PrunableLeafs([]pruneParams)                       {}
func (al *AssertionListener) MootableStakes([]recoverStakeMootedParams)         {}
func (al *AssertionListener) OldStakes([]recoverStakeOldParams)                 {}

func (al *AssertionListener) AdvancedKnownValidNode([32]byte) {}
func (al *AssertionListener) AdvancedKnownAssertion(assertion *protocol.ExecutionAssertion, txHash [32]byte) {
	al.CompletedAssertionChan <- FinalizedAssertion{
		Assertion:     assertion,
		OnChainTxHash: txHash,
	}
}
