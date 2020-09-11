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

inline void mpz_export_and_pad32(uint8_t* output, mpz_t input) {
    size_t countp;
    mpz_export(output, &countp, 1, 1, 1, 0, input);

    if (countp < 32) {
        uint8_t difference = 32 - countp;

        for (int i = 0; i < 32 - difference; i++) {
            output[32 - difference - i] = output[32 - difference - i - 1];
        }

        for (int i = 0; i < difference; i++) {
            output[i] = 0;
        }
    }
}

inline G1Point toArbPoint(G1<alt_bn128_pp> P) {
    P.to_affine_coordinates();

    alt_bn128_Fq X = P.X;
    alt_bn128_Fq Y = P.Y;

    bigint<alt_bn128_q_limbs> xbi = X.as_bigint();
    bigint<alt_bn128_q_limbs> ybi = Y.as_bigint();

    mpz_t mpx, mpy;
    mpz_inits(mpx, mpy, NULL);
    xbi.to_mpz(mpx);
    ybi.to_mpz(mpy);

    uint8_t xbytes[32];
    uint8_t ybytes[32];

    mpz_export_and_pad32(xbytes, mpx);
    mpz_export_and_pad32(ybytes, mpy);

    auto x_int = intx::be::load<uint256_t>(xbytes);
    auto y_int = intx::be::load<uint256_t>(ybytes);

    mpz_clears(mpx, mpy, NULL);

    return {x_int, y_int};
}

inline G2Point toArbPoint(G2<alt_bn128_pp> P) {
    P.to_affine_coordinates();
    alt_bn128_Fq2 X = P.X;
    alt_bn128_Fq2 Y = P.Y;

    alt_bn128_Fq xc0 = X.c0;
    alt_bn128_Fq xc1 = X.c1;

    alt_bn128_Fq yc0 = Y.c0;
    alt_bn128_Fq yc1 = Y.c1;

    bigint<alt_bn128_q_limbs> xc0bi = xc0.as_bigint();
    bigint<alt_bn128_q_limbs> xc1bi = xc1.as_bigint();
    bigint<alt_bn128_q_limbs> yc0bi = yc0.as_bigint();
    bigint<alt_bn128_q_limbs> yc1bi = yc1.as_bigint();

    mpz_t mpzxc0, mpzxc1, mpzyc0, mpzyc1;
    mpz_inits(mpzxc0, mpzxc1, mpzyc0, mpzyc1, NULL);

    xc0bi.to_mpz(mpzxc0);
    xc1bi.to_mpz(mpzxc1);
    yc0bi.to_mpz(mpzyc0);
    yc1bi.to_mpz(mpzyc1);

    uint8_t raw_bytes[32];
    mpz_export_and_pad32(raw_bytes, mpzxc0);
    auto x0_int = intx::be::load<uint256_t>(raw_bytes);
    mpz_export_and_pad32(raw_bytes, mpzxc1);
    auto x1_int = intx::be::load<uint256_t>(raw_bytes);
    mpz_export_and_pad32(raw_bytes, mpzyc0);
    auto y0_int = intx::be::load<uint256_t>(raw_bytes);
    mpz_export_and_pad32(raw_bytes, mpzyc1);
    auto y1_int = intx::be::load<uint256_t>(raw_bytes);

    mpz_clears(mpzxc0, mpzxc1, mpzyc0, mpzyc1, NULL);
    return {x0_int, x1_int, y0_int, y1_int};
}

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
