/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

#include "referencecount.hpp"

#include <data_storage/readwritetransaction.hpp>
#include <data_storage/storageresult.hpp>

#include <rocksdb/utilities/transaction_db.h>

#include <boost/algorithm/hex.hpp>

#include <iostream>

namespace {
struct RefCountResult {
    uint32_t reference_count;
    rocksdb::Status status;
};

RefCountResult getRefCountData(const ReadTransaction& tx,
                               const rocksdb::Slice& hash_key) {
    std::string return_value;
    auto get_status = tx.refCountedGet(hash_key, &return_value);

    if (!get_status.ok()) {
        auto unsuccessful = rocksdb::Status::NotFound();
        return {0, unsuccessful};
    }
    if (return_value.empty()) {
        return {0, get_status};
    }
    const char* c_string = return_value.c_str();
    uint32_t ref_count;
    memcpy(&ref_count, c_string, sizeof(ref_count));
    return {ref_count, get_status};
}
}  // namespace

SaveResults saveValueWithRefCount(ReadWriteTransaction& tx,
                                  uint32_t updated_ref_count,
                                  const rocksdb::Slice& hash_key,
                                  const std::vector<unsigned char>& value) {
    std::vector<unsigned char> updated_entry;
    updated_entry.reserve(sizeof(updated_ref_count) + value.size());
    auto count_data = reinterpret_cast<unsigned char*>(&updated_ref_count);
    updated_entry.insert(updated_entry.end(), count_data,
                         count_data + sizeof(updated_ref_count));
    updated_entry.insert(updated_entry.end(), value.begin(), value.end());

    rocksdb::Slice value_slice(reinterpret_cast<char*>(updated_entry.data()),
                               updated_entry.size());
    auto status = tx.refCountedPut(hash_key, value_slice);

    if (status.ok()) {
        return SaveResults{updated_ref_count, status};
    } else {
        return SaveResults{0, status};
    }
}

SaveResults incrementReference(ReadWriteTransaction& tx,
                               const rocksdb::Slice& hash_key,
                               uint32_t new_references) {
    std::string return_value;
    auto get_status = tx.refCountedGet(hash_key, &return_value);

    if (!get_status.ok()) {
        return SaveResults{0, rocksdb::Status::NotFound()};
    }
    if (return_value.empty()) {
        return SaveResults{0, rocksdb::Status::NotFound()};
    }

    const char* c_string = return_value.c_str();
    uint32_t ref_count;
    memcpy(&ref_count, c_string, sizeof(ref_count));
    ref_count += new_references;
    for (size_t i = 0; i < sizeof(ref_count); i++) {
        return_value[i] = *(reinterpret_cast<char*>(&ref_count) + i);
    }
    auto status = tx.refCountedPut(hash_key, return_value);
    if (status.ok()) {
        return SaveResults{ref_count, status};
    } else {
        return SaveResults{0, status};
    }
}

SaveResults saveRefCountedDataReplaced(ReadWriteTransaction& tx,
                                       const rocksdb::Slice& hash_key,
                                       const std::vector<unsigned char>& value,
                                       uint32_t new_references) {
    auto results = getRefCountData(tx, hash_key);
    uint32_t ref_count;

    if (results.status.ok()) {
        ref_count = results.reference_count + new_references;
    } else {
        ref_count = new_references;
    }
    return saveValueWithRefCount(tx, ref_count, hash_key, value);
}

DeleteResults deleteRefCountedData(ReadWriteTransaction& tx,
                                   const rocksdb::Slice& hash_key,
                                   uint32_t deleted_references) {
    auto results = getRefCountedData(tx, hash_key);

    if (results.status.ok()) {
        if (results.reference_count <= deleted_references) {
            auto delete_status = tx.refCountedDelete(hash_key);
            return DeleteResults{0, delete_status,
                                 std::move(results.stored_value)};
        } else {
            auto updated_ref_count =
                results.reference_count - deleted_references;
            auto update_result = saveValueWithRefCount(
                tx, updated_ref_count, hash_key, results.stored_value);
            return DeleteResults{updated_ref_count, update_result.status,
                                 std::move(results.stored_value)};
        }
    } else {
        return DeleteResults{0, results.status,
                             std::move(results.stored_value)};
    }
}

GetResults getRefCountedData(const ReadTransaction& tx,
                             const rocksdb::Slice& hash_key) {
    std::string return_value;
    auto get_status = tx.refCountedGet(hash_key, &return_value);

    if (!get_status.ok()) {
        auto unsuccessful = rocksdb::Status::NotFound();
        return GetResults{0, unsuccessful, std::vector<unsigned char>()};
    }
    if (return_value.empty()) {
        return GetResults{0, get_status, {}};
    }
    const char* c_string = return_value.c_str();
    uint32_t ref_count;
    memcpy(&ref_count, c_string, sizeof(ref_count));
    std::vector<unsigned char> saved_value(
        return_value.begin() + sizeof(ref_count), return_value.end());
    return GetResults{ref_count, get_status, saved_value};
}
