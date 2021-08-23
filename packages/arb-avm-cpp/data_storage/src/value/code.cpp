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
    sizeof(uint32_t) + sizeof(uint64_t) + sizeof(uint64_t) + sizeof(uint64_t);
constexpr auto segment_chunk_key_size =
    segment_key_prefix.size() + sizeof(uint64_t) + sizeof(uint64_t) + 1;

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

std::array<unsigned char, segment_chunk_key_size> segment_op_chunk_key_raw(
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
    key.back() = 0;
    return key;
}

std::array<unsigned char, segment_chunk_key_size> segment_op_chunk_key(
    uint64_t segment_id,
    uint64_t chunk) {
    auto key = segment_op_chunk_key_raw(segment_id, chunk);
    key.back() = 0;
    return key;
}

std::array<unsigned char, segment_chunk_key_size> segment_hash_chunk_key(
    uint64_t segment_id,
    uint64_t chunk) {
    auto key = segment_op_chunk_key_raw(segment_id, chunk);
    key.back() = 1;
    return key;
}

struct CodeSegmentMetadata {
    uint32_t ref_count;
    uint64_t op_count;
    uint64_t hash_count;
    uint64_t chunk_count;

    static CodeSegmentMetadata load(const rocksdb::PinnableSlice& slice) {
        auto iter = slice.data();
        uint32_t ref_count;
        memcpy(&ref_count, iter, sizeof(ref_count));
        ref_count = boost::endian::big_to_native(ref_count);
        iter += sizeof(ref_count);
        auto op_count = deserialize_uint64_t(iter);
        auto hash_count = deserialize_uint64_t(iter);
        auto chunk_count = deserialize_uint64_t(iter);
        return {ref_count, op_count, hash_count, chunk_count};
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

        auto big_hash_count = boost::endian::native_to_big(hash_count);
        auto big_hash_count_ptr =
            reinterpret_cast<const char*>(&big_hash_count);
        it = std::copy(big_hash_count_ptr,
                       big_hash_count_ptr + sizeof(big_hash_count), it);

        auto big_chunk_count = boost::endian::native_to_big(chunk_count);
        auto big_chunk_count_ptr =
            reinterpret_cast<const char*>(&big_chunk_count);
        it = std::copy(big_chunk_count_ptr,
                       big_chunk_count_ptr + sizeof(big_chunk_count), it);
        return value;
    }
};

struct RawOperation {
    OpCode opcode;
    std::optional<ParsedSerializedVal> parsed_immediate;
};

RawOperation extractRawOperation(const char*& buf) {
    bool is_immediate = *buf;
    ++buf;
    auto opcode = static_cast<OpCode>(*buf);
    ++buf;
    if (!is_immediate) {
        return {opcode, std::nullopt};
    }
    return {opcode, parseRecord(buf)};
}

void extractRawOperations(std::vector<RawOperation>& cps, const char*& buf) {
    auto op_count = deserialize_uint64_t(buf);
    for (uint64_t i = 0; i < op_count; i++) {
        cps.push_back(extractRawOperation(buf));
    }
}

void extractRawHashes(std::vector<uint256_t>& hashes, const char*& ptr) {
    auto hash_count = deserialize_uint64_t(ptr);
    for (uint64_t i = 0; i < hash_count; i++) {
        hashes.push_back(deserializeUint256t(ptr));
    }
}

struct RawCodeSegmentData {
    CodeSegmentMetadata metadata;
    std::vector<unsigned char> op_data;
    std::vector<unsigned char> hash_data;
};

RawCodeSegmentData prepareToSaveCodeSegment(
    ReadWriteTransaction& tx,
    const CodeSegmentSnapshot& snapshot,
    std::map<uint64_t, uint64_t>& segment_counts) {
    uint64_t segment_id = snapshot.segmentID();
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
        return {metadata, {}, {}};
    }

    std::vector<unsigned char> op_data;
    marshal_uint64_t(snapshot.op_count - metadata.op_count, op_data);
    for (uint64_t i = metadata.op_count; i < snapshot.op_count; ++i) {
        auto& op = snapshot.loadOperation(i);
        op_data.push_back(op.immediate ? 1 : 0);
        op_data.push_back(static_cast<unsigned char>(op.opcode));
        if (op.immediate) {
            auto values = serializeValue(tx.getSecretHashSeed(), *op.immediate,
                                         op_data, segment_counts);
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
    std::vector<unsigned char> hash_data;
    marshal_uint64_t(snapshot.cached_hash_count - metadata.hash_count,
                     hash_data);
    for (uint64_t i = metadata.hash_count; i < snapshot.cached_hash_count;
         i++) {
        marshal_uint256_t(snapshot.loadCachedHash(i), hash_data);
    }
    metadata.op_count = snapshot.op_count;
    metadata.hash_count = snapshot.cached_hash_count;
    return {metadata, std::move(op_data), std::move(hash_data)};
}

std::vector<RawOperation> loadRawOperations(
    const ReadTransaction& tx,
    uint64_t segment_id,
    const CodeSegmentMetadata& metadata) {
    std::vector<RawOperation> raw_ops;
    raw_ops.reserve(metadata.op_count);
    for (uint64_t i = 0; i < metadata.chunk_count; i++) {
        auto key = segment_op_chunk_key(segment_id, i);
        rocksdb::PinnableSlice val;
        auto s = tx.refCountedGet(vecToSlice(key), &val);
        if (!s.ok()) {
            throw std::runtime_error("failed to load segment op chunk");
        }
        auto buf = val.data();
        extractRawOperations(raw_ops, buf);
        val.Reset();
    }
    return raw_ops;
}

std::vector<uint256_t> loadHashes(const ReadTransaction& tx,
                                  uint64_t segment_id,
                                  const CodeSegmentMetadata& metadata) {
    std::vector<uint256_t> hashes;
    hashes.reserve(metadata.chunk_count / 10 + 1);
    for (uint64_t i = 0; i < metadata.chunk_count; i++) {
        auto key = segment_hash_chunk_key(segment_id, i);
        rocksdb::PinnableSlice val;
        auto s = tx.refCountedGet(vecToSlice(key), &val);
        if (!s.ok()) {
            throw std::runtime_error("failed to load segment hash chunk");
        }
        auto hashes_data = val.data();
        extractRawHashes(hashes, hashes_data);
        val.Reset();
    }
    return hashes;
}
}  // namespace

void restoreCodeSegments(const ReadTransaction& transaction,
                         const std::shared_ptr<CoreCode>& core_code,
                         ValueCache& value_cache,
                         std::set<uint64_t> segment_ids) {
    bool loaded_segment = true;
    while (loaded_segment) {
        loaded_segment = false;
        std::set<uint64_t> next_segment_ids;
        for (auto it = segment_ids.rbegin(); it != segment_ids.rend(); ++it) {
            if (core_code->containsSegment(*it)) {
                // If the segment is already loaded, no need to restore it
                continue;
            }
            auto segment = getCodeSegment(transaction, *it, next_segment_ids,
                                          value_cache, ENABLE_LAZY_LOADING);
            core_code->restoreExistingSegment(std::move(segment));
            loaded_segment = true;
        }
        segment_ids = std::move(next_segment_ids);
    }
}

std::shared_ptr<UnsafeCodeSegment> getCodeSegment(
    const ReadTransaction& tx,
    uint64_t segment_id,
    std::set<uint64_t>& segment_ids,
    ValueCache& value_cache,
    bool lazy_load) {
    auto key_vec = segment_metadata_key(segment_id);
    rocksdb::PinnableSlice val;
    auto s = tx.refCountedGet(vecToSlice(key_vec), &val);
    if (!s.ok()) {
        throw std::runtime_error("failed to load segment metadata");
    }
    auto metadata = CodeSegmentMetadata::load(val);
    val.Reset();

    auto raw_ops = loadRawOperations(tx, segment_id, metadata);
    std::vector<Operation> ops;
    ops.reserve(raw_ops.size());
    for (const auto& raw_op : raw_ops) {
        if (!raw_op.parsed_immediate) {
            ops.emplace_back(raw_op.opcode);
        } else {
            auto imm = getValueRecord(tx, *raw_op.parsed_immediate, segment_ids,
                                      value_cache, lazy_load);
            if (std::holds_alternative<rocksdb::Status>(imm)) {
                throw std::runtime_error("failed to load immediate value");
            }
            ops.emplace_back(raw_op.opcode,
                             std::get<CountedData<value>>(imm).data);
        }
    }

    auto next_hashes = loadHashes(tx, segment_id, metadata);
    return std::make_shared<UnsafeCodeSegment>(
        segment_id, CodeSegmentData{std::move(ops), std::move(next_hashes)});
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

template <typename Map = std::unordered_map<uint64_t, uint64_t>, typename Func>
Map breadthFirstSearch(std::map<uint64_t, uint64_t>& segment_counts,
                       Func&& func) {
    Map total_segment_counts{};
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
            auto ops =
                loadRawOperations(tx, segment_id, current_value_it->second);
            for (const auto& op : ops) {
                if (op.parsed_immediate) {
                    deleteValueRecord(tx, *op.parsed_immediate,
                                      next_segment_counts);
                }
            }
            return true;
        });

    // Now that we've handled all reference of removed segments, decrement all
    // reference counts
    for (const auto& item : total_deleted_segment_references) {
        auto& metadata = current_values[item.first];
        auto key_vec = segment_metadata_key(item.first);

        if (item.second >= metadata.ref_count) {
            // No more references so delete the segment
            auto s = tx.refCountedDelete(vecToSlice(key_vec));
            if (!s.ok()) {
                return s;
            }
            for (uint64_t i = 0; i < metadata.chunk_count; i++) {
                auto chunk_op_key = segment_op_chunk_key(item.first, i);
                auto s = tx.refCountedDelete(vecToSlice(chunk_op_key));
                if (!s.ok()) {
                    return s;
                }

                auto chunk_hash_key = segment_hash_chunk_key(item.first, i);
                s = tx.refCountedDelete(vecToSlice(chunk_hash_key));
                if (!s.ok()) {
                    return s;
                }
            }
        } else {
            metadata.ref_count -= item.second;
            auto metadata_raw = metadata.toData();
            auto s =
                tx.refCountedPut(vecToSlice(key_vec), vecToSlice(metadata_raw));
            if (!s.ok()) {
                return s;
            }
        }
    }
    return rocksdb::Status::OK();
}

rocksdb::Status saveCode(ReadWriteTransaction& tx,
                         const Code& code,
                         std::map<uint64_t, uint64_t>& segment_counts) {
    auto snapshots = code.snapshot();
    saveNextSegmentID(tx, snapshots.next_segment_num);

    std::unordered_map<uint64_t, RawCodeSegmentData> code_segments_to_save{};

    auto total_segment_counts =
        breadthFirstSearch<std::map<uint64_t, uint64_t>>(
            segment_counts,
            [&](uint64_t segment_id, uint64_t total_reference_count,
                std::map<uint64_t, uint64_t>& next_segment_counts) {
                if (total_reference_count == 0) {
                    // If there are no references, don't bother saving
                    return false;
                }
                if (code_segments_to_save.find(segment_id) !=
                    code_segments_to_save.end()) {
                    // If we've already saved this segment, there's nothing to
                    // do
                    return false;
                }
                uint64_t current_segment_count =
                    next_segment_counts[segment_id];
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

        if (!item.second.op_data.empty()) {
            item.second.metadata.chunk_count++;
        }

        auto key_vec = segment_metadata_key(item.first);
        auto metadata_raw = item.second.metadata.toData();
        auto s =
            tx.refCountedPut(vecToSlice(key_vec), vecToSlice(metadata_raw));
        if (!s.ok()) {
            return s;
        }

        if (!item.second.op_data.empty()) {
            auto chunk_op_key = segment_op_chunk_key(item.first, chunk);
            s = tx.refCountedPut(vecToSlice(chunk_op_key),
                                 vecToSlice(item.second.op_data));
            if (!s.ok()) {
                return s;
            }

            auto chunk_hash_key = segment_hash_chunk_key(item.first, chunk);
            s = tx.refCountedPut(vecToSlice(chunk_hash_key),
                                 vecToSlice(item.second.hash_data));
            if (!s.ok()) {
                return s;
            }
        }
    }
    segment_counts = std::move(total_segment_counts);
    return rocksdb::Status::OK();
}
