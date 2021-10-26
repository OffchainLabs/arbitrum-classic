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

#include "helper.hpp"

#include <data_storage/arbstorage.hpp>
#include <data_storage/readwritetransaction.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machinestate/datastack.hpp>

#include <rocksdb/status.h>

#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

std::string dbpath =
    boost::filesystem::current_path().generic_string() + "/machineDb";

void checkGetTupleResult(const DbResult<value>& res,
                         uint256_t expected_count,
                         uint256_t expected_hash) {
    REQUIRE(std::holds_alternative<CountedData<value>>(res));
    REQUIRE(
        std::holds_alternative<Tuple>(std::get<CountedData<value>>(res).data));
    REQUIRE(std::get<CountedData<value>>(res).reference_count ==
            expected_count);
    REQUIRE(hash_value(std::get<CountedData<value>>(res).data) ==
            expected_hash);
}

void initializeDatastack(const ReadTransaction& transaction,
                         uint256_t tuple_hash,
                         uint256_t expected_hash,
                         uint64_t expected_size) {
    ValueCache value_cache{1, 0};
    auto results = ::getValue(transaction, tuple_hash, value_cache, false);
    REQUIRE(std::holds_alternative<CountedData<value>>(results));
    REQUIRE(std::holds_alternative<Tuple>(
        std::get<CountedData<value>>(results).data));

    Datastack data_stack(
        std::get<Tuple>(std::get<CountedData<value>>(results).data));

    REQUIRE(data_stack.hash() == expected_hash);
    REQUIRE(data_stack.stacksize() == expected_size);
}

void saveDataStack(const Datastack& data_stack) {
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    std::vector<CodePoint> code;
    auto transaction = storage.makeReadWriteTransaction();

    auto tuple_ret = data_stack.getTupleRepresentation();
    auto results = saveValue(*transaction, tuple_ret);

    REQUIRE(transaction->commit().ok());
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
}

void saveDataStackTwice(const Datastack& data_stack) {
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    std::vector<CodePoint> code;
    auto transaction = storage.makeReadWriteTransaction();

    auto tuple_ret = data_stack.getTupleRepresentation();
    auto results = saveValue(*transaction, tuple_ret);
    auto results2 = saveValue(*transaction, tuple_ret);

    REQUIRE(transaction->commit().ok());
    REQUIRE(results2.status.ok());
    REQUIRE(results2.reference_count == 2);
}

void saveAndGetDataStack(ReadWriteTransaction& transaction,
                         const Datastack& data_stack,
                         uint256_t expected_hash) {
    auto tuple_ret = data_stack.getTupleRepresentation();
    auto results = saveValue(transaction, tuple_ret);
    REQUIRE(results.status.ok());
    transaction.commit();

    ValueCache value_cache{1, 0};
    auto get_results = getValue(transaction, expected_hash, value_cache, false);
    checkGetTupleResult(get_results, 1, expected_hash);
}

void saveTwiceAndGetDataStack(ReadWriteTransaction& transaction,
                              const Datastack& data_stack,
                              uint256_t expected_hash) {
    auto tuple_ret = data_stack.getTupleRepresentation();
    auto results = saveValue(transaction, tuple_ret);
    auto results2 = saveValue(transaction, tuple_ret);
    REQUIRE(results.status.ok());
    REQUIRE(results2.status.ok());
    transaction.commit();

    ValueCache value_cache{1, 0};
    auto get_results = getValue(transaction, expected_hash, value_cache, false);
    checkGetTupleResult(get_results, 2, expected_hash);
}

TEST_CASE("Initialize datastack") {
    DBDeleter deleter;
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    auto transaction = storage.makeReadWriteTransaction();
    Datastack data_stack;

    SECTION("default") {
        auto tuple_ret = data_stack.getTupleRepresentation();
        auto results = saveValue(*transaction, tuple_ret);
        transaction->commit();
        initializeDatastack(*transaction, hash(tuple_ret), data_stack.hash(),
                            0);
    }
    SECTION("push empty tuple") {
        Tuple tuple;
        data_stack.push(tuple);
        auto tuple_ret = data_stack.getTupleRepresentation();
        auto results = saveValue(*transaction, tuple_ret);
        transaction->commit();
        initializeDatastack(*transaction, hash(tuple_ret), data_stack.hash(),
                            1);
    }
    SECTION("push num, tuple") {
        CodePointStub code_point_stub{{0, 0}, 3452345};
        uint256_t num = 1;
        auto tuple = Tuple::createTuple(code_point_stub);
        data_stack.push(num);
        data_stack.push(tuple);
        auto tuple_ret = data_stack.getTupleRepresentation();
        auto results = saveValue(*transaction, tuple_ret);
        transaction->commit();
        initializeDatastack(*transaction, hash(tuple_ret), data_stack.hash(),
                            2);
    }
    SECTION("push codepoint, tuple") {
        CodePointStub code_point_stub{{0, 0}, 3452345};
        uint256_t num = 1;
        auto tuple = Tuple::createTuple(num);
        data_stack.push(code_point_stub);
        data_stack.push(tuple);
        auto tuple_ret = data_stack.getTupleRepresentation();
        auto results = saveValue(*transaction, tuple_ret);
        transaction->commit();
        initializeDatastack(*transaction, hash(tuple_ret), data_stack.hash(),
                            2);
    }
}

TEST_CASE("Save datastack") {
    DBDeleter deleter;
    Datastack datastack;

    SECTION("save empty") { saveDataStack(datastack); }
    SECTION("save empty twice") { saveDataStackTwice(datastack); }
    SECTION("save with values") {
        uint256_t num = 1;
        uint256_t intVal = 5435;
        auto tuple = Tuple::createTuple(intVal);
        datastack.push(num);
        datastack.push(tuple);
        Tuple tup0;
        auto tup1 = Tuple(num, tup0);
        auto tup_rep = Tuple(tuple, tup1);
        saveDataStack(datastack);
    }
    SECTION("save with values, twice") {
        uint256_t num = 1;
        uint256_t intVal = 5435;
        auto tuple = Tuple::createTuple(intVal);
        datastack.push(num);
        datastack.push(tuple);
        Tuple tup0;
        auto tup1 = Tuple(num, tup0);
        auto tup_rep = Tuple(tuple, tup1);
        saveDataStackTwice(datastack);
    }
}

TEST_CASE("Save and get datastack") {
    DBDeleter deleter;
    ArbCoreConfig coreConfig{};
    ArbStorage storage(dbpath, coreConfig);
    Datastack datastack;

    SECTION("save datastack and get") {
        uint256_t intVal = 5435;
        auto transaction = storage.makeReadWriteTransaction();
        uint256_t num = 1;
        auto tuple = Tuple::createTuple(intVal);
        datastack.push(num);
        datastack.push(tuple);
        Tuple tup0;
        auto tup1 = Tuple(num, tup0);
        auto tup_rep = Tuple(tuple, tup1);
        saveAndGetDataStack(*transaction, datastack, hash(tup_rep));
    }
    SECTION("save datastack twice and get") {
        uint256_t intVal = 5435;
        auto transaction = storage.makeReadWriteTransaction();
        uint256_t num = 1;
        auto tuple = Tuple::createTuple(intVal);
        datastack.push(num);
        datastack.push(tuple);
        Tuple tup0;
        auto tup1 = Tuple(num, tup0);
        auto tup_rep = Tuple(tuple, tup1);
        saveTwiceAndGetDataStack(*transaction, datastack, hash(tup_rep));
    }
}

TEST_CASE("Initial VM Values") {
    SECTION("parse invalid path") {
        CHECK_THROWS(loadExecutable("nonexistent/path"));
    }
    boost::filesystem::remove_all(dbpath);
}
