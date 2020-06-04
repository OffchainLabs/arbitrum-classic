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
	"runtime"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
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

func (checkpoint *CheckpointStorage) CloseCheckpointStorage() bool {
	success := C.closeCheckpointStorage(checkpoint.c)

	return success == 1
}

func cDestroyCheckpointStorage(cCheckpointStorage *CheckpointStorage) {
	C.destroyCheckpointStorage(cCheckpointStorage.c)
}

func (checkpoint *CheckpointStorage) GetInitialMachine() (machine.Machine, error) {
	cMachine := C.getInitialMachine(checkpoint.c)

	if cMachine == nil {
		return nil, fmt.Errorf("error getting initial machine from checkpointstorage")
	}

	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (checkpoint *CheckpointStorage) GetMachine(machineHash common.Hash) (machine.Machine, error) {
	cMachine := C.getMachine(checkpoint.c, unsafe.Pointer(&machineHash[0]))

	if cMachine == nil {
		return nil, fmt.Errorf("error getting machine from checkpointstorage")
	}

	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (checkpoint *CheckpointStorage) DeleteCheckpoint(machineHash common.Hash) bool {
	success := C.deleteCheckpoint(checkpoint.c, unsafe.Pointer(&machineHash[0]))

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

func (checkpoint *CheckpointStorage) GetValue(hashValue common.Hash) value.Value {
	cData := C.getValue(checkpoint.c, unsafe.Pointer(&hashValue[0]))
	if cData.data == nil {
		return nil
	}

	dataBuff := toByteSlice(cData)

	val, err := value.UnmarshalValue(bytes.NewReader(dataBuff[:]))
	if err != nil {
		panic(err)
	}

	return val
}

func (checkpoint *CheckpointStorage) DeleteValue(hashValue common.Hash) bool {
	success := C.deleteValue(checkpoint.c, unsafe.Pointer(&hashValue[0]))

	return success == 1
}

func (checkpoint *CheckpointStorage) SaveData(key []byte, data []byte) bool {
	if len(key) == 0 {
		return false
	}

	if len(data) == 0 {
		success := C.saveData(checkpoint.c,
			unsafe.Pointer(&key[0]),
			C.int(len(key)),
			unsafe.Pointer(nil),
			C.int(0),
		)
		return success == 1
	}

	success := C.saveData(checkpoint.c,
		unsafe.Pointer(&key[0]),
		C.int(len(key)),
		unsafe.Pointer(&data[0]),
		C.int(len(data)))

	return success == 1
}

func (checkpoint *CheckpointStorage) GetData(key []byte) []byte {
	cData := C.getData(checkpoint.c, unsafe.Pointer(&key[0]), C.int(len(key)))

	if cData.found == 0 {
		return nil
	}

	return toByteSlice(cData.slice)
}

func (checkpoint *CheckpointStorage) DeleteData(key []byte) bool {
	success := C.deleteData(checkpoint.c, unsafe.Pointer(&key[0]), C.int(len(key)))

	return success == 1
}

func (checkpoint *CheckpointStorage) GetBlockStore() machine.BlockStore {
	bs := C.createBlockStore(checkpoint.c)

	return NewBlockStore(bs)
}

func (checkpoint *CheckpointStorage) GetNodeStore() machine.NodeStore {
	bs := C.createNodeStore(checkpoint.c)

	return NewNodeStore(bs)
}
