/*
 * Copyright 2019, Offchain Labs, Inc.
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

#include <data_storage/checkpoint/machinestatefetcher.hpp>

#include <avm_values/tuple.hpp>
#include <data_storage/checkpoint/checkpointutils.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/transaction.hpp>

namespace {
rocksdb::Slice vecToSlice(const std::vector<unsigned char>& vec) {
    return {reinterpret_cast<const char*>(vec.data()), vec.size()};
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
    auto value_vectors = checkpoint::utils::parseTuple(results.stored_value);

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
