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

#include <avm/machine.hpp>

#include <boost/dll.hpp>
#include <catch2/catch.hpp>

auto execution_path = boost::dll::program_location().parent_path();
auto save_path = execution_path.generic_string() + "/rocksDb";
auto contract_path =
    execution_path.parent_path().parent_path().generic_string() +
    "/tests/contract.ao";

void checkpointState(CheckpointStorage& storage, Machine& machine) {
    auto results = machine.checkpoint(storage);

    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
    REQUIRE(results.storage_key == GetHashKey(machine.hash()));
}

void checkpointStateTwice(CheckpointStorage& storage, Machine& machine) {
    auto results = machine.checkpoint(storage);
    auto results2 = machine.checkpoint(storage);

    REQUIRE(results2.status.ok());
    REQUIRE(results2.reference_count == 2);
    REQUIRE(results2.storage_key == GetHashKey(machine.hash()));
}

void deleteCheckpoint(CheckpointStorage& storage,
                      Machine& machine,
                      const std::vector<unsigned char>& checkpoint_key) {
    auto results = machine.deleteCheckpoint(storage, checkpoint_key);
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 0);
}

void restoreCheckpoint(CheckpointStorage& storage,
                       Machine& expected_machine,
                       const std::vector<unsigned char>& checkpoint_key) {
    Machine machine;
    machine.initializeMachine(contract_path);
    auto success = machine.restoreCheckpoint(storage, checkpoint_key);

    REQUIRE(success);
    REQUIRE(machine.hash() == expected_machine.hash());
}

TEST_CASE("Checkpoint State") {
    SECTION("default") {
        CheckpointStorage storage(save_path);
        Machine machine;
        machine.initializeMachine(contract_path);

        checkpointState(storage, machine);
    }
    SECTION("save twice") {
        CheckpointStorage storage(save_path);
        Machine machine;
        machine.initializeMachine(contract_path);

        checkpointStateTwice(storage, machine);
    }
}

TEST_CASE("Delete machine checkpoint") {
    SECTION("default") {
        CheckpointStorage storage(save_path);
        Machine machine;
        machine.initializeMachine(contract_path);
        auto results = machine.checkpoint(storage);

        deleteCheckpoint(storage, machine, results.storage_key);
    }
}

TEST_CASE("Restore checkpoint fails") {
    SECTION("default") {
        TuplePool pool;
        CheckpointStorage storage(save_path);
        Machine machine;
        machine.initializeMachine(contract_path);
        auto results = machine.checkpoint(storage);

        restoreCheckpoint(storage, machine, results.storage_key);
    }
}

// testing
//
//        auto saver = MachineStateSaver(&storage, &pool,
//        machine.machine_state.code);
//        auto res =
//        saver.saveValue(machine.machine_state.staticVal);
//        auto res2 =
//        saver.getValue(res.storage_key);
//
//        auto tup1 = nonstd::get<Tuple>(machine.machine_state.staticVal);
//        auto tup2 = nonstd::get<Tuple>(res2.data);
//
//        for(auto i = 0 ; i < tup1.tuple_size(); i++){
//
//            auto item = nonstd::get<CodePoint>(tup1.get_element(i));
//            auto saved_item = nonstd::get<CodePoint>(tup2.get_element(i));
//            auto actual_item = machine.machine_state.code[item.pc];
//
//
//            REQUIRE(item.pc == saved_item.pc);
//            auto actual_nexthash =
//            machine.machine_state.code[item.pc].nextHash;
//            REQUIRE(actual_nexthash ==
//            hash(machine.machine_state.code[item.pc+1]));
//
//            REQUIRE(actual_nexthash == saved_item.nextHash);
//            REQUIRE(actual_nexthash == item.nextHash);
//            REQUIRE(item.nextHash == saved_item.nextHash);
//            REQUIRE(hash(item) == hash(saved_item));
//
//        }
//
//        restoreCheckpoint(storage, machine, results.storage_key);
//    }
//}
