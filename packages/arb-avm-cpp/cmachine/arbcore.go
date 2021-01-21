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

package cmachine

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -L../build/rocksdb -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb -ldl
#include "../cavm/arbcore.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"runtime"
	"unsafe"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ArbCore struct {
	c unsafe.Pointer
}

func deleteArbCore(ac *ArbCore) {
	C.deleteArbCore(ac.c)
}

func NewArbCore(c unsafe.Pointer) *ArbCore {
	as := &ArbCore{c: c}
	runtime.SetFinalizer(as, deleteArbCore)
	return as
}

func (ac *ArbCore) GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error) {
	result := C.arbCoreGetSends(ac.c, intToData(startIndex), intToData(count))
	if result.found == 0 {
		return nil, errors.New("failed to get log")
	}
	return toByteSliceArray(result.slice), nil
}

func (ac *ArbCore) GetMessages(startIndex *big.Int, count *big.Int) ([]inbox.InboxMessage, error) {
	result := C.arbCoreGetMessages(ac.c, intToData(startIndex), intToData(count))
	if result.found == 0 {
		return nil, errors.New("failed to get log")
	}
	return toByteInboxArray(result.slice), nil
}

func (as *AggregatorStore) MessageCount() (uint64, error) {
	result := C.aggregatorMessageCount(as.c)
	if result.found == 0 {
		return 0, errors.New("failed to load l2message count")
	}
	return uint64(result.value), nil
}

func (as *AggregatorStore) SaveMessage(buf []byte) error {
	cData := C.CBytes(buf)
	defer C.free(cData)
	if C.aggregatorSaveMessage(as.c, cData, C.uint64_t(len(buf))) == 0 {
		return errors.New("failed to save l2message")
	}

	return nil
}

func (as *AggregatorStore) GetMessage(index uint64) (value.Value, error) {
	result := C.aggregatorGetMessage(as.c, C.uint64_t(index))
	if result.found == 0 {
		return nil, errors.New("failed to get l2message")
	}
	logBytes := toByteSlice(result.slice)
	return value.UnmarshalValue(bytes.NewBuffer(logBytes))
}

func parseBlockData(data []byte) (*types.Header, *uint64, error) {
	blockType := data[0]
	data = data[1:]
	var logIndex *uint64
	if blockType == 1 {
		index := binary.BigEndian.Uint64(data)
		logIndex = &index
		data = data[8:]
	}
	header := &types.Header{}
	if err := header.UnmarshalJSON(data); err != nil {
		return nil, nil, err
	}
	return header, logIndex, nil
}

func (as *AggregatorStore) LatestBlock() (*common.BlockId, error) {
	result := C.aggregatorLatestBlock(as.c)
	if result.found == 0 {
		return nil, errors.New("failed to load block count")
	}

	header, _, err := parseBlockData(toByteSlice(result.data))
	if err != nil {
		return nil, err
	}
	return &common.BlockId{
		Height:     common.NewTimeBlocks(new(big.Int).SetUint64(uint64(result.height))),
		HeaderHash: common.NewHashFromEth(header.Hash()),
	}, nil
}

func (as *AggregatorStore) SaveBlock(header *types.Header, logIndex uint64) error {
	blockData := []byte{1}

	logIndexData := make([]byte, 8)
	binary.BigEndian.PutUint64(logIndexData[:], logIndex)
	blockData = append(blockData, logIndexData...)

	headerJSON, err := header.MarshalJSON()
	if err != nil {
		return err
	}
	blockData = append(blockData, headerJSON...)
	cBlockData := C.CBytes(blockData)
	defer C.free(cBlockData)

	if C.aggregatorSaveBlock(as.c, C.uint64_t(header.Number.Uint64()), cBlockData, C.int(len(blockData))) == 0 {
		return errors.New("failed to save block")
	}
	return nil
}

func (as *AggregatorStore) SaveEmptyBlock(header *types.Header) error {
	blockData := []byte{0}
	headerJSON, err := header.MarshalJSON()
	if err != nil {
		return err
	}
	blockData = append(blockData, headerJSON...)
	cBlockData := C.CBytes(blockData)
	defer C.free(cBlockData)

	if C.aggregatorSaveBlock(as.c, C.uint64_t(header.Number.Uint64()), cBlockData, C.int(len(blockData))) == 0 {
		return errors.New("failed to save block")
	}
	return nil
}

func (as *AggregatorStore) GetBlock(height uint64) (*machine.BlockInfo, error) {
	blockData := C.aggregatorGetBlock(as.c, C.uint64_t(height))
	if blockData.found == 0 {
		return nil, nil
	}
	header, logIndex, err := parseBlockData(toByteSlice(blockData.data))
	if err != nil {
		return nil, err
	}
	info := &machine.BlockInfo{
		Header: header,
	}
	if logIndex != nil {
		avmLog, err := as.GetLog(*logIndex)
		if err != nil {
			return nil, err
		}
		info.BlockLog = avmLog
	}
	return info, nil
}

func (as *AggregatorStore) Reorg(height uint64, sendCount uint64, logCount uint64) error {
	if C.aggregatorReorg(
		as.c,
		C.uint64_t(height),
		C.uint64_t(sendCount),
		C.uint64_t(logCount),
	) == 0 {
		return errors.New("failed to restore block")
	}
	return nil
}

func (as *AggregatorStore) GetPossibleRequestInfo(requestId common.Hash) *uint64 {
	cHash := hashToData(requestId)
	defer C.free(cHash)

	result := C.aggregatorGetPossibleRequestInfo(as.c, cHash)
	if result.found == 0 {
		return nil
	}
	index := uint64(result.value)
	return &index
}

func (as *AggregatorStore) SaveRequest(requestId common.Hash, logIndex uint64) error {
	cHash := hashToData(requestId)
	defer C.free(cHash)

	if C.aggregatorSaveRequest(as.c, cHash, C.uint64_t(logIndex)) == 0 {
		return errors.New("failed to save request")
	}
	return nil
}

func (as *AggregatorStore) GetPossibleBlock(blockHash common.Hash) *uint64 {
	cHash := hashToData(blockHash)
	defer C.free(cHash)

	result := C.aggregatorGetPossibleBlock(as.c, cHash)
	if result.found == 0 {
		return nil
	}
	index := uint64(result.value)
	return &index
}

func (as *AggregatorStore) SaveBlockHash(blockHash common.Hash, blockHeight uint64) error {
	cHash := hashToData(blockHash)
	defer C.free(cHash)

	if C.aggregatorSaveBlockHash(as.c, cHash, C.uint64_t(blockHeight)) == 0 {
		return errors.New("failed to save request")
	}
	return nil
}
