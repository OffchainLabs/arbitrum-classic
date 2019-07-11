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

	"github.com/offchainlabs/arb-util/machine"
	"github.com/offchainlabs/arb-util/protocol"

	"github.com/offchainlabs/arb-validator/valmessage"
)

type Config struct {
	Address             common.Address
	ChallengeEverything bool
	VMConfig            *valmessage.VMConfiguration
}

func NewValidatorConfig(address common.Address, config *valmessage.VMConfiguration, challengeEverything bool) *Config {
	return &Config{
		address,
		challengeEverything,
		config,
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

func (c *Core) OffchainAssert(
	messages []protocol.Message,
	timeBounds protocol.TimeBounds,
	maxSteps int32,
) (*Core, *protocol.Assertion) {
	newState := c.machine.Clone()
	newState.SendOffchainMessages(messages)
	assDef, _ := newState.ExecuteAssertion(
		maxSteps,
		timeBounds,
	)
	return &Core{
		machine: newState,
	}, assDef.GetAssertion()
}

func (c *Core) GetCore() *Core {
	return c
}

func (c *Core) SendMessageToVM(msg protocol.Message) {
	c.machine.SendOnchainMessage(msg)
}

func (c *Core) DeliverMessagesToVM() {
	c.machine.DeliverOnchainMessage()
}

func (c *Core) GetMachine() machine.Machine {
	return c.machine
}

//func (c *Core) GetInbox() *protocol.Inbox {
//	return c.inbox
//}
//
//func (c *Core) GetBalance() *protocol.BalanceTracker {
//	return c.balance
//}

func (c *Core) CreateDisputableDefender(beginTime, length uint64, maxSteps int32) (machine.Machine, machine.AssertionDefender) {
	endTime := beginTime + length
	assDef, _ := c.machine.ExecuteAssertion(
		maxSteps,
		[2]uint64{beginTime, endTime},
	)
	newState := c.machine
	c.machine = assDef.GetMachineState()
	return newState, assDef
}

func (c *Core) ValidateAssertion(pre *protocol.Precondition, time uint64) bool {
	if !c.machine.CheckPrecondition(pre) {
		return false
	}

	if time < pre.TimeBounds[0] || time > pre.TimeBounds[1] {
		return false
	}

	return true
}
