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
	"github.com/ethereum/go-ethereum/common"

	"github.com/offchainlabs/arb-validator/valmessage"

	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-avm/vm"
)

type validatorConfig struct {
	address             common.Address
	challengeEverything bool
	config              *valmessage.VMConfiguration
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
	inbox   *protocol.Inbox
	balance *protocol.BalanceTracker
	machine *vm.Machine
}

func (c *validatorCore) Clone() *validatorCore {
	return &validatorCore{
		inbox:   c.inbox.Clone(),
		balance: c.balance.Clone(),
		machine: c.machine.Clone(),
	}
}

func (c *validatorCore) OffchainAssert(
	mq *protocol.MessageQueue,
	timeBounds protocol.TimeBounds,
) (*validatorCore, *protocol.Assertion) {
	inbox := c.inbox.Clone()
	inbox.InsertMessageQueue(mq)
	newState := c.machine.Clone()
	assCtx := protocol.NewMachineAssertionContext(
		newState,
		c.balance,
		timeBounds,
		inbox.Receive(),
	)
	newState.RunUntilStop()
	assDef := assCtx.Finalize(newState)

	newAssertion := assDef.GetAssertion()
	newBalance := c.balance.Clone()
	// This spend is guaranteed to be correct since the VM made sure to only produce on outgoing if it could spend
	_ = newBalance.SpendAll(protocol.NewBalanceTrackerFromMessages(newAssertion.OutMsgs))
	return &validatorCore{
		inbox:   inbox,
		balance: newBalance,
		machine: newState,
	}, newAssertion
}

func (c *validatorCore) GetCore() *validatorCore {
	return c
}

func (c *validatorCore) SendMessageToVM(msg protocol.Message) {
	c.inbox.SendMessage(msg)
}

func (c *validatorCore) DeliverMessagesToVM() {
	c.balance.AddAll(c.inbox.PendingQueue.Balance)
	c.inbox.DeliverMessages()
}

func (c *validatorCore) GetMachine() *vm.Machine {
	return c.machine
}

func (c *validatorCore) GetInbox() *protocol.Inbox {
	return c.inbox
}

func (c *validatorCore) GetBalance() *protocol.BalanceTracker {
	return c.balance
}

func (c *validatorCore) GeneratePrecondition(beginTime, endTime uint64, includePendingMessages bool) *protocol.Precondition {
	var inboxValue value.Value
	if includePendingMessages {
		inboxValue = c.inbox.ReceivePending()
	} else {
		inboxValue = c.inbox.Receive()
	}
	return &protocol.Precondition{
		BeforeHash:    c.machine.Hash(),
		TimeBounds:    [2]uint64{beginTime, endTime},
		BeforeBalance: c.balance,
		BeforeInbox:   value.NewHashOnlyValueFromValue(inboxValue),
	}
}

func (c *validatorCore) createDisputableDefender(beginTime, length uint64, includePendingMessages bool, maxSteps int32) (*vm.Machine, protocol.AssertionDefender) {
	endTime := beginTime + length
	var inboxValue value.Value
	if includePendingMessages {
		inboxValue = c.inbox.ReceivePending()
	} else {
		inboxValue = c.inbox.Receive()
	}
	newState := c.machine.Clone()
	assCtx := protocol.NewMachineAssertionContext(
		newState,
		c.balance,
		[2]uint64{beginTime, endTime},
		inboxValue,
	)
	newState.Run(maxSteps)
	assDef := assCtx.Finalize(c.machine)
	return newState, assDef
}

func (c *validatorCore) ValidateAssertion(pre *protocol.Precondition, time uint64) bool {
	if pre.BeforeInbox.Hash() != c.inbox.ReceivePending().Hash() && pre.BeforeInbox.Hash() != c.inbox.Receive().Hash() {
		return false
	}

	if pre.BeforeHash != c.machine.Hash() {
		return false
	}

	if time < pre.TimeBounds[0] || time > pre.TimeBounds[1] {
		return false
	}

	if !c.balance.CanSpendAll(pre.BeforeBalance) {
		return false
	}
	return true
}
