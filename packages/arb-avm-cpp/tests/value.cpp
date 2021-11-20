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

#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

#include <catch2/catch.hpp>
#include <nlohmann/json.hpp>

#include <boost/algorithm/hex.hpp>

#include <fstream>
#include <iostream>

std::vector<char> hexStringToBytes(const std::string& hexstr) {
    std::vector<char> bytes;
    boost::algorithm::unhex(hexstr.begin(), hexstr.end(),
                            std::back_inserter(bytes));
    return bytes;
}

TEST_CASE("Value hashing") {
    std::ifstream i(test_cases_path);
    nlohmann::json j;
    i >> j;
    for (auto valtest : j) {
        // Don't include codepoints in test
        if (valtest["is_codepoint"]) {
            continue;
        }
        DYNAMIC_SECTION("Test " << valtest["name"].get<std::string>()) {
            auto valBytes =
                hexStringToBytes(valtest["value"].get<std::string>());
            auto valRaw = reinterpret_cast<const char*>(valBytes.data());
            uint256_t givenHash = intx::from_string<uint256_t>(
                std::string("0x") + valtest["hash"].get<std::string>());
            auto val = deserialize_value(valRaw);
            auto calcHash = hash_value(val);
            REQUIRE(givenHash == calcHash);
        }
    }
}

// Test is disabled since it it incompatible with the new codepoint system
TEST_CASE("Value marshaling") {
    std::ifstream i(test_cases_path);
    nlohmann::json j;
    i >> j;
    for (auto valtest : j) {
        // Don't include codepoints in test
        if (valtest["is_codepoint"]) {
            continue;
        }
        DYNAMIC_SECTION("Test " << valtest["name"].get<std::string>()) {
            auto valBytes =
                hexStringToBytes(valtest["value"].get<std::string>());
            auto valRaw = reinterpret_cast<const char*>(valBytes.data());
            auto val = deserialize_value(valRaw);
            std::vector<unsigned char> buf;
            marshal_value(val, buf);
            auto valptr = (const char*)&buf[0];
            auto newval = deserialize_value(valptr);
            REQUIRE(values_equal(val, newval));
        }
    }
}

TEST_CASE("UnloadedValue equality") {
    Tuple tup;
    UnloadedValue uv(BigUnloadedValue{TUPLE, ::hash(tup), tup.getSize()});
    REQUIRE(values_equal(value(tup), value(uv)));
}

void checkUvIntegrity(BigUnloadedValue big, bool shouldHeap) {
    UnloadedValue uv(big);
    REQUIRE(uv.hash() == big.hash);
    REQUIRE(uv.value_size() == big.value_size);
    REQUIRE(uv.type() == big.type);
    REQUIRE(uv.isHeaped() == shouldHeap);

    UnloadedValue copy(uv);
    REQUIRE(copy.hash() == big.hash);
    REQUIRE(copy.value_size() == big.value_size);
    REQUIRE(copy.type() == big.type);
    REQUIRE(copy.isHeaped() == shouldHeap);

    UnloadedValue move(std::move(copy));
    REQUIRE(move.hash() == big.hash);
    REQUIRE(move.value_size() == big.value_size);
    REQUIRE(move.type() == big.type);
    REQUIRE(move.isHeaped() == shouldHeap);

    // Make sure uv isn't moved (though C++ probably wouldn't allow that
    // anyways)
    REQUIRE(uv.hash() == big.hash);
}

TEST_CASE("UnloadedValue inlining") {
    Tuple tup;
    checkUvIntegrity({TUPLE, ::hash(tup), 8}, false);
    checkUvIntegrity({HASH_PRE_IMAGE, ::hash(tup), 1}, true);
    checkUvIntegrity({HASH_PRE_IMAGE, ::hash(tup), uint256_t(1) << 128}, true);
}
