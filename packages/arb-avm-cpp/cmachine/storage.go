/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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
#cgo LDFLAGS: -L. -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb -lkeccak -ldl
#include "../cavm/carbstorage.h"
#include "../cavm/cvaluecache.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type ArbStorage struct {
	c unsafe.Pointer
}

func NewArbStorage(dbPath string) (*ArbStorage, error) {
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

func (s *ArbStorage) GetArbCore() core.ArbCore {
	ac := C.createArbCore(s.c)
	return NewArbCore(ac, s)
}

func (s *ArbStorage) GetNodeStore() machine.NodeStore {
	as := C.createAggregatorStore(s.c)
	return NewNodeStore(as)
}
