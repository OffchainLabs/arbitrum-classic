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

#include "config.hpp"
#include "helper.hpp"

#include <avm/machinestate/machinestate.hpp>
#include <avm_values/vmValueParser.hpp>
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/machinestatedeleter.hpp>
#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/checkpoint/machinestatesaver.hpp>

#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

void saveValue(Transaction& transaction,
               const value& val,
               int expected_ref_count,
               bool expected_status) {
    auto results = saveValue(transaction, val);
    transaction.commit();
    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void getValue(MachineStateFetcher& fetcher,
              std::vector<unsigned char>& hash_key,
              int expected_ref_count,
              uint256_t& expected_hash,
              ValueTypes expected_value_type,
              bool expected_status) {
    auto results = fetcher.getValue(hash_key);
    auto serialized_val = checkpoint::utils::serializeValue(results.data);
    auto type = (ValueTypes)serialized_val[0];

    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(type == expected_value_type);
    REQUIRE(hash(results.data) == expected_hash);
}

void saveTuple(Transaction& transaction,
               const Tuple& tup,
               int expected_ref_count,
               bool expected_status) {
    auto results = saveValue(transaction, tup);
    transaction.commit();
    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void getTuple(MachineStateFetcher& fetcher,
              std::vector<unsigned char>& hash_key,
              int expected_ref_count,
              uint256_t& expected_hash,
              int expected_tuple_size,
              bool expected_status) {
    auto results = fetcher.getValue(hash_key);

    REQUIRE(nonstd::holds_alternative<Tuple>(results.data));

    auto tuple = nonstd::get<Tuple>(results.data);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(tuple.calculateHash() == expected_hash);
    REQUIRE(tuple.tuple_size() == expected_tuple_size);
    REQUIRE(results.status.ok() == expected_status);
}

void getTupleValues(MachineStateFetcher& fetcher,
                    std::vector<unsigned char>& hash_key,
                    std::vector<uint256_t> value_hashes) {
    auto results = fetcher.getValue(hash_key);
    REQUIRE(nonstd::holds_alternative<Tuple>(results.data));

    auto tuple = nonstd::get<Tuple>(results.data);
    REQUIRE(tuple.tuple_size() == value_hashes.size());

    for (size_t i = 0; i < value_hashes.size(); i++) {
        REQUIRE(hash(tuple.get_element(i)) == value_hashes[i]);
    }
}

TEST_CASE("Save value") {
    DBDeleter deleter;
    CheckpointStorage storage(dbpath, test_contract_path);
    auto transaction = storage.makeTransaction();

    SECTION("save 1 num tuple") {
        TuplePool pool;
        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        saveValue(*transaction, tuple, 1, true);
    }
    SECTION("save num") {
        uint256_t num = 1;
        saveValue(*transaction, num, 1, true);
    }
    SECTION("save codepoint") {
        auto code_point = CodePoint(1, Operation(), 0);
        saveValue(*transaction, code_point, 1, true);
    }
}

TEST_CASE("Save tuple") {
    DBDeleter deleter;
    CheckpointStorage storage(dbpath, test_contract_path);
    auto transaction = storage.makeTransaction();

    TuplePool pool;

    SECTION("save 1 num tuple") {
        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        saveTuple(*transaction, tuple, 1, true);
    }
    SECTION("save 2, 1 num tuples") {
        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        saveTuple(*transaction, tuple, 1, true);
        saveTuple(*transaction, tuple, 2, true);
    }
    SECTION("saved tuple in tuple") {
        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);
        saveTuple(*transaction, tuple, 1, true);
        saveTuple(*transaction, tuple, 2, true);
    }
}

TEST_CASE("Save and get value") {
    SECTION("save empty tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        auto tuple = Tuple();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        saveValue(*transaction, tuple, 1, true);
        getValue(fetcher, hash_key, 1, tup_hash, TUPLE, true);
    }
    SECTION("save tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        TuplePool pool;
        auto tuple = Tuple(num, &pool);
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        saveValue(*transaction, tuple, 1, true);
        getValue(fetcher, hash_key, 1, tup_hash, TUPLE, true);
    }
    SECTION("save num") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto hash_key = GetHashKey(num);
        auto num_hash = hash(num);

        saveValue(*transaction, num, 1, true);
        getValue(fetcher, hash_key, 1, num_hash, NUM, true);
    }
    SECTION("save codepoint") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        auto hash_key = GetHashKey(code_point);
        auto cp_hash = hash(code_point);
        saveValue(*transaction, code_point, 1, true);
        getValue(fetcher, hash_key, 1, cp_hash, CODEPT, true);
    }
    SECTION("save err codepoint") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto code_point = getErrCodePoint();

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        auto hash_key = GetHashKey(code_point);
        auto cp_hash = hash(code_point);
        saveValue(*transaction, code_point, 1, true);
        getValue(fetcher, hash_key, 1, cp_hash, CODEPT, true);
    }
}

TEST_CASE("Save and get tuple values") {
    SECTION("save num tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        TuplePool pool;
        auto tuple = Tuple(num, &pool);

        saveTuple(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(num)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    SECTION("save codepoint tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        CodePoint code_point = storage.getInitialVmValues().code[1];

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        TuplePool pool;
        auto tuple = Tuple(code_point, &pool);

        saveTuple(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    SECTION("save codepoint tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        TuplePool pool;
        auto tuple = Tuple(code_point, &pool);

        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    SECTION("save nested tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        auto inner_tuple = Tuple();
        TuplePool pool;
        auto tuple = Tuple(inner_tuple, &pool);

        saveTuple(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    SECTION("save multiple valued tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);
        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        auto inner_tuple = Tuple();
        uint256_t num = 1;
        auto tuple = Tuple(inner_tuple, num, code_point, &pool);

        saveTuple(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple), hash(num),
                                      hash(code_point)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
    SECTION("save multiple valued tuple, saveValue()") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[2];

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        auto inner_tuple = Tuple();
        uint256_t num = 1;
        auto tuple = Tuple(inner_tuple, num, code_point, &pool);

        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple), hash(num),
                                      hash(code_point)};
        auto hash_key = GetHashKey(tuple);

        getTupleValues(fetcher, hash_key, hashes);
    }
}

TEST_CASE("Save And Get Tuple") {
    SECTION("save 1 num tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(*transaction, tuple, 1, true);
        getTuple(fetcher, hash_key, 1, tup_hash, 1, true);
    }
    SECTION("save codepoint in tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto code_point = storage.getInitialVmValues().code[0];

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        auto tuple = Tuple(code_point, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(*transaction, tuple, 1, true);
        getTuple(fetcher, hash_key, 1, tup_hash, 1, true);
    }
    SECTION("save 1 num tuple twice") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto transaction2 = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(*transaction, tuple, 1, true);
        saveTuple(*transaction2, tuple, 2, true);
        getTuple(fetcher, hash_key, 2, tup_hash, 1, true);
    }
    SECTION("save 2 num tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);
        std::vector<CodePoint> code;
        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        uint256_t num2 = 2;
        auto tuple = Tuple(num, num2, &pool);
        auto tup_hash = tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);

        saveTuple(*transaction, tuple, 1, true);
        getTuple(fetcher, hash_key, 1, tup_hash, 2, true);
    }
    SECTION("save tuple in tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);
        saveTuple(*transaction, tuple, 1, true);

        auto inner_hash_key = GetHashKey(inner_tuple);
        auto inner_tup_hash = inner_tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        getTuple(fetcher, hash_key, 1, tup_hash, 1, true);
        getTuple(fetcher, inner_hash_key, 1, inner_tup_hash, 1, true);
    }
    SECTION("save 2 tuples in tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        uint256_t num2 = 2;
        auto inner_tuple2 = Tuple(num2, &pool);
        auto tuple = Tuple(inner_tuple, inner_tuple2, &pool);
        saveTuple(*transaction, tuple, 1, true);

        auto inner_hash_key = GetHashKey(inner_tuple);
        auto inner_tup_hash = inner_tuple.calculateHash();
        auto inner_hash_key2 = GetHashKey(inner_tuple2);
        auto inner_tup_hash2 = inner_tuple2.calculateHash();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        getTuple(fetcher, hash_key, 1, tup_hash, 2, true);
        getTuple(fetcher, inner_hash_key, 1, inner_tup_hash, 1, true);
        getTuple(fetcher, inner_hash_key2, 1, inner_tup_hash2, 1, true);
    }
    SECTION("save saved tuple in tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto transaction2 = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);
        auto inner_hash_key = GetHashKey(inner_tuple);
        auto inner_tup_hash = inner_tuple.calculateHash();
        auto hash_key = GetHashKey(tuple);
        auto tup_hash = tuple.calculateHash();

        saveTuple(*transaction, inner_tuple, 1, true);
        getTuple(fetcher, inner_hash_key, 1, inner_tup_hash, 1, true);
        saveTuple(*transaction2, tuple, 1, true);
        getTuple(fetcher, hash_key, 1, tup_hash, 1, true);
        getTuple(fetcher, inner_hash_key, 2, inner_tup_hash, 1, true);
    }
}

void saveState(Transaction& transaction,
               MachineStateKeys storage_data,
               std::vector<unsigned char> checkpoint_name) {
    auto results = saveMachineState(transaction, storage_data, checkpoint_name);
    auto status = transaction.commit();

    REQUIRE(results.reference_count == 1);
    REQUIRE(results.status.ok());
}

void getSavedState(MachineStateFetcher& fetcher,
                   std::vector<unsigned char> checkpoint_name,
                   MachineStateKeys expected_data,
                   int expected_ref_count,
                   std::vector<std::vector<unsigned char>> keys) {
    auto results = fetcher.getMachineState(checkpoint_name);

    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == expected_ref_count);

    auto data = results.data;

    REQUIRE(data.status_char == expected_data.status_char);
    REQUIRE(data.pc_key == expected_data.pc_key);
    REQUIRE(data.datastack_key == expected_data.datastack_key);
    REQUIRE(data.auxstack_key == expected_data.auxstack_key);
    REQUIRE(data.register_val_key == expected_data.register_val_key);

    for (auto& key : keys) {
        auto res = fetcher.getValue(key);
        REQUIRE(res.status.ok());
    }
}

void deleteCheckpoint(CheckpointStorage& storage,
                      MachineStateFetcher& fetcher,
                      std::vector<unsigned char> checkpoint_name,
                      std::vector<std::vector<unsigned char>> deleted_values) {
    auto res = deleteCheckpoint(storage, checkpoint_name);
    auto results = fetcher.getMachineState(checkpoint_name);
    REQUIRE(results.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = fetcher.getValue(hash_key);
        REQUIRE(res.status.ok() == false);
    }
}

void deleteCheckpointSavedTwice(
    CheckpointStorage& storage,
    MachineStateFetcher& fetcher,
    std::vector<unsigned char> checkpoint_name,
    std::vector<std::vector<unsigned char>> deleted_values) {
    auto res = deleteCheckpoint(storage, checkpoint_name);
    auto res2 = deleteCheckpoint(storage, checkpoint_name);
    auto results = fetcher.getMachineState(checkpoint_name);

    REQUIRE(results.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = fetcher.getValue(hash_key);
        REQUIRE(res.status.ok() == false);
    }
}

void deleteCheckpointSavedTwiceReordered(
    CheckpointStorage& storage,
    MachineStateFetcher& fetcher,
    std::vector<unsigned char> checkpoint_name,
    std::vector<std::vector<unsigned char>> deleted_values) {
    auto resultsx = fetcher.getMachineState(checkpoint_name);
    for (auto& hash_key : deleted_values) {
        auto res = fetcher.getValue(hash_key);
        REQUIRE(res.status.ok());
    }
    auto res = deleteCheckpoint(storage, checkpoint_name);
    auto results = fetcher.getMachineState(checkpoint_name);
    REQUIRE(results.status.ok() == true);

    for (auto& hash_key : deleted_values) {
        auto res = fetcher.getValue(hash_key);
        REQUIRE(res.status.ok());
    }
    auto res2 = deleteCheckpoint(storage, checkpoint_name);
    auto results2 = fetcher.getMachineState(checkpoint_name);
    REQUIRE(results2.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = fetcher.getValue(hash_key);
        REQUIRE(res.status.ok() == false);
    }
}

MachineStateKeys makeStorageData(Transaction& transaction,
                                 value registerVal,
                                 Datastack stack,
                                 Datastack auxstack,
                                 Status state,
                                 CodePoint pc,
                                 CodePoint err_pc,
                                 BlockReason blockReason) {
    TuplePool pool;

    auto datastack_results = stack.checkpointState(transaction, &pool);
    auto auxstack_results = auxstack.checkpointState(transaction, &pool);

    auto register_val_results = saveValue(transaction, registerVal);
    auto pc_results = saveValue(transaction, pc);
    auto err_pc_results = saveValue(transaction, err_pc);

    auto status_str = (unsigned char)state;

    return MachineStateKeys{
        register_val_results.storage_key, datastack_results.storage_key,
        auxstack_results.storage_key,     pc_results.storage_key,
        err_pc_results.storage_key,       status_str};
}

MachineStateKeys getStateValues(Transaction& transaction) {
    TuplePool pool;
    uint256_t register_val = 100;
    auto static_val = Tuple(register_val, Tuple(), &pool);

    auto code_point = CodePoint(1, Operation(), 0);
    auto tup1 = Tuple(register_val, &pool);
    auto tup2 = Tuple(code_point, tup1, &pool);

    Datastack data_stack;
    data_stack.push(register_val);
    Datastack aux_stack;
    aux_stack.push(register_val);
    aux_stack.push(code_point);

    CodePoint pc_codepoint(0, Operation(), 0);
    CodePoint err_pc_codepoint(0, Operation(), 0);
    Status state = Status::Extensive;

    auto inbox_blocked = InboxBlocked(0);

    auto saved_data =
        makeStorageData(transaction, register_val, data_stack, aux_stack, state,
                        pc_codepoint, err_pc_codepoint, inbox_blocked);

    return saved_data;
}

MachineStateKeys getDefaultValues(Transaction& transaction) {
    TuplePool pool;
    auto register_val = Tuple();
    auto data_stack = Tuple();
    auto aux_stack = Tuple();

    Status state = Status::Extensive;
    CodePoint code_point(0, Operation(), 0);

    auto data = makeStorageData(transaction, Tuple(), Datastack(), Datastack(),
                                state, code_point, code_point, NotBlocked());

    return data;
}

std::vector<std::vector<unsigned char>> getHashKeys(MachineStateKeys data) {
    std::vector<std::vector<unsigned char>> hash_keys;

    hash_keys.push_back(data.auxstack_key);
    hash_keys.push_back(data.datastack_key);
    hash_keys.push_back(data.pc_key);
    hash_keys.push_back(data.err_pc_key);
    hash_keys.push_back(data.register_val_key);

    return hash_keys;
}

TEST_CASE("Save Machinestatedata") {
    SECTION("default") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto transaction = storage.makeTransaction();
        auto data_values = getDefaultValues(*transaction);
        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveState(*transaction, data_values, checkpoint_key);
    }
    SECTION("with values") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];
        CodePoint code_point2 = storage.getInitialVmValues().code[1];

        auto transaction = storage.makeTransaction();
        auto state_data = getStateValues(*transaction);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveState(*transaction, state_data, checkpoint_key);
    }
}

TEST_CASE("Get Machinestate data") {
    SECTION("default") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        auto data_values = getDefaultValues(*transaction);
        auto keys = getHashKeys(data_values);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveMachineState(*transaction, data_values, checkpoint_key);
        transaction->commit();
        getSavedState(fetcher, checkpoint_key, data_values, 1, keys);
    }
    SECTION("with values") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];
        CodePoint code_point2 = storage.getInitialVmValues().code[1];

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        auto state_data = getStateValues(*transaction);
        auto keys = getHashKeys(state_data);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveState(*transaction, state_data, checkpoint_key);
        getSavedState(fetcher, checkpoint_key, state_data, 1, keys);
    }
}

TEST_CASE("Delete checkpoint") {
    SECTION("default") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);
        auto data_values = getDefaultValues(*transaction);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveMachineState(*transaction, data_values, checkpoint_key);
        transaction->commit();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpoint(storage, fetcher, checkpoint_key, hash_keys);
    }
    SECTION("with actual state values") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];
        CodePoint code_point2 = storage.getInitialVmValues().code[2];

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);
        auto data_values = getStateValues(*transaction);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveMachineState(*transaction, data_values, checkpoint_key);
        transaction->commit();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpoint(storage, fetcher, checkpoint_key, hash_keys);
    }
    SECTION("delete checkpoint saved twice") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];
        CodePoint code_point2 = storage.getInitialVmValues().code[1];

        auto fetcher = MachineStateFetcher(storage);
        auto transaction = storage.makeTransaction();
        auto data_values = getStateValues(*transaction);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveMachineState(*transaction, data_values, checkpoint_key);
        saveMachineState(*transaction, data_values, checkpoint_key);
        transaction->commit();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpointSavedTwice(storage, fetcher, checkpoint_key, hash_keys);
    }
    SECTION("delete checkpoint saved twice, reordered") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePoint code_point = storage.getInitialVmValues().code[0];
        CodePoint code_point2 = storage.getInitialVmValues().code[1];

        auto fetcher = MachineStateFetcher(storage);
        auto transaction = storage.makeTransaction();
        auto transaction2 = storage.makeTransaction();
        auto data_values = getStateValues(*transaction);

        std::vector<unsigned char> checkpoint_key = {'k', 'e', 'y'};

        saveMachineState(*transaction, data_values, checkpoint_key);
        transaction->commit();
        saveMachineState(*transaction2, data_values, checkpoint_key);
        transaction2->commit();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpointSavedTwiceReordered(storage, fetcher, checkpoint_key,
                                            hash_keys);
    }
}
