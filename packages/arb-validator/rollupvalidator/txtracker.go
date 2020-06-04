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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/valprotocol"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollup"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
	"log"
	"math/big"
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
	msg  evm.EthBridgeMessage
	Logs []evm.Log
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

type assertionInfo struct {
	TxLogs            []logsInfo
	LogsAccHashes     []string
	LogsValHashes     []string
	OutMessages       []value.Value
	AssertNodeHash    common.Hash
	TransactionHashes []common.Hash
}

type logResponse struct {
	Log evm.Log
	Msg evm.EthBridgeMessage
}

func newAssertionInfo() *assertionInfo {
	return &assertionInfo{}
}

func (a *assertionInfo) FindLogs(address *big.Int, topics []common.Hash) []logResponse {
	logs := make([]logResponse, 0)
	for _, txLogs := range a.TxLogs {
		for _, evmLog := range txLogs.Logs {
			if address != nil && !value.NewIntValue(address).Equal(evmLog.ContractID) {
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
				logs = append(logs, logResponse{evmLog, txLogs.msg})
			}
		}
	}
	return logs
}

type txTracker struct {
	*sync.RWMutex
	rollup.NoopListener
	transactions  map[common.Hash]txInfo
	assertionInfo []*assertionInfo
	assertionMap  map[common.Hash]*assertionInfo
	chainAddress  common.Address
}

func newTxTracker(
	chainAddress common.Address,
) *txTracker {
	return &txTracker{
		transactions:  make(map[common.Hash]txInfo),
		assertionInfo: make([]*assertionInfo, 0),
		assertionMap:  make(map[common.Hash]*assertionInfo),
		chainAddress:  chainAddress,
	}
}

// Delete assertion and transaction data from the reorged blocks
func (tr *txTracker) RestartingFromLatestValid(_ context.Context, chain *rollup.ChainObserver, node *structures.Node) {
	reorgDepth := node.Depth()
	go func() {
		tr.Lock()
		defer tr.Unlock()
		assertionsToReorg := tr.assertionInfo[reorgDepth:]
		for _, assertion := range assertionsToReorg {
			delete(tr.assertionMap, assertion.AssertNodeHash)
			for _, txHash := range assertion.TransactionHashes {
				delete(tr.transactions, txHash)
			}
		}
		tr.assertionInfo = tr.assertionInfo[:reorgDepth]

	}()
}

func (tr *txTracker) AdvancedKnownNode(_ context.Context, _ *rollup.ChainObserver, node *structures.Node) {
	go func() {
		assertionInfo, transactions := processNode(node, tr.chainAddress)
		tr.Lock()
		defer tr.Unlock()
		for _, tx := range transactions {
			tr.transactions[tx.transactionHash] = tx
		}
		tr.assertionInfo = append(tr.assertionInfo, assertionInfo)
		tr.assertionMap[assertionInfo.AssertNodeHash] = assertionInfo
	}()
}

func (tr *txTracker) AssertionCount() int {
	tr.RLock()
	defer tr.RUnlock()
	return len(tr.assertionInfo) - 1
}

func (tr *txTracker) OutputMsgVal(assertionHash common.Hash, msgIndex int64) value.Value {
	tr.RLock()
	defer tr.RUnlock()
	assertionVal, ok := tr.assertionMap[assertionHash]
	if !ok || msgIndex < 0 || msgIndex >= int64(len(assertionVal.OutMessages)) {
		return nil
	}
	return assertionVal.OutMessages[msgIndex]
}

func (tr *txTracker) TxInfo(txHash common.Hash) txInfo {
	tr.RLock()
	defer tr.RUnlock()
	tx, ok := tr.transactions[txHash]
	if !ok {
		return txInfo{Found: false}
	}
	return tx
}

func (tr *txTracker) FindLogs(
	fromHeight *int64,
	toHeight *int64,
	address *big.Int,
	topics []common.Hash,
) []*validatorserver.LogInfo {
	tr.RLock()
	defer tr.RUnlock()
	startHeight := int64(0)
	endHeight := int64(len(tr.assertionInfo))
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
	if startHeight >= int64(len(tr.assertionInfo)) {
		return logs
	}
	assertions := tr.assertionInfo[startHeight:endHeight]

	for i, assertion := range assertions {
		assertionLogs := assertion.FindLogs(address, topics)
		for j, evmLog := range assertionLogs {
			addressBytes := evmLog.Log.ContractID.ToBytes()
			topicStrings := make([]string, 0, len(evmLog.Log.Topics))
			for _, topic := range evmLog.Log.Topics {
				topicStrings = append(topicStrings, hexutil.Encode(topic[:]))
			}

			logs = append(logs, &validatorserver.LogInfo{
				Address:          hexutil.Encode(addressBytes[12:]),
				BlockHash:        hexutil.Encode(evmLog.Msg.TxHash[:]),
				BlockNumber:      "0x" + strconv.FormatInt(startHeight+int64(i), 16),
				Data:             hexutil.Encode(evmLog.Log.Data[:]),
				LogIndex:         "0x" + strconv.FormatInt(int64(j), 16),
				Topics:           topicStrings,
				TransactionIndex: "0x0",
				TransactionHash:  hexutil.Encode(evmLog.Msg.TxHash[:]),
			})
		}
	}
	return logs
}

func processNode(node *structures.Node, chain common.Address) (*assertionInfo, []txInfo) {
	assertionInfo := newAssertionInfo()
	assertionInfo.AssertNodeHash = node.Hash()

	if node.LinkType() != valprotocol.ValidChildType {
		return assertionInfo, nil
	}

	zero := common.Hash{}
	logsPreHash := hexutil.Encode(zero[:])

	txHash := node.AssertionTxHash()
	disputableTxHash := hexutil.Encode(txHash[:])

	assertion := node.Assertion()

	logs := assertion.Logs

	assertionInfo.OutMessages = assertion.OutMsgs
	assertionInfo.LogsValHashes = make([]string, 0, len(logs))
	assertionInfo.LogsAccHashes = make([]string, 0, len(logs))

	acc := common.Hash{}
	for _, logsVal := range logs {
		logsValHash := logsVal.Hash()
		assertionInfo.LogsValHashes = append(assertionInfo.LogsValHashes,
			hexutil.Encode(logsValHash[:]))
		acc = hashing.SoliditySHA3(
			hashing.Bytes32(acc),
			hashing.Bytes32(logsValHash),
		)
		assertionInfo.LogsAccHashes = append(assertionInfo.LogsAccHashes,
			hexutil.Encode(acc.Bytes()))
	}

	var logsPostHash string
	if len(logs) > 0 {
		logsPostHash = assertionInfo.LogsAccHashes[len(assertionInfo.LogsAccHashes)-1]
	} else {
		logsPostHash = hexutil.Encode(zero[:])
	}

	assertionInfo.TransactionHashes = make([]common.Hash, 0, len(logs))
	transactions := make([]txInfo, 0, len(logs))

	for i, logVal := range logs {
		if i > 0 {
			logsPreHash = assertionInfo.LogsAccHashes[i-1] // Previous acc hash
		}
		logsValHashes := assertionInfo.LogsValHashes[i+1:] // log acc hashes after logVal

		evmVal, err := evm.ProcessLog(logVal, chain)
		if err != nil {
			log.Printf("VM produced invalid evm result: %v\n", err)
			continue
		}
		switch evmVal := evmVal.(type) {
		case evm.Stop:
			assertionInfo.TxLogs = append(assertionInfo.TxLogs, logsInfo{evmVal.Msg, evmVal.Logs})
		case evm.Return:
			assertionInfo.TxLogs = append(assertionInfo.TxLogs, logsInfo{evmVal.Msg, evmVal.Logs})
		case evm.Revert:
			log.Printf("*********** evm.Revert occurred with message \"%v\"\n", string(evmVal.ReturnVal))
		}

		msg := evmVal.GetEthMsg()
		log.Println("Coordinator got response for", hexutil.Encode(msg.TxHash[:]))
		info := txInfo{
			Found:           true,
			transactionHash: msg.TxHash,
			assertionIndex:  node.Depth(),
			RawVal:          logVal,
			LogsPreHash:     logsPreHash,
			LogsPostHash:    logsPostHash,
			LogsValHashes:   logsValHashes,
			OnChainTxHash:   disputableTxHash,
		}
		transactions = append(transactions, info)
		assertionInfo.TransactionHashes = append(assertionInfo.TransactionHashes, info.transactionHash)
	}
	return assertionInfo, transactions
}
