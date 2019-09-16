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

package chainvalidator

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/state"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/disputable"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/challenge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/core"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethconnection"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type Validator struct {
	Name        string
	actions     chan func(*Validator, bridge.ArbVMBridge)
	maybeAssert chan bool

	// Run loop only
	bot                      state.ChainState
	challengeBot             challenge.State
	latestHeader             *types.Header
	pendingDisputableRequest *disputable.AssertionRequest
}

func NewValidator(
	name string,
	address common.Address,
	latestHeader *types.Header,
	balance *protocol.BalanceTracker,
	config *valmessage.VMConfiguration,
	machine machine.Machine,
	challengeEverything bool,
	maxCallSteps int32,
) *Validator {
	actions := make(chan func(*Validator, bridge.ArbVMBridge), 100)
	maybeAssert := make(chan bool, 100)
	c := core.NewCore(
		machine,
		balance,
	)

	valConfig := core.NewValidatorConfig(address, config, challengeEverything, maxCallSteps)
	return &Validator{
		name,
		actions,
		maybeAssert,
		state.NewWaiting(valConfig, c),
		nil,
		latestHeader,
		nil,
	}
}

func (validator *Validator) RequestCall(msg protocol.Message) (<-chan value.Value, <-chan error) {
	resultChan := make(chan value.Value, 1)
	errChan := make(chan error, 1)
	validator.actions <- func(validator *Validator, bridge bridge.ArbVMBridge) {
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
	validator.actions <- func(validator *Validator, bridge bridge.ArbVMBridge) {
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
	validator.actions <- func(validator *Validator, bridge bridge.ArbVMBridge) {
		resultChan <- validator.canRun()
	}
	return resultChan
}

func (validator *Validator) CanContinueRunning() chan bool {
	resultChan := make(chan bool, 1)
	validator.actions <- func(validator *Validator, bridge bridge.ArbVMBridge) {
		if !validator.canRun() {
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
	validator.actions <- func(validator *Validator, bridge bridge.ArbVMBridge) {
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
	validator.actions <- func(validator *Validator, b bridge.ArbVMBridge) {
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
			validator.actions <- func(validator *Validator, b bridge.ArbVMBridge) {
				validator.pendingDisputableRequest = request
				validator.maybeAssert <- true
			}
		}()
	}
	return resultChan, errChan
}

func (validator *Validator) Run(recvChan <-chan ethconnection.Notification, bridge bridge.Bridge, ctx context.Context) {
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
					if !validator.bot.GetCore().ValidateAssertion(pre, newHeader.Number.Uint64()) {
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
	newBot, err := validator.bot.ChainUpdateTime(validator.latestHeader.Number.Uint64(), bridge)
	if err != nil {
		fmt.Printf("%v: Error %v responding to event by %T\n", validator.Name, err, newBot)
		return
	}
	validator.bot = newBot
}

func (validator *Validator) eventUpdate(ev ethconnection.VMEvent, header *types.Header, bridge bridge.Bridge) {
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
		newBot, challengeBot, err := validator.bot.ChainUpdateState(ev, header.Number.Uint64(), bridge)
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
