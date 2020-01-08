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
	"github.com/gogo/protobuf/proto"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"
)

type CheckpointContext interface {
	AddValue(value.Value)
	AddMachine(machine.Machine)
	Manifest() *CheckpointManifest
	Values() map[[32]byte]value.Value
	Machines() map[[32]byte]machine.Machine
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

type RollupCheckpointer struct {
	versionsToKeep *big.Int
	cp             CheckpointerWithMetadata
}

func NewRollupCheckpointer(kind string, versionsToKeep int64) *RollupCheckpointer {
	switch kind {
	case "dummy":
		return &RollupCheckpointer{big.NewInt(versionsToKeep), NewDummyCheckpointer()}
	default:
		return nil
	}
}

func (rcp *RollupCheckpointer) SaveCheckpoint(
	blockHeight *big.Int,
	contents []byte,
	ctx CheckpointContext,
) error {
	var metadataBuf *CheckpointMetadata
	var oldestInCp *big.Int
	var newestInCp *big.Int
	rawMetadata := rcp.cp.RestoreMetadata()
	if rawMetadata == nil || len(rawMetadata) == 0 {
		oldestInCp = blockHeight
		newestInCp = blockHeight
		metadataBuf = &CheckpointMetadata{
			FormatVersion:     1,
			OldestBlockHeight: utils.MarshalBigInt(oldestInCp),
			NewestBlockHeight: utils.MarshalBigInt(newestInCp),
		}
		buf, err := proto.Marshal(metadataBuf)
		if err != nil {
			return err
		}
		rcp.cp.SaveMetadata(buf)
	} else {
		metadataBuf = &CheckpointMetadata{}
		if err := proto.Unmarshal(rawMetadata, metadataBuf); err != nil {
			return err
		}
		oldestInCp = utils.UnmarshalBigInt(metadataBuf.OldestBlockHeight)
		newestInCp = utils.UnmarshalBigInt(metadataBuf.NewestBlockHeight)
		if blockHeight.Cmp(newestInCp) > 0 {
			metadataBuf.NewestBlockHeight = utils.MarshalBigInt(blockHeight)
			buf, err := proto.Marshal(metadataBuf)
			if err != nil {
				return err
			}
			rcp.cp.SaveMetadata(buf)
		}
	}
	rcp.cp.SaveCheckpoint(blockHeight, contents, ctx.Manifest(), ctx.Values(), ctx.Machines())

	if oldestInCp.Cmp(new(big.Int).Sub(newestInCp, rcp.versionsToKeep)) < 0 {
		go func() {
			//TODO: clean up old versions
		}()
	}
	return nil
}

func (rcp *RollupCheckpointer) RestoreCheckpoint(blockHeight *big.Int) ([]byte, RestoreContext, error) {
	var metadataBuf *CheckpointMetadata
	var oldestInCp *big.Int
	var newestInCp *big.Int
	rawMetadata := rcp.cp.RestoreMetadata()
	if rawMetadata == nil {
		return nil, nil, nil
	}

	metadataBuf = &CheckpointMetadata{}
	if err := proto.Unmarshal(rawMetadata, metadataBuf); err != nil {
		return nil, nil, err
	}
	oldestInCp = utils.UnmarshalBigInt(metadataBuf.OldestBlockHeight)
	newestInCp = utils.UnmarshalBigInt(metadataBuf.NewestBlockHeight)

	if blockHeight.Cmp(oldestInCp) < 0 || blockHeight.Cmp(newestInCp) > 0 {
		return nil, nil, nil
	}

	buf, ctx := rcp.cp.RestoreCheckpoint(blockHeight)
	return buf, ctx, nil
}

type CheckpointerWithMetadata interface {
	SaveMetadata([]byte)
	RestoreMetadata() []byte
	SaveCheckpoint(
		blockHeight *big.Int,
		contents []byte,
		manifest *CheckpointManifest,
		values map[[32]byte]value.Value,
		machines map[[32]byte]machine.Machine,
	)
	RestoreCheckpoint(blockHeight *big.Int) ([]byte, RestoreContext) // returns nil, nil if no data at blockHeight
	DeleteCheckpoint(blockHeight *big.Int)
}

type DummyCheckpointer struct {
	metadata []byte
	cp       map[*big.Int]*dummyCheckpoint
}

type dummyCheckpoint struct {
	contents []byte
	manifest *CheckpointManifest
	values   map[[32]byte]value.Value
	machines map[[32]byte]machine.Machine
}

func (dcp *dummyCheckpoint) GetValue(h [32]byte) value.Value {
	return dcp.values[h]
}

func (dcp *dummyCheckpoint) GetMachine(h [32]byte) machine.Machine {
	return dcp.machines[h]
}

func NewDummyCheckpointer() CheckpointerWithMetadata {
	return &DummyCheckpointer{nil, make(map[*big.Int]*dummyCheckpoint)}
}

func (cp *DummyCheckpointer) SaveMetadata(data []byte) {
	cp.metadata = append([]byte{}, data...)
}

func (cp *DummyCheckpointer) RestoreMetadata() []byte {
	return append([]byte{}, cp.metadata...)
}

func (cp *DummyCheckpointer) SaveCheckpoint(
	blockHeight *big.Int,
	contents []byte,
	manifest *CheckpointManifest,
	values map[[32]byte]value.Value,
	machines map[[32]byte]machine.Machine,
) {
	cp.cp[blockHeight] = &dummyCheckpoint{contents, manifest, values, machines}
}

func (cp *DummyCheckpointer) RestoreCheckpoint(blockHeight *big.Int) ([]byte, RestoreContext) {
	dcp := cp.cp[blockHeight]
	if dcp == nil {
		return nil, nil
	} else {
		return dcp.contents, dcp
	}
}

func (cp *DummyCheckpointer) DeleteCheckpoint(blockHeight *big.Int) {
	delete(cp.cp, blockHeight)
}
