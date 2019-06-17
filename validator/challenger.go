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

package validator

import (
	"github.com/offchainlabs/arb-validator/ethbridge"
	"math/rand"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"
	"github.com/offchainlabs/arb-validator/valmessage"
)

type WaitingContinuingChallenger struct {
	*validatorConfig
	challengedPrecondition *protocol.Precondition
	challengedAssertion    *protocol.AssertionStub
	challengedInbox        value.Value
	startState             *vm.Machine
	deadline               uint64
}

func (bot WaitingContinuingChallenger) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time > bot.deadline {
		return TimedOutAsserterChallenger{bot.validatorConfig},
			[]valmessage.OutgoingMessage{valmessage.SendAsserterTimedOutChallengeMessage{
				Deadline:     bot.deadline,
				Precondition: bot.challengedPrecondition,
				Assertion:    bot.challengedAssertion,
			}}, nil
	} else {
		return bot, nil, nil
	}
}

func (bot WaitingContinuingChallenger) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case ethbridge.BisectionEvent:
		preconditions := protocol.GeneratePreconditions(bot.challengedPrecondition, ev.Assertions)
		assertionNum, machine := protocol.ChooseAssertionToChallenge(bot.startState, ev.Assertions, preconditions, bot.challengedInbox)
		if assertionNum == uint16(len(ev.Assertions)) && bot.challengeEverything {
			assertionNum = uint16(rand.Int31n(int32(len(ev.Assertions))))
			machine = bot.startState.Clone()
			for i := uint16(0); i < assertionNum; i++ {
				machine.Run(int32(ev.Assertions[i].NumSteps))
			}
		}
		deadline := time + bot.config.GracePeriod
		if assertionNum < uint16(len(ev.Assertions)) {
			return ContinuingChallenger{
					validatorConfig: bot.validatorConfig,
					challengedState: machine,
					deadline:        deadline,
					preconditions:   preconditions,
					assertions:      ev.Assertions,
					challengedInbox: bot.challengedInbox,
				},
				[]valmessage.OutgoingMessage{valmessage.SendContinueChallengeMessage{
					AssertionToChallenge: assertionNum,
					Deadline:             deadline,
					Preconditions:        preconditions,
					Assertions:           ev.Assertions,
				}}, nil
		} else {
			return nil, nil, &Error{nil, "ERROR: WaitingContinuingChallenger: Critical bug: All segments in false assertion are valid"}
		}
	case ethbridge.OneStepProofEvent:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: WaitingContinuingChallenger: VM state got unsynchronized"}
	}
}

type ContinuingChallenger struct {
	*validatorConfig
	challengedState *vm.Machine
	deadline        uint64
	preconditions   []*protocol.Precondition
	assertions      []*protocol.AssertionStub
	challengedInbox value.Value
}

func (bot ContinuingChallenger) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time > bot.deadline {
		// msg := SendChallengerTimedOutChallengeMessage{
		//	bot.deadline,
		//	bot.preconditions,
		//	bot.assertions,
		//}
		// Currently not sending a timeout valmessage if this validator timed out
		return TimedOutChallengerChallenger{bot.validatorConfig}, nil, nil
	} else {
		return bot, nil, nil
	}
}

func (bot ContinuingChallenger) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		deadline := time + bot.config.GracePeriod
		return WaitingContinuingChallenger{
			bot.validatorConfig,
			bot.preconditions[ev.ChallengedAssertion],
			bot.assertions[ev.ChallengedAssertion],
			bot.challengedInbox,
			bot.challengedState,
			deadline,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: ContinuingChallenger: VM state got unsynchronized"}
	}
}

type TimedOutAsserterChallenger struct {
	*validatorConfig
}

func (bot TimedOutAsserterChallenger) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot TimedOutAsserterChallenger) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: TimedOutAsserterChallenger: VM state got unsynchronized"}
	}
}

type TimedOutChallengerChallenger struct {
	*validatorConfig
}

func (bot TimedOutChallengerChallenger) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot TimedOutChallengerChallenger) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case ethbridge.ChallengerTimeoutEvent:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: TimedOutChallengerChallenger: VM state got unsynchronized"}
	}
}
