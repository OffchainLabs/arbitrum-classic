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
#include <data_storage/readtransaction.hpp>
#include <data_storage/storageresult.hpp>

#include <openssl/rand.h>
#include <rocksdb/convenience.h>
#include <rocksdb/filter_policy.h>

#include <iostream>
#include <string>
#include <thread>

std::atomic<bool> DataStorage::shutting_down;
std::atomic<uint64_t> DataStorage::concurrent_database_access_counter;

DataStorage::DataStorage(const std::string& db_path) {
    std::atomic_init(&shutting_down, false);
    std::atomic_init(&concurrent_database_access_counter, 0);
    // Make sure database isn't closed while constructor still running
    auto counter = getCounter();

    rocksdb::TransactionDBOptions txn_options{};
    rocksdb::Options options{};
    rocksdb::ColumnFamilyOptions cf_options{};
    rocksdb::ColumnFamilyOptions small_cf_options{};
    rocksdb::ColumnFamilyOptions hashkey_cf_options{};
    rocksdb::BlockBasedTableOptions table_options{};
    options.create_if_missing = true;
    options.create_missing_column_families = true;

    options.compression = rocksdb::CompressionType::kZSTD;
    cf_options.compression = rocksdb::CompressionType::kZSTD;
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

    status = updateSecretHashSeed();
    if (!status.ok()) {
        throw std::runtime_error(status.ToString());
    }
}

DataStorage::~DataStorage() {
    auto status = closeDb();
    if (!status.ok()) {
        std::cerr << "error closing DataStorage: " << status.ToString()
                  << std::endl;
    }
}

rocksdb::Status DataStorage::flushNextColumn() {
    // Make sure database isn't closed while it is being used
    auto counter = getCounter();

    next_column_to_flush = (next_column_to_flush + 1) % column_handles.size();
    return txn_db->Flush(flush_options, column_handles[next_column_to_flush]);
}

rocksdb::Status DataStorage::closeDb() {
    if (txn_db) {
        std::cerr << "closing ArbStorage" << std::endl;
        shutting_down = true;
        auto last_concurrent_counter =
            concurrent_database_access_counter.load();
        for (std::chrono::seconds counter_seconds_left =
                 std::chrono::minutes(10);
             last_concurrent_counter > 0 && counter_seconds_left.count() > 0;
             counter_seconds_left -= std::chrono::seconds(1)) {
            if (counter_seconds_left.count() % 10 == 0) {
                // Print message every 10 seconds
                auto output_minutes =
                    std::chrono::duration_cast<std::chrono::minutes>(
                        counter_seconds_left)
                        .count();
                auto output_seconds =
                    std::chrono::duration_cast<std::chrono::seconds>(
                        counter_seconds_left % std::chrono::minutes(1))
                        .count();
                std::cerr
                    << "waiting up to " << output_minutes << " minutes, "
                    << output_seconds << " seconds for "
                    << last_concurrent_counter
                    << " database operation(s) to finish before shutting down"
                    << std::endl;
            }
            std::this_thread::sleep_for(std::chrono::seconds(1));
            last_concurrent_counter = concurrent_database_access_counter.load();
        }
        for (auto handle : column_handles) {
            auto status = txn_db->DestroyColumnFamilyHandle(handle);
            if (!status.ok()) {
                return status;
            }
        }

        txn_db->SyncWAL();
        auto s = txn_db->Close();
        for (std::chrono::seconds close_seconds_left = std::chrono::minutes(10);
             s.IsAborted() && close_seconds_left.count() > 0;
             close_seconds_left -= std::chrono::seconds(1)) {
            // Try to close database once a second for 30 minutes
            if (close_seconds_left.count() % 10 == 0) {
                // Print message every 10 seconds
                auto output_minutes =
                    std::chrono::duration_cast<std::chrono::minutes>(
                        close_seconds_left)
                        .count();
                auto output_seconds =
                    std::chrono::duration_cast<std::chrono::seconds>(
                        close_seconds_left % std::chrono::minutes(1))
                        .count();
                std::cerr << "waiting up to " << output_minutes << " minutes, "
                          << output_seconds
                          << " seconds for rocksdb snapshots to be freed"
                          << std::endl;
            }
            std::this_thread::sleep_for(std::chrono::seconds(1));
            s = txn_db->Close();
        }
        if (s.IsAborted()) {
            std::cerr << "rocksdb snapshots not freed" << std::endl;
        }
        txn_db.reset();
        std::cerr << "closed ArbStorage" << std::endl;
        return s;
    }
    return rocksdb::Status::OK();
}

std::unique_ptr<Transaction> Transaction::makeTransaction(
    std::shared_ptr<DataStorage> store) {
    // Make sure database isn't closed while it is being used
    auto counter = store->getCounter();

    auto tx = store->beginTransaction();
    return std::make_unique<Transaction>(std::move(store), std::move(tx));
}

rocksdb::Status DataStorage::clearDBExceptInbox() {
    // Make sure database isn't closed while it is being used
    auto counter = getCounter();

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

rocksdb::Status DataStorage::updateSecretHashSeed() {
    std::string key("secretHashSeed");

    // Make sure database isn't closed while it is being used
    auto counter = getCounter();

    rocksdb::PinnableSlice value;
    rocksdb::ReadOptions read_opts;
    auto status =
        txn_db->Get(read_opts, column_handles[STATE_COLUMN], key, &value);
    if (status.IsNotFound()) {
        secret_hash_seed.resize(32);
        RAND_bytes(secret_hash_seed.data(),
                   static_cast<int>(secret_hash_seed.size()));
        rocksdb::WriteOptions write_opts;
        rocksdb::Slice value_slice(
            reinterpret_cast<const char*>(secret_hash_seed.data()),
            secret_hash_seed.size());
        status = txn_db->Put(write_opts, column_handles[STATE_COLUMN], key,
                             value_slice);
    } else if (status.ok()) {
        secret_hash_seed = std::vector<unsigned char>(
            value.data(), value.data() + value.size());
        value.Reset();
    }
    return status;
}
std::unique_ptr<DataStorage::ConcurrentCounter> DataStorage::getCounter() {
    return std::make_unique<ConcurrentCounter>();
}
