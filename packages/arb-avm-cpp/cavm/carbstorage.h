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
#ifndef carbstorage_h
#define carbstorage_h

#include "ctypes.h"
#include "cmachine.h"

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    int32_t message_process_count;
    int32_t checkpoint_load_gas_cost;
    int32_t min_gas_checkpoint_frequency;
    int32_t cache_expiration_seconds;
    int32_t lru_cache_size;
    int32_t debug;
    int32_t save_rocksdb_interval;
    const char* save_rocksdb_path;
    int32_t lazy_load_core_machine;
    int32_t lazy_load_archive_queries;
    int32_t profile_reorg_to;
    int32_t profile_run_until;
    int32_t profile_load_count;
    int32_t profile_reset_db_except_inbox;
    int32_t profile_just_metadata;
} CArbCoreConfig;

CArbStorage* createArbStorage(const char* db_path,
                              CArbCoreConfig arb_core_config);
int initializeArbStorage(CArbStorage* storage_ptr, const char* executable_path);
int arbStorageInitialized(CArbStorage* storage_ptr);
void destroyArbStorage(CArbStorage* storage);
int closeArbStorage(CArbStorage* storage_ptr);

int initializeArbStorageForWasm(CArbStorage* storage_ptr, CMachine *machine);

CArbCore* createArbCore(CArbStorage* storage_ptr);
CAggregatorStore* createAggregatorStore(CArbStorage* storage_ptr);

#ifdef __cplusplus
}
#endif

#endif /* carbstorage_h */
