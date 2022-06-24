/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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

package cmachine

/*
#include "../cavm/caggregator.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"encoding/binary"
	"math/big"
	"runtime"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type NodeStore struct {
	c unsafe.Pointer
}

func deleteNodeStore(bs *NodeStore) {
	C.deleteAggregatorStore(bs.c)
}

func NewNodeStore(c unsafe.Pointer) *NodeStore {
	as := &NodeStore{c: c}
	runtime.SetFinalizer(as, deleteNodeStore)
	return as
}

func parseBlockData(data []byte) (*machine.BlockInfo, error) {
	index := binary.BigEndian.Uint64(data)
	data = data[8:]
	count := binary.BigEndian.Uint64(data)
	data = data[8:]
	header := &types.Header{}
	if err := header.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	return &machine.BlockInfo{
		BlockLog: index,
		LogCount: count,
		Header:   header,
	}, nil
}

func serializeBlockData(info *machine.BlockInfo) ([]byte, error) {
	var blockData []byte

	logIndexData := make([]byte, 8)
	binary.BigEndian.PutUint64(logIndexData[:], info.BlockLog)
	blockData = append(blockData, logIndexData...)

	logCountData := make([]byte, 8)
	binary.BigEndian.PutUint64(logCountData[:], info.LogCount)
	blockData = append(blockData, logCountData...)

	headerJSON, err := info.Header.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return append(blockData, headerJSON...), nil
}

func (as *NodeStore) SaveMessageBatch(batchNum *big.Int, logIndex uint64) error {
	defer runtime.KeepAlive(as)
	result := C.aggregatorSaveMessageBatch(as.c, unsafeDataPointer(math.U256Bytes(batchNum)), C.uint64_t(logIndex))
	if result == 0 {
		return errors.New("failed to save message batch")
	}

	return nil
}

func (as *NodeStore) GetMessageBatch(batchNum *big.Int) *uint64 {
	defer runtime.KeepAlive(as)
	result := C.aggregatorGetMessageBatch(as.c, unsafeDataPointer(math.U256Bytes(batchNum)))
	if result.found == 0 {
		return nil
	}
	index := uint64(result.value)
	return &index
}

func (as *NodeStore) SaveBlock(info *machine.BlockInfo, requests []machine.EVMRequestInfo) error {
	defer runtime.KeepAlive(as)
	blockData, err := serializeBlockData(info)
	if err != nil {
		return err
	}

	rawRequestIds := make([][]byte, 0, len(requests))
	logIndexes := make([]C.uint64_t, 0, len(requests))
	for _, request := range requests {
		rawRequestId := new(big.Int).SetBytes(request.RequestId.Bytes())
		rawRequestIds = append(rawRequestIds, math.U256Bytes(rawRequestId))
		logIndexes = append(logIndexes, C.uint64_t(request.LogIndex))
	}
	byteSlices := make([]C.ByteSlice, 0, len(rawRequestIds))
	for _, data := range rawRequestIds {
		byteSlices = append(byteSlices, toByteSliceView(data))
	}

	var logIndexesPtr *C.uint64_t
	if len(logIndexes) > 0 {
		logIndexesPtr = &logIndexes[0]
	}
	headerHash := info.Header.Hash().Bytes()
	if C.aggregatorSaveBlock(
		as.c,
		C.uint64_t(info.Header.Number.Uint64()),
		unsafeDataPointer(headerHash),
		toByteSliceArrayView(byteSlices),
		logIndexesPtr,
		unsafeDataPointer(blockData),
		C.int(len(blockData))) == 0 {
		return errors.New("failed to save block")
	}

	return nil
}

func (as *NodeStore) BlockCount() (uint64, error) {
	defer runtime.KeepAlive(as)
	result := C.aggregatorBlockCount(as.c)
	if result.found == 0 {
		return 0, errors.New("failed to load block count")

	}
	return uint64(result.value), nil
}

func (as *NodeStore) GetBlockInfo(height uint64) (*machine.BlockInfo, error) {
	defer runtime.KeepAlive(as)
	blockData := C.aggregatorGetBlock(as.c, C.uint64_t(height))
	if blockData.found == 0 {
		return nil, nil
	}
	return parseBlockData(receiveByteSlice(blockData.data))
}

func (as *NodeStore) Reorg(height uint64) error {
	defer runtime.KeepAlive(as)
	status := C.aggregatorReorg(as.c, C.uint64_t(height))
	if status == 0 {
		return errors.New("failed to reset node height")
	}
	return nil
}

func (as *NodeStore) GetPossibleRequestInfo(requestId common.Hash) *uint64 {
	defer runtime.KeepAlive(as)
	result := C.aggregatorGetPossibleRequestInfo(as.c, unsafeDataPointer(requestId.Bytes()))
	if result.found == 0 {
		return nil
	}
	index := uint64(result.value)
	return &index
}

func (as *NodeStore) GetPossibleBlock(blockHash common.Hash) *uint64 {
	defer runtime.KeepAlive(as)
	result := C.aggregatorGetPossibleBlock(as.c, unsafeDataPointer(blockHash.Bytes()))
	if result.found == 0 {
		return nil
	}
	index := uint64(result.value)
	return &index
}
