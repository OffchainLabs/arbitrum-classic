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

#include <data_storage/value/value.hpp>
#include <utility>

#include "referencecount.hpp"
#include "utils.hpp"

#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>

#include <avm_values/tuple.hpp>
#include <cstdint>
#include <data_storage/readtransaction.hpp>
#include <vector>

DbResult<value> getValue(const ReadTransaction& tx,
                         const uint256_t value_hash,
                         ValueCache& value_cache) {
    throw std::runtime_error("TODO: unimplemented");
}

SaveResults saveValue(ReadWriteTransaction& tx, const value& val) {
    bool first = true;
    SaveResults ret{};
    std::vector<value> items_to_save{val};
    while (!items_to_save.empty()) {
        auto next_item = std::move(items_to_save.back());
        items_to_save.pop_back();
        auto hash = hash_value(next_item);
        std::vector<unsigned char> hash_key;
        marshal_uint256_t(hash, hash_key);
        auto key = vecToSlice(hash_key);
        auto results = getRefCountedData(tx, key);
        SaveResults save_ret;
        auto existing_references = 0;
        if (results.status.ok()) {
            existing_references = results.references_count > 0;
        }
        auto exists = existing_references > 0;
        if (auto code = std::get_if<CodeSegment>(&next_item) && exists) {
            const char* buf =
                reinterpret_cast<const char*>(results.stored_value.data());
            if (*buf++ != CODE_SEGMENT) {
                throw new std::runtime_error(
                    "DB corruption: non-code-segment found in code segment "
                    "key");
            }
            auto segment_id = deserialize_uint64_t(buf);
            if (segment_id != code->segmentID()) {
                throw new std::runtime_error(
                    "DB corruption: code segment ID didn't match key");
            }
            auto existing_len = deserialize_uint64_t(buf);
            exists = existing_len >= code->load().size();
        }
        if (exists) {
            save_ret = incrementReference(tx, key);
        } else {
            std::vector<unsigned char> value_vector{};
            auto new_items_to_save = serializeValue(next_item, value_vector);
            items_to_save.insert(items_to_save.end(), new_items_to_save.begin(),
                                 new_items_to_save.end());
            save_ret = saveRefCountedData(tx, key, value_vector,
                                          existing_reference_count + 1);
        }
        if (!save_ret.status.ok()) {
            return save_ret;
        }
        if (first) {
            ret = save_ret;
            first = false;
        }
    }
    return ret;
}

// TODO should operate on deserialized values
DeleteResults deleteValues(ReadWriteTransaction& tx,
                           std::vector<uint256_t> items_to_delete) {
    bool first = true;
    DeleteResults ret{};
    while (!items_to_delete.empty()) {
        auto next_item = items_to_delete.back();
        items_to_delete.pop_back();
        std::vector<unsigned char> hash_key;
        marshal_uint256_t(next_item, hash_key);
        auto key = vecToSlice(hash_key);
        auto results = deleteRefCountedData(tx, key);
        if (results.status.ok() && results.reference_count == 0) {
            auto it = results.stored_value.cbegin();
            throw std::runtime_error("TODO: deleting unimplemented");
        }
        if (first) {
            ret = results;
            first = false;
        }
    }
    return ret;
}

DeleteResults deleteValue(ReadWriteTransaction& tx, uint256_t value_hash) {
    std::vector<uint256_t> items_to_delete{value_hash};
    return deleteValues(tx, std::move(items_to_delete));
}
