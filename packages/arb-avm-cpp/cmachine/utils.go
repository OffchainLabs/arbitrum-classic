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
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"math/big"
	"unsafe"
)

func intToData(val *big.Int) unsafe.Pointer {
	var lowerBoundBlockBuf bytes.Buffer

	// Potential error can be ignored, bytes.Buffer is safe
	_ = value.NewIntValue(val).Marshal(&lowerBoundBlockBuf)
	return C.CBytes(lowerBoundBlockBuf.Bytes())
}

func dataToInt(ptr unsafe.Pointer) *big.Int {
	dataBuff := C.GoBytes(ptr, 32)
	buf := bytes.NewBuffer(dataBuff)

	// Potential error can be ignored, bytes.Buffer is safe
	intVal, _ := value.NewIntValueFromReader(buf)
	return intVal.BigInt()
}

func hashToData(val common.Hash) unsafe.Pointer {
	var lowerBoundBlockBuf bytes.Buffer

	// Potential error can be ignored, bytes.Buffer is safe
	_ = value.NewIntValue(new(big.Int).SetBytes(val[:])).Marshal(&lowerBoundBlockBuf)
	return C.CBytes(lowerBoundBlockBuf.Bytes())
}

func toByteSlice(slice C.ByteSlice) []byte {
	defer C.free(unsafe.Pointer(slice.data))
	return C.GoBytes(unsafe.Pointer(slice.data), slice.length)
}

func toByteSliceArray(sliceArray C.ByteSliceArray) [][]byte {
	defer C.free(unsafe.Pointer(sliceArray.slices))
	dataSlices := (*[1 << 30]C.struct_ByteSliceStruct)(unsafe.Pointer(sliceArray.slices))[:sliceArray.count:sliceArray.count]
	slices := make([][]byte, sliceArray.count)
	for i := range dataSlices {
		slices[i] = toByteSlice(dataSlices[i])
	}
	return slices
}

func encodeByteSliceArray(goSlices [][]byte) C.ByteSliceArray {
	sliceArrayData := C.malloc(C.size_t(C.sizeof_struct_ByteSliceStruct * len(goSlices)))
	sliceArray := (*[1 << 30]C.struct_ByteSliceStruct)(sliceArrayData)[:len(goSlices):len(goSlices)]

	for i := range goSlices {
		sliceData := C.malloc(C.size_t(len(goSlices[i])))
		slice := (*[1 << 30]byte)(sliceData)[:len(goSlices[i]):len(goSlices[i])]
		copy(goSlices[i], slice)
		sliceArray[i].length = C.int(len(slice))
		sliceArray[i].data = C.CBytes(slice)
	}

	return C.ByteSliceArray{
		count:  C.int(len(goSlices)),
		slices: sliceArrayData,
	}
}

func freeEncodedByteSliceArray(byteSliceArray C.ByteSliceArray) {
	sliceArray := (*[1 << 30]C.struct_ByteSliceStruct)(byteSliceArray.slices)[:byteSliceArray.count:byteSliceArray.count]

	for _, slice := range sliceArray {
		C.free(slice.data)
	}

	C.free(byteSliceArray.slices)
}
