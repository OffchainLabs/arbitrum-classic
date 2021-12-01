/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
#include "../cavm/ctypes.h"
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"math/big"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
)

func unsafeDataPointer(data []byte) unsafe.Pointer {
	return unsafe.Pointer(&data[0])
}

func receiveBigInt(ptr unsafe.Pointer) *big.Int {
	data := receive32Bytes(ptr)
	return new(big.Int).SetBytes(data[:])
}

func receive32Bytes(ptr unsafe.Pointer) common.Hash {
	defer C.free(ptr)
	dataBuff := C.GoBytes(ptr, 32)
	rd := bytes.NewBuffer(dataBuff)
	var data common.Hash
	// Potential error can be ignored, bytes.Buffer is safe
	_, _ = rd.Read(data[:])
	return data
}

func receiveByteSlice(slice C.ByteSlice) []byte {
	defer C.free(unsafe.Pointer(slice.data))
	return C.GoBytes(unsafe.Pointer(slice.data), slice.length)
}

func receiveByteSliceArray(sliceArray C.ByteSliceArray) [][]byte {
	defer C.free(unsafe.Pointer(sliceArray.slices))
	dataSlices := (*[1 << 30]C.struct_ByteSliceStruct)(unsafe.Pointer(sliceArray.slices))[:sliceArray.count:sliceArray.count]
	slices := make([][]byte, sliceArray.count)
	for i := range dataSlices {
		slices[i] = receiveByteSlice(dataSlices[i])
	}
	return slices
}

// Note: the slices field of the returned struct needs manually freed by C.free
func bytesArrayToByteSliceArray(bytes [][]byte) C.struct_ByteSliceArrayStruct {
	byteSlices := encodeByteSliceList(bytes)
	sliceArrayData := C.malloc(C.size_t(C.sizeof_struct_ByteSliceStruct * len(byteSlices)))
	sliceArray := (*[1 << 30]C.struct_ByteSliceStruct)(sliceArrayData)[:len(byteSlices):len(byteSlices)]
	for i, data := range byteSlices {
		sliceArray[i] = data
	}
	return C.struct_ByteSliceArrayStruct{slices: sliceArrayData, count: C.int(len(byteSlices))}
}

func freeByteSliceArray(sliceArray C.struct_ByteSliceArrayStruct) {
	dataSlices := (*[1 << 30]C.struct_ByteSliceStruct)(unsafe.Pointer(sliceArray.slices))[:sliceArray.count:sliceArray.count]
	for _, slice := range dataSlices {
		C.free(slice.data)
	}
	C.free(sliceArray.slices)
}

func toByteSliceView(data []byte) C.ByteSlice {
	return C.struct_ByteSliceStruct{data: C.CBytes(data), length: C.int(len(data))}
}

func toByteSliceArrayView(slices []C.ByteSlice) C.ByteSliceArray {
	if len(slices) == 0 {
		return C.struct_ByteSliceArrayStruct{slices: nil, count: 0}
	}
	return C.struct_ByteSliceArrayStruct{slices: unsafe.Pointer(&slices[0]), count: C.int(len(slices))}
}

func encodeByteSliceList(goSlices [][]byte) []C.ByteSlice {
	byteSlices := make([]C.ByteSlice, 0, len(goSlices))
	for _, data := range goSlices {
		byteSlices = append(byteSlices, toByteSliceView(data))
	}
	return byteSlices
}

func encodeMachineInboxMessages(inboxMessages []inbox.InboxMessage) [][]byte {
	data := make([][]byte, 0, len(inboxMessages))
	for _, msg := range inboxMessages {
		machineMsg := inbox.MachineMessage{
			Accumulator: common.Hash{},
			Message:     msg,
		}
		data = append(data, machineMsg.ToBytes())
	}
	return data
}

func encodeInboxMessages(inboxMessages []inbox.InboxMessage) [][]byte {
	data := make([][]byte, 0, len(inboxMessages))
	for _, msg := range inboxMessages {
		data = append(data, msg.ToBytes())
	}
	return data
}

func encodeSequencerBatchItems(seqBatchItems []inbox.SequencerBatchItem) [][]byte {
	data := make([][]byte, 0, len(seqBatchItems))
	for _, msg := range seqBatchItems {
		data = append(data, msg.ToBytesWithSeqNum())
	}
	return data
}

func encodeDelayedMessages(delayedMessages []inbox.DelayedMessage) [][]byte {
	data := make([][]byte, 0, len(delayedMessages))
	for _, msg := range delayedMessages {
		data = append(data, msg.ToBytesWithSeqNum())
	}
	return data
}
