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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"math/big"
)

type DummyCheckpointer struct {
	initialMachine machine.Machine
}

func NewDummyCheckpointer() *DummyCheckpointer {
	return &DummyCheckpointer{nil}
}

func (dcp *DummyCheckpointer) Initialize(arbitrumCodefilePath string) error {
	mach, err := loader.LoadMachineFromFile(arbitrumCodefilePath, true, "cpp")
	if err != nil {
		return err
	}
	dcp.initialMachine = mach
	return nil
}

func (dcp *DummyCheckpointer) Initialized() bool {
	return dcp.initialMachine != nil
}

func (dcp *DummyCheckpointer) GetCheckpointDB() machine.CheckpointStorage {
	return nil
}

func (dcp *DummyCheckpointer) GetConfirmedNodeStore() machine.ConfirmedNodeStore {
	return nil
}

func (dcp *DummyCheckpointer) HasCheckpointedState() bool {
	return false
}

func (dcp *DummyCheckpointer) RestoreLatestState(context.Context, arbbridge.ChainTimeGetter, func([]byte, ckptcontext.RestoreContext) error) error {
	return errors.New("no checkpoints in database")
}

func (dcp *DummyCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return dcp.initialMachine.Clone(), nil
}

func (dcp *DummyCheckpointer) AsyncSaveCheckpoint(_ *common.BlockId, _ []byte, _ *ckptcontext.CheckpointContext) {
}

func (dcp *DummyCheckpointer) CheckpointConfirmedNode(nodeHash common.Hash, depth uint64, nodeData []byte, cpCtx *ckptcontext.CheckpointContext) error {
	return nil
}

func (dcp DummyCheckpointer) MaxReorgHeight() *big.Int {
	return big.NewInt(100)
}
