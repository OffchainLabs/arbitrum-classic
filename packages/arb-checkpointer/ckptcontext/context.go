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

package ckptcontext

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
)

type RestoreContext interface {
	GetValue(common.Hash) (value.Value, error)
	GetMachine(common.Hash) (machine.Machine, error)
}

type CheckpointContext struct {
	values   map[common.Hash]value.Value
	machines map[common.Hash]machine.Machine
}

func NewCheckpointContext() *CheckpointContext {
	return &CheckpointContext{
		values:   make(map[common.Hash]value.Value),
		machines: make(map[common.Hash]machine.Machine),
	}
}

func (ctx *CheckpointContext) AddValue(val value.Value) {
	ctx.values[val.Hash()] = val
}

func (ctx *CheckpointContext) AddValueWithHash(val value.Value, valueHash common.Hash) {
	ctx.values[valueHash] = val
}

func (ctx *CheckpointContext) AddMachine(mach machine.Machine) {
	if ctx.machines[mach.Hash()] == nil {
		ctx.machines[mach.Hash()] = mach.Clone()
	}
}

func (ctx *CheckpointContext) Manifest() *CheckpointManifest {
	vals := make([]*common.HashBuf, 0, len(ctx.values))
	for h := range ctx.values {
		vals = append(vals, h.MarshalToBuf())
	}
	machines := make([]*common.HashBuf, 0, len(ctx.machines))
	for h := range ctx.machines {
		machines = append(machines, h.MarshalToBuf())
	}
	return &CheckpointManifest{Values: vals, Machines: machines}
}

func (ctx *CheckpointContext) Values() map[common.Hash]value.Value {
	return ctx.values
}

func (ctx *CheckpointContext) Machines() map[common.Hash]machine.Machine {
	return ctx.machines
}

func (ctx *CheckpointContext) GetValue(h common.Hash) (value.Value, error) {
	if val, ok := ctx.values[h]; ok {
		return val, nil
	}

	return nil, &machine.ValueNotFoundError{HashValue: h}
}

func (ctx *CheckpointContext) GetMachine(h common.Hash) (machine.Machine, error) {
	if mach, ok := ctx.machines[h]; ok {
		return mach, nil
	}

	return nil, &machine.MachineNotFoundError{HashValue: h}
}

func SaveCheckpointContext(db machine.ArbStorage, ckpCtx *CheckpointContext) error {
	for _, val := range ckpCtx.Values() {
		if ok := db.SaveValue(val); !ok {
			return errors.New("failed to write value to checkpoint db")
		}
	}
	for _, mach := range ckpCtx.Machines() {
		if ok := mach.Checkpoint(db); !ok {
			return errors.New("failed to write machine to checkpoint db")
		}
	}
	return nil
}
