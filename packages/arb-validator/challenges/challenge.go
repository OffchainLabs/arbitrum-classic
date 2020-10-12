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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type ChallengeState uint8

const (
	ChallengeContinuing ChallengeState = iota
	ChallengeAsserterWon
	ChallengeAsserterTimedOut
	ChallengeChallengerTimedOut
	ChallengerDiscontinued
	DefenderDiscontinued
)

var replayTimeout = time.Second

var challengeNoEvents = errors.New("challenge event channel terminated unexpectedly")

func getSegmentCount(count, segments, index uint64) uint64 {
	if index == 0 {
		return count/segments + count%segments
	} else {
		return count / segments
	}
}

func getSegmentStart(count, segments, index uint64) uint64 {
	start := uint64(0)
	for i := uint64(0); i < index; i++ {
		start += getSegmentCount(count, segments, 0)
	}
	return start
}

func getAfterState(event arbbridge.Event) ChallengeState {
	switch event.(type) {
	case arbbridge.AsserterTimeoutEvent:
		return ChallengeAsserterTimedOut
	case arbbridge.ChallengerTimeoutEvent:
		return ChallengeChallengerTimedOut
	}
	return ChallengeContinuing
}

func getNextEvent(eventChan <-chan arbbridge.Event) (arbbridge.Event, ChallengeState, error) {
	event, ok := <-eventChan
	if !ok {
		return nil, 0, challengeNoEvents
	}
	return event, getAfterState(event), nil
}

func getNextEventWithTimeout(
	ctx context.Context,
	eventChan <-chan arbbridge.Event,
	deadline common.TimeTicks,
	contract arbbridge.Challenge,
	client arbbridge.ArbClient,
) (arbbridge.Event, ChallengeState, error) {
	ticker := time.NewTicker(common.NewTimeBlocksInt(2).Duration())
	for {
		select {
		case <-ctx.Done():
			return nil, 0, errors.New("context cancelled while waiting for event")
		case <-ticker.C:
			blockId, err := client.BlockIdForHeight(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			if common.TicksFromBlockNum(blockId.Height).Cmp(deadline) > 0 {
				err := contract.TimeoutChallenge(ctx)
				if err != nil {
					return nil, 0, err
				}
				ticker.Stop()
			}
		case event, ok := <-eventChan:
			if !ok {
				return nil, 0, challengeNoEvents
			}
			return event, getAfterState(event), nil
		}
	}
}

func getNextEventIfExists(ctx context.Context, eventChan <-chan arbbridge.Event, timeout time.Duration) (bool, arbbridge.Event, ChallengeState, error) {
	for {
		select {
		case event, ok := <-eventChan:
			if !ok {
				return false, nil, ChallengeContinuing, challengeNoEvents
			} else {
				return false, event, getAfterState(event), nil
			}
		case <-time.After(timeout):
			return true, nil, ChallengeContinuing, nil
		case <-ctx.Done():
			return false, nil, ChallengeContinuing, errors.New("context cancelled while waiting for event")
		}
	}
}
