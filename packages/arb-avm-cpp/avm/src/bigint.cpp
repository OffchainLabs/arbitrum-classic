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

#include "avm/bigint.hpp"

#include "bigint_utils.hpp"
#include "util.hpp"

#include <sstream>

uint256_t hash(const uint256_t& val) {
    std::array<unsigned char, 32> intData;
    to_big_endian(val, intData.begin());

    std::array<unsigned char, 32> hashData;
    evm::Keccak_256(intData.data(), 32, hashData.data());
    return from_big_endian(hashData.begin(), hashData.end());
}
