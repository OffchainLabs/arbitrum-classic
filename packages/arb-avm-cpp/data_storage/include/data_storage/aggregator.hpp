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

#include <optional>
#include "datastorage.hpp"
#include "storageresult.hpp"

class DataStorage;

class AggregatorStore {
    std::shared_ptr<DataStorage> data_storage;

   public:
    explicit AggregatorStore(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)) {}

    std::pair<uint64_t, std::vector<char>> latestBlock() const;
    void saveBlock(uint64_t height, const std::vector<char>& data);
    std::vector<char> getBlock(uint64_t height) const;

    std::optional<uint64_t> getPossibleRequestInfo(
        const uint256_t& request_id) const;
    void saveRequest(const uint256_t& request_id, uint64_t log_index);
    std::optional<uint64_t> getPossibleBlock(const uint256_t& block_hash) const;
    void saveBlockHash(const uint256_t& block_hash, uint64_t block_height);

    void reorg(uint64_t block_height);
    ValueResult<uint256_t> logsProcessedCount() const;
    rocksdb::Status updateLogsProcessedCount(uint256_t& count);
};

#endif /* aggregator_hpp */
