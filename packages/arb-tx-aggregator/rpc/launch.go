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

package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/machineobserver"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"time"
)

func LaunchAggregator(
	ctx context.Context,
	client ethutils.EthClient,
	auth *bind.TransactOpts,
	rollupAddress common.Address,
	executable string,
	dbPath string,
	aggPort string,
	web3Port string,
	flags utils2.RPCFlags,
	maxBatchTime time.Duration,
) error {
	arbClient := ethbridge.NewEthClient(client)
	db, err := machineobserver.RunObserver(ctx, rollupAddress, arbClient, executable, dbPath)
	if err != nil {
		return err
	}

	authClient := ethbridge.NewEthAuthClient(client, auth)
	rollupContract, err := arbClient.NewRollupWatcher(rollupAddress)
	if err != nil {
		return err
	}
	inboxAddress, err := rollupContract.InboxAddress(context.Background())
	if err != nil {
		return err
	}
	globalInbox, err := authClient.NewGlobalInbox(inboxAddress, rollupAddress)
	if err != nil {
		return err
	}

	batch := batcher.NewBatcher(ctx, db, rollupAddress, client, globalInbox, maxBatchTime)

	srv := aggregator.NewServer(client, batch, rollupAddress, db)
	errChan := make(chan error, 1)

	aggServer, err := aggregator.GenerateRPCServer(srv)
	if err != nil {
		return err
	}

	web3Server, err := web3.GenerateWeb3Server(srv)
	if err != nil {
		return err
	}

	if aggPort != "" {
		go func() {
			errChan <- utils2.LaunchRPC(aggServer, aggPort, flags)
		}()
	}
	if web3Port != "" {
		go func() {
			errChan <- utils2.LaunchRPC(web3Server, web3Port, flags)
		}()
	}

	return <-errChan
}
