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
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"math/big"
)

type RollupCheckpointer interface {
	Initialize(arbitrumCodeFilePath string) error
	Initialized() bool
	HasCheckpointedState() bool
	RestoreLatestState(context.Context, arbbridge.ChainTimeGetter, func([]byte, ckptcontext.RestoreContext, *common.BlockId) error) error
	GetInitialMachine(valueCache machine.ValueCache) (machine.Machine, error)
	AsyncSaveCheckpoint(blockId *common.BlockId, contents []byte, cpCtx *ckptcontext.CheckpointContext) <-chan error

	MaxReorgHeight() *big.Int
}

const checkpointDatabasePathBase = "/tmp/arb-validator-checkpoint-"

func MakeCheckpointDatabasePath(rollupAddr common.Address) string {
	return checkpointDatabasePathBase + rollupAddr.Hex()[2:]
}
