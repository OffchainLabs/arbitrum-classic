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
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type ChallengeState uint8

const (
	ChallengeContinuing PendingTopState = iota
	ChallengeAsserterWon
	ChallengeAsserterTimedOut
	ChallengeChallengerTimedOut
)

var challengeNoEvents = errors.New("PendingTopChallengeContract notification channel terminated unexpectedly")

func handleNextEvent(note ethbridge.Notification) (outNote ethbridge.Notification, state PendingTopState, err error) {
	switch note.Event.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return note, ChallengeAsserterTimedOut, nil
	case ethbridge.ChallengerTimeoutEvent:
		return note, ChallengeChallengerTimedOut, nil
	}
	return note, 0, nil
}

func getNextEvent(outChan chan ethbridge.Notification) (note ethbridge.Notification, state PendingTopState, err error) {
	note, ok := <-outChan
	if !ok {
		return note, 0, challengeNoEvents
	}
	return handleNextEvent(note)
}

func getNextEventWithTimeout(
	auth *bind.TransactOpts,
	outChan chan ethbridge.Notification,
	deadline *big.Int,
	contract ethbridge.ChallengeContract,
	client *ethclient.Client,
) (note ethbridge.Notification, state PendingTopState, err error) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			header, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				return note, 0, err
			}
			if header.Number.Cmp(deadline) >= 0 {
				_, err := contract.TimeoutChallenge(auth)
				if err != nil {
					return note, 0, err
				}
				ticker.Stop()
			}
		case note, ok := <-outChan:
			if !ok {
				return note, 0, challengeNoEvents
			}
			return handleNextEvent(note)
		}
	}
}
