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
	"log"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func DefendPendingTopClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *structures.BlockId,
	startLogIndex uint,
	pendingInbox *structures.MessageStack,
	afterPendingTop common.Hash,
	topPending common.Hash,
	bisectionCount uint64,
) (ChallengeState, error) {
	contractWatcher, err := client.NewPendingTopChallengeWatcher(address)
	if err != nil {
		return 0, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockId, startLogIndex, contractWatcher)

	contract, err := client.NewPendingTopChallenge(address)
	if err != nil {
		return 0, err
	}
	log.Println("=======> defending pending top claim")

	return defendPendingTop(
		reorgCtx,
		eventChan,
		contract,
		client,
		pendingInbox,
		afterPendingTop,
		topPending,
		bisectionCount,
	)
}

func ChallengePendingTopClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *structures.BlockId,
	startLogIndex uint,
	pendingInbox *structures.MessageStack,
) (ChallengeState, error) {
	contractWatcher, err := client.NewPendingTopChallengeWatcher(address)
	if err != nil {
		return 0, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockId, startLogIndex, contractWatcher)

	contract, err := client.NewPendingTopChallenge(address)
	if err != nil {
		return 0, err
	}
	log.Println("=======> challenging pending top claim")
	return challengePendingTop(
		reorgCtx,
		eventChan,
		contract,
		client,
		pendingInbox,
	)
}

func defendPendingTop(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.PendingTopChallenge,
	client arbbridge.ArbClient,
	pendingInbox *structures.MessageStack,
	afterPendingTop common.Hash,
	topPending common.Hash,
	bisectionCount uint64,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("PendingTopChallenge defender expected InitiateChallengeEvent but got %T", event)
	}

	startState := afterPendingTop
	endState := topPending

	for {
		messageCount, err := pendingInbox.SegmentSize(startState, endState)
		if err != nil {
			return 0, err
		}

		if messageCount == 1 {
			timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
			if timedOut {
				nextHash, valueHash, err := pendingInbox.GenerateOneStepProof(startState)
				if err != nil {
					return 0, err
				}
				err = contract.OneStepProof(ctx, startState, nextHash, valueHash)
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
				return 0, fmt.Errorf("PendingTopChallenge defender expected OneStepProof but got %T", event)
			}
			return ChallengeAsserterWon, nil
		}

		timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
		if timedOut {
			chainHashes, err := pendingInbox.GenerateBisection(startState, endState, bisectionCount)
			if err != nil {
				return 0, err
			}
			err = contract.Bisect(ctx, chainHashes, new(big.Int).SetUint64(messageCount))
			if err != nil {
				return 0, err
			}
			event, state, err = getNextEvent(eventChan)
		}

		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := event.(arbbridge.PendingTopBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("PendingTopChallenge defender expected PendingTopBisectionEvent but got %T", event)
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
			return 0, fmt.Errorf("PendingTopChallenge defender expected ContinueChallengeEvent but got %T", event)
		}
		startState = ev.ChainHashes[contEv.SegmentIndex.Uint64()]
		endState = ev.ChainHashes[contEv.SegmentIndex.Uint64()+1]
	}
}

func challengePendingTop(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.PendingTopChallenge,
	client arbbridge.ArbClient,
	pendingInbox *structures.MessageStack,
) (ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("PendingTopChallenge challenger expected InitiateChallengeEvent but got %T", event)
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

		ev, ok := event.(arbbridge.PendingTopBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("PendingTopChallenge challenger expected PendingTopBisectionEvent but got %T", event)
		}

		// Wait to check if we've already chosen a segment
		timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
		if timedOut {
			err = nil
			challengedSegment, err := pendingInbox.CheckBisection(ev.ChainHashes)
			if err != nil {
				return 0, err
			}
			err = contract.ChooseSegment(ctx, uint16(challengedSegment), ev.ChainHashes, ev.TotalLength.Uint64())
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
			return 0, fmt.Errorf("PendingTopChallenge challenger expected ContinueChallengeEvent but got %T", event)
		}
		deadline = contEv.Deadline
	}
}
