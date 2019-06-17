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

type waitingContinuingChallenger struct {
	*validatorConfig
	challengedPrecondition *protocol.Precondition
	challengedAssertion    *protocol.AssertionStub
	challengedInbox        value.Value
	startState             *vm.Machine
	deadline               uint64
}

func (bot waitingContinuingChallenger) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time <= bot.deadline {
		return bot, nil, nil
	}

	return timedOutAsserterChallenger{bot.validatorConfig},
		[]valmessage.OutgoingMessage{valmessage.SendAsserterTimedOutChallengeMessage{
			Deadline:     bot.deadline,
			Precondition: bot.challengedPrecondition,
			Assertion:    bot.challengedAssertion,
		}}, nil
}

func (bot waitingContinuingChallenger) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
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
		if assertionNum >= uint16(len(ev.Assertions)) {
			return nil, nil, &Error{nil, "ERROR: waitingContinuingChallenger: Critical bug: All segments in false assertion are valid"}

		}
		deadline := time + bot.config.GracePeriod
		return continuingChallenger{
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
	case ethbridge.OneStepProofEvent:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: waitingContinuingChallenger: VM state got unsynchronized"}
	}
}

type continuingChallenger struct {
	*validatorConfig
	challengedState *vm.Machine
	deadline        uint64
	preconditions   []*protocol.Precondition
	assertions      []*protocol.AssertionStub
	challengedInbox value.Value
}

func (bot continuingChallenger) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time <= bot.deadline {
		return bot, nil, nil
	}

	// msg := SendChallengerTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.preconditions,
	//	bot.assertions,
	//}
	// Currently not sending a timeout valmessage if this validator timed out
	return timedOutChallengerChallenger{bot.validatorConfig}, nil, nil
}

func (bot continuingChallenger) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		deadline := time + bot.config.GracePeriod
		return waitingContinuingChallenger{
			bot.validatorConfig,
			bot.preconditions[ev.ChallengedAssertion],
			bot.assertions[ev.ChallengedAssertion],
			bot.challengedInbox,
			bot.challengedState,
			deadline,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: continuingChallenger: VM state got unsynchronized"}
	}
}

type timedOutAsserterChallenger struct {
	*validatorConfig
}

func (bot timedOutAsserterChallenger) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot timedOutAsserterChallenger) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: timedOutAsserterChallenger: VM state got unsynchronized"}
	}
}

type timedOutChallengerChallenger struct {
	*validatorConfig
}

func (bot timedOutChallengerChallenger) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot timedOutChallengerChallenger) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case ethbridge.ChallengerTimeoutEvent:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: timedOutChallengerChallenger: VM state got unsynchronized"}
	}
}
