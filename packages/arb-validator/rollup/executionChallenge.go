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

package rollup

import (
	"context"
	"errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

func DefendExecutionClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	startDefender machine.AssertionDefender,
) (ChallengeState, error) {
	contract, err := ethbridge.NewExecutionChallenge(address, client)
	if err != nil {
		return ChallengeContinuing, err
	}
	ctx := context.TODO()
	noteChan := make(chan ethbridge.Notification, 1024)

	go ethbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return defendExecution(
		auth,
		client,
		contract,
		noteChan,
		startDefender,
	)
}

func ChallengeExecutionClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	startMachine machine.Machine,
	startPrecondition *protocol.Precondition,
) (ChallengeState, error) {
	contract, err := ethbridge.NewExecutionChallenge(address, client)
	if err != nil {
		return 0, err
	}
	ctx := context.TODO()
	noteChan := make(chan ethbridge.Notification, 1024)

	go ethbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return challengeExecution(
		auth,
		client,
		contract,
		noteChan,
		startMachine,
		startPrecondition,
	)
}

func defendExecution(
	auth *bind.TransactOpts,
	client *ethclient.Client,
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
			_, err = contract.OneStepProof(auth, defender.GetPrecondition(), defender.GetAssertion().Stub(), proof)
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
		assertions := make([]*protocol.AssertionStub, 0, len(defenders))
		for _, defender := range defenders {
			assertions = append(assertions, defender.GetAssertion().Stub())
		}
		_, err := contract.BisectAssertion(auth, defender.GetPrecondition(), assertions)

		note, state, err := getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := note.Event.(ethbridge.ExecutionBisectionEvent)
		if !ok {
			return 0, errors.New("ExecutionChallenge expected ExecutionBisectionEvent")
		}

		note, state, err = getNextEventWithTimeout(
			auth,
			outChan,
			ev.DeadlineTicks,
			contract,
			client,
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
	auth *bind.TransactOpts,
	client *ethclient.Client,
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
	deadline := ev.DeadlineTicks
	for {
		note, state, err := getNextEventWithTimeout(
			auth,
			outChan,
			deadline,
			contract,
			client,
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
		challengedAssertionNum, m, err := machine.ChooseAssertionToChallenge(mach, ev.Assertions, startPrecondition.TimeBounds)
		if err != nil {
			return 0, err
		}
		preconditions := protocol.GeneratePreconditions(precondition, ev.Assertions)
		_, err = contract.ChooseSegment(
			auth,
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
		deadline = contEv.DeadlineTicks
	}
}
