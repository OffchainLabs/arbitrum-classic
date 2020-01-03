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
	messageCount uint64,
) (asserterWon bool, err error) {
	contract, err := ethbridge.NewPendingTopChallenge(address, client)
	if err != nil {
		return false, err
	}
	ctx := context.TODO()
	noteChan := make(chan ethbridge.Notification, 1024)

	go handleBlockchainNotifications(ctx, noteChan, contract)
	return defendChallenge(
		auth,
		client,
		contract,
		pendingInbox,
		noteChan,
		pendingTopClaim,
		messageCount,
	)
}

var pendingTopNoEvents = errors.New("PendingTopChallenge notification channel terminated unexpectedly")

func handleNextEvent(note ethbridge.Notification) (outNote ethbridge.Notification, terminated bool, asserterWon bool, err error) {
	switch note.Event.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return note, true, false, nil
	case ethbridge.ChallengerTimeoutEvent:
		return note, true, true, nil
	}
	return note, false, false, nil
}

func getNextEvent(outChan chan ethbridge.Notification) (note ethbridge.Notification, terminated bool, asserterWon bool, err error) {
	note, ok := <-outChan
	if !ok {
		return note, false, false, pendingTopNoEvents
	}
	return handleNextEvent(note)
}

func getNextEventWithTimeout(
	auth *bind.TransactOpts,
	outChan chan ethbridge.Notification,
	deadline *big.Int,
	contract *ethbridge.PendingTopChallenge,
	client *ethclient.Client,
) (note ethbridge.Notification, terminated bool, asserterWon bool, err error) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			header, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				return note, false, false, err
			}
			if header.Number.Cmp(deadline) >= 0 {
				_, err := contract.TimeoutChallenge(auth)
				if err != nil {
					return note, false, false, err
				}
				ticker.Stop()
			}
		case note, ok := <-outChan:
			if !ok {
				return note, false, false, pendingTopNoEvents
			}
			return handleNextEvent(note)
		}
	}
}

func defendChallenge(
	auth *bind.TransactOpts,
	client *ethclient.Client,
	contract *ethbridge.PendingTopChallenge,
	pendingInbox *PendingInbox,
	outChan chan ethbridge.Notification,
	pendingTopClaim ethbridge.PendingTopOutput,
	messageCount uint64,
) (bool, error) {
	note, ok := <-outChan
	if !ok {
		return false, pendingTopNoEvents
	}
	_, ok = note.Event.(ethbridge.InitiateChallengeEvent)
	if !ok {
		return false, errors.New("PendingTopChallenge expected InitiateChallengeEvent")
	}

	startState := pendingTopClaim.AfterPendingTop

	for {
		if messageCount == 1 {
			nextHash, valueHash, err := pendingInbox.GenerateOneStepProof(startState)
			if err != nil {
				return false, err
			}
			_, err = contract.OneStepProof(auth, startState, nextHash, valueHash)
			if err != nil {
				return false, err
			}
			note, terminated, asserterWon, err := getNextEvent(outChan)
			if err != nil || terminated {
				return asserterWon, err
			}
			_, ok = note.Event.(ethbridge.OneStepProof)
			if !ok {
				return false, errors.New("PendingTopChallenge expected OneStepProof")
			}
			return true, nil
		}
		chainHashes, err := pendingInbox.GenerateBisection(startState, messageCount, 100)
		if err != nil {
			return false, err
		}
		_, err = contract.Bisect(auth, chainHashes, new(big.Int).SetUint64(messageCount))
		if err != nil {
			return false, err
		}

		note, terminated, asserterWon, err := getNextEvent(outChan)
		if err != nil || terminated {
			return asserterWon, err
		}
		ev, ok := note.Event.(ethbridge.PendingTopBisectionEvent)
		if !ok {
			return false, errors.New("PendingTopChallenge expected PendingTopBisectionEvent")
		}

		note, terminated, asserterWon, err = getNextEventWithTimeout(
			auth,
			outChan,
			ev.DeadlineTicks,
			contract,
			client,
		)
		if err != nil || terminated {
			return asserterWon, err
		}
		contEv, ok := note.Event.(ethbridge.ContinueChallengeEvent)
		if !ok {
			return false, errors.New("PendingTopChallenge expected ContinueChallengeEvent")
		}
		startState = chainHashes[contEv.SegmentIndex.Uint64()]
	}

}
