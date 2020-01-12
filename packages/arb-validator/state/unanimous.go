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
	"fmt"

	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
)

type attemptingUnanimousClosing struct {
	*core.Config
	*core.Core
	assertion *protocol.ExecutionAssertion
	retChan   chan<- bool
	errChan   chan<- error
}

func (bot attemptingUnanimousClosing) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot, nil
}

func (bot attemptingUnanimousClosing) ChannelUpdateState(ev arbbridge.Event, time uint64, bridge bridge.Bridge) (ChannelState, error) {
	switch ev.(type) {
	case ethbridge.PendingUnanimousAssertEvent:
		// Someone proposed a pending update
		// Final update has already been sent
		return bot, nil
	case ethbridge.PendingDisputableAssertionEvent:
		// Someone proposed a disputable ExecutionAssertion
		// Final update has already been sent
		return bot, nil
	case ethbridge.FinalizedUnanimousAssertEvent:
		bot.Core.DeliverMessagesToVM(bridge)
		if bot.retChan != nil {
			bot.retChan <- true
		}
		return NewWaiting(bot.Config, bot.Core), nil
	default:
		err := &Error{nil, "ERROR: attemptingUnanimousClosing: VM state got unsynchronized"}
		if bot.errChan != nil {
			bot.errChan <- err
		}
		return nil, err
	}
}

type attemptingOffchainClosing struct {
	*core.Config
	*core.Core
	sequenceNum uint64
	assertion   *protocol.ExecutionAssertion
	retChan     chan<- bool
	errChan     chan<- error
}

func (bot attemptingOffchainClosing) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot, nil
}

func (bot attemptingOffchainClosing) ChannelUpdateState(ev arbbridge.Event, time uint64, bridge bridge.Bridge) (ChannelState, error) {
	switch ev := ev.(type) {
	case ethbridge.PendingUnanimousAssertEvent:
		if ev.SequenceNum < bot.sequenceNum {
			// Someone proposed an old update
			// Newer update has already been sent
			fmt.Println("Someone proposed an old update")
			return bot, nil
		} else if ev.SequenceNum > bot.sequenceNum {
			err := errors.New("unanimous ExecutionAssertion unexpectedly superseded")
			if bot.errChan != nil {
				bot.errChan <- err
			}
			return nil, err
		} else {
			return waitingOffchainClosing{
				Config:    bot.Config,
				Core:      bot.GetCore(),
				assertion: bot.assertion,
				deadline:  ev.Deadline,
				retChan:   bot.retChan,
				errChan:   bot.errChan,
			}, nil
		}
	case ethbridge.PendingDisputableAssertionEvent:
		// Someone proposed a disputable ExecutionAssertion
		// Unanimous proposal has already been sent
		return bot, nil
	case ethbridge.FinalizedUnanimousAssertEvent:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, errors.New("unanimous ExecutionAssertion unexpectedly superseded by final assert")
	default:
		if bot.retChan != nil {
			bot.retChan <- false
		}
		return nil, &Error{nil, "ERROR: attemptingOffchainClosing: VM state got unsynchronized"}
	}
}

type waitingOffchainClosing struct {
	*core.Config
	*core.Core
	assertion *protocol.ExecutionAssertion
	deadline  uint64
	retChan   chan<- bool
	errChan   chan<- error
}

func (bot waitingOffchainClosing) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	if time <= bot.deadline {
		return bot, nil
	}
	_, err := bridge.ConfirmUnanimousAsserted(
		context.Background(),
		bot.Core.GetMachine().InboxHash().Hash(),
		bot.assertion,
	)

	if err != nil {
		isPending, err2 := bridge.IsPendingUnanimous(context.Background())
		if err2 != nil {
			return nil, err2
		}
		if isPending {
			return nil, err
		}
	}
	return finalizingOffchainClosing{
		Config:  bot.Config,
		Core:    bot.Core,
		retChan: bot.retChan,
	}, nil
}

func (bot waitingOffchainClosing) ChannelUpdateState(ev arbbridge.Event, time uint64, bridge bridge.Bridge) (ChannelState, error) {
	switch ev.(type) {
	case ethbridge.PendingUnanimousAssertEvent:
		err := errors.New("unanimous ExecutionAssertion unexpectedly superseded by sequence number")
		if bot.errChan != nil {
			bot.errChan <- err
		}
		return nil, err
	case ethbridge.FinalizedUnanimousAssertEvent:
		err := errors.New("unanimous ExecutionAssertion unexpectedly superseded by final assert")
		if bot.errChan != nil {
			bot.errChan <- err
		}
		return nil, err
	default:
		err := &Error{nil, "ERROR: waitingOffchainClosing: VM state got unsynchronized"}
		if bot.errChan != nil {
			bot.errChan <- err
		}
		return nil, err
	}
}

type finalizingOffchainClosing struct {
	*core.Config
	*core.Core
	retChan chan<- bool
}

func (bot finalizingOffchainClosing) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot, nil
}

func (bot finalizingOffchainClosing) ChannelUpdateState(ev arbbridge.Event, time uint64, bridge bridge.Bridge) (ChannelState, error) {
	switch ev.(type) {
	case ethbridge.ConfirmedUnanimousAssertEvent:
		bot.GetCore().DeliverMessagesToVM(bridge)
		if bot.retChan != nil {
			bot.retChan <- true
		}
		return NewWaiting(bot.Config, bot.Core), nil
	default:
		return nil, &Error{nil, "ERROR: finalizingOffchainClosing: VM state got unsynchronized"}
	}
}
