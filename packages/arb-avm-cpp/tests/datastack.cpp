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
#include <data_storage/storageresult.hpp>
#include <data_storage/value/value.hpp>

#include <avm/machinestate/datastack.hpp>

#include <rocksdb/status.h>

#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

std::string dbpath =
    boost::filesystem::current_path().generic_string() + "/machineDb";

void initializeDatastack(const Transaction& transaction,
                         uint256_t tuple_hash,
                         uint256_t expected_hash,
                         uint64_t expected_size) {
    ValueCache value_cache{};
    auto results = ::getValue(transaction, tuple_hash, value_cache);
    REQUIRE(results.status.ok());
    REQUIRE(nonstd::holds_alternative<Tuple>(results.data));

    Datastack data_stack(nonstd::get<Tuple>(results.data));

    REQUIRE(data_stack.hash() == expected_hash);
    REQUIRE(data_stack.stacksize() == expected_size);
}

void saveDataStack(const Datastack& data_stack) {
    ArbStorage storage(dbpath);
    std::vector<CodePoint> code;
    auto transaction = storage.makeTransaction();

    auto tuple_ret = data_stack.getTupleRepresentation();
    auto results = saveValue(*transaction, tuple_ret);

    REQUIRE(transaction->commit().ok());
    REQUIRE(results.status.ok());
    REQUIRE(results.reference_count == 1);
}

void saveDataStackTwice(const Datastack& data_stack) {
    ArbStorage storage(dbpath);
    std::vector<CodePoint> code;
    auto transaction = storage.makeTransaction();

    auto tuple_ret = data_stack.getTupleRepresentation();
    auto results = saveValue(*transaction, tuple_ret);
    auto results2 = saveValue(*transaction, tuple_ret);

    REQUIRE(transaction->commit().ok());
    REQUIRE(results2.status.ok());
    REQUIRE(results2.reference_count == 2);
}

void saveAndGetDataStack(Transaction& transaction,
                         const Datastack& data_stack,
                         uint256_t expected_hash) {
    auto tuple_ret = data_stack.getTupleRepresentation();
    auto results = saveValue(transaction, tuple_ret);
    REQUIRE(results.status.ok());
    transaction.commit();

    ValueCache value_cache{};
    auto get_results = getValue(transaction, expected_hash, value_cache);

    REQUIRE(nonstd::holds_alternative<Tuple>(get_results.data));
    REQUIRE(get_results.status.ok());
    REQUIRE(get_results.reference_count == 1);
    REQUIRE(hash_value(get_results.data) == expected_hash);
}

void saveTwiceAndGetDataStack(Transaction& transaction,
                              const Datastack& data_stack,
                              uint256_t expected_hash) {
    auto tuple_ret = data_stack.getTupleRepresentation();
    auto results = saveValue(transaction, tuple_ret);
    auto results2 = saveValue(transaction, tuple_ret);
    REQUIRE(results.status.ok());
    REQUIRE(results2.status.ok());
    transaction.commit();

    ValueCache value_cache{};
    auto get_results = getValue(transaction, expected_hash, value_cache);

    REQUIRE(nonstd::holds_alternative<Tuple>(get_results.data));
    REQUIRE(get_results.status.ok());
    REQUIRE(get_results.reference_count == 2);
    REQUIRE(hash_value(get_results.data) == expected_hash);
}

TEST_CASE("Initialize datastack") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    auto transaction = storage.makeTransaction();
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
        auto tuple = Tuple(code_point_stub);
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
        auto tuple = Tuple(num);
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
        auto tuple = Tuple(intVal);
        datastack.push(num);
        datastack.push(tuple);
        Tuple tup0;
        auto tup1 = Tuple(tuple, tup0);
        auto tup_rep = Tuple(num, tup1);
        saveDataStack(datastack);
    }
    SECTION("save with values, twice") {
        uint256_t num = 1;
        uint256_t intVal = 5435;
        auto tuple = Tuple(intVal);
        datastack.push(num);
        datastack.push(tuple);
        Tuple tup0;
        auto tup1 = Tuple(tuple, tup0);
        auto tup_rep = Tuple(num, tup1);
        saveDataStackTwice(datastack);
    }
}

TEST_CASE("Save and get datastack") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    Datastack datastack;

    SECTION("save datastack and get") {
        uint256_t intVal = 5435;
        auto transaction = storage.makeTransaction();
        uint256_t num = 1;
        auto tuple = Tuple(intVal);
        datastack.push(num);
        datastack.push(tuple);
        Tuple tup0;
        auto tup1 = Tuple(tuple, tup0);
        auto tup_rep = Tuple(num, tup1);
        saveAndGetDataStack(*transaction, datastack, hash(tup_rep));
    }
    SECTION("save datastack twice and get") {
        uint256_t intVal = 5435;
        auto transaction = storage.makeTransaction();
        uint256_t num = 1;
        auto tuple = Tuple(intVal);
        datastack.push(num);
        datastack.push(tuple);
        Tuple tup0;
        auto tup1 = Tuple(tuple, tup0);
        auto tup_rep = Tuple(num, tup1);
        saveTwiceAndGetDataStack(*transaction, datastack, hash(tup_rep));
    }
}

TEST_CASE("Initial VM Values") {
    SECTION("parse invalid path") {
        CHECK_THROWS(loadExecutable("nonexistent/path"));
    }
    boost::filesystem::remove_all(dbpath);
}
