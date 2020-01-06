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
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -L../build/rocksdb -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb
#include "../cavm/ccheckpointstorage.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"bytes"
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"runtime"
	"unsafe"
)

type CheckpointStorage struct {
	c unsafe.Pointer
}

func NewCheckpoint(dbPath string, contractPath string) (*CheckpointStorage, error) {
	cDbPath := C.CString(dbPath)
	cContractPath := C.CString(contractPath)
	cCheckpointStorage := C.createCheckpointStorage(cDbPath, cContractPath)

	if cCheckpointStorage == nil {
		return nil, fmt.Errorf("error creating CheckpointStorage %v", dbPath)
	}

	returnVal := &CheckpointStorage{cCheckpointStorage}
	runtime.SetFinalizer(returnVal, cDestroyCheckpointStorage)

	C.free(unsafe.Pointer(cDbPath))
	C.free(unsafe.Pointer(cContractPath))

	return returnVal, nil
}

func cDestroyCheckpointStorage(cCheckpointStorage *CheckpointStorage) {
	C.destroyCheckpointStorage(cCheckpointStorage.c)
}

func (checkpoint *CheckpointStorage) GetInitialMachine() (machine.Machine, error) {
	cMachine := C.getInitialMachine(checkpoint.c)

	if cMachine == nil {
		return nil, fmt.Errorf("error getting machine from checkpointstorage")
	}

	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (checkpoint *CheckpointStorage) DeleteCheckpoint(checkpointName string) bool {
	cCheckpointName := C.CString(checkpointName)
	success := C.deleteCheckpoint(checkpoint.c, cCheckpointName)

	C.free(unsafe.Pointer(cCheckpointName))

	return success == 1
}

func (checkpoint *CheckpointStorage) SaveValue(val value.Value) bool {
	var buf bytes.Buffer

	err := value.MarshalValue(val, &buf)
	if err != nil {
		panic(err)
	}

	valData := buf.Bytes()
	success := C.saveValue(checkpoint.c, unsafe.Pointer(&valData[0]))

	return success == 1
}

func (checkpoint *CheckpointStorage) GetValue(hashValue [32]byte) value.Value {
	cData := C.getValue(checkpoint.c, unsafe.Pointer(&hashValue[0]))
	dataBuff := C.GoBytes(unsafe.Pointer(cData.data), cData.length)

	C.free(unsafe.Pointer(cData.data))

	val, err := value.UnmarshalValue(bytes.NewReader(dataBuff[:]))
	if err != nil {
		panic(err)
	}

	return val
}

func (checkpoint *CheckpointStorage) DeleteValue(hashValue [32]byte) bool {
	success := C.deleteValue(checkpoint.c, unsafe.Pointer(&hashValue[0]))

	return success == 1
}

func (checkpoint *CheckpointStorage) SaveData(key []byte, data []byte) bool {
	success := C.saveData(checkpoint.c,
		unsafe.Pointer(&key[0]),
		C.int(len(key)),
		unsafe.Pointer(&data[0]),
		C.int(len(data)))

	return success == 1
}

func (checkpoint *CheckpointStorage) GetData(key []byte) []byte {
	cData := C.getData(checkpoint.c, unsafe.Pointer(&key[0]), C.int(len(key)))
	dataBuff := C.GoBytes(unsafe.Pointer(cData.data), cData.length)

	C.free(unsafe.Pointer(cData.data))

	return dataBuff
}

func (checkpoint *CheckpointStorage) DeleteData(key []byte) bool {
	success := C.deleteData(checkpoint.c, unsafe.Pointer(&key[0]), C.int(len(key)))

	return success == 1
}
