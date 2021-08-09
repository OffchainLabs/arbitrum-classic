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

package ethbridge

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/pkg/errors"
)

type NodeWatcher struct {
	con          *ethbridgecontracts.INode
	baseCallOpts bind.CallOpts
}

func NewNodeWatcher(address ethcommon.Address, client ethutils.EthClient, callOpts bind.CallOpts) (*NodeWatcher, error) {
	con, err := ethbridgecontracts.NewINode(address, client)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &NodeWatcher{
		con:          con,
		baseCallOpts: callOpts,
	}, nil
}

func (n *NodeWatcher) getCallOpts(ctx context.Context) *bind.CallOpts {
	opts := n.baseCallOpts
	opts.Context = ctx
	return &opts
}

func (n *NodeWatcher) Prev(ctx context.Context) (*big.Int, error) {
	prev, err := n.con.Prev(n.getCallOpts(ctx))
	return prev, errors.WithStack(err)
}

func (n *NodeWatcher) DeadlineBlock(ctx context.Context) (*big.Int, error) {
	block, err := n.con.DeadlineBlock(n.getCallOpts(ctx))
	return block, errors.WithStack(err)
}

func (n *NodeWatcher) StakerCount(ctx context.Context) (*big.Int, error) {
	count, err := n.con.StakerCount(n.getCallOpts(ctx))
	return count, errors.WithStack(err)
}
