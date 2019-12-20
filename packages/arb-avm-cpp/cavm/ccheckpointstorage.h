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

typedef void CCheckpointStorage;

CCheckpointStorage* createCheckpointStorage(const char* filename);
void destroyCheckpointStorage(CCheckpointStorage* storage);
int deleteCheckpoint(CCheckpointStorage* storage_ptr,
                     const char* checkpoint_name);
int saveValue(CCheckpointStorage* storage_ptr, void* value_data);
char* getValue(CCheckpointStorage* storage_ptr, const char* key);
int deleteValue(CCheckpointStorage* storage_ptr, const char* key);
int saveData(CCheckpointStorage* storage_ptr,
             const char* key,
             const char* value);
char* getData(CCheckpointStorage* storage_ptr, const char* key);
int deleteData(CCheckpointStorage* storage_ptr, const char* key);

#ifdef __cplusplus
}
#endif

#endif /* ccheckpointstorage_hpp */
