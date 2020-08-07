package web3

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type BlockNumberArgs struct{}

type AccountInfoArgs struct {
	Address  *common.Address
	BlockNum *ethrpc.BlockNumber
}

func (n *AccountInfoArgs) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONArray(buf, []interface{}{&n.Address, &n.BlockNum})
}

type GetBlockByNumberArgs struct {
	BlockNum      *ethrpc.BlockNumber
	IncludeTxData bool
}

func (n *GetBlockByNumberArgs) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONArray(buf, []interface{}{&n.BlockNum, &n.IncludeTxData})
}

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

type CallArgs struct {
	CallArgs *CallTxArgs
	BlockNum *ethrpc.BlockNumber
}

func (n *CallArgs) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONArray(buf, []interface{}{&n.CallArgs, &n.BlockNum})
}

type EmptyArgs struct{}

type SendTransactionArgs struct {
	Data *hexutil.Bytes
}

func (n *SendTransactionArgs) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONArray(buf, []interface{}{&n.Data})
}

type GetTransactionReceiptArgs struct {
	Data *hexutil.Bytes
}

func (n *GetTransactionReceiptArgs) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONArray(buf, []interface{}{&n.Data})
}

// Receipt represents the results of a transaction.
type GetTransactionReceiptResult struct {
	TransactionHash   hexutil.Bytes  `json:"transactionHash"`
	TransactionIndex  hexutil.Uint64 `json:"transactionIndex"`
	BlockHash         hexutil.Bytes  `json:"blockHash"`
	BlockNumber       *hexutil.Big   `json:"blockNumber"`
	From              hexutil.Bytes  `json:"from"`
	To                *hexutil.Bytes `json:"to"`
	GasUsed           hexutil.Uint64 `json:"gasUsed"`
	CumulativeGasUsed hexutil.Uint64 `json:"cumulativeGasUsed"`
	ContractAddress   *hexutil.Bytes `json:"contractAddress"`
	Logs              []LogResult    `json:"logs"`
	Status            hexutil.Uint64 `json:"status"`
	LogsBloom         hexutil.Bytes  `json:"logsBloom"`
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

type GetLogsArgs struct {
	FromBlock *ethrpc.BlockNumber `json:"fromBlock"`
	ToBlock   *ethrpc.BlockNumber `json:"toBlock"`
	Address   *common.Address     `json:"address"`
	Topics    []common.Hash       `json:"topics"`
	BlockHash *common.Hash        `json:"blockHash"`
}

type LogResult struct {
	Removed          bool          `json:"removed"`
	LogIndex         *string       `json:"logIndex"`
	TransactionIndex *string       `json:"transactionIndex"`
	TransactionHash  *common.Hash  `json:"transactionHash"`
	BlockHash        *common.Hash  `json:"blockHash"`
	BlockNumber      *string       `json:"blockNumber"`
	Address          string        `json:"address"`
	Data             string        `json:"data"`
	Topics           []common.Hash `json:"topics"`
}

func unmarshalJSONArray(buf []byte, fields []interface{}) error {
	wantLen := len(fields)
	if err := json.Unmarshal(buf, &fields); err != nil {
		return err
	}
	if g, e := len(fields), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in CallArgs: %d != %d", g, e)
	}
	return nil
}
