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

#include <data_storage/storageresult.hpp>
#include <data_storage/transaction.hpp>

#include <rocksdb/utilities/transaction_db.h>

#include <boost/algorithm/hex.hpp>

#include <iostream>

std::tuple<uint32_t, std::vector<unsigned char>> parseCountAndValue(
    const std::string& string_value) {
    if (string_value.empty()) {
        return std::make_tuple(0, std::vector<unsigned char>());
    } else {
        const char* c_string = string_value.c_str();
        uint32_t ref_count;
        memcpy(&ref_count, c_string, sizeof(ref_count));
        std::vector<unsigned char> saved_value(
            string_value.begin() + sizeof(ref_count), string_value.end());

        return std::make_tuple(ref_count, saved_value);
    }
}

std::vector<unsigned char> serializeCountAndValue(
    int32_t count,
    const std::vector<unsigned char>& value) {
    std::vector<unsigned char> output_vector(sizeof(count));
    memcpy(&output_vector[0], &count, sizeof(count));
    output_vector.insert(output_vector.end(), value.begin(), value.end());

    return output_vector;
}

Transaction::Transaction(std::unique_ptr<rocksdb::Transaction> transaction_)
    : transaction(std::move(transaction_)) {}

Transaction::~Transaction() {
    assert(transaction->GetState() == rocksdb::Transaction::COMMITED ||
           transaction->GetState() == rocksdb::Transaction::ROLLEDBACK);
}

rocksdb::Status Transaction::commit() {
    return transaction->Commit();
}

rocksdb::Status Transaction::rollBack() {
    return transaction->Rollback();
}

SaveResults Transaction::incrementReference(
    const std::vector<unsigned char>& hash_key) {
    auto results = getData(hash_key);

    if (results.status.ok()) {
        auto updated_count = results.reference_count + 1;
        return saveValueWithRefCount(updated_count, hash_key,
                                     results.stored_value);
    } else {
        return SaveResults{0, results.status, hash_key};
    }
}

SaveResults Transaction::saveData(const std::vector<unsigned char>& hash_key,
                                  const std::vector<unsigned char>& value) {
    auto results = getData(hash_key);
    int ref_count;

    if (results.status.ok()) {
        if (results.stored_value != value) {
            std::cout << "Different value for key: ";
            boost::algorithm::hex(hash_key.begin(), hash_key.end(),
                                  std::ostream_iterator<char>{std::cout, ""});
            std::cout << "\nPrevious value: ";
            boost::algorithm::hex(results.stored_value.begin(),
                                  results.stored_value.end(),
                                  std::ostream_iterator<char>{std::cout, ""});
            std::cout << "\nNew Value: ";
            boost::algorithm::hex(value.begin(), value.end(),
                                  std::ostream_iterator<char>{std::cout, ""});
            std::cout << std::endl;
        }
        assert(results.stored_value == value);
        ref_count = results.reference_count + 1;
    } else {
        ref_count = 1;
    }
    return saveValueWithRefCount(ref_count, hash_key, value);
}

DeleteResults Transaction::deleteData(
    const std::vector<unsigned char>& hash_key) {
    auto results = getData(hash_key);

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

GetResults Transaction::getData(
    const std::vector<unsigned char>& hash_key) const {
    auto read_options = rocksdb::ReadOptions();
    std::string return_value;
    std::string key_str(hash_key.begin(), hash_key.end());
    auto get_status = transaction->Get(read_options, key_str, &return_value);

    if (get_status.ok()) {
        auto parsed_values = parseCountAndValue(return_value);
        auto stored_val = std::get<1>(parsed_values);
        auto ref_count = std::get<0>(parsed_values);

        return GetResults{ref_count, get_status, stored_val};
    } else {
        auto unsuccessful = rocksdb::Status().NotFound();
        return GetResults{0, unsuccessful, std::vector<unsigned char>()};
    }
}

// private ------------------------------------------------------------------
SaveResults Transaction::saveValueWithRefCount(
    uint32_t updated_ref_count,
    const std::vector<unsigned char>& hash_key,
    const std::vector<unsigned char>& value) {
    auto updated_entry = serializeCountAndValue(updated_ref_count, value);

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
