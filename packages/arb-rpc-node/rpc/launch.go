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
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/staker"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/batcher"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	utils2 "github.com/offchainlabs/arbitrum/packages/arb-rpc-node/utils"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
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
	InboxReader                *staker.InboxReader
	DelayedMessagesTargetDelay *big.Int
}

func (b SequencerBatcherMode) isBatcherMode() {}

func LaunchNode(
	ctx context.Context,
	client ethutils.EthClient,
	rollupAddress common.Address,
	db *txdb.TxDB,
	web3RPCPort string,
	web3WSPort string,
	flags utils2.RPCFlags,
	maxBatchTime time.Duration,
	batcherMode BatcherMode,
) error {
	l2ChainID := message.ChainAddressToID(rollupAddress)

	var batch batcher.TransactionBatcher
	switch batcherMode := batcherMode.(type) {
	case ForwarderBatcherMode:
		var err error
		batch, err = batcher.NewForwarder(ctx, batcherMode.NodeURL)
		if err != nil {
			return err
		}
	case StatelessBatcherMode:
		auth, err := ethbridge.NewTransactAuth(ctx, client, batcherMode.Auth)
		if err != nil {
			return err
		}
		inbox, err := ethbridge.NewStandardInbox(batcherMode.InboxAddress.ToEthAddress(), client, auth)
		if err != nil {
			return err
		}
		batch = batcher.NewStatelessBatcher(ctx, db, l2ChainID, client, inbox, maxBatchTime)
	case StatefulBatcherMode:
		auth, err := ethbridge.NewTransactAuth(ctx, client, batcherMode.Auth)
		if err != nil {
			return err
		}
		inbox, err := ethbridge.NewStandardInbox(batcherMode.InboxAddress.ToEthAddress(), client, auth)
		if err != nil {
			return err
		}
		batch, err = batcher.NewStatefulBatcher(ctx, db, l2ChainID, client, inbox, maxBatchTime)
		if err != nil {
			return err
		}
	case SequencerBatcherMode:
		rollup, err := ethbridgecontracts.NewRollup(rollupAddress.ToEthAddress(), client)
		if err != nil {
			return err
		}
		callOpts := &bind.CallOpts{Context: ctx}
		seqInboxAddr, err := rollup.SequencerBridge(callOpts)
		if err != nil {
			return err
		}
		seqInbox, err := ethbridgecontracts.NewSequencerInbox(seqInboxAddr, client)
		if err != nil {
			return err
		}
		batch, err = batcher.NewSequencerBatcher(ctx, batcherMode.Core, batcherMode.InboxReader, client, batcherMode.DelayedMessagesTargetDelay, seqInbox, batcherMode.Auth)
		if err != nil {
			return err
		}
	default:
		return errors.New("unexpected batcher type")
	}

	return LaunchNodeAdvanced(
		db,
		rollupAddress,
		web3RPCPort,
		web3WSPort,
		flags,
		batch,
		nil,
		false,
		make(map[string]interface{}),
	)
}

func LaunchNodeAdvanced(
	db *txdb.TxDB,
	rollupAddress common.Address,
	web3RPCPort string,
	web3WSPort string,
	flags utils2.RPCFlags,
	batch batcher.TransactionBatcher,
	privateKeys []*ecdsa.PrivateKey,
	ganacheMode bool,
	plugins map[string]interface{},
) error {
	srv := aggregator.NewServer(batch, rollupAddress, db)
	errChan := make(chan error, 1)

	web3Server, err := web3.GenerateWeb3Server(srv, privateKeys, ganacheMode, plugins)
	if err != nil {
		return err
	}

	if web3RPCPort != "" {
		go func() {
			errChan <- utils2.LaunchRPC(web3Server, web3RPCPort, flags)
		}()
	}
	if web3WSPort != "" {
		go func() {
			errChan <- utils2.LaunchWS(web3Server, web3WSPort, flags)
		}()
	}

	return <-errChan
}
