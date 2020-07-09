/*
 * Copyright 2020, Offchain Labs, Inc.
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

#include <data_storage/value/code.hpp>

#include "referencecount.hpp"
#include "utils.hpp"

#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/value.hpp>

#include <avm_values/code.hpp>

#include <boost/endian/conversion.hpp>

#include <vector>

namespace {

const char* max_code_segment_key = "max_code_segment";
constexpr auto segment_key_prefix = std::array<char, 1>{87};
constexpr auto segment_key_size = segment_key_prefix.size() + sizeof(uint64_t);

std::array<unsigned char, segment_key_size> segment_key(uint64_t segment_id) {
    std::array<unsigned char, segment_key_size> key;
    auto it = std::copy(segment_key_prefix.begin(), segment_key_prefix.end(),
                        key.begin());
    auto big_id = boost::endian::native_to_big(segment_id);
    auto big_id_ptr = reinterpret_cast<const char*>(&big_id);
    std::copy(big_id_ptr, big_id_ptr + sizeof(big_id), it);
    return key;
}
}  // namespace

DeleteResults deleteCodeSegment(
    Transaction& transaction,
    uint64_t segment_id,
    std::unordered_map<uint64_t, uint64_t>& segment_counts) {
    uint64_t deleted_ref_count = segment_counts[segment_id];
    auto key_vec = segment_key(segment_id);
    auto key = vecToSlice(key_vec);
    auto results = getRefCountedData(*transaction.transaction, key);

    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    auto delete_results =
        deleteRefCountedData(*transaction.transaction, key, deleted_ref_count);

    if (delete_results.reference_count < 1) {
        auto iter = results.stored_value.begin();
        auto ptr = reinterpret_cast<const char*>(&*iter);
        auto cp_count = checkpoint::utils::deserialize_uint64(ptr);
        for (uint64_t i = 0; i < cp_count; i++) {
            bool is_immediate = static_cast<bool>(*ptr);
            ptr += 34;
            if (is_immediate) {
                uint256_t value_hash = deserializeUint256t(ptr);
                deleteValueImpl(transaction, value_hash, segment_counts);
            }
        }
    }
    return delete_results;
}

std::vector<unsigned char> prepareToSaveCodeSegment(
    Transaction& transaction,
    const CodeSegmentSnapshot& snapshot,
    std::map<uint64_t, uint64_t>& segment_counts) {
    uint64_t segment_id = snapshot.segment->segmentID();
    uint64_t added_ref_count = segment_counts[segment_id];
    auto key_vec = segment_key(segment_id);
    auto key = vecToSlice(key_vec);

    auto results = getRefCountedData(*transaction.transaction, key);

    auto incr_ref_count = results.status.ok() && results.reference_count > 0;

    uint64_t existing_cp_count = 0;
    if (incr_ref_count) {
        auto iter = results.stored_value.begin();
        auto ptr = reinterpret_cast<const char*>(&*iter);
        existing_cp_count = checkpoint::utils::deserialize_uint64(ptr);
        if (existing_cp_count >= snapshot.op_count) {
            // If this segment is already saved with at least as many ops as is
            // currently contains, just increment the reference count
            return std::move(results.stored_value);
        }
    }

    std::vector<unsigned char> serialized_code;
    marshal_uint64_t(snapshot.op_count, serialized_code);
    for (uint64_t i = 0; i < snapshot.op_count; ++i) {
        const auto& cp = (*snapshot.segment)[i];
        // Ignore referemces to other code segments
        serialized_code.push_back(cp.op.immediate ? 1 : 0);
        serialized_code.push_back(static_cast<unsigned char>(cp.op.opcode));
        marshal_uint256_t(cp.nextHash, serialized_code);
        if (cp.op.immediate) {
            // Only save the immediate value, if it hasn't been already saved
            // for this code segment
            if (i >= existing_cp_count) {
                saveValueImpl(transaction, *cp.op.immediate, segment_counts);
            }
            marshal_uint256_t(hash_value(*cp.op.immediate), serialized_code);
        }
    }
    // Ignore internal references to this segment
    segment_counts[segment_id] = added_ref_count;
    return serialized_code;
}

std::shared_ptr<CodeSegment> getCodeSegment(const Transaction& transaction,
                                            uint64_t segment_id,
                                            TuplePool* pool,
                                            std::set<uint64_t>& segment_ids) {
    auto key_vec = segment_key(segment_id);
    auto key = vecToSlice(key_vec);
    auto results = getRefCountedData(*transaction.transaction, key);

    if (!results.status.ok()) {
        throw std::runtime_error("failed to load segment");
    }

    auto iter = results.stored_value.begin();
    auto ptr = reinterpret_cast<const char*>(&*iter);
    auto cp_count = checkpoint::utils::deserialize_uint64(ptr);
    std::vector<CodePoint> cps;
    cps.reserve(cp_count);
    for (uint64_t i = 0; i < cp_count; i++) {
        bool is_immediate = static_cast<bool>(*ptr);
        ++ptr;
        OpCode opcode = static_cast<OpCode>(*ptr);
        ++ptr;
        uint256_t next_hash = deserializeUint256t(ptr);
        if (is_immediate) {
            uint256_t value_hash = deserializeUint256t(ptr);
            auto imm = getValueImpl(transaction, value_hash, pool, segment_ids);
            if (!imm.status.ok()) {
                throw std::runtime_error("failed to load immediate value");
            }
            cps.push_back({{opcode, imm.data}, next_hash});
        } else {
            cps.push_back({{opcode}, next_hash});
        }
    }
    return std::make_shared<CodeSegment>(segment_id, std::move(cps));
}

void deleteCode(Transaction& transaction,
                std::unordered_map<uint64_t, uint64_t>& segment_counts) {
    std::vector<uint64_t> segment_ids;
    segment_ids.reserve(segment_counts.size());
    for (const auto& item : segment_counts) {
        segment_ids.push_back(item.first);
    }

    // Sort segments in reverse order by segment ID since later segments could
    // reference earlier ones
    std::sort(segment_ids.begin(), segment_ids.end(),
              [](uint64_t first, uint64_t second) { return first > second; });
    for (uint64_t segment_id : segment_ids) {
        if (segment_counts[segment_id] > 0) {
            deleteCodeSegment(transaction, segment_id, segment_counts);
        }
    }
}

uint64_t getNextSegmentID(const Transaction& transaction) {
    std::string segment_id_raw;
    auto s = transaction.transaction->Get(rocksdb::ReadOptions(),
                                          rocksdb::Slice(max_code_segment_key),
                                          &segment_id_raw);
    if (s.IsNotFound()) {
        return 0;
    }
    if (!s.ok()) {
        throw std::runtime_error("couldn't load segment id");
    }
    auto ptr = segment_id_raw.data();
    return checkpoint::utils::deserialize_uint64(ptr);
}

void saveCode(Transaction& transaction,
              const Code& code,
              std::map<uint64_t, uint64_t>& segment_counts) {
    auto snapshots = code.snapshot();

    std::vector<unsigned char> value_data;
    marshal_uint64_t(snapshots.op_count, value_data);
    auto value_slice = vecToSlice(value_data);
    auto status = transaction.transaction->Put(
        rocksdb::Slice(max_code_segment_key), value_slice);
    if (!status.ok()) {
        throw std::runtime_error("failed to size mac code segment");
    }

    std::map<uint64_t, uint64_t> total_segment_counts;
    auto current_segment_counts = segment_counts;
    std::unordered_map<uint64_t, std::vector<unsigned char>>
        code_segments_to_save;
    bool saved_segment = true;

    while (saved_segment) {
        saved_segment = false;
        std::map<uint64_t, uint64_t> next_segment_counts;
        for (auto it = current_segment_counts.rbegin();
             it != current_segment_counts.rend(); ++it) {
            uint64_t segment_id = it->first;
            if (current_segment_counts[segment_id] == 0) {
                // If there are no references, don't bother saving
                continue;
            }
            if (code_segments_to_save.find(segment_id) !=
                code_segments_to_save.end()) {
                // If we've already saved this segment, there's nothing to do
                continue;
            }
            code_segments_to_save[segment_id] = prepareToSaveCodeSegment(
                transaction, snapshots.segments[segment_id],
                next_segment_counts);
            saved_segment = true;
        }
        // Aggregate total reference count for saving
        for (const auto& item : next_segment_counts) {
            total_segment_counts[item.first] += item.second;
        }
        current_segment_counts = std::move(next_segment_counts);
    }

    // Now that we've handled all references, save all the serialized segments
    for (const auto& item : code_segments_to_save) {
        auto key_vec = segment_key(item.first);
        auto key = vecToSlice(key_vec);
        saveRefCountedData(*transaction.transaction, key, item.second,
                           total_segment_counts[item.first], true);
    }
}
