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

#include "avm/checkpointstorage.hpp"

#include "rocksdb/options.h"
#include "rocksdb/utilities/transaction.h"

using UCharVec = std::vector<unsigned char>;

std::string dbPath = "tmp/rocksDbPath";
std::string machine_code_key = "machine_code";

CheckpointStorage::CheckpointStorage() {
    rocksdb::Options options;
    rocksdb::TransactionDBOptions txn_options;
    options.create_if_missing = true;

    rocksdb::TransactionDB::Open(options, txn_options, dbPath, &txn_db);
};

CheckpointStorage::~CheckpointStorage() {
    delete txn_db;
}

SaveResults CheckpointStorage::incrementReference(
    std::vector<unsigned char> hash_key) {
    auto results = getRawData(hash_key);

    if (results.status.ok()) {
        auto updated_count = results.reference_count + 1;
        auto updated_entry =
            SerializeCountAndValue(updated_count, results.stored_value);
        std::string key_str(hash_key.begin(), hash_key.end());
        auto commit_status = SaveKeyValuePair(updated_entry, key_str);

        if (commit_status.ok()) {
            return SaveResults{updated_count, commit_status};
        } else {
            return SaveResults{results.reference_count, commit_status,
                               hash_key};
        }
    } else {
        return SaveResults{-1, results.status, hash_key};
    }
}

SaveResults CheckpointStorage::saveValue(std::string val,
                                         std::vector<unsigned char> hash_key) {
    auto results = getRawData(hash_key);
    auto ref_count = results.reference_count;

    if (!results.status.ok() || ref_count < 1) {
        ref_count = 1;
    } else {
        ref_count += 1;
    }

    auto updated_entry = SerializeCountAndValue(ref_count, val);
    std::string key_str(hash_key.begin(), hash_key.end());

    auto commit_status = SaveKeyValuePair(updated_entry, key_str);

    if (commit_status.ok()) {
        SaveResults save_results{ref_count, commit_status, hash_key};

        return save_results;
    } else {
        auto unsuccessful = rocksdb::Status().NotFound();
        ref_count -= 1;
        SaveResults save_results{ref_count, unsuccessful, hash_key};

        // log
        return save_results;
    }
};

// DeleteResults CheckpointStorage::decrementReference(std::vector<unsigned
// char> hash_key){
//    auto results = getRawData(hash_key);
//
//    if(results.status.ok()){
//        auto updated_count = results.reference_count-1;
//
//        if(updated_count < 1){
//
//        }
//    }else{
//
//    }
//}

DeleteResults CheckpointStorage::deleteStoredValue(
    std::vector<unsigned char> hash_key) {
    auto results = getRawData(hash_key);

    if (results.status.ok()) {
        auto value = results.stored_value;

        if (results.reference_count < 2) {
            auto delete_status =
                DeleteValue(std::string(hash_key.begin(), hash_key.end()));
            return DeleteResults{0, delete_status};

        } else {
            auto updated_ref_count = results.reference_count - 1;
            auto updated_entry =
                SerializeCountAndValue(updated_ref_count, value);
            std::string key_str(hash_key.begin(), hash_key.end());
            auto status = SaveKeyValuePair(updated_entry, key_str);

            return DeleteResults{updated_ref_count, status};
        }
    } else {
        return DeleteResults{0, rocksdb::Status().NotFound()};
    }
}

GetResults CheckpointStorage::getStoredValue(
    std::vector<unsigned char> hash_key) {
    auto results = getRawData(hash_key);

    return GetResults{results.reference_count, results.status,
                      results.stored_value};
}

RawData CheckpointStorage::getRawData(std::vector<unsigned char> hash_key) {
    auto read_options = rocksdb::ReadOptions();
    std::string return_value;
    std::string key_str(hash_key.begin(), hash_key.end());
    auto get_status = txn_db->Get(read_options, key_str, &return_value);

    if (get_status.ok()) {
        auto tuple = ParseCountAndValue(return_value);

        RawData results{std::get<0>(tuple), get_status, hash_key,
                        std::get<1>(tuple)};

        return results;
    } else {
        // make sure this is correct
        auto unsuccessful = rocksdb::Status().NotFound();
        RawData results{0, unsuccessful, std::vector<unsigned char>(),
                        std::string()};

        return results;
    }
}

std::tuple<int, std::string> CheckpointStorage::ParseCountAndValue(
    std::string string_value) {
    // is max 256 references good enough?
    const char* c_string = string_value.c_str();
    uint16_t ref_count;
    memcpy(&ref_count, c_string, sizeof(ref_count));
    auto saved_value = string_value.substr(1, string_value.size() - 1);

    return std::make_tuple(ref_count, saved_value);
}

std::string CheckpointStorage::SerializeCountAndValue(int count,
                                                      std::string value) {
    std::string str;
    if (count > 255) {
        // error
    } else {
        auto count_as_char = (char)count;
        // does not replace
        value.insert(value.begin(), count_as_char);
    }

    return str;
}

rocksdb::Status CheckpointStorage::SaveKeyValuePair(std::string value,
                                                    std::string key) {
    rocksdb::WriteOptions writeOptions;
    rocksdb::Transaction* transaction = txn_db->BeginTransaction(writeOptions);
    assert(transaction);

    auto put_status = transaction->Put(key, value);
    assert(put_status.ok());

    auto commit_status = transaction->Commit();
    assert(commit_status.ok());

    delete transaction;

    return commit_status;
}

rocksdb::Status CheckpointStorage::DeleteValue(std::string key) {
    rocksdb::WriteOptions writeOptions;
    rocksdb::Transaction* transaction = txn_db->BeginTransaction(writeOptions);
    assert(transaction);

    auto delete_status = transaction->Delete(key);
    assert(delete_status.ok());

    auto commit_status = transaction->Commit();
    assert(commit_status.ok());

    delete transaction;

    return commit_status;
}
