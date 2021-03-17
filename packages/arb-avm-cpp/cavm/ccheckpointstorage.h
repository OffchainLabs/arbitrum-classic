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
#ifndef ccheckpointstorage_h
#define ccheckpointstorage_h

#include "ctypes.h"

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

CCheckpointStorage* createCheckpointStorage(const char* filename);
int initializeCheckpointStorage(CCheckpointStorage* storage_ptr,
                                const char* executable_path);
int checkpointStorageInitialized(CCheckpointStorage* storage_ptr);
void destroyCheckpointStorage(CCheckpointStorage* storage);
CMachine* getInitialMachine(const CCheckpointStorage* storage_ptr,
                            CValueCache* value_cache_ptr);
CMachine* getMachine(const CCheckpointStorage* storage_ptr,
                     const void* machine_hash,
                     CValueCache* value_cache_ptr);
int flushCheckpointStorage(CCheckpointStorage* storage_ptr);
int closeCheckpointStorage(CCheckpointStorage* storage_ptr);
int deleteCheckpoint(CCheckpointStorage* storage_ptr, const void* machine_hash);
int saveValue(CCheckpointStorage* storage_ptr, const void* value_data);
ByteSlice getValue(const CCheckpointStorage* storage_ptr,
                   const void* hash_key,
                   CValueCache* value_cache_ptr);
int deleteValue(CCheckpointStorage* storage_ptr, const void* hash_key);
int saveData(CCheckpointStorage* storage_ptr,
             const void* key,
             int key_length,
             const void* data,
             int data_length);
ByteSliceResult getData(CCheckpointStorage* storage_ptr,
                        const void* key,
                        int key_length);
int deleteData(CCheckpointStorage* storage_ptr,
               const void* key,
               int key_length);

CBlockStore* createBlockStore(CCheckpointStorage* storage_ptr);
CAggregatorStore* createAggregatorStore(CCheckpointStorage* storage_ptr);

#ifdef __cplusplus
}
#endif

#endif /* ccheckpointstorage_hpp */
