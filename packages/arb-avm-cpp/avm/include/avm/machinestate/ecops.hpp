//
//  ecops.hpp
//  algebra_bilinearity_test
//
//  Created by Harry Kalodner on 9/8/20.
//

#ifndef ecops_hpp
#define ecops_hpp

#include <avm_values/bigint.hpp>

#include <libff/algebra/curves/alt_bn128/alt_bn128_pp.hpp>
#include <libff/algebra/curves/public_params.hpp>

#include <vector>

struct G1Point {
    uint256_t x;
    uint256_t y;
};

struct G2Point {
    uint256_t x0;
    uint256_t x1;
    uint256_t y0;
    uint256_t y1;
};

libff::G1<libff::alt_bn128_pp> g1PfromBytes(const G1Point& point);
libff::G2<libff::alt_bn128_pp> g2PfromBytes(const G2Point& point);

libff::alt_bn128_GT ecpairing_internal(std::vector<uint8_t> input);

int ecpairing(std::vector<uint8_t> input);

#endif /* ecops_hpp */
