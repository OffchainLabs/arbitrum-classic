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
	"math"
	"math/big"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/state"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ChannelValidator struct {
	*Validator
	channelBot *ChannelBot
}

func NewChannelValidator(
	name string,
	b bridge.Bridge,
	address common.Address,
	latestHeader *types.Header,
	config *valmessage.VMConfiguration,
	challengeEverything bool,
	machine machine.Machine,
	maxCallSteps int32,
) *ChannelValidator {
	c := core.NewCore(
		machine,
	)

	valConfig := core.NewValidatorConfig(address, config, challengeEverything, maxCallSteps)
	channelBot := &ChannelBot{state.NewWaiting(valConfig, c), b}
	val := NewValidator(
		name,
		channelBot,
		latestHeader,
	)
	return &ChannelValidator{
		val,
		channelBot,
	}
}

func (validator *ChannelValidator) HasOpenAssertion() chan bool {
	resultChan := make(chan bool, 1)
	validator.actions <- func() {
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
	Assertion *protocol.ExecutionAssertion

	ShouldFinalize func(*protocol.ExecutionAssertion) bool

	ResultChan chan<- valmessage.UnanimousUpdateResults
	ErrChan    chan<- error
}

func (validator *ChannelValidator) InitiateUnanimousRequest(
	length uint64,
	messages []protocol.Message,
	messageHashes [][]byte,
	final bool,
	maxSteps int32,
	shouldFinalize func(*protocol.ExecutionAssertion) bool,
) (
	<-chan valmessage.UnanimousRequest,
	<-chan valmessage.UnanimousUpdateResults,
	<-chan error,
) {
	unanRequestChan := make(chan valmessage.UnanimousRequest, 1)
	updateResultChan := make(chan valmessage.UnanimousUpdateResults, 1)
	errChan := make(chan error, 1)

	validator.actions <- func() {
		if !validator.canRun() {
			errChan <- errors.New("can't unanimous assert when not running")
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
		timeBounds := protocol.NewTimeBoundsBlocks(validator.latestHeader.Number.Uint64(), validator.latestHeader.Number.Uint64()+length)
		seqNum := bot.OffchainContext(final)
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
	shouldFinalize func(*protocol.ExecutionAssertion) bool,
) (<-chan valmessage.UnanimousUpdateResults, <-chan error) {
	resultChan := make(chan valmessage.UnanimousUpdateResults, 1)
	errChan := make(chan error, 1)
	validator.actions <- func() {
		if !validator.canRun() {
			errChan <- errors.New("can't unanimous assert when not running")
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

		_ = bot.OffchainContext(request.SequenceNum == math.MaxUint64)
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
	validator.actions <- func() {
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
	validator.actions <- func() {
		bot, ok := validator.channelBot.ChannelState.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("recieved unanimous confirm request, but was in the wrong state to handle it: %T", validator.channelBot.ChannelState)
			return
		}
		if err := bot.ValidateUnanimousAssertion(request); err != nil {
			errChan <- err
			return
		}

		proposalResults := bot.ProposalResults()
		newBot, err := bot.FinalizePendingUnanimous(signatures)
		if err != nil {
			errChan <- err
			return
		}

		validator.channelBot.bridge.FinalizedAssertion(
			nil,
			[]byte{},
			signatures,
			&proposalResults,
		)

		if request.SequenceNum == math.MaxUint64 {
			if canClose {
				_, err := newBot.CloseUnanimous(validator.channelBot.bridge)
				if err != nil {
					errChan <- errors2.Wrap(err, "Error closing unanimous assertion")
					return
				}
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
	validator.actions <- func() {
		bot, ok := validator.channelBot.ChannelState.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("can't close unanimous request, but was in the wrong state to handle it: %T", validator.channelBot.ChannelState)
			return
		}
		_, err := bot.CloseUnanimous(validator.channelBot.bridge)
		if err != nil {
			errChan <- err
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

func (validator *ChannelValidator) ClosingUnanimousAssertionRequest() (<-chan bool, <-chan error) {
	resultChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	validator.actions <- func() {
		bot, ok := validator.channelBot.ChannelState.(state.Waiting)
		if !ok {
			errChan <- fmt.Errorf("can't close unanimous request. Validator was in the wrong state to handle it: %T", validator.channelBot.ChannelState)
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
