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
#include "../cavm/carbcore.h"
#include "../cavm/cvaluecache.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"runtime"
	"unsafe"

	"github.com/pkg/errors"
)

type ArbCore struct {
	c unsafe.Pointer
}

func deleteArbCore(ac *ArbCore) {
	C.deleteArbCore(ac.c)
}

func NewArbCore(c unsafe.Pointer) *ArbCore {
	ac := &ArbCore{c: c}
	runtime.SetFinalizer(ac, deleteArbCore)
	return ac
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
	return toByteInboxArray(result.slice)
}

func (ac *ArbCore) GetInboxDelta(startIndex *big.Int, count *big.Int) (*big.Int, error) {
	result := C.arbCoreGetInboxDelta(ac.c, intToData(startIndex), intToData(count))
	if result.found == 0 {
		return nil, errors.New("failed to get inbox delta")
	}
	return dataToInt(result.value), nil
}

func (ac *ArbCore) GetSendAcc(startHash *big.Int, startIndex *big.Int, count *big.Int) (*big.Int, error) {
	result := C.arbCoreGetSendAcc(ac.c, intToData(startHash), intToData(startIndex), intToData(count))
	if result.found == 0 {
		return nil, errors.New("failed to get inbox delta")
	}
	return dataToInt(result.value), nil
}

func (ac *ArbCore) GetLogAcc(startHash *big.Int, startIndex *big.Int, count *big.Int, valueCache ValueCache) (*big.Int, error) {
	result := C.arbCoreGetLogAcc(ac.c, intToData(startHash), intToData(startIndex), intToData(count), valueCache.c)
	if result.found == 0 {
		return nil, errors.New("failed to get inbox delta")
	}
	return dataToInt(result.value), nil
}

func toByteInboxArray(sliceArray C.ByteSliceArray) ([]inbox.InboxMessage, error) {
	defer C.free(unsafe.Pointer(sliceArray.data))
	dataSlices := (*[1 << 30]C.struct_ByteSliceStruct)(unsafe.Pointer(sliceArray.data))[:sliceArray.length:sliceArray.length]
	messages := make([]inbox.InboxMessage, sliceArray.length)
	for i := range dataSlices {
		var err error
		messages[i], err = inbox.NewInboxMessageFromData(dataSlices[i])
		if err != nil {
			return nil, err
		}
	}
	return messages, nil
}
