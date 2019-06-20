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
	"github.com/offchainlabs/arb-validator/challenge"
	"github.com/offchainlabs/arb-validator/core"
	"math/rand"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"

	"github.com/offchainlabs/arb-validator/bridge"
	"github.com/offchainlabs/arb-validator/ethbridge"
)

func New(
	config *core.Config,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	inbox value.Value,
	machine *vm.Machine,
	deadline uint64,
) challenge.State {
	return waitingContinuing{
		config,
		precondition,
		assertion,
		inbox,
		machine,
		deadline,
	}
}

type waitingContinuing struct {
	*core.Config
	challengedPrecondition *protocol.Precondition
	challengedAssertion    *protocol.AssertionStub
	challengedInbox        value.Value
	startState             *vm.Machine
	deadline               uint64
}

func (bot waitingContinuing) UpdateTime(time uint64, bridge bridge.Bridge) (challenge.State, error) {
	if time <= bot.deadline {
		return bot, nil
	}
	bridge.TimeoutAsserter(
		bot.challengedPrecondition,
		bot.challengedAssertion,
		bot.deadline,
	)
	return challenge.TimedOutAsserter{bot.Config}, nil
}

func (bot waitingContinuing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.BisectionEvent:
		preconditions := protocol.GeneratePreconditions(bot.challengedPrecondition, ev.Assertions)
		assertionNum, machine := protocol.ChooseAssertionToChallenge(bot.startState, ev.Assertions, preconditions, bot.challengedInbox)
		if assertionNum == uint16(len(ev.Assertions)) && bot.ChallengeEverything {
			assertionNum = uint16(rand.Int31n(int32(len(ev.Assertions))))
			machine = bot.startState.Clone()
			for i := uint16(0); i < assertionNum; i++ {
				machine.Run(int32(ev.Assertions[i].NumSteps))
			}
		}
		if assertionNum >= uint16(len(ev.Assertions)) {
			return nil, &challenge.Error{nil, "ERROR: waitingContinuing: Critical bug: All segments in false Assertion are valid"}
		}
		deadline := time + bot.VMConfig.GracePeriod
		bridge.ContinueChallenge(
			assertionNum,
			preconditions,
			ev.Assertions,
			deadline,
		)
		return continuing{
			Config:          bot.Config,
			challengedState: machine,
			deadline:        deadline,
			preconditions:   preconditions,
			assertions:      ev.Assertions,
			challengedInbox: bot.challengedInbox,
		}, nil
	case ethbridge.OneStepProofEvent:
		return nil, nil
	default:
		return nil, &challenge.Error{nil, "ERROR: waitingContinuing: VM state got unsynchronized"}
	}
}

type continuing struct {
	*core.Config
	challengedState *vm.Machine
	deadline        uint64
	preconditions   []*protocol.Precondition
	assertions      []*protocol.AssertionStub
	challengedInbox value.Value
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
	return challenge.TimedOutChallenger{bot.Config}, nil
}

func (bot continuing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		deadline := time + bot.VMConfig.GracePeriod
		return waitingContinuing{
			bot.Config,
			bot.preconditions[ev.ChallengedAssertion],
			bot.assertions[ev.ChallengedAssertion],
			bot.challengedInbox,
			bot.challengedState,
			deadline,
		}, nil
	default:
		return nil, &challenge.Error{nil, "ERROR: continuing: VM state got unsynchronized"}
	}
}
