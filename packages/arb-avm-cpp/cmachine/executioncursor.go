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
#include "../cavm/cexecutioncursor.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/pkg/errors"
)

type ExecutionCursor struct {
	c unsafe.Pointer
}

func deleteExecutionCursor(ac *ExecutionCursor) {
	C.deleteExecutionCursor(ac.c)
}

func NewExecutionCursor(c unsafe.Pointer) *ExecutionCursor {
	ec := &ExecutionCursor{c: c}
	runtime.SetFinalizer(ec, deleteExecutionCursor)
	return ec
}

func (ec *ExecutionCursor) MessageCount() (uint64, error) {
	result := C.executionCursorMachineHash(ec.c)
	if result.found == 0 {
		return 0, errors.New("failed to load l2message count")
	}
	return uint64(result.value), nil
}
