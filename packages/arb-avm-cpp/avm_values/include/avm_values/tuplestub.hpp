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

#ifndef tuplestub_hpp
#define tuplestub_hpp

#include <avm_values/bigint.hpp>

#include <array>
#include <memory>
#include <vector>

class HashPreImage {
   private:
    std::array<unsigned char, 32> firstHash;
    uint256_t valueSize;

   public:
    HashPreImage(std::array<unsigned char, 32> _firstHash,
                 uint256_t _valueSize) {
        firstHash = _firstHash;
        valueSize = _valueSize;
    }
    std::array<unsigned char, 32> getFirstHash() const { return firstHash; }
    uint256_t getSize() const { return valueSize; }
    void marshal(std::vector<unsigned char>& buf) const;
    uint256_t hash() const;
    uint256_t secretHash(const std::vector<unsigned char>& seed) const;
};

inline uint256_t hash(const HashPreImage& hv) {
    return hv.hash();
}

inline uint256_t hash(const std::shared_ptr<HashPreImage>& hv) {
    return hv->hash();
}

inline bool operator==(const HashPreImage& val1, const HashPreImage& val2) {
    return val1.hash() == val2.hash();
}

inline bool operator!=(const HashPreImage& val1, const HashPreImage& val2) {
    return val1.hash() != val2.hash();
}

std::ostream& operator<<(std::ostream& os, const HashPreImage& val);

#endif /* tuplestub_hpp */
