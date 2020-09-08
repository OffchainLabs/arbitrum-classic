//
//  ecops.cpp
//  algebra_bilinearity_test
//
//  Created by Harry Kalodner on 9/8/20.
//

#include <avm/machinestate/ecops.hpp>

#include <gmp.h>
#include <libff/algebra/curves/alt_bn128/alt_bn128_pp.hpp>

#include <catch2/catch.hpp>

using namespace libff;

void mpz_export_and_pad32(uint8_t* output, mpz_t input) {
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

G1Point toArbPoint(G1<alt_bn128_pp> P) {
    P.to_affine_coordinates();

    alt_bn128_Fq X = P.X;
    alt_bn128_Fq Y = P.Y;

    bigint<4L> xbi = X.as_bigint();
    bigint<4l> ybi = Y.as_bigint();

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

G2Point toArbPoint(G2<alt_bn128_pp> P) {
    P.to_affine_coordinates();
    alt_bn128_Fq2 X = P.X;
    alt_bn128_Fq2 Y = P.Y;

    alt_bn128_Fq xc0 = X.c0;
    alt_bn128_Fq xc1 = X.c1;

    alt_bn128_Fq yc0 = Y.c0;
    alt_bn128_Fq yc1 = Y.c1;

    bigint<4L> xc0bi = xc0.as_bigint();
    bigint<4l> xc1bi = xc1.as_bigint();
    bigint<4L> yc0bi = yc0.as_bigint();
    bigint<4l> yc1bi = yc1.as_bigint();

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

struct PairingTestCase {
    G1<alt_bn128_pp> a;
    G2<alt_bn128_pp> b;
    G1<alt_bn128_pp> c;
    G2<alt_bn128_pp> d;
};

TEST_CASE("ECPairing: ecpairing") {
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

    for (const auto& testCase : cases) {
        auto pg1 = toArbPoint(testCase.a);
        auto sqg2 = toArbPoint(testCase.b);
        auto spg1 = toArbPoint(testCase.c);
        auto qg2 = toArbPoint(testCase.d);

        std::vector<std::array<uint256_t, 6>> all_points = {
            {pg1.x, pg1.y, sqg2.x0, sqg2.x1, sqg2.y0, sqg2.y1},
            {spg1.x, spg1.y, qg2.x0, qg2.x1, qg2.y0, qg2.y1}};

        auto res = ecpairing(all_points);
        REQUIRE(nonstd::holds_alternative<bool>(res));
        REQUIRE(res.get<bool>());
    }
}
