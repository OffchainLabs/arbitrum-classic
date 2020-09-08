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

#include <nonstd/variant.hpp>

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

nonstd::variant<libff::G1<libff::alt_bn128_pp>, std::string> g1PfromBytes(
    const G1Point& point);
nonstd::variant<libff::G2<libff::alt_bn128_pp>, std::string> g2PfromBytes(
    const G2Point& point);

nonstd::variant<libff::alt_bn128_GT, std::string> ecpairing_internal(
    const std::vector<std::array<uint256_t, 6>>& input);

nonstd::variant<bool, std::string> ecpairing(
    const std::vector<std::array<uint256_t, 6>>& input);

#endif /* ecops_hpp */
