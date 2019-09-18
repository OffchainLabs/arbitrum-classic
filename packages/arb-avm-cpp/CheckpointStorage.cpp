//
//  CheckpointStorage.cpp
//  avm
//
//  Created by Minh Truong on 9/12/19.
//

#include "CheckpointStorage.hpp"
#include "CheckpointDataLayer.hpp"
#include "avm/tuple.hpp"
#include "avm/value.hpp"
#include "rocksdb/db.h"

rocksdb::Status CheckpointStorage::SaveMachineCode(
    std::vector<CodePoint> code) {
    auto code_to_save = CheckpointDataLayer::ConvertMachineCode(code);
    auto status = SaveValue(code_to_save, machine_code_key);

    return status;
};

rocksdb::Status CheckpointStorage::SaveValue(value& val, TuplePool& pool) {
    // nonstd::visit([&](const auto& val) { return SaveValue(val); }, val);
    auto value_type = GetType(val);
    auto hash_key = hash(val);

    std::vector<unsigned char> value_vector;
    marshal_value(val, value_vector);
    auto value_buffer = reinterpret_cast<char*>(value_vector.data());
    rocksdb::Status status;

    switch (value_type) {
        case TUPLE: {
            auto tup_size = get_tuple_size(value_buffer);
            auto tuple = deserialize_tuple(value_buffer, tup_size, pool);
            status = SaveValue(tuple, pool);
        }
        case NUM: {
            auto num = deserialize_int(value_buffer);
            status = SaveValue(num);
        }
        case CODEPT: {
            auto codepint = deserializeCodePoint(value_buffer, pool);
            status = SaveValue(codepint);
        }
        default:
            // error
    }

    return status;
};

struct {
    TuplePool& pool;
    CheckpointStorage cp_storage;

    rocksdb::Status operator()(const Tuple& value) {
        return SaveValue(value, pool);
    }

    rocksdb::Status operator()(const uint256_t value) {
        return SaveValue(value);
    }

    rocksdb::Status operator()(const CodePoint& value) {
        return SaveValue(value);
    }
}

rocksdb::Status
CheckpointStorage::SaveValue(Tuple& val, TuplePool& pool) {
    auto hash_key = hash(val);

    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash_key, hash_key_vector);

    auto ref_key = MakeReferenceKey(hash_key_vector);
    auto key = MakeKey(hash_key_vector);

    auto ref_count_value = GetValue(txn_db, ref_key);

    if (ref_count_value.empty()) {
        // restriction on tuple size?
        for (uint64_t i = 0; i < val.tuple_size(); i++) {
            auto item = val.get_element(i);
            auto save_status = SaveValue(item, pool);

            if (save_status.ok()) {
                auto ref_value = std::to_string(1);
                auto status = SaveValue(ref_value, ref_key);
            }
        }
    } else {
        auto ref_count = std::stoi(ref_count_value);
        ref_count += 1;
        SaveValue(std::to_string(ref_count), ref_key);
    }
};

rocksdb::Status CheckpointStorage::SaveValue(uint256_t& val) {
    auto hash_key = hash(val);

    std::vector<unsigned char> hash_key_vector;
    std::vector<unsigned char> value_vector;
    marshal_value(hash_key, hash_key_vector);
    marshal_value(val, value_vector);

    std::string value_str(value_vector.begin(), value_vector.end());
    auto ref_key = MakeReferenceKey(hash_key_vector);
    auto key = MakeKey(hash_key_vector);

    auto ref_count_value = GetValue(txn_db, ref_key);

    if (ref_count_value.empty()) {
        // make all one transaction, currently bad
        SaveValue(value_str, key);
        auto ref_value = std::to_string(1);
        SaveValue(ref_value, ref_key);
    } else {
        auto ref_count = std::stoi(ref_count_value);
        ref_count += 1;
        SaveValue(std::to_string(ref_count), ref_key);
    }
};

rocksdb::Status CheckpointStorage::SaveValue(CodePoint& val) {
    auto hash_key = hash(val);

    std::vector<unsigned char> hash_key_vector;
    marshal_value(val, hash_key_vector);

    auto ref_key = MakeReferenceKey(hash_key_vector);
    auto key = MakeKey(hash_key_vector);

    std::vector<unsigned char> pc_value_vector;
    pc_value_vector.push_back(val.pc);

    std::string pc_value_to_store(pc_value_vector.begin(),
                                  pc_value_vector.end());

    auto ref_count_value = GetValue(txn_db, ref_key);
    if (ref_count_value.empty()) {
        // make all one transaction, currently bad
        SaveValue(pc_value_to_store, key);
        auto ref_value = std::to_string(1);
        SaveValue(ref_value, ref_key);
    } else {
        auto ref_count = std::stoi(ref_count_value);
        ref_count += 1;
        SaveValue(std::to_string(ref_count), ref_key);
    }
};

// private function
// --------------------------------------------------------------------

std::string CheckpointStorage::MakeReferenceKey(
    std::vector<unsigned char>& key) {
    key.push_back((unsigned char)0);

    std::string str(key.begin(), key.end());
    return str;
}

std::string CheckpointStorage::MakeKey(std::vector<unsigned char>& key) {
    key.push_back((unsigned char)1);

    std::string str(key.begin(), key.end());
    return str;
}

rocksdb::Status CheckpointStorage::UpdateReferenceCount(
    std::vector<unsigned char>& hash_key,
    int count) {
    auto ref_key = MakeReferenceKey(hash_key);
    auto status = SaveValue(std::to_string(count), ref_key);

    return status;
};

int CheckpointStorage::GetReferenceCount(std::vector<unsigned char>& hash_key) {
    auto ref_key = MakeReferenceKey(hash_key);
    auto ref_count_value = GetValue(txn_db, ref_key);

    if (ref_count_value.empty()) {
        return 0;
    } else {
        return std::stoi(ref_count_value);
    }
};

rocksdb::Status CheckpointStorage::SaveValAndUpdateRef(
    std::vector<unsigned char>& hash_key,
    std::string value) {
    auto ref_count = GetReferenceCount(hash_key);

    if (ref_count == 0) {
        // make sure this is correct
        rocksdb::WriteOptions writeOptions;
        rocksdb::Transaction* transaction =
            txn_db->BeginTransaction(writeOptions);
        assert(transaction);

        var

    } else {
        ref_count += 1;
        auto status = UpdateReferenceCount(hash_key, ref_count);
    }
};

rocksdb::Status CheckpointStorage::SaveValue(std::string val, std::string key) {
    // make sure this is correct
    rocksdb::WriteOptions writeOptions;
    rocksdb::Transaction* transaction = txn_db->BeginTransaction(writeOptions);
    assert(transaction);

    auto put_status = transaction->Put(key, val);
    assert(put_status.ok());

    auto commit_status = transaction->Commit();
    assert(commit_status.ok());

    return commit_status;
};

// verify transaction will be updated without having to return
rocksdb::Status PutInTransaction(rocksdb::Transaction* transaction,
                                 std::string key,
                                 std::string value){

};

// use variant to return status error or value
std::string CheckpointStorage::GetValue(rocksdb::TransactionDB* txn_db,
                                        std::string key) {
    rocksdb::ReadOptions read_options;
    std::string return_value;

    auto get_status = txn_db->Get(read_options, key, &return_value);
    if (get_status.ok()) {
        return return_value;
    } else {
        return std::string();
    }
}

rocksdb::Status CheckpointStorage::DeleteValue(rocksdb::TransactionDB* txn_db,
                                               std::string key) {
    rocksdb::WriteOptions writeOptions;
    rocksdb::Transaction* transaction = txn_db->BeginTransaction(writeOptions);
    assert(transaction);

    auto delete_status = transaction->Delete(key);
    assert(delete_status.ok());

    auto commit_status = transaction->Commit();
    assert(commit_status.ok());

    return commit_status;
}
