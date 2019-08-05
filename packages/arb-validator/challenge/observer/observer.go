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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

func New(
	config *core.Config,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	deadline uint64,
) challenge.State {
	return waitingChallenge{
		Config:       config,
		precondition: precondition,
		assertion:    assertion,
		deadline:     deadline,
	}
}

type waitingChallenge struct {
	*core.Config
	precondition *protocol.Precondition
	assertion    *protocol.AssertionStub
	deadline     uint64
}

func (bot waitingChallenge) UpdateTime(time uint64, bridge bridge.Bridge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	// timeoutMsg := SendAsserterTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.precondition,
	//	bot.Assertion,
	//}
	return challenge.TimedOutAsserter{Config: bot.Config}, nil
}

func (bot waitingChallenge) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.BisectionEvent:
		deadline := time + bot.VMConfig.GracePeriod
		preconditions := protocol.GeneratePreconditions(bot.precondition, ev.Assertions)
		return waitingBisected{bot.Config, deadline, preconditions, ev.Assertions}, nil
	default:
		return nil, &challenge.Error{Message: "ERROR: waitingChallenge: VM state got unsynchronized"}
	}
}

type waitingBisected struct {
	*core.Config
	deadline      uint64
	preconditions []*protocol.Precondition
	assertions    []*protocol.AssertionStub
}

func (bot waitingBisected) UpdateTime(time uint64, bridge bridge.Bridge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	// msg := SendChallengerTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.preconditions,
	//	bot.assertions,
	//}
	return challenge.TimedOutChallenger{Config: bot.Config}, nil
}

func (bot waitingBisected) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		deadline := time + bot.VMConfig.GracePeriod
		return waitingChallenge{
			bot.Config,
			bot.preconditions[ev.ChallengedAssertion],
			bot.assertions[ev.ChallengedAssertion],
			deadline,
		}, nil
	default:
		return nil, &challenge.Error{Message: "ERROR: waitingBisected: VM state got unsynchronized"}
	}
}
