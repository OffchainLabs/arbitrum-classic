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

// Receipt represents the results of a transaction.
type GetTransactionReceiptResult struct {
	Status            hexutil.Uint64 `json:"status"`
	CumulativeGasUsed hexutil.Uint64 `json:"cumulativeGasUsed"`
	Bloom             hexutil.Bytes  `json:"logsBloom"`
	Logs              []*types.Log   `json:"logs"`
	// They are stored in the chain database.
	TxHash          common.Hash     `json:"transactionHash"`
	ContractAddress *common.Address `json:"contractAddress"`
	GasUsed         hexutil.Uint64  `json:"gasUsed"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        common.Hash    `json:"blockHash"`
	BlockNumber      *hexutil.Big   `json:"blockNumber"`
	TransactionIndex hexutil.Uint64 `json:"transactionIndex"`

	// Arbitrum Specific Fields
	ReturnCode hexutil.Uint64 `json:"returnCode"`
	ReturnData hexutil.Bytes  `json:"returnData"`
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
	L1SeqNum        *hexutil.Big `json:"l1SequenceNumber"`
	ParentRequestId *common.Hash `json:"parentRequestId"`
	IndexInParent   *hexutil.Big `json:"indexInParent"`
}
