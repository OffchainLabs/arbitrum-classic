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

#ifndef ctypes_h
#define ctypes_h

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct ByteSliceStruct {
    void* data;
    int length;
} ByteSlice;

typedef struct {
    ByteSlice slice;
    int found;
} ByteSliceResult;

typedef struct {
    void* data;
    int count;
} HashList;

typedef struct ByteSliceArrayStruct {
    void* slices;
    int count;
} ByteSliceArray;

typedef struct {
    ByteSliceArray array;
    int found;
} ByteSliceArrayResult;

typedef struct Uint64ResultStruct {
    uint64_t value;
    int found;
} Uint64Result;

typedef struct HashResultStruct {
    void* value;
    int found;
} Uint256Result;

typedef struct {
    uint64_t inbox_messages_consumed;
    ByteSlice sends;
    int sendCount;
    void* sendAcc;
    ByteSlice logs;
    int logCount;
    void* logAcc;
    ByteSlice debugPrints;
    int debugPrintCount;
    uint64_t numSteps;
    uint64_t numGas;
} RawAssertion;

typedef void CAggregatorStore;
typedef void CArbCore;
typedef void CArbStorage;
typedef void CBlockStore;
typedef void CCheckpointedMachine;
typedef void CExecutionCursor;
typedef void CMachine;
typedef void CValueCache;

#ifdef __cplusplus
}
#endif

#endif /* ctypes_h */
