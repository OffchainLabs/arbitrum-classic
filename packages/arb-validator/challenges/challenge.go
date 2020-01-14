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
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

type ChallengeState uint8

const (
	ChallengeContinuing ChallengeState = iota
	ChallengeAsserterWon
	ChallengeAsserterTimedOut
	ChallengeChallengerTimedOut
)

var replayTimeout = time.Second

var challengeNoEvents = errors.New("PendingTopChallengeContract notification channel terminated unexpectedly")

func getAfterState(note arbbridge.Notification) ChallengeState {
	switch note.Event.(type) {
	case arbbridge.AsserterTimeoutEvent:
		return ChallengeAsserterTimedOut
	case arbbridge.ChallengerTimeoutEvent:
		return ChallengeChallengerTimedOut
	}
	return ChallengeContinuing
}

func getNextEvent(outChan chan arbbridge.Notification) (arbbridge.Notification, ChallengeState, error) {
	note, ok := <-outChan
	if !ok {
		return note, 0, challengeNoEvents
	}
	return note, getAfterState(note), nil
}

func getNextEventWithTimeout(
	ctx context.Context,
	outChan chan arbbridge.Notification,
	deadline common.TimeTicks,
	contract arbbridge.Challenge,
	client arbbridge.ArbClient,
) (arbbridge.Notification, ChallengeState, error) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return arbbridge.Notification{}, 0, errors.New("context cancelled while waiting for event")
		case <-ticker.C:
			blockId, err := client.CurrentBlockId(ctx)
			if err != nil {
				return arbbridge.Notification{}, 0, err
			}
			if common.TimeFromBlockNum(blockId.Height).Cmp(deadline) >= 0 {
				err := contract.TimeoutChallenge(ctx)
				if err != nil {
					return arbbridge.Notification{}, 0, err
				}
				ticker.Stop()
			}
		case note, ok := <-outChan:
			if !ok {
				return note, 0, challengeNoEvents
			}
			return note, getAfterState(note), nil
		}
	}
}

func getNextEventIfExists(ctx context.Context, outChan chan arbbridge.Notification, timeout time.Duration) (bool, arbbridge.Notification, ChallengeState, error) {
	for {
		select {
		case note, ok := <-outChan:
			if !ok {
				return false, note, 0, challengeNoEvents
			}
			return false, note, getAfterState(note), nil
		case <-time.After(timeout):
			return true, arbbridge.Notification{}, 0, nil
		case <-ctx.Done():
			return false, arbbridge.Notification{}, 0, challengeNoEvents
		}
	}
}
