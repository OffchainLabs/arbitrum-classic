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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"github.com/pkg/errors"
)

func DefendExecutionClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	startMachine machine.Machine,
	assertion *valprotocol.ExecutionAssertionStub,
	inboxStack *structures.MessageStack,
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
		logger.Fatal().Msg("nil startMachine in DefendExecutionClaim")
	}
	return defendExecution(
		reorgCtx,
		eventChan,
		contract,
		client,
		NewAssertionDefender(
			numSteps,
			startMachine,
			inboxStack,
			assertion,
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
		return 0, errors.Errorf("ExecutionChallenge expected InitiateChallengeEvent but got %T", event)
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
			return 0, errors.Errorf("ExecutionChallenge defender expected ExecutionBisectionEvent but got %T", event)
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
			return 0, errors.Errorf("ExecutionChallenge defender expected ContinueChallengeEvent but got %T", event)
		}

		if bisected {
			// Freshly bisected assertion
			defender = defenders[continueEvent.SegmentIndex.Uint64()]
		} else {
			// Replayed from existing event
			defenderPointer, err := defender.MoveDefender(bisectionEvent, continueEvent)
			if err != nil {
				return 0, err
			}
			defender = *defenderPointer
		}
	}
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
		defenders = defender.NBisect(uint64(bisectionCount))
		assertions := make([]*valprotocol.ExecutionAssertionStub, 0, len(defenders))
		for _, def := range defenders {
			assertions = append(assertions, def.AssertionStub())
		}
		err := contract.BisectAssertion(
			ctx,
			assertions,
			defender.NumSteps())
		if err != nil {
			return nil, 0, defenders, makeBisection, err
		}
		event, state, err = getNextEvent(eventChan)
		if err != nil {
			return nil, 0, defenders, makeBisection, err
		}
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
		proof, msg, err := defender.SolidityOneStepProof()
		if err != nil {
			return 0, err
		}
		if msg != nil {
			err = contract.OneStepProofWithMessage(
				ctx,
				defender.AssertionStub(),
				proof,
				*msg,
			)
		} else {
			err = contract.OneStepProof(
				ctx,
				defender.AssertionStub(),
				proof,
			)
		}
		if err != nil {
			return 0, err
		}

		event, state, err = getNextEvent(eventChan)
		if err != nil {
			return 0, err
		}
	}

	if challengeEnded(state, err) {
		return state, err
	}

	_, ok := event.(arbbridge.OneStepProofEvent)
	if !ok {
		return 0, errors.Errorf("ExecutionChallenge defender expected OneStepProof but got %T", event)
	}
	return ChallengeAsserterWon, nil
}
