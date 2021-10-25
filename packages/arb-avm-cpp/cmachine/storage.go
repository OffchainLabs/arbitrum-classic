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

func boolToCInt(b bool) C.int {
	x := 0
	if b {
		x = 1
	}
	return C.int(x)
}

func NewArbStorage(dbPath string, coreConfig *configuration.Core) (*ArbStorage, error) {
	cDbPath := C.CString(dbPath)
	defer C.free(unsafe.Pointer(cDbPath))

	cSaveRocksdbPath := C.CString(coreConfig.SaveRocksdbPath)
	defer C.free(unsafe.Pointer(cSaveRocksdbPath))

	cacheExpirationSeconds := int(coreConfig.Cache.TimedExpire.Seconds())
	saveRocksdbIntervalSeconds := int(coreConfig.SaveRocksdbInterval.Seconds())
	cConfig := C.CArbCoreConfig{
		message_process_count:         C.int(coreConfig.MessageProcessCount),
		checkpoint_load_gas_cost:      C.int(coreConfig.CheckpointLoadGasCost),
		min_gas_checkpoint_frequency:  C.int(coreConfig.GasCheckpointFrequency),
		cache_expiration_seconds:      C.int(cacheExpirationSeconds),
		lru_cache_size:                C.int(coreConfig.Cache.LRUSize),
		debug:                         boolToCInt(coreConfig.Debug),
		save_rocksdb_interval:         C.int(saveRocksdbIntervalSeconds),
		save_rocksdb_path:             cSaveRocksdbPath,
		lazy_load_core_machine:        boolToCInt(coreConfig.LazyLoadCoreMachine),
		lazy_load_archive_queries:     boolToCInt(coreConfig.LazyLoadArchiveQueries),
		profile_reorg_to:              C.int(coreConfig.Profile.ReorgTo),
		profile_run_until:             C.int(coreConfig.Profile.RunUntil),
		profile_load_count:            C.int(coreConfig.Profile.LoadCount),
		profile_reset_db_except_inbox: boolToCInt(coreConfig.Profile.ResetAllExceptInbox),
	}

	cArbStorage := C.createArbStorage(cDbPath, cConfig)

	if cArbStorage == nil {
		return nil, errors.Errorf("error creating ArbStorage %v", dbPath)
	}

	returnVal := &ArbStorage{cArbStorage}
	runtime.SetFinalizer(returnVal, cDestroyArbStorage)

	return returnVal, nil
}

func (s *ArbStorage) Initialize(contractPath string) error {
	defer runtime.KeepAlive(s)
	cContractPath := C.CString(contractPath)
	defer C.free(unsafe.Pointer(cContractPath))
	success := C.initializeArbStorage(s.c, cContractPath)

	if success == 0 {
		return errors.Errorf("failed to initialize storage with mexe '%v', possibly corrupt database or incorrect L1 node?", contractPath)
	}
	return nil
}

func (s *ArbStorage) Initialized() bool {
	defer runtime.KeepAlive(s)
	return C.arbStorageInitialized(s.c) == 1
}

func (s *ArbStorage) CloseArbStorage() bool {
	defer runtime.KeepAlive(s)
	return C.closeArbStorage(s.c) == 1
}

func cDestroyArbStorage(cArbStorage *ArbStorage) {
	C.destroyArbStorage(cArbStorage.c)
}

func (s *ArbStorage) GetArbCore() core.ArbCore {
	defer runtime.KeepAlive(s)
	ac := C.createArbCore(s.c)
	return NewArbCore(ac, s)
}

func (s *ArbStorage) GetNodeStore() machine.NodeStore {
	defer runtime.KeepAlive(s)
	as := C.createAggregatorStore(s.c)
	return NewNodeStore(as)
}
