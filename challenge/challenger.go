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

package challenge

import (
	"fmt"
	"github.com/offchainlabs/arb-validator/core"
	"math/rand"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"

	"github.com/offchainlabs/arb-validator/bridge"
	"github.com/offchainlabs/arb-validator/ethbridge"
)

type Error struct {
	err     error
	message string
}

func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%v: %v", e.message, e.err)
	}
	return e.message
}

type waitingContinuingChallenger struct {
	*core.Config
	challengedPrecondition *protocol.Precondition
	challengedAssertion    *protocol.AssertionStub
	challengedInbox        value.Value
	startState             *vm.Machine
	deadline               uint64
}

func NewChallenger(
	config *core.Config,
	precondition *protocol.Precondition,
	assertion *protocol.AssertionStub,
	inbox value.Value,
	machine *vm.Machine,
	deadline uint64,
) State {
	return waitingContinuingChallenger{
		config,
		precondition,
		assertion,
		inbox,
		machine,
		deadline,
	}
}

func (bot waitingContinuingChallenger) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	if time <= bot.deadline {
		return bot, nil
	}
	bridge.TimeoutAsserter(
		bot.challengedPrecondition,
		bot.challengedAssertion,
		bot.deadline,
	)
	return timedOutAsserterChallenger{bot.Config}, nil
}

func (bot waitingContinuingChallenger) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
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
			return nil, &Error{nil, "ERROR: waitingContinuingChallenger: Critical bug: All segments in false Assertion are valid"}
		}
		deadline := time + bot.VMConfig.GracePeriod
		bridge.ContinueChallenge(
			assertionNum,
			preconditions,
			ev.Assertions,
			deadline,
		)
		return continuingChallenger{
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
		return nil, &Error{nil, "ERROR: waitingContinuingChallenger: VM state got unsynchronized"}
	}
}

type continuingChallenger struct {
	*core.Config
	challengedState *vm.Machine
	deadline        uint64
	preconditions   []*protocol.Precondition
	assertions      []*protocol.AssertionStub
	challengedInbox value.Value
}

func (bot continuingChallenger) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	// msg := SendChallengerTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.preconditions,
	//	bot.assertions,
	//}
	// Currently not sending a timeout valmessage if this validator timed out
	return timedOutChallengerChallenger{bot.Config}, nil
}

func (bot continuingChallenger) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		deadline := time + bot.VMConfig.GracePeriod
		return waitingContinuingChallenger{
			bot.Config,
			bot.preconditions[ev.ChallengedAssertion],
			bot.assertions[ev.ChallengedAssertion],
			bot.challengedInbox,
			bot.challengedState,
			deadline,
		}, nil
	default:
		return nil, &Error{nil, "ERROR: continuingChallenger: VM state got unsynchronized"}
	}
}

type timedOutAsserterChallenger struct {
	*core.Config
}

func (bot timedOutAsserterChallenger) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot timedOutAsserterChallenger) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
	switch ev.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return nil, nil
	default:
		return nil, &Error{nil, "ERROR: timedOutAsserterChallenger: VM state got unsynchronized"}
	}
}

type timedOutChallengerChallenger struct {
	*core.Config
}

func (bot timedOutChallengerChallenger) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot timedOutChallengerChallenger) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, error) {
	switch ev.(type) {
	case ethbridge.ChallengerTimeoutEvent:
		return nil, nil
	default:
		return nil, &Error{nil, "ERROR: timedOutChallengerChallenger: VM state got unsynchronized"}
	}
}
