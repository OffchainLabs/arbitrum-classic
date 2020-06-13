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

void initializeDatastack(const Transaction& transaction,
                         uint256_t tuple_hash,
                         uint256_t expected_hash,
                         int expected_size) {
    TuplePool pool;

    auto results = ::getValue(transaction, tuple_hash, &pool);
    REQUIRE(results.status.ok());
    REQUIRE(nonstd::holds_alternative<Tuple>(results.data));

    Datastack data_stack(nonstd::get<Tuple>(results.data));

    REQUIRE(data_stack.hash() == expected_hash);
    REQUIRE(data_stack.stacksize() == expected_size);
}

void saveDataStack(Datastack data_stack) {
    TuplePool pool;
    CheckpointStorage storage(dbpath, test_contract_path);
    std::vector<CodePoint> code;
    auto transaction = storage.makeTransaction();

    auto tuple_ret = data_stack.getTupleRepresentation(&pool);
    auto results = saveValue(*transaction, tuple_ret);

    REQUIRE(transaction->commit().ok());
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
}

void saveDataStackTwice(Datastack data_stack) {
    TuplePool pool;
    CheckpointStorage storage(dbpath, test_contract_path);
    std::vector<CodePoint> code;
    auto transaction = storage.makeTransaction();

    auto tuple_ret = data_stack.getTupleRepresentation(&pool);
    auto results = saveValue(*transaction, tuple_ret);
    auto results2 = saveValue(*transaction, tuple_ret);

    REQUIRE(transaction->commit().ok());
    REQUIRE(results2.status.ok());
    REQUIRE(results2.reference_count == 2);
}

void saveAndGetDataStack(Transaction& transaction,
                         Datastack data_stack,
                         uint256_t expected_hash) {
    TuplePool pool;
    auto tuple_ret = data_stack.getTupleRepresentation(&pool);
    auto results = saveValue(transaction, tuple_ret);
    REQUIRE(results.status.ok());
    transaction.commit();

    auto get_results = getValue(transaction, expected_hash, &pool);

    REQUIRE(nonstd::holds_alternative<Tuple>(get_results.data));
    REQUIRE(get_results.status.ok());
    REQUIRE(get_results.reference_count == 1);
    REQUIRE(nonstd::get<Tuple>(get_results.data).calculateHash() ==
            expected_hash);
}

void saveTwiceAndGetDataStack(Transaction& transaction,
                              Datastack data_stack,
                              uint256_t expected_hash) {
    TuplePool pool;

    auto tuple_ret = data_stack.getTupleRepresentation(&pool);
    auto results = saveValue(transaction, tuple_ret);
    auto results2 = saveValue(transaction, tuple_ret);
    REQUIRE(results.status.ok());
    REQUIRE(results2.status.ok());
    transaction.commit();

    auto get_results = getValue(transaction, expected_hash, &pool);

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

        Datastack data_stack;

        auto tuple_ret = data_stack.getTupleRepresentation(&pool);
        auto results = saveValue(*transaction, tuple_ret);
        transaction->commit();

        initializeDatastack(*transaction, hash(tuple_ret), data_stack.hash(),
                            0);
    }
    boost::filesystem::remove_all(dbpath);

    SECTION("push empty tuple") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        auto transaction = storage.makeTransaction();

        Datastack data_stack;
        Tuple tuple;
        data_stack.push(tuple);

        auto tuple_ret = data_stack.getTupleRepresentation(&pool);
        auto results = saveValue(*transaction, tuple_ret);
        transaction->commit();

        initializeDatastack(*transaction, hash(tuple_ret), data_stack.hash(),
                            1);
    }
    boost::filesystem::remove_all(dbpath);

    SECTION("push num, tuple") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePointStub code_point_stub{0, 3452345};

        auto transaction = storage.makeTransaction();

        Datastack data_stack;
        uint256_t num = 1;
        auto tuple = Tuple(code_point_stub, &pool);

        data_stack.push(num);
        data_stack.push(tuple);

        auto tuple_ret = data_stack.getTupleRepresentation(&pool);
        auto results = saveValue(*transaction, tuple_ret);
        transaction->commit();

        initializeDatastack(*transaction, hash(tuple_ret), data_stack.hash(),
                            2);
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("push codepoint, tuple") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        CodePointStub code_point_stub{0, 3452345};

        auto transaction = storage.makeTransaction();

        Datastack data_stack;

        uint256_t num = 1;
        auto tuple = Tuple(num, &pool);

        data_stack.push(code_point_stub);
        data_stack.push(tuple);

        auto tuple_ret = data_stack.getTupleRepresentation(&pool);
        auto results = saveValue(*transaction, tuple_ret);
        transaction->commit();

        initializeDatastack(*transaction, hash(tuple_ret), data_stack.hash(),
                            2);
    }
    boost::filesystem::remove_all(dbpath);
}

TEST_CASE("Save datastack") {
    SECTION("save empty") {
        Datastack datastack;
        saveDataStack(datastack);
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("save empty twice") {
        Datastack datastack;
        saveDataStackTwice(datastack);
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

        Tuple tup0;
        auto tup1 = Tuple(tuple, tup0, &pool);
        auto tup_rep = Tuple(num, tup1, &pool);

        saveDataStack(datastack);
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

        Tuple tup0;
        auto tup1 = Tuple(tuple, tup0, &pool);
        auto tup_rep = Tuple(num, tup1, &pool);

        saveDataStackTwice(datastack);
    }
    boost::filesystem::remove_all(dbpath);
}

TEST_CASE("Save and get datastack") {
    SECTION("save datastack and get") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        uint256_t intVal = 5435;

        auto transaction = storage.makeTransaction();

        uint256_t num = 1;
        auto tuple = Tuple(intVal, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        Tuple tup0;
        auto tup1 = Tuple(tuple, tup0, &pool);
        auto tup_rep = Tuple(num, tup1, &pool);

        saveAndGetDataStack(*transaction, datastack, hash(tup_rep));
    }
    boost::filesystem::remove_all(dbpath);
    SECTION("save datastack twice and get") {
        TuplePool pool;
        CheckpointStorage storage(dbpath, test_contract_path);

        uint256_t intVal = 5435;

        auto transaction = storage.makeTransaction();

        uint256_t num = 1;
        auto tuple = Tuple(intVal, &pool);

        Datastack datastack;
        datastack.push(num);
        datastack.push(tuple);

        Tuple tup0;
        auto tup1 = Tuple(tuple, tup0, &pool);
        auto tup_rep = Tuple(num, tup1, &pool);

        saveTwiceAndGetDataStack(*transaction, datastack, hash(tup_rep));
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
