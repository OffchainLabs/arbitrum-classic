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

GetResults CheckpointStorage::GetMachineState(std::string machine_name) {
    auto machine_state_results = GetValue(machine_name);
    auto data_hash = machine_state_results.stored_value;
    auto data_results = GetValue(data_hash);

    return
}

GetResults CheckpointStorage::SaveMachineState(
    std::string checkpoint_name,
    const Tuple& tuple,
    std::vector<unsigned char> state_data) {
    auto tuple_save_status = SaveValue(tuple);
    auto value_key = GetHashKey(tuple);

    auto state_save_status = SaveValueToDb(state_data, checkpoint_name);

    auto map_status = SaveValueToDb(value_key, checkpoint_name);

    return map_status;
}

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

GetResults CheckpointStorage::SaveValueToDb(std::string val, std::string key) {
    auto results = GetValue(key);
    auto ref_count = results.reference_count;
    auto value = results.stored_value;

    if (!results.status.ok() && ref_count < 1) {
        value = val;
        ref_count = 1;
    } else {
        ref_count += 1;
    }

    auto updated_value = SerializeCountAndValue(ref_count, value);

    // make sure this is correct
    rocksdb::WriteOptions writeOptions;
    rocksdb::Transaction* transaction = txn_db->BeginTransaction(writeOptions);
    assert(transaction);

    auto put_status = transaction->Put(key, updated_value);
    assert(put_status.ok());

    auto commit_status = transaction->Commit();
    assert(commit_status.ok());

    if (commit_status.ok()) {
        GetResults save_results{ref_count, commit_status, key, val};

        return save_results;
    } else {
        auto unsuccessful = rocksdb::Status().NotFound();
        GetResults save_results{--ref_count, unsuccessful, key, val};

        // log
    }
};

// use variant to return status error or value
GetResults CheckpointStorage::GetValue(std::string hash_key) {
    rocksdb::ReadOptions read_options;
    std::string return_value;

    auto get_status = txn_db->Get(read_options, hash_key, &return_value);

    if (get_status.ok()) {
        auto tuple = ParseCountAndValue(return_value);

        GetResults results{std::get<0>(tuple), get_status, hash_key,
                           std::get<1>(tuple)};

        return results;
    } else {
        // make sure this is correct
        auto unsuccessful = rocksdb::Status().NotFound();
        GetResults results{0, unsuccessful, std::string(), std::string()};

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

std::string CheckpointStorage::GetHashKey(const value& val) {
    auto hash_key = hash(val);

    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash_key, hash_key_vector);

    return std::string(hash_key_vector.begin(), hash_key_vector.end());
}

std::tuple<int, std::string> ParseCountAndValue(std::string string_value) {
    // is max 256 references good enough?
    const char* c_string = string_value.c_str();
    auto ref_count = (int)c_string[0];

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
