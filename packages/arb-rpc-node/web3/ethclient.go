package web3

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type EthClient struct {
	srv    *Server
	events *filters.EventSystem
	filter *filters.PublicFilterAPI
}

func NewEthClient(srv *aggregator.Server, ganacheMode bool) *EthClient {
	return &EthClient{
		srv:    NewServer(srv, ganacheMode),
		events: filters.NewEventSystem(srv, false),
		filter: filters.NewPublicFilterAPI(srv, false, 2*time.Minute),
	}
}

func blockNum(blockNumber *big.Int) rpc.BlockNumberOrHash {
	var blockNum *rpc.BlockNumber
	if blockNumber != nil {
		tmp := blockNumber.Int64()
		blockNum = (*rpc.BlockNumber)(&tmp)
	} else {
		pending := rpc.PendingBlockNumber
		blockNum = &pending
	}
	return rpc.BlockNumberOrHash{BlockNumber: blockNum}
}

func (c *EthClient) BalanceAt(_ context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	bal, err := c.srv.GetBalance(&account, blockNum(blockNumber))
	if err != nil {
		return nil, err
	}
	return bal.ToInt(), nil
}

func (c *EthClient) CodeAt(_ context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return c.srv.GetCode(&contract, blockNum(blockNumber))
}

func (c *EthClient) CallContract(_ context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	args := CallTxArgs{
		From:     &call.From,
		To:       call.To,
		Gas:      (*hexutil.Uint64)(&call.Gas),
		GasPrice: (*hexutil.Big)(call.GasPrice),
		Value:    (*hexutil.Big)(call.Value),
		Data:     (*hexutil.Bytes)(&call.Data),
	}
	return c.srv.Call(args, blockNum(blockNumber))
}

func (c *EthClient) PendingCodeAt(_ context.Context, account common.Address) ([]byte, error) {
	pending := rpc.PendingBlockNumber
	block := rpc.BlockNumberOrHash{BlockNumber: &pending}
	return c.srv.GetCode(&account, block)
}

// Treats a null blockNumber as the latest block, not pending
func (c *EthClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	var rpcBlockNumber rpc.BlockNumber
	if blockNumber == nil {
		rpcBlockNumber = rpc.BlockNumber(rpc.LatestBlockNumber)
	} else {
		if !blockNumber.IsInt64() {
			return 0, errors.New("block number is not int64")
		}
		rpcBlockNumber = rpc.BlockNumber(blockNumber.Int64())
	}
	block := rpc.BlockNumberOrHash{BlockNumber: &rpcBlockNumber}
	count, err := c.srv.GetTransactionCount(ctx, &account, block)
	if err != nil {
		return 0, err
	}
	return uint64(count), err
}

func (c *EthClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	pending := rpc.PendingBlockNumber
	block := rpc.BlockNumberOrHash{BlockNumber: &pending}
	count, err := c.srv.GetTransactionCount(ctx, &account, block)
	if err != nil {
		return 0, err
	}
	return uint64(count), err
}

func (c *EthClient) SuggestGasPrice(_ context.Context) (*big.Int, error) {
	gasPriceRaw, err := c.srv.GasPrice()
	return (*big.Int)(gasPriceRaw), err
}

func (c *EthClient) ChainID(_ context.Context) (*big.Int, error) {
	return c.srv.srv.ChainId(), nil
}

func (c *EthClient) EstimateGas(_ context.Context, call ethereum.CallMsg) (uint64, error) {
	args := CallTxArgs{
		From:     &call.From,
		To:       call.To,
		Gas:      (*hexutil.Uint64)(&call.Gas),
		GasPrice: (*hexutil.Big)(call.GasPrice),
		Value:    (*hexutil.Big)(call.Value),
		Data:     (*hexutil.Bytes)(&call.Data),
	}
	gas, err := c.srv.EstimateGas(args)
	if err != nil {
		return 0, err
	}
	return uint64(gas), err
}

func (c *EthClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.srv.srv.SendTransaction(ctx, tx)
}

func (c *EthClient) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	logs, err := c.filter.GetLogs(ctx, filters.FilterCriteria(query))
	if err != nil {
		return nil, err
	}
	parsedLogs := make([]types.Log, 0, len(logs))
	for _, l := range logs {
		parsedLogs = append(parsedLogs, *l)
	}
	return parsedLogs, nil
}

func (c *EthClient) SubscribeFilterLogs(_ context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	logsChan := make(chan []*types.Log)
	sub, err := c.events.SubscribeLogs(query, logsChan)
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
		for logs := range logsChan {
			for _, l := range logs {
				ch <- *l
			}
		}
	}()
	return sub, nil
}

func (c *EthClient) TransactionReceipt(_ context.Context, txHash common.Hash) (*types.Receipt, error) {
	res, block, err := c.srv.getTransactionInfoByHash(txHash.Bytes())
	if err != nil || res == nil {
		return nil, err
	}
	return res.ToEthReceipt(arbcommon.NewHashFromEth(block.Header.Hash())), nil
}

func (c *EthClient) TransactionByHash(_ context.Context, txHash common.Hash) (*types.Transaction, bool, error) {
	res, _, err := c.srv.getTransactionInfoByHash(txHash.Bytes())
	if err != nil || res == nil {
		return nil, false, err
	}
	tx, err := evm.GetTransaction(res)
	if err != nil {
		return nil, false, err
	}
	return tx.Tx, false, nil
}

func (c *EthClient) BlockByHash(_ context.Context, hash common.Hash) (*types.Block, error) {
	info, err := c.srv.srv.BlockInfoByHash(arbcommon.NewHashFromEth(hash))
	if err != nil || info == nil {
		return nil, err
	}
	_, results, err := c.srv.srv.GetMachineBlockResults(info)
	if err != nil || results == nil {
		return nil, err
	}
	processedTxes := evm.FilterEthTxResults(results)
	txes := make([]*types.Transaction, 0, len(processedTxes))
	receipts := make([]*types.Receipt, 0, len(processedTxes))
	for _, res := range processedTxes {
		txes = append(txes, res.Tx)
		receipts = append(receipts, res.Result.ToEthReceipt(arbcommon.NewHashFromEth(hash)))
	}
	return types.NewBlock(info.Header, txes, nil, receipts, new(trie.Trie)), nil
}

func (c *EthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	if number == nil {
		return c.srv.srv.LatestBlockHeader()
	}
	info, err := c.srv.srv.BlockInfoByNumber(number.Uint64())
	if err != nil || info == nil {
		return nil, err
	}
	return info.Header, nil
}

func (c *EthClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return big.NewInt(0), nil
}
