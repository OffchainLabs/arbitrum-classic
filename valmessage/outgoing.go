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

package valmessage

import (
	"math/big"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
)

type Signature struct {
	R [32]byte
	S [32]byte
	V uint8
}

type VMMessageRequest struct {
	Data        value.Value
	TokenType   [21]byte
	Currency    *big.Int
	Destination [32]byte
	SequenceNum *big.Int
}

type OutgoingMessage interface {
	IsOutgoingMessage()
}

type SendAssertMessage struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.Assertion
}

func (SendAssertMessage) IsOutgoingMessage() {}

type SendUnanimousAssertMessage struct {
	NewInboxHash [32]byte
	TimeBounds   protocol.TimeBounds
	Assertion    *protocol.Assertion
	Signatures   []Signature
}

func (SendUnanimousAssertMessage) IsOutgoingMessage() {}

type SendProposeUnanimousAssertMessage struct {
	NewInboxHash [32]byte
	TimeBounds   protocol.TimeBounds
	Assertion    *protocol.Assertion
	SequenceNum  uint64
	Signatures   []Signature
}

func (SendProposeUnanimousAssertMessage) IsOutgoingMessage() {}

type SendConfirmUnanimousAssertedMessage struct {
	NewInboxHash [32]byte
	Assertion    *protocol.Assertion
}

func (SendConfirmUnanimousAssertedMessage) IsOutgoingMessage() {}

type SendInitiateChallengeMessage struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.AssertionStub
}

func (SendInitiateChallengeMessage) IsOutgoingMessage() {}

type SendBisectionMessage struct {
	Deadline     uint64
	Precondition *protocol.Precondition
	Assertions   []*protocol.Assertion
}

func (SendBisectionMessage) IsOutgoingMessage() {}

type SendContinueChallengeMessage struct {
	AssertionToChallenge uint16
	Deadline             uint64
	Preconditions        []*protocol.Precondition
	Assertions           []*protocol.AssertionStub
}

func (SendContinueChallengeMessage) IsOutgoingMessage() {}

type SendOneStepProofMessage struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.Assertion
	Proof        []byte
	Deadline     uint64
}

func (SendOneStepProofMessage) IsOutgoingMessage() {}

type SendConfirmedAssertMessage struct {
	Precondition *protocol.Precondition
	Assertion    *protocol.Assertion
}

func (SendConfirmedAssertMessage) IsOutgoingMessage() {}

type SendAsserterTimedOutChallengeMessage struct {
	Deadline     uint64
	Precondition *protocol.Precondition
	Assertion    *protocol.AssertionStub
}

func (SendAsserterTimedOutChallengeMessage) IsOutgoingMessage() {}

type SendChallengerTimedOutChallengeMessage struct {
	Deadline      uint64
	Preconditions []*protocol.Precondition
	Assertions    []*protocol.AssertionStub
}

func (SendChallengerTimedOutChallengeMessage) IsOutgoingMessage() {}

type FinalizedAssertion struct {
	Assertion   *protocol.Assertion
	NewLogCount int
}

func (FinalizedAssertion) IsOutgoingMessage() {}

func (f FinalizedAssertion) NewLogs() []value.Value {
	return f.Assertion.Logs[len(f.Assertion.Logs)-f.NewLogCount:]
}