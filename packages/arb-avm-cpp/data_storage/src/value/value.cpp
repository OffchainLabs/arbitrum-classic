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

#include <data_storage/storageresult.hpp>
#include <data_storage/value/checkpointutils.hpp>
#include <data_storage/value/transaction.hpp>

#include <avm_values/tuple.hpp>

constexpr int TUP_TUPLE_LENGTH = 33;
constexpr int TUP_NUM_LENGTH = 33;
constexpr int TUP_CODEPT_LENGTH = 41;

namespace {
rocksdb::Slice vecToSlice(const std::vector<unsigned char>& vec) {
    return {reinterpret_cast<const char*>(vec.data()), vec.size()};
}

std::vector<unsigned char> getHashKey(const value& val) {
    auto hash_key = hash_value(val);
    std::vector<unsigned char> hash_key_vector;
    marshal_uint256_t(hash_key, hash_key_vector);

    return hash_key_vector;
}

struct ValueSerializer {
    std::vector<unsigned char> operator()(const Tuple& val) const {
        std::vector<unsigned char> value_vector;
        value_vector.push_back(TUPLE);
        auto hash_key = hash_value(val);
        marshal_uint256_t(hash_key, value_vector);

        return value_vector;
    }

    std::vector<unsigned char> operator()(const uint256_t& val) const {
        std::vector<unsigned char> value_vector;
        value_vector.push_back(NUM);
        marshal_uint256_t(val, value_vector);

        return value_vector;
    }

    std::vector<unsigned char> operator()(const CodePointStub& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = static_cast<unsigned char>(CODEPT);
        value_vector.push_back(type_code);
        checkpoint::utils::serializeCodePointStub(val, value_vector);
        return value_vector;
    }

    std::vector<unsigned char> operator()(const HashPreImage& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = static_cast<unsigned char>(HASH_PRE_IMAGE);
        value_vector.push_back(type_code);
        val.marshal(value_vector);

        return value_vector;
    }
};

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
            case CODEPT: {
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

CodePointStub deserializeCodePointStub(const char*& bufptr) {
    auto pc_val = checkpoint::utils::deserialize_uint64(bufptr);
    auto hash_val = deserializeUint256t(bufptr);
    return {pc_val, hash_val};
}

DbResult<value> getTuple(const Transaction& transaction,
                         const GetResults& results,
                         TuplePool* pool);

DbResult<value> getTuple(const Transaction& transaction,
                         const rocksdb::Slice& key,
                         TuplePool* pool) {
    auto results = transaction.getData(key);

    if (!results.status.ok()) {
        return DbResult<value>{results.status, results.reference_count,
                               Tuple()};
    }
    return getTuple(transaction, results, pool);
}

DbResult<value> getTuple(const Transaction& transaction,
                         const GetResults& results,
                         TuplePool* pool) {
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
            case CODEPT: {
                auto code_point = deserializeCodePointStub(buf);
                values.push_back(code_point);
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
                auto tuple = getTuple(transaction, tupKey, pool).data;
                values.push_back(tuple);
                break;
            }
        }
    }
    auto tuple = Tuple(values, pool);
    return DbResult<value>{results.status, results.reference_count, tuple};
}

SaveResults saveTuple(Transaction& transaction, const Tuple& val) {
    auto hash_key = getHashKey(val);
    auto key = vecToSlice(hash_key);
    auto results = transaction.getData(key);

    auto incr_ref_count = results.status.ok() && results.reference_count > 0;

    if (incr_ref_count) {
        return transaction.incrementReference(key);
    } else {
        std::vector<unsigned char> value_vector;
        value_vector.push_back(TUPLE + val.tuple_size());

        for (uint64_t i = 0; i < val.tuple_size(); i++) {
            auto current_val = val.get_element(i);
            auto serialized_val = nonstd::visit(ValueSerializer{}, current_val);
            value_vector.insert(value_vector.end(), serialized_val.begin(),
                                serialized_val.end());

            auto type = static_cast<ValueTypes>(serialized_val[0]);
            if (type == TUPLE) {
                auto tup_val = nonstd::get<Tuple>(current_val);
                auto tuple_save_results = saveTuple(transaction, tup_val);
            }
        }
        return transaction.saveData(key, value_vector);
    }
}

DeleteResults deleteTuple(Transaction& transaction,
                          const rocksdb::Slice& hash_key);

DeleteResults deleteTuple(Transaction& transaction,
                          const rocksdb::Slice& hash_key,
                          GetResults results) {
    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    if (results.reference_count == 1) {
        auto value_vectors = parseTuple(results.stored_value);

        for (const auto& vec : value_vectors) {
            if (static_cast<ValueTypes>(vec[0]) == TUPLE) {
                rocksdb::Slice tupKey{
                    reinterpret_cast<const char*>(vec.data()) + 1,
                    vec.size() - 1};
                auto delete_status = deleteTuple(transaction, tupKey);
            }
        }
    }
    return transaction.deleteData(hash_key);
}

DeleteResults deleteTuple(Transaction& transaction,
                          const rocksdb::Slice& hash_key) {
    auto results = transaction.getData(hash_key);
    return deleteTuple(transaction, hash_key, results);
}

DeleteResults deleteValue(Transaction& transaction,
                          const rocksdb::Slice& hash_key) {
    auto results = transaction.getData(hash_key);

    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    auto type = static_cast<ValueTypes>(results.stored_value[0]);

    if (type == TUPLE) {
        return deleteTuple(transaction, hash_key, results);
    } else {
        return transaction.deleteData(hash_key);
    }
}
}  // namespace

DbResult<value> getValue(const Transaction& transaction,
                         uint256_t value_hash,
                         TuplePool* pool) {
    std::vector<unsigned char> hash_key;
    marshal_uint256_t(value_hash, hash_key);
    auto key = vecToSlice(hash_key);
    auto results = transaction.getData(key);

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
        case CODEPT: {
            auto code_point = deserializeCodePointStub(buf);
            return DbResult<value>{results.status, results.reference_count,
                                   code_point};
        }
        case HASH_PRE_IMAGE: {
            throw std::runtime_error("HASH_ONLY item");
            return DbResult<value>();
        }
        default: {
            if (value_type - TUPLE > 8) {
                throw std::runtime_error("can't get value with invalid type");
            }
            return getTuple(transaction, results, pool);
        }
    }
}

SaveResults saveValue(Transaction& transaction, const value& val) {
    if (nonstd::holds_alternative<Tuple>(val)) {
        auto tuple = nonstd::get<Tuple>(val);
        return saveTuple(transaction, tuple);
    } else {
        auto serialized_value = nonstd::visit(ValueSerializer{}, val);
        auto hash_key = getHashKey(val);
        auto key = vecToSlice(hash_key);
        return transaction.saveData(key, serialized_value);
    }
}

DeleteResults deleteValue(Transaction& transaction, uint256_t value_hash) {
    std::vector<unsigned char> hash_key;
    marshal_uint256_t(value_hash, hash_key);
    auto key = vecToSlice(hash_key);
    return deleteValue(transaction, key);
}
