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
	"github.com/offchainlabs/arb-validator/bridge"
	"github.com/offchainlabs/arb-validator/ethbridge"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-validator/valmessage"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/vm"
)

type validatorState interface {
	UpdateTime(uint64, bridge.Bridge) (validatorState, error)
	UpdateState(ethbridge.Event, uint64, bridge.Bridge) (validatorState, challengeState, error)

	SendMessageToVM(msg protocol.Message)
	GetCore() *validatorCore
	GetConfig() *validatorConfig
}

type challengeState interface {
	UpdateTime(uint64, bridge.Bridge) (challengeState, error)
	UpdateState(ethbridge.Event, uint64, bridge.Bridge) (challengeState, error)
}

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

type Validator struct {
	Name        string
	requests    chan interface{}
	maybeAssert chan bool

	// Run loop only
	bot                      validatorState
	challengeBot             challengeState
	latestHeader             *types.Header
	pendingDisputableRequest *disputableAssertionRequest
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
		newWaitingObserver(valConfig, core),
		nil,
		nil,
		nil,
	}
}

func (validator *Validator) RequestCall(msg protocol.Message) (<-chan value.Value, <-chan error) {
	resultChan := make(chan value.Value, 1)
	errorChan := make(chan error, 1)
	validator.requests <- callRequest{
		Message:    msg,
		ResultChan: resultChan,
		ErrorChan:  errorChan,
	}
	return resultChan, errorChan
}

func (validator *Validator) HasPendingMessages() chan bool {
	retChan := make(chan bool, 1)
	validator.requests <- pendingMessageCheck{ResultChan: retChan}
	return retChan
}

func (validator *Validator) RequestVMState() <-chan valmessage.VMStateData {
	resultChan := make(chan valmessage.VMStateData)
	validator.requests <- vmStateRequest{ResultChan: resultChan}
	return resultChan
}

func (validator *Validator) RequestDisputableAssertion(length uint64, includePendingMessages bool) <-chan bool {
	resultChan := make(chan bool)
	validator.requests <- disputableDefenderRequest{
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
	validator.requests <- initiateUnanimousRequest{
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
	validator.requests <- followUnanimousRequest{
		UnanimousRequestData: request,
		NewMessages:          messages,
		ResultChan:           resultChan,
		ErrChan:              errChan,
	}
	return resultChan, errChan
}

func (validator *Validator) ConfirmOffchainUnanimousAssertion(
	unanRequest valmessage.UnanimousRequestData,
	signatures [][]byte,
) (<-chan bool, <-chan error) {
	resultChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	validator.requests <- unanimousConfirmRequest{
		UnanimousRequestData: unanRequest,
		Signatures:           signatures,
		ResultChan:           resultChan,
		ErrChan:              errChan,
	}
	return resultChan, errChan
}

func (validator *Validator) CloseUnanimousAssertionRequest() <-chan bool {
	resultChan := make(chan bool, 1)
	validator.requests <- closeUnanimousAssertionRequest{
		ResultChan: resultChan,
	}
	return resultChan
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
						pre := validator.pendingDisputableRequest.GetPrecondition()
						if !validator.bot.GetCore().ValidateAssertion(pre, newHeader.Number.Uint64()) {
							validator.pendingDisputableRequest.NotifyInvalid()
							validator.pendingDisputableRequest = nil
						}
					}
				}

				switch ev := notification.Event.(type) {
				case ethbridge.NewTimeEvent:
					break
				case ethbridge.VMEvent:
					validator.eventUpdate(ev, notification.Header, bridge)
				case ethbridge.MessageDeliveredEvent:
					validator.bot.SendMessageToVM(ev.Msg)

					// Invalidate assertions that included pending messages
					if validator.pendingDisputableRequest != nil && validator.pendingDisputableRequest.IncludedPendingInbox() {
						validator.pendingDisputableRequest.NotifyInvalid()
						validator.pendingDisputableRequest = nil
					}
				default:
					panic("Should never recieve other kinds of events")
				}
				validator.tryToAssert(bridge)
			case request := <-validator.requests:
				switch request := request.(type) {
				case initiateUnanimousRequest:
					if bot, ok := validator.bot.(waitingObserver); ok {
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
							validator.requests <- unanimousUpdateRequest{
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
				case followUnanimousRequest:
					if bot, ok := validator.bot.(waitingObserver); ok {
						if err := bot.validateUnanimousRequest(request.UnanimousRequestData); err != nil {
							request.ErrChan <- err
							break
						}

						mq, _, _ := bot.OffchainContext(request.NewMessages, request.TimeBounds, request.SequenceNum == math.MaxUint64)
						clonedCore := bot.GetCore().Clone()
						go func() {
							newCore, assertion := clonedCore.OffchainAssert(mq, request.TimeBounds)
							validator.requests <- unanimousUpdateRequest{
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
				case unanimousUpdateRequest:
					if bot, ok := validator.bot.(waitingObserver); ok {
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
				case unanimousConfirmRequest:
					if bot, ok := validator.bot.(waitingObserver); ok {
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
						bridge.FinalizedAssertion(
							proposal.assertion,
							proposal.newLogCount,
						)
						request.ResultChan <- true
					} else {
						request.ErrChan <- fmt.Errorf("recieved unanimous confirm request, but was in the wrong state to handle it: %T", validator.bot)
						break
					}
				case closeUnanimousAssertionRequest:
					if bot, ok := validator.bot.(waitingObserver); ok {
						_ = bot.GetCore()
						newBot, err := bot.CloseUnanimous(bridge, request.ResultChan)
						if err != nil {
							request.ErrChan <- err
							break
						}

						validator.bot = newBot
					} else {
						request.ErrChan <- fmt.Errorf("can't close unanimous request, but was in the wrong state to handle it: %T", validator.bot)
					}
				case disputableDefenderRequest:
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
						validator.requests <- disputableAssertionRequest{
							State:           machine,
							Defender:        defender,
							IncludedPending: request.IncludePendingMessages,
							ResultChan:      request.ResultChan,
						}
					}()
				case disputableAssertionRequest:
					validator.pendingDisputableRequest = &request
					validator.maybeAssert <- true
				case vmStateRequest:
					core := validator.bot.GetCore()
					machineHash := core.machine.Hash()
					request.ResultChan <- valmessage.VMStateData{
						MachineState: machineHash,
						Config:       *validator.bot.GetConfig().config,
					}
				case pendingMessageCheck:
					core := validator.bot.GetCore()
					request.ResultChan <- !core.GetInbox().PendingQueue.IsEmpty()
				case callRequest:
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
				validator.tryToAssert(bridge)
			case <-validator.maybeAssert:
				validator.tryToAssert(bridge)
			}
		}
	}()
}

func (validator *Validator) tryToAssert(bridge bridge.Bridge) {
	if validator.pendingDisputableRequest != nil && validator.canDisputableAssert() {
		validator.bot = attemptingAssertDefender{
			validator.bot.GetCore(),
			validator.bot.GetConfig(),
			*validator.pendingDisputableRequest,
		}
		bridge.DisputableAssert(
			validator.pendingDisputableRequest.Defender.GetPrecondition(),
			validator.pendingDisputableRequest.Defender.GetAssertion(),
		)
		validator.pendingDisputableRequest = nil
	}
}

func (validator *Validator) canDisputableAssert() bool {
	switch validator.bot.(type) {
	case waitingObserver:
		return true
	default:
		return false
	}
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
