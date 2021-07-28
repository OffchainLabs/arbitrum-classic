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
#include "../cavm/carbstorage.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
)

type ArbStorage struct {
	c unsafe.Pointer
}

func NewArbStorage(dbPath string, coreConfig *configuration.Core) (*ArbStorage, error) {
	cDbPath := C.CString(dbPath)
	defer C.free(unsafe.Pointer(cDbPath))

	debugInt := 0
	if coreConfig.Debug {
		debugInt = 1
	}

	cacheExpirationSeconds := int(coreConfig.Cache.TimedExpire.Seconds())
	cArbStorage := C.createArbStorage(
		cDbPath,
		C.int(coreConfig.MessageProcessCount),
		C.int(coreConfig.CheckpointLoadGasCost),
		C.int(coreConfig.GasCheckpointFrequency),
		C.int(cacheExpirationSeconds),
		C.int(coreConfig.Cache.LRUSize),
		C.int(debugInt),
	)

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
		return errors.Errorf("failed to initialize storage with mexe '%v', possibly incorrect L1 node?", contractPath)
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

func (s *ArbStorage) GetArbCore() core.ArbCore {
	ac := C.createArbCore(s.c)
	return NewArbCore(ac, s)
}

func (s *ArbStorage) GetNodeStore() machine.NodeStore {
	as := C.createAggregatorStore(s.c)
	return NewNodeStore(as)
}
