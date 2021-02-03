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
	"container/list"
	"context"
	"github.com/pkg/errors"
	"math/big"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var logger = log.With().Caller().Str("component", "batcher").Logger()

const maxBatchSize ethcommon.StorageSize = 120000

type txResponse int

const (
	SKIP = iota
	ACCEPT
	REMOVE
	FULL
)

type l2TxSender interface {
	SendL2MessageFromOrigin(ctx context.Context, data []byte) (common.Hash, error)
}

type batch interface {
	newFromExisting() batch
	validateTx(tx *types.Transaction) (txResponse, error)
	isFull() bool
	getAppliedTxes() []*types.Transaction
	addIncludedTx(tx *types.Transaction) error
	updateCurrentSnap(pendingSentBatches *list.List)
	getLatestSnap() *snapshot.Snapshot
}

type TransactionBatcher interface {
	// Return nil if no pending transaction count is available
	PendingTransactionCount(ctx context.Context, account common.Address) *uint64

	SendTransaction(ctx context.Context, tx *types.Transaction) error

	SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription

	// Return nil if no pending snapshot is available
	PendingSnapshot() *snapshot.Snapshot
}

type pendingSentBatch struct {
	txHash common.Hash
	txes   []*types.Transaction
}

type Batcher struct {
	signer types.Signer

	sync.Mutex

	queuedTxes         *txQueues
	pendingBatch       batch
	pendingSentBatches *list.List
	newTxFeed          event.Feed
}

func NewStatefulBatcher(
	ctx context.Context,
	db *txdb.TxDB,
	chainId *big.Int,
	receiptFetcher ethutils.ReceiptFetcher,
	globalInbox l2TxSender,
	maxBatchTime time.Duration,
) *Batcher {
	signer := types.NewEIP155Signer(chainId)
	return newBatcher(
		ctx,
		chainId,
		receiptFetcher,
		globalInbox,
		maxBatchTime,
		newStatefulBatch(db, maxBatchSize, signer),
	)
}

func NewStatelessBatcher(
	ctx context.Context,
	db *txdb.TxDB,
	chainId *big.Int,
	receiptFetcher ethutils.ReceiptFetcher,
	globalInbox l2TxSender,
	maxBatchTime time.Duration,
) *Batcher {
	signer := types.NewEIP155Signer(chainId)
	return newBatcher(
		ctx,
		chainId,
		receiptFetcher,
		globalInbox,
		maxBatchTime,
		newStatelessBatch(db, maxBatchSize, signer),
	)
}

func newBatcher(
	ctx context.Context,
	chainId *big.Int,
	receiptFetcher ethutils.ReceiptFetcher,
	globalInbox l2TxSender,
	maxBatchTime time.Duration,
	pendingBatch batch,
) *Batcher {
	server := &Batcher{
		signer:             types.NewEIP155Signer(chainId),
		queuedTxes:         newTxQueues(),
		pendingBatch:       pendingBatch,
		pendingSentBatches: list.New(),
	}

	go func() {
		lastBatch := time.Now()
		ticker := time.NewTicker(time.Millisecond * 500)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				server.Lock()
				for {
					tx, accountIndex, cont := popRandomTx(server.pendingBatch, server.queuedTxes)
					if tx != nil {
						err := server.pendingBatch.addIncludedTx(tx)
						server.queuedTxes.maybeRemoveAccountAtIndex(accountIndex)
						if err != nil {
							logger.Error().Stack().Err(err).Msg("Aggregator ignored invalid tx")
							continue
						}
					}
					if server.pendingBatch.isFull() || (!cont && time.Since(lastBatch) > maxBatchTime) {
						lastBatch = time.Now()
						server.sendBatch(ctx, globalInbox)
					}

					if !cont {
						// If we didn't fill the last batch, pause for more transactions
						server.Unlock()
						break
					}
				}

			}
		}
	}()

	go func() {
		ticker := time.NewTicker(maxBatchTime)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				server.Lock()
				// Note: this loop is the only place where items can be removed
				// from pendingSentBatches, so pendingSentBatches.Front() is
				// guaranteed not to change when the server lock is released
				for server.pendingSentBatches.Len() > 0 {
					batch := server.pendingSentBatches.Front().Value.(*pendingSentBatch)
					txHash := batch.txHash.ToEthHash()
					server.Unlock()
					receipt, err := ethbridge.WaitForReceiptWithResultsSimple(ctx, receiptFetcher, txHash)
					if err != nil || receipt.Status != 1 {
						// batch failed
						logger.Fatal().Stack().Err(err).Msg("Error submitted batch")
					}

					receiptJSON, err := receipt.MarshalJSON()
					if err != nil {
						logger.Error().Stack().Err(err).Msg("failed to generate json for receipt")
					} else {
						logger.Info().RawJSON("receipt", receiptJSON).Msg("batch receipt")
					}

					// batch succeeded
					server.Lock()
					server.pendingSentBatches.Remove(server.pendingSentBatches.Front())
				}
				server.Unlock()
			}
		}
	}()
	return server
}

func (m *Batcher) sendBatch(ctx context.Context, inbox l2TxSender) {
	txes := m.pendingBatch.getAppliedTxes()
	if len(txes) == 0 {
		return
	}
	batchTxes := make([]message.AbstractL2Message, 0, len(txes))
	for _, tx := range txes {
		batchTxes = append(batchTxes, message.NewCompressedECDSAFromEth(tx))
	}
	batchTx, err := message.NewTransactionBatchFromMessages(batchTxes)
	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("transaction aggregator failed")
	}
	logger.Info().Int("txcount", len(batchTxes)).Msg("Submitting batch")
	txHash, err := inbox.SendL2MessageFromOrigin(
		ctx,
		message.NewSafeL2Message(batchTx).AsData(),
	)

	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("transaction aggregator failed")
		return
	}

	m.pendingBatch = m.pendingBatch.newFromExisting()
	m.pendingSentBatches.PushBack(&pendingSentBatch{
		txHash: txHash,
		txes:   txes,
	})
}

func (m *Batcher) PendingSnapshot() *snapshot.Snapshot {
	m.Lock()
	defer m.Unlock()
	m.pendingBatch.updateCurrentSnap(m.pendingSentBatches)
	return m.pendingBatch.getLatestSnap()
}

func (m *Batcher) PendingTransactionCount(_ context.Context, account common.Address) *uint64 {
	m.Lock()
	defer m.Unlock()
	q, ok := m.queuedTxes.queues[account.ToEthAddress()]
	if !ok {
		return nil
	}
	count := q.maxNonce + 1
	return &count
}

// SendTransaction takes a request signed transaction l2message from a client
// and puts it in a queue to be included in the next transaction batch
func (m *Batcher) SendTransaction(_ context.Context, tx *types.Transaction) error {
	sender, err := types.Sender(m.signer, tx)
	if err != nil {
		logger.Error().Stack().Err(err).Msg("error processing user transaction")
		return err
	}

	m.Lock()
	defer m.Unlock()

	action, err := m.pendingBatch.validateTx(tx)
	if action == REMOVE {
		return errors.Wrap(err, "transaction rejected")
	}

	m.pendingBatch.updateCurrentSnap(m.pendingSentBatches)

	if err := m.queuedTxes.addTransaction(tx, sender); err != nil {
		return err
	}

	m.newTxFeed.Send(core.NewTxsEvent{Txs: []*types.Transaction{tx}})

	txJSON, err := tx.MarshalJSON()
	if err != nil {
		logger.Error().Stack().Err(err).Msg("failed to marshal tx into json")
	} else {
		logger.Info().RawJSON("tx", txJSON).Hex("sender", sender.Bytes()).Msg("user tx")
	}
	return nil
}

func (m *Batcher) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return m.newTxFeed.Subscribe(ch)
}
