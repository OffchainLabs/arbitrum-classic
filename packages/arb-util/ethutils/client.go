/*
 * Copyright 2020, Offchain Labs, Inc.
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

package ethutils

import (
	"context"
	"encoding/json"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const maxErrCount = 5

type ReceiptFetcher interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
}

type EthClient interface {
	bind.ContractBackend
	ReceiptFetcher

	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
	BlockInfoByNumber(ctx context.Context, number *big.Int) (*BlockInfo, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error)
	PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error)
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
}

type RPCEthClient struct {
	sync.RWMutex

	url      string
	eth      *ethclient.Client
	rpc      *rpc.Client
	errCount uint64
}

type BlockInfo struct {
	Hash       common.Hash    `json:"hash"`
	ParentHash common.Hash    `json:"parentHash"`
	Time       hexutil.Uint64 `json:"timestamp"`
	Number     *hexutil.Big   `json:"number"`
}

func NewRPCEthClient(url string) (*RPCEthClient, error) {
	r := &RPCEthClient{url: url}
	err := r.reconnect()
	return r, err
}

func (r *RPCEthClient) reconnect() error {
	r.Lock()
	defer r.Unlock()
	if r.errCount < maxErrCount {
		// We must have already reconnected
		return nil
	}
	rpccl, err := rpc.Dial(r.url)
	if err != nil {
		return err
	}
	r.eth = ethclient.NewClient(rpccl)
	r.rpc = rpccl
	r.errCount = 0
	return nil
}

func (r *RPCEthClient) handleCallErr(err error) error {
	if err == nil {
		// Reset err count if any call is working since we're looking for a connection error
		atomic.StoreUint64(&r.errCount, 0)
		return nil
	}
	totalErrCount := atomic.AddUint64(&r.errCount, 1)

	// If we've had above a threshold number of errors, reinitialize the connection
	if totalErrCount >= maxErrCount {
		if err := r.reconnect(); err != nil {
			return err
		}
	}
	return err
}

func (r *RPCEthClient) BlockInfoByNumber(ctx context.Context, number *big.Int) (*BlockInfo, error) {
	var raw json.RawMessage
	var numParam string
	if number != nil {
		numParam = hexutil.EncodeBig(number)
	} else {
		numParam = "latest"
	}
	if err := r.rpc.CallContext(ctx, &raw, "eth_getBlockByNumber", numParam, false); err != nil {
		return nil, err
	}
	if len(raw) == 0 {
		return nil, ethereum.NotFound
	}
	var ret BlockInfo
	if err := json.Unmarshal(raw, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (r *RPCEthClient) ChainID(ctx context.Context) (*big.Int, error) {
	r.RLock()
	val, err := r.eth.ChainID(ctx)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	r.RLock()
	val, err := r.eth.CodeAt(ctx, account, blockNumber)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	r.RLock()
	val, err := r.eth.BalanceAt(ctx, account, blockNumber)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	r.RLock()
	val, err := r.eth.CallContract(ctx, msg, blockNumber)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	r.RLock()
	val, err := r.eth.HeaderByNumber(ctx, number)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	r.RLock()
	val, err := r.eth.PendingCodeAt(ctx, account)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	r.RLock()
	val, err := r.eth.PendingNonceAt(ctx, account)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) PendingCallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	r.RLock()
	val, err := r.eth.PendingCallContract(ctx, msg)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	r.RLock()
	val, err := r.eth.SuggestGasPrice(ctx)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	r.RLock()
	val, err := r.eth.SuggestGasTipCap(ctx)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	r.RLock()
	val, err := r.eth.EstimateGas(ctx, msg)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	r.RLock()
	err := r.eth.SendTransaction(ctx, tx)
	r.RUnlock()
	return r.handleCallErr(err)
}

func (r *RPCEthClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	r.RLock()
	val, err := r.eth.FilterLogs(ctx, q)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	r.RLock()
	val, err := r.eth.SubscribeFilterLogs(ctx, q, ch)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	r.RLock()
	val, err := r.eth.TransactionReceipt(ctx, txHash)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	r.RLock()
	val, err := r.eth.NonceAt(ctx, account, blockNumber)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error) {
	r.RLock()
	val, err := r.eth.HeaderByHash(ctx, hash)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error) {
	r.RLock()
	val, err := r.eth.BlockByHash(ctx, hash)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

func (r *RPCEthClient) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	r.RLock()
	tx, isPending, err = r.eth.TransactionByHash(ctx, hash)
	r.RUnlock()
	return tx, isPending, r.handleCallErr(err)
}

func (r *RPCEthClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	r.RLock()
	val, err := r.eth.TransactionInBlock(ctx, blockHash, index)
	r.RUnlock()
	return val, r.handleCallErr(err)
}

type SimulatedEthClient struct {
	*backends.SimulatedBackend
}

func (r *SimulatedEthClient) BlockInfoByNumber(ctx context.Context, number *big.Int) (*BlockInfo, error) {
	header, err := r.SimulatedBackend.HeaderByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	return &BlockInfo{
		Hash:       header.Hash(),
		ParentHash: header.ParentHash,
		Time:       hexutil.Uint64(header.Time),
		Number:     (*hexutil.Big)(header.Number),
	}, nil
}
