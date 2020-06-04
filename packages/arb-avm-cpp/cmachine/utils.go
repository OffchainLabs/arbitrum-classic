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
	_ = value.NewIntValue(val).Marshal(&lowerBoundBlockBuf)
	return C.CBytes(lowerBoundBlockBuf.Bytes())
}

func dataToInt(ptr unsafe.Pointer) *big.Int {
	dataBuff := C.GoBytes(ptr, 32)
	buf := bytes.NewBuffer(dataBuff)
	intVal, _ := value.NewIntValueFromReader(buf)
	return intVal.BigInt()
}

func hashToData(val common.Hash) unsafe.Pointer {
	var lowerBoundBlockBuf bytes.Buffer
	_ = value.NewIntValue(new(big.Int).SetBytes(val[:])).Marshal(&lowerBoundBlockBuf)
	return C.CBytes(lowerBoundBlockBuf.Bytes())
}

func bytesArrayToVals(data []byte, valCount int) []value.Value {
	rd := bytes.NewReader(data)
	vals := make([]value.Value, 0, valCount)
	for i := 0; i < valCount; i++ {
		val, err := value.UnmarshalValue(rd)
		if err != nil {
			panic(err)
		}
		vals = append(vals, val)
	}
	return vals
}

func toByteSlice(slice C.ByteSlice) []byte {
	defer C.free(unsafe.Pointer(slice.data))
	return C.GoBytes(unsafe.Pointer(slice.data), slice.length)
}
