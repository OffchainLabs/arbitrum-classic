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
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"
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
	requests    chan interface{}
	maybeAssert chan bool

	// Run loop only
	bot                      state.State
	challengeBot             challenge.State
	latestHeader             *types.Header
	pendingDisputableRequest *state.DisputableAssertionRequest
}

func NewValidator(name string, address common.Address, balance *protocol.BalanceTracker, config *valmessage.VMConfiguration, machine machine.Machine, challengeEverything bool, maxCallSteps int32) *Validator {
	requests := make(chan interface{}, 10)
	maybeAssert := make(chan bool, 100)
	c := core.NewCore(
		machine,
		balance,
	)

	// TODO: latestHeader starts as nil which isn't valid. This needs to be properly initialized
	valConfig := core.NewValidatorConfig(address, config, challengeEverything, maxCallSteps)
	return &Validator{
		name,
		requests,
		maybeAssert,
		state.NewWaiting(valConfig, c),
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

func (validator *Validator) HasOpenAssertion() chan bool {
	retChan := make(chan bool, 1)
	validator.requests <- openAssertionCheck{ResultChan: retChan}
	return retChan
}

func (validator *Validator) RequestVMState() <-chan valmessage.VMStateData {
	resultChan := make(chan valmessage.VMStateData)
	validator.requests <- vmStateRequest{ResultChan: resultChan}
	return resultChan
}

func (validator *Validator) RequestDisputableAssertion(length uint64) <-chan bool {
	resultChan := make(chan bool)
	validator.requests <- disputableDefenderRequest{
		Length:     length,
		ResultChan: resultChan,
	}
	return resultChan
}

func (validator *Validator) InitiateUnanimousRequest(
	length uint64,
	messages []protocol.Message,
	messageHashes [][]byte,
	final bool,
	maxSteps int32,
) (
	<-chan valmessage.UnanimousRequest,
	<-chan valmessage.UnanimousUpdateResults,
	<-chan error,
) {
	unanRequestChan := make(chan valmessage.UnanimousRequest, 1)
	updateResultChan := make(chan valmessage.UnanimousUpdateResults, 1)
	errChan := make(chan error, 1)
	validator.requests <- initiateUnanimousRequest{
		TimeLength:    length,
		NewMessages:   messages,
		MessageHashes: messageHashes,
		Final:         final,
		MaxSteps:      maxSteps,
		RequestChan:   unanRequestChan,
		ResultChan:    updateResultChan,
		ErrChan:       errChan,
	}
	return unanRequestChan, updateResultChan, errChan
}

func (validator *Validator) RequestFollowUnanimous(
	request valmessage.UnanimousRequestData,
	messages []protocol.Message,
	maxSteps int32,
) (<-chan valmessage.UnanimousUpdateResults, <-chan error) {
	resultChan := make(chan valmessage.UnanimousUpdateResults, 1)
	errChan := make(chan error, 1)
	validator.requests <- followUnanimousRequest{
		UnanimousRequestData: request,
		NewMessages:          messages,
		MaxSteps:             maxSteps,
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

func (validator *Validator) CloseUnanimousAssertionRequest() (<-chan bool, <-chan error) {
	resultChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	validator.requests <- closeUnanimousAssertionRequest{
		ResultChan: resultChan,
		ErrChan:    errChan,
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
				default:
					panic("Should never recieve other kinds of events")
				}
				validator.tryToAssert(bridge)
			case request := <-validator.requests:
				switch request := request.(type) {
				case initiateUnanimousRequest:
					if bot, ok := validator.bot.(state.Waiting); ok {
						newMessages := make([]protocol.Message, 0, len(request.NewMessages))
						messageRecords := make([]protocol.Message, 0, len(request.NewMessages))
						for i, msg := range request.NewMessages {
							msgHashInt := new(big.Int).SetBytes(request.MessageHashes[i])
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
						timeBounds := [2]uint64{validator.latestHeader.Number.Uint64(), validator.latestHeader.Number.Uint64() + request.TimeLength}
						tb, seqNum := bot.OffchainContext(timeBounds, request.Final)
						clonedMachine := bot.GetCore().GetMachine().Clone()
						requestData := valmessage.UnanimousRequestData{
							BeforeHash:  clonedMachine.Hash(),
							BeforeInbox: clonedMachine.InboxHash().Hash(),
							SequenceNum: seqNum,
							TimeBounds:  tb,
						}

						request.RequestChan <- valmessage.UnanimousRequest{UnanimousRequestData: requestData, NewMessages: messageRecords}
						go func() {
							clonedMachine.SendOffchainMessages(newMessages)
							assertion := clonedMachine.ExecuteAssertion(
								request.MaxSteps,
								timeBounds,
							)
							validator.requests <- state.UnanimousUpdateRequest{
								UnanimousRequestData: requestData,
								NewMessages:          newMessages,
								Machine:              clonedMachine,
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
					if bot, ok := validator.bot.(state.Waiting); ok {
						if err := bot.ValidateUnanimousRequest(request.UnanimousRequestData); err != nil {
							request.ErrChan <- err
							break
						}

						_, _ = bot.OffchainContext(request.TimeBounds, request.SequenceNum == math.MaxUint64)
						clonedMachine := bot.GetCore().GetMachine().Clone()
						go func() {
							clonedMachine.SendOffchainMessages(request.NewMessages)
							assertion := clonedMachine.ExecuteAssertion(
								request.MaxSteps,
								request.TimeBounds,
							)
							validator.requests <- state.UnanimousUpdateRequest{
								UnanimousRequestData: request.UnanimousRequestData,
								NewMessages:          request.NewMessages,
								Machine:              clonedMachine,
								Assertion:            assertion,
								ResultChan:           request.ResultChan,
								ErrChan:              request.ErrChan,
							}
						}()
					} else {
						request.ErrChan <- fmt.Errorf("recieved follow unanimous request, but was in the wrong state to handle it: %T", validator.bot)
						break
					}
				case state.UnanimousUpdateRequest:
					if bot, ok := validator.bot.(state.Waiting); ok {
						if err := bot.ValidateUnanimousRequest(request.UnanimousRequestData); err != nil {
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
					if bot, ok := validator.bot.(state.Waiting); ok {
						if err := bot.ValidateUnanimousRequest(request.UnanimousRequestData); err != nil {
							request.ErrChan <- err
							break
						}

						newBot, proposal, err := bot.FinalizePendingUnanimous(request.SequenceNum, request.Signatures)
						if err != nil {
							request.ErrChan <- err
							break
						}
						validator.bot = newBot
						proposalResults := bot.ProposalResults()
						bridge.FinalizedAssertion(
							proposal.Assertion,
							proposal.NewLogCount,
							request.Signatures,
							&proposalResults,
							[]byte{},
						)
						request.ResultChan <- true
					} else {
						request.ErrChan <- fmt.Errorf("recieved unanimous confirm request, but was in the wrong state to handle it: %T", validator.bot)
						break
					}
				case closeUnanimousAssertionRequest:
					if bot, ok := validator.bot.(state.Waiting); ok {
						newBot, err := bot.CloseUnanimous(bridge, request.ResultChan, request.ErrChan)
						if err != nil {
							request.ErrChan <- err
							break
						}
						validator.bot = newBot
					} else {
						request.ErrChan <- fmt.Errorf("can't close unanimous request, but was in the wrong state to handle it: %T", validator.bot)
					}
				case disputableDefenderRequest:
					c := validator.bot.GetCore()
					mClone := c.GetMachine().Clone()
					maxSteps := validator.bot.GetConfig().VMConfig.MaxExecutionStepCount
					startTime := validator.latestHeader.Number.Uint64()
					balance := c.GetBalance().Clone()
					go func() {
						endTime := startTime + request.Length
						tb := [2]uint64{startTime, endTime}
						pre := &protocol.Precondition{
							BeforeHash:    mClone.Hash(),
							TimeBounds:    tb,
							BeforeBalance: balance,
							BeforeInbox:   mClone.InboxHash(),
						}
						assertion := mClone.ExecuteAssertion(int32(maxSteps), tb)
						validator.requests <- state.DisputableAssertionRequest{
							AfterState:   mClone,
							Precondition: pre,
							Assertion:    assertion,
							ResultChan:   request.ResultChan,
						}
					}()
				case state.DisputableAssertionRequest:
					validator.pendingDisputableRequest = &request
					validator.maybeAssert <- true
				case vmStateRequest:
					c := validator.bot.GetCore()
					machineHash := c.GetMachine().Hash()
					request.ResultChan <- valmessage.VMStateData{
						MachineState: machineHash,
						Config:       *validator.bot.GetConfig().VMConfig,
					}
				case pendingMessageCheck:
					c := validator.bot.GetCore()
					request.ResultChan <- c.GetMachine().HasPendingMessages()
				case openAssertionCheck:
					bot, ok := validator.bot.(state.Waiting)
					if !ok {
						request.ResultChan <- false
					} else {
						request.ResultChan <- bot.HasOpenAssertion()
					}
				case callRequest:
					c := validator.bot.GetCore()
					updatedState := c.GetMachine().Clone()
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
							request.ErrorChan <- errors.New("Call produced no output")
							return
						}
						lastLogVal := results[len(results)-1]
						lastLog, err := evm.ProcessLog(lastLogVal)
						if err != nil {
							request.ErrorChan <- err
							return
						}
						logHash := lastLog.GetEthMsg().Data.TxHash
						if !bytes.Equal(logHash[:], messageHash) {
							// Last produced log is not the call we sent
							request.ErrorChan <- errors.New("Call took too long to execute")
							return
						}

						request.ResultChan <- results[len(results)-1]
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
	if bot, ok := validator.bot.(state.Waiting); ok && validator.pendingDisputableRequest != nil {
		validator.bot = bot.AttemptAssertion(*validator.pendingDisputableRequest, bridge)
		validator.pendingDisputableRequest = nil
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
