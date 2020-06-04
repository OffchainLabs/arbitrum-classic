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
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"strconv"
	"sync"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/validatorserver"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
)

type logsInfo struct {
	Logs    []evm.Log
	TxIndex uint64
	TxHash  common.Hash
}

type nodeInfo struct {
	TxLogs            []logsInfo
	LogsAccHashes     []string
	LogsValHashes     []string
	OutLogs           []value.Value
	OutMessages       []value.Value
	NodeHash          common.Hash
	NodeHeight        uint64
	TransactionHashes []common.Hash
	OnChainTxHash     string
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

func (a *nodeInfo) FindLogs(address *common.Address, topics []common.Hash) []logResponse {
	logs := make([]logResponse, 0)
	for _, txLogs := range a.TxLogs {
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
	*sync.RWMutex
	rollup.NoopListener
	txDB          *txDB
	maxNodeHeight uint64
	chainAddress  common.Address
}

func newTxTracker(
	db machine.CheckpointStorage,
	ns machine.NodeStore,
	chainAddress common.Address,
) *txTracker {
	return &txTracker{
		txDB:          newTxDB(db, ns, chainAddress),
		chainAddress:  chainAddress,
		maxNodeHeight: 0,
	}
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

func (tr *txTracker) AssertionCount() uint64 {
	tr.RLock()
	defer tr.RUnlock()
	return tr.maxNodeHeight
}

func (tr *txTracker) OutputMsgVal(nodeHash common.Hash, msgIndex int64) value.Value {
	tr.RLock()
	defer tr.RUnlock()
	nodeData, err := tr.txDB.lookupNodeWithHash(nodeHash)
	if err != nil {
		return nil
	}

	if msgIndex < 0 || msgIndex >= int64(len(nodeData.OutMessages)) {
		return nil
	}
	return nodeData.OutMessages[msgIndex]
}

func (tr *txTracker) TxInfo(txHash common.Hash) txInfo {
	tr.RLock()
	defer tr.RUnlock()
	tx, err := tr.txDB.lookupTxRecord(txHash)
	if err != nil {
		return txInfo{Found: false}
	}
	nodeInfo, err := tr.txDB.lookupNodeWithHeight(tx.NodeHeight)
	if err != nil {
		return txInfo{Found: false}
	}
	return getTxInfo(txHash, nodeInfo, tx)
}

func (tr *txTracker) FindLogs(
	fromHeight *int64,
	toHeight *int64,
	address *common.Address,
	topics []common.Hash,
) []*validatorserver.LogInfo {
	tr.RLock()
	defer tr.RUnlock()
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
		assertion, err := tr.txDB.lookupNodeWithHeight(uint64(i))
		if err != nil {
			continue
		}
		assertionLogs := assertion.FindLogs(address, topics)
		for j, evmLog := range assertionLogs {
			topicStrings := make([]string, 0, len(evmLog.Log.Topics))
			for _, topic := range evmLog.Log.Topics {
				topicStrings = append(topicStrings, hexutil.Encode(topic[:]))
			}

			logs = append(logs, &validatorserver.LogInfo{
				Address:          hexutil.Encode(evmLog.Log.Address[:]),
				BlockHash:        hexutil.Encode(assertion.NodeHash.Bytes()),
				BlockNumber:      "0x" + strconv.FormatInt(int64(assertion.NodeHeight), 16),
				Data:             hexutil.Encode(evmLog.Log.Data[:]),
				LogIndex:         "0x" + strconv.FormatInt(int64(j), 16),
				Topics:           topicStrings,
				TransactionIndex: "0x" + strconv.FormatInt(int64(evmLog.TxIndex), 16),
				TransactionHash:  hexutil.Encode(evmLog.TxHash[:]),
			})
		}
	}
	return logs
}

type txRecordInfo struct {
	record *TxRecord
	txHash common.Hash
}

func processNode(node *structures.Node, chain common.Address) (*nodeInfo, []txRecordInfo) {
	nodeInfo := newNodeInfo()
	nodeInfo.NodeHash = node.Hash()
	nodeInfo.NodeHeight = node.Depth()
	txHash := node.AssertionTxHash()
	nodeInfo.OnChainTxHash = hexutil.Encode(txHash[:])

	if node.LinkType() != valprotocol.ValidChildType {
		return nodeInfo, nil
	}

	assertion := node.Assertion()

	logs := assertion.Logs

	nodeInfo.OutMessages = assertion.OutMsgs
	nodeInfo.OutLogs = assertion.Logs
	nodeInfo.LogsValHashes = make([]string, 0, len(logs))
	nodeInfo.LogsAccHashes = make([]string, 0, len(logs))

	acc := common.Hash{}
	for _, logsVal := range logs {
		logsValHash := logsVal.Hash()
		nodeInfo.LogsValHashes = append(nodeInfo.LogsValHashes,
			hexutil.Encode(logsValHash[:]))
		acc = hashing.SoliditySHA3(
			hashing.Bytes32(acc),
			hashing.Bytes32(logsValHash),
		)
		nodeInfo.LogsAccHashes = append(nodeInfo.LogsAccHashes,
			hexutil.Encode(acc.Bytes()))
	}

	nodeInfo.TransactionHashes = make([]common.Hash, 0, len(logs))
	transactions := make([]txRecordInfo, 0, len(logs))

	for i, logVal := range logs {
		evmVal, err := evm.ProcessLog(logVal, chain)
		if err != nil {
			log.Printf("VM produced invalid evm result: %v\n", err)
			continue
		}
		msg := evmVal.GetEthMsg()
		nodeInfo.TxLogs = append(nodeInfo.TxLogs, logsInfo{
			Logs:    evmVal.GetLogs(),
			TxIndex: uint64(i),
			TxHash:  msg.TxHash,
		})

		if evmVal, ok := evmVal.(evm.Revert); ok {
			log.Printf("*********** evm.Revert occurred with message \"%v\"\n", string(evmVal.ReturnVal))
		}

		log.Println("Coordinator got response for", hexutil.Encode(msg.TxHash[:]))
		record := &TxRecord{
			NodeHeight:       node.Depth(),
			NodeHash:         node.Hash().MarshalToBuf(),
			TransactionIndex: uint64(i),
		}
		info := txRecordInfo{
			record: record,
			txHash: msg.TxHash,
		}
		transactions = append(transactions, info)
		nodeInfo.TransactionHashes = append(nodeInfo.TransactionHashes, info.txHash)
	}
	return nodeInfo, transactions
}

func getTxInfo(txHash common.Hash, nodeInfo *nodeInfo, txRecord *TxRecord) txInfo {
	zero := common.Hash{}

	var logsPostHash string
	if len(nodeInfo.LogsAccHashes) > 0 {
		logsPostHash = nodeInfo.LogsAccHashes[len(nodeInfo.LogsAccHashes)-1]
	} else {
		logsPostHash = hexutil.Encode(zero[:])
	}

	logsPreHash := hexutil.Encode(zero[:])
	if txRecord.TransactionIndex > 0 {
		logsPreHash = nodeInfo.LogsAccHashes[txRecord.TransactionIndex-1] // Previous acc hash
	}
	logsValHashes := nodeInfo.LogsValHashes[txRecord.TransactionIndex+1:] // log acc hashes after logVal

	return txInfo{
		Found:           true,
		transactionHash: txHash,
		assertionIndex:  txRecord.NodeHeight,
		RawVal:          nodeInfo.OutLogs[txRecord.TransactionIndex],
		LogsPreHash:     logsPreHash,
		LogsPostHash:    logsPostHash,
		LogsValHashes:   logsValHashes,
		OnChainTxHash:   nodeInfo.OnChainTxHash,
	}
}
