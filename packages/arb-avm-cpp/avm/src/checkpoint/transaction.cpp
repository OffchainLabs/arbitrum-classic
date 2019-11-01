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

#include <avm/checkpoint/transaction.hpp>

#include <avm/checkpoint/checkpointresult.hpp>
#include <avm/checkpoint/checkpointutils.hpp>

#include <rocksdb/utilities/transaction_db.h>

Transaction::Transaction(rocksdb::Transaction* transaction_) {
    transaction = std::unique_ptr<rocksdb::Transaction>(transaction_);
}

Transaction::~Transaction() {
    assert(transaction->GetState() == rocksdb::Transaction::COMMITED ||
           rocksdb::Transaction::ROLLEDBACK);
}

rocksdb::Status Transaction::commit() {
    return transaction->Commit();
}

rocksdb::Status Transaction::rollBack() {
    return transaction->Rollback();
}

SaveResults Transaction::incrementReference(
    const std::vector<unsigned char>& hash_key) {
    auto results = getValue(hash_key);

    if (results.status.ok()) {
        auto updated_count = results.reference_count + 1;
        return saveValueWithRefCount(updated_count, hash_key,
                                     results.stored_value);
    } else {
        return SaveResults{0, results.status, hash_key};
    }
}

SaveResults Transaction::saveValue(const std::vector<unsigned char>& hash_key,
                                   const std::vector<unsigned char>& value) {
    auto results = getValue(hash_key);
    int ref_count;

    if (results.status.ok()) {
        assert(results.stored_value == value);
        ref_count = results.reference_count + 1;
    } else {
        ref_count = 1;
    }
    return saveValueWithRefCount(ref_count, hash_key, value);
};

DeleteResults Transaction::deleteValue(
    const std::vector<unsigned char>& hash_key) {
    auto results = getValue(hash_key);

    if (results.status.ok()) {
        if (results.reference_count < 2) {
            auto delete_status = deleteKeyValuePair(hash_key);
            return DeleteResults{0, delete_status};

        } else {
            auto updated_ref_count = results.reference_count - 1;
            auto update_result = saveValueWithRefCount(
                updated_ref_count, hash_key, results.stored_value);
            return DeleteResults{updated_ref_count, update_result.status};
        }
    } else {
        return DeleteResults{0, results.status};
    }
}

GetResults Transaction::getValue(
    const std::vector<unsigned char>& hash_key) const {
    auto read_options = rocksdb::ReadOptions();
    std::string return_value;
    std::string key_str(hash_key.begin(), hash_key.end());
    auto get_status = transaction->Get(read_options, key_str, &return_value);

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

// private
// ----------------------------------------------------------------------
SaveResults Transaction::saveValueWithRefCount(
    uint32_t updated_ref_count,
    const std::vector<unsigned char>& hash_key,
    const std::vector<unsigned char>& value) {
    auto updated_entry =
        checkpoint::storage::serializeCountAndValue(updated_ref_count, value);

    auto status = saveKeyValuePair(hash_key, updated_entry);

    if (status.ok()) {
        return SaveResults{updated_ref_count, status, hash_key};
    } else {
        return SaveResults{0, status, hash_key};
    }
}

rocksdb::Status Transaction::saveKeyValuePair(
    const std::vector<unsigned char>& key,
    const std::vector<unsigned char>& value) {
    std::string value_str(value.begin(), value.end());
    std::string key_str(key.begin(), key.end());
    return transaction->Put(key_str, value_str);
}

rocksdb::Status Transaction::deleteKeyValuePair(
    const std::vector<unsigned char>& key) {
    std::string key_str(key.begin(), key.end());
    return transaction->Delete(key_str);
}
