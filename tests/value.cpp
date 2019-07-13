//
//  value.cpp
//  AVMtest
//
//  Created by Harry Kalodner on 6/23/19.
//

#include "config.hpp"

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
        DYNAMIC_SECTION("Test " << valtest["name"]) {
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
        DYNAMIC_SECTION("Test " << valtest["name"]) {
            auto valBytes =
                hexStringToBytes(valtest["value"].get<std::string>());
            auto valRaw = valBytes.data();
            TuplePool pool;
            auto val = deserialize_value(valRaw, pool);
            std::vector<unsigned char> buf;
            marshal_value(val, buf);
            char* valptr = (char*)&buf[0];
            auto newval = deserialize_value(valptr, pool);
            REQUIRE(val == newval);
        }
    }
}
