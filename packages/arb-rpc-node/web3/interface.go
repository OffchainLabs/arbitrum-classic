/*
 * Copyright 2021, Offchain Labs, Inc.
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

package web3

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type GetBlockResult struct {
	Number           *hexutil.Big      `json:"number"`
	Hash             hexutil.Bytes     `json:"hash"`
	ParentHash       hexutil.Bytes     `json:"parentHash"`
	MixDigest        hexutil.Bytes     `json:"mixHash"`
	Nonce            *types.BlockNonce `json:"nonce"`
	Sha3Uncles       hexutil.Bytes     `json:"sha3Uncles"`
	LogsBloom        hexutil.Bytes     `json:"logsBloom"`
	TransactionsRoot hexutil.Bytes     `json:"transactionsRoot"`
	StateRoot        hexutil.Bytes     `json:"stateRoot"`
	ReceiptsRoot     hexutil.Bytes     `json:"receiptsRoot"`
	Miner            hexutil.Bytes     `json:"miner"`
	Difficulty       *hexutil.Big      `json:"difficulty"`
	TotalDifficulty  *hexutil.Big      `json:"totalDifficulty"`
	ExtraData        *hexutil.Bytes    `json:"extraData"`
	Size             *hexutil.Uint64   `json:"size"`
	GasLimit         *hexutil.Uint64   `json:"gasLimit"`
	GasUsed          *hexutil.Uint64   `json:"gasUsed"`
	Timestamp        *hexutil.Uint64   `json:"timestamp"`
	Transactions     interface{}       `json:"transactions"`
	Uncles           *[]hexutil.Bytes  `json:"uncles"`

	L1BlockNumber *hexutil.Big `json:"l1BlockNumber"`
}

type CallTxArgs struct {
	From       *common.Address `json:"from"`
	To         *common.Address `json:"to"`
	Gas        *hexutil.Uint64 `json:"gas"`
	GasPrice   *hexutil.Big    `json:"gasPrice"`
	Value      *hexutil.Big    `json:"value"`
	Data       *hexutil.Bytes  `json:"data"`
	Aggregator *common.Address `json:"aggregator"`
}

type FeeSetResult struct {
	L1Transaction *hexutil.Big `json:"l1Transaction"`
	L1Calldata    *hexutil.Big `json:"l1Calldata"`
	L2Storage     *hexutil.Big `json:"l2Storage"`
	L2Computation *hexutil.Big `json:"l2Computation"`
}

type FeeStatsResult struct {
	Prices    *FeeSetResult `json:"prices"`
	UnitsUsed *FeeSetResult `json:"unitsUsed"`
	Paid      *FeeSetResult `json:"paid"`
}

type L1InboxBatchInfo struct {
	Confirmations *hexutil.Big   `json:"confirmations"`
	BlockNumber   *hexutil.Big   `json:"blockNumber"`
	LogAddress    common.Address `json:"logAddress"`
	LogTopics     []common.Hash  `json:"logTopics"`
	LogData       hexutil.Bytes  `json:"logData"`
}

// Receipt represents the results of a transaction.
type GetTransactionReceiptResult struct {
	TransactionHash   common.Hash     `json:"transactionHash"`
	TransactionIndex  hexutil.Uint64  `json:"transactionIndex"`
	BlockHash         common.Hash     `json:"blockHash"`
	BlockNumber       *hexutil.Big    `json:"blockNumber"`
	From              common.Address  `json:"from"`
	To                *common.Address `json:"to"`
	CumulativeGasUsed hexutil.Uint64  `json:"cumulativeGasUsed"`
	GasUsed           hexutil.Uint64  `json:"gasUsed"`
	EffectiveGasPrice hexutil.Uint64  `json:"effectiveGasPrice"`
	ContractAddress   *common.Address `json:"contractAddress"`
	Logs              []*types.Log    `json:"logs"`
	LogsBloom         hexutil.Bytes   `json:"logsBloom"`
	Status            hexutil.Uint64  `json:"status"`

	// Arbitrum Specific Fields
	ReturnCode       hexutil.Uint64    `json:"returnCode"`
	ReturnData       hexutil.Bytes     `json:"returnData"`
	FeeStats         *FeeStatsResult   `json:"feeStats"`
	L1BlockNumber    *hexutil.Big      `json:"l1BlockNumber"`
	L1InboxBatchInfo *L1InboxBatchInfo `json:"l1InboxBatchInfo"`
}

type ArbGetTxReceiptOpts struct {
	ReturnL1InboxBatchInfo bool `json:"returnL1InboxBatchInfo"`
}

type TransactionResult struct {
	BlockHash        *common.Hash    `json:"blockHash"`
	BlockNumber      *hexutil.Big    `json:"blockNumber"`
	From             common.Address  `json:"from"`
	Gas              hexutil.Uint64  `json:"gas"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	Hash             common.Hash     `json:"hash"`
	Input            hexutil.Bytes   `json:"input"`
	Nonce            hexutil.Uint64  `json:"nonce"`
	To               *common.Address `json:"to"`
	TransactionIndex *hexutil.Uint64 `json:"transactionIndex"`
	Value            *hexutil.Big    `json:"value"`
	V                *hexutil.Big    `json:"v"`
	R                *hexutil.Big    `json:"r"`
	S                *hexutil.Big    `json:"s"`

	// Arbitrum Specific Fields
	L1SeqNum        *hexutil.Big    `json:"l1SequenceNumber"`
	ParentRequestId *common.Hash    `json:"parentRequestId"`
	IndexInParent   *hexutil.Big    `json:"indexInParent"`
	ArbType         hexutil.Uint64  `json:"arbType"`
	ArbSubType      *hexutil.Uint64 `json:"arbSubType"`
	L1BlockNumber   *hexutil.Big    `json:"l1BlockNumber"`
}
