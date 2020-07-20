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

package txdb

import (
	"context"
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"log"
)

type TxDB struct {
	mach         machine.Machine
	as           *cmachine.AggregatorStore
	checkpointer checkpointing.RollupCheckpointer
	timeGetter   arbbridge.ChainTimeGetter
}

type blockInbox struct {
	inbox value.TupleValue
	block *common.BlockId
}

func New(
	ctx context.Context,
	clnt arbbridge.ChainTimeGetter,
	checkpointer checkpointing.RollupCheckpointer,
	as *cmachine.AggregatorStore,
	initialBlockId *common.BlockId,
) (*TxDB, error) {
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
			return nil, err
		}
		if err := as.RestoreBlock(blockId.Height.AsInt().Uint64()); err != nil {
			return nil, err
		}
		return &TxDB{
			mach:         mach,
			as:           as,
			checkpointer: checkpointer,
			timeGetter:   clnt,
		}, nil
	}

	mach, err := checkpointer.GetInitialMachine()
	if err != nil {
		return nil, err
	}
	prevBlockId, err := clnt.BlockIdForHeight(ctx, common.NewTimeBlocksInt(initialBlockId.Height.AsInt().Int64()-1))
	if err != nil {
		return nil, err
	}
	// Save initial machine so that next time we will have a checkpointed state
	saveChan := saveMach(mach, prevBlockId, checkpointer)
	select {
	case err := <-saveChan:
		if err != nil {
			return nil, err
		}
	case <-ctx.Done():
		return nil, errors.New("timed out saving checkpoint")
	}
	if err := as.SaveBlock(prevBlockId); err != nil {
		return nil, err
	}
	return &TxDB{
		mach:         mach,
		as:           as,
		checkpointer: checkpointer,
		timeGetter:   clnt,
	}, nil
}

// The first
func (txdb *TxDB) AddMessages(ctx context.Context, msgs []arbbridge.MessageDeliveredEvent) error {
	blockInboxes, err := txdb.breakByBlock(ctx, msgs)
	if err != nil {
		return err
	}
	for _, bi := range blockInboxes {
		assertion, numSteps := txdb.mach.ExecuteAssertion(1000000000000, bi.inbox, 0)
		if err := txdb.addAssertion(assertion, numSteps, bi.block); err != nil {
			return nil
		}
	}
	return nil
}

func (txdb *TxDB) BlockCount() (uint64, error) {
	return txdb.as.BlockCount()
}

func (txdb *TxDB) addAssertion(assertion *protocol.ExecutionAssertion, numSteps uint64, block *common.BlockId) error {
	for _, avmMessage := range assertion.ParseOutMessages() {
		if err := txdb.as.SaveMessage(avmMessage); err != nil {
			return err
		}
	}
	for _, avmLog := range assertion.ParseLogs() {
		if err := txdb.as.SaveLog(avmLog); err != nil {
			return err
		}

		res, err := evm.NewResultFromValue(avmLog)
		if err != nil {
			log.Println("Error parsing log result", err)
			continue
		}
		newLogCount, err := txdb.as.LogCount()
		if err != nil {
			return err
		}
		if err := txdb.as.SaveRequest(res.L1Message.MessageID(), newLogCount); err != nil {
			return err
		}
		// save requestId -> currentLogIndex
		if err := txdb.as.SaveBlock(block); err != nil {
			return err
		}
	}
	// Don't bother saving if we didn't run at all
	if numSteps > 0 {
		saveMach(txdb.mach, block, txdb.checkpointer)
	}
	return nil
}

func saveMach(mach machine.Machine, id *common.BlockId, checkpointer checkpointing.RollupCheckpointer) <-chan error {
	ctx := ckptcontext.NewCheckpointContext()
	ctx.AddMachine(mach)
	machHash := mach.Hash()
	cpData := make([]byte, 32)
	copy(cpData[:], machHash[:])
	return checkpointer.AsyncSaveCheckpoint(id, cpData, ctx)
}

// makeBlockInbox assumes that all messages are from the same block
func makeBlockInbox(msgs []arbbridge.MessageDeliveredEvent) blockInbox {
	inbox := value.NewEmptyTuple()
	for _, msg := range msgs {
		inbox = value.NewTuple2(inbox, msg.Message.AsValue())
	}
	return blockInbox{
		inbox: inbox,
		block: msgs[0].BlockId,
	}
}

func (txdb *TxDB) breakByBlock(ctx context.Context, msgs []arbbridge.MessageDeliveredEvent) ([]blockInbox, error) {
	currentBlockCount, err := txdb.as.BlockCount()
	if err != nil {
		return nil, err
	}
	blocks := make([]blockInbox, 0)
	stack := make([]arbbridge.MessageDeliveredEvent, 0)
	for _, msg := range msgs {
		if len(stack) == 0 || msg.BlockId.Height.Cmp(stack[0].BlockId.Height) == 0 {
			stack = append(stack, msg)
		} else {
			nextBlockIndex := makeBlockInbox(stack)
			for nextBlockIndex.block.Height.AsInt().Uint64() > currentBlockCount {
				blockId, err := txdb.timeGetter.BlockIdForHeight(ctx, common.NewTimeBlocksInt(int64(currentBlockCount)))
				if err != nil {
					return nil, err
				}
				blocks = append(blocks, blockInbox{
					inbox: value.NewEmptyTuple(),
					block: blockId,
				})
				currentBlockCount++
			}
			blocks = append(blocks, makeBlockInbox(stack))
			stack = make([]arbbridge.MessageDeliveredEvent, 0)
		}
	}
	if len(stack) > 0 {
		blocks = append(blocks, makeBlockInbox(stack))
	}
	return blocks, nil
}
