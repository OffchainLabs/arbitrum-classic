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

void MachineStateSaver::setStorage(CheckpointStorage* storage,
                                   TuplePool* _pool) {
    checkpoint_storage = storage;
    _pool = pool;
}

SaveResults MachineStateSaver::SaveValue(const value& val) {
    SaveResults save_results;
    auto serialized_value = SerializeValue(val);

    if (serialized_value.type == TUPLE_TYPE) {
        auto tuple = nonstd::get<Tuple>(val);
        auto save_results = SaveTuple(tuple);

        if (!save_results.status.ok()) {
            // log
        }
    } else {
        auto hash_key = GetHashKey(val);
        auto save_results = checkpoint_storage->saveValue(
            serialized_value.string_value, hash_key);
    }

    return save_results;
}

DeleteResults MachineStateSaver::Delete(Tuple& tuple) {
    auto hash_key = GetHashKey(tuple);
    auto results = checkpoint_storage->getStoredValue(hash_key);

    if (results.status.ok()) {
        if (results.reference_count == 1) {
            for (uint64_t i = 0; i < tuple.tuple_size(); i++) {
                auto current_val = tuple.get_element(i);
                // doesnt need to serialize, just figure out if tuple
                auto serialized_value = SerializeValue(current_val);

                if (serialized_value.type == TUPLE_TYPE) {
                    auto del_res = Delete(nonstd::get<Tuple>(current_val));
                }
            }
        }

        return checkpoint_storage->deleteStoredValue(hash_key);
    } else {
        return DeleteResults{0, rocksdb::Status().NotFound()};
    }
}

SaveResults MachineStateSaver::SaveTuple(const Tuple& val) {
    auto hash_key = GetHashKey(val);
    auto results = checkpoint_storage->getStoredValue(hash_key);

    if (results.status.ok() && results.reference_count > 0) {
        return checkpoint_storage->incrementReference(hash_key);
    } else {
        std::vector<unsigned char> value_to_store;

        for (uint64_t i = 0; i < val.tuple_size(); i++) {
            auto current_val = val.get_element(i);
            auto serialized_value = SerializeValue(current_val);

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

CodePoint MachineStateSaver::getCodePoint(std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);
    std::vector<unsigned char> code_pt_vector(std::begin(results.stored_value),
                                              std::end(results.stored_value));
    auto code_point = deserializeCheckpointCodePt(code_pt_vector);

    return code_point;
}

uint256_t MachineStateSaver::getInt256(std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);
    std::vector<unsigned char> uint256_vector(std::begin(results.stored_value),
                                              std::end(results.stored_value));
    auto num256 = deserializeCheckpoint256(uint256_vector);

    return num256;
}

ValueResult MachineStateSaver::getValue(std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);
    auto iter = results.stored_value.begin();

    switch ((valueTypes)*iter) {
        case TUPLE_TYPE: {
            auto tuple_res = getTuple(hash_key);
            return ValueResult{tuple_res.reference_count, tuple_res.tuple};
        }
        case NUM_TYPE: {
            std::vector<unsigned char> data_vector(
                std::begin(results.stored_value),
                std::end(results.stored_value));
            auto val = deserializeCheckpoint256(data_vector);
            return ValueResult{results.reference_count, val};
        }
        case CODEPT_TYPE: {
            std::vector<unsigned char> data_vector(
                std::begin(results.stored_value),
                std::end(results.stored_value));
            auto val = deserializeCheckpointCodePt(data_vector);
            return ValueResult{results.reference_count, val};
        }
    }
}

TupleResult MachineStateSaver::getTuple(std::vector<unsigned char> hash_key) {
    std::vector<value> values;

    auto results = checkpoint_storage->getStoredValue(hash_key);

    std::vector<unsigned char> data_vector(results.stored_value.begin(),
                                           results.stored_value.end());
    auto value_vectors = breakIntoValues(data_vector);

    for (auto& current_vector : value_vectors) {
        switch (current_vector[0]) {
            case TUPLE: {
                std::vector<unsigned char> tup_hash(current_vector.begin() + 1,
                                                    current_vector.end());
                auto tup = getTuple(tup_hash).tuple;
                values.push_back(tup);
            }
            case NUM: {
                auto num = deserializeCheckpoint256(current_vector);
                values.push_back(num);
            }
            case CODEPT: {
                auto codept = deserializeCheckpointCodePt(current_vector);
                values.push_back(codept);
            }
            case HASH_ONLY: {
                // error?
            }
        }
    }

    auto tuple = Tuple(values, pool);
    return TupleResult{results.reference_count, tuple};
};

// make sure correct
std::vector<std::vector<unsigned char>> MachineStateSaver::breakIntoValues(
    std::vector<unsigned char> data_vector) {
    std::vector<std::vector<unsigned char>> return_vector;

    auto it = data_vector.begin();

    while (it != data_vector.end()) {
        auto val = *it;
        std::vector<unsigned char> current;

        switch (val) {
            case TUPLE: {
                current.insert(current.end(), it, it + 34);
                it += 34;
            }
            case NUM: {
                current.insert(current.end(), it, it + 34);
                it += 34;
            }
            case CODEPT: {
                current.insert(current.end(), it, it + 9);
                it += 9;
            }
            default: {
            }
        }

        return_vector.push_back(current);
    }

    return return_vector;
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

    state_data_vector.insert(state_data_vector.end(),
                             state_data.inbox_results.storage_key.begin(),
                             state_data.inbox_results.storage_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.pending_results.storage_key.begin(),
                             state_data.pending_results.storage_key.end());

    state_data_vector.insert(state_data_vector.end(),
                             state_data.pc_results.storage_key.begin(),
                             state_data.pc_results.storage_key.end());
    return state_data_vector;
}

DeleteResults MachineStateSaver::DeleteCheckpoint(std::string checkpoint_name) {
    std::vector<unsigned char> name_vector(std::begin(checkpoint_name),
                                           std::end(checkpoint_name));
    auto machine_state_results = GetStringValue(name_vector);

    if (machine_state_results.status.ok()) {
        std::vector<unsigned char> stored_state(
            std::begin(machine_state_results.stored_value),
            std::end(machine_state_results.stored_value));

        auto parsed_state = parseCheckpointState(stored_state);
        auto delete_static_res =
            checkpoint_storage->deleteStoredValue(parsed_state.static_val_key);
        auto delete_register_res = checkpoint_storage->deleteStoredValue(
            parsed_state.register_val_key);
        auto delete_cp_key =
            checkpoint_storage->deleteStoredValue(parsed_state.pc_key);

        return checkpoint_storage->deleteStoredValue(name_vector);

    } else {
    }
}

MachineStateFetchedData MachineStateSaver::GetMachineStateData(
    std::string checkpoint_name) {
    std::vector<unsigned char> name_vector(checkpoint_name.begin(),
                                           checkpoint_name.end());
    auto machine_state_results = GetStringValue(name_vector);

    if (machine_state_results.status.ok()) {
        std::vector<unsigned char> stored_state(
            std::begin(machine_state_results.stored_value),
            std::end(machine_state_results.stored_value));

        auto parsed_state = parseCheckpointState(stored_state);
        return deserializeCheckpointState(parsed_state);
    } else {
        // return variant failure?
    }
}

ParsedCheckpointState parseCheckpointState(
    std::vector<unsigned char> stored_state) {
    auto iter = stored_state.begin();
    auto status = (unsigned char)(*iter);
    iter += 1;

    // blockreason
    auto block_type = (BlockType)*iter;
    auto length_of_block_reason = blockreason_type_length[block_type];
    std::vector<unsigned char> blockreason_vector(
        iter, iter + length_of_block_reason);

    iter += length_of_block_reason;

    // balancetracker
    unsigned int balance_tracker_length;
    memcpy(&balance_tracker_length, &(*iter), sizeof(balance_tracker_length));
    iter += sizeof(unsigned int);

    auto total_len = 54 * balance_tracker_length;
    std::vector<unsigned char> balance_track_vector(iter, iter + total_len);
    iter += total_len;

    // staticval
    std::vector<unsigned char> static_val(iter, iter + 33);
    iter += 33;
    std::vector<unsigned char> register_val(iter, iter + 33);
    iter += 33;
    std::vector<unsigned char> datastack(iter, iter + 33);
    iter += 33;
    std::vector<unsigned char> auxstack(iter, iter + 33);
    iter += 33;
    std::vector<unsigned char> inbox(iter, iter + 33);
    iter += 33;
    std::vector<unsigned char> pending(iter, iter + 33);
    iter += 33;
    std::vector<unsigned char> pc(iter, iter + 33);

    return ParsedCheckpointState{static_val,
                                 register_val,
                                 datastack,
                                 auxstack,
                                 inbox,
                                 pending,
                                 pc,
                                 status,
                                 blockreason_vector,
                                 balance_track_vector};
}

MachineStateFetchedData MachineStateSaver::deserializeCheckpointState(
    // status
    ParsedCheckpointState stored_state) {
    // staticval
    auto static_val_results = getValue(stored_state.static_val_key);
    auto register_val_ressults = getValue(stored_state.auxstack_key);
    auto datastack_results = getTuple(stored_state.datastack_key);
    auto auxstack_results = getTuple(stored_state.auxstack_key);
    auto inbox_results = getTuple(stored_state.inbox_key);
    auto pending_results = getTuple(stored_state.pending_key);
    auto pc_results = getCodePoint(stored_state.pc_key);

    return MachineStateFetchedData{static_val_results.val,
                                   register_val_ressults.val,
                                   datastack_results.tuple,
                                   auxstack_results.tuple,
                                   inbox_results.tuple,
                                   pending_results.tuple,
                                   pc_results,
                                   stored_state.status_char,
                                   stored_state.blockreason_str,
                                   stored_state.balancetracker_str};
}

SaveResults MachineStateSaver::SaveMachineState(
    MachineStateStorageData state_data,
    std::string checkpoint_name) {
    auto serialized_state = serializeState(state_data);
    std::vector<unsigned char> checkpoint_name_vector(
        std::begin(checkpoint_name), std::end(checkpoint_name));

    return SaveStringValue(
        std::string(serialized_state.begin(), serialized_state.end()),
        checkpoint_name_vector);
}

SaveResults MachineStateSaver::SaveStringValue(
    const std::string value,
    const std::vector<unsigned char> key) {
    return checkpoint_storage->saveValue(value, key);
}

GetResults MachineStateSaver::GetStringValue(
    const std::vector<unsigned char> key) {
    return checkpoint_storage->getStoredValue(key);
}
