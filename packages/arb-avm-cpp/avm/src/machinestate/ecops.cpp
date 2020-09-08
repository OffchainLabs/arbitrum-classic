//
//  ecops.cpp
//  algebra_bilinearity_test
//
//  Created by Harry Kalodner on 9/8/20.
//

#include <avm/machinestate/ecops.hpp>

#include <gmp.h>
#include <libff/algebra/curves/alt_bn128/alt_bn128_pp.hpp>

using namespace libff;

// assumes bytes are big endian
// also assumes 64 bytes, 0..31 for X and 32...64 for Y, representing a curve
// point using affine coordinates if either X or Y is less than 32 bytes, they
// are assumed to be padded with leading 0s
G1<alt_bn128_pp> g1PfromBytes(std::vector<uint8_t> input) {
    if (input.size() != 64) {
        throw -1;  // change to throw AVM exception
    }
    uint8_t* xbytes = &input[0];
    uint8_t* ybytes = &input[32];

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
        throw -2;  // change to throw AVM exception
    }

    if (mpz_cmp(mpzy, modulus) >= 0) {
        throw -3;  // change to throw AVM exception
    }

    bigint<4L> xbi(mpzx);
    bigint<4L> ybi(mpzy);
    alt_bn128_Fq X(mpzx);
    alt_bn128_Fq Y(mpzy);

    // assumes affine coordinates
    G1<alt_bn128_pp> P = alt_bn128_G1(X, Y, alt_bn128_Fq::one());

    mpz_clears(mpzx, mpzy, modulus, NULL);
    if (!P.is_well_formed()) {
        throw -1;  // change to throw AVM exception
    }
    return P;
}

// assumes bytes are big endian
// also assumes 128 bytes representing a curve point using affine coordinates
// if either X or Y is less than 64 bytes, they are assumed to be padded with
// leading 0s
G2<alt_bn128_pp> g2PfromBytes(std::vector<uint8_t> input) {
    if (input.size() != 128) {
        throw -1;  // change to throw AVM exception
    }
    uint8_t* xc0bytes = &input[0];
    uint8_t* xc1bytes = &input[32];
    uint8_t* yc0bytes = &input[64];
    uint8_t* yc1bytes = &input[96];

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
        throw -2;  // change to throw AVM exception
    }

    if (mpz_cmp(mpzxc1, modulus) >= 0) {
        throw -3;  // change to throw AVM exception
    }

    if (mpz_cmp(mpzyc0, modulus) >= 0) {
        throw -4;  // change to throw AVM exception
    }

    if (mpz_cmp(mpzyc1, modulus) >= 0) {
        throw -5;  // change to throw AVM exception
    }

    bigint<4L> xc0bi(mpzxc0);
    bigint<4L> xc1bi(mpzxc1);
    bigint<4L> yc0bi(mpzyc0);
    bigint<4L> yc1bi(mpzyc1);

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
        throw -1;  // change to throw AVM exception
    }
    return P;
}

alt_bn128_GT ecpairing_internal(std::vector<uint8_t> input) {
    if (input.size() % 192 != 0) {
        throw -1;  // change to throw AVM exception
    }

    uint64_t numPairs =
        input.size() / 192;  // TODO: can you give so many pairs to overflow?
    if (numPairs > 20) {
        throw -2;  // change to throw AVM exception
    }

    alt_bn128_pp::init_public_params();

    alt_bn128_Fq12 prod = alt_bn128_Fq12::one();

    std::vector<uint8_t>::const_iterator first;
    std::vector<uint8_t>::const_iterator last;
    for (uint8_t i = 0; i < numPairs; i++) {
        first = input.begin() + 192 * i;
        last = input.begin() + 192 * i + 64;
        std::vector<uint8_t> g1Vec(first, last);

        first = input.begin() + 192 * i + 64;
        last = input.begin() + 192 * i + 192;
        std::vector<uint8_t> g2Vec(first, last);

        prod = prod *
               alt_bn128_pp::pairing(g1PfromBytes(g1Vec), g2PfromBytes(g2Vec));
    }

    const alt_bn128_GT result = alt_bn128_final_exponentiation(prod);
    return result;
}

int ecpairing(std::vector<uint8_t> input) {
    return (ecpairing_internal(input) == GT<alt_bn128_pp>::one());
}

// used only for testing
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

// used only for testing
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

// used only for testing
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

void testg1PfromBytes() {
    alt_bn128_pp::init_public_params();

    G1<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G1<alt_bn128_pp>::one();

    G1<alt_bn128_pp> outP = g1PfromBytes(toBytes(P));

    assert(outP == P);
}

void testg2PfromBytes() {
    alt_bn128_pp::init_public_params();

    G2<alt_bn128_pp> P =
        (Fr<alt_bn128_pp>::random_element()) * G2<alt_bn128_pp>::one();

    G2<alt_bn128_pp> outP = g2PfromBytes(toBytes(P));

    assert(outP == P);
}

void testecpairing_internal() {
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
    assert(prod == ecpairing_internal(allbytes));
}

void testecpairing() {
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

    assert(ecpairing(allbytes) != 0);

    pbytes = toBytes(P);
    sqbytes = toBytes(negone * sQ);
    spbytes = toBytes(sP);
    qbytes = toBytes(Q);

    allbytes.clear();

    allbytes.insert(allbytes.end(), pbytes.begin(), pbytes.end());
    allbytes.insert(allbytes.end(), sqbytes.begin(), sqbytes.end());
    allbytes.insert(allbytes.end(), spbytes.begin(), spbytes.end());
    allbytes.insert(allbytes.end(), qbytes.begin(), qbytes.end());

    assert(ecpairing(allbytes) != 0);

    pbytes = toBytes(P);
    sqbytes = toBytes(sQ);
    spbytes = toBytes(negone * sP);
    qbytes = toBytes(Q);

    allbytes.clear();

    allbytes.insert(allbytes.end(), pbytes.begin(), pbytes.end());
    allbytes.insert(allbytes.end(), sqbytes.begin(), sqbytes.end());
    allbytes.insert(allbytes.end(), spbytes.begin(), spbytes.end());
    allbytes.insert(allbytes.end(), qbytes.begin(), qbytes.end());

    assert(ecpairing(allbytes) != 0);

    pbytes = toBytes(P);
    sqbytes = toBytes(sQ);
    spbytes = toBytes(sP);
    qbytes = toBytes(negone * Q);

    allbytes.clear();

    allbytes.insert(allbytes.end(), pbytes.begin(), pbytes.end());
    allbytes.insert(allbytes.end(), sqbytes.begin(), sqbytes.end());
    allbytes.insert(allbytes.end(), spbytes.begin(), spbytes.end());
    allbytes.insert(allbytes.end(), qbytes.begin(), qbytes.end());

    assert(ecpairing(allbytes) != 0);
}
