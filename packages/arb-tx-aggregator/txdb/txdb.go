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

import "C"
import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type TxDB struct {
	mach         machine.Machine
	as           *cmachine.AggregatorStore
	checkpointer checkpointing.RollupCheckpointer
	timeGetter   arbbridge.ChainTimeGetter

	callMut     sync.Mutex
	callMach    machine.Machine
	lastBlockId *common.BlockId
}

func New(
	ctx context.Context,
	clnt arbbridge.ChainTimeGetter,
	checkpointer checkpointing.RollupCheckpointer,
	as *cmachine.AggregatorStore,
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
	// We failed to restore from a checkpoint
	mach, err := checkpointer.GetInitialMachine()
	if err != nil {
		return nil, err
	}
	txdb.mach = mach
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
	blockInfo, err := txdb.GetBlock(blockId.Height.AsInt().Uint64())
	if err != nil {
		return err
	}
	if blockInfo == nil {
		return errors.New("should only checkpoint at non-empty blocks")
	}
	block, err := evm.NewBlockResultFromValue(blockInfo.BlockLog)
	if err != nil {
		return err
	}
	if err := txdb.as.Reorg(
		blockId.Height.AsInt().Uint64(),
		block.ChainStats.AVMSendCount.Uint64(),
		block.ChainStats.AVMLogCount.Uint64(),
	); err != nil {
		return err
	}

	txdb.mach = mach
	txdb.callMut.Lock()
	defer txdb.callMut.Unlock()
	txdb.callMach = mach.Clone()
	txdb.lastBlockId = blockId
	return nil
}

func (txdb *TxDB) AddMessages(ctx context.Context, msgs []arbbridge.MessageDeliveredEvent, finishedBlock *common.BlockId) error {
	type resultInfo struct {
		logIndex uint64
		result   *evm.TxResult
	}

	messages := make([]inbox.InboxMessage, 0, len(msgs))
	for _, msg := range msgs {
		messages = append(messages, msg.Message)
	}

	nextBlockHeight := new(big.Int).Add(finishedBlock.Height.AsInt(), big.NewInt(1))
	// TODO: Give ExecuteCallServerAssertion the ability to run unbounded until it blocks
	// The max steps here is a hack since it should just run until it blocks
	assertion, _ := txdb.mach.ExecuteCallServerAssertion(1000000000000, messages, value.NewIntValue(nextBlockHeight), 0)
	for _, avmMessage := range assertion.ParseOutMessages() {
		if err := txdb.as.SaveMessage(avmMessage); err != nil {
			return err
		}
	}

	var lastBlock *common.BlockId
	results := make([]resultInfo, 0)
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

		switch res := res.(type) {
		case *evm.TxResult:
			log.Println("Got result for", res.IncomingRequest.MessageID, res.ResultCode)
			results = append(results, resultInfo{
				logIndex: logIndex,
				result:   res,
			})
		case *evm.BlockInfo:
			ethLogs := make([]*types.Log, 0)
			for _, item := range results {
				for _, evmLog := range item.result.EVMLogs {
					ethLogs = append(ethLogs, &types.Log{
						Address: evmLog.Address.ToEthAddress(),
						Topics:  common.NewEthHashesFromHashes(evmLog.Topics),
					})
				}
			}

			block, err := txdb.timeGetter.BlockIdForHeight(ctx, common.NewTimeBlocks(res.BlockNum))
			if err != nil {
				return err
			}
			log.Println("Saving logs into bloom", ethLogs)
			logBloom := types.BytesToBloom(types.LogsBloom(ethLogs).Bytes())
			log.Println("Saving bloom", hexutil.Encode(logBloom.Bytes()))
			if err := txdb.as.SaveBlock(block, logIndex, logBloom); err != nil {
				return err
			}

			for _, item := range results {
				if err := txdb.as.SaveRequest(item.result.IncomingRequest.MessageID, item.logIndex); err != nil {
					return err
				}
			}
			lastBlock = block
			results = make([]resultInfo, 0)
			txdb.callMut.Lock()
			txdb.callMach = txdb.mach.Clone()
			txdb.lastBlockId = block
			txdb.callMut.Unlock()
		}
	}
	if lastBlock != nil {
		saveMach(txdb.mach, lastBlock, txdb.checkpointer)
	}

	txdb.callMut.Lock()
	if txdb.callMach == nil || txdb.callMach.Hash() != txdb.mach.Hash() {
		txdb.callMach = txdb.mach.Clone()
	}
	txdb.lastBlockId = finishedBlock
	txdb.callMut.Unlock()
	return nil
}

func (txdb *TxDB) CallInfo() (machine.Machine, *common.BlockId) {
	txdb.callMut.Lock()
	defer txdb.callMut.Unlock()
	return txdb.callMach.Clone(), txdb.lastBlockId
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
	res, err := evm.NewTxResultFromValue(logVal)
	if err != nil {
		return nil, err
	}
	if res.IncomingRequest.MessageID != requestId {
		return nil, errors.New("request not found")
	}
	return logVal, nil
}

func (txdb *TxDB) GetBlock(height uint64) (*machine.BlockInfo, error) {
	return txdb.as.GetBlock(height)
}

func (txdb *TxDB) LatestBlockId() *common.BlockId {
	txdb.callMut.Lock()
	defer txdb.callMut.Unlock()
	return txdb.lastBlockId
}

func (txdb *TxDB) FindLogs(
	ctx context.Context,
	fromHeight *uint64,
	toHeight *uint64,
	address []common.Address,
	topics [][]common.Hash,
) ([]evm.FullLog, error) {
	latestBlock := txdb.LatestBlockId()
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
		if blockInfo == nil {
			// No arbitrum txes in this block
			continue
		}
		if !maybeMatchesLogQuery(blockInfo.Bloom, address, topics) {
			continue
		}

		res, err := evm.NewBlockResultFromValue(blockInfo.BlockLog)
		if err != nil {
			return nil, err
		}

		first := res.FirstAVMLog().Uint64()
		for j := uint64(0); j < res.BlockStats.AVMLogCount.Uint64(); j++ {
			logVal, err := txdb.as.GetLog(first + j)
			if err != nil {
				return nil, err
			}

			res, err := evm.NewTxResultFromValue(logVal)
			if err != nil {
				return nil, err
			}

			logIndex := uint64(0)
			for _, evmLog := range res.EVMLogs {
				if evmLog.MatchesQuery(address, topics) {
					logs = append(logs, evm.FullLog{
						Log:     evmLog,
						TxIndex: j,
						TxHash:  res.IncomingRequest.MessageID,
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

func saveMach(mach machine.Machine, id *common.BlockId, checkpointer checkpointing.RollupCheckpointer) <-chan error {
	ctx := ckptcontext.NewCheckpointContext()
	ctx.AddMachine(mach)
	machHash := mach.Hash()
	cpData := make([]byte, 32)
	copy(cpData[:], machHash[:])
	return checkpointer.AsyncSaveCheckpoint(id, cpData, ctx)
}
