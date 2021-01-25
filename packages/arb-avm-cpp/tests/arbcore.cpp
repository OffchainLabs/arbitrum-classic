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

#include <data_storage/aggregator.hpp>
#include <data_storage/datastorage.hpp>

#include <catch2/catch.hpp>

TEST_CASE("ArbCore tests") {
    DBDeleter deleter;
    auto storage = std::make_shared<DataStorage>(dbpath);
    auto store = std::make_unique<AggregatorStore>(storage);

    /* TODO
    SECTION("requests") {
        REQUIRE(!store->getPossibleRequestInfo(10).has_value());
        store->saveRequest(10, 5);
        auto requestIndex = store->getPossibleRequestInfo(10);
        REQUIRE(requestIndex.has_value());
        REQUIRE(*requestIndex == 5);
        REQUIRE(!store->getPossibleRequestInfo(8).has_value());
    }

    SECTION("blocks") {
        CHECK_THROWS(store->latestBlock());
        std::vector<unsigned char> data{1, 2, 3, 4};
        auto tx = storage->beginTransaction();
        store->saveLog(*tx, data);
        store->saveLog(*tx, data);
        store->saveSend(*tx, data);
        store->saveSend(*tx, data);
        store->saveSend(*tx, data);
        tx = nullptr;
        std::vector<char> block_data{1, 2, 3, 4};
        store->saveBlock(50, block_data);
        {
            auto latest = store->latestBlock();
            REQUIRE(latest.first == 50);
            REQUIRE(latest.second == block_data);
        }

        tx = storage->beginTransaction();
        store->saveLog(*tx, data);
        store->saveLog(*tx, data);
        store->saveSend(*tx, data);
        tx = nullptr;
        std::vector<char> block_data2{1, 2, 3, 5};
        store->saveBlock(52, block_data2);
        {
            auto block = store->getBlock(52);
            REQUIRE(block == block_data2);
        }

        {
            // Latest is now updated
            auto latest = store->latestBlock();
            REQUIRE(latest.first == 52);
        }
    }
    */
}
