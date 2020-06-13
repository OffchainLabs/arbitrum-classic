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

void initializeDatastack(MachineStateFetcher& fetcher,
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
    CheckpointStorage storage(dbpath, test_contract_path);
    std::vector<CodePoint> code;
    auto transaction = storage.makeTransaction();

    auto results = data_stack.checkpointState(*transaction, &pool);
    transaction->commit();

    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
    REQUIRE(results.storage_key == expected_hash_key);
}

void saveDataStackTwice(Datastack data_stack,
                        std::vector<unsigned char> expected_hash_key) {
    TuplePool pool;
    CheckpointStorage storage(dbpath, test_contract_path);
    std::vector<CodePoint> code;
    auto transaction = storage.makeTransaction();

    auto results = data_stack.checkpointState(*transaction, &pool);
    auto results2 = data_stack.checkpointState(*transaction, &pool);

    transaction->commit();

    REQUIRE(results2.status.ok());
    REQUIRE(results2.reference_count == 2);
    REQUIRE(results2.storage_key == expected_hash_key);
}

void saveAndGetDataStack(Transaction& transaction,
                         MachineStateFetcher& fetcher,
                         Datastack data_stack,
                         std::vector<unsigned char> hash_key,
                         uint256_t expected_hash) {
    TuplePool pool;
    data_stack.checkpointState(transaction, &pool);
    transaction.commit();

    auto get_results = fetcher.getValue(hash_key);

    REQUIRE(nonstd::holds_alternative<Tuple>(get_results.data));
    REQUIRE(get_results.status.ok());
    REQUIRE(get_results.reference_count == 1);
    REQUIRE(nonstd::get<Tuple>(get_results.data).calculateHash() ==
            expected_hash);
}

void saveTwiceAndGetDataStack(Transaction& transaction,
                              MachineStateFetcher& fetcher,
                              Datastack data_stack,
                              std::vector<unsigned char> hash_key,
                              uint256_t expected_hash) {
    TuplePool pool;

    data_stack.checkpointState(transaction, &pool);
    data_stack.checkpointState(transaction, &pool);
    transaction.commit();

    auto get_results = fetcher.getValue(hash_key);

    REQUIRE(nonstd::holds_alternative<Tuple>(get_results.data));
    REQUIRE(get_results.status.ok());
    REQUIRE(get_results.reference_count == 2);
    REQUIRE(nonstd::get<Tuple>(get_results.data).calculateHash() ==
            expected_hash);
}

TEST_CASE("Initialize datastack") {
    SECTION("default") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        Datastack data_stack;

        auto results = data_stack.checkpointState(*transaction, &pool);
        transaction->commit();
        auto stack_hash = data_stack.hash();

        initializeDatastack(fetcher, results.storage_key, stack_hash, 0);
    }
    boost::filesystem::remove_all(dbpath);

    SECTION("push empty tuple") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        Datastack data_stack;
        Tuple tuple;
        data_stack.push(tuple);

        auto results = data_stack.checkpointState(*transaction, &pool);
        transaction->commit();
        auto stack_hash = data_stack.hash();

        initializeDatastack(fetcher, results.storage_key, stack_hash, 1);
    }
    boost::filesystem::remove_all(dbpath);

    SECTION("push num, tuple") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePointStub code_point_stub{0, 3452345};

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        Datastack data_stack;
        uint256_t num = 1;
        auto tuple = Tuple(code_point_stub, &pool);

        data_stack.push(num);
        data_stack.push(tuple);

        auto results = data_stack.checkpointState(*transaction, &pool);
        transaction->commit();
        auto stack_hash = data_stack.hash();

        initializeDatastack(fetcher, results.storage_key, stack_hash, 2);
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("push codepoint, tuple") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePointStub code_point_stub{0, 3452345};

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        Datastack data_stack;

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);

        data_stack.push(code_point_stub);
        data_stack.push(tuple);

        auto results = data_stack.checkpointState(*transaction, &pool);
        transaction->commit();
        auto stack_hash = data_stack.hash();

        initializeDatastack(fetcher, results.storage_key, stack_hash, 2);
    }
    boost::filesystem::remove_all(dbpath);
}

TEST_CASE("Save datastack") {
    SECTION("save empty") {
        Datastack datastack;
        std::vector<unsigned char> hash_key_vector;
        marshal_uint256_t(Tuple().calculateHash(), hash_key_vector);

        saveDataStack(datastack, hash_key_vector);
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("save empty twice") {
        Datastack datastack;
        std::vector<unsigned char> hash_key_vector;
        marshal_uint256_t(Tuple().calculateHash(), hash_key_vector);

        saveDataStackTwice(datastack, hash_key_vector);
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("save with values") {
        TuplePool pool;

        uint256_t num = 1;
        uint256_t intVal = 5435;
        auto tuple = Tuple(intVal, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        auto tup1 = Tuple(num, &pool);
        auto tup_rep = Tuple(tuple, tup1, &pool);

        std::vector<unsigned char> hash_key_vector;
        marshal_uint256_t(tup_rep.calculateHash(), hash_key_vector);

        saveDataStack(datastack, hash_key_vector);
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("save with values, twice") {
        TuplePool pool;

        uint256_t num = 1;
        uint256_t intVal = 5435;
        auto tuple = Tuple(intVal, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        auto tup1 = Tuple(num, &pool);
        auto tup_rep = Tuple(tuple, tup1, &pool);

        std::vector<unsigned char> hash_key_vector;
        marshal_uint256_t(tup_rep.calculateHash(), hash_key_vector);

        saveDataStackTwice(datastack, hash_key_vector);
    }
    boost::filesystem::remove_all(dbpath);
}

TEST_CASE("Save and get datastack") {
    SECTION("save datastack and get") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        uint256_t intVal = 5435;

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto tuple = Tuple(intVal, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        auto tup1 = Tuple(num, &pool);
        auto tup_rep = Tuple(tuple, tup1, &pool);

        std::vector<unsigned char> hash_key_vector;
        marshal_uint256_t(tup_rep.calculateHash(), hash_key_vector);

        saveAndGetDataStack(*transaction, fetcher, datastack, hash_key_vector,
                            tup_rep.calculateHash());
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("save datastack twice and get") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        uint256_t intVal = 5435;

        auto transaction = storage.makeTransaction();
        auto fetcher = MachineStateFetcher(storage);

        uint256_t num = 1;
        auto tuple = Tuple(intVal, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        auto tup1 = Tuple(num, &pool);
        auto tup_rep = Tuple(tuple, tup1, &pool);

        std::vector<unsigned char> hash_key_vector;
        marshal_uint256_t(tup_rep.calculateHash(), hash_key_vector);

        saveTwiceAndGetDataStack(*transaction, fetcher, datastack,
                                 hash_key_vector, tup_rep.calculateHash());
    }
    boost::filesystem::remove_all(dbpath);
}

TEST_CASE("Initial VM Values") {
    SECTION("parse invalid path") {
        TuplePool pool = TuplePool();
        TuplePool& pool_ref = pool;
        auto values = parseInitialVmValues("nonexistent/path", pool_ref);
    }
    boost::filesystem::remove_all(dbpath);
}
