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

package challenges

import (
	"errors"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

type AssertionDefender struct {
	numSteps  uint64
	initState machine.Machine
	inbox     *structures.MessageStack
	assertion *valprotocol.ExecutionAssertionStub
}

func NewAssertionDefender(numSteps uint64, initState machine.Machine, inbox *structures.MessageStack, assertion *valprotocol.ExecutionAssertionStub) AssertionDefender {
	return AssertionDefender{
		numSteps:  numSteps,
		initState: initState.Clone(),
		inbox:     inbox,
		assertion: assertion,
	}
}

func (ad AssertionDefender) NumSteps() uint64 {
	return ad.numSteps
}

func (ad AssertionDefender) AssertionStub() *valprotocol.ExecutionAssertionStub {
	return ad.assertion
}

func (ad AssertionDefender) MoveDefender(bisectionEvent arbbridge.ExecutionBisectionEvent, continueEvent arbbridge.ContinueChallengeEvent) (*AssertionDefender, error) {
	segmentCount := uint64(len(bisectionEvent.AssertionHashes))
	stepsToSkip := computeStepsUpTo(continueEvent.SegmentIndex.Uint64(), segmentCount, ad.numSteps)
	steps := valprotocol.CalculateBisectionStepCount(
		continueEvent.SegmentIndex.Uint64(),
		segmentCount,
		ad.numSteps,
	)

	// Update mach, precondition, deadline
	messages, err := ad.inbox.GetAssertionMessages(ad.assertion.BeforeInboxHash, ad.assertion.AfterInboxHash)
	if err != nil {
		return nil, fmt.Errorf("assertion defender must have valid messages: %s %s, %s", ad.assertion.BeforeInboxHash, ad.assertion.AfterInboxHash, err)
	}

	// Last value returned is the number of steps executed
	skippedAssertion, _ := ad.initState.ExecuteAssertion(
		stepsToSkip,
		messages,
		0,
	)
	skippedAssertionStub := structures.NewExecutionAssertionStubFromAssertion(
		skippedAssertion,
		ad.assertion.BeforeInboxHash,
		ad.assertion.FirstLogHash,
		ad.assertion.FirstMessageHash,
		ad.inbox,
	)

	// Last value returned is the number of steps executed
	assertion, _ := ad.initState.Clone().ExecuteAssertion(steps, messages[skippedAssertion.InboxMessagesConsumed:], 0)
	assertionStub := structures.NewExecutionAssertionStubFromAssertion(
		assertion,
		skippedAssertionStub.AfterInboxHash,
		skippedAssertionStub.LastLogHash,
		skippedAssertionStub.LastMessageHash,
		ad.inbox,
	)
	assertionDefender := NewAssertionDefender(steps, ad.initState, ad.inbox, assertionStub)
	return &assertionDefender, nil
}

func (ad AssertionDefender) NBisect(slices uint64) []AssertionDefender {
	nsteps := ad.NumSteps()
	if nsteps < slices {
		slices = nsteps
	}
	defenders := make([]AssertionDefender, 0, slices)
	m := ad.initState.Clone()

	beforeInboxHash := ad.assertion.BeforeInboxHash
	firstLogHash := ad.assertion.FirstLogHash
	firstMessageHash := ad.assertion.FirstMessageHash

	for i := uint64(0); i < slices; i++ {
		steps := valprotocol.CalculateBisectionStepCount(i, slices, nsteps)
		initState := m.Clone()

		inboxMessages, err := ad.inbox.GetAssertionMessages(beforeInboxHash, ad.assertion.AfterInboxHash)
		if err != nil {
			log.Fatal("inbox messages must exist for assertion that you're defending ", beforeInboxHash, ad.assertion.AfterInboxHash)
		}

		// Last value returned is the number of steps executed
		assertion, numSteps := m.ExecuteAssertion(
			steps,
			inboxMessages,
			0,
		)
		stub := structures.NewExecutionAssertionStubFromAssertion(assertion, beforeInboxHash, firstLogHash, firstMessageHash, ad.inbox)
		defenders = append(defenders, NewAssertionDefender(
			numSteps,
			initState,
			ad.inbox,
			stub,
		))
		beforeInboxHash = stub.AfterInboxHash
		firstLogHash = stub.LastLogHash
		firstMessageHash = stub.LastMessageHash
	}
	return defenders
}

func (ad AssertionDefender) SolidityOneStepProof() ([]byte, *inbox.InboxMessage, error) {
	proofData, err := ad.initState.MarshalForProof()
	if err != nil {
		return nil, nil, err
	}

	messages, err := ad.inbox.GetAssertionMessages(ad.assertion.BeforeInboxHash, ad.assertion.AfterInboxHash)
	if err != nil {
		return nil, nil, err
	}

	if len(messages) > 1 {
		return nil, nil, errors.New("can't prove assertion with more than one message")
	}
	if len(messages) == 1 {
		return proofData, &messages[0], nil
	}
	return proofData, nil, nil
}
