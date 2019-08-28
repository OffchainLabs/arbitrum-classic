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

#include "bigint_utils.hpp"

#include <avm/tuple.hpp>
#include <avm/value.hpp>

#include <catch2/catch.hpp>
#include <nlohmann/json.hpp>

#include <boost/algorithm/hex.hpp>

#include <fstream>
#include <iostream>

std::vector<char> hexStringToBytes(const std::string& hexstr) {
    std::vector<char> bytes;
    bytes.reserve(hexstr.size() / 2);
    boost::algorithm::unhex(hexstr.begin(), hexstr.end(), bytes.begin());
    return bytes;
}

TEST_CASE("Value hashing") {
    std::ifstream i(test_cases_path);
    nlohmann::json j;
    i >> j;
    for (auto valtest : j) {
        DYNAMIC_SECTION("Test " << valtest["name"].get<std::string>()) {
            auto valBytes =
                hexStringToBytes(valtest["value"].get<std::string>());
            auto valRaw = valBytes.data();
            uint256_t givenHash = from_hex_str(valtest["hash"]);
            TuplePool pool;
            auto val = deserialize_value(valRaw, pool);
            auto calcHash = hash(val);
            REQUIRE(givenHash == calcHash);
        }
    }

    //    SECTION("Non overlow is correct") { testBinaryOp(4, 3, 1,
    //    OpCode::SUB); }
    //
    //    SECTION("Overlow is correct") { testBinaryOp(3, 4, -1, OpCode::SUB); }
}

TEST_CASE("Value marshaling") {
    std::ifstream i(test_cases_path);
    nlohmann::json j;
    i >> j;
    for (auto valtest : j) {
        DYNAMIC_SECTION("Test " << valtest["name"].get<std::string>()) {
            auto valBytes =
                hexStringToBytes(valtest["value"].get<std::string>());
            auto valRaw = valBytes.data();
            TuplePool pool;
            auto val = deserialize_value(valRaw, pool);
            std::vector<unsigned char> buf;
            marshal_value(val, buf);
            char* valptr = (char*)&buf[0];
            auto newval = deserialize_value(valptr, pool);
            auto valsEqual = val == newval;
            REQUIRE(valsEqual);
            // REQUIRE(val == newval); junit output broken with map::at error
        }
    }
}
