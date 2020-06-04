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

#ifndef datastorage_hpp
#define datastorage_hpp

#include <memory>
#include <string>
#include <vector>

#include <data_storage/keyvaluestore.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/transaction.hpp>

#include <rocksdb/utilities/transaction.h>
#include <rocksdb/utilities/transaction_db.h>

class BlockStore;
class NodeStore;

class DataStorage {
   private:
    std::string txn_db_path;
    std::shared_ptr<rocksdb::TransactionDB> txn_db;
    std::shared_ptr<rocksdb::ColumnFamilyHandle> default_column;
    std::shared_ptr<rocksdb::ColumnFamilyHandle> blocks_column;
    std::shared_ptr<rocksdb::ColumnFamilyHandle> nodes_column;

   public:
    DataStorage(const std::string& db_path);
    rocksdb::Status closeDb();
    GetResults getValue(const std::vector<unsigned char>& hash_key) const;
    std::unique_ptr<Transaction> makeTransaction();
    std::unique_ptr<KeyValueStore> makeKeyValueStore();

    std::unique_ptr<BlockStore> getBlockStore() const;
    std::unique_ptr<NodeStore> getNodeStore() const;
};

#endif /* datastorage_hpp */
