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

package ethbridge

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgetestcontracts"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func WaitForBalance(ctx context.Context, client ethutils.EthClient, tokenAddress common.Address, userAddress common.Address) error {
	emptyAddress := common.Address{}
	if tokenAddress == emptyAddress {
		balance, err := client.BalanceAt(ctx, userAddress.ToEthAddress(), nil)
		if err != nil {
			return errors.WithStack(err)
		}
		if balance.Cmp(big.NewInt(0)) > 0 {
			return nil
		}
		logger.Info().Hex("account", userAddress.Bytes()).Msg("Waiting for account to receive ETH")
		timer := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ctx.Done():
				return errors.New("timed out waiting for balance")
			case <-timer.C:
				balance, err := client.BalanceAt(ctx, userAddress.ToEthAddress(), nil)
				if err != nil {
					return errors.WithStack(err)
				}
				if balance.Cmp(big.NewInt(0)) > 0 {
					return nil
				}
			}
		}
	} else {
		erc20, err := ethbridgetestcontracts.NewIERC20(tokenAddress.ToEthAddress(), client)
		if err != nil {
			return errors.WithStack(err)
		}
		balance, err := erc20.BalanceOf(&bind.CallOpts{Context: ctx}, userAddress.ToEthAddress())
		if err != nil {
			return errors.WithStack(err)
		}
		if balance.Cmp(big.NewInt(0)) > 0 {
			return nil
		}
		logger.Info().
			Hex("account", userAddress.Bytes()).
			Hex("contract", tokenAddress.Bytes()).
			Msg("Waiting for account to receive ERC-20 token from contract")
		timer := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ctx.Done():
				return errors.New("timed out waiting for balance")
			case <-timer.C:
				balance, err := erc20.BalanceOf(&bind.CallOpts{Context: ctx}, userAddress.ToEthAddress())
				if err != nil {
					return errors.WithStack(err)
				}
				if balance.Cmp(big.NewInt(0)) > 0 {
					return nil
				}
			}
		}
	}
}
