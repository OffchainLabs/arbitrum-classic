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
    cases.push_back({negone * P, sQ, sP, Q});
    cases.push_back({P, negone * sQ, sP, Q});
    cases.push_back({P, sQ, negone * sP, Q});
    cases.push_back({P, sQ, sP, negone * Q});
    return cases;
}

std::vector<ECAddTestCase> prepareECAddCases() {
    alt_bn128_pp::init_public_params();

    G1<alt_bn128_pp> Pff =
        (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();
    G1<alt_bn128_pp> Qff =
        (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();

    auto P = toArbPoint(Pff);
    auto Q = toArbPoint(Qff);
    G1Point sum = toArbPoint(Pff + Qff);
    return {{P, Q, sum}};
}

PairingTestCase::PairingTestCase(const std::string& data) {
    a = {hexToInt({data.begin(), data.begin() + 64}),
         hexToInt({data.begin() + 64, data.begin() + 64 * 2})};

    b = {hexToInt({data.begin() + 64 * 2, data.begin() + 64 * 3}),
         hexToInt({data.begin() + 64 * 3, data.begin() + 64 * 4}),
         hexToInt({data.begin() + 64 * 4, data.begin() + 64 * 5}),
         hexToInt({data.begin() + 64 * 5, data.begin() + 64 * 6})};

    c = {hexToInt({data.begin() + 64 * 6, data.begin() + 64 * 7}),
         hexToInt({data.begin() + 64 * 7, data.begin() + 64 * 8})};

    d = {hexToInt({data.begin() + 64 * 8, data.begin() + 64 * 9}),
         hexToInt({data.begin() + 64 * 9, data.begin() + 64 * 10}),
         hexToInt({data.begin() + 64 * 10, data.begin() + 64 * 11}),
         hexToInt({data.begin() + 64 * 11, data.begin() + 64 * 12})};
}

TEST_CASE("ECOp: g1PfromBytes") {
    initEcOps();
    G1<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();

    auto outP = g1PfromBytes(toArbPoint(P));
    REQUIRE(nonstd::holds_alternative<G1<alt_bn128_pp>>(outP));
    REQUIRE(outP.get<G1<alt_bn128_pp>>() == P);
}

TEST_CASE("ECOp: g2PfromBytes") {
    initEcOps();
    G2<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G2<alt_bn128_pp>::one();

    auto outP = g2PfromBytes(toArbPoint(P));
    REQUIRE(nonstd::holds_alternative<G2<alt_bn128_pp>>(outP));
    REQUIRE(outP.get<G2<alt_bn128_pp>>() == P);
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

        auto g1 = toArbPoint(P[i]);
        auto g2 = toArbPoint(Q[i]);

        all_points.push_back({g1, g2});
    }

    GT<alt_bn128_pp> prod = GT<alt_bn128_pp>::one();

    for (int i = 0; i < numPairs; i++) {
        prod = prod * alt_bn128_pp::reduced_pairing(P[i], Q[i]);
    }
    auto res = ecpairing_internal(all_points);
    REQUIRE(nonstd::holds_alternative<alt_bn128_GT>(res));
    REQUIRE(prod == res.get<alt_bn128_GT>());
}

TEST_CASE("ECOp: ecpairing") {
    for (const auto& testCase : preparePairingCases()) {
        std::vector<std::pair<G1Point, G2Point>> all_points = {
            {testCase.a, testCase.b}, {testCase.c, testCase.d}};

        auto res = ecpairing(all_points);
        std::string msg;
        if (nonstd::holds_alternative<std::string>(res)) {
            msg = res.get<std::string>();
        }
        INFO(msg);
        REQUIRE(nonstd::holds_alternative<bool>(res));
        REQUIRE(res.get<bool>());
    }
}

TEST_CASE("ECOp: ecadd") {
    for (const auto& test_case : prepareECAddCases()) {
        auto res = ecadd(test_case.a, test_case.b);
        REQUIRE(!nonstd::holds_alternative<std::string>(res));
        REQUIRE(nonstd::holds_alternative<G1Point>(res));
        REQUIRE(test_case.res.x == res.get<G1Point>().x);
        REQUIRE(test_case.res.y == res.get<G1Point>().y);
    }
}

TEST_CASE("ECOp: ecmul") {
    for (const auto& test_case : prepareECMulCases()) {
        auto res = ecmul(test_case.a, test_case.k);
        REQUIRE(nonstd::holds_alternative<G1Point>(res));
        REQUIRE(test_case.res.x == res.get<G1Point>().x);
        REQUIRE(test_case.res.y == res.get<G1Point>().y);
    }
}
