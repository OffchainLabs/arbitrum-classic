package web3

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
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
	Status            uint64       `json:"status"`
	CumulativeGasUsed uint64       `json:"cumulativeGasUsed"`
	Bloom             string       `json:"logsBloom"`
	Logs              []*types.Log `json:"logs"`
	// They are stored in the chain database.
	TxHash          common.Hash `json:"transactionHash"`
	ContractAddress string      `json:"contractAddress"`
	GasUsed         uint64      `json:"gasUsed"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        common.Hash `json:"blockHash"`
	BlockNumber      *big.Int    `json:"blockNumber"`
	TransactionIndex uint        `json:"transactionIndex"`
}

type TransactionResult struct {
	BlockHash        *common.Hash `json:"blockHash"`
	BlockNumber      *string      `json:"blockNumber"`
	From             string       `json:"from"`
	Gas              string       `json:"gas"`
	GasPrice         string       `json:"gasPrice"`
	Hash             common.Hash  `json:"hash"`
	Input            string       `json:"input"`
	Nonce            string       `json:"nonce"`
	To               *string      `json:"to"`
	TransactionIndex *uint64      `json:"transactionIndex"`
	Value            string       `json:"value"`
	V                string       `json:"v"`
	R                string       `json:"r"`
	S                string       `json:"s"`
}
