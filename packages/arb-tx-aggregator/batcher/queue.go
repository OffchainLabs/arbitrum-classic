package batcher

import (
	"container/heap"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
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

func (q *txQueues) addTransaction(tx *types.Transaction, signer types.Signer) error {
	sender, _ := types.Sender(signer, tx)
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

func (q *txQueues) getRandomTx() (*types.Transaction, int) {
	index := int(rand.Int31n(int32(len(q.accounts))))
	nextAccount := q.queues[q.accounts[index]]
	tx := nextAccount.Peek()
	return tx, index
}

type pendingBatch struct {
	snap        *snapshot.Snapshot
	txCounts    map[common.Address]uint64
	appliedTxes []*types.Transaction
	sizeBytes   common.StorageSize
	maxSize     common.StorageSize
	full        bool
	signer      types.Signer
}

func newPendingBatch(snap *snapshot.Snapshot, maxSize common.StorageSize, signer types.Signer) *pendingBatch {
	return &pendingBatch{
		snap:        snap,
		txCounts:    make(map[common.Address]uint64),
		appliedTxes: nil,
		sizeBytes:   0,
		maxSize:     maxSize,
		full:        false,
		signer:      signer,
	}
}

func newPendingBatchFromExisting(batch *pendingBatch, maxSize common.StorageSize) *pendingBatch {
	return &pendingBatch{
		snap:        batch.snap,
		txCounts:    batch.txCounts,
		appliedTxes: nil,
		sizeBytes:   0,
		maxSize:     maxSize,
		full:        false,
		signer:      batch.signer,
	}
}

func (p *pendingBatch) getTxCount(account common.Address) uint64 {
	count, ok := p.txCounts[account]
	if !ok {
		txCount, err := p.snap.GetTransactionCount(arbcommon.NewAddressFromEth(account))
		if err != nil {
			panic(err)
		}
		count = txCount.Uint64()
		p.txCounts[account] = count
	}
	return count
}

func (p *pendingBatch) popRandomTx(queuedTxes *txQueues, signer types.Signer) (*types.Transaction, int, bool) {
	queuedCount := int32(len(queuedTxes.accounts))
	if queuedCount == 0 {
		return nil, 0, false
	}
	index := int(rand.Int31n(queuedCount))
	first := true
	lastIndex := index
	index--
	for {
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

		sender, _ := types.Sender(signer, tx)
		nextValidNonce := p.getTxCount(sender)
		if tx.Nonce() > nextValidNonce {
			continue
		}
		if p.sizeBytes+tx.Size() > p.maxSize {
			p.full = true
			return nil, 0, true
		}
		queuedTxes.removeTxFromAccountAtIndex(index)

		if tx.Nonce() < nextValidNonce {
			// Just discard this tx since it is old
			queuedTxes.maybeRemoveAccountAtIndex(index)
			return nil, 0, true
		}

		return tx, index, true
	}
}

func snapWithTx(snap *snapshot.Snapshot, tx *types.Transaction, signer types.Signer) (*snapshot.Snapshot, error) {
	msg, err := message.NewL2Message(message.SignedTransaction{Tx: tx})
	if err != nil {
		return nil, err
	}

	sender, _ := types.Sender(signer, tx)
	_, err = snap.AddMessage(msg, arbcommon.NewAddressFromEth(sender), arbcommon.NewHashFromEth(tx.Hash()))
	return snap, err
}

func (p *pendingBatch) addUpdatedSnap(tx *types.Transaction, newSnap *snapshot.Snapshot) {
	p.snap = newSnap
	p.addIncludedTx(tx)
}

func (p *pendingBatch) addIncludedTx(tx *types.Transaction) {
	p.appliedTxes = append(p.appliedTxes, tx)
	p.sizeBytes += tx.Size()
	sender, _ := types.Sender(p.signer, tx)
	p.txCounts[sender] = tx.Nonce() + 1
}

func (p *pendingBatch) checkValidForQueue(tx *types.Transaction) error {
	//ethSender, _ := types.Sender(p.signer, tx)
	//sender := arbcommon.NewAddressFromEth(ethSender)
	//txCount, err := p.snap.GetTransactionCount(sender)
	//if err != nil {
	//	return err
	//}
	//
	//if tx.Nonce() < txCount.Uint64() {
	//	return core.ErrNonceTooLow
	//}
	//
	//amount, err := p.snap.GetBalance(sender)
	//if err != nil {
	//	return err
	//}
	//
	//if tx.Cost().Cmp(amount) > 0 {
	//	log.Println("tx rejected for insufficient funds:", tx.Value(), tx.GasPrice(), tx.Gas(), amount)
	//	return core.ErrInsufficientFunds
	//}
	return nil
}
