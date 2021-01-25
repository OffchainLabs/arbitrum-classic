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

#include <avm_values/bigint.hpp>

#include <ethash/keccak.hpp>

#include <sstream>
#include <vector>

uint256_t hash(const uint256_t& val) {
    auto int_data = intx::be::store<ethash::hash256>(val);
    auto hash_val = ethash::keccak256(int_data);
    return intx::be::load<uint256_t>(hash_val);
}

uint256_t hash(const std::vector<unsigned char>& data) {
    auto hash_val = ethash::keccak256(data.data(), data.size());
    return intx::be::load<uint256_t>(hash_val);
}

uint256_t hash(const uint256_t& val1, const uint256_t& val2) {
    unsigned char data1[32];
    unsigned char data2[32];
    unsigned char data[64];
    intx::be::store<256>(data1, val1);
    intx::be::store<256>(data2, val2);
    memcpy(data, data1, 32);
    memcpy(data + 32, data2, 32);
    auto hash_val = ethash::keccak256(data, 64);
    return intx::be::load<uint256_t>(hash_val);
}

void marshal_uint256_t(const uint256_t& val, std::vector<unsigned char>& buf) {
    buf.resize(buf.size() + 32);
    to_big_endian(val, &*(buf.end() - 32));
}
