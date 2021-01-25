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

#include <data_storage/datastorage.hpp>

#include "value/utils.hpp"

#include <avm_values/value.hpp>
#include <data_storage/blockstore.hpp>
#include <data_storage/storageresult.hpp>

#include <string>

DataStorage::DataStorage(const std::string& db_path) {
    rocksdb::TransactionDBOptions txn_options;
    rocksdb::Options options;
    options.create_if_missing = true;
    options.create_missing_column_families = true;

    txn_db_path = db_path;

    std::vector<rocksdb::ColumnFamilyDescriptor> column_families;
    column_families.emplace_back(rocksdb::kDefaultColumnFamilyName,
                                 rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("states", rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("blocks", rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("nodes", rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("checkpoints", rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("messageentries",
                                 rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("logs", rocksdb::ColumnFamilyOptions());
    column_families.emplace_back("sends", rocksdb::ColumnFamilyOptions());

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
    blocks_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[2]);
    node_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[3]);
    checkpoint_column =
        std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[4]);
    messageentry_column =
        std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[5]);
    log_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[6]);
    send_column = std::unique_ptr<rocksdb::ColumnFamilyHandle>(handles[7]);
}

rocksdb::Status DataStorage::closeDb() {
    default_column.reset();
    state_column.reset();
    blocks_column.reset();
    node_column.reset();
    checkpoint_column.reset();
    messageentry_column.reset();
    log_column.reset();
    send_column.reset();
    auto s = txn_db->Close();
    txn_db.reset();
    return s;
}

std::unique_ptr<Transaction> Transaction::makeTransaction(
    std::shared_ptr<DataStorage> store) {
    auto tx = store->beginTransaction();
    return std::make_unique<Transaction>(std::move(store), std::move(tx));
}

ValueResult<std::vector<std::vector<unsigned char>>>
getVectorVectorUsingFamilyAndKey(rocksdb::Transaction& tx,
                                 rocksdb::ColumnFamilyHandle* family,
                                 rocksdb::Slice first_key_slice,
                                 size_t count) {
    auto it = std::unique_ptr<rocksdb::Iterator>(
        tx.GetIterator(rocksdb::ReadOptions(), family));

    // Find first message
    it->Seek(vecToSlice(first_key_slice));
    if (!it->status().ok()) {
        return {it->status(), {}};
    }

    std::vector<std::vector<unsigned char>> vectors;
    for (size_t i = 0; i < count; i++) {
        if (!it->Valid()) {
            if (!it->status().ok()) {
                return {it->status(), {}};
            }
            return {rocksdb::Status::NotFound(), {}};
        }
        vectors.emplace_back(it->value().data(),
                             it->value().data() + it->value().size());
    }

    return {rocksdb::Status::OK(), vectors};
}

ValueResult<std::vector<unsigned char>> getVectorUsingFamilyAndKey(
    rocksdb::Transaction& tx,
    rocksdb::ColumnFamilyHandle* family,
    rocksdb::Slice key_slice) {
    std::string returned_value;

    auto status =
        tx.Get(rocksdb::ReadOptions(), family, key_slice, &returned_value);
    if (!status.ok()) {
        return {status, {}};
    }

    std::vector<unsigned char> saved_value(returned_value.begin(),
                                           returned_value.end());

    return {status, saved_value};
}

ValueResult<std::vector<uint256_t>> getUint256VectorUsingFamilyAndKey(
    rocksdb::Transaction& tx,
    rocksdb::ColumnFamilyHandle* family,
    rocksdb::Slice first_key_slice,
    size_t count) {
    auto it = std::unique_ptr<rocksdb::Iterator>(
        tx.GetIterator(rocksdb::ReadOptions(), family));

    // Find first message
    it->Seek(vecToSlice(first_key_slice));
    if (!it->status().ok()) {
        return {it->status(), {}};
    }

    std::vector<uint256_t> vectors;
    for (size_t i = 0; i < count; i++) {
        if (!it->Valid()) {
            if (!it->status().ok()) {
                return {it->status(), {}};
            }
            return {rocksdb::Status::NotFound(), {}};
        }

        auto data = reinterpret_cast<const char*>(it->value().data());
        vectors.push_back(deserializeUint256t(data));
    }

    return {rocksdb::Status::OK(), vectors};
}

ValueResult<uint256_t> getUint256UsingFamilyAndKey(
    rocksdb::Transaction& tx,
    rocksdb::ColumnFamilyHandle* family,
    rocksdb::Slice key_slice) {
    auto result = getVectorUsingFamilyAndKey(tx, family, key_slice);
    if (!result.status.ok()) {
        return {result.status, {}};
    }

    auto data = reinterpret_cast<const char*>(result.data.data());
    return {result.status, deserializeUint256t(data)};
}
