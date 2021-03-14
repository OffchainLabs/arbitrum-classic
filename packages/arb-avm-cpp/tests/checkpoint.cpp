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

#include "helper.hpp"

#include <data_storage/arbstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>

#define CATCH_CONFIG_ENABLE_BENCHMARKING 1
#include <catch2/catch.hpp>

void saveValue(ReadWriteTransaction& transaction,
               const value& val,
               uint32_t expected_ref_count,
               bool expected_status) {
    auto results = saveValue(transaction, val);
    transaction.commit();
    REQUIRE(results.status.ok() == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

DbResult<value> getValue(const ReadTransaction& transaction,
                         const value& value_target,
                         uint32_t expected_ref_count,
                         bool expected_status,
                         ValueCache& value_cache) {
    auto res = getValue(transaction, hash_value(value_target), value_cache);
    if (expected_status) {
        REQUIRE(std::holds_alternative<CountedData<value>>(res));
        REQUIRE(std::get<CountedData<value>>(res).reference_count ==
                expected_ref_count);
        REQUIRE(hash_value(std::get<CountedData<value>>(res).data) ==
                hash_value(value_target));
    } else {
        REQUIRE(std::holds_alternative<rocksdb::Status>(res));
    }
    return res;
}

void getTuple(const ReadTransaction& transaction,
              const value& val,
              uint32_t expected_ref_count,
              bool expected_status,
              ValueCache& value_cache) {
    auto res = getValue(transaction, val, expected_ref_count, expected_status,
                        value_cache);
    const auto& tuple = std::get<Tuple>(val);
    if (expected_status) {
        REQUIRE(std::holds_alternative<Tuple>(
            std::get<CountedData<value>>(res).data));
        REQUIRE(std::get<Tuple>(std::get<CountedData<value>>(res).data) ==
                tuple);
    }
}

void getTupleValues(const ReadTransaction& transaction,
                    uint256_t tuple_hash,
                    std::vector<uint256_t> value_hashes,
                    ValueCache& value_cache) {
    auto results = getValue(transaction, tuple_hash, value_cache);
    REQUIRE(std::holds_alternative<CountedData<value>>(results));
    auto val = std::get<CountedData<value>>(results).data;
    REQUIRE(std::holds_alternative<Tuple>(val));
    auto tuple = std::get<Tuple>(val);
    REQUIRE(tuple.tuple_size() == value_hashes.size());

    for (size_t i = 0; i < value_hashes.size(); i++) {
        REQUIRE(hash_value(tuple.get_element(i)) == value_hashes[i]);
    }
}

TEST_CASE("Save value") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    auto transaction = storage.makeReadWriteTransaction();

    SECTION("save 1 num tuple") {
        uint256_t num = 1;
        auto tuple = Tuple::createTuple(num);
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
    ArbStorage storage(dbpath);
    auto transaction = storage.makeReadWriteTransaction();

    SECTION("save 1 num tuple") {
        uint256_t num = 1;
        auto tuple = Tuple::createTuple(num);
        saveValue(*transaction, tuple, 1, true);
    }
    SECTION("save 2, 1 num tuples") {
        uint256_t num = 1;
        auto tuple = Tuple::createTuple(num);
        saveValue(*transaction, tuple, 1, true);
        saveValue(*transaction, tuple, 2, true);
    }
    SECTION("saved tuple in tuple") {
        uint256_t num = 1;
        value inner_tuple = Tuple::createTuple(num);
        auto tuple = Tuple::createTuple(inner_tuple);
        saveValue(*transaction, tuple, 1, true);
        saveValue(*transaction, tuple, 2, true);
    }
}

TEST_CASE("Save and get value") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    auto transaction = storage.makeReadWriteTransaction();
    ValueCache value_cache{1, 0};

    SECTION("save empty tuple") {
        auto tuple = Tuple();
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true, value_cache);
    }
    SECTION("save num") {
        uint256_t num = 1;
        saveValue(*transaction, num, 1, true);
        getValue(*transaction, num, 1, true, value_cache);
    }
    SECTION("save codepoint") {
        CodePointStub code_point_stub({0, 1}, 654546);
        saveValue(*transaction, code_point_stub, 1, true);
        getValue(*transaction, code_point_stub, 1, true, value_cache);
    }
    SECTION("save err codepoint") {
        CodePointStub code_point_stub({0, 1}, 654546);
        saveValue(*transaction, code_point_stub, 1, true);
        getValue(*transaction, code_point_stub, 1, true, value_cache);
    }
}

TEST_CASE("Save and get tuple values") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    auto transaction = storage.makeReadWriteTransaction();
    ValueCache value_cache{1, 0};

    SECTION("save num tuple") {
        uint256_t num = 1;
        auto tuple = Tuple::createTuple(num);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(num)};
        getTupleValues(*transaction, hash(tuple), hashes, value_cache);
    }
    SECTION("save codepoint tuple") {
        CodePointStub code_point_stub({0, 1}, 654546);
        auto tuple = Tuple::createTuple(code_point_stub);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point_stub)};
        getTupleValues(*transaction, hash(tuple), hashes, value_cache);
    }
    SECTION("save codepoint tuple") {
        CodePointStub code_point_stub({0, 1}, 654546);
        auto tuple = Tuple::createTuple(code_point_stub);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(code_point_stub)};
        getTupleValues(*transaction, hash(tuple), hashes, value_cache);
    }
    SECTION("save nested tuple") {
        value inner_tuple = Tuple();
        value tuple = Tuple::createTuple(inner_tuple);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash_value(inner_tuple)};
        getTupleValues(*transaction, hash_value(tuple), hashes, value_cache);
    }
    SECTION("save multiple valued tuple") {
        CodePointStub code_point_stub({0, 1}, 654546);
        value inner_tuple = Tuple();
        uint256_t num = 1;
        value tuple = Tuple(inner_tuple, num, code_point_stub);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash_value(inner_tuple), hash(num),
                                      hash(code_point_stub)};
        getTupleValues(*transaction, hash_value(tuple), hashes, value_cache);
    }
    SECTION("save multiple valued tuple, saveValue()") {
        CodePointStub code_point_stub({0, 1}, 654546);
        auto inner_tuple = Tuple();
        uint256_t num = 1;
        auto tuple = Tuple(inner_tuple, num, code_point_stub);
        saveValue(*transaction, tuple, 1, true);
        std::vector<uint256_t> hashes{hash(inner_tuple), hash(num),
                                      hash(code_point_stub)};
        getTupleValues(*transaction, hash(tuple), hashes, value_cache);
    }
}

TEST_CASE("Save And Get Tuple") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    auto transaction = storage.makeReadWriteTransaction();

    SECTION("save 1 num tuple") {
        ValueCache value_cache{1, 0};
        uint256_t num = 1;
        auto tuple = Tuple::createTuple(num);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true, value_cache);
    }
    SECTION("save codepoint in tuple") {
        ValueCache value_cache{1, 0};
        CodePointStub code_point_stub({0, 1}, 654546);
        auto tuple = Tuple::createTuple(code_point_stub);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true, value_cache);
    }
    SECTION("save 1 num tuple twice") {
        ValueCache value_cache{1, 0};
        auto transaction2 = storage.makeReadWriteTransaction();
        uint256_t num = 1;
        auto tuple = Tuple::createTuple(num);
        saveValue(*transaction, tuple, 1, true);
        saveValue(*transaction2, tuple, 2, true);
        getTuple(*transaction, tuple, 2, true, value_cache);

        // Test cache
        getTuple(*transaction, tuple, 0, true, value_cache);
    }
    SECTION("save 2 num tuple") {
        ValueCache value_cache{1, 0};
        std::vector<CodePoint> code;
        uint256_t num = 1;
        uint256_t num2 = 2;
        auto tuple = Tuple(num, num2);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true, value_cache);
    }
    SECTION("save tuple in tuple") {
        ValueCache value_cache{1, 0};
        uint256_t num = 1;
        auto inner_tuple = Tuple::createTuple(num);
        auto tuple = Tuple::createTuple(inner_tuple);
        REQUIRE(hash(tuple) != hash(inner_tuple));
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true, value_cache);
        getTuple(*transaction, inner_tuple, 1, true, value_cache);
    }
    SECTION("save 2 tuples in tuple") {
        ValueCache value_cache{1, 0};
        uint256_t num = 1;
        value inner_tuple = Tuple::createTuple(num);
        uint256_t num2 = 2;
        value inner_tuple2 = Tuple::createTuple(num2);
        auto tuple = Tuple(inner_tuple, inner_tuple2);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true, value_cache);
        getTuple(*transaction, inner_tuple, 1, true, value_cache);
        getTuple(*transaction, inner_tuple2, 1, true, value_cache);
    }
    SECTION("save saved tuple in tuple") {
        ValueCache value_cache{1, 0};
        auto transaction2 = storage.makeReadWriteTransaction();
        uint256_t num = 1;
        value inner_tuple = Tuple::createTuple(num);
        value tuple = Tuple::createTuple(inner_tuple);
        saveValue(*transaction, inner_tuple, 1, true);
        getTuple(*transaction, inner_tuple, 1, true, value_cache);
        saveValue(*transaction, tuple, 1, true);
        getTuple(*transaction, tuple, 1, true, value_cache);

        // Use different cache to get real reference count
        ValueCache value_cache2{1, 0};
        getTuple(*transaction, inner_tuple, 2, true, value_cache2);

        // Test cache
        getTuple(*transaction, inner_tuple, 0, true, value_cache2);
    }
}

TEST_CASE("Checkpoint Benchmark") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    auto transaction = storage.makeReadWriteTransaction();
    uint256_t num = 1;
    value tuple = Tuple::createTuple(num);
    for (uint64_t i = 1; i < 100000; i++) {
        tuple = Tuple::createTuple(tuple);
    }
    saveValue(*transaction, tuple);
    ValueCache value_cache{1, 0};

    auto tuple_hash = hash_value(tuple);
    // Initial get to populate cache
    getValue(*transaction, tuple_hash, value_cache);

    BENCHMARK_ADVANCED("restoreCheckpoint1")
    (Catch::Benchmark::Chronometer meter) {
        meter.measure([&transaction, tuple_hash, &value_cache] {
            return getValue(*transaction, tuple_hash, value_cache);
        });
    };
}

void saveState(ReadWriteTransaction& transaction,
               const Machine& machine,
               uint256_t expected_ref_count) {
    auto results = saveMachine(transaction, machine);
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(transaction.commit().ok());
}

void checkSavedState(const ReadWriteTransaction& transaction,
                     const Machine& expected_machine,
                     uint32_t expected_ref_count) {
    auto expected_hash = expected_machine.hash();
    REQUIRE(expected_hash);
    auto results = getMachineStateKeys(transaction, *expected_hash);
    REQUIRE(std::holds_alternative<CountedData<MachineStateKeys>>(results));
    auto res = std::get<CountedData<MachineStateKeys>>(results);
    REQUIRE(res.reference_count == expected_ref_count);

    auto data = res.data;
    REQUIRE(data.status == expected_machine.machine_state.state);
    REQUIRE(data.pc.pc == expected_machine.machine_state.pc);
    REQUIRE(
        data.datastack_hash ==
        hash(expected_machine.machine_state.stack.getTupleRepresentation()));
    REQUIRE(
        data.auxstack_hash ==
        hash(expected_machine.machine_state.auxstack.getTupleRepresentation()));
    REQUIRE(data.register_hash ==
            hash_value(expected_machine.machine_state.registerVal));

    ValueCache value_cache{1, 0};
    REQUIRE(!std::holds_alternative<rocksdb::Status>(
        getValue(transaction, data.datastack_hash, value_cache)));
    REQUIRE(!std::holds_alternative<rocksdb::Status>(
        getValue(transaction, data.auxstack_hash, value_cache)));
    REQUIRE(!std::holds_alternative<rocksdb::Status>(
        getValue(transaction, data.register_hash, value_cache)));
}

void checkDeletedCheckpoint(ReadTransaction& transaction,
                            const Machine& deleted_machine) {
    auto deleted_hash = deleted_machine.hash();
    REQUIRE(deleted_hash);
    auto results = getMachineStateKeys(transaction, *deleted_hash);
    REQUIRE(std::holds_alternative<rocksdb::Status>(results));

    auto datastack_tup =
        deleted_machine.machine_state.stack.getTupleRepresentation();
    auto auxstack_tup =
        deleted_machine.machine_state.auxstack.getTupleRepresentation();
    ValueCache value_cache{1, 0};
    REQUIRE(std::holds_alternative<rocksdb::Status>(
        getValue(transaction, hash(datastack_tup), value_cache)));
    REQUIRE(std::holds_alternative<rocksdb::Status>(
        getValue(transaction, hash(auxstack_tup), value_cache)));
    REQUIRE(std::holds_alternative<rocksdb::Status>(getValue(
        transaction, hash_value(deleted_machine.machine_state.registerVal),
        value_cache)));
}

void deleteCheckpoint(ReadWriteTransaction& transaction,
                      const Machine& deleted_machine) {
    auto deleted_hash = deleted_machine.hash();
    REQUIRE(deleted_hash);
    auto res = deleteMachine(transaction, *deleted_hash);
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

    Datastack data_stack;
    data_stack.push(register_val);
    Datastack aux_stack;
    aux_stack.push(register_val);
    aux_stack.push(code_point_stub);

    uint256_t arb_gas_remaining = 534574678365;

    CodePointRef pc{0, 0};
    CodePointStub err_pc({0, 0}, 968769876);
    Status state = Status::Extensive;

    auto output = MachineOutput{{42, 54}, 23, 54, 12, 65, 76, 43, 65};

    staged_variant staged_message;

    return Machine(MachineState(std::move(code), register_val,
                                std::move(static_val), data_stack, aux_stack,
                                arb_gas_remaining, state, pc, err_pc,
                                std::move(staged_message), output));
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
    auto output = MachineOutput{{42, 54}, 23, 54, 12, 65, 76, 43, 34};
    staged_variant staged_message;
    return Machine(MachineState(std::move(code), register_val,
                                std::move(static_val), data_stack, aux_stack,
                                arb_gas_remaining, state, pc, err_pc,
                                staged_message, output));
}

TEST_CASE("Save Machinestatedata") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    auto transaction = storage.makeReadWriteTransaction();

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
    ArbStorage storage(dbpath);
    auto transaction = storage.makeReadWriteTransaction();

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
    ArbStorage storage(dbpath);
    auto transaction = storage.makeReadWriteTransaction();

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
            auto transaction2 = storage.makeReadWriteTransaction();
            saveState(*transaction2, machine, 2);
        }
        auto transaction3 = storage.makeReadWriteTransaction();
        auto machine_hash = machine.hash();
        REQUIRE(machine_hash);
        auto res = deleteMachine(*transaction3, *machine.hash());
        REQUIRE(res.status.ok());
        auto res2 = deleteMachine(*transaction3, *machine.hash());
        REQUIRE(res2.status.ok());
        checkDeletedCheckpoint(*transaction3, machine);
    }
    SECTION("delete checkpoint saved twice, reordered") {
        auto transaction2 = storage.makeReadWriteTransaction();
        auto machine = getComplexMachine();
        saveState(*transaction, machine, 1);
        saveState(*transaction2, machine, 2);

        checkSavedState(*transaction, machine, 2);
        auto machine_hash = machine.hash();
        REQUIRE(machine_hash);
        auto res = deleteMachine(*transaction, *machine.hash());
        checkSavedState(*transaction, machine, 1);
        auto res2 = deleteMachine(*transaction, *machine.hash());
        checkDeletedCheckpoint(*transaction, machine);
    }
}
