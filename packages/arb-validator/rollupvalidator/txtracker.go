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

// txTracker is thread safe
type txTracker struct {
	rollup.NoopListener
	chainAddress common.Address

	// The RWMutex protects the variables listed below it
	sync.RWMutex
	txDB          *txDB
	maxNodeHeight uint64
	initialized   bool
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
		initialized:   false,
	}, nil
}

// Delete assertion and transaction data from the reorged blocks if there are any
func (tr *txTracker) RestartingFromLatestValid(_ context.Context, _ *rollup.ChainObserver, node *structures.Node) {
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
	}()
}

// AddedToChain is called when this listener is initially added to the
// chain. If the listener was already added to a previous chain observer, we
// must be restarting after a reorg and this function does nothing. When this
// method is called for the first time, it processes all nodes that are valid,
// but have not yet been confirmed and saved into the longterm db
func (tr *txTracker) AddedToChain(_ context.Context, chain *rollup.ChainObserver) {
	tr.Lock()
	defer tr.Unlock()
	if tr.initialized {
		return
	}
	tr.initialized = true
	nodesToProcess := chain.PendingCorrectNodes()
	go func() {
		tr.Lock()
		defer tr.Unlock()
		for _, node := range nodesToProcess {
			tr.processNextNode(node)
		}
	}()
}

func (tr *txTracker) AdvancedKnownNode(_ context.Context, _ *rollup.ChainObserver, node *structures.Node) {
	go func() {
		tr.Lock()
		defer tr.Unlock()
		tr.processNextNode(node)
	}()
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

// processNextNode must be called with a write lock
func (tr *txTracker) processNextNode(node *structures.Node) {
	// We must have already processed this node if it is olded than the latest
	// node that we've seen
	sawOldNode := node.Depth() < tr.maxNodeHeight
	if sawOldNode {
		return
	}
	nodeInfo, transactions := processNode(node, tr.chainAddress)
	tr.txDB.addUnconfirmedNode(nodeInfo, transactions)
	tr.maxNodeHeight = node.Depth()
}

func (tr *txTracker) OutputMsgVal(ctx context.Context, nodeHash common.Hash, msgIndex int64) (value.Value, error) {
	tr.RLock()
	defer tr.RUnlock()
	select {
	case <-ctx.Done():
		return nil, errors.New("call timed out")
	default:
	}

	height, err := tr.txDB.lookupNodeHeight(nodeHash)
	if err != nil {
		return nil, err
	}

	nodeData, err := tr.txDB.lookupNodeRecord(height, nodeHash)
	if err != nil {
		return nil, err
	}

	if msgIndex < 0 || msgIndex >= int64(len(nodeData.AVMMessages)) {
		return nil, err
	}
	return nodeData.AVMMessages[msgIndex], nil
}

func (tr *txTracker) TxInfo(ctx context.Context, txHash common.Hash) (txInfo, error) {
	tr.RLock()
	defer tr.RUnlock()
	select {
	case <-ctx.Done():
		return txInfo{Found: false}, errors.New("call timed out")
	default:
	}
	tx, err := tr.txDB.lookupTxRecord(txHash)
	if err != nil || tx == nil {
		return txInfo{Found: false}, err
	}
	nodeInfo, err := tr.txDB.lookupNodeRecord(tx.NodeHeight, tx.NodeHash.Unmarshal())
	if err != nil {
		return txInfo{Found: false}, nil
	}
	return getTxInfo(txHash, nodeInfo, tx), nil
}

func (tr *txTracker) AssertionCount(ctx context.Context) (uint64, error) {
	tr.RLock()
	defer tr.RUnlock()
	select {
	case <-ctx.Done():
		return 0, errors.New("call timed out")
	default:
	}
	return tr.maxNodeHeight, nil
}

func (tr *txTracker) FindLogs(
	ctx context.Context,
	fromHeight *int64,
	toHeight *int64,
	address *common.Address,
	topics []common.Hash,
) ([]*validatorserver.LogInfo, error) {
	tr.RLock()
	defer tr.RUnlock()
	select {
	case <-ctx.Done():
		return nil, errors.New("call timed out")
	default:
	}
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
		return logs, nil
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
	return logs, nil
}
