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
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"
	"runtime"
	"unsafe"

	"github.com/pkg/errors"
)

type ArbCore struct {
	c unsafe.Pointer
}

func NewArbCore(c unsafe.Pointer) *ArbCore {
	// ArbCore has same lifetime as ArbStorage, no need to have finalizer
	return &ArbCore{c: c}
}

func (ac *ArbCore) StartThread() bool {
	status := C.arbCoreStartThread(ac.c)
	if status == 0 {
		return false
	}
	return true
}

func (ac *ArbCore) StopThread() {
	C.arbCoreAbortThread(ac.c)
}

func (ac *ArbCore) DeliverMessages(messages []inbox.InboxMessage, previousInboxHash *big.Int) {
	cPreviousInboxHash := intToData(previousInboxHash)
	defer C.free(cPreviousInboxHash)

	msgDataC := C.CBytes(encodeInboxMessages(messages))

	C.arbCoreDeliverMessages(ac.c, msgDataC, cPreviousInboxHash)
}

func (ac *ArbCore) GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error) {
	cStartIndex := intToData(startIndex)
	defer C.free(cStartIndex)
	cCount := intToData(count)
	defer C.free(cCount)

	result := C.arbCoreGetSends(ac.c, cStartIndex, cCount)
	if result.found == 0 {
		return nil, errors.New("failed to get log")
	}
	defer freeEncodedByteSliceArray(result.array)

	return toByteSliceArray(result.array), nil
}

func (ac *ArbCore) GetMessages(startIndex *big.Int, count *big.Int) ([]inbox.InboxMessage, error) {
	cStartIndex := intToData(startIndex)
	defer C.free(cStartIndex)
	cCount := intToData(count)
	defer C.free(cCount)

	result := C.arbCoreGetMessages(ac.c, cStartIndex, cCount)
	if result.found == 0 {
		return nil, errors.New("failed to get log")
	}
	defer freeEncodedByteSliceArray(result.array)

	data := toByteSliceArray(result.array)
	messages := make([]inbox.InboxMessage, len(data))
	for i, slice := range data {
		var err error
		messages[i], err = inbox.NewInboxMessageFromData(slice)
		if err != nil {
			return nil, err
		}
	}

	return messages, nil
}

func (ac *ArbCore) GetInboxDelta(startIndex *big.Int, count *big.Int) (ret common.Hash, err error) {
	cStartIndex := intToData(startIndex)
	defer C.free(cStartIndex)
	cCount := intToData(count)
	defer C.free(cCount)

	status := C.arbCoreGetInboxDelta(ac.c, cStartIndex, cCount, unsafe.Pointer(&ret[0]))
	if status == 0 {
		err = errors.New("failed to get inbox delta")
		return
	}

	return
}

func (ac *ArbCore) GetInboxAcc(index *big.Int) (ret common.Hash, err error) {
	cIndex := intToData(index)
	defer C.free(cIndex)

	status := C.arbCoreGetInboxAcc(ac.c, cIndex, unsafe.Pointer(&ret[0]))
	if status == 0 {
		err = errors.New("failed to get inbox delta")
	}

	return
}

func (ac *ArbCore) GetSendAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (ret common.Hash, err error) {
	cStartAcc := hashToData(startAcc)
	defer C.free(cStartAcc)
	cStartIndex := intToData(startIndex)
	defer C.free(cStartIndex)
	cCount := intToData(count)
	defer C.free(cCount)

	status := C.arbCoreGetSendAcc(ac.c, cStartAcc, cStartIndex, cCount, unsafe.Pointer(&ret[0]))
	if status == 0 {
		err = errors.New("failed to get inbox delta")
	}

	return
}

func (ac *ArbCore) GetLogAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (ret common.Hash, err error) {
	cStartAcc := hashToData(startAcc)
	defer C.free(cStartAcc)
	cStartIndex := intToData(startIndex)
	defer C.free(cStartIndex)
	cCount := intToData(count)
	defer C.free(cCount)

	status := C.arbCoreGetLogAcc(ac.c, cStartAcc, cStartIndex, cCount, unsafe.Pointer(&ret[0]))
	if status == 0 {
		err = errors.New("failed to get inbox delta")
	}

	return
}

func (ac *ArbCore) GetExecutionCursor(totalGasUsed *big.Int) (*ExecutionCursor, error) {
	cTotalGasUsed := intToData(totalGasUsed)
	defer C.free(cTotalGasUsed)

	cExecutionCursor := C.arbCoreGetExecutionCursor(ac.c, cTotalGasUsed)

	if cExecutionCursor == nil {
		return nil, errors.Errorf("error creating execution cursor")
	}
	ret, err := NewExecutionCursor(cExecutionCursor)
	if err != nil {
		return nil, errors.Errorf("Error call NewExecutionCursor")
	}
	runtime.SetFinalizer(ret, deleteExecutionCursor)
	return ret, nil
}

func (ac *ArbCore) AdvanceExecutionCursor(executionCursor core.ExecutionCursor, maxGas *big.Int, goOverGas bool) error {
	cursor, ok := executionCursor.(*ExecutionCursor)
	if !ok {
		return errors.New("unsupported execution cursor type")
	}
	cMaxGas := intToData(maxGas)
	defer C.free(cMaxGas)

	goOverGasInt := 0
	if goOverGas {
		goOverGasInt = 1
	}

	status := C.arbCoreAdvanceExecutionCursor(ac.c, cursor.c, cMaxGas, C.int(goOverGasInt))
	if status == 0 {
		return errors.New("failed to advance")
	}

	return cursor.updateValues()
}

func (ac *ArbCore) LogsCursorRequest(count *big.Int) error {
	cCount := intToData(count)
	defer C.free(cCount)

	status := C.arbCoreLogsCursorRequest(ac.c, cCount)
	if status == 0 {
		return errors.New("failed to send logs cursor request")
	}

	return nil
}

func (ac *ArbCore) LogsCursorGetLogs() ([]value.Value, error) {
	result := C.arbCoreLogsCursorGetLogs(ac.c)
	if result.found == 0 {
		// Nothing found, try again later
		return nil, nil
	}
	defer freeEncodedByteSliceArray(result.array)

	data := toByteSliceArray(result.array)
	logs := make([]value.Value, len(data))
	for i, slice := range data {
		var err error
		logs[i], err = value.UnmarshalValue(bytes.NewReader(slice[:]))
		if err != nil {
			return nil, err
		}
	}
	return logs, nil
}

func (ac *ArbCore) LogsCursorSetNextIndex(count *big.Int) error {
	cCount := intToData(count)
	defer C.free(cCount)

	status := C.arbCoreLogsCursorSetNextIndex(ac.c, cCount)
	if status == 0 {
		return errors.New("failed to send logs cursor set next index")
	}

	return nil
}

func (ac *ArbCore) LogsCursorCheckError() bool {
	status := C.arbCoreLogsCursorCheckError(ac.c)
	if status == 0 {
		return false
	}

	return true
}

func (ac *ArbCore) LogsCursorClearError() (string, error) {
	cStr := C.arbCoreLogsCursorClearError(ac.c)
	if cStr == nil {
		return "", errors.New("no error string present")
	}
	defer C.free(unsafe.Pointer(cStr))

	return C.GoString(cStr), nil
}
