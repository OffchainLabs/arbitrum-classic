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

type WaitingObserver struct {
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

func NewWaitingObserver(config *validatorConfig, core *validatorCore) WaitingObserver {
	return WaitingObserver{
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

func (bot WaitingObserver) SlowCloseUnanimous() valmessage.SendProposeUnanimousAssertMessage {
	return valmessage.SendProposeUnanimousAssertMessage{
		NewInboxHash: bot.GetCore().inbox.Receive().Hash(),
		TimeBounds:   bot.timeBounds,
		Assertion:    bot.assertion,
		SequenceNum:  bot.sequenceNum,
		Signatures:   bot.signatures,
	}
}

func (bot WaitingObserver) FastCloseUnanimous() valmessage.SendUnanimousAssertMessage {
	inboxHash := bot.GetCore().inbox.Receive().Hash()
	return valmessage.SendUnanimousAssertMessage{
		NewInboxHash: inboxHash,
		TimeBounds:   bot.timeBounds,
		Assertion:    bot.assertion,
		Signatures:   bot.signatures,
	}
}

func (bot WaitingObserver) CloseUnanimous(retChan chan<- bool) (validatorState, []valmessage.OutgoingMessage, error) {
	if bot.assertion == nil {
		return bot, nil, errors.New("couldn't close since no assertion is open")
	}

	if bot.sequenceNum == math.MaxUint64 {
		return AttemptingUnanimousClosing{
				bot.validatorConfig,
				bot.GetCore(),
				bot.assertion,
				retChan,
			},
			[]valmessage.OutgoingMessage{bot.FastCloseUnanimous()},
			nil
	} else {
		return AttemptingOffchainClosing{
				bot.validatorConfig,
				bot.GetCore(),
				bot.sequenceNum,
				bot.assertion,
				retChan,
			},
			[]valmessage.OutgoingMessage{bot.SlowCloseUnanimous()},
			nil
	}
}

func (bot WaitingObserver) OrigInbox() *protocol.Inbox {
	return protocol.NewInbox(
		bot.origMessages,
		bot.pendingMessages,
	)
}

func (bot WaitingObserver) ProposedInbox() *protocol.Inbox {
	curInbox := bot.GetCore().inbox
	updatedQueues := curInbox.Accepted.WithAddedQueue(bot.proposed.messages)
	return protocol.NewInbox(
		updatedQueues,
		bot.pendingMessages,
	)
}

func (bot WaitingObserver) ProposalResults() valmessage.UnanimousUpdateResults {
	return valmessage.UnanimousUpdateResults{
		SequenceNum:       bot.proposed.sequenceNum,
		BeforeHash:        bot.origMachine.Hash(),
		TimeBounds:        bot.timeBounds,
		NewInboxHash:      bot.ProposedInbox().Receive().Hash(),
		OriginalInboxHash: bot.OrigInbox().Receive().Hash(),
		Assertion:         bot.proposed.assertion,
	}
}

func (bot WaitingObserver) Clone() WaitingObserver {
	return WaitingObserver{
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

func (bot WaitingObserver) GetCore() *validatorCore {
	if bot.assertion != nil {
		return &validatorCore{
			inbox:   protocol.NewInbox(bot.acceptedMessages, bot.pendingMessages),
			balance: bot.acceptedBalance,
			machine: bot.acceptedMachine,
		}
	} else {
		return &validatorCore{
			inbox:   protocol.NewInbox(bot.origMessages, bot.pendingMessages),
			balance: bot.origBalance,
			machine: bot.origMachine,
		}
	}
}

func (bot WaitingObserver) SendMessageToVM(msg protocol.Message) {
	bot.pendingMessages.AddMessage(msg)
}

func (bot WaitingObserver) OffchainContext(
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

func (bot WaitingObserver) validateUnanimousRequest(request valmessage.UnanimousRequestData) error {
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

func (bot WaitingObserver) PreparePendingUnanimous(request valmessage.UnanimousUpdateRequest) (WaitingObserver, error) {
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

	return WaitingObserver{
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

func (bot WaitingObserver) FinalizePendingUnanimous(signatures []valmessage.Signature) (validatorState, *proposedUpdate, error) {
	if bot.proposed == nil {
		return nil, nil, errors.New("no pending assertion")
	}

	core := bot.GetCore()
	balance := core.balance.Clone()

	// This spend is guaranteed to be correct since the VM made sure to only produce on outgoing if it could spend
	_ = balance.SpendAll(protocol.NewBalanceTrackerFromMessages(bot.proposed.assertion.OutMsgs))

	return WaitingObserver{
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

func (bot WaitingObserver) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot WaitingObserver) UpdateState(ev valmessage.IncomingMessage, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case valmessage.FinalUnanimousAssertMessage:
		if bot.sequenceNum != math.MaxUint64 {
			return nil, nil, nil, errors.New("waiting observer saw signed final unanimous proposal that it doesn't remember")
		}
		core := bot.GetCore()
		core.DeliverMessagesToVM()
		return NewWaitingObserver(bot.validatorConfig, core), nil, nil, nil
	case valmessage.ProposedUnanimousAssertMessage:
		if bot.acceptedMachine == nil || ev.SequenceNum > bot.sequenceNum {
			return nil, nil, nil, errors.New("waiting observer saw signed unanimous proposal that it doesn't remember")
		} else if ev.SequenceNum < bot.sequenceNum {
			newBot, msgs, err := bot.CloseUnanimous(nil)
			return newBot, nil, msgs, err
		} else {
			return WaitingOffchainClosing{
				bot.validatorConfig,
				bot.GetCore(),
				bot.assertion,
				time + bot.config.GracePeriod,
				nil,
			}, nil, nil, nil
		}
	case valmessage.AssertMessage:
		if bot.acceptedMachine != nil {
			newBot, msgs, err := bot.CloseUnanimous(nil)
			return newBot, nil, msgs, err
		} else {
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
			return WatchingAssertionObserver{
				core,
				bot.validatorConfig,
				inboxVal,
				updatedState,
				deadline,
				ev.Precondition,
				ad.GetAssertion(),
			}, nil, msgs, nil
		}
	default:
		return nil, nil, nil, &Error{nil, fmt.Sprintf("ERROR: WaitingObserver: VM state got unsynchronized with valmessage %T", ev)}
	}
}

type WatchingAssertionObserver struct {
	*validatorCore
	*validatorConfig
	inboxVal     value.Value
	pendingState *vm.Machine
	deadline     uint64
	precondition *protocol.Precondition
	assertion    *protocol.Assertion
}

func (bot WatchingAssertionObserver) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	if time > bot.deadline {
		newBalance := bot.balance.Clone()
		err := newBalance.SpendAll(bot.precondition.BeforeBalance)
		if err != nil {
			log.Fatal("Ethbridge admited assertion with more than available balance")
		}
		return FinalizingAssertObserver{
			validatorCore: &validatorCore{
				inbox:   bot.inbox,
				balance: newBalance,
				machine: bot.machine,
			},
			validatorConfig: bot.validatorConfig,
			ResultChan:      nil,
		}, []valmessage.OutgoingMessage{valmessage.FinalizedAssertion{Assertion: bot.assertion, NewLogCount: len(bot.assertion.Logs)}}, nil
	} else {
		return bot, nil, nil
	}
}

func (bot WatchingAssertionObserver) UpdateState(ev valmessage.IncomingMessage, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case valmessage.InitiateChallengeMessage:
		deadline := time + bot.config.GracePeriod
		var challenge challengeState = nil
		if ev.Challenger == bot.address {
			challenge = WaitingContinuingChallenger{
				bot.validatorConfig,
				bot.precondition,
				bot.assertion.Stub(),
				bot.inboxVal,
				bot.machine.Clone(),
				deadline,
			}
		} else {
			challenge = WaitingChallengeObserver{
				bot.validatorConfig,
				bot.precondition,
				bot.assertion.Stub(),
				deadline,
			}
		}
		return NewWaitingObserver(bot.validatorConfig, bot.validatorCore), challenge, nil, nil

	default:
		return nil, nil, nil, &Error{nil, "ERROR: WaitingValidObserver: VM state got unsynchronized"}
	}
}

type AttemptingAssertDefender struct {
	*validatorCore
	*validatorConfig
	request valmessage.DisputableAssertionRequest
}

func (bot AttemptingAssertDefender) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot AttemptingAssertDefender) UpdateState(ev valmessage.IncomingMessage, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev := ev.(type) {
	case valmessage.AssertMessage:
		if ev.Asserter == bot.address {
			deadline := time + bot.config.GracePeriod
			return WaitingAssertDefender{
				bot.validatorCore,
				bot.validatorConfig,
				bot.request,
				deadline,
			}, nil, nil, nil
		} else {
			fmt.Println("AttemptingAssertDefender: Other assertion got in before ours")
			return NewWaitingObserver(bot.validatorConfig, bot.validatorCore).UpdateState(ev, time)
		}
	default:
		return nil, nil, nil, &Error{nil, "ERROR: AttemptingAssertDefender: VM state got unsynchronized"}
	}
}

type WaitingAssertDefender struct {
	*validatorCore
	*validatorConfig
	request  valmessage.DisputableAssertionRequest
	deadline uint64
}

func (bot WaitingAssertDefender) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	if time > bot.deadline {
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
		return FinalizingAssertObserver{
			&validatorCore{
				inbox:   bot.inbox,
				balance: newBalance,
				machine: bot.request.State,
			},
			bot.validatorConfig,
			bot.request.ResultChan,
		}, []valmessage.OutgoingMessage{evmRes, conf}, nil
	} else {
		return bot, nil, nil
	}
}

func (bot WaitingAssertDefender) UpdateState(ev valmessage.IncomingMessage, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case valmessage.InitiateChallengeMessage:
		bot.request.NotifyInvalid()
		ct, msgs, err := NewDefendingValidator(bot.validatorConfig, bot.request.Defender, time)
		return NewWaitingObserver(bot.validatorConfig, bot.validatorCore), ct, msgs, err

	default:
		return nil, nil, nil, &Error{nil, "ERROR: WaitingAssertDefender: VM state got unsynchronized"}
	}
}

type FinalizingAssertObserver struct {
	*validatorCore
	*validatorConfig
	ResultChan chan<- bool
}

func (bot FinalizingAssertObserver) UpdateTime(time uint64) (validatorState, []valmessage.OutgoingMessage, error) {
	return bot, nil, nil
}

func (bot FinalizingAssertObserver) UpdateState(ev valmessage.IncomingMessage, time uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error) {
	switch ev.(type) {
	case valmessage.ConfirmedAssertMessage:
		if bot.ResultChan != nil {
			bot.ResultChan <- true
		}
		bot.GetCore().DeliverMessagesToVM()
		return NewWaitingObserver(bot.validatorConfig, bot.validatorCore), nil, nil, nil
	default:
		return nil, nil, nil, &Error{nil, "ERROR: FinalizingAssertDefender: VM state got unsynchronized"}
	}
}
