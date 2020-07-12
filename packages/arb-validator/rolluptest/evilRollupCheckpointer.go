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

package rolluptest

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ckptcontext"
	"math/big"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/checkpointing"
)

func NewEvilRollupCheckpointer(
	rollupAddr common.Address,
	databasePath string,
	maxReorgDepth *big.Int,
	forceFreshStart bool,
) *EvilRollupCheckpointer {
	return &EvilRollupCheckpointer{
		checkpointing.NewIndexedCheckpointer(
			rollupAddr,
			databasePath,
			maxReorgDepth,
			forceFreshStart,
		),
	}
}

type EvilRollupCheckpointer struct {
	cp checkpointing.RollupCheckpointer
}

func (e *EvilRollupCheckpointer) Initialize(arbitrumCodeFilePath string) error {
	return e.cp.Initialize(arbitrumCodeFilePath)
}

func (e *EvilRollupCheckpointer) Initialized() bool {
	return e.cp.Initialized()
}

func (e EvilRollupCheckpointer) GetValue(h common.Hash) value.Value {
	return e.cp.(ckptcontext.RestoreContext).GetValue(h)
}

func (e EvilRollupCheckpointer) GetMachine(h common.Hash) machine.Machine {
	return NewEvilMachine(e.cp.(ckptcontext.RestoreContext).GetMachine(h).(*cmachine.Machine))
}

func (e *EvilRollupCheckpointer) GetCheckpointDB() machine.CheckpointStorage {
	return e.cp.GetCheckpointDB()
}

func (e *EvilRollupCheckpointer) GetConfirmedNodeStore() machine.ConfirmedNodeStore {
	return e.cp.GetConfirmedNodeStore()
}

func (e EvilRollupCheckpointer) HasCheckpointedState() bool {
	return e.cp.HasCheckpointedState()
}

func (e EvilRollupCheckpointer) RestoreLatestState(
	ctx context.Context,
	clnt arbbridge.ChainTimeGetter,
	unmarshalFunc func([]byte, ckptcontext.RestoreContext) error,
) error {
	return e.cp.RestoreLatestState(
		ctx,
		clnt,
		func(contents []byte, resCtx ckptcontext.RestoreContext) error {
			return unmarshalFunc(contents, &evilRestoreContext{resCtx})
		},
	)
}

type evilRestoreContext struct {
	rc ckptcontext.RestoreContext
}

func (erc *evilRestoreContext) GetValue(h common.Hash) value.Value {
	return erc.rc.GetValue(h)
}

func (erc *evilRestoreContext) GetMachine(h common.Hash) machine.Machine {
	return NewEvilMachine(erc.rc.GetMachine(h).(*cmachine.Machine))
}

func (e EvilRollupCheckpointer) GetInitialMachine() (machine.Machine, error) {
	m, err := e.cp.GetInitialMachine()
	if err != nil {
		return m, err
	}
	return NewEvilMachine(m.(*cmachine.Machine)), nil
}

func (e EvilRollupCheckpointer) AsyncSaveCheckpoint(blockId *common.BlockId, contents []byte, cpCtx *ckptcontext.CheckpointContext) {
	e.cp.AsyncSaveCheckpoint(blockId, contents, cpCtx)
}

func (e EvilRollupCheckpointer) CheckpointConfirmedNode(nodeHash common.Hash, depth uint64, nodeData []byte, cpCtx *ckptcontext.CheckpointContext) error {
	return e.cp.CheckpointConfirmedNode(nodeHash, depth, nodeData, cpCtx)
}

func (e EvilRollupCheckpointer) MaxReorgHeight() *big.Int {
	return e.cp.MaxReorgHeight()
}
