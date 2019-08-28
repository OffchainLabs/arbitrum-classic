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

#ifndef bigint_utils_hpp
#define bigint_utils_hpp

#include <boost/endian/conversion.hpp>
#include <boost/multiprecision/cpp_int.hpp>

#include <limits>
#include <sstream>
#include <string>

using uint256_t = boost::multiprecision::uint256_t;

inline auto from_hex_str(const std::string& s) {
    std::stringstream ss;
    ss << std::hex << s;
    uint256_t v;
    ss >> v;
    return v;
}

inline auto to_hex_str(const uint256_t& v) {
    std::stringstream ss;
    ss << "0x" << std::hex << v;
    return ss.str();
}

template <typename Iterator>
auto from_big_endian(const Iterator begin, const Iterator end) {
    uint256_t v;
    // imports in big endian by default
    boost::multiprecision::import_bits(
        v, begin, end, std::numeric_limits<uint8_t>::digits, true);
    return v;
}

template <typename Iterator>
inline void to_big_endian(uint256_t v, Iterator out) {
    // boost::multiprecision::export_bits() does not work here, because it
    // doesn't support fixed width export.
    uint64_t* o = reinterpret_cast<uint64_t*>(&*out);
    constexpr uint64_t mask64 = 0xffffffff'ffffffff;

    for (size_t i = 4; i-- > 0;) {
        uint64_t n = static_cast<uint64_t>(v & mask64);
        v >>= 64;
        o[i] = boost::endian::native_to_big(n);
    }
}

#endif /* bigint_utils_hpp */
