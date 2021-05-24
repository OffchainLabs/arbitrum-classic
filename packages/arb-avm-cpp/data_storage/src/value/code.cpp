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

#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/utils.hpp>

#include <avm_values/code.hpp>

#include <boost/endian/conversion.hpp>

#include <utility>
#include <vector>

namespace {

constexpr auto max_code_segment_key = "max_code_segment";
constexpr auto segment_key_prefix = std::array<char, 1>{87};
constexpr auto segment_key_size = segment_key_prefix.size() + sizeof(uint64_t);

std::array<unsigned char, segment_key_size> segment_key(uint64_t segment_id) {
    std::array<unsigned char, segment_key_size> key{};
    auto it = std::copy(segment_key_prefix.begin(), segment_key_prefix.end(),
                        key.begin());
    auto big_id = boost::endian::native_to_big(segment_id);
    auto big_id_ptr = reinterpret_cast<const char*>(&big_id);
    std::copy(big_id_ptr, big_id_ptr + sizeof(big_id), it);
    return key;
}

struct RawCodePoint {
    OpCode opcode;
    std::optional<ParsedSerializedVal> parsed_immediate;
    uint256_t next_hash;
};

RawCodePoint extractRawCodePoint(
    std::vector<unsigned char>::const_iterator& it) {
    bool is_immediate = *it;
    ++it;
    auto opcode = static_cast<OpCode>(*it);
    ++it;
    auto ptr = reinterpret_cast<const char*>(&*it);
    uint256_t next_hash = deserializeUint256t(ptr);
    it += 32;
    if (!is_immediate) {
        return {opcode, std::nullopt, next_hash};
    }
    return {opcode, parseRecord(it), next_hash};
}

std::vector<RawCodePoint> extractRawCodeSegment(
    const std::vector<unsigned char>& stored_value) {
    auto iter = stored_value.cbegin();
    auto ptr = reinterpret_cast<const char*>(&*iter);
    auto cp_count = deserialize_uint64_t(ptr);
    iter += sizeof(cp_count);
    std::vector<RawCodePoint> cps;
    cps.reserve(cp_count);
    for (uint64_t i = 0; i < cp_count; i++) {
        cps.push_back(extractRawCodePoint(iter));
    }
    return cps;
}

void serializeCodePoint(const CodePoint& cp,
                        std::vector<unsigned char>& serialized_code) {
    // Ignore referemces to other code segments
    serialized_code.push_back(cp.op.immediate ? 1 : 0);
    serialized_code.push_back(static_cast<unsigned char>(cp.op.opcode));
    marshal_uint256_t(cp.nextHash, serialized_code);
}

std::vector<unsigned char> prepareToSaveCodeSegment(
    ReadWriteTransaction& tx,
    const CodeSegmentSnapshot& snapshot,
    std::map<uint64_t, uint64_t>& segment_counts) {
    uint64_t segment_id = snapshot.segment->segmentID();
    auto key = segment_key(segment_id);
    uint64_t existing_cp_count = 0;
    rocksdb::PinnableSlice val;
    auto s = tx.refCountedGet(vecToSlice(key), &val);
    std::vector<unsigned char> serialized_code;
    if (s.ok()) {
        auto iter = val.data();
        iter += sizeof(uint32_t);
        existing_cp_count = deserialize_uint64_t(iter);
        if (existing_cp_count >= snapshot.op_count) {
            // If this segment is already saved with at least as many ops as is
            // currently contains, just increment the reference count
            val.Reset();
            return {};
        }
        marshal_uint64_t(snapshot.op_count, serialized_code);
        auto offset = sizeof(uint32_t) + sizeof(uint64_t);
        serialized_code.insert(serialized_code.end(), val.data() + offset,
                               val.data() + val.size());
    } else {
        marshal_uint64_t(snapshot.op_count, serialized_code);
    }
    val.Reset();

    for (uint64_t i = existing_cp_count; i < snapshot.op_count; ++i) {
        auto cp = snapshot.segment->loadCodePoint(i);
        serializeCodePoint(cp, serialized_code);
        if (cp.op.immediate) {
            auto values = serializeValue(*cp.op.immediate, serialized_code,
                                         segment_counts);
            // Save the immediate values, that weren't already saved for this
            // code segment
            for (const auto& val : values) {
                auto result = saveValueImpl(tx, val, segment_counts);
                if (!result.status.ok()) {
                    throw std::runtime_error("failed to save immediate value");
                }
            }
        }
    }
    return serialized_code;
}
}  // namespace

std::shared_ptr<CodeSegment> getCodeSegment(const ReadTransaction& tx,
                                            uint64_t segment_id,
                                            std::set<uint64_t>& segment_ids,
                                            ValueCache& value_cache) {
    auto key_vec = segment_key(segment_id);
    auto key = vecToSlice(key_vec);
    auto results = getRefCountedData(tx, key);

    if (!results.status.ok()) {
        throw std::runtime_error("failed to load segment");
    }

    auto raw_cps = extractRawCodeSegment(results.stored_value);
    std::vector<Operation> ops;
    std::vector<uint256_t> next_hashes;
    ops.reserve(raw_cps.size());
    for (const auto& raw_cp : raw_cps) {
        if (!raw_cp.parsed_immediate) {
            ops.emplace_back(raw_cp.opcode);
        } else {
            auto imm = getValueRecord(tx, *raw_cp.parsed_immediate, segment_ids,
                                      value_cache);
            if (std::holds_alternative<rocksdb::Status>(imm)) {
                throw std::runtime_error("failed to load immediate value");
            }
            ops.emplace_back(raw_cp.opcode,
                             std::get<CountedData<value>>(imm).data);
        }
        if (ops.size() > 1 && ops.size() % 10 == 1) {
            next_hashes.push_back(raw_cp.next_hash);
        }
    }
    return std::make_shared<CodeSegment>(segment_id, std::move(ops),
                                         std::move(next_hashes));
}

void saveNextSegmentID(ReadWriteTransaction& tx, uint64_t next_segment_id) {
    std::vector<unsigned char> value_data;
    marshal_uint64_t(next_segment_id, value_data);
    auto value_slice = vecToSlice(value_data);
    auto status =
        tx.defaultPut(rocksdb::Slice(max_code_segment_key), value_slice);
    if (!status.ok()) {
        throw std::runtime_error("failed to size mac code segment");
    }
}

uint64_t getNextSegmentID(std::shared_ptr<DataStorage> store) {
    ReadTransaction tx(std::move(store));
    std::string segment_id_raw;
    auto s =
        tx.defaultGet(rocksdb::Slice(max_code_segment_key), &segment_id_raw);
    if (s.IsNotFound()) {
        return 0;
    }
    if (!s.ok()) {
        throw std::runtime_error("couldn't load segment id");
    }
    auto ptr = reinterpret_cast<const char*>(segment_id_raw.data());
    return deserialize_uint64_t(ptr);
}

template <typename Func>
std::unordered_map<uint64_t, uint64_t> breadthFirstSearch(
    std::map<uint64_t, uint64_t>& segment_counts,
    Func&& func) {
    std::unordered_map<uint64_t, uint64_t> total_segment_counts{};
    auto current_segment_counts = segment_counts;

    bool found = true;
    while (found) {
        found = false;
        std::map<uint64_t, uint64_t> next_segment_counts;
        for (auto it = current_segment_counts.rbegin();
             it != current_segment_counts.rend(); ++it) {
            uint64_t segment_id = it->first;
            total_segment_counts[segment_id] += it->second;
            uint64_t total_reference_count = total_segment_counts[segment_id];
            if (func(segment_id, total_reference_count, next_segment_counts)) {
                found = true;
            }
        }
        current_segment_counts = std::move(next_segment_counts);
    }
    return total_segment_counts;
}

rocksdb::Status deleteCode(ReadWriteTransaction& tx,
                           std::map<uint64_t, uint64_t>& segment_counts) {
    std::unordered_map<uint64_t, GetResults> current_values{};

    auto total_deleted_segment_references = breadthFirstSearch(
        segment_counts, [&](uint64_t segment_id, uint64_t total_reference_count,
                            std::map<uint64_t, uint64_t>& next_segment_counts) {
            auto current_value_it = current_values.find(segment_id);
            // Load the segment if it isn't already loaded
            if (current_value_it == current_values.end()) {
                auto key_vec = segment_key(segment_id);
                auto key = vecToSlice(key_vec);
                auto inserted = current_values.insert(
                    std::make_pair(segment_id, getRefCountedData(tx, key)));
                current_value_it = inserted.first;
            }

            if (total_reference_count <
                current_value_it->second.reference_count) {
                // There are still other references to this section, so we won't
                // delete it
                return false;
            }
            auto cps =
                extractRawCodeSegment(current_value_it->second.stored_value);
            for (const auto& cp : cps) {
                if (cp.parsed_immediate) {
                    deleteValueRecord(tx, *cp.parsed_immediate,
                                      next_segment_counts);
                }
            }
            return true;
        });

    // Now that we've handled all reference of removed segments, decrement all
    // reference counts
    for (const auto& item : total_deleted_segment_references) {
        auto key_vec = segment_key(item.first);
        auto key = vecToSlice(key_vec);
        auto result = deleteRefCountedData(tx, key, item.second);
        if (!result.status.ok()) {
            return result.status;
        }
    }

    return rocksdb::Status::OK();
}

rocksdb::Status saveCode(ReadWriteTransaction& tx,
                         const Code& code,
                         std::map<uint64_t, uint64_t>& segment_counts) {
    auto snapshots = code.snapshot();
    saveNextSegmentID(tx, snapshots.op_count);

    std::unordered_map<uint64_t, std::vector<unsigned char>>
        code_segments_to_save{};

    auto total_segment_counts = breadthFirstSearch(
        segment_counts, [&](uint64_t segment_id, uint64_t total_reference_count,
                            std::map<uint64_t, uint64_t>& next_segment_counts) {
            if (total_reference_count == 0) {
                // If there are no references, don't bother saving
                return false;
            }
            if (code_segments_to_save.find(segment_id) !=
                code_segments_to_save.end()) {
                // If we've already saved this segment, there's nothing to do
                return false;
            }
            uint64_t current_segment_count = next_segment_counts[segment_id];
            code_segments_to_save[segment_id] = prepareToSaveCodeSegment(
                tx, snapshots.segments[segment_id], next_segment_counts);
            // Ignore internal references to this segment
            next_segment_counts[segment_id] = current_segment_count;
            return true;
        });

    // Now that we've handled all references, save all the serialized segments
    for (const auto& item : code_segments_to_save) {
        auto key_vec = segment_key(item.first);
        SaveResults results;
        if (item.second.empty()) {
            results = incrementReference(tx, vecToSlice(key_vec));
        } else {
            results =
                saveRefCountedDataReplaced(tx, vecToSlice(key_vec), item.second,
                                           total_segment_counts[item.first]);
        }
        if (!results.status.ok()) {
            return results.status;
        }
    }

    return rocksdb::Status::OK();
}
