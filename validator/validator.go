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
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-validator/valmessage"
	"github.com/pkg/errors"
	"math"
	"math/big"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/vm"
)

type validatorState interface {
	UpdateTime(uint64) (validatorState, []valmessage.OutgoingMessage, error)
	UpdateState(valmessage.IncomingMessage, uint64) (validatorState, challengeState, []valmessage.OutgoingMessage, error)

	SendMessageToVM(msg protocol.Message)
	GetCore() *validatorCore
	GetConfig() *validatorConfig
}

type challengeState interface {
	UpdateTime(uint64) (challengeState, []valmessage.OutgoingMessage, error)
	UpdateState(valmessage.IncomingMessage, uint64) (challengeState, []valmessage.OutgoingMessage, error)
}

type Error struct {
	err     error
	message string
}

func (e *Error) Error() string {
	if e.err != nil {
		return fmt.Sprintf("%v: %v", e.message, e.err)
	} else {
		return e.message
	}
}

type Validator struct {
	Name        string
	requests    chan interface{}
	maybeAssert chan bool

	// Run loop only
	bot                      validatorState
	challengeBot             challengeState
	latestHeader             *types.Header
	pendingDisputableRequest *valmessage.DisputableAssertionRequest
}

func NewValidator(name string, address common.Address, inbox *protocol.Inbox, balance *protocol.BalanceTracker, config *valmessage.VMConfiguration, machine *vm.Machine, challengeEverything bool) *Validator {
	requests := make(chan interface{}, 10)
	maybeAssert := make(chan bool, 100)
	core := &validatorCore{
		inbox,
		balance,
		machine,
	}

	// TODO: latestHeader starts as nil which isn't valid. This needs to be properly initialized
	valConfig := NewValidatorConfig(address, config, challengeEverything)
	return &Validator{
		name,
		requests,
		maybeAssert,
		NewWaitingObserver(valConfig, core),
		nil,
		nil,
		nil,
	}
}

func (validator *Validator) RequestCall(msg protocol.Message) (<-chan value.Value, <-chan error) {
	resultChan := make(chan value.Value, 1)
	errorChan := make(chan error, 1)
	validator.requests <- valmessage.CallRequest{
		Message:    msg,
		ResultChan: resultChan,
		ErrorChan:  errorChan,
	}
	return resultChan, errorChan
}

func (validator *Validator) HasPendingMessages() chan bool {
	retChan := make(chan bool, 1)
	validator.requests <- valmessage.PendingMessageCheck{ResultChan: retChan}
	return retChan
}

func (validator *Validator) RequestVMState() <-chan valmessage.VMStateData {
	resultChan := make(chan valmessage.VMStateData)
	validator.requests <- valmessage.VMStateRequest{ResultChan: resultChan}
	return resultChan
}

func (validator *Validator) RequestDisputableAssertion(length uint64, includePendingMessages bool) <-chan bool {
	resultChan := make(chan bool)
	validator.requests <- valmessage.DisputableDefenderRequest{
		Length:                 length,
		IncludePendingMessages: includePendingMessages,
		ResultChan:             resultChan,
	}
	return resultChan
}

func (validator *Validator) InitiateUnanimousRequest(
	length uint64,
	messages []protocol.Message,
	final bool,
) (
	<-chan valmessage.UnanimousRequest,
	<-chan valmessage.UnanimousUpdateResults,
	<-chan error,
) {
	unanRequestChan := make(chan valmessage.UnanimousRequest, 1)
	updateResultChan := make(chan valmessage.UnanimousUpdateResults, 1)
	errChan := make(chan error, 1)
	validator.requests <- valmessage.InitiateUnanimousRequest{
		TimeLength:  length,
		NewMessages: messages,
		Final:       final,
		RequestChan: unanRequestChan,
		ResultChan:  updateResultChan,
		ErrChan:     errChan,
	}
	return unanRequestChan, updateResultChan, errChan
}

func (validator *Validator) RequestFollowUnanimous(
	request valmessage.UnanimousRequestData,
	messages []protocol.Message,
) (<-chan valmessage.UnanimousUpdateResults, <-chan error) {
	resultChan := make(chan valmessage.UnanimousUpdateResults, 1)
	errChan := make(chan error, 1)
	validator.requests <- valmessage.FollowUnanimousRequest{
		UnanimousRequestData: request,
		NewMessages:          messages,
		ResultChan:           resultChan,
		ErrChan:              errChan,
	}
	return resultChan, errChan
}

func (validator *Validator) ConfirmOffchainUnanimousAssertion(
	unanRequest valmessage.UnanimousRequestData,
	signatures []valmessage.Signature,
) (<-chan bool, <-chan error) {
	resultChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	validator.requests <- valmessage.UnanimousConfirmRequest{
		UnanimousRequestData: unanRequest,
		Signatures:           signatures,
		ResultChan:           resultChan,
		ErrChan:              errChan,
	}
	return resultChan, errChan
}

func (validator *Validator) CloseUnanimousAssertionRequest() <-chan bool {
	resultChan := make(chan bool, 1)
	validator.requests <- valmessage.CloseUnanimousAssertionRequest{
		ResultChan: resultChan,
	}
	return resultChan
}

func (validator *Validator) Run(recvChan <-chan valmessage.IncomingValidatorMessage, sendChan chan<- valmessage.OutgoingMessage) {
	go func() {
		defer fmt.Printf("%v: Exiting\n", validator.Name)
		defer close(sendChan)
		for {
			select {
			case event, ok := <-recvChan:
				//fmt.Printf("Got valmessage %T: %v\n", event, event)
				if !ok {
					fmt.Printf("%v: Error in recvChan\n", validator.Name)
					return
				}

				newHeader := event.GetHeader()
				if validator.latestHeader == nil || newHeader.Number.Uint64() >= validator.latestHeader.Number.Uint64() && newHeader.Hash() != validator.latestHeader.Hash() {
					validator.latestHeader = newHeader
					validator.timeUpdate(sendChan)

					if validator.pendingDisputableRequest != nil {
						pre := validator.pendingDisputableRequest.GetPrecondition()
						if !validator.bot.GetCore().ValidateAssertion(pre, newHeader.Number.Uint64()) {
							validator.pendingDisputableRequest.NotifyInvalid()
							validator.pendingDisputableRequest = nil
						}
					}
				}

				switch ev := event.(type) {
				case valmessage.TimeUpdateMessage:
					break
				case valmessage.BridgeMessage:
					validator.eventUpdate(ev, sendChan)
				case valmessage.IncomingMessageMessage:
					validator.bot.SendMessageToVM(ev.Msg)

					// Invalidate assertions that included pending messages
					if validator.pendingDisputableRequest != nil && validator.pendingDisputableRequest.IncludedPendingInbox() {
						validator.pendingDisputableRequest.NotifyInvalid()
						validator.pendingDisputableRequest = nil
					}
				default:
					panic("Should never recieve other kinds of events")
				}
				validator.tryToAssert(sendChan)
			case request := <-validator.requests:
				switch request := request.(type) {
				case valmessage.InitiateUnanimousRequest:
					if bot, ok := validator.bot.(WaitingObserver); ok {
						newMessages := make([]protocol.Message, 0, len(request.NewMessages))
						messageRecords := make([]protocol.Message, 0, len(request.NewMessages))
						for _, msg := range request.NewMessages {
							messageHash := solsha3.SoliditySHA3(
								solsha3.Bytes32(msg.Destination),
								solsha3.Bytes32(msg.Data.Hash()),
								solsha3.Uint256(msg.Currency),
								msg.TokenType[:],
							)
							msgHashInt := new(big.Int).SetBytes(messageHash[:])
							val, _ := value.NewTupleFromSlice([]value.Value{
								msg.Data,
								value.NewIntValue(validator.latestHeader.Time),
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
						timeBounds := [2]uint64{validator.latestHeader.Number.Uint64(), validator.latestHeader.Number.Uint64() + request.TimeLength}
						mq, tb, seqNum := bot.OffchainContext(newMessages, timeBounds, request.Final)
						clonedCore := bot.GetCore().Clone()
						requestData := valmessage.UnanimousRequestData{
							BeforeHash:  clonedCore.machine.Hash(),
							BeforeInbox: clonedCore.inbox.Receive().Hash(),
							SequenceNum: seqNum,
							TimeBounds:  tb,
						}

						request.RequestChan <- valmessage.UnanimousRequest{UnanimousRequestData: requestData, NewMessages: messageRecords}
						go func() {
							newCore, assertion := clonedCore.OffchainAssert(mq, timeBounds)
							validator.requests <- valmessage.UnanimousUpdateRequest{
								UnanimousRequestData: requestData,
								NewMessages:          newMessages,
								Inbox:                newCore.inbox,
								Machine:              newCore.machine,
								Assertion:            assertion,
								ResultChan:           request.ResultChan,
								ErrChan:              request.ErrChan,
							}
						}()
					} else {
						request.ErrChan <- fmt.Errorf("recieved initiate unanimous request, but was in the wrong state to handle it: %T", validator.bot)
						break
					}
				case valmessage.FollowUnanimousRequest:
					if bot, ok := validator.bot.(WaitingObserver); ok {
						if err := bot.validateUnanimousRequest(request.UnanimousRequestData); err != nil {
							request.ErrChan <- err
							break
						}

						mq, _, _ := bot.OffchainContext(request.NewMessages, request.TimeBounds, request.SequenceNum == math.MaxUint64)
						clonedCore := bot.GetCore().Clone()
						go func() {
							newCore, assertion := clonedCore.OffchainAssert(mq, request.TimeBounds)
							validator.requests <- valmessage.UnanimousUpdateRequest{
								UnanimousRequestData: request.UnanimousRequestData,
								NewMessages:          request.NewMessages,
								Inbox:                newCore.inbox,
								Machine:              newCore.machine,
								Assertion:            assertion,
								ResultChan:           request.ResultChan,
								ErrChan:              request.ErrChan,
							}
						}()
					} else {
						request.ErrChan <- fmt.Errorf("recieved follow unanimous request, but was in the wrong state to handle it: %T", validator.bot)
						break
					}
				case valmessage.UnanimousUpdateRequest:
					if bot, ok := validator.bot.(WaitingObserver); ok {
						if err := bot.validateUnanimousRequest(request.UnanimousRequestData); err != nil {
							request.ErrChan <- err
							break
						}

						newBot, err := bot.PreparePendingUnanimous(request)
						if err != nil {
							request.ErrChan <- err
							break
						}
						request.ResultChan <- newBot.ProposalResults()
						validator.bot = newBot

					} else {
						request.ErrChan <- fmt.Errorf("recieved unanimous update request, but was in the wrong state to handle it: %T", validator.bot)
						break
					}
				case valmessage.UnanimousConfirmRequest:
					if bot, ok := validator.bot.(WaitingObserver); ok {
						if err := bot.validateUnanimousRequest(request.UnanimousRequestData); err != nil {
							request.ErrChan <- err
							break
						}

						newBot, proposal, err := bot.FinalizePendingUnanimous(request.Signatures)
						if err != nil {
							request.ErrChan <- err
							break
						}
						validator.bot = newBot
						sendChan <- valmessage.FinalizedAssertion{Assertion: proposal.assertion, NewLogCount: proposal.newLogCount}
						request.ResultChan <- true
					} else {
						request.ErrChan <- fmt.Errorf("recieved unanimous confirm request, but was in the wrong state to handle it: %T", validator.bot)
						break
					}
				case valmessage.CloseUnanimousAssertionRequest:
					if bot, ok := validator.bot.(WaitingObserver); ok {
						_ = bot.GetCore()
						newBot, msgs, err := bot.CloseUnanimous(request.ResultChan)

						if err != nil {
							request.ErrChan <- err
							break
						}

						validator.bot = newBot
						for _, msg := range msgs {
							sendChan <- msg
						}
					} else {
						request.ErrChan <- fmt.Errorf("can't close unanimous request, but was in the wrong state to handle it: %T", validator.bot)
					}
				case valmessage.DisputableDefenderRequest:
					core := validator.bot.GetCore()
					maxSteps := validator.bot.GetConfig().config.MaxExecutionStepCount
					startTime := validator.latestHeader.Number.Uint64()
					go func() {
						machine, defender := core.createDisputableDefender(
							startTime,
							request.Length,
							request.IncludePendingMessages,
							int32(maxSteps),
						)
						validator.requests <- valmessage.DisputableAssertionRequest{
							State:           machine,
							Defender:        defender,
							IncludedPending: request.IncludePendingMessages,
							ResultChan:      request.ResultChan,
						}
					}()
				case valmessage.DisputableAssertionRequest:
					validator.pendingDisputableRequest = &request
					validator.maybeAssert <- true
				case valmessage.VMStateRequest:
					core := validator.bot.GetCore()
					machineHash := core.machine.Hash()
					request.ResultChan <- valmessage.VMStateData{
						MachineState: machineHash,
						Config:       *validator.bot.GetConfig().config,
					}
				case valmessage.PendingMessageCheck:
					core := validator.bot.GetCore()
					request.ResultChan <- !core.GetInbox().PendingQueue.IsEmpty()
				case valmessage.CallRequest:
					core := validator.bot.GetCore()
					updatedState := core.machine.Clone()
					box := core.GetInbox().Clone()
					balance := core.balance.Clone()
					startTime := validator.latestHeader.Number.Uint64()
					msg := request.Message
					messageHash := solsha3.SoliditySHA3(
						solsha3.Bytes32(msg.Destination),
						solsha3.Bytes32(msg.Data.Hash()),
						solsha3.Uint256(msg.Currency),
						msg.TokenType[:],
					)
					msgHashInt := new(big.Int).SetBytes(messageHash[:])
					val, _ := value.NewTupleFromSlice([]value.Value{
						msg.Data,
						value.NewIntValue(validator.latestHeader.Time),
						value.NewIntValue(validator.latestHeader.Number),
						value.NewIntValue(msgHashInt),
					})
					callingMessage := protocol.Message{
						Data:        val.Clone(),
						TokenType:   msg.TokenType,
						Currency:    msg.Currency,
						Destination: msg.Destination,
					}
					go func() {
						box.InsertMessageGroup([]protocol.Message{callingMessage})
						actx := protocol.NewMachineAssertionContext(
							updatedState,
							balance,
							[2]uint64{startTime, startTime + 1},
							box.Receive(),
						)
						updatedState.RunUntilStop()
						ad := actx.Finalize(updatedState)
						results := ad.GetAssertion().Logs
						if len(results) > 0 {
							request.ResultChan <- results[len(results)-1]
						} else {
							request.ErrorChan <- errors.New("Call produced no output")
						}

					}()
				default:
					fmt.Printf("Unahandled validator request %T: %v\n", request, request)
				}
				validator.tryToAssert(sendChan)
			case <-validator.maybeAssert:
				validator.tryToAssert(sendChan)
			}

		}
	}()
}

func (validator *Validator) tryToAssert(sendChan chan<- valmessage.OutgoingMessage) {
	if validator.pendingDisputableRequest != nil && validator.canDisputableAssert() {
		validator.bot = AttemptingAssertDefender{
			validator.bot.GetCore(),
			validator.bot.GetConfig(),
			*validator.pendingDisputableRequest,
		}
		sendChan <- valmessage.SendAssertMessage{
			Precondition: validator.pendingDisputableRequest.Defender.GetPrecondition(),
			Assertion:    validator.pendingDisputableRequest.Defender.GetAssertion(),
		}
		validator.pendingDisputableRequest = nil
	}
}

func (validator *Validator) canDisputableAssert() bool {
	switch validator.bot.(type) {
	case WaitingObserver:
		return true
	default:
		return false
	}
}

func (validator *Validator) timeUpdate(sendChan chan<- valmessage.OutgoingMessage) {
	if validator.challengeBot != nil {
		newBot, msgs, err := validator.challengeBot.UpdateTime(validator.latestHeader.Number.Uint64())
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
			return
		}
		for _, msg := range msgs {
			sendChan <- msg
		}
		validator.challengeBot = newBot
	}
	newBot, msgs, err := validator.bot.UpdateTime(validator.latestHeader.Number.Uint64())
	if err != nil {
		fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
		return
	}
	for _, msg := range msgs {
		sendChan <- msg
	}
	validator.bot = newBot
}

func (validator *Validator) eventUpdate(ev valmessage.BridgeMessage, sendChan chan<- valmessage.OutgoingMessage) {
	if ev.Message.GetIncomingMessageType() == valmessage.ChallengeMessage {
		if validator.challengeBot == nil {
			panic("challengeBot can't be nil if challenge message is recieved")
		}

		newBot, msgs, err := validator.challengeBot.UpdateState(ev.Message, ev.GetHeader().Number.Uint64())
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
			return
		}
		for _, msg := range msgs {
			sendChan <- msg
		}
		validator.challengeBot = newBot
	} else {
		newBot, challengeBot, msgs, err := validator.bot.UpdateState(ev.Message, ev.GetHeader().Number.Uint64())
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, validator.bot)
			return
		}
		for _, msg := range msgs {
			sendChan <- msg
		}
		validator.bot = newBot
		if challengeBot != nil {
			validator.challengeBot = challengeBot
		}
	}
}
