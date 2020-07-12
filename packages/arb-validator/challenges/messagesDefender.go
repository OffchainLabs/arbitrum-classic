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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	errors2 "github.com/pkg/errors"
	"log"
	"math/big"
)

func DefendMessagesClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	inbox *structures.MessageStack,
	beforeInbox common.Hash,
	messageCount *big.Int,
	bisectionCount uint64,
) (ChallengeState, error) {
	contractWatcher, err := client.NewMessagesChallengeWatcher(address)
	if err != nil {
		return 0, err
	}

	reorgCtx, eventChan := arbbridge.HandleBlockchainEvents(ctx, client, startBlockId, startLogIndex, contractWatcher)

	contract, err := client.NewMessagesChallenge(address)
	if err != nil {
		return 0, err
	}

	return defendMessages(
		reorgCtx,
		eventChan,
		contract,
		client,
		inbox,
		beforeInbox,
		messageCount.Uint64(),
		bisectionCount,
	)
}

func defendMessages(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.MessagesChallenge,
	client arbbridge.ArbClient,
	inbox *structures.MessageStack,
	beforeInbox common.Hash,
	messageCount uint64,
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

	vmInbox, err := inbox.GenerateVMInbox(beforeInbox, messageCount)
	if err != nil {
		return 0, err
	}

	log.Println("Inbox", inbox)
	log.Println("VM inbox", vmInbox)

	startInbox := beforeInbox
	inboxStartCount := uint64(0)

	tuple := value.NewEmptyTuple()
	hashPreImage := tuple.GetPreImage()
	for {
		log.Println(inboxStartCount, messageCount)
		if messageCount == 1 {
			return runMsgsOneStepProof(
				ctx,
				eventChan,
				inbox,
				contract,
				startInbox,
				hashPreImage)
		}

		event, state, preImages, err := msgsDefenderUpdate(
			ctx,
			eventChan,
			contract,
			inbox,
			startInbox,
			messageCount,
			bisectionCount,
			inboxStartCount,
			vmInbox)

		if challengeEnded(state, err) {
			return state, err
		}

		bisectionEvent, ok := event.(arbbridge.MessagesBisectionEvent)
		if !ok {
			return 0, fmt.Errorf("MessagesChallenge defender expected MessagesBisectionEvent but got %T", event)
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
			return 0, fmt.Errorf("MessagesChallenge defender expected ContinueChallengeEvent but got %T", event)
		}

		startInbox, hashPreImage, inboxStartCount, messageCount = updateMsgChallengeData(
			bisectionEvent,
			continueEvent,
			preImages,
			messageCount,
			inboxStartCount)

		log.Println("messageCount", messageCount, uint64(len(bisectionEvent.ChainHashes))-1, continueEvent.SegmentIndex.Uint64())
	}
}

func updateMsgChallengeData(
	bisectionEvent arbbridge.MessagesBisectionEvent,
	continueEvent arbbridge.ContinueChallengeEvent,
	preImages []value.HashPreImage,
	messageCount uint64,
	inboxStartCount uint64,
) (common.Hash, value.HashPreImage, uint64, uint64) {
	startInbox := bisectionEvent.ChainHashes[continueEvent.SegmentIndex.Uint64()]
	hashPreImage := preImages[continueEvent.SegmentIndex.Uint64()]

	inboxStartCount += getSegmentStart(
		messageCount,
		uint64(len(bisectionEvent.ChainHashes))-1,
		continueEvent.SegmentIndex.Uint64())

	messageCount = getSegmentCount(
		messageCount,
		uint64(len(bisectionEvent.ChainHashes))-1,
		continueEvent.SegmentIndex.Uint64())

	return startInbox, hashPreImage, inboxStartCount, messageCount
}

func msgsDefenderUpdate(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.MessagesChallenge,
	inbox *structures.MessageStack,
	startInbox common.Hash,
	messageCount uint64,
	bisectionCount uint64,
	inboxStartCount uint64,
	vmInbox *structures.VMInbox,
) (arbbridge.Event, ChallengeState, []value.HashPreImage, error) {
	makeBisection, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	if makeBisection {
		chainHashes, err := inbox.GenerateBisection(startInbox, bisectionCount, messageCount)
		preImages, err := vmInbox.GenerateBisection(inboxStartCount, bisectionCount, messageCount)
		if err != nil {
			return nil, 0, nil, err
		}

		simpleHashes := make([]common.Hash, 0, len(preImages))

		for _, h := range preImages {
			simpleHashes = append(simpleHashes, h.Hash())
		}

		err = contract.Bisect(ctx, chainHashes, simpleHashes, new(big.Int).SetUint64(messageCount))
		if err != nil {
			return nil, 0, nil, errors2.Wrap(err, "failing making bisection")
		}

		event, state, err = getNextEvent(eventChan)

		return event, state, preImages, err
	}

	return event, state, make([]value.HashPreImage, 0), err
}

func runMsgsOneStepProof(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	inbox *structures.MessageStack,
	contract arbbridge.MessagesChallenge,
	startInbox common.Hash,
	hashPreImage value.HashPreImage,
) (ChallengeState, error) {
	timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	if timedOut {
		msg, err := inbox.GenerateOneStepProof(startInbox)
		if err != nil {
			return 0, err
		}
		if err := contract.OneStepProof(ctx, startInbox, hashPreImage, msg); err != nil {
			return 0, errors2.Wrap(err, "failing making one step proof")
		}
		event, state, err = getNextEvent(eventChan)
	}

	if challengeEnded(state, err) {
		return state, err
	}
	_, ok := event.(arbbridge.OneStepProofEvent)
	if !ok {
		return 0, fmt.Errorf("MessagesChallenge defender expected OneStepProof but got %T", event)
	}
	return ChallengeAsserterWon, nil
}
