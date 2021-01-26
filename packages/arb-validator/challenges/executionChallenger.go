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
	"github.com/pkg/errors"
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func ChallengeExecutionClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	inboxStack *structures.MessageStack,
	numSteps uint64,
	startMachine machine.Machine,
	beforeInboxHash common.Hash,
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

	messages, err := inboxStack.GetAllMessagesAfter(beforeInboxHash)
	if err != nil {
		logger.Fatal().Msg("before inbox hash must be valid")
	}

	// Last value returned is not an error type
	assertion, _, _ := startMachine.Clone().ExecuteAssertion(numSteps, true, messages, true)
	stub := structures.NewExecutionAssertionStubFromWholeAssertion(assertion, beforeInboxHash, inboxStack)

	return challengeExecution(
		reorgCtx,
		eventChan,
		contract,
		client,
		NewAssertionDefender(
			numSteps,
			startMachine,
			inboxStack,
			stub,
		),
		challengeEverything,
		challengeType,
	)
}

func challengeExecution(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.ExecutionChallenge,
	client arbbridge.ArbClient,
	defender AssertionDefender,
	challengeEverything bool,
	challengeType ExecutionChallengeInfo,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.Errorf("ExecutionChallenge challenger expected InitiateChallengeEvent but got %T", event)
	}

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
			return 0, errors.Errorf("ExecutionChallenge challenger expected ExecutionBisectionEvent but got %T", event)
		}

		chooseSegment, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)

		if chooseSegment {
			var challengedAssertionNum int
			challengedAssertionNum, defender, err = chooseDefender(defender, bisectionEvent, challengeEverything)
			if err != nil {
				return state, err
			}
			if err := contract.ChooseSegment(
				ctx,
				uint16(challengedAssertionNum),
				bisectionEvent.AssertionHashes,
			); err != nil {
				return state, err
			}
			event, state, err = getNextEvent(eventChan)
		}

		if challengeEnded(state, err) {
			return state, err
		}

		continueEvent, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, errors.Errorf("ExecutionChallenge challenger expected ContinueChallengeEvent but got %T", event)
		}

		// Update mach, precondition, deadline
		if !chooseSegment {
			// Replayed from existing event
			defenderPointer, err := defender.MoveDefender(bisectionEvent, continueEvent)
			if err != nil {
				return 0, err
			}
			defender = *defenderPointer
		}
		deadline = continueEvent.Deadline
	}
}

func computeStepsUpTo(
	segmentsToSkip,
	totalSegmentCount,
	totalSteps uint64,
) uint64 {
	stepCount := uint64(0)
	for i := uint64(0); i < segmentsToSkip; i++ {
		stepCount += valprotocol.CalculateBisectionStepCount(
			i,
			totalSegmentCount,
			totalSteps,
		)
	}
	return stepCount
}

func chooseDefender(
	defender AssertionDefender,
	bisectionEvent arbbridge.ExecutionBisectionEvent,
	challengeEverything bool,
) (int, AssertionDefender, error) {
	defenders := defender.NBisect(uint64(len(bisectionEvent.AssertionHashes)))
	for i, defender := range defenders {
		if valprotocol.ExecutionDataHash(defender.numSteps, defender.assertion) != bisectionEvent.AssertionHashes[i] {
			return i, defender, nil
		}
	}
	if !challengeEverything {
		return 0, AssertionDefender{}, errors.New("all assertions were valid")
	}

	challengedAssertionNum := rand.Int31n(int32(len(bisectionEvent.AssertionHashes)))
	return int(challengedAssertionNum), defenders[challengedAssertionNum], nil
}
