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
#cgo LDFLAGS: -L. -L../build/rocksdb -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb
#include "../cavm/cnodestore.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"runtime"
	"unsafe"
)

type NodeStore struct {
	c unsafe.Pointer
}

func deleteNodeStore(ns *NodeStore) {
	C.deleteNodeStore(ns.c)
}

func NewNodeStore(c unsafe.Pointer) *NodeStore {
	bs := &NodeStore{c: c}
	runtime.SetFinalizer(bs, deleteNodeStore)
	return bs
}

func (ns *NodeStore) PutNode(height uint64, hash common.Hash, data []byte) error {
	cHash := hashToData(hash)
	defer C.free(cHash)
	cData := C.CBytes(data)
	defer C.free(cData)

	success := C.putNode(
		ns.c,
		C.uint64_t(height),
		cHash,
		cData,
		C.int(len(data)),
	)

	if success == 0 {
		return errors.New("write failed")
	}
	return nil
}

func (ns *NodeStore) GetNode(height uint64, hash common.Hash) ([]byte, error) {
	cHash := hashToData(hash)
	defer C.free(cHash)

	result := C.getNode(
		ns.c,
		C.uint64_t(height),
		cHash,
	)

	if result.found == 0 {
		return nil, errors.New("not found")
	}

	return toByteSlice(result.slice), nil
}

func (ns *NodeStore) GetNodeHeight(hash common.Hash) (uint64, error) {
	cHash := hashToData(hash)
	defer C.free(cHash)

	result := C.getNodeHeight(
		ns.c,
		cHash,
	)

	if result.found == 0 {
		return 0, errors.New("not found")
	}

	return uint64(result.value), nil
}

func (ns *NodeStore) GetNodeHash(height uint64) (common.Hash, error) {
	result := C.getNodeHash(
		ns.c,
		C.uint64_t(height),
	)

	if result.found == 0 {
		return common.Hash{}, errors.New("not found")
	}

	return dataToHash(result.value), nil
}

func (ns *NodeStore) Empty() bool {
	return C.isNodeStoreEmpty(ns.c) == 1
}

func (ns *NodeStore) MaxHeight() uint64 {
	return uint64(C.maxNodeHeight(ns.c))
}
