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
#include <avm/machinestate/value/codepoint.hpp>
#include <avm/machinestate/value/tuple.hpp>

MachineStateSaver::MachineStateSaver(CheckpointStorage* storage,
                                     TuplePool* _pool) {
    checkpoint_storage = storage;
    pool = _pool;
}

SaveResults MachineStateSaver::SaveValue(const value& val) {
    SaveResults save_results;
    auto serialized_value = StateSaverUtils::serializeValue(val);

    if (serialized_value.type == TUPLE_TYPE) {
        auto tuple = nonstd::get<Tuple>(val);
        save_results = SaveTuple(tuple);

        if (!save_results.status.ok()) {
            // log
        }
    } else {
        auto hash_key = GetHashKey(val);
        save_results = checkpoint_storage->saveValue(
            serialized_value.string_value, hash_key);
    }
    return save_results;
}

SaveResults MachineStateSaver::SaveTuple(const Tuple& val) {
    auto hash_key = GetHashKey(val);
    auto results = checkpoint_storage->getStoredValue(hash_key);

    if (results.status.ok() && results.reference_count > 0) {
        return checkpoint_storage->incrementReference(hash_key);
    } else {
        auto value_type = (unsigned char)TUPLE_TYPE;
        std::vector<unsigned char> value_to_store;
        value_to_store.push_back(value_type);

        for (uint64_t i = 0; i < val.tuple_size(); i++) {
            auto current_val = val.get_element(i);
            auto serialized_value =
                StateSaverUtils::serializeValue(current_val);

            value_to_store.insert(value_to_store.end(),
                                  std::begin(serialized_value.string_value),
                                  std::end(serialized_value.string_value));

            if (serialized_value.type == TUPLE_TYPE) {
                auto tuple_save_results =
                    SaveTuple(nonstd::get<Tuple>(current_val));

                if (!tuple_save_results.status.ok()) {
                    // error
                }
            }
        }

        std::string val_str(value_to_store.begin(), value_to_store.end());
        auto save_results = checkpoint_storage->saveValue(val_str, hash_key);

        return save_results;
    }
};

ValueResult MachineStateSaver::getValue(std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);
    auto iter = results.stored_value.begin();

    switch ((valueTypes)*iter) {
        case TUPLE_TYPE: {
            auto tuple_res = getTuple(hash_key);
            return ValueResult{tuple_res.status, tuple_res.reference_count,
                               tuple_res.tuple};
        }
        case NUM_TYPE: {
            std::vector<unsigned char> data_vector(
                std::begin(results.stored_value),
                std::end(results.stored_value));
            data_vector.erase(data_vector.begin());
            auto val = StateSaverUtils::deserializeCheckpoint256(data_vector);
            return ValueResult{results.status, results.reference_count, val};
        }
        case CODEPT_TYPE: {
            std::vector<unsigned char> data_vector(
                std::begin(results.stored_value),
                std::end(results.stored_value));
            auto val =
                StateSaverUtils::deserializeCheckpointCodePt(data_vector);
            return ValueResult{results.status, results.reference_count, val};
        }
    }
}

TupleResult MachineStateSaver::getTuple(std::vector<unsigned char> hash_key) {
    std::vector<value> values;
    auto results = checkpoint_storage->getStoredValue(hash_key);

    if (results.status.ok()) {
        std::vector<unsigned char> data_vector(results.stored_value.begin(),
                                               results.stored_value.end());
        auto value_vectors = StateSaverUtils::parseSerializedTuple(data_vector);

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
                        auto num = StateSaverUtils::deserializeCheckpoint256(
                            current_vector);
                        values.push_back(num);
                        break;
                    }
                    case CODEPT_TYPE: {
                        auto codept =
                            StateSaverUtils::deserializeCheckpointCodePt(
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

MachineStateFetchedData MachineStateSaver::GetMachineStateData(
    std::string checkpoint_name) {
    std::vector<unsigned char> name_vector(checkpoint_name.begin(),
                                           checkpoint_name.end());

    auto machine_state_results =
        checkpoint_storage->getStoredValue(name_vector);

    if (machine_state_results.status.ok()) {
        std::vector<unsigned char> stored_state(
            std::begin(machine_state_results.stored_value),
            std::end(machine_state_results.stored_value));

        auto parsed_state = StateSaverUtils::parseCheckpointState(stored_state);
        return deserializeCheckpointState(parsed_state);
    } else {
        // return variant failure?
        return MachineStateFetchedData();
    }
}

DeleteResults MachineStateSaver::DeleteCheckpoint(std::string checkpoint_name) {
    std::vector<unsigned char> name_vector(std::begin(checkpoint_name),
                                           std::end(checkpoint_name));

    auto machine_state_results =
        checkpoint_storage->getStoredValue(name_vector);

    if (machine_state_results.status.ok()) {
        std::vector<unsigned char> stored_state(
            std::begin(machine_state_results.stored_value),
            std::end(machine_state_results.stored_value));

        auto parsed_state = StateSaverUtils::parseCheckpointState(stored_state);

        auto delete_static_res = deleteValue(parsed_state.static_val_key);
        auto delete_register_res = deleteValue(parsed_state.register_val_key);
        auto delete_cp_key = deleteValue(parsed_state.pc_key);
        auto delete_datastack_res = deleteTuple(parsed_state.datastack_key);
        auto delete_auxstack_res = deleteTuple(parsed_state.auxstack_key);
        auto delete_inbox_res = deleteTuple(parsed_state.inbox_key);
        auto delete_pendinginbox_res = deleteTuple(parsed_state.pending_key);

        if (delete_static_res.status.ok() && delete_register_res.status.ok() &&
            delete_cp_key.status.ok() && delete_datastack_res.status.ok() &&
            delete_auxstack_res.status.ok() && delete_inbox_res.status.ok() &&
            delete_pendinginbox_res.status.ok()) {
            return checkpoint_storage->deleteStoredValue(name_vector);
        } else {
            // undo parshal delete?
            // make these things atomic?
        }
    } else {
        return DeleteResults{0, machine_state_results.status};
    }
}

SaveResults MachineStateSaver::SaveMachineState(
    MachineStateStorageData state_data,
    std::string checkpoint_name) {
    auto serialized_state = serializeState(state_data);
    std::vector<unsigned char> checkpoint_name_vector(
        std::begin(checkpoint_name), std::end(checkpoint_name));

    return checkpoint_storage->saveValue(
        std::string(serialized_state.begin(), serialized_state.end()),
        checkpoint_name_vector);
}

// private
// -------------------------------------------------------------------------------

DeleteResults MachineStateSaver::deleteValue(
    std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);

    if (results.status.ok()) {
        auto type = (valueTypes)results.stored_value[0];

        if (type == TUPLE_TYPE) {
            return deleteTuple(hash_key);
        } else {
            return checkpoint_storage->deleteStoredValue(hash_key);
        }
    } else {
        return DeleteResults{0, rocksdb::Status().NotFound()};
    }
}

DeleteResults MachineStateSaver::deleteTuple(
    std::vector<unsigned char> hash_key) {
    // reduce extra db calls
    auto results = checkpoint_storage->getStoredValue(hash_key);

    if (results.status.ok()) {
        if (results.reference_count == 1) {
            std::vector<unsigned char> data_vector(results.stored_value.begin(),
                                                   results.stored_value.end());
            auto value_vectors =
                StateSaverUtils::parseSerializedTuple(data_vector);

            for (auto& vector : value_vectors) {
                if ((valueTypes)vector[0] == TUPLE_TYPE) {
                    auto delete_stat = deleteTuple(std::vector<unsigned char>(
                        vector.begin() + 1, vector.end()));
                }
            }
        }
        return checkpoint_storage->deleteStoredValue(hash_key);
    } else {
        return DeleteResults{0, rocksdb::Status().NotFound()};
    }
}

CodePoint MachineStateSaver::getCodePoint(std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);
    std::vector<unsigned char> code_pt_vector(std::begin(results.stored_value),
                                              std::end(results.stored_value));
    auto code_point =
        StateSaverUtils::deserializeCheckpointCodePt(code_pt_vector);

    return code_point;
}

uint256_t MachineStateSaver::getInt256(std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);
    std::vector<unsigned char> uint256_vector(std::begin(results.stored_value),
                                              std::end(results.stored_value));
    auto num256 = StateSaverUtils::deserializeCheckpoint256(uint256_vector);

    return num256;
}

std::vector<unsigned char> MachineStateSaver::serializeState(
    MachineStateStorageData state_data) {
    std::vector<unsigned char> state_data_vector;
    state_data_vector.push_back(state_data.status_char);

    state_data_vector.insert(state_data_vector.end(),
                             state_data.blockreason_str.begin(),
                             state_data.blockreason_str.end());

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
    auto register_val_results = getValue(stored_state.auxstack_key);
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
