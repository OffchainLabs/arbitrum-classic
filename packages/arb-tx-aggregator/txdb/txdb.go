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
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/checkpointing"
	"github.com/offchainlabs/arbitrum/packages/arb-checkpointer/ckptcontext"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

type TxDB struct {
	View
	mach         machine.Machine
	checkpointer checkpointing.RollupCheckpointer
	timeGetter   arbbridge.ChainTimeGetter
	chain        common.Address

	callMut            sync.Mutex
	callMach           machine.Machine
	lastBlockProcessed *common.BlockId
	lastInboxSeq       *big.Int
}

func New(
	ctx context.Context,
	clnt arbbridge.ChainTimeGetter,
	checkpointer checkpointing.RollupCheckpointer,
	as *cmachine.AggregatorStore,
	chain common.Address,
) (*TxDB, error) {
	txdb := &TxDB{
		View:         View{as: as},
		checkpointer: checkpointer,
		timeGetter:   clnt,
		chain:        chain,
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
	txdb.callMach = mach.Clone()
	txdb.lastInboxSeq = big.NewInt(0)
	return txdb, nil
}

func (txdb *TxDB) RestoreFromCheckpoint(ctx context.Context) error {
	var mach machine.Machine
	var blockId *common.BlockId
	var lastInboxSeq *big.Int
	if err := txdb.checkpointer.RestoreLatestState(ctx, txdb.timeGetter, func(chainObserverBytes []byte, restoreCtx ckptcontext.RestoreContext, restoreBlockId *common.BlockId) error {
		var machineHash common.Hash
		copy(machineHash[:], chainObserverBytes)
		lastInboxSeq = new(big.Int).SetBytes(chainObserverBytes[32:])
		mach = restoreCtx.GetMachine(machineHash)
		blockId = restoreBlockId
		return nil
	}); err != nil {
		return err
	}
	blockInfo, err := txdb.as.GetBlock(blockId.Height.AsInt().Uint64())
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
	txdb.lastBlockProcessed = blockId
	txdb.lastInboxSeq = lastInboxSeq
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
			logBloom := types.BytesToBloom(types.LogsBloom(ethLogs).Bytes())
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
			txdb.lastBlockProcessed = block
			txdb.callMut.Unlock()
		}
	}

	txdb.callMut.Lock()
	if txdb.callMach == nil || txdb.callMach.Hash() != txdb.mach.Hash() {
		txdb.callMach = txdb.mach.Clone()
	}
	if len(messages) > 0 {
		txdb.lastInboxSeq = messages[len(messages)-1].InboxSeqNum
	}
	txdb.lastBlockProcessed = finishedBlock
	lastInboxSeq := txdb.lastInboxSeq
	txdb.callMut.Unlock()

	if lastBlock != nil {
		ctx := ckptcontext.NewCheckpointContext()
		ctx.AddMachine(txdb.mach)
		machHash := txdb.mach.Hash()
		cpData := make([]byte, 64)
		copy(cpData[:], machHash[:])
		copy(cpData[32:], math.U256Bytes(lastInboxSeq))
		txdb.checkpointer.AsyncSaveCheckpoint(lastBlock, cpData, ctx)
	}
	return nil
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
	latest := txdb.LatestBlock()
	if height > latest.Height.AsInt().Uint64() {
		return nil, nil
	}
	return txdb.as.GetBlock(height)
}

func (txdb *TxDB) LatestBlock() *common.BlockId {
	block, err := txdb.as.LatestBlock()
	if err != nil {
		return txdb.lastBlockProcessed
	}
	return block
}

func (txdb *TxDB) LatestSnapshot() *snapshot.Snapshot {
	txdb.callMut.Lock()
	defer txdb.callMut.Unlock()
	currentTime := inbox.ChainTime{
		BlockNum:  txdb.lastBlockProcessed.Height,
		Timestamp: big.NewInt(time.Now().Unix()),
	}
	return snapshot.NewSnapshot(txdb.callMach.Clone(), currentTime, txdb.lastInboxSeq, message.ChainAddressToID(txdb.chain))
}

func (txdb *TxDB) LatestBlockId() *common.BlockId {
	txdb.callMut.Lock()
	defer txdb.callMut.Unlock()
	return txdb.lastBlockProcessed
}
