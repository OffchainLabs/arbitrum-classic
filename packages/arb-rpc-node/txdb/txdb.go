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

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcore "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

var logger = log.With().Caller().Str("component", "txdb").Logger()

type AggregatorStore interface {
	GetPossibleRequestInfo(requestId common.Hash) *uint64
	GetPossibleBlock(blockHash common.Hash) *uint64
	GetBlockHeader(height uint64) (*machine.BlockInfo, error)
	EarliestBlock() (*common.BlockId, error)
	LatestBlock() (*common.BlockId, error)

	SaveBlock(header *types.Header, logIndex uint64) error
	SaveEmptyBlock(header *types.Header) error
	SaveBlockHash(blockHash common.Hash, blockHeight uint64) error
	SaveRequest(requestId common.Hash, logIndex uint64) error
	Reorg(height uint64) error
}

type ChainTimeGetter interface {
	BlockIdForHeight(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error)
	TimestampForBlockHash(ctx context.Context, hash common.Hash) (*big.Int, error)
}

type TxDB struct {
	lookup     core.ArbCoreLookup
	as         AggregatorStore
	timeGetter ChainTimeGetter
	chain      common.Address

	rmLogsFeed      event.Feed
	chainFeed       event.Feed
	chainSideFeed   event.Feed
	chainHeadFeed   event.Feed
	logsFeed        event.Feed
	pendingLogsFeed event.Feed
	blockProcFeed   event.Feed

	callMut sync.Mutex
}

func (db *TxDB) GetBlockResults(res *evm.BlockInfo) ([]*evm.TxResult, error) {
	avmLogs, err := db.lookup.GetLogs(res.FirstAVMLog(), res.BlockStats.AVMLogCount)
	if err != nil {
		return nil, err
	}
	results := make([]*evm.TxResult, 0, len(avmLogs))
	for _, avmLog := range avmLogs {
		res, err := evm.NewTxResultFromValue(avmLog)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}

func (db *TxDB) CurrentLogCount() (*big.Int, error) {
	//return db.as.CurrentLogCount()
	// TODO
	return big.NewInt(0), nil
}

func (db *TxDB) UpdateCurrentLogCount() (*big.Int, error) {
	//return db.as.UpdateCurrentLogCount()
	// TODO
	return big.NewInt(0), nil
}

func (db *TxDB) AddLogs(avmLogs []value.Value) error {
	ctx := context.Background()

	for _, avmLog := range avmLogs {
		if err := db.HandleLog(ctx, avmLog); err != nil {
			return err
		}
	}

	return nil
}

func (db *TxDB) DeleteLogs(avmLogs []value.Value) error {
	// Collect all logs that will be removed so they can be sent to rmLogs subscription
	lastResultIndex := len(avmLogs) - 1
	var currentBlockHeight uint64
	blocksFound := false
	for i, _ := range avmLogs {
		// Parse L2 transaction receipts in reverse
		res, err := evm.NewResultFromValue(avmLogs[lastResultIndex-i])
		if err != nil {
			return err
		}
		txRes, ok := res.(*evm.TxResult)
		if !ok {
			continue
		}

		blocksFound = true

		currentBlockHeight = txRes.IncomingRequest.ChainTime.BlockNum.AsInt().Uint64()
		logBlockInfo, err := db.GetBlock(currentBlockHeight)
		if err != nil {
			return err
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
	if !blocksFound {
		return nil
	}

	if currentBlockHeight > 0 {
		currentBlockHeight--
	}

	// Reset block height
	err := db.as.Reorg(currentBlockHeight)
	if err != nil {
		return err
	}

	return nil
}

func (db *TxDB) HandleLog(ctx context.Context, avmLog value.Value) error {
	res, err := evm.NewResultFromValue(avmLog)
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Error parsing log result")
		return nil
	}
	blockInfo, ok := res.(*evm.BlockInfo)
	if !ok {
		return nil
	}

	logger.Debug().
		Uint64("number", blockInfo.BlockNum.Uint64()).
		Uint64("block_logcount", blockInfo.ChainStats.AVMLogCount.Uint64()).
		Uint64("block_sendcount", blockInfo.ChainStats.AVMSendCount.Uint64()).
		Msg("produced l2 block")

	if err := db.fillEmptyBlocks(ctx, blockInfo.BlockNum); err != nil {
		return err
	}

	txResults, err := db.GetBlockResults(blockInfo)
	if err != nil {
		return err
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

	id, err := db.timeGetter.BlockIdForHeight(ctx, common.NewTimeBlocks(blockInfo.BlockNum))
	if err != nil {
		return err
	}
	prev, err := db.GetBlock(blockInfo.BlockNum.Uint64() - 1)
	if err != nil {
		return err
	}
	if prev == nil {
		return errors.Errorf("trying to add block %v, but prev header was not found", blockInfo.BlockNum.Uint64())
	}
	header := &types.Header{
		ParentHash: prev.Header.Hash(),
		Difficulty: big.NewInt(0),
		Number:     new(big.Int).Set(blockInfo.BlockNum),
		GasLimit:   blockInfo.GasLimit().Uint64(),
		GasUsed:    blockInfo.BlockStats.GasUsed.Uint64(),
		Time:       blockInfo.Timestamp.Uint64(),
		Extra:      id.HeaderHash.Bytes(),
	}

	block := types.NewBlock(header, ethTxes, nil, ethReceipts, new(trie.Trie))
	avmLogIndex := blockInfo.ChainStats.AVMLogCount.Uint64() - 1
	logger.Debug().
		Uint64("number", block.Header().Number.Uint64()).
		Uint64("block_logcount", blockInfo.ChainStats.AVMLogCount.Uint64()).
		Uint64("block_messagecount", blockInfo.ChainStats.AVMSendCount.Uint64()).
		Msg("saved l2 block")
	if err := db.as.SaveBlock(block.Header(), avmLogIndex); err != nil {
		return err
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

		if err := db.as.SaveRequest(txRes.IncomingRequest.MessageID, blockInfo.FirstAVMLog().Uint64()+uint64(i)); err != nil {
			return err
		}
	}

	if err := db.as.SaveBlockHash(common.NewHashFromEth(block.Hash()), block.Number().Uint64()); err != nil {
		return err
	}

	db.chainFeed.Send(ethcore.ChainEvent{Block: block, Hash: block.Hash(), Logs: ethLogs})
	db.chainHeadFeed.Send(ethcore.ChainEvent{Block: block, Hash: block.Hash(), Logs: ethLogs})
	if len(ethLogs) > 0 {
		db.logsFeed.Send(ethLogs)
	}
	return nil
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
	return nil
}

func (db *TxDB) AddInitialBlock(ctx context.Context, initialBlockHeight *big.Int) error {
	return db.saveEmptyBlock(ctx, ethcommon.Hash{}, initialBlockHeight)
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

func (db *TxDB) GetBlockWithHash(blockHash common.Hash) (*machine.BlockInfo, error) {
	blockHeight := db.as.GetPossibleBlock(blockHash)
	if blockHeight == nil {
		return nil, nil
	}
	info, err := db.as.GetBlockHeader(*blockHeight)
	if err != nil {
		return nil, err
	}
	if info.Header.Hash() != blockHash.ToEthHash() {
		return nil, nil
	}
	return info, err
}

func (db *TxDB) GetRequest(requestId common.Hash) (value.Value, error) {
	requestCandidate := db.as.GetPossibleRequestInfo(requestId)
	if requestCandidate == nil {
		return nil, nil
	}
	logVals, err := db.lookup.GetLogs(new(big.Int).SetUint64(*requestCandidate), big.NewInt(1))
	if err != nil {
		return nil, err
	}
	if len(logVals) != 1 {
		return nil, errors.New("unexpected log count")
	}
	logVal := logVals[0]
	res, err := evm.NewTxResultFromValue(logVal)
	if err != nil {
		return nil, err
	}
	if res.IncomingRequest.MessageID != requestId {
		return nil, nil
	}
	return logVal, nil
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

func (db *TxDB) GetBlock(height uint64) (*machine.BlockInfo, error) {
	latest, err := db.LatestBlock()
	if err != nil {
		return nil, err
	}
	if height > latest.Height.AsInt().Uint64() {
		return nil, nil
	}
	return db.as.GetBlockHeader(height)
}

func (db *TxDB) EarliestBlock() (*common.BlockId, error) {
	return db.as.EarliestBlock()
}

func (db *TxDB) LatestBlock() (*common.BlockId, error) {
	return db.as.LatestBlock()
}

func (db *TxDB) getSnapshotForInfo(info *machine.BlockInfo) (*snapshot.Snapshot, error) {
	mach, err := db.lookup.GetMachineForSideload(info.Header.Number.Uint64())
	if err != nil || mach == nil {
		return nil, err
	}
	currentTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocks(new(big.Int).Set(info.Header.Number)),
		Timestamp: new(big.Int).SetUint64(info.Header.Time),
	}
	snap := snapshot.NewSnapshot(mach, currentTime, message.ChainAddressToID(db.chain), big.NewInt(1<<60))
	return snap, nil
}

func (db *TxDB) GetSnapshot(blockHeight uint64) (*snapshot.Snapshot, error) {
	info, err := db.GetBlock(blockHeight)
	if err != nil || info == nil {
		return nil, err
	}
	return db.getSnapshotForInfo(info)
}

func (db *TxDB) LatestSnapshot() *snapshot.Snapshot {
	block, err := db.LatestBlock()
	if err != nil || block == nil {
		return nil
	}
	snap, err := db.GetSnapshot(block.Height.AsInt().Uint64())
	if err != nil {
		return nil
	}
	return snap
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
