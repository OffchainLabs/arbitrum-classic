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
#ifndef ccheckpointstorage_h
#define ccheckpointstorage_h

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
    void* data;
    int length;
} ByteSlice;

using CMachine = void;
using CCheckpointStorage = void;

auto createCheckpointStorage(const char* filename, const char* contract_path)
    -> CCheckpointStorage*;
void destroyCheckpointStorage(CCheckpointStorage* storage);
auto getInitialMachine(const CCheckpointStorage* storage_ptr) -> CMachine*;
auto closeCheckpointStorage(CCheckpointStorage* storage_ptr) -> int;
auto deleteCheckpoint(CCheckpointStorage* storage_ptr, const void* machine_hash)
    -> int;
auto saveValue(CCheckpointStorage* storage_ptr, const void* value_data) -> int;
auto getValue(const CCheckpointStorage* storage_ptr, const void* hash_key)
    -> ByteSlice;
auto deleteValue(CCheckpointStorage* storage_ptr, const void* hash_key) -> int;
auto saveData(CCheckpointStorage* storage_ptr,
              const void* key,
              int key_length,
              const void* data,
              int data_length) -> int;
auto getData(CCheckpointStorage* storage_ptr, const void* key, int key_length)
    -> ByteSlice;
auto deleteData(CCheckpointStorage* storage_ptr,
                const void* key,
                int key_length) -> int;

#ifdef __cplusplus
}
#endif

#endif /* ccheckpointstorage_hpp */
