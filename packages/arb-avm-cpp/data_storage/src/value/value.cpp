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

#include "referencecount.hpp"
#include "utils.hpp"

#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>

#include <avm_values/tuple.hpp>

constexpr int TUP_TUPLE_LENGTH = 33;
constexpr int TUP_NUM_LENGTH = 33;
constexpr int TUP_CODEPT_LENGTH = 49;

namespace {

struct ValueHash {
    uint256_t hash;
};

using ParsedTupVal = nonstd::variant<uint256_t, CodePointStub, ValueHash>;

using ParsedSerializedVal =
    nonstd::variant<uint256_t, CodePointStub, std::vector<ParsedTupVal>>;

std::vector<ParsedTupVal> parseTuple(const std::vector<unsigned char>& data) {
    std::vector<ParsedTupVal> return_vector;

    auto iter = data.begin();
    uint8_t count = *iter - TUPLE;
    ++iter;

    for (uint8_t i = 0; i < count; i++) {
        auto value_type = static_cast<ValueTypes>(*iter);
        auto buf = reinterpret_cast<const char*>(&*iter);
        ++buf;

        switch (value_type) {
            case NUM: {
                return_vector.emplace_back(deserializeUint256t(buf));
                iter += TUP_NUM_LENGTH;
                break;
            }
            case CODE_POINT_STUB: {
                return_vector.emplace_back(deserializeCodePointStub(buf));
                iter += TUP_CODEPT_LENGTH;
                break;
            }
            case HASH_PRE_IMAGE: {
                throw std::runtime_error("HASH_ONLY item");
            }
            case TUPLE: {
                return_vector.emplace_back(ValueHash{deserializeUint256t(buf)});
                iter += TUP_TUPLE_LENGTH;
                break;
            }
            default: {
                throw std::runtime_error(
                    "tried to parse tuple value with invalid typecode");
            }
        }
    }
    return return_vector;
}

ParsedSerializedVal parseRecord(const std::vector<unsigned char>& data) {
    auto buf = reinterpret_cast<const char*>(data.data());
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
        default: {
            if (value_type - TUPLE > 8) {
                throw std::runtime_error("can't get value with invalid type");
            }
            return parseTuple(data);
        }
    }
}

std::vector<value> serializeValue(const value& val,
                                  std::vector<unsigned char>& value_vector,
                                  std::map<uint64_t, uint64_t>& segment_counts);
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
    std::vector<value> ret;
    value_vector.push_back(TUPLE + val.tuple_size());
    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto nested = val.get_element_unsafe(i);
        if (nonstd::holds_alternative<Tuple>(nested)) {
            const auto& nested_tup = nested.get<Tuple>();
            value_vector.push_back(TUPLE);
            marshal_uint256_t(hash(nested_tup), value_vector);
            ret.push_back(nested);
        } else {
            serializeValue(nested, value_vector, segment_counts);
        }
    }
    return ret;
}
std::vector<value> serializeValue(const HashPreImage&,
                                  std::vector<unsigned char>&,
                                  std::map<uint64_t, uint64_t>&) {
    throw std::runtime_error("Can't serialize hash preimage in db");
}
std::vector<value> serializeValue(
    const value& val,
    std::vector<unsigned char>& value_vector,
    std::map<uint64_t, uint64_t>& segment_counts) {
    return nonstd::visit(
        [&](const auto& val) {
            return serializeValue(val, value_vector, segment_counts);
        },
        val);
}

// Returns a list of value hashes to be deleted
void deleteParsedValue(const uint256_t&,
                       std::vector<uint256_t>&,
                       std::map<uint64_t, uint64_t>&) {}
void deleteParsedValue(const CodePointStub& cp,
                       std::vector<uint256_t>&,
                       std::map<uint64_t, uint64_t>& segment_counts) {
    segment_counts[cp.pc.segment]++;
}
void deleteParsedValue(const std::vector<ParsedTupVal>& tup,
                       std::vector<uint256_t>& vals_to_delete,
                       std::map<uint64_t, uint64_t>&) {
    for (const auto& val : tup) {
        // We only need to delete tuples since other values are recorded inline
        if (nonstd::holds_alternative<ValueHash>(val)) {
            vals_to_delete.push_back(val.get<ValueHash>().hash);
        }
    }
}
}  // namespace

GetResults processVal(const Transaction& transaction,
                      const ValueHash& val_hash,
                      std::vector<std::vector<DeserializedValue>>& vals_stack,
                      std::vector<std::vector<ParsedTupVal>>& raw_vals_stack,
                      std::vector<uint32_t>& reference_count_stack,
                      std::set<uint64_t>& segment_ids,
                      ValueCache& val_cache);

GetResults processVal(const Transaction&,
                      const uint256_t& val,
                      std::vector<std::vector<DeserializedValue>>& vals_stack,
                      std::vector<std::vector<ParsedTupVal>>&,
                      std::vector<uint32_t>& reference_count_stack,
                      std::set<uint64_t>&,
                      ValueCache& val_cache) {
    if (reference_count_stack.back() > 1) {
        val_cache[hash_value(val)] = val;
    }

    vals_stack.back().push_back(value{val});
    return GetResults{1, rocksdb::Status::OK(), {}};
}

GetResults processVal(const Transaction&,
                      const CodePointStub& val,
                      std::vector<std::vector<DeserializedValue>>& vals_stack,
                      std::vector<std::vector<ParsedTupVal>>&,
                      std::vector<uint32_t>& reference_count_stack,
                      std::set<uint64_t>& segment_ids,
                      ValueCache& val_cache) {
    if (reference_count_stack.back() > 1) {
        val_cache[hash_value(val)] = val;
    }

    segment_ids.insert(val.pc.segment);
    vals_stack.back().push_back(value{val});
    return GetResults{1, rocksdb::Status::OK(), {}};
}

GetResults processVal(const Transaction&,
                      const std::vector<ParsedTupVal>& val,
                      std::vector<std::vector<DeserializedValue>>& vals_stack,
                      std::vector<std::vector<ParsedTupVal>>& raw_vals_stack,
                      std::vector<uint32_t>& reference_count_stack,
                      std::set<uint64_t>&,
                      ValueCache&) {
    vals_stack.emplace_back();
    vals_stack.back().push_back(
        TuplePlaceholder{static_cast<uint8_t>(val.size())});
    raw_vals_stack.emplace_back();
    raw_vals_stack.back().insert(raw_vals_stack.back().end(), val.rbegin(),
                                 val.rend());
    reference_count_stack.push_back(0);
    return GetResults{1, rocksdb::Status::OK(), {}};
}

GetResults processVal(const Transaction& transaction,
                      const ValueHash& val_hash,
                      std::vector<std::vector<DeserializedValue>>& vals_stack,
                      std::vector<std::vector<ParsedTupVal>>& raw_vals_stack,
                      std::vector<uint32_t>& reference_count_stack,
                      std::set<uint64_t>& segment_ids,
                      ValueCache& val_cache) {
    auto cached_value = val_cache.find(val_hash.hash);
    if (cached_value != val_cache.end()) {
        vals_stack.back().push_back(cached_value->second);
        return {0, rocksdb::Status::OK(), {}};
    }

    std::vector<unsigned char> hash_key;
    marshal_uint256_t(val_hash.hash, hash_key);
    auto key = vecToSlice(hash_key);
    auto results = getRefCountedData(*transaction.transaction, key);
    reference_count_stack.back() = results.reference_count;
    if (!results.status.ok()) {
        return results;
    }

    auto record = parseRecord(results.stored_value);

    auto nested_ret = nonstd::visit(
        [&](const auto& val) {
            return processVal(transaction, val, vals_stack, raw_vals_stack,
                              reference_count_stack, segment_ids, val_cache);
        },
        record);
    if (!nested_ret.status.ok()) {
        return nested_ret;
    }

    return results;
}

DbResult<value> getValueImpl(const Transaction& transaction,
                             uint256_t value_hash,
                             std::set<uint64_t>& segment_ids,
                             ValueCache& value_cache) {
    std::vector<std::vector<DeserializedValue>> vals_stack{{}};
    std::vector<std::vector<ParsedTupVal>> raw_vals_stack{
        {ValueHash{value_hash}}};
    std::vector<uint32_t> reference_count_stack{0};
    while ((raw_vals_stack.size() > 1) || !raw_vals_stack[0].empty()) {
        if (!raw_vals_stack.back().empty()) {
            auto next = std::move(raw_vals_stack.back().back());
            raw_vals_stack.back().pop_back();

            auto results = nonstd::visit(
                [&](const auto& val) {
                    return processVal(transaction, val, vals_stack,
                                      raw_vals_stack, reference_count_stack,
                                      segment_ids, value_cache);
                },
                next);
            if (!results.status.ok()) {
                return {results.status, 0, Tuple()};
            }
        } else {
            // All child values resolved
            raw_vals_stack.pop_back();

            auto value = assembleSingleValueFromDeserialized(
                std::move(vals_stack.back()));
            vals_stack.pop_back();

            reference_count_stack.pop_back();
            auto reference_count = reference_count_stack.back();

            if (reference_count > 1) {
                value_cache[hash_value(value)] = value;
            }

            vals_stack.back().emplace_back(std::move(value));
        }
    }

    return {rocksdb::Status::OK(), reference_count_stack.back(),
            std::move(vals_stack.back().back().get<value>())};
}

DbResult<value> getValue(const Transaction& transaction,
                         uint256_t value_hash,
                         ValueCache& value_cache) {
    std::set<uint64_t> segment_ids;
    return getValueImpl(transaction, value_hash, segment_ids, value_cache);
}

SaveResults saveValueImpl(Transaction& transaction,
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
        auto results = getRefCountedData(*transaction.transaction, key);
        SaveResults save_ret;
        if (results.status.ok() && results.reference_count > 0) {
            save_ret = incrementReference(*transaction.transaction, key);
        } else {
            std::vector<unsigned char> value_vector;
            auto new_items_to_save =
                serializeValue(next_item, value_vector, segment_counts);
            items_to_save.insert(items_to_save.end(), new_items_to_save.begin(),
                                 new_items_to_save.end());
            save_ret =
                saveRefCountedData(*transaction.transaction, key, value_vector);
        }
        if (first) {
            ret = save_ret;
            first = false;
        }
    }
    return ret;
}

SaveResults saveValue(Transaction& transaction, const value& val) {
    std::map<uint64_t, uint64_t> segment_counts;
    return saveValueImpl(transaction, val, segment_counts);
}

DeleteResults deleteValueImpl(Transaction& transaction,
                              const uint256_t& value_hash,
                              std::map<uint64_t, uint64_t>& segment_counts) {
    bool first = true;
    DeleteResults ret{};
    std::vector<uint256_t> items_to_delete{value_hash};
    while (!items_to_delete.empty()) {
        auto next_item = items_to_delete.back();
        items_to_delete.pop_back();
        std::vector<unsigned char> hash_key;
        marshal_uint256_t(next_item, hash_key);
        auto key = vecToSlice(hash_key);
        auto results = deleteRefCountedData(*transaction.transaction, key);
        if (results.status.ok() && results.reference_count == 0) {
            nonstd::visit(
                [&](const auto& val) {
                    deleteParsedValue(val, items_to_delete, segment_counts);
                },
                parseRecord(results.stored_value));
        }
        if (first) {
            ret = results;
            first = false;
        }
    }
    return ret;
}

DeleteResults deleteValue(Transaction& transaction, uint256_t value_hash) {
    std::map<uint64_t, uint64_t> segment_counts;
    return deleteValueImpl(transaction, value_hash, segment_counts);
}
