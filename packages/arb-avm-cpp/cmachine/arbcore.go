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
#include "../cavm/carbcore.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"bytes"
	"math/big"
	"runtime"
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

func (ac *ArbCore) MachineMessagesRead() *big.Int {
	return receiveBigInt(C.arbCoreMachineMessagesRead(ac.c))
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

func sequencerBatchItemsToByteSliceArray(batchItems []inbox.SequencerBatchItem) C.struct_ByteSliceArrayStruct {
	return bytesArrayToByteSliceArray(encodeSequencerBatchItems(batchItems))
}

func delayedMessagesToByteSliceArray(delayedMessages []inbox.DelayedMessage) C.struct_ByteSliceArrayStruct {
	return bytesArrayToByteSliceArray(encodeDelayedMessages(delayedMessages))
}

func u256ArrayToByteSliceArray(nums []*big.Int) C.struct_ByteSliceArrayStruct {
	var bytes [][]byte
	for _, num := range nums {
		bytes = append(bytes, math.U256Bytes(num))
	}
	return bytesArrayToByteSliceArray(bytes)
}

func (ac *ArbCore) DeliverMessages(previousMessageCount *big.Int, previousSeqBatchAcc common.Hash, seqBatchItems []inbox.SequencerBatchItem, delayedMessages []inbox.DelayedMessage, reorgSeqBatchItemCount *big.Int) bool {
	previousMessageCountPtr := unsafeDataPointer(math.U256Bytes(previousMessageCount))
	previousSeqBatchAccPtr := unsafeDataPointer(previousSeqBatchAcc.Bytes())
	seqBatchItemsSlice := sequencerBatchItemsToByteSliceArray(seqBatchItems)
	delayedMessagesSlice := delayedMessagesToByteSliceArray(delayedMessages)

	var cReorgSeqBatchItemCount unsafe.Pointer
	if reorgSeqBatchItemCount != nil {
		reorgSeqBatchItemCount := math.U256Bytes(reorgSeqBatchItemCount)
		cReorgSeqBatchItemCount = unsafeDataPointer(reorgSeqBatchItemCount)
	}

	status := C.arbCoreDeliverMessages(ac.c, previousMessageCountPtr, previousSeqBatchAccPtr, seqBatchItemsSlice, delayedMessagesSlice, cReorgSeqBatchItemCount)
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

func (ac *ArbCore) GetDelayedMessageCount() (*big.Int, error) {
	result := C.arbCoreGetDelayedMessageCount(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to load send count")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetTotalDelayedMessagesSequenced() (*big.Int, error) {
	result := C.arbCoreGetTotalDelayedMessagesSequenced(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to load send count")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetDelayedMessagesToSequence(maxBlock *big.Int) (*big.Int, error) {
	maxBlockData := math.U256Bytes(maxBlock)
	result := C.arbCoreGetDelayedMessagesToSequence(ac.c, unsafeDataPointer(maxBlockData))
	if result.found == 0 {
		return nil, errors.New("failed to load delayed messages to sequence")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error) {
	startIndexData := math.U256Bytes(startIndex)
	countData := math.U256Bytes(count)
	result := C.arbCoreGetSends(ac.c, unsafeDataPointer(startIndexData), unsafeDataPointer(countData))
	if result.found == 0 {
		return nil, errors.New("failed to get sends")
	}

	return receiveByteSliceArray(result.array), nil
}

func (ac *ArbCore) GetLogs(startIndex *big.Int, count *big.Int) ([]value.Value, error) {
	if count.Cmp(big.NewInt(0)) == 0 {
		return nil, nil
	}
	startIndexData := math.U256Bytes(startIndex)
	countData := math.U256Bytes(count)
	result := C.arbCoreGetLogs(ac.c, unsafeDataPointer(startIndexData), unsafeDataPointer(countData))
	if result.found == 0 {
		return nil, errors.New("failed to get logs")
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
		return nil, errors.New("failed to get messages")
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

func (ac *ArbCore) GetSequencerBatchItems(startIndex *big.Int) ([]inbox.SequencerBatchItem, error) {
	startIndexData := math.U256Bytes(startIndex)

	result := C.arbCoreGetSequencerBatchItems(ac.c, unsafeDataPointer(startIndexData))
	if result.found == 0 {
		return nil, errors.New("failed to get messages")
	}

	data := receiveByteSliceArray(result.array)
	items := make([]inbox.SequencerBatchItem, len(data))
	for i, slice := range data {
		var err error
		items[i], err = inbox.NewSequencerBatchItemFromData(slice)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (ac *ArbCore) GetSequencerBlockNumberAt(index *big.Int) (*big.Int, error) {
	indexData := math.U256Bytes(index)

	res := C.arbCoreGetSequencerBlockNumberAt(ac.c, unsafeDataPointer(indexData))
	if res.found == 0 {
		return nil, errors.Errorf("failed to get sequencer block number for %v", index)
	}

	return receiveBigInt(res.value), nil
}

func (ac *ArbCore) GenInboxProof(seqNum *big.Int, batchIndex *big.Int, batchEndCount *big.Int) ([]byte, error) {
	seqNumData := math.U256Bytes(seqNum)
	batchEndCountData := math.U256Bytes(batchEndCount)
	batchIndexData := math.U256Bytes(batchIndex)

	res := C.arbCoreGenInboxProof(ac.c, unsafeDataPointer(seqNumData), unsafeDataPointer(batchIndexData), unsafeDataPointer(batchEndCountData))
	if res.found == 0 {
		return nil, errors.Errorf("failed to generate inbox proof for %v", seqNum)
	}

	return receiveByteSlice(res.slice), nil
}

func (ac *ArbCore) GetInboxAcc(index *big.Int) (ret common.Hash, err error) {
	startIndexData := math.U256Bytes(index)

	status := C.arbCoreGetInboxAcc(ac.c, unsafeDataPointer(startIndexData), unsafe.Pointer(&ret[0]))
	if status == 0 {
		err = errors.Errorf("failed to get inbox acc for %v", index)
	}

	return
}

func (ac *ArbCore) GetDelayedInboxAcc(index *big.Int) (ret common.Hash, err error) {
	startIndexData := math.U256Bytes(index)

	status := C.arbCoreGetDelayedInboxAcc(ac.c, unsafeDataPointer(startIndexData), unsafe.Pointer(&ret[0]))
	if status == 0 {
		err = errors.Errorf("failed to get delayed inbox acc for %v", index)
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

func (ac *ArbCore) CountMatchingBatchAccs(lastSeqNums []*big.Int, accs []common.Hash) (ret int, err error) {
	if len(lastSeqNums) != len(accs) {
		return -1, errors.New("mismatching lengths when counting matching batches")
	}
	bytes := make([]byte, 0, len(lastSeqNums)*64)
	for i := 0; i < len(lastSeqNums); i++ {
		bytes = append(bytes, math.U256Bytes(lastSeqNums[i])...)
		bytes = append(bytes, accs[i].Bytes()...)
	}
	ret = int(C.arbCoreCountMatchingBatchAccs(ac.c, toByteSliceView(bytes)))
	if ret < 0 {
		err = errors.New("failed to get matching batch accs")
	}

	return
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
		return errors.Errorf("unsupported execution cursor type %T", executionCursor)
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

func (ac *ArbCore) GetLastMachine() (machine.Machine, error) {
	cMachine := C.arbCoreGetLastMachine(ac.c)
	if cMachine == nil {
		return nil, errors.Errorf("error getting last machine")
	}
	ret := &Machine{cMachine}

	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (ac *ArbCore) GetLastMachineTotalGas() (*big.Int, error) {
	result := C.arbCoreGetLastMachineTotalGas(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to get last machine total gas")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) TakeMachine(executionCursor core.ExecutionCursor) (machine.Machine, error) {
	cursor, ok := executionCursor.(*ExecutionCursor)
	if !ok {
		return nil, errors.Errorf("unsupported execution cursor type %T", executionCursor)
	}
	cMachine := C.arbCoreTakeMachine(ac.c, cursor.c)
	if cMachine == nil {
		return nil, errors.Errorf("error taking machine from execution cursor")
	}
	ret := &Machine{cMachine}

	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (ac *ArbCore) LogsCursorPosition(cursorIndex *big.Int) (*big.Int, error) {
	cursorIndexData := math.U256Bytes(cursorIndex)
	result := C.arbCoreLogsCursorGetPosition(ac.c, unsafeDataPointer(cursorIndexData))
	if result.found == 0 {
		return nil, errors.New("failed to load logs cursor position")
	}

	return receiveBigInt(result.value), nil
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

func (ac *ArbCore) LogsCursorGetLogs(cursorIndex *big.Int) (*big.Int, []value.Value, []value.Value, error) {
	cursorIndexData := math.U256Bytes(cursorIndex)
	result := C.arbCoreLogsCursorGetLogs(ac.c, unsafeDataPointer(cursorIndexData))
	if result.found == 0 {
		err := ac.LogsCursorCheckError(cursorIndex)
		if err != nil {
			return nil, nil, nil, err
		}

		// Nothing found, try again later
		return nil, nil, nil, nil
	}

	firstIndex := receiveBigInt(result.first_index)
	data := receiveByteSliceArray(result.first_array)
	logs := make([]value.Value, len(data))
	for i, slice := range data {
		var err error
		logs[i], err = value.UnmarshalValue(bytes.NewReader(slice[:]))
		if err != nil {
			return nil, nil, nil, err
		}
	}

	deletedData := receiveByteSliceArray(result.second_array)
	deletedLogs := make([]value.Value, len(deletedData))
	for i, slice := range deletedData {
		var err error
		deletedLogs[i], err = value.UnmarshalValue(bytes.NewReader(slice[:]))
		if err != nil {
			return nil, nil, nil, err
		}
	}

	if len(logs) == 0 && len(deletedLogs) == 0 {
		return nil, nil, nil, errors.New("logs cursor missing response")
	}

	return firstIndex, logs, deletedLogs, nil
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

func (ac *ArbCore) GetMachineForSideload(blockNumber uint64, allowSlowLookup bool) (machine.Machine, error) {
	CallowSlowLookup := 0
	if allowSlowLookup {
		CallowSlowLookup = 1
	}
	cMachineResult := C.arbCoreGetMachineForSideload(ac.c, C.uint64_t(blockNumber), C.int(CallowSlowLookup))

	if cMachineResult.slow_error == 1 {
		return nil, errors.Errorf("missing trie node 0000000000000000000000000000000000000000000000000000000000000000 (path )")
	}

	if cMachineResult.machine == nil {
		return nil, errors.Errorf("error getting machine for sideload")
	}

	return WrapCMachine(cMachineResult.machine), nil
}
