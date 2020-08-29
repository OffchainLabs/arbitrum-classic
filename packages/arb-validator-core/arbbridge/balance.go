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

package arbbridge

import (
	"context"
	"errors"
	"log"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func WaitForBalance(ctx context.Context, client ArbClient, tokenAddress common.Address, userAddress common.Address) error {
	emptyAddress := common.Address{}
	if tokenAddress == emptyAddress {
		balance, err := client.GetBalance(ctx, userAddress)
		if err != nil {
			return err
		}
		if balance.Cmp(big.NewInt(0)) > 0 {
			return nil
		}
		log.Println("Waiting for account", userAddress, "to receive ETH")
		timer := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ctx.Done():
				return errors.New("timed out waiting for balance")
			case <-timer.C:
				balance, err := client.GetBalance(ctx, userAddress)
				if err != nil {
					return err
				}
				if balance.Cmp(big.NewInt(0)) > 0 {
					return nil
				}
			}
		}
	} else {
		erc20, err := client.NewIERC20Watcher(tokenAddress)
		if err != nil {
			return err
		}
		balance, err := erc20.BalanceOf(ctx, userAddress)
		if err != nil {
			return err
		}
		if balance.Cmp(big.NewInt(0)) > 0 {
			return nil
		}
		log.Println("Waiting for account", userAddress, "to receive ERC-20 token from contract", tokenAddress)
		timer := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ctx.Done():
				return errors.New("timed out waiting for balance")
			case <-timer.C:
				balance, err := erc20.BalanceOf(ctx, userAddress)
				if err != nil {
					return err
				}
				if balance.Cmp(big.NewInt(0)) > 0 {
					return nil
				}
			}
		}
	}
}
