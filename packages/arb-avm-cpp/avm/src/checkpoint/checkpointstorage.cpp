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

#include <avm/checkpoint/checkpointstorage.hpp>

#include <avm/checkpoint/checkpointresult.hpp>
#include <avm/checkpoint/checkpointutils.hpp>
#include <avm/checkpoint/transaction.hpp>

#include <avm/value/codepoint.hpp>
#include <avm/value/tuple.hpp>

#include <rocksdb/options.h>
#include <rocksdb/utilities/transaction.h>
#include <rocksdb/utilities/transaction_db.h>

CheckpointStorage::CheckpointStorage(std::string db_path) {
    rocksdb::TransactionDBOptions txn_options;
    rocksdb::Options options;
    options.create_if_missing = true;

    txn_db_path = std::move(db_path);
    rocksdb::TransactionDB* db = nullptr;
    auto status =
        rocksdb::TransactionDB::Open(options, txn_options, txn_db_path, &db);
    if (!status.ok()) {
        std::cout << "db status " << status.ToString() << std::endl;
    }
    assert(status.ok());
    txn_db = std::unique_ptr<rocksdb::TransactionDB>(db);
}

CheckpointStorage::~CheckpointStorage() {
    txn_db->Close();
}

GetResults CheckpointStorage::getValue(
    const std::vector<unsigned char>& hash_key) const {
    auto read_options = rocksdb::ReadOptions();
    std::string return_value;
    std::string key_str(hash_key.begin(), hash_key.end());
    auto get_status = txn_db->Get(read_options, key_str, &return_value);

    if (get_status.ok()) {
        auto tuple = checkpoint::storage::parseCountAndValue(return_value);
        auto stored_val = std::get<1>(tuple);
        auto ref_count = std::get<0>(tuple);

        return GetResults{ref_count, get_status, stored_val};
    } else {
        auto unsuccessful = rocksdb::Status().NotFound();
        return GetResults{0, unsuccessful, std::vector<unsigned char>()};
    }
}

std::unique_ptr<Transaction> CheckpointStorage::makeTransaction() {
    rocksdb::WriteOptions writeOptions;
    rocksdb::Transaction* transaction = txn_db->BeginTransaction(writeOptions);
    return std::make_unique<Transaction>(transaction);
}
