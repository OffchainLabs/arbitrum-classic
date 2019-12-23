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

#include <avm/machinestate/datastack.hpp>

#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/machinestatefetcher.hpp>
#include <data_storage/checkpoint/machinestatesaver.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/transaction.hpp>

#include <rocksdb/status.h>

#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

std::string dbpath =
    boost::filesystem::current_path().generic_string() + "/machineDb";

void initializeDatastack(MachineStateSaver& msSaver,
                         MachineStateFetcher& fetcher,
                         std::vector<unsigned char> hash_key,
                         uint256_t expected_hash,
                         int expected_size) {
    Datastack data_stack;
    data_stack.initializeDataStack(fetcher, hash_key);

    REQUIRE(data_stack.hash() == expected_hash);
    REQUIRE(data_stack.stacksize() == expected_size);
}

void saveDataStack(Datastack data_stack,
                   std::vector<unsigned char> expected_hash_key) {
    TuplePool pool;
    auto state = checkpoint::getInitialVmState(test_contract_path);
    CheckpointStorage storage(dbpath, state);
    std::vector<CodePoint> code;
    auto saver = MachineStateSaver(storage.makeTransaction());

    auto results = data_stack.checkpointState(saver, &pool);
    saver.commitTransaction();

    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
    REQUIRE(results.storage_key == expected_hash_key);
}

void saveDataStackTwice(Datastack data_stack,
                        std::vector<unsigned char> expected_hash_key) {
    TuplePool pool;
    auto state = checkpoint::getInitialVmState(test_contract_path);
    CheckpointStorage storage(dbpath, state);
    std::vector<CodePoint> code;
    auto saver = MachineStateSaver(storage.makeTransaction());

    auto results = data_stack.checkpointState(saver, &pool);
    auto results2 = data_stack.checkpointState(saver, &pool);

    saver.commitTransaction();

    REQUIRE(results2.status.ok());
    REQUIRE(results2.reference_count == 2);
    REQUIRE(results2.storage_key == expected_hash_key);
}

void saveAndGetDataStack(MachineStateSaver& saver,
                         MachineStateFetcher& fetcher,
                         Datastack data_stack,
                         std::vector<unsigned char> hash_key,
                         uint256_t expected_hash) {
    TuplePool pool;
    data_stack.checkpointState(saver, &pool);
    saver.commitTransaction();

    auto get_results = fetcher.getTuple(hash_key);

    REQUIRE(get_results.status.ok());
    REQUIRE(get_results.reference_count == 1);
    REQUIRE(get_results.data.calculateHash() == expected_hash);
}

void saveTwiceAndGetDataStack(MachineStateSaver& saver,
                              MachineStateFetcher& fetcher,
                              Datastack data_stack,
                              std::vector<unsigned char> hash_key,
                              uint256_t expected_hash) {
    TuplePool pool;

    data_stack.checkpointState(saver, &pool);
    data_stack.checkpointState(saver, &pool);
    saver.commitTransaction();

    auto get_results = fetcher.getTuple(hash_key);

    REQUIRE(get_results.status.ok());
    REQUIRE(get_results.reference_count == 2);
    REQUIRE(get_results.data.calculateHash() == expected_hash);
}

TEST_CASE("Initialize datastack") {
    SECTION("default") {
        TuplePool pool;
        auto state = checkpoint::getInitialVmState(test_contract_path);
        CheckpointStorage storage(dbpath, state);

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        Datastack data_stack;

        auto results = data_stack.checkpointState(saver, &pool);
        saver.commitTransaction();
        auto stack_hash = data_stack.hash();

        initializeDatastack(saver, fetcher, results.storage_key, stack_hash, 0);
    }
    boost::filesystem::remove_all(dbpath);

    SECTION("push empty tuple") {
        TuplePool pool;
        auto state = checkpoint::getInitialVmState(test_contract_path);
        CheckpointStorage storage(dbpath, state);

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        Datastack data_stack;
        Tuple tuple;
        data_stack.push(tuple);

        auto results = data_stack.checkpointState(saver, &pool);
        saver.commitTransaction();
        auto stack_hash = data_stack.hash();

        initializeDatastack(saver, fetcher, results.storage_key, stack_hash, 1);
    }
    boost::filesystem::remove_all(dbpath);

    SECTION("push num, tuple") {
        TuplePool pool;
        auto state = checkpoint::getInitialVmState(test_contract_path);
        CheckpointStorage storage(dbpath, state);

        CodePoint code_point = state.code[0];
        CodePoint code_point2 = state.code[1];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        Datastack data_stack;
        uint256_t num = 1;
        auto tuple = Tuple(code_point, &pool);

        data_stack.push(num);
        data_stack.push(tuple);

        auto results = data_stack.checkpointState(saver, &pool);
        saver.commitTransaction();
        auto stack_hash = data_stack.hash();

        initializeDatastack(saver, fetcher, results.storage_key, stack_hash, 2);
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("push codepoint, tuple") {
        TuplePool pool;
        auto state = checkpoint::getInitialVmState(test_contract_path);
        CheckpointStorage storage(dbpath, state);

        CodePoint code_point = state.code[0];
        CodePoint code_point2 = state.code[1];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        Datastack data_stack;

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);

        data_stack.push(code_point2);
        data_stack.push(tuple);

        auto results = data_stack.checkpointState(saver, &pool);
        saver.commitTransaction();
        auto stack_hash = data_stack.hash();

        initializeDatastack(saver, fetcher, results.storage_key, stack_hash, 2);
    }
    boost::filesystem::remove_all(dbpath);
}

TEST_CASE("Save datastack") {
    SECTION("save empty") {
        Datastack datastack;
        std::vector<unsigned char> hash_key_vector;
        marshal_value(Tuple().calculateHash(), hash_key_vector);

        saveDataStack(datastack, hash_key_vector);
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("save empty twice") {
        Datastack datastack;
        std::vector<unsigned char> hash_key_vector;
        marshal_value(Tuple().calculateHash(), hash_key_vector);

        saveDataStackTwice(datastack, hash_key_vector);
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("save with values") {
        TuplePool pool;

        uint256_t num = 1;
        auto code_point = CodePoint(1, Operation(), 0);
        auto tuple = Tuple(code_point, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        auto tup1 = Tuple(num, &pool);
        auto tup_rep = Tuple(tuple, tup1, &pool);

        std::vector<unsigned char> hash_key_vector;
        marshal_value(tup_rep.calculateHash(), hash_key_vector);

        saveDataStack(datastack, hash_key_vector);
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("save with values, twice") {
        TuplePool pool;

        uint256_t num = 1;
        auto code_point = CodePoint(1, Operation(), 0);
        auto tuple = Tuple(code_point, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        auto tup1 = Tuple(num, &pool);
        auto tup_rep = Tuple(tuple, tup1, &pool);

        std::vector<unsigned char> hash_key_vector;
        marshal_value(tup_rep.calculateHash(), hash_key_vector);

        saveDataStackTwice(datastack, hash_key_vector);
    }
    boost::filesystem::remove_all(dbpath);
}

TEST_CASE("Save and get datastack") {
    SECTION("save datastack and get") {
        TuplePool pool;
        auto state = checkpoint::getInitialVmState(test_contract_path);
        CheckpointStorage storage(dbpath, state);

        auto code_point = state.code[0];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto tuple = Tuple(code_point, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        auto tup1 = Tuple(num, &pool);
        auto tup_rep = Tuple(tuple, tup1, &pool);

        std::vector<unsigned char> hash_key_vector;
        marshal_value(tup_rep.calculateHash(), hash_key_vector);

        saveAndGetDataStack(saver, fetcher, datastack, hash_key_vector,
                            tup_rep.calculateHash());
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("save datastack twice and get") {
        TuplePool pool;
        auto state = checkpoint::getInitialVmState(test_contract_path);
        CheckpointStorage storage(dbpath, state);

        auto code_point = state.code[0];

        auto saver = MachineStateSaver(storage.makeTransaction());
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto tuple = Tuple(code_point, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        auto tup1 = Tuple(num, &pool);
        auto tup_rep = Tuple(tuple, tup1, &pool);

        std::vector<unsigned char> hash_key_vector;
        marshal_value(tup_rep.calculateHash(), hash_key_vector);

        saveTwiceAndGetDataStack(saver, fetcher, datastack, hash_key_vector,
                                 tup_rep.calculateHash());
    }
    boost::filesystem::remove_all(dbpath);
}
