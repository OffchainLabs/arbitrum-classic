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

type WaitingObserver struct {
	*validatorConfig

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
		pendingMessages:  core.Inbox.PendingQueue,
		origMessages:     core.Inbox.Accepted,
		origBalance:      core.Balance,
		origMachine:      core.Machine,
	}
}

func (bot WaitingObserver) SlowCloseUnanimous(bridge bridge.Bridge) {
	bridge.UnanimousAssert(
		bot.GetCore().Inbox.Receive().Hash(),
		bot.timeBounds,
		bot.assertion,
		bot.sequenceNum,
		bot.signatures,
	)
}

func (bot WaitingObserver) FastCloseUnanimous(bridge bridge.Bridge) {
	inboxHash := bot.GetCore().Inbox.Receive().Hash()
	bridge.FinalUnanimousAssert(
		inboxHash,
		bot.timeBounds,
		bot.assertion,
		bot.signatures,
	)
}

func (bot WaitingObserver) CloseUnanimous(bridge bridge.Bridge, retChan chan<- bool) (ValidatorState, error) {
	if bot.assertion == nil {
		return bot, errors.New("couldn't close since no Assertion is open")
	}

	if bot.sequenceNum == math.MaxUint64 {
		bot.FastCloseUnanimous(bridge)
		return attemptingUnanimousClosing{
				bot.validatorConfig,
				bot.GetCore(),
				bot.assertion,
				retChan,
			},
			nil
	}
	bot.SlowCloseUnanimous(bridge)
	return attemptingOffchainClosing{
			bot.validatorConfig,
			bot.GetCore(),
			bot.sequenceNum,
			bot.assertion,
			retChan,
		},
		nil
}

func (bot WaitingObserver) AttemptAssertion(request DisputableAssertionRequest, bridge bridge.Bridge) ValidatorState {
	bridge.DisputableAssert(
		request.Defender.GetPrecondition(),
		request.Defender.GetAssertion(),
	)

	return attemptingAssertDefender{
		bot.GetCore(),
		bot.GetConfig(),
		request,
	}
}

func (bot WaitingObserver) OrigInbox() *protocol.Inbox {
	return protocol.NewInbox(
		bot.origMessages,
		bot.pendingMessages,
	)
}

func (bot WaitingObserver) ProposedInbox() *protocol.Inbox {
	curInbox := bot.GetCore().Inbox
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
		Assertion:         bot.proposed.Assertion,
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
			Inbox:   protocol.NewInbox(bot.acceptedMessages, bot.pendingMessages),
			Balance: bot.acceptedBalance,
			Machine: bot.acceptedMachine,
		}
	}
	return &validatorCore{
		Inbox:   protocol.NewInbox(bot.origMessages, bot.pendingMessages),
		Balance: bot.origBalance,
		Machine: bot.origMachine,
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

func (bot WaitingObserver) ValidateUnanimousRequest(request valmessage.UnanimousRequestData) error {
	core := bot.GetCore()

	if request.BeforeHash != core.Machine.Hash() {
		return errors.New("recieved unanimous request with invalid before hash")
	}
	if request.BeforeInbox != core.Inbox.Receive().Hash() {
		fmt.Println("validateUnanimousRequest", core.Inbox.Receive())
		return errors.New("recieved unanimous request with invalid before Inbox")
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

func (bot WaitingObserver) PreparePendingUnanimous(request UnanimousUpdateRequest) (WaitingObserver, error) {
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

func (bot WaitingObserver) FinalizePendingUnanimous(signatures [][]byte) (ValidatorState, *proposedUpdate, error) {
	if bot.proposed == nil {
		return nil, nil, errors.New("no pending Assertion")
	}

	core := bot.GetCore()
	balance := core.Balance.Clone()

	// This spend is guaranteed to be correct since the VM made sure to only produce on outgoing if it could spend
	_ = balance.SpendAll(protocol.NewBalanceTrackerFromMessages(bot.proposed.Assertion.OutMsgs))

	return WaitingObserver{
		validatorConfig:  bot.validatorConfig,
		proposed:         nil,
		acceptedMachine:  bot.proposed.machine,
		acceptedMessages: core.Inbox.Accepted.WithAddedQueue(bot.proposed.messages),
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

func (bot WaitingObserver) UpdateTime(time uint64, bridge bridge.Bridge) (ValidatorState, error) {
	return bot, nil
}

func (bot WaitingObserver) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (ValidatorState, ChallengeState, error) {
	switch ev := ev.(type) {
	case ethbridge.FinalUnanimousAssertEvent:
		if bot.sequenceNum != math.MaxUint64 {
			return nil, nil, errors.New("waiting observer saw signed final unanimous proposal that it doesn't remember")
		}
		core := bot.GetCore()
		core.DeliverMessagesToVM()
		return NewWaitingObserver(bot.validatorConfig, core), nil, nil
	case ethbridge.ProposedUnanimousAssertEvent:
		if bot.acceptedMachine == nil || ev.SequenceNum > bot.sequenceNum {
			return nil, nil, errors.New("waiting observer saw signed unanimous proposal that it doesn't remember")
		} else if ev.SequenceNum < bot.sequenceNum {
			newBot, err := bot.CloseUnanimous(bridge, nil)
			return newBot, nil, err
		} else {
			return waitingOffchainClosing{
				bot.validatorConfig,
				bot.GetCore(),
				bot.assertion,
				time + bot.Config.GracePeriod,
				nil,
			}, nil, nil
		}
	case ethbridge.DisputableAssertionEvent:
		if bot.acceptedMachine != nil {
			newBot, err := bot.CloseUnanimous(bridge, nil)
			return newBot, nil, err
		}

		core := bot.GetCore()
		deadline := time + bot.Config.GracePeriod
		inbox := core.Inbox
		var inboxVal value.Value
		if inbox.Receive().Hash() == ev.Precondition.BeforeInbox.Hash() {
			inboxVal = inbox.Receive()
		} else if inbox.ReceivePending().Hash() == ev.Precondition.BeforeInbox.Hash() {
			inboxVal = inbox.ReceivePending()
		} else {
			return nil, nil, errors.New("waiting observer has incorrect valmessage")
		}
		updatedState := core.Machine.Clone()
		actx := protocol.NewMachineAssertionContext(
			updatedState,
			core.Balance,
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
		return watchingAssertionObserver{
			core,
			bot.validatorConfig,
			inboxVal,
			updatedState,
			deadline,
			ev.Precondition,
			ad.GetAssertion(),
		}, nil, nil
	default:
		return nil, nil, &Error{nil, fmt.Sprintf("ERROR: WaitingObserver: VM state got unsynchronized with valmessage %T", ev)}
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

func (bot watchingAssertionObserver) UpdateTime(time uint64, bridge bridge.Bridge) (ValidatorState, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	newBalance := bot.Balance.Clone()
	err := newBalance.SpendAll(bot.precondition.BeforeBalance)
	if err != nil {
		log.Fatal("Ethbridge admited Assertion with more than available Balance")
	}
	bridge.FinalizedAssertion(bot.assertion, len(bot.assertion.Logs))
	return finalizingAssertObserver{
		validatorCore: &validatorCore{
			Inbox:   bot.Inbox,
			Balance: newBalance,
			Machine: bot.Machine,
		},
		validatorConfig: bot.validatorConfig,
		ResultChan:      nil,
	}, nil
}

func (bot watchingAssertionObserver) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (ValidatorState, ChallengeState, error) {
	switch ev := ev.(type) {
	case ethbridge.InitiateChallengeEvent:
		deadline := time + bot.Config.GracePeriod
		var challenge ChallengeState
		if ev.Challenger == bot.Address {
			challenge = waitingContinuingChallenger{
				bot.validatorConfig,
				bot.precondition,
				bot.assertion.Stub(),
				bot.inboxVal,
				bot.Machine.Clone(),
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
		return NewWaitingObserver(bot.validatorConfig, bot.validatorCore), challenge, nil

	default:
		return nil, nil, &Error{nil, "ERROR: WaitingValidObserver: VM state got unsynchronized"}
	}
}

type attemptingAssertDefender struct {
	*validatorCore
	*validatorConfig
	request DisputableAssertionRequest
}

func (bot attemptingAssertDefender) UpdateTime(time uint64, bridge bridge.Bridge) (ValidatorState, error) {
	return bot, nil
}

func (bot attemptingAssertDefender) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (ValidatorState, ChallengeState, error) {
	switch ev := ev.(type) {
	case ethbridge.DisputableAssertionEvent:
		if ev.Asserter != bot.Address {
			fmt.Println("attemptingAssertDefender: Other Assertion got in before ours")
			return NewWaitingObserver(bot.validatorConfig, bot.validatorCore).UpdateState(ev, time, bridge)
		}

		deadline := time + bot.Config.GracePeriod
		return waitingAssertDefender{
			bot.validatorCore,
			bot.validatorConfig,
			bot.request,
			deadline,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: attemptingAssertDefender: VM state got unsynchronized"}
	}
}

type waitingAssertDefender struct {
	*validatorCore
	*validatorConfig
	request  DisputableAssertionRequest
	deadline uint64
}

func (bot waitingAssertDefender) UpdateTime(time uint64, bridge bridge.Bridge) (ValidatorState, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	newBalance := bot.Balance.Clone()
	err := newBalance.SpendAll(protocol.NewBalanceTrackerFromMessages(bot.request.Defender.GetAssertion().OutMsgs))
	if err != nil {
		log.Fatal("Ethbridge admited Assertion with more than available Balance")
	}
	bridge.ConfirmDisputableAssertion(
		bot.request.Defender.GetPrecondition(),
		bot.request.Defender.GetAssertion(),
	)
	assertion := bot.request.Defender.GetAssertion()
	bridge.FinalizedAssertion(assertion, len(assertion.Logs))
	return finalizingAssertObserver{
		&validatorCore{
			Inbox:   bot.Inbox,
			Balance: newBalance,
			Machine: bot.request.State,
		},
		bot.validatorConfig,
		bot.request.ResultChan,
	}, nil
}

func (bot waitingAssertDefender) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (ValidatorState, ChallengeState, error) {
	switch ev.(type) {
	case ethbridge.InitiateChallengeEvent:
		bot.request.NotifyInvalid()
		ct, err := defenseValidator(bot.validatorConfig, bot.request.Defender, time, bridge)
		return NewWaitingObserver(bot.validatorConfig, bot.validatorCore), ct, err

	default:
		return nil, nil, &Error{nil, "ERROR: waitingAssertDefender: VM state got unsynchronized"}
	}
}

type finalizingAssertObserver struct {
	*validatorCore
	*validatorConfig
	ResultChan chan<- bool
}

func (bot finalizingAssertObserver) UpdateTime(time uint64, bridge bridge.Bridge) (ValidatorState, error) {
	return bot, nil
}

func (bot finalizingAssertObserver) UpdateState(ev ethbridge.Event, time uint64, bridge bridge.Bridge) (ValidatorState, ChallengeState, error) {
	switch ev.(type) {
	case ethbridge.ConfirmedAssertEvent:
		if bot.ResultChan != nil {
			bot.ResultChan <- true
		}
		bot.GetCore().DeliverMessagesToVM()
		return NewWaitingObserver(bot.validatorConfig, bot.validatorCore), nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: FinalizingAssertDefender: VM state got unsynchronized"}
	}
}
