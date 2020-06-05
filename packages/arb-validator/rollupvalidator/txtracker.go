/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package rollupvalidator

import (
	"context"
	"errors"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"strconv"
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/validatorserver"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
)

type logsInfo struct {
	Logs    []evm.Log
	TxIndex uint64
	TxHash  common.Hash
}

type txInfo struct {
	Found           bool
	assertionIndex  uint64
	transactionHash common.Hash
	RawVal          value.Value
	LogsPreHash     string
	LogsPostHash    string
	LogsValHashes   []string
	OnChainTxHash   string
}

type logResponse struct {
	Log     evm.Log
	TxIndex uint64
	TxHash  common.Hash
}

func newNodeInfo() *nodeInfo {
	return &nodeInfo{}
}

func (ni *nodeInfo) calculateBloomFilter() types.Bloom {
	ethLogs := make([]*types.Log, 0)
	logIndex := uint(0)
	for _, logsInfo := range ni.EVMLogs {
		for _, ethLog := range logsInfo.Logs {
			topics := make([]ethcommon.Hash, 0, len(ethLog.Topics))
			for _, topic := range ethLog.Topics {
				topics = append(topics, topic.ToEthHash())
			}
			ethLogs = append(ethLogs, &types.Log{
				Address:     ethLog.Address.ToEthAddress(),
				Topics:      topics,
				Data:        ethLog.Data,
				BlockNumber: ni.NodeHeight,
				TxHash:      logsInfo.TxHash.ToEthHash(),
				TxIndex:     uint(logsInfo.TxIndex),
				BlockHash:   ni.NodeHash.ToEthHash(),
				Index:       logIndex,
			})
			logIndex++
		}
	}
	return types.BytesToBloom(types.LogsBloom(ethLogs).Bytes())
}

func (x *NodeMetadata) MaybeMatchesLogQuery(address *common.Address, topics []common.Hash) bool {
	logFilter := types.BytesToBloom(x.LogBloom)
	if address != nil && !logFilter.TestBytes(address[:]) {
		return false
	}
	for _, topic := range topics {
		if !logFilter.TestBytes(topic.Bytes()) {
			return false
		}
	}
	return true
}

func (ni *nodeInfo) FindLogs(address *common.Address, topics []common.Hash) []logResponse {
	logs := make([]logResponse, 0)
	for _, txLogs := range ni.EVMLogs {
		for _, evmLog := range txLogs.Logs {
			if address != nil && *address != evmLog.Address {
				continue
			}

			if len(topics) > len(evmLog.Topics) {
				continue
			}

			match := true
			for i, topic := range topics {
				if topic != evmLog.Topics[i] {
					match = false
					break
				}
			}
			if match {
				logs = append(logs, logResponse{
					Log:     evmLog,
					TxIndex: txLogs.TxIndex,
					TxHash:  txLogs.TxHash,
				})
			}
		}
	}
	return logs
}

type txTracker struct {
	sync.RWMutex
	rollup.NoopListener
	txDB          *txDB
	maxNodeHeight uint64
	chainAddress  common.Address
}

func newTxTracker(
	db machine.CheckpointStorage,
	ns machine.NodeStore,
	chainAddress common.Address,
) (*txTracker, error) {
	txdb, err := newTxDB(db, ns, chainAddress)
	if err != nil {
		return nil, err
	}
	return &txTracker{
		txDB:          txdb,
		chainAddress:  chainAddress,
		maxNodeHeight: 0,
	}, nil
}

// Delete assertion and transaction data from the reorged blocks if there are any
func (tr *txTracker) RestartingFromLatestValid(_ context.Context, chain *rollup.ChainObserver, node *structures.Node) {
	startDepth := node.Depth()
	go func() {
		tr.Lock()
		defer tr.Unlock()
		// First remove any data from reorged nodes
		for i := tr.maxNodeHeight; i > startDepth; i-- {
			if err := tr.txDB.removeUnconfirmedNode(i); err != nil {
				continue
			}
		}

		tr.maxNodeHeight = startDepth

		// Next process data for any nodes which have not yet been processed
		chain.ReplayNodesToLatestValid(tr.processNextNode)
	}()
}

func (tr *txTracker) AdvancedKnownNode(_ context.Context, _ *rollup.ChainObserver, node *structures.Node) {
	go tr.processNextNode(node)
}

func (tr *txTracker) ConfirmedNode(_ context.Context, _ *rollup.ChainObserver, ev arbbridge.ConfirmedEvent) {
	go func() {
		tr.Lock()
		defer tr.Unlock()

		if err := tr.txDB.confirmNode(ev.NodeHash); err != nil {
			log.Println(err)
			return
		}
	}()
}

func (tr *txTracker) processNextNode(node *structures.Node) {
	// We must have already processed this node
	if node.Depth() < tr.maxNodeHeight {
		return
	}
	nodeInfo, transactions := processNode(node, tr.chainAddress)
	tr.Lock()
	defer tr.Unlock()
	tr.txDB.addUnconfirmedNode(nodeInfo, transactions)
	tr.maxNodeHeight = node.Depth()
}

func (tr *txTracker) OutputMsgVal(ctx context.Context, nodeHash common.Hash, msgIndex int64) (value.Value, error) {
	ret, err := tr.callOrCancel(ctx, func() interface{} {
		height, err := tr.txDB.lookupNodeHeight(nodeHash)
		if err != nil {
			return nil
		}

		nodeData, err := tr.txDB.lookupNodeRecord(height, nodeHash)
		if err != nil {
			return nil
		}

		if msgIndex < 0 || msgIndex >= int64(len(nodeData.AVMMessages)) {
			return nil
		}
		return nodeData.AVMMessages[msgIndex]
	})
	return ret.(value.Value), err
}

func (tr *txTracker) TxInfo(ctx context.Context, txHash common.Hash) (txInfo, error) {
	ret, err := tr.callOrCancel(ctx, func() interface{} {
		tx, err := tr.txDB.lookupTxRecord(txHash)
		if err != nil {
			return txInfo{Found: false}
		}
		nodeInfo, err := tr.txDB.lookupNodeRecord(tx.NodeHeight, tx.NodeHash.Unmarshal())
		if err != nil {
			return txInfo{Found: false}
		}
		return getTxInfo(txHash, nodeInfo, tx)
	})
	return ret.(txInfo), err
}

func (tr *txTracker) AssertionCount(ctx context.Context) (uint64, error) {
	ret, err := tr.callOrCancel(ctx, func() interface{} {
		return tr.maxNodeHeight
	})
	return ret.(uint64), err
}

func (tr *txTracker) FindLogs(
	ctx context.Context,
	fromHeight *int64,
	toHeight *int64,
	address *common.Address,
	topics []common.Hash,
) ([]*validatorserver.LogInfo, error) {
	ret, err := tr.callOrCancel(ctx, func() interface{} {
		startHeight := int64(0)
		endHeight := int64(tr.maxNodeHeight)
		if fromHeight != nil && *fromHeight > int64(0) {
			startHeight = *fromHeight
		}
		if toHeight != nil {
			altEndHeight := *toHeight + 1
			if endHeight > altEndHeight {
				endHeight = altEndHeight
			}
		}
		logs := make([]*validatorserver.LogInfo, 0)
		if startHeight >= int64(tr.maxNodeHeight) {
			return logs
		}

		for i := startHeight; i < endHeight; i++ {
			nodeHash, err := tr.txDB.lookupNodeHash(uint64(i))
			if err != nil {
				continue
			}
			metadata, err := tr.txDB.lookupNodeMetadata(uint64(i), nodeHash)
			if err != nil {
				continue
			}

			if !metadata.MaybeMatchesLogQuery(address, topics) {
				continue
			}

			info, err := tr.txDB.lookupNodeRecord(uint64(i), nodeHash)
			if err != nil {
				continue
			}
			assertionLogs := info.FindLogs(address, topics)
			for j, evmLog := range assertionLogs {
				topicStrings := make([]string, 0, len(evmLog.Log.Topics))
				for _, topic := range evmLog.Log.Topics {
					topicStrings = append(topicStrings, hexutil.Encode(topic[:]))
				}

				logs = append(logs, &validatorserver.LogInfo{
					Address:          hexutil.Encode(evmLog.Log.Address[:]),
					BlockHash:        hexutil.Encode(info.NodeHash.Bytes()),
					BlockNumber:      "0x" + strconv.FormatInt(int64(info.NodeHeight), 16),
					Data:             hexutil.Encode(evmLog.Log.Data[:]),
					LogIndex:         "0x" + strconv.FormatInt(int64(j), 16),
					Topics:           topicStrings,
					TransactionIndex: "0x" + strconv.FormatInt(int64(evmLog.TxIndex), 16),
					TransactionHash:  hexutil.Encode(evmLog.TxHash[:]),
				})
			}
		}
		return logs
	})
	return ret.([]*validatorserver.LogInfo), err
}

func (tr *txTracker) callOrCancel(ctx context.Context, method func() interface{}) (interface{}, error) {
	retChan := make(chan interface{}, 1)
	go func() {
		defer close(retChan)
		tr.RLock()
		defer tr.RUnlock()
		select {
		case <-ctx.Done():
			return
		default:
		}
		retChan <- method()
	}()

	select {
	case ret := <-retChan:
		return ret, nil
	case <-ctx.Done():
		return nil, errors.New("call timed out")
	}
}
