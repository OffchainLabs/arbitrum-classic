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

#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

auto execution_path = boost::filesystem::current_path();

void checkpointState(CheckpointStorage& storage, Machine& machine) {
    auto transaction = storage.makeTransaction();
    auto results = saveMachine(*transaction, machine);
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
    REQUIRE(transaction->commit().ok());
}

void checkpointStateTwice(CheckpointStorage& storage, Machine& machine) {
    auto transaction1 = storage.makeTransaction();
    auto results = saveMachine(*transaction1, machine);
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
    REQUIRE(transaction1->commit().ok());

    auto transaction2 = storage.makeTransaction();
    auto results2 = saveMachine(*transaction2, machine);
    REQUIRE(results2.status.ok());
    REQUIRE(results2.reference_count == 2);
    REQUIRE(transaction2->commit().ok());
}

void deleteCheckpoint(Transaction& transaction, Machine& machine) {
    auto results = deleteMachine(transaction, machine.hash());
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 0);
}

void restoreCheckpoint(CheckpointStorage& storage, Machine& expected_machine) {
    auto mach = storage.getMachine(expected_machine.hash());
    REQUIRE(mach.hash() == expected_machine.hash());
}

TEST_CASE("Checkpoint State") {
    DBDeleter deleter;
    TuplePool pool;
    CheckpointStorage storage(dbpath);
    storage.initialize(test_contract_path);

    auto machine = storage.getInitialMachine();
    machine.run(1, TimeBounds{}, Tuple(), std::chrono::seconds{0});

    SECTION("default") { checkpointState(storage, machine); }
    SECTION("save twice") { checkpointStateTwice(storage, machine); }
    SECTION("assert machine hash") {
        auto hash1 = machine.hash();
        auto transaction = storage.makeTransaction();
        auto results = saveMachine(*transaction, machine);
        REQUIRE(results.status.ok());
        REQUIRE(transaction->commit().ok());
        auto machine2 = storage.getMachine(hash1);
        auto hash2 = machine2.hash();
        REQUIRE(hash2 == hash1);
    }
}

TEST_CASE("Delete machine checkpoint") {
    DBDeleter deleter;
    TuplePool pool;
    CheckpointStorage storage(dbpath);
    storage.initialize(test_contract_path);

    SECTION("default") {
        auto machine = storage.getInitialMachine();
        machine.run(1, TimeBounds{}, Tuple(), std::chrono::seconds{0});
        auto transaction = storage.makeTransaction();
        auto results = saveMachine(*transaction, machine);
        deleteCheckpoint(*transaction, machine);
        REQUIRE(transaction->commit().ok());
    }
}

TEST_CASE("Restore checkpoint") {
    DBDeleter deleter;
    TuplePool pool;
    CheckpointStorage storage(dbpath);
    storage.initialize(test_contract_path);

    SECTION("default") {
        auto machine = storage.getInitialMachine();
        auto transaction = storage.makeTransaction();
        auto results = saveMachine(*transaction, machine);
        REQUIRE(results.status.ok());
        REQUIRE(transaction->commit().ok());
        restoreCheckpoint(storage, machine);
    }
}

TEST_CASE("Proof") {
    auto machine = Machine::loadFromFile(test_contract_path);
    while (true) {
        auto assertion = machine.run(1, {}, {}, std::chrono::seconds{0});
        machine.marshalForProof();
        if (assertion.stepCount == 0) {
            break;
        }
    }
}

TEST_CASE("Clone") {
    auto machine = Machine::loadFromFile(test_contract_path);
    for (int i = 0; i < 100; i++) {
        machine.machine_state.stack.push(Tuple(uint256_t{3}, uint256_t{6},
                                               uint256_t{7}, uint256_t{2},
                                               &machine.getPool()));
        machine.machine_state.auxstack.push(Tuple(uint256_t{3}, uint256_t{6},
                                                  uint256_t{7}, uint256_t{2},
                                                  &machine.getPool()));
    }

    for (int i = 0; i < 1000; i++) {
        Machine m = machine;
        REQUIRE(m.hash() != 3242);
    }
}

TEST_CASE("Machine hash") {
    MachineState machine = MachineState::loadFromFile(test_contract_path);
    auto pcHash = ::hash(machine.loadCurrentInstruction());
    auto stackHash = machine.stack.hash();
    auto auxstackHash = machine.auxstack.hash();
    auto registerHash = ::hash_value(machine.registerVal);
    auto staticHash = ::hash_value(machine.static_val);
    auto errHash = ::hash(machine.errpc);
    auto machineHash = machine.hash();

    REQUIRE(pcHash == uint256_t("9437065110668622075464824926507979877827393212"
                                "8194559331492955050019282050496"));
    REQUIRE(stackHash == uint256_t("4251290975118555612292311539115420848775231"
                                   "0613213055089416300774052282720344"));
    REQUIRE(auxstackHash == uint256_t("4251290975118555612292311539115420848775"
                                      "2310613213055089416300774052282720344"));
    REQUIRE(registerHash == uint256_t("4251290975118555612292311539115420848775"
                                      "2310613213055089416300774052282720344"));
    REQUIRE(staticHash == uint256_t("113182352889449210665994027227588754290969"
                                    "798016938687372921424809289618385856"));
    REQUIRE(errHash == uint256_t("817555893843236912662725763451298816577059146"
                                 "21008081459572116739688988488432"));
    REQUIRE(machineHash == uint256_t("56208326812724912066026123588383649819390"
                                     "601658448049319166841561743369815863"));
}

TEST_CASE("MachineTestVectors") {
    DBDeleter deleter;
    TuplePool pool;

    std::vector<std::string> files = {
        "opcodetestarbgas",   "opcodetestdup",   "opcodetestecrecover",
        "opcodetestethhash2", "opcodetesthash",  "opcodetestlogic",
        "opcodetestmath",     "opcodeteststack", "opcodetesttuple",
        "opcodetestcode"};

    for (const auto& filename : files) {
        DYNAMIC_SECTION(filename) {
            auto test_file =
                std::string{machine_test_cases_path} + "/" + filename + ".mexe";

            auto mach = Machine::loadFromFile(test_file);
            while (
                nonstd::holds_alternative<NotBlocked>(mach.isBlocked(false))) {
                mach.run(1, TimeBounds{}, Tuple(), std::chrono::seconds{0});
            }
            REQUIRE(mach.currentStatus() == Status::Halted);
        }
    }
}
