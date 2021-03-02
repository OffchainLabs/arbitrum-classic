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
}

type CallTxArgs struct {
	From     *common.Address `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Data     *hexutil.Bytes  `json:"data"`
}

type FeeStatsResult struct {
	WeiPerTx        *hexutil.Big `json:"weiPerTx"`
	WeiPerCalldata  *hexutil.Big `json:"weiPerCalldata"`
	WeiPerStorage   *hexutil.Big `json:"weiPerStorage"`
	WeiPerArbGas    *hexutil.Big `json:"weiPerArbGas"`
	PaidForTx       *hexutil.Big `json:"paidForTx"`
	PaidForCalldata *hexutil.Big `json:"paidForCalldata"`
	PaidForStorage  *hexutil.Big `json:"paidForStorage"`
	PaidForCompute  *hexutil.Big `json:"paidForCompute"`
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
	ContractAddress   *common.Address `json:"contractAddress"`
	Logs              []*types.Log    `json:"logs"`
	LogsBloom         hexutil.Bytes   `json:"logsBloom"`
	Status            hexutil.Uint64  `json:"status"`

	// Arbitrum Specific Fields
	ReturnCode hexutil.Uint64  `json:"returnCode"`
	ReturnData hexutil.Bytes   `json:"returnData"`
	FeeStats   *FeeStatsResult `json:"feeStats"`
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
}
