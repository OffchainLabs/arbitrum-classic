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

#include <nonstd/variant.hpp>

#include <vector>

const mp_size_t BIG_INT_FOR_UINT256 = (256 + GMP_NUMB_BITS - 1) / GMP_NUMB_BITS;

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

nonstd::variant<libff::alt_bn128_G1, std::string> ecadd(
    const std::array<uint256_t, 4>& input);

nonstd::variant<libff::alt_bn128_G1, std::string> ecsmult(
    const std::array<uint256_t, 3>& input);

#endif /* ecops_hpp */
