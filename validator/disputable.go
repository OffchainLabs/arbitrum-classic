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
	"fmt"
	"github.com/offchainlabs/arb-validator/ethbridge"
	"log"
	"math"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"
	"github.com/offchainlabs/arb-validator/valmessage"
)

type proposedUpdate struct {
	machine     *vm.Machine
	messages    *protocol.MessageQueue
	assertion   *protocol.Assertion
	sequenceNum uint64
	newLogCount int
}

func (p *proposedUpdate) clone() *proposedUpdate {
	return &proposedUpdate{
		machine:     p.machine.Clone(),
		messages:    p.messages.Clone(),
		assertion:   p.assertion,
		sequenceNum: p.sequenceNum,
	}
}

type waitingObserver struct {
	*validatorConfig

	proposed *proposedUpdate

	acceptedMachine  *vm.Machine
	acceptedMessages *protocol.MessageQueues
	acceptedBalance  *protocol.BalanceTracker
	assertion        *protocol.Assertion
	sequenceNum      uint64
	signatures       []valmessage.Signature

	timeBounds      protocol.TimeBounds
	pendingMessages *protocol.MessageQueue
	origMessages    *protocol.MessageQueues
	origBalance     *protocol.BalanceTracker
	origMachine     *vm.Machine
}

func newWaitingObserver(config *validatorConfig, core *validatorCore) waitingObserver {
	return waitingObserver{
		validatorConfig:  config,
		proposed:         nil,
		acceptedMachine:  nil,
		acceptedMessages: nil,
		acceptedBalance:  nil,
		assertion:        nil,
		sequenceNum:      0,
		signatures:       nil,
		timeBounds:       protocol.TimeBounds{},
		pendingMessages:  core.inbox.PendingQueue,
		origMessages:     core.inbox.Accepted,
		origBalance:      core.balance,
		origMachine:      core.machine,
	}
}

func (bot waitingObserver) SlowCloseUnanimous() valmessage.SendProposeUnanimousAssertMessage {
	return valmessage.SendProposeUnanimousAssertMessage{
		NewInboxHash: bot.GetCore().inbox.Receive().Hash(),
		TimeBounds:   bot.timeBounds,
		Assertion:    bot.assertion,
		SequenceNum:  bot.sequenceNum,
		Signatures:   bot.signatures,
	}
}

func (bot waitingObserver) FastCloseUnanimous() valmessage.SendUnanimousAssertMessage {
	inboxHash := bot.GetCore().inbox.Receive().Hash()
	return valmessage.SendUnanimousAssertMessage{
		NewInboxHash: inboxHash,
		TimeBounds:   bot.timeBounds,
		Assertion:    bot.assertion,
		Signatures:   bot.signatures,
	}
}

func (bot waitingObserver) CloseUnanimous(retChan chan<- bool) (validatorState, []valmessage.OutgoingMessage, error) {
	if bot.assertion == nil {
		return bot, nil, errors.New("couldn't close since no assertion is open")
	}

	if bot.sequenceNum == math.MaxUint64 {
		return attemptingUnanimousClosing{
				bot.validatorConfig,
				bot.GetCore(),
				bot.assertion,
				retChan,
			},
			[]valmessage.OutgoingMessage{bot.FastCloseUnanimous()},
			nil
	}
	return attemptingOffchainClosing{
			bot.validatorConfig,
			bot.GetCore(),
			bot.sequenceNum,
			bot.assertion,
			retChan,
		},
		[]valmessage.OutgoingMessage{bot.SlowCloseUnanimous()},
		nil
}

func (bot waitingObserver) OrigInbox() *protocol.Inbox {
	return protocol.NewInbox(
		bot.origMessages,
		bot.pendingMessages,
	)
}

func (bot waitingObserver) ProposedInbox() *protocol.Inbox {
	curInbox := bot.GetCore().inbox
	updatedQueues := curInbox.Accepted.WithAddedQueue(bot.proposed.messages)
	return protocol.NewInbox(
		updatedQueues,
		bot.pendingMessages,
	)
}

func (bot waitingObserver) ProposalResults() valmessage.UnanimousUpdateResults {
	return valmessage.UnanimousUpdateResults{
		SequenceNum:       bot.proposed.sequenceNum,
		BeforeHash:        bot.origMachine.Hash(),
		TimeBounds:        bot.timeBounds,
		NewInboxHash:      bot.ProposedInbox().Receive().Hash(),
		OriginalInboxHash: bot.OrigInbox().Receive().Hash(),
		Assertion:         bot.proposed.assertion,
	}
}

func (bot waitingObserver) Clone() waitingObserver {
	return waitingObserver{
		validatorConfig:  bot.validatorConfig,
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

func (bot waitingObserver) GetCore() *validatorCore {
	if bot.assertion != nil {
		return &validatorCore{
			inbox:   protocol.NewInbox(bot.acceptedMessages, bot.pendingMessages),
			balance: bot.acceptedBalance,
			machine: bot.acceptedMachine,
		}
	}
	return &validatorCore{
		inbox:   protocol.NewInbox(bot.origMessages, bot.pendingMessages),
		balance: bot.origBalance,
		machine: bot.origMachine,
	}
}

func (bot waitingObserver) SendMessageToVM(msg protocol.Message) {
	bot.pendingMessages.AddMessage(msg)
}

func (bot waitingObserver) OffchainContext(
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

func (bot waitingObserver) validateUnanimousRequest(request valmessage.UnanimousRequestData) error {
	core := bot.GetCore()

	if request.BeforeHash != core.machine.Hash() {
		return errors.New("recieved unanimous request with invalid before hash")
	}
	if request.BeforeInbox != core.inbox.Receive().Hash() {
		fmt.Println("validateUnanimousRequest", core.inbox.Receive())
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

func (bot waitingObserver) PreparePendingUnanimous(request unanimousUpdateRequest) (waitingObserver, error) {
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

	return waitingObserver{
		validatorConfig: bot.validatorConfig,
		proposed: &proposedUpdate{
			machine:     request.Machine,
			messages:    mq,
			assertion:   assertion,
			sequenceNum: request.SequenceNum,
			newLogCount: newLogCount,
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

func (bot waitingObserver) FinalizePendingUnanimous(signatures []valmessage.Signature) (validatorState, *proposedUpdate, error) {
	if bot.proposed == nil {
		return nil, nil, errors.New("no pending assertion")
	}

	core := bot.GetCore()
	balance := core.balance.Clone()

	// This spend is guaranteed to be correct since the VM made sure to only produce on outgoing if it could spend
	_ = balance.SpendAll(protocol.NewBalanceTrackerFromMessages(bot.proposed.assertion.OutMsgs))

	return waitingObserver{
		validatorConfig:  bot.validatorConfig,
		proposed:         nil,
		acceptedMachine:  bot.proposed.machine,
		acceptedMessages: core.inbox.Accepted.WithAddedQueue(bot.proposed.messages),
		acceptedBalance:  balance,
		assertion:        bot.proposed.assertion,
		sequenceNum:      bot.proposed.sequenceNum,
		signatures:       signatures,
		timeBounds:       bot.timeBounds,
		pendingMessages:  bot.pendingMessages,
		origMessages:     bot.origMessages,
		origBalance:      bot.origBalance,
		origMachine:      bot.origMachine,
	}, bot.proposed, nil
}

func (bot waitingObserver) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot waitingObserver) UpdateState(ev ethbridge.Event, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case ethbridge.FinalUnanimousAssertEvent:
		if bot.sequenceNum != math.MaxUint64 {
			return nil, nil, nil, errors.New("waiting observer saw signed final unanimous proposal that it doesn't remember")
		}
		core := bot.GetCore()
		core.DeliverMessagesToVM()
		return newWaitingObserver(bot.validatorConfig, core), nil, nil, nil
	case ethbridge.ProposedUnanimousAssertEvent:
		if bot.acceptedMachine == nil || ev.SequenceNum > bot.sequenceNum {
			return nil, nil, nil, errors.New("waiting observer saw signed unanimous proposal that it doesn't remember")
		} else if ev.SequenceNum < bot.sequenceNum {
			newBot, msgs, err := bot.CloseUnanimous(nil)
			return newBot, nil, msgs, err
		} else {
			return waitingOffchainClosing{
				bot.validatorConfig,
				bot.GetCore(),
				bot.assertion,
				time + bot.config.GracePeriod,
				nil,
			}, nil, nil, nil
		}
	case ethbridge.DisputableAssertionEvent:
		if bot.acceptedMachine != nil {
			newBot, msgs, err := bot.CloseUnanimous(nil)
			return newBot, nil, msgs, err
		}

		core := bot.GetCore()
		deadline := time + bot.config.GracePeriod
		inbox := core.inbox
		var inboxVal value.Value
		if inbox.Receive().Hash() == ev.Precondition.BeforeInbox.Hash() {
			inboxVal = inbox.Receive()
		} else if inbox.ReceivePending().Hash() == ev.Precondition.BeforeInbox.Hash() {
			inboxVal = inbox.ReceivePending()
		} else {
			return nil, nil, nil, errors.New("waiting observer has incorrect valmessage")
		}
		updatedState := core.machine.Clone()
		actx := protocol.NewMachineAssertionContext(
			updatedState,
			core.balance,
			ev.Precondition.TimeBounds,
			inboxVal,
		)
		updatedState.Run(int32(ev.Assertion.NumSteps))
		ad := actx.Finalize(updatedState)

		var msgs []valmessage.OutgoingMessage
		if !ad.GetAssertion().Stub().Equals(ev.Assertion) || bot.challengeEverything {
			msgs = append(msgs, valmessage.SendInitiateChallengeMessage{
				Precondition: ev.Precondition,
				Assertion:    ev.Assertion,
			})
		}
		return watchingAssertionObserver{
			core,
			bot.validatorConfig,
			inboxVal,
			updatedState,
			deadline,
			ev.Precondition,
			ad.GetAssertion(),
		}, nil, msgs, nil
	default:
		return nil, nil, nil, &Error{nil, fmt.Sprintf("ERROR: waitingObserver: VM state got unsynchronized with valmessage %T", ev)}
	}
}

type watchingAssertionObserver struct {
	*validatorCore
	*validatorConfig
	inboxVal     value.Value
	pendingState *vm.Machine
	deadline     uint64
	precondition *protocol.Precondition
	assertion    *protocol.Assertion
}

func (bot watchingAssertionObserver) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	if time <= bot.deadline {
		return bot, nil, nil
	}

	newBalance := bot.balance.Clone()
	err := newBalance.SpendAll(bot.precondition.BeforeBalance)
	if err != nil {
		log.Fatal("Ethbridge admited assertion with more than available balance")
	}
	return finalizingAssertObserver{
		validatorCore: &validatorCore{
			inbox:   bot.inbox,
			balance: newBalance,
			machine: bot.machine,
		},
		validatorConfig: bot.validatorConfig,
		ResultChan:      nil,
	}, []valmessage.OutgoingMessage{valmessage.FinalizedAssertion{Assertion: bot.assertion, NewLogCount: len(bot.assertion.Logs)}}, nil
}

func (bot watchingAssertionObserver) UpdateState(ev ethbridge.Event, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case ethbridge.InitiateChallengeEvent:
		deadline := time + bot.config.GracePeriod
		var challenge challengeState
		if ev.Challenger == bot.address {
			challenge = waitingContinuingChallenger{
				bot.validatorConfig,
				bot.precondition,
				bot.assertion.Stub(),
				bot.inboxVal,
				bot.machine.Clone(),
				deadline,
			}
		} else {
			challenge = waitingChallengeObserver{
				bot.validatorConfig,
				bot.precondition,
				bot.assertion.Stub(),
				deadline,
			}
		}
		return newWaitingObserver(bot.validatorConfig, bot.validatorCore), challenge, nil, nil

	default:
		return nil, nil, nil, &Error{nil, "ERROR: WaitingValidObserver: VM state got unsynchronized"}
	}
}

type attemptingAssertDefender struct {
	*validatorCore
	*validatorConfig
	request disputableAssertionRequest
}

func (bot attemptingAssertDefender) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot attemptingAssertDefender) UpdateState(ev ethbridge.Event, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case ethbridge.DisputableAssertionEvent:
		if ev.Asserter != bot.address {
			fmt.Println("attemptingAssertDefender: Other assertion got in before ours")
			return newWaitingObserver(bot.validatorConfig, bot.validatorCore).UpdateState(ev, time)
		}

		deadline := time + bot.config.GracePeriod
		return waitingAssertDefender{
			bot.validatorCore,
			bot.validatorConfig,
			bot.request,
			deadline,
		}, nil, nil, nil
	default:
		return nil, nil, nil, &Error{nil, "ERROR: attemptingAssertDefender: VM state got unsynchronized"}
	}
}

type waitingAssertDefender struct {
	*validatorCore
	*validatorConfig
	request  disputableAssertionRequest
	deadline uint64
}

func (bot waitingAssertDefender) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	if time <= bot.deadline {
		return bot, nil, nil
	}

	assertion := bot.request.Defender.GetAssertion()
	evmRes := valmessage.FinalizedAssertion{Assertion: assertion, NewLogCount: len(assertion.Logs)}
	conf := valmessage.SendConfirmedAssertMessage{
		Precondition: bot.request.Defender.GetPrecondition(),
		Assertion:    bot.request.Defender.GetAssertion(),
	}
	newBalance := bot.balance.Clone()
	err := newBalance.SpendAll(protocol.NewBalanceTrackerFromMessages(bot.request.Defender.GetAssertion().OutMsgs))
	if err != nil {
		log.Fatal("Ethbridge admited assertion with more than available balance")
	}
	return finalizingAssertObserver{
		&validatorCore{
			inbox:   bot.inbox,
			balance: newBalance,
			machine: bot.request.State,
		},
		bot.validatorConfig,
		bot.request.ResultChan,
	}, []valmessage.OutgoingMessage{evmRes, conf}, nil
}

func (bot waitingAssertDefender) UpdateState(ev ethbridge.Event, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case ethbridge.InitiateChallengeEvent:
		bot.request.NotifyInvalid()
		ct, msgs, err := defenseValidator(bot.validatorConfig, bot.request.Defender, time)
		return newWaitingObserver(bot.validatorConfig, bot.validatorCore), ct, msgs, err

	default:
		return nil, nil, nil, &Error{nil, "ERROR: waitingAssertDefender: VM state got unsynchronized"}
	}
}

type finalizingAssertObserver struct {
	*validatorCore
	*validatorConfig
	ResultChan chan<- bool
}

func (bot finalizingAssertObserver) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot finalizingAssertObserver) UpdateState(ev ethbridge.Event, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case ethbridge.ConfirmedAssertEvent:
		if bot.ResultChan != nil {
			bot.ResultChan <- true
		}
		bot.GetCore().DeliverMessagesToVM()
		return newWaitingObserver(bot.validatorConfig, bot.validatorCore), nil, nil, nil
	default:
		return nil, nil, nil, &Error{nil, "ERROR: FinalizingAssertDefender: VM state got unsynchronized"}
	}
}
