//
//  MachineStateSaver.cpp
//  avm
//
//  Created by Minh Truong on 9/30/19.
//

#include "avm/machinestatesaver.hpp"

GetResults MachineStateSaver::SaveTuple(const Tuple& val) {
    auto hash_key = GetHashKey(val);
    auto value_to_store = std::string();

    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto item = val.get_element(i);
        auto serialized_value = SerializeValue(item);

        switch (serialized_value.type) {
            case TUPLE: {
                value_to_store += serialized_value.string_value;
                auto tup = nonstd::get<Tuple>(item);
                auto results = SaveValue(tup);

                if (!results.status.ok()) {
                    // log
                }
            }
            case NUM: {
                value_to_store += serialized_value.string_value;
            }
            case CODEPT: {
                value_to_store += serialized_value.string_value;
            }
            case HASH_ONLY: {
                // huh? error
            }
        }
    }

    auto save_results = storage.SaveValueToDb(value_to_store, hash_key);

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

// rocksdb::Status MachineStateSaver::SaveMachineState(
//                                                    std::string
//                                                    checkpoint_name, const
//                                                    Tuple& tuple,
//                                                    std::vector<unsigned char>
//                                                    state_data) {
//    auto tuple_save_results = SaveValue(tuple);
//
//    auto serialized_state =
//    SerializeMachineData(tuple_save_results.storage_key, state_data);
//    auto state_save_results =
//    storage.SaveKeyValuePair(serialized_state, checkpoint_name);
//
//    return state_save_results;
//}

GetResults MachineStateSaver::SaveStringValue(
    const std::string value,
    const std::vector<unsigned char> key) {
    return storage.SaveValueToDb(value, key);
}
