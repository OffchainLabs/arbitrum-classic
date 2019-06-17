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
	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-validator/ethbridge"
	"github.com/offchainlabs/arb-validator/valmessage"
)

type waitingChallengeObserver struct {
	*validatorConfig
	precondition *protocol.Precondition
	assertion    *protocol.AssertionStub
	deadline     uint64
}

func (bot waitingChallengeObserver) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time <= bot.deadline {
		return bot, nil, nil
	}

	// timeoutMsg := SendAsserterTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.precondition,
	//	bot.assertion,
	//}
	return waitingAsserterTimeoutObserver{bot.validatorConfig}, nil, nil
}

func (bot waitingChallengeObserver) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case ethbridge.BisectionEvent:
		deadline := time + bot.config.GracePeriod
		preconditions := protocol.GeneratePreconditions(bot.precondition, ev.Assertions)
		return waitingBisectedObserver{bot.validatorConfig, deadline, preconditions, ev.Assertions}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: waitingChallengeObserver: VM state got unsynchronized"}
	}
}

type waitingBisectedObserver struct {
	*validatorConfig
	deadline      uint64
	preconditions []*protocol.Precondition
	assertions    []*protocol.AssertionStub
}

func (bot waitingBisectedObserver) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time <= bot.deadline {
		return bot, nil, nil
	}

	// msg := SendChallengerTimedOutChallengeMessage{
	//	bot.deadline,
	//	bot.preconditions,
	//	bot.assertions,
	//}
	return waitingChallengerTimeoutObserver{bot.validatorConfig}, nil, nil
}

func (bot waitingBisectedObserver) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		deadline := time + bot.config.GracePeriod
		return waitingChallengeObserver{
			bot.validatorConfig,
			bot.preconditions[ev.ChallengedAssertion],
			bot.assertions[ev.ChallengedAssertion],
			deadline,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: waitingBisectedObserver: VM state got unsynchronized"}
	}
}

type waitingChallengerTimeoutObserver struct {
	*validatorConfig
}

func (bot waitingChallengerTimeoutObserver) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot waitingChallengerTimeoutObserver) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case ethbridge.ChallengerTimeoutEvent:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: waitingChallengerTimeoutObserver: VM state got unsynchronized"}
	}
}

type waitingAsserterTimeoutObserver struct {
	*validatorConfig
}

func (bot waitingAsserterTimeoutObserver) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot waitingAsserterTimeoutObserver) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: waitingAsserterTimeoutObserver: VM state got unsynchronized"}
	}
}
