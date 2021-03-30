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

#ifndef bigint_hpp
#define bigint_hpp

#include <intx/intx.hpp>

#include <vector>

using uint256_t = intx::uint256;
using uint512_t = intx::uint512;

inline int get_sign(uint256_t v) {
    return (v >> 255) ? -1 : 1;
}

uint256_t hash(const uint256_t& val);
uint256_t hash(const std::vector<unsigned char>& data);
uint256_t hash(const std::array<unsigned char, 32>& data);
uint256_t hash(const uint256_t& val, const uint256_t& val2);

void marshal_uint256_t(const uint256_t& val, std::vector<unsigned char>& buf);
void marshal_uint256_t(const uint256_t& val,
                       std::array<unsigned char, 32>& buf);

inline char* to_big_endian(const uint256_t& val, char* it) {
    intx::be::unsafe::store(reinterpret_cast<unsigned char*>(it), val);
    return it + 32;
}

inline unsigned char* to_big_endian(const uint256_t& val, unsigned char* it) {
    intx::be::unsafe::store(it, val);
    return it + 32;
}

#endif /* bigint_hpp */
