/*
 * Copyright 2019, Offchain Labs, Inc.
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
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type BlockInfo struct {
	BlockLog value.Value
	Header   *types.Header
}

type AggregatorStore interface {
	GetMessage(index uint64) (value.Value, error)
	GetLog(index uint64) (value.Value, error)
	GetPossibleRequestInfo(requestId common.Hash) *uint64
	GetPossibleBlock(blockHash common.Hash) *uint64
	GetBlock(height uint64) (*BlockInfo, error)
	LatestBlock() (*common.BlockId, error)
	LogCount() (uint64, error)
	MessageCount() (uint64, error)

	SaveLog(val value.Value) error
	SaveMessage(val []byte) error
	SaveBlock(header *types.Header, logIndex uint64) error
	SaveEmptyBlock(header *types.Header) error
	SaveBlockHash(blockHash common.Hash, blockHeight uint64) error
	SaveRequest(requestId common.Hash, logIndex uint64) error
	Reorg(height uint64, messageCount uint64, logCount uint64) error
}

type BlockEntry struct {
	header   *types.Header
	logIndex *uint64
}

type InMemoryAggregatorStore struct {
	sync.Mutex
	messages     [][]byte
	logs         []value.Value
	blocks       map[uint64]*BlockEntry
	latestBlock  uint64
	requestIndex map[common.Hash]uint64
	blockIndex   map[common.Hash]uint64
}

func NewInMemoryAggregatorStore() *InMemoryAggregatorStore {
	return &InMemoryAggregatorStore{
		blocks:       make(map[uint64]*BlockEntry),
		requestIndex: make(map[common.Hash]uint64),
		blockIndex:   make(map[common.Hash]uint64),
	}
}

func (as *InMemoryAggregatorStore) GetMessage(index uint64) (value.Value, error) {
	as.Lock()
	defer as.Unlock()
	if index >= uint64(len(as.messages)) {
		return nil, errors.New("failed to get l2message")
	}
	panic("UNSUPPORTED")
	//return as.messages[index], nil
}

func (as *InMemoryAggregatorStore) GetLog(index uint64) (value.Value, error) {
	as.Lock()
	defer as.Unlock()
	if index >= uint64(len(as.logs)) {
		return nil, errors.New("failed to get log")
	}
	return as.logs[index], nil
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
	rawBlock, ok := as.blocks[height]
	if !ok {
		return nil, nil
	}
	var blockLog value.Value
	if rawBlock.logIndex != nil {
		if *rawBlock.logIndex >= uint64(len(as.logs)) {
			panic("out of bounds")
		}
		blockLog = as.logs[*rawBlock.logIndex]
	}
	return &BlockInfo{
		BlockLog: blockLog,
		Header:   rawBlock.header,
	}, nil
}

func (as *InMemoryAggregatorStore) LatestBlock() (*common.BlockId, error) {
	as.Lock()
	defer as.Unlock()
	if len(as.blocks) == 0 {
		return nil, errors.New("No blocks")
	}
	block := as.blocks[as.latestBlock]
	return &common.BlockId{
		Height:     common.NewTimeBlocksInt(int64(as.latestBlock)),
		HeaderHash: common.NewHashFromEth(block.header.Hash()),
	}, nil
}

func (as *InMemoryAggregatorStore) SaveLog(val value.Value) error {
	as.Lock()
	defer as.Unlock()
	as.logs = append(as.logs, val)
	return nil
}

func (as *InMemoryAggregatorStore) SaveMessage(val []byte) error {
	as.Lock()
	defer as.Unlock()
	as.messages = append(as.messages, val)
	return nil
}

func (as *InMemoryAggregatorStore) SaveBlock(header *types.Header, logIndex uint64) error {
	as.Lock()
	defer as.Unlock()
	if logIndex >= uint64(len(as.logs)) {
		return errors.New("bad log index")
	}
	as.blocks[header.Number.Uint64()] = &BlockEntry{
		header:   header,
		logIndex: &logIndex,
	}
	as.latestBlock = header.Number.Uint64()
	return nil
}

func (as *InMemoryAggregatorStore) SaveEmptyBlock(header *types.Header) error {
	as.Lock()
	defer as.Unlock()
	as.blocks[header.Number.Uint64()] = &BlockEntry{
		header: header,
	}
	as.latestBlock = header.Number.Uint64()
	return nil
}

func (as *InMemoryAggregatorStore) SaveBlockHash(blockHash common.Hash, blockHeight uint64) error {
	as.Lock()
	defer as.Unlock()
	as.blockIndex[blockHash] = blockHeight
	return nil
}

func (as *InMemoryAggregatorStore) SaveRequest(requestId common.Hash, logIndex uint64) error {
	as.Lock()
	defer as.Unlock()
	as.requestIndex[requestId] = logIndex
	return nil
}

func (as *InMemoryAggregatorStore) Reorg(height uint64, messageCount uint64, logCount uint64) error {
	as.Lock()
	defer as.Unlock()
	log.Info().
		Uint64("height", height).
		Uint64("messageCount", messageCount).
		Uint64("logCount", logCount).
		Msg("aggregator triggered reorg")
	as.messages = as.messages[:messageCount]
	as.logs = as.logs[:logCount]
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
		if block.logIndex != nil {
			i := *block.logIndex
			if i >= uint64(len(as.logs)) {
				panic("bad block")
			}
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
