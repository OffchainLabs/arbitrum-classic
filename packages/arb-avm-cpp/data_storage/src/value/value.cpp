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
#include <data_storage/value/deserialize.hpp>
#include <data_storage/value/serialize.hpp>

#include <avm_values/tuple.hpp>
#include <cstdint>
#include <data_storage/readtransaction.hpp>
#include <vector>

constexpr auto next_code_segment_key = "next_code_segment";

ValueResult<uint64_t> getDbNextSegmentId(const ReadTransaction& tx) {
    auto result = tx.defaultGetUint64(next_code_segment_key);
    if (result.status.IsNotFound()) {
        return {rocksdb::Status::OK(), 0};
    }
    return result;
}

rocksdb::Status maybeUpdateNextSegmentId(ReadWriteTransaction& tx,
                                         uint64_t known_segment_id) {
    ValueResult<uint64_t> get_result = getDbNextSegmentId(tx);
    if (!get_result.status.ok() || known_segment_id < get_result.data) {
        return get_result.status;
    }
    std::vector<unsigned char> next_segment_id_bytes;
    marshal_uint64_t(known_segment_id + 1, next_segment_id_bytes);
    auto value_slice = vecToSlice(next_segment_id_bytes);
    return tx.defaultPut(next_code_segment_key, value_slice);
}

DbResult<value> getValue(const ReadTransaction& tx,
                         uint256_t value_hash,
                         ValueCache& value_cache) {
    value result;
    std::vector<Slot> slots(1, {SlotPointer(&result), value_hash});
    bool first = true;
    uint32_t first_reference_count = 0;
    while (!slots.empty()) {
        auto slot = slots.back();
        slots.pop_back();
        if (auto val = value_cache.loadIfExists(slot.hash)) {
            if (auto val_ptr = std::get_if<value*>(&slot.ptr)) {
                **val_ptr = *val;
            } else if (auto buf_ptr = std::get_if<Buffer*>(&slot.ptr)) {
                **buf_ptr = std::get<Buffer>(*val);
            } else if (auto code_segment_ptr =
                           std::get_if<CodeSegment*>(&slot.ptr)) {
                **code_segment_ptr = std::get<CodeSegment>(*val);
            } else {
                throw std::runtime_error("unexpected slot pointer type");
            }
        } else {
            std::vector<unsigned char> hash_key;
            marshal_uint256_t(slot.hash, hash_key);
            auto key = vecToSlice(hash_key);
            auto results = getRefCountedData(tx, key);
            if (results.status.ok()) {
                assert(results.reference_count > 0);
                if (first) {
                    first_reference_count = results.reference_count;
                    first = false;
                }
                auto bytes = results.stored_value.cbegin();
                std::visit(
                    [&](const auto& ptr) {
                        deserializeValue(bytes, ptr, slots);
                    },
                    slot.ptr);
                std::visit(
                    [&](const auto& ptr) {
                        value_cache.maybeSave(*ptr, slot.hash);
                    },
                    slot.ptr);
            } else {
                return results.status;
            }
        }
    }
    if (hash_value(result) != value_hash) {
        throw std::runtime_error("Retrieved wrong hash from database");
    }
    return CountedData<value>{first_reference_count, result};
}

SaveResults saveValue(ReadWriteTransaction& tx, const value& val) {
    bool first = true;
    SaveResults ret{};
    ValueCounter items_to_save;
    items_to_save[val] = 1;
    while (!items_to_save.empty()) {
        auto items_to_save_it = items_to_save.begin();
        auto next_item = std::move(items_to_save_it->first);
        auto new_references = items_to_save_it->second;
        items_to_save.erase(items_to_save_it);
        auto hash = hash_value(next_item);
        std::vector<unsigned char> hash_key;
        marshal_uint256_t(hash, hash_key);
        auto key = vecToSlice(hash_key);
        auto results = getRefCountedData(tx, key);
        SaveResults save_ret;
        auto existing_references = 0;
        if (results.status.ok()) {
            existing_references = results.reference_count > 0;
        }
        auto exists = existing_references > 0;
        uint64_t existing_segment_length = 0;
        if (auto code = std::get_if<CodeSegment>(&next_item)) {
            if (exists) {
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
                existing_segment_length = deserialize_uint64_t(buf);
                exists = existing_segment_length >= code->load().size();
            } else {
                auto status = maybeUpdateNextSegmentId(tx, code->segmentID());
                if (!status.ok()) {
                    return {0, status};
                }
            }
        }
        if (exists) {
            save_ret = incrementReference(tx, key);
        } else {
            std::vector<unsigned char> value_vector{};
            serializeValue(next_item, value_vector);
            if (auto code = std::get_if<CodeSegment>(&next_item)) {
                getCodeSegmentDependencies(*code, items_to_save,
                                           existing_segment_length);
            } else {
                getValueDependencies(next_item, items_to_save);
            }
            save_ret = saveRefCountedData(tx, key, value_vector,
                                          existing_references + new_references);
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

DeleteResults deleteValues(ReadWriteTransaction& tx,
                           ValueCounter items_to_delete) {
    bool first = true;
    DeleteResults ret{};
    while (!items_to_delete.empty()) {
        auto items_to_delete_it = items_to_delete.begin();
        auto next_item = std::move(items_to_delete_it->first);
        auto deleted_references = items_to_delete_it->second;
        items_to_delete.erase(items_to_delete_it);
        auto hash = hash_value(next_item);
        std::vector<unsigned char> hash_key;
        marshal_uint256_t(hash, hash_key);
        auto key = vecToSlice(hash_key);
        auto results = deleteRefCountedData(tx, key, deleted_references);
        if (results.status.ok() && results.reference_count == 0) {
            getValueDependencies(next_item, items_to_delete);
        }
        if (first) {
            ret = results;
            first = false;
        }
    }
    return ret;
}

DeleteResults deleteValue(ReadWriteTransaction& tx,
                          const uint256_t& value_hash) {
    ValueCache cache{1, 0};
    auto value_result = getValue(tx, value_hash, cache);
    if (auto status = std::get_if<rocksdb::Status>(&value_result)) {
        return {0, *status};
    }
    auto val = std::get<CountedData<value>>(value_result).data;
    ValueCounter counter;
    counter[val] = 1;
    return deleteValues(tx, counter);
}
