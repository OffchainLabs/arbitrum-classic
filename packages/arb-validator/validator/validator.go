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
	"errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/state"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/disputable"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ChainBot struct {
	state.ChainState
}

func (bot *ChainBot) updateBot(state state.ChainState) {
	bot.ChainState = state
}

type ChannelBot struct {
	state.ChannelState
}

func (bot *ChannelBot) updateBot(state state.ChannelState) {
	bot.ChannelState = state
}

type Bot interface {
	state.State
}

type Validator struct {
	Name        string
	actions     chan func(bridge.ArbVMBridge)
	maybeAssert chan bool

	// Run loop only
	bot                      Bot
	challengeBot             challenge.State
	latestHeader             *types.Header
	pendingDisputableRequest *disputable.AssertionRequest
}

func NewValidator(
	name string,
	bot state.State,
	latestHeader *types.Header,
) *Validator {
	actions := make(chan func(bridge.ArbVMBridge), 100)
	maybeAssert := make(chan bool, 100)
	return &Validator{
		name,
		actions,
		maybeAssert,
		bot,
		nil,
		latestHeader,
		nil,
	}
}

func (validator *Validator) RequestCall(msg protocol.Message) (<-chan value.Value, <-chan error) {
	resultChan := make(chan value.Value, 1)
	errChan := make(chan error, 1)
	validator.actions <- func(bridge bridge.ArbVMBridge) {
		if !validator.canRun() {
			errChan <- errors.New("Cannot call when ArbChannel is not running")
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
	validator.actions <- func(bridge bridge.ArbVMBridge) {
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
	validator.actions <- func(bridge bridge.ArbVMBridge) {
		resultChan <- validator.canRun()
	}
	return resultChan
}

func (validator *Validator) CanContinueRunning() chan bool {
	resultChan := make(chan bool, 1)
	validator.actions <- func(bridge bridge.ArbVMBridge) {
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
	validator.actions <- func(bridge bridge.ArbVMBridge) {
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
	validator.actions <- func(b bridge.ArbVMBridge) {
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
			request := &disputable.AssertionRequest{
				AfterCore:    core.NewCore(mClone, balance),
				Precondition: pre,
				Assertion:    assertion,
				ResultChan:   resultChan,
				ErrorChan:    errChan,
			}
			validator.actions <- func(b bridge.ArbVMBridge) {
				validator.pendingDisputableRequest = request
				validator.maybeAssert <- true
			}
		}()
	}
	return resultChan, errChan
}
