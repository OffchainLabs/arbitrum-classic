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

	"github.com/offchainlabs/arbitrum/packages/arb-validator/disputable"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge/challenger"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge/defender"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge/observer"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
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
	machine         machine.Machine
	newMessageCount uint64
	Assertion       *protocol.Assertion
	timeBounds      protocol.TimeBounds
	sequenceNum     uint64
	NewLogCount     int
}

func (p *proposedUpdate) clone() *proposedUpdate {
	return &proposedUpdate{
		machine:         p.machine.Clone(),
		newMessageCount: p.newMessageCount,
		Assertion:       p.Assertion,
		timeBounds:      p.timeBounds,
		sequenceNum:     p.sequenceNum,
		NewLogCount:     0,
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

func (bot Waiting) OrigHash() [32]byte {
	return bot.orig.GetMachine().Hash()
}

func (bot Waiting) OrigInboxHash() [32]byte {
	return bot.orig.GetMachine().InboxHash().Hash()
}

func (bot Waiting) HasOpenAssertion() bool {
	return bot.assertion != nil
}

func (bot Waiting) CloseUnanimous(bridge bridge.Bridge) (chan *types.Receipt, chan error) {
	if bot.sequenceNum == math.MaxUint64 {
		return bridge.FinalizedUnanimousAssert(
			context.Background(),
			bot.GetCore().GetMachine().InboxHash().Hash(),
			bot.assertion,
			bot.signatures,
		)
	} else {
		return bridge.PendingUnanimousAssert(
			context.Background(),
			bot.GetCore().GetMachine().InboxHash().Hash(),
			bot.assertion,
			bot.sequenceNum,
			bot.signatures,
		)
	}
}

func (bot Waiting) ClosingUnanimous(retChan chan<- bool, errChan chan<- error) (ChannelState, error) {
	// If there is no active unanimous assertion, there is nothing to close
	// TODO: Validator should refuse to unanimous assert again from the same start point
	if bot.assertion == nil {
		err := errors.New("couldn't close since no Assertion is open")
		if errChan != nil {
			errChan <- err
		}
		return bot, err
	}

	if bot.sequenceNum == math.MaxUint64 {
		return attemptingUnanimousClosing{
				bot.Config,
				bot.GetCore(),
				bot.assertion,
				retChan,
				errChan,
			},
			nil
	} else {
		return attemptingOffchainClosing{
				bot.Config,
				bot.GetCore(),
				bot.sequenceNum,
				bot.assertion,
				retChan,
				errChan,
			},
			nil
	}
}

func (bot Waiting) AttemptAssertion(ctx context.Context, request disputable.AssertionRequest, bridge bridge.ArbVMBridge) ChainState {
	bridge.PendingDisputableAssert(
		ctx,
		request.Precondition,
		request.Assertion,
	)

	return attemptingAssertion{&disputableAssertCore{
		bot.GetCore(),
		bot.GetConfig(),
		request.AfterCore,
		request.Precondition,
		request.Assertion,
		request.ResultChan,
		request.ErrorChan,
	}}
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
		UnanimousRequestData: valmessage.UnanimousRequestData{
			BeforeHash:  bot.orig.GetMachine().Hash(),
			BeforeInbox: bot.orig.GetMachine().InboxHash().Hash(),
			SequenceNum: bot.proposed.sequenceNum,
			TimeBounds:  bot.proposed.timeBounds,
		},
		NewInboxHash: bot.proposed.machine.InboxHash().Hash(),
		Assertion:    bot.proposed.Assertion,
		NewLogCount:  bot.proposed.NewLogCount,
	}
}

func (bot Waiting) ProposedMessageCount() uint64 {
	return bot.proposed.newMessageCount
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
) uint64 {
	var seqNum uint64
	if bot.accepted != nil {
		seqNum = bot.sequenceNum + 1
	} else {
		seqNum = 0
	}

	if final {
		seqNum = math.MaxUint64
	}

	return seqNum
}

func (bot Waiting) ValidateUnanimousRequest(request valmessage.UnanimousRequestData) error {
	if bot.proposed != nil {
		return errors.New("can't process unanimous request while request is pending")
	}
	if request.BeforeHash != bot.orig.GetMachine().Hash() {
		return errors.New("recieved unanimous request with invalid before hash")
	}
	if request.BeforeInbox != bot.orig.GetMachine().InboxHash().Hash() {
		return errors.New("recieved unanimous request with invalid before inbox")
	}

	if bot.accepted != nil {
		if request.TimeBounds[0] < bot.timeBounds[0] {
			return errors.New("unanimous assertion request starting time bound may only increase")
		}
		if request.SequenceNum <= bot.sequenceNum {
			return errors.New("recieved unanimous request with invalid sequence number")
		}
	}
	return nil
}

func (bot Waiting) ValidateUnanimousAssertion(request valmessage.UnanimousRequestData) error {
	if bot.proposed == nil {
		return errors.New("validator unanimous assertion without pending request")
	}

	if request.BeforeHash != bot.orig.GetMachine().Hash() {
		return errors.New("recieved unanimous update with invalid before hash")
	}
	if request.BeforeInbox != bot.orig.GetMachine().InboxHash().Hash() {
		return errors.New("recieved unanimous update with invalid before inbox")
	}

	if bot.accepted != nil {
		if request.TimeBounds[0] < bot.timeBounds[0] {
			return errors.New("unanimous assertion starting time bound may only increase")
		}
	}

	if request.SequenceNum < bot.proposed.sequenceNum {
		return errors.New("recieved unanimous update with invalid sequence number")
	}
	return nil
}

func (bot Waiting) PreparePendingUnanimous(
	newAssertion *protocol.Assertion,
	messages []protocol.Message,
	machine machine.Machine,
	sequenceNum uint64,
	timeBounds protocol.TimeBounds,
	shouldFinalize func(*protocol.Assertion) bool,
) (Waiting, error) {
	newLogCount := len(newAssertion.Logs)
	if bot.assertion != nil {
		newAssertion.NumSteps += bot.assertion.NumSteps
		newAssertion.OutMsgs = append(bot.assertion.OutMsgs, newAssertion.OutMsgs...)
		newAssertion.Logs = append(bot.assertion.Logs, newAssertion.Logs...)
	}

	if shouldFinalize(newAssertion) {
		sequenceNum = math.MaxUint64
	}

	return Waiting{
		Config: bot.Config,
		proposed: &proposedUpdate{
			machine:         machine,
			newMessageCount: uint64(len(messages)),
			Assertion:       newAssertion,
			timeBounds:      timeBounds,
			sequenceNum:     sequenceNum,
			NewLogCount:     newLogCount,
		},
		accepted:    bot.accepted,
		assertion:   bot.assertion,
		sequenceNum: bot.sequenceNum,
		signatures:  bot.signatures,
		timeBounds:  bot.timeBounds,
		orig:        bot.orig,
	}, nil
}

func (bot Waiting) FinalizePendingUnanimous(signatures [][]byte) (Waiting, error) {
	if bot.proposed == nil {
		return Waiting{}, errors.New("no pending Assertion")
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
		sequenceNum: bot.proposed.sequenceNum,
		signatures:  signatures,
		timeBounds:  bot.proposed.timeBounds,
		orig:        bot.orig,
	}, nil
}

func (bot Waiting) updateState(ev ethconnection.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, challenge.State, error) {
	switch ev := ev.(type) {
	case ethconnection.PendingDisputableAssertionEvent:
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
		balance := c.GetBalance()
		_ = balance.SpendAll(protocol.NewBalanceTrackerFromMessages(assertion.OutMsgs))
		return watchingAssertion{
			c,
			bot.Config,
			inboxVal,
			core.NewCore(updatedState, balance),
			deadline,
			ev.Precondition,
			assertion,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, fmt.Sprintf("ERROR: Waiting: VM state got unsynchronized with valmessage %T", ev)}
	}
}

func (bot Waiting) ChainUpdateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot, nil
}

func (bot Waiting) ChainUpdateState(ev ethconnection.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, challenge.State, error) {
	return bot.updateState(ev, time, bridge)
}

func (bot Waiting) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot, nil
}

func (bot Waiting) ChannelUpdateState(ev ethconnection.Event, time uint64, bridge bridge.Bridge) (ChannelState, challenge.State, error) {
	switch ev := ev.(type) {
	case ethconnection.PendingUnanimousAssertEvent:
		if bot.accepted == nil || ev.SequenceNum > bot.sequenceNum {
			return nil, nil, errors.New("waiting observer saw pending unanimous assertion that it doesn't remember")
		} else if ev.SequenceNum < bot.sequenceNum {
			bot.CloseUnanimous(bridge)
			newBot, err := bot.ClosingUnanimous(nil, nil)
			return newBot, nil, err
		} else {
			return waitingOffchainClosing{
				bot.Config,
				bot.GetCore(),
				bot.assertion,
				time + bot.VMConfig.GracePeriod,
				nil,
				nil,
			}, nil, nil
		}
	default:
		return bot.updateState(ev, time, bridge)
	}
}

type watchingAssertion struct {
	*core.Core
	*core.Config
	inboxVal     value.Value
	pending      *core.Core
	deadline     uint64
	precondition *protocol.Precondition
	assertion    *protocol.Assertion
}

func (bot watchingAssertion) SendMessageToVM(msg protocol.Message) {
	bot.Core.SendMessageToVM(msg)
	bot.pending.SendMessageToVM(msg)
}

func (bot watchingAssertion) updateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	if time <= bot.deadline {
		return nil, nil
	}

	return finalizingAssertion{
		Core:       bot.pending,
		Config:     bot.Config,
		ResultChan: nil,
		assertion:  bot.assertion,
	}, nil
}

func (bot watchingAssertion) updateState(ev ethconnection.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, challenge.State, error) {
	switch ev := ev.(type) {
	case ethconnection.InitiateChallengeEvent:
		deadline := time + bot.VMConfig.GracePeriod
		var challengeState challenge.State
		if ev.Challenger == bot.Address {
			challengeState = challenger.New(
				bot.Config,
				bot.precondition,
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

func (bot watchingAssertion) ChainUpdateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot.updateTime(time, bridge)
}

func (bot watchingAssertion) ChainUpdateState(ev ethconnection.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, challenge.State, error) {
	return bot.updateState(ev, time, bridge)
}

func (bot watchingAssertion) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot.updateTime(time, bridge)
}

func (bot watchingAssertion) ChannelUpdateState(ev ethconnection.Event, time uint64, bridge bridge.Bridge) (ChannelState, challenge.State, error) {
	return bot.updateState(ev, time, bridge)
}

type disputableAssertCore struct {
	*core.Core
	*core.Config
	afterCore    *core.Core
	precondition *protocol.Precondition
	assertion    *protocol.Assertion
	resultChan   chan<- bool
	errorChan    chan<- error
}

func (d *disputableAssertCore) SendMessageToVM(msg protocol.Message) {
	d.Core.SendMessageToVM(msg)
	d.afterCore.SendMessageToVM(msg)
}

type attemptingAssertion struct {
	*disputableAssertCore
}

func (bot attemptingAssertion) ChainUpdateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot, nil
}

func (bot attemptingAssertion) ChainUpdateState(ev ethconnection.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, challenge.State, error) {
	switch ev := ev.(type) {
	case ethconnection.PendingDisputableAssertionEvent:
		if ev.Asserter != bot.Address {
			bot.errorChan <- errors.New("attemptingAssertion: Other Assertion got in before ours")
			close(bot.errorChan)
			close(bot.resultChan)
			return NewWaiting(bot.Config, bot.Core).ChainUpdateState(ev, time, bridge)
		}

		deadline := time + bot.VMConfig.GracePeriod
		return waitingAssertion{
			bot.disputableAssertCore,
			deadline,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: attemptingAssertion: VM state got unsynchronized"}
	}
}

func (bot attemptingAssertion) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot, nil
}

func (bot attemptingAssertion) ChannelUpdateState(ev ethconnection.Event, time uint64, bridge bridge.Bridge) (ChannelState, challenge.State, error) {
	switch ev := ev.(type) {
	case ethconnection.PendingDisputableAssertionEvent:
		if ev.Asserter != bot.Address {
			bot.errorChan <- errors.New("attemptingAssertion: Other Assertion got in before ours")
			close(bot.errorChan)
			close(bot.resultChan)
			return NewWaiting(bot.Config, bot.Core).ChannelUpdateState(ev, time, bridge)
		}

		deadline := time + bot.VMConfig.GracePeriod
		return waitingAssertion{
			bot.disputableAssertCore,
			deadline,
		}, nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: attemptingAssertion: VM state got unsynchronized"}
	}
}

type waitingAssertion struct {
	*disputableAssertCore
	deadline uint64
}

func (bot waitingAssertion) updateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	bridge.ConfirmDisputableAsserted(
		context.Background(),
		bot.precondition,
		bot.assertion,
	)
	return finalizingAssertion{
		Core:       bot.afterCore,
		Config:     bot.Config,
		ResultChan: bot.resultChan,
		assertion:  bot.assertion,
	}, nil
}

func (bot waitingAssertion) updateState(ev ethconnection.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, challenge.State, error) {
	switch ev.(type) {
	case ethconnection.InitiateChallengeEvent:
		bot.resultChan <- false
		ct, err := defender.New(
			bot.Config,
			machine.NewAssertionDefender(
				bot.assertion,
				bot.precondition,
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

func (bot waitingAssertion) ChainUpdateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot.updateTime(time, bridge)
}

func (bot waitingAssertion) ChainUpdateState(ev ethconnection.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, challenge.State, error) {
	return bot.updateState(ev, time, bridge)
}

func (bot waitingAssertion) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot.updateTime(time, bridge)
}

func (bot waitingAssertion) ChannelUpdateState(ev ethconnection.Event, time uint64, bridge bridge.Bridge) (ChannelState, challenge.State, error) {
	return bot.updateState(ev, time, bridge)
}

type finalizingAssertion struct {
	*core.Core
	*core.Config
	ResultChan chan<- bool
	assertion  *protocol.Assertion
}

func (bot finalizingAssertion) updateState(ev ethconnection.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, challenge.State, error) {
	switch ev := ev.(type) {
	case ethconnection.ConfirmedDisputableAssertEvent:
		if bot.ResultChan != nil {
			bot.ResultChan <- true
		}
		bridge.FinalizedAssertion(
			bot.assertion,
			ev.TxHash[:],
			[][]byte{},
			nil,
		)
		bot.GetCore().DeliverMessagesToVM(bridge)
		return NewWaiting(bot.Config, bot.Core), nil, nil
	default:
		return nil, nil, &Error{nil, "ERROR: FinalizingAssertDefender: VM state got unsynchronized"}
	}
}

func (bot finalizingAssertion) ChainUpdateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot, nil
}

func (bot finalizingAssertion) ChainUpdateState(ev ethconnection.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, challenge.State, error) {
	return bot.updateState(ev, time, bridge)
}

func (bot finalizingAssertion) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot, nil
}

func (bot finalizingAssertion) ChannelUpdateState(ev ethconnection.Event, time uint64, bridge bridge.Bridge) (ChannelState, challenge.State, error) {
	return bot.updateState(ev, time, bridge)
}
