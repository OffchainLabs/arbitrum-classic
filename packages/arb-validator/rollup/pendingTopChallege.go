/*
* Copyright 2019, Offchain Labs, Inc.
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
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type PendingTopState uint8

const (
	PendingTopContinuing PendingTopState = iota
	PendingTopAsserterWon
	PendingTopAsserterTimedOut
	PendingTopChallengerTimedOut
)

func handleBlockchainNotifications(ctx context.Context, noteChan chan ethbridge.Notification, contract *ethbridge.PendingTopChallenge) {
	outChan := make(chan ethbridge.Notification, 1024)
	errChan := make(chan error, 1024)
	defer close(outChan)
	defer close(errChan)
	if err := contract.StartConnection(ctx, outChan, errChan); err != nil {
		return
	}
	for {
		hitError := false
		select {
		case <-ctx.Done():
			break
		case notification, ok := <-outChan:
			if !ok {
				hitError = true
				break
			}
			noteChan <- notification
		case <-errChan:
			hitError = true
		}

		if hitError {
			// Ignore error and try to reset connection
			for {
				if err := contract.StartConnection(ctx, outChan, errChan); err == nil {
					break
				}
				log.Println("Error: Can't connect to blockchain")
				time.Sleep(5 * time.Second)
			}
		}
	}
}

func DefendPendingTopClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	pendingInbox *PendingInbox,
	pendingTopClaim ethbridge.PendingTopOutput,
	topPending [32]byte,
) (PendingTopState, error) {
	contract, err := ethbridge.NewPendingTopChallenge(address, client)
	if err != nil {
		return PendingTopContinuing, err
	}
	ctx := context.TODO()
	noteChan := make(chan ethbridge.Notification, 1024)

	go handleBlockchainNotifications(ctx, noteChan, contract)
	return defendPendingTop(
		auth,
		client,
		contract,
		pendingInbox,
		noteChan,
		pendingTopClaim,
		topPending,
	)
}

func ChallengePendingTopClaim(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	address common.Address,
	pendingInbox *PendingInbox,
	pendingTopClaim ethbridge.PendingTopOutput,
	topPending [32]byte,
) (PendingTopState, error) {
	contract, err := ethbridge.NewPendingTopChallenge(address, client)
	if err != nil {
		return PendingTopContinuing, err
	}
	ctx := context.TODO()
	noteChan := make(chan ethbridge.Notification, 1024)

	go handleBlockchainNotifications(ctx, noteChan, contract)
	return challengePendingTop(
		auth,
		client,
		contract,
		pendingInbox,
		noteChan,
		pendingTopClaim,
		topPending,
	)
}

var pendingTopNoEvents = errors.New("PendingTopChallenge notification channel terminated unexpectedly")

func handleNextEvent(note ethbridge.Notification) (outNote ethbridge.Notification, state PendingTopState, err error) {
	switch note.Event.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return note, PendingTopAsserterTimedOut, nil
	case ethbridge.ChallengerTimeoutEvent:
		return note, PendingTopChallengerTimedOut, nil
	}
	return note, PendingTopContinuing, nil
}

func getNextEvent(outChan chan ethbridge.Notification) (note ethbridge.Notification, state PendingTopState, err error) {
	note, ok := <-outChan
	if !ok {
		return note, PendingTopContinuing, pendingTopNoEvents
	}
	return handleNextEvent(note)
}

func getNextEventWithTimeout(
	auth *bind.TransactOpts,
	outChan chan ethbridge.Notification,
	deadline *big.Int,
	contract *ethbridge.PendingTopChallenge,
	client *ethclient.Client,
) (note ethbridge.Notification, state PendingTopState, err error) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			header, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				return note, PendingTopContinuing, err
			}
			if header.Number.Cmp(deadline) >= 0 {
				_, err := contract.TimeoutChallenge(auth)
				if err != nil {
					return note, PendingTopContinuing, err
				}
				ticker.Stop()
			}
		case note, ok := <-outChan:
			if !ok {
				return note, PendingTopContinuing, pendingTopNoEvents
			}
			return handleNextEvent(note)
		}
	}
}

func defendPendingTop(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	contract *ethbridge.PendingTopChallenge,
	pendingInbox *PendingInbox,
	outChan chan ethbridge.Notification,
	pendingTopClaim ethbridge.PendingTopOutput,
	topPending [32]byte,
) (PendingTopState, error) {
	note, ok := <-outChan
	if !ok {
		return PendingTopContinuing, pendingTopNoEvents
	}
	_, ok = note.Event.(ethbridge.InitiateChallengeEvent)
	if !ok {
		return PendingTopContinuing, errors.New("PendingTopChallenge expected InitiateChallengeEvent")
	}

	startState := pendingTopClaim.AfterPendingTop
	endState := topPending

	for {
		messageCount, err := pendingInbox.SegmentSize(startState, endState)
		if err != nil {
			return PendingTopContinuing, err
		}

		if messageCount == 1 {
			nextHash, valueHash, err := pendingInbox.GenerateOneStepProof(startState)
			if err != nil {
				return PendingTopContinuing, err
			}
			_, err = contract.OneStepProof(auth, startState, nextHash, valueHash)
			if err != nil {
				return PendingTopContinuing, err
			}
			note, state, err := getNextEvent(outChan)
			if err != nil || state != PendingTopContinuing {
				return state, err
			}
			_, ok = note.Event.(ethbridge.OneStepProof)
			if !ok {
				return PendingTopContinuing, errors.New("PendingTopChallenge expected OneStepProof")
			}
			return PendingTopAsserterWon, nil
		}

		chainHashes, err := pendingInbox.GenerateBisection(startState, endState, 100)
		if err != nil {
			return PendingTopContinuing, err
		}
		_, err = contract.Bisect(auth, chainHashes, new(big.Int).SetUint64(messageCount))
		if err != nil {
			return PendingTopContinuing, err
		}

		note, state, err := getNextEvent(outChan)
		if err != nil || state != PendingTopContinuing {
			return state, err
		}
		ev, ok := note.Event.(ethbridge.PendingTopBisectionEvent)
		if !ok {
			return PendingTopContinuing, errors.New("PendingTopChallenge expected PendingTopBisectionEvent")
		}

		note, state, err = getNextEventWithTimeout(
			auth,
			outChan,
			ev.DeadlineTicks,
			contract,
			client,
		)
		if err != nil || state != PendingTopContinuing {
			return state, err
		}
		contEv, ok := note.Event.(ethbridge.ContinueChallengeEvent)
		if !ok {
			return PendingTopContinuing, errors.New("PendingTopChallenge expected ContinueChallengeEvent")
		}
		startState = chainHashes[contEv.SegmentIndex.Uint64()]
		endState = chainHashes[contEv.SegmentIndex.Uint64()+1]
	}
}

func challengePendingTop(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	contract *ethbridge.PendingTopChallenge,
	pendingInbox *PendingInbox,
	outChan chan ethbridge.Notification,
	pendingTopClaim ethbridge.PendingTopOutput,
	topPending [32]byte,
) (PendingTopState, error) {
	note, ok := <-outChan
	if !ok {
		return PendingTopContinuing, pendingTopNoEvents
	}
	ev, ok := note.Event.(ethbridge.InitiateChallengeEvent)
	if !ok {
		return PendingTopContinuing, errors.New("PendingTopChallenge expected InitiateChallengeEvent")
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
		if err != nil || state != PendingTopContinuing {
			return state, err
		}

		if _, ok := note.Event.(ethbridge.OneStepProof); ok {
			return PendingTopAsserterWon, nil
		}

		ev, ok := note.Event.(ethbridge.PendingTopBisectionEvent)
		if !ok {
			return PendingTopContinuing, errors.New("PendingTopChallenge expected PendingTopBisectionEvent")
		}
		challengedSegment, err := pendingInbox.CheckBisection(ev.ChainHashes)
		if err != nil {
			return PendingTopContinuing, err
		}
		_, err = contract.ChooseSegment(auth, uint16(challengedSegment), ev.ChainHashes)
		if err != nil {
			return PendingTopContinuing, err
		}
		note, state, err = getNextEvent(outChan)
		if err != nil || state != PendingTopContinuing {
			return state, err
		}
		contEv, ok := note.Event.(ethbridge.ContinueChallengeEvent)
		if !ok {
			return PendingTopContinuing, errors.New("PendingTopChallenge expected ContinueChallengeEvent")
		}
		deadline = contEv.DeadlineTicks
	}
}
