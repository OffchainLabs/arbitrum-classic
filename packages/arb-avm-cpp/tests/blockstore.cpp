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

#include <data_storage/blockstore.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>

#include <catch2/catch.hpp>

TEST_CASE("BlockStore tests") {
    DBDeleter deleter;
    auto storage = std::make_shared<DataStorage>(dbpath);
    auto store = std::make_unique<BlockStore>(storage);

    SECTION("BlockStore min, max, and empty") {
        REQUIRE(store->isEmpty());
        REQUIRE(store->maxHeight() == 0);
        REQUIRE(store->minHeight() == 0);

        REQUIRE(store->putBlock(10, 30, {1, 2, 3}).ok());
        REQUIRE(!store->isEmpty());
        REQUIRE(store->maxHeight() == 10);
        REQUIRE(store->minHeight() == 10);

        REQUIRE(store->putBlock(20, 30, {1, 2, 3}).ok());
        REQUIRE(!store->isEmpty());
        REQUIRE(store->maxHeight() == 20);
        REQUIRE(store->minHeight() == 10);

        REQUIRE(store->putBlock(5, 30, {1, 2, 3}).ok());
        REQUIRE(!store->isEmpty());
        REQUIRE(store->maxHeight() == 20);
        REQUIRE(store->minHeight() == 5);

        REQUIRE(store->putBlock(15, 30, {1, 2, 3}).ok());
        REQUIRE(!store->isEmpty());
        REQUIRE(store->maxHeight() == 20);
        REQUIRE(store->minHeight() == 5);

        REQUIRE(store->deleteBlock(20, 30).ok());
        REQUIRE(store->maxHeight() == 15);
        REQUIRE(store->minHeight() == 5);

        REQUIRE(store->deleteBlock(5, 30).ok());
        REQUIRE(store->maxHeight() == 15);
        REQUIRE(store->minHeight() == 10);
    }

    SECTION("BlockStore block hashes at height") {
        REQUIRE(store->blockHashesAtHeight(10).empty());

        REQUIRE(store->putBlock(10, 30, {1, 2, 3}).ok());
        REQUIRE(store->blockHashesAtHeight(10) == std::vector<uint256_t>{30});
        REQUIRE(store->blockHashesAtHeight(9).empty());
        REQUIRE(store->blockHashesAtHeight(11).empty());

        REQUIRE(store->putBlock(10, 40, {1, 2, 3}).ok());
        REQUIRE(store->blockHashesAtHeight(10) ==
                std::vector<uint256_t>{30, 40});

        REQUIRE(store->deleteBlock(10, 30).ok());
        REQUIRE(store->blockHashesAtHeight(10) == std::vector<uint256_t>{40});
    }

    SECTION("BlockStore put and get") {
        auto result = store->getBlock(10, 30);
        REQUIRE(!result.status.ok());
        REQUIRE(result.data.empty());

        REQUIRE(store->putBlock(10, 30, {1, 2, 3}).ok());
        result = store->getBlock(10, 30);
        REQUIRE(result.status.ok());
        REQUIRE(result.data == std::vector<unsigned char>{1, 2, 3});

        REQUIRE(store->putBlock(10, 30, {2, 2, 3, 4}).ok());
        result = store->getBlock(10, 30);
        REQUIRE(result.status.ok());
        REQUIRE(result.data == std::vector<unsigned char>{2, 2, 3, 4});

        REQUIRE(store->deleteBlock(10, 30).ok());
        result = store->getBlock(10, 30);
        REQUIRE(!result.status.ok());
        REQUIRE(result.data.empty());
    }
}
