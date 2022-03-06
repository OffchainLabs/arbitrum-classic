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
#include <data_storage/datastorage.hpp>
#include <data_storage/readsnapshottransaction.hpp>
#include <data_storage/storageresult.hpp>

#include <rocksdb/utilities/transaction.h>

#include <optional>

class DataStorage;

class AggregatorStore {
    std::shared_ptr<DataStorage> data_storage;

   public:
    explicit AggregatorStore(std::shared_ptr<DataStorage> data_storage_);

    [[nodiscard]] uint64_t blockCount() const;
    void saveBlock(uint64_t height,
                   const uint256_t& block_hash,
                   const std::vector<uint256_t>& requests,
                   const uint64_t* log_indexes,
                   const std::vector<char>& data);
    [[nodiscard]] std::vector<char> getBlock(uint64_t height) const;
    [[nodiscard]] std::optional<uint64_t> getPossibleRequestInfo(
        const uint256_t& request_id) const;
    [[nodiscard]] std::optional<uint64_t> getPossibleBlock(
        const uint256_t& block_hash) const;

    void reorg(uint64_t block_height);
    void saveMessageBatch(const uint256_t& batchNum, const uint64_t& logIndex);
    std::optional<uint64_t> getMessageBatch(const uint256_t& batchNum);
};

#endif /* aggregator_hpp */
