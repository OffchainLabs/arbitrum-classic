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

package observer

import (
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
)

func New(
	deadline uint64,
) challenge.State {
	return waitingChallenge{
		deadline: deadline,
	}
}

type waitingChallenge struct {
	deadline uint64
}

func (bot waitingChallenge) UpdateTime(time uint64, bridge bridge.Challenge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	// timeoutMsg := SendAsserterTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.precondition,
	//	bot.Assertion,
	//}
	return challenge.TimedOutAsserter{}, nil
}

func (bot waitingChallenge) UpdateState(ev arbbridge.Event, time uint64, brdg bridge.Challenge) (challenge.State, error) {
	switch ev := ev.(type) {
	case arbbridge.ExecutionBisectionEvent:
		return waitingBisected{ev.Deadline}, nil
	default:
		return nil, &bridge.Error{Message: "ERROR: waitingChallenge: VM state got unsynchronized"}
	}
}

type waitingBisected struct {
	deadline uint64
}

func (bot waitingBisected) UpdateTime(time uint64, bridge bridge.Challenge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	// msg := SendChallengerTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.preconditions,
	//	bot.assertions,
	//}
	return challenge.TimedOutChallenger{}, nil
}

func (bot waitingBisected) UpdateState(ev arbbridge.Event, time uint64, brdg bridge.Challenge) (challenge.State, error) {
	switch ev := ev.(type) {
	case arbbridge.ContinueChallengeEvent:
		return waitingChallenge{
			ev.Deadline,
		}, nil
	default:
		return nil, &bridge.Error{Message: "ERROR: waitingBisected: VM state got unsynchronized"}
	}
}
