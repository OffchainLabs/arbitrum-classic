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

#include <avm/machinestate/ecops.hpp>

#include <gmpxx.h>
#include <libff/algebra/curves/alt_bn128/alt_bn128_pp.hpp>
#include <libff/common/profiling.hpp>

#include <nonstd/variant.hpp>

using namespace libff;

namespace {
struct Init {
    Init() {
        libff::inhibit_profiling_counters = true;
        alt_bn128_pp::init_public_params();
    }
};
static Init init;
}  // namespace

// assumes bytes are big endian
// also assumes 64 bytes, 0..31 for X and 32...64 for Y, representing a curve
// point using affine coordinates if either X or Y is less than 32 bytes, they
// are assumed to be padded with leading 0s
nonstd::variant<G1<alt_bn128_pp>, std::string> g1PfromBytes(
    const G1Point& point) {
    uint8_t xbytes[32];
    intx::be::store(xbytes, point.x);
    uint8_t ybytes[32];
    intx::be::store(ybytes, point.y);

    mpz_t mpzx, mpzy, modulus;
    mpz_inits(mpzx, mpzy, modulus, NULL);

    mpz_import(mpzx, 32, 1, 1, 1, 0, xbytes);
    mpz_import(mpzy, 32, 1, 1, 1, 0, ybytes);

    alt_bn128_Fq::mod.to_mpz(modulus);

    if (mpz_sgn(mpzx) == 0 && mpz_sgn(mpzy) == 0) {
        mpz_clears(mpzx, mpzy, modulus, NULL);
        G1<alt_bn128_pp> P(alt_bn128_Fq::zero(), alt_bn128_Fq::one(),
                           alt_bn128_Fq::zero());
        return P;
    }

    if (mpz_cmp(mpzx, modulus) >= 0) {
        return std::string("bad x");
    }

    if (mpz_cmp(mpzy, modulus) >= 0) {
        return std::string("bad y");
    }

    bigint<alt_bn128_q_limbs> xbi(mpzx);
    bigint<alt_bn128_q_limbs> ybi(mpzy);
    alt_bn128_Fq X(mpzx);
    alt_bn128_Fq Y(mpzy);

    // assumes affine coordinates
    G1<alt_bn128_pp> P = alt_bn128_G1(X, Y, alt_bn128_Fq::one());

    mpz_clears(mpzx, mpzy, modulus, NULL);
    if (!P.is_well_formed()) {
        return std::string("badly formed g1 point");
    }
    return P;
}

// assumes bytes are big endian
// also assumes 128 bytes representing a curve point using affine coordinates
// if either X or Y is less than 64 bytes, they are assumed to be padded with
// leading 0s
nonstd::variant<G2<alt_bn128_pp>, std::string> g2PfromBytes(
    const G2Point& point) {
    uint8_t xc0bytes[32];
    intx::be::store(xc0bytes, point.x0);
    uint8_t xc1bytes[32];
    intx::be::store(xc1bytes, point.x1);
    uint8_t yc0bytes[32];
    intx::be::store(yc0bytes, point.y0);
    uint8_t yc1bytes[32];
    intx::be::store(yc1bytes, point.y1);

    mpz_t mpzxc0, mpzxc1, mpzyc0, mpzyc1, modulus;
    mpz_inits(mpzxc0, mpzxc1, mpzyc0, mpzyc1, modulus, NULL);
    mpz_import(mpzxc0, 32, 1, 1, 1, 0, xc0bytes);
    mpz_import(mpzxc1, 32, 1, 1, 1, 0, xc1bytes);
    mpz_import(mpzyc0, 32, 1, 1, 1, 0, yc0bytes);
    mpz_import(mpzyc1, 32, 1, 1, 1, 0, yc1bytes);

    alt_bn128_Fq::mod.to_mpz(modulus);

    if (mpz_sgn(mpzxc0) == 0 && mpz_sgn(mpzxc1) == 0 && mpz_sgn(mpzyc0) == 0 &&
        mpz_sgn(mpzyc1) == 0) {
        mpz_clears(mpzxc0, mpzxc1, mpzyc0, mpzyc1, modulus, NULL);
        G2<alt_bn128_pp> P(alt_bn128_Fq2::zero(), alt_bn128_Fq2::one(),
                           alt_bn128_Fq2::zero());
        return P;
    }

    if (mpz_cmp(mpzxc0, modulus) >= 0) {
        return std::string("bad x0");
    }

    if (mpz_cmp(mpzxc1, modulus) >= 0) {
        return std::string("bad x1");
    }

    if (mpz_cmp(mpzyc0, modulus) >= 0) {
        return std::string("bad y0");
    }

    if (mpz_cmp(mpzyc1, modulus) >= 0) {
        return std::string("bad y1");
    }

    bigint<alt_bn128_q_limbs> xc0bi(mpzxc0);
    bigint<alt_bn128_q_limbs> xc1bi(mpzxc1);
    bigint<alt_bn128_q_limbs> yc0bi(mpzyc0);
    bigint<alt_bn128_q_limbs> yc1bi(mpzyc1);

    alt_bn128_Fq Xc0(mpzxc0);
    alt_bn128_Fq Xc1(mpzxc1);
    alt_bn128_Fq Yc0(mpzyc0);
    alt_bn128_Fq Yc1(mpzyc1);

    alt_bn128_Fq2 X(Xc0, Xc1);
    alt_bn128_Fq2 Y(Yc0, Yc1);

    // assumes affine coordinates
    G2<alt_bn128_pp> P = alt_bn128_G2(X, Y, alt_bn128_Fq2::one());

    mpz_clears(mpzxc0, mpzxc1, mpzyc0, mpzyc1, modulus, NULL);
    if (!P.is_well_formed()) {
        return std::string("badly formed g2 point");
    }
    return P;
}

nonstd::variant<alt_bn128_GT, std::string> ecpairing_internal(
    const std::vector<std::array<uint256_t, 6>>& input) {
    alt_bn128_Fq12 prod = alt_bn128_Fq12::one();

    for (const auto& item : input) {
        auto g1 = g1PfromBytes({item[0], item[1]});
        auto g2 = g2PfromBytes({item[2], item[3], item[4], item[5]});
        if (nonstd::holds_alternative<std::string>(g1)) {
            return g1.get<std::string>();
        }
        if (nonstd::holds_alternative<std::string>(g2)) {
            return g2.get<std::string>();
        }
        prod = prod * alt_bn128_pp::pairing(g1.get<G1<alt_bn128_pp>>(),
                                            g2.get<G2<alt_bn128_pp>>());
    }

    return alt_bn128_final_exponentiation(prod);
}

nonstd::variant<bool, std::string> ecpairing(
    const std::vector<std::array<uint256_t, 6>>& input) {
    auto res = ecpairing_internal(input);
    if (nonstd::holds_alternative<std::string>(res)) {
        return res.get<std::string>();
    }
    return res.get<alt_bn128_GT>() == GT<alt_bn128_pp>::one();
}

nonstd::variant<alt_bn128_G1, std::string> ecadd(
    const std::array<uint256_t, 4>& input) {
    auto a = g1PfromBytes({input[0], input[1]});
    auto b = g1PfromBytes({input[2], input[3]});
    if (nonstd::holds_alternative<std::string>(a)) {
        return a.get<std::string>();
    }
    if (nonstd::holds_alternative<std::string>(b)) {
        return b.get<std::string>();
    }
    return a.get<G1<alt_bn128_pp>>() + b.get<G1<alt_bn128_pp>>();
}

nonstd::variant<alt_bn128_G1, std::string> ecsmult(
    const std::array<uint256_t, 3>& input) {
    auto a = g1PfromBytes({input[0], input[1]});

    if (nonstd::holds_alternative<std::string>(a)) {
        return a.get<std::string>();
    }

    uint8_t sbytes[32];
    intx::be::store(sbytes, input[2]);

    mpz_t mpzs;
    mpz_init(mpzs);
    mpz_import(mpzs, 32, 1, 1, 1, 0, sbytes);
    bigint<BIG_INT_FOR_UINT256> s(mpzs);
    mpz_clear(mpzs);
    return s * a.get<G1<alt_bn128_pp>>();
}
