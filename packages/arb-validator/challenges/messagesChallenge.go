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
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	afterPending common.Hash,
	importedMessagesSlice common.Hash,
	bisectionCount uint64,
) (ChallengeState, error) {
	contract, err := client.NewMessagesChallenge(address)
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
	return defendMessages(
		ctx,
		noteChan,
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
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	afterPending common.Hash,
) (ChallengeState, error) {
	contract, err := client.NewMessagesChallenge(address)
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
	return challengeMessages(
		ctx,
		noteChan,
		contract,
		client,
		pendingInbox,
		beforePending,
		afterPending,
	)
}

func defendMessages(
	ctx context.Context,
	outChan chan arbbridge.Notification,
	contract arbbridge.MessagesChallenge,
	client arbbridge.ArbClient,
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	afterPending common.Hash,
	importedMessagesSlice common.Hash,
	bisectionCount uint64,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = note.Event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("MessagesChallenge defender expected InitiateChallengeEvent but got %T", note.Event)
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
			note, state, err := getNextEvent(outChan)
			if err != nil || state != ChallengeContinuing {
				return state, err
			}
			_, ok = note.Event.(arbbridge.OneStepProofEvent)
			if !ok {
				return 0, fmt.Errorf("MessagesChallenge defender expected OneStepProof but got %T", note.Event)
			}
			return ChallengeAsserterWon, nil
		}

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

		note, state, err := getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := note.Event.(arbbridge.MessagesBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge defender expected MessagesBisectionEvent but got %T", note.Event)
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
			return 0, fmt.Errorf("MessagesChallenge defender expected ContinueChallengeEvent but got %T", note.Event)
		}
		startPending = chainHashes[contEv.SegmentIndex.Uint64()]
		endPending = chainHashes[contEv.SegmentIndex.Uint64()+1]
		startMessages = stackHashes[contEv.SegmentIndex.Uint64()]
		endMessages = stackHashes[contEv.SegmentIndex.Uint64()+1]
	}
}

func challengeMessages(
	ctx context.Context,
	outChan chan arbbridge.Notification,
	contract arbbridge.MessagesChallenge,
	client arbbridge.ArbClient,
	pendingInbox *structures.MessageStack,
	beforePending common.Hash,
	afterPending common.Hash,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := note.Event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("MessagesChallenge challenger expected InitiateChallengeEvent but got %T", note.Event)
	}

	messagesStack, err := pendingInbox.Substack(beforePending, afterPending)
	if err != nil {
		return 0, err
	}

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

		ev, ok := note.Event.(arbbridge.MessagesBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge challenger expected MessagesBisectionEvent but got %T", note.Event)
		}
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
		note, state, err = getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge challenger expected ContinueChallengeEvent but got %T", note.Event)
		}
		deadline = contEv.Deadline
	}
}
