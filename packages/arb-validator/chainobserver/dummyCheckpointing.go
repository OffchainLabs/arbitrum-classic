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

package chainobserver

import (
	"context"
	"github.com/pkg/errors"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
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

func (dcp *DummyCheckpointer) HasCheckpointedState() bool {
	return false
}

func (dcp *DummyCheckpointer) RestoreLatestState(context.Context, arbbridge.ChainTimeGetter, func([]byte, ckptcontext.RestoreContext, *common.BlockId) error) error {
	return errors.New("no checkpoints in database")
}

func (dcp *DummyCheckpointer) GetInitialMachine(valueCache machine.ValueCache) (machine.Machine, error) {
	return dcp.initialMachine.Clone(), nil
}

func (dcp *DummyCheckpointer) AsyncSaveCheckpoint(_ *common.BlockId, _ []byte, _ *ckptcontext.CheckpointContext) <-chan error {
	return nil
}

func (dcp DummyCheckpointer) MaxReorgHeight() *big.Int {
	return big.NewInt(100)
}
