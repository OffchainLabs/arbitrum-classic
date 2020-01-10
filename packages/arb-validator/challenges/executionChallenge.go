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
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
)

func DefendExecutionClaim(
	auth *bind.TransactOpts,
	client arbbridge.ArbClient,
	address common.Address,
	precondition *protocol.Precondition,
	numSteps uint32,
	startMachine machine.Machine,
) (ChallengeState, error) {
	contract, err := client.NewExecutionChallenge(address, auth)
	if err != nil {
		return ChallengeContinuing, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan arbbridge.Notification, 1024)

	go arbbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return defendExecution(
		ctx,
		contract,
		noteChan,
		machine.NewAssertionDefender(
			precondition,
			numSteps,
			startMachine,
		),
	)
}

func ChallengeExecutionClaim(
	auth *bind.TransactOpts,
	client arbbridge.ArbClient,
	address common.Address,
	startPrecondition *protocol.Precondition,
	startMachine machine.Machine,
) (ChallengeState, error) {
	contract, err := client.NewExecutionChallenge(address, auth)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan arbbridge.Notification, 1024)

	go arbbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return challengeExecution(
		ctx,
		contract,
		noteChan,
		startMachine,
		startPrecondition,
	)
}

func defendExecution(
	ctx context.Context,
	contract arbbridge.ExecutionChallenge,
	outChan chan arbbridge.Notification,
	startDefender machine.AssertionDefender,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = note.Event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.New("ExecutionChallenge expected InitiateChallengeEvent")
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
			err = contract.OneStepProof(ctx, defender.GetPrecondition(), assertion.Stub(), proof)
			if err != nil {
				return 0, err
			}
			note, state, err := getNextEvent(outChan)
			if err != nil || state != ChallengeContinuing {
				return state, err
			}
			_, ok = note.Event.(arbbridge.OneStepProof)
			if !ok {
				return 0, errors.New("ExecutionChallenge expected OneStepProof")
			}
			return ChallengeAsserterWon, nil
		}

		defenders, assertions := defender.NBisect(50)
		err := contract.BisectAssertion(ctx, defender.GetPrecondition(), assertions, defender.NumSteps())

		note, state, err := getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := note.Event.(arbbridge.ExecutionBisectionEvent)
		if !ok {
			return 0, errors.New("ExecutionChallenge expected ExecutionBisectionEvent")
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
			return 0, errors.New("ExecutionChallenge expected ContinueChallengeEvent")
		}
		defender = defenders[contEv.SegmentIndex.Uint64()]
	}
}

func challengeExecution(
	ctx context.Context,
	contract arbbridge.ExecutionChallenge,
	outChan chan arbbridge.Notification,
	startMachine machine.Machine,
	startPrecondition *protocol.Precondition,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := note.Event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.New("ExecutionChallenge expected InitiateChallengeEvent")
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

		if _, ok := note.Event.(arbbridge.OneStepProof); ok {
			return ChallengeAsserterWon, nil
		}

		ev, ok := note.Event.(arbbridge.ExecutionBisectionEvent)
		if !ok {
			return 0, errors.New("ExecutionChallenge expected ExecutionBisectionEvent")
		}
		challengedAssertionNum, m, err := machine.ChooseAssertionToChallenge(mach, startPrecondition, ev.Assertions, ev.TotalSteps)
		if err != nil {
			return 0, err
		}
		preconditions := protocol.GeneratePreconditions(precondition, ev.Assertions)
		err = contract.ExecutionChallengeChooseSegment(
			ctx,
			challengedAssertionNum,
			preconditions,
			ev.Assertions,
		)
		note, state, err = getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, errors.New("ExecutionChallenge expected ContinueChallengeEvent")
		}
		mach = m
		precondition = preconditions[contEv.SegmentIndex.Uint64()]
		deadline = contEv.Deadline
	}
}
