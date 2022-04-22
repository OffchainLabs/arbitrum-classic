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

#include <avm_values/tuplestub.hpp>

#include <avm_values/valuetype.hpp>

#include <ethash/keccak.hpp>

#include <atomic>
#include <iostream>

uint256_t HashPreImage::hash() const {
    std::array<unsigned char, 65> tupData2;
    tupData2[0] = TUPLE;
    auto iter = tupData2.begin();
    iter++;

    iter = std::copy(firstHash.begin(), firstHash.end(), iter);
    to_big_endian(valueSize, iter);

    auto hash_val = ethash::keccak256(tupData2.data(), tupData2.size());
    return intx::be::load<uint256_t>(hash_val.bytes);
}

void memcpyAtomic(uint8_t* dest, const uint8_t* source, size_t length) {
    for (size_t i = 0; i < length; i++) {
        reinterpret_cast<std::atomic<uint8_t>*>(&dest[i])->store(
            source[i], std::memory_order_relaxed);
    }
}

void HashPreImage::writeAtomic(const HashPreImage& other) {
    memcpyAtomic(firstHash.begin(), other.firstHash.begin(), 32);
    memcpyAtomic(reinterpret_cast<uint8_t*>(&valueSize),
                 reinterpret_cast<const uint8_t*>(&other.valueSize), 32);
}

uint256_t HashPreImage::secretHash(
    const std::vector<unsigned char>& seed) const {
    std::vector<unsigned char> tupData2;
    tupData2.push_back(TUPLE);
    tupData2.insert(tupData2.end(), seed.begin(), seed.end());

    tupData2.insert(tupData2.end(), firstHash.begin(), firstHash.end());
    marshal_uint256_t(valueSize, tupData2);

    auto hash_val = ethash::keccak256(tupData2.data(), tupData2.size());
    return intx::be::load<uint256_t>(hash_val.bytes);
}

void HashPreImage::marshal(std::vector<unsigned char>& buf) const {
    buf.insert(buf.end(), firstHash.begin(), firstHash.end());
    marshal_uint256_t(valueSize, buf);
}

std::ostream& operator<<(std::ostream& os, const HashPreImage& val) {
    os << "HashPreImage(" << intx::to_string(val.hash()) << ")";
    return os;
}
