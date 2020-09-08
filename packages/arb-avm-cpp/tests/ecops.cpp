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

std::vector<uint8_t> toBytes(G1<alt_bn128_pp> P) {
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

    std::vector<uint8_t> xvec(xbytes, xbytes + 32);
    std::vector<uint8_t> yvec(ybytes, ybytes + 32);
    xvec.insert(xvec.end(), yvec.begin(), yvec.end());

    mpz_clears(mpx, mpy, NULL);

    return xvec;
}

std::vector<uint8_t> toBytes(G2<alt_bn128_pp> P) {
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

    uint8_t uxbytes[64];
    uint8_t uybytes[64];

    mpz_export_and_pad32(uxbytes, mpzxc0);
    mpz_export_and_pad32(uxbytes + 32, mpzxc1);
    mpz_export_and_pad32(uybytes, mpzyc0);
    mpz_export_and_pad32(uybytes + 32, mpzyc1);

    std::vector<uint8_t> xvec(uxbytes, uxbytes + 64);
    std::vector<uint8_t> yvec(uybytes, uybytes + 64);
    xvec.insert(xvec.end(), yvec.begin(), yvec.end());

    mpz_clears(mpzxc0, mpzxc1, mpzyc0, mpzyc1, NULL);
    return xvec;
}

TEST_CASE("ECPairing: g1PfromBytes") {
    alt_bn128_pp::init_public_params();

    G1<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();

    G1<alt_bn128_pp> outP = g1PfromBytes(toArbPoint(P));

    REQUIRE(outP == P);
}

TEST_CASE("ECPairing: g2PfromBytes") {
    alt_bn128_pp::init_public_params();

    G2<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G2<alt_bn128_pp>::one();

    G2<alt_bn128_pp> outP = g2PfromBytes(toArbPoint(P));

    REQUIRE(outP == P);
}

TEST_CASE("ECPairing: ecpairing_internal") {
    alt_bn128_pp::init_public_params();

    int numPairs = 3;

    G1<alt_bn128_pp> P[numPairs];
    G2<alt_bn128_pp> Q[numPairs];

    std::vector<uint8_t> allbytes;

    for (int i = 0; i < numPairs; i++) {
        P[i] = (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();
        Q[i] = (Fr<alt_bn128_pp>::random_element()) * G2<alt_bn128_pp>::one();

        std::vector<uint8_t> pbytes = toBytes(P[i]);
        std::vector<uint8_t> qbytes = toBytes(Q[i]);

        allbytes.insert(allbytes.end(), pbytes.begin(), pbytes.end());
        allbytes.insert(allbytes.end(), qbytes.begin(), qbytes.end());
    }

    GT<alt_bn128_pp> prod = GT<alt_bn128_pp>::one();

    for (int i = 0; i < numPairs; i++) {
        prod = prod * alt_bn128_pp::reduced_pairing(P[i], Q[i]);
    }
    REQUIRE(prod == ecpairing_internal(allbytes));
}

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

    std::vector<uint8_t> pbytes = toBytes(negone * P);
    std::vector<uint8_t> sqbytes = toBytes(sQ);

    std::vector<uint8_t> spbytes = toBytes(sP);
    std::vector<uint8_t> qbytes = toBytes(Q);

    std::vector<uint8_t> allbytes;
    allbytes.insert(allbytes.end(), pbytes.begin(), pbytes.end());
    allbytes.insert(allbytes.end(), sqbytes.begin(), sqbytes.end());
    allbytes.insert(allbytes.end(), spbytes.begin(), spbytes.end());
    allbytes.insert(allbytes.end(), qbytes.begin(), qbytes.end());

    REQUIRE(ecpairing(allbytes) != 0);

    pbytes = toBytes(P);
    sqbytes = toBytes(negone * sQ);
    spbytes = toBytes(sP);
    qbytes = toBytes(Q);

    allbytes.clear();

    allbytes.insert(allbytes.end(), pbytes.begin(), pbytes.end());
    allbytes.insert(allbytes.end(), sqbytes.begin(), sqbytes.end());
    allbytes.insert(allbytes.end(), spbytes.begin(), spbytes.end());
    allbytes.insert(allbytes.end(), qbytes.begin(), qbytes.end());

    REQUIRE(ecpairing(allbytes) != 0);

    pbytes = toBytes(P);
    sqbytes = toBytes(sQ);
    spbytes = toBytes(negone * sP);
    qbytes = toBytes(Q);

    allbytes.clear();

    allbytes.insert(allbytes.end(), pbytes.begin(), pbytes.end());
    allbytes.insert(allbytes.end(), sqbytes.begin(), sqbytes.end());
    allbytes.insert(allbytes.end(), spbytes.begin(), spbytes.end());
    allbytes.insert(allbytes.end(), qbytes.begin(), qbytes.end());

    REQUIRE(ecpairing(allbytes) != 0);

    pbytes = toBytes(P);
    sqbytes = toBytes(sQ);
    spbytes = toBytes(sP);
    qbytes = toBytes(negone * Q);

    allbytes.clear();

    allbytes.insert(allbytes.end(), pbytes.begin(), pbytes.end());
    allbytes.insert(allbytes.end(), sqbytes.begin(), sqbytes.end());
    allbytes.insert(allbytes.end(), spbytes.begin(), spbytes.end());
    allbytes.insert(allbytes.end(), qbytes.begin(), qbytes.end());

    REQUIRE(ecpairing(allbytes) != 0);
}
