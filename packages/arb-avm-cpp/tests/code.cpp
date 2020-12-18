/*
 * Copyright 2020, Offchain Labs, Inc.
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

#include <data_storage/arbstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>

#include <avm/machine.hpp>

#include <avm_values/code.hpp>

#include <rocksdb/status.h>

#include <catch2/catch.hpp>

Machine generateTestMachine() {
    auto code = std::make_shared<Code>();
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

    Machine mach{std::move(code), Tuple()};
    for (int i = 0; i < 4; i++) {
        mach.machine_state.stack.push(uint256_t{1});
    }
    return mach;
}

void checkRun(Machine& mach, uint64_t step_count_target = 6) {
    auto assertion = mach.run(10000, {}, std::chrono::seconds{0});
    REQUIRE(assertion.stepCount == step_count_target);
    auto val = mach.machine_state.stack.pop();
    REQUIRE(val == value{uint256_t{4}});
    REQUIRE(mach.machine_state.stack.stacksize() == 0);
}

TEST_CASE("Code works correctly") {
    auto mach = generateTestMachine();
    checkRun(mach);
}

TEST_CASE("Code serialization") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    auto mach = generateTestMachine();
    auto tx = storage.makeTransaction();
    ValueCache value_cache{};

    SECTION("Save and load") {
        saveMachine(*tx, mach);
        REQUIRE(tx->commit().ok());
        auto mach2 = storage.getMachine(mach.hash(), value_cache);
        checkRun(mach2);
    }

    SECTION("Save different and load") {
        auto mach2 = mach;
        mach2.run(2, {}, std::chrono::seconds{0});
        saveMachine(*tx, mach);
        saveMachine(*tx, mach2);

        SECTION("Delete first") {
            deleteMachine(*tx, mach.hash());
            REQUIRE(tx->commit().ok());
            auto mach3 = storage.getMachine(mach2.hash(), value_cache);
            checkRun(mach3, 4);
        }

        SECTION("Delete second") {
            deleteMachine(*tx, mach2.hash());
            REQUIRE(tx->commit().ok());
            auto mach3 = storage.getMachine(mach.hash(), value_cache);
            checkRun(mach3);
        }
    }

    SECTION("Save twice, delete and load") {
        saveMachine(*tx, mach);
        saveMachine(*tx, mach);
        deleteMachine(*tx, mach.hash());
        REQUIRE(tx->commit().ok());
        auto mach2 = storage.getMachine(mach.hash(), value_cache);
        checkRun(mach2);
    }
}
