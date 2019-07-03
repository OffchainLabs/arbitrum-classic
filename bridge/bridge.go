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
	"github.com/offchainlabs/arb-util/protocol"
)

type Bridge interface {
	FinalizedAssertion(assertion *protocol.Assertion, newLogCount int)
	FinalUnanimousAssert(newInboxHash [32]byte, timeBounds protocol.TimeBounds, assertion *protocol.Assertion, signatures [][]byte)
	UnanimousAssert(newInboxHash [32]byte, timeBounds protocol.TimeBounds, assertion *protocol.Assertion, sequenceNum uint64, signatures [][]byte)
	ConfirmUnanimousAssertion(newInboxHash [32]byte, assertion *protocol.Assertion)
	DisputableAssert(precondition *protocol.Precondition, assertion *protocol.Assertion)
	ConfirmDisputableAssertion(precondition *protocol.Precondition, assertion *protocol.Assertion)
	InitiateChallenge(precondition *protocol.Precondition, assertion *protocol.AssertionStub)
	BisectAssertion(precondition *protocol.Precondition, assertions []*protocol.Assertion, deadline uint64)
	ContinueChallenge(assertionToChallenge uint16, preconditions []*protocol.Precondition, assertions []*protocol.AssertionStub, deadline uint64)
	OneStepProof(precondition *protocol.Precondition, assertion *protocol.Assertion, proof []byte, deadline uint64)
	TimeoutAsserter(precondition *protocol.Precondition, assertion *protocol.AssertionStub, deadline uint64)
	TimeoutChallenger(preconditions []*protocol.Precondition, assertions []*protocol.AssertionStub, deadline uint64)
}
