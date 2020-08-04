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
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
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
	inboxMessages []inbox.InboxMessage,
	startMachine machine.Machine,
	challengeEverything bool,
	challengeType ExecutionChallengeInfo,
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
		inboxMessages,
		challengeEverything,
		challengeType,
	)
}

func challengeExecution(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.ExecutionChallenge,
	client arbbridge.ArbClient,
	startMachine machine.Machine,
	inboxMessages []inbox.InboxMessage,
	challengeEverything bool,
	challengeType ExecutionChallengeInfo,
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
	deadline := ev.Deadline
	for {
		cont := ContinueChallenge(challengeType)

		if cont {
			if challengeType.isDiscontinueType {
				challengeType.currentRound += 1
			}
		} else {
			return ChallengerDiscontinued, nil
		}

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

		chooseSegment, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)

		if chooseSegment {
			mach, inboxMessages, err = executionChallengerUpdate(
				ctx,
				mach,
				contract,
				inboxMessages,
				bisectionEvent,
				challengeEverything,
			)
			if err != nil {
				return state, err
			}
			event, state, err = getNextEvent(eventChan)
		}

		if challengeEnded(state, err) {
			return state, err
		}

		continueEvent, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge challenger expected ContinueChallengeEvent but got %T", event)
		}

		// Update mach, precondition, deadline
		if !chooseSegment {
			// Replayed from existing event
			totalSteps := computeSteps(continueEvent, bisectionEvent)
			assertion, _ := mach.ExecuteAssertion(
				totalSteps,
				inboxMessages,
				0,
			)
			inboxMessages = inboxMessages[assertion.InboxMessagesConsumed:]
		}
		deadline = continueEvent.Deadline
	}
}

func computeSteps(
	continueEvent arbbridge.ContinueChallengeEvent,
	bisectionEvent arbbridge.ExecutionBisectionEvent,
) uint64 {
	totalSteps := uint64(0)
	for i := uint64(0); i < continueEvent.SegmentIndex.Uint64(); i++ {
		totalSteps += valprotocol.CalculateBisectionStepCount(
			i,
			uint64(len(bisectionEvent.Assertions)),
			bisectionEvent.TotalSteps)
	}
	return totalSteps
}

func executionChallengerUpdate(
	ctx context.Context,
	mach machine.Machine,
	contract arbbridge.ExecutionChallenge,
	inboxMessages []inbox.InboxMessage,
	bisectionEvent arbbridge.ExecutionBisectionEvent,
	challengeEverything bool,
) (machine.Machine, []inbox.InboxMessage, error) {
	challengedAssertionNum, newMachine, err := ChooseAssertionToChallenge(
		mach.Clone(),
		inboxMessages,
		bisectionEvent.Assertions,
		bisectionEvent.TotalSteps)

	if err != nil {
		if !challengeEverything {
			return nil, nil, err
		} else {
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
					inboxMessages,
					0,
				)
				inboxMessages = inboxMessages[assertion.InboxMessagesConsumed:]
			}
			err = nil
		}
	}

	err = contract.ChooseSegment(
		ctx,
		challengedAssertionNum,
		bisectionEvent.Assertions,
		bisectionEvent.TotalSteps,
	)
	return newMachine, inboxMessages, err
}
