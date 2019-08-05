/*
 * Copyright 2019, Offchain Labs, Inc.
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

package bridge

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type Bridge interface {
	FinalizedAssertion(
		assertion *protocol.Assertion,
		newLogCount int,
		signatures [][]byte,
		proposalResults *valmessage.UnanimousUpdateResults,
		onChainTxHash []byte,
	)
	FinalizedUnanimousAssert(
		newInboxHash [32]byte,
		timeBounds protocol.TimeBounds,
		assertion *protocol.Assertion,
		signatures [][]byte,
	)
	PendingUnanimousAssert(
		newInboxHash [32]byte,
		timeBounds protocol.TimeBounds,
		assertion *protocol.Assertion,
		sequenceNum uint64,
		signatures [][]byte,
	)
	ConfirmUnanimousAsserted(
		newInboxHash [32]byte,
		assertion *protocol.Assertion,
	)
	PendingDisputableAssert(
		precondition *protocol.Precondition,
		assertion *protocol.Assertion,
	)
	ConfirmDisputableAsserted(
		precondition *protocol.Precondition,
		assertion *protocol.Assertion,
	)
	InitiateChallenge(
		precondition *protocol.Precondition,
		assertion *protocol.AssertionStub,
	)
	BisectAssertion(
		precondition *protocol.Precondition,
		assertions []*protocol.AssertionStub,
		deadline uint64,
	)
	ContinueChallenge(
		assertionToChallenge uint16,
		preconditions []*protocol.Precondition,
		assertions []*protocol.AssertionStub,
		deadline uint64,
	)
	OneStepProof(
		precondition *protocol.Precondition,
		assertion *protocol.AssertionStub,
		proof []byte,
		deadline uint64,
	)
	AsserterTimedOut(
		precondition *protocol.Precondition,
		assertion *protocol.AssertionStub,
		deadline uint64,
	)
	ChallengerTimedOut(
		preconditions []*protocol.Precondition,
		assertions []*protocol.AssertionStub,
		deadline uint64,
	)
}
