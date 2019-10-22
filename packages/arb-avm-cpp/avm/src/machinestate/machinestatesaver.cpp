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

#include <variant>

#include <avm/machinestate/machinestatesaver.hpp>
#include <avm/machinestate/tokenTracker.hpp>
#include <avm/value/codepoint.hpp>
#include <avm/value/tuple.hpp>

MachineStateSaver::MachineStateSaver(CheckpointStorage* storage,
                                     TuplePool* pool_,
                                     std::vector<CodePoint> code_) {
    checkpoint_storage = storage;
    pool = pool_;
    code = code_;
}

DbResult<CodePoint> MachineStateSaver::getCodePoint(
    const std::vector<unsigned char>& hash_key) {
    auto results = checkpoint_storage->getValue(hash_key);

    if (results.status.ok()) {
        auto code_point =
            Checkpoint::Utils::deserializeCodepoint(results.stored_value, code);
        return DbResult<CodePoint>{results.status, results.reference_count,
                                   code_point};
    } else {
        return DbResult<CodePoint>{results.status, results.reference_count,
                                   CodePoint()};
    }
}

DbResult<uint256_t> MachineStateSaver::getUint256_t(
    const std::vector<unsigned char>& hash_key) {
    auto results = checkpoint_storage->getValue(hash_key);

    if (results.status.ok()) {
        auto num =
            Checkpoint::Utils::deserializeUint256_t(results.stored_value);
        return DbResult<uint256_t>{results.status, results.reference_count,
                                   num};
    } else {
        return DbResult<uint256_t>{results.status, results.reference_count, 0};
    }
}

SaveResults MachineStateSaver::saveValue(const value& val) {
    auto serialized_value = Checkpoint::Utils::serializeValue(val);
    auto type = (valueTypes)serialized_value[0];

    if (type == TUPLE_TYPE) {
        auto tuple = nonstd::get<Tuple>(val);
        return saveTuple(tuple);
    } else {
        auto hash_key = GetHashKey(val);
        return checkpoint_storage->saveValue(serialized_value, hash_key);
    }
}

DbResult<value> MachineStateSaver::getValue(
    const std::vector<unsigned char>& hash_key) {
    auto results = checkpoint_storage->getValue(hash_key);

    if (results.status.ok()) {
        auto value_type = (valueTypes)results.stored_value[0];

        switch (value_type) {
            case TUPLE_TYPE: {
                auto tuple_res = getTuple(hash_key);
                return DbResult<value>{tuple_res.status,
                                       tuple_res.reference_count,
                                       tuple_res.data};
            }
            case NUM_TYPE: {
                auto val = Checkpoint::Utils::deserializeUint256_t(
                    results.stored_value);
                return DbResult<value>{results.status, results.reference_count,
                                       val};
            }
            case CODEPT_TYPE: {
                auto code_point = Checkpoint::Utils::deserializeCodepoint(
                    results.stored_value, code);
                return DbResult<value>{results.status, results.reference_count,
                                       code_point};
            }
        }
    } else {
        auto error_res = DbResult<value>();
        error_res.status = results.status;
        error_res.reference_count = results.reference_count;
        return error_res;
    }
}

SaveResults MachineStateSaver::saveTuple(const Tuple& val) {
    auto hash_key = GetHashKey(val);
    auto results = checkpoint_storage->getValue(hash_key);

    auto incr_ref_count = results.status.ok() && results.reference_count > 0;

    if (incr_ref_count) {
        return checkpoint_storage->incrementReference(hash_key);
    } else {
        std::vector<unsigned char> value_vector{(unsigned char)TUPLE_TYPE};

        for (uint64_t i = 0; i < val.tuple_size(); i++) {
            auto current_val = val.get_element(i);
            auto serialized_val =
                Checkpoint::Utils::serializeValue(current_val);

            value_vector.insert(value_vector.end(), serialized_val.begin(),
                                serialized_val.end());

            auto type = (valueTypes)serialized_val[0];
            if (type == TUPLE_TYPE) {
                auto tup_val = nonstd::get<Tuple>(current_val);
                auto tuple_save_results = saveTuple(tup_val);
            }
        }
        return checkpoint_storage->saveValue(value_vector, hash_key);
    }
};

DbResult<Tuple> MachineStateSaver::getTuple(
    const std::vector<unsigned char>& hash_key) {
    std::vector<value> values;
    auto results = checkpoint_storage->getValue(hash_key);

    if (results.status.ok()) {
        auto value_vectors =
            Checkpoint::Utils::parseSerializedTuple(results.stored_value);

        if (value_vectors.empty()) {
            return DbResult<Tuple>{results.status, results.reference_count,
                                   Tuple()};
        } else {
            for (auto& current_vector : value_vectors) {
                auto value_type = (valueTypes)current_vector[0];

                switch (value_type) {
                    case TUPLE_TYPE: {
                        current_vector.erase(current_vector.begin());
                        auto tuple = getTuple(current_vector).data;
                        values.push_back(tuple);
                        break;
                    }
                    case NUM_TYPE: {
                        auto num = Checkpoint::Utils::deserializeUint256_t(
                            current_vector);
                        values.push_back(num);
                        break;
                    }
                    case CODEPT_TYPE: {
                        auto code_point =
                            Checkpoint::Utils::deserializeCodepoint(
                                current_vector, code);
                        values.push_back(code_point);
                        break;
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

DbResult<ParsedState> MachineStateSaver::getMachineState(
    const std::vector<unsigned char>& checkpoint_name) {
    std::vector<unsigned char> name_vector(checkpoint_name.begin(),
                                           checkpoint_name.end());

    auto results = checkpoint_storage->getValue(name_vector);

    if (results.status.ok()) {
        auto parsed_state = Checkpoint::Utils::parseState(results.stored_value);

        return DbResult<ParsedState>{results.status, results.reference_count,
                                     parsed_state};
    } else {
        return DbResult<ParsedState>{results.status, results.reference_count,
                                     ParsedState()};
    }
}

SaveResults MachineStateSaver::saveMachineState(
    ParsedState state_data,
    const std::vector<unsigned char>& checkpoint_name) {
    auto serialized_state = Checkpoint::Utils::serializeState(state_data);
    return checkpoint_storage->saveValue(serialized_state, checkpoint_name);
}
