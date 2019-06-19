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
	"github.com/offchainlabs/arb-validator/bridge"
	"github.com/offchainlabs/arb-validator/ethbridge"

	"github.com/offchainlabs/arb-avm/protocol"
)

type attemptingUnanimousClosing struct {
	*validatorConfig
	*validatorCore
	assertion *protocol.Assertion
	retChan   chan<- bool
}

func (bot attemptingUnanimousClosing) UpdateTime(time uint64, bridge bridge.Bridge) (validatorState, error) {
	return bot, nil
}

func (bot attemptingUnanimousClosing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (validatorState, challengeState, error) {
	switch ev.(type) {
	case ethbridge.ProposedUnanimousAssertEvent:
		// Someone proposed an non-final update
		// Final update has already been sent
		return bot, nil, nil
	case ethbridge.DisputableAssertionEvent:
		// Someone proposed a disputable assertion
		// Final update has already been sent
		return bot, nil, nil
	case ethbridge.FinalUnanimousAssertEvent:
		if bot.retChan != nil {
			bot.retChan <- true
		}
		bot.validatorCore.DeliverMessagesToVM()
		return newWaitingObserver(bot.validatorConfig, bot.validatorCore), nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: waitingAssertDefender: VM state got unsynchronized"}
	}
}

type attemptingOffchainClosing struct {
	*validatorConfig
	*validatorCore
	sequenceNum uint64
	assertion   *protocol.Assertion
	retChan     chan<- bool
}

func (bot attemptingOffchainClosing) UpdateTime(time uint64, bridge bridge.Bridge) (validatorState, error) {
	return bot, nil
}

func (bot attemptingOffchainClosing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (validatorState, challengeState, error) {
	switch ev := ev.(type) {
	case ethbridge.ProposedUnanimousAssertEvent:
		if ev.SequenceNum < bot.sequenceNum {
			// Someone proposed an old update
			// Newer update has already been sent
			return bot, nil, nil
		} else if ev.SequenceNum > bot.sequenceNum {
			if bot.retChan != nil {
				bot.retChan <- false
			}
			return nil, nil, errors.New("unanimous assertion unexpectedly superseded")
		} else {
			return waitingOffchainClosing{
				bot.validatorConfig,
				bot.GetCore(),
				bot.assertion,
				time + bot.config.GracePeriod,
				bot.retChan,
			}, nil, nil
		}
	case ethbridge.DisputableAssertionEvent:
		// Someone proposed a disputable assertion
		// Unanimous proposal has already been sent
		return bot, nil, nil
	case ethbridge.FinalUnanimousAssertEvent:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, errors.New("unanimous assertion unexpectedly superseded by final assert")
	default:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, &Error{nil, "ERROR: waitingAssertDefender: VM state got unsynchronized"}
	}
}

type waitingOffchainClosing struct {
	*validatorConfig
	*validatorCore
	assertion *protocol.Assertion
	deadline  uint64
	retChan   chan<- bool
}

func (bot waitingOffchainClosing) UpdateTime(time uint64, bridge bridge.Bridge) (validatorState, error) {
	if time <= bot.deadline {
		return bot, nil
	}
	bridge.ConfirmUnanimousAssertion(
		bot.validatorCore.inbox.Receive().Hash(),
		bot.assertion,
	)
	return finalizingOffchainClosing{
		validatorConfig: bot.validatorConfig,
		validatorCore:   bot.validatorCore,
		retChan:         bot.retChan,
	}, nil
}

func (bot waitingOffchainClosing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (validatorState, challengeState, error) {
	switch ev.(type) {
	case ethbridge.ProposedUnanimousAssertEvent:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, errors.New("unanimous assertion unexpectedly superseded by sequence number")
	case ethbridge.FinalUnanimousAssertEvent:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, errors.New("unanimous assertion unexpectedly superseded by final assert")
	default:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, &Error{nil, "ERROR: waitingAssertDefender: VM state got unsynchronized"}
	}
}

type finalizingOffchainClosing struct {
	*validatorConfig
	*validatorCore
	retChan chan<- bool
}

func (bot finalizingOffchainClosing) UpdateTime(time uint64, bridge bridge.Bridge) (validatorState, error) {
	return bot, nil
}

func (bot finalizingOffchainClosing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (validatorState, challengeState, error) {
	switch ev.(type) {
	case ethbridge.ConfirmedUnanimousAssertEvent:
		bot.GetCore().DeliverMessagesToVM()
		if bot.retChan != nil {
			bot.retChan <- true
		}
		return newWaitingObserver(bot.validatorConfig, bot.validatorCore), nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: waitingAssertDefender: VM state got unsynchronized"}
	}
}
