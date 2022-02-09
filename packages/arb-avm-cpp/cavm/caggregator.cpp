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

#include "caggregator.h"

#include "utils.hpp"

#include <data_storage/aggregator.hpp>

#include <iostream>

void deleteAggregatorStore(CAggregatorStore* agg) {
    delete static_cast<AggregatorStore*>(agg);
}

Uint64Result aggregatorBlockCount(const CAggregatorStore* agg) {
    try {
        auto count = static_cast<const AggregatorStore*>(agg)->blockCount();
        return {count, true};
    } catch (const std::exception& e) {
        std::cerr << "Exception loading block count: " << e.what() << std::endl;
        return {0, false};
    }
}

int aggregatorSaveMessageBatch(CAggregatorStore* agg_ptr,
                               const void* batch_num_ptr,
                               const uint64_t log_index) {
    try {
        auto agg = static_cast<AggregatorStore*>(agg_ptr);
        auto batch_num = receiveUint256(batch_num_ptr);

        agg->saveMessageBatch(batch_num, log_index);

        return true;
    } catch (const std::exception& e) {
        std::cerr << "aggregatorSaveBlock error: " << e.what() << std::endl;
        return false;
    }
}

Uint64Result aggregatorGetMessageBatch(CAggregatorStore* agg_ptr,
                                       const void* batch_num_ptr) {
    try {
        auto agg = static_cast<AggregatorStore*>(agg_ptr);
        auto batch_num = receiveUint256(batch_num_ptr);
        auto index = agg->getMessageBatch(batch_num);
        if (index) {
            return {*index, true};
        } else {
            return {0, false};
        }
    } catch (const std::exception& e) {
        std::cerr << "aggregatorGetMessageBatch error: " << e.what()
                  << std::endl;
        return {0, false};
    }
}

int aggregatorSaveBlock(CAggregatorStore* agg_ptr,
                        const uint64_t height,
                        const void* block_hash_ptr,
                        const ByteSliceArray requests_data,
                        const uint64_t* log_indexes,
                        const void* block_data,
                        const int block_data_length) {
    try {
        auto agg = static_cast<AggregatorStore*>(agg_ptr);
        auto block_hash = receiveUint256(block_hash_ptr);
        auto request_ids = receiveUint256Array(requests_data);
        auto block_ptr = reinterpret_cast<const char*>(block_data);

        agg->saveBlock(height, block_hash, request_ids, log_indexes,
                       {block_ptr, block_ptr + block_data_length});

        return true;
    } catch (const std::exception& e) {
        std::cerr << "aggregatorSaveBlock error: " << e.what() << std::endl;
        return false;
    }
}

CBlockData aggregatorGetBlock(const CAggregatorStore* agg, uint64_t height) {
    try {
        auto block = static_cast<const AggregatorStore*>(agg)->getBlock(height);
        return {true, height, returnCharVector(block)};
    } catch (const std::exception&) {
        return {false, 0, ByteSlice{nullptr, 0}};
    }
}

int aggregatorReorg(CAggregatorStore* agg, uint64_t block_height) {
    try {
        static_cast<AggregatorStore*>(agg)->reorg(block_height);
        return true;
    } catch (const std::exception& e) {
        std::cerr << "aggregatorRestoreBlock error: " << e.what() << std::endl;
        return false;
    }
}

// request_id is 32 bytes long
Uint64Result aggregatorGetPossibleRequestInfo(const CAggregatorStore* agg,
                                              const void* request_id) {
    auto index =
        static_cast<const AggregatorStore*>(agg)->getPossibleRequestInfo(
            receiveUint256(request_id));
    if (index) {
        return {*index, true};
    } else {
        return {0, false};
    }
}

// block_hash is 32 bytes long
Uint64Result aggregatorGetPossibleBlock(const CAggregatorStore* agg,
                                        const void* block_hash) {
    auto index = static_cast<const AggregatorStore*>(agg)->getPossibleBlock(
        receiveUint256(block_hash));
    if (index) {
        return {*index, true};
    } else {
        return {0, false};
    }
}
