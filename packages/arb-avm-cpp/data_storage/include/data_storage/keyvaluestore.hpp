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

#ifndef keyvaluestore_hpp
#define keyvaluestore_hpp

#include <data_storage/storageresultfwd.hpp>

#include <memory>
#include <vector>

class DataStorage;

namespace rocksdb {
class Transaction;
class Status;
struct Slice;
class ColumnFamilyHandle;
}  // namespace rocksdb

class KeyValueStore {
   private:
    std::shared_ptr<DataStorage> data_storage;

   public:
    KeyValueStore(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(std::move(data_storage_)) {}
    rocksdb::Status saveData(const rocksdb::Slice& key,
                             const std::vector<unsigned char>& value);
    rocksdb::Status deleteData(const rocksdb::Slice& key);
    DataResults getData(const rocksdb::Slice& key) const;
};

#endif /* keyvaluestore_hpp */
