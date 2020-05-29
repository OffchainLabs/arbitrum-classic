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

#include <data_storage/datastorage.hpp>
#include <data_storage/nodestore.hpp>
#include <data_storage/storageresult.hpp>

#include <catch2/catch.hpp>

TEST_CASE("NodeStore tests") {
    DBDeleter deleter;
    auto storage = std::make_shared<DataStorage>(dbpath);
    auto store = std::make_unique<NodeStore>(storage);

    SECTION("isEmpty") {
        REQUIRE(store->isEmpty());
        REQUIRE(store->putNode(0, 30, {1, 2, 3}).ok());
        REQUIRE(!store->isEmpty());
    }

    SECTION("NodeStore isEmpty and maxNodeHeight") {
        // It should start out empty with max height 0
        REQUIRE(store->maxNodeHeight() == 0);

        // After adding a node at height 0, max height should be the same
        REQUIRE(store->putNode(0, 30, {1, 2, 3}).ok());
        REQUIRE(store->maxNodeHeight() == 0);

        // After adding a node at height 1, the max height is 1
        REQUIRE(store->putNode(1, 31, {1, 2, 3}).ok());
        REQUIRE(store->maxNodeHeight() == 1);

        // After adding a node at height 10, the max height is 10
        REQUIRE(store->putNode(10, 32, {1, 2, 3}).ok());
        REQUIRE(store->maxNodeHeight() == 10);
    }

    SECTION("getHeight") {
        REQUIRE(store->putNode(3, 30, {1, 2, 3}).ok());
        auto res = store->getHeight(30);
        REQUIRE(res.status.ok());
        REQUIRE(res.data == 3);
    }

    SECTION("getHash") {
        REQUIRE(store->putNode(3, 30, {1, 2, 3}).ok());
        auto res = store->getHash(3);
        REQUIRE(res.status.ok());
        REQUIRE(res.data == 30);
    }

    SECTION("getNode") {
        REQUIRE(store->putNode(3, 30, {1, 2, 3}).ok());
        auto res = store->getNode(3, 30);
        REQUIRE(res.status.ok());
        REQUIRE(res.data == std::vector<unsigned char>{1, 2, 3});
    }

    SECTION("overwrite node") {
        REQUIRE(store->putNode(3, 30, {1, 2, 3}).ok());
        REQUIRE(store->putNode(3, 30, {2, 3, 4}).ok());
        auto res = store->getNode(3, 30);
        REQUIRE(res.status.ok());
        REQUIRE(res.data == std::vector<unsigned char>{2, 3, 4});
    }
}
