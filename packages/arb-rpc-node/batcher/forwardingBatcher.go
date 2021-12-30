/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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

package batcher

import (
	"context"
	"encoding/json"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
)

type Forwarder struct {
	client     *ethclient.Client
	aggregator *common.Address
}

type AggregatorInfo struct {
	Address *ethcommon.Address `json:"address"`
}

func NewForwarder(ctx context.Context, config configuration.Forwarder) (*Forwarder, error) {
	client, err := ethclient.DialContext(ctx, config.Target)
	if err != nil {
		return nil, err
	}

	var agg *common.Address
	if config.Submitter != "" {
		tmp := common.HexToAddress(config.Submitter)
		agg = &tmp
	} else {
		rpcClient, err := rpc.DialContext(ctx, config.Target)
		if err != nil {
			return nil, err
		}
		var raw json.RawMessage
		if err := rpcClient.CallContext(ctx, &raw, "arb_getAggregator"); err != nil {
			return nil, err
		}
		if len(raw) == 0 {
			return nil, ethereum.NotFound
		}
		var ret AggregatorInfo
		if err := json.Unmarshal(raw, &ret); err != nil {
			return nil, err
		}
		if ret.Address != nil {
			tmp := common.NewAddressFromEth(*ret.Address)
			agg = &tmp
		}
	}

	return &Forwarder{client: client, aggregator: agg}, nil
}

// Return nil if no pending transaction count is available
func (b *Forwarder) PendingTransactionCount(ctx context.Context, account common.Address) (*uint64, error) {
	nonce, err := b.client.PendingNonceAt(ctx, account.ToEthAddress())
	if err != nil {
		return nil, errors.Wrap(err, "error fetching pending nonce from forwarding target")
	}
	return &nonce, nil
}

func (b *Forwarder) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	logger.Info().Str("hash", tx.Hash().String()).Msg("got user tx")
	return b.client.SendTransaction(ctx, tx)
}

func (b *Forwarder) PendingSnapshot(_ context.Context) (*snapshot.Snapshot, error) {
	return nil, nil
}

func (b *Forwarder) Aggregator() *common.Address {
	return b.aggregator
}

func (m *Forwarder) Start(ctx context.Context) {
}
