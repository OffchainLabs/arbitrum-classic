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
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

func DefendPendingTopClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	pendingInbox *structures.PendingInbox,
	afterPendingTop [32]byte,
	topPending [32]byte,
) (ChallengeState, error) {
	contract, err := ethbridge.NewPendingTopChallenge(address, client, auth)
	if err != nil {
		return ChallengeContinuing, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan ethbridge.Notification, 1024)

	go ethbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return defendPendingTop(
		ctx,
		noteChan,
		contract,
		pendingInbox,
		afterPendingTop,
		topPending,
	)
}

func ChallengePendingTopClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	pendingInbox *structures.PendingInbox,
) (ChallengeState, error) {
	contract, err := ethbridge.NewPendingTopChallenge(address, client, auth)
	if err != nil {
		return 0, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	noteChan := make(chan ethbridge.Notification, 1024)

	go ethbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return challengePendingTop(
		ctx,
		noteChan,
		contract,
		pendingInbox,
	)
}

func defendPendingTop(
	ctx context.Context,
	outChan chan ethbridge.Notification,
	contract *ethbridge.PendingTopChallenge,
	pendingInbox *structures.PendingInbox,
	afterPendingTop [32]byte,
	topPending [32]byte,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	_, ok = note.Event.(ethbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.New("PendingTopChallenge expected InitiateChallengeEvent")
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
			_, err = contract.OneStepProof(ctx, startState, nextHash, valueHash)
			if err != nil {
				return 0, err
			}
			note, state, err := getNextEvent(outChan)
			if err != nil || state != ChallengeContinuing {
				return state, err
			}
			_, ok = note.Event.(ethbridge.OneStepProof)
			if !ok {
				return 0, errors.New("PendingTopChallenge expected OneStepProof")
			}
			return ChallengeAsserterWon, nil
		}

		chainHashes, err := pendingInbox.GenerateBisection(startState, endState, 100)
		if err != nil {
			return 0, err
		}
		_, err = contract.Bisect(ctx, chainHashes, new(big.Int).SetUint64(messageCount))
		if err != nil {
			return 0, err
		}

		note, state, err := getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		ev, ok := note.Event.(ethbridge.PendingTopBisectionEvent)
		if !ok {
			return 0, errors.New("PendingTopChallenge expected PendingTopBisectionEvent")
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
		contEv, ok := note.Event.(ethbridge.ContinueChallengeEvent)
		if !ok {
			return 0, errors.New("PendingTopChallenge expected ContinueChallengeEvent")
		}
		startState = chainHashes[contEv.SegmentIndex.Uint64()]
		endState = chainHashes[contEv.SegmentIndex.Uint64()+1]
	}
}

func challengePendingTop(
	ctx context.Context,
	outChan chan ethbridge.Notification,
	contract *ethbridge.PendingTopChallenge,
	pendingInbox *structures.PendingInbox,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := note.Event.(ethbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.New("PendingTopChallenge expected InitiateChallengeEvent")
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

		if _, ok := note.Event.(ethbridge.OneStepProof); ok {
			return ChallengeAsserterWon, nil
		}

		ev, ok := note.Event.(ethbridge.PendingTopBisectionEvent)
		if !ok {
			return 0, errors.New("PendingTopChallenge expected PendingTopBisectionEvent")
		}
		challengedSegment, err := pendingInbox.CheckBisection(ev.ChainHashes)
		if err != nil {
			return 0, err
		}
		_, err = contract.ChooseSegment(ctx, uint16(challengedSegment), ev.ChainHashes)
		if err != nil {
			return 0, err
		}
		note, state, err = getNextEvent(outChan)
		if err != nil || state != ChallengeContinuing {
			return state, err
		}
		contEv, ok := note.Event.(ethbridge.ContinueChallengeEvent)
		if !ok {
			return 0, errors.New("PendingTopChallenge expected ContinueChallengeEvent")
		}
		deadline = contEv.Deadline
	}
}
