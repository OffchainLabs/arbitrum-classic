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

package challenges

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

func ChallengeExecutionClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	startPrecondition *valprotocol.Precondition,
	startMachine machine.Machine,
	challengeEverything bool,
) (ChallengeState, error) {
	contractWatcher, err := client.NewExecutionChallengeWatcher(address)
	if err != nil {
		return 0, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockId, startLogIndex, contractWatcher)

	contract, err := client.NewExecutionChallenge(address)
	if err != nil {
		return 0, err
	}

	return challengeExecution(
		reorgCtx,
		eventChan,
		contract,
		client,
		startMachine,
		startPrecondition,
		challengeEverything,
	)
}

func challengeExecution(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.ExecutionChallenge,
	client arbbridge.ArbClient,
	startMachine machine.Machine,
	startPrecondition *valprotocol.Precondition,
	challengeEverything bool,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("ExecutionChallenge challenger expected InitiateChallengeEvent but got %T", event)
	}

	mach := startMachine
	precondition := startPrecondition
	deadline := ev.Deadline
	for {
		// get defender update
		event, state, err := getNextEventWithTimeout(
			ctx,
			eventChan,
			deadline,
			contract,
			client,
		)

		if challengeEnded(state, err) {
			return state, err
		}

		if _, ok := event.(arbbridge.OneStepProofEvent); ok {
			return ChallengeAsserterWon, nil
		}

		bisectionEvent, ok := event.(arbbridge.ExecutionBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge challenger expected ExecutionBisectionEvent but got %T", event)
		}

		newMachine, preconditions, event, state, chooseSegment, err := executionChallengerUpdate(
			ctx,
			eventChan,
			mach,
			contract,
			precondition,
			bisectionEvent,
			challengeEverything)

		if challengeEnded(state, err) {
			return state, err
		}

		continueEvent, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge challenger expected ContinueChallengeEvent but got %T", event)
		}

		// Update mach, precondition, deadline
		if chooseSegment {
			// Freshly bisected assertion
			mach = newMachine
			precondition = preconditions[continueEvent.SegmentIndex.Uint64()]
		} else {
			// Replayed from existing event
			precondition = setPreCondition(
				continueEvent,
				bisectionEvent,
				mach,
				startPrecondition,
				precondition)
		}
		deadline = continueEvent.Deadline
	}
}

func setPreCondition(
	continueEvent arbbridge.ContinueChallengeEvent,
	bisectionEvent arbbridge.ExecutionBisectionEvent,
	mach machine.Machine,
	startPrecondition *valprotocol.Precondition,
	precondition *valprotocol.Precondition,
) *valprotocol.Precondition {
	totalSteps := uint64(0)
	for i := uint64(0); i < continueEvent.SegmentIndex.Uint64(); i++ {
		totalSteps += valprotocol.CalculateBisectionStepCount(
			i,
			uint64(len(bisectionEvent.Assertions)),
			bisectionEvent.TotalSteps)
	}
	assertion, _ := mach.ExecuteAssertion(
		totalSteps,
		startPrecondition.TimeBounds,
		startPrecondition.BeforeInbox,
		0,
	)
	assertStub := valprotocol.NewExecutionAssertionStubFromAssertion(assertion)
	precondition = precondition.GeneratePostcondition(assertStub)

	return precondition
}

func executionChallengerUpdate(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	mach machine.Machine,
	contract arbbridge.ExecutionChallenge,
	precondition *valprotocol.Precondition,
	bisectionEvent arbbridge.ExecutionBisectionEvent,
	challengeEverything bool,
) (machine.Machine, []*valprotocol.Precondition, arbbridge.Event, ChallengeState, bool, error) {
	chooseSegment, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	var preconditions []*valprotocol.Precondition
	var newMachine machine.Machine
	if chooseSegment {
		var challengedAssertionNum uint16
		challengedAssertionNum, newMachine, err = ChooseAssertionToChallenge(
			mach.Clone(),
			precondition,
			bisectionEvent.Assertions,
			bisectionEvent.TotalSteps)

		if err != nil {
			if !challengeEverything {
				return nil, nil, nil, 0, chooseSegment, err
			} else {
				pre := precondition
				cMach := mach.Clone()
				challengedAssertionNum = uint16(rand.Int31n(int32(len(bisectionEvent.Assertions))))

				for i := 0; i < len(bisectionEvent.Assertions); i++ {
					stepCount := valprotocol.CalculateBisectionStepCount(
						uint64(i),
						uint64(len(bisectionEvent.Assertions)),
						bisectionEvent.TotalSteps)

					newMachine = cMach.Clone()
					assertion, _ := cMach.ExecuteAssertion(
						stepCount,
						pre.TimeBounds,
						pre.BeforeInbox,
						0,
					)
					newAssertStub := valprotocol.NewExecutionAssertionStubFromAssertion(assertion)
					pre = precondition.GeneratePostcondition(newAssertStub)
				}
				err = nil
			}
		}

		preconditions = valprotocol.GeneratePreconditions(precondition, bisectionEvent.Assertions)
		err = contract.ChooseSegment(
			ctx,
			challengedAssertionNum,
			preconditions,
			bisectionEvent.Assertions,
			bisectionEvent.TotalSteps,
		)
		event, state, err = getNextEvent(eventChan)
	}

	return newMachine, preconditions, event, state, chooseSegment, err
}
