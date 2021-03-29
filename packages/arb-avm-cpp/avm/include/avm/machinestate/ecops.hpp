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

#ifndef ecops_hpp
#define ecops_hpp

#include <avm_values/bigint.hpp>

#include <libff/algebra/curves/alt_bn128/alt_bn128_pp.hpp>
#include <libff/algebra/curves/public_params.hpp>

#include <variant>
#include <vector>

using namespace libff;

const mp_size_t BIG_INT_FOR_UINT256 = (256 + GMP_NUMB_BITS - 1) / GMP_NUMB_BITS;

void initEcOps();

struct G1Point {
    uint256_t x;
    uint256_t y;
};

std::ostream& operator<<(std::ostream& os, const G1Point& val);

struct G2Point {
    uint256_t x0;
    uint256_t x1;
    uint256_t y0;
    uint256_t y1;
};

std::ostream& operator<<(std::ostream& os, const G2Point& val);

void mpz_export_and_pad32(uint8_t* output, mpz_t input);
G1Point toG1ArbPoint(G1<alt_bn128_pp> P);
// Probably unsafe as EIP-196 doesn't define an encoding for this.
// In particular, the zero point is encoded as non-zero bytes.
G2Point toG2ArbPoint(G2<alt_bn128_pp> P);

std::variant<libff::G1<libff::alt_bn128_pp>, std::string> g1PfromBytes(
    const G1Point& point);
std::variant<libff::G2<libff::alt_bn128_pp>, std::string> g2PfromBytes(
    const G2Point& point);

std::variant<alt_bn128_GT, std::string> ecpairing_internal(
    const std::vector<std::pair<G1Point, G2Point>>& input);

std::variant<bool, std::string> ecpairing(
    const std::vector<std::pair<G1Point, G2Point>>& input);

std::variant<G1Point, std::string> ecadd(const G1Point& input_a,
                                         const G1Point& input_b);

std::variant<G1Point, std::string> ecmul(const G1Point& point,
                                         const uint256_t& factor);

#endif /* ecops_hpp */
