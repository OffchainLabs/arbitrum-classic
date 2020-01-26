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
	"context"
	"errors"
	"log"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
)

type DummyCheckpointer struct {
	initialMachine machine.Machine
}

func NewDummyCheckpointer(arbitrumCodefilePath string) *DummyCheckpointer {
	theMachine, err := loader.LoadMachineFromFile(arbitrumCodefilePath, true, "test")
	if err != nil {
		log.Fatal("newDummyCheckpointer: error loading ", arbitrumCodefilePath)
	}
	return &DummyCheckpointer{theMachine}
}

func (dcp *DummyCheckpointer) HasCheckpointedState() bool {
	return false
}

func (dcp *DummyCheckpointer) RestoreLatestState(ctx context.Context, client arbbridge.ArbClient, contractAddr common.Address, beOpinionated bool) ([]byte, structures.RestoreContext, error) {
	return nil, nil, errors.New("no checkpoints in database")
}

func (dcp *DummyCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return dcp.initialMachine.Clone(), nil
}

func (dcp *DummyCheckpointer) AsyncSaveCheckpoint(blockId *structures.BlockId, contents []byte, cpCtx structures.CheckpointContext, closeWhenDone chan struct{}) {
	if closeWhenDone != nil {
		closeWhenDone <- struct{}{}
	}
}

type dummyCheckpointer struct {
	metadata       []byte
	cp             map[*structures.BlockId]*dummyCheckpoint
	initialMachine machine.Machine
}

func newDummyCheckpointer(contractPath string) *dummyCheckpointer {
	theMachine, err := loader.LoadMachineFromFile(contractPath, true, "test")
	if err != nil {
		log.Fatal("newDummyCheckpointer: error loading ", contractPath)
	}
	return &dummyCheckpointer{
		nil,
		make(map[*structures.BlockId]*dummyCheckpoint),
		theMachine,
	}
}

type dummyCheckpoint struct {
	contents []byte
	manifest *structures.CheckpointManifest
	values   map[common.Hash]value.Value
	machines map[common.Hash]machine.Machine
}

func (dcp *dummyCheckpoint) GetValue(h common.Hash) value.Value {
	return dcp.values[h]
}

func (dcp *dummyCheckpoint) GetMachine(h common.Hash) machine.Machine {
	return dcp.machines[h]
}

func (cp *dummyCheckpointer) SaveMetadata(data []byte) {
	cp.metadata = append([]byte{}, data...)
}

func (cp *dummyCheckpointer) RestoreMetadata() []byte {
	return append([]byte{}, cp.metadata...)
}

func (cp *dummyCheckpointer) SaveCheckpoint(
	id *structures.BlockId,
	contents []byte,
	manifest *structures.CheckpointManifest,
	values map[common.Hash]value.Value,
	machines map[common.Hash]machine.Machine,
) {
	cp.cp[id] = &dummyCheckpoint{contents, manifest, values, machines}
}

func (cp *dummyCheckpointer) RestoreCheckpoint(blockId *structures.BlockId) ([]byte, structures.RestoreContext) {
	dcp := cp.cp[blockId]
	if dcp == nil {
		return nil, nil
	} else {
		return dcp.contents, dcp
	}
}

func (cp *dummyCheckpointer) DeleteOldCheckpoints(earliestRollbackPoint *big.Int) {
	// ignore cleanup requests; we're being simple and inefficient for debugging
}

func (cp *dummyCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return cp.initialMachine.Clone(), nil
}
