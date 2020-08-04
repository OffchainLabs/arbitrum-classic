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
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
)

func DefendExecutionClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	inboxMessages []inbox.InboxMessage,
	startMachine machine.Machine,
	numSteps uint64,
	bisectionCount uint32,
	challengeType ExecutionChallengeInfo,
) (ChallengeState, error) {
	contractWatcher, err := client.NewExecutionChallengeWatcher(address)
	if err != nil {
		return 0, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockId, startLogIndex, contractWatcher)

	contract, err := client.NewExecutionChallenge(address)
	if err != nil {
		return ChallengeContinuing, err
	}

	if startMachine == nil {
		log.Fatal("nil startMachine in DefendExecutionClaim")
	}
	return defendExecution(
		reorgCtx,
		eventChan,
		contract,
		client,
		NewAssertionDefender(
			inboxMessages,
			numSteps,
			startMachine,
		),
		bisectionCount,
		challengeType,
	)
}

func defendExecution(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.ExecutionChallenge,
	client arbbridge.ArbClient,
	startDefender AssertionDefender,
	bisectionCount uint32,
	challengeType ExecutionChallengeInfo,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("ExecutionChallenge expected InitiateChallengeEvent but got %T", event)
	}

	defender := startDefender

	for {
		cont := ContinueChallenge(challengeType)

		if cont {
			if challengeType.isDiscontinueType {
				challengeType.currentRound += 1
			}
		} else {
			return DefenderDiscontinued, nil
		}

		if defender.NumSteps() == 1 {
			return runExecutionOneStepProof(ctx, eventChan, defender, contract)
		}

		event, state, defenders, bisected, err := executionDefenderUpdate(
			ctx,
			eventChan,
			contract,
			defender,
			bisectionCount)

		if challengeEnded(state, err) {
			return state, err
		}

		bisectionEvent, ok := event.(arbbridge.ExecutionBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge defender expected ExecutionBisectionEvent but got %T", event)
		}

		// get challenger update
		event, state, err = getNextEventWithTimeout(
			ctx,
			eventChan,
			bisectionEvent.Deadline,
			contract,
			client,
		)
		if challengeEnded(state, err) {
			return state, err
		}

		continueEvent, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge defender expected ContinueChallengeEvent but got %T", event)
		}

		if bisected {
			// Freshly bisected assertion
			defender = defenders[continueEvent.SegmentIndex.Uint64()]
		} else {
			// Replayed from existing event
			defender = updateExecutionData(continueEvent, defender, bisectionEvent)
		}
	}
}

func updateExecutionData(
	continueEvent arbbridge.ContinueChallengeEvent,
	defender AssertionDefender,
	bisectionEvent arbbridge.ExecutionBisectionEvent,
) AssertionDefender {
	totalSteps := computeSteps(continueEvent, bisectionEvent)

	mach := defender.initState
	inboxMessages := defender.inboxMessages
	// Update mach, precondition, deadline
	assertion, _ := mach.ExecuteAssertion(
		totalSteps,
		inboxMessages,
		0,
	)
	inboxMessages = inboxMessages[assertion.InboxMessagesConsumed:]

	steps := valprotocol.CalculateBisectionStepCount(
		continueEvent.SegmentIndex.Uint64(),
		uint64(len(bisectionEvent.Assertions)),
		bisectionEvent.TotalSteps)

	return NewAssertionDefender(inboxMessages, steps, mach)
}

func executionDefenderUpdate(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.ExecutionChallenge,
	defender AssertionDefender,
	bisectionCount uint32,
) (arbbridge.Event, ChallengeState, []AssertionDefender, bool, error) {
	makeBisection, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	var defenders []AssertionDefender = nil
	if makeBisection {
		var assertions []*valprotocol.ExecutionAssertionStub
		defenders, assertions = defender.NBisect(uint64(bisectionCount))
		err := contract.BisectAssertion(
			ctx,
			assertions,
			defender.NumSteps())
		if err != nil {
			return nil, 0, defenders, makeBisection, err
		}
		event, state, err = getNextEvent(eventChan)
	}

	return event, state, defenders, makeBisection, err
}

func runExecutionOneStepProof(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	defender AssertionDefender,
	contract arbbridge.ExecutionChallenge,
) (ChallengeState, error) {
	timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	if timedOut {
		proof, err := defender.SolidityOneStepProof()
		if err != nil {
			return 0, err
		}
		assertion, _ := defender.GetMachineState().ExecuteAssertion(
			1,
			defender.GetInboxMessages(),
			0,
		)

		err = contract.OneStepProof(
			ctx,
			valprotocol.NewExecutionAssertionStubFromAssertion(assertion, defender.GetInboxMessages()),
			proof,
		)
		if err != nil {
			return 0, err
		}
		event, state, err = getNextEvent(eventChan)
	}

	if challengeEnded(state, err) {
		return state, err
	}

	_, ok := event.(arbbridge.OneStepProofEvent)
	if !ok {
		return 0, fmt.Errorf("ExecutionChallenge defender expected OneStepProof but got %T", event)
	}
	return ChallengeAsserterWon, nil
}
