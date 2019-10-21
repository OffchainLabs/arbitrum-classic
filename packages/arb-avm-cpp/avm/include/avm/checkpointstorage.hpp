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

#ifndef checkpointstorage_hpp
#define checkpointstorage_hpp

#include <rocksdb/db.h>
#include <rocksdb/utilities/transaction_db.h>
#include <vector>

struct GetResults {
    int reference_count;
    rocksdb::Status status;
    std::vector<unsigned char> stored_value;
};

struct SaveResults {
    int reference_count;
    rocksdb::Status status;
    std::vector<unsigned char> storage_key;
};

struct DeleteResults {
    int reference_count;
    rocksdb::Status status;
};

class CheckpointStorage {
   private:
    std::string txn_db_path;
    std::unique_ptr<rocksdb::TransactionDB> txn_db;
    std::unique_ptr<rocksdb::Transaction> makeTransaction();
    rocksdb::Status saveValue(const std::string& value, const std::string& key);
    rocksdb::Status deleteValue(const std::string& key);
    std::tuple<int, std::vector<unsigned char>> parseCountAndValue(
        const std::string& string_value);
    std::vector<unsigned char> serializeCountAndValue(
        int count,
        const std::vector<unsigned char>& value);
    SaveResults saveValueWithRefCount(
        int new_count,
        const std::vector<unsigned char>& hash_key,
        const std::vector<unsigned char>& value);

   public:
    CheckpointStorage(std::string db_path);
    ~CheckpointStorage();

    SaveResults incrementReference(const std::vector<unsigned char>& hash_key);
    SaveResults saveValue(const std::vector<unsigned char>& value,
                          const std::vector<unsigned char>& hash_key);
    GetResults getValue(const std::vector<unsigned char>& hash_key);
    DeleteResults deleteValue(const std::vector<unsigned char>& hash_key);
};

#endif /* checkpointstorage_hpp */
