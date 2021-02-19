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

#include "ecops.hpp"

#include <catch2/catch.hpp>

using namespace libff;

std::vector<PairingTestCase> preparePairingCases() {
    G1<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();
    G2<alt_bn128_pp> Q =
        (Fr<alt_bn128_pp>::random_element()) * G2<alt_bn128_pp>::one();

    Fr<alt_bn128_pp> s = Fr<alt_bn128_pp>::random_element();
    Fr<alt_bn128_pp> negone = Fr<alt_bn128_pp>(-1);

    // Fr<ppT> s = Fr<ppT>("2");
    G1<alt_bn128_pp> sP = s * P;
    G2<alt_bn128_pp> sQ = s * Q;

    std::vector<PairingTestCase> cases;
    cases.push_back({{{negone * P, sQ}, {sP, Q}}, true});
    cases.push_back({{{P, negone * sQ}, {sP, Q}}, true});
    cases.push_back({{{P, sQ}, {negone * sP, Q}}, true});
    cases.push_back({{{P, sQ}, {sP, negone * Q}}, true});
    return cases;
}

std::vector<ECAddTestCase> prepareECAddCases() {
    alt_bn128_pp::init_public_params();

    G1<alt_bn128_pp> Pff =
        (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();
    G1<alt_bn128_pp> Qff =
        (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();

    auto P = toG1ArbPoint(Pff);
    auto Q = toG1ArbPoint(Qff);
    G1Point sum = toG1ArbPoint(Pff + Qff);
    return {{P, Q, sum}};
}

PairingTestCase::PairingTestCase(const std::string& data, bool valid_)
    : valid(valid_) {
    G1Point a = {hexToInt({data.begin(), data.begin() + 64}),
                 hexToInt({data.begin() + 64, data.begin() + 64 * 2})};

    G2Point b = {hexToInt({data.begin() + 64 * 2, data.begin() + 64 * 3}),
                 hexToInt({data.begin() + 64 * 3, data.begin() + 64 * 4}),
                 hexToInt({data.begin() + 64 * 4, data.begin() + 64 * 5}),
                 hexToInt({data.begin() + 64 * 5, data.begin() + 64 * 6})};

    G1Point c = {hexToInt({data.begin() + 64 * 6, data.begin() + 64 * 7}),
                 hexToInt({data.begin() + 64 * 7, data.begin() + 64 * 8})};

    G2Point d = {hexToInt({data.begin() + 64 * 8, data.begin() + 64 * 9}),
                 hexToInt({data.begin() + 64 * 9, data.begin() + 64 * 10}),
                 hexToInt({data.begin() + 64 * 10, data.begin() + 64 * 11}),
                 hexToInt({data.begin() + 64 * 11, data.begin() + 64 * 12})};

    points.push_back({a, b});
    points.push_back({c, d});
}

TEST_CASE("ECOp: g1PfromBytes") {
    initEcOps();
    G1<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();

    auto outP = g1PfromBytes(toG1ArbPoint(P));
    REQUIRE(std::holds_alternative<G1<alt_bn128_pp>>(outP));
    REQUIRE(std::get<G1<alt_bn128_pp>>(outP) == P);
}

TEST_CASE("ECOp: g2PfromBytes") {
    initEcOps();
    G2<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G2<alt_bn128_pp>::one();

    auto outP = g2PfromBytes(toG2ArbPoint(P));
    REQUIRE(std::holds_alternative<G2<alt_bn128_pp>>(outP));
    REQUIRE(std::get<G2<alt_bn128_pp>>(outP) == P);
}

TEST_CASE("ECOp: ecpairing_internal") {
    initEcOps();
    constexpr int numPairs = 3;

    std::array<G1<alt_bn128_pp>, numPairs> P;
    std::array<G2<alt_bn128_pp>, numPairs> Q;

    std::vector<std::pair<G1Point, G2Point>> all_points;

    for (int i = 0; i < numPairs; i++) {
        P[i] = (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();
        Q[i] = (Fr<alt_bn128_pp>::random_element()) * G2<alt_bn128_pp>::one();

        auto g1 = toG1ArbPoint(P[i]);
        auto g2 = toG2ArbPoint(Q[i]);

        all_points.push_back({g1, g2});
    }

    GT<alt_bn128_pp> prod = GT<alt_bn128_pp>::one();

    for (int i = 0; i < numPairs; i++) {
        prod = prod * alt_bn128_pp::reduced_pairing(P[i], Q[i]);
    }
    auto res = ecpairing_internal(all_points);
    REQUIRE(std::holds_alternative<alt_bn128_GT>(res));
    REQUIRE(prod == std::get<alt_bn128_GT>(res));
}

TEST_CASE("ECOp: ecpairing") {
    for (const auto& test_case : preparePairingCases()) {
        auto res = ecpairing(test_case.points);
        std::string msg;
        if (std::holds_alternative<std::string>(res)) {
            msg = std::get<std::string>(res);
        }
        INFO(msg);
        REQUIRE(std::holds_alternative<bool>(res));
        REQUIRE(std::get<bool>(res));
    }
}

TEST_CASE("ECOp: ecadd") {
    for (const auto& test_case : prepareECAddCases()) {
        auto res = ecadd(test_case.a, test_case.b);
        REQUIRE(!std::holds_alternative<std::string>(res));
        REQUIRE(std::holds_alternative<G1Point>(res));
        REQUIRE(test_case.res.x == std::get<G1Point>(res).x);
        REQUIRE(test_case.res.y == std::get<G1Point>(res).y);
    }
}

TEST_CASE("ECOp: ecmul") {
    for (const auto& test_case : prepareECMulCases()) {
        auto res = ecmul(test_case.a, test_case.k);
        REQUIRE(std::holds_alternative<G1Point>(res));
        REQUIRE(test_case.res.x == std::get<G1Point>(res).x);
        REQUIRE(test_case.res.y == std::get<G1Point>(res).y);
    }
}
