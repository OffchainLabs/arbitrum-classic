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

package rollup

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

func DefendMessagesClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	pendingInbox *PendingInbox,
	beforePending [32]byte,
	afterPending [32]byte,
	messagesOutput ethbridge.MessagesOutput,
) (ChallengeState, error) {
	contract, err := ethbridge.NewMessagesChallenge(address, client)
	if err != nil {
		return 0, err
	}
	ctx := context.TODO()
	noteChan := make(chan ethbridge.Notification, 1024)

	go ethbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return defendMessages(
		auth,
		client,
		contract,
		pendingInbox,
		noteChan,
		beforePending,
		afterPending,
		messagesOutput,
	)
}

func ChallengeMessagesClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	pendingInbox *PendingInbox,
	beforePending [32]byte,
	afterPending [32]byte,
) (ChallengeState, error) {
	contract, err := ethbridge.NewMessagesChallenge(address, client)
	if err != nil {
		return 0, err
	}
	ctx := context.TODO()
	noteChan := make(chan ethbridge.Notification, 1024)

	go ethbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return challengeMessages(
		auth,
		client,
		contract,
		pendingInbox,
		noteChan,
		beforePending,
		afterPending,
	)
}

func defendMessages(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	contract *ethbridge.MessagesChallenge,
	pendingInbox *PendingInbox,
	outChan chan ethbridge.Notification,
	beforePending [32]byte,
	afterPending [32]byte,
	messagesOutput ethbridge.MessagesOutput,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = note.Event.(ethbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.New("MessagesChallenge expected InitiateChallengeEvent")
	}

	messagesStack, err := pendingInbox.Substack(beforePending, afterPending)
	if err != nil {
		return 0, err
	}

	startPending := beforePending
	endPending := afterPending
	startMessages := messagesStack.hashOfRest
	endMessages := messagesOutput.ImportedMessagesSlice

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
			_, err = contract.OneStepProof(auth, startPending, pendingNextHash, startMessages, messagesNextHash, pendingValueHash)
			if err != nil {
				return 0, err
			}
			note, state, err := getNextEvent(outChan)
			if err != nil || state != ChallengeContinuing {
				return state, err
			}
			_, ok = note.Event.(ethbridge.OneStepProof)
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
		_, err = contract.Bisect(auth, chainHashes, stackHashes, new(big.Int).SetUint64(messageCount))
		if err != nil {
			return 0, err
		}

		note, state, err := getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := note.Event.(ethbridge.MessagesBisectionEvent)
		if !ok {
			return 0, errors.New("MessagesChallenge expected MessagesBisectionEvent")
		}

		note, state, err = getNextEventWithTimeout(
			auth,
			outChan,
			ev.DeadlineTicks,
			contract,
			client,
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(ethbridge.ContinueChallengeEvent)
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
	auth *bind.TransactOpts,
	client *ethclient.Client,
	contract *ethbridge.MessagesChallenge,
	pendingInbox *PendingInbox,
	outChan chan ethbridge.Notification,
	beforePending [32]byte,
	afterPending [32]byte,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := note.Event.(ethbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.New("MessagesChallenge expected InitiateChallengeEvent")
	}

	messagesStack, err := pendingInbox.Substack(beforePending, afterPending)
	if err != nil {
		return 0, err
	}

	deadline := ev.DeadlineTicks
	for {
		note, state, err := getNextEventWithTimeout(
			auth,
			outChan,
			deadline,
			contract,
			client,
		)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}

		if _, ok := note.Event.(ethbridge.OneStepProof); ok {
			return ChallengeAsserterWon, nil
		}

		ev, ok := note.Event.(ethbridge.MessagesBisectionEvent)
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

		_, err = contract.ChooseSegment(auth, uint16(maxSegment), ev.ChainHashes)
		if err != nil {
			return 0, err
		}
		note, state, err = getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(ethbridge.ContinueChallengeEvent)
		if !ok {
			return 0, errors.New("MessagesChallenge expected ContinueChallengeEvent")
		}
		deadline = contEv.DeadlineTicks
	}
}
