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

package aggregator

import (
	"context"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/machineobserver"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

func GenerateRPCServer(server *Server) (*rpc.Server, error) {
	s := rpc.NewServer()
	s.RegisterCodec(
		json.NewCodec(),
		"application/json",
	)
	s.RegisterCodec(
		json.NewCodec(),
		"application/json;charset=UTF-8",
	)

	if err := s.RegisterService(server, "Aggregator"); err != nil {
		return nil, err
	}

	return s, nil
}

func LaunchAggregator(
	ctx context.Context,
	client arbbridge.ArbAuthClient,
	rollupAddress common.Address,
	executable string,
	dbPath string,
	port string,
	flags utils2.RPCFlags,
) error {
	db, err := machineobserver.RunObserver(ctx, rollupAddress, client, executable, dbPath)
	if err != nil {
		return err
	}
	rollupContract, err := client.NewRollupWatcher(rollupAddress)
	if err != nil {
		return err
	}
	inboxAddress, err := rollupContract.InboxAddress(context.Background())
	if err != nil {
		return err
	}
	globalInbox, err := client.NewGlobalInbox(inboxAddress, rollupAddress)
	if err != nil {
		return err
	}
	server := NewServer(ctx, globalInbox, rollupAddress, db)
	s, err := GenerateRPCServer(server)
	if err != nil {
		return err
	}

	return utils2.LaunchRPC(s, port, flags)
}
