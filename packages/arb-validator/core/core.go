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

package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type Config struct {
	Address             common.Address
	ChallengeEverything bool
	VMConfig            *valmessage.VMConfiguration
	MaxCallSteps        int32
}

func NewValidatorConfig(address common.Address, config *valmessage.VMConfiguration, challengeEverything bool, maxCallSteps int32) *Config {
	return &Config{
		address,
		challengeEverything,
		config,
		maxCallSteps,
	}
}

func (c *Config) GetConfig() *Config {
	return c
}

type Core struct {
	machine machine.Machine
	balance *protocol.BalanceTracker
}

func NewCore(machine machine.Machine, balance *protocol.BalanceTracker) *Core {
	return &Core{
		machine: machine,
		balance: balance,
	}
}

func (c *Core) Clone() *Core {
	return &Core{
		machine: c.machine.Clone(),
		balance: c.balance.Clone(),
	}
}

func (c *Core) GetCore() *Core {
	return c
}

func (c *Core) SendMessageToVM(msg protocol.Message) {
	c.machine.SendOnchainMessage(msg)
	c.balance.Add(msg.TokenType, msg.Currency)
}

func (c *Core) DeliverMessagesToVM(bridge bridge.Bridge) {
	bridge.AddedNewMessages(c.machine.PendingMessageCount())
	c.machine.DeliverOnchainMessage()
}

func (c *Core) GetMachine() machine.Machine {
	return c.machine
}

func (c *Core) GetBalance() *protocol.BalanceTracker {
	return c.balance
}

func (c *Core) ValidateAssertion(pre *protocol.Precondition, time uint64) bool {
	if pre.BeforeHash != c.machine.Hash() {
		return false
	}

	if !pre.BeforeInbox.Equal(c.machine.InboxHash()) {
		return false
	}

	if c.balance.CanSpendAll(pre.BeforeBalance) {
		return false
	}

	if time < pre.TimeBounds[0] || time > pre.TimeBounds[1] {
		return false
	}

	return true
}
