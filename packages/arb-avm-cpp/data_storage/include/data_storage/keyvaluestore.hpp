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

#include <memory>
#include <vector>

struct DataResults;

namespace rocksdb {
class Transaction;
class Status;
struct Slice;
class ColumnFamilyHandle;
}  // namespace rocksdb

class KeyValueStore {
   private:
    std::unique_ptr<rocksdb::Transaction> transaction;
    std::shared_ptr<rocksdb::ColumnFamilyHandle> column;

   public:
    KeyValueStore(std::unique_ptr<rocksdb::Transaction> transaction_,
                  std::shared_ptr<rocksdb::ColumnFamilyHandle> column_)
        : transaction(std::move(transaction_)), column(std::move(column_)) {}
    rocksdb::Status saveData(const rocksdb::Slice& key,
                             const std::vector<unsigned char>& value);
    rocksdb::Status deleteData(const rocksdb::Slice& key);
    DataResults getData(const rocksdb::Slice& key) const;
};

#endif /* keyvaluestore_hpp */
