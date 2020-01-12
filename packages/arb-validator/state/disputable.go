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

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/disputable"
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
	machine         machine.Machine
	newMessageCount uint64
	Assertion       *protocol.ExecutionAssertion
	timeBounds      *protocol.TimeBounds
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
	assertion   *protocol.ExecutionAssertion
	sequenceNum uint64
	signatures  [][]byte

	timeBounds *protocol.TimeBounds
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
		timeBounds:  nil,
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

func (bot Waiting) CloseUnanimous(bridge bridge.Bridge) (*types.Receipt, error) {
	if bot.sequenceNum == math.MaxUint64 {
		return bridge.FinalizedUnanimousAssert(
			context.Background(),
			bot.GetCore().GetMachine().InboxHash().Hash(),
			bot.assertion,
			bot.signatures,
		)
	}

	return bridge.PendingUnanimousAssert(
		context.Background(),
		bot.GetCore().GetMachine().InboxHash().Hash(),
		bot.assertion,
		bot.sequenceNum,
		bot.signatures,
	)
}

func (bot Waiting) ClosingUnanimous(retChan chan<- bool, errChan chan<- error) (ChannelState, error) {
	// If there is no active unanimous assertion, there is nothing to close
	// TODO: Validator should refuse to unanimous assert again from the same start point
	if bot.assertion == nil {
		err := errors.New("couldn't close since no ExecutionAssertion is open")
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

func (bot Waiting) AttemptAssertion(ctx context.Context, request disputable.AssertionRequest, bridge bridge.ArbVMBridge) (ChainState, error) {
	_, err := bridge.PendingDisputableAssert(
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
	}}, err
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
		if request.TimeBounds.StartTime < bot.timeBounds.StartTime {
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
		if request.TimeBounds.StartTime < bot.timeBounds.StartTime {
			return errors.New("unanimous assertion starting time bound may only increase")
		}
	}

	if request.SequenceNum < bot.proposed.sequenceNum {
		return errors.New("recieved unanimous update with invalid sequence number")
	}
	return nil
}

func (bot Waiting) PreparePendingUnanimous(
	newAssertion *protocol.ExecutionAssertion,
	messages []protocol.Message,
	machine machine.Machine,
	sequenceNum uint64,
	timeBounds *protocol.TimeBounds,
	shouldFinalize func(*protocol.ExecutionAssertion) bool,
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
		return Waiting{}, errors.New("no pending ExecutionAssertion")
	}

	return Waiting{
		Config:   bot.Config,
		proposed: nil,
		accepted: core.NewCore(
			bot.proposed.machine,
		),
		assertion:   bot.proposed.Assertion,
		sequenceNum: bot.proposed.sequenceNum,
		signatures:  signatures,
		timeBounds:  bot.proposed.timeBounds,
		orig:        bot.orig,
	}, nil
}

func (bot Waiting) updateState(ev arbbridge.Event, time uint64, brdg bridge.ArbVMBridge) (ChainState, error) {
	switch ev := ev.(type) {
	case ethbridge.PendingDisputableAssertionEvent:
		c := bot.GetCore()
		var inboxVal value.Value
		if c.GetMachine().InboxHash().Hash() != ev.Precondition.BeforeInboxValue() {
			return nil, errors.New("waiting observer has incorrect valmessage")
		}
		updatedState := c.GetMachine().Clone()
		assertion := updatedState.ExecuteAssertion(
			int32(ev.Assertion.NumSteps),
			ev.Precondition.TimeBounds,
		)
		if !assertion.Stub().Equals(ev.Assertion) || bot.ChallengeEverything {
			_, err := brdg.InitiateChallenge(
				context.Background(),
				ev.Precondition,
				ev.Assertion,
			)
			if err != nil {
				return nil, &bridge.Error{err, "ERROR: InitiateChallenge: failed", false}
			}
		}
		return watchingAssertion{
			c,
			bot.Config,
			inboxVal,
			core.NewCore(updatedState),
			ev.Deadline,
			ev.Precondition,
			assertion,
		}, nil
	default:
		return nil, &Error{nil, fmt.Sprintf("ERROR: Waiting: VM state got unsynchronized with valmessage %T", ev)}
	}
}

func (bot Waiting) ChainUpdateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot, nil
}

func (bot Waiting) ChainUpdateState(ev arbbridge.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot.updateState(ev, time, bridge)
}

func (bot Waiting) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot, nil
}

func (bot Waiting) ChannelUpdateState(ev arbbridge.Event, time uint64, bridge bridge.Bridge) (ChannelState, error) {
	switch ev := ev.(type) {
	case ethbridge.PendingUnanimousAssertEvent:
		if bot.accepted == nil || ev.SequenceNum > bot.sequenceNum {
			return nil, errors.New("waiting observer saw pending unanimous assertion that it doesn't remember")
		} else if ev.SequenceNum < bot.sequenceNum {
			_, err := bot.CloseUnanimous(bridge)
			if err != nil {
				// add error handling
				return nil, err
			}
			newBot, err := bot.ClosingUnanimous(nil, nil)
			return newBot, err
		} else {
			return waitingOffchainClosing{
				bot.Config,
				bot.GetCore(),
				bot.assertion,
				ev.Deadline,
				nil,
				nil,
			}, nil
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
	assertion    *protocol.ExecutionAssertion
}

func (bot watchingAssertion) SendMessageToVM(msg protocol.Message) {
	bot.Core.SendMessageToVM(msg)
	bot.pending.SendMessageToVM(msg)
}

func (bot watchingAssertion) updateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	return finalizingAssertion{
		Core:       bot.pending,
		Config:     bot.Config,
		ResultChan: nil,
		assertion:  bot.assertion,
	}, nil
}

func (bot watchingAssertion) updateState(ev arbbridge.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	switch ev := ev.(type) {
	case ethbridge.ChallengeLaunchedEvent:
		var err error
		if ev.Challenger == bot.Address {
			err = bridge.Challenge(
				context.Background(),
				ev.ChallengeAddress,
				bot.precondition,
				bot.GetMachine().Clone(),
			)
		} else {
			err = bridge.ObserveChallenge(
				context.Background(),
				ev.ChallengeAddress,
			)
		}
		return NewWaiting(bot.Config, bot.Core), err

	default:
		return nil, &Error{nil, "ERROR: WaitingValidObserver: VM state got unsynchronized"}
	}
}

func (bot watchingAssertion) ChainUpdateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot.updateTime(time, bridge)
}

func (bot watchingAssertion) ChainUpdateState(ev arbbridge.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot.updateState(ev, time, bridge)
}

func (bot watchingAssertion) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot.updateTime(time, bridge)
}

func (bot watchingAssertion) ChannelUpdateState(ev arbbridge.Event, time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot.updateState(ev, time, bridge)
}

type disputableAssertCore struct {
	*core.Core
	*core.Config
	afterCore    *core.Core
	precondition *protocol.Precondition
	assertion    *protocol.ExecutionAssertion
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

func (bot attemptingAssertion) updateState(ev arbbridge.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	switch ev := ev.(type) {
	case ethbridge.PendingDisputableAssertionEvent:
		if ev.Asserter != bot.Address {
			bot.errorChan <- fmt.Errorf("attemptingAssertion: Other ExecutionAssertion by %v got in before ours by %v", hexutil.Encode(ev.Asserter[:]), hexutil.Encode(bot.Address[:]))
			close(bot.errorChan)
			close(bot.resultChan)
			return NewWaiting(bot.Config, bot.Core).ChainUpdateState(ev, time, bridge)
		}

		return waitingAssertion{
			bot.disputableAssertCore,
			ev.Deadline,
		}, nil
	default:
		return nil, &Error{nil, "ERROR: attemptingAssertion: VM state got unsynchronized"}
	}
}

func (bot attemptingAssertion) ChainUpdateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot, nil
}

func (bot attemptingAssertion) ChainUpdateState(ev arbbridge.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot.updateState(ev, time, bridge)
}

func (bot attemptingAssertion) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot, nil
}

func (bot attemptingAssertion) ChannelUpdateState(ev arbbridge.Event, time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot.updateState(ev, time, bridge)
}

type waitingAssertion struct {
	*disputableAssertCore
	deadline uint64
}

func (bot waitingAssertion) updateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	if time <= bot.deadline {
		return bot, nil
	}

	_, err := bridge.ConfirmDisputableAsserted(
		context.Background(),
		bot.precondition,
		bot.assertion,
	)
	return finalizingAssertion{
		Core:       bot.afterCore,
		Config:     bot.Config,
		ResultChan: bot.resultChan,
		assertion:  bot.assertion,
	}, err
}

func (bot waitingAssertion) updateState(ev arbbridge.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	switch ev := ev.(type) {
	case ethbridge.ChallengeLaunchedEvent:
		bot.resultChan <- false
		err := bridge.DefendChallenge(
			context.Background(),
			ev.ChallengeAddress,
			machine.NewAssertionDefender(
				bot.assertion,
				bot.precondition,
				bot.GetMachine().Clone(),
			),
		)
		return NewWaiting(bot.Config, bot.Core), err

	default:
		return nil, &Error{nil, "disputable.go ERROR: waitingAssertion: VM state got unsynchronized"}
	}
}

func (bot waitingAssertion) ChainUpdateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot.updateTime(time, bridge)
}

func (bot waitingAssertion) ChainUpdateState(ev arbbridge.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot.updateState(ev, time, bridge)
}

func (bot waitingAssertion) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot.updateTime(time, bridge)
}

func (bot waitingAssertion) ChannelUpdateState(ev arbbridge.Event, time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot.updateState(ev, time, bridge)
}

type finalizingAssertion struct {
	*core.Core
	*core.Config
	ResultChan chan<- bool
	assertion  *protocol.ExecutionAssertion
}

func (bot finalizingAssertion) updateState(ev arbbridge.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	switch ev := ev.(type) {
	case ethbridge.ConfirmedDisputableAssertEvent:
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
		return NewWaiting(bot.Config, bot.Core), nil
	default:
		return nil, &Error{nil, "ERROR: FinalizingAssertDefender: VM state got unsynchronized"}
	}
}

func (bot finalizingAssertion) ChainUpdateTime(time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot, nil
}

func (bot finalizingAssertion) ChainUpdateState(ev arbbridge.Event, time uint64, bridge bridge.ArbVMBridge) (ChainState, error) {
	return bot.updateState(ev, time, bridge)
}

func (bot finalizingAssertion) ChannelUpdateTime(time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot, nil
}

func (bot finalizingAssertion) ChannelUpdateState(ev arbbridge.Event, time uint64, bridge bridge.Bridge) (ChannelState, error) {
	return bot.updateState(ev, time, bridge)
}
