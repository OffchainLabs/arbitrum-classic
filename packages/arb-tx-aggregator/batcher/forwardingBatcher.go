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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"log"
)

type Forwarder struct {
	client *ethclient.Client
}

func NewForwarder(client *ethclient.Client) *Forwarder {
	return &Forwarder{client: client}
}

// Return nil if no pending transaction count is available
func (b *Forwarder) PendingTransactionCount(ctx context.Context, account common.Address) *uint64 {
	nonce, err := b.client.PendingNonceAt(ctx, account.ToEthAddress())
	if err != nil {
		log.Println("Error fetching pending nice")
		return nil
	}
	return &nonce
}

func (b *Forwarder) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return b.client.SendTransaction(ctx, tx)
}

func (b *Forwarder) PendingSnapshot() *snapshot.Snapshot {
	return nil
}
