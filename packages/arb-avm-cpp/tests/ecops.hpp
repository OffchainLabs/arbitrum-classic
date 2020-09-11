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

#ifndef test_ecops_hpp
#define test_ecops_hpp

#include <gmpxx.h>
#include <avm/machinestate/ecops.hpp>
#include <boost/algorithm/hex.hpp>
#include <libff/algebra/curves/alt_bn128/alt_bn128_pp.hpp>

using namespace libff;

struct PairingTestCase {
    G1Point a;
    G2Point b;
    G1Point c;
    G2Point d;

    PairingTestCase(const G1<alt_bn128_pp>& a_,
                    const G2<alt_bn128_pp>& b_,
                    const G1<alt_bn128_pp>& c_,
                    const G2<alt_bn128_pp>& d_)
        : a(toArbPoint(a_)),
          b(toArbPoint(b_)),
          c(toArbPoint(c_)),
          d(toArbPoint(d_)) {}
};

inline std::vector<PairingTestCase> prepareCases() {
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

inline uint256_t hexToInt(const std::string& hexstr) {
    std::vector<unsigned char> bytes;
    bytes.resize(hexstr.size() / 2);
    boost::algorithm::unhex(hexstr.begin(), hexstr.end(), bytes.begin());
    return intx::be::unsafe::load<uint256_t>(bytes.data());
}

#endif /* test_ecops_hpp */
