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
	"errors"
	"log"
	"sort"
	"sync"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
)

const maxTransactions = 200

type TransactionBatcher interface {
	PendingTransactionCount(account common.Address) *uint64
	SendTransaction(tx *types.Transaction) (common.Hash, error)
	PendingSnapshot() (*snapshot.Snapshot, bool)
}

type Batcher struct {
	signer      types.Signer
	client      ethutils.EthClient
	globalInbox arbbridge.GlobalInbox

	db *txdb.TxDB

	sync.Mutex
	valid bool

	queuedTxes      *txQueues
	pendingSnapshot *snapshot.Snapshot
	pendingBatch    []*types.Transaction
}

func NewBatcher(
	ctx context.Context,
	db *txdb.TxDB,
	rollupAddress common.Address,
	client ethutils.EthClient,
	globalInbox arbbridge.GlobalInbox,
	maxBatchTime time.Duration,
) *Batcher {
	server := &Batcher{
		signer:      types.NewEIP155Signer(message.ChainAddressToID(rollupAddress)),
		client:      client,
		globalInbox: globalInbox,
		db:          db,
		valid:       true,
		queuedTxes:  newTxQueues(),
	}

	go func() {
		ticker := time.NewTicker(maxBatchTime)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				server.Lock()
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
				// Keep sending in spin loop until we can't anymore

				//sentFull := false
				//for server.valid && len(server.transactions) >= maxTransactions {
				//	server.sendBatch(ctx)
				//	sentFull = true
				//}
				//// If we have've sent any batches, send a partial
				//if !sentFull && server.valid && len(server.transactions) > 0 {
				//	server.sendBatch(ctx)
				//}
				server.Unlock()
			}
		}
	}()
	return server
}

// prepareTransactions reorders the transactions such that the position of each
// user is maintained, but the transactions of that user are swapped to be in
// sequence number order
func prepareTransactions(signer types.Signer, txes []*types.Transaction) []*types.Transaction {
	transactionsBySender := make(map[ethcommon.Address][]*types.Transaction)
	for _, tx := range txes {
		sender, _ := types.Sender(signer, tx)
		transactionsBySender[sender] = append(transactionsBySender[sender], tx)
	}

	for _, txes := range transactionsBySender {
		sort.Slice(txes, func(i, j int) bool {
			return txes[i].Nonce() < txes[j].Nonce()
		})
	}

	batchTxes := make([]*types.Transaction, 0, len(txes))
	for _, tx := range txes {
		sender, _ := types.Sender(signer, tx)
		nextTx := transactionsBySender[sender][0]
		transactionsBySender[sender] = transactionsBySender[sender][1:]
		batchTxes = append(batchTxes, nextTx)
	}
	return batchTxes
}

func (m *Batcher) sendBatch(ctx context.Context) {
	var txes []*types.Transaction

	//if len(m.transactions) > maxTransactions {
	//	txes = m.transactions[:maxTransactions]
	//	m.transactions = m.transactions[maxTransactions:]
	//} else {
	//	txes = m.transactions
	//	m.transactions = nil
	//}
	m.Unlock()

	log.Println("Submitting batch with", len(txes), "transactions")

	batch := prepareTransactions(m.signer, txes)

	batchTxes := make([]message.AbstractL2Message, 0, len(batch))
	for _, tx := range batch {
		batchTxes = append(batchTxes, message.SignedTransaction{Tx: tx})
	}
	batchTx, err := message.NewTransactionBatchFromMessages(batchTxes)
	if err != nil {
		log.Println("transaction aggregator failed: ", err)
		m.valid = false
		return
	}
	txHash, err := m.globalInbox.SendL2MessageNoWait(
		ctx,
		message.NewSafeL2Message(batchTx).AsData(),
	)

	m.Lock()
	if err != nil {
		log.Println("transaction aggregator failed: ", err)
		m.valid = false
		return
	}

	go func() {
		receipt, err := ethbridge.WaitForReceiptWithResultsSimple(ctx, m.client, txHash.ToEthHash())
		if err != nil || receipt.Status != 1 {
			// batch failed
			log.Fatal("Error submitted batch", err)
		} else {
			m.Lock()
			defer m.Unlock()
			// batch succeeded
			for _, tx := range batch {
				_, _ = types.Sender(m.signer, tx)
				//m.pendingTxes[sender]--
			}
		}
	}()
}

func (m *Batcher) PendingTransactionCount(account common.Address) *uint64 {
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
func (m *Batcher) SendTransaction(tx *types.Transaction) (common.Hash, error) {
	ethSender, err := types.Sender(m.signer, tx)
	if err != nil {
		log.Println("Error processing transaction", err)
		return common.Hash{}, err
	}

	txHash := common.NewHashFromEth(tx.Hash())
	log.Println("Got tx: with hash", txHash, "from", ethSender.Hex())

	m.Lock()
	defer m.Unlock()

	if !m.valid {
		return common.Hash{}, errors.New("tx aggregator is not running")
	}

	if err := m.validateTxAgainstPending(tx); err != nil {
		return common.Hash{}, err
	}

	if err := m.queuedTxes.addTransaction(tx); err != nil {
		return common.Hash{}, err
	}

	return txHash, nil
}

func (m *Batcher) validateTxAgainstPending(tx *types.Transaction) error {
	ethSender, _ := types.Sender(m.signer, tx)
	sender := common.NewAddressFromEth(ethSender)
	txCount, err := m.pendingSnapshot.GetTransactionCount(sender)
	if err != nil {
		return err
	}

	if tx.Nonce() < txCount.Uint64() {
		return core.ErrNonceTooLow
	}

	amount, err := m.pendingSnapshot.GetBalance(sender)
	if err != nil {
		return err
	}

	if tx.Cost().Cmp(amount) < 0 {
		return core.ErrInsufficientFunds
	}
	return nil
}

func (m *Batcher) addTxToBatch(tx *types.Transaction) error {
	ethSender, _ := types.Sender(m.signer, tx)
	sender := common.NewAddressFromEth(ethSender)
	msg, err := message.NewL2Message(message.SignedTransaction{Tx: tx})
	if err != nil {
		return err
	}
	newSnap, _, err := snapshot.NewSnapshotWithMessage(m.pendingSnapshot, msg, sender)
	if err != nil {
		return err
	}
	m.pendingSnapshot = newSnap
	m.pendingBatch = append(m.pendingBatch, tx)
	return nil
}

func (m *Batcher) PendingSnapshot() (*snapshot.Snapshot, bool) {
	m.setupPending()
	latest := m.db.LatestSnapshot()
	if m.pendingSnapshot == nil || m.pendingSnapshot.Height().Cmp(latest.Height()) < 0 {
		m.pendingSnapshot = latest
		return latest, true
	}
	return m.pendingSnapshot, false
}

func (m *Batcher) setupPending() bool {
	latest := m.db.LatestSnapshot()
	if m.pendingSnapshot == nil || m.pendingSnapshot.Height().Cmp(latest.Height()) < 0 {
		m.pendingSnapshot = latest
		return true
	}
	return false
}
