/*
 * Copyright 2019, Offchain Labs, Inc.
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

package coordinator

import (
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	solsha3 "github.com/miguelmota/go-solidity-sha3"

	"github.com/offchainlabs/arbitrum/packages/arb-util/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type validatorRequest interface {
}

type assertionCountRequest struct {
	resultChan chan<- int
}

type vmCreatedTxHashRequest struct {
	resultChan chan<- [32]byte
}

type txRequest struct {
	txHash     [32]byte
	resultChan chan<- txInfo
}

type findLogsRequest struct {
	fromHeight *int64
	toHeight   *int64
	address    *big.Int
	topics     [][32]byte

	resultChan chan<- []*LogInfo
}

type logsInfo struct {
	msg  evm.EthMsg
	Logs []evm.Log
}

type txInfo struct {
	Found          bool
	assertionIndex int
	RawVal         value.Value
	LogsPreHash    string
	LogsPostHash   string
	LogsValHashes  []string
	ValidatorSigs  []string
	PartialHash    string
	OnChainTxHash  string
}

type assertionInfo struct {
	TxLogs            []logsInfo
	LogsAccHashes     []string
	LogsValHashes     []string
	SequenceNum       uint64
	BeforeHash        [32]byte
	OriginalInboxHash [32]byte
}

type logResponse struct {
	Log evm.Log
	Msg evm.EthMsg
}

func (a *assertionInfo) FindLogs(address *big.Int, topics [][32]byte) []logResponse {
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

func newAssertionInfo() *assertionInfo {
	return &assertionInfo{}
}

type txTracker struct {
	txRequestIndex  int
	transactions    map[[32]byte]txInfo
	assertionInfo   []*assertionInfo
	accountNonces   map[common.Address]uint64
	vmID            [32]byte
	vmCreatedTxHash [32]byte
	requests        chan validatorRequest
}

func newTxTracker(
	vmID [32]byte,
	vmCreatedTxHash [32]byte,
) *txTracker {
	requests := make(chan validatorRequest, 100)
	return &txTracker{
		txRequestIndex:  0,
		transactions:    make(map[[32]byte]txInfo),
		assertionInfo:   make([]*assertionInfo, 0),
		accountNonces:   make(map[common.Address]uint64),
		vmID:            vmID,
		vmCreatedTxHash: vmCreatedTxHash,
		requests:        requests,
	}
}

func (tr *txTracker) AssertionCount() <-chan int {
	req := make(chan int, 1)
	tr.requests <- assertionCountRequest{req}
	return req
}

func (tr *txTracker) VMCreatedTxHash() <-chan [32]byte {
	req := make(chan [32]byte, 1)
	tr.requests <- vmCreatedTxHashRequest{req}
	return req
}

func (tr *txTracker) TxInfo(txHash [32]byte) <-chan txInfo {
	req := make(chan txInfo, 1)
	tr.requests <- txRequest{txHash, req}
	return req
}

func (tr *txTracker) FindLogs(
	fromHeight *int64,
	toHeight *int64,
	address *big.Int,
	topics [][32]byte,
) <-chan []*LogInfo {
	req := make(chan []*LogInfo, 1)
	tr.requests <- findLogsRequest{fromHeight, toHeight, address, topics, req}
	return req
}

func (tr *txTracker) processFinalizedAssertion(assertion valmessage.FinalizedAssertion) {
	info := newAssertionInfo()

	var partialHash string
	var sigs []string
	var disputableTxHash string
	zero := [32]byte{}
	logsPreHash := hexutil.Encode(zero[:])
	prop := assertion.ProposalResults
	if prop != nil {
		partialHashBytes, err := hashing.UnanimousAssertPartialHash(
			prop.SequenceNum,
			prop.BeforeHash,
			prop.NewInboxHash,
			prop.BeforeInbox,
			prop.Assertion,
		)
		if err != nil {
			panic("Could not create partial hash")
		}
		partialHash = hexutil.Encode(partialHashBytes[:])

		// Encode assertion.Signatures as []string
		sigs = make([]string, 0, len(assertion.Signatures))
		for _, sig := range assertion.Signatures {
			sigs = append(sigs, hexutil.Encode(sig))
		}

		if len(tr.assertionInfo) > 0 {
			prev := tr.assertionInfo[len(tr.assertionInfo)-1]
			if prop.SequenceNum > prev.SequenceNum &&
				prop.BeforeHash == prev.BeforeHash &&
				prop.BeforeInbox == prev.OriginalInboxHash &&
				len(prev.LogsAccHashes) > 0 {
				logsPreHash = prev.LogsAccHashes[len(prev.LogsAccHashes)-1]
			}
		}
		info.SequenceNum = prop.SequenceNum
		info.BeforeHash = prop.BeforeHash
		info.OriginalInboxHash = prop.BeforeInbox
	} else {
		disputableTxHash = hexutil.Encode(assertion.OnChainTxHash)
	}

	logs := assertion.NewLogs()
	info.LogsValHashes = make([]string, 0, len(logs))
	info.LogsAccHashes = make([]string, 0, len(logs))
	acc, _ := hexutil.Decode(logsPreHash)
	for _, logsVal := range logs {
		logsValHash := logsVal.Hash()
		info.LogsValHashes = append(info.LogsValHashes,
			hexutil.Encode(logsValHash[:]))
		acc = solsha3.SoliditySHA3(
			solsha3.Bytes32(acc),
			solsha3.Bytes32(logsValHash),
		)
		info.LogsAccHashes = append(info.LogsAccHashes,
			hexutil.Encode(acc))
	}

	var logsPostHash string
	if len(logs) > 0 {
		logsPostHash = info.LogsAccHashes[len(info.LogsAccHashes)-1]
	} else {
		logsPostHash = hexutil.Encode(zero[:])
	}

	for i, logVal := range logs {
		if i > 0 {
			logsPreHash = info.LogsAccHashes[i-1] // Previous acc hash
		}
		logsValHashes := info.LogsValHashes[i+1:] // log acc hashes after logVal

		txInfo := txInfo{
			Found:          true,
			assertionIndex: len(tr.assertionInfo),
			RawVal:         logVal,
			LogsPreHash:    logsPreHash,
			LogsPostHash:   logsPostHash,
			LogsValHashes:  logsValHashes,
			ValidatorSigs:  sigs,
			PartialHash:    partialHash,
			OnChainTxHash:  disputableTxHash,
		}

		evmVal, err := evm.ProcessLog(logVal)
		if err != nil {
			log.Printf("VM produced invalid evm result: %v\n", err)
			continue
		}
		switch evmVal := evmVal.(type) {
		case evm.Stop:
			info.TxLogs = append(info.TxLogs, logsInfo{evmVal.Msg, evmVal.Logs})
		case evm.Return:
			info.TxLogs = append(info.TxLogs, logsInfo{evmVal.Msg, evmVal.Logs})
		case evm.Revert:
		}

		msg := evmVal.GetEthMsg()
		log.Println("Coordinator got response for", hexutil.Encode(msg.Data.TxHash[:]))
		tr.transactions[msg.Data.TxHash] = txInfo
	}
	tr.assertionInfo = append(tr.assertionInfo, info)
}

func (tr *txTracker) processRequest(request validatorRequest) {
	switch request := request.(type) {
	case vmCreatedTxHashRequest:
		request.resultChan <- tr.vmCreatedTxHash
	case assertionCountRequest:
		request.resultChan <- len(tr.assertionInfo) - 1
	case txRequest:
		tx, ok := tr.transactions[request.txHash]
		if ok {
			request.resultChan <- tx
		} else {
			request.resultChan <- txInfo{Found: false}
		}
	case findLogsRequest:
		startHeight := int64(0)
		endHeight := int64(len(tr.assertionInfo))
		if request.fromHeight != nil && *request.fromHeight > int64(0) {
			startHeight = *request.fromHeight
		}
		if request.toHeight != nil {
			altEndHeight := *request.toHeight + 1
			if endHeight > altEndHeight {
				endHeight = altEndHeight
			}
		}
		logs := make([]*LogInfo, 0)
		if startHeight >= int64(len(tr.assertionInfo)) {
			request.resultChan <- logs
			break
		}
		assertions := tr.assertionInfo[startHeight:endHeight]

		for i, assertion := range assertions {
			assertionLogs := assertion.FindLogs(request.address, request.topics)
			for j, evmLog := range assertionLogs {
				addressBytes := evmLog.Log.ContractID.ToBytes()
				topicStrings := make([]string, 0, len(evmLog.Log.Topics))
				for _, topic := range evmLog.Log.Topics {
					topicStrings = append(topicStrings, hexutil.Encode(topic[:]))
				}
				txHash := evmLog.Msg.MsgHash(tr.vmID)
				logs = append(logs, &LogInfo{
					Address:          hexutil.Encode(addressBytes[12:]),
					BlockHash:        hexutil.Encode(txHash[:]),
					BlockNumber:      "0x" + strconv.FormatInt(startHeight+int64(i), 16),
					Data:             hexutil.Encode(evmLog.Log.Data[:]),
					LogIndex:         "0x" + strconv.FormatInt(int64(j), 16),
					Topics:           topicStrings,
					TransactionIndex: "0x0",
					TransactionHash:  hexutil.Encode(txHash[:]),
				})
			}
		}
		request.resultChan <- logs
	}
}

func (tr *txTracker) handleTxResults(completedCalls chan valmessage.FinalizedAssertion) {
	for {
		select {
		case finalizedAssertion := <-completedCalls:
			tr.processFinalizedAssertion(finalizedAssertion)
		case request := <-tr.requests:
			tr.processRequest(request)
		}
	}
}
