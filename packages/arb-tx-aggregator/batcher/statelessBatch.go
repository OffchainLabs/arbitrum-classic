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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-tx-aggregator/snapshot"
)

type statelessBatch struct {
	appliedTxes []*types.Transaction
	sizeBytes   common.StorageSize
	maxSize     common.StorageSize
	full        bool
}

func newStatelessBatch(maxSize common.StorageSize) *statelessBatch {
	return &statelessBatch{
		appliedTxes: nil,
		sizeBytes:   0,
		maxSize:     maxSize,
		full:        false,
	}
}

func (p *statelessBatch) newFromExisting() batch {
	return &statelessBatch{
		appliedTxes: nil,
		sizeBytes:   0,
		maxSize:     p.maxSize,
		full:        false,
	}
}

func (p *statelessBatch) getAppliedTxes() []*types.Transaction {
	return p.appliedTxes
}

func (p *statelessBatch) checkValidForQueue(*types.Transaction) error {
	return nil
}

func (p *statelessBatch) updateFromCurrentSnap(*snapshot.Snapshot, *list.List) {

}

func (p *statelessBatch) getLatestSnap() *snapshot.Snapshot {
	return nil
}

func (p *statelessBatch) addIncludedTx(tx *types.Transaction) error {
	p.appliedTxes = append(p.appliedTxes, tx)
	p.sizeBytes += tx.Size()
	return nil
}

func (p *statelessBatch) validateTx(tx *types.Transaction) txResponse {
	if p.sizeBytes+tx.Size() > p.maxSize {
		p.full = true
		return FULL
	}
	return ACCEPT
}

func (p *statelessBatch) isFull() bool {
	return p.full
}
