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
    int reference_count = 0;
    rocksdb::Status status;
    std::vector<unsigned char> storage_key;
    std::string stored_value;
};

struct SaveResults {
    rocksdb::Status status;
    std::vector<unsigned char> storage_key;
    std::string stored_value;
};

struct DeleteResults {
    rocksdb::Status status;
    std::vector<unsigned char> storage_key;
    std::string stored_value;
};

class CheckpointStorage {
   private:
    rocksdb::TransactionDB* txn_db;
    rocksdb::Status SaveKeyValuePair(std::string value, std::string key);
    std::tuple<int, std::string> ParseCountAndValue(std::string string_value);
    std::string SerializeCountAndValue(int count, std::string value);
    rocksdb::Status DeleteValue(std::string key);

   public:
    CheckpointStorage();
    ~CheckpointStorage();

    GetResults saveValue(std::string val, std::vector<unsigned char> hash_key);
    GetResults getStoredValue(std::vector<unsigned char> hash_key);
    GetResults deleteStoredValue(std::vector<unsigned char> hash_key);
};

#endif /* checkpointstorage_hpp */
