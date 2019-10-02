//
//  MachineStateSaver.cpp
//  avm
//
//  Created by Minh Truong on 9/30/19.
//

#include "avm/machinestatesaver.hpp"
#include <avm/codepoint.hpp>
#include <avm/tokenTracker.hpp>
#include <avm/tuple.hpp>

void MachineStateSaver::setStorage(CheckpointStorage* storage) {
    checkpoint_storage = storage;
}

GetResults MachineStateSaver::SaveValue(const value& val) {
    GetResults save_results;
    auto serialized_value = SerializeValue(val);

    if (serialized_value.type == TUPLE) {
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

// can store somes values directly if too slow
GetResults MachineStateSaver::SaveTuple(const Tuple& val) {
    std::vector<unsigned char> value_to_store;

    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto current_val = val.get_element(i);
        auto serialized_value = SerializeValue(val);

        value_to_store.insert(value_to_store.end(),
                              std::begin(serialized_value.string_value),
                              std::end(serialized_value.string_value));

        if (serialized_value.type == TUPLE) {
            auto tuple_save_results =
                SaveTuple(nonstd::get<Tuple>(current_val));

            if (!tuple_save_results.status.ok()) {
                // error
            }
        }
    }

    std::string val_str(value_to_store.begin(), value_to_store.end());

    auto hash_key = GetHashKey(val);
    auto save_results = checkpoint_storage->saveValue(val_str, hash_key);

    return save_results;
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

value MachineStateSaver::getValue(std::vector<unsigned char> hash_key) {
    auto results = checkpoint_storage->getStoredValue(hash_key);
    auto iter = results.stored_value.begin();

    switch ((types)*iter) {
        case TUPLE: {
            return getTuple(hash_key);
        }
        case NUM: {
            std::vector<unsigned char> data_vector(
                std::begin(results.stored_value),
                std::end(results.stored_value));
            auto val = deserializeCheckpoint256(data_vector);
            return val;
        }
        case CODEPT: {
            std::vector<unsigned char> data_vector(
                std::begin(results.stored_value),
                std::end(results.stored_value));
            auto val = deserializeCheckpointCodePt(data_vector);
            return val;
        }
        case HASH_ONLY: {
            // error?
        }
    }
}

Tuple MachineStateSaver::getTuple(std::vector<unsigned char> hash_key) {
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
                auto tup = getTuple(tup_hash);
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

    auto tup = Tuple(values, pool);

    return tup;
};

// make sure correct
std::vector<std::vector<unsigned char>> MachineStateSaver::breakIntoValues(
    std::vector<unsigned char> data_vecgtor) {
    std::vector<std::vector<unsigned char>> return_vector;

    auto it = data_vecgtor.begin();

    while (it != data_vecgtor.end()) {
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

std::string SerializeMachineData(std::vector<unsigned char> tuple_key,
                                 std::vector<unsigned char> state_data) {
    tuple_key.insert(tuple_key.end(), state_data.begin(), state_data.end());
    std::string str(tuple_key.begin(), tuple_key.end());

    return str;
}

std::vector<unsigned char> MachineStateSaver::serializeState(
    MachineStateStorageData state_data) {
    std::vector<unsigned char> state_data_vector;
    state_data_vector.push_back(state_data.status_str);
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

MachineStateFetchedData MachineStateSaver::GetMachineStateData(
    std::string checkpoint_name) {
    std::vector<unsigned char> name_vector(checkpoint_name.begin(),
                                           checkpoint_name.end());
    auto machine_state_results = GetStringValue(name_vector);

    std::vector<unsigned char> stored_state(
        std::begin(machine_state_results.stored_value),
        std::end(machine_state_results.stored_value));

    auto return_state_data = deserializeState(stored_state);
    return return_state_data;
}

MachineStateFetchedData MachineStateSaver::deserializeState(
    // status
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
    auto static_val_results = getValue(static_val);
    iter += 33;
    std::vector<unsigned char> register_val(iter, iter + 33);
    auto register_val_ressults = getValue(register_val);
    iter += 33;
    std::vector<unsigned char> datastack(iter, iter + 33);
    auto datastack_results = getTuple(datastack);
    iter += 33;
    std::vector<unsigned char> auxstack(iter, iter + 33);
    auto auxstack_results = getTuple(auxstack);
    iter += 33;
    std::vector<unsigned char> inbox(iter, iter + 33);
    auto inbox_results = getTuple(inbox);
    iter += 33;
    std::vector<unsigned char> pending(iter, iter + 33);
    iter += 33;
    auto pending_results = getTuple(pending);
    std::vector<unsigned char> pc(iter, iter + 33);
    auto pc_results = getCodePoint(pc);

    return MachineStateFetchedData{static_val_results, register_val_ressults,
                                   datastack_results,  auxstack_results,
                                   inbox_results,      pending_results,
                                   pc_results,         status,
                                   blockreason_vector, balance_track_vector};
}

GetResults MachineStateSaver::SaveMachineState(
    MachineStateStorageData state_data,
    std::string checkpoint_name) {
    auto serialized_state = serializeState(state_data);
    std::vector<unsigned char> checkpoint_name_vector(
        std::begin(checkpoint_name), std::end(checkpoint_name));

    return SaveStringValue(
        std::string(serialized_state.begin(), serialized_state.end()),
        checkpoint_name_vector);
}

GetResults MachineStateSaver::SaveStringValue(
    const std::string value,
    const std::vector<unsigned char> key) {
    return checkpoint_storage->saveValue(value, key);
}

GetResults MachineStateSaver::GetStringValue(
    const std::vector<unsigned char> key) {
    return checkpoint_storage->getStoredValue(key);
}
