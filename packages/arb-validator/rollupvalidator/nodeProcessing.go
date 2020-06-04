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
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
)

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

type txRecordInfo struct {
	record *TxRecord
	txHash common.Hash
}

func processNode(node *structures.Node, chain common.Address) (*nodeInfo, []txRecordInfo) {
	return processNodeImpl(
		node.Hash(),
		node.Depth(),
		node.AssertionTxHash(),
		node.LinkType(),
		node.Assertion(),
		chain,
	)
}

func processNodeImpl(
	nodeHash common.Hash,
	nodeHeight uint64,
	assertionTxHash common.Hash,
	linkType valprotocol.ChildType,
	assertion *protocol.ExecutionAssertion,
	chain common.Address,
) (*nodeInfo, []txRecordInfo) {
	nodeInfo := newNodeInfo()
	nodeInfo.NodeHash = nodeHash
	nodeInfo.NodeHeight = nodeHeight
	txHash := assertionTxHash
	nodeInfo.OnChainTxHash = hexutil.Encode(txHash[:])

	if linkType != valprotocol.ValidChildType {
		return nodeInfo, nil
	}

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
			NodeHeight:       nodeHeight,
			NodeHash:         nodeHash.MarshalToBuf(),
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
