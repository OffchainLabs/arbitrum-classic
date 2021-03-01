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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

type ChainTimeGetter interface {
	BlockIdForHeight(ctx context.Context, height *common.TimeBlocks) (*common.BlockId, error)
	TimestampForBlockHash(ctx context.Context, hash common.Hash) (*big.Int, error)
}

type TxDB struct {
	lookup core.ArbOutputLookup
	as     machine.AggregatorStore
	chain  common.Address

	rmLogsFeed      event.Feed
	chainFeed       event.Feed
	chainSideFeed   event.Feed
	chainHeadFeed   event.Feed
	logsFeed        event.Feed
	pendingLogsFeed event.Feed
	blockProcFeed   event.Feed

	callMut sync.Mutex
}

func New(
	core core.ArbOutputLookup,
	as machine.AggregatorStore,
	chain common.Address,
) (*TxDB, error) {
	return &TxDB{
		lookup: core,
		as:     as,
		chain:  chain,
	}, nil
}

func (db *TxDB) GetBlockResults(res *evm.BlockInfo) ([]*evm.TxResult, error) {
	avmLogs, err := db.lookup.GetLogs(res.FirstAVMLog(), res.BlockStats.TxCount)
	if err != nil {
		return nil, err
	}
	fmt.Println("GetBlockResults", res.BlockNum, res.FirstAVMLog(), res.BlockStats.AVMLogCount, res.BlockStats.TxCount, len(avmLogs))
	results := make([]*evm.TxResult, 0, len(avmLogs))
	for _, avmLog := range avmLogs {
		res, err := evm.NewResultFromValue(avmLog)
		if err != nil {
			return nil, err
		}
		txRes, ok := res.(*evm.TxResult)
		if !ok {
			continue
		}
		results = append(results, txRes)
	}
	return results, nil
}

func (db *TxDB) CurrentLogCount() (*big.Int, error) {
	return db.as.CurrentLogCount()
}

func (db *TxDB) UpdateCurrentLogCount(count *big.Int) error {
	return db.as.UpdateCurrentLogCount(count)
}

func (db *TxDB) AddLogs(avmLogs []value.Value) error {
	for _, avmLog := range avmLogs {
		if err := db.HandleLog(avmLog); err != nil {
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

		currentBlockHeight = txRes.IncomingRequest.L2BlockNumber.Uint64()
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

func (db *TxDB) HandleLog(avmLog value.Value) error {
	res, err := evm.NewResultFromValue(avmLog)
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Error parsing log result")
		return nil
	}
	blockInfo, ok := res.(*evm.BlockInfo)
	if !ok {
		return nil
	}

	totalLogCount, err := db.lookup.GetLogCount()
	if err != nil {
		return err
	}
	fmt.Println("Total log count", totalLogCount)

	logger.Debug().
		Uint64("number", blockInfo.BlockNum.Uint64()).
		Uint64("block_txcount", blockInfo.BlockStats.TxCount.Uint64()).
		Uint64("block_logcount", blockInfo.BlockStats.AVMLogCount.Uint64()).
		Uint64("block_sendcount", blockInfo.BlockStats.AVMSendCount.Uint64()).
		Msg("produced l2 block")

	txResults, err := db.GetBlockResults(blockInfo)
	if err != nil {
		return err
	}
	fmt.Println("Block results for", blockInfo.BlockNum)
	for _, res := range txResults {
		fmt.Println("Got res", res.IncomingRequest.MessageID)
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
		if txRes.ResultCode != evm.ReturnCode && txRes.ResultCode != evm.RevertCode {
			// If this log was for an invalid transaction, only save the request if it hasn't been saved before
			if db.as.GetPossibleRequestInfo(txRes.IncomingRequest.MessageID) != nil {
				continue
			}
		}

		requests = append(requests, machine.EVMRequestInfo{
			RequestId: txRes.IncomingRequest.MessageID,
			LogIndex:  blockInfo.FirstAVMLog().Uint64() + uint64(i),
		})
	}

	if err := db.as.SaveBlock(block.Header(), avmLogIndex, requests); err != nil {
		return err
	}

	db.chainFeed.Send(ethcore.ChainEvent{Block: block, Hash: block.Hash(), Logs: ethLogs})
	db.chainHeadFeed.Send(ethcore.ChainEvent{Block: block, Hash: block.Hash(), Logs: ethLogs})
	if len(ethLogs) > 0 {
		db.logsFeed.Send(ethLogs)
	}
	return nil
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

func (db *TxDB) GetRequest(requestId common.Hash) (*evm.TxResult, error) {
	requestCandidate := db.as.GetPossibleRequestInfo(requestId)
	if requestCandidate == nil {
		return nil, nil
	}
	logVal, err := core.GetSingleLog(db.lookup, new(big.Int).SetUint64(*requestCandidate))
	if err != nil {
		return nil, err
	}
	res, err := evm.NewTxResultFromValue(logVal)
	if err != nil {
		return nil, err
	}
	if res.IncomingRequest.MessageID != requestId {
		return nil, nil
	}
	return res, nil
}

func (db *TxDB) GetMachineBlockResults(block *machine.BlockInfo) ([]*evm.TxResult, error) {
	blockLog, err := core.GetSingleLog(db.lookup, new(big.Int).SetUint64(block.BlockLog))
	if err != nil {
		return nil, err
	}
	res, err := evm.NewBlockResultFromValue(blockLog)
	if err != nil {
		return nil, err
	}
	return db.GetBlockResults(res)
}

func (db *TxDB) GetBlock(height uint64) (*machine.BlockInfo, error) {
	count, err := db.BlockCount()
	if err != nil {
		return nil, err
	}
	if height >= count {
		return nil, nil
	}
	return db.as.GetBlockInfo(height)
}

func (db *TxDB) BlockCount() (uint64, error) {
	return db.as.BlockCount()
}

func (db *TxDB) LatestBlock() (uint64, error) {
	blockCount, err := db.as.BlockCount()
	if err != nil {
		return 0, err
	}
	if blockCount == 0 {
		return 0, errors.New("no blocks")
	}
	return blockCount - 1, nil
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

func (db *TxDB) LatestSnapshot() (*snapshot.Snapshot, error) {
	block, err := db.LatestBlock()
	if err != nil {
		return nil, err
	}
	return db.GetSnapshot(block)
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
