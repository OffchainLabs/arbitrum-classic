//
//  checkpointstorage.cpp
//  avm
//
//  Created by Minh Truong on 9/23/19.
//

#include "avm/checkpointstorage.hpp"
#include <vector>
#include "avm/checkpointutils.hpp"
#include "avm/tuple.hpp"
#include "avm/value.hpp"

using UCharVec = std::vector<unsigned char>;

std::string dbPath = "tmp/rocksDbPath";
std::string machine_code_key = "machine_code";

bool CheckpointStorage::Intialize() {
    rocksdb::Options options;
    rocksdb::TransactionDBOptions txn_options;
    options.create_if_missing = true;

    auto status =
        rocksdb::TransactionDB::Open(options, txn_options, dbPath, &txn_db);

    return status.ok();
};

void CheckpointStorage::Close() {
    delete txn_db;
}

std::string SerializeMachineData(std::vector<unsigned char> tuple_key,
                                 std::vector<unsigned char> state_data) {
    tuple_key.insert(tuple_key.end(), state_data.begin(), state_data.end());
    std::string str(tuple_key.begin(), tuple_key.end());

    return str;
}

std::tuple<Tuple, SerializedStateData> CheckpointStorage::GetMachineState(
    std::string checkpoint_name) {
    std::vector<unsigned char> name_vector(checkpoint_name.begin(),
                                           checkpoint_name.end());
    auto machine_state_results = GetValue(name_vector);

    std::vector<unsigned char> tuple_hash(
        machine_state_results.stored_value.begin(),
        machine_state_results.stored_value.begin() + 33);
    std::vector<unsigned char> state_data(
        machine_state_results.stored_value.begin() + 33,
        machine_state_results.stored_value.end());

    auto state_data_object = Deserialize(state_data);
    auto tup = GetTuple(tuple_hash);

    return std::make_tuple(tup, state_data_object);
}

rocksdb::Status CheckpointStorage::SaveMachineState(
    std::string checkpoint_name,
    const Tuple& tuple,
    std::vector<unsigned char> state_data) {
    auto tuple_save_results = SaveValue(tuple);

    auto serialized_state =
        SerializeMachineData(tuple_save_results.storage_key, state_data);
    auto state_save_results =
        SaveKeyValuePair(serialized_state, checkpoint_name);

    return state_save_results;
}

std::vector<std::vector<unsigned char>> breakIntoValues(
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

Tuple CheckpointStorage::GetTuple(std::vector<unsigned char> hash_key) {
    std::vector<value> values;

    auto results = GetValue(hash_key);

    std::vector<unsigned char> data_vector(results.stored_value.begin(),
                                           results.stored_value.end());
    auto value_vectors = breakIntoValues(data_vector);

    for (auto& vec : value_vectors) {
        auto it = vec.begin() + 1;
        std::vector<unsigned char> current(it, vec.end());

        switch (vec[0]) {
            case TUPLE: {
                auto tup = GetTuple(current);
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

    // get pool
    TuplePool* pool;
    auto tup = Tuple(values, pool);

    return tup;
};

GetResults CheckpointStorage::SaveValue(const Tuple& val) {
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

    auto save_results = SaveValueToDb(value_to_store, hash_key);

    return save_results;
};

rocksdb::Status CheckpointStorage::SaveKeyValuePair(std::string value,
                                                    std::string key) {
    rocksdb::WriteOptions writeOptions;
    rocksdb::Transaction* transaction = txn_db->BeginTransaction(writeOptions);
    assert(transaction);

    auto put_status = transaction->Put(key, value);
    assert(put_status.ok());

    auto commit_status = transaction->Commit();
    assert(commit_status.ok());

    return commit_status;
}

GetResults CheckpointStorage::SaveValueToDb(
    std::string val,
    std::vector<unsigned char> hash_key) {
    auto results = GetValue(hash_key);
    auto ref_count = results.reference_count;
    auto value = results.stored_value;

    if (!results.status.ok() || ref_count < 1) {
        value = val;
        ref_count = 1;
    } else {
        ref_count += 1;
    }

    auto updated_value = SerializeCountAndValue(ref_count, value);

    std::string key_str(hash_key.begin(), hash_key.end());

    auto commit_status = SaveKeyValuePair(updated_value, key_str);
    assert(commit_status.ok());

    if (commit_status.ok()) {
        GetResults save_results{ref_count, commit_status, hash_key, val};

        return save_results;
    } else {
        auto unsuccessful = rocksdb::Status().NotFound();
        GetResults save_results{--ref_count, unsuccessful, hash_key, val};

        // log
        return save_results;
    }
};

// use variant to return status error or value
GetResults CheckpointStorage::GetValue(std::vector<unsigned char> hash_key) {
    rocksdb::ReadOptions read_options;
    std::string return_value;

    std::string key_str(hash_key.begin(), hash_key.end());
    auto get_status = txn_db->Get(read_options, key_str, &return_value);

    if (get_status.ok()) {
        auto tuple = ParseCountAndValue(return_value);

        GetResults results{std::get<0>(tuple), get_status, hash_key,
                           std::get<1>(tuple)};

        return results;
    } else {
        // make sure this is correct
        auto unsuccessful = rocksdb::Status().NotFound();
        GetResults results{0, unsuccessful, std::vector<unsigned char>(),
                           std::string()};

        return results;
    }
}

rocksdb::Status CheckpointStorage::DeleteValue(std::string key) {
    rocksdb::WriteOptions writeOptions;
    rocksdb::Transaction* transaction = txn_db->BeginTransaction(writeOptions);
    assert(transaction);

    auto delete_status = transaction->Delete(key);
    assert(delete_status.ok());

    auto commit_status = transaction->Commit();
    assert(commit_status.ok());

    return commit_status;
}

std::vector<unsigned char> CheckpointStorage::GetHashKey(const value& val) {
    auto hash_key = hash(val);
    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash_key, hash_key_vector);

    return hash_key_vector;
}

std::tuple<int, std::string> ParseCountAndValue(std::string string_value) {
    // is max 256 references good enough?
    const char* c_string = string_value.c_str();
    //    auto ref_count = (int)c_string[0];
    //    auto ref_count = static_cast<int>(c_string[0]);
    uint16_t ref_count;
    memcpy(&ref_count, c_string, sizeof(ref_count));
    //    auto ref_count = *reinterpret_cast<const int *>(&c_string[0]);

    // skips exactly the first char(byte) in order to extract value saved?
    auto saved_value = string_value.substr(1, string_value.size() - 1);

    return std::make_tuple(ref_count, saved_value);
}

std::string SerializeCountAndValue(int count, std::string value) {
    std::string str;
    if (count > 255) {
        // error
    } else {
        auto count_as_char = (char)count;
        // does not replace
        value.insert(value.begin(), count_as_char);
    }

    return str;
}
