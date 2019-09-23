//
//  checkpointdatalayer.cpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#include "avm/checkpointdatalayer.hpp"
#include <vector>
#include "avm/checkpointutils.hpp"
#include "avm/tuple.hpp"
#include "avm/value.hpp"

using UCharVec = std::vector<unsigned char>;

std::string dbPath = "tmp/rocksDbPath";
std::string machine_code_key = "machine_code";

bool CheckpointDataLayer::Intialize() {
    rocksdb::Options options;
    rocksdb::TransactionDBOptions txn_options;
    options.create_if_missing = true;

    auto status =
        rocksdb::TransactionDB::Open(options, txn_options, dbPath, &txn_db);

    return status.ok();
};

void CheckpointDataLayer::Close() {
    delete txn_db;
}

rocksdb::Status CheckpointDataLayer::GetMachineState(std::string machine_name) {
    auto hash_results = GetValueAndCount(machine_name);
    auto data_hash = hash_results.result_value;
    auto data_results = GetValueAndCount(data_hash);
}

rocksdb::Status CheckpointDataLayer::SaveValueAndMapToKey(
    const Tuple& val,
    std::string hash_key) {
    auto status = SaveValue(val);
    auto value_key = GetHashKey(val);
    auto map_status = SaveValue(value_key, hash_key);

    return map_status;
}

rocksdb::Status CheckpointDataLayer::SaveValue(const Tuple& val) {
    // somtime it says no conversion
    auto hash_key = GetHashKey(val);
    auto value_to_store = std::string();

    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto item = val.get_element(i);
        auto serialized_value = SerializeValue(item);

        switch (serialized_value.type) {
            case TUPLE: {
                value_to_store += serialized_value.string_value;
                auto tup = nonstd::get<Tuple>(item);
                auto status = SaveValue(tup);

                if (!status.ok()) {
                }
            }
            case NUM: {
                value_to_store += serialized_value.string_value;
            }
            case CODEPT: {
                value_to_store += serialized_value.string_value;
            }
        }
    }

    auto save_status = SaveValue(value_to_store, hash_key);

    return save_status;
};

rocksdb::Status CheckpointDataLayer::SaveValue(std::string val,
                                               std::string key) {
    auto results = GetValueAndCount(key);
    auto ref_count = results.reference_count;
    auto value = results.result_value;

    if (ref_count < 1) {
        value = val;
    }
    ref_count += 1;

    auto updated_value = SerializeCountAndValue(ref_count, value);

    // make sure this is correct
    rocksdb::WriteOptions writeOptions;
    rocksdb::Transaction* transaction = txn_db->BeginTransaction(writeOptions);
    assert(transaction);

    auto put_status = transaction->Put(key, updated_value);
    assert(put_status.ok());

    auto commit_status = transaction->Commit();
    assert(commit_status.ok());

    return commit_status;
};

// use variant to return status error or value
GetResults CheckpointDataLayer::GetValueAndCount(std::string hash_key) {
    rocksdb::ReadOptions read_options;
    std::string return_value;

    auto get_status = txn_db->Get(read_options, hash_key, &return_value);

    if (get_status.ok()) {
        return ParseCountAndValue(return_value);
    } else {
        GetResults results{0, std::string()};

        return results;
    }
}

rocksdb::Status CheckpointDataLayer::DeleteValue(std::string key) {
    rocksdb::WriteOptions writeOptions;
    rocksdb::Transaction* transaction = txn_db->BeginTransaction(writeOptions);
    assert(transaction);

    auto delete_status = transaction->Delete(key);
    assert(delete_status.ok());

    auto commit_status = transaction->Commit();
    assert(commit_status.ok());

    return commit_status;
}

std::string CheckpointDataLayer::GetHashKey(const value& val) {
    auto hash_key = hash(val);

    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash_key, hash_key_vector);

    return std::string(hash_key_vector.begin(), hash_key_vector.end());
}

GetResults ParseCountAndValue(std::string string_value) {
    // is max 256 references good enough?
    const char* c_string = string_value.c_str();
    auto ref_count = (int)c_string[0];

    // skips exactly the first char(byte) in order to extract value saved?
    auto saved_value = string_value.substr(1, string_value.size() - 1);

    GetResults results{ref_count, saved_value};

    return results;
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
