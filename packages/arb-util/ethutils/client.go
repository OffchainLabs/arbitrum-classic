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

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ReceiptFetcher interface {
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
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
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
}

type RPCEthClient struct {
	*ethclient.Client
	rpc *rpc.Client
}

type BlockInfo struct {
	Hash       common.Hash    `json:"hash"`
	ParentHash common.Hash    `json:"parentHash"`
	Time       hexutil.Uint64 `json:"timestamp"`
	Number     *hexutil.Big   `json:"number"`
}

func NewRPCEthClient(url string) (*RPCEthClient, error) {
	ethcl, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	rpccl, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	return &RPCEthClient{
		Client: ethcl,
		rpc:    rpccl,
	}, nil
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
