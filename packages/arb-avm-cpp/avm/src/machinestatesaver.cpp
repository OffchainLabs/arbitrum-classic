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

#include "avm/machinestate/machinestatesaver.hpp"
#include <avm/machinestate/tokenTracker.hpp>
#include <avm/value/codepoint.hpp>
#include <avm/value/tuple.hpp>
#include <variant>

MachineStateSaver::MachineStateSaver(CheckpointStorage* storage,
                                     TuplePool* _pool) {
    checkpoint_storage = storage;
    pool = _pool;
}

SaveResults MachineStateSaver::saveValue(const value& val) {
    auto serialized_value = StateSaverUtils::serializeValue(val);
    auto type = (valueTypes)serialized_value[0];

    if (type == TUPLE_TYPE) {
        auto tuple = nonstd::get<Tuple>(val);
        return saveTuple(tuple);
    } else {
        auto hash_key = GetHashKey(val);
        return checkpoint_storage->saveValue(serialized_value, hash_key);
    }
}

SaveResults MachineStateSaver::saveTuple(const Tuple& val) {
    auto hash_key = GetHashKey(val);
    auto results = checkpoint_storage->getStoredValue(hash_key);

    auto incr_ref_count = results.status.ok() && results.reference_count > 0;

    if (incr_ref_count) {
        return checkpoint_storage->incrementReference(hash_key);
    } else {
        std::vector<unsigned char> value_vector{(unsigned char)TUPLE_TYPE};

        for (uint64_t i = 0; i < val.tuple_size(); i++) {
            auto current_val = val.get_element(i);
            auto serialized_val = StateSaverUtils::serializeValue(current_val);

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

ValueResult MachineStateSaver::getValue(std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);

    if (results.status.ok()) {
        auto value_type = (valueTypes)results.stored_value[0];
        results.stored_value.erase(results.stored_value.begin());

        switch (value_type) {
            case TUPLE_TYPE: {
                auto tuple_res = getTuple(hash_key);
                return ValueResult{tuple_res.status, tuple_res.reference_count,
                                   tuple_res.tuple};
            }
            case NUM_TYPE: {
                auto val =
                    StateSaverUtils::deserializeUint256(results.stored_value);
                return ValueResult{results.status, results.reference_count,
                                   val};
            }
            case CODEPT_TYPE: {
                auto val =
                    StateSaverUtils::deserializeCodepoint(results.stored_value);
                return ValueResult{results.status, results.reference_count,
                                   val};
            }
        }
    } else {
        auto error_res = ValueResult();
        error_res.status = results.status;
        error_res.reference_count = results.reference_count;
        return error_res;
    }
}

TupleResult MachineStateSaver::getTuple(std::vector<unsigned char> hash_key) {
    std::vector<value> values;
    auto results = checkpoint_storage->getStoredValue(hash_key);

    if (results.status.ok()) {
        auto value_vectors =
            StateSaverUtils::parseSerializedTuple(results.stored_value);

        if (value_vectors.empty()) {
            return TupleResult{results.status, results.reference_count,
                               Tuple()};
        } else {
            for (auto& current_vector : value_vectors) {
                auto value_type = (valueTypes)current_vector[0];
                current_vector.erase(current_vector.begin());

                switch (value_type) {
                    case TUPLE_TYPE: {
                        auto tuple = getTuple(current_vector).tuple;
                        values.push_back(tuple);
                        break;
                    }
                    case NUM_TYPE: {
                        auto num =
                            StateSaverUtils::deserializeUint256(current_vector);
                        values.push_back(num);
                        break;
                    }
                    case CODEPT_TYPE: {
                        auto codept = StateSaverUtils::deserializeCodepoint(
                            current_vector);
                        values.push_back(codept);
                        break;
                    }
                }
            }
            auto tuple = Tuple(values, pool);
            return TupleResult{results.status, results.reference_count, tuple};
        }
    } else {
        return TupleResult{results.status, results.reference_count, Tuple()};
    }
};

StateResult MachineStateSaver::getMachineStateData(
    std::string checkpoint_name) {
    std::vector<unsigned char> name_vector(checkpoint_name.begin(),
                                           checkpoint_name.end());

    auto results = checkpoint_storage->getStoredValue(name_vector);

    if (results.status.ok()) {
        auto parsed_state =
            StateSaverUtils::parseCheckpointState(results.stored_value);

        return StateResult{results.status, results.reference_count,
                           deserializeCheckpointState(parsed_state)};
    } else {
        return StateResult{results.status, results.reference_count,
                           MachineStateFetchedData()};
    }
}

DeleteResults MachineStateSaver::deleteCheckpoint(std::string checkpoint_name) {
    std::vector<unsigned char> name_vector(checkpoint_name.begin(),
                                           checkpoint_name.end());

    auto results = checkpoint_storage->getStoredValue(name_vector);

    if (results.status.ok()) {
        auto parsed_state =
            StateSaverUtils::parseCheckpointState(results.stored_value);

        auto delete_static_res = deleteValue(parsed_state.static_val_key);
        auto delete_register_res = deleteValue(parsed_state.register_val_key);
        auto delete_cp_key = deleteValue(parsed_state.pc_key);
        auto delete_datastack_res = deleteTuple(parsed_state.datastack_key);
        auto delete_auxstack_res = deleteTuple(parsed_state.auxstack_key);
        auto delete_inbox_res = deleteTuple(parsed_state.inbox_key);
        auto delete_inbox_count = deleteValue(parsed_state.inbox_count_key);
        auto delete_pendinginbox_res = deleteTuple(parsed_state.pending_key);
        auto delete_pending_count = deleteValue(parsed_state.pending_count_key);

        if (delete_static_res.status.ok() && delete_register_res.status.ok() &&
            delete_cp_key.status.ok() && delete_datastack_res.status.ok() &&
            delete_auxstack_res.status.ok() && delete_inbox_res.status.ok() &&
            delete_pendinginbox_res.status.ok() &&
            delete_inbox_count.status.ok() &&
            delete_pending_count.status.ok()) {
        }
        return checkpoint_storage->deleteStoredValue(name_vector);
    } else {
        return DeleteResults{0, results.status};
    }
}

SaveResults MachineStateSaver::saveMachineState(
    MachineStateStorageData state_data,
    std::string checkpoint_name) {
    auto serialized_state = serializeState(state_data);
    std::vector<unsigned char> name_vector(checkpoint_name.begin(),
                                           checkpoint_name.end());

    return checkpoint_storage->saveValue(serialized_state, name_vector);
}

// private ------------------------------------------------------------------

DeleteResults MachineStateSaver::deleteValue(
    std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);

    if (results.status.ok()) {
        auto type = (valueTypes)results.stored_value[0];

        if (type == TUPLE_TYPE) {
            return deleteTuple(hash_key, results);
        } else {
            return checkpoint_storage->deleteStoredValue(hash_key);
        }
    } else {
        return DeleteResults{0, results.status};
    }
}

DeleteResults MachineStateSaver::deleteTuple(
    std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);

    return deleteTuple(hash_key, results);
}

DeleteResults MachineStateSaver::deleteTuple(
    std::vector<unsigned char> hash_key,
    GetResults& results) {
    if (results.status.ok()) {
        if (results.reference_count == 1) {
            auto value_vectors =
                StateSaverUtils::parseSerializedTuple(results.stored_value);

            for (auto& vector : value_vectors) {
                if ((valueTypes)vector[0] == TUPLE_TYPE) {
                    vector.erase(vector.begin());
                    auto delete_stat = deleteTuple(vector);
                }
            }
        }
        return checkpoint_storage->deleteStoredValue(hash_key);
    } else {
        return DeleteResults{0, results.status};
    }
}

CodePoint MachineStateSaver::getCodePoint(std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);
    return StateSaverUtils::deserializeCodepoint(results.stored_value);
}

uint256_t MachineStateSaver::getInt256(std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);
    return StateSaverUtils::deserializeUint256(results.stored_value);
}

std::vector<unsigned char> MachineStateSaver::serializeState(
    MachineStateStorageData state_data) {
    std::vector<unsigned char> state_data_vector;
    state_data_vector.push_back(state_data.status_char);

    state_data_vector.insert(state_data_vector.end(),
                             state_data.blockreason_str.begin(),
                             state_data.blockreason_str.end());

    unsigned int tracker_length = state_data.balancetracker_str.size();
    std::vector<unsigned char> tracker_len_vector(sizeof(tracker_length));
    memcpy(&tracker_len_vector[0], &tracker_length, sizeof(tracker_length));

    state_data_vector.insert(state_data_vector.end(),
                             tracker_len_vector.begin(),
                             tracker_len_vector.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.balancetracker_str.begin(),
                             state_data.balancetracker_str.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.static_val_results.storage_key.begin(),
                             state_data.static_val_results.storage_key.end());
    state_data_vector.insert(
        state_data_vector.end(),
        state_data.register_val_results.storage_key.begin(),
        state_data.register_val_results.storage_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.datastack_results.storage_key.begin(),
                             state_data.datastack_results.storage_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.auxstack_results.storage_key.begin(),
                             state_data.auxstack_results.storage_key.end());

    state_data_vector.insert(
        state_data_vector.end(),
        state_data.inbox_messages_results.storage_key.begin(),
        state_data.inbox_messages_results.storage_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.inbox_count_results.storage_key.begin(),
                             state_data.inbox_count_results.storage_key.end());

    state_data_vector.insert(
        state_data_vector.end(),
        state_data.pending_messages_results.storage_key.begin(),
        state_data.pending_messages_results.storage_key.end());

    state_data_vector.insert(
        state_data_vector.end(),
        state_data.pending_count_results.storage_key.begin(),
        state_data.pending_count_results.storage_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.pc_results.storage_key.begin(),
                             state_data.pc_results.storage_key.end());
    return state_data_vector;
}

MachineStateFetchedData MachineStateSaver::deserializeCheckpointState(
    ParsedCheckpointState stored_state) {
    auto static_val_results = getValue(stored_state.static_val_key);
    auto register_val_results = getValue(stored_state.register_val_key);
    auto datastack_results = getTuple(stored_state.datastack_key);
    auto auxstack_results = getTuple(stored_state.auxstack_key);
    auto inbox_results = getTuple(stored_state.inbox_key);
    auto inbox_count = getValue(stored_state.inbox_count_key);
    auto pending_results = getTuple(stored_state.pending_key);
    auto pending_count = getValue(stored_state.pending_count_key);
    auto pc_results = getCodePoint(stored_state.pc_key);

    return MachineStateFetchedData{static_val_results.val,
                                   register_val_results.val,
                                   datastack_results.tuple,
                                   auxstack_results.tuple,
                                   inbox_results.tuple,
                                   inbox_count.val,
                                   pending_results.tuple,
                                   pending_count.val,
                                   pc_results,
                                   stored_state.status_char,
                                   stored_state.blockreason_str,
                                   stored_state.balancetracker_str};
}
