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

constexpr uint64_t tupleInlineNumerator = 0;
constexpr uint64_t tupleInlineDenominator = 100;
bool shouldInlineValue(const value& val,
                       const std::vector<unsigned char>& secret_hash_seed) {
    if (auto tuple = std::get_if<Tuple>(&val)) {
        auto hash = tuple->getHashPreImage().secretHash(secret_hash_seed);
        return tuple->tuple_size() == 0 ||
               (hash % tupleInlineDenominator) < tupleInlineNumerator;
    } else if (std::holds_alternative<CodePointStub>(val) ||
               std::holds_alternative<UnloadedValue>(val)) {
        return false;
    } else {
        return true;
    }
}

namespace {

template <class T>
T parseBuffer(const char* buf, int& len) {
    uint8_t depth = buf[0];
    len++;
    buf += 1;
    if (depth == 0) {
        len += RawBuffer::leaf_size;
        const auto* data = reinterpret_cast<const unsigned char*>(buf);
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
    // Begin by attempting to fill a single tuple of size `count`
    std::vector<std::pair<ParsedTupValVector, uint8_t>> tuple_stack(
        1, std::make_pair(ParsedTupValVector(), count));
    while (true) {
        // Collapse full tuples into their parents (or if the root, return it)
        while (tuple_stack.back().first.size() >= tuple_stack.back().second) {
            ParsedTupValVector created_tup =
                std::move(tuple_stack.back().first);
            tuple_stack.pop_back();
            if (tuple_stack.empty()) {
                return created_tup;
            }
            tuple_stack.back().first.push_back(created_tup);
        }

        auto value_type = static_cast<ValueTypes>(*buf);
        ++buf;
        ParsedTupVal val;

        switch (value_type) {
            case BUFFER: {
                int len = 0;
                auto res = parseBuffer<ParsedTupVal>(buf, len);
                val = res;
                buf += len;
                break;
            }
            case NUM: {
                val = deserializeUint256t(buf);
                break;
            }
            case CODE_POINT_STUB: {
                val = deserializeCodePointStub(buf);
                break;
            }
            case HASH_PRE_IMAGE: {
                auto ty = static_cast<ValueTypes>(*buf);
                ++buf;
                auto hash = deserializeUint256t(buf);
                auto size = deserializeUint256t(buf);
                val = BigUnloadedValue{ty, hash, size};
                break;
            }
            default: {
                auto inner_count = value_type - TUPLE;
                if (inner_count > 8) {
                    throw std::runtime_error(
                        "can't get tuple value with invalid type");
                }
                // Before continuing with the parent, fill in this tuple first
                tuple_stack.emplace_back(ParsedTupValVector(), inner_count);
                // Don't attempt to put a value in this tuple yet
                continue;
            }
        }
        // Continue filling in tuple by adding this value
        // (if it was full, it'd have been collapsed earlier)
        tuple_stack.back().first.push_back(val);
    }
}

std::vector<value> serializeValue(const std::vector<unsigned char>&,
                                  const uint256_t& val,
                                  std::vector<unsigned char>& value_vector,
                                  std::map<uint64_t, uint64_t>&) {
    value_vector.push_back(NUM);
    marshal_uint256_t(val, value_vector);
    return {};
}
std::vector<value> serializeValue(
    const std::vector<unsigned char>&,
    const CodePointStub& val,
    std::vector<unsigned char>& value_vector,
    std::map<uint64_t, uint64_t>& segment_counts) {
    value_vector.push_back(CODE_POINT_STUB);
    val.marshal(value_vector);
    ++segment_counts[val.pc.segment];
    return {};
}
std::vector<value> serializeValue(
    const std::vector<unsigned char>& secret_hash_seed,
    const Tuple& root,
    std::vector<unsigned char>& value_vector,
    std::map<uint64_t, uint64_t>& segment_counts) {
    std::vector<value> ret{};
    value_vector.push_back(TUPLE + root.tuple_size());
    // `to_serialize` is a stack, so we populate it in reverse order
    std::vector<value> to_serialize(root.rbegin(), root.rend());
    while (!to_serialize.empty()) {
        // Pull from the end of `to_serialize` as its contents are reversed
        value nested = std::move(to_serialize.back());
        to_serialize.pop_back();
        if (shouldInlineValue(nested, secret_hash_seed)) {
            if (const auto& nested_tup = std::get_if<Tuple>(&nested)) {
                // Write the tuple header, then add the tuple contents
                // to the stack of values to serialize (again in reverse)
                value_vector.push_back(TUPLE + nested_tup->tuple_size());
                to_serialize.insert(to_serialize.end(), nested_tup->rbegin(),
                                    nested_tup->rend());
            } else {
                auto new_ret = serializeValue(secret_hash_seed, nested,
                                              value_vector, segment_counts);
                ret.insert(ret.end(), new_ret.begin(), new_ret.end());
            }
        } else {
            // Serialize reference to value
            value_vector.push_back(HASH_PRE_IMAGE);
            value_vector.push_back(
                static_cast<uint8_t>(std::visit(ValueTypeVisitor{}, nested)));
            marshal_uint256_t(hash_value(nested), value_vector);
            marshal_uint256_t(::getSize(nested), value_vector);
            // Mark value for separate saving
            ret.push_back(nested);
        }
    }
    return ret;
}
std::vector<value> serializeValue(const std::vector<unsigned char>&,
                                  const std::shared_ptr<HashPreImage>&,
                                  std::vector<unsigned char>&,
                                  std::map<uint64_t, uint64_t>&) {
    throw std::runtime_error("Can't serialize hash preimage in db");
}
std::vector<value> serializeValue(const std::vector<unsigned char>&,
                                  const UnloadedValue&,
                                  std::vector<unsigned char>&,
                                  std::map<uint64_t, uint64_t>&) {
    throw std::runtime_error("Can't serialize unloaded value in db");
}
std::vector<value> serializeValue(const std::vector<unsigned char>&,
                                  const Buffer& b,
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
        // We only need to delete tuples and buffers as all other values
        // are "primitives"/"trivial" types in that they're recorded inline
        // and don't reference anything else
        if (std::holds_alternative<BigUnloadedValue>(val)) {
            // Delete the referenced hash in a later pass
            vals_to_delete.push_back(std::get<BigUnloadedValue>(val).hash);
        } else if (std::holds_alternative<ParsedBuffer>(val)) {
            // Delete any buffer nodes in a later pass
            auto parsed = std::get<ParsedBuffer>(val);
            for (const auto& val2 : parsed.nodes) {
                vals_to_delete.push_back(val2);
            }
        } else if (std::holds_alternative<ParsedTupValVector>(val)) {
            // Descend into the tuple by merging in its contents into ours
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
    const std::vector<unsigned char>& secret_hash_seed,
    const value& val,
    std::vector<unsigned char>& value_vector,
    std::map<uint64_t, uint64_t>& segment_counts) {
    return std::visit(
        [&](const auto& val) {
            return serializeValue(secret_hash_seed, val, value_vector,
                                  segment_counts);
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
                      const BigUnloadedValue& val_info,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>& segment_ids,
                      uint32_t,
                      ValueCache& val_cache,
                      bool lazy_load);

GetResults processVal(const ReadTransaction& tx,
                      const ParsedBuffer& val,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>&,
                      uint32_t reference_count,
                      ValueCache& val_cache,
                      bool lazy_load);

GetResults getStoredValue(const ReadTransaction& tx,
                          const BigUnloadedValue& val_info) {
    std::array<unsigned char, 32> hash_key{};
    marshal_uint256_t(val_info.hash, hash_key);
    auto key = vecToSlice(hash_key);
    auto results = getRefCountedData(tx, key);
    return results;
}

GetResults processVal(const ReadTransaction&,
                      const uint256_t& val,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>&,
                      const uint32_t reference_count,
                      ValueCache&,
                      bool) {
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
                      ValueCache&,
                      bool) {
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
        auto unloaded_val = BigUnloadedValue{BUFFER, node_hash, 1};
        // Value not in cache, so need to load from database
        auto results = getStoredValue(tx, unloaded_val);
        if (!results.status.ok()) {
            std::cerr << "Error loading buffer record "
                      << static_cast<uint64_t>(node_hash) << std::endl;
            return {};
        }
        auto record_string =
            reinterpret_cast<const char*>(results.stored_value.data());
        auto record = parseRecord(record_string);

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
            return {};
        }
    }

    return {std::move(vec[0]), std::move(vec[1])};
}

GetResults processVal(const ReadTransaction& tx,
                      const ParsedBuffer& val,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>&,
                      const uint32_t reference_count,
                      ValueCache& val_cache,
                      bool) {
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
                      ValueCache&,
                      bool) {
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
                      ValueCache&,
                      bool) {
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
                      value_cache, false);
}

GetResults processVal(const ReadTransaction& tx,
                      const BigUnloadedValue& val_info,
                      std::vector<ValueBeingParsed>& val_stack,
                      std::set<uint64_t>& segment_ids,
                      const uint32_t,
                      ValueCache& val_cache,
                      bool lazy_load) {
    if (auto val = val_cache.loadIfExists(val_info.hash)) {
        // Use cached value
        return applyValue(std::move(*val), 0, val_stack);
    }

    if (lazy_load && val_info.type == TUPLE) {
        // Don't load value immediately; load on demand
        return applyValue(UnloadedValue(val_info), 0, val_stack);
    }

    // Value not in cache, so need to load from database
    auto results = getStoredValue(tx, val_info);
    if (!results.status.ok()) {
        return results;
    }

    auto buf = reinterpret_cast<const char*>(results.stored_value.data());
    auto record = parseRecord(buf);

    return std::visit(
        [&](const auto& val) {
            return processVal(tx, val, val_stack, segment_ids,
                              results.reference_count, val_cache, lazy_load);
        },
        record);
}

GetResults processFirstVal(const ReadTransaction& tx,
                           const BigUnloadedValue& val_info,
                           std::vector<ValueBeingParsed>& val_stack,
                           std::set<uint64_t>& segment_ids,
                           const uint32_t,
                           ValueCache& val_cache) {
    if (auto val = val_cache.loadIfExists(val_info.hash)) {
        // Use cached value
        val_stack.emplace_back(std::move(*val), 0);
        return GetResults{0, rocksdb::Status::OK(), {}};
    }

    // Value not in cache, so need to load from database
    auto results = getStoredValue(tx, val_info);
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
                                 ValueCache& value_cache,
                                 bool lazy_load) {
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
                                      value_cache, lazy_load);
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
                               ValueCache& value_cache,
                               bool lazy_load) {
    std::vector<ValueBeingParsed> val_stack{};
    std::visit(
        [&](const auto& val) {
            return processFirstVal(tx, val, val_stack, segment_ids, 1,
                                   value_cache);
        },
        record);
    return getValueInternal(tx, std::move(val_stack), segment_ids, value_cache,
                            lazy_load);
}

DbResult<value> getValueImpl(const ReadTransaction& tx,
                             const uint256_t value_hash,
                             std::set<uint64_t>& segment_ids,
                             ValueCache& value_cache,
                             bool lazy_load) {
    std::vector<ValueBeingParsed> val_stack{};

    // Initialize val_stack with first value from database
    // Note: only the hash field of this UnloadedValue is relevant
    auto result =
        processFirstVal(tx, BigUnloadedValue{HASH_PRE_IMAGE, value_hash, 1},
                        val_stack, segment_ids, 0, value_cache);
    if (!result.status.ok()) {
        return result.status;
    }
    auto res = getValueInternal(tx, std::move(val_stack), segment_ids,
                                value_cache, lazy_load);
    if (std::holds_alternative<rocksdb::Status>(res)) {
        return res;
    }
    auto& val = std::get<CountedData<value>>(res).data;
    if (std::holds_alternative<UnloadedValue>(val)) {
        throw std::runtime_error(
            "attempting to get value resulted in unloaded value");
    }
    assert(hash_value(val) == value_hash);
    return res;
}

DbResult<value> getValue(const ReadTransaction& tx,
                         const uint256_t value_hash,
                         ValueCache& value_cache,
                         bool lazy_load) {
    std::set<uint64_t> segment_ids{};
    return getValueImpl(tx, value_hash, segment_ids, value_cache, lazy_load);
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
            if (std::holds_alternative<UnloadedValue>(next_item)) {
                throw std::runtime_error(
                    "Attempted to save unknown unloaded value");
            }
            std::vector<unsigned char> value_vector{};
            auto new_items_to_save =
                serializeValue(tx.getSecretHashSeed(), next_item, value_vector,
                               segment_counts);
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
