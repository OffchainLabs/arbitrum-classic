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

#include <bigint_utils.hpp>

#include <ethash/keccak.hpp>

uint256_t HashPreImage::hash() const {
    std::array<unsigned char, 65> tupData2;
    tupData2[0] = TUPLE;
    auto iter = tupData2.begin();
    iter++;

    iter = std::copy(firstHash.begin(), firstHash.end(), iter);
    to_big_endian(valueSize, iter);

    auto hash_val = ethash::keccak256(tupData2.data(), tupData2.size());
    return from_big_endian(&hash_val.bytes[0], &hash_val.bytes[32]);
}

void HashPreImage::marshal(std::vector<unsigned char>& buf) const {
    buf.insert(buf.end(), firstHash.begin(), firstHash.end());

    std::array<unsigned char, 32> tmpbuf;
    to_big_endian(valueSize, tmpbuf.begin());
    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

std::ostream& operator<<(std::ostream& os, const HashPreImage& val) {
    os << "HashPreImage(" << val.hash() << ")";
    return os;
}
