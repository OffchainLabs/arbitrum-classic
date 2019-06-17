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

type WaitingChallengeObserver struct {
	*validatorConfig
	precondition *protocol.Precondition
	assertion    *protocol.AssertionStub
	deadline     uint64
}

func (bot WaitingChallengeObserver) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time > bot.deadline {
		// timeoutMsg := SendAsserterTimedOutChallengeMessage{
		//	bot.deadline,
		//	bot.precondition,
		//	bot.assertion,
		//}
		return WaitingAsserterTimeoutObserver{bot.validatorConfig}, nil, nil
	} else {
		return bot, nil, nil
	}
}

func (bot WaitingChallengeObserver) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case ethbridge.BisectionEvent:
		deadline := time + bot.config.GracePeriod
		preconditions := protocol.GeneratePreconditions(bot.precondition, ev.Assertions)
		return WaitingBisectedObserver{bot.validatorConfig, deadline, preconditions, ev.Assertions}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: WaitingChallengeObserver: VM state got unsynchronized"}
	}
}

type WaitingBisectedObserver struct {
	*validatorConfig
	deadline      uint64
	preconditions []*protocol.Precondition
	assertions    []*protocol.AssertionStub
}

func (bot WaitingBisectedObserver) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	if time > bot.deadline {
		// msg := SendChallengerTimedOutChallengeMessage{
		//	bot.deadline,
		//	bot.preconditions,
		//	bot.assertions,
		//}
		return WaitingChallengerTimeoutObserver{bot.validatorConfig}, nil, nil
	} else {
		return bot, nil, nil
	}
}

func (bot WaitingBisectedObserver) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case ethbridge.ContinueChallengeEvent:
		deadline := time + bot.config.GracePeriod
		return WaitingChallengeObserver{
			bot.validatorConfig,
			bot.preconditions[ev.ChallengedAssertion],
			bot.assertions[ev.ChallengedAssertion],
			deadline,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: WaitingBisectedObserver: VM state got unsynchronized"}
	}
}

type WaitingChallengerTimeoutObserver struct {
	*validatorConfig
}

func (bot WaitingChallengerTimeoutObserver) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot WaitingChallengerTimeoutObserver) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case ethbridge.ChallengerTimeoutEvent:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: WaitingChallengerTimeoutObserver: VM state got unsynchronized"}
	}
}

type WaitingAsserterTimeoutObserver struct {
	*validatorConfig
}

func (bot WaitingAsserterTimeoutObserver) UpdateTime(time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot WaitingAsserterTimeoutObserver) UpdateState(ev ethbridge.Event, time uint64) (challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case ethbridge.AsserterTimeoutEvent:
		return nil, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: WaitingAsserterTimeoutObserver: VM state got unsynchronized"}
	}
}
