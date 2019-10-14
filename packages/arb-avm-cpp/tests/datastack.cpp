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

#include "avm/machinestate/datastack.hpp"
#include <catch2/catch.hpp>
#include "avm/machinestate/machinestatesaver.hpp"

std::string dbpath =
    "/Users/minhtruong/Dev/arbitrum/packages/arb-avm-cpp/build/tests/rocksDb";

void initializeDatastack(Tuple tuple,
                         uint256_t expected_hash,
                         int expected_size) {
    Datastack data_stack(tuple);

    REQUIRE(data_stack.hash() == expected_hash);
    REQUIRE(data_stack.stacksize() == expected_size);
}

TEST_CASE("Initialize datastack") {
    TuplePool pool;
    SECTION("Empty stack") {
        Datastack data_stack;
        auto stack_hash = data_stack.hash();

        Tuple tuple;

        initializeDatastack(tuple, stack_hash, 0);
    }

    SECTION("push empty tuple") {
        Datastack data_stack;
        Tuple tuple;
        data_stack.push(tuple);
        auto stack_hash = data_stack.hash();

        auto tuple_rep = Tuple(tuple, &pool);

        initializeDatastack(tuple_rep, stack_hash, 1);
    }

    SECTION("push num, tuple") {
        TuplePool pool;
        Datastack data_stack;

        uint256_t num = 1;
        auto code_point = CodePoint(1, Operation(), 0);
        auto tuple = Tuple(code_point, &pool);

        auto tup1 = Tuple(num, &pool);
        auto tup2 = Tuple(tuple, tup1, &pool);

        data_stack.push(num);
        data_stack.push(tuple);

        auto stack_hash = data_stack.hash();

        initializeDatastack(tup2, stack_hash, 2);
    }
    SECTION("push codepoint, tuple") {
        TuplePool pool;
        Datastack data_stack;

        auto code_point = CodePoint(1, Operation(), 0);
        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);

        auto tup1 = Tuple(code_point, &pool);
        auto tup2 = Tuple(tuple, tup1, &pool);

        data_stack.push(code_point);
        data_stack.push(tuple);

        auto stack_hash = data_stack.hash();

        initializeDatastack(tup2, stack_hash, 2);
    }
}

void saveDataStack(Datastack data_stack,
                   int expected_ref_count,
                   std::vector<unsigned char> expected_hash_key) {
    TuplePool pool;
    CheckpointStorage storage(dbpath);
    auto saver = MachineStateSaver(&storage, &pool);

    auto results = data_stack.checkpointState(saver, &pool);

    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(results.storage_key == expected_hash_key);
}

void saveDataStackTwice(Datastack data_stack,
                        int expected_ref_count,
                        std::vector<unsigned char> expected_hash_key) {
    TuplePool pool;
    CheckpointStorage storage(dbpath);
    auto saver = MachineStateSaver(&storage, &pool);
    auto results = data_stack.checkpointState(saver, &pool);
    auto results2 = data_stack.checkpointState(saver, &pool);

    REQUIRE(results2.status.ok());
    REQUIRE(results2.reference_count == expected_ref_count);
    REQUIRE(results2.storage_key == expected_hash_key);
}

void saveAndGetDataStack(Datastack data_stack,
                         std::vector<unsigned char> hash_key,
                         int expected_ref_count,
                         uint256_t expected_hash) {
    TuplePool pool;
    CheckpointStorage storage(dbpath);
    auto saver = MachineStateSaver(&storage, &pool);

    data_stack.checkpointState(saver, &pool);
    auto get_results = saver.getTuple(hash_key);

    REQUIRE(get_results.status.ok());
    REQUIRE(get_results.reference_count == expected_ref_count);
    REQUIRE(get_results.tuple.calculateHash() == expected_hash);
}

void saveTwiceAndGetDataStack(Datastack data_stack,
                              std::vector<unsigned char> hash_key,
                              int expected_ref_count,
                              uint256_t expected_hash) {
    TuplePool pool;
    CheckpointStorage storage(dbpath);
    auto saver = MachineStateSaver(&storage, &pool);

    data_stack.checkpointState(saver, &pool);
    data_stack.checkpointState(saver, &pool);
    auto get_results = saver.getTuple(hash_key);

    REQUIRE(get_results.status.ok());
    REQUIRE(get_results.reference_count == expected_ref_count);
    REQUIRE(get_results.tuple.calculateHash() == expected_hash);
}

void saveAndRecreate(Datastack data_stack) {
    TuplePool pool;
    CheckpointStorage storage(dbpath);
    auto saver = MachineStateSaver(&storage, &pool);

    auto save_results = data_stack.checkpointState(saver, &pool);
    auto results = saver.getTuple(save_results.storage_key);

    Datastack recreated_stack(results.tuple);

    REQUIRE(results.reference_count == 1);
    REQUIRE(data_stack.hash() == recreated_stack.hash());
}

void saveTwiceAndRecreate(Datastack data_stack) {
    TuplePool pool;
    CheckpointStorage storage(dbpath);
    auto saver = MachineStateSaver(&storage, &pool);

    auto save_results = data_stack.checkpointState(saver, &pool);
    auto save_results2 = data_stack.checkpointState(saver, &pool);
    auto results = saver.getTuple(save_results.storage_key);

    Datastack recreated_stack(results.tuple);

    REQUIRE(results.reference_count == 2);
    REQUIRE(data_stack.hash() == recreated_stack.hash());
}

TEST_CASE("Save datastack") {
    SECTION("save empty") {
        Datastack datastack;
        std::vector<unsigned char> hash_key_vector;
        marshal_value(Tuple().calculateHash(), hash_key_vector);

        saveDataStack(datastack, 1, hash_key_vector);
    }
    SECTION("save empty twice") {
        Datastack datastack;
        std::vector<unsigned char> hash_key_vector;
        marshal_value(Tuple().calculateHash(), hash_key_vector);

        saveDataStackTwice(datastack, 2, hash_key_vector);
    }
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

        saveDataStack(datastack, 1, hash_key_vector);
    }
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

        saveDataStackTwice(datastack, 2, hash_key_vector);
    }
}

TEST_CASE("Save and get datastack") {
    SECTION("save datastack and get tuple respresentation") {
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

        saveAndGetDataStack(datastack, hash_key_vector, 1,
                            tup_rep.calculateHash());
    }
    SECTION("save datastack twice and get tuple respresentation") {
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

        saveTwiceAndGetDataStack(datastack, hash_key_vector, 2,
                                 tup_rep.calculateHash());
    }
}

TEST_CASE("Save and recreate") {
    SECTION("Save empty") { saveAndRecreate(Datastack()); }
    SECTION("Save with values") {
        TuplePool pool;

        uint256_t num = 1;
        auto code_point = CodePoint(1, Operation(), 0);
        auto tuple = Tuple(code_point, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        saveAndRecreate(datastack);
    }
    SECTION("Save with values twice") {
        TuplePool pool;

        uint256_t num = 1;
        auto code_point = CodePoint(1, Operation(), 0);
        auto tuple = Tuple(code_point, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        saveTwiceAndRecreate(datastack);
    }
}
