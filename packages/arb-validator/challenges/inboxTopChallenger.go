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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func ChallengeInboxTopClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	challengeAddress common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	inbox *structures.MessageStack,
	challengeEverything bool,
) (ChallengeState, error) {
	contractWatcher, err := client.NewInboxTopChallengeWatcher(challengeAddress)
	if err != nil {
		return 0, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockId, startLogIndex, contractWatcher)

	contract, err := client.NewInboxTopChallenge(challengeAddress)
	if err != nil {
		return 0, err
	}
	logger.Info().Msg("=======> challenging inbox top claim")
	return challengeInboxTop(
		reorgCtx,
		eventChan,
		contract,
		client,
		inbox,
		challengeEverything,
	)
}

func challengeInboxTop(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.InboxTopChallenge,
	client arbbridge.ArbClient,
	inbox *structures.MessageStack,
	challengeEverything bool,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.Errorf("InboxTopChallenge challenger expected InitiateChallengeEvent but got %T", event)
	}

	deadline := ev.Deadline
	for {
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

		bisectEvent, ok := event.(arbbridge.InboxTopBisectionEvent)
		if !ok {
			return 0, errors.Errorf("InboxTopChallenge challenger expected InboxTopBisectionEvent but got %T", event)
		}

		event, state, err = inboxChallengerUpdate(
			ctx,
			eventChan,
			contract,
			inbox,
			challengeEverything,
			bisectEvent)

		if challengeEnded(state, err) {
			return state, err
		}

		continueEvent, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, errors.Errorf("InboxTopChallenge challenger expected ContinueChallengeEvent but got %T", event)
		}
		deadline = continueEvent.Deadline
	}
}

func inboxChallengerUpdate(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.InboxTopChallenge,
	inbox *structures.MessageStack,
	challengeEverything bool,
	bisectionEvent arbbridge.InboxTopBisectionEvent,
) (arbbridge.Event, ChallengeState, error) {
	// Wait to check if we've already chosen a segment
	timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	if !timedOut {
		return event, state, err
	}

	bisectionLength := bisectionEvent.TotalLength.Uint64()
	segments, err := getSegments(
		inbox,
		bisectionEvent)
	if err != nil {
		return nil, 0, err
	}

	segmentToChallenge, found := findSegmentToChallenge(segments, bisectionEvent.ChainHashes)

	if !found {
		if challengeEverything {
			segmentToChallenge = uint64(rand.Int31n(int32(len(bisectionEvent.ChainHashes) - 1)))
		} else {
			return nil, 0, errors.New("can't find inbox segment to challenge")
		}
	}
	err = contract.ChooseSegment(
		ctx,
		uint16(segmentToChallenge),
		bisectionEvent.ChainHashes,
		bisectionLength)

	if err != nil {
		return nil, 0, err
	}
	return getNextEvent(eventChan)
}

func getSegments(
	inbox *structures.MessageStack,
	bisectionEvent arbbridge.InboxTopBisectionEvent,
) ([]common.Hash, error) {
	bisectionLength := bisectionEvent.TotalLength.Uint64()
	return inbox.GenerateBisection(
		bisectionEvent.ChainHashes[0],
		uint64(len(bisectionEvent.ChainHashes))-1,
		bisectionLength)
}
