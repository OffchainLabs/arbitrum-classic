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
#include <data_storage/value/transaction.hpp>

#include <rocksdb/db.h>

#include <catch2/catch.hpp>

#include <boost/filesystem.hpp>

#include <iostream>

std::vector<unsigned char> hash_key1 = {1};
std::vector<unsigned char> hash_key2 = {2};
std::vector<unsigned char> value1 = {'v', 'a', 'l', 'u', 'e'};
std::vector<unsigned char> value2 = {'v', 'a', 'l', 'u', 'e', '2'};

void saveVal(CheckpointStorage& storage,
             std::vector<unsigned char> val,
             std::vector<unsigned char> hash_key,
             int expected_ref_count,
             bool expected_status) {
    auto trans = storage.makeTransaction();
    rocksdb::Slice key{reinterpret_cast<const char*>(hash_key.data()),
                       hash_key.size()};
    auto results = trans->saveData(key, val);
    auto status = trans->commit();
    auto success = results.status.ok() && status.ok();

    REQUIRE(success == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void getVal(CheckpointStorage& storage,
            std::vector<unsigned char> hash_key,
            int expected_ref_count,
            bool expected_status,
            std::vector<unsigned char> expected_val) {
    auto trans = storage.makeTransaction();
    rocksdb::Slice key{reinterpret_cast<const char*>(hash_key.data()),
                       hash_key.size()};
    auto results = trans->getData(key);
    auto status = trans->commit();
    auto success = results.status.ok() && status.ok();

    REQUIRE(success == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
    REQUIRE(results.stored_value == expected_val);
}

void incrementRef(CheckpointStorage& storage,
                  std::vector<unsigned char> hash_key,
                  int expected_ref_count,
                  bool expected_status) {
    auto trans = storage.makeTransaction();
    rocksdb::Slice key{reinterpret_cast<const char*>(hash_key.data()),
                       hash_key.size()};
    auto results = trans->incrementReference(key);
    auto status = trans->commit();
    auto success = results.status.ok() && status.ok();

    REQUIRE(success == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

void deleteVal(CheckpointStorage& storage,
               std::vector<unsigned char> hash_key,
               int expected_ref_count,
               bool expected_status) {
    auto trans = storage.makeTransaction();
    rocksdb::Slice key{reinterpret_cast<const char*>(hash_key.data()),
                       hash_key.size()};
    auto results = trans->deleteData(key);
    auto status = trans->commit();
    auto success = results.status.ok() && status.ok();

    REQUIRE(success == expected_status);
    REQUIRE(results.reference_count == expected_ref_count);
}

TEST_CASE("Checkpointstorage initialize") {
    DBDeleter deleter;
    TuplePool pool;
    CheckpointStorage storage(dbpath, test_contract_path);
    SECTION("construction fails") {
        bool fails;
        try {
            CheckpointStorage storage2(dbpath, test_contract_path);
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
            CheckpointStorage storage2(dbpath, test_contract_path);
            fails = false;
        } catch (std::exception ex) {
            fails = true;
        }
        REQUIRE(fails == false);
    }
}

TEST_CASE("Save and get values") {
    DBDeleter deleter;
    TuplePool pool;
    CheckpointStorage storage(dbpath, test_contract_path);

    SECTION("save and get") {
        saveVal(storage, value1, hash_key1, 1, true);
        getVal(storage, hash_key1, 1, true, value1);
    }
    SECTION("db cleared") {
        getVal(storage, hash_key1, 0, false, std::vector<unsigned char>());
    }
    SECTION("save, increment, get") {
        saveVal(storage, value1, hash_key1, 1, true);
        incrementRef(storage, hash_key1, 2, true);
        getVal(storage, hash_key1, 2, true, value1);
    }
    SECTION("save, delete, get") {
        saveVal(storage, value1, hash_key1, 1, true);
        saveVal(storage, value2, hash_key2, 1, true);
        getVal(storage, hash_key2, 1, true, value2);
        getVal(storage, hash_key1, 1, true, value1);
        deleteVal(storage, hash_key1, 0, true);
        getVal(storage, hash_key1, 0, false, std::vector<unsigned char>());
        getVal(storage, hash_key2, 1, true, value2);
    }
    SECTION("save, increment, delete, get") {
        saveVal(storage, value1, hash_key1, 1, true);
        saveVal(storage, value2, hash_key2, 1, true);
        getVal(storage, hash_key2, 1, true, value2);
        getVal(storage, hash_key1, 1, true, value1);
        incrementRef(storage, hash_key1, 2, true);
        deleteVal(storage, hash_key1, 1, true);
        getVal(storage, hash_key1, 1, true, value1);
        getVal(storage, hash_key2, 1, true, value2);
    }
    SECTION("save, increment, delete, get") {
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
    SECTION("save, increment, get") {
        saveVal(storage, value1, hash_key1, 1, true);
        incrementRef(storage, hash_key1, 2, true);
        deleteVal(storage, hash_key1, 1, true);
        incrementRef(storage, hash_key1, 2, true);
        incrementRef(storage, hash_key1, 3, true);
        getVal(storage, hash_key1, 3, true, value1);
    }
    SECTION("save, delete, increment, get") {
        saveVal(storage, value1, hash_key1, 1, true);
        deleteVal(storage, hash_key1, 0, true);
        incrementRef(storage, hash_key1, 0, false);
        getVal(storage, hash_key1, 0, false, std::vector<unsigned char>());
    }
}
