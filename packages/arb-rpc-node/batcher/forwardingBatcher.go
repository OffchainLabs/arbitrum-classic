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

package batcher

import (
	"context"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type Forwarder struct {
	client    *ethclient.Client
	newTxFeed event.Feed
}

func NewForwarder(client *ethclient.Client) *Forwarder {
	return &Forwarder{client: client}
}

// Return nil if no pending transaction count is available
func (b *Forwarder) PendingTransactionCount(ctx context.Context, account common.Address) *uint64 {
	nonce, err := b.client.PendingNonceAt(ctx, account.ToEthAddress())
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Error fetching pending nonce")
		return nil
	}
	return &nonce
}

func (b *Forwarder) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	txes := []*types.Transaction{tx}
	b.newTxFeed.Send(core.NewTxsEvent{Txs: txes})
	return b.client.SendTransaction(ctx, tx)
}

func (b *Forwarder) PendingSnapshot() (*snapshot.Snapshot, error) {
	return nil, nil
}

func (b *Forwarder) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return b.newTxFeed.Subscribe(ch)
}
