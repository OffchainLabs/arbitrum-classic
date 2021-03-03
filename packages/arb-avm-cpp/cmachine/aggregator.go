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
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -L../build/rocksdb -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb -ldl
#include "../cavm/caggregator.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"encoding/binary"
	"github.com/pkg/errors"
	"math/big"
	"runtime"
	"unsafe"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type AggregatorStore struct {
	c unsafe.Pointer
}

func deleteAggregatorStore(bs *AggregatorStore) {
	C.deleteAggregatorStore(bs.c)
}

func NewAggregatorStore(c unsafe.Pointer) *AggregatorStore {
	as := &AggregatorStore{c: c}
	runtime.SetFinalizer(as, deleteAggregatorStore)
	return as
}

func parseBlockData(data []byte) (*machine.BlockInfo, error) {
	index := binary.BigEndian.Uint64(data)
	data = data[8:]
	header := &types.Header{}
	if err := header.UnmarshalJSON(data); err != nil {
		return nil, err
	}
	return &machine.BlockInfo{
		BlockLog: index,
		Header:   header,
	}, nil
}

func serializeBlockData(header *types.Header, logIndex uint64) ([]byte, error) {
	var blockData []byte

	logIndexData := make([]byte, 8)
	binary.BigEndian.PutUint64(logIndexData[:], logIndex)
	blockData = append(blockData, logIndexData...)

	headerJSON, err := header.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return append(blockData, headerJSON...), nil
}

func (as *AggregatorStore) SaveBlock(header *types.Header, logIndex uint64, requests []machine.EVMRequestInfo) error {
	blockData, err := serializeBlockData(header, logIndex)
	if err != nil {
		return err
	}

	if C.aggregatorSaveBlock(as.c, C.uint64_t(header.Number.Uint64()), unsafeDataPointer(blockData), C.int(len(blockData))) == 0 {
		return errors.New("failed to save block")
	}

	for _, request := range requests {
		if C.aggregatorSaveRequest(as.c, unsafeDataPointer(request.RequestId.Bytes()), C.uint64_t(request.LogIndex)) == 0 {
			return errors.New("failed to save request")
		}
	}

	headerHash := header.Hash()
	if C.aggregatorSaveBlockHash(as.c, unsafeDataPointer(headerHash.Bytes()), C.uint64_t(header.Number.Uint64())) == 0 {
		return errors.New("failed to save request")
	}

	return nil
}

func (as *AggregatorStore) BlockCount() (uint64, error) {
	result := C.aggregatorBlockCount(as.c)
	if result.found == 0 {
		return 0, errors.New("failed to load block count")

	}
	return uint64(result.value), nil
}

func (as *AggregatorStore) GetBlockInfo(height uint64) (*machine.BlockInfo, error) {
	blockData := C.aggregatorGetBlock(as.c, C.uint64_t(height))
	if blockData.found == 0 {
		return nil, nil
	}
	return parseBlockData(receiveByteSlice(blockData.data))
}

func (as *AggregatorStore) Reorg(height uint64) error {
	status := C.aggregatorReorg(as.c, C.uint64_t(height))
	if status == 0 {
		return errors.New("failed to reset aggregator height")
	}
	return nil
}

func (as *AggregatorStore) GetPossibleRequestInfo(requestId common.Hash) *uint64 {
	result := C.aggregatorGetPossibleRequestInfo(as.c, unsafeDataPointer(requestId.Bytes()))
	if result.found == 0 {
		return nil
	}
	index := uint64(result.value)
	return &index
}

func (as *AggregatorStore) GetPossibleBlock(blockHash common.Hash) *uint64 {
	result := C.aggregatorGetPossibleBlock(as.c, unsafeDataPointer(blockHash.Bytes()))
	if result.found == 0 {
		return nil
	}
	index := uint64(result.value)
	return &index
}

func (as *AggregatorStore) CurrentLogCount() (*big.Int, error) {
	result := C.aggregatorLogsProcessedCount(as.c)
	if result.found == 0 {
		return nil, errors.New("failed to get processed log count")
	}
	return receiveBigInt(result.value), nil
}

func (as *AggregatorStore) UpdateCurrentLogCount(count *big.Int) error {
	countData := math.U256Bytes(count)
	status := C.aggregatorUpdateLogsProcessedCount(as.c, unsafeDataPointer(countData))
	if status == 0 {
		return errors.New("failed to update processed log count")
	}
	return nil
}
