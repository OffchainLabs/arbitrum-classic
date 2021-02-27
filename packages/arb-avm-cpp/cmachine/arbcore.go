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
#include "../cavm/carbcore.h"
#include "../cavm/cvaluecache.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"bytes"
	"math/big"
	"unsafe"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"

	"github.com/pkg/errors"
)

type ArbCore struct {
	c       unsafe.Pointer
	storage *ArbStorage
}

func NewArbCore(c unsafe.Pointer, storage *ArbStorage) *ArbCore {
	// ArbCore has same lifetime as ArbStorage, no need to have finalizer
	// Keeping a reference to ArbStorage makes sure that ArbCore isn't
	// destroyed too early, as ArbStorage owns ArbCore, not this struct
	return &ArbCore{c: c, storage: storage}
}

func (ac *ArbCore) StartThread() bool {
	status := C.arbCoreStartThread(ac.c)
	return status == 1
}

func (ac *ArbCore) StopThread() {
	C.arbCoreAbortThread(ac.c)
}

func (ac *ArbCore) MachineIdle() bool {
	status := C.arbCoreMachineIdle(ac.c)
	return status == 1
}

func (ac *ArbCore) MessagesStatus() (core.MessageStatus, error) {
	statusRaw := C.arbCoreMessagesStatus(ac.c)
	status := core.MessageStatus(int(statusRaw))
	if status == core.MessagesError {
		cStr := C.arbCoreMessagesClearError(ac.c)
		defer C.free(unsafe.Pointer(cStr))
		return core.MessagesError, errors.New(C.GoString(cStr))
	}
	return status, nil
}

func (ac *ArbCore) DeliverMessages(messages []inbox.InboxMessage, previousInboxAcc common.Hash, lastBlockComplete bool) bool {
	rawInboxData := encodeInboxMessages(messages)
	byteSlices := encodeByteSliceList(rawInboxData)

	sliceArrayData := C.malloc(C.size_t(C.sizeof_struct_ByteSliceStruct * len(byteSlices)))
	sliceArray := (*[1 << 30]C.struct_ByteSliceStruct)(sliceArrayData)[:len(byteSlices):len(byteSlices)]
	for i, data := range byteSlices {
		sliceArray[i] = data
	}
	defer C.free(sliceArrayData)
	msgData := C.struct_ByteSliceArrayStruct{slices: sliceArrayData, count: C.int(len(byteSlices))}

	cLastBlockComplete := 0
	if lastBlockComplete {
		cLastBlockComplete = 1
	}

	status := C.arbCoreDeliverMessages(ac.c, msgData, unsafeDataPointer(previousInboxAcc.Bytes()), C.int(cLastBlockComplete))
	return status == 1
}

func (ac *ArbCore) GetSendCount() (*big.Int, error) {
	result := C.arbCoreGetSendCount(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to load send count")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetLogCount() (*big.Int, error) {
	result := C.arbCoreGetLogCount(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to load log count")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetMessageCount() (*big.Int, error) {
	result := C.arbCoreGetMessageCount(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to load send count")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error) {
	startIndexData := math.U256Bytes(startIndex)
	countData := math.U256Bytes(count)
	result := C.arbCoreGetSends(ac.c, unsafeDataPointer(startIndexData), unsafeDataPointer(countData))
	if result.found == 0 {
		return nil, errors.New("failed to get log")
	}

	return receiveByteSliceArray(result.array), nil
}

func (ac *ArbCore) GetLogs(startIndex *big.Int, count *big.Int) ([]value.Value, error) {
	startIndexData := math.U256Bytes(startIndex)
	countData := math.U256Bytes(count)
	result := C.arbCoreGetLogs(ac.c, unsafeDataPointer(startIndexData), unsafeDataPointer(countData))
	if result.found == 0 {
		return nil, errors.New("failed to get log")
	}

	marshaledValues := receiveByteSliceArray(result.array)
	logVals := make([]value.Value, 0, len(marshaledValues))
	for _, marshaledValue := range marshaledValues {
		val, err := value.UnmarshalValue(bytes.NewReader(marshaledValue))
		if err != nil {
			return nil, err
		}
		logVals = append(logVals, val)
	}
	return logVals, nil
}

func (ac *ArbCore) GetMessages(startIndex *big.Int, count *big.Int) ([]inbox.InboxMessage, error) {
	startIndexData := math.U256Bytes(startIndex)
	countData := math.U256Bytes(count)

	result := C.arbCoreGetMessages(ac.c, unsafeDataPointer(startIndexData), unsafeDataPointer(countData))
	if result.found == 0 {
		return nil, errors.New("failed to get log")
	}

	data := receiveByteSliceArray(result.array)
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

func (ac *ArbCore) GetInboxAcc(index *big.Int) (ret common.Hash, err error) {
	startIndexData := math.U256Bytes(index)

	status := C.arbCoreGetInboxAcc(ac.c, unsafeDataPointer(startIndexData), unsafe.Pointer(&ret[0]))
	if status == 0 {
		err = errors.New("failed to get inbox acc")
	}

	return
}

func (ac *ArbCore) GetInboxAccPair(index1 *big.Int, index2 *big.Int) (ret1 common.Hash, ret2 common.Hash, err error) {
	startIndex1Data := math.U256Bytes(index1)
	startIndex2Data := math.U256Bytes(index2)

	status := C.arbCoreGetInboxAccPair(ac.c, unsafeDataPointer(startIndex1Data), unsafeDataPointer(startIndex2Data), unsafe.Pointer(&ret1[0]), unsafe.Pointer(&ret2[0]))
	if status == 0 {
		err = errors.New("failed to get inbox acc")
	}

	return
}

func (ac *ArbCore) GetSendAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (ret common.Hash, err error) {
	startIndexData := math.U256Bytes(startIndex)
	countData := math.U256Bytes(count)

	status := C.arbCoreGetSendAcc(
		ac.c,
		unsafeDataPointer(startAcc.Bytes()),
		unsafeDataPointer(startIndexData),
		unsafeDataPointer(countData),
		unsafe.Pointer(&ret[0]),
	)
	if status == 0 {
		err = errors.New("failed to get send acc")
	}

	return
}

func (ac *ArbCore) GetLogAcc(startAcc common.Hash, startIndex *big.Int, count *big.Int) (ret common.Hash, err error) {
	startIndexData := math.U256Bytes(startIndex)
	countData := math.U256Bytes(count)

	status := C.arbCoreGetLogAcc(
		ac.c,
		unsafeDataPointer(startAcc.Bytes()),
		unsafeDataPointer(startIndexData),
		unsafeDataPointer(countData),
		unsafe.Pointer(&ret[0]),
	)
	if status == 0 {
		err = errors.New("failed to get log acc")
	}

	return
}

func (ac *ArbCore) GetInboxHash(index *big.Int) (ret common.Hash, err error) {
	panic("unimplemented method")
}

func (ac *ArbCore) GetExecutionCursor(totalGasUsed *big.Int) (core.ExecutionCursor, error) {
	totalGasUsedData := math.U256Bytes(totalGasUsed)

	cExecutionCursor := C.arbCoreGetExecutionCursor(ac.c, unsafeDataPointer(totalGasUsedData))

	if cExecutionCursor == nil {
		return nil, errors.Errorf("error creating execution cursor")
	}
	return NewExecutionCursor(cExecutionCursor)
}

func (ac *ArbCore) AdvanceExecutionCursor(executionCursor core.ExecutionCursor, maxGas *big.Int, goOverGas bool) error {
	cursor, ok := executionCursor.(*ExecutionCursor)
	if !ok {
		return errors.New("unsupported execution cursor type")
	}
	maxGasData := math.U256Bytes(maxGas)

	goOverGasInt := 0
	if goOverGas {
		goOverGasInt = 1
	}

	status := C.arbCoreAdvanceExecutionCursor(ac.c, cursor.c, unsafeDataPointer(maxGasData), C.int(goOverGasInt))
	if status == 0 {
		return errors.New("failed to advance")
	}

	return cursor.updateValues()
}

func (ac *ArbCore) LogsCursorRequest(cursorIndex *big.Int, count *big.Int) error {
	cursorIndexData := math.U256Bytes(cursorIndex)
	countData := math.U256Bytes(count)

	status := C.arbCoreLogsCursorRequest(ac.c, unsafeDataPointer(cursorIndexData), unsafeDataPointer(countData))
	if status == 0 {
		err := ac.LogsCursorCheckError(cursorIndex)
		if err != nil {
			return err
		}

		return errors.New("failed to send logs cursor request")
	}

	return nil
}

func (ac *ArbCore) LogsCursorGetLogs(cursorIndex *big.Int) (*big.Int, []value.Value, error) {
	cursorIndexData := math.U256Bytes(cursorIndex)
	result := C.arbCoreLogsCursorGetLogs(ac.c, unsafeDataPointer(cursorIndexData))
	if result.found == 0 {
		err := ac.LogsCursorCheckError(cursorIndex)
		if err != nil {
			return nil, nil, err
		}

		// Nothing found, try again later
		return nil, nil, nil
	}

	firstIndex := receiveBigInt(result.first_index)
	data := receiveByteSliceArray(result.array)
	logs := make([]value.Value, len(data))
	for i, slice := range data {
		var err error
		logs[i], err = value.UnmarshalValue(bytes.NewReader(slice[:]))
		if err != nil {
			return nil, nil, err
		}
	}

	return firstIndex, logs, nil
}

func (ac *ArbCore) LogsCursorGetDeletedLogs(cursorIndex *big.Int) (*big.Int, []value.Value, error) {
	cursorIndexData := math.U256Bytes(cursorIndex)
	result := C.arbCoreLogsCursorGetDeletedLogs(ac.c, unsafeDataPointer(cursorIndexData))
	if result.found == 0 {
		err := ac.LogsCursorCheckError(cursorIndex)
		if err != nil {
			return nil, nil, err
		}

		// Nothing found, try again later
		return nil, nil, nil
	}

	firstIndex := receiveBigInt(result.first_index)
	data := receiveByteSliceArray(result.array)
	logs := make([]value.Value, len(data))
	for i, slice := range data {
		var err error
		logs[i], err = value.UnmarshalValue(bytes.NewReader(slice[:]))
		if err != nil {
			return nil, nil, err
		}
	}
	return firstIndex, logs, nil
}

func (ac *ArbCore) LogsCursorCheckError(cursorIndex *big.Int) error {
	cursorIndexData := math.U256Bytes(cursorIndex)
	status := C.arbCoreLogsCursorCheckError(ac.c, unsafeDataPointer(cursorIndexData))
	if status == 0 {
		return nil
	}

	cStr := C.arbCoreLogsCursorClearError(ac.c, unsafeDataPointer(cursorIndexData))
	if cStr == nil {
		return errors.New("Error occurred but no error string present")
	}
	defer C.free(unsafe.Pointer(cStr))

	return errors.New(C.GoString(cStr))
}

func (ac *ArbCore) LogsCursorConfirmReceived(cursorIndex *big.Int) (bool, error) {
	cursorIndexData := math.U256Bytes(cursorIndex)
	status := C.arbCoreLogsCursorConfirmReceived(ac.c, unsafeDataPointer(cursorIndexData))
	if status == 0 {
		err := ac.LogsCursorCheckError(cursorIndex)
		if err != nil {
			return false, err
		}

		// Still have more logs to retrieve
		return false, nil
	}

	return true, nil
}

func (ac *ArbCore) GetMachineForSideload(blockNumber uint64) (machine.Machine, error) {
	cMachine := C.arbCoreGetMachineForSideload(ac.c, C.uint64_t(blockNumber))

	if cMachine == nil {
		return nil, errors.Errorf("error getting machine for sideload")
	}

	return WrapCMachine(cMachine), nil
}
