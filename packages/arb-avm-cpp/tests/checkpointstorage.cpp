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

#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/transaction.hpp>

#include <rocksdb/db.h>

#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

#include <iostream>

std::vector<unsigned char> hash_key1 = {1};
std::vector<unsigned char> hash_key2 = {2};
std::vector<unsigned char> value1 = {'v', 'a', 'l', 'u', 'e'};
std::vector<unsigned char> value2 = {'v', 'a', 'l', 'u', 'e', '2'};

auto dbPath = boost::filesystem::current_path().generic_string() + "/machineDb";

void saveVal(CheckpointStorage& storage,
             const std::vector<unsigned char>& val,
             const std::vector<unsigned char>& hash_key,
             int expected_ref_count,
             bool expected_status) {
    auto trans = storage.makeTransaction();
    auto results = trans->saveData(hash_key, val);
    auto status = trans->commit();
    auto success = results.status.ok() && status.ok();

    REQUIRE(success == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void getVal(CheckpointStorage& storage,
            const std::vector<unsigned char>& hash_key,
            int expected_ref_count,
            bool expected_status,
            const std::vector<unsigned char>& expected_val) {
    auto trans = storage.makeTransaction();
    auto results = trans->getData(hash_key);
    auto status = trans->commit();
    auto success = results.status.ok() && status.ok();

    REQUIRE(success == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(results.stored_value == expected_val);
}

void incrementRef(CheckpointStorage& storage,
                  const std::vector<unsigned char>& hash_key,
                  int expected_ref_count,
                  bool expected_status) {
    auto trans = storage.makeTransaction();
    auto results = trans->incrementReference(hash_key);
    auto status = trans->commit();
    auto success = results.status.ok() && status.ok();

    REQUIRE(success == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void deleteVal(CheckpointStorage& storage,
               const std::vector<unsigned char>& hash_key,
               int expected_ref_count,
               bool expected_status) {
    auto trans = storage.makeTransaction();
    auto results = trans->deleteData(hash_key);
    auto status = trans->commit();
    auto success = results.status.ok() && status.ok();

    REQUIRE(success == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

TEST_CASE("Checkpointstorage initialize") {
    TuplePool pool;
    CheckpointStorage storage(dbPath, test_contract_path);
    SECTION("construction fails") {
        bool fails;
        try {
            CheckpointStorage storage2(dbPath, test_contract_path);
            fails = false;
        } catch (std::exception ex) {
            fails = true;
        }
        REQUIRE(fails == true);
    }
    SECTION("get") {
        getVal(storage, hash_key1, 0, false, std::vector<unsigned char>());
    }
    SECTION("get") {
        getVal(storage, hash_key2, 0, false, std::vector<unsigned char>());
    }
    SECTION("save") { saveVal(storage, value1, hash_key1, 1, true); }
    SECTION("increment") { incrementRef(storage, hash_key1, 0, false); }
    SECTION("delete") { deleteVal(storage, hash_key1, 0, false); }
    SECTION("construction succeeds") {
        bool fails;
        try {
            storage.closeCheckpointStorage();
            CheckpointStorage storage2(dbPath, test_contract_path);
            fails = false;
        } catch (std::exception ex) {
            fails = true;
        }
        REQUIRE(fails == false);
    }

    boost::filesystem::remove_all(dbPath);
}

TEST_CASE("Save and get values") {
    SECTION("save and get") {
        TuplePool pool;
        CheckpointStorage storage(dbPath, test_contract_path);
        saveVal(storage, value1, hash_key1, 1, true);
        getVal(storage, hash_key1, 1, true, value1);
    }
    boost::filesystem::remove_all(dbPath);
    SECTION("db cleared") {
        TuplePool pool;
        CheckpointStorage storage(dbPath, test_contract_path);
        getVal(storage, hash_key1, 0, false, std::vector<unsigned char>());
    }
    boost::filesystem::remove_all(dbPath);
    SECTION("save, increment, get") {
        TuplePool pool;
        CheckpointStorage storage(dbPath, test_contract_path);
        saveVal(storage, value1, hash_key1, 1, true);
        incrementRef(storage, hash_key1, 2, true);
        getVal(storage, hash_key1, 2, true, value1);
    }
    boost::filesystem::remove_all(dbPath);
    SECTION("save, delete, get") {
        TuplePool pool;
        CheckpointStorage storage(dbPath, test_contract_path);
        saveVal(storage, value1, hash_key1, 1, true);
        saveVal(storage, value2, hash_key2, 1, true);
        getVal(storage, hash_key2, 1, true, value2);
        getVal(storage, hash_key1, 1, true, value1);
        deleteVal(storage, hash_key1, 0, true);
        getVal(storage, hash_key1, 0, false, std::vector<unsigned char>());
        getVal(storage, hash_key2, 1, true, value2);
    }
    boost::filesystem::remove_all(dbPath);
    SECTION("save, increment, delete, get") {
        TuplePool pool;
        CheckpointStorage storage(dbPath, test_contract_path);
        saveVal(storage, value1, hash_key1, 1, true);
        saveVal(storage, value2, hash_key2, 1, true);
        getVal(storage, hash_key2, 1, true, value2);
        getVal(storage, hash_key1, 1, true, value1);
        incrementRef(storage, hash_key1, 2, true);
        deleteVal(storage, hash_key1, 1, true);
        getVal(storage, hash_key1, 1, true, value1);
        getVal(storage, hash_key2, 1, true, value2);
    }
    boost::filesystem::remove_all(dbPath);
    SECTION("save, increment, delete, get") {
        TuplePool pool;
        CheckpointStorage storage(dbPath, test_contract_path);
        saveVal(storage, value1, hash_key1, 1, true);
        saveVal(storage, value2, hash_key2, 1, true);
        getVal(storage, hash_key2, 1, true, value2);
        getVal(storage, hash_key1, 1, true, value1);
        incrementRef(storage, hash_key1, 2, true);
        incrementRef(storage, hash_key1, 3, true);
        deleteVal(storage, hash_key1, 2, true);
        getVal(storage, hash_key1, 2, true, value1);
        getVal(storage, hash_key2, 1, true, value2);
    }
    boost::filesystem::remove_all(dbPath);
    SECTION("save, increment, get") {
        TuplePool pool;
        CheckpointStorage storage(dbPath, test_contract_path);
        saveVal(storage, value1, hash_key1, 1, true);
        incrementRef(storage, hash_key1, 2, true);
        deleteVal(storage, hash_key1, 1, true);
        incrementRef(storage, hash_key1, 2, true);
        incrementRef(storage, hash_key1, 3, true);
        getVal(storage, hash_key1, 3, true, value1);
    }
    boost::filesystem::remove_all(dbPath);
    SECTION("save, delete, increment, get") {
        TuplePool pool;
        CheckpointStorage storage(dbPath, test_contract_path);
        saveVal(storage, value1, hash_key1, 1, true);
        deleteVal(storage, hash_key1, 0, true);
        incrementRef(storage, hash_key1, 0, false);
        getVal(storage, hash_key1, 0, false, std::vector<unsigned char>());
    }
    boost::filesystem::remove_all(dbPath);
}
