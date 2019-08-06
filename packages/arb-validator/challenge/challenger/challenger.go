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

package challenger

import (
	"context"
	"math/rand"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
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
	machine machine.Machine,
	deadline uint64,
) challenge.State {
	return waitingContinuing{
		config,
		precondition,
		assertion,
		machine,
		deadline,
	}
}

type waitingContinuing struct {
	*core.Config
	challengedPrecondition *protocol.Precondition
	challengedAssertion    *protocol.AssertionStub
	startState             machine.Machine
	deadline               uint64
}

func (bot waitingContinuing) UpdateTime(time uint64, bridge bridge.Bridge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}
	bridge.AsserterTimedOut(
		context.Background(),
		bot.challengedPrecondition,
		bot.challengedAssertion,
		bot.deadline,
	)
	return challenge.TimedOutAsserter{Config: bot.Config}, nil
}

func (bot waitingContinuing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.BisectionEvent:
		preconditions := protocol.GeneratePreconditions(bot.challengedPrecondition, ev.Assertions)
		assertionNum, m, err := machine.ChooseAssertionToChallenge(bot.startState, ev.Assertions, preconditions)
		if err != nil && bot.ChallengeEverything {
			assertionNum = uint16(rand.Int31n(int32(len(ev.Assertions))))
			m = bot.startState
			for i := uint16(0); i < assertionNum; i++ {
				m.ExecuteAssertion(
					int32(ev.Assertions[i].NumSteps),
					preconditions[i].TimeBounds,
				)
			}
			err = nil
		}
		if err != nil {
			return nil, &challenge.Error{Message: "ERROR: waitingContinuing: Critical bug: All segments in false Assertion are valid"}
		}
		deadline := time + bot.VMConfig.GracePeriod
		bridge.ContinueChallenge(
			context.Background(),
			assertionNum,
			preconditions,
			ev.Assertions,
			deadline,
		)
		return continuing{
			Config:          bot.Config,
			challengedState: m,
			deadline:        deadline,
			preconditions:   preconditions,
			assertions:      ev.Assertions,
		}, nil
	case ethbridge.OneStepProofEvent:
		return nil, nil
	default:
		return nil, &challenge.Error{Message: "ERROR: waitingContinuing: VM state got unsynchronized"}
	}
}

type continuing struct {
	*core.Config
	challengedState machine.Machine
	deadline        uint64
	preconditions   []*protocol.Precondition
	assertions      []*protocol.AssertionStub
}

func (bot continuing) UpdateTime(time uint64, bridge bridge.Bridge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	// msg := SendChallengerTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.preconditions,
	//	bot.assertions,
	//}
	// Currently not sending a timeout valmessage if this validator timed out
	return challenge.TimedOutChallenger{Config: bot.Config}, nil
}

func (bot continuing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		deadline := time + bot.VMConfig.GracePeriod
		return waitingContinuing{
			bot.Config,
			bot.preconditions[ev.ChallengedAssertion],
			bot.assertions[ev.ChallengedAssertion],
			bot.challengedState,
			deadline,
		}, nil
	default:
		return nil, &challenge.Error{Message: "ERROR: continuing: VM state got unsynchronized"}
	}
}
