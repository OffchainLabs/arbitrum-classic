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
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/l2message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
)

const maxTransactions = 200

type DecodedBatchTx struct {
	tx     l2message.SignedTransaction
	sender common.Address
}

type Batcher struct {
	rollupAddress common.Address
	globalInbox   arbbridge.GlobalInbox

	sync.Mutex
	valid        bool
	transactions []DecodedBatchTx
}

func NewBatcher(
	ctx context.Context,
	globalInbox arbbridge.GlobalInbox,
	rollupAddress common.Address,
) *Batcher {
	server := &Batcher{
		rollupAddress: rollupAddress,
		globalInbox:   globalInbox,
		valid:         true,
	}

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				server.Lock()
				// Keep sending in spin loop until we can't anymore
				sentFull := false
				for server.valid && len(server.transactions) >= maxTransactions {
					server.sendBatch(ctx)
					sentFull = true
				}
				// If we have've sent any batches, send a partial
				if !sentFull && server.valid && len(server.transactions) > 0 {
					server.sendBatch(ctx)
				}
				server.Unlock()
			}
		}
	}()
	return server
}

// prepareTransactions reorders the transactions such that the position of each
// user is maintained, but the transactions of that user are swapped to be in
// sequence number order
func prepareTransactions(txes []DecodedBatchTx) l2message.TransactionBatch {
	transactionsBySender := make(map[common.Address][]DecodedBatchTx)
	for _, tx := range txes {
		transactionsBySender[tx.sender] = append(transactionsBySender[tx.sender], tx)
	}

	for _, txes := range transactionsBySender {
		sort.SliceStable(txes, func(i, j int) bool {
			return txes[i].tx.Transaction.SequenceNum.Cmp(txes[j].tx.Transaction.SequenceNum) < 0
		})
	}

	batchTxes := make([]l2message.AbstractL2Message, 0, len(txes))
	for _, tx := range txes {
		nextTx := transactionsBySender[tx.sender][0]
		transactionsBySender[tx.sender] = transactionsBySender[tx.sender][1:]
		batchTxes = append(batchTxes, nextTx.tx)
	}
	return l2message.NewTransactionBatchFromMessages(batchTxes)
}

func (m *Batcher) sendBatch(ctx context.Context) {
	var txes []DecodedBatchTx

	if len(m.transactions) > maxTransactions {
		txes = m.transactions[:maxTransactions]
		m.transactions = m.transactions[maxTransactions:]
	} else {
		txes = m.transactions
		m.transactions = nil
	}
	m.Unlock()

	log.Println("Submitting batch with", len(txes), "transactions")

	err := m.globalInbox.SendL2MessageNoWait(
		ctx,
		m.rollupAddress,
		message.L2Message{Data: l2message.L2MessageAsData(prepareTransactions(txes))},
	)

	m.Lock()
	if err != nil {
		log.Println("transaction aggregator failed: ", err)
		m.valid = false
	}
}

// SendTransaction takes a request signed transaction l2message from a client
// and puts it in a queue to be included in the next transaction batch
func (m *Batcher) SendTransaction(tx *types.Transaction) (common.Hash, error) {
	chainId := l2message.ChainAddressToID(m.rollupAddress)
	signer := types.NewEIP155Signer(chainId)
	ethSender, err := signer.Sender(tx)
	if err != nil {
		log.Println("Error processing transaction", err)
		log.Printf("Tx chain id was %v and rollup chain's is %v", tx.ChainId(), chainId)
		return common.Hash{}, err
	}
	sender := common.NewAddressFromEth(ethSender)
	batchTx := l2message.NewSignedTransactionFromEth(tx)

	txHash := tx.Hash()
	log.Println("Got tx: ", batchTx.Transaction, "with hash", txHash.Hex(), "from", sender.Hex())

	m.Lock()
	defer m.Unlock()

	if !m.valid {
		return common.Hash{}, errors.New("tx aggregator is not running")
	}

	m.transactions = append(m.transactions, DecodedBatchTx{
		tx:     batchTx,
		sender: sender,
	})
	return common.NewHashFromEth(txHash), nil
}
