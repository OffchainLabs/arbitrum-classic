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

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/bridge"
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
}

func NewCore(machine machine.Machine) *Core {
	return &Core{
		machine: machine,
	}
}

func (c *Core) Clone() *Core {
	return &Core{
		machine: c.machine.Clone(),
	}
}

func (c *Core) GetCore() *Core {
	return c
}

func (c *Core) SendMessageToVM(msg protocol.Message) {
	c.machine.SendOnchainMessage(msg)
}

func (c *Core) DeliverMessagesToVM(bridge bridge.ArbVMBridge) {
	c.machine.DeliverOnchainMessage()
}

func (c *Core) GetMachine() machine.Machine {
	return c.machine
}

func (c *Core) ValidateAssertion(pre *protocol.Precondition, time uint64) bool {
	if pre.BeforeHashValue() != c.machine.Hash() {
		return false
	}

	if pre.BeforeInboxValue() != c.machine.InboxHash().Hash() {
		return false
	}

	if pre.TimeBounds.IsValidTime(time) != nil {
		return false
	}

	return true
}
