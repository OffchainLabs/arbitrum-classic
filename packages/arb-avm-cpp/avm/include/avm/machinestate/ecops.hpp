//
//  ecops.hpp
//  algebra_bilinearity_test
//
//  Created by Harry Kalodner on 9/8/20.
//

#ifndef ecops_hpp
#define ecops_hpp

#include <libff/algebra/curves/alt_bn128/alt_bn128_pp.hpp>
#include <libff/algebra/curves/public_params.hpp>

#include <vector>

libff::G1<libff::alt_bn128_pp> g1PfromBytes(std::vector<uint8_t> input);
libff::G2<libff::alt_bn128_pp> g2PfromBytes(std::vector<uint8_t> input);

libff::alt_bn128_GT ecpairing_internal(std::vector<uint8_t> input);

int ecpairing(std::vector<uint8_t> input);

#endif /* ecops_hpp */
