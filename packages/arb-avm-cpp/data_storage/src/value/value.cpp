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
constexpr int TUP_BUFFER_LENGTH = 1;
constexpr int TUP_CODEPT_LENGTH = 49;

namespace {

struct ValueHash {
    uint256_t hash;
};

using ParsedTupVal = nonstd::variant<uint256_t, CodePointStub, Buffer, ValueHash>;

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
            case BUFFER: {
                return_vector.push_back(Buffer());
                iter += TUP_BUFFER_LENGTH;
                break;
            }
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
            case TUPLE: {
                return_vector.push_back(ValueHash{deserializeUint256t(buf)});
                iter += TUP_TUPLE_LENGTH;
                break;
            }
            default: {
                throw std::runtime_error(
                    "tried to parse tuple value with invalid typecode");
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
std::vector<value> serializeValue(const Buffer&,
                                  std::vector<unsigned char>& value_vector,
                                  std::map<uint64_t, uint64_t>&) {
    value_vector.push_back(BUFFER);
    return {};
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
                      std::vector<DeserializedValue>& vals,
                      std::vector<ParsedTupVal>& raw_vals,
                      std::set<uint64_t>& segment_ids);

GetResults processVal(const Transaction&,
                      const uint256_t& val,
                      std::vector<DeserializedValue>& vals,
                      std::vector<ParsedTupVal>&,
                      std::set<uint64_t>&) {
    vals.push_back(value{val});
    return GetResults{1, rocksdb::Status::OK(), {}};
}

GetResults processVal(const Transaction&,
                      const Buffer& val,
                      std::vector<DeserializedValue>& vals,
                      std::vector<ParsedTupVal>&,
                      std::set<uint64_t>&) {
    vals.push_back(value{val});
    return GetResults{1, rocksdb::Status::OK(), {}};
}

GetResults processVal(const Transaction&,
                      const CodePointStub& val,
                      std::vector<DeserializedValue>& vals,
                      std::vector<ParsedTupVal>&,
                      std::set<uint64_t>& segment_ids) {
    segment_ids.insert(val.pc.segment);
    vals.push_back(value{val});
    return GetResults{1, rocksdb::Status::OK(), {}};
}

GetResults processVal(const Transaction&,
                      const std::vector<ParsedTupVal>& val,
                      std::vector<DeserializedValue>& vals,
                      std::vector<ParsedTupVal>& raw_vals,
                      std::set<uint64_t>&) {
    vals.push_back(TuplePlaceholder{static_cast<uint8_t>(val.size())});
    raw_vals.insert(raw_vals.end(), val.rbegin(), val.rend());
    return GetResults{1, rocksdb::Status::OK(), {}};
}

GetResults processVal(const Transaction& transaction,
                      const ValueHash& val_hash,
                      std::vector<DeserializedValue>& vals,
                      std::vector<ParsedTupVal>& raw_vals,
                      std::set<uint64_t>& segment_ids) {
    std::vector<unsigned char> hash_key;
    marshal_uint256_t(val_hash.hash, hash_key);
    auto key = vecToSlice(hash_key);
    auto results = getRefCountedData(*transaction.transaction, key);
    if (!results.status.ok()) {
        return results;
    }

    auto record = parseRecord(results.stored_value);
    auto nested_ret = nonstd::visit(
        [&](const auto& val) {
            return processVal(transaction, val, vals, raw_vals, segment_ids);
        },
        record);
    if (!nested_ret.status.ok()) {
        return nested_ret;
    }
    return results;
}

DbResult<value> getValueImpl(const Transaction& transaction,
                             uint256_t value_hash,
                             std::set<uint64_t>& segment_ids) {
    std::vector<DeserializedValue> vals;
    std::vector<ParsedTupVal> raw_vals{ValueHash{value_hash}};
    bool first = true;
    GetResults ret;
    while (!raw_vals.empty()) {
        auto next = std::move(raw_vals.back());
        raw_vals.pop_back();
        auto results = nonstd::visit(
            [&](const auto& val) {
                return processVal(transaction, val, vals, raw_vals,
                                  segment_ids);
            },
            next);
        if (!results.status.ok()) {
            return {results.status, 0, Tuple()};
        }
        if (first) {
            ret = results;
            first = false;
        }
    }
    return {ret.status, ret.reference_count,
            assembleValueFromDeserialized(std::move(vals))};
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
        // std::cerr << "got item " << next_item << std::endl;
        auto hash = hash_value(next_item);
        std::vector<unsigned char> hash_key;
        // std::cerr << "got item 2" << std::endl;
        marshal_uint256_t(hash, hash_key);
        auto key = vecToSlice(hash_key);
        auto results = getRefCountedData(*transaction.transaction, key);
        // std::cerr << "got item 3" << std::endl;
        SaveResults save_ret;
        if (results.status.ok() && results.reference_count > 0) {
            // std::cerr << "got item inc" << std::endl;
            save_ret = incrementReference(*transaction.transaction, key);
        } else {
            std::vector<unsigned char> value_vector;
            // std::cerr << "got item ser" << std::endl;
            auto new_items_to_save =
                serializeValue(next_item, value_vector, segment_counts);
            // std::cerr << "got item ser" << std::endl;
            items_to_save.insert(items_to_save.end(), new_items_to_save.begin(),
                                 new_items_to_save.end());
            // std::cerr << "got item ser" << std::endl;
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
