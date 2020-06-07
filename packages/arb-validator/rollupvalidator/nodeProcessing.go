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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
)

type nodeInfo struct {
	EVMLogs              []logsInfo
	EVMTransactionHashes []common.Hash

	// These members contain the logs and messages from the assertion in this
	// node if there was one, otherwise they are empty lists
	AVMLogs     []value.Value
	AVMMessages []value.Value

	// These members are generated from the AVMLogs and stored as an
	// optimization since there are expensive to generate
	AVMLogsAccHashes []string
	AVMLogsValHashes []string

	NodeHash   common.Hash
	NodeHeight uint64

	// This is the transaction hash of the l1 transaction that was responsible
	// for creating this node
	L1TxHash common.Hash
}

func (ni *nodeInfo) FindLogs(addresses []common.Address, topics [][]common.Hash) []logResponse {
	logs := make([]logResponse, 0)
	for _, txLogs := range ni.EVMLogs {
		for _, evmLog := range txLogs.Logs {
			if evmLog.MatchesQuery(addresses, topics) {
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

func valueSlicesEqual(a []value.Value, b []value.Value) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if !value.Eq(t, b[i]) {
			return false
		}
	}
	return true
}

func stringSlicesEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if t != b[i] {
			return false
		}
	}
	return true
}

func logSlicesEqual(a []logsInfo, b []logsInfo) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if !t.Equals(b[i]) {
			return false
		}
	}
	return true
}

func hashSlicesEqual(a []common.Hash, b []common.Hash) bool {
	if len(a) != len(b) {
		return false
	}
	for i, t := range a {
		if t != b[i] {
			return false
		}
	}
	return true
}

func (e *nodeInfo) Equals(o *nodeInfo) bool {
	return logSlicesEqual(e.EVMLogs, o.EVMLogs) &&
		hashSlicesEqual(e.EVMTransactionHashes, o.EVMTransactionHashes) &&
		valueSlicesEqual(e.AVMLogs, o.AVMLogs) &&
		valueSlicesEqual(e.AVMMessages, o.AVMMessages) &&
		stringSlicesEqual(e.AVMLogsAccHashes, o.AVMLogsAccHashes) &&
		stringSlicesEqual(e.AVMLogsValHashes, o.AVMLogsValHashes) &&
		e.NodeHash == o.NodeHash &&
		e.NodeHeight == o.NodeHeight
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
	nodeInfo.L1TxHash = txHash

	if node.LinkType() != valprotocol.ValidChildType {
		return nodeInfo, nil
	}

	logs := node.Assertion().Logs

	nodeInfo.AVMMessages = node.Assertion().OutMsgs
	nodeInfo.AVMLogs = node.Assertion().Logs
	nodeInfo.AVMLogsValHashes = make([]string, 0, len(logs))
	nodeInfo.AVMLogsAccHashes = make([]string, 0, len(logs))

	acc := common.Hash{}
	for _, logsVal := range logs {
		logsValHash := logsVal.Hash()
		nodeInfo.AVMLogsValHashes = append(nodeInfo.AVMLogsValHashes,
			hexutil.Encode(logsValHash[:]))
		acc = hashing.SoliditySHA3(
			hashing.Bytes32(acc),
			hashing.Bytes32(logsValHash),
		)
		nodeInfo.AVMLogsAccHashes = append(nodeInfo.AVMLogsAccHashes,
			hexutil.Encode(acc.Bytes()))
	}

	nodeInfo.EVMTransactionHashes = make([]common.Hash, 0, len(logs))
	transactions := make([]txRecordInfo, 0, len(logs))

	for i, logVal := range logs {
		evmVal, err := evm.ProcessLog(logVal, chain)
		if err != nil {
			log.Printf("VM produced invalid evm result: %v\n", err)
			continue
		}
		msg := evmVal.GetEthMsg()
		nodeInfo.EVMLogs = append(nodeInfo.EVMLogs, logsInfo{
			Logs:    evmVal.GetLogs(),
			TxIndex: uint64(i),
			TxHash:  msg.TxHash(),
		})

		if evmVal, ok := evmVal.(evm.Revert); ok {
			log.Printf("*********** evm.Revert occurred with message \"%v\"\n", string(evmVal.ReturnVal))
		}

		record := &TxRecord{
			NodeHeight:       node.Depth(),
			NodeHash:         node.Hash().MarshalToBuf(),
			TransactionIndex: uint64(i),
		}
		info := txRecordInfo{
			record: record,
			txHash: msg.TxHash(),
		}
		transactions = append(transactions, info)
		nodeInfo.EVMTransactionHashes = append(nodeInfo.EVMTransactionHashes, info.txHash)
	}
	return nodeInfo, transactions
}

func getTxInfo(txHash common.Hash, nodeInfo *nodeInfo, txRecord *TxRecord) evm.TxInfo {
	zero := common.Hash{}

	var logsPostHash string
	if len(nodeInfo.AVMLogsAccHashes) > 0 {
		logsPostHash = nodeInfo.AVMLogsAccHashes[len(nodeInfo.AVMLogsAccHashes)-1]
	} else {
		logsPostHash = hexutil.Encode(zero[:])
	}

	logsPreHash := hexutil.Encode(zero[:])
	if txRecord.TransactionIndex > 0 {
		logsPreHash = nodeInfo.AVMLogsAccHashes[txRecord.TransactionIndex-1] // Previous acc hash
	}
	logsValHashes := nodeInfo.AVMLogsValHashes[txRecord.TransactionIndex+1:] // log acc hashes after logVal

	return evm.TxInfo{
		Found:           true,
		TransactionHash: txHash,
		NodeHeight:      txRecord.NodeHeight,
		NodeHash:        txRecord.NodeHash.Unmarshal(),
		RawVal:          nodeInfo.AVMLogs[txRecord.TransactionIndex],
		LogsPreHash:     logsPreHash,
		LogsPostHash:    logsPostHash,
		LogsValHashes:   logsValHashes,
		OnChainTxHash:   nodeInfo.L1TxHash,
	}
}
