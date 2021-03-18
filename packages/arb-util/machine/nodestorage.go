/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

package machine

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"math/big"
)

type BlockInfo struct {
	BlockLog uint64
	LogCount uint64
	Header   *types.Header
}

func (b *BlockInfo) InitialLogIndex() uint64 {
	return b.BlockLog - b.LogCount
}

type EVMRequestInfo struct {
	RequestId common.Hash
	LogIndex  uint64
}

type NodeStore interface {
	GetPossibleRequestInfo(requestId common.Hash) *uint64
	GetPossibleBlock(blockHash common.Hash) *uint64
	GetBlockInfo(height uint64) (*BlockInfo, error)
	BlockCount() (uint64, error)

	SaveMessageBatch(batchNum *big.Int, logIndex uint64) error
	GetMessageBatch(batchNum *big.Int) *uint64
	SaveBlock(header *types.Header, logIndex, logCount uint64, requests []EVMRequestInfo) error
	Reorg(height uint64) error
}
