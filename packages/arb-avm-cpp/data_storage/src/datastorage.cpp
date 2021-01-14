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

#include <data_storage/blockstore.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <string>

#include <iostream>

DataStorage::DataStorage(const std::string& db_path) {
    rocksdb::TransactionDBOptions txn_options;
    rocksdb::Options options;
    options.create_if_missing = true;
    options.create_missing_column_families = true;

    txn_db_path = db_path;

    std::vector<rocksdb::ColumnFamilyDescriptor> column_families;
    column_families.emplace_back(rocksdb::kDefaultColumnFamilyName,
                                 rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("indexes", rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("blocks", rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("nodes", rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("checkpoints", rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("messageentries",
                                 rocksdb::ColumnFamilyOptions());

    rocksdb::TransactionDB* db = nullptr;
    std::vector<rocksdb::ColumnFamilyHandle*> handles;
    auto status = rocksdb::TransactionDB::Open(
        options, txn_options, txn_db_path, column_families, &handles, &db);

    if (!status.ok()) {
        std::cerr << "rocksdb construction status: " << status.ToString()
                  << std::endl;

        throw std::exception();
    }
    assert(status.ok());
    txn_db = std::unique_ptr<rocksdb::TransactionDB>(db);
    default_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[0]);
    index_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[1]);
    blocks_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[2]);
    node_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[3]);
    checkpoint_column =
        std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[4]);
    messageentry_column =
        std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[5]);
}

rocksdb::Status DataStorage::closeDb() {
    default_column.reset();
    index_column.reset();
    blocks_column.reset();
    node_column.reset();
    checkpoint_column.reset();
    messageentry_column.reset();
    auto s = txn_db->Close();
    txn_db.reset();
    return s;
}

std::unique_ptr<Transaction> Transaction::makeTransaction(
    std::shared_ptr<DataStorage> store) {
    auto tx = store->beginTransaction();
    return std::make_unique<Transaction>(std::move(store), std::move(tx));
}
