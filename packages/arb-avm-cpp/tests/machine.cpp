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

#include <data_storage/arbstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machine.hpp>

#define CATCH_CONFIG_ENABLE_BENCHMARKING 1
#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

auto execution_path = boost::filesystem::current_path();

void checkpointState(ArbStorage& storage, Machine& machine) {
    auto transaction = storage.makeTransaction();
    auto results = saveMachine(*transaction, machine);
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
    REQUIRE(transaction->commit().ok());
}

void checkpointStateTwice(ArbStorage& storage, Machine& machine) {
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

void restoreCheckpoint(ArbStorage& storage,
                       Machine& expected_machine,
                       ValueCache& value_cache) {
    auto mach = storage.getMachine(expected_machine.hash(), value_cache);
    REQUIRE(mach->hash() == expected_machine.hash());
}

TEST_CASE("Checkpoint State") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    REQUIRE(storage.initialize(test_contract_path).ok());
    ValueCache value_cache{};

    auto machine = storage.getInitialMachine(value_cache);
    MachineExecutionConfig execConfig;
    execConfig.max_gas = 3;
    execConfig.next_block_height = 3;
    machine->run(execConfig);

    SECTION("default") { checkpointState(storage, *machine); }
    SECTION("save twice") { checkpointStateTwice(storage, *machine); }
    SECTION("assert machine hash") {
        auto hash1 = machine->hash();
        auto transaction = storage.makeTransaction();
        auto results = saveMachine(*transaction, *machine);
        REQUIRE(results.status.ok());
        REQUIRE(transaction->commit().ok());
        auto machine2 = storage.getMachine(hash1, value_cache);
        auto hash2 = machine2->hash();
        REQUIRE(hash2 == hash1);
    }
}

TEST_CASE("Delete machine checkpoint") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    REQUIRE(storage.initialize(test_contract_path).ok());
    ValueCache value_cache{};

    SECTION("default") {
        auto machine = storage.getInitialMachine(value_cache);
        MachineExecutionConfig execConfig;
        execConfig.max_gas = 4;
        execConfig.next_block_height = 3;
        machine->run(execConfig);
        auto transaction = storage.makeTransaction();
        saveMachine(*transaction, *machine);
        execConfig.max_gas = 0;
        machine->run(execConfig);
        saveMachine(*transaction, *machine);
        deleteCheckpoint(*transaction, *machine);
        REQUIRE(transaction->commit().ok());
    }
}

TEST_CASE("Restore checkpoint") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    REQUIRE(storage.initialize(test_contract_path).ok());
    ValueCache value_cache{};

    SECTION("default") {
        auto machine = storage.getInitialMachine(value_cache);
        auto transaction = storage.makeTransaction();
        auto results = saveMachine(*transaction, *machine);
        REQUIRE(results.status.ok());
        REQUIRE(transaction->commit().ok());
        restoreCheckpoint(storage, *machine, value_cache);
    }
}

TEST_CASE("Proof") {
    auto machine = Machine::loadFromFile(test_contract_path);
    while (true) {
        MachineExecutionConfig execConfig;
        execConfig.max_gas = 3;
        execConfig.next_block_height = 3;
        auto assertion = machine.run(execConfig);
        machine.marshalForProof();
        if (assertion.stepCount == 0) {
            break;
        }
    }
}

TEST_CASE("Clone") {
    auto machine = Machine::loadFromFile(test_contract_path);
    for (int i = 0; i < 100; i++) {
        machine.machine_state.stack.push(
            Tuple(uint256_t{3}, uint256_t{6}, uint256_t{7}, uint256_t{2}));
        machine.machine_state.auxstack.push(
            Tuple(uint256_t{3}, uint256_t{6}, uint256_t{7}, uint256_t{2}));
    }

    for (int i = 0; i < 1000; i++) {
        REQUIRE(machine.hash() != 3242);
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

    REQUIRE(pcHash == intx::from_string<uint256_t>(
                          "9437065110668622075464824926507979877827393212819455"
                          "9331492955050019282050496"));
    REQUIRE(stackHash == intx::from_string<uint256_t>(
                             "4251290975118555612292311539115420848775231061321"
                             "3055089416300774052282720344"));
    REQUIRE(auxstackHash == intx::from_string<uint256_t>(
                                "4251290975118555612292311539115420848775231061"
                                "3213055089416300774052282720344"));
    REQUIRE(registerHash == intx::from_string<uint256_t>(
                                "4251290975118555612292311539115420848775231061"
                                "3213055089416300774052282720344"));
    REQUIRE(staticHash == intx::from_string<uint256_t>(
                              "113182352889449210665994027227588754290969798016"
                              "938687372921424809289618385856"));
    REQUIRE(errHash == intx::from_string<uint256_t>(
                           "817555893843236912662725763451298816577059146210080"
                           "81459572116739688988488432"));
    REQUIRE(machineHash == intx::from_string<uint256_t>(
                               "12818298244055256727021237632105567892940754157"
                               "945856618644698870485816765145"));
}

TEST_CASE("MachineTestVectors") {
    DBDeleter deleter;

    std::vector<std::string> files = {
        "opcodetestarbgas",   "opcodetestdup",     "opcodetestecops",
        "opcodetestethhash2", "opcodetesthash",    "opcodetestlogic",
        "opcodetestmath",     "opcodeteststack",   "opcodetesttuple",
        "opcodetestcode",     "opcodetestkeccakf", "opcodetestsha256f"};

    for (const auto& filename : files) {
        DYNAMIC_SECTION(filename) {
            auto test_file =
                std::string{machine_test_cases_path} + "/" + filename + ".mexe";

            auto mach = Machine::loadFromFile(test_file);
            MachineExecutionConfig execConfig;
            while (std::holds_alternative<NotBlocked>(mach.isBlocked(false))) {
                mach.run(execConfig);
            }
            REQUIRE(mach.currentStatus() == Status::Halted);
        }
    }
}

TEST_CASE("Stopping on sideload") {
    auto orig_machine = Machine::loadFromFile(
        std::string(machine_test_cases_path) + "/sideloadtest.mexe");
    MachineExecutionConfig execConfig;

    // First, test running straight past the sideload
    Machine machine = orig_machine;
    auto assertion = machine.run(execConfig);
    REQUIRE(machine.currentStatus() == Status::Error);
    REQUIRE(!assertion.sideloadBlockNumber);
    REQUIRE(assertion.gasCount == 13);

    // Next, test running past the sideload with a value specified
    machine = orig_machine;
    execConfig.sideloads.emplace_back(InboxMessage());
    execConfig.stop_on_sideload = true;  // Shouldn't matter
    assertion = machine.run(execConfig);
    REQUIRE(machine.currentStatus() == Status::Halted);
    REQUIRE(!assertion.sideloadBlockNumber);
    REQUIRE(assertion.gasCount == 23);

    // Next, test stopping on the sideload but continuing
    machine = orig_machine;
    execConfig.sideloads.clear();
    execConfig.stop_on_sideload = true;
    assertion = machine.run(execConfig);
    REQUIRE(machine.currentStatus() == Status::Extensive);
    REQUIRE(assertion.sideloadBlockNumber == uint256_t(0x321));
    REQUIRE(assertion.gasCount == 1);

    assertion = machine.run(execConfig);
    REQUIRE(machine.currentStatus() == Status::Error);
    REQUIRE(!assertion.sideloadBlockNumber);
    REQUIRE(assertion.gasCount == 12);

    // Next, test stopping on the sideload and adding a value
    machine = orig_machine;
    assertion = machine.run(execConfig);
    REQUIRE(machine.currentStatus() == Status::Extensive);
    REQUIRE(assertion.sideloadBlockNumber == uint256_t(0x321));
    REQUIRE(assertion.gasCount == 1);

    execConfig.sideloads.emplace_back(InboxMessage());
    assertion = machine.run(execConfig);
    REQUIRE(machine.currentStatus() == Status::Halted);
    REQUIRE(!assertion.sideloadBlockNumber);
    REQUIRE(assertion.gasCount == 22);
}
