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
    auto transaction = storage.makeReadWriteTransaction();
    auto results = saveTestMachine(*transaction, machine);
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
    REQUIRE(transaction->commit().ok());
}

void checkpointStateTwice(ArbStorage& storage, Machine& machine) {
    auto transaction1 = storage.makeReadWriteTransaction();
    auto results = saveTestMachine(*transaction1, machine);
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
    REQUIRE(transaction1->commit().ok());

    auto transaction2 = storage.makeReadWriteTransaction();
    auto results2 = saveTestMachine(*transaction2, machine);
    REQUIRE(results2.status.ok());
    REQUIRE(results2.reference_count == 2);
    REQUIRE(transaction2->commit().ok());
}

void deleteCheckpoint(ReadWriteTransaction& transaction, Machine& machine) {
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
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    REQUIRE(storage.initialize(test_contract_path).ok());
    ValueCache value_cache{1, 0};

    auto machine = storage.getInitialMachine();
    MachineExecutionConfig execConfig;
    execConfig.max_gas = 3;
    machine->machine_state.context = AssertionContext(execConfig);
    machine->run();

    SECTION("default") { checkpointState(storage, *machine); }
    SECTION("save twice") { checkpointStateTwice(storage, *machine); }
    SECTION("assert machine hash") {
        auto transaction = storage.makeReadWriteTransaction();
        auto results = saveTestMachine(*transaction, *machine);
        REQUIRE(results.status.ok());
        REQUIRE(transaction->commit().ok());
        auto machine2 = storage.getMachine(machine->hash(), value_cache);
        REQUIRE(machine2->hash() == machine->hash());
    }
}

TEST_CASE("Delete machine checkpoint") {
    DBDeleter deleter;
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    REQUIRE(storage.initialize(test_contract_path).ok());

    SECTION("default") {
        auto machine = storage.getInitialMachine();
        MachineExecutionConfig execConfig;
        execConfig.max_gas = 4;
        machine->machine_state.context = AssertionContext(execConfig);
        machine->run();
        auto transaction = storage.makeReadWriteTransaction();
        saveTestMachine(*transaction, *machine);
        execConfig.max_gas = 0;
        machine->machine_state.context = AssertionContext(execConfig);
        machine->run();
        saveTestMachine(*transaction, *machine);
        deleteCheckpoint(*transaction, *machine);
        REQUIRE(transaction->commit().ok());
    }
}

TEST_CASE("Restore checkpoint") {
    DBDeleter deleter;
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    REQUIRE(storage.initialize(test_contract_path).ok());
    ValueCache value_cache{1, 0};

    SECTION("default") {
        auto machine = storage.getInitialMachine();
        auto transaction = storage.makeReadWriteTransaction();
        auto results = saveTestMachine(*transaction, *machine);
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
        machine.machine_state.context = AssertionContext(execConfig);
        auto assertion = machine.run();
        machine.marshalForProof();
        if (assertion.step_count == 0) {
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
                               "56208326812724912066026123588383649819390601658"
                               "448049319166841561743369815863"));
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
                mach.run();
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
    machine.machine_state.context = AssertionContext(execConfig);
    auto assertion = machine.run();
    REQUIRE(machine.currentStatus() == Status::Error);
    REQUIRE(!assertion.sideload_block_number);
    REQUIRE(assertion.gas_count == 13);

    // Next, test running past the sideload with a value specified
    machine = orig_machine;
    execConfig.sideloads.emplace_back(InboxMessage());
    execConfig.stop_on_sideload = true;  // Shouldn't matter
    machine.machine_state.context = AssertionContext(execConfig);
    assertion = machine.run();
    REQUIRE(machine.currentStatus() == Status::Halted);
    REQUIRE(!assertion.sideload_block_number);
    REQUIRE(assertion.gas_count == 23);

    // Next, test stopping on the sideload but continuing
    machine = orig_machine;
    execConfig.sideloads.clear();
    execConfig.stop_on_sideload = true;
    machine.machine_state.context = AssertionContext(execConfig);
    assertion = machine.run();
    REQUIRE(machine.currentStatus() == Status::Extensive);
    REQUIRE(assertion.sideload_block_number == uint256_t(0x321));
    REQUIRE(assertion.gas_count == 1);
    machine.machine_state.context = AssertionContext(execConfig);
    assertion = machine.run();
    REQUIRE(machine.currentStatus() == Status::Error);
    REQUIRE(!assertion.sideload_block_number);
    REQUIRE(assertion.gas_count == 12);

    // Next, test stopping on the sideload and adding a value
    machine = orig_machine;
    machine.machine_state.context = AssertionContext(execConfig);
    assertion = machine.run();
    REQUIRE(machine.currentStatus() == Status::Extensive);
    REQUIRE(assertion.sideload_block_number == uint256_t(0x321));
    REQUIRE(assertion.gas_count == 1);

    execConfig.sideloads.emplace_back(InboxMessage());
    machine.machine_state.context = AssertionContext(execConfig);
    assertion = machine.run();
    REQUIRE(machine.currentStatus() == Status::Halted);
    REQUIRE(!assertion.sideload_block_number);
    REQUIRE(assertion.gas_count == 22);
}
