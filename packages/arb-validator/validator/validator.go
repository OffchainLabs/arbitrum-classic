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
	"log"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"

	"github.com/ethereum/go-ethereum/core/types"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/disputable"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/state"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ChainBot struct {
	state.ChainState
	bridge bridge.ArbVMBridge
}

func (bot *ChainBot) updateBot(state state.ChainState) {
	bot.ChainState = state
}

func (bot *ChainBot) getBridge() bridge.ArbVMBridge {
	return bot.bridge
}

func (bot *ChainBot) attemptDisputableAssertion(ctx context.Context, request *disputable.AssertionRequest) (bool, error) {
	if waitingBot, ok := bot.ChainState.(state.Waiting); ok && request != nil {
		newBot, err := waitingBot.AttemptAssertion(ctx, *request, bot.bridge)
		if err != nil {
			return false, err
		}
		bot.ChainState = newBot
		return true, nil
	}
	return false, nil
}

func (bot *ChainBot) updateTime(time uint64) error {
	newBot, err := bot.ChainState.ChainUpdateTime(time, bot.bridge)
	if err != nil {
		return err
	}
	bot.ChainState = newBot
	return nil
}

func (bot *ChainBot) updateState(ev arbbridge.Event, time uint64) error {
	newBot, err := bot.ChainState.ChainUpdateState(ev, time, bot.bridge)
	if err != nil {
		return err
	}
	bot.ChainState = newBot
	return nil
}

type ChannelBot struct {
	state.ChannelState
	bridge bridge.Bridge
}

func (bot *ChannelBot) updateBot(state state.ChannelState) {
	bot.ChannelState = state
}

func (bot *ChannelBot) getBridge() bridge.ArbVMBridge {
	return bot.bridge
}

func (bot *ChannelBot) attemptDisputableAssertion(ctx context.Context, request *disputable.AssertionRequest) (bool, error) {
	if waitingBot, ok := bot.ChannelState.(state.Waiting); ok && request != nil {
		newBot, err := waitingBot.AttemptAssertion(ctx, *request, bot.bridge)
		if err != nil {
			return false, err
		}
		bot.ChannelState = newBot
		return true, nil
	}
	return false, nil
}

func (bot *ChannelBot) updateTime(time uint64) error {
	newBot, err := bot.ChannelState.ChannelUpdateTime(time, bot.bridge)
	if err != nil {
		return err
	}
	bot.ChannelState = newBot
	return nil
}

func (bot *ChannelBot) updateState(ev arbbridge.Event, time uint64) error {
	newBot, err := bot.ChannelState.ChannelUpdateState(ev, time, bot.bridge)
	if err != nil {
		return err
	}
	bot.ChannelState = newBot
	return nil
}

type Bot interface {
	state.State
	updateTime(uint64) error
	updateState(arbbridge.Event, uint64) error
	getBridge() bridge.ArbVMBridge
	attemptDisputableAssertion(ctx context.Context, request *disputable.AssertionRequest) (bool, error)
}

type Validator struct {
	Name        string
	actions     chan func()
	maybeAssert chan bool

	// Run loop only
	bot                      Bot
	latestHeader             *types.Header
	pendingDisputableRequest *disputable.AssertionRequest
}

func NewValidator(
	name string,
	bot Bot,
	latestHeader *types.Header,
) *Validator {
	actions := make(chan func(), 100)
	maybeAssert := make(chan bool, 100)
	return &Validator{
		name,
		actions,
		maybeAssert,
		bot,
		latestHeader,
		nil,
	}
}

func (validator *Validator) RequestCall(msg protocol.Message) (<-chan value.Value, <-chan error) {
	resultChan := make(chan value.Value, 1)
	errChan := make(chan error, 1)
	validator.actions <- func() {
		if !validator.canRun() {
			errChan <- errors.New("cannot call when ArbChannel is not running")
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
				protocol.NewTimeBoundsBlocks(startTime, startTime+1),
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
	validator.actions <- func() {
		c := validator.bot.GetCore()
		resultChan <- c.GetMachine().PendingMessageCount()
	}
	return resultChan
}

func (validator *Validator) canRun() bool {
	return validator.bot.GetCore().GetMachine().CurrentStatus() == machine.Extensive
}

func (validator *Validator) CanRun() chan bool {
	resultChan := make(chan bool, 1)
	validator.actions <- func() {
		resultChan <- validator.canRun()
	}
	return resultChan
}

func (validator *Validator) CanContinueRunning() chan bool {
	resultChan := make(chan bool, 1)
	validator.actions <- func() {
		if !validator.canRun() {
			resultChan <- false
		} else {
			currentTime := validator.latestHeader.Number.Uint64()
			mach := validator.bot.GetCore().GetMachine()
			resultChan <- !machine.IsMachineBlocked(mach, currentTime)
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
	validator.actions <- func() {
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
	validator.actions <- func() {
		if !validator.canRun() {
			errChan <- errors.New("can't disputable assert when not running")
			return
		}
		c := validator.bot.GetCore()
		mClone := c.GetMachine().Clone()
		maxSteps := validator.bot.GetConfig().VMConfig.MaxExecutionStepCount
		startTime := validator.latestHeader.Number.Uint64()
		go func() {
			endTime := startTime + length
			tb := protocol.NewTimeBoundsBlocks(startTime, endTime)
			beforeHash := mClone.Hash()
			assertion := mClone.ExecuteAssertion(int32(maxSteps), tb)
			pre := valprotocol.NewPrecondition(beforeHash, tb, mClone.InboxHash())
			request := &disputable.AssertionRequest{
				AfterCore:    core.NewCore(mClone),
				Precondition: pre,
				Assertion:    assertion,
				ResultChan:   resultChan,
				ErrorChan:    errChan,
			}
			validator.actions <- func() {
				validator.pendingDisputableRequest = request
				validator.maybeAssert <- true
			}
		}()
	}
	return resultChan, errChan
}

func (validator *Validator) validatorClosing() {
	fmt.Printf("%v: Exiting\n", validator.Name)
	validator.bot.getBridge().SendMonitorErr(bridge.Error{errors.New("WARNING: validator closing"), "WARNING: validator closing", false})
}

func (validator *Validator) Run(ctx context.Context, recvChan <-chan arbbridge.Notification) {
	defer validator.validatorClosing()
	for {
		select {
		case <-ctx.Done():
			break
		case notification, ok := <-recvChan:
			//log.Printf("validator %v got notification %T: %v\n", validator.Name, notification, notification)
			if !ok {
				fmt.Printf("%v: Error in recvChan\n", validator.Name)
				return
			}

			newHeader := notification.Header
			if validator.latestHeader == nil || newHeader.Number.Uint64() >= validator.latestHeader.Number.Uint64() && newHeader.Hash() != validator.latestHeader.Hash() {
				validator.latestHeader = newHeader
				err := validator.bot.updateTime(validator.latestHeader.Number.Uint64())
				if err != nil {
					//log.Printf("Validator %v: Error processing time update - %v\n", validator.Name, err)
					if errstat, ok := err.(*bridge.Error); ok {
						if !errstat.Recoverable {
							//log.Printf("Validator %v: non recoverable error\n", validator.Name)
							validator.bot.getBridge().SendMonitorErr(*errstat)
							return
						} else {
							//log.Printf("Validator %v: recoverable error - contiuing\n", validator.Name)
							validator.bot.getBridge().SendMonitorErr(*errstat)
						}
					} else {
						validator.bot.getBridge().SendMonitorErr(bridge.Error{err, "non recoverable error - exiting", false})
						return
					}
				}
				if validator.pendingDisputableRequest != nil {
					pre := validator.pendingDisputableRequest.Precondition
					if !validator.bot.GetCore().ValidateAssertion(pre, newHeader.Number.Uint64()) {
						validator.pendingDisputableRequest.ErrorChan <- errors.New("precondition was invalidated")
						close(validator.pendingDisputableRequest.ErrorChan)
						close(validator.pendingDisputableRequest.ResultChan)
						validator.pendingDisputableRequest = nil
					}
				}
			}

			switch ev := notification.Event.(type) {
			case arbbridge.NewTimeEvent:
				break
			case arbbridge.VMEvent:
				err := validator.bot.updateState(ev, notification.Header.Number.Uint64())
				if err != nil {
					//log.Printf("*****Validator %v: error - %v\n", validator.Name, err)
					if errstat, ok := err.(*bridge.Error); ok {
						if !errstat.Recoverable {
							//log.Printf("Validator %v: non recoverable error - %v\n", validator.Name, err)
							validator.bot.getBridge().SendMonitorErr(*errstat)
							return
						} else {
							//log.Printf("Validator %v: recoverable error - %v - contiuing\n", validator.Name, err)
							validator.bot.getBridge().SendMonitorErr(*errstat)
						}
					} else {
						log.Println("Error processing event update", err)
						return
					}
				}
			case arbbridge.MessageDeliveredEvent:
				validator.bot.SendMessageToVM(ev.Msg)
			default:
				panic("Should never receive other kinds of events")
			}
		case action := <-validator.actions:
			action()
		case <-validator.maybeAssert:
		}

		if asserted, err := validator.bot.attemptDisputableAssertion(ctx, validator.pendingDisputableRequest); asserted || err != nil {
			if err != nil {
				log.Printf("Validator %v: Failed to disputable assert - %v\n", validator.Name, err)
				validator.bot.getBridge().SendMonitorErr(bridge.Error{err, "ERROR: failed to create disputable assertion", true})
			}
			validator.pendingDisputableRequest = nil
		}
	}
}
