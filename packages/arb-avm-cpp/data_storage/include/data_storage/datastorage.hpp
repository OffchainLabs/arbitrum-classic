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

#include <rocksdb/utilities/transaction.h>
#include <rocksdb/utilities/transaction_db.h>

class Transaction;
class ReadTransaction;
class ReadWriteTransaction;

class DataStorage {
    friend Transaction;

   public:
    std::string txn_db_path;
    std::unique_ptr<rocksdb::TransactionDB> txn_db;
    std::unique_ptr<rocksdb::ColumnFamilyHandle> default_column;
    std::unique_ptr<rocksdb::ColumnFamilyHandle> state_column;
    std::unique_ptr<rocksdb::ColumnFamilyHandle> checkpoint_column;
    std::unique_ptr<rocksdb::ColumnFamilyHandle> messageentry_column;
    std::unique_ptr<rocksdb::ColumnFamilyHandle> log_column;
    std::unique_ptr<rocksdb::ColumnFamilyHandle> send_column;
    std::unique_ptr<rocksdb::ColumnFamilyHandle> sideload_column;
    std::unique_ptr<rocksdb::ColumnFamilyHandle> aggregator_column;

    explicit DataStorage(const std::string& db_path);
    rocksdb::Status closeDb();

   private:
    [[nodiscard]] std::unique_ptr<rocksdb::Transaction> beginTransaction()
        const {
        return std::unique_ptr<rocksdb::Transaction>{
            txn_db->BeginTransaction(rocksdb::WriteOptions())};
    }
};

class Transaction {
    friend ReadTransaction;
    friend ReadWriteTransaction;

   public:
    std::shared_ptr<DataStorage> datastorage;
    std::unique_ptr<rocksdb::Transaction> transaction;

    Transaction(std::shared_ptr<DataStorage> datastorage_,
                std::unique_ptr<rocksdb::Transaction> transaction_)
        : datastorage(std::move(datastorage_)),
          transaction(std::move(transaction_)) {}

    rocksdb::Status commit() { return transaction->Commit(); }

    rocksdb::Status rollback() { return transaction->Rollback(); }

   private:
    static std::unique_ptr<Transaction> makeTransaction(
        std::shared_ptr<DataStorage> store);
};

#endif /* datastorage_hpp */
