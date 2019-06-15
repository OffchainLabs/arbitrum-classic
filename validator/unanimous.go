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
	"errors"
	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-validator/valmessage"
)

type AttemptingUnanimousClosing struct {
	*validatorConfig
	*validatorCore
	assertion *protocol.Assertion
	retChan   chan<- bool
}

func (bot AttemptingUnanimousClosing) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot AttemptingUnanimousClosing) UpdateState(ev valmessage.IncomingMessage, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case valmessage.ProposedUnanimousAssertMessage:
		// Someone proposed an non-final update
		// Final update has already been sent
		return bot, nil, nil, nil
	case valmessage.AssertMessage:
		// Someone proposed a disputable assertion
		// Final update has already been sent
		return bot, nil, nil, nil
	case valmessage.FinalUnanimousAssertMessage:
		if bot.retChan != nil {
			bot.retChan <- true
		}
		bot.validatorCore.DeliverMessagesToVM()
		return NewWaitingObserver(bot.validatorConfig, bot.validatorCore), nil, nil, nil
	default:
		return nil, nil, nil, &Error{nil, "ERROR: WaitingAssertDefender: VM state got unsynchronized"}
	}
}

type AttemptingOffchainClosing struct {
	*validatorConfig
	*validatorCore
	sequenceNum uint64
	assertion   *protocol.Assertion
	retChan     chan<- bool
}

func (bot AttemptingOffchainClosing) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot AttemptingOffchainClosing) UpdateState(ev valmessage.IncomingMessage, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case valmessage.ProposedUnanimousAssertMessage:
		if ev.SequenceNum < bot.sequenceNum {
			// Someone proposed an old update
			// Newer update has already been sent
			return bot, nil, nil, nil
		} else if ev.SequenceNum > bot.sequenceNum {
			if bot.retChan != nil {
				bot.retChan <- false
			}
			return nil, nil, nil, errors.New("Unanimous assertion unexpectedly superseded")
		} else {
			return WaitingOffchainClosing{
				bot.validatorConfig,
				bot.GetCore(),
				bot.assertion,
				time + bot.config.GracePeriod,
				bot.retChan,
			}, nil, nil, nil
		}
	case valmessage.AssertMessage:
		// Someone proposed a disputable assertion
		// Unanimous proposal has already been sent
		return bot, nil, nil, nil
	case valmessage.FinalUnanimousAssertMessage:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, nil, errors.New("Unanimous assertion unexpectedly superseded by final assert")
	default:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, nil, &Error{nil, "ERROR: WaitingAssertDefender: VM state got unsynchronized"}
	}
}

type WaitingOffchainClosing struct {
	*validatorConfig
	*validatorCore
	assertion *protocol.Assertion
	deadline  uint64
	retChan   chan<- bool
}

func (bot WaitingOffchainClosing) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	if time > bot.deadline {
		return FinalizingOffchainClosing{
				validatorConfig: bot.validatorConfig,
				validatorCore:   bot.validatorCore,
				retChan:         bot.retChan,
			},
			[]valmessage.OutgoingMessage{valmessage.SendConfirmUnanimousAssertedMessage{
				NewInboxHash: bot.validatorCore.inbox.Receive().Hash(),
				Assertion:    bot.assertion,
			}},
			nil
	} else {
		return bot, nil, nil
	}
}

func (bot WaitingOffchainClosing) UpdateState(ev valmessage.IncomingMessage, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case valmessage.ProposedUnanimousAssertMessage:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, nil, errors.New("Unanimous assertion unexpectedly superseded by sequence number")
	case valmessage.FinalUnanimousAssertMessage:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, nil, errors.New("Unanimous assertion unexpectedly superseded by final assert")
	default:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, nil, &Error{nil, "ERROR: WaitingAssertDefender: VM state got unsynchronized"}
	}
}

type FinalizingOffchainClosing struct {
	*validatorConfig
	*validatorCore
	retChan chan<- bool
}

func (bot FinalizingOffchainClosing) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot FinalizingOffchainClosing) UpdateState(ev valmessage.IncomingMessage, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case valmessage.ConfirmedUnanimousAssertMessage:
		bot.GetCore().DeliverMessagesToVM()
		if bot.retChan != nil {
			bot.retChan <- true
		}
		return NewWaitingObserver(bot.validatorConfig, bot.validatorCore), nil, nil, nil
	default:
		return nil, nil, nil, &Error{nil, "ERROR: WaitingAssertDefender: VM state got unsynchronized"}
	}
}
