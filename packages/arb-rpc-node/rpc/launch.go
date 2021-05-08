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
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-rpc-node/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
)

type BatcherMode interface {
	isBatcherMode()
}

type ForwarderBatcherMode struct {
	NodeURL string
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
	Auth                       *bind.TransactOpts
	Core                       core.ArbCore
	InboxReader                *monitor.InboxReader
	DelayedMessagesTargetDelay *big.Int
}

func (b SequencerBatcherMode) isBatcherMode() {}

func SetupBatcher(
	ctx context.Context,
	client ethutils.EthClient,
	rollupAddress common.Address,
	db *txdb.TxDB,
	maxBatchTime time.Duration,
	batcherMode BatcherMode,
) (batcher.TransactionBatcher, error) {
	l2ChainID := message.ChainAddressToID(rollupAddress)
	switch batcherMode := batcherMode.(type) {
	case ForwarderBatcherMode:
		return batcher.NewForwarder(ctx, batcherMode.NodeURL)
	case StatelessBatcherMode:
		auth, err := ethbridge.NewTransactAuth(ctx, client, batcherMode.Auth)
		if err != nil {
			return nil, err
		}
		inbox, err := ethbridge.NewStandardInbox(batcherMode.InboxAddress.ToEthAddress(), client, auth)
		if err != nil {
			return nil, err
		}
		return batcher.NewStatelessBatcher(ctx, db, l2ChainID, client, inbox, maxBatchTime), nil
	case StatefulBatcherMode:
		auth, err := ethbridge.NewTransactAuth(ctx, client, batcherMode.Auth)
		if err != nil {
			return nil, err
		}
		inbox, err := ethbridge.NewStandardInbox(batcherMode.InboxAddress.ToEthAddress(), client, auth)
		if err != nil {
			return nil, err
		}
		return batcher.NewStatefulBatcher(ctx, db, l2ChainID, client, inbox, maxBatchTime)
	case SequencerBatcherMode:
		rollup, err := ethbridgecontracts.NewRollup(rollupAddress.ToEthAddress(), client)
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
		seqBatcher, err := batcher.NewSequencerBatcher(ctx, batcherMode.Core, l2ChainID, batcherMode.InboxReader, client, batcherMode.DelayedMessagesTargetDelay, seqInbox, batcherMode.Auth)
		if err != nil {
			return nil, err
		}
		go seqBatcher.Start(ctx)
		return seqBatcher, nil
	default:
		return nil, errors.New("unexpected batcher type")
	}
}

func LaunchPublicServer(ctx context.Context, web3Server *rpc.Server, web3RPCPort string, web3WSPort string) error {
	errChan := make(chan error, 1)
	if web3RPCPort != "" {
		go func() {
			errChan <- utils2.LaunchRPC(ctx, web3Server, web3RPCPort)
		}()
	}
	if web3WSPort != "" {
		go func() {
			errChan <- utils2.LaunchWS(ctx, web3Server, web3WSPort)
		}()
	}
	return <-errChan
}
