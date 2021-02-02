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
#cgo LDFLAGS: -L. -L../build/rocksdb -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb -lkeccak -ldl
#include "../cavm/carbstorage.h"
#include "../cavm/cvaluecache.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/pkg/errors"
	"runtime"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type ArbStorage struct {
	c unsafe.Pointer
}

func NewCheckpoint(dbPath string) (*ArbStorage, error) {
	cDbPath := C.CString(dbPath)
	defer C.free(unsafe.Pointer(cDbPath))

	cArbStorage := C.createArbStorage(cDbPath)

	if cArbStorage == nil {
		return nil, errors.Errorf("error creating ArbStorage %v", dbPath)
	}

	returnVal := &ArbStorage{cArbStorage}
	runtime.SetFinalizer(returnVal, cDestroyArbStorage)

	return returnVal, nil
}

func (s *ArbStorage) Initialize(contractPath string) error {
	cContractPath := C.CString(contractPath)
	defer C.free(unsafe.Pointer(cContractPath))
	success := C.initializeArbStorage(s.c, cContractPath)

	if success == 0 {
		return errors.New("failed to initialize storage")
	}
	return nil
}

func (s *ArbStorage) Initialized() bool {
	return C.arbStorageInitialized(s.c) == 1
}

func (s *ArbStorage) CloseArbStorage() bool {
	return C.closeArbStorage(s.c) == 1
}

func cDestroyArbStorage(cArbStorage *ArbStorage) {
	C.destroyArbStorage(cArbStorage.c)
}

func (s *ArbStorage) GetInitialMachine() (machine.Machine, error) {
	cMachine := C.getInitialMachine(s.c)
	if cMachine == nil {
		return nil, errors.Errorf("error getting initial machine from arbstorage")
	}

	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (s *ArbStorage) GetMachine(machineHash common.Hash) (machine.Machine, error) {
	cMachine := C.getMachine(s.c, unsafe.Pointer(&machineHash[0]))

	if cMachine == nil {
		return nil, &machine.MachineNotFoundError{HashValue: machineHash}
	}

	ret := &Machine{cMachine}
	runtime.SetFinalizer(ret, cdestroyVM)
	return ret, nil
}

func (s *ArbStorage) DeleteCheckpoint(machineHash common.Hash) bool {
	success := C.deleteCheckpoint(s.c, unsafe.Pointer(&machineHash[0]))

	return success == 1
}

func (s *ArbStorage) SaveValue(val value.Value) bool {
	var buf bytes.Buffer

	err := value.MarshalValue(val, &buf)
	if err != nil {
		panic(err)
	}

	valData := buf.Bytes()
	success := C.saveValue(s.c, unsafe.Pointer(&valData[0]))

	return success == 1
}

func (s *ArbStorage) GetValue(hashValue common.Hash) (value.Value, error) {
	cData := C.getValue(s.c, unsafe.Pointer(&hashValue[0]))
	if cData.data == nil {
		return nil, &machine.ValueNotFoundError{HashValue: hashValue}
	}

	dataBuff := receiveByteSlice(cData)

	val, err := value.UnmarshalValue(bytes.NewReader(dataBuff[:]))
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (s *ArbStorage) DeleteValue(hashValue common.Hash) bool {
	success := C.deleteValue(s.c, unsafe.Pointer(&hashValue[0]))

	return success == 1
}

func (s *ArbStorage) SaveData(key []byte, data []byte) bool {
	if len(key) == 0 {
		return false
	}

	if len(data) == 0 {
		success := C.saveData(s.c,
			unsafe.Pointer(&key[0]),
			C.int(len(key)),
			unsafe.Pointer(nil),
			C.int(0),
		)
		return success == 1
	}

	success := C.saveData(s.c,
		unsafe.Pointer(&key[0]),
		C.int(len(key)),
		unsafe.Pointer(&data[0]),
		C.int(len(data)))

	return success == 1
}

func (s *ArbStorage) GetData(key []byte) ([]byte, error) {
	cData := C.getData(s.c, unsafe.Pointer(&key[0]), C.int(len(key)))

	if cData.found == 0 {
		return nil, &machine.DataNotFoundError{Key: key}
	}

	return receiveByteSlice(cData.slice), nil
}

func (s *ArbStorage) DeleteData(key []byte) bool {
	success := C.deleteData(s.c, unsafe.Pointer(&key[0]), C.int(len(key)))

	return success == 1
}

func (s *ArbStorage) GetArbCore() core.ArbCoreLookup {
	ac := C.createArbCore(s.c)
	return NewArbCore(ac)
}

func (s *ArbStorage) GetAggregatorStore() *AggregatorStore {
	as := C.createAggregatorStore(s.c)
	return NewAggregatorStore(as)
}
