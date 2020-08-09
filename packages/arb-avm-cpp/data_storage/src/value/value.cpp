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
struct SerializedTupStub {
    uint8_t size;
    uint256_t hash;
};

using ParsedTupVal =
    nonstd::variant<uint256_t, CodePointStub, SerializedTupStub>;

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
                return_vector.push_back(deserializeUint256t(buf));
                iter += TUP_NUM_LENGTH;
                break;
            }
            case CODE_POINT_STUB: {
                return_vector.push_back(deserializeCodePointStub(buf));
                iter += TUP_CODEPT_LENGTH;
                break;
            }
            case HASH_PRE_IMAGE: {
                throw std::runtime_error("HASH_ONLY item");
            }
            default: {
                uint8_t tup_size = value_type - TUPLE;
                if (tup_size > 8) {
                    throw std::runtime_error(
                        "tried to parse tuple with invalid typecode");
                }
                return_vector.push_back(
                    SerializedTupStub{tup_size, deserializeUint256t(buf)});
                iter += TUP_TUPLE_LENGTH;
                break;
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

DbResult<value> getTuple(const Transaction& transaction,
                         const GetResults& results,
                         std::set<uint64_t>& segment_ids);

struct TupleGetter {
    const Transaction& transaction;
    std::set<uint64_t>& segment_ids;

    DbResult<value> operator()(const uint256_t& val) const {
        return {rocksdb::Status::OK(), 1, val};
    }

    DbResult<value> operator()(const CodePointStub& val) const {
        segment_ids.insert(val.pc.segment);
        return {rocksdb::Status::OK(), 1, val};
    }

    DbResult<value> operator()(const SerializedTupStub& val) const {
        std::vector<unsigned char> hash_key;
        marshal_uint256_t(val.hash, hash_key);
        auto key = vecToSlice(hash_key);
        auto results = getRefCountedData(*transaction.transaction, key);

        if (!results.status.ok()) {
            return DbResult<value>{results.status, results.reference_count,
                                   Tuple()};
        }
        return getTuple(transaction, results, segment_ids);
    }
};

DbResult<value> getTuple(const Transaction& transaction,
                         const GetResults& results,
                         std::set<uint64_t>& segment_ids) {
    std::vector<value> values;
    auto nested_values = parseTuple(results.stored_value);
    for (auto& current_vector : nested_values) {
        auto val = nonstd::visit(TupleGetter{transaction, segment_ids},
                                 current_vector);
        if (!val.status.ok()) {
            return DbResult<value>{val.status, val.reference_count, Tuple()};
        }
        values.push_back(std::move(val.data));
    }
    auto tuple = Tuple(std::move(values));
    return DbResult<value>{results.status, results.reference_count, tuple};
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
            value_vector.push_back(TUPLE + nested_tup.tuple_size());
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
std::vector<uint256_t> deleteParsedValue(const uint256_t&,
                                         std::map<uint64_t, uint64_t>&) {
    return {};
}
std::vector<uint256_t> deleteParsedValue(
    const CodePointStub& cp,
    std::map<uint64_t, uint64_t>& segment_counts) {
    segment_counts[cp.pc.segment]++;
    return {};
}
std::vector<uint256_t> deleteParsedValue(const std::vector<ParsedTupVal>& tup,
                                         std::map<uint64_t, uint64_t>&) {
    std::vector<uint256_t> vals_to_delete;
    for (const auto& val : tup) {
        // We only need to delete tuples since other values are recorded inline
        if (nonstd::holds_alternative<SerializedTupStub>(val)) {
            vals_to_delete.push_back(val.get<SerializedTupStub>().hash);
        }
    }
    return vals_to_delete;
}
}  // namespace

DbResult<value> getValueImpl(const Transaction& transaction,
                             uint256_t value_hash,
                             std::set<uint64_t>& segment_ids) {
    std::vector<unsigned char> hash_key;
    marshal_uint256_t(value_hash, hash_key);
    auto key = vecToSlice(hash_key);
    auto results = getRefCountedData(*transaction.transaction, key);

    if (!results.status.ok()) {
        return DbResult<value>{results.status, results.reference_count,
                               Tuple()};
    }

    auto buf = reinterpret_cast<const char*>(results.stored_value.data());
    auto value_type = static_cast<ValueTypes>(*buf);
    ++buf;

    switch (value_type) {
        case NUM: {
            auto val = deserializeUint256t(buf);
            return DbResult<value>{results.status, results.reference_count,
                                   val};
        }
        case CODE_POINT_STUB: {
            auto code_point = deserializeCodePointStub(buf);
            segment_ids.insert(code_point.pc.segment);
            return DbResult<value>{results.status, results.reference_count,
                                   code_point};
        }
        case HASH_PRE_IMAGE: {
            throw std::runtime_error("HASH_ONLY item");
        }
        default: {
            if (value_type - TUPLE > 8) {
                throw std::runtime_error("can't get value with invalid type");
            }
            return getTuple(transaction, results, segment_ids);
        }
    }
}

DbResult<value> getValue(const Transaction& transaction, uint256_t value_hash) {
    std::set<uint64_t> segment_ids;
    return getValueImpl(transaction, value_hash, segment_ids);
}

SaveResults saveValueImpl(Transaction& transaction,
                          const value& val,
                          std::map<uint64_t, uint64_t>& segment_counts) {
    bool first = true;
    SaveResults ret;
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
    DeleteResults ret;
    std::vector<uint256_t> items_to_delete{value_hash};
    while (!items_to_delete.empty()) {
        auto next_item = std::move(items_to_delete.back());
        items_to_delete.pop_back();
        std::vector<unsigned char> hash_key;
        marshal_uint256_t(next_item, hash_key);
        auto key = vecToSlice(hash_key);
        auto results = deleteRefCountedData(*transaction.transaction, key);
        if (results.status.ok() && results.reference_count == 0) {
            auto new_items_to_delete = nonstd::visit(
                [&](const auto& val) {
                    return deleteParsedValue(val, segment_counts);
                },
                parseRecord(results.stored_value));
            items_to_delete.insert(items_to_delete.end(),
                                   new_items_to_delete.begin(),
                                   new_items_to_delete.end());
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
