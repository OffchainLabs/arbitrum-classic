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
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"unsafe"
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

func toByteSliceView(data []byte) C.ByteSlice {
	return C.struct_ByteSliceStruct{data: unsafeDataPointer(data), length: C.int(len(data))}
}

func toByteSliceArrayView(slices []C.ByteSlice) C.ByteSliceArray {
	return C.struct_ByteSliceArrayStruct{slices: unsafe.Pointer(&slices[0]), count: C.int(len(slices))}
}

func encodeByteSliceList(goSlices [][]byte) []C.ByteSlice {
	byteSlices := make([]C.ByteSlice, 0, len(goSlices))
	for _, data := range goSlices {
		byteSlices = append(byteSlices, toByteSliceView(data))
	}
	return byteSlices
}

func encodeInboxMessages(inboxMessages []inbox.InboxMessage) [][]byte {
	data := make([][]byte, 0, len(inboxMessages))
	for _, msg := range inboxMessages {
		data = append(data, msg.ToBytes())
	}
	return data
}
