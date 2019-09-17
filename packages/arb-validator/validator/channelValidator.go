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
	"context"
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/state"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ChannelValidator struct {
	*Validator
	channelBot     *ChannelBot
	channelActions chan func(bridge.Bridge)
}

func NewChannelValidator(
	name string,
	address common.Address,
	latestHeader *types.Header,
	balance *protocol.BalanceTracker,
	config *valmessage.VMConfiguration,
	machine machine.Machine,
	challengeEverything bool,
	maxCallSteps int32,
) *ChannelValidator {
	actions := make(chan func(bridge.Bridge), 100)
	c := core.NewCore(
		machine,
		balance,
	)

	valConfig := core.NewValidatorConfig(address, config, challengeEverything, maxCallSteps)
	channelBot := &ChannelBot{state.NewWaiting(valConfig, c)}
	val := NewValidator(
		name,
		channelBot,
		latestHeader,
	)
	return &ChannelValidator{
		val,
		channelBot,
		actions,
	}
}

func (validator *ChannelValidator) HasOpenAssertion() chan bool {
	resultChan := make(chan bool, 1)
	validator.channelActions <- func(bridge bridge.Bridge) {
		bot, ok := validator.channelBot.ChannelState.(state.Waiting)
		if !ok {
			resultChan <- false
		} else {
			resultChan <- bot.HasOpenAssertion()
		}
	}

	return resultChan
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

func (validator *ChannelValidator) InitiateUnanimousRequest(
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

	validator.channelActions <- func(bridge bridge.Bridge) {
		if !validator.canRun() {
			errChan <- errors.New("Can't unanimous assert when not running")
			return
		}
		bot, ok := validator.channelBot.ChannelState.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("recieved initiate unanimous request, but was in the wrong state to handle it: %T", validator.channelBot.ChannelState)
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

func (validator *ChannelValidator) RequestFollowUnanimous(
	request valmessage.UnanimousRequestData,
	messages []protocol.Message,
	maxSteps int32,
	shouldFinalize func(*protocol.Assertion) bool,
) (<-chan valmessage.UnanimousUpdateResults, <-chan error) {
	resultChan := make(chan valmessage.UnanimousUpdateResults, 1)
	errChan := make(chan error, 1)
	validator.channelActions <- func(bridge bridge.Bridge) {
		if !validator.canRun() {
			errChan <- errors.New("Can't unanimous assert when not running")
			return
		}
		bot, ok := validator.channelBot.ChannelState.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("recieved follow unanimous request, but was in the wrong state to handle it: %T", validator.channelBot.ChannelState)
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

func (validator *ChannelValidator) requestUnanimousUpdate(request unanimousUpdateRequest) {
	validator.channelActions <- func(bridge bridge.Bridge) {
		bot, ok := validator.channelBot.ChannelState.(state.Waiting)
		if !ok {
			request.ErrChan <- fmt.Errorf("recieved unanimous update request, but was in the wrong state to handle it: %T", validator.channelBot.ChannelState)
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
		validator.channelBot.updateBot(newBot)
	}
}

func (validator *ChannelValidator) ConfirmOffchainUnanimousAssertion(
	request valmessage.UnanimousRequestData,
	signatures [][]byte,
	canClose bool,
) (<-chan bool, <-chan error) {
	resultChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	validator.channelActions <- func(bridge bridge.Bridge) {
		bot, ok := validator.channelBot.ChannelState.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("recieved unanimous confirm request, but was in the wrong state to handle it: %T", validator.channelBot.ChannelState)
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
			validator.channelBot.updateBot(newBot2)
		} else {
			validator.channelBot.updateBot(newBot)
			resultChan <- true
		}
	}
	return resultChan, errChan
}

func (validator *ChannelValidator) CloseUnanimousAssertionRequest() (<-chan bool, <-chan error) {
	resultChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	validator.channelActions <- func(bridge bridge.Bridge) {
		bot, ok := validator.channelBot.ChannelState.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("can't close unanimous request, but was in the wrong state to handle it: %T", validator.channelBot.ChannelState)
			return
		}
		bot.CloseUnanimous(bridge)
		newBot, err := bot.ClosingUnanimous(resultChan, errChan)
		if err != nil {
			errChan <- err
			return
		}
		validator.channelBot.updateBot(newBot)
	}
	return resultChan, errChan
}

func (validator *ChannelValidator) ClosingUnanimousAssertionRequest() (<-chan bool, <-chan error) {
	resultChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	validator.channelActions <- func(bridge bridge.Bridge) {
		bot, ok := validator.channelBot.ChannelState.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("can't close unanimous request. ChannelValidator was in the wrong state to handle it: %T", validator.channelBot.ChannelState)
			return
		}
		newBot, err := bot.ClosingUnanimous(resultChan, errChan)
		if err != nil {
			errChan <- err
			return
		}
		validator.channelBot.updateBot(newBot)
	}
	return resultChan, errChan
}

func (validator *ChannelValidator) Run(recvChan <-chan ethconnection.Notification, bridge bridge.Bridge, ctx context.Context) {
	defer fmt.Printf("%v: Exiting\n", validator.Name)
	for {
		select {
		case <-ctx.Done():
			break
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
					if !validator.channelBot.GetCore().ValidateAssertion(pre, newHeader.Number.Uint64()) {
						validator.pendingDisputableRequest.ErrorChan <- errors.New("Precondition was invalidated")
						close(validator.pendingDisputableRequest.ErrorChan)
						close(validator.pendingDisputableRequest.ResultChan)
						validator.pendingDisputableRequest = nil
					}
				}
			}

			switch ev := notification.Event.(type) {
			case ethconnection.NewTimeEvent:
				break
			case ethconnection.VMEvent:
				validator.eventUpdate(ev, notification.Header, bridge)
			case ethconnection.MessageDeliveredEvent:
				validator.channelBot.SendMessageToVM(ev.Msg)
			default:
				panic("Should never recieve other kinds of events")
			}
		case action := <-validator.actions:
			action(bridge)
		case action := <-validator.channelActions:
			action(bridge)
		case <-validator.maybeAssert:
		}

		if bot, ok := validator.channelBot.ChannelState.(state.Waiting); ok && validator.pendingDisputableRequest != nil {
			validator.channelBot.updateBot(bot.AttemptAssertion(context.Background(), *validator.pendingDisputableRequest, bridge))
			validator.pendingDisputableRequest = nil
		}
	}
}

func (validator *ChannelValidator) timeUpdate(bridge bridge.Bridge) {
	if validator.challengeBot != nil {
		newBot, err := validator.challengeBot.UpdateTime(validator.latestHeader.Number.Uint64(), bridge)
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
			return
		}
		validator.challengeBot = newBot
	}
	newBot, err := validator.channelBot.ChannelUpdateTime(validator.latestHeader.Number.Uint64(), bridge)
	if err != nil {
		fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
		return
	}
	validator.channelBot.updateBot(newBot)
}

func (validator *ChannelValidator) eventUpdate(ev ethconnection.VMEvent, header *types.Header, bridge bridge.Bridge) {
	if ev.GetIncomingMessageType() == ethconnection.ChallengeMessage {
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
		newBot, challengeBot, err := validator.channelBot.ChannelUpdateState(ev, header.Number.Uint64(), bridge)
		if err != nil {
			fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, validator.channelBot.ChannelState)
			return
		}
		validator.channelBot.updateBot(newBot)
		if challengeBot != nil {
			validator.challengeBot = challengeBot
		}
	}
}
