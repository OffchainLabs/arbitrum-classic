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

#ifndef caggregator_h
#define caggregator_h

#include "ctypes.h"

#include <cstdio>

#ifdef __cplusplus
extern "C" {
#endif

struct CBlockDataStruct {
    int found;
    uint64_t height;
    ByteSlice data;
};

typedef struct CBlockDataStruct CBlockData;

struct CBlockIdStruct {
    int found;
    const void* hash;
    uint64_t height;
};

typedef struct CBlockIdStruct CBlockId;

struct CRequestInfoStruct {
    int found;
    uint64_t log_index;
    uint64_t evm_start_log_index;
};

typedef struct CRequestInfoStruct CRequestInfo;

void deleteAggregatorStore(CAggregatorStore* m);

Uint64Result aggregatorBlockCount(const CAggregatorStore* agg);
int aggregatorSaveMessageBatch(CAggregatorStore* agg_ptr,
                               const void* batch_num_ptr,
                               uint64_t log_index);
Uint64Result aggregatorGetMessageBatch(CAggregatorStore* agg_ptr,
                                       const void* batch_num_ptr);
int aggregatorSaveBlock(CAggregatorStore* agg_ptr,
                        uint64_t height,
                        const void* block_hash,
                        ByteSliceArray requests_data,
                        const uint64_t* log_indexes,
                        const void* block_data,
                        int block_data_length);
CBlockData aggregatorGetBlock(const CAggregatorStore* agg, uint64_t height);
int aggregatorReorg(CAggregatorStore* agg, uint64_t block_height);

// request_id is 32 bytes long
Uint64Result aggregatorGetPossibleRequestInfo(const CAggregatorStore* agg,
                                              const void* request_id);
// block_hash is 32 bytes long
Uint64Result aggregatorGetPossibleBlock(const CAggregatorStore* agg,
                                        const void* block_hash);

Uint256Result aggregatorLogsProcessedCount(CAggregatorStore* agg);
int aggregatorUpdateLogsProcessedCount(CAggregatorStore* agg, void* count_ptr);

#ifdef __cplusplus
}
#endif

#endif /* caggregator_h */
