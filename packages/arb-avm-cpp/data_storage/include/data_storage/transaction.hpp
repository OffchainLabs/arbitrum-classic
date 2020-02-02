/*
 * Copyright 2019, Offchain Labs, Inc.
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

#ifndef transaction_hpp
#define transaction_hpp

#include <memory>
#include <vector>

struct DeleteResults;
struct GetResults;
struct SaveResults;

namespace rocksdb {
class Transaction;
class Status;
}  // namespace rocksdb

auto parseCountAndValue(const std::string& string_value)
    -> std::tuple<uint32_t, std::vector<unsigned char>>;

class Transaction {
   private:
    std::unique_ptr<rocksdb::Transaction> transaction;

    auto saveKeyValuePair(const std::vector<unsigned char>& key,
                          const std::vector<unsigned char>& value)
        -> rocksdb::Status;
    auto deleteKeyValuePair(const std::vector<unsigned char>& key)
        -> rocksdb::Status;
    auto saveValueWithRefCount(uint32_t updated_ref_count,
                               const std::vector<unsigned char>& hash_key,
                               const std::vector<unsigned char>& value)
        -> SaveResults;

   public:
    Transaction(rocksdb::Transaction* transaction_);
    Transaction(const Transaction&) = delete;
    Transaction(Transaction&&) = default;
    Transaction& operator=(const Transaction&) = delete;
    Transaction& operator=(Transaction&&) = default;
    ~Transaction();

    auto incrementReference(const std::vector<unsigned char>& hash_key)
        -> SaveResults;
    auto saveData(const std::vector<unsigned char>& hash_key,
                  const std::vector<unsigned char>& value) -> SaveResults;
    auto getData(const std::vector<unsigned char>& hash_key) const
        -> GetResults;
    auto deleteData(const std::vector<unsigned char>& hash_key)
        -> DeleteResults;
    auto commit() -> rocksdb::Status;
    auto rollBack() -> rocksdb::Status;
};

#endif /* transaction_hpp */
