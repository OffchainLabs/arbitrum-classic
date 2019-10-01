//
//  MachineStateSaver.cpp
//  avm
//
//  Created by Minh Truong on 9/30/19.
//

#include "avm/machinestatesaver.hpp"
#include <avm/codepoint.hpp>
#include <avm/tuple.hpp>

GetResults MachineStateSaver::SaveValue(const value& val) {
    auto serialized_value = SerializeValue(val);
    auto hash_key = GetHashKey(val);

    if (serialized_value.type == TUPLE) {
        auto tup = nonstd::get<Tuple>(val);
        auto results = SaveValue(tup);

        if (!results.status.ok()) {
            // log
        }
    }

    auto save_results =
        storage.SaveValueToDb(serialized_value.string_value, hash_key);

    return save_results;
}

// can store somes values directly
GetResults MachineStateSaver::SaveTuple(const Tuple& val) {
    auto hash_key = GetHashKey(val);
    std::vector<unsigned char> value_to_store;

    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto item = val.get_element(i);

        auto saved_val_result = SaveValue(item);

        if (!saved_val_result.status.ok()) {
            // log
        }

        value_to_store.insert(value_to_store.end(),
                              std::begin(saved_val_result.stored_value),
                              std::end(saved_val_result.stored_value));
    }

    std::string val_str(value_to_store.begin(), value_to_store.end());
    auto save_results = storage.SaveValueToDb(val_str, hash_key);

    return save_results;
};

Tuple MachineStateSaver::getTuple(std::vector<unsigned char> hash_key) {
    std::vector<value> values;

    auto results = storage.getStoredValue(hash_key);

    std::vector<unsigned char> data_vector(results.stored_value.begin(),
                                           results.stored_value.end());
    auto value_vectors = breakIntoValues(data_vector);

    for (auto& vec : value_vectors) {
        auto it = vec.begin() + 1;
        std::vector<unsigned char> current(it, vec.end());

        switch (vec[0]) {
            case TUPLE: {
                auto tup = getTuple(current);
                values.push_back(tup);
            }
            case NUM: {
                auto buff = reinterpret_cast<char*>(&current[0]);
                auto num = deserialize_int(buff);
                values.push_back(num);
            }
            case CODEPT: {
                // not the rest?
                auto buff = reinterpret_cast<char*>(&current[0]);
                auto pc = deserialize_int64(buff);
                auto code = CodePoint();
                code.pc = pc;
                values.push_back(code);
            }
        }
    }

    // get pool;
    auto tup = Tuple(values, pool);

    return tup;
};

std::vector<std::vector<unsigned char>> MachineStateSaver::breakIntoValues(
    std::vector<unsigned char> data_vecgtor) {
    std::vector<std::vector<unsigned char>> return_vector;

    auto it = data_vecgtor.begin();

    while (it != data_vecgtor.end()) {
        auto val = *it;
        std::vector<unsigned char> current;
        current.push_back(val);

        it++;

        switch (val) {
            case TUPLE: {
                current.insert(current.end(), it, it + 33);
                it += 33;
            }
            case NUM: {
                current.insert(current.end(), it, it + 33);
                it += 33;
            }
            case CODEPT: {
                current.insert(current.end(), it, it + 8);
                it += 8;
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
    MachineStateData state_data) {
    std::vector<unsigned char> state_data_vector;
    state_data_vector.insert(state_data_vector.end(),
                             state_data.status_str.begin(),
                             state_data.status_str.end());
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

MachineStateData MachineStateSaver::deserializeState(
    std::vector<unsigned char> stored_state) {
    auto iter = stored_state.begin();
    auto status = std::vector<unsigned char>(*iter);
    iter += 1;
}

GetResults MachineStateSaver::SaveMachineState(MachineStateData state_data,
                                               std::string checkpoint_name) {
    auto serialized_state = serializeState(state_data);
    std::vector<unsigned char> checkpoint_name_vector(
        std::begin(checkpoint_name), std::end(checkpoint_name));

    return SaveStringValue(
        std::string(serialized_state.begin(), serialized_state.end()),
        checkpoint_name_vector);
}

// MachineLoadData MachineStateSaver::GetMachineState(std::string
// checkpoint_name) {
//
//    std::vector<unsigned char> name_vector(checkpoint_name.begin(),
//    checkpoint_name.end());
//
//    auto machine_state_results = storage.getStoredValue(name_vector);
//    std::vector<unsigned char> tuple_hash(
//                                          machine_state_results.stored_value.begin(),
//                                          machine_state_results.stored_value.begin()
//                                          + 33);
//    std::vector<unsigned char> state_data(
//                                          machine_state_results.stored_value.begin()
//                                          + 33,
//                                          machine_state_results.stored_value.end());
//
//    auto state_data_object = Deserialize(state_data);
//    auto tup = getTuple(tuple_hash);
//
//    return MachineLoadData{
//        tup,
//        state_data_object
//    };
//}

GetResults MachineStateSaver::SaveStringValue(
    const std::string value,
    const std::vector<unsigned char> key) {
    return storage.SaveValueToDb(value, key);
}
