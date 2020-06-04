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
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"log"
)

type RestoreContext interface {
	GetValue(common.Hash) value.Value
	GetMachine(common.Hash) machine.Machine
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

func (ctx *CheckpointContext) GetValue(h common.Hash) value.Value {
	return ctx.values[h]
}

func (ctx *CheckpointContext) GetMachine(h common.Hash) machine.Machine {
	return ctx.machines[h]
}

func SaveCheckpointContext(db machine.CheckpointStorage, ckpCtx *CheckpointContext) error {
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

type SimpleRestore struct {
	db machine.CheckpointStorage
}

func NewSimpleRestore(db machine.CheckpointStorage) *SimpleRestore {
	return &SimpleRestore{db: db}
}

func (sr *SimpleRestore) GetValue(h common.Hash) value.Value {
	return sr.db.GetValue(h)
}

func (sr *SimpleRestore) GetMachine(h common.Hash) machine.Machine {
	ret, err := sr.db.GetMachine(h)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}
