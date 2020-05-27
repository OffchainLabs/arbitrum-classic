/*
 * Copyright 2020, Offchain Labs, Inc.
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

#ifndef cblockstore_h
#define cblockstore_h

#include "ctypes.h"

#include <stdio.h>

#ifdef __cplusplus
extern "C" {
#endif

void deleteBlockStore(CBlockStore* m);

int putBlock(CBlockStore* storage_ptr,
             const void* height,
             const void* hash,
             const void* data,
             int data_length);
int deleteBlock(CBlockStore* storage_ptr, const void* height, const void* hash);
ByteSliceResult getBlock(const CBlockStore* storage_ptr,
                         const void* height,
                         const void* hash);
HashList blockHashesAtHeight(const CBlockStore* storage_ptr,
                             const void* height);
int isBlockStoreEmpty(const CBlockStore* storage_ptr);
void* maxBlockStoreHeight(const CBlockStore* storage_ptr);
void* minBlockStoreHeight(const CBlockStore* storage_ptr);

#ifdef __cplusplus
}
#endif

#endif /* cblockstore_h */
