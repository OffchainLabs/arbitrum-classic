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

#include <avm/checkpoint/machinestatefetcher.hpp>

#include <avm/checkpoint/checkpointresult.hpp>
#include <avm/checkpoint/checkpointstorage.hpp>
#include <avm/checkpoint/transaction.hpp>
#include <avm/value/codepoint.hpp>
#include <avm/value/tuple.hpp>

MachineStateFetcher::MachineStateFetcher(const CheckpointStorage& storage,
                                         TuplePool* pool_,
                                         const std::vector<CodePoint> code_)
    : checkpoint_storage(storage), pool(pool_), code(std::move(code_)) {}

DbResult<ParsedState> MachineStateFetcher::getMachineState(
    const std::vector<unsigned char>& checkpoint_name) const {
    auto results = checkpoint_storage.getValue(checkpoint_name);

    if (results.status.ok()) {
        auto parsed_state = checkpoint::utils::parseState(results.stored_value);

        return DbResult<ParsedState>{results.status, results.reference_count,
                                     parsed_state};
    } else {
        return DbResult<ParsedState>{results.status, results.reference_count,
                                     ParsedState()};
    }
}

DbResult<CodePoint> MachineStateFetcher::getCodePoint(
    const std::vector<unsigned char>& hash_key) const {
    auto results = checkpoint_storage.getValue(hash_key);

    if (results.status.ok()) {
        auto code_point =
            checkpoint::utils::deserializeCodepoint(results.stored_value, code);
        return DbResult<CodePoint>{results.status, results.reference_count,
                                   code_point};
    } else {
        return DbResult<CodePoint>{results.status, results.reference_count,
                                   CodePoint()};
    }
}

DbResult<uint256_t> MachineStateFetcher::getUint256_t(
    const std::vector<unsigned char>& hash_key) const {
    auto results = checkpoint_storage.getValue(hash_key);

    if (results.status.ok()) {
        auto num =
            checkpoint::utils::deserializeUint256_t(results.stored_value);
        return DbResult<uint256_t>{results.status, results.reference_count,
                                   num};
    } else {
        return DbResult<uint256_t>{results.status, results.reference_count, 0};
    }
}

DbResult<Tuple> MachineStateFetcher::getTuple(
    const std::vector<unsigned char>& hash_key) const {
    std::vector<value> values;
    auto results = checkpoint_storage.getValue(hash_key);

    if (results.status.ok()) {
        auto value_vectors =
            checkpoint::utils::parseTuple(results.stored_value);

        if (value_vectors.empty()) {
            return DbResult<Tuple>{results.status, results.reference_count,
                                   Tuple()};
        } else {
            for (auto& current_vector : value_vectors) {
                auto value_type = static_cast<ValueTypes>(current_vector[0]);

                switch (value_type) {
                    case TUPLE: {
                        current_vector.erase(current_vector.begin());
                        auto tuple = getTuple(current_vector).data;
                        values.push_back(tuple);
                        break;
                    }
                    case NUM: {
                        auto num = checkpoint::utils::deserializeUint256_t(
                            current_vector);
                        values.push_back(num);
                        break;
                    }
                    case CODEPT: {
                        auto code_point =
                            checkpoint::utils::deserializeCodepoint(
                                current_vector, code);
                        values.push_back(code_point);
                        break;
                    }
                    case HASH_ONLY: {
                        throw std::runtime_error("HASH_ONLY item");
                    }
                }
            }
            auto tuple = Tuple(values, pool);
            return DbResult<Tuple>{results.status, results.reference_count,
                                   tuple};
        }
    } else {
        return DbResult<Tuple>{results.status, results.reference_count,
                               Tuple()};
    }
};

DbResult<value> MachineStateFetcher::getValue(
    const std::vector<unsigned char>& hash_key) const {
    auto results = checkpoint_storage.getValue(hash_key);

    if (results.status.ok()) {
        auto value_type = static_cast<ValueTypes>(results.stored_value[0]);

        switch (value_type) {
            case TUPLE: {
                auto tuple_res = getTuple(hash_key);
                return DbResult<value>{tuple_res.status,
                                       tuple_res.reference_count,
                                       tuple_res.data};
            }
            case NUM: {
                auto val = checkpoint::utils::deserializeUint256_t(
                    results.stored_value);
                return DbResult<value>{results.status, results.reference_count,
                                       val};
            }
            case CODEPT: {
                auto code_point = checkpoint::utils::deserializeCodepoint(
                    results.stored_value, code);
                return DbResult<value>{results.status, results.reference_count,
                                       code_point};
            }
            case HASH_ONLY: {
                throw std::runtime_error("HASH_ONLY item");
            }
        }
    } else {
        auto error_res = DbResult<value>();
        error_res.status = results.status;
        error_res.reference_count = results.reference_count;
        return error_res;
    }
}
