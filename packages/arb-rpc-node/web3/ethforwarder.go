/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/pkg/errors"
)

const nonMutatingModeError = "mutating transactions are disabled on this node"

type ForwarderServer struct {
	srv    *aggregator.Server
	ethSrv *Server
	mode   RpcMode
}

func NewForwarderServer(
	srv *aggregator.Server,
	ethSrv *Server,
	mode RpcMode,
) *ForwarderServer {
	return &ForwarderServer{
		srv:    srv,
		ethSrv: ethSrv,
		mode:   mode,
	}
}

func (f *ForwarderServer) GetTransactionCount(ctx context.Context, address *common.Address, blockNum rpc.BlockNumberOrHash) (hexutil.Uint64, error) {
	return f.ethSrv.getTransactionCountInner(ctx, address, blockNum, f.mode == ForwardingOnlyMode)
}

func (f *ForwarderServer) SendRawTransaction(ctx context.Context, data hexutil.Bytes) (hexutil.Bytes, error) {
	if f.mode == NonMutatingMode {
		return nil, errors.New(nonMutatingModeError)
	}

	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(data, tx); err != nil {
		return nil, err
	}
	err := f.srv.SendTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	return tx.Hash().Bytes(), nil
}
