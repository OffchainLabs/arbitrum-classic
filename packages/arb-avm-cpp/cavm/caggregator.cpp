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

Uint64Result aggregatorLogCount(const CAggregatorStore* agg) {
    try {
        return {static_cast<const AggregatorStore*>(agg)->logCount(), true};
    } catch (const std::exception&) {
        return {0, false};
    }
}

int aggregatorSaveLog(CAggregatorStore* agg,
                      const void* data,
                      uint64_t length) {
    try {
        auto ptr = reinterpret_cast<const char*>(data);
        static_cast<AggregatorStore*>(agg)->saveLog({ptr, ptr + length});
        return 1;
    } catch (const std::exception&) {
        return 0;
    }
}

ByteSliceResult aggregatorGetLog(const CAggregatorStore* agg, uint64_t index) {
    try {
        auto data = returnCharVector(
            static_cast<const AggregatorStore*>(agg)->getLog(index));
        return {data, true};
    } catch (const std::exception&) {
        return {{nullptr, 0}, false};
    }
}

Uint64Result aggregatorMessageCount(const CAggregatorStore* agg) {
    try {
        return {static_cast<const AggregatorStore*>(agg)->messageCount(), true};
    } catch (const std::exception&) {
        return {0, false};
    }
}

int aggregatorSaveMessage(CAggregatorStore* agg,
                          const void* data,
                          uint64_t length) {
    try {
        auto ptr = reinterpret_cast<const char*>(data);
        static_cast<AggregatorStore*>(agg)->saveMessage({ptr, ptr + length});
        return 1;
    } catch (const std::exception&) {
        return 0;
    }
}

ByteSliceResult aggregatorGetMessage(const CAggregatorStore* agg,
                                     uint64_t index) {
    try {
        auto data = returnCharVector(
            static_cast<const AggregatorStore*>(agg)->getMessage(index));
        return {data, true};
    } catch (const std::exception&) {
        return {{nullptr, 0}, false};
    }
}

CBlockData aggregatorLatestBlock(const CAggregatorStore* agg) {
    try {
        auto latest = static_cast<const AggregatorStore*>(agg)->latestBlock();
        return {true, latest.first, returnCharVector(latest.second)};
    } catch (const std::exception& e) {
        std::cerr << "aggregatorLatestBlock error: " << e.what() << std::endl;
        return {false, 0, ByteSlice{nullptr, 0}};
    }
}

int aggregatorSaveBlock(CAggregatorStore* agg,
                        uint64_t height,
                        const void* data,
                        int data_length) {
    try {
        auto ptr = reinterpret_cast<const char*>(data);
        static_cast<AggregatorStore*>(agg)->saveBlock(height,
                                                      {ptr, ptr + data_length});
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

int aggregatorReorg(CAggregatorStore* agg,
                    uint64_t block_height,
                    uint64_t message_count,
                    uint64_t log_count) {
    try {
        static_cast<AggregatorStore*>(agg)->reorg(block_height, message_count,
                                                  log_count);
        return true;
    } catch (const std::exception& e) {
        std::cerr << "aggregatorRestoreBlock error: " << e.what() << std::endl;
        return false;
    }
}

// request_id is 32 bytes long
Uint64Result aggregatorGetPossibleRequestInfo(const CAggregatorStore* agg,
                                              const void* request_id) {
    try {
        auto index =
            static_cast<const AggregatorStore*>(agg)->getPossibleRequestInfo(
                receiveUint256(request_id));
        return {index, true};
    } catch (const std::exception&) {
        return {0, false};
    }
}

int aggregatorSaveRequest(CAggregatorStore* agg,
                          const void* request_id,
                          uint64_t log_index) {
    try {
        static_cast<AggregatorStore*>(agg)->saveRequest(
            receiveUint256(request_id), log_index);
        return 1;
    } catch (const std::exception&) {
        return 0;
    }
}
