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

#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/utils.hpp>

#include <avm_values/tuple.hpp>
#include <cstdint>
#include <data_storage/readtransaction.hpp>
#include <vector>

struct ValueBeingParsed {
    value val;
    uint32_t reference_count;
    std::vector<ParsedTupVal> raw_vals;

    ValueBeingParsed(value&& v, uint32_t count)
        : val{std::move(v)}, reference_count{count}, raw_vals{} {}

    void setTupleElement(uint64_t pos, value&& newval) {
        if (std::holds_alternative<Tuple>(val)) {
            std::get<Tuple>(val).unsafe_set_element(pos, std::move(newval));
        }
    }
};

constexpr uint64_t tupleInlineNumerator = 7;
constexpr uint64_t tupleInlineDenominator = 8;
bool shouldInlineTuple(const Tuple& tuple) {
    return (::hash(tuple) % tupleInlineDenominator) < tupleInlineNumerator;
}

namespace {

template <class T>
T parseBuffer(const char* buf, int& len) {
    uint8_t depth = buf[0];
    len++;
    buf += 1;
    if (depth == 0) {
        len += RawBuffer::leaf_size;
        const unsigned char* data = reinterpret_cast<const unsigned char*>(buf);
        RawBuffer::LeafData leaf;
        std::copy(data, data + RawBuffer::leaf_size, leaf.begin());
        return Buffer{leaf};
    }
    auto res = std::vector<uint256_t>();
    for (uint64_t i = 0; i < RawBuffer::children_size; i++) {
        uint256_t hash = deserializeUint256t(buf);
        res.push_back(hash);
        len += 32;
    }
    return ParsedBuffer{depth, res};
}

ParsedTupValVector parseTupleData(const char*& buf, uint8_t count) {
    ParsedTupValVector return_vector{};
    for (uint8_t i = 0; i < count; i++) {
        auto value_type = static_cast<ValueTypes>(*buf);
        ++buf;

        switch (value_type) {
            case BUFFER: {
                int len = 0;
                auto res = parseBuffer<ParsedTupVal>(buf, len);
                return_vector.push_back(res);
                buf += len;
                break;
            }
            case NUM: {
                return_vector.emplace_back(deserializeUint256t(buf));
                break;
            }
            case CODE_POINT_STUB: {
                return_vector.emplace_back(deserializeCodePointStub(buf));
                break;
            }
            case HASH_PRE_IMAGE: {
                return_vector.emplace_back(ValueHash{deserializeUint256t(buf)});
                break;
            }
            default: {
                if (value_type - TUPLE > 8) {
                    throw std::runtime_error(
                        "can't get tuple value with invalid type");
                }
                return_vector.emplace_back(
                    parseTupleData(buf, value_type - TUPLE));
            }
        }
    }
    return return_vector;
}

std::vector<value> serializeValue(const uint256_t& val,
                                  std::vector<unsigned char>& value_vector,
                                  std::map<uint64_t, uint64_t>&) {
    value_vector.push_back(NUM);
    marshal_uint256_t(val, value_vector);
    return {};
}
std::vector<value> serializeValue(
    const CodePointStub& val,
    std::vector<unsigned char>& value_vector,
    std::map<uint64_t, uint64_t>& segment_counts) {
    value_vector.push_back(CODE_POINT_STUB);
    val.marshal(value_vector);
    ++segment_counts[val.pc.segment];
    return {};
}
std::vector<value> serializeValue(
    const Tuple& val,
    std::vector<unsigned char>& value_vector,
    std::map<uint64_t, uint64_t>& segment_counts) {
    std::vector<value> ret{};
    value_vector.push_back(TUPLE + val.tuple_size());
    std::vector<value> to_serialize(val.rbegin(), val.rend());
    while (!to_serialize.empty()) {
        value nested = std::move(to_serialize.back());
        to_serialize.pop_back();
        if (std::holds_alternative<Tuple>(nested)) {
            const auto& nested_tup = std::get<Tuple>(nested);
            if (shouldInlineTuple(nested_tup)) {
                value_vector.push_back(TUPLE + nested_tup.tuple_size());
                to_serialize.insert(to_serialize.end(), nested_tup.rbegin(),
                                    nested_tup.rend());
            } else {
                value_vector.push_back(HASH_PRE_IMAGE);
                marshal_uint256_t(hash(nested_tup), value_vector);
                ret.push_back(nested);
            }
        } else {
            auto res = serializeValue(nested, value_vector, segment_counts);
            for (const auto& re : res) {
                ret.push_back(re);
            }
        }
    }
    return ret;
}
std::vector<value> serializeValue(const std::shared_ptr<HashPreImage>&,
                                  std::vector<unsigned char>&,
                                  std::map<uint64_t, uint64_t>&) {
    throw std::runtime_error("Can't serialize hash preimage in db");
}
std::vector<value> serializeValue(const Buffer& b,
                                  std::vector<unsigned char>& value_vector,
                                  std::map<uint64_t, uint64_t>&) {
    value_vector.push_back(BUFFER);
    std::vector<Buffer> res = b.serialize(value_vector);
    std::vector<value> ret{};
    ret.reserve(res.size());
    for (auto& re : res) {
        ret.emplace_back(Buffer(re));
    }
    return ret;
}

// Returns a list of value hashes to be deleted
void deleteParsedValue(const uint256_t&,
                       std::vector<uint256_t>&,
                       std::map<uint64_t, uint64_t>&) {}
void deleteParsedValue(const Buffer&,
                       std::vector<uint256_t>&,
                       std::map<uint64_t, uint64_t>&) {}
void deleteParsedValue(const ParsedBuffer& parsed,
                       std::vector<uint256_t>& vals_to_delete,
                       std::map<uint64_t, uint64_t>&) {
    for (const auto& val : parsed.nodes) {
        vals_to_delete.push_back(val);
    }
}
void deleteParsedValue(const CodePointStub& cp,
                       std::vector<uint256_t>&,
                       std::map<uint64_t, uint64_t>& segment_counts) {
    segment_counts[cp.pc.segment]++;
}
void deleteParsedValue(std::vector<ParsedTupVal> tup,
                       std::vector<uint256_t>& vals_to_delete,
                       std::map<uint64_t, uint64_t>&) {
    while (!tup.empty()) {
        ParsedTupVal val = std::move(tup.back());
        tup.pop_back();
        // We only need to delete tuples since other values are recorded inline
        if (std::holds_alternative<ValueHash>(val)) {
            vals_to_delete.push_back(std::get<ValueHash>(val).hash);
        } else if (std::holds_alternative<ParsedBuffer>(val)) {
            auto parsed = std::get<ParsedBuffer>(val);
            for (const auto& val2 : parsed.nodes) {
                vals_to_delete.push_back(val2);
            }
        } else if (std::holds_alternative<ParsedTupValVector>(val)) {
            const auto& inner = std::get<ParsedTupValVector>(val);
            tup.insert(tup.end(), inner.begin(), inner.end());
        }
    }
}
}  // namespace

ParsedSerializedVal parseRecord(const char*& buf) {
    auto value_type = static_cast<ValueTypes>(*buf);
    ++buf;

    switch (value_type) {
        case NUM: {
            return deserializeUint256t(buf);
        }
        case CODE_POINT_STUB: {
            return deserializeCodePointStub(buf);
        }
        case HASH_PRE_IMAGE: {
            throw std::runtime_error("HASH_ONLY item");
        }
        case BUFFER: {
            int len = 0;
            auto res = parseBuffer<ParsedSerializedVal>(buf, len);
            buf += len;
            return res;
        }
        default: {
            if (value_type - TUPLE > 8) {
                throw std::runtime_error("can't get value with invalid type");
            }
            return parseTupleData(buf, value_type - TUPLE);
        }
    }
}

std::vector<value> serializeValue(
    const value& val,
    std::vector<unsigned char>& value_vector,
    std::map<uint64_t, uint64_t>& segment_counts) {
    return std::visit(
        [&](const auto& val) {
            return serializeValue(val, value_vector, segment_counts);
        },
        val);
}

GetResults applyValue(value&& val,
                      const uint32_t reference_count,
                      std::vector<ValueBeingParsed>& val_stack) {
    auto& current = val_stack.back();

    // This function is only called when populating an existing tuple
    auto tuple_size = std::get<Tuple>(current.val).tuple_size();
    uint64_t tuple_index = tuple_size - current.raw_vals.size() - 1;
    if (tuple_index >= tuple_size) {
        throw std::runtime_error("Tuple and raw_vals size mismatch");
    }

    //  Add value to Tuple that is currently being populated
    current.setTupleElement(tuple_index, std::move(val));
    return GetResults{reference_count, rocksdb::Status::OK(), {}};
}

GetResults processVal(const ReadTransaction& tx,
                      const ValueHash& val_hash,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>& segment_ids,
                      const uint32_t,
                      ValueCache& val_cache);

GetResults processVal(const ReadTransaction& tx,
                      const ParsedBuffer& val,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>&,
                      const uint32_t reference_count,
                      ValueCache& val_cache);

GetResults getStoredValue(const ReadTransaction& tx,
                          const ValueHash& val_hash) {
    std::array<unsigned char, 32> hash_key;
    marshal_uint256_t(val_hash.hash, hash_key);
    auto key = vecToSlice(hash_key);
    auto results = getRefCountedData(tx, key);
    return results;
}

GetResults processVal(const ReadTransaction&,
                      const uint256_t& val,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>&,
                      const uint32_t reference_count,
                      ValueCache&) {
    return applyValue(val, reference_count, val_stack);
}

GetResults processFirstVal(const ReadTransaction&,
                           const uint256_t& val,
                           std::vector<ValueBeingParsed>& val_stack,
                           std::set<uint64_t>&,
                           const uint32_t reference_count,
                           ValueCache&) {
    // Single number requested
    val_stack.emplace_back(val, reference_count);

    return GetResults{reference_count, rocksdb::Status::OK(), {}};
}

GetResults processVal(const ReadTransaction&,
                      const Buffer& val,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>&,
                      const uint32_t reference_count,
                      ValueCache&) {
    return applyValue(val, reference_count, val_stack);
}

GetResults processFirstVal(const ReadTransaction&,
                           const Buffer& val,
                           std::vector<ValueBeingParsed>& val_stack,
                           std::set<uint64_t>&,
                           const uint32_t reference_count,
                           ValueCache&) {
    // Single number requested
    val_stack.emplace_back(val, reference_count);

    return GetResults{reference_count, rocksdb::Status::OK(), {}};
}

Buffer processBuffer(const ReadTransaction& tx,
                     const ParsedBuffer& val,
                     ValueCache& val_cache) {
    std::vector<Buffer> vec;
    for (const auto& node_hash : val.nodes) {
        if (auto cached_val = val_cache.loadIfExists(node_hash)) {
            vec.push_back(std::get<Buffer>(cached_val.value()));
            continue;
        }
        auto val_hash = ValueHash{node_hash};
        // Value not in cache, so need to load from database
        auto results = getStoredValue(tx, val_hash);
        if (!results.status.ok()) {
            std::cerr << "Error loading buffer record "
                      << static_cast<uint64_t>(val_hash.hash) << std::endl;
            return Buffer();
        }
        auto buf = reinterpret_cast<const char*>(results.stored_value.data());
        auto record = parseRecord(buf);

        if (std::holds_alternative<Buffer>(record)) {
            Buffer buf = std::get<Buffer>(record);
            // Check that it has correct height
            val_cache.maybeSave(buf);
            vec.push_back(buf);
        } else if (std::holds_alternative<ParsedBuffer>(record)) {
            Buffer buf =
                processBuffer(tx, std::get<ParsedBuffer>(record), val_cache);
            val_cache.maybeSave(buf);
            vec.push_back(buf);
        } else {
            std::cerr << "Error loading buffer from record" << std::endl;
            return Buffer();
        }
    }

    return {std::move(vec[0]), std::move(vec[1])};
}

GetResults processVal(const ReadTransaction& tx,
                      const ParsedBuffer& val,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>&,
                      const uint32_t reference_count,
                      ValueCache& val_cache) {
    return applyValue(processBuffer(tx, val, val_cache), reference_count,
                      val_stack);
}

GetResults processFirstVal(const ReadTransaction& tx,
                           const ParsedBuffer& val,
                           std::vector<ValueBeingParsed>& val_stack,
                           std::set<uint64_t>&,
                           const uint32_t reference_count,
                           ValueCache& val_cache) {
    val_stack.emplace_back(processBuffer(tx, val, val_cache), reference_count);
    return GetResults{reference_count, rocksdb::Status::OK(), {}};
}

GetResults processVal(const ReadTransaction&,
                      const CodePointStub& val,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>& segment_ids,
                      const uint32_t reference_count,
                      ValueCache&) {
    segment_ids.insert(val.pc.segment);

    return applyValue(val, reference_count, val_stack);
}

GetResults processFirstVal(const ReadTransaction&,
                           const CodePointStub& val,
                           std::vector<ValueBeingParsed>& val_stack,
                           std::set<uint64_t>& segment_ids,
                           const uint32_t reference_count,
                           ValueCache&) {
    // Single segment requested
    segment_ids.insert(val.pc.segment);

    val_stack.emplace_back(val, reference_count);

    return GetResults{reference_count, rocksdb::Status::OK(), {}};
}

GetResults processVal(const ReadTransaction&,
                      const std::vector<ParsedTupVal>& val,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>&,
                      const uint32_t reference_count,
                      ValueCache&) {
    // Add empty tuple to stack, will be filled in as values are processed
    val_stack.emplace_back(Tuple::createSizedTuple(val.size()),
                           reference_count);

    // Fill new vector with list of elements that will populate tuple
    val_stack.back().raw_vals.insert(val_stack.back().raw_vals.end(),
                                     val.rbegin(), val.rend());

    return GetResults{reference_count, rocksdb::Status::OK(), {}};
}

GetResults processFirstVal(const ReadTransaction& tx,
                           const std::vector<ParsedTupVal>& val,
                           std::vector<ValueBeingParsed>& val_stack,
                           std::set<uint64_t>& segment_ids,
                           const uint32_t reference_count,
                           ValueCache& value_cache) {
    return processVal(tx, val, val_stack, segment_ids, reference_count,
                      value_cache);
}

GetResults processVal(const ReadTransaction& tx,
                      const ValueHash& val_hash,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>& segment_ids,
                      const uint32_t,
                      ValueCache& val_cache) {
    if (auto val = val_cache.loadIfExists(val_hash.hash)) {
        // Use cached value
        return applyValue(std::move(*val), 0, val_stack);
    }

    // Value not in cache, so need to load from database
    auto results = getStoredValue(tx, val_hash);
    if (!results.status.ok()) {
        return results;
    }

    auto buf = reinterpret_cast<const char*>(results.stored_value.data());
    auto record = parseRecord(buf);

    return std::visit(
        [&](const auto& val) {
            return processVal(tx, val, val_stack, segment_ids,
                              results.reference_count, val_cache);
        },
        record);
}

GetResults processFirstVal(const ReadTransaction& tx,
                           const ValueHash& val_hash,
                           std::vector<ValueBeingParsed>& val_stack,
                           std::set<uint64_t>& segment_ids,
                           const uint32_t,
                           ValueCache& val_cache) {
    if (auto val = val_cache.loadIfExists(val_hash.hash)) {
        // Use cached value
        val_stack.emplace_back(std::move(*val), 0);
        return GetResults{0, rocksdb::Status::OK(), {}};
    }

    // Value not in cache, so need to load from database
    auto results = getStoredValue(tx, val_hash);
    if (!results.status.ok()) {
        return results;
    }
    auto buf = reinterpret_cast<const char*>(results.stored_value.data());
    auto record = parseRecord(buf);

    return std::visit(
        [&](const auto& val) {
            return processFirstVal(tx, val, val_stack, segment_ids,
                                   results.reference_count, val_cache);
        },
        record);
}

DbResult<value> getValueInternal(const ReadTransaction& tx,
                                 std::vector<ValueBeingParsed> val_stack,
                                 std::set<uint64_t>& segment_ids,
                                 ValueCache& value_cache) {
    if (val_stack[0].raw_vals.empty()) {
        // First value has no child values, so just return single value without
        // populating cache
        return CountedData<value>{val_stack[0].reference_count,
                                  std::move(val_stack[0].val)};
    }

    // This should always be true
    while (!val_stack.empty()) {
        auto& current = val_stack.back();
        if (!current.raw_vals.empty()) {
            // Take next value to process
            auto next = std::move(current.raw_vals.back());
            current.raw_vals.pop_back();

            auto results = std::visit(
                [&](const auto& val) {
                    return processVal(tx, val, val_stack, segment_ids, 0,
                                      value_cache);
                },
                next);
            if (!results.status.ok()) {
                return results.status;
            }
        } else {
            // All child values have been resolved
            auto val = std::move(current.val);
            auto reference_count = current.reference_count;
            val_stack.pop_back();

            if (val_stack.empty()) {
                // All values resolved
                value_cache.maybeSave(val);
                return CountedData<value>{reference_count, std::move(val)};
            }

            if (reference_count > 1) {
                value_cache.maybeSave(val);
            }

            applyValue(std::move(val), reference_count, val_stack);
        }
    }

    throw std::runtime_error("val_stack loop should never finish");
}

DbResult<value> getValueRecord(const ReadTransaction& tx,
                               const ParsedSerializedVal& record,
                               std::set<uint64_t>& segment_ids,
                               ValueCache& value_cache) {
    std::vector<ValueBeingParsed> val_stack{};
    std::visit(
        [&](const auto& val) {
            return processFirstVal(tx, val, val_stack, segment_ids, 1,
                                   value_cache);
        },
        record);
    return getValueInternal(tx, std::move(val_stack), segment_ids, value_cache);
}

DbResult<value> getValueImpl(const ReadTransaction& tx,
                             const uint256_t value_hash,
                             std::set<uint64_t>& segment_ids,
                             ValueCache& value_cache) {
    std::vector<ValueBeingParsed> val_stack{};

    // Initialize val_stack with first value from database
    auto result = processFirstVal(tx, ValueHash{value_hash}, val_stack,
                                  segment_ids, 0, value_cache);
    if (!result.status.ok()) {
        return result.status;
    }
    auto res =
        getValueInternal(tx, std::move(val_stack), segment_ids, value_cache);
    if (std::holds_alternative<rocksdb::Status>(res)) {
        return res;
    }
    auto res_hash = hash_value(std::get<CountedData<value>>(res).data);
    assert(res_hash == value_hash);
    if (res_hash != value_hash) {
        throw std::runtime_error("deserialized with incorrect hash");
    }
    return res;
}

DbResult<value> getValue(const ReadTransaction& tx,
                         const uint256_t value_hash,
                         ValueCache& value_cache) {
    std::set<uint64_t> segment_ids{};
    return getValueImpl(tx, value_hash, segment_ids, value_cache);
}

SaveResults saveValueImpl(ReadWriteTransaction& tx,
                          const value& val,
                          std::map<uint64_t, uint64_t>& segment_counts) {
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
        SaveResults save_ret = incrementReference(tx, key);
        if (save_ret.status.IsNotFound()) {
            std::vector<unsigned char> value_vector{};
            auto new_items_to_save =
                serializeValue(next_item, value_vector, segment_counts);
            items_to_save.insert(items_to_save.end(), new_items_to_save.begin(),
                                 new_items_to_save.end());
            save_ret = saveValueWithRefCount(tx, 1, key, value_vector);
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

SaveResults saveValue(ReadWriteTransaction& tx, const value& val) {
    std::map<uint64_t, uint64_t> segment_counts{};
    return saveValueImpl(tx, val, segment_counts);
}

DeleteResults deleteValues(ReadWriteTransaction& tx,
                           std::vector<uint256_t> items_to_delete,
                           std::map<uint64_t, uint64_t>& segment_counts) {
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
            auto buf =
                reinterpret_cast<const char*>(results.stored_value.data());
            std::visit(
                [&](const auto& val) {
                    deleteParsedValue(val, items_to_delete, segment_counts);
                },
                parseRecord(buf));
        }
        if (first) {
            ret = results;
            first = false;
        }
    }
    return ret;
}

DeleteResults deleteValueImpl(ReadWriteTransaction& tx,
                              const uint256_t& value_hash,
                              std::map<uint64_t, uint64_t>& segment_counts) {
    std::vector<uint256_t> items_to_delete{value_hash};
    return deleteValues(tx, std::move(items_to_delete), segment_counts);
}

DeleteResults deleteValueRecord(ReadWriteTransaction& tx,
                                const ParsedSerializedVal& val,
                                std::map<uint64_t, uint64_t>& segment_counts) {
    std::vector<uint256_t> items_to_delete{};
    std::visit(
        [&](const auto& val) {
            deleteParsedValue(val, items_to_delete, segment_counts);
        },
        val);
    return deleteValues(tx, std::move(items_to_delete), segment_counts);
}

DeleteResults deleteValue(ReadWriteTransaction& tx, uint256_t value_hash) {
    std::map<uint64_t, uint64_t> segment_counts{};
    return deleteValueImpl(tx, value_hash, segment_counts);
}
