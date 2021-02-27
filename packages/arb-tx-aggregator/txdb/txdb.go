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
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/arbosmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var logger = log.With().Caller().Str("component", "txdb").Logger()

var snapshotCacheSize = 100

type TxDB struct {
	View
	mach              machine.Machine
	checkpointer      checkpointing.RollupCheckpointer
	timeGetter        arbbridge.ChainTimeGetter
	chain             common.Address
	EventCreated      arbbridge.ChainInfo
	CreationTimestamp *big.Int

	rmLogsFeed      event.Feed
	chainFeed       event.Feed
	chainSideFeed   event.Feed
	chainHeadFeed   event.Feed
	logsFeed        event.Feed
	pendingLogsFeed event.Feed
	blockProcFeed   event.Feed

	callMut            sync.Mutex
	lastBlockProcessed *common.BlockId
	lastInboxSeq       *big.Int
	snapCache          *snapshotCache
}

func New(
	clnt arbbridge.ChainTimeGetter,
	checkpointer checkpointing.RollupCheckpointer,
	as machine.AggregatorStore,
	chain common.Address,
	eventCreated arbbridge.ChainInfo,
	creationTimestamp *big.Int,
) (*TxDB, error) {
	if eventCreated.BlockId.Height.AsInt().Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("chain can't be created in 0 block")
	}

	return &TxDB{
		View:              View{as: as},
		checkpointer:      checkpointer,
		timeGetter:        clnt,
		chain:             chain,
		snapCache:         newSnapshotCache(snapshotCacheSize),
		EventCreated:      eventCreated,
		CreationTimestamp: creationTimestamp,
	}, nil
}

func (db *TxDB) Load(ctx context.Context) (bool, error) {
	if db.checkpointer.HasCheckpointedState() {
		err := db.restoreFromCheckpoint(ctx)
		if err == nil {
			return false, nil
		}
		logger.Error().Stack().Err(err).Msg("Failed to restore from checkpoint, falling back to fresh start")
	}
	// We failed to restore from a checkpoint
	logger.Info().Msg("Starting database from scratch")

	if err := db.as.Reorg(0, 0, 0); err != nil {
		return false, err
	}

	valueCache, err := cmachine.NewValueCache()
	if err != nil {
		return false, err
	}

	mach, err := db.checkpointer.GetInitialMachine(valueCache)
	if err != nil {
		return false, err
	}

	initialHeight := new(big.Int).Sub(db.EventCreated.BlockId.Height.AsInt(), big.NewInt(1))

	db.mach = arbosmachine.New(mach)
	db.callMut.Lock()
	db.lastBlockProcessed = nil
	db.lastInboxSeq = big.NewInt(0)
	db.callMut.Unlock()
	return true, db.saveEmptyBlock(ctx, ethcommon.Hash{}, initialHeight)
}

// addSnap must be called with callMut locked or during construction
func (db *TxDB) addSnap(mach machine.Machine, blockNum *big.Int, timestamp *big.Int) {
	currentTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks(new(big.Int).Set(blockNum)),
		Timestamp: new(big.Int).Set(timestamp),
	}
	snap := snapshot.NewSnapshot(mach, currentTime, message.ChainAddressToID(db.chain), new(big.Int).Set(db.lastInboxSeq))
	db.snapCache.addSnapshot(snap)
}

func (db *TxDB) restoreFromCheckpoint(ctx context.Context) error {
	var mach machine.Machine
	var blockId *common.BlockId
	var lastInboxSeq *big.Int
	if err := db.checkpointer.RestoreLatestState(ctx, db.timeGetter, func(chainObserverBytes []byte, restoreCtx ckptcontext.RestoreContext, restoreBlockId *common.BlockId) error {
		var machineHash common.Hash
		copy(machineHash[:], chainObserverBytes)
		lastInboxSeq = new(big.Int).SetBytes(chainObserverBytes[32:])
		rawMach, err := restoreCtx.GetMachine(machineHash)
		if err != nil {
			return err
		}
		mach = arbosmachine.New(rawMach)
		blockId = restoreBlockId
		return nil
	}); err != nil {
		return err
	}

	logger.Debug().
		Uint64("height", blockId.Height.AsInt().Uint64()).
		Hex("machine", mach.Hash().Bytes()).
		Uint64("lastInboxSeq", lastInboxSeq.Uint64()).
		Msg("loaded checkpoint")

	restoreHeight := blockId.Height.AsInt().Uint64()
	// Find the previous block checkout that included an AVM log to find the max
	// avm log and avm send index at restore point
	var blockLog value.Value
	for blockLog == nil {
		blockInfo, err := db.as.GetBlock(restoreHeight)
		if err != nil {
			return err
		}
		if blockInfo == nil {
			return errors.Errorf("no block saved at height %v", restoreHeight)
		}
		blockLog = blockInfo.BlockLog
		restoreHeight--
	}

	block, err := evm.NewBlockResultFromValue(blockLog)
	if err != nil {
		return err
	}

	// Collect all logs that will be removed so they can be sent to rmLogs subscription
	latest, err := db.as.LatestBlock()
	if err == nil {
		oldEthLogs := make([]*types.Log, 0)
		currentHeight := latest.Height.AsInt().Uint64()
		blocksToReorg := currentHeight - restoreHeight
		for i := uint64(0); i < blocksToReorg; i++ {
			height := latest.Height.AsInt().Uint64() - i
			logBlockInfo, err := db.as.GetBlock(height)
			if err != nil {
				return err
			}
			if logBlockInfo == nil {
				// No block at this height so go to the next
				continue
			}

			results, err := db.GetMachineBlockResults(logBlockInfo)
			if err != nil {
				return err
			}

			for i := range results {
				result := results[len(results)-1-i]
				logs := result.EthLogs(common.NewHashFromEth(logBlockInfo.Header.Hash()))
				for j := range logs {
					// Add logs in reverse
					oldEthLogs = append(oldEthLogs, logs[len(logs)-1-j])
				}
			}
		}
		if len(oldEthLogs) > 0 {
			db.rmLogsFeed.Send(core.RemovedLogsEvent{Logs: oldEthLogs})
		}
	}

	if err := db.as.Reorg(
		blockId.Height.AsInt().Uint64(),
		block.ChainStats.AVMSendCount.Uint64(),
		block.ChainStats.AVMLogCount.Uint64(),
	); err != nil {
		return err
	}

	db.mach = mach
	db.callMut.Lock()
	defer db.callMut.Unlock()
	db.lastBlockProcessed = blockId
	db.lastInboxSeq = lastInboxSeq
	db.addSnap(mach.Clone(), block.BlockNum, block.Timestamp)
	return nil
}

func (db *TxDB) AddMessages(ctx context.Context, msgs []inbox.InboxMessage, finishedBlock *common.BlockId) ([]*evm.TxResult, error) {
	var results []*evm.TxResult
	timestamp, err := db.timeGetter.TimestampForBlockHash(ctx, finishedBlock.HeaderHash)
	if err != nil {
		return nil, err
	}

	db.blockProcFeed.Send(true)
	defer db.blockProcFeed.Send(false)

	var lastBlock *evm.BlockInfo
	for _, msg := range msgs {
		assertion, _, _ := db.mach.ExecuteAssertion(1000000000000, true, []inbox.InboxMessage{msg}, true)
		db.callMut.Lock()
		db.lastInboxSeq = msg.InboxSeqNum
		db.callMut.Unlock()
		processedAssertion, err := db.processAssertion(assertion)
		if err != nil {
			return nil, err
		}
		txResults, err := db.saveAssertion(ctx, processedAssertion)
		if err != nil {
			return nil, err
		}
		results = append(results, txResults...)
		if len(processedAssertion.blocks) > 0 {
			block := processedAssertion.blocks[len(processedAssertion.blocks)-1]
			db.callMut.Lock()
			db.addSnap(db.mach.Clone(), block.BlockNum, block.Timestamp)
			db.callMut.Unlock()
			lastBlock = block
		}
	}

	db.callMut.Lock()
	db.lastBlockProcessed = finishedBlock
	lastInboxSeq := new(big.Int).Set(db.lastInboxSeq)

	latestSnap := db.snapCache.latest()
	if latestSnap == nil || latestSnap.Height().Cmp(finishedBlock.Height) < 0 {
		db.addSnap(db.mach.Clone(), finishedBlock.Height.AsInt(), timestamp)
	}
	db.callMut.Unlock()

	if err := db.fillEmptyBlocks(ctx, new(big.Int).Add(finishedBlock.Height.AsInt(), big.NewInt(1))); err != nil {
		return nil, err
	}

	if lastBlock != nil {
		ctx := ckptcontext.NewCheckpointContext()
		ctx.AddMachine(db.mach)
		machHash := db.mach.Hash()
		cpData := make([]byte, 64)
		copy(cpData[:], machHash[:])
		copy(cpData[32:], math.U256Bytes(lastInboxSeq))
		logger.Debug().
			Uint64("height", finishedBlock.Height.AsInt().Uint64()).
			Hex("machine", machHash.Bytes()).
			Uint64("lastInboxSeq", lastInboxSeq.Uint64()).
			Msg("saving checkpoint")
		db.checkpointer.AsyncSaveCheckpoint(finishedBlock, cpData, ctx)
	}
	return results, nil
}

type processedAssertion struct {
	avmLogs   []value.Value
	blocks    []*evm.BlockInfo
	assertion *protocol.ExecutionAssertion
}

func (db *TxDB) processAssertion(assertion *protocol.ExecutionAssertion) (processedAssertion, error) {
	startLogCount, err := db.as.LogCount()
	if err != nil {
		return processedAssertion{}, err
	}
	blocks := make([]*evm.BlockInfo, 0)
	avmLogs := assertion.ParseLogs()
	for i, avmLog := range avmLogs {
		res, err := evm.NewResultFromValue(avmLog)
		if err != nil {
			logger.Error().Stack().Err(err).Msg("Error parsing log result")
			continue
		}

		blockInfo, ok := res.(*evm.BlockInfo)
		if !ok {
			continue
		}

		logger.Debug().
			Uint64("number", blockInfo.BlockNum.Uint64()).
			Uint64("block_logcount", blockInfo.ChainStats.AVMLogCount.Uint64()).
			Uint64("block_messagecount", blockInfo.ChainStats.AVMSendCount.Uint64()).
			Msg("produced l2 block")

		lastLogIndex := blockInfo.ChainStats.AVMLogCount.Uint64() - 1
		if lastLogIndex != startLogCount+uint64(i) {
			return processedAssertion{}, errors.New("unexpected block count")
		}

		blocks = append(blocks, blockInfo)
	}

	return processedAssertion{
		avmLogs:   avmLogs,
		blocks:    blocks,
		assertion: assertion,
	}, nil
}

func (db *TxDB) saveEmptyBlock(ctx context.Context, prev ethcommon.Hash, number *big.Int) error {
	blockId, err := db.timeGetter.BlockIdForHeight(ctx, common.NewTimeBlocks(number))
	if err != nil {
		return err
	}
	time, err := db.timeGetter.TimestampForBlockHash(ctx, blockId.HeaderHash)
	if err != nil {
		return err
	}
	header := &types.Header{
		ParentHash: prev,
		Difficulty: big.NewInt(0),
		Number:     new(big.Int).Set(number),
		GasLimit:   10000000,
		GasUsed:    0,
		Time:       time.Uint64(),
		Extra:      blockId.HeaderHash.Bytes(),
	}
	block := types.NewBlock(header, nil, nil, nil, new(trie.Trie))
	if err := db.as.SaveEmptyBlock(block.Header()); err != nil {
		return err
	}

	if err := db.as.SaveBlockHash(common.NewHashFromEth(block.Hash()), block.NumberU64()); err != nil {
		return err
	}
	db.callMut.Lock()
	defer db.callMut.Unlock()
	db.lastBlockProcessed = blockId
	return nil
}

func (db *TxDB) fillEmptyBlocks(ctx context.Context, max *big.Int) error {
	latest, err := db.as.LatestBlock()
	if err != nil {
		return err
	}
	next := new(big.Int).Add(latest.Height.AsInt(), big.NewInt(1))
	// Fill in empty blocks
	for next.Cmp(max) < 0 {
		prev, err := db.GetBlock(next.Uint64() - 1)
		if err != nil {
			return err
		}
		if prev == nil {
			return errors.Errorf("trying to add block %v, but prev header was not found", next)
		}
		if err := db.saveEmptyBlock(ctx, prev.Header.Hash(), next); err != nil {
			return err
		}
		next = next.Add(next, big.NewInt(1))
	}
	return nil
}

func (db *TxDB) GetBlockResults(res *evm.BlockInfo) ([]*evm.TxResult, error) {
	txCount := res.BlockStats.TxCount.Uint64()
	startLog := res.FirstAVMLog().Uint64()
	results := make([]*evm.TxResult, 0, txCount)
	for i := uint64(0); i < txCount; i++ {
		avmLog, err := db.GetLog(startLog + i)
		if err != nil {
			return nil, err
		}
		res, err := evm.NewTxResultFromValue(avmLog)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func (db *TxDB) GetMachineBlockResults(block *machine.BlockInfo) ([]*evm.TxResult, error) {
	if block.BlockLog == nil {
		// No arb block at this height
		return nil, nil
	}

	res, err := evm.NewBlockResultFromValue(block.BlockLog)
	if err != nil {
		return nil, err
	}
	return db.GetBlockResults(res)
}

func (db *TxDB) GetReceipts(_ context.Context, blockHash ethcommon.Hash) (types.Receipts, error) {
	info, err := db.GetBlockWithHash(common.NewHashFromEth(blockHash))
	if err != nil || info == nil {
		return nil, err
	}

	results, err := db.GetMachineBlockResults(info)
	if err != nil {
		return nil, err
	}
	receipts := make(types.Receipts, 0, len(results))
	for _, res := range results {
		receipts = append(receipts, res.ToEthReceipt(common.NewHashFromEth(blockHash)))
	}
	return receipts, nil
}

func (db *TxDB) GetLogs(_ context.Context, blockHash ethcommon.Hash) ([][]*types.Log, error) {
	info, err := db.GetBlockWithHash(common.NewHashFromEth(blockHash))
	if err != nil || info == nil {
		return nil, err
	}

	results, err := db.GetMachineBlockResults(info)
	if err != nil {
		return nil, err
	}
	logs := make([][]*types.Log, 0, len(results))
	for _, res := range results {
		logs = append(logs, res.EthLogs(common.NewHashFromEth(blockHash)))
	}
	return logs, nil
}

func (db *TxDB) saveAssertion(ctx context.Context, processed processedAssertion) ([]*evm.TxResult, error) {
	for _, avmLog := range processed.avmLogs {
		if err := db.as.SaveLog(avmLog); err != nil {
			return nil, err
		}
	}

	logCount, err := db.as.LogCount()
	if err != nil {
		return nil, err
	}

	for _, avmMessage := range processed.assertion.ParseOutMessages() {
		if err := db.as.SaveMessage(avmMessage); err != nil {
			return nil, err
		}
	}

	messageCount, err := db.as.MessageCount()
	if err != nil {
		return nil, err
	}

	var results []*evm.TxResult
	finalBlockIndex := len(processed.blocks) - 1
	for blockIndex, info := range processed.blocks {
		if err := db.fillEmptyBlocks(ctx, info.BlockNum); err != nil {
			return nil, err
		}

		startLog := info.FirstAVMLog().Uint64()
		// Instead of pulling from DB everytime, should use what we already have
		txResults, err := db.GetBlockResults(info)
		if err != nil {
			return nil, err
		}

		processedResults := evm.FilterEthTxResults(txResults)

		ethTxes := make([]*types.Transaction, 0, len(txResults))
		ethReceipts := make([]*types.Receipt, 0, len(txResults))
		for _, res := range processedResults {
			ethTxes = append(ethTxes, res.Tx)
			ethReceipts = append(ethReceipts, res.Result.ToEthReceipt(common.Hash{}))
			results = append(results, res.Result)

			logger.Debug().
				Hex("hash", res.Result.IncomingRequest.MessageID.Bytes()).
				Int("resulttype", int(res.Result.ResultCode)).
				Msg("got tx result")

			if res.Result.ResultCode == evm.RevertCode {
				logger := logger.Warn().
					Hex("hash", res.Result.IncomingRequest.MessageID.Bytes()).
					Hex("result", res.Result.ReturnData).
					Uint64("gas_used", res.Result.GasUsed.Uint64()).
					Uint64("gas_limit", res.Tx.Gas())
				revertReason, unpackError := abi.UnpackRevert(res.Result.ReturnData)
				if unpackError == nil {
					logger = logger.Str("result_message", revertReason)
				}
				logger.Msg("tx reverted")
			}
		}

		id, err := db.timeGetter.BlockIdForHeight(ctx, common.NewTimeBlocks(info.BlockNum))
		if err != nil {
			return nil, err
		}
		prev, err := db.GetBlock(info.BlockNum.Uint64() - 1)
		if err != nil {
			return nil, err
		}
		if prev == nil {
			return nil, errors.Errorf("trying to add block %v, but prev header was not found", info.BlockNum.Uint64())
		}
		header := &types.Header{
			ParentHash: prev.Header.Hash(),
			Difficulty: big.NewInt(0),
			Number:     new(big.Int).Set(info.BlockNum),
			GasLimit:   info.GasLimit.Uint64(),
			GasUsed:    info.BlockStats.GasUsed.Uint64(),
			Time:       info.Timestamp.Uint64(),
			Extra:      id.HeaderHash.Bytes(),
		}

		block := types.NewBlock(header, ethTxes, nil, ethReceipts, new(trie.Trie))
		avmLogIndex := info.ChainStats.AVMLogCount.Uint64() - 1
		logger.Debug().
			Uint64("number", block.Header().Number.Uint64()).
			Uint64("block_logcount", info.ChainStats.AVMLogCount.Uint64()).
			Uint64("block_messagecount", info.ChainStats.AVMSendCount.Uint64()).
			Uint64("logsize", logCount).
			Uint64("messagesize", messageCount).
			Msg("saved l2 block")
		if err := db.as.SaveBlock(block.Header(), avmLogIndex); err != nil {
			return nil, err
		}

		ethLogs := make([]*types.Log, 0)
		for _, res := range processedResults {
			ethLogs = append(ethLogs, res.Result.EthLogs(common.NewHashFromEth(block.Hash()))...)
		}

		for i, txRes := range txResults {
			if txRes.ResultCode != evm.ReturnCode && txRes.ResultCode != evm.RevertCode {
				// If this log was for an invalid transaction, only save the request if it hasn't been saved before
				if db.as.GetPossibleRequestInfo(txRes.IncomingRequest.MessageID) != nil {
					continue
				}
			}

			if err := db.as.SaveRequest(txRes.IncomingRequest.MessageID, startLog+uint64(i)); err != nil {
				return nil, err
			}
		}

		if err := db.as.SaveBlockHash(common.NewHashFromEth(block.Hash()), block.Number().Uint64()); err != nil {
			return nil, err
		}

		db.chainFeed.Send(core.ChainEvent{Block: block, Hash: block.Hash(), Logs: ethLogs})
		if finalBlockIndex == blockIndex {
			db.chainHeadFeed.Send(core.ChainEvent{Block: block, Hash: block.Hash(), Logs: ethLogs})
		}
		if len(ethLogs) > 0 {
			db.logsFeed.Send(ethLogs)
		}
	}
	return results, nil
}

func (db *TxDB) GetMessage(index uint64) (value.Value, error) {
	return db.as.GetMessage(index)
}

func (db *TxDB) GetLog(index uint64) (value.Value, error) {
	return db.as.GetLog(index)
}

func (db *TxDB) GetBlock(height uint64) (*machine.BlockInfo, error) {
	latest := db.LatestBlock()
	if height > latest.Height.AsInt().Uint64() {
		return nil, nil
	}
	return db.as.GetBlock(height)
}

func (db *TxDB) LatestBlock() *common.BlockId {
	block, err := db.as.LatestBlock()
	if err != nil {
		return db.lastBlockProcessed
	}
	return block
}

func (db *TxDB) LatestSnapshot() *snapshot.Snapshot {
	db.callMut.Lock()
	defer db.callMut.Unlock()
	return db.snapCache.latest()
}

func (db *TxDB) GetSnapshot(time inbox.ChainTime) *snapshot.Snapshot {
	db.callMut.Lock()
	defer db.callMut.Unlock()
	return db.snapCache.getSnapshot(time)
}

func (db *TxDB) LatestBlockId() *common.BlockId {
	db.callMut.Lock()
	defer db.callMut.Unlock()
	return db.lastBlockProcessed
}

func (db *TxDB) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	return db.chainFeed.Subscribe(ch)
}

func (db *TxDB) SubscribeChainHeadEvent(ch chan<- core.ChainEvent) event.Subscription {
	return db.chainHeadFeed.Subscribe(ch)
}

func (db *TxDB) SubscribeChainSideEvent(ch chan<- core.ChainEvent) event.Subscription {
	return db.chainSideFeed.Subscribe(ch)
}

func (db *TxDB) SubscribeRemovedLogsEvent(ch chan<- core.RemovedLogsEvent) event.Subscription {
	return db.rmLogsFeed.Subscribe(ch)
}

func (db *TxDB) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return db.logsFeed.Subscribe(ch)
}

func (db *TxDB) SubscribePendingLogsEvent(ch chan<- []*types.Log) event.Subscription {
	return db.pendingLogsFeed.Subscribe(ch)
}

func (db *TxDB) SubscribeBlockProcessingEvent(ch chan<- []*types.Log) event.Subscription {
	return db.blockProcFeed.Subscribe(ch)
}
