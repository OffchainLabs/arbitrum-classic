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

#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/storageresult.hpp>

#include <boost/filesystem.hpp>
#include <catch2/catch.hpp>

std::string pathc =
    boost::filesystem::current_path().generic_string() + "/keyvaluestoreDb";

namespace {
std::vector<unsigned char> hash_key1 = {1};
std::vector<unsigned char> hash_key2 = {2};
std::vector<unsigned char> value1 = {'v', 'a', 'l', 'u', 'e'};
std::vector<unsigned char> value4 = {};
std::vector<unsigned char> value2 = {'v', 'a', 'l', 'u', 'e', '2'};
}  // namespace

TEST_CASE("Save") {
    CheckpointStorage storage(pathc, test_contract_path);
    auto store = storage.makeKeyValueStore();

    auto status = store->saveData(hash_key1, value1);
    REQUIRE(status.ok() == true);

    auto res = store->getData(hash_key1);
    REQUIRE(res.data == value1);

    auto res2 = store->getData(hash_key2);

    store->saveData(hash_key2, value4);

    auto x = store->getData(hash_key2);
}
