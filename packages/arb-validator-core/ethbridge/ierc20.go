/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/ierc20"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type IERC20 struct {
	IERC20 *ierc20.IERC20
	client *ethclient.Client
	auth   *TransactAuth
}

func newIERC20(address ethcommon.Address, client *ethclient.Client, auth *TransactAuth) (*IERC20, error) {
	ierc20Contract, err := ierc20.NewIERC20(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to IERC20")
	}
	return &IERC20{ierc20Contract, client, auth}, nil
}

func (con *IERC20) Approve(ctx context.Context, spender common.Address, amount *big.Int) error {
	con.auth.Lock()
	defer con.auth.Unlock()
	tx, err := con.IERC20.Approve(
		con.auth.getAuth(ctx),
		spender.ToEthAddress(),
		amount,
	)
	if err != nil {
		return err
	}
	return con.waitForReceipt(ctx, tx, "Approve")
}

func (con *IERC20) waitForReceipt(ctx context.Context, tx *types.Transaction, methodName string) error {
	return waitForReceipt(ctx, con.client, con.auth.auth.From, tx, methodName)
}

type IERC20Watcher struct {
	IERC20 *ierc20.IERC20
	client ethutils.EthClient
}

func newIERC20Watcher(address ethcommon.Address, client ethutils.EthClient) (*IERC20Watcher, error) {
	ierc20Contract, err := ierc20.NewIERC20(address, client)
	if err != nil {
		return nil, errors2.Wrap(err, "Failed to connect to IERC20")
	}
	return &IERC20Watcher{ierc20Contract, client}, nil
}

func (con *IERC20Watcher) BalanceOf(ctx context.Context, account common.Address) (*big.Int, error) {
	return con.IERC20.BalanceOf(&bind.CallOpts{Context: ctx}, account.ToEthAddress())
}

func (con *IERC20Watcher) Allowance(ctx context.Context, owner, spender common.Address) (*big.Int, error) {
	return con.IERC20.Allowance(&bind.CallOpts{Context: ctx}, owner.ToEthAddress(), spender.ToEthAddress())
}
