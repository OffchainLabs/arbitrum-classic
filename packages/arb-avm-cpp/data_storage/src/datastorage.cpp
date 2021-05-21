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

#include <rocksdb/convenience.h>
#include <rocksdb/filter_policy.h>
#include <data_storage/storageresult.hpp>

#include <iostream>
#include <string>

DataStorage::DataStorage(const std::string& db_path) {
    rocksdb::TransactionDBOptions txn_options{};
    rocksdb::Options options{};
    rocksdb::ColumnFamilyOptions cf_options{};
    rocksdb::ColumnFamilyOptions small_cf_options{};
    rocksdb::ColumnFamilyOptions hashkey_cf_options{};
    rocksdb::BlockBasedTableOptions table_options{};
    options.create_if_missing = true;
    options.create_missing_column_families = true;

    cf_options.target_file_size_multiplier = 10;

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
    options.table_factory.reset(
        rocksdb::NewBlockBasedTableFactory(table_options));
    table_options.format_version = 4;

    options.compression = rocksdb::CompressionType::kZSTD;

    // No need to keep old log files
    options.keep_log_file_num = 3;

    // Settings for small tables
    small_cf_options = cf_options;
    small_cf_options.num_levels = 2;
    small_cf_options.OptimizeForSmallDb();

    // Settings for refcounted data table
    hashkey_cf_options = cf_options;
    //

    // OptimizeForPointLookup slows down when database gets large
    // hashkey_cf_options.OptimizeForPointLookup(16);
    // hashkey_cf_options.level_compaction_dynamic_level_bytes = true;

    txn_db_path = db_path;

    std::vector<rocksdb::ColumnFamilyDescriptor> column_descriptors{
        FAMILY_COLUMN_COUNT};
    column_descriptors[DEFAULT_COLUMN] = {rocksdb::kDefaultColumnFamilyName,
                                          small_cf_options};
    column_descriptors[STATE_COLUMN] = {"states", small_cf_options};
    column_descriptors[CHECKPOINT_COLUMN] = {"checkpoints", cf_options};
    column_descriptors[DELAYEDMESSAGE_COLUMN] = {"delayedmessages", cf_options};
    column_descriptors[SEQUENCERBATCHITEM_COLUMN] = {"sequencerbatchitems",
                                                     cf_options};
    column_descriptors[SEQUENCERBATCH_COLUMN] = {"sequencerbatches",
                                                 cf_options};
    column_descriptors[LOG_COLUMN] = {"logs", cf_options};
    column_descriptors[SEND_COLUMN] = {"sends", cf_options};
    column_descriptors[SIDELOAD_COLUMN] = {"sideloads", cf_options};
    column_descriptors[AGGREGATOR_COLUMN] = {"aggregator", cf_options};
    column_descriptors[REFCOUNTED_COLUMN] = {"refcounted", hashkey_cf_options};

    rocksdb::TransactionDB* db = nullptr;
    auto status =
        rocksdb::TransactionDB::Open(options, txn_options, txn_db_path,
                                     column_descriptors, &column_handles, &db);

    if (!status.ok()) {
        throw std::runtime_error(status.ToString());
    }
    assert(status.ok());

    // Compact all family columns on startup
    // Disabled because of concern over leaving the db in a bad state if
    // terminated during compaction

    //    auto cr_options = rocksdb::CompactRangeOptions();
    //    for (size_t i = 0; i < FAMILY_COLUMN_COUNT; i++) {
    //        db->CompactRange(cr_options, column_handles[i], nullptr, nullptr);
    //    }

    txn_db = std::unique_ptr<rocksdb::TransactionDB>(db);
}

rocksdb::Status DataStorage::flushNextColumn() {
    next_column_to_flush = (next_column_to_flush + 1) % column_handles.size();
    return txn_db->Flush(flush_options, column_handles[next_column_to_flush]);
}

rocksdb::Status DataStorage::closeDb() {
    column_handles.clear();
    if (txn_db) {
        auto s = txn_db->Close();
        txn_db.reset();
        return s;
    }

    return rocksdb::Status::OK();
}

std::unique_ptr<Transaction> Transaction::makeTransaction(
    std::shared_ptr<DataStorage> store) {
    auto tx = store->beginTransaction();
    return std::make_unique<Transaction>(std::move(store), std::move(tx));
}

rocksdb::Status DataStorage::clearDBExceptInbox() {
    for (int i = 0; i < FAMILY_COLUMN_COUNT; i++) {
        if (i == DEFAULT_COLUMN || i == DELAYEDMESSAGE_COLUMN ||
            i == SEQUENCERBATCHITEM_COLUMN || i == SEQUENCERBATCH_COLUMN) {
            continue;
        }
        auto s = txn_db->DropColumnFamily(column_handles[i]);
        if (!s.ok()) {
            return s;
        }
    }
    return rocksdb::Status::OK();
}
