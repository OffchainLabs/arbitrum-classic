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
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type statefulBatch struct {
	*statelessBatch
	snap     *snapshot.Snapshot
	txCounts map[common.Address]uint64
}

func newStatefulBatch(db *txdb.TxDB, maxSize common.StorageSize, signer types.Signer) (*statefulBatch, error) {
	snap, err := db.LatestSnapshot()
	if err != nil {
		return nil, err
	}
	return &statefulBatch{
		statelessBatch: newStatelessBatch(db, maxSize, signer),
		snap:           snap,
		txCounts:       make(map[common.Address]uint64),
	}, nil
}

func (p *statefulBatch) newFromExisting() batch {
	return &statefulBatch{
		statelessBatch: p.statelessBatch.newFromExisting().(*statelessBatch),
		snap:           p.snap,
		txCounts:       p.txCounts,
	}
}

func (p *statefulBatch) getTxCount(account common.Address) (uint64, error) {
	count, ok := p.txCounts[account]
	if !ok {
		txCount, err := p.snap.GetTransactionCount(arbcommon.NewAddressFromEth(account))
		if err != nil {
			return 0, err
		}
		count = txCount.Uint64()
		p.txCounts[account] = count
	}
	return count, nil
}

func (p *statefulBatch) validateTx(tx *types.Transaction) (txResponse, error) {
	sender, err := types.Sender(p.signer, tx)
	if err != nil {
		return REMOVE, errors.New("invalid signature")
	}
	nextValidNonce, err := p.getTxCount(sender)
	if err != nil {
		return SKIP, err
	}
	if tx.Nonce() > nextValidNonce {
		return SKIP, errors.WithStack(core.ErrNonceTooHigh)
	}
	if tx.Nonce() < nextValidNonce {
		// Just discard this tx since it is old
		return REMOVE, errors.WithStack(core.ErrNonceTooLow)
	}

	amount, err := p.snap.GetBalance(arbcommon.NewAddressFromEth(sender))
	if err != nil {
		return REMOVE, err
	}

	if tx.Cost().Cmp(amount) > 0 {
		logger.Warn().
			Str("value", tx.Value().String()).
			Str("gasPrice", tx.GasPrice().String()).
			Uint64("Gas", tx.Gas()).
			Str("amount", amount.String()).
			Msg("tx rejected for insufficient funds")
		return REMOVE, errors.WithStack(core.ErrInsufficientFunds)
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

func (p *statefulBatch) updateCurrentSnap(pendingSentBatches *list.List) error {
	snap, err := p.db.LatestSnapshot()
	if err != nil {
		return err
	}
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
	return nil
}
