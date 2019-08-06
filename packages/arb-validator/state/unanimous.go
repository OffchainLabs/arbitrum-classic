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

package state

import (
	"context"
	"errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type attemptingUnanimousClosing struct {
	*core.Config
	*core.Core
	assertion *protocol.Assertion
	retChan   chan<- bool
}

func (bot attemptingUnanimousClosing) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot attemptingUnanimousClosing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev.(type) {
	case ethbridge.PendingUnanimousAssertEvent:
		// Someone proposed a pending update
		// Final update has already been sent
		return bot, nil, nil
	case ethbridge.PendingDisputableAssertionEvent:
		// Someone proposed a disputable Assertion
		// Final update has already been sent
		return bot, nil, nil
	case ethbridge.FinalizedUnanimousAssertEvent:
		if bot.retChan != nil {
			bot.retChan <- true
		}
		bot.Core.DeliverMessagesToVM()
		return NewWaiting(bot.Config, bot.Core), nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: waitingAssertion: VM state got unsynchronized"}
	}
}

type attemptingOffchainClosing struct {
	*core.Config
	*core.Core
	sequenceNum uint64
	assertion   *protocol.Assertion
	retChan     chan<- bool
}

func (bot attemptingOffchainClosing) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot attemptingOffchainClosing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.PendingUnanimousAssertEvent:
		if ev.SequenceNum < bot.sequenceNum {
			// Someone proposed an old update
			// Newer update has already been sent
			return bot, nil, nil
		} else if ev.SequenceNum > bot.sequenceNum {
			if bot.retChan != nil {
				bot.retChan <- false
			}
			return nil, nil, errors.New("unanimous Assertion unexpectedly superseded")
		} else {
			return waitingOffchainClosing{
				bot.Config,
				bot.GetCore(),
				bot.assertion,
				time + bot.VMConfig.GracePeriod,
				bot.retChan,
			}, nil, nil
		}
	case ethbridge.PendingDisputableAssertionEvent:
		// Someone proposed a disputable Assertion
		// Unanimous proposal has already been sent
		return bot, nil, nil
	case ethbridge.FinalizedUnanimousAssertEvent:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, errors.New("unanimous Assertion unexpectedly superseded by final assert")
	default:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, &Error{nil, "ERROR: waitingAssertion: VM state got unsynchronized"}
	}
}

type waitingOffchainClosing struct {
	*core.Config
	*core.Core
	assertion *protocol.Assertion
	deadline  uint64
	retChan   chan<- bool
}

func (bot waitingOffchainClosing) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	if time <= bot.deadline {
		return bot, nil
	}
	bridge.ConfirmUnanimousAsserted(
		context.Background(),
		bot.Core.GetMachine().InboxHash().Hash(),
		bot.assertion,
	)
	return finalizingOffchainClosing{
		Config:  bot.Config,
		Core:    bot.Core,
		retChan: bot.retChan,
	}, nil
}

func (bot waitingOffchainClosing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev.(type) {
	case ethbridge.PendingUnanimousAssertEvent:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, errors.New("unanimous Assertion unexpectedly superseded by sequence number")
	case ethbridge.FinalizedUnanimousAssertEvent:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, errors.New("unanimous Assertion unexpectedly superseded by final assert")
	default:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, nil, &Error{nil, "ERROR: waitingAssertion: VM state got unsynchronized"}
	}
}

type finalizingOffchainClosing struct {
	*core.Config
	*core.Core
	retChan chan<- bool
}

func (bot finalizingOffchainClosing) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot finalizingOffchainClosing) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev.(type) {
	case ethbridge.ConfirmedUnanimousAssertEvent:
		bot.GetCore().DeliverMessagesToVM()
		if bot.retChan != nil {
			bot.retChan <- true
		}
		return NewWaiting(bot.Config, bot.Core), nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: waitingAssertion: VM state got unsynchronized"}
	}
}
