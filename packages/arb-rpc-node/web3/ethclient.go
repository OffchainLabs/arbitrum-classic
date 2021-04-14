package web3

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
	"time"
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

func (c *EthClient) BalanceAt(_ context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var blockNum *int64
	if blockNumber != nil {
		tmp := blockNumber.Int64()
		blockNum = &tmp
	}
	bal, err := c.srv.GetBalance(&account, (*rpc.BlockNumber)(blockNum))
	if err != nil {
		return nil, err
	}
	return bal.ToInt(), nil
}

func (c *EthClient) CodeAt(_ context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	var blockNum *int64
	if blockNumber != nil {
		tmp := blockNumber.Int64()
		blockNum = &tmp
	}
	return c.srv.GetCode(&contract, (*rpc.BlockNumber)(blockNum))
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
	var blockNum *int64
	if blockNumber != nil {
		tmp := blockNumber.Int64()
		blockNum = &tmp
	}
	return c.srv.Call(args, (*rpc.BlockNumber)(blockNum))
}

func (c *EthClient) PendingCodeAt(_ context.Context, account common.Address) ([]byte, error) {
	blockNum := rpc.PendingBlockNumber
	return c.srv.GetCode(&account, &blockNum)
}

func (c *EthClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	blockNum := rpc.PendingBlockNumber
	count, err := c.srv.GetTransactionCount(ctx, &account, &blockNum)
	if err != nil {
		return 0, err
	}
	return uint64(count), err
}

func (c *EthClient) SuggestGasPrice(_ context.Context) (*big.Int, error) {
	return (*big.Int)(c.srv.GasPrice()), nil
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
