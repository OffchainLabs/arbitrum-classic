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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func DefendPendingTopClaim(
	client arbbridge.ArbAuthClient,
	address common.Address,
	pendingInbox *structures.MessageStack,
	afterPendingTop common.Hash,
	topPending common.Hash,
	bisectionCount uint64,
) (ChallengeState, error) {
	contractWatcher, err := client.NewPendingTopChallengeWatcher(address)
	if err != nil {
		return ChallengeContinuing, err
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
	contract, err := client.NewPendingTopChallenge(address)
	if err != nil {
		return ChallengeContinuing, err
	}
	return defendPendingTop(
		ctx,
		noteChan,
		contract,
		client,
		pendingInbox,
		afterPendingTop,
		topPending,
		bisectionCount,
	)
}

func ChallengePendingTopClaim(
	client arbbridge.ArbAuthClient,
	address common.Address,
	pendingInbox *structures.MessageStack,
) (ChallengeState, error) {
	contractWatcher, err := client.NewPendingTopChallengeWatcher(address)
	if err != nil {
		return ChallengeContinuing, err
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
	contract, err := client.NewPendingTopChallenge(address)
	if err != nil {
		return 0, err
	}
	return challengePendingTop(
		ctx,
		noteChan,
		contract,
		client,
		pendingInbox,
	)
}

func defendPendingTop(
	ctx context.Context,
	outChan chan arbbridge.Notification,
	contract arbbridge.PendingTopChallenge,
	client arbbridge.ArbClient,
	pendingInbox *structures.MessageStack,
	afterPendingTop common.Hash,
	topPending common.Hash,
	bisectionCount uint64,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = note.Event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("PendingTopChallenge defender expected InitiateChallengeEvent but got %T", note.Event)
	}

	startState := afterPendingTop
	endState := topPending

	for {
		messageCount, err := pendingInbox.SegmentSize(startState, endState)
		if err != nil {
			return 0, err
		}

		if messageCount == 1 {
			nextHash, valueHash, err := pendingInbox.GenerateOneStepProof(startState)
			if err != nil {
				return 0, err
			}
			err = contract.OneStepProof(ctx, startState, nextHash, valueHash)
			if err != nil {
				return 0, err
			}
			note, state, err := getNextEvent(outChan)
			if err != nil || state != ChallengeContinuing {
				return state, err
			}
			_, ok = note.Event.(arbbridge.OneStepProofEvent)
			if !ok {
				return 0, fmt.Errorf("PendingTopChallenge defender expected OneStepProof but got %T", note.Event)
			}
			return ChallengeAsserterWon, nil
		}

		chainHashes, err := pendingInbox.GenerateBisection(startState, endState, bisectionCount)
		if err != nil {
			return 0, err
		}
		err = contract.Bisect(ctx, chainHashes, new(big.Int).SetUint64(messageCount))
		if err != nil {
			return 0, err
		}

		note, state, err := getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := note.Event.(arbbridge.PendingTopBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("PendingTopChallenge defender expected PendingTopBisectionEvent but got %T", note.Event)
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
			return 0, fmt.Errorf("PendingTopChallenge defender expected ContinueChallengeEvent but got %T", note.Event)
		}
		startState = chainHashes[contEv.SegmentIndex.Uint64()]
		endState = chainHashes[contEv.SegmentIndex.Uint64()+1]
	}
}

func challengePendingTop(
	ctx context.Context,
	outChan chan arbbridge.Notification,
	contract arbbridge.PendingTopChallenge,
	client arbbridge.ArbClient,
	pendingInbox *structures.MessageStack,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := note.Event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("PendingTopChallenge challenger expected InitiateChallengeEvent but got %T", note.Event)
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

		ev, ok := note.Event.(arbbridge.PendingTopBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("PendingTopChallenge challenger expected PendingTopBisectionEvent but got %T", note.Event)
		}
		challengedSegment, err := pendingInbox.CheckBisection(ev.ChainHashes)
		if err != nil {
			return 0, err
		}
		err = contract.ChooseSegment(ctx, uint16(challengedSegment), ev.ChainHashes, uint32(ev.TotalLength.Uint64()))
		if err != nil {
			return 0, err
		}
		note, state, err = getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("PendingTopChallenge challenger expected ContinueChallengeEvent but got %T", note.Event)
		}
		deadline = contEv.Deadline
	}
}
