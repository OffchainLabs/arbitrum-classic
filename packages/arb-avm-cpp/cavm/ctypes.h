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

struct Uint64ResultStruct {
    uint64_t value;
    int found;
};

typedef struct Uint64ResultStruct Uint64Result;

struct HashResultStruct {
    void* value;
    int found;
};

typedef struct {
    uint64_t inbox_messages_consumed;
    ByteSlice outMessages;
    int outMessageCount;
    ByteSlice logs;
    int logCount;
    ByteSlice debugPrints;
    int debugPrintCount;
    uint64_t numSteps;
    uint64_t numGas;
} RawAssertion;

typedef struct HashResultStruct HashResult;

typedef void CMachine;
typedef void CCheckpointedMachine;
typedef void CCheckpointStorage;
typedef void CBlockStore;
typedef void CAggregatorStore;
typedef void CValueCache;

#ifdef __cplusplus
}
#endif

#endif /* ctypes_h */
