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
	"math/big"
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
)

func DefendExecutionClaim(
	client arbbridge.ArbAuthClient,
	address common.Address,
	precondition *valprotocol.Precondition,
	startMachine machine.Machine,
	numSteps uint32,
	bisectionCount uint32,
) (ChallengeState, error) {
	contractWatcher, err := client.NewExecutionChallengeWatcher(address)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan arbbridge.Notification, 1024)
	defer close(noteChan)

	parsingChan := arbbridge.HandleBlockchainNotifications(ctx, common.NewTimeBlocks(big.NewInt(0)), contractWatcher)
	go func() {
		for event := range parsingChan {
			_, ok := event.Event.(arbbridge.NewTimeEvent)
			if !ok {
				noteChan <- event
			}
		}
	}()

	contract, err := client.NewExecutionChallenge(address)
	if err != nil {
		return ChallengeContinuing, err
	}
	return defendExecution(
		ctx,
		contract,
		client,
		noteChan,
		NewAssertionDefender(
			precondition,
			numSteps,
			startMachine,
		),
		bisectionCount,
	)
}

func ChallengeExecutionClaim(
	client arbbridge.ArbAuthClient,
	address common.Address,
	startPrecondition *valprotocol.Precondition,
	startMachine machine.Machine,
	challengeEverything bool,
) (ChallengeState, error) {
	contractWatcher, err := client.NewExecutionChallengeWatcher(address)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan arbbridge.Notification, 1024)
	defer close(noteChan)

	parsingChan := arbbridge.HandleBlockchainNotifications(ctx, common.NewTimeBlocks(big.NewInt(0)), contractWatcher)
	go func() {
		for event := range parsingChan {
			_, ok := event.Event.(arbbridge.NewTimeEvent)
			if !ok {
				noteChan <- event
			}
		}
	}()

	contract, err := client.NewExecutionChallenge(address)
	if err != nil {
		return 0, err
	}
	return challengeExecution(
		ctx,
		contract,
		client,
		noteChan,
		startMachine,
		startPrecondition,
		challengeEverything,
	)
}

func defendExecution(
	ctx context.Context,
	contract arbbridge.ExecutionChallenge,
	client arbbridge.ArbClient,
	outChan chan arbbridge.Notification,
	startDefender AssertionDefender,
	bisectionCount uint32,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = note.Event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("ExecutionChallenge expected InitiateChallengeEvent but got %T", note.Event)
	}

	defender := startDefender

	for {
		if defender.NumSteps() == 1 {
			timedOut, note, state, err := getNextEventIfExists(ctx, outChan, replayTimeout)
			if timedOut {
				proof, err := defender.SolidityOneStepProof()
				if err != nil {
					return 0, err
				}
				pre := defender.GetPrecondition()
				assertion, _ := defender.GetMachineState().ExecuteAssertion(1, pre.TimeBounds, pre.BeforeInbox.(value.TupleValue))
				err = contract.OneStepProof(
					ctx,
					defender.GetPrecondition(),
					valprotocol.NewExecutionAssertionStubFromAssertion(assertion),
					proof,
				)
				if err != nil {
					return 0, err
				}
				note, state, err = getNextEvent(outChan)
			}

			if err != nil || state != ChallengeContinuing {
				return state, err
			}
			_, ok = note.Event.(arbbridge.OneStepProofEvent)
			if !ok {
				return 0, fmt.Errorf("ExecutionChallenge defender expected OneStepProof but got %T", note.Event)
			}
			return ChallengeAsserterWon, nil
		}
		timedOut, note, state, err := getNextEventIfExists(ctx, outChan, replayTimeout)
		var defenders []AssertionDefender = nil
		if timedOut {
			var assertions []*valprotocol.ExecutionAssertionStub
			defenders, assertions = defender.NBisect(bisectionCount)
			err := contract.BisectAssertion(ctx, defender.GetPrecondition(), assertions, defender.NumSteps())
			if err != nil {
				return 0, err
			}
			note, state, err = getNextEvent(outChan)
		}

		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := note.Event.(arbbridge.ExecutionBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge defender expected ExecutionBisectionEvent but got %T", note.Event)
		}

		note, state, err = getNextEventWithTimeout(
			ctx,
			outChan,
			ev.Deadline,
			contract,
			client,
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge defender expected ContinueChallengeEvent but got %T", note.Event)
		}

		if timedOut {
			// Freshly bisected assertion
			defender = defenders[contEv.SegmentIndex.Uint64()]
		} else {
			// Replayed from existing event
			totalSteps := uint32(0)
			for i := uint32(0); i < uint32(contEv.SegmentIndex.Uint64()); i++ {
				totalSteps += structures.CalculateBisectionStepCount(i, uint32(len(ev.Assertions)), ev.TotalSteps)
			}

			mach := defender.initState
			pre := defender.precondition
			// Update mach, precondition, deadline
			assertion, _ := mach.ExecuteAssertion(totalSteps, pre.TimeBounds, pre.BeforeInbox.(value.TupleValue))
			pre = pre.GeneratePostcondition(valprotocol.NewExecutionAssertionStubFromAssertion(assertion))

			steps := structures.CalculateBisectionStepCount(uint32(contEv.SegmentIndex.Uint64()), uint32(len(ev.Assertions)), ev.TotalSteps)
			defender = NewAssertionDefender(pre, steps, mach)
		}
	}
}

func challengeExecution(
	ctx context.Context,
	contract arbbridge.ExecutionChallenge,
	client arbbridge.ArbClient,
	outChan chan arbbridge.Notification,
	startMachine machine.Machine,
	startPrecondition *valprotocol.Precondition,
	challengeEverything bool,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := note.Event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("ExecutionChallenge challenger expected InitiateChallengeEvent but got %T", note.Event)
	}

	mach := startMachine
	precondition := startPrecondition
	deadline := ev.Deadline
	for {
		note, state, err := getNextEventWithTimeout(
			ctx,
			outChan,
			deadline,
			contract,
			client,
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}

		if _, ok := note.Event.(arbbridge.OneStepProofEvent); ok {
			return ChallengeAsserterWon, nil
		}

		ev, ok := note.Event.(arbbridge.ExecutionBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge challenger expected ExecutionBisectionEvent but got %T", note.Event)
		}
		timedOut, note, state, err := getNextEventIfExists(ctx, outChan, replayTimeout)
		var preconditions []*valprotocol.Precondition
		var m machine.Machine
		if timedOut {
			var challengedAssertionNum uint16
			challengedAssertionNum, m, err = ChooseAssertionToChallenge(mach.Clone(), precondition, ev.Assertions, ev.TotalSteps)
			if err != nil && challengeEverything {
				pre := precondition
				cMach := mach.Clone()
				challengedAssertionNum = uint16(rand.Int31n(int32(len(ev.Assertions))))
				for i := 0; i < len(ev.Assertions); i++ {
					stepCount := structures.CalculateBisectionStepCount(uint32(i), uint32(len(ev.Assertions)), ev.TotalSteps)
					m = cMach.Clone()
					assertion, _ := cMach.ExecuteAssertion(stepCount, pre.TimeBounds, pre.BeforeInbox.(value.TupleValue))
					pre = pre.GeneratePostcondition(valprotocol.NewExecutionAssertionStubFromAssertion(assertion))
				}
				err = nil
			}
			if err != nil {
				return 0, err
			}
			preconditions = valprotocol.GeneratePreconditions(precondition, ev.Assertions)
			err = contract.ChooseSegment(
				ctx,
				challengedAssertionNum,
				preconditions,
				ev.Assertions,
				ev.TotalSteps,
			)
			note, state, err = getNextEvent(outChan)
		}

		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge challenger expected ContinueChallengeEvent but got %T", note.Event)
		}

		// Update mach, precondition, deadline
		if timedOut {
			// Freshly bisected assertion
			mach = m
			precondition = preconditions[contEv.SegmentIndex.Uint64()]
		} else {
			// Replayed from existing event
			totalSteps := uint32(0)
			for i := uint32(0); i < uint32(contEv.SegmentIndex.Uint64()); i++ {
				totalSteps += structures.CalculateBisectionStepCount(i, uint32(len(ev.Assertions)), ev.TotalSteps)
			}
			assertion, _ := mach.ExecuteAssertion(totalSteps, startPrecondition.TimeBounds, startPrecondition.BeforeInbox.(value.TupleValue))
			precondition = precondition.GeneratePostcondition(valprotocol.NewExecutionAssertionStubFromAssertion(assertion))
		}
		deadline = contEv.Deadline
	}
}
