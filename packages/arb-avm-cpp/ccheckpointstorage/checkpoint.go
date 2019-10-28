/*
 * Copyright 2019, Offchain Labs, Inc.
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

package ccheckpointstorage

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lcavm -lavm -lstdc++ -lrocksdb
#include "../cavm/ccheckpointstorage.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type CheckpointStorage struct {
	c unsafe.Pointer
}

func NewCheckpoint(dbPath string) (*CheckpointStorage, error) {
	cDbPath := C.CString(dbPath)
	cCheckpointStorage := C.createCheckpointStorage(cDbPath)

	if cCheckpointStorage == nil {
		return nil, fmt.Errorf("error creating CheckpointStorage %v", dbPath)
	}

	returnVal := &CheckpointStorage{cCheckpointStorage}
	runtime.SetFinalizer(returnVal, cDestroyCheckpointStorage)
	C.free(unsafe.Pointer(cDbPath))

	return returnVal, nil
}

func cDestroyCheckpointStorage(cCheckpointStorage *CheckpointStorage) {
	C.destroyCheckpointStorage(cCheckpointStorage.c)
}

func (checkpoint *CheckpointStorage) DeleteCheckpoint(checkpointName string) bool {
	cCheckpointName := C.CString(checkpointName)
	success := C.deleteCheckpoint(checkpoint.c, cCheckpointName)

	return success == 1
}

func (checkpoint *CheckpointStorage) GetCStorage() unsafe.Pointer {
	return checkpoint.c
}
