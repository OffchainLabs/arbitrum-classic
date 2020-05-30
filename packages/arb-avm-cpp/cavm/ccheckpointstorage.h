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

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct ByteSliceStruct {
    void* data;
    int length;
};
typedef struct ByteSliceStruct ByteSlice;

typedef struct {
    ByteSlice slice;
    int found;
} ByteSliceResult;

typedef struct {
    void* data;
    int count;
} HashList;

typedef struct {
    ByteSlice* slices;
    int count;
} ByteSliceArray;

typedef void CMachine;
typedef void CCheckpointStorage;

CCheckpointStorage* createCheckpointStorage(const char* filename,
                                            const char* contract_path);
void destroyCheckpointStorage(CCheckpointStorage* storage);
CMachine* getInitialMachine(const CCheckpointStorage* storage_ptr);
CMachine* getMachine(const CCheckpointStorage* storage_ptr,
                     const void* machine_hash);
int closeCheckpointStorage(CCheckpointStorage* storage_ptr);
int deleteCheckpoint(CCheckpointStorage* storage_ptr, const void* machine_hash);
int saveValue(CCheckpointStorage* storage_ptr, const void* value_data);
ByteSlice getValue(const CCheckpointStorage* storage_ptr, const void* hash_key);
int deleteValue(CCheckpointStorage* storage_ptr, const void* hash_key);
int saveData(CCheckpointStorage* storage_ptr,
             const void* key,
             int key_length,
             const void* data,
             int data_length);
ByteSlice getData(CCheckpointStorage* storage_ptr,
                  const void* key,
                  int key_length);
int deleteData(CCheckpointStorage* storage_ptr,
               const void* key,
               int key_length);

int putBlock(CCheckpointStorage* storage_ptr,
             const void* height,
             const void* hash,
             const void* data,
             int data_length);
int deleteBlock(CCheckpointStorage* storage_ptr,
                const void* height,
                const void* hash);
ByteSliceResult getBlock(const CCheckpointStorage* storage_ptr,
                         const void* height,
                         const void* hash);
HashList blockHashesAtHeight(const CCheckpointStorage* storage_ptr,
                             const void* height);
int isBlockStoreEmpty(const CCheckpointStorage* storage_ptr);
void* maxBlockStoreHeight(const CCheckpointStorage* storage_ptr);
void* minBlockStoreHeight(const CCheckpointStorage* storage_ptr);

#ifdef __cplusplus
}
#endif

#endif /* ccheckpointstorage_hpp */
