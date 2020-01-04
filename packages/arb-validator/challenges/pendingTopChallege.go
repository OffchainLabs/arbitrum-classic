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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

func DefendPendingTopClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	pendingInbox *rollup.PendingInbox,
	pendingTopClaim ethbridge.PendingTopOutput,
	topPending [32]byte,
) (ChallengeState, error) {
	contract, err := ethbridge.NewPendingTopChallenge(address, client)
	if err != nil {
		return ChallengeContinuing, err
	}
	ctx := context.TODO()
	noteChan := make(chan ethbridge.Notification, 1024)

	go ethbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return defendPendingTop(
		auth,
		client,
		noteChan,
		contract,
		pendingInbox,
		pendingTopClaim,
		topPending,
	)
}

func ChallengePendingTopClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	pendingInbox *rollup.PendingInbox,
) (ChallengeState, error) {
	contract, err := ethbridge.NewPendingTopChallenge(address, client)
	if err != nil {
		return 0, err
	}
	ctx := context.TODO()
	noteChan := make(chan ethbridge.Notification, 1024)

	go ethbridge.HandleBlockchainNotifications(ctx, noteChan, contract)
	return challengePendingTop(
		auth,
		client,
		noteChan,
		contract,
		pendingInbox,
	)
}

func defendPendingTop(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	outChan chan ethbridge.Notification,
	contract *ethbridge.PendingTopChallenge,
	pendingInbox *rollup.PendingInbox,
	pendingTopClaim ethbridge.PendingTopOutput,
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

	startState := pendingTopClaim.AfterPendingTop
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
			_, err = contract.OneStepProof(auth, startState, nextHash, valueHash)
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
		_, err = contract.Bisect(auth, chainHashes, new(big.Int).SetUint64(messageCount))
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
			return 0, errors.New("PendingTopChallenge expected ContinueChallengeEvent")
		}
		startState = chainHashes[contEv.SegmentIndex.Uint64()]
		endState = chainHashes[contEv.SegmentIndex.Uint64()+1]
	}
}

func challengePendingTop(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	outChan chan ethbridge.Notification,
	contract *ethbridge.PendingTopChallenge,
	pendingInbox *rollup.PendingInbox,
) (ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return 0, challengeNoEvents
	}
	ev, ok := note.Event.(ethbridge.InitiateChallengeEvent)
	if !ok {
		return 0, errors.New("PendingTopChallenge expected InitiateChallengeEvent")
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

		ev, ok := note.Event.(ethbridge.PendingTopBisectionEvent)
		if !ok {
			return 0, errors.New("PendingTopChallenge expected PendingTopBisectionEvent")
		}
		challengedSegment, err := pendingInbox.CheckBisection(ev.ChainHashes)
		if err != nil {
			return 0, err
		}
		_, err = contract.ChooseSegment(auth, uint16(challengedSegment), ev.ChainHashes)
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
		deadline = contEv.DeadlineTicks
	}
}
