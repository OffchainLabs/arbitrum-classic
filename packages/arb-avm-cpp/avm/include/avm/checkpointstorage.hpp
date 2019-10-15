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

#include <vector>
#include "rocksdb/db.h"
#include "rocksdb/utilities/transaction_db.h"

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
    std::unique_ptr<rocksdb::TransactionDB*> txn_db;
    rocksdb::Transaction* makeTransaction();
    rocksdb::Status saveValueToDb(std::string value, std::string key);
    rocksdb::Status deleteValueFromDb(std::string key);
    std::tuple<int, std::string> parseCountAndValue(std::string string_value);
    std::vector<unsigned char> serializeCountAndValue(
        int count,
        std::vector<unsigned char> value);
    SaveResults saveValueWithRefCount(int new_count,
                                      std::vector<unsigned char> hash_key,
                                      std::vector<unsigned char> value);

   public:
    CheckpointStorage(std::string db_path);
    ~CheckpointStorage();

    SaveResults incrementReference(std::vector<unsigned char> hash_key);
    SaveResults saveValue(std::vector<unsigned char> value,
                          std::vector<unsigned char> hash_key);
    GetResults getStoredValue(std::vector<unsigned char> hash_key);
    DeleteResults deleteStoredValue(std::vector<unsigned char> hash_key);
};

#endif /* checkpointstorage_hpp */
