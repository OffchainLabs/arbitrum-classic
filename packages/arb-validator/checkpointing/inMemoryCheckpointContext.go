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

package checkpointing

import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type InMemoryCheckpointContext struct {
	vals  map[common.Hash]value.Value
	machs map[common.Hash]machine.Machine
}

func NewInMemoryCheckpointContext() *InMemoryCheckpointContext {
	return &InMemoryCheckpointContext{
		make(map[common.Hash]value.Value),
		make(map[common.Hash]machine.Machine),
	}
}

func (c *InMemoryCheckpointContext) AddValue(v value.Value) {
	c.vals[v.Hash()] = v
}

func (c *InMemoryCheckpointContext) AddMachine(m machine.Machine) {
	c.machs[m.Hash()] = m.Clone()
}

func (c *InMemoryCheckpointContext) Manifest() *CheckpointManifest {
	valHashes := []*common.HashBuf{}
	for h, _ := range c.vals {
		valHashes = append(valHashes, h.MarshalToBuf())
	}
	machHashes := []*common.HashBuf{}
	for h, _ := range c.machs {
		machHashes = append(machHashes, h.MarshalToBuf())
	}
	return &CheckpointManifest{
		Values:   valHashes,
		Machines: machHashes,
	}
}

func (c *InMemoryCheckpointContext) Values() map[common.Hash]value.Value {
	return c.vals
}

func (c *InMemoryCheckpointContext) Machines() map[common.Hash]machine.Machine {
	return c.machs
}

func (c *InMemoryCheckpointContext) GetValue(h common.Hash) value.Value {
	return c.vals[h]
}

func (c *InMemoryCheckpointContext) GetMachine(h common.Hash) machine.Machine {
	return c.machs[h]
}
