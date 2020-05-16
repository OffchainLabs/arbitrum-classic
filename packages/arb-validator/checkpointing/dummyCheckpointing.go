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
	"github.com/offchainlabs/arbitrum/packages/arb-validator/loader"
	"log"
)

type DummyCheckpointerFactory struct {
	initialMachine machine.Machine
}

func NewDummyCheckpointerFactory(arbitrumCodefilePath string) RollupCheckpointerFactory {
	theMachine, err := loader.LoadMachineFromFile(arbitrumCodefilePath, true, "test")
	if err != nil {
		log.Fatal("newDummyCheckpointer: error loading ", arbitrumCodefilePath)
	}
	return &DummyCheckpointerFactory{theMachine}
}

func (fac *DummyCheckpointerFactory) New(context.Context) RollupCheckpointer {
	return &DummyCheckpointer{fac}
}

type DummyCheckpointer struct {
	fac *DummyCheckpointerFactory
}

func (dcp *DummyCheckpointer) HasCheckpointedState() bool {
	return false
}

func (dcp *DummyCheckpointer) RestoreLatestState(context.Context, arbbridge.ArbClient, func([]byte, RestoreContext) error) error {
	return errors.New("no checkpoints in database")
}

func (dcp *DummyCheckpointer) GetInitialMachine() (machine.Machine, error) {
	return dcp.fac.initialMachine.Clone(), nil
}

func (dcp *DummyCheckpointer) AsyncSaveCheckpoint(_ *common.BlockId, _ []byte, _ *CheckpointContext, closeWhenDone chan struct{}) {
	if closeWhenDone != nil {
		closeWhenDone <- struct{}{}
	}
}
