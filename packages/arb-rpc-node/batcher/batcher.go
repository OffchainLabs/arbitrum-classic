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
	"container/list"
	"context"
	"math/big"
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbtransaction"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/monitor"
	"github.com/offchainlabs/arbitrum/packages/arb-util/transactauth"
)

var logger = log.With().Caller().Stack().Str("component", "batcher").Logger()

const maxBatchSize ethcommon.StorageSize = 120000

type txResponse int

var errFailedBatch = errors.New("submitted of batch failed with revert")

const (
	SKIP = iota
	ACCEPT
	REMOVE
	FULL
)

type l2TxSender interface {
	SendL2MessageFromOrigin(ctx context.Context, data []byte) (*arbtransaction.ArbTransaction, error)
	Sender() common.Address
}

type batch interface {
	newFromExisting() batch
	validateTx(tx *types.Transaction) (txResponse, error)
	isFull() bool
	getAppliedTxes() []*types.Transaction
	addIncludedTx(tx *types.Transaction) error
	updateCurrentSnap(pendingSentBatches *list.List) error
	getLatestSnap() *snapshot.Snapshot
}

type TransactionBatcher interface {
	// Return nil if no pending transaction count is available
	PendingTransactionCount(ctx context.Context, account common.Address) (*uint64, error)

	SendTransaction(ctx context.Context, tx *types.Transaction) error

	SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription

	// Return nil if no pending snapshot is available
	PendingSnapshot() (*snapshot.Snapshot, error)

	Aggregator() *common.Address

	Start(context.Context)
}

type pendingSentBatch struct {
	batchTx *arbtransaction.ArbTransaction
	txes    []*types.Transaction
}

type Batcher struct {
	signer types.Signer
	sender common.Address

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
	receiptFetcher transactauth.TransactAuth,
	globalInbox l2TxSender,
	maxBatchTime time.Duration,
) (*Batcher, error) {
	signer := types.NewEIP155Signer(chainId)
	batch, err := newStatefulBatch(db, maxBatchSize, signer)
	if err != nil {
		return nil, err
	}
	return newBatcher(
		ctx,
		chainId,
		receiptFetcher,
		globalInbox,
		maxBatchTime,
		batch,
	), nil
}

func NewStatelessBatcher(
	ctx context.Context,
	db *txdb.TxDB,
	chainId *big.Int,
	receiptFetcher transactauth.ArbReceiptFetcher,
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
	receiptFetcher transactauth.ArbReceiptFetcher,
	globalInbox l2TxSender,
	maxBatchTime time.Duration,
	pendingBatch batch,
) *Batcher {
	server := &Batcher{
		signer:             types.NewEIP155Signer(chainId),
		sender:             globalInbox.Sender(),
		queuedTxes:         newTxQueues(),
		pendingBatch:       pendingBatch,
		pendingSentBatches: list.New(),
	}

	go func() {
		lastBatch := time.Now()
		checkForFinish := true
		for {
			if checkForFinish {
				select {
				case <-ctx.Done():
					return
				default:
				}
				checkForFinish = false
			}
			server.Lock()
			moreTxesWaiting := server.handleNextTx()
			submittedBatch, err := server.maybeSubmitBatch(ctx, maxBatchTime, lastBatch, globalInbox, moreTxesWaiting)
			if err != nil {
				logger.Error().Err(err).Msg("failed submitting batch")
				time.Sleep(2 * time.Second)
				continue
			}

			if submittedBatch {
				lastBatch = time.Now()
				checkForFinish = true
			}
			if !moreTxesWaiting {
				// If we didn't fill the last batch, pause for more transactions
				time.Sleep(time.Millisecond * 500)
				checkForFinish = true
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
				if server.pendingSentBatches.Len() > 0 {
					if err := server.checkForNextBatch(ctx, receiptFetcher); err != nil {
						log.Error().Err(err).Msg("error checking for submitted batch")
					}
				}
				server.Unlock()
			}
		}
	}()
	return server
}

func (m *Batcher) handleNextTx() bool {
	tx, accountIndex, cont := popRandomTx(m.pendingBatch, m.queuedTxes)
	if tx != nil {
		err := m.pendingBatch.addIncludedTx(tx)
		m.queuedTxes.maybeRemoveAccountAtIndex(accountIndex)
		if err != nil {
			logger.Error().Err(err).Msg("Aggregator ignored invalid tx")
		}
	}
	return cont
}

func (m *Batcher) maybeSubmitBatch(ctx context.Context, maxBatchTime time.Duration, lastBatch time.Time, globalInbox l2TxSender, moreTxesWaiting bool) (bool, error) {
	txes := m.pendingBatch.getAppliedTxes()
	full := m.pendingBatch.isFull()
	m.Unlock()

	if !full && !(len(txes) > 0 && !moreTxesWaiting && time.Since(lastBatch) > maxBatchTime) {
		return false, nil
	}
	lastBatch = time.Now()
	batchTxes := make([]message.AbstractL2Message, 0, len(txes))
	for _, tx := range txes {
		batchTxes = append(batchTxes, message.NewCompressedECDSAFromEth(tx))
	}
	batchTx, err := message.NewTransactionBatchFromMessages(batchTxes)
	if err != nil {
		return false, errors.Wrap(err, "invalid transaction in batch")
	}

	logger.Info().Int("txcount", len(txes)).Msg("Submitting batch")
	batchData := message.NewSafeL2Message(batchTx).AsData()
	tx, err := globalInbox.SendL2MessageFromOrigin(ctx, batchData)
	if err != nil {
		return false, errors.Wrap(err, "error calling SendL2MessageFromOrigin")
	}

	for _, l2tx := range txes {
		monitor.GlobalMonitor.IncludedInBatch(common.NewHashFromEth(l2tx.Hash()), common.NewHashFromEth(l2tx.Hash()))
	}
	monitor.GlobalMonitor.SubmittedBatch(common.NewHashFromEth(tx.Hash()))

	m.Lock()
	m.pendingBatch = m.pendingBatch.newFromExisting()
	m.pendingSentBatches.PushBack(&pendingSentBatch{
		batchTx: tx,
		txes:    txes,
	})
	m.Unlock()
	return true, nil
}

// checkForNextBatch expects the mutex to be held on entry and leaves it unlocked on return
func (m *Batcher) checkForNextBatch(ctx context.Context, receiptFetcher transactauth.ArbReceiptFetcher) error {
	// Note: this is the only place where items can be removed
	// from pendingSentBatches, so pendingSentBatches.Front() is
	// guaranteed not to change when the server lock is released
	if m.pendingSentBatches.Len() == 0 {
		return nil
	}

	batch := m.pendingSentBatches.Front().Value.(*pendingSentBatch)
	m.Unlock()
	receipt, err := transactauth.WaitForReceiptWithResultsSimple(ctx, receiptFetcher, batch.batchTx)
	if err != nil {
		m.Lock()
		return err
	}
	if receipt.Status != 1 {
		// batch failed unexpectedly
		m.Lock()
		return errFailedBatch
	}

	monitor.GlobalMonitor.BatchAccepted(common.NewHashFromEth(receipt.TxHash))

	logger.Info().
		Str("hash", receipt.TxHash.Hex()).
		Uint64("status", receipt.Status).
		Uint64("gasUsed", receipt.GasUsed).
		Str("blockHash", receipt.BlockHash.Hex()).
		Uint64("blockNumber", receipt.BlockNumber.Uint64()).
		Msg("batch receipt")

	// batch succeeded
	m.Lock()
	m.pendingSentBatches.Remove(m.pendingSentBatches.Front())
	return nil
}

func (m *Batcher) PendingSnapshot() (*snapshot.Snapshot, error) {
	m.Lock()
	defer m.Unlock()
	if err := m.pendingBatch.updateCurrentSnap(m.pendingSentBatches); err != nil {
		return nil, err
	}
	return m.pendingBatch.getLatestSnap(), nil
}

func (m *Batcher) PendingTransactionCount(_ context.Context, account common.Address) (*uint64, error) {
	m.Lock()
	defer m.Unlock()
	q, ok := m.queuedTxes.queues[account.ToEthAddress()]
	if !ok {
		return nil, nil
	}
	count := q.maxNonce + 1
	return &count, nil
}

// SendTransaction takes a request signed transaction l2message from a client
// and puts it in a queue to be included in the next transaction batch
func (m *Batcher) SendTransaction(_ context.Context, tx *types.Transaction) error {
	sender, err := types.Sender(m.signer, tx)
	if err != nil {
		logger.Warn().Err(err).Msg("error processing user transaction")
		return err
	}

	monitor.GlobalMonitor.GotTransactionFromUser(common.NewHashFromEth(tx.Hash()))

	m.Lock()
	defer m.Unlock()

	action, err := m.pendingBatch.validateTx(tx)
	if action == REMOVE {
		return err
	}

	if err := m.pendingBatch.updateCurrentSnap(m.pendingSentBatches); err != nil {
		return err
	}

	if err := m.queuedTxes.addTransaction(tx, sender); err != nil {
		return err
	}

	m.newTxFeed.Send(core.NewTxsEvent{Txs: []*types.Transaction{tx}})

	logItem := logger.Info().
		Str("sender", sender.Hex()).
		Uint64("nonce", tx.Nonce()).
		Uint64("gas", tx.Gas()).
		Str("gasPrice", tx.GasPrice().String()).
		Int("calldatasize", len(tx.Data())).
		Str("value", tx.Value().String()).
		Str("hash", tx.Hash().Hex())
	if tx.To() != nil {
		logItem = logItem.Str("dest", tx.To().Hex())
	} else {
		logItem = logItem.Str("dest", "contract-creation")
	}
	logItem.Msg("user tx")
	return nil
}

func (m *Batcher) SubscribeNewTxsEvent(ch chan<- core.NewTxsEvent) event.Subscription {
	return m.newTxFeed.Subscribe(ch)
}

func (m *Batcher) Aggregator() *common.Address {
	return &m.sender
}

func (m *Batcher) Start(_ context.Context) {
}
