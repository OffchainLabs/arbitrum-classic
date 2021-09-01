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

using namespace libff;

void initEcOps() {
    static bool initialized;
    if (!initialized) {
        inhibit_profiling_counters = true;
        alt_bn128_pp::init_public_params();
        initialized = true;
    }
}

void mpz_export_and_pad32(uint8_t* output, mpz_t input) {
    uint8_t tmp[32];
    size_t countp;
    mpz_export(tmp, &countp, 1, 1, 1, 0, input);

    size_t padding = 32 - countp;
    for (size_t i = 0; i < padding; i++) {
        output[i] = 0;
    }
    for (size_t i = 0; i < countp; i++) {
        output[i + padding] = tmp[i];
    }
}

G1Point toG1ArbPoint(G1<alt_bn128_pp> P) {
    if (P.is_zero()) {
        // This is a special case defined in EIP-196.
        // libff would encode this as {0, 1}
        return {0, 0};
    }

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

G2Point toG2ArbPoint(G2<alt_bn128_pp> P) {
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

// assumes bytes are big endian
// also assumes 64 bytes, 0..31 for X and 32...64 for Y, representing a curve
// point using affine coordinates if either X or Y is less than 32 bytes, they
// are assumed to be padded with leading 0s
std::variant<G1<alt_bn128_pp>, std::string> g1PfromBytes(const G1Point& point) {
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
        return G1<alt_bn128_pp>::zero();
    }

    if (mpz_cmp(mpzx, modulus) >= 0) {
        mpz_clears(mpzx, mpzy, modulus, NULL);
        return std::string("bad x");
    }

    if (mpz_cmp(mpzy, modulus) >= 0) {
        mpz_clears(mpzx, mpzy, modulus, NULL);
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
std::variant<G2<alt_bn128_pp>, std::string> g2PfromBytes(const G2Point& point) {
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
        return G2<alt_bn128_pp>::zero();
    }

    if (mpz_cmp(mpzxc0, modulus) >= 0) {
        mpz_clears(mpzxc0, mpzxc1, mpzyc0, mpzyc1, modulus, NULL);
        return std::string("bad x0");
    }

    if (mpz_cmp(mpzxc1, modulus) >= 0) {
        mpz_clears(mpzxc0, mpzxc1, mpzyc0, mpzyc1, modulus, NULL);
        return std::string("bad x1");
    }

    if (mpz_cmp(mpzyc0, modulus) >= 0) {
        mpz_clears(mpzxc0, mpzxc1, mpzyc0, mpzyc1, modulus, NULL);
        return std::string("bad y0");
    }

    if (mpz_cmp(mpzyc1, modulus) >= 0) {
        mpz_clears(mpzxc0, mpzxc1, mpzyc0, mpzyc1, modulus, NULL);
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

std::variant<alt_bn128_GT, std::string> ecpairing_internal(
    const std::vector<std::pair<G1Point, G2Point>>& input) {
    alt_bn128_Fq12 prod = alt_bn128_Fq12::one();

    for (const auto& item : input) {
        auto g1 = g1PfromBytes(item.first);
        auto g2 = g2PfromBytes(item.second);
        if (std::holds_alternative<std::string>(g1)) {
            return std::get<std::string>(g1);
        }
        if (std::holds_alternative<std::string>(g2)) {
            return std::get<std::string>(g2);
        }
        auto g1p = std::get<G1<alt_bn128_pp>>(g1);
        auto g2p = std::get<G2<alt_bn128_pp>>(g2);
        if (!g1p.is_zero() && !g2p.is_zero()) {
            prod = prod * alt_bn128_pp::pairing(g1p, g2p);
        }
    }

    return alt_bn128_final_exponentiation(prod);
}

std::variant<bool, std::string> ecpairing(
    const std::vector<std::pair<G1Point, G2Point>>& input) {
    initEcOps();
    auto res = ecpairing_internal(input);
    if (std::holds_alternative<std::string>(res)) {
        return std::get<std::string>(res);
    }
    return std::get<alt_bn128_GT>(res) == GT<alt_bn128_pp>::one();
}

std::variant<G1Point, std::string> ecadd(const G1Point& input_a,
                                         const G1Point& input_b) {
    initEcOps();
    auto a = g1PfromBytes(input_a);
    auto b = g1PfromBytes(input_b);
    if (std::holds_alternative<std::string>(a)) {
        return std::get<std::string>(a);
    }
    if (std::holds_alternative<std::string>(b)) {
        return std::get<std::string>(b);
    }
    return toG1ArbPoint(std::get<G1<alt_bn128_pp>>(a) +
                        std::get<G1<alt_bn128_pp>>(b));
}

std::variant<G1Point, std::string> ecmul(const G1Point& point,
                                         const uint256_t& factor) {
    initEcOps();
    auto a = g1PfromBytes(point);

    if (std::holds_alternative<std::string>(a)) {
        return std::get<std::string>(a);
    }

    uint8_t sbytes[32];
    intx::be::store(sbytes, factor);

    mpz_t mpzs;
    mpz_init(mpzs);
    mpz_import(mpzs, 32, 1, 1, 1, 0, sbytes);
    bigint<BIG_INT_FOR_UINT256> s(mpzs);
    mpz_clear(mpzs);
    return toG1ArbPoint(s * std::get<G1<alt_bn128_pp>>(a));
}

std::ostream& operator<<(std::ostream& os, const G1Point& val) {
    os << "(" << intx::to_string(val.x) << ", " << intx::to_string(val.y)
       << ")";
    return os;
}

std::ostream& operator<<(std::ostream& os, const G2Point& val) {
    os << "(" << intx::to_string(val.x0) << ", " << intx::to_string(val.x1)
       << ", " << intx::to_string(val.y0) << ", " << intx::to_string(val.y1)
       << ")";
    return os;
}
