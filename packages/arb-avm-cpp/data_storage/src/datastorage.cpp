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
#include <rocksdb/filter_policy.h>
#include <avm_values/value.hpp>
#include <data_storage/storageresult.hpp>

#include <string>

DataStorage::DataStorage(const std::string& db_path) {
    rocksdb::TransactionDBOptions txn_options;
    rocksdb::Options options = rocksdb::Options();
    rocksdb::ColumnFamilyOptions cf_options;
    rocksdb::ColumnFamilyOptions small_cf_options;
    rocksdb::ColumnFamilyOptions refcounted_cf_options;
    rocksdb::BlockBasedTableOptions table_options;
    rocksdb::BlockBasedTableOptions bloom_table_options;
    options.create_if_missing = true;
    options.create_missing_column_families = true;

    // As recommended for new applications by
    // https://github.com/facebook/rocksdb/wiki/Setup-Options-and-Basic-Tuning
    cf_options.level_compaction_dynamic_level_bytes = true;
    options.max_background_compactions = 4;
    options.max_background_flushes = 2;
    options.bytes_per_sync = 1048576;
    options.compaction_pri = rocksdb::kMinOverlappingRatio;
    table_options.block_size = 16 * 1024;
    table_options.cache_index_and_filter_blocks = true;
    table_options.pin_l0_filter_and_index_blocks_in_cache = true;
    table_options.format_version = 5;
    options.table_factory.reset(
        rocksdb::NewBlockBasedTableFactory(table_options));

    // Increase the number of threads to open files to offset slow disk access
    options.max_file_opening_threads = 50;

    // Decrease the WAL log size to 50MB so that DB is flushed regularly
    options.max_total_wal_size = 52428800;

    // No need to wait for manual flush to finish
    flush_options.wait = false;

    // Settings for small tables
    small_cf_options = cf_options;
    small_cf_options.OptimizeForSmallDb();

    bloom_table_options = table_options;
    // bloom_table_options.filter_policy.reset(
    //   rocksdb::NewBloomFilterPolicy(10, false));
    // bloom_table_options.optimize_filters_for_memory = true;

    // Settings for refcounted data table using bloom filters and no iterators
    refcounted_cf_options = cf_options;
    refcounted_cf_options.OptimizeForPointLookup(16);
    refcounted_cf_options.level_compaction_dynamic_level_bytes = true;
    refcounted_cf_options.table_factory =
        std::unique_ptr<rocksdb::TableFactory>(
            rocksdb::NewBlockBasedTableFactory(bloom_table_options));

    txn_db_path = db_path;

    std::vector<rocksdb::ColumnFamilyDescriptor> column_descriptors{
        FAMILY_COLUMN_COUNT};
    column_descriptors[DEFAULT_COLUMN] = {rocksdb::kDefaultColumnFamilyName,
                                          cf_options};
    column_descriptors[STATE_COLUMN] = {"states", small_cf_options};
    column_descriptors[CHECKPOINT_COLUMN] = {"checkpoints", cf_options};
    column_descriptors[MESSAGEENTRY_COLUMN] = {"messageentries", cf_options};
    column_descriptors[LOG_COLUMN] = {"logs", cf_options};
    column_descriptors[SEND_COLUMN] = {"sends", cf_options};
    column_descriptors[SIDELOAD_COLUMN] = {"sideloads", cf_options};
    column_descriptors[AGGREGATOR_COLUMN] = {"aggregator", cf_options};
    column_descriptors[REFCOUNTED_COLUMN] = {"refcounted",
                                             refcounted_cf_options};

    rocksdb::TransactionDB* db = nullptr;
    auto status =
        rocksdb::TransactionDB::Open(options, txn_options, txn_db_path,
                                     column_descriptors, &column_handles, &db);

    if (!status.ok()) {
        throw std::runtime_error(status.ToString());
    }
    assert(status.ok());
    txn_db = std::unique_ptr<rocksdb::TransactionDB>(db);
}

rocksdb::Status DataStorage::flushNextColumn() {
    next_column_to_flush = (next_column_to_flush + 1) % column_handles.size();
    return txn_db->Flush(flush_options, column_handles[next_column_to_flush]);
}

rocksdb::Status DataStorage::closeDb() {
    column_handles.clear();
    auto s = txn_db->Close();
    txn_db.reset();
    return s;
}

std::unique_ptr<Transaction> Transaction::makeTransaction(
    std::shared_ptr<DataStorage> store) {
    auto tx = store->beginTransaction();
    return std::make_unique<Transaction>(std::move(store), std::move(tx));
}
