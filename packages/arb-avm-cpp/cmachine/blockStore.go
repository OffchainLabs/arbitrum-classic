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
#cgo LDFLAGS: -L. -L../build/rocksdb -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb -lkeccak -ldl
#include "../cavm/cblockstore.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"github.com/pkg/errors"
	"runtime"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

type BlockStore struct {
	c unsafe.Pointer
}

func deleteBlockStore(bs *BlockStore) {
	C.deleteBlockStore(bs.c)
}

func NewBlockStore(c unsafe.Pointer) *BlockStore {
	bs := &BlockStore{c: c}
	runtime.SetFinalizer(bs, deleteBlockStore)
	return bs
}

func (bs *BlockStore) PutBlock(id *common.BlockId, data []byte) error {
	cHeight := intToData(id.Height.AsInt())
	defer C.free(cHeight)
	cHash := hashToData(id.HeaderHash)
	defer C.free(cHash)
	cData := C.CBytes(data)
	defer C.free(cData)

	success := C.putBlock(
		bs.c,
		cHeight,
		cHash,
		cData,
		C.int(len(data)),
	)

	if success == 0 {
		return errors.New("write failed")
	}
	return nil
}

func (bs *BlockStore) DeleteBlock(id *common.BlockId) error {
	cHeight := intToData(id.Height.AsInt())
	defer C.free(cHeight)
	cHash := hashToData(id.HeaderHash)
	defer C.free(cHash)

	success := C.deleteBlock(
		bs.c,
		cHeight,
		cHash,
	)

	if success == 0 {
		return errors.New("delete failed")
	}
	return nil
}

func (bs *BlockStore) GetBlock(id *common.BlockId) ([]byte, error) {
	cHeight := intToData(id.Height.AsInt())
	defer C.free(cHeight)
	cHash := hashToData(id.HeaderHash)
	defer C.free(cHash)

	result := C.getBlock(
		bs.c,
		cHeight,
		cHash,
	)

	if result.found == 0 {
		return nil, errors.New("block not found in block store")
	}

	return toByteSlice(result.slice), nil
}

func (bs *BlockStore) BlocksAtHeight(height *common.TimeBlocks) []*common.BlockId {
	cHeight := intToData(height.AsInt())
	defer C.free(cHeight)

	cHashList := C.blockHashesAtHeight(bs.c, cHeight)
	defer C.free(cHashList.data)

	if cHashList.count == 0 {
		return nil
	}

	data := C.GoBytes(unsafe.Pointer(cHashList.data), cHashList.count*32)
	ret := make([]*common.BlockId, 0, int(cHashList.count))
	for i := 0; i < int(cHashList.count); i++ {
		var hashVal common.Hash
		copy(hashVal[:], data[i*32:])
		ret = append(ret, &common.BlockId{
			Height:     height,
			HeaderHash: hashVal,
		})
	}

	return ret
}

func (bs *BlockStore) IsBlockStoreEmpty() bool {
	return C.isBlockStoreEmpty(bs.c) == 1
}

func (bs *BlockStore) MaxBlockStoreHeight() *common.TimeBlocks {
	cHeight := C.maxBlockStoreHeight(bs.c)
	defer C.free(cHeight)
	return common.NewTimeBlocks(dataToInt(cHeight))
}

func (bs *BlockStore) MinBlockStoreHeight() *common.TimeBlocks {
	cHeight := C.minBlockStoreHeight(bs.c)
	defer C.free(cHeight)
	return common.NewTimeBlocks(dataToInt(cHeight))
}
