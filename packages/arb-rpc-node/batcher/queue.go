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
	"container/heap"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"math/rand"
)

// An TxHeap is a min-heap of transactions sorted by nonce.
type TxHeap []*types.Transaction

func (h TxHeap) Len() int           { return len(h) }
func (h TxHeap) Less(i, j int) bool { return h[i].Nonce() < h[j].Nonce() }
func (h TxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TxHeap) Push(x interface{}) {
	*h = append(*h, x.(*types.Transaction))
}

func (h *TxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type txQueue struct {
	txes        TxHeap
	txesByNonce map[uint64]*types.Transaction
	maxNonce    uint64
}

func newTxQueue() *txQueue {
	return &txQueue{
		txes:        nil,
		txesByNonce: make(map[uint64]*types.Transaction),
		maxNonce:    0,
	}
}

func (q *txQueue) addTransaction(tx *types.Transaction) error {
	if _, ok := q.txesByNonce[tx.Nonce()]; ok {
		return errors.New("transaction replacement not supported")
	}

	q.txesByNonce[tx.Nonce()] = tx
	heap.Push(&q.txes, tx)

	if tx.Nonce() > q.maxNonce {
		q.maxNonce = tx.Nonce()
	}
	return nil
}

func (q *txQueue) Empty() bool {
	return len(q.txes) == 0
}

func (q *txQueue) Peek() *types.Transaction {
	if q == nil || len(q.txes) == 0 {
		return nil
	}
	return q.txes[0]
}

func (q *txQueue) Pop() *types.Transaction {
	tx := heap.Pop(&q.txes).(*types.Transaction)
	delete(q.txesByNonce, tx.Nonce())
	return tx
}

type txQueues struct {
	queues   map[common.Address]*txQueue
	accounts []common.Address
}

func newTxQueues() *txQueues {
	return &txQueues{
		queues:   make(map[common.Address]*txQueue),
		accounts: nil,
	}
}

func (q *txQueues) addTransaction(tx *types.Transaction, sender common.Address) error {
	queue, ok := q.queues[sender]
	if !ok {
		queue = newTxQueue()
		q.queues[sender] = queue
		q.accounts = append(q.accounts, sender)
	}
	return queue.addTransaction(tx)
}

func (q *txQueues) removeTxFromAccountAtIndex(i int) {
	q.queues[q.accounts[i]].Pop()
}

func (q *txQueues) maybeRemoveAccountAtIndex(i int) {
	account := q.accounts[i]
	if q.queues[account].Empty() {
		delete(q.queues, account)
		q.accounts[i] = q.accounts[len(q.accounts)-1]
		q.accounts = q.accounts[:len(q.accounts)-1]
	}
}

func popRandomTx(b batch, queuedTxes *txQueues) (*types.Transaction, int, bool) {
	queuedCount := int32(len(queuedTxes.accounts))
	if queuedCount == 0 {
		return nil, 0, false
	}
	index := int(rand.Int31n(queuedCount))
	first := true
	lastIndex := index
	index--
	for {
		if len(queuedTxes.accounts) == 0 {
			return nil, 0, false
		}
		index++
		if index == len(queuedTxes.accounts) {
			index = 0
		}
		if !first && index == lastIndex {
			return nil, 0, false
		}

		first = false
		account := queuedTxes.accounts[index]
		nextAccount := queuedTxes.queues[account]
		tx := nextAccount.Peek()
		// No tx in account
		if tx == nil {
			queuedTxes.maybeRemoveAccountAtIndex(index)
			continue
		}

		// err param can be ignored
		action, _ := b.validateTx(tx)
		switch action {
		case REMOVE:
			queuedTxes.removeTxFromAccountAtIndex(index)
		case SKIP:
		case FULL:
			return nil, 0, true
		case ACCEPT:
			queuedTxes.removeTxFromAccountAtIndex(index)
			return tx, index, true
		}
	}
}
