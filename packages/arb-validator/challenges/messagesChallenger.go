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
	"fmt"
	errors2 "github.com/pkg/errors"
	"math/big"
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func ChallengeMessagesClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	inbox *structures.MessageStack,
	beforeInbox common.Hash,
	messageCount *big.Int,
	challengeEverything bool,
) (ChallengeState, error) {
	contractWatcher, err := client.NewMessagesChallengeWatcher(address)
	if err != nil {
		return ChallengeContinuing, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockId, startLogIndex, contractWatcher)

	contract, err := client.NewMessagesChallenge(address)
	if err != nil {
		return 0, err
	}

	return challengeMessages(
		reorgCtx,
		eventChan,
		contract,
		client,
		inbox,
		beforeInbox,
		messageCount.Uint64(),
		challengeEverything,
	)
}

func challengeMessages(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.MessagesChallenge,
	client arbbridge.ArbClient,
	inbox *structures.MessageStack,
	beforeInbox common.Hash,
	messageCount uint64,
	challengeEverything bool,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("MessagesChallenge challenger expected InitiateChallengeEvent but got %T", event)
	}

	vmInbox, err := inbox.GenerateVMInbox(beforeInbox, messageCount)
	if err != nil {
		return 0, errors2.Wrap(err, "challenger error generating vm inbox")
	}

	startInbox := uint64(0)
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

		bisectionEvent, ok := event.(arbbridge.MessagesBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge challenger expected MessagesBisectionEvent but got %T", event)
		}

		event, state, err = msgsChallengerUpdate(
			ctx,
			eventChan,
			contract,
			inbox,
			bisectionEvent,
			startInbox,
			vmInbox,
			challengeEverything)

		if challengeEnded(state, err) {
			return state, err
		}

		continueEvent, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge challenger expected ContinueChallengeEvent but got %T", event)
		}
		deadline = continueEvent.Deadline
	}
}

func msgsChallengerUpdate(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.MessagesChallenge,
	inbox *structures.MessageStack,
	bisectionEvent arbbridge.MessagesBisectionEvent,
	startInbox uint64,
	vmInbox *structures.VMInbox,
	challengeEverything bool,
) (arbbridge.Event, ChallengeState, error) {
	// Wait to check if we've already chosen a segment
	timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	if timedOut {
		totalMessages := bisectionEvent.TotalLength.Uint64()
		segmentCount := uint64(len(bisectionEvent.SegmentHashes)) - 1

		inboxSegments, err := inbox.GenerateBisectionReverse(
			bisectionEvent.ChainHashes[0],
			segmentCount,
			totalMessages,
		)
		if err != nil {
			return nil, 0, err
		}

		vmInboxHashes, err := vmInbox.GenerateBisection(
			startInbox,
			segmentCount,
			totalMessages,
		)
		if err != nil {
			return nil, 0, err
		}

		segmentToChallenge, found := findSegmentToChallenge(inboxSegments, bisectionEvent.ChainHashes)
		if !found {
			segmentToChallenge, found = findSegmentToChallenge(vmInboxHashes, bisectionEvent.SegmentHashes)
		}

		if !found {
			if challengeEverything {
				segmentToChallenge = uint64(rand.Int31n(int32(segmentCount)))
			} else {
				return nil, 0, errors.New("Nothing to challenge")
			}
		}
		err = contract.ChooseSegment(ctx, uint16(segmentToChallenge), bisectionEvent.ChainHashes, bisectionEvent.SegmentHashes, bisectionEvent.TotalLength)
		if err != nil {
			return nil, 0, err
		}
		event, state, err = getNextEvent(eventChan)
	}
	return event, state, err
}
