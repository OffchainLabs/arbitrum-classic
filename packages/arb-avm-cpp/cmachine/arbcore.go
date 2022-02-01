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

const slowLookupErrorString = "missing trie node 0000000000000000000000000000000000000000000000000000000000000000 (path )"

func NewArbCore(c unsafe.Pointer, storage *ArbStorage) *ArbCore {
	// ArbCore has same lifetime as ArbStorage, no need to have finalizer
	// Keeping a reference to ArbStorage makes sure that ArbCore isn't
	// destroyed too early, as ArbStorage owns ArbCore, not this struct
	return &ArbCore{c: c, storage: storage}
}

func (ac *ArbCore) StartThread() bool {
	defer runtime.KeepAlive(ac)
	status := C.arbCoreStartThread(ac.c)
	return status == 1
}

func (ac *ArbCore) StopThread() {
	defer runtime.KeepAlive(ac)
	C.arbCoreAbortThread(ac.c)
}

func (ac *ArbCore) MachineIdle() bool {
	defer runtime.KeepAlive(ac)
	status := C.arbCoreMachineIdle(ac.c)
	return status == 1
}

func (ac *ArbCore) SaveRocksdbCheckpoint() {
	defer runtime.KeepAlive(ac)
	C.arbCoreSaveRocksdbCheckpoint(ac.c)
}

func (ac *ArbCore) MachineMessagesRead() *big.Int {
	defer runtime.KeepAlive(ac)
	return receiveBigInt(C.arbCoreMachineMessagesRead(ac.c))
}

func (ac *ArbCore) MessagesStatus() (core.MessageStatus, error) {
	defer runtime.KeepAlive(ac)
	statusRaw := C.arbCoreMessagesStatus(ac.c)
	status := core.MessageStatus(int(statusRaw))
	if status == core.MessagesError {
		cStr := C.arbCoreMessagesClearError(ac.c)
		defer C.free(unsafe.Pointer(cStr))
		return core.MessagesError, errors.New(C.GoString(cStr))
	}
	return status, nil
}

func (ac *ArbCore) PrintCoreThreadBacktrace() {
	defer runtime.KeepAlive(ac)
	C.arbCorePrintCoreThreadBacktrace(ac.c)
}

// Note: the slices field of the returned struct needs manually freed by C.free
func sequencerBatchItemsToByteSliceArray(batchItems []inbox.SequencerBatchItem) C.struct_ByteSliceArrayStruct {
	return bytesArrayToByteSliceArray(encodeSequencerBatchItems(batchItems))
}

// Note: the slices field of the returned struct needs manually freed by C.free
func delayedMessagesToByteSliceArray(delayedMessages []inbox.DelayedMessage) C.struct_ByteSliceArrayStruct {
	return bytesArrayToByteSliceArray(encodeDelayedMessages(delayedMessages))
}

func (ac *ArbCore) DeliverMessages(previousMessageCount *big.Int, previousSeqBatchAcc common.Hash, seqBatchItems []inbox.SequencerBatchItem, delayedMessages []inbox.DelayedMessage, reorgSeqBatchItemCount *big.Int) bool {
	defer runtime.KeepAlive(ac)
	previousMessageCountPtr := unsafeDataPointer(math.U256Bytes(previousMessageCount))
	previousSeqBatchAccPtr := unsafeDataPointer(previousSeqBatchAcc.Bytes())
	seqBatchItemsSlice := sequencerBatchItemsToByteSliceArray(seqBatchItems)
	defer freeByteSliceArray(seqBatchItemsSlice)
	delayedMessagesSlice := delayedMessagesToByteSliceArray(delayedMessages)
	defer freeByteSliceArray(delayedMessagesSlice)

	var cReorgSeqBatchItemCount unsafe.Pointer
	if reorgSeqBatchItemCount != nil {
		reorgSeqBatchItemCount := math.U256Bytes(reorgSeqBatchItemCount)
		cReorgSeqBatchItemCount = unsafeDataPointer(reorgSeqBatchItemCount)
	}

	status := C.arbCoreDeliverMessages(ac.c, previousMessageCountPtr, previousSeqBatchAccPtr, seqBatchItemsSlice, delayedMessagesSlice, cReorgSeqBatchItemCount)
	return status == 1
}

func (ac *ArbCore) GetSendCount() (*big.Int, error) {
	defer runtime.KeepAlive(ac)
	result := C.arbCoreGetSendCount(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to load send count")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetLogCount() (*big.Int, error) {
	defer runtime.KeepAlive(ac)
	result := C.arbCoreGetLogCount(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to load log count")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetMessageCount() (*big.Int, error) {
	defer runtime.KeepAlive(ac)
	result := C.arbCoreGetMessageCount(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to load send count")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetDelayedMessageCount() (*big.Int, error) {
	defer runtime.KeepAlive(ac)
	result := C.arbCoreGetDelayedMessageCount(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to load send count")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetTotalDelayedMessagesSequenced() (*big.Int, error) {
	defer runtime.KeepAlive(ac)
	result := C.arbCoreGetTotalDelayedMessagesSequenced(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to load send count")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetDelayedMessagesToSequence(maxBlock *big.Int) (*big.Int, error) {
	defer runtime.KeepAlive(ac)
	maxBlockData := math.U256Bytes(maxBlock)
	result := C.arbCoreGetDelayedMessagesToSequence(ac.c, unsafeDataPointer(maxBlockData))
	if result.found == 0 {
		return nil, errors.New("failed to load delayed messages to sequence")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) GetSends(startIndex *big.Int, count *big.Int) ([][]byte, error) {
	defer runtime.KeepAlive(ac)
	startIndexData := math.U256Bytes(startIndex)
	countData := math.U256Bytes(count)
	result := C.arbCoreGetSends(ac.c, unsafeDataPointer(startIndexData), unsafeDataPointer(countData))
	if result.found == 0 {
		return nil, errors.New("failed to get sends")
	}

	return receiveByteSliceArray(result.array), nil
}

func unmarshalLog(marshaled []byte) (core.ValueAndInbox, error) {
	reader := bytes.NewReader(marshaled)
	var inboxData [64]byte
	_, err := reader.Read(inboxData[:])
	if err != nil {
		return core.ValueAndInbox{}, err
	}
	inboxCount := new(big.Int).SetBytes(inboxData[:32])
	var inboxAccumulator common.Hash
	copy(inboxAccumulator[:], inboxData[32:])

	val, err := value.UnmarshalValue(reader)
	if err != nil {
		return core.ValueAndInbox{}, err
	}

	return core.ValueAndInbox{
		Value: val,
		Inbox: core.InboxState{
			Count:       inboxCount,
			Accumulator: inboxAccumulator,
		},
	}, nil
}

func (ac *ArbCore) GetLogs(startIndex *big.Int, count *big.Int) ([]core.ValueAndInbox, error) {
	defer runtime.KeepAlive(ac)
	if count.Cmp(big.NewInt(0)) == 0 {
		return nil, nil
	}
	startIndexData := math.U256Bytes(startIndex)
	countData := math.U256Bytes(count)
	result := C.arbCoreGetLogs(ac.c, unsafeDataPointer(startIndexData), unsafeDataPointer(countData))
	if result.found == 0 {
		return nil, errors.New("failed to get logs")
	}

	marshaledLogs := receiveByteSliceArray(result.array)
	logVals := make([]core.ValueAndInbox, 0, len(marshaledLogs))
	for _, marshaledLog := range marshaledLogs {
		log, err := unmarshalLog(marshaledLog)
		if err != nil {
			return nil, err
		}
		logVals = append(logVals, log)
	}
	return logVals, nil
}

func (ac *ArbCore) GetMessages(startIndex *big.Int, count *big.Int) ([]inbox.InboxMessage, error) {
	defer runtime.KeepAlive(ac)
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
	defer runtime.KeepAlive(ac)
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
	defer runtime.KeepAlive(ac)
	indexData := math.U256Bytes(index)

	res := C.arbCoreGetSequencerBlockNumberAt(ac.c, unsafeDataPointer(indexData))
	if res.found == 0 {
		return nil, errors.Errorf("failed to get sequencer block number for %v", index)
	}

	return receiveBigInt(res.value), nil
}

func (ac *ArbCore) GenInboxProof(seqNum *big.Int, batchIndex *big.Int, batchEndCount *big.Int) ([]byte, error) {
	defer runtime.KeepAlive(ac)
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
	defer runtime.KeepAlive(ac)
	startIndexData := math.U256Bytes(index)

	status := C.arbCoreGetInboxAcc(ac.c, unsafeDataPointer(startIndexData), unsafe.Pointer(&ret[0]))
	if status == 0 {
		err = errors.Errorf("failed to get inbox acc for %v", index)
	}

	return
}

func (ac *ArbCore) GetDelayedInboxAcc(index *big.Int) (ret common.Hash, err error) {
	defer runtime.KeepAlive(ac)
	startIndexData := math.U256Bytes(index)

	status := C.arbCoreGetDelayedInboxAcc(ac.c, unsafeDataPointer(startIndexData), unsafe.Pointer(&ret[0]))
	if status == 0 {
		err = errors.Errorf("failed to get delayed inbox acc for %v", index)
	}

	return
}

func (ac *ArbCore) GetInboxAccPair(index1 *big.Int, index2 *big.Int) (ret1 common.Hash, ret2 common.Hash, err error) {
	defer runtime.KeepAlive(ac)
	startIndex1Data := math.U256Bytes(index1)
	startIndex2Data := math.U256Bytes(index2)

	status := C.arbCoreGetInboxAccPair(ac.c, unsafeDataPointer(startIndex1Data), unsafeDataPointer(startIndex2Data), unsafe.Pointer(&ret1[0]), unsafe.Pointer(&ret2[0]))
	if status == 0 {
		err = errors.New("failed to get inbox acc")
	}

	return
}

func (ac *ArbCore) CountMatchingBatchAccs(lastSeqNums []*big.Int, accs []common.Hash) (ret int, err error) {
	defer runtime.KeepAlive(ac)
	if len(lastSeqNums) != len(accs) {
		return -1, errors.New("mismatching lengths when counting matching batches")
	}
	data := make([]byte, 0, len(lastSeqNums)*64)
	for i := 0; i < len(lastSeqNums); i++ {
		data = append(data, math.U256Bytes(lastSeqNums[i])...)
		data = append(data, accs[i].Bytes()...)
	}
	ret = int(C.arbCoreCountMatchingBatchAccs(ac.c, toByteSliceView(data)))
	if ret < 0 {
		err = errors.New("failed to get matching batch accs")
	}

	return
}

func (ac *ArbCore) GetExecutionCursor(totalGasUsed *big.Int, allowSlowLookup bool) (core.ExecutionCursor, error) {
	defer runtime.KeepAlive(ac)
	totalGasUsedData := math.U256Bytes(totalGasUsed)

	cExecutionCursor := C.arbCoreGetExecutionCursor(ac.c, unsafeDataPointer(totalGasUsedData), boolToCInt(allowSlowLookup))

	if cExecutionCursor == nil {
		return nil, errors.Errorf("error creating execution cursor")
	}
	return NewExecutionCursor(cExecutionCursor)
}

func (ac *ArbCore) AdvanceExecutionCursor(executionCursor core.ExecutionCursor, maxGas *big.Int, goOverGas bool, allowSlowLookup bool) error {
	defer runtime.KeepAlive(ac)
	defer runtime.KeepAlive(executionCursor)
	cursor, ok := executionCursor.(*ExecutionCursor)
	if !ok {
		return errors.Errorf("unsupported execution cursor type %T", executionCursor)
	}
	maxGasData := math.U256Bytes(maxGas)

	status := C.arbCoreAdvanceExecutionCursor(ac.c, cursor.c, unsafeDataPointer(maxGasData), boolToCInt(goOverGas), boolToCInt(allowSlowLookup))
	if status == 0 {
		return errors.New("failed to advance")
	}

	return cursor.updateValues()
}

func (ac *ArbCore) AdvanceExecutionCursorWithTracing(executionCursor core.ExecutionCursor, maxGas *big.Int, goOverGas bool, allowSlowLookup bool, logNumberStart, logNumberEnd *big.Int) ([]core.MachineEmission, error) {
	defer runtime.KeepAlive(ac)
	cursor, ok := executionCursor.(*ExecutionCursor)
	if !ok {
		return nil, errors.Errorf("unsupported execution cursor type %T", executionCursor)
	}
	maxGasData := math.U256Bytes(maxGas)
	logNumberStartData := math.U256Bytes(logNumberStart)
	logNumberEndData := math.U256Bytes(logNumberEnd)

	result := C.arbCoreAdvanceExecutionCursorWithTracing(
		ac.c,
		cursor.c,
		unsafeDataPointer(maxGasData),
		boolToCInt(goOverGas),
		boolToCInt(allowSlowLookup),
		unsafeDataPointer(logNumberStartData),
		unsafeDataPointer(logNumberEndData),
	)
	if result.found == 0 {
		return nil, errors.New("failed to advance cursor with tracing")
	}
	runtime.KeepAlive(cursor)

	valCount := uint64(result.data.count)
	rd := bytes.NewReader(receiveByteSlice(result.data.slice))
	vals := make([]core.MachineEmission, 0, valCount)
	for i := uint64(0); i < valCount; i++ {
		var logCountData [32]byte
		_, err := rd.Read(logCountData[:])
		if err != nil {
			return nil, err
		}
		val, err := value.UnmarshalValue(rd)
		if err != nil {
			return nil, err
		}
		vals = append(vals, core.MachineEmission{
			Value:    val,
			LogCount: new(big.Int).SetBytes(logCountData[:]),
		})
	}
	return vals, cursor.updateValues()
}

func (ac *ArbCore) GetLastMachine() (machine.Machine, error) {
	defer runtime.KeepAlive(ac)
	cMachine := C.arbCoreGetLastMachine(ac.c)
	if cMachine == nil {
		return nil, errors.Errorf("error getting last machine")
	}
	ret := &Machine{cMachine}

	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (ac *ArbCore) GetLastMachineTotalGas() (*big.Int, error) {
	defer runtime.KeepAlive(ac)
	result := C.arbCoreGetLastMachineTotalGas(ac.c)
	if result.found == 0 {
		return nil, errors.New("failed to get last machine total gas")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) UpdateCheckpointPruningGas(gas *big.Int) {
	defer runtime.KeepAlive(ac)
	gasData := math.U256Bytes(gas)
	C.arbCoreUpdateCheckpointPruningGas(ac.c, unsafeDataPointer(gasData))
}

func (ac *ArbCore) TakeMachine(executionCursor core.ExecutionCursor) (machine.Machine, error) {
	defer runtime.KeepAlive(ac)
	defer runtime.KeepAlive(executionCursor)
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
	defer runtime.KeepAlive(ac)
	cursorIndexData := math.U256Bytes(cursorIndex)
	result := C.arbCoreLogsCursorGetPosition(ac.c, unsafeDataPointer(cursorIndexData))
	if result.found == 0 {
		return nil, errors.New("failed to load logs cursor position")
	}

	return receiveBigInt(result.value), nil
}

func (ac *ArbCore) LogsCursorRequest(cursorIndex *big.Int, count *big.Int) error {
	defer runtime.KeepAlive(ac)
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

func (ac *ArbCore) LogsCursorGetLogs(cursorIndex *big.Int) (*big.Int, []core.ValueAndInbox, []core.ValueAndInbox, error) {
	defer runtime.KeepAlive(ac)
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
	logs := make([]core.ValueAndInbox, len(data))
	for i, slice := range data {
		var err error
		logs[i], err = unmarshalLog(slice[:])
		if err != nil {
			return nil, nil, nil, err
		}
	}

	deletedData := receiveByteSliceArray(result.second_array)
	deletedLogs := make([]core.ValueAndInbox, len(deletedData))
	for i, slice := range deletedData {
		var err error
		deletedLogs[i], err = unmarshalLog(slice[:])
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
	defer runtime.KeepAlive(ac)
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
	defer runtime.KeepAlive(ac)
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

func (ac *ArbCore) GetExecutionCursorAtEndOfBlock(blockNumber uint64, allowSlowLookup bool) (core.ExecutionCursor, error) {
	defer runtime.KeepAlive(ac)
	cExecutionCursorResult := C.arbCoreGetExecutionCursorAtEndOfBlock(ac.c, C.uint64_t(blockNumber), boolToCInt(allowSlowLookup))

	if cExecutionCursorResult.slow_error == 1 {
		return nil, errors.Errorf(slowLookupErrorString)
	}

	if cExecutionCursorResult.execution_cursor == nil {
		return nil, errors.Errorf("error creating execution cursor")
	}
	return NewExecutionCursor(cExecutionCursorResult.execution_cursor)
}
