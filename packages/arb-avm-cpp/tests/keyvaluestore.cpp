/*
 * Copyright 2020, Offchain Labs, Inc.
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

#include <carbstorage.h>

#include <data_storage/arbstorage.hpp>
#include <data_storage/storageresult.hpp>

#include <boost/filesystem.hpp>
#include <catch2/catch.hpp>

namespace {
std::vector<char> hash_key1_vec = {1};
rocksdb::Slice hash_key1{hash_key1_vec.data(), hash_key1_vec.size()};
std::vector<char> hash_key2_vec = {2};
rocksdb::Slice hash_key2{hash_key2_vec.data(), hash_key2_vec.size()};
std::vector<unsigned char> value1 = {'v', 'a', 'l', 'u', 'e'};
std::vector<unsigned char> value4 = {};
std::vector<unsigned char> value2 = {'v', 'a', 'l', 'u', 'e', '2'};
}  // namespace

TEST_CASE("KeyValueStore test") {
    DBDeleter deleter;
    ArbStorage storage(dbpath);
    auto store = storage.makeKeyValueStore();

    auto status = store->saveData(hash_key1, value1);
    REQUIRE(status.ok() == true);

    auto res = store->getData(hash_key1);
    REQUIRE(res.data == value1);

    auto res2 = store->getData(hash_key2);
    REQUIRE(res2.status.ok() == false);

    store->saveData(hash_key2, value4);
    res2 = store->getData(hash_key2);
    REQUIRE(res2.data == value4);
}

TEST_CASE("CArbStorage test") {
    DBDeleter deleter;
    auto store = createArbStorage(dbpath.c_str());
    auto res = getData(store, hash_key2.data(), hash_key2.size());

    REQUIRE(res.found == false);
    REQUIRE(res.slice.length == 0);
    destroyArbStorage(store);
}
