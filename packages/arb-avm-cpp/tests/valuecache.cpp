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

#include <data_storage/value/valuecache.hpp>

#include <avm/inboxmessage.hpp>

#include <avm_values/vmValueParser.hpp>

#include <catch2/catch.hpp>
#include <nlohmann/json.hpp>

TEST_CASE("ValueCache reload") {
    ValueCache cache{4, 0};

    Value val1{CodePointStub({0, 0}, 0)};
    // Cache 1
    cache.maybeSave(val1);
    auto retrieved_val = cache.loadIfExists(hash_value(val1));
    REQUIRE(retrieved_val != std::nullopt);
    REQUIRE(hash_value(val1) == hash_value(*retrieved_val));

    // Cache 2
    cache.nextCache();
    retrieved_val = cache.loadIfExists(hash_value(val1));
    REQUIRE(retrieved_val != std::nullopt);
    REQUIRE(hash_value(val1) == hash_value(*retrieved_val));

    // Cache 3
    // Leave cache empty
    cache.nextCache();

    // Cache 4
    // Leave cache empty
    cache.nextCache();

    // Cache 1
    cache.nextCache();
    // Loads from cache 2
    retrieved_val = cache.loadIfExists(hash_value(val1));
    REQUIRE(retrieved_val != std::nullopt);
    REQUIRE(hash_value(val1) == hash_value(*retrieved_val));

    // Cache 2
    cache.nextCache();
    // Loads from cache 1
    retrieved_val = cache.loadIfExists(hash_value(val1));
    REQUIRE(retrieved_val != std::nullopt);
    REQUIRE(hash_value(val1) == hash_value(*retrieved_val));
}

TEST_CASE("ValueCache reset") {
    ValueCache cache{2, 0};

    Value val1{CodePointStub({0, 0}, 0)};
    // Cache 1
    cache.maybeSave(val1);
    REQUIRE(cache.loadIfExists(hash_value(val1)) != std::nullopt);
    REQUIRE(values_equal(*cache.loadIfExists(hash_value(val1)), val1));

    // Cache 2
    // Leave cache empty
    cache.nextCache();

    // Cache 1
    cache.nextCache();
    REQUIRE(cache.loadIfExists(hash_value(val1)) == std::nullopt);

    // Cache 2
    cache.nextCache();
    REQUIRE(cache.loadIfExists(hash_value(val1)) == std::nullopt);
}

TEST_CASE("ValueCache reload max") {
    ValueCache cache{4, 2};

    Value val1{CodePointStub({0, 0}, 0)};
    Value val2{uint256_t(42)};
    // Cache 1
    cache.maybeSave(val1);
    cache.maybeSave(val2);

    // Cache 2
    // max elements in cache, now using cache 2
    // loads from cache 1
    auto retrieved_val = cache.loadIfExists(hash_value(val1));
    REQUIRE(retrieved_val != std::nullopt);
    REQUIRE(hash_value(val1) == hash_value(*retrieved_val));
    retrieved_val = cache.loadIfExists(hash_value(val2));
    REQUIRE(retrieved_val != std::nullopt);
    REQUIRE(hash_value(val2) == hash_value(*retrieved_val));

    // Cache 3
    // max elements in cache, now using cache 3
    // Leave cache empty
    cache.nextCache();

    // Cache 4
    // Leave cache empty
    cache.nextCache();

    // Cache 1
    // Leave cache empty
    cache.nextCache();

    // Cache 2
    // Loads from cache 2
    retrieved_val = cache.loadIfExists(hash_value(val1));
    REQUIRE(retrieved_val == std::nullopt);
}

TEST_CASE("ValueCache disabled") {
    ValueCache cache{0, 0};

    Value val1{CodePointStub({0, 0}, 0)};
    // Cache 1
    cache.maybeSave(val1);
    REQUIRE(cache.loadIfExists(hash_value(val1)) == std::nullopt);
    cache.nextCache();
    REQUIRE(cache.loadIfExists(hash_value(val1)) == std::nullopt);
}
