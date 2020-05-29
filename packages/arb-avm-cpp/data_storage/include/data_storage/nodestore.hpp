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

#ifndef nodestore_hpp
#define nodestore_hpp

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

class NodeStore {
   private:
    std::shared_ptr<DataStorage> data_storage;

   public:
    NodeStore() = default;
    NodeStore(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)) {}
    rocksdb::Status putNode(uint64_t height,
                            const uint256_t& hash,
                            const std::vector<char>& data);
    DataResults getNode(uint64_t height, const uint256_t& hash) const;

    ValueResult<uint64_t> getHeight(const uint256_t& hash) const;
    ValueResult<uint256_t> getHash(uint64_t height) const;

    bool isEmpty() const;
    uint64_t maxNodeHeight() const;
};

#endif /* nodestore_hpp */
