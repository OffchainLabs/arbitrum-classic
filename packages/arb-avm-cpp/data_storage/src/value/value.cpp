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
std::vector<unsigned char> getHashKey(const value& val) {
    auto hash_key = hash_value(val);
    std::vector<unsigned char> hash_key_vector;
    marshal_uint256_t(hash_key, hash_key_vector);

    return hash_key_vector;
}

std::vector<std::vector<unsigned char>> parseTuple(
    const std::vector<unsigned char>& data) {
    std::vector<std::vector<unsigned char>> return_vector;

    auto iter = data.begin();
    uint8_t count = *iter - TUPLE;
    ++iter;

    for (uint8_t i = 0; i < count; i++) {
        auto value_type = static_cast<ValueTypes>(*iter);
        std::vector<unsigned char> current;

        switch (value_type) {
            case NUM: {
                auto next_it = iter + TUP_NUM_LENGTH;
                current.insert(current.end(), iter, next_it);
                iter = next_it;
                break;
            }
            case CODE_POINT_STUB: {
                auto next_it = iter + TUP_CODEPT_LENGTH;
                current.insert(current.end(), iter, next_it);
                iter = next_it;
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
                auto next_it = iter + TUP_TUPLE_LENGTH;
                current.insert(current.end(), iter, next_it);
                iter = next_it;
                break;
            }
        }
        return_vector.push_back(current);
    }
    return return_vector;
}

void parseTupleForDeletion(const std::vector<unsigned char>& tuple_data,
                           std::vector<uint256_t>& tuples_to_delete,
                           std::map<uint64_t, uint64_t>& segment_counts) {
    auto iter = tuple_data.begin();
    uint8_t count = *iter - TUPLE;
    ++iter;

    for (uint8_t i = 0; i < count; i++) {
        auto value_type = static_cast<ValueTypes>(*iter);
        auto buf = reinterpret_cast<const char*>(&*iter);
        switch (value_type) {
            case NUM: {
                iter += TUP_NUM_LENGTH;
                break;
            }
            case CODE_POINT_STUB: {
                ++buf;
                auto cp = deserializeCodePointStub(buf);
                ++segment_counts[cp.pc.segment];
                iter += TUP_CODEPT_LENGTH;
                break;
            }
            case HASH_PRE_IMAGE: {
                throw std::runtime_error("HASH_ONLY item");
            }
            case TUPLE: {
                ++buf;
                tuples_to_delete.push_back(deserializeUint256t(buf));
                iter += TUP_TUPLE_LENGTH;
                break;
            }
            default: {
                throw std::runtime_error("invalid value type in tuple");
            }
        }
    }
}

DbResult<value> getTuple(const Transaction& transaction,
                         const GetResults& results,
                         std::set<uint64_t>& segment_ids) {
    auto value_vectors = parseTuple(results.stored_value);

    if (value_vectors.empty()) {
        return DbResult<value>{results.status, results.reference_count,
                               Tuple()};
    }

    std::vector<value> values;
    for (auto& current_vector : value_vectors) {
        auto buf = reinterpret_cast<const char*>(current_vector.data());
        auto value_type = static_cast<ValueTypes>(*buf);
        ++buf;

        switch (value_type) {
            case NUM: {
                values.push_back(deserializeUint256t(buf));
                break;
            }
            case CODE_POINT_STUB: {
                auto code_point = deserializeCodePointStub(buf);
                values.push_back(code_point);
                segment_ids.insert(code_point.pc.segment);
                break;
            }
            case HASH_PRE_IMAGE: {
                throw std::runtime_error("HASH_ONLY item");
            }
            default: {
                auto tup_size = value_type - TUPLE;
                if (tup_size > 8) {
                    throw std::runtime_error(
                        "tried to get value inside tuple with invalid "
                        "typecode");
                }
                rocksdb::Slice tupKey(buf, 32);
                auto results =
                    getRefCountedData(*transaction.transaction, tupKey);

                if (!results.status.ok()) {
                    return DbResult<value>{results.status,
                                           results.reference_count, Tuple()};
                }
                auto tuple = getTuple(transaction, results, segment_ids).data;
                values.push_back(tuple);
                break;
            }
        }
    }
    auto tuple = Tuple(std::move(values));
    return DbResult<value>{results.status, results.reference_count, tuple};
}

struct ValueSerializer {
    std::vector<unsigned char>& value_vector;
    std::map<uint64_t, uint64_t>& segment_counts;

    void operator()(const Tuple& val) const {
        value_vector.push_back(TUPLE);
        auto hash_key = hash_value(val);
        marshal_uint256_t(hash_key, value_vector);
    }

    void operator()(const uint256_t& val) const {
        value_vector.push_back(NUM);
        marshal_uint256_t(val, value_vector);
    }

    void operator()(const CodePointStub& val) const {
        value_vector.push_back(CODE_POINT_STUB);
        val.marshal(value_vector);
        ++segment_counts[val.pc.segment];
    }

    void operator()(const HashPreImage& val) const {
        value_vector.push_back(HASH_PRE_IMAGE);
        val.marshal(value_vector);
    }
};

struct BasicValChecker {
    Transaction& transaction;

    bool operator()(const value& val) const {
        return nonstd::visit(*this, val);
    }
    bool operator()(const Tuple& val) const {
        auto hash_key = getHashKey(val);
        auto key = vecToSlice(hash_key);
        auto results = getRefCountedData(*transaction.transaction, key);
        return results.status.ok() && results.reference_count > 0;
    }

    template <typename T>
    bool operator()(const T&) const {
        return true;
    }
};

SaveResults saveTupleRecord(Transaction& transaction,
                            const Tuple& val,
                            std::map<uint64_t, uint64_t>& segment_counts) {
    auto hash_key = getHashKey(val);
    auto key = vecToSlice(hash_key);
    std::vector<unsigned char> value_vector;
    value_vector.push_back(TUPLE + val.tuple_size());

    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto current_val = val.get_element(i);
        nonstd::visit(ValueSerializer{value_vector, segment_counts},
                      current_val);
    }
    return saveRefCountedData(*transaction.transaction, key, value_vector);
}

bool isSaved(Transaction& transaction, const Tuple& val) {
    auto hash_key = getHashKey(val);
    auto key = vecToSlice(hash_key);
    auto results = getRefCountedData(*transaction.transaction, key);
    return results.status.ok() && results.reference_count > 0;
}

SaveResults saveTuple(Transaction& transaction,
                      const Tuple& val,
                      std::map<uint64_t, uint64_t>& segment_counts) {
    SaveResults ret;
    std::vector<Tuple> tups{val};
    while (!tups.empty()) {
        auto tup = tups.back();
        if (isSaved(transaction, val)) {
            auto hash_key = getHashKey(val);
            auto key = vecToSlice(hash_key);
            ret = incrementReference(*transaction.transaction, key);
            tups.pop_back();
        } else {
            bool found_complex = false;
            for (uint64_t i = 0; i < tup.tuple_size(); ++i) {
                auto& elem = tup.get_element_unsafe(i);
                if (!nonstd::holds_alternative<Tuple>(elem)) {
                    continue;
                }
                if (!isSaved(transaction, elem.get<Tuple>())) {
                    found_complex = true;
                    tups.push_back(tup.get_element(i).get<Tuple>());
                }
            }
            if (!found_complex) {
                ret = saveTupleRecord(transaction, tup, segment_counts);
                tups.pop_back();
            }
        }
    }
    return ret;
}

struct ValueSaver {
    Transaction& transaction;
    std::map<uint64_t, uint64_t>& segment_counts;

    template <typename T>
    SaveResults saveImpl(const T& val, bool allow_replacement) const {
        std::vector<unsigned char> serialized_value;
        ValueSerializer{serialized_value, segment_counts}(val);
        auto hash_key = getHashKey(val);
        auto key = vecToSlice(hash_key);
        return saveRefCountedData(*transaction.transaction, key,
                                  serialized_value, 1, allow_replacement);
    }

    SaveResults operator()(const Tuple& val) const {
        return saveTuple(transaction, val, segment_counts);
    }

    SaveResults operator()(const CodePointStub& val) const {
        // The same code point can exist in different segments with different
        // serializations mapping to the same hash. If this occurs, the
        // different versions are interchangeable
        return saveImpl(val, true);
    }

    template <typename T>
    SaveResults operator()(const T& val) const {
        return saveImpl(val, false);
    }
};

DeleteResults deleteTuple(Transaction& transaction,
                          const uint256_t& tuple_hash,
                          std::map<uint64_t, uint64_t>& segment_counts) {
    std::vector<uint256_t> tuples_to_delete{tuple_hash};
    DeleteResults ret;
    while (!tuples_to_delete.empty()) {
        auto tup_hash = std::move(tuples_to_delete.back());
        tuples_to_delete.pop_back();
        std::vector<unsigned char> hash_key;
        marshal_uint256_t(tup_hash, hash_key);
        auto key = vecToSlice(hash_key);

        auto results = getRefCountedData(*transaction.transaction, key);
        if (!results.status.IsNotFound()) {
            // value was already deleted, this shouldn't happen, but we can
            // continue
            continue;
        }

        if (!results.status.ok()) {
            // Some unexpected error occured
            return DeleteResults{0, results.status};
        }

        if (results.reference_count == 1) {
            // This was the last reference to this tuple so decrement the
            // reference count of the references values
            parseTupleForDeletion(results.stored_value, tuples_to_delete,
                                  segment_counts);
        }

        ret = deleteRefCountedData(*transaction.transaction, key);
    }
    return ret;
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
    return nonstd::visit(ValueSaver{transaction, segment_counts}, val);
}

SaveResults saveValue(Transaction& transaction, const value& val) {
    std::map<uint64_t, uint64_t> segment_counts;
    return saveValueImpl(transaction, val, segment_counts);
}

DeleteResults deleteValueImpl(Transaction& transaction,
                              const uint256_t& value_hash,
                              std::map<uint64_t, uint64_t>& segment_counts) {
    std::vector<unsigned char> hash_key;
    marshal_uint256_t(value_hash, hash_key);
    auto key = vecToSlice(hash_key);

    auto results = getRefCountedData(*transaction.transaction, key);

    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    auto ret = deleteRefCountedData(*transaction.transaction, key);

    auto value_type = static_cast<ValueTypes>(*results.stored_value.data());

    if (results.reference_count == 1 &&
        (value_type >= TUPLE && value_type <= TUPLE + 8)) {
        // The value we deleted was a tuple
        std::vector<uint256_t> tuples_to_delete;
        parseTupleForDeletion(results.stored_value, tuples_to_delete,
                              segment_counts);
        for (const auto& tup : tuples_to_delete) {
            deleteTuple(transaction, tup, segment_counts);
        }
    }

    return ret;
}

DeleteResults deleteValue(Transaction& transaction, uint256_t value_hash) {
    std::map<uint64_t, uint64_t> segment_counts;
    return deleteValueImpl(transaction, value_hash, segment_counts);
}
