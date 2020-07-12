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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
)

type nodeInfo struct {
	EVMLogs              [][]evm.Log
	EVMTransactionHashes []common.Hash

	// These members contain the logs and messages from the assertion in this
	// node if there was one, otherwise they are empty lists
	AVMLogs     []value.Value
	AVMMessages []value.Value

	// These members are generated from the AVMLogs and stored as an
	// optimization since there are expensive to generate
	AVMLogsAccHashes []string
	AVMLogsValHashes []string

	Location *evm.NodeLocation
}

func newNodeInfo() *nodeInfo {
	return &nodeInfo{}
}

func (ni *nodeInfo) FindLogs(addresses []common.Address, topics [][]common.Hash) []evm.FullLog {
	logs := make([]evm.FullLog, 0)
	for _, evmLog := range ni.fullLogs() {
		if evmLog.MatchesQuery(addresses, topics) {
			logs = append(logs, evmLog)
		}
	}
	return logs
}

func (ni *nodeInfo) fullLogs() []evm.FullLog {
	logs := make([]evm.FullLog, 0)
	logIndex := uint64(0)
	for i, txLogs := range ni.EVMLogs {
		for _, evmLog := range txLogs {
			l := evm.FullLog{
				Log:      evmLog,
				TxIndex:  uint64(i),
				TxHash:   ni.EVMTransactionHashes[i],
				Location: ni.Location,
				Index:    logIndex,
			}
			logs = append(logs, l)
			logIndex++
		}
	}
	return logs
}

func (ni *nodeInfo) calculateBloomFilter() types.Bloom {
	ethLogs := make([]*types.Log, 0)
	for _, evmLog := range ni.fullLogs() {
		ethLogs = append(ethLogs, evmLog.ToEVMLog())
	}
	return types.BytesToBloom(types.LogsBloom(ethLogs).Bytes())
}

func (ni *nodeInfo) Equals(o *nodeInfo) bool {
	return nestedLogSlicesEqual(ni.EVMLogs, o.EVMLogs) &&
		hashSlicesEqual(ni.EVMTransactionHashes, o.EVMTransactionHashes) &&
		valueSlicesEqual(ni.AVMLogs, o.AVMLogs) &&
		valueSlicesEqual(ni.AVMMessages, o.AVMMessages) &&
		stringSlicesEqual(ni.AVMLogsAccHashes, o.AVMLogsAccHashes) &&
		stringSlicesEqual(ni.AVMLogsValHashes, o.AVMLogsValHashes) &&
		ni.Location.Equals(o.Location)
}

func (ni *nodeInfo) getTxInfo(txIndex uint64) *evm.TxInfo {
	zero := common.Hash{}

	var logsPostHash string
	if len(ni.AVMLogsAccHashes) > 0 {
		logsPostHash = ni.AVMLogsAccHashes[len(ni.AVMLogsAccHashes)-1]
	} else {
		logsPostHash = hexutil.Encode(zero[:])
	}

	logsPreHash := hexutil.Encode(zero[:])
	if txIndex > 0 {
		logsPreHash = ni.AVMLogsAccHashes[txIndex-1] // Previous acc hash
	}
	logsValHashes := ni.AVMLogsValHashes[txIndex+1:] // log acc hashes after logVal

	startLogIndex := uint64(0)
	for _, logs := range ni.EVMLogs[:txIndex] {
		startLogIndex += uint64(len(logs))
	}

	return &evm.TxInfo{
		TransactionHash:  ni.EVMTransactionHashes[txIndex],
		TransactionIndex: txIndex,
		RawVal:           ni.AVMLogs[txIndex],
		Location:         ni.Location,
		StartLogIndex:    startLogIndex,
		Proof: &evm.AVMLogProof{
			LogPreHash:   logsPreHash,
			LogPostHash:  logsPostHash,
			LogValHashes: logsValHashes,
		},
	}
}

func processNode(node *structures.Node) (*nodeInfo, error) {
	nodeInfo := newNodeInfo()

	l1TxHashString := ""
	l1TxHash := node.AssertionTxHash()
	emtpyHash := common.Hash{}
	if l1TxHash != emtpyHash {
		l1TxHashString = l1TxHash.String()
	}

	nodeInfo.Location = &evm.NodeLocation{
		NodeHeight: node.Depth(),
		NodeHash:   node.Hash().String(),
		L1TxHash:   l1TxHashString,
	}

	if node.LinkType() != valprotocol.ValidChildType {
		return nodeInfo, nil
	}

	assertion := node.Assertion()

	nodeInfo.AVMMessages = assertion.ParseOutMessages()
	nodeInfo.AVMLogs = assertion.ParseLogs()
	nodeInfo.AVMLogsValHashes = make([]string, 0, len(nodeInfo.AVMLogs))
	nodeInfo.AVMLogsAccHashes = make([]string, 0, len(nodeInfo.AVMLogs))

	acc := common.Hash{}
	for _, logsVal := range nodeInfo.AVMLogs {
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

	nodeInfo.EVMTransactionHashes = make([]common.Hash, 0, len(nodeInfo.AVMLogs))

	for _, logVal := range nodeInfo.AVMLogs {
		evmVal, err := evm.ProcessLog(logVal)
		if err != nil {
			log.Printf("VM produced invalid evm result: %v\n", err)
			continue
		}
		nodeInfo.EVMLogs = append(nodeInfo.EVMLogs, evmVal.GetLogs())

		if evmVal, ok := evmVal.(evm.Revert); ok {
			log.Printf("*********** evm.Revert occurred with message \"%v\"\n", string(evmVal.ReturnVal))
		}

		delivered, err := message.UnmarshalRawDelivered(evmVal.GetDeliveredMessage())
		if err != nil {
			return nil, err
		}
		nodeInfo.EVMTransactionHashes = append(nodeInfo.EVMTransactionHashes, delivered.TxHash())
	}
	return nodeInfo, nil
}
