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

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func DefendMessagesClaim(
	client arbbridge.ArbAuthClient,
	address common.Address,
	startHeight *common.TimeBlocks,
	startLogIndex uint,
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	afterPending common.Hash,
	importedMessagesSlice common.Hash,
	bisectionCount uint64,
) (ChallengeState, error) {
	contractWatcher, err := client.NewMessagesChallengeWatcher(address)
	if err != nil {
		return ChallengeContinuing, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventChan := arbbridge.HandleBlockchainNotifications(ctx, startHeight, startLogIndex, contractWatcher)
	contract, err := client.NewMessagesChallenge(address)
	if err != nil {
		return 0, err
	}
	return defendMessages(
		ctx,
		eventChan,
		contract,
		client,
		pendingInbox,
		beforePending,
		afterPending,
		importedMessagesSlice,
		bisectionCount,
	)
}

func ChallengeMessagesClaim(
	client arbbridge.ArbAuthClient,
	address common.Address,
	startHeight *common.TimeBlocks,
	startLogIndex uint,
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	afterPending common.Hash,
) (ChallengeState, error) {
	contractWatcher, err := client.NewMessagesChallengeWatcher(address)
	if err != nil {
		return ChallengeContinuing, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eventChan := arbbridge.HandleBlockchainNotifications(ctx, startHeight, startLogIndex, contractWatcher)
	contract, err := client.NewMessagesChallenge(address)
	if err != nil {
		return 0, err
	}
	return challengeMessages(
		ctx,
		eventChan,
		contract,
		client,
		pendingInbox,
		beforePending,
		afterPending,
	)
}

func defendMessages(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.MessagesChallenge,
	client arbbridge.ArbClient,
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	afterPending common.Hash,
	importedMessagesSlice common.Hash,
	bisectionCount uint64,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("MessagesChallenge defender expected InitiateChallengeEvent but got %T", event)
	}

	messagesStack, err := pendingInbox.Substack(beforePending, afterPending)
	if err != nil {
		return 0, err
	}

	startPending := beforePending
	endPending := afterPending
	startMessages := value.NewEmptyTuple().Hash()
	endMessages := importedMessagesSlice

	for {
		messageCount, err := pendingInbox.SegmentSize(startPending, endPending)
		if err != nil {
			return 0, err
		}

		if messageCount == 1 {
			timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
			if timedOut {
				pendingNextHash, pendingValueHash, err := pendingInbox.GenerateOneStepProof(startPending)
				if err != nil {
					return 0, err
				}
				messagesNextHash, _, err := messagesStack.GenerateOneStepProof(startMessages)
				if err != nil {
					return 0, err
				}
				err = contract.OneStepProof(ctx, startPending, pendingNextHash, startMessages, messagesNextHash, pendingValueHash)
				if err != nil {
					return 0, err
				}
				event, state, err = getNextEvent(eventChan)
			}

			if err != nil || state != ChallengeContinuing {
				return state, err
			}
			_, ok = event.(arbbridge.OneStepProofEvent)
			if !ok {
				return 0, fmt.Errorf("MessagesChallenge defender expected OneStepProof but got %T", event)
			}
			return ChallengeAsserterWon, nil
		}

		timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
		if timedOut {
			chainHashes, err := pendingInbox.GenerateBisection(startPending, endPending, bisectionCount)
			if err != nil {
				return 0, err
			}
			stackHashes, err := messagesStack.GenerateBisection(startMessages, endMessages, bisectionCount)
			if err != nil {
				return 0, err
			}
			err = contract.Bisect(ctx, chainHashes, stackHashes, new(big.Int).SetUint64(messageCount))
			if err != nil {
				return 0, err
			}

			event, state, err = getNextEvent(eventChan)
		}

		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := event.(arbbridge.MessagesBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge defender expected MessagesBisectionEvent but got %T", event)
		}

		event, state, err = getNextEventWithTimeout(
			ctx,
			eventChan,
			ev.Deadline,
			contract,
			client,
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge defender expected ContinueChallengeEvent but got %T", event)
		}
		startPending = ev.ChainHashes[contEv.SegmentIndex.Uint64()]
		endPending = ev.ChainHashes[contEv.SegmentIndex.Uint64()+1]
		startMessages = ev.SegmentHashes[contEv.SegmentIndex.Uint64()]
		endMessages = ev.SegmentHashes[contEv.SegmentIndex.Uint64()+1]
	}
}

func challengeMessages(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.MessagesChallenge,
	client arbbridge.ArbClient,
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	afterPending common.Hash,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("MessagesChallenge challenger expected InitiateChallengeEvent but got %T", event)
	}

	messagesStack, err := pendingInbox.Substack(beforePending, afterPending)
	if err != nil {
		return 0, err
	}

	deadline := ev.Deadline
	for {
		event, state, err := getNextEventWithTimeout(
			ctx,
			eventChan,
			deadline,
			contract,
			client,
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}

		if _, ok := event.(arbbridge.OneStepProofEvent); ok {
			return ChallengeAsserterWon, nil
		}

		ev, ok := event.(arbbridge.MessagesBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge challenger expected MessagesBisectionEvent but got %T", event)
		}

		timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
		if timedOut {
			pendingChallengedSegment, err := pendingInbox.CheckBisection(ev.ChainHashes)
			if err != nil {
				return 0, err
			}
			messagesChallengedSegment, err := messagesStack.CheckBisection(ev.SegmentHashes)
			if err != nil {
				return 0, err
			}
			maxSegment := pendingChallengedSegment
			if messagesChallengedSegment > maxSegment {
				maxSegment = messagesChallengedSegment
			}

			err = contract.ChooseSegment(ctx, uint16(maxSegment), ev.ChainHashes, ev.SegmentHashes, ev.TotalLength)
			if err != nil {
				return 0, err
			}
			event, state, err = getNextEvent(eventChan)
		}
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge challenger expected ContinueChallengeEvent but got %T", event)
		}
		deadline = contEv.Deadline
	}
}
