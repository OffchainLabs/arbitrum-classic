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

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

CArbStorage* createArbStorage(const char* filename);
int initializeArbStorage(CArbStorage* storage_ptr, const char* executable_path);
int arbStorageInitialized(CArbStorage* storage_ptr);
void destroyArbStorage(CArbStorage* storage);
CMachine* getInitialMachine(const CArbStorage* storage_ptr,
                            CValueCache* value_cache_ptr);
CMachine* getMachine(const CArbStorage* storage_ptr,
                     const void* machine_hash,
                     CValueCache* value_cache_ptr);
int closeArbStorage(CArbStorage* storage_ptr);
int deleteCheckpoint(CArbStorage* storage_ptr, const void* machine_hash);
int saveValue(CArbStorage* storage_ptr, const void* value_data);
ByteSlice getValue(const CArbStorage* storage_ptr,
                   const void* hash_key,
                   CValueCache* value_cache_ptr);
int deleteValue(CArbStorage* storage_ptr, const void* hash_key);
int saveData(CArbStorage* storage_ptr,
             const void* key,
             int key_length,
             const void* data,
             int data_length);
ByteSliceResult getData(CArbStorage* storage_ptr,
                        const void* key,
                        int key_length);
int deleteData(CArbStorage* storage_ptr, const void* key, int key_length);

CBlockStore* createBlockStore(CArbStorage* storage_ptr);
CAggregatorStore* createAggregatorStore(CArbStorage* storage_ptr);

RawAssertion ArbExecuteAssertion(CArbStorage* storage_ptr,
                                 uint64_t maxSteps,
                                 void* inbox_messages,
                                 uint64_t message_count,
                                 uint64_t wallLimit);

#ifdef __cplusplus
}
#endif

#endif /* carbstorage_h */
