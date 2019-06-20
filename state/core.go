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
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"

	"github.com/offchainlabs/arb-validator/bridge"
	"github.com/offchainlabs/arb-validator/ethbridge"
	"github.com/offchainlabs/arb-validator/valmessage"
)

type ValidatorState interface {
	UpdateTime(uint64, bridge.Bridge) (ValidatorState, error)
	UpdateState(ethbridge.Event, uint64, bridge.Bridge) (ValidatorState, ChallengeState, error)

	SendMessageToVM(msg protocol.Message)
	GetCore() *validatorCore
	GetConfig() *validatorConfig
}

type ChallengeState interface {
	UpdateTime(uint64, bridge.Bridge) (ChallengeState, error)
	UpdateState(ethbridge.Event, uint64, bridge.Bridge) (ChallengeState, error)
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

type validatorConfig struct {
	Address             common.Address
	ChallengeEverything bool
	Config              *valmessage.VMConfiguration
}

func NewValidatorConfig(address common.Address, config *valmessage.VMConfiguration, challengeEverything bool) *validatorConfig {
	return &validatorConfig{
		address,
		challengeEverything,
		config,
	}
}

func (c *validatorConfig) GetConfig() *validatorConfig {
	return c
}

type validatorCore struct {
	Inbox   *protocol.Inbox
	Balance *protocol.BalanceTracker
	Machine *vm.Machine
}

func NewValidatorCore(inbox *protocol.Inbox, balance *protocol.BalanceTracker, machine *vm.Machine) *validatorCore {
	return &validatorCore{
		Inbox:   inbox,
		Balance: balance,
		Machine: machine,
	}
}

func (c *validatorCore) Clone() *validatorCore {
	return &validatorCore{
		Inbox:   c.Inbox.Clone(),
		Balance: c.Balance.Clone(),
		Machine: c.Machine.Clone(),
	}
}

func (c *validatorCore) OffchainAssert(
	mq *protocol.MessageQueue,
	timeBounds protocol.TimeBounds,
) (*validatorCore, *protocol.Assertion) {
	inbox := c.Inbox.Clone()
	inbox.InsertMessageQueue(mq)
	newState := c.Machine.Clone()
	assCtx := protocol.NewMachineAssertionContext(
		newState,
		c.Balance,
		timeBounds,
		inbox.Receive(),
	)
	newState.RunUntilStop()
	assDef := assCtx.Finalize(newState)

	newAssertion := assDef.GetAssertion()
	newBalance := c.Balance.Clone()
	// This spend is guaranteed to be correct since the VM made sure to only produce on outgoing if it could spend
	_ = newBalance.SpendAll(protocol.NewBalanceTrackerFromMessages(newAssertion.OutMsgs))
	return &validatorCore{
		Inbox:   inbox,
		Balance: newBalance,
		Machine: newState,
	}, newAssertion
}

func (c *validatorCore) GetCore() *validatorCore {
	return c
}

func (c *validatorCore) SendMessageToVM(msg protocol.Message) {
	c.Inbox.SendMessage(msg)
}

func (c *validatorCore) DeliverMessagesToVM() {
	c.Balance.AddAll(c.Inbox.PendingQueue.Balance)
	c.Inbox.DeliverMessages()
}

func (c *validatorCore) GetMachine() *vm.Machine {
	return c.Machine
}

func (c *validatorCore) GetInbox() *protocol.Inbox {
	return c.Inbox
}

func (c *validatorCore) GetBalance() *protocol.BalanceTracker {
	return c.Balance
}

func (c *validatorCore) GeneratePrecondition(beginTime, endTime uint64, includePendingMessages bool) *protocol.Precondition {
	var inboxValue value.Value
	if includePendingMessages {
		inboxValue = c.Inbox.ReceivePending()
	} else {
		inboxValue = c.Inbox.Receive()
	}
	return &protocol.Precondition{
		BeforeHash:    c.Machine.Hash(),
		TimeBounds:    [2]uint64{beginTime, endTime},
		BeforeBalance: c.Balance,
		BeforeInbox:   value.NewHashOnlyValueFromValue(inboxValue),
	}
}

func (c *validatorCore) CreateDisputableDefender(beginTime, length uint64, includePendingMessages bool, maxSteps int32) (*vm.Machine, protocol.AssertionDefender) {
	endTime := beginTime + length
	var inboxValue value.Value
	if includePendingMessages {
		inboxValue = c.Inbox.ReceivePending()
	} else {
		inboxValue = c.Inbox.Receive()
	}
	newState := c.Machine.Clone()
	assCtx := protocol.NewMachineAssertionContext(
		newState,
		c.Balance,
		[2]uint64{beginTime, endTime},
		inboxValue,
	)
	newState.Run(maxSteps)
	assDef := assCtx.Finalize(c.Machine)
	return newState, assDef
}

func (c *validatorCore) ValidateAssertion(pre *protocol.Precondition, time uint64) bool {
	if pre.BeforeInbox.Hash() != c.Inbox.ReceivePending().Hash() && pre.BeforeInbox.Hash() != c.Inbox.Receive().Hash() {
		return false
	}

	if pre.BeforeHash != c.Machine.Hash() {
		return false
	}

	if time < pre.TimeBounds[0] || time > pre.TimeBounds[1] {
		return false
	}

	if !c.Balance.CanSpendAll(pre.BeforeBalance) {
		return false
	}
	return true
}
