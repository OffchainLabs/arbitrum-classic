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

#ifndef datastorage_hpp
#define datastorage_hpp

#include <memory>
#include <string>
#include <utility>
#include <vector>

#include <avm_values/bigint.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/util.hpp>

#include <rocksdb/utilities/transaction.h>
#include <rocksdb/utilities/transaction_db.h>

class Transaction;
class ReadTransaction;
class DataStorage;

class DbLockShared {
    friend DataStorage;

   private:
    const DataStorage* storage;

    DbLockShared(const DataStorage* storage);
    DbLockShared(const DbLockShared& other) = delete;
    DbLockShared(DbLockShared&& other) noexcept;
    DbLockShared& operator=(const DbLockShared& other) = delete;
    DbLockShared& operator=(DbLockShared&& other) noexcept;

   public:
    ~DbLockShared();
};

class DataStorage {
    friend Transaction;
    friend DbLockShared;

   public:
    enum column_family_indexes {
        DEFAULT_COLUMN = 0,
        STATE_COLUMN,
        CHECKPOINT_COLUMN,
        DELAYEDMESSAGE_COLUMN,
        SEQUENCERBATCHITEM_COLUMN,
        SEQUENCERBATCH_COLUMN,
        LOG_COLUMN,
        SEND_COLUMN,
        SIDELOAD_COLUMN,
        AGGREGATOR_COLUMN,
        REFCOUNTED_COLUMN,
        FAMILY_COLUMN_COUNT
    };
    std::string txn_db_path;
    std::unique_ptr<rocksdb::TransactionDB> txn_db;
    std::vector<rocksdb::ColumnFamilyHandle*> column_handles;
    std::vector<uint8_t> secret_hash_seed;

    class shutting_down_exception : public std::exception {};

    explicit DataStorage(const std::string& db_path,
                         const ArbCoreConfig& coreConfig);
    ~DataStorage();

    rocksdb::Status closeDb();
    rocksdb::Status clearDBExceptInbox();
    [[nodiscard]] DbLockShared tryLockShared() const;
    rocksdb::Status cleanupValidator();
    rocksdb::Status compact(bool aggressive);

   private:
    std::atomic<bool> shutting_down{false};
    mutable std::atomic<int64_t> concurrent_database_access_counter{0};

    rocksdb::Status updateSecretHashSeed();

    [[nodiscard]] std::unique_ptr<rocksdb::Transaction> beginTransaction() {
        // Make sure database isn't closed while it is being used
        auto counter = tryLockShared();

        return std::unique_ptr<rocksdb::Transaction>{
            txn_db->BeginTransaction(rocksdb::WriteOptions())};
    }
};

class Transaction {
    friend ReadTransaction;

   public:
    std::shared_ptr<DataStorage> datastorage;
    std::unique_ptr<rocksdb::Transaction> transaction;

    Transaction(std::shared_ptr<DataStorage> datastorage_,
                std::unique_ptr<rocksdb::Transaction> transaction_)
        : datastorage(std::move(datastorage_)),
          transaction(std::move(transaction_)) {}

    rocksdb::Status commit() {
        // Make sure database isn't closed while it is being used
        auto counter = datastorage->tryLockShared();

        return transaction->Commit();
    }

    rocksdb::Status rollback() {
        // Make sure database isn't closed while it is being used
        auto counter = datastorage->tryLockShared();

        return transaction->Rollback();
    }

   private:
    static std::unique_ptr<Transaction> makeTransaction(
        std::shared_ptr<DataStorage> store);
};

#endif /* datastorage_hpp */
