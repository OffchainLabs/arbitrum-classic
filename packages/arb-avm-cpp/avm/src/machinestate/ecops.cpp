//
//  ecops.cpp
//  algebra_bilinearity_test
//
//  Created by Harry Kalodner on 9/8/20.
//

#include <avm/machinestate/ecops.hpp>

#include <gmpxx.h>
#include <libff/algebra/curves/alt_bn128/alt_bn128_pp.hpp>
#include <libff/common/profiling.hpp>

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
G1<alt_bn128_pp> g1PfromBytes(const G1Point& point) {
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
G2<alt_bn128_pp> g2PfromBytes(const G2Point& point) {
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

    alt_bn128_Fq12 prod = alt_bn128_Fq12::one();

    for (uint8_t i = 0; i < numPairs; i++) {
        auto g1x =
            intx::be::unsafe::load<uint256_t>(input.data() + 192 * i + 32 * 0);
        auto g1y =
            intx::be::unsafe::load<uint256_t>(input.data() + 192 * i + 32 * 1);
        auto g2x0 =
            intx::be::unsafe::load<uint256_t>(input.data() + 192 * i + 32 * 2);
        auto g2x1 =
            intx::be::unsafe::load<uint256_t>(input.data() + 192 * i + 32 * 3);
        auto g2y0 =
            intx::be::unsafe::load<uint256_t>(input.data() + 192 * i + 32 * 4);
        auto g2y1 =
            intx::be::unsafe::load<uint256_t>(input.data() + 192 * i + 32 * 5);

        prod = prod *
               alt_bn128_pp::pairing(g1PfromBytes({g1x, g1y}),
                                     g2PfromBytes({g2x0, g2x1, g2y0, g2y1}));
    }

    const alt_bn128_GT result = alt_bn128_final_exponentiation(prod);
    return result;
}

int ecpairing(std::vector<uint8_t> input) {
    return ecpairing_internal(input) == GT<alt_bn128_pp>::one();
}
