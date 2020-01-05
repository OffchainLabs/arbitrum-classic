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

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

func DefendExecutionClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	precondition *protocol.Precondition,
	numSteps uint32,
	assertion *protocol.ExecutionAssertionStub,
	startMachine machine.Machine,
) (ChallengeState, error) {
	contract, err := ethbridge.NewExecutionChallenge(address, client, auth)
	if err != nil {
		return ChallengeContinuing, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan ethbridge.Notification, 1024)

	go ethbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return defendExecution(
		ctx,
		contract,
		noteChan,
		machine.NewAssertionDefender(
			precondition,
			numSteps,
			assertion,
			startMachine,
		),
	)
}

func ChallengeExecutionClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	startPrecondition *protocol.Precondition,
	startMachine machine.Machine,
) (ChallengeState, error) {
	contract, err := ethbridge.NewExecutionChallenge(address, client, auth)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan ethbridge.Notification, 1024)

	go ethbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
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
	contract *ethbridge.ExecutionChallenge,
	outChan chan ethbridge.Notification,
	startDefender machine.AssertionDefender,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = note.Event.(ethbridge.InitiateChallengeEvent)
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
			_, err = contract.OneStepProof(ctx, defender.GetPrecondition(), defender.GetAssertion(), proof)
			if err != nil {
				return 0, err
			}
			note, state, err := getNextEvent(outChan)
			if err != nil || state != ChallengeContinuing {
				return state, err
			}
			_, ok = note.Event.(ethbridge.OneStepProof)
			if !ok {
				return 0, errors.New("ExecutionChallenge expected OneStepProof")
			}
			return ChallengeAsserterWon, nil
		}

		defenders := defender.NBisect(50)
		assertions := make([]*protocol.ExecutionAssertionStub, 0, len(defenders))
		for _, defender := range defenders {
			assertions = append(assertions, defender.GetAssertion())
		}
		_, err := contract.BisectAssertion(ctx, defender.GetPrecondition(), assertions, defender.NumSteps())

		note, state, err := getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := note.Event.(ethbridge.ExecutionBisectionEvent)
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
		contEv, ok := note.Event.(ethbridge.ContinueChallengeEvent)
		if !ok {
			return 0, errors.New("ExecutionChallenge expected ContinueChallengeEvent")
		}
		defender = defenders[contEv.SegmentIndex.Uint64()]
	}
}

func challengeExecution(
	ctx context.Context,
	contract *ethbridge.ExecutionChallenge,
	outChan chan ethbridge.Notification,
	startMachine machine.Machine,
	startPrecondition *protocol.Precondition,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := note.Event.(ethbridge.InitiateChallengeEvent)
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

		if _, ok := note.Event.(ethbridge.OneStepProof); ok {
			return ChallengeAsserterWon, nil
		}

		ev, ok := note.Event.(ethbridge.ExecutionBisectionEvent)
		if !ok {
			return 0, errors.New("ExecutionChallenge expected ExecutionBisectionEvent")
		}
		challengedAssertionNum, m, err := machine.ChooseAssertionToChallenge(mach, ev.Assertions, startPrecondition.TimeBounds, ev.TotalSteps)
		if err != nil {
			return 0, err
		}
		preconditions := protocol.GeneratePreconditions(precondition, ev.Assertions)
		_, err = contract.ChooseSegment(
			ctx,
			challengedAssertionNum,
			preconditions,
			ev.Assertions,
		)
		note, state, err = getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(ethbridge.ContinueChallengeEvent)
		if !ok {
			return 0, errors.New("ExecutionChallenge expected ContinueChallengeEvent")
		}
		mach = m
		precondition = preconditions[contEv.SegmentIndex.Uint64()]
		deadline = contEv.Deadline
	}
}
