/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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

#include <avm/machine.hpp>

#include <avm_values/code.hpp>

#include <rocksdb/status.h>

#include <catch2/catch.hpp>

void generateTestMachine(std::unique_ptr<Machine>& mach) {
    auto& code = mach->machine_state.code;
    auto stub1 = code->addSegment();
    auto stub2 = code->addSegment();
    auto stub3 = code->addSegment();

    auto add_op1 = [&](Operation op) {
        stub1 = code->addOperation(stub1.pc, std::move(op));
    };
    auto add_op2 = [&](Operation op) {
        stub2 = code->addOperation(stub2.pc, std::move(op));
    };
    auto add_op3 = [&](Operation op) {
        stub3 = code->addOperation(stub3.pc, std::move(op));
    };

    add_op2(Operation{OpCode::HALT});
    add_op2(Operation{OpCode::ADD});

    add_op3(Operation{OpCode::JUMP, stub2});
    add_op3(Operation{OpCode::ADD});

    add_op1(Operation{OpCode::JUMP, stub3});
    add_op1(Operation{OpCode::ADD});

    for (int i = 0; i < 4; i++) {
        mach->machine_state.stack.push(uint256_t{1});
    }
    mach->machine_state.pc = CodePointRef(1, 2);
}

void checkRun(Machine& mach, uint64_t gas_count_target = 27) {
    MachineExecutionConfig execConfig;
    execConfig.max_gas = gas_count_target;
    mach.machine_state.context = AssertionContext(execConfig);
    auto assertion = mach.run();
    REQUIRE(assertion.gas_count <= gas_count_target);
    auto val = mach.machine_state.stack.pop();
    REQUIRE(values_equal(val, Value{uint256_t{4}}));
    REQUIRE(mach.machine_state.stack.stacksize() == 0);
}

TEST_CASE("Code works correctly") {
    DBDeleter deleter;
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    auto result = storage.initialize(LoadedExecutable(
        std::make_shared<UnsafeCodeSegment>(0), Value{Tuple()}));
    REQUIRE(result.status.ok());
    REQUIRE(result.finished == false);
    auto mach = storage.getInitialMachine();
    generateTestMachine(mach);
    checkRun(*mach);
}

TEST_CASE("Code serialization") {
    DBDeleter deleter;
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    auto result = storage.initialize(LoadedExecutable(
        std::make_shared<UnsafeCodeSegment>(0), Value{Tuple()}));
    REQUIRE(result.status.ok());
    REQUIRE(result.finished == false);
    ValueCache value_cache{1, 0};

    auto mach = storage.getInitialMachine();
    generateTestMachine(mach);
    auto tx = storage.makeReadWriteTransaction();

    SECTION("Save and load") {
        auto save_ret = saveTestMachine(*tx, *mach);
        REQUIRE(save_ret.status.ok());
        REQUIRE(tx->commit().ok());
        auto mach2 = storage.getMachine(mach->hash(), value_cache);
        checkRun(*mach2);
    }

    SECTION("Save different and load") {
        auto save_ret = saveTestMachine(*tx, *mach);
        REQUIRE(save_ret.status.ok());

        auto mach2{*mach};
        MachineExecutionConfig execConfig;
        execConfig.max_gas = 7;
        mach2.machine_state.context = AssertionContext(execConfig);
        mach2.run();
        save_ret = saveTestMachine(*tx, mach2);
        REQUIRE(save_ret.status.ok());

        SECTION("Delete first") {
            auto del_ret = deleteMachine(*tx, mach->hash());
            REQUIRE(del_ret.status.ok());
            REQUIRE(tx->commit().ok());
            auto mach3 = storage.getMachine(mach2.hash(), value_cache);
            checkRun(*mach3);
        }

        SECTION("Delete second") {
            auto del_ret = deleteMachine(*tx, mach2.hash());
            REQUIRE(del_ret.status.ok());
            REQUIRE(tx->commit().ok());
            auto mach3 = storage.getMachine(mach->hash(), value_cache);
            checkRun(*mach3);
        }
    }

    SECTION("Save twice, delete and load") {
        saveTestMachine(*tx, *mach);
        saveTestMachine(*tx, *mach);
        deleteMachine(*tx, mach->hash());
        REQUIRE(tx->commit().ok());
        auto mach2 = storage.getMachine(mach->hash(), value_cache);
        checkRun(*mach2);
    }
}

TEST_CASE("Code forks are identical to original") {
    CoreCode code;
    std::vector<CodePointStub> stubs(1, code.addSegment());
    constexpr size_t num_ops = 45;
    for (size_t i = 0; i < num_ops; i++) {
        stubs.push_back(
            code.addOperation(stubs.back().pc, Operation{OpCode::NOP}));
    }
    for (size_t i = 0; i < stubs.size(); i++) {
        auto new_stub = stubs[i];
        for (size_t j = i; j < num_ops; j++) {
            new_stub = code.addOperation(new_stub.pc, Operation{OpCode::NOP});
            REQUIRE(::hash(code.loadCodePoint(new_stub.pc)) == new_stub.hash);
        }
        REQUIRE(new_stub.hash == stubs.back().hash);
    }
}
