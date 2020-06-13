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

#include <avm/machine.hpp>
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/machinestatedeleter.hpp>
#include <data_storage/storageresult.hpp>

#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

auto execution_path = boost::filesystem::current_path();

void checkpointState(CheckpointStorage& storage, Machine& machine) {
    auto results = storage.saveMachine(machine);

    std::vector<unsigned char> hash_vector;
    marshal_uint256_t(machine.hash(), hash_vector);

    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
    REQUIRE(results.storage_key == hash_vector);
}

void checkpointStateTwice(CheckpointStorage& storage, Machine& machine) {
    auto results = storage.saveMachine(machine);
    auto results2 = storage.saveMachine(machine);

    std::vector<unsigned char> hash_vector;
    marshal_uint256_t(machine.hash(), hash_vector);

    REQUIRE(results2.status.ok());
    REQUIRE(results2.reference_count == 2);
    REQUIRE(results2.storage_key == hash_vector);
}

void deleteCheckpoint(Transaction& transaction,
                      Machine& machine,
                      const std::vector<unsigned char>& checkpoint_key) {
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
    SECTION("default") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto ret = Machine::loadFromFile(test_contract_path);
        REQUIRE(ret.second);

        checkpointState(storage, ret.first);
    }
    SECTION("save twice") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto ret = Machine::loadFromFile(test_contract_path);
        REQUIRE(ret.second);

        checkpointStateTwice(storage, ret.first);
    }
    SECTION("assert machine hash") {
        DBDeleter deleter;
        auto ret = Machine::loadFromFile(test_contract_path);
        REQUIRE(ret.second);
        auto machine = std::move(ret.first);

        CheckpointStorage storage(dbpath, test_contract_path);
        auto initial_machine = storage.getInitialVmValues();
        Machine machine2(initial_machine.code, initial_machine.staticVal,
                         storage.pool);

        auto hash1 = machine.hash();
        auto hash2 = machine2.hash();

        storage.saveMachine(machine);

        auto ret2 = storage.getMachine(hash1);
        REQUIRE(ret2.second);
        auto machine3 = std::move(ret2.first);

        auto hash3 = machine3.hash();

        REQUIRE(hash3 == hash2);
        REQUIRE(hash1 == hash2);
    }
}

TEST_CASE("Delete machine checkpoint") {
    SECTION("default") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto ret = Machine::loadFromFile(test_contract_path);
        REQUIRE(ret.second);
        auto results = storage.saveMachine(ret.first);
        auto transaction = storage.makeTransaction();
        deleteCheckpoint(*transaction, ret.first, results.storage_key);
        REQUIRE(transaction->commit().ok());
    }
}

TEST_CASE("Restore checkpoint") {
    SECTION("default") {
        DBDeleter deleter;
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);
        auto ret = Machine::loadFromFile(test_contract_path);
        REQUIRE(ret.second);
        auto results = storage.saveMachine(ret.first);

        restoreCheckpoint(storage, ret.first);
    }
}
