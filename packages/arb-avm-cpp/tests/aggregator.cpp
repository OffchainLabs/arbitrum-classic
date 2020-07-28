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

TEST_CASE("Aggregator tests") {
    DBDeleter deleter;
    auto storage = std::make_shared<DataStorage>(dbpath);
    auto store = std::make_unique<AggregatorStore>(storage);

    SECTION("logs") {
        REQUIRE(store->logCount() == 0);
        CHECK_THROWS(store->getLog(0));
        std::vector<char> sample_log{1, 2, 3, 4};
        store->saveLog(sample_log);
        REQUIRE(store->logCount() == 1);
        REQUIRE(store->getLog(0) == sample_log);
        CHECK_THROWS(store->getLog(1));
    }

    SECTION("messages") {
        REQUIRE(store->messageCount() == 0);
        CHECK_THROWS(store->getMessage(0));
        std::vector<char> sample_message{1, 2, 3, 4};
        store->saveMessage(sample_message);
        REQUIRE(store->messageCount() == 1);
        REQUIRE(store->getMessage(0) == sample_message);
        CHECK_THROWS(store->getMessage(1));
    }

    SECTION("requests") {
        CHECK_THROWS(store->getPossibleRequestInfo(10));
        store->saveRequest(10, 5);
        REQUIRE(store->getPossibleRequestInfo(10) == 5);
        CHECK_THROWS(store->getPossibleRequestInfo(8));
    }

    SECTION("blocks") {
        CHECK_THROWS(store->latestBlock());
        std::vector<char> data{1, 2, 3, 4};
        store->saveLog(data);
        store->saveLog(data);
        store->saveMessage(data);
        store->saveMessage(data);
        store->saveMessage(data);
        std::vector<char> block_data{1, 2, 3, 4};
        store->saveBlock(50, block_data);
        {
            auto latest = store->latestBlock();
            REQUIRE(latest.first == 50);
            REQUIRE(latest.second == block_data);
        }

        CHECK_THROWS(store->saveBlock(52, block_data));
        {
            // Latest is unmodified
            auto latest = store->latestBlock();
            REQUIRE(latest.first == 50);
            REQUIRE(latest.second == block_data);
        }

        store->saveLog(data);
        store->saveLog(data);
        store->saveMessage(data);
        std::vector<char> block_data2{1, 2, 3, 5};
        store->saveBlock(51, block_data2);
        {
            auto block = store->getBlock(51);
            REQUIRE(block == block_data2);
        }

        {
            // Latest is now updated
            auto latest = store->latestBlock();
            REQUIRE(latest.first == 51);
        }
    }
}
