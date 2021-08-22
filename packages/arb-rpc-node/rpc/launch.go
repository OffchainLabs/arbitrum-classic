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
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks"
	"math/big"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-util/broadcaster"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-rpc-node/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
)

type BatcherMode interface {
	isBatcherMode()
}

type ForwarderBatcherMode struct {
	Config configuration.Forwarder
}

func (b ForwarderBatcherMode) isBatcherMode() {}

type StatefulBatcherMode struct {
	Auth         *bind.TransactOpts
	InboxAddress common.Address
}

func (b StatefulBatcherMode) isBatcherMode() {}

type StatelessBatcherMode struct {
	Auth         *bind.TransactOpts
	InboxAddress common.Address
}

func (b StatelessBatcherMode) isBatcherMode() {}

type SequencerBatcherMode struct {
	Auth        *bind.TransactOpts
	Core        core.ArbCore
	InboxReader *monitor.InboxReader
}

func (b SequencerBatcherMode) isBatcherMode() {}

func SetupBatcher(
	ctx context.Context,
	client ethutils.EthClient,
	rollupAddress common.Address,
	l2ChainId *big.Int,
	db *txdb.TxDB,
	maxBatchTime time.Duration,
	batcherMode BatcherMode,
	dataSigner func([]byte) ([]byte, error),
	config *configuration.Config,
	walletConfig *configuration.Wallet,
) (batcher.TransactionBatcher, error) {
	switch batcherMode := batcherMode.(type) {
	case ForwarderBatcherMode:
		return batcher.NewForwarder(ctx, batcherMode.Config)
	case StatelessBatcherMode:
		var auth *ethbridge.TransactAuth
		var fb *fireblocks.Fireblocks
		var err error
		if len(config.Wallet.FireblocksSSLKey) > 0 {
			auth, fb, err = ethbridge.NewFireblocksTransactAuth(ctx, client, batcherMode.Auth, config, walletConfig)
		} else {
			auth, err = ethbridge.NewTransactAuth(ctx, client, batcherMode.Auth, config, walletConfig)

		}
		if err != nil {
			return nil, err
		}
		inbox, err := ethbridge.NewStandardInbox(batcherMode.InboxAddress.ToEthAddress(), client, auth)
		if err != nil {
			return nil, err
		}
		return batcher.NewStatelessBatcher(ctx, db, l2ChainId, client, inbox, maxBatchTime, fb), nil
	case StatefulBatcherMode:
		var auth *ethbridge.TransactAuth
		var fb *fireblocks.Fireblocks
		var err error
		if len(config.Wallet.FireblocksSSLKey) > 0 {
			auth, fb, err = ethbridge.NewFireblocksTransactAuth(ctx, client, batcherMode.Auth, config, walletConfig)
		} else {
			auth, err = ethbridge.NewTransactAuth(ctx, client, batcherMode.Auth, config, walletConfig)

		}
		if err != nil {
			return nil, err
		}
		inbox, err := ethbridge.NewStandardInbox(batcherMode.InboxAddress.ToEthAddress(), client, auth)
		if err != nil {
			return nil, err
		}
		return batcher.NewStatefulBatcher(ctx, db, l2ChainId, client, inbox, maxBatchTime, fb)
	case SequencerBatcherMode:
		rollup, err := ethbridgecontracts.NewRollupUserFacet(rollupAddress.ToEthAddress(), client)
		if err != nil {
			return nil, err
		}
		callOpts := &bind.CallOpts{Context: ctx}
		seqInboxAddr, err := rollup.SequencerBridge(callOpts)
		if err != nil {
			return nil, err
		}
		seqInbox, err := ethbridgecontracts.NewSequencerInbox(seqInboxAddr, client)
		if err != nil {
			return nil, err
		}
		feedBroadcaster := broadcaster.NewBroadcaster(config.Feed.Output)
		seqBatcher, err := batcher.NewSequencerBatcher(
			ctx,
			batcherMode.Core,
			l2ChainId,
			batcherMode.InboxReader,
			client,
			seqInbox,
			batcherMode.Auth,
			dataSigner,
			feedBroadcaster,
			config,
			walletConfig)
		if err != nil {
			return nil, err
		}

		err = feedBroadcaster.Start(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error starting feed broadcaster")
		}
		return seqBatcher, nil
	default:
		return nil, errors.New("unexpected batcher type")
	}
}

func LaunchPublicServer(ctx context.Context, web3Server *rpc.Server, rpc configuration.RPC, ws configuration.WS) error {
	if rpc.Port == ws.Port && rpc.Port != "" {
		if rpc.Addr != ws.Addr {
			return errors.New("if serving on same port, rpc and ws addreses must be the same")
		}
		if rpc.Path == ws.Path {
			return errors.New("if serving on same port, ws and rpc path must be different")
		}
		return utils2.LaunchRPCAndWS(ctx, web3Server, rpc.Addr, rpc.Port, rpc.Path, ws.Path)
	}

	errChan := make(chan error, 1)
	if rpc.Port != "" {
		go func() {
			errChan <- utils2.LaunchRPC(ctx, web3Server, rpc.Addr, rpc.Port, rpc.Path)
		}()
	}
	if ws.Port != "" {
		go func() {
			errChan <- utils2.LaunchWS(ctx, web3Server, ws.Addr, ws.Port, ws.Path)
		}()
	}
	return <-errChan
}
