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
constexpr auto metadata_value_size =
    sizeof(uint32_t) + sizeof(uint64_t) + sizeof(uint64_t);
constexpr auto segment_chunk_key_size =
    segment_key_prefix.size() + sizeof(uint64_t) + sizeof(uint64_t);

std::array<unsigned char, segment_key_size> segment_metadata_key(
    uint64_t segment_id) {
    std::array<unsigned char, segment_key_size> key{};
    auto it = std::copy(segment_key_prefix.begin(), segment_key_prefix.end(),
                        key.begin());
    auto big_id = boost::endian::native_to_big(segment_id);
    auto big_id_ptr = reinterpret_cast<const char*>(&big_id);
    std::copy(big_id_ptr, big_id_ptr + sizeof(big_id), it);
    return key;
}

std::array<unsigned char, segment_chunk_key_size> segment_chunk_key(
    uint64_t segment_id,
    uint64_t chunk) {
    std::array<unsigned char, segment_chunk_key_size> key{};
    auto it = std::copy(segment_key_prefix.begin(), segment_key_prefix.end(),
                        key.begin());

    auto big_id = boost::endian::native_to_big(segment_id);
    auto big_id_ptr = reinterpret_cast<const char*>(&big_id);
    it = std::copy(big_id_ptr, big_id_ptr + sizeof(big_id), it);

    auto big_chunk = boost::endian::native_to_big(chunk);
    auto big_chunk_ptr = reinterpret_cast<const char*>(&big_chunk);
    std::copy(big_chunk_ptr, big_chunk_ptr + sizeof(big_chunk), it);

    return key;
}

struct CodeSegmentMetadata {
    uint32_t ref_count;
    uint64_t op_count;
    uint64_t chunk_count;

    static CodeSegmentMetadata load(const rocksdb::PinnableSlice& slice) {
        auto iter = slice.data();
        uint32_t ref_count;
        memcpy(&ref_count, iter, sizeof(ref_count));
        ref_count = boost::endian::big_to_native(ref_count);
        iter += sizeof(ref_count);
        auto op_count = deserialize_uint64_t(iter);
        auto chunk_count = deserialize_uint64_t(iter);
        return {ref_count, op_count, chunk_count};
    }

    std::array<unsigned char, metadata_value_size> toData() const {
        std::array<unsigned char, metadata_value_size> value{};
        auto it = value.begin();

        auto big_ref_count = boost::endian::native_to_big(ref_count);
        auto big_ref_count_ptr = reinterpret_cast<const char*>(&big_ref_count);
        it = std::copy(big_ref_count_ptr,
                       big_ref_count_ptr + sizeof(big_ref_count), it);

        auto big_op_count = boost::endian::native_to_big(op_count);
        auto big_op_count_ptr = reinterpret_cast<const char*>(&big_op_count);
        it = std::copy(big_op_count_ptr,
                       big_op_count_ptr + sizeof(big_op_count), it);

        auto big_chunk_count = boost::endian::native_to_big(chunk_count);
        auto big_chunk_count_ptr =
            reinterpret_cast<const char*>(&big_chunk_count);
        it = std::copy(big_chunk_count_ptr,
                       big_chunk_count_ptr + sizeof(big_chunk_count), it);
        return value;
    }
};

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

void extractRawCodeSegment(std::vector<RawCodePoint>& cps,
                           const std::vector<unsigned char>& stored_value) {
    auto iter = stored_value.cbegin();
    auto ptr = reinterpret_cast<const char*>(&*iter);
    auto cp_count = deserialize_uint64_t(ptr);
    iter += sizeof(cp_count);
    for (uint64_t i = 0; i < cp_count; i++) {
        cps.push_back(extractRawCodePoint(iter));
    }
}

struct CodeSegmentData {
    CodeSegmentMetadata metadata;
    std::vector<unsigned char> data;
};

CodeSegmentData prepareToSaveCodeSegment(
    ReadWriteTransaction& tx,
    const CodeSegmentSnapshot& snapshot,
    std::map<uint64_t, uint64_t>& segment_counts) {
    uint64_t segment_id = snapshot.segment->segmentID();
    auto key = segment_metadata_key(segment_id);
    CodeSegmentMetadata metadata{};
    rocksdb::PinnableSlice val;
    auto s = tx.refCountedGet(vecToSlice(key), &val);
    if (s.ok()) {
        metadata = CodeSegmentMetadata::load(val);
        val.Reset();
    }

    if (metadata.op_count >= snapshot.op_count) {
        // If this segment is already saved with at least as many ops as is
        // currently contains, just increment the reference count
        return {metadata, {}};
    }

    std::vector<unsigned char> serialized_code;
    marshal_uint64_t(snapshot.op_count - metadata.op_count, serialized_code);
    for (uint64_t i = metadata.op_count; i < snapshot.op_count; ++i) {
        if (i > 1 && i % 10 == 0) {
            auto cp = snapshot.segment->loadCodePoint(i);
            serialized_code.push_back(cp.op.immediate ? 1 : 0);
            serialized_code.push_back(static_cast<unsigned char>(cp.op.opcode));
            marshal_uint256_t(cp.nextHash, serialized_code);
        } else {
            auto& op = snapshot.segment->loadOperation(i);
            serialized_code.push_back(op.immediate ? 1 : 0);
            serialized_code.push_back(static_cast<unsigned char>(op.opcode));
            marshal_uint256_t(0, serialized_code);
        }
        auto& op = snapshot.segment->loadOperation(i);
        if (op.immediate) {
            auto values =
                serializeValue(*op.immediate, serialized_code, segment_counts);
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
    metadata.op_count = snapshot.op_count;
    return CodeSegmentData{metadata, std::move(serialized_code)};
}

std::vector<RawCodePoint> loadRawCodePoints(
    const ReadTransaction& tx,
    uint64_t segment_id,
    const CodeSegmentMetadata& metadata) {
    std::vector<RawCodePoint> raw_cps;
    raw_cps.reserve(metadata.op_count);
    for (uint64_t i = 0; i < metadata.chunk_count; i++) {
        auto key = segment_chunk_key(segment_id, i);
        rocksdb::PinnableSlice val;
        auto s = tx.refCountedGet(vecToSlice(key), &val);
        if (!s.ok()) {
            throw std::runtime_error("failed to load segment chunk");
        }
        std::vector<unsigned char> data{val.data(), val.data() + val.size()};
        val.Reset();
        extractRawCodeSegment(raw_cps, data);
    }
    return raw_cps;
}
}  // namespace

std::shared_ptr<CodeSegment> getCodeSegment(const ReadTransaction& tx,
                                            uint64_t segment_id,
                                            std::set<uint64_t>& segment_ids,
                                            ValueCache& value_cache) {
    auto key_vec = segment_metadata_key(segment_id);
    auto key = vecToSlice(key_vec);
    rocksdb::PinnableSlice val;
    auto s = tx.refCountedGet(key, &val);
    if (!s.ok()) {
        throw std::runtime_error("failed to load segment metadata");
    }
    auto metadata = CodeSegmentMetadata::load(val);
    val.Reset();

    auto raw_cps = loadRawCodePoints(tx, segment_id, metadata);

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
    std::unordered_map<uint64_t, CodeSegmentMetadata> current_values{};

    auto total_deleted_segment_references = breadthFirstSearch(
        segment_counts, [&](uint64_t segment_id, uint64_t total_reference_count,
                            std::map<uint64_t, uint64_t>& next_segment_counts) {
            auto current_value_it = current_values.find(segment_id);
            // Load the segment if it isn't already loaded
            if (current_value_it == current_values.end()) {
                auto key_vec = segment_metadata_key(segment_id);
                rocksdb::PinnableSlice val;
                auto s = tx.refCountedGet(vecToSlice(key_vec), &val);
                if (!s.ok()) {
                    std::cerr
                        << "Couldn't load code segment metadata when deleting"
                        << std::endl;
                    return false;
                }
                auto metadata = CodeSegmentMetadata::load(val);
                val.Reset();
                auto inserted =
                    current_values.insert(std::make_pair(segment_id, metadata));
                current_value_it = inserted.first;
            }

            if (total_reference_count < current_value_it->second.ref_count) {
                // There are still other references to this section, so we won't
                // delete it
                return false;
            }
            auto cps =
                loadRawCodePoints(tx, segment_id, current_value_it->second);
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
        auto key_vec = segment_metadata_key(item.first);
        auto result =
            deleteRefCountedData(tx, vecToSlice(key_vec), item.second);
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

    std::unordered_map<uint64_t, CodeSegmentData> code_segments_to_save{};

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
    for (auto& item : code_segments_to_save) {
        uint64_t chunk = item.second.metadata.chunk_count;
        item.second.metadata.ref_count += total_segment_counts[item.first];

        if (!item.second.data.empty()) {
            item.second.metadata.chunk_count++;
        }

        auto key_vec = segment_metadata_key(item.first);
        auto metadata_raw = item.second.metadata.toData();
        auto s =
            tx.refCountedPut(vecToSlice(key_vec), vecToSlice(metadata_raw));
        if (!s.ok()) {
            return s;
        }

        if (!item.second.data.empty()) {
            auto chunk_key = segment_chunk_key(item.first, chunk);
            s = tx.refCountedPut(vecToSlice(chunk_key),
                                 vecToSlice(item.second.data));
            if (!s.ok()) {
                return s;
            }
        }
    }
    return rocksdb::Status::OK();
}
