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

#include <nonstd/optional.hpp>

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
    void saveLog(rocksdb::Transaction& tx,
                 const std::vector<unsigned char>& log);
    std::vector<char> getLog(uint64_t index) const;

    uint64_t messageCount() const;
    void saveMessage(rocksdb::Transaction& tx,
                     const std::vector<unsigned char>& output);
    std::vector<char> getMessage(uint64_t index) const;

    std::pair<uint64_t, std::vector<char>> latestBlock() const;
    void saveBlock(uint64_t height, const std::vector<char>& data);
    std::vector<char> getBlock(uint64_t height) const;

    nonstd::optional<uint64_t> getPossibleRequestInfo(
        const uint256_t& request_id) const;
    void saveRequest(const uint256_t& request_id, uint64_t log_index);

    nonstd::optional<uint64_t> getPossibleBlock(
        const uint256_t& block_hash) const;
    void saveBlockHash(const uint256_t& block_hash, uint64_t block_height);

    void reorg(uint64_t block_height,
               uint64_t message_count,
               uint64_t log_count);
};

#endif /* aggregator_hpp */
