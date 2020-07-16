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

void marshal_uint256_t(const uint256_t& val, std::vector<unsigned char>& buf) {
    buf.resize(buf.size() + 32);
    to_big_endian(val, &*(buf.end() - 32));
}
