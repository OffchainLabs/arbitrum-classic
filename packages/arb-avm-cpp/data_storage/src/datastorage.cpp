/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

#include <data_storage/datastorage.hpp>

#include "value/utils.hpp"

#include <rocksdb/convenience.h>
#include <avm_values/value.hpp>
#include <data_storage/storageresult.hpp>

#include <string>

DataStorage::DataStorage(const std::string& db_path) {
    rocksdb::TransactionDBOptions txn_options;
    rocksdb::Options options = rocksdb::Options();
    rocksdb::ColumnFamilyOptions cf_options;
    rocksdb::BlockBasedTableOptions table_options;
    options.create_if_missing = true;
    options.create_missing_column_families = true;

    // As recommended for new applications by
    // https://github.com/facebook/rocksdb/wiki/Setup-Options-and-Basic-Tuning
    //    cf_options.level_compaction_dynamic_level_bytes = true;
    //    options.max_background_compactions = 4;
    //    options.max_background_flushes = 2;
    //    options.bytes_per_sync = 1048576;
    //    options.compaction_pri = rocksdb::kMinOverlappingRatio;
    //    table_options.block_size = 16 * 1024;
    //    table_options.cache_index_and_filter_blocks = true;
    //    table_options.pin_l0_filter_and_index_blocks_in_cache = true;
    //    table_options.format_version = 5;
    //    options.table_factory.reset(
    //        rocksdb::NewBlockBasedTableFactory(table_options));

    txn_db_path = db_path;

    std::vector<rocksdb::ColumnFamilyDescriptor> column_families;
    column_families.emplace_back(rocksdb::kDefaultColumnFamilyName, cf_options);
    column_families.emplace_back("states", cf_options);
    column_families.emplace_back("checkpoints", cf_options);
    column_families.emplace_back("messageentries", cf_options);
    column_families.emplace_back("logs", cf_options);
    column_families.emplace_back("sends", cf_options);
    column_families.emplace_back("sideloads", cf_options);
    column_families.emplace_back("aggregator", cf_options);

    rocksdb::TransactionDB* db = nullptr;
    std::vector<rocksdb::ColumnFamilyHandle*> handles;
    auto status = rocksdb::TransactionDB::Open(
        options, txn_options, txn_db_path, column_families, &handles, &db);

    if (!status.ok()) {
        throw std::runtime_error(status.ToString());
    }
    assert(status.ok());
    txn_db = std::unique_ptr<rocksdb::TransactionDB>(db);
    default_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[0]);
    state_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[1]);
    checkpoint_column =
        std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[2]);
    messageentry_column =
        std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[3]);
    log_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[4]);
    send_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[5]);
    sideload_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[6]);
    aggregator_column =
        std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[7]);
}

rocksdb::Status DataStorage::closeDb() {
    default_column.reset();
    state_column.reset();
    checkpoint_column.reset();
    messageentry_column.reset();
    log_column.reset();
    send_column.reset();
    sideload_column.reset();
    aggregator_column.reset();
    auto s = txn_db->Close();
    txn_db.reset();
    return s;
}

std::unique_ptr<Transaction> Transaction::makeTransaction(
    std::shared_ptr<DataStorage> store) {
    auto tx = store->beginTransaction();
    return std::make_unique<Transaction>(std::move(store), std::move(tx));
}
