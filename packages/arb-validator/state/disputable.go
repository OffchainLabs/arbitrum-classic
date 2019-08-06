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
	"math"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge/challenger"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge/defender"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge/observer"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
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

type proposedUpdate struct {
	machine     machine.Machine
	messages    *protocol.MessageQueue
	Assertion   *protocol.Assertion
	sequenceNum uint64
	NewLogCount int
}

func (p *proposedUpdate) clone() *proposedUpdate {
	return &proposedUpdate{
		machine:     p.machine.Clone(),
		messages:    p.messages.Clone(),
		Assertion:   p.Assertion,
		sequenceNum: p.sequenceNum,
	}
}

type Waiting struct {
	*core.Config

	proposed    *proposedUpdate
	accepted    *core.Core
	assertion   *protocol.Assertion
	sequenceNum uint64
	signatures  [][]byte

	timeBounds protocol.TimeBounds
	orig       *core.Core
}

func NewWaiting(config *core.Config, c *core.Core) Waiting {
	return Waiting{
		Config:      config,
		proposed:    nil,
		accepted:    nil,
		assertion:   nil,
		sequenceNum: 0,
		signatures:  nil,
		timeBounds:  protocol.TimeBounds{},
		orig:        c,
	}
}

func (bot Waiting) SlowCloseUnanimous(bridge bridge.Bridge) {
	bridge.PendingUnanimousAssert(
		context.Background(),
		bot.GetCore().GetMachine().InboxHash().Hash(),
		bot.timeBounds,
		bot.assertion,
		bot.sequenceNum,
		bot.signatures,
	)
}

func (bot Waiting) FastCloseUnanimous(bridge bridge.Bridge) {
	inboxHash := bot.GetCore().GetMachine().InboxHash()
	bridge.FinalizedUnanimousAssert(
		context.Background(),
		inboxHash.Hash(),
		bot.timeBounds,
		bot.assertion,
		bot.signatures,
	)
}

func (bot Waiting) CloseUnanimous(bridge bridge.Bridge, retChan chan<- bool) (State, error) {
	if bot.assertion == nil {
		return bot, errors.New("couldn't close since no Assertion is open")
	}

	if bot.sequenceNum == math.MaxUint64 {
		bot.FastCloseUnanimous(bridge)
		return attemptingUnanimousClosing{
				bot.Config,
				bot.GetCore(),
				bot.assertion,
				retChan,
			},
			nil
	}
	bot.SlowCloseUnanimous(bridge)
	return attemptingOffchainClosing{
			bot.Config,
			bot.GetCore(),
			bot.sequenceNum,
			bot.assertion,
			retChan,
		},
		nil
}

func (bot Waiting) AttemptAssertion(request DisputableAssertionRequest, bridge bridge.Bridge) State {
	bridge.PendingDisputableAssert(
		context.Background(),
		request.Precondition,
		request.Assertion,
	)

	return attemptingAssertion{
		bot.GetCore(),
		bot.GetConfig(),
		request,
	}
}

func (bot Waiting) GetCore() *core.Core {
	if bot.assertion != nil {
		return bot.accepted
	}
	return bot.orig
}

func (bot Waiting) SendMessageToVM(msg protocol.Message) {
	bot.orig.SendMessageToVM(msg)
	if bot.proposed != nil {
		bot.proposed.machine.SendOnchainMessage(msg)
	}
	if bot.accepted != nil {
		bot.accepted.SendMessageToVM(msg)
	}
}

func (bot Waiting) ProposalResults() valmessage.UnanimousUpdateResults {
	return valmessage.UnanimousUpdateResults{
		SequenceNum:       bot.proposed.sequenceNum,
		BeforeHash:        bot.orig.GetMachine().Hash(),
		TimeBounds:        bot.timeBounds,
		NewInboxHash:      bot.proposed.machine.InboxHash().Hash(),
		OriginalInboxHash: bot.orig.GetMachine().InboxHash().Hash(),
		Assertion:         bot.proposed.Assertion,
	}
}

func (bot Waiting) Clone() Waiting {
	return Waiting{
		Config:      bot.Config,
		proposed:    bot.proposed.clone(),
		accepted:    bot.accepted.Clone(),
		assertion:   bot.assertion,
		sequenceNum: bot.sequenceNum,
		signatures:  bot.signatures,
		timeBounds:  bot.timeBounds,
		orig:        bot.orig,
	}
}

func (bot Waiting) OffchainContext(
	timeBounds protocol.TimeBounds,
	final bool,
) (protocol.TimeBounds, uint64) {
	var tb protocol.TimeBounds
	var seqNum uint64
	if bot.accepted != nil {
		tb = bot.timeBounds
		seqNum = bot.sequenceNum + 1
	} else {
		tb = timeBounds
		seqNum = 0
	}

	if final {
		seqNum = math.MaxUint64
	}

	return tb, seqNum
}

func (bot Waiting) ValidateUnanimousRequest(request valmessage.UnanimousRequestData) error {
	c := bot.GetCore()

	if request.BeforeHash != c.GetMachine().Hash() {
		return errors.New("recieved unanimous request with invalid before hash")
	}
	if request.BeforeInbox != c.GetMachine().InboxHash().Hash() {
		return errors.New("recieved unanimous request with invalid before inbox")
	}

	var tb protocol.TimeBounds
	var seqNum uint64
	if bot.accepted != nil {
		tb = bot.timeBounds
		seqNum = bot.sequenceNum + 1
	} else {
		tb = request.TimeBounds
		seqNum = 0
	}

	if request.TimeBounds != tb {
		return errors.New("recieved unanimous request with invalid timebounds")
	}

	if request.SequenceNum < seqNum {
		return errors.New("recieved unanimous request with invalid sequence number")
	}
	return nil
}

func (bot Waiting) PreparePendingUnanimous(request UnanimousUpdateRequest) (Waiting, error) {
	assertion := request.Assertion
	newLogCount := len(assertion.Logs)
	if bot.assertion != nil {
		assertion.NumSteps += bot.assertion.NumSteps
		assertion.OutMsgs = append(bot.assertion.OutMsgs, assertion.OutMsgs...)
		assertion.Logs = append(bot.assertion.Logs, assertion.Logs...)
	}

	mq := protocol.NewMessageQueue()
	for _, msg := range request.NewMessages {
		mq.AddMessage(msg)
	}

	return Waiting{
		Config: bot.Config,
		proposed: &proposedUpdate{
			machine:     request.Machine,
			messages:    mq,
			Assertion:   assertion,
			sequenceNum: request.SequenceNum,
			NewLogCount: newLogCount,
		},
		accepted:    bot.accepted,
		assertion:   bot.assertion,
		sequenceNum: bot.sequenceNum,
		signatures:  bot.signatures,
		timeBounds:  request.TimeBounds,
		orig:        bot.orig,
	}, nil
}

func (bot Waiting) FinalizePendingUnanimous(sequenceNum uint64, signatures [][]byte) (State, *proposedUpdate, error) {
	if bot.proposed == nil {
		return nil, nil, errors.New("no pending Assertion")
	}

	balance := bot.GetCore().GetBalance()
	_ = balance.SpendAll(protocol.NewBalanceTrackerFromMessages(bot.proposed.Assertion.OutMsgs))

	return Waiting{
		Config:   bot.Config,
		proposed: nil,
		accepted: core.NewCore(
			bot.proposed.machine,
			balance,
		),
		assertion:   bot.proposed.Assertion,
		sequenceNum: sequenceNum,
		signatures:  signatures,
		timeBounds:  bot.timeBounds,
		orig:        bot.orig,
	}, bot.proposed, nil
}

func (bot Waiting) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot Waiting) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.FinalizedUnanimousAssertEvent:
		if bot.sequenceNum != math.MaxUint64 {
			return nil, nil, errors.New("waiting observer saw signed final unanimous proposal that it doesn't remember")
		}
		c := bot.GetCore()
		c.DeliverMessagesToVM()
		return NewWaiting(bot.Config, c), nil, nil
	case ethbridge.PendingUnanimousAssertEvent:
		if bot.accepted == nil || ev.SequenceNum > bot.sequenceNum {
			return nil, nil, errors.New("waiting observer saw pending unanimous assertion that it doesn't remember")
		} else if ev.SequenceNum < bot.sequenceNum {
			newBot, err := bot.CloseUnanimous(bridge, nil)
			return newBot, nil, err
		} else {
			return waitingOffchainClosing{
				bot.Config,
				bot.GetCore(),
				bot.assertion,
				time + bot.VMConfig.GracePeriod,
				nil,
			}, nil, nil
		}
	case ethbridge.PendingDisputableAssertionEvent:
		if bot.accepted != nil {
			newBot, err := bot.CloseUnanimous(bridge, nil)
			return newBot, nil, err
		}

		c := bot.GetCore()
		deadline := time + bot.VMConfig.GracePeriod
		var inboxVal value.Value
		if c.GetMachine().InboxHash().Hash() != ev.Precondition.BeforeInbox.Hash() {
			return nil, nil, errors.New("waiting observer has incorrect valmessage")
		}
		updatedState := c.GetMachine().Clone()
		assertion := updatedState.ExecuteAssertion(
			int32(ev.Assertion.NumSteps),
			ev.Precondition.TimeBounds,
		)
		if !assertion.Stub().Equals(ev.Assertion) || bot.ChallengeEverything {
			bridge.InitiateChallenge(
				context.Background(),
				ev.Precondition,
				ev.Assertion,
			)
		}
		return watchingAssertion{
			c,
			bot.Config,
			inboxVal,
			updatedState,
			deadline,
			ev.Precondition,
			assertion,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, fmt.Sprintf("ERROR: Waiting: VM state got unsynchronized with valmessage %T", ev)}
	}
}

type watchingAssertion struct {
	*core.Core
	*core.Config
	inboxVal     value.Value
	pendingState machine.Machine
	deadline     uint64
	precondition *protocol.Precondition
	assertion    *protocol.Assertion
}

func (bot watchingAssertion) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	balance := bot.GetBalance()
	_ = balance.SpendAll(protocol.NewBalanceTrackerFromMessages(bot.assertion.OutMsgs))

	return finalizingAssertion{
		Core:       core.NewCore(bot.pendingState, balance),
		Config:     bot.Config,
		ResultChan: nil,
		assertion:  bot.assertion,
	}, nil
}

func (bot watchingAssertion) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.InitiateChallengeEvent:
		deadline := time + bot.VMConfig.GracePeriod
		var challengeState challenge.State
		if ev.Challenger == bot.Address {
			challengeState = challenger.New(
				bot.Config,
				bot.precondition,
				bot.assertion.Stub(),
				bot.GetMachine().Clone(),
				deadline,
			)
		} else {
			challengeState = observer.New(
				bot.Config,
				bot.precondition,
				bot.assertion.Stub(),
				deadline,
			)
		}
		return NewWaiting(bot.Config, bot.Core), challengeState, nil

	default:
		return nil, nil, &Error{nil, "ERROR: WaitingValidObserver: VM state got unsynchronized"}
	}
}

type attemptingAssertion struct {
	*core.Core
	*core.Config
	request DisputableAssertionRequest
}

func (bot attemptingAssertion) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot attemptingAssertion) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.PendingDisputableAssertionEvent:
		if ev.Asserter != bot.Address {
			fmt.Println("attemptingAssertion: Other Assertion got in before ours")
			return NewWaiting(bot.Config, bot.Core).UpdateState(ev, time, bridge)
		}

		deadline := time + bot.VMConfig.GracePeriod
		return waitingAssertion{
			bot.Core,
			bot.Config,
			bot.request,
			deadline,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: attemptingAssertion: VM state got unsynchronized"}
	}
}

type waitingAssertion struct {
	*core.Core
	*core.Config
	request  DisputableAssertionRequest
	deadline uint64
}

func (bot waitingAssertion) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	bridge.ConfirmDisputableAsserted(
		context.Background(),
		bot.request.Precondition,
		bot.request.Assertion,
	)
	assertion := bot.request.Assertion
	balance := bot.GetBalance()
	_ = balance.SpendAll(protocol.NewBalanceTrackerFromMessages(assertion.OutMsgs))
	return finalizingAssertion{
		Core: core.NewCore(
			bot.request.AfterState,
			balance,
		),
		Config:     bot.Config,
		ResultChan: bot.request.ResultChan,
		assertion:  assertion,
	}, nil
}

func (bot waitingAssertion) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev.(type) {
	case ethbridge.InitiateChallengeEvent:
		bot.request.NotifyInvalid()
		ct, err := defender.New(
			bot.Config,
			machine.NewAssertionDefender(
				bot.request.Assertion,
				bot.request.Precondition,
				bot.GetMachine().Clone(),
			),
			time,
			bridge,
		)
		return NewWaiting(bot.Config, bot.Core), ct, err

	default:
		return nil, nil, &Error{nil, "ERROR: waitingAssertion: VM state got unsynchronized"}
	}
}

type finalizingAssertion struct {
	*core.Core
	*core.Config
	ResultChan chan<- bool
	assertion  *protocol.Assertion
}

func (bot finalizingAssertion) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot finalizingAssertion) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.ConfirmedDisputableAssertEvent:
		if bot.ResultChan != nil {
			bot.ResultChan <- true
		}
		bridge.FinalizedAssertion(
			bot.assertion,
			len(bot.assertion.Logs),
			[][]byte{},
			nil,
			ev.TxHash[:],
		)
		bot.GetCore().DeliverMessagesToVM()
		return NewWaiting(bot.Config, bot.Core), nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: FinalizingAssertDefender: VM state got unsynchronized"}
	}
}
