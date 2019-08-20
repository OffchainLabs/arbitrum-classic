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
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/state"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type Validator struct {
	Name        string
	actions     chan func(*Validator, bridge.Bridge)
	maybeAssert chan bool

	// Run loop only
	bot                      state.State
	challengeBot             challenge.State
	latestHeader             *types.Header
	pendingDisputableRequest *state.DisputableAssertionRequest
	isCreated                bool
}

func NewValidator(name string, address common.Address, balance *protocol.BalanceTracker, config *valmessage.VMConfiguration, machine machine.Machine, challengeEverything bool, maxCallSteps int32) *Validator {
	actions := make(chan func(*Validator, bridge.Bridge), 100)
	maybeAssert := make(chan bool, 100)
	c := core.NewCore(
		machine,
		balance,
	)

	// TODO: latestHeader starts as nil which isn't valid. This needs to be properly initialized
	valConfig := core.NewValidatorConfig(address, config, challengeEverything, maxCallSteps)
	return &Validator{
		name,
		actions,
		maybeAssert,
		state.NewWaiting(valConfig, c),
		nil,
		nil,
		nil,
		false,
	}
}

func (validator *Validator) RequestCall(msg protocol.Message) (<-chan value.Value, <-chan error) {
	resultChan := make(chan value.Value, 1)
	errChan := make(chan error, 1)
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		if !validator.canRun() {
			errChan <- errors.New("Cannot call when VM is not running")
			return
		}
		c := validator.bot.GetCore()
		updatedState := c.GetMachine().Clone()
		startTime := validator.latestHeader.Number.Uint64()
		messageHash := solsha3.SoliditySHA3(
			solsha3.Bytes32(msg.Destination),
			solsha3.Bytes32(msg.Data.Hash()),
			solsha3.Uint256(msg.Currency),
			msg.TokenType[:],
		)
		msgHashInt := new(big.Int).SetBytes(messageHash[:])
		val, _ := value.NewTupleFromSlice([]value.Value{
			msg.Data,
			value.NewIntValue(new(big.Int).SetUint64(validator.latestHeader.Time)),
			value.NewIntValue(validator.latestHeader.Number),
			value.NewIntValue(msgHashInt),
		})
		callingMessage := protocol.Message{
			Data:        val.Clone(),
			TokenType:   msg.TokenType,
			Currency:    msg.Currency,
			Destination: msg.Destination,
		}
		maxCallSteps := validator.bot.GetConfig().MaxCallSteps
		go func() {
			updatedState.SendOffchainMessages([]protocol.Message{callingMessage})
			assertion := updatedState.ExecuteAssertion(
				maxCallSteps,
				[2]uint64{startTime, startTime + 1},
			)
			results := assertion.Logs
			if len(results) == 0 {
				errChan <- errors.New("call produced no output")
				return
			}
			lastLogVal := results[len(results)-1]
			lastLog, err := evm.ProcessLog(lastLogVal)
			if err != nil {
				errChan <- err
				return
			}
			logHash := lastLog.GetEthMsg().Data.TxHash
			if !bytes.Equal(logHash[:], messageHash) {
				// Last produced log is not the call we sent
				errChan <- errors.New("call took too long to execute")
				return
			}

			resultChan <- results[len(results)-1]
		}()
	}

	return resultChan, errChan
}

func (validator *Validator) PendingMessageCount() chan uint64 {
	resultChan := make(chan uint64, 1)
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		c := validator.bot.GetCore()
		resultChan <- c.GetMachine().PendingMessageCount()
	}
	return resultChan
}

func (validator *Validator) HasOpenAssertion() chan bool {
	resultChan := make(chan bool, 1)
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		bot, ok := validator.bot.(state.Waiting)
		if !ok {
			resultChan <- false
		} else {
			resultChan <- bot.HasOpenAssertion()
		}
	}

	return resultChan
}

func (validator *Validator) canRun() bool {
	c := validator.bot.GetCore()
	return validator.isCreated && c.GetMachine().CurrentStatus() == machine.Extensive
}

func (validator *Validator) CanRun() chan bool {
	resultChan := make(chan bool, 1)
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		resultChan <- validator.canRun()
	}
	return resultChan
}

func (validator *Validator) CanContinueRunning() chan bool {
	resultChan := make(chan bool, 1)
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		if !validator.canRun() {
			resultChan <- false
		} else if validator.latestHeader == nil {
			resultChan <- false
		} else {
			currentTime := validator.latestHeader.Number.Uint64()
			resultChan <- !machine.IsMachineBlocked(validator.bot.GetCore().GetMachine(), currentTime)
		}

	}
	return resultChan
}

type VMStateData struct {
	MachineState [32]byte
	Config       valmessage.VMConfiguration
}

func (validator *Validator) RequestVMState() <-chan VMStateData {
	resultChan := make(chan VMStateData)
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		c := validator.bot.GetCore()
		machineHash := c.GetMachine().Hash()
		resultChan <- VMStateData{
			MachineState: machineHash,
			Config:       *validator.bot.GetConfig().VMConfig,
		}
	}
	return resultChan
}

func (validator *Validator) RequestDisputableAssertion(length uint64) (<-chan bool, <-chan error) {
	resultChan := make(chan bool)
	errChan := make(chan error)
	validator.actions <- func(validator *Validator, b bridge.Bridge) {
		if !validator.canRun() {
			errChan <- errors.New("Can't disputable assert when not running")
			return
		}
		c := validator.bot.GetCore()
		mClone := c.GetMachine().Clone()
		maxSteps := validator.bot.GetConfig().VMConfig.MaxExecutionStepCount
		startTime := validator.latestHeader.Number.Uint64()
		go func() {
			endTime := startTime + length
			tb := [2]uint64{startTime, endTime}
			beforeHash := mClone.Hash()
			assertion := mClone.ExecuteAssertion(int32(maxSteps), tb)
			spentBalance := protocol.NewBalanceTrackerFromMessages(assertion.OutMsgs)
			balance := c.GetBalance()
			_ = balance.SpendAll(spentBalance)

			pre := &protocol.Precondition{
				BeforeHash:    beforeHash,
				TimeBounds:    tb,
				BeforeBalance: spentBalance,
				BeforeInbox:   mClone.InboxHash(),
			}
			request := &state.DisputableAssertionRequest{
				AfterCore:    core.NewCore(mClone, balance),
				Precondition: pre,
				Assertion:    assertion,
				ResultChan:   resultChan,
				ErrorChan:    errChan,
			}
			validator.actions <- func(validator *Validator, b bridge.Bridge) {
				validator.pendingDisputableRequest = request
				validator.maybeAssert <- true
			}
		}()
	}
	return resultChan, errChan
}

type unanimousUpdateRequest struct {
	valmessage.UnanimousRequestData

	NewMessages []protocol.Message

	Machine   machine.Machine
	Assertion *protocol.Assertion

	ShouldFinalize func(*protocol.Assertion) bool

	ResultChan chan<- valmessage.UnanimousUpdateResults
	ErrChan    chan<- error
}

func (validator *Validator) InitiateUnanimousRequest(
	length uint64,
	messages []protocol.Message,
	messageHashes [][]byte,
	final bool,
	maxSteps int32,
	shouldFinalize func(*protocol.Assertion) bool,
) (
	<-chan valmessage.UnanimousRequest,
	<-chan valmessage.UnanimousUpdateResults,
	<-chan error,
) {
	unanRequestChan := make(chan valmessage.UnanimousRequest, 1)
	updateResultChan := make(chan valmessage.UnanimousUpdateResults, 1)
	errChan := make(chan error, 1)

	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		if !validator.canRun() {
			errChan <- errors.New("Can't unanimous assert when not running")
			return
		}
		bot, ok := validator.bot.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("recieved initiate unanimous request, but was in the wrong state to handle it: %T", validator.bot)
			return
		}
		newMessages := make([]protocol.Message, 0, len(messages))
		messageRecords := make([]protocol.Message, 0, len(messages))
		for i, msg := range messages {
			msgHashInt := new(big.Int).SetBytes(messageHashes[i])
			val, _ := value.NewTupleFromSlice([]value.Value{
				msg.Data,
				value.NewIntValue(new(big.Int).SetUint64(validator.latestHeader.Time)),
				value.NewIntValue(validator.latestHeader.Number),
				value.NewIntValue(msgHashInt),
			})
			newMessages = append(newMessages, protocol.Message{
				Data:        val,
				TokenType:   msg.TokenType,
				Currency:    msg.Currency,
				Destination: msg.Destination,
			})
			messageRecords = append(messageRecords, protocol.Message{
				Data:        val.Clone(),
				TokenType:   msg.TokenType,
				Currency:    msg.Currency,
				Destination: msg.Destination,
			})
		}
		timeBounds := [2]uint64{validator.latestHeader.Number.Uint64(), validator.latestHeader.Number.Uint64() + length}
		seqNum := bot.OffchainContext(timeBounds, final)
		clonedMachine := bot.GetCore().GetMachine().Clone()
		requestData := valmessage.UnanimousRequestData{
			BeforeHash:  bot.OrigHash(),
			BeforeInbox: bot.OrigInboxHash(),
			SequenceNum: seqNum,
			TimeBounds:  timeBounds,
		}

		unanRequestChan <- valmessage.UnanimousRequest{UnanimousRequestData: requestData, NewMessages: messageRecords}
		go func() {
			clonedMachine.SendOffchainMessages(newMessages)
			assertion := clonedMachine.ExecuteAssertion(
				maxSteps,
				timeBounds,
			)
			validator.requestUnanimousUpdate(unanimousUpdateRequest{
				UnanimousRequestData: requestData,
				NewMessages:          newMessages,
				Machine:              clonedMachine,
				Assertion:            assertion,
				ShouldFinalize:       shouldFinalize,
				ResultChan:           updateResultChan,
				ErrChan:              errChan,
			})
		}()
	}
	return unanRequestChan, updateResultChan, errChan
}

func (validator *Validator) RequestFollowUnanimous(
	request valmessage.UnanimousRequestData,
	messages []protocol.Message,
	maxSteps int32,
	shouldFinalize func(*protocol.Assertion) bool,
) (<-chan valmessage.UnanimousUpdateResults, <-chan error) {
	resultChan := make(chan valmessage.UnanimousUpdateResults, 1)
	errChan := make(chan error, 1)
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		if !validator.canRun() {
			errChan <- errors.New("Can't unanimous assert when not running")
			return
		}
		bot, ok := validator.bot.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("recieved follow unanimous request, but was in the wrong state to handle it: %T", validator.bot)
			return
		}

		if err := bot.ValidateUnanimousRequest(request); err != nil {
			errChan <- err
			return
		}

		_ = bot.OffchainContext(request.TimeBounds, request.SequenceNum == math.MaxUint64)
		clonedMachine := bot.GetCore().GetMachine().Clone()
		go func() {
			clonedMachine.SendOffchainMessages(messages)
			assertion := clonedMachine.ExecuteAssertion(
				maxSteps,
				request.TimeBounds,
			)
			validator.requestUnanimousUpdate(unanimousUpdateRequest{
				UnanimousRequestData: request,
				NewMessages:          messages,
				Machine:              clonedMachine,
				Assertion:            assertion,
				ShouldFinalize:       shouldFinalize,
				ResultChan:           resultChan,
				ErrChan:              errChan,
			})
		}()
	}
	return resultChan, errChan
}

func (validator *Validator) requestUnanimousUpdate(request unanimousUpdateRequest) {
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		bot, ok := validator.bot.(state.Waiting)
		if !ok {
			request.ErrChan <- fmt.Errorf("recieved unanimous update request, but was in the wrong state to handle it: %T", validator.bot)
			return
		}

		newBot, err := bot.PreparePendingUnanimous(
			request.Assertion,
			request.NewMessages,
			request.Machine,
			request.SequenceNum,
			request.TimeBounds,
			request.ShouldFinalize,
		)

		if err != nil {
			request.ErrChan <- err
			return
		}
		request.ResultChan <- newBot.ProposalResults()
		validator.bot = newBot
	}
}

func (validator *Validator) ConfirmOffchainUnanimousAssertion(
	request valmessage.UnanimousRequestData,
	signatures [][]byte,
	canClose bool,
) (<-chan bool, <-chan error) {
	resultChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		bot, ok := validator.bot.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("recieved unanimous confirm request, but was in the wrong state to handle it: %T", validator.bot)
			return
		}
		if err := bot.ValidateUnanimousAssertion(request); err != nil {
			errChan <- err
			return
		}

		bridge.AddedNewMessages(bot.ProposedMessageCount())

		proposalResults := bot.ProposalResults()
		newBot, err := bot.FinalizePendingUnanimous(signatures)
		if err != nil {
			errChan <- err
			return
		}

		bridge.FinalizedAssertion(
			nil,
			[]byte{},
			signatures,
			&proposalResults,
		)

		if request.SequenceNum == math.MaxUint64 {
			if canClose {
				newBot.CloseUnanimous(bridge)
			}
			// Can only error if there is no pending assertion which is guaranteed here
			newBot2, _ := newBot.ClosingUnanimous(resultChan, errChan)
			validator.bot = newBot2
		} else {
			validator.bot = newBot
			resultChan <- true
		}
	}
	return resultChan, errChan
}

func (validator *Validator) CloseUnanimousAssertionRequest() (<-chan bool, <-chan error) {
	resultChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		bot, ok := validator.bot.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("can't close unanimous request, but was in the wrong state to handle it: %T", validator.bot)
			return
		}
		bot.CloseUnanimous(bridge)
		newBot, err := bot.ClosingUnanimous(resultChan, errChan)
		if err != nil {
			errChan <- err
			return
		}
		validator.bot = newBot
	}
	return resultChan, errChan
}

func (validator *Validator) ClosingUnanimousAssertionRequest() (<-chan bool, <-chan error) {
	resultChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	validator.actions <- func(validator *Validator, bridge bridge.Bridge) {
		bot, ok := validator.bot.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("can't close unanimous request. Validator was in the wrong state to handle it: %T", validator.bot)
			return
		}
		newBot, err := bot.ClosingUnanimous(resultChan, errChan)
		if err != nil {
			errChan <- err
			return
		}
		validator.bot = newBot
	}
	return resultChan, errChan
}

func (validator *Validator) Run(recvChan <-chan ethbridge.Notification, bridge bridge.Bridge) {
	go func() {
		defer fmt.Printf("%v: Exiting\n", validator.Name)
		for {
			select {
			case notification, ok := <-recvChan:
				// fmt.Printf("Got valmessage %T: %v\n", event, event)
				if !ok {
					fmt.Printf("%v: Error in recvChan\n", validator.Name)
					return
				}

				newHeader := notification.Header
				if validator.latestHeader == nil || newHeader.Number.Uint64() >= validator.latestHeader.Number.Uint64() && newHeader.Hash() != validator.latestHeader.Hash() {
					validator.latestHeader = newHeader
					validator.timeUpdate(bridge)

					if validator.pendingDisputableRequest != nil {
						pre := validator.pendingDisputableRequest.Precondition
						if !validator.bot.GetCore().ValidateAssertion(pre, newHeader.Number.Uint64()) {
							validator.pendingDisputableRequest.ErrorChan <- errors.New("Precondition was invalidated")
							close(validator.pendingDisputableRequest.ErrorChan)
							close(validator.pendingDisputableRequest.ResultChan)
							validator.pendingDisputableRequest = nil
						}
					}
				}

				switch ev := notification.Event.(type) {
				case ethbridge.NewTimeEvent:
					break
				case ethbridge.VMCreatedEvent:
					validator.isCreated = true
				case ethbridge.VMEvent:
					validator.eventUpdate(ev, notification.Header, bridge)
				case ethbridge.MessageDeliveredEvent:
					validator.bot.SendMessageToVM(ev.Msg)
				default:
					panic("Should never recieve other kinds of events")
				}
			case action := <-validator.actions:
				action(validator, bridge)
			case <-validator.maybeAssert:
			}

			if bot, ok := validator.bot.(state.Waiting); ok && validator.pendingDisputableRequest != nil {
				validator.bot = bot.AttemptAssertion(context.Background(), *validator.pendingDisputableRequest, bridge)
				validator.pendingDisputableRequest = nil
			}
		}
	}()
}

func (validator *Validator) timeUpdate(bridge bridge.Bridge) {
	if validator.challengeBot != nil {
		newBot, err := validator.challengeBot.UpdateTime(validator.latestHeader.Number.Uint64(), bridge)
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
			return
		}
		validator.challengeBot = newBot
	}
	newBot, err := validator.bot.UpdateTime(validator.latestHeader.Number.Uint64(), bridge)
	if err != nil {
		fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
		return
	}
	validator.bot = newBot
}

func (validator *Validator) eventUpdate(ev ethbridge.VMEvent, header *types.Header, bridge bridge.Bridge) {
	if ev.GetIncomingMessageType() == ethbridge.ChallengeMessage {
		if validator.challengeBot == nil {
			panic("challengeBot can't be nil if challenge message is recieved")
		}

		newBot, err := validator.challengeBot.UpdateState(ev, header.Number.Uint64(), bridge)
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
			return
		}
		validator.challengeBot = newBot
	} else {
		newBot, challengeBot, err := validator.bot.UpdateState(ev, header.Number.Uint64(), bridge)
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, validator.bot)
			return
		}
		validator.bot = newBot
		if challengeBot != nil {
			validator.challengeBot = challengeBot
		}
	}
}
