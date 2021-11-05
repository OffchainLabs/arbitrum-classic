/*
* Copyright 2020-2021, Offchain Labs, Inc.
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
	"fmt"
	"math/big"
	"strings"

	lru "github.com/hashicorp/golang-lru"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/blockcache"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/monitor"
)

var logger = log.With().Caller().Stack().Str("component", "txdb").Logger()

type TxDB struct {
	Lookup          core.ArbCoreLookup
	allowSlowLookup bool
	as              machine.NodeStore
	logReader       *core.LogReader

	newTxsFeed      event.Feed
	rmLogsFeed      event.Feed
	chainFeed       event.Feed
	chainSideFeed   event.Feed
	chainHeadFeed   event.Feed
	logsFeed        event.Feed
	pendingLogsFeed event.Feed
	blockProcFeed   event.Feed

	snapshotLRUCache   *lru.Cache
	blockInfoLRUCache  *lru.Cache
	snapshotTimedCache *blockcache.BlockCache
}

func New(
	ctx context.Context,
	arbCore core.ArbCore,
	as machine.NodeStore,
	nodeConfig *configuration.Node,
) (*TxDB, <-chan error, error) {
	var snapshotLRUCache *lru.Cache
	var blockInfoLRUCache *lru.Cache
	if nodeConfig.Cache.LRUSize > 0 {
		var err error
		snapshotLRUCache, err = lru.New(nodeConfig.Cache.LRUSize)
		if err != nil {
			return nil, nil, err
		}
	}
	if nodeConfig.Cache.BlockInfoLRUSize > 0 {
		var err error
		blockInfoLRUCache, err = lru.New(nodeConfig.Cache.BlockInfoLRUSize)
		if err != nil {
			return nil, nil, err
		}
	}
	snapshotTimedCache, err := blockcache.New(nodeConfig.Cache.TimedInitialSize, nodeConfig.Cache.TimedExpire)
	if err != nil {
		return nil, nil, err
	}
	db := &TxDB{
		Lookup:             arbCore,
		as:                 as,
		snapshotLRUCache:   snapshotLRUCache,
		blockInfoLRUCache:  blockInfoLRUCache,
		snapshotTimedCache: snapshotTimedCache,
		allowSlowLookup:    nodeConfig.Cache.AllowSlowLookup,
	}
	logReader := core.NewLogReader(db, arbCore, big.NewInt(0), big.NewInt(int64(nodeConfig.LogProcessCount)), nodeConfig.LogIdleSleep)
	errChan := logReader.Start(ctx)
	db.logReader = logReader
	return db, errChan, nil
}

func (db *TxDB) Close() {
	db.logReader.Stop()
}

func (db *TxDB) GetBlockResults(block *machine.BlockInfo) (*evm.BlockInfo, []*evm.TxResult, error) {
	startLog := new(big.Int).SetUint64(block.InitialLogIndex())
	logCount := new(big.Int).SetUint64(block.LogCount + 1)

	avmLogs, err := db.Lookup.GetLogs(startLog, logCount)
	if err != nil {
		return nil, nil, err
	}
	if uint64(len(avmLogs)) != block.LogCount+1 {
		logger.Warn().Msg("reorged getting block results")
		return nil, nil, nil
	}
	l2Block, err := evm.NewBlockResultFromValue(avmLogs[len(avmLogs)-1].Value)
	if err != nil {
		logger.Warn().Msg("reorged getting block results")
		return nil, nil, nil
	}
	if l2Block.BlockStats.AVMLogCount.Cmp(new(big.Int).SetUint64(block.LogCount)) != 0 ||
		l2Block.BlockNum.Cmp(block.Header.Number) != 0 {
		fmt.Println(l2Block.BlockStats.AVMLogCount, block.LogCount)
		fmt.Println(l2Block.BlockNum, block.Header.Number)
		logger.Warn().Msg("reorged getting block results")
		return nil, nil, nil
	}
	txResults, err := processBlockResults(l2Block, avmLogs[:l2Block.BlockStats.TxCount.Uint64()])
	return l2Block, txResults, err
}

func (db *TxDB) getBlockResultsUnsafe(res *evm.BlockInfo) ([]*evm.TxResult, error) {
	avmLogs, err := db.Lookup.GetLogs(res.FirstAVMLog(), res.BlockStats.TxCount)
	if err != nil {
		return nil, err
	}
	return processBlockResults(res, avmLogs)
}

func processBlockResults(block *evm.BlockInfo, avmLogs []core.ValueAndInbox) ([]*evm.TxResult, error) {
	results := make([]*evm.TxResult, 0, len(avmLogs))
	for _, avmLog := range avmLogs {
		res, err := evm.NewResultFromValue(avmLog.Value)
		if err != nil {
			return nil, err
		}
		txRes, ok := res.(*evm.TxResult)
		if !ok {
			return nil, errors.Errorf("expected tx result but got %T", res)
		}
		if txRes.ResultCode != evm.RevertCode && txRes.IncomingRequest.L2BlockNumber.Cmp(block.BlockNum) != 0 {
			return nil, errors.New("tx from wrong block")
		}
		results = append(results, txRes)
	}
	return results, nil
}

func (db *TxDB) AddLogs(initialLogIndex *big.Int, avmLogs []core.ValueAndInbox) error {
	logger.Info().Str("start", initialLogIndex.String()).Int("count", len(avmLogs)).Msg("adding logs")
	logIndex := initialLogIndex.Uint64()
	for _, avmLog := range avmLogs {
		if err := db.HandleLog(logIndex, avmLog); err != nil {
			return err
		}
		logIndex++
	}
	return nil
}

func (db *TxDB) DeleteLogs(avmLogs []core.ValueAndInbox) error {
	logger.Info().Int("count", len(avmLogs)).Msg("deleting logs")
	oldHeight, err := db.BlockCount()
	if err != nil {
		return err
	}
	// Collect all logs that will be removed so they can be sent to rmLogs subscription
	var reorgBlockHeight uint64
	blockReceiptFound := false
	for _, avmLog := range avmLogs {
		// L2 transaction receipts already provided in reverse
		res, err := evm.NewResultFromValue(avmLog.Value)
		if err != nil {
			return err
		}
		txRes, ok := res.(*evm.TxResult)
		if !ok {
			blockRes, ok := res.(*evm.BlockInfo)
			if ok {
				blockReceiptFound = true
				reorgBlockHeight = blockRes.BlockNum.Uint64()
			}
			continue
		}

		currentBlockHeight := txRes.IncomingRequest.L2BlockNumber.Uint64()
		logBlockInfo, err := db.GetBlock(currentBlockHeight)
		if err != nil {
			return err
		}
		if logBlockInfo == nil {
			continue
		}
		logs := txRes.EthLogs(common.NewHashFromEth(logBlockInfo.Header.Hash()))
		lastLogIndex := len(logs) - 1
		oldEthLogs := make([]*types.Log, 0)
		for j := range logs {
			// Add logs in reverse
			oldEthLogs = append(oldEthLogs, logs[lastLogIndex-j])
		}

		if len(oldEthLogs) > 0 {
			db.rmLogsFeed.Send(ethcore.RemovedLogsEvent{Logs: oldEthLogs})
		}
	}

	if blockReceiptFound {
		// Reset block height
		err := db.as.Reorg(reorgBlockHeight)
		if err != nil {
			return err
		}

		if db.snapshotLRUCache != nil {
			for i := oldHeight; i > reorgBlockHeight; i-- {
				db.snapshotLRUCache.Remove(i)
			}
			db.snapshotLRUCache.Remove(reorgBlockHeight)
		}
		if db.blockInfoLRUCache != nil {
			for i := oldHeight; i > reorgBlockHeight; i-- {
				db.blockInfoLRUCache.Remove(i)
			}
			db.blockInfoLRUCache.Remove(reorgBlockHeight)
		}
		db.snapshotTimedCache.Reorg(reorgBlockHeight)
	}

	return nil
}

func (db *TxDB) HandleLog(logIndex uint64, avmLog core.ValueAndInbox) error {
	res, err := evm.NewResultFromValue(avmLog.Value)
	if err != nil {
		logger.Error().Err(err).Msg("Error parsing log result")
		return nil
	}

	switch res := res.(type) {
	case *evm.BlockInfo:
		return db.handleBlockReceipt(res)
	case *evm.MerkleRootResult:
		return db.as.SaveMessageBatch(res.BatchNumber, logIndex)
	case *evm.TxResult:
		monitor.GlobalMonitor.GotLog(res.IncomingRequest.MessageID)
		tx, err := evm.GetTransaction(res)
		if err != nil {
			logger.Warn().Err(err).Msg("error pulling transaction from receipt")
		} else {
			db.newTxsFeed.Send(ethcore.NewTxsEvent{Txs: []*types.Transaction{tx.Tx}})
		}
	}
	return nil
}

func (db *TxDB) handleBlockReceipt(blockInfo *evm.BlockInfo) error {
	logger.Debug().
		Uint64("number", blockInfo.BlockNum.Uint64()).
		Uint64("block_txcount", blockInfo.BlockStats.TxCount.Uint64()).
		Uint64("block_logcount", blockInfo.BlockStats.AVMLogCount.Uint64()).
		Uint64("block_sendcount", blockInfo.BlockStats.AVMSendCount.Uint64()).
		Msg("produced l2 block")

	txResults, err := db.getBlockResultsUnsafe(blockInfo)
	if err != nil {
		return err
	}

	if uint64(len(txResults)) != blockInfo.BlockStats.TxCount.Uint64() {
		logger.Warn().
			Uint64("block", blockInfo.BlockNum.Uint64()).
			Int("real", len(txResults)).
			Uint64("claimed", blockInfo.BlockStats.TxCount.Uint64()).
			Msg("expected to get same number of results")
	}
	if blockInfo.BlockStats.AVMLogCount.Cmp(big.NewInt(0)) == 0 {
		logger.Warn().
			Uint64("block", blockInfo.BlockNum.Uint64()).
			Msg("found empty block")
	}

	processedResults := evm.FilterEthTxResults(txResults)

	var results []*evm.TxResult
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

	prevHash := ethcommon.Hash{}
	if blockInfo.BlockNum.Cmp(big.NewInt(0)) > 0 {
		prev, err := db.GetBlock(blockInfo.BlockNum.Uint64() - 1)
		if err != nil {
			return err
		}
		if prev == nil {
			return errors.Errorf("trying to add block %v, but prev header was not found", blockInfo.BlockNum.Uint64())
		}
		prevHash = prev.Header.Hash()
	}

	header := &types.Header{
		ParentHash: prevHash,
		Difficulty: big.NewInt(0),
		Number:     new(big.Int).Set(blockInfo.BlockNum),
		GasLimit:   blockInfo.GasLimit().Uint64(),
		GasUsed:    blockInfo.BlockStats.GasUsed.Uint64(),
		Time:       blockInfo.Timestamp.Uint64(),
		Extra:      nil,
	}

	block := types.NewBlock(header, ethTxes, nil, ethReceipts, new(trie.Trie))
	avmLogIndex := blockInfo.ChainStats.AVMLogCount.Uint64() - 1
	ethLogs := make([]*types.Log, 0)
	for _, res := range processedResults {
		ethLogs = append(ethLogs, res.Result.EthLogs(common.NewHashFromEth(block.Hash()))...)
	}

	requests := make([]machine.EVMRequestInfo, 0, len(txResults))

	for i, txRes := range txResults {
		// && txRes.ResultCode != evm.RevertCode
		if txRes.ResultCode != evm.ReturnCode {
			// If this log was for an invalid transaction, only save the request if it hasn't been saved before
			orig, _, _, err := db.GetRequest(txRes.IncomingRequest.MessageID)
			if err != nil {
				return err
			}
			if orig != nil {
				continue
			}
		}

		requests = append(requests, machine.EVMRequestInfo{
			RequestId: txRes.IncomingRequest.MessageID,
			LogIndex:  blockInfo.FirstAVMLog().Uint64() + uint64(i),
		})
	}

	arbBlockInfo := &machine.BlockInfo{
		Header:   block.Header(),
		BlockLog: avmLogIndex,
		LogCount: blockInfo.BlockStats.AVMLogCount.Uint64(),
	}
	if err := db.as.SaveBlock(arbBlockInfo, requests); err != nil {
		return err
	}
	if db.blockInfoLRUCache != nil {
		db.blockInfoLRUCache.Add(header.Number.Uint64(), arbBlockInfo)
	}

	db.chainFeed.Send(ethcore.ChainEvent{Block: block, Hash: block.Hash(), Logs: ethLogs})
	db.chainHeadFeed.Send(ethcore.ChainEvent{Block: block, Hash: block.Hash(), Logs: ethLogs})
	if len(ethLogs) > 0 {
		db.logsFeed.Send(ethLogs)
	}
	return nil
}

func (db *TxDB) GetMessageBatch(index *big.Int) (*evm.MerkleRootResult, error) {
	logIndex := db.as.GetMessageBatch(index)
	if logIndex == nil {
		return nil, nil
	}
	logVal, err := core.GetZeroOrOneLog(db.Lookup, new(big.Int).SetUint64(*logIndex))
	if err != nil || logVal.Value == nil {
		return nil, err
	}
	res, err := evm.NewResultFromValue(logVal.Value)
	if err != nil {
		return nil, err
	}
	merkleRes, ok := res.(*evm.MerkleRootResult)
	if !ok {
		return nil, errors.Errorf("expected merkle root result but got %T at log index %v", res, *logIndex)
	}
	if merkleRes.BatchNumber.Cmp(index) != 0 {
		return nil, nil
	}
	return merkleRes, nil
}

func (db *TxDB) GetBlockWithHash(blockHash common.Hash) (*machine.BlockInfo, error) {
	blockHeight := db.as.GetPossibleBlock(blockHash)
	if blockHeight == nil {
		return nil, nil
	}
	info, err := db.as.GetBlockInfo(*blockHeight)
	if err != nil {
		return nil, err
	}
	if info.Header.Hash() != blockHash.ToEthHash() {
		return nil, nil
	}
	return info, err
}

func (db *TxDB) GetRequest(requestId common.Hash) (*evm.TxResult, core.InboxState, *big.Int, error) {
	requestCandidate := db.as.GetPossibleRequestInfo(requestId)
	if requestCandidate == nil {
		return nil, core.InboxState{}, nil, nil
	}
	logNumber := new(big.Int).SetUint64(*requestCandidate)
	logVal, err := core.GetZeroOrOneLog(db.Lookup, logNumber)
	if err != nil || logVal.Value == nil {
		return nil, core.InboxState{}, nil, err
	}
	res, err := evm.NewResultFromValue(logVal.Value)
	if err != nil {
		return nil, core.InboxState{}, nil, err
	}
	txRes, ok := res.(*evm.TxResult)
	if !ok {
		return nil, core.InboxState{}, nil, nil
	}
	if txRes.IncomingRequest.MessageID != requestId {
		return nil, core.InboxState{}, nil, nil
	}
	return txRes, logVal.Inbox, logNumber, nil
}

func (db *TxDB) GetL2Block(block *machine.BlockInfo) (*evm.BlockInfo, error) {
	blockLog, err := core.GetZeroOrOneLog(db.Lookup, new(big.Int).SetUint64(block.BlockLog))
	if err != nil || blockLog.Value == nil {
		return nil, err
	}
	return evm.NewBlockResultFromValue(blockLog.Value)
}

func (db *TxDB) GetBlock(height uint64) (*machine.BlockInfo, error) {
	if db.blockInfoLRUCache != nil {
		cachedInfo, ok := db.blockInfoLRUCache.Get(height)
		if ok {
			return cachedInfo.(*machine.BlockInfo), nil
		}
	}
	count, err := db.BlockCount()
	if err != nil {
		return nil, err
	}
	if height >= count {
		return nil, nil
	}
	info, err := db.as.GetBlockInfo(height)
	if err == nil && db.blockInfoLRUCache != nil {
		db.blockInfoLRUCache.Add(info.Header.Number.Uint64(), info)
	}
	return info, err
}

func (db *TxDB) BlockCount() (uint64, error) {
	return db.as.BlockCount()
}

func (db *TxDB) LatestBlock() (*machine.BlockInfo, error) {
	blockCount, err := db.as.BlockCount()
	if err != nil {
		return nil, err
	}
	totalLogCountBig, err := db.Lookup.GetLogCount()
	if err != nil {
		return nil, err
	}
	totalLogCount := totalLogCountBig.Uint64()
	for blockCount > 0 {
		blockData, err := db.as.GetBlockInfo(blockCount - 1)
		if err != nil {
			return nil, err
		}
		if blockData.BlockLog < totalLogCount {
			return blockData, nil
		}
		blockCount--
	}
	return nil, errors.New("can't get latest block because there are no blocks")
}

func (db *TxDB) getSnapshotForInfo(info *machine.BlockInfo) (*snapshot.Snapshot, error) {
	if db.snapshotLRUCache != nil {
		cachedSnap, found := db.snapshotLRUCache.Get(info.Header.Number.Uint64())
		if found {
			return cachedSnap.(*snapshot.Snapshot), nil
		}
	}
	_, cachedSnap := db.snapshotTimedCache.Get(info.Header.Number.Uint64())
	if cachedSnap != nil {
		return cachedSnap, nil
	}
	cursor, err := db.Lookup.GetExecutionCursorAtEndOfBlock(info.Header.Number.Uint64(), db.allowSlowLookup)
	if err != nil {
		return nil, err
	}
	mach, err := db.Lookup.TakeMachine(cursor)
	if err != nil {
		return nil, err
	}

	currentTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks(new(big.Int).Set(info.Header.Number)),
		Timestamp: new(big.Int).SetUint64(info.Header.Time),
	}
	snap, err := snapshot.NewSnapshot(mach, currentTime, big.NewInt(1<<60))
	if err != nil {
		return nil, err
	}
	if db.snapshotLRUCache != nil {
		db.snapshotLRUCache.Add(info.Header.Number.Uint64(), snap)
	}
	db.snapshotTimedCache.Add(info.Header, snap)
	return snap, nil
}

func (db *TxDB) GetSnapshot(blockHeight uint64) (*snapshot.Snapshot, error) {
	info, err := db.GetBlock(blockHeight)
	if err != nil || info == nil {
		return nil, err
	}
	return db.getSnapshotForInfo(info)
}

func (db *TxDB) LatestSnapshot() (*snapshot.Snapshot, error) {
	block, err := db.LatestBlock()
	if err != nil {
		return nil, err
	}
	snap, err := db.getSnapshotForInfo(block)
	if err != nil {
		if strings.Contains(err.Error(), "block not in cache") {
			logger.Error().Hex("block", block.Header.Number.Bytes()).Msg("latest block is not in cache")
		}

		return nil, err
	}

	return snap, nil
}

func (db *TxDB) SubscribeNewTxsEvent(ch chan<- ethcore.NewTxsEvent) event.Subscription {
	return db.newTxsFeed.Subscribe(ch)
}

func (db *TxDB) SubscribeChainEvent(ch chan<- ethcore.ChainEvent) event.Subscription {
	return db.chainFeed.Subscribe(ch)
}

func (db *TxDB) SubscribeChainHeadEvent(ch chan<- ethcore.ChainEvent) event.Subscription {
	return db.chainHeadFeed.Subscribe(ch)
}

func (db *TxDB) SubscribeChainSideEvent(ch chan<- ethcore.ChainEvent) event.Subscription {
	return db.chainSideFeed.Subscribe(ch)
}

func (db *TxDB) SubscribeRemovedLogsEvent(ch chan<- ethcore.RemovedLogsEvent) event.Subscription {
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
