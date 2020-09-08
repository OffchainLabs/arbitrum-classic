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

TEST_CASE("ECPairing: g1PfromBytes") {
    alt_bn128_pp::init_public_params();

    G1<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();

    auto outP = g1PfromBytes(toArbPoint(P));
    REQUIRE(nonstd::holds_alternative<G1<alt_bn128_pp>>(outP));
    REQUIRE(outP.get<G1<alt_bn128_pp>>() == P);
}

TEST_CASE("ECPairing: g2PfromBytes") {
    alt_bn128_pp::init_public_params();

    G2<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G2<alt_bn128_pp>::one();

    auto outP = g2PfromBytes(toArbPoint(P));
    REQUIRE(nonstd::holds_alternative<G2<alt_bn128_pp>>(outP));
    REQUIRE(outP.get<G2<alt_bn128_pp>>() == P);
}

TEST_CASE("ECPairing: ecpairing_internal") {
    alt_bn128_pp::init_public_params();

    constexpr int numPairs = 3;

    std::array<G1<alt_bn128_pp>, numPairs> P;
    std::array<G2<alt_bn128_pp>, numPairs> Q;

    std::vector<std::array<uint256_t, 6>> all_points;

    for (int i = 0; i < numPairs; i++) {
        P[i] = (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();
        Q[i] = (Fr<alt_bn128_pp>::random_element()) * G2<alt_bn128_pp>::one();

        auto g1 = toArbPoint(P[i]);
        auto g2 = toArbPoint(Q[i]);

        all_points.push_back({g1.x, g1.y, g2.x0, g2.x1, g2.y0, g2.y1});
    }

    GT<alt_bn128_pp> prod = GT<alt_bn128_pp>::one();

    for (int i = 0; i < numPairs; i++) {
        prod = prod * alt_bn128_pp::reduced_pairing(P[i], Q[i]);
    }
    auto res = ecpairing_internal(all_points);
    REQUIRE(nonstd::holds_alternative<alt_bn128_GT>(res));
    REQUIRE(prod == res.get<alt_bn128_GT>());
}

TEST_CASE("ECPairing: ecpairing") {
    for (const auto& testCase : prepareCases()) {
        std::vector<std::array<uint256_t, 6>> all_points = {
            {testCase.a.x, testCase.a.y, testCase.b.x0, testCase.b.x1,
             testCase.b.y0, testCase.b.y1},
            {testCase.c.x, testCase.c.y, testCase.d.x0, testCase.d.x1,
             testCase.d.y0, testCase.d.y1}};

        auto res = ecpairing(all_points);
        REQUIRE(nonstd::holds_alternative<bool>(res));
        REQUIRE(res.get<bool>());
    }
}
