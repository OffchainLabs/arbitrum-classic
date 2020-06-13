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

#include <avm_values/codepoint.hpp>
#include <avm_values/tuple.hpp>
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/checkpointutils.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/transaction.hpp>

namespace {
rocksdb::Slice vecToSlice(const std::vector<unsigned char>& vec) {
    return {reinterpret_cast<const char*>(vec.data()), vec.size()};
}

DbResult<Tuple> getTuple(const Transaction& transaction,
                         const std::vector<unsigned char>& hash_key,
                         TuplePool* pool) {
    auto key = vecToSlice(hash_key);
    std::vector<value> values;

    auto results = transaction.getData(key);

    if (!results.status.ok()) {
        return DbResult<Tuple>{results.status, results.reference_count,
                               Tuple()};
    }

    auto value_vectors = checkpoint::utils::parseTuple(results.stored_value);

    if (value_vectors.empty()) {
        return DbResult<Tuple>{results.status, results.reference_count,
                               Tuple()};
    }

    for (auto& current_vector : value_vectors) {
        auto value_type = static_cast<ValueTypes>(current_vector[0]);

        switch (value_type) {
            case NUM: {
                auto num =
                    checkpoint::utils::deserializeUint256_t(current_vector);
                values.push_back(num);
                break;
            }
            case CODEPT: {
                auto code_point =
                    checkpoint::utils::deserializeCodePointStub(current_vector);
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
                current_vector.erase(current_vector.begin());
                auto tuple = getTuple(transaction, current_vector, pool).data;
                values.push_back(tuple);
                break;
            }
        }
    }
    auto tuple = Tuple(values, pool);
    return DbResult<Tuple>{results.status, results.reference_count, tuple};
}
}  // namespace

DbResult<value> getValue(const Transaction& transaction,
                         const std::vector<unsigned char>& hash_key,
                         TuplePool* pool) {
    auto key = vecToSlice(hash_key);
    auto results = transaction.getData(key);

    if (!results.status.ok()) {
        auto error_res = DbResult<value>();
        error_res.status = results.status;
        error_res.reference_count = results.reference_count;
        return error_res;
    }

    auto value_type = static_cast<ValueTypes>(results.stored_value[0]);

    switch (value_type) {
        case NUM: {
            auto val =
                checkpoint::utils::deserializeUint256_t(results.stored_value);
            return DbResult<value>{results.status, results.reference_count,
                                   val};
        }
        case CODEPT: {
            auto code_point = checkpoint::utils::deserializeCodePointStub(
                results.stored_value);
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
            auto tuple_res = getTuple(transaction, hash_key, pool);
            return DbResult<value>{tuple_res.status, tuple_res.reference_count,
                                   tuple_res.data};
        }
    }
}

DbResult<MachineStateKeys> getMachineState(
    const Transaction& transaction,
    const std::vector<unsigned char>& checkpoint_name) {
    auto key = vecToSlice(checkpoint_name);
    auto results = transaction.getData(key);

    if (!results.status.ok()) {
        return DbResult<MachineStateKeys>{
            results.status, results.reference_count, MachineStateKeys()};
    }
    auto parsed_state =
        checkpoint::utils::extractStateKeys(results.stored_value);

    return DbResult<MachineStateKeys>{results.status, results.reference_count,
                                      parsed_state};
}
