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

#include <data_storage/storageresultfwd.hpp>

#include <memory>
#include <vector>

namespace rocksdb {
class Transaction;
class Status;
}  // namespace rocksdb

std::tuple<uint32_t, std::vector<unsigned char>> parseCountAndValue(
    const std::string& string_value);

class Transaction {
   private:
    std::unique_ptr<rocksdb::Transaction> transaction;

    rocksdb::Status saveKeyValuePair(const std::vector<unsigned char>& key,
                                     const std::vector<unsigned char>& value);
    rocksdb::Status deleteKeyValuePair(const std::vector<unsigned char>& key);
    SaveResults saveValueWithRefCount(
        uint32_t updated_ref_count,
        const std::vector<unsigned char>& hash_key,
        const std::vector<unsigned char>& value);

   public:
    Transaction(std::unique_ptr<rocksdb::Transaction> transaction_);
    ~Transaction();
    SaveResults incrementReference(const std::vector<unsigned char>& hash_key);
    SaveResults saveData(const std::vector<unsigned char>& hash_key,
                         const std::vector<unsigned char>& value);
    GetResults getData(const std::vector<unsigned char>& hash_key) const;
    DeleteResults deleteData(const std::vector<unsigned char>& hash_key);
    rocksdb::Status commit();
    rocksdb::Status rollBack();
};

#endif /* transaction_hpp */
