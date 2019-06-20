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
	"errors"
	"fmt"
	"github.com/offchainlabs/arb-validator/bridge"
	"github.com/offchainlabs/arb-validator/challenge"
	"github.com/offchainlabs/arb-validator/core"
	"github.com/offchainlabs/arb-validator/ethbridge"
	"log"
	"math"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"
	"github.com/offchainlabs/arb-validator/valmessage"
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
	machine     *vm.Machine
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

	proposed *proposedUpdate

	acceptedMachine  *vm.Machine
	acceptedMessages *protocol.MessageQueues
	acceptedBalance  *protocol.BalanceTracker
	assertion        *protocol.Assertion
	sequenceNum      uint64
	signatures       [][]byte

	timeBounds      protocol.TimeBounds
	pendingMessages *protocol.MessageQueue
	origMessages    *protocol.MessageQueues
	origBalance     *protocol.BalanceTracker
	origMachine     *vm.Machine
}

func NewWaiting(config *core.Config, c *core.Core) Waiting {
	return Waiting{
		Config:           config,
		proposed:         nil,
		acceptedMachine:  nil,
		acceptedMessages: nil,
		acceptedBalance:  nil,
		assertion:        nil,
		sequenceNum:      0,
		signatures:       nil,
		timeBounds:       protocol.TimeBounds{},
		pendingMessages:  c.GetInbox().PendingQueue,
		origMessages:     c.GetInbox().Accepted,
		origBalance:      c.GetBalance(),
		origMachine:      c.GetMachine(),
	}
}

func (bot Waiting) SlowCloseUnanimous(bridge bridge.Bridge) {
	bridge.UnanimousAssert(
		bot.GetCore().GetInbox().Receive().Hash(),
		bot.timeBounds,
		bot.assertion,
		bot.sequenceNum,
		bot.signatures,
	)
}

func (bot Waiting) FastCloseUnanimous(bridge bridge.Bridge) {
	inboxHash := bot.GetCore().GetInbox().Receive().Hash()
	bridge.FinalUnanimousAssert(
		inboxHash,
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
	bridge.DisputableAssert(
		request.Defender.GetPrecondition(),
		request.Defender.GetAssertion(),
	)

	return attemptingAssertion{
		bot.GetCore(),
		bot.GetConfig(),
		request,
	}
}

func (bot Waiting) OrigInbox() *protocol.Inbox {
	return protocol.NewInbox(
		bot.origMessages,
		bot.pendingMessages,
	)
}

func (bot Waiting) ProposedInbox() *protocol.Inbox {
	curInbox := bot.GetCore().GetInbox()
	updatedQueues := curInbox.Accepted.WithAddedQueue(bot.proposed.messages)
	return protocol.NewInbox(
		updatedQueues,
		bot.pendingMessages,
	)
}

func (bot Waiting) ProposalResults() valmessage.UnanimousUpdateResults {
	return valmessage.UnanimousUpdateResults{
		SequenceNum:       bot.proposed.sequenceNum,
		BeforeHash:        bot.origMachine.Hash(),
		TimeBounds:        bot.timeBounds,
		NewInboxHash:      bot.ProposedInbox().Receive().Hash(),
		OriginalInboxHash: bot.OrigInbox().Receive().Hash(),
		Assertion:         bot.proposed.Assertion,
	}
}

func (bot Waiting) Clone() Waiting {
	return Waiting{
		Config:           bot.Config,
		proposed:         bot.proposed.clone(),
		acceptedMachine:  bot.acceptedMachine.Clone(),
		acceptedMessages: bot.acceptedMessages,
		acceptedBalance:  bot.acceptedBalance.Clone(),
		assertion:        bot.assertion,
		sequenceNum:      bot.sequenceNum,
		signatures:       bot.signatures,
		timeBounds:       bot.timeBounds,
		pendingMessages:  bot.pendingMessages.Clone(),
		origMessages:     bot.origMessages,
		origBalance:      bot.origBalance,
		origMachine:      bot.origMachine,
	}
}

func (bot Waiting) GetCore() *core.Core {
	if bot.assertion != nil {
		return core.NewCore(
			protocol.NewInbox(bot.acceptedMessages, bot.pendingMessages),
			bot.acceptedBalance,
			bot.acceptedMachine,
		)
	}
	return core.NewCore(
		protocol.NewInbox(bot.origMessages, bot.pendingMessages),
		bot.origBalance,
		bot.origMachine,
		)
}

func (bot Waiting) SendMessageToVM(msg protocol.Message) {
	bot.pendingMessages.AddMessage(msg)
}

func (bot Waiting) OffchainContext(
	newMessages []protocol.Message,
	timeBounds protocol.TimeBounds,
	final bool,
) (*protocol.MessageQueue, protocol.TimeBounds, uint64) {
	var tb protocol.TimeBounds
	var seqNum uint64
	mq := protocol.NewMessageQueue()
	for _, msg := range newMessages {
		mq.AddMessage(msg)
	}
	if bot.acceptedMachine != nil {
		tb = bot.timeBounds
		seqNum = bot.sequenceNum + 1
	} else {
		tb = timeBounds
		seqNum = 0
	}

	if final {
		seqNum = math.MaxUint64
	}

	return mq, tb, seqNum
}

func (bot Waiting) ValidateUnanimousRequest(request valmessage.UnanimousRequestData) error {
	core := bot.GetCore()

	if request.BeforeHash != core.GetMachine().Hash() {
		return errors.New("recieved unanimous request with invalid before hash")
	}
	if request.BeforeInbox != core.GetInbox().Receive().Hash() {
		fmt.Println("validateUnanimousRequest", core.GetInbox().Receive())
		return errors.New("recieved unanimous request with invalid before inbox")
	}

	var tb protocol.TimeBounds
	var seqNum uint64
	if bot.acceptedMachine != nil {
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
		acceptedMachine:  bot.acceptedMachine,
		acceptedMessages: bot.acceptedMessages,
		acceptedBalance:  bot.acceptedBalance,
		assertion:        bot.assertion,
		sequenceNum:      bot.sequenceNum,
		signatures:       bot.signatures,
		timeBounds:       request.TimeBounds,
		pendingMessages:  bot.pendingMessages,
		origMessages:     bot.origMessages,
		origBalance:      bot.origBalance,
		origMachine:      bot.origMachine,
	}, nil
}

func (bot Waiting) FinalizePendingUnanimous(signatures [][]byte) (State, *proposedUpdate, error) {
	if bot.proposed == nil {
		return nil, nil, errors.New("no pending Assertion")
	}

	core := bot.GetCore()
	balance := core.GetBalance().Clone()

	// This spend is guaranteed to be correct since the VM made sure to only produce on outgoing if it could spend
	_ = balance.SpendAll(protocol.NewBalanceTrackerFromMessages(bot.proposed.Assertion.OutMsgs))

	return Waiting{
		Config:           bot.Config,
		proposed:         nil,
		acceptedMachine:  bot.proposed.machine,
		acceptedMessages: core.GetInbox().Accepted.WithAddedQueue(bot.proposed.messages),
		acceptedBalance:  balance,
		assertion:        bot.proposed.Assertion,
		sequenceNum:      bot.proposed.sequenceNum,
		signatures:       signatures,
		timeBounds:       bot.timeBounds,
		pendingMessages:  bot.pendingMessages,
		origMessages:     bot.origMessages,
		origBalance:      bot.origBalance,
		origMachine:      bot.origMachine,
	}, bot.proposed, nil
}

func (bot Waiting) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot Waiting) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.FinalUnanimousAssertEvent:
		if bot.sequenceNum != math.MaxUint64 {
			return nil, nil, errors.New("waiting observer saw signed final unanimous proposal that it doesn't remember")
		}
		core := bot.GetCore()
		core.DeliverMessagesToVM()
		return NewWaiting(bot.Config, core), nil, nil
	case ethbridge.ProposedUnanimousAssertEvent:
		if bot.acceptedMachine == nil || ev.SequenceNum > bot.sequenceNum {
			return nil, nil, errors.New("waiting observer saw signed unanimous proposal that it doesn't remember")
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
	case ethbridge.DisputableAssertionEvent:
		if bot.acceptedMachine != nil {
			newBot, err := bot.CloseUnanimous(bridge, nil)
			return newBot, nil, err
		}

		core := bot.GetCore()
		deadline := time + bot.VMConfig.GracePeriod
		inbox := core.GetInbox()
		var inboxVal value.Value
		if inbox.Receive().Hash() == ev.Precondition.BeforeInbox.Hash() {
			inboxVal = inbox.Receive()
		} else if inbox.ReceivePending().Hash() == ev.Precondition.BeforeInbox.Hash() {
			inboxVal = inbox.ReceivePending()
		} else {
			return nil, nil, errors.New("waiting observer has incorrect valmessage")
		}
		updatedState := core.GetMachine().Clone()
		actx := protocol.NewMachineAssertionContext(
			updatedState,
			core.GetBalance(),
			ev.Precondition.TimeBounds,
			inboxVal,
		)
		updatedState.Run(int32(ev.Assertion.NumSteps))
		ad := actx.Finalize(updatedState)

		if !ad.GetAssertion().Stub().Equals(ev.Assertion) || bot.ChallengeEverything {
			bridge.InitiateChallenge(
				ev.Precondition,
				ev.Assertion,
			)
		}
		return watchingAssertion{
			core,
			bot.Config,
			inboxVal,
			updatedState,
			deadline,
			ev.Precondition,
			ad.GetAssertion(),
		}, nil, nil
	default:
		return nil, nil, &Error{nil, fmt.Sprintf("ERROR: Waiting: VM state got unsynchronized with valmessage %T", ev)}
	}
}

type watchingAssertion struct {
	*core.Core
	*core.Config
	inboxVal     value.Value
	pendingState *vm.Machine
	deadline     uint64
	precondition *protocol.Precondition
	assertion    *protocol.Assertion
}

func (bot watchingAssertion) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	newBalance := bot.GetBalance().Clone()
	err := newBalance.SpendAll(bot.precondition.BeforeBalance)
	if err != nil {
		log.Fatal("Ethbridge admited Assertion with more than available balance")
	}
	bridge.FinalizedAssertion(bot.assertion, len(bot.assertion.Logs))
	return finalizingAssertion{
		Core: core.NewCore(
			bot.GetInbox(),
			newBalance,
			bot.GetMachine(),
		),
		Config:     bot.Config,
		ResultChan: nil,
	}, nil
}

func (bot watchingAssertion) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev := ev.(type) {
	case ethbridge.InitiateChallengeEvent:
		deadline := time + bot.VMConfig.GracePeriod
		var challengeState challenge.State
		if ev.Challenger == bot.Address {
			challengeState = challenge.NewChallenger(
				bot.Config,
				bot.precondition,
				bot.assertion.Stub(),
				bot.inboxVal,
				bot.GetMachine().Clone(),
				deadline,
			)
		} else {
			challengeState = challenge.NewObserver(
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
	case ethbridge.DisputableAssertionEvent:
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

	newBalance := bot.GetBalance().Clone()
	err := newBalance.SpendAll(protocol.NewBalanceTrackerFromMessages(bot.request.Defender.GetAssertion().OutMsgs))
	if err != nil {
		log.Fatal("Ethbridge admited Assertion with more than available balance")
	}
	bridge.ConfirmDisputableAssertion(
		bot.request.Defender.GetPrecondition(),
		bot.request.Defender.GetAssertion(),
	)
	assertion := bot.request.Defender.GetAssertion()
	bridge.FinalizedAssertion(assertion, len(assertion.Logs))
	return finalizingAssertion{
		core.NewCore(
			bot.GetInbox(),
			newBalance,
			bot.request.State,
		),
		bot.Config,
		bot.request.ResultChan,
	}, nil
}

func (bot waitingAssertion) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev.(type) {
	case ethbridge.InitiateChallengeEvent:
		bot.request.NotifyInvalid()
		ct, err := challenge.NewDefender(bot.Config, bot.request.Defender, time, bridge)
		return NewWaiting(bot.Config, bot.Core), ct, err

	default:
		return nil, nil, &Error{nil, "ERROR: waitingAssertion: VM state got unsynchronized"}
	}
}

type finalizingAssertion struct {
	*core.Core
	*core.Config
	ResultChan chan<- bool
}

func (bot finalizingAssertion) UpdateTime(time uint64, bridge bridge.Bridge) (State, error) {
	return bot, nil
}

func (bot finalizingAssertion) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (State, challenge.State, error) {
	switch ev.(type) {
	case ethbridge.ConfirmedAssertEvent:
		if bot.ResultChan != nil {
			bot.ResultChan <- true
		}
		bot.GetCore().DeliverMessagesToVM()
		return NewWaiting(bot.Config, bot.Core), nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: FinalizingAssertDefender: VM state got unsynchronized"}
	}
}
