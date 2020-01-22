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
	"fmt"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type CheckpointContext interface {
	AddValue(value.Value)
	AddMachine(machine.Machine)
	Manifest() *CheckpointManifest
	Values() map[common.Hash]value.Value
	Machines() map[common.Hash]machine.Machine
}

type RestoreContext interface {
	GetValue(common.Hash) value.Value
	GetMachine(common.Hash) machine.Machine
}

type CheckpointContextImpl struct {
	values   map[common.Hash]value.Value
	machines map[common.Hash]machine.Machine
}

func NewCheckpointContextImpl() *CheckpointContextImpl {
	return &CheckpointContextImpl{
		values:   make(map[common.Hash]value.Value),
		machines: make(map[common.Hash]machine.Machine),
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
	vals := []*common.HashBuf{}
	for h, _ := range ctx.values {
		vals = append(vals, h.MarshalToBuf())
	}
	machines := []*common.HashBuf{}
	for h, _ := range ctx.machines {
		machines = append(machines, h.MarshalToBuf())
	}
	return &CheckpointManifest{Values: vals, Machines: machines}
}

func (ctx *CheckpointContextImpl) Values() map[common.Hash]value.Value {
	return ctx.values
}

func (ctx *CheckpointContextImpl) Machines() map[common.Hash]machine.Machine {
	return ctx.machines
}

func (ctx *CheckpointContextImpl) GetValue(h common.Hash) value.Value {
	return ctx.values[h]
}

func (ctx *CheckpointContextImpl) GetMachine(h common.Hash) machine.Machine {
	return ctx.machines[h]
}

type SimpleRestoreContext struct {
	values   map[common.Hash]value.Value
	machines map[common.Hash]machine.Machine
}

func NewSimpleRestoreContext() *SimpleRestoreContext {
	return &SimpleRestoreContext{
		values:   make(map[common.Hash]value.Value),
		machines: make(map[common.Hash]machine.Machine),
	}
}

func (src *SimpleRestoreContext) GetValue(h common.Hash) value.Value {
	return src.values[h]
}

func (src *SimpleRestoreContext) GetMachine(h common.Hash) machine.Machine {
	return src.machines[h]
}

func (src *SimpleRestoreContext) AddValue(val value.Value) {
	src.values[val.Hash()] = val
}

func (src *SimpleRestoreContext) AddMachine(mach machine.Machine) {
	src.machines[mach.Hash()] = mach
}

type BlockId struct {
	Height     *common.TimeBlocks
	HeaderHash common.Hash
}

func (id *BlockId) Clone() *BlockId {
	return &BlockId{
		Height:     id.Height.Clone(),
		HeaderHash: id.HeaderHash,
	}
}

func (id *BlockId) MarshalToBuf() *BlockIdBuf {
	return &BlockIdBuf{
		Height:     id.Height.Marshal(),
		HeaderHash: id.HeaderHash.MarshalToBuf(),
	}
}

func (idb *BlockIdBuf) Unmarshal() *BlockId {
	return &BlockId{
		Height:     idb.Height.Unmarshal(),
		HeaderHash: idb.HeaderHash.Unmarshal(),
	}
}

func (id *BlockId) String() string {
	return fmt.Sprintf("Block(%v, %v)", id.Height.AsInt(), id.HeaderHash)
}
