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
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type BlockInfo struct {
	BlockLog uint64
	Header   *types.Header
}

type EVMRequestInfo struct {
	RequestId common.Hash
	LogIndex  uint64
}

type AggregatorStore interface {
	GetPossibleRequestInfo(requestId common.Hash) *uint64
	GetPossibleBlock(blockHash common.Hash) *uint64
	GetBlockInfo(height uint64) (*BlockInfo, error)
	BlockCount() (uint64, error)

	SaveBlock(header *types.Header, logIndex uint64, requests []EVMRequestInfo) error
	Reorg(height uint64) error

	CurrentLogCount() (*big.Int, error)
	UpdateCurrentLogCount(count *big.Int) error
}

type InMemoryAggregatorStore struct {
	sync.Mutex
	messages     [][]byte
	logs         []value.Value
	blocks       map[uint64]*BlockInfo
	latestBlock  uint64
	requestIndex map[common.Hash]uint64
	blockIndex   map[common.Hash]uint64
}

func NewInMemoryAggregatorStore() *InMemoryAggregatorStore {
	return &InMemoryAggregatorStore{
		blocks:       make(map[uint64]*BlockInfo),
		requestIndex: make(map[common.Hash]uint64),
		blockIndex:   make(map[common.Hash]uint64),
	}
}

func (as *InMemoryAggregatorStore) GetPossibleRequestInfo(requestId common.Hash) *uint64 {
	as.Lock()
	defer as.Unlock()
	request, ok := as.requestIndex[requestId]
	if !ok {
		return nil
	}
	return &request
}

func (as *InMemoryAggregatorStore) GetPossibleBlock(blockHash common.Hash) *uint64 {
	as.Lock()
	defer as.Unlock()
	block, ok := as.blockIndex[blockHash]
	if !ok {
		return nil
	}
	return &block
}

func (as *InMemoryAggregatorStore) GetBlock(height uint64) (*BlockInfo, error) {
	as.Lock()
	defer as.Unlock()
	blockInfo, ok := as.blocks[height]
	if !ok {
		return nil, nil
	}
	return blockInfo, nil
}

func (as *InMemoryAggregatorStore) LatestBlockInfo() (*BlockInfo, error) {
	as.Lock()
	defer as.Unlock()
	if len(as.blocks) == 0 {
		return nil, errors.New("No blocks")
	}
	return as.blocks[as.latestBlock], nil
}

func (as *InMemoryAggregatorStore) SaveBlock(header *types.Header, logIndex uint64, requests []EVMRequestInfo) error {
	as.Lock()
	defer as.Unlock()
	if logIndex >= uint64(len(as.logs)) {
		return errors.New("bad log index")
	}
	as.blocks[header.Number.Uint64()] = &BlockInfo{
		Header:   header,
		BlockLog: logIndex,
	}
	as.latestBlock = header.Number.Uint64()
	as.blockIndex[common.NewHashFromEth(header.Hash())] = header.Number.Uint64()

	for _, request := range requests {
		as.requestIndex[request.RequestId] = request.LogIndex
	}
	return nil
}

func (as *InMemoryAggregatorStore) Reorg(height uint64) error {
	as.Lock()
	defer as.Unlock()
	log.Info().
		Uint64("height", height).
		Msg("aggregator triggered reorg")
	for i := as.latestBlock; i > height; i-- {
		delete(as.blocks, i)
		if i == 0 {
			break
		}
	}
	as.latestBlock = height
	for blockHeight, block := range as.blocks {
		if blockHeight > height {
			panic("bad height")
		}
		i := block.BlockLog
		if i >= uint64(len(as.logs)) {
			panic("bad block")
		}
	}
	return nil
}

func (as *InMemoryAggregatorStore) LogCount() (uint64, error) {
	as.Lock()
	defer as.Unlock()
	return uint64(len(as.logs)), nil
}

func (as *InMemoryAggregatorStore) MessageCount() (uint64, error) {
	as.Lock()
	defer as.Unlock()
	return uint64(len(as.messages)), nil
}
