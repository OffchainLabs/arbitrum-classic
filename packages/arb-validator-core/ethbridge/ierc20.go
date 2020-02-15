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

	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge/ierc20"

	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

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
