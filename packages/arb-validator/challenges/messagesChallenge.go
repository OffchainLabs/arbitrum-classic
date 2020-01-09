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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arb"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func DefendMessagesClaim(
	auth *bind.TransactOpts,
	client arbbridge.ArbClient,
	address common.Address,
	pendingInbox *structures.PendingInbox,
	beforePending [32]byte,
	afterPending [32]byte,
	importedMessagesSlice [32]byte,
) (ChallengeState, error) {
	contract, err := arb.NewMessagesChallenge(address, client, auth)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan arbbridge.Notification, 1024)

	go arbbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return defendMessages(
		ctx,
		noteChan,
		contract,
		pendingInbox,
		beforePending,
		afterPending,
		importedMessagesSlice,
	)
}

func ChallengeMessagesClaim(
	auth *bind.TransactOpts,
	client arbbridge.ArbClient,
	address common.Address,
	pendingInbox *structures.PendingInbox,
	beforePending [32]byte,
	afterPending [32]byte,
) (ChallengeState, error) {
	contract, err := arb.NewMessagesChallenge(address, client, auth)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan arbbridge.Notification, 1024)

	go arbbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return challengeMessages(
		ctx,
		noteChan,
		contract,
		pendingInbox,
		beforePending,
		afterPending,
	)
}

func defendMessages(
	ctx context.Context,
	outChan chan arbbridge.Notification,
	contract arbbridge.MessagesChallenge,
	pendingInbox *structures.PendingInbox,
	beforePending [32]byte,
	afterPending [32]byte,
	importedMessagesSlice [32]byte,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = note.Event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.New("MessagesChallenge expected InitiateChallengeEvent")
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
			messagesNextHash, _, err := pendingInbox.GenerateOneStepProof(startMessages)
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
			_, ok = note.Event.(arbbridge.OneStepProof)
			if !ok {
				return 0, errors.New("MessagesChallenge expected OneStepProof")
			}
			return ChallengeAsserterWon, nil
		}

		chainHashes, err := pendingInbox.GenerateBisection(startPending, endPending, 100)
		if err != nil {
			return 0, err
		}
		stackHashes, err := messagesStack.GenerateBisection(startMessages, endMessages, 100)
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
			return 0, errors.New("MessagesChallenge expected MessagesBisectionEvent")
		}

		note, state, err = getNextEventWithTimeout(
			ctx,
			outChan,
			ev.Deadline,
			contract,
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, errors.New("MessagesChallenge expected ContinueChallengeEvent")
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
	pendingInbox *structures.PendingInbox,
	beforePending [32]byte,
	afterPending [32]byte,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := note.Event.(arbbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.New("MessagesChallenge expected InitiateChallengeEvent")
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
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}

		if _, ok := note.Event.(arbbridge.OneStepProof); ok {
			return ChallengeAsserterWon, nil
		}

		ev, ok := note.Event.(arbbridge.MessagesBisectionEvent)
		if !ok {
			return 0, errors.New("MessagesChallenge expected MessagesBisectionEvent")
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

		err = contract.ChooseSegment(ctx, uint16(maxSegment), ev.ChainHashes)
		if err != nil {
			return 0, err
		}
		note, state, err = getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(arbbridge.ContinueChallengeEvent)
		if !ok {
			return 0, errors.New("MessagesChallenge expected ContinueChallengeEvent")
		}
		deadline = contEv.Deadline
	}
}
