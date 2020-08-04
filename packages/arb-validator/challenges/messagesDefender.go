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

func DefendMessagesClaim(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	address common.Address,
	startBlockId *common.BlockId,
	startLogIndex uint,
	inbox *structures.MessageStack,
	beforeInboxTop common.Hash,
	afterInboxTop common.Hash,
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
		beforeInboxTop,
		afterInboxTop,
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
	beforeInboxTop common.Hash,
	afterInboxTop common.Hash,
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

	vmInbox, err := inbox.GenerateVMInbox(beforeInboxTop, messageCount)
	if err != nil {
		return 0, errors2.Wrap(err, "defender error generating vm inbox")
	}

	beforeGlobalInbox := afterInboxTop
	afterGlobalInbox := beforeInboxTop
	inboxStartCount := uint64(0)

	beforeInboxHash := vmInbox.Hash()
	for {
		if messageCount == 1 {
			return runMsgsOneStepProof(
				ctx,
				eventChan,
				inbox,
				contract,
				afterGlobalInbox,
				beforeInboxHash,
			)
		}

		event, state, vmInboxHashes, err := msgsDefenderUpdate(
			ctx,
			eventChan,
			contract,
			inbox,
			beforeGlobalInbox,
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

		log.Println("Chose segment", continueEvent.SegmentIndex)

		beforeGlobalInbox, afterGlobalInbox, beforeInboxHash, inboxStartCount, messageCount = updateMsgChallengeData(
			bisectionEvent,
			continueEvent,
			vmInboxHashes,
			messageCount,
			inboxStartCount,
		)
	}
}

func updateMsgChallengeData(
	bisectionEvent arbbridge.MessagesBisectionEvent,
	continueEvent arbbridge.ContinueChallengeEvent,
	vmInboxHashes []common.Hash,
	messageCount uint64,
	inboxStartCount uint64,
) (common.Hash, common.Hash, common.Hash, uint64, uint64) {
	beforeGlobalInbox := bisectionEvent.ChainHashes[continueEvent.SegmentIndex.Uint64()]
	afterGlobalInbox := bisectionEvent.ChainHashes[continueEvent.SegmentIndex.Uint64()+1]
	beforeInboxHash := vmInboxHashes[continueEvent.SegmentIndex.Uint64()]

	inboxStartCount += getSegmentStart(
		messageCount,
		uint64(len(bisectionEvent.ChainHashes))-1,
		continueEvent.SegmentIndex.Uint64())

	messageCount = getSegmentCount(
		messageCount,
		uint64(len(bisectionEvent.ChainHashes))-1,
		continueEvent.SegmentIndex.Uint64())

	return beforeGlobalInbox, afterGlobalInbox, beforeInboxHash, inboxStartCount, messageCount
}

func msgsDefenderUpdate(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	contract arbbridge.MessagesChallenge,
	inbox *structures.MessageStack,
	afterGlobalInbox common.Hash,
	messageCount uint64,
	bisectionCount uint64,
	inboxStartCount uint64,
	vmInbox *structures.VMInbox,
) (arbbridge.Event, ChallengeState, []common.Hash, error) {
	log.Println("Bisecting from", afterGlobalInbox, bisectionCount, messageCount)
	chainHashes, err := inbox.GenerateBisectionReverse(afterGlobalInbox, bisectionCount, messageCount)
	log.Println("Bisection", chainHashes)
	vmInboxHashes, err := vmInbox.GenerateBisection(inboxStartCount, bisectionCount, messageCount)
	if err != nil {
		return nil, 0, nil, err
	}

	makeBisection, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	if makeBisection {
		err = contract.Bisect(ctx, chainHashes, vmInboxHashes, new(big.Int).SetUint64(messageCount))
		if err != nil {
			return nil, 0, nil, errors2.Wrap(err, "failing making messages bisection")
		}

		event, state, err = getNextEvent(eventChan)
	}
	return event, state, vmInboxHashes, err
}

func runMsgsOneStepProof(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	inbox *structures.MessageStack,
	contract arbbridge.MessagesChallenge,
	afterGlobalInbox common.Hash,
	beforeVmInbox common.Hash,
) (ChallengeState, error) {
	timedOut, event, state, err := getNextEventIfExists(ctx, eventChan, replayTimeout)
	if timedOut {
		msg, err := inbox.InboxMessageAfter(afterGlobalInbox)
		if err != nil {
			return 0, err
		}
		if err := contract.OneStepProof(ctx, afterGlobalInbox, beforeVmInbox, msg); err != nil {
			log.Println("afterGlobalInbox", afterGlobalInbox)
			log.Println("msg", msg.AsValue().Hash())
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
