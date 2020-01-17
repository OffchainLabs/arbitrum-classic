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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type FinalizedAssertion struct {
	Assertion     *protocol.ExecutionAssertion // Disputable assertion
	OnChainTxHash common.Hash                  // Disputable assertion on-chain Tx hash
}

type AssertionListener struct {
	CompletedAssertionChan chan FinalizedAssertion
}

func (al *AssertionListener) StakeCreated(*ChainObserver, arbbridge.StakeCreatedEvent)  {}
func (al *AssertionListener) StakeRemoved(*ChainObserver, arbbridge.StakeRefundedEvent) {}
func (al *AssertionListener) StakeMoved(*ChainObserver, arbbridge.StakeMovedEvent)      {}
func (al *AssertionListener) StartedChallenge(*ChainObserver, arbbridge.ChallengeStartedEvent, *Node, *Node) {
}
func (al *AssertionListener) CompletedChallenge(observer *ChainObserver, event arbbridge.ChallengeCompletedEvent) {
}
func (al *AssertionListener) SawAssertion(*ChainObserver, arbbridge.AssertedEvent, *common.TimeBlocks, common.Hash) {
}
func (al *AssertionListener) ConfirmedNode(*ChainObserver, arbbridge.ConfirmedEvent)           {}
func (al *AssertionListener) PrunedLeaf(*ChainObserver, arbbridge.PrunedEvent)                 {}
func (al *AssertionListener) MessageDelivered(*ChainObserver, arbbridge.MessageDeliveredEvent) {}

func (al *AssertionListener) AssertionPrepared(*ChainObserver, *preparedAssertion)              {}
func (al *AssertionListener) ValidNodeConfirmable(*ChainObserver, *confirmValidOpportunity)     {}
func (al *AssertionListener) InvalidNodeConfirmable(*ChainObserver, *confirmInvalidOpportunity) {}
func (al *AssertionListener) PrunableLeafs(*ChainObserver, []pruneParams)                       {}
func (al *AssertionListener) MootableStakes(*ChainObserver, []recoverStakeMootedParams)         {}
func (al *AssertionListener) OldStakes(*ChainObserver, []recoverStakeOldParams)                 {}

func (al *AssertionListener) AdvancedKnownValidNode(*ChainObserver, common.Hash) {}
func (al *AssertionListener) AdvancedKnownAssertion(chain *ChainObserver, assertion *protocol.ExecutionAssertion, txHash common.Hash) {
	al.CompletedAssertionChan <- FinalizedAssertion{
		Assertion:     assertion,
		OnChainTxHash: txHash,
	}
}
