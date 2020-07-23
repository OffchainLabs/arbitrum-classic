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
	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type TxDB struct {
	mach         machine.Machine
	as           *cmachine.AggregatorStore
	checkpointer checkpointing.RollupCheckpointer
	timeGetter   arbbridge.ChainTimeGetter

	callMut   sync.Mutex
	callMach  machine.Machine
	callBlock *common.BlockId
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
	blockCreated *common.BlockId,
) (*TxDB, error) {
	txdb := &TxDB{
		as:           as,
		checkpointer: checkpointer,
		timeGetter:   clnt,
	}
	if checkpointer.HasCheckpointedState() {
		if err := txdb.RestoreFromCheckpoint(ctx); err == nil {
			return txdb, nil
		} else {
			log.Println("Failed to restore from checkpoint, falling back to fresh start")
		}
	}

	mach, err := checkpointer.GetInitialMachine()
	if err != nil {
		return nil, err
	}
	prevBlockId, err := clnt.BlockIdForHeight(ctx, common.NewTimeBlocksInt(blockCreated.Height.AsInt().Int64()-1))
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
	if err := as.SaveBlock(prevBlockId, types.Bloom{}); err != nil {
		return nil, err
	}
	txdb.mach = mach
	txdb.callMach = mach.Clone()
	txdb.callBlock = prevBlockId
	return txdb, nil
}

func (txdb *TxDB) RestoreFromCheckpoint(ctx context.Context) error {
	var mach machine.Machine
	var blockId *common.BlockId
	if err := txdb.checkpointer.RestoreLatestState(ctx, txdb.timeGetter, func(chainObserverBytes []byte, restoreCtx ckptcontext.RestoreContext, restoreBlockId *common.BlockId) error {
		var machineHash common.Hash
		copy(machineHash[:], chainObserverBytes)
		mach = restoreCtx.GetMachine(machineHash)
		blockId = restoreBlockId
		return nil
	}); err != nil {
		return err
	}
	if err := txdb.as.RestoreBlock(blockId.Height.AsInt().Uint64()); err != nil {
		return err
	}
	txdb.mach = mach
	txdb.callMut.Lock()
	defer txdb.callMut.Unlock()
	txdb.callMach = mach.Clone()
	txdb.callBlock = blockId
	return nil
}

// The first
func (txdb *TxDB) AddMessages(ctx context.Context, msgs []arbbridge.MessageDeliveredEvent, lastBlock uint64) error {
	log.Println("Processing new messages", msgs)
	blockInboxes, err := txdb.breakByBlock(ctx, msgs, lastBlock)
	if err != nil {
		return err
	}
	for _, bi := range blockInboxes {
		assertion, numSteps := txdb.mach.ExecuteAssertion(1000000000000, bi.inbox, 0)
		if err := txdb.addAssertion(assertion, numSteps, bi.block); err != nil {
			return nil
		}
		txdb.callMut.Lock()
		// If we didn't run, no need to update the machine
		if numSteps > 0 {
			txdb.callMach = txdb.mach.Clone()
		}
		txdb.callBlock = bi.block
		txdb.callMut.Unlock()
	}
	return nil
}

func (txdb *TxDB) CallInfo() (machine.Machine, *common.BlockId) {
	txdb.callMut.Lock()
	defer txdb.callMut.Unlock()
	return txdb.callMach.Clone(), txdb.callBlock
}

func (txdb *TxDB) GetMessage(index uint64) (value.Value, error) {
	return txdb.as.GetMessage(index)
}

func (txdb *TxDB) GetLog(index uint64) (value.Value, error) {
	return txdb.as.GetLog(index)
}

func (txdb *TxDB) GetRequest(requestId common.Hash) (value.Value, error) {
	requestCandidate, err := txdb.as.GetPossibleRequestInfo(requestId)
	if err != nil {
		return nil, err
	}
	logVal, err := txdb.as.GetLog(requestCandidate)
	if err != nil {
		return nil, err
	}
	res, err := evm.NewResultFromValue(logVal)
	if err != nil {
		return nil, err
	}
	if res.L1Message.MessageID() != requestId {
		return nil, errors.New("request not found")
	}
	return logVal, nil
}

func (txdb *TxDB) GetBlock(height uint64) (machine.BlockInfo, error) {
	return txdb.as.GetBlock(height)
}

func (txdb *TxDB) LatestBlock() (*common.BlockId, error) {
	return txdb.as.LatestBlock()
}

func (txdb *TxDB) FindLogs(
	ctx context.Context,
	fromHeight *uint64,
	toHeight *uint64,
	address []common.Address,
	topics [][]common.Hash,
) ([]evm.FullLog, error) {
	latestBlock, err := txdb.LatestBlock()
	if err != nil {
		return nil, err
	}
	startHeight := uint64(0)
	endHeight := latestBlock.Height.AsInt().Uint64()
	if fromHeight != nil && *fromHeight > 0 {
		startHeight = *fromHeight
	}
	if toHeight != nil {
		altEndHeight := *toHeight + 1
		if endHeight > altEndHeight {
			endHeight = altEndHeight
		}
	}
	logs := make([]evm.FullLog, 0)
	if startHeight >= endHeight {
		return logs, nil
	}

	for i := startHeight; i <= endHeight; i++ {
		select {
		case <-ctx.Done():
			return nil, errors.New("call timed out")
		default:
		}
		blockInfo, err := txdb.GetBlock(i)
		if err != nil {
			return nil, err
		}
		if !maybeMatchesLogQuery(blockInfo.Bloom, address, topics) {
			continue
		}

		for j := uint64(0); j < blockInfo.LogCount; j++ {
			logVal, err := txdb.as.GetLog(blockInfo.StartLog + j)
			if err != nil {
				return nil, err
			}

			res, err := evm.NewResultFromValue(logVal)
			if err != nil {
				return nil, err
			}

			logIndex := uint64(0)
			for _, evmLog := range res.EVMLogs {
				if evmLog.MatchesQuery(address, topics) {
					logs = append(logs, evm.FullLog{
						Log:     evmLog,
						TxIndex: j,
						TxHash:  res.L1Message.MessageID(),
						Index:   logIndex,
						Block: &common.BlockId{
							Height:     common.NewTimeBlocks(new(big.Int).SetUint64(i)),
							HeaderHash: blockInfo.Hash,
						},
					})
				}
				logIndex++
			}
		}
	}
	return logs, nil
}

func maybeMatchesLogQuery(logFilter types.Bloom, addresses []common.Address, topics [][]common.Hash) bool {
	if len(addresses) > 0 {
		match := false
		for _, addr := range addresses {
			if logFilter.TestBytes(addr[:]) {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}

	for _, topicGroup := range topics {
		if len(topicGroup) == 0 {
			continue
		}
		match := false
		for _, topic := range topicGroup {
			if logFilter.TestBytes(topic[:]) {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}
	return true
}

func (txdb *TxDB) addAssertion(assertion *protocol.ExecutionAssertion, numSteps uint64, block *common.BlockId) error {
	for _, avmMessage := range assertion.ParseOutMessages() {
		if err := txdb.as.SaveMessage(avmMessage); err != nil {
			return err
		}
	}
	ethLogs := make([]*types.Log, 0)
	for _, avmLog := range assertion.ParseLogs() {
		logIndex, err := txdb.as.LogCount()
		if err != nil {
			return err
		}

		if err := txdb.as.SaveLog(avmLog); err != nil {
			return err
		}

		res, err := evm.NewResultFromValue(avmLog)
		if err != nil {
			log.Println("Error parsing log result", err)
			continue
		}

		if res.L1Message.Kind == message.L2Type {
			l2msg, err := l2message.NewL2MessageFromData(res.L1Message.Data)
			if err != nil {
				log.Println("error NewL2MessageFromData", err)
			}
			log.Printf("msg type%T\n", l2msg)
		}

		if err := txdb.as.SaveRequest(res.L1Message.MessageID(), logIndex); err != nil {
			return err
		}

		for _, evmLog := range res.EVMLogs {
			ethLogs = append(ethLogs, &types.Log{
				Address: evmLog.Address.ToEthAddress(),
				Topics:  common.NewEthHashesFromHashes(evmLog.Topics),
			})
		}
	}
	logBloom := types.BytesToBloom(types.LogsBloom(ethLogs).Bytes())
	if err := txdb.as.SaveBlock(block, logBloom); err != nil {
		return err
	}

	saveMach(txdb.mach, block, txdb.checkpointer)
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

func (txdb *TxDB) breakByBlock(ctx context.Context, msgs []arbbridge.MessageDeliveredEvent, lastBlock uint64) ([]blockInbox, error) {
	blocks := make([]blockInbox, 0)
	latestBlock, err := txdb.as.LatestBlock()
	if err != nil {
		return nil, err
	}
	currentBlockHeight := latestBlock.Height.AsInt().Uint64() + 1
	addEmptyBlock := func() error {
		blockId, err := txdb.timeGetter.BlockIdForHeight(ctx, common.NewTimeBlocksInt(int64(currentBlockHeight)))
		if err != nil {
			return err
		}
		blocks = append(blocks, blockInbox{
			inbox: value.NewEmptyTuple(),
			block: blockId,
		})
		currentBlockHeight++
		return nil
	}

	startHeight := lastBlock
	if len(msgs) > 0 {
		startHeight = msgs[0].BlockId.Height.AsInt().Uint64()
	}

	for currentBlockHeight < startHeight {
		if err := addEmptyBlock(); err != nil {
			return nil, err
		}
	}

	stack := make([]arbbridge.MessageDeliveredEvent, 0)
	for _, msg := range msgs {
		if len(stack) == 0 || msg.BlockId.Height.Cmp(stack[0].BlockId.Height) == 0 {
			stack = append(stack, msg)
		} else {
			nextBlockIndex := makeBlockInbox(stack)
			for nextBlockIndex.block.Height.AsInt().Uint64() > currentBlockHeight {
				if err := addEmptyBlock(); err != nil {
					return nil, err
				}
			}
			blocks = append(blocks, makeBlockInbox(stack))
			stack = make([]arbbridge.MessageDeliveredEvent, 0)
			currentBlockHeight++
		}
	}
	if len(stack) > 0 {
		blocks = append(blocks, makeBlockInbox(stack))
		currentBlockHeight++
	}
	for currentBlockHeight <= lastBlock {
		if err := addEmptyBlock(); err != nil {
			return nil, err
		}
	}
	return blocks, nil
}
