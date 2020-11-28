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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
)

func DefendInboxTopClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	challengeAddress common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	inbox *structures.MessageStack,
	inboxTopInitial common.Hash,
	messageCount *big.Int,
	bisectionCount uint64,
) (ChallengeState, error) {
	contractWatcher, err := client.NewInboxTopChallengeWatcher(challengeAddress)
	if err != nil {
		return 0, err
	}

	reorgCtx, challengeEvent := arbbridge.HandleBlockchainEvents(ctx, client, startBlockId, startLogIndex, contractWatcher)

	contract, err := client.NewInboxTopChallenge(challengeAddress)
	if err != nil {
		return 0, err
	}
	log.Println("=======> defending inbox top claim")

	return defendInboxTop(
		reorgCtx,
		challengeEvent,
		contract,
		client,
		inbox,
		inboxTopInitial,
		messageCount.Uint64(),
		bisectionCount,
	)
}

func defendInboxTop(
	ctx context.Context,
	challengeEvent <-chan arbbridge.Event,
	contract arbbridge.InboxTopChallenge,
	client arbbridge.ArbClient,
	inbox *structures.MessageStack,
	inboxTopInitial common.Hash,
	messageCount uint64,
	bisectionCount uint64,
) (ChallengeState, error) {
	event, ok := <-challengeEvent
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, fmt.Errorf("InboxTopChallenge defender expected InitiateChallengeEvent but got %T", event)
	}

	currentStartState := inboxTopInitial
	for {
		if messageCount == 1 {
			return runInboxOneStepProof(
				ctx,
				challengeEvent,
				currentStartState,
				inbox,
				contract)
		}

		event, state, err := inboxDefenderUpdate(
			ctx,
			challengeEvent,
			contract,
			inbox,
			currentStartState,
			messageCount,
			bisectionCount)

		if challengeEnded(state, err) {
			return state, err
		}

		bisectionEvent, ok := event.(arbbridge.InboxTopBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("InboxTopChallenge defender expected InboxTopBisectionEvent but got %T", event)
		}

		// get challenger update
		event, state, err = getNextEventWithTimeout(
			ctx,
			challengeEvent,
			bisectionEvent.Deadline,
			contract,
			client,
		)

		if challengeEnded(state, err) {
			return state, err
		}

		challengeContEvent, ok := event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, fmt.Errorf("InboxTopChallenge defender expected ContinueChallengeEvent but got %T", event)
		}

		currentStartState, messageCount = updateInboxChallengeData(challengeContEvent, bisectionEvent, messageCount)
	}
}

func updateInboxChallengeData(
	challengeContEvent arbbridge.ContinueChallengeEvent,
	bisectionEvent arbbridge.InboxTopBisectionEvent,
	messageCount uint64,
) (common.Hash, uint64) {
	bisectionIndex := challengeContEvent.SegmentIndex.Uint64()
	currentStartState := bisectionEvent.ChainHashes[bisectionIndex]
	messageCount = getSegmentCount(messageCount, uint64(len(bisectionEvent.ChainHashes))-1, bisectionIndex)

	return currentStartState, messageCount
}

func inboxDefenderUpdate(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.InboxTopChallenge,
	inbox *structures.MessageStack,
	currentStartState common.Hash,
	messageCount uint64,
	bisectionCount uint64,
) (arbbridge.Event, ChallengeState, error) {
	// Wait to check if we've already committed bisection
	makeTransaction, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	if err != nil {
		return nil, state, err
	}

	if makeTransaction {
		chainHashes, err := inbox.GenerateBisection(currentStartState, bisectionCount, messageCount)
		if err != nil {
			return nil, 0, err
		}
		err = contract.Bisect(ctx, chainHashes, new(big.Int).SetUint64(messageCount))
		if err != nil {
			return nil, 0, errors2.WithStack(errors2.Wrap(err, "Error bisecting"))
		}
		event, state, err = getNextEvent(eventChan)
		if err != nil {
			return nil, state, err
		}
	}

	return event, state, err
}

func runInboxOneStepProof(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	currentStartState common.Hash,
	inbox *structures.MessageStack,
	contract arbbridge.InboxTopChallenge,
) (ChallengeState, error) {
	timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	if timedOut {
		msg, err := inbox.InboxMessageAfter(currentStartState)
		if err != nil {
			return 0, err
		}
		err = contract.OneStepProof(ctx, currentStartState, msg.CommitmentHash())
		if err != nil {
			return 0, errors2.WithStack(errors2.Wrap(err, "Error making one step proof"))
		}
		event, state, err = getNextEvent(eventChan)
		if err != nil {
			return 0, errors2.WithStack(errors2.Wrap(err, "Error getting next event"))
		}
	}

	if challengeEnded(state, err) {
		return state, err
	}

	_, ok := event.(arbbridge.OneStepProofEvent)
	if !ok {
		return 0, fmt.Errorf("InboxTopChallenge defender expected OneStepProof but got %T", event)
	}
	return ChallengeAsserterWon, nil
}
