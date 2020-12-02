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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/txdb"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"log"
)

type statefulBatch struct {
	*statelessBatch
	snap     *snapshot.Snapshot
	txCounts map[common.Address]uint64
}

func newStatefulBatch(db *txdb.TxDB, maxSize common.StorageSize, signer types.Signer) *statefulBatch {
	return &statefulBatch{
		statelessBatch: newStatelessBatch(db, maxSize, signer),
		snap:           db.LatestSnapshot(),
		txCounts:       make(map[common.Address]uint64),
	}
}

func (p *statefulBatch) newFromExisting() batch {
	return &statefulBatch{
		statelessBatch: p.statelessBatch.newFromExisting().(*statelessBatch),
		snap:           p.snap,
		txCounts:       p.txCounts,
	}
}

func (p *statefulBatch) getTxCount(account common.Address) uint64 {
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

func (p *statefulBatch) validateTx(tx *types.Transaction) txResponse {
	sender, err := types.Sender(p.signer, tx)
	if err != nil {
		return REMOVE
	}
	nextValidNonce := p.getTxCount(sender)
	if tx.Nonce() > nextValidNonce {
		return SKIP
	}
	if tx.Nonce() < nextValidNonce {
		// Just discard this tx since it is old
		return REMOVE
	}
	return p.statelessBatch.validateTx(tx)
}

func snapWithTx(snap *snapshot.Snapshot, tx *types.Transaction, signer types.Signer) (*snapshot.Snapshot, error) {
	msg, err := message.NewL2Message(message.SignedTransaction{Tx: tx})
	if err != nil {
		return nil, err
	}

	sender, err := types.Sender(signer, tx)
	if err != nil {
		return nil, err
	}

	_, err = snap.AddMessage(msg, arbcommon.NewAddressFromEth(sender), arbcommon.NewHashFromEth(tx.Hash()))
	return snap, err
}

func (p *statefulBatch) getLatestSnap() *snapshot.Snapshot {
	return p.snap
}

func (p *statefulBatch) addIncludedTx(tx *types.Transaction) error {
	newSnap := p.snap.Clone()
	newSnap, err := snapWithTx(newSnap, tx, p.signer)
	if err != nil {
		return err
	}

	sender, err := types.Sender(p.signer, tx)
	if err != nil {
		return err
	}

	if err := p.statelessBatch.addIncludedTx(tx); err != nil {
		return err
	}

	p.snap = newSnap
	p.txCounts[sender] = tx.Nonce() + 1
	return nil
}

func (p *statefulBatch) checkValidForQueue(tx *types.Transaction) error {
	ethSender, err := types.Sender(p.signer, tx)
	if err != nil {
		return err
	}
	sender := arbcommon.NewAddressFromEth(ethSender)
	txCount, err := p.snap.GetTransactionCount(sender)
	if err != nil {
		return err
	}

	if tx.Nonce() < txCount.Uint64() {
		return core.ErrNonceTooLow
	}

	amount, err := p.snap.GetBalance(sender)
	if err != nil {
		return err
	}

	if tx.Cost().Cmp(amount) > 0 {
		log.Println("tx rejected for insufficient funds:", tx.Value(), tx.GasPrice(), tx.Gas(), amount)
		return core.ErrInsufficientFunds
	}
	return nil
}

func (p *statefulBatch) updateCurrentSnap(pendingSentBatches *list.List) {
	snap := p.db.LatestSnapshot().Clone()
	if p.snap.Height().Cmp(snap.Height()) < 0 {
		// Add all of the already broadcast transactions to the snapshot
		// If they were already included, they'll be ignored because they will
		// have invalid sequence numbers
		n := pendingSentBatches.Front()
		for n != nil {
			item := n.Value.(*pendingSentBatch)
			for _, tx := range item.txes {
				var err error
				newSnap, err := snapWithTx(snap, tx, p.signer)
				if err != nil {
					continue
				}
				snap = newSnap
			}
			n = n.Next()
		}
		for _, tx := range p.appliedTxes {
			var err error
			newSnap, err := snapWithTx(snap, tx, p.signer)
			if err != nil {
				continue
			}
			snap = newSnap
		}
		p.snap = snap
	}
}
