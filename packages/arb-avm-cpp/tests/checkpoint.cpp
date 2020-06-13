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
#include <data_storage/checkpoint/checkpointutils.hpp>
#include <data_storage/checkpoint/machine.hpp>
#include <data_storage/checkpoint/machinestatedeleter.hpp>
#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/checkpoint/machinestatesaver.hpp>
#include <data_storage/storageresult.hpp>

#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

using namespace checkpoint::utils;

void saveValue(Transaction& transaction,
               const value& val,
               int expected_ref_count,
               bool expected_status) {
    auto results = saveValue(transaction, val);
    transaction.commit();
    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void getValue(const Transaction& transaction,
              const value& value,
              int expected_ref_count,
              ValueTypes expected_value_type,
              bool expected_status) {
    TuplePool pool;
    auto results = getValue(transaction, hash_value(value), &pool);
    auto serialized_val = checkpoint::utils::serializeValue(results.data);
    auto type = (ValueTypes)serialized_val[0];

    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(type == expected_value_type);
    REQUIRE(hash_value(results.data) == hash_value(value));
}

void getTuple(const Transaction& transaction,
              const Tuple& tuple,
              int expected_ref_count,
              bool expected_status) {
    TuplePool pool;
    auto results = getValue(transaction, hash(tuple), &pool);

    REQUIRE(nonstd::holds_alternative<Tuple>(results.data));

    auto loadedTuple = nonstd::get<Tuple>(results.data);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(loadedTuple.calculateHash() == tuple.calculateHash());
    REQUIRE(loadedTuple.tuple_size() == tuple.tuple_size());
    REQUIRE(results.status.ok() == expected_status);
}

void getTupleValues(const Transaction& transaction,
                    uint256_t tuple_hash,
                    std::vector<uint256_t> value_hashes) {
    TuplePool pool;
    auto results = getValue(transaction, tuple_hash, &pool);
    REQUIRE(nonstd::holds_alternative<Tuple>(results.data));

    auto tuple = nonstd::get<Tuple>(results.data);
    REQUIRE(tuple.tuple_size() == value_hashes.size());

    for (size_t i = 0; i < value_hashes.size(); i++) {
        REQUIRE(hash_value(tuple.get_element(i)) == value_hashes[i]);
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
        auto code_point = CodePointStub(1, 654546);
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
        saveValue(*transaction, tuple, 1, true);
    }
    SECTION("save 2, 1 num tuples") {
        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);
        saveValue(*transaction, tuple, 1, true);
        saveValue(*transaction, tuple, 2, true);
    }
    SECTION("saved tuple in tuple") {
        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);
        saveValue(*transaction, tuple, 1, true);
        saveValue(*transaction, tuple, 2, true);
    }
}

TEST_CASE("Save and get value") {
    SECTION("save empty tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto transaction = storage.makeTransaction();

        auto tuple = Tuple();

        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
    }
    SECTION("save tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto transaction = storage.makeTransaction();

        uint256_t num = 1;
        TuplePool pool;
        auto tuple = Tuple(num, &pool);

        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
    }
    SECTION("save num") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto transaction = storage.makeTransaction();

        uint256_t num = 1;

        saveValue(*transaction, num, 1, true);
        getValue(*transaction, num, 1, NUM, true);
    }
    SECTION("save codepoint") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        CodePointStub code_point_stub(1, 654546);

        auto transaction = storage.makeTransaction();

        saveValue(*transaction, code_point_stub, 1, true);
        getValue(*transaction, code_point_stub, 1, CODEPT, true);
    }
    SECTION("save err codepoint") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        CodePointStub code_point_stub(1, 654546);

        auto transaction = storage.makeTransaction();

        saveValue(*transaction, code_point_stub, 1, true);
        getValue(*transaction, code_point_stub, 1, CODEPT, true);
    }
}

TEST_CASE("Save and get tuple values") {
    SECTION("save num tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto transaction = storage.makeTransaction();

        uint256_t num = 1;
        TuplePool pool;
        auto tuple = Tuple(num, &pool);

        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(num)};

        getTupleValues(*transaction, hash(tuple), hashes);
    }
    SECTION("save codepoint tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        CodePointStub code_point_stub(1, 654546);

        auto transaction = storage.makeTransaction();

        TuplePool pool;
        auto tuple = Tuple(code_point_stub, &pool);

        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point_stub)};

        getTupleValues(*transaction, hash(tuple), hashes);
    }
    SECTION("save codepoint tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);
        CodePointStub code_point_stub(1, 654546);

        auto transaction = storage.makeTransaction();

        TuplePool pool;
        auto tuple = Tuple(code_point_stub, &pool);

        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point_stub)};

        getTupleValues(*transaction, hash(tuple), hashes);
    }
    SECTION("save nested tuple") {
        DBDeleter deleter;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();

        auto inner_tuple = Tuple();
        TuplePool pool;
        auto tuple = Tuple(inner_tuple, &pool);

        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple)};

        getTupleValues(*transaction, hash(tuple), hashes);
    }
    SECTION("save multiple valued tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);
        CodePointStub code_point_stub(1, 654546);

        auto transaction = storage.makeTransaction();

        auto inner_tuple = Tuple();
        uint256_t num = 1;
        auto tuple = Tuple(inner_tuple, num, code_point_stub, &pool);

        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple), hash(num),
                                      hash(code_point_stub)};

        getTupleValues(*transaction, hash(tuple), hashes);
    }
    SECTION("save multiple valued tuple, saveValue()") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);
        CodePointStub code_point_stub(1, 654546);

        auto transaction = storage.makeTransaction();

        auto inner_tuple = Tuple();
        uint256_t num = 1;
        auto tuple = Tuple(inner_tuple, num, code_point_stub, &pool);

        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple), hash(num),
                                      hash(code_point_stub)};

        getTupleValues(*transaction, hash(tuple), hashes);
    }
}

TEST_CASE("Save And Get Tuple") {
    SECTION("save 1 num tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);

        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
    }
    SECTION("save codepoint in tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePointStub code_point_stub(1, 654546);

        auto transaction = storage.makeTransaction();

        auto tuple = Tuple(code_point_stub, &pool);

        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
    }
    SECTION("save 1 num tuple twice") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto transaction2 = storage.makeTransaction();

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);

        saveValue(*transaction, tuple, 1, true);
        saveValue(*transaction2, tuple, 2, true);
        getTuple(*transaction, tuple, 2, true);
    }
    SECTION("save 2 num tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);
        std::vector<CodePoint> code;
        auto transaction = storage.makeTransaction();

        uint256_t num = 1;
        uint256_t num2 = 2;
        auto tuple = Tuple(num, num2, &pool);

        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
    }
    SECTION("save tuple in tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);
        saveValue(*transaction, tuple, 1, true);

        getTuple(*transaction, tuple, 1, true);
        getTuple(*transaction, inner_tuple, 1, true);
    }
    SECTION("save 2 tuples in tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        uint256_t num2 = 2;
        auto inner_tuple2 = Tuple(num2, &pool);
        auto tuple = Tuple(inner_tuple, inner_tuple2, &pool);
        saveValue(*transaction, tuple, 1, true);

        getTuple(*transaction, tuple, 1, true);
        getTuple(*transaction, inner_tuple, 1, true);
        getTuple(*transaction, inner_tuple2, 1, true);
    }
    SECTION("save saved tuple in tuple") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto transaction2 = storage.makeTransaction();

        uint256_t num = 1;
        auto inner_tuple = Tuple(num, &pool);
        auto tuple = Tuple(inner_tuple, &pool);

        saveValue(*transaction, inner_tuple, 1, true);
        getTuple(*transaction, inner_tuple, 1, true);
        saveValue(*transaction2, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
        getTuple(*transaction, inner_tuple, 2, true);
    }
}

void saveState(Transaction& transaction,
               MachineStateKeys storage_data,
               uint256_t machine_hash) {
    auto results = saveMachineState(transaction, storage_data, machine_hash);
    auto status = transaction.commit();

    REQUIRE(results.reference_count == 1);
    REQUIRE(results.status.ok());
}

void getSavedState(const Transaction& transaction,
                   uint256_t machine_hash,
                   MachineStateKeys expected_data,
                   int expected_ref_count,
                   std::vector<uint256_t> keys) {
    TuplePool pool;
    auto results = getMachineState(transaction, machine_hash);

    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == expected_ref_count);

    auto data = results.data;

    REQUIRE(data.status_char == expected_data.status_char);
    REQUIRE(data.pc == expected_data.pc);
    REQUIRE(data.datastack_hash == expected_data.datastack_hash);
    REQUIRE(data.auxstack_hash == expected_data.auxstack_hash);
    REQUIRE(data.register_hash == expected_data.register_hash);

    for (const auto& key : keys) {
        auto res = getValue(transaction, key, &pool);
        REQUIRE(res.status.ok());
    }
}

void deleteCheckpoint(Transaction& transaction,
                      uint256_t machine_hash,
                      std::vector<uint256_t> deleted_values) {
    TuplePool pool;
    auto res = deleteMachine(transaction, machine_hash);
    auto results = getMachineState(transaction, machine_hash);
    REQUIRE(results.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = getValue(transaction, hash_key, &pool);
        REQUIRE(res.status.ok() == false);
    }
}

void deleteCheckpointSavedTwice(Transaction& transaction,
                                uint256_t machine_hash,
                                std::vector<uint256_t> deleted_values) {
    auto res = deleteMachine(transaction, machine_hash);
    auto res2 = deleteMachine(transaction, machine_hash);
    auto results = getMachineState(transaction, machine_hash);

    TuplePool pool;

    REQUIRE(results.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = getValue(transaction, hash_key, &pool);
        REQUIRE(res.status.ok() == false);
    }
}

void deleteCheckpointSavedTwiceReordered(
    Transaction& transaction,
    uint256_t machine_hash,
    std::vector<uint256_t> deleted_values) {
    TuplePool pool;
    auto resultsx = getMachineState(transaction, machine_hash);
    for (auto& hash_key : deleted_values) {
        auto res = getValue(transaction, hash_key, &pool);
        REQUIRE(res.status.ok());
    }
    auto res = deleteMachine(transaction, machine_hash);
    auto results = getMachineState(transaction, machine_hash);
    REQUIRE(results.status.ok() == true);

    for (auto& hash_key : deleted_values) {
        auto res = getValue(transaction, hash_key, &pool);
        REQUIRE(res.status.ok());
    }
    auto res2 = deleteMachine(transaction, machine_hash);
    auto results2 = getMachineState(transaction, machine_hash);
    REQUIRE(results2.status.ok() == false);

    for (auto& hash_key : deleted_values) {
        auto res = getValue(transaction, hash_key, &pool);
        REQUIRE(res.status.ok() == false);
    }
}

MachineStateKeys makeStorageData(Transaction& transaction,
                                 value registerVal,
                                 Datastack stack,
                                 Datastack auxstack,
                                 Status state,
                                 CodePointStub pc,
                                 CodePointStub err_pc,
                                 BlockReason blockReason) {
    TuplePool pool;
    auto register_val_results = saveValue(transaction, registerVal);
    auto stack_tuple = stack.getTupleRepresentation(&pool);
    auto datastack_results = saveValue(transaction, stack_tuple);
    auto auxstack_tuple = auxstack.getTupleRepresentation(&pool);
    auto auxstack_results = saveValue(transaction, auxstack_tuple);

    auto status_str = (unsigned char)state;

    return MachineStateKeys{hash_value(registerVal),
                            hash(stack_tuple),
                            hash(auxstack_tuple),
                            pc,
                            err_pc,
                            status_str};
}

MachineStateKeys getStateValues(Transaction& transaction) {
    TuplePool pool;
    uint256_t register_val = 100;
    auto static_val = Tuple(register_val, Tuple(), &pool);

    CodePointStub code_point_stub(1, 654546);
    auto tup1 = Tuple(register_val, &pool);
    auto tup2 = Tuple(code_point_stub, tup1, &pool);

    Datastack data_stack;
    data_stack.push(register_val);
    Datastack aux_stack;
    aux_stack.push(register_val);
    aux_stack.push(code_point_stub);

    CodePointStub pc_codepoint_stub(0, 645357);
    CodePointStub err_codepoint_stub(0, 968769876);
    Status state = Status::Extensive;

    auto inbox_blocked = InboxBlocked(0);

    auto saved_data =
        makeStorageData(transaction, register_val, data_stack, aux_stack, state,
                        pc_codepoint_stub, err_codepoint_stub, inbox_blocked);

    return saved_data;
}

MachineStateKeys getDefaultValues(Transaction& transaction) {
    TuplePool pool;
    auto register_val = Tuple();
    auto data_stack = Tuple();
    auto aux_stack = Tuple();

    Status state = Status::Extensive;
    CodePointStub code_point_stub(1, 654546);

    auto data =
        makeStorageData(transaction, Tuple(), Datastack(), Datastack(), state,
                        code_point_stub, code_point_stub, NotBlocked());

    return data;
}

std::vector<uint256_t> getHashKeys(MachineStateKeys data) {
    std::vector<uint256_t> hash_keys;

    hash_keys.push_back(data.auxstack_hash);
    hash_keys.push_back(data.datastack_hash);
    hash_keys.push_back(data.register_hash);

    return hash_keys;
}

TEST_CASE("Save Machinestatedata") {
    SECTION("default") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto data_values = getDefaultValues(*transaction);

        uint256_t machine_hash = 5445765;

        saveState(*transaction, data_values, machine_hash);
    }
    SECTION("with values") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto state_data = getStateValues(*transaction);

        uint256_t machine_hash = 5445765;

        saveState(*transaction, state_data, machine_hash);
    }
}

TEST_CASE("Get Machinestate data") {
    SECTION("default") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();

        auto data_values = getDefaultValues(*transaction);
        auto keys = getHashKeys(data_values);

        uint256_t machine_hash = 6546457;

        saveMachineState(*transaction, data_values, machine_hash);
        transaction->commit();
        getSavedState(*transaction, machine_hash, data_values, 1, keys);
    }
    SECTION("with values") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();

        auto state_data = getStateValues(*transaction);
        auto keys = getHashKeys(state_data);

        uint256_t machine_hash = 6546457;

        saveState(*transaction, state_data, machine_hash);
        getSavedState(*transaction, machine_hash, state_data, 1, keys);
    }
}

TEST_CASE("Delete checkpoint") {
    SECTION("default") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto data_values = getDefaultValues(*transaction);

        uint256_t machine_hash = 6546457;

        saveMachineState(*transaction, data_values, machine_hash);
        transaction->commit();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpoint(*transaction, machine_hash, hash_keys);
    }
    SECTION("with actual state values") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto data_values = getStateValues(*transaction);

        uint256_t machine_hash = 6546457;

        saveMachineState(*transaction, data_values, machine_hash);
        transaction->commit();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpoint(*transaction, machine_hash, hash_keys);
    }
    SECTION("delete checkpoint saved twice") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto data_values = getStateValues(*transaction);

        uint256_t machine_hash = 6546457;

        saveMachineState(*transaction, data_values, machine_hash);
        saveMachineState(*transaction, data_values, machine_hash);
        transaction->commit();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpointSavedTwice(*transaction, machine_hash, hash_keys);
    }
    SECTION("delete checkpoint saved twice, reordered") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto transaction2 = storage.makeTransaction();
        auto data_values = getStateValues(*transaction);

        uint256_t machine_hash = 6546457;

        saveMachineState(*transaction, data_values, machine_hash);
        transaction->commit();
        saveMachineState(*transaction2, data_values, machine_hash);
        transaction2->commit();
        auto hash_keys = getHashKeys(data_values);

        deleteCheckpointSavedTwiceReordered(*transaction, machine_hash,
                                            hash_keys);
    }
}
