/*
 * Copyright 2019, Offchain Labs, Inc.
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

// Ported from
// https://github.com/microsoft/eEVM/blob/master/include/eEVM/bigint.h Copyright
// (c) Microsoft Corporation. All rights reserved. Licensed under the MIT
// License.

#ifndef bigint_hpp
#define bigint_hpp

#include <boost/multiprecision/cpp_int.hpp>

/* Boosts big integers behave (sort of) unexpected in the following sense.
numeric_limits<int256_t>::max() == numeric_limits<uint256_t>::max() == (1 <<
256) -1 I.e., the sign is stored in a separate bit.
*/

using uint128_t = boost::multiprecision::uint128_t;
using uint256_t = boost::multiprecision::uint256_t;
using uint512_t = boost::multiprecision::uint512_t;

using int128_t = boost::multiprecision::int128_t;
using int256_t = boost::multiprecision::int256_t;
using int512_t = boost::multiprecision::int512_t;

inline int get_sign(uint256_t v) {
    return (v >> 255) ? -1 : 1;
}

inline auto power(uint256_t b, uint64_t e) {
    return boost::multiprecision::pow(b, static_cast<unsigned int>(e));
}

uint256_t hash(const uint256_t& val);

inline bool bit(uint256_t x, int i) {
    return boost::multiprecision::bit_test(x, i);
}

#endif /* bigint_hpp */
