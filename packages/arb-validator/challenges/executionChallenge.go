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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"

	"github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

func DefendExecutionClaim(
	client arbbridge.ArbAuthClient,
	address common.Address,
	bisectionCount uint32,
	precondition *valprotocol.Precondition,
	numSteps uint32,
	startMachine machine.Machine,
) (ChallengeState, error) {
	contract, err := client.NewExecutionChallenge(address)
	if err != nil {
		return ChallengeContinuing, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan arbbridge.Notification, 1024)
	defer close(noteChan)

	parsingChan := arbbridge.HandleBlockchainNotifications(ctx, contract)
	go func() {
		for event := range parsingChan {
			_, ok := event.Event.(arbbridge.NewTimeEvent)
			if !ok {
				noteChan <- event
			}
		}
	}()
	return defendExecution(
		ctx,
		contract,
		noteChan,
		bisectionCount,
		NewAssertionDefender(
			precondition,
			numSteps,
			startMachine,
		),
	)
}

func ChallengeExecutionClaim(
	client arbbridge.ArbAuthClient,
	address common.Address,
	startPrecondition *valprotocol.Precondition,
	startMachine machine.Machine,
	challengeEverything bool,
) (ChallengeState, error) {
	contract, err := client.NewExecutionChallenge(address)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan arbbridge.Notification, 1024)
	defer close(noteChan)

	parsingChan := arbbridge.HandleBlockchainNotifications(ctx, contract)
	go func() {
		for event := range parsingChan {
			_, ok := event.Event.(arbbridge.NewTimeEvent)
			if !ok {
				noteChan <- event
			}
		}
	}()
	return challengeExecution(
		ctx,
		contract,
		noteChan,
		startMachine,
		startPrecondition,
		challengeEverything,
	)
}

func defendExecution(
	ctx context.Context,
	contract arbbridge.ExecutionChallenge,
	outChan chan arbbridge.Notification,
	bisectionCount uint32,
	startDefender AssertionDefender,
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
			note, state, err := getNextEvent(outChan)
			if err != nil || state != ChallengeContinuing {
				return state, err
			}
			_, ok = note.Event.(arbbridge.OneStepProofEvent)
			if !ok {
				return 0, fmt.Errorf("ExecutionChallenge defender expected OneStepProof but got %T", note.Event)
			}
			return ChallengeAsserterWon, nil
		}

		defenders, assertions := defender.NBisect(bisectionCount)
		err := contract.BisectAssertion(ctx, defender.GetPrecondition(), assertions, defender.NumSteps())
		if err != nil {
			return 0, err
		}
		note, state, err := getNextEvent(outChan)
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
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge defender expected ContinueChallengeEvent but got %T", note.Event)
		}
		defender = defenders[contEv.SegmentIndex.Uint64()]
	}
}

func challengeExecution(
	ctx context.Context,
	contract arbbridge.ExecutionChallenge,
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
		challengedAssertionNum, m, err := ChooseAssertionToChallenge(mach.Clone(), startPrecondition, ev.Assertions, ev.TotalSteps)
		if err != nil && challengeEverything {
			cMach := mach.Clone()
			challengedAssertionNum = uint16(rand.Int31n(int32(len(ev.Assertions))))
			for i := 0; i < len(ev.Assertions); i++ {
				stepCount := CalculateBisectionStepCount(uint32(i), uint32(len(ev.Assertions)), ev.TotalSteps)
				m = cMach.Clone()
				assertion, _ := cMach.ExecuteAssertion(stepCount, startPrecondition.TimeBounds, startPrecondition.BeforeInbox.(value.TupleValue))
				startPrecondition = startPrecondition.GeneratePostcondition(valprotocol.NewExecutionAssertionStubFromAssertion(assertion))
			}
			err = nil
		}
		if err != nil {
			return 0, err
		}
		preconditions := valprotocol.GeneratePreconditions(precondition, ev.Assertions)
		err = contract.ChooseSegment(
			ctx,
			challengedAssertionNum,
			preconditions,
			ev.Assertions,
			ev.TotalSteps,
		)
		note, state, err = getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("ExecutionChallenge challenger expected ContinueChallengeEvent but got %T", note.Event)
		}
		mach = m
		precondition = preconditions[contEv.SegmentIndex.Uint64()]
		deadline = contEv.Deadline
	}
}
