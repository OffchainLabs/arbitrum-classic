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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"log"
	"math/big"
	"sync"
)

var snapshotCacheSize = 10

type TxDB struct {
	View
	mach         machine.Machine
	checkpointer checkpointing.RollupCheckpointer
	timeGetter   arbbridge.ChainTimeGetter
	chain        common.Address

	callMut            sync.Mutex
	lastBlockProcessed *common.BlockId
	lastInboxSeq       *big.Int
	snapCache          *snapshotCache
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
		snapCache:    newSnapshotCache(snapshotCacheSize),
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
	txdb.lastInboxSeq = big.NewInt(0)
	return txdb, nil
}

// addSnap must be called with callMut locked or during construction
func (txdb *TxDB) addSnap(blockNum *big.Int, timestamp *big.Int) {
	currentTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks(new(big.Int).Set(blockNum)),
		Timestamp: new(big.Int).Set(timestamp),
	}
	snap := snapshot.NewSnapshot(txdb.mach.Clone(), currentTime, message.ChainAddressToID(txdb.chain), new(big.Int).Set(txdb.lastInboxSeq))
	txdb.snapCache.addSnapshot(snap)
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
	txdb.lastBlockProcessed = blockId
	txdb.lastInboxSeq = lastInboxSeq
	txdb.addSnap(block.BlockNum, block.Timestamp)
	return nil
}

func (txdb *TxDB) AddMessages(ctx context.Context, msgs []arbbridge.MessageDeliveredEvent, finishedBlock *common.BlockId) error {
	timestamp, err := txdb.timeGetter.TimestampForBlockHash(ctx, finishedBlock.HeaderHash)
	if err != nil {
		return err
	}

	var lastBlock *common.BlockId
	for _, msg := range msgs {
		block, err := txdb.processMessage(ctx, msg.Message)
		if err != nil {
			return err
		}
		if block != nil {
			lastBlock = block
		}
	}

	nextBlockHeight := new(big.Int).Add(finishedBlock.Height.AsInt(), big.NewInt(1))
	// TODO: Give ExecuteCallServerAssertion the ability to run unbounded until it blocks
	// The max steps here is a hack since it should just run until it blocks
	assertion, _ := txdb.mach.ExecuteCallServerAssertion(1000000000000, nil, value.NewIntValue(nextBlockHeight), 0)
	block, err := txdb.processAssertion(ctx, assertion)
	if err != nil {
		return err
	}
	if block != nil {
		lastBlock = block
	}

	txdb.callMut.Lock()
	txdb.lastBlockProcessed = finishedBlock
	lastInboxSeq := new(big.Int).Set(txdb.lastInboxSeq)

	latestSnap := txdb.snapCache.latest()
	if latestSnap == nil || latestSnap.Height().Cmp(finishedBlock.Height) < 0 {
		txdb.addSnap(finishedBlock.Height.AsInt(), timestamp)
	}
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

func (txdb *TxDB) processMessage(ctx context.Context, msg inbox.InboxMessage) (*common.BlockId, error) {
	// TODO: Give ExecuteAssertion the ability to run unbounded until it blocks
	// The max steps here is a hack since it should just run until it blocks
	assertion, _ := txdb.mach.ExecuteAssertion(1000000000000, []inbox.InboxMessage{msg}, 0)
	txdb.callMut.Lock()
	txdb.lastInboxSeq = msg.InboxSeqNum
	txdb.callMut.Unlock()
	return txdb.processAssertion(ctx, assertion)
}

func (txdb *TxDB) processAssertion(ctx context.Context, assertion *protocol.ExecutionAssertion) (*common.BlockId, error) {
	for _, avmMessage := range assertion.ParseOutMessages() {
		if err := txdb.as.SaveMessage(avmMessage); err != nil {
			return nil, err
		}
	}

	var lastBlock *common.BlockId
	avmLogs := assertion.ParseLogs()
	for i, avmLog := range avmLogs {
		logIndex, err := txdb.as.LogCount()
		if err != nil {
			return nil, err
		}

		if err := txdb.as.SaveLog(avmLog); err != nil {
			return nil, err
		}

		res, err := evm.NewResultFromValue(avmLog)
		if err != nil {
			log.Println("Error parsing log result", err)
			continue
		}

		switch res := res.(type) {
		case *evm.TxResult:
			log.Println("Got result for", res.IncomingRequest.MessageID, res.ResultCode)
			if err := txdb.as.SaveRequest(res.IncomingRequest.MessageID, logIndex); err != nil {
				return nil, err
			}
		case *evm.BlockInfo:
			if i != len(avmLogs)-1 {
				return nil, errors.New("block info should only come at end of assertion")
			}

			txCount := res.BlockStats.TxCount.Uint64()
			startLog := res.FirstAVMLog().Uint64()
			ethLogs := make([]*types.Log, 0)
			txResults := make([]*evm.TxResult, 0, txCount)
			for i := uint64(0); i < txCount; i++ {
				avmLog, err := txdb.GetLog(startLog + i)
				if err != nil {
					return nil, err
				}
				txRes, err := evm.NewTxResultFromValue(avmLog)
				if err != nil {
					return nil, err
				}
				txResults = append(txResults, txRes)
			}

			for _, txRes := range txResults {
				for _, evmLog := range txRes.EVMLogs {
					ethLogs = append(ethLogs, &types.Log{
						Address: evmLog.Address.ToEthAddress(),
						Topics:  common.NewEthHashesFromHashes(evmLog.Topics),
					})
				}
			}

			block, err := txdb.timeGetter.BlockIdForHeight(ctx, common.NewTimeBlocks(res.BlockNum))
			if err != nil {
				return nil, err
			}
			logBloom := types.BytesToBloom(types.LogsBloom(ethLogs).Bytes())
			if err := txdb.as.SaveBlock(block, logIndex, logBloom); err != nil {
				return nil, err
			}

			for i, txRes := range txResults {
				if err := txdb.as.SaveRequest(txRes.IncomingRequest.MessageID, startLog+uint64(i)); err != nil {
					return nil, err
				}
			}

			lastBlock = block
			txdb.callMut.Lock()
			txdb.addSnap(res.BlockNum, res.Timestamp)
			txdb.callMut.Unlock()
		}
	}
	return lastBlock, nil
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
	return txdb.snapCache.latest()
}

func (txdb *TxDB) GetSnapshot(time inbox.ChainTime) *snapshot.Snapshot {
	txdb.callMut.Lock()
	defer txdb.callMut.Unlock()
	return txdb.snapCache.getSnapshot(time)
}

func (txdb *TxDB) LatestBlockId() *common.BlockId {
	txdb.callMut.Lock()
	defer txdb.callMut.Unlock()
	return txdb.lastBlockProcessed
}
