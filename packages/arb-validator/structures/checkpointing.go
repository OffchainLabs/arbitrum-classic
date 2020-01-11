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
	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type CheckpointContext interface {
	AddValue(value.Value)
	AddMachine(machine.Machine)
	Manifest() *CheckpointManifest
	Values() map[[32]byte]value.Value
	Machines() map[[32]byte]machine.Machine
}

type RestoreContext interface {
	GetValue([32]byte) value.Value
	GetMachine([32]byte) machine.Machine
}

type CheckpointContextImpl struct {
	values   map[[32]byte]value.Value
	machines map[[32]byte]machine.Machine
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
	if ctx.machines[mach.Hash()] == nil {
		ctx.machines[mach.Hash()] = mach.Clone()
	}
}

func (ctx *CheckpointContextImpl) Manifest() *CheckpointManifest {
	vals := []*value.HashBuf{}
	for h, _ := range ctx.values {
		vals = append(vals, utils.MarshalHash(h))
	}
	machines := []*value.HashBuf{}
	for h, _ := range ctx.machines {
		machines = append(machines, utils.MarshalHash(h))
	}
	return &CheckpointManifest{Values: vals, Machines: machines}
}

func (ctx *CheckpointContextImpl) Values() map[[32]byte]value.Value {
	return ctx.values
}

func (ctx *CheckpointContextImpl) Machines() map[[32]byte]machine.Machine {
	return ctx.machines
}

func (ctx *CheckpointContextImpl) GetValue(h [32]byte) value.Value {
	return ctx.values[h]
}

func (ctx *CheckpointContextImpl) GetMachine(h [32]byte) machine.Machine {
	return ctx.machines[h]
}

type SimpleRestoreContext struct {
	values   map[[32]byte]value.Value
	machines map[[32]byte]machine.Machine
}

func NewSimpleRestoreContext() *SimpleRestoreContext {
	return &SimpleRestoreContext{
		values:   make(map[[32]byte]value.Value),
		machines: make(map[[32]byte]machine.Machine),
	}
}

func (src *SimpleRestoreContext) GetValue(h [32]byte) value.Value {
	return src.values[h]
}

func (src *SimpleRestoreContext) GetMachine(h [32]byte) machine.Machine {
	return src.machines[h]
}

func (src *SimpleRestoreContext) AddValue(val value.Value) {
	src.values[val.Hash()] = val
}

func (src *SimpleRestoreContext) AddMachine(mach machine.Machine) {
	src.machines[mach.Hash()] = mach
}
