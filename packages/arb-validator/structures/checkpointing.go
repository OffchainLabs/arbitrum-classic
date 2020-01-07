/*
* Copyright 2020, Offchain Labs, Inc.
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

package structures

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type CheckpointContext interface {
	AddValue(value.Value)
	AddMachine(machine.Machine)
}

type CheckpointContextImpl struct {
	values   map[[32]byte]value.Value
	machines map[[32]byte]machine.Machine
}

type RestoreContext interface {
	GetValue([32]byte) value.Value
	GetMachine([32]byte) machine.Machine
}

func NewCheckpointContextImpl() *CheckpointContextImpl {
	return &CheckpointContextImpl{
		values:   make(map[[32]byte]value.Value),
		machines: make(map[[32]byte]machine.Machine),
	}
}

func (ctx *CheckpointContextImpl) AddValue(val value.Value) {
	ctx.values[val.Hash()] = val
}

func (ctx *CheckpointContextImpl) AddMachine(mach machine.Machine) {
	ctx.machines[mach.Hash()] = mach
}

func (ctx *CheckpointContextImpl) GetValue(h [32]byte) value.Value {
	return ctx.values[h]
}

func (ctx *CheckpointContextImpl) GetMachine(h [32]byte) machine.Machine {
	return ctx.machines[h]
}
