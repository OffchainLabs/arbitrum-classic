/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

#include <data_storage/checkpointstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>

#include <avm_values/vmValueParser.hpp>

#include <catch2/catch.hpp>

void saveValue(Transaction& transaction,
               const value& val,
               uint32_t expected_ref_count,
               bool expected_status) {
    auto results = saveValue(transaction, val);
    transaction.commit();
    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void getValue(const Transaction& transaction,
              const value& value,
              uint32_t expected_ref_count,
              bool expected_status) {
    auto results = getValue(transaction, hash_value(value));

    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(hash_value(results.data) == hash_value(value));
}

void getTuple(const Transaction& transaction,
              const value& val,
              uint32_t expected_ref_count,
              bool expected_status) {
    auto tuple = val.get<Tuple>();
    auto results = getValue(transaction, hash(tuple));

    REQUIRE(nonstd::holds_alternative<Tuple>(results.data));

    auto loadedTuple = nonstd::get<Tuple>(results.data);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(loadedTuple == tuple);
    REQUIRE(loadedTuple.tuple_size() == tuple.tuple_size());
    REQUIRE(results.status.ok() == expected_status);
}

void getTupleValues(const Transaction& transaction,
                    uint256_t tuple_hash,
                    std::vector<uint256_t> value_hashes) {
    auto results = getValue(transaction, tuple_hash);
    REQUIRE(results.status.ok());
    REQUIRE(nonstd::holds_alternative<Tuple>(results.data));

    auto tuple = nonstd::get<Tuple>(results.data);
    REQUIRE(tuple.tuple_size() == value_hashes.size());

    for (size_t i = 0; i < value_hashes.size(); i++) {
        REQUIRE(hash_value(tuple.get_element(i)) == value_hashes[i]);
    }
}

TEST_CASE("Save value") {
    DBDeleter deleter;
    CheckpointStorage storage(dbpath);
    auto transaction = storage.makeTransaction();

    SECTION("save 1 num tuple") {
        uint256_t num = 1;
        auto tuple = Tuple(num);
        saveValue(*transaction, tuple, 1, true);
    }
    SECTION("save num") {
        uint256_t num = 1;
        saveValue(*transaction, num, 1, true);
    }
    SECTION("save codepoint") {
        CodePointStub code_point_stub({0, 1}, 654546);
        saveValue(*transaction, code_point_stub, 1, true);
    }
}

TEST_CASE("Save tuple") {
    DBDeleter deleter;
    CheckpointStorage storage(dbpath);
    auto transaction = storage.makeTransaction();

    SECTION("save 1 num tuple") {
        uint256_t num = 1;
        auto tuple = Tuple(num);
        saveValue(*transaction, tuple, 1, true);
    }
    SECTION("save 2, 1 num tuples") {
        uint256_t num = 1;
        auto tuple = Tuple(num);
        saveValue(*transaction, tuple, 1, true);
        saveValue(*transaction, tuple, 2, true);
    }
    SECTION("saved tuple in tuple") {
        uint256_t num = 1;
        auto inner_tuple = Tuple(num);
        auto tuple = Tuple(inner_tuple);
        saveValue(*transaction, tuple, 1, true);
        saveValue(*transaction, tuple, 2, true);
    }
}

TEST_CASE("Save and get value") {
    DBDeleter deleter;
    CheckpointStorage storage(dbpath);
    auto transaction = storage.makeTransaction();

    SECTION("save empty tuple") {
        auto tuple = Tuple();
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
    }
    SECTION("save num") {
        uint256_t num = 1;
        saveValue(*transaction, num, 1, true);
        getValue(*transaction, num, 1, true);
    }
    SECTION("save codepoint") {
        CodePointStub code_point_stub({0, 1}, 654546);
        saveValue(*transaction, code_point_stub, 1, true);
        getValue(*transaction, code_point_stub, 1, true);
    }
    SECTION("save err codepoint") {
        CodePointStub code_point_stub({0, 1}, 654546);
        saveValue(*transaction, code_point_stub, 1, true);
        getValue(*transaction, code_point_stub, 1, true);
    }
}

TEST_CASE("Save and get tuple values") {
    DBDeleter deleter;
    CheckpointStorage storage(dbpath);
    auto transaction = storage.makeTransaction();

    SECTION("save num tuple") {
        uint256_t num = 1;
        auto tuple = Tuple(num);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(num)};
        getTupleValues(*transaction, hash(tuple), hashes);
    }
    SECTION("save codepoint tuple") {
        CodePointStub code_point_stub({0, 1}, 654546);
        auto tuple = Tuple(code_point_stub);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point_stub)};
        getTupleValues(*transaction, hash(tuple), hashes);
    }
    SECTION("save codepoint tuple") {
        CodePointStub code_point_stub({0, 1}, 654546);
        auto tuple = Tuple(code_point_stub);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point_stub)};
        getTupleValues(*transaction, hash(tuple), hashes);
    }
    SECTION("save nested tuple") {
        value inner_tuple = Tuple();
        value tuple = Tuple(inner_tuple);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash_value(inner_tuple)};
        getTupleValues(*transaction, hash_value(tuple), hashes);
    }
    SECTION("save multiple valued tuple") {
        CodePointStub code_point_stub({0, 1}, 654546);
        value inner_tuple = Tuple();
        uint256_t num = 1;
        value tuple = Tuple(inner_tuple, num, code_point_stub);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash_value(inner_tuple), hash(num),
                                      hash(code_point_stub)};
        getTupleValues(*transaction, hash_value(tuple), hashes);
    }
    SECTION("save multiple valued tuple, saveValue()") {
        CodePointStub code_point_stub({0, 1}, 654546);
        auto inner_tuple = Tuple();
        uint256_t num = 1;
        auto tuple = Tuple(inner_tuple, num, code_point_stub);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple), hash(num),
                                      hash(code_point_stub)};
        getTupleValues(*transaction, hash(tuple), hashes);
    }
}

TEST_CASE("Save And Get Tuple") {
    DBDeleter deleter;
    CheckpointStorage storage(dbpath);
    auto transaction = storage.makeTransaction();

    SECTION("save 1 num tuple") {
        uint256_t num = 1;
        auto tuple = Tuple(num);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
    }
    SECTION("save codepoint in tuple") {
        CodePointStub code_point_stub({0, 1}, 654546);
        auto tuple = Tuple(code_point_stub);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
    }
    SECTION("save 1 num tuple twice") {
        auto transaction2 = storage.makeTransaction();
        uint256_t num = 1;
        auto tuple = Tuple(num);
        saveValue(*transaction, tuple, 1, true);
        saveValue(*transaction2, tuple, 2, true);
        getTuple(*transaction, tuple, 2, true);
    }
    SECTION("save 2 num tuple") {
        std::vector<CodePoint> code;
        uint256_t num = 1;
        uint256_t num2 = 2;
        auto tuple = Tuple(num, num2);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
    }
    SECTION("save tuple in tuple") {
        uint256_t num = 1;
        auto inner_tuple = Tuple(num);
        auto tuple = Tuple(inner_tuple);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
        getTuple(*transaction, inner_tuple, 1, true);
    }
    SECTION("save 2 tuples in tuple") {
        uint256_t num = 1;
        auto inner_tuple = Tuple(num);
        uint256_t num2 = 2;
        auto inner_tuple2 = Tuple(num2);
        auto tuple = Tuple(inner_tuple, inner_tuple2);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
        getTuple(*transaction, inner_tuple, 1, true);
        getTuple(*transaction, inner_tuple2, 1, true);
    }
    SECTION("save saved tuple in tuple") {
        auto transaction2 = storage.makeTransaction();
        uint256_t num = 1;
        value inner_tuple = Tuple(num);
        value tuple = Tuple(inner_tuple);
        saveValue(*transaction, inner_tuple, 1, true);
        getTuple(*transaction, inner_tuple, 1, true);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true);
        getTuple(*transaction, inner_tuple, 2, true);
    }
}

void saveState(Transaction& transaction,
               const Machine& machine,
               uint256_t expected_ref_count) {
    auto results = saveMachine(transaction, machine);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(results.status.ok());
    REQUIRE(transaction.commit().ok());
}

void checkSavedState(const Transaction& transaction,
                     const Machine& expected_machine,
                     uint32_t expected_ref_count) {
    auto results = getMachineState(transaction, expected_machine.hash());
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == expected_ref_count);

    auto data = results.data;
    REQUIRE(data.status == expected_machine.machine_state.state);
    REQUIRE(data.pc == expected_machine.machine_state.pc);
    REQUIRE(
        data.datastack_hash ==
        hash(expected_machine.machine_state.stack.getTupleRepresentation()));
    REQUIRE(
        data.auxstack_hash ==
        hash(expected_machine.machine_state.auxstack.getTupleRepresentation()));
    REQUIRE(data.register_hash ==
            hash_value(expected_machine.machine_state.registerVal));

    REQUIRE(getValue(transaction, data.datastack_hash).status.ok());
    REQUIRE(getValue(transaction, data.auxstack_hash).status.ok());
    REQUIRE(getValue(transaction, data.register_hash).status.ok());
}

void checkDeletedCheckpoint(Transaction& transaction,
                            const Machine& deleted_machine) {
    auto results = getMachineState(transaction, deleted_machine.hash());
    REQUIRE(!results.status.ok());

    auto datastack_tup =
        deleted_machine.machine_state.stack.getTupleRepresentation();
    auto auxstack_tup =
        deleted_machine.machine_state.auxstack.getTupleRepresentation();
    REQUIRE(!getValue(transaction, hash(datastack_tup)).status.ok());
    REQUIRE(!getValue(transaction, hash(auxstack_tup)).status.ok());
    REQUIRE(!getValue(transaction,
                      hash_value(deleted_machine.machine_state.registerVal))
                 .status.ok());
}

void deleteCheckpoint(Transaction& transaction,
                      const Machine& deleted_machine) {
    auto res = deleteMachine(transaction, deleted_machine.hash());
    REQUIRE(res.status.ok());
    checkDeletedCheckpoint(transaction, deleted_machine);
}

Machine getComplexMachine() {
    auto code = std::make_shared<Code>();
    auto stub = code->addSegment();
    stub = code->addOperation(stub.pc, Operation(OpCode::ADD));
    stub = code->addOperation(stub.pc, Operation(OpCode::MUL));
    stub = code->addOperation(stub.pc, Operation(OpCode::SUB));
    uint256_t register_val = 100;
    auto static_val = Tuple(register_val, Tuple());

    CodePointStub code_point_stub({0, 1}, 654546);
    auto tup1 = Tuple(register_val);
    auto tup2 = Tuple(code_point_stub, tup1);

    Datastack data_stack;
    data_stack.push(register_val);
    Datastack aux_stack;
    aux_stack.push(register_val);
    aux_stack.push(code_point_stub);

    uint256_t arb_gas_remaining = 534574678365;

    CodePointRef pc{0, 0};
    CodePointStub err_pc({0, 0}, 968769876);
    Status state = Status::Extensive;

    Tuple staged_message(uint256_t{100}, uint256_t{200});

    return Machine(MachineState(std::move(code), register_val,
                                std::move(static_val), data_stack, aux_stack,
                                arb_gas_remaining, state, pc, err_pc,
                                std::move(staged_message)));
}

Machine getDefaultMachine() {
    auto code = std::make_shared<Code>();
    code->addSegment();
    auto static_val = Tuple();
    auto register_val = Tuple();
    auto data_stack = Tuple();
    auto aux_stack = Tuple();
    uint256_t arb_gas_remaining = 534574678365;
    CodePointRef pc(0, 0);
    CodePointStub err_pc({0, 0}, 968769876);
    Status state = Status::Extensive;
    Tuple staged_message;
    return Machine(MachineState(
        std::move(code), register_val, std::move(static_val), data_stack,
        aux_stack, arb_gas_remaining, state, pc, err_pc, staged_message));
}

TEST_CASE("Save Machinestatedata") {
    DBDeleter deleter;
    CheckpointStorage storage(dbpath);
    auto transaction = storage.makeTransaction();

    SECTION("default") {
        auto machine = getDefaultMachine();
        saveState(*transaction, machine, 1);
    }
    SECTION("with values") {
        auto machine = getComplexMachine();
        saveState(*transaction, machine, 1);
    }
}

TEST_CASE("Get Machinestate data") {
    DBDeleter deleter;
    CheckpointStorage storage(dbpath);
    auto transaction = storage.makeTransaction();

    SECTION("default") {
        auto machine = getDefaultMachine();
        saveState(*transaction, machine, 1);
        checkSavedState(*transaction, machine, 1);
    }
    SECTION("with values") {
        auto machine = getComplexMachine();
        saveState(*transaction, machine, 1);
        checkSavedState(*transaction, machine, 1);
    }
}

TEST_CASE("Delete checkpoint") {
    DBDeleter deleter;
    CheckpointStorage storage(dbpath);
    auto transaction = storage.makeTransaction();

    SECTION("default") {
        auto machine = getDefaultMachine();
        saveState(*transaction, machine, 1);
        deleteCheckpoint(*transaction, machine);
    }
    SECTION("with actual state values") {
        auto machine = getComplexMachine();
        saveState(*transaction, machine, 1);
        deleteCheckpoint(*transaction, machine);
    }
    SECTION("delete checkpoint saved twice") {
        auto machine = getComplexMachine();
        saveState(*transaction, machine, 1);
        {
            auto transaction2 = storage.makeTransaction();
            saveState(*transaction2, machine, 2);
        }
        auto transaction3 = storage.makeTransaction();
        auto res = deleteMachine(*transaction3, machine.hash());
        REQUIRE(res.status.ok());
        auto res2 = deleteMachine(*transaction3, machine.hash());
        REQUIRE(res2.status.ok());
        checkDeletedCheckpoint(*transaction3, machine);
    }
    SECTION("delete checkpoint saved twice, reordered") {
        auto transaction2 = storage.makeTransaction();
        auto machine = getComplexMachine();
        saveState(*transaction, machine, 1);
        saveState(*transaction2, machine, 2);

        checkSavedState(*transaction, machine, 2);
        auto res = deleteMachine(*transaction, machine.hash());
        checkSavedState(*transaction, machine, 1);
        auto res2 = deleteMachine(*transaction, machine.hash());
        checkDeletedCheckpoint(*transaction, machine);
    }
}
