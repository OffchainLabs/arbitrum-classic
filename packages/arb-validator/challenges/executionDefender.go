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
	precondition *valprotocol.Precondition,
	startMachine machine.Machine,
	numSteps uint64,
	bisectionCount uint32,
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
			precondition,
			numSteps,
			startMachine,
		),
		bisectionCount,
	)
}

func defendExecution(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.ExecutionChallenge,
	client arbbridge.ArbClient,
	startDefender AssertionDefender,
	bisectionCount uint32,
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
	totalSteps := uint64(0)
	assertionLen := uint64(len(bisectionEvent.Assertions))
	segmentIndex := continueEvent.SegmentIndex.Uint64()

	for i := uint64(0); i < segmentIndex; i++ {
		totalSteps += valprotocol.CalculateBisectionStepCount(
			i,
			assertionLen,
			bisectionEvent.TotalSteps)
	}

	mach := defender.initState
	pre := defender.precondition
	// Update mach, precondition, deadline
	assertion, _ := mach.ExecuteAssertion(
		totalSteps,
		pre.BeforeInbox,
		0,
	)
	pre = pre.GeneratePostcondition(valprotocol.NewExecutionAssertionStubFromAssertion(assertion))

	steps := valprotocol.CalculateBisectionStepCount(
		segmentIndex,
		assertionLen,
		bisectionEvent.TotalSteps)

	return NewAssertionDefender(pre, steps, mach)
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
			defender.GetPrecondition(),
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
		pre := defender.GetPrecondition()
		assertion, _ := defender.GetMachineState().ExecuteAssertion(
			1,
			pre.BeforeInbox,
			0,
		)

		err = contract.OneStepProof(
			ctx,
			defender.GetPrecondition(),
			valprotocol.NewExecutionAssertionStubFromAssertion(assertion),
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
