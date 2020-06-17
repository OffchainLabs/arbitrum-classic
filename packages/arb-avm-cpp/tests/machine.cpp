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
    auto ret = storage.getMachine(expected_machine.hash());
    REQUIRE(ret.second);
    REQUIRE(ret.first.hash() == expected_machine.hash());
}

TEST_CASE("Checkpoint State") {
    DBDeleter deleter;
    TuplePool pool;
    CheckpointStorage storage(dbpath, test_contract_path);

    SECTION("default") {
        auto ret = Machine::loadFromFile(test_contract_path);
        REQUIRE(ret.second);
        checkpointState(storage, ret.first);
    }
    SECTION("save twice") {
        auto ret = Machine::loadFromFile(test_contract_path);
        REQUIRE(ret.second);
        checkpointStateTwice(storage, ret.first);
    }
    SECTION("assert machine hash") {
        auto ret = Machine::loadFromFile(test_contract_path);
        REQUIRE(ret.second);
        auto machine = std::move(ret.first);
        Machine machine2 = storage.getInitialMachine();
        auto hash1 = machine.hash();
        auto hash2 = machine2.hash();
        auto transaction = storage.makeTransaction();
        auto results = saveMachine(*transaction, machine);
        REQUIRE(results.status.ok());
        REQUIRE(transaction->commit().ok());
        auto ret2 = storage.getMachine(hash1);
        REQUIRE(ret2.second);
        auto machine3 = std::move(ret2.first);
        auto hash3 = machine3.hash();
        REQUIRE(hash3 == hash2);
        REQUIRE(hash1 == hash2);
    }
}

TEST_CASE("Delete machine checkpoint") {
    DBDeleter deleter;
    TuplePool pool;
    CheckpointStorage storage(dbpath, test_contract_path);

    SECTION("default") {
        auto ret = Machine::loadFromFile(test_contract_path);
        REQUIRE(ret.second);
        auto transaction = storage.makeTransaction();
        auto results = saveMachine(*transaction, ret.first);
        deleteCheckpoint(*transaction, ret.first);
        REQUIRE(transaction->commit().ok());
    }
}

TEST_CASE("Restore checkpoint") {
    DBDeleter deleter;
    TuplePool pool;
    CheckpointStorage storage(dbpath, test_contract_path);

    SECTION("default") {
        auto ret = Machine::loadFromFile(test_contract_path);
        REQUIRE(ret.second);
        auto transaction = storage.makeTransaction();
        auto results = saveMachine(*transaction, ret.first);
        REQUIRE(results.status.ok());
        REQUIRE(transaction->commit().ok());
        restoreCheckpoint(storage, ret.first);
    }
}

TEST_CASE("Proof") {
    auto ret = Machine::loadFromFile(test_contract_path);
    REQUIRE(ret.second);
    while (true) {
        auto assertion = ret.first.run(1, {}, {}, std::chrono::seconds{0});
        ret.first.marshalForProof();
        if (assertion.stepCount == 0) {
            break;
        }
    }
}

TEST_CASE("Clone") {
    auto ret = Machine::loadFromFile(test_contract_path);
    REQUIRE(ret.second);
    auto machine = std::move(ret.first);
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
    DBDeleter deleter;
    TuplePool pool;
    CheckpointStorage storage(dbpath, test_contract_path);
    MachineState machine = MachineState::loadFromFile(test_contract_path).first;
    auto pcHash = ::hash(machine.static_values->code[machine.pc]);
    auto stackHash = machine.stack.hash();
    auto auxstackHash = machine.auxstack.hash();
    auto registerHash = ::hash_value(machine.registerVal);
    auto staticHash = ::hash_value(machine.static_values->staticVal);
    auto errHash = ::hash(machine.static_values->code[machine.errpc]);
    auto machineHash = machine.hash();

    REQUIRE(pcHash == uint256_t("7737343943613437755395141441291826898796163866"
                                "8492301359045565991555588221763"));
    REQUIRE(stackHash == uint256_t("4251290975118555612292311539115420848775231"
                                   "0613213055089416300774052282720344"));
    REQUIRE(auxstackHash == uint256_t("4251290975118555612292311539115420848775"
                                      "2310613213055089416300774052282720344"));
    REQUIRE(registerHash == uint256_t("4251290975118555612292311539115420848775"
                                      "2310613213055089416300774052282720344"));
    REQUIRE(staticHash == uint256_t("832315546794065743621037540472311669512067"
                                    "66645942016669157541405145405171869"));
    REQUIRE(errHash == uint256_t("817555893843236912662725763451298816577059146"
                                 "21008081459572116739688988488432"));
    REQUIRE(machineHash == uint256_t("38086450045779233370084791113969759535500"
                                     "380524553750909191655579112918186895"));
}
