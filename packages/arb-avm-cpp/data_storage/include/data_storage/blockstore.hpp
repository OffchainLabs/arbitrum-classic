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

#ifndef blockstore_hpp
#define blockstore_hpp

#include <avm_values/bigint.hpp>
#include <data_storage/storageresultfwd.hpp>

#include <memory>
#include <vector>

class DataStorage;

namespace rocksdb {
class TransactionDB;
class Status;
struct Slice;
class ColumnFamilyHandle;
}  // namespace rocksdb

class BlockStore {
   private:
    std::shared_ptr<DataStorage> data_storage;

   public:
    BlockStore() = default;
    BlockStore(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)) {}
    rocksdb::Status putBlock(const uint256_t& height,
                             const uint256_t& hash,
                             const std::vector<char>& value);
    rocksdb::Status deleteBlock(const uint256_t& height, const uint256_t& hash);
    DataResults getBlock(const uint256_t& height, const uint256_t& hash) const;

    std::vector<uint256_t> blockHashesAtHeight(const uint256_t& height) const;
    bool isEmpty() const;
    uint256_t maxHeight() const;
    uint256_t minHeight() const;
};

#endif /* blockstore_hpp */
