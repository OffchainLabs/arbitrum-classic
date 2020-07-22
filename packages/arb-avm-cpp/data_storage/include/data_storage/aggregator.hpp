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

#ifndef aggregator_hpp
#define aggregator_hpp

#include <avm_values/bigint.hpp>

#include <rocksdb/utilities/transaction.h>

class DataStorage;

struct BlockData {
    uint64_t start_log;
    uint64_t log_count;

    uint64_t start_message;
    uint64_t message_count;

    std::vector<char> data;
};

class AggregatorStore {
    std::shared_ptr<DataStorage> data_storage;

   public:
    AggregatorStore(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)) {}

    uint64_t logCount() const;
    void saveLog(const std::vector<char>& log);
    std::vector<char> getLog(uint64_t index) const;

    uint64_t messageCount() const;
    void saveMessage(const std::vector<char>& log);
    std::vector<char> getMessage(uint64_t index) const;

    std::pair<uint64_t, BlockData> latestBlock() const;
    uint64_t getInitialBlock() const;
    void saveBlock(uint64_t height, const std::vector<char>& data);
    BlockData getBlock(uint64_t height) const;
    void restoreBlock(uint64_t height);

    uint64_t getPossibleRequestInfo(const uint256_t& request_id) const;
    void saveRequest(const uint256_t& request_id, uint64_t log_index);
};

#endif /* aggregator_hpp */
