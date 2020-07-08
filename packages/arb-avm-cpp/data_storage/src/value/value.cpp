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

DbResult<value> getTuple(const Transaction& transaction,
                         const GetResults& results,
                         TuplePool* pool,
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
                auto tuple =
                    getTuple(transaction, results, pool, segment_ids).data;
                values.push_back(tuple);
                break;
            }
        }
    }
    auto tuple = Tuple(values, pool);
    return DbResult<value>{results.status, results.reference_count, tuple};
}

struct ValueSerializer {
    std::vector<unsigned char>& value_vector;
    std::unordered_map<uint64_t, uint64_t>& segment_counts;

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

SaveResults saveTuple(Transaction& transaction,
                      const Tuple& val,
                      std::unordered_map<uint64_t, uint64_t>& segment_counts) {
    auto hash_key = getHashKey(val);
    auto key = vecToSlice(hash_key);
    auto results = getRefCountedData(*transaction.transaction, key);

    auto incr_ref_count = results.status.ok() && results.reference_count > 0;

    if (incr_ref_count) {
        return incrementReference(*transaction.transaction, key);
    }

    std::vector<unsigned char> value_vector;
    value_vector.push_back(TUPLE + val.tuple_size());

    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto current_val = val.get_element(i);
        nonstd::visit(ValueSerializer{value_vector, segment_counts},
                      current_val);

        if (nonstd::holds_alternative<Tuple>(current_val)) {
            auto tup_val = nonstd::get<Tuple>(current_val);
            auto tuple_save_results =
                saveTuple(transaction, tup_val, segment_counts);
        }
    }
    return saveRefCountedData(*transaction.transaction, key, value_vector);
}

struct ValueSaver {
    Transaction& transaction;
    std::unordered_map<uint64_t, uint64_t>& segment_counts;

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

DeleteResults deleteValue(
    Transaction& transaction,
    const rocksdb::Slice& hash_key,
    std::unordered_map<uint64_t, uint64_t>& segment_counts) {
    auto results = getRefCountedData(*transaction.transaction, hash_key);

    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    if (results.reference_count > 1) {
        auto buf = reinterpret_cast<const char*>(results.stored_value.data());
        auto value_type = static_cast<ValueTypes>(*buf);
        ++buf;

        if (value_type == TUPLE) {
            auto value_vectors = parseTuple(results.stored_value);

            for (const auto& vec : value_vectors) {
                rocksdb::Slice tupKey{
                    reinterpret_cast<const char*>(vec.data()) + 1,
                    vec.size() - 1};
                auto delete_status =
                    deleteValue(transaction, tupKey, segment_counts);
            }
        } else if (value_type == CODE_POINT_STUB) {
            auto code_point = deserializeCodePointStub(buf);
            ++segment_counts[code_point.pc.segment];
        }
    }

    return deleteRefCountedData(*transaction.transaction, hash_key);
}
}  // namespace

DbResult<value> getValueImpl(const Transaction& transaction,
                             uint256_t value_hash,
                             TuplePool* pool,
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
            return getTuple(transaction, results, pool, segment_ids);
        }
    }
}

DbResult<value> getValue(const Transaction& transaction,
                         uint256_t value_hash,
                         TuplePool* pool) {
    std::set<uint64_t> segment_ids;
    return getValueImpl(transaction, value_hash, pool, segment_ids);
}

SaveResults saveValueImpl(
    Transaction& transaction,
    const value& val,
    std::unordered_map<uint64_t, uint64_t>& segment_counts) {
    return nonstd::visit(ValueSaver{transaction, segment_counts}, val);
}

SaveResults saveValue(Transaction& transaction, const value& val) {
    std::unordered_map<uint64_t, uint64_t> segment_counts;
    return saveValueImpl(transaction, val, segment_counts);
}

DeleteResults deleteValueImpl(
    Transaction& transaction,
    const uint256_t& value_hash,
    std::unordered_map<uint64_t, uint64_t>& segment_counts) {
    std::vector<unsigned char> hash_key;
    marshal_uint256_t(value_hash, hash_key);
    auto key = vecToSlice(hash_key);
    return deleteValue(transaction, key, segment_counts);
}

DeleteResults deleteValue(Transaction& transaction, uint256_t value_hash) {
    std::unordered_map<uint64_t, uint64_t> segment_counts;
    return deleteValueImpl(transaction, value_hash, segment_counts);
}
