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

package machineobserver

import (
	"context"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

func RestoreLatestMachine(
	ctx context.Context,
	clnt arbbridge.ArbClient,
	checkpointer checkpointing.RollupCheckpointer,
	initialBlockId *common.BlockId,
) (machine.Machine, *common.BlockId, error) {
	if checkpointer.HasCheckpointedState() {
		var mach machine.Machine
		var blockId *common.BlockId
		if err := checkpointer.RestoreLatestState(ctx, clnt, func(chainObserverBytes []byte, restoreCtx ckptcontext.RestoreContext, restoreBlockId *common.BlockId) error {
			var machineHash common.Hash
			copy(machineHash[:], chainObserverBytes)
			mach = restoreCtx.GetMachine(machineHash)
			blockId = restoreBlockId
			return nil
		}); err != nil {
			return nil, nil, err
		}
		return mach, blockId, nil
	}

	mach, err := checkpointer.GetInitialMachine()
	return mach, initialBlockId, err
}

type MachineObserver struct {
	mach           machine.Machine
	logsChan       chan value.Value
	currentBlockId *common.BlockId
	prevBlockId    *common.BlockId
	checkpointer   checkpointing.RollupCheckpointer
}

func (mo *MachineObserver) processNextMessage(ev arbbridge.MessageDeliveredEvent) {
	inbox := value.NewTuple2(value.NewEmptyTuple(), ev.Message.AsValue())
	assertion, _ := mo.mach.ExecuteAssertion(100000000000000, inbox, 0)

	for _, vmLog := range assertion.ParseLogs() {
		mo.logsChan <- vmLog
	}

	if ev.BlockId.Height.Cmp(mo.prevBlockId.Height) > 0 {
		mo.currentBlockId = mo.prevBlockId
		mo.prevBlockId = ev.BlockId

		ctx := ckptcontext.NewCheckpointContext()
		ctx.AddMachine(mo.mach)
		machHash := mo.mach.Hash()
		cpData := make([]byte, 32)
		copy(cpData[:], machHash[:])
		mo.checkpointer.AsyncSaveCheckpoint(mo.prevBlockId, cpData, ctx)
	}
}
