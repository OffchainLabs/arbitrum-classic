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
	"github.com/ethereum/go-ethereum/core"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	arbcommon "github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type statelessBatch struct {
	db          *txdb.TxDB
	signer      types.Signer
	appliedTxes []*types.Transaction
	sizeBytes   common.StorageSize
	maxSize     common.StorageSize
	full        bool
}

func newStatelessBatch(db *txdb.TxDB, maxSize common.StorageSize, signer types.Signer) *statelessBatch {
	return &statelessBatch{
		db:          db,
		signer:      signer,
		appliedTxes: nil,
		sizeBytes:   0,
		maxSize:     maxSize,
		full:        false,
	}
}

func (p *statelessBatch) newFromExisting() batch {
	return &statelessBatch{
		db:          p.db,
		signer:      p.signer,
		appliedTxes: nil,
		sizeBytes:   0,
		maxSize:     p.maxSize,
		full:        false,
	}
}

func (p *statelessBatch) getAppliedTxes() []*types.Transaction {
	return p.appliedTxes
}

func (p *statelessBatch) updateCurrentSnap(*list.List) {

}

func (p *statelessBatch) getLatestSnap() *snapshot.Snapshot {
	return nil
}

func (p *statelessBatch) addIncludedTx(tx *types.Transaction) error {
	p.appliedTxes = append(p.appliedTxes, tx)
	p.sizeBytes += tx.Size()
	return nil
}

func (p *statelessBatch) validateTx(tx *types.Transaction) (txResponse, error) {
	// If we don't have access to a db, skip this check
	rejectLogger := logger.With().Hex("tx", tx.Hash().Bytes()).Logger()
	rejectMsg := "rejected user tx"
	if p.db != nil {
		sender, err := types.Sender(p.signer, tx)
		if err != nil {
			rejectLogger.Info().Stack().Err(err).Str("reason", "sender").Msg(rejectMsg)
			return REMOVE, errors.New("couldn't recover sender")
		}

		txCount, err := p.db.LatestSnapshot().GetTransactionCount(arbcommon.NewAddressFromEth(sender))
		if err != nil {
			rejectLogger.Info().Stack().Err(err).Str("reason", "snapshot").Msg(rejectMsg)
			return REMOVE, errors.New("aggregator failed to verify nonce")
		}

		// If the transaction's nonce is less than the latest state tx count, we can ignore
		if tx.Nonce() < txCount.Uint64() {
			rejectLogger.Info().
				Str("reason", "nonce").
				Uint64("nonce", tx.Nonce()).
				Uint64("txcount", txCount.Uint64()).
				Msg(rejectMsg)
			return REMOVE, errors.WithStack(core.ErrNonceTooLow)
		}
	}

	if p.sizeBytes+tx.Size() > p.maxSize {
		p.full = true
		return FULL, nil
	}
	return ACCEPT, nil
}

func (p *statelessBatch) isFull() bool {
	return p.full
}
