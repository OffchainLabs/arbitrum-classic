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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type FinalizedAssertion struct {
	Assertion     *protocol.ExecutionAssertion // Disputable assertion
	OnChainTxHash [32]byte                     // Disputable assertion on-chain Tx hash
}

type AssertionListener struct {
	CompletedAssertionChan chan FinalizedAssertion
}

func (al *AssertionListener) StakeCreated(ethbridge.StakeCreatedEvent)                       {}
func (al *AssertionListener) StakeRemoved(ethbridge.StakeRefundedEvent)                      {}
func (al *AssertionListener) StakeMoved(ethbridge.StakeMovedEvent)                           {}
func (al *AssertionListener) StartedChallenge(ethbridge.ChallengeStartedEvent, *Node, *Node) {}
func (al *AssertionListener) CompletedChallenge(event ethbridge.ChallengeCompletedEvent)     {}

func (al *AssertionListener) AssertionPrepared(*preparedAssertion)      {}
func (al *AssertionListener) PrunableLeafs([]pruneParams)               {}
func (al *AssertionListener) MootableStakes([]recoverStakeMootedParams) {}
func (al *AssertionListener) OldStakes([]recoverStakeOldParams)         {}

func (lis *AssertionListener) AdvancedKnownAssertion(assertion *protocol.ExecutionAssertion, txHash [32]byte) {
	lis.CompletedAssertionChan <- FinalizedAssertion{
		Assertion:     assertion,
		OnChainTxHash: txHash,
	}
}
