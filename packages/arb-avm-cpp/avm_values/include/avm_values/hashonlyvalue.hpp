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

#ifndef hashonlyvalue_hpp
#define hashonlyvalue_hpp

#include <avm_values/value.hpp>

class HashOnly {
    uint256_t hash;
    int size;

   public:
    HashOnly() = default;
    HashOnly(const value& val);

    uint256_t getHash() const { return hash; }
    int getSize() const { return size; }
    void marshal(std::vector<unsigned char>& buf) const;
};

inline uint256_t hash(const HashOnly& hv) {
    return hv.getHash();
}

inline bool operator==(const HashOnly& val1, const HashOnly& val2) {
    return val1.getHash() == val2.getHash();
}

inline bool operator!=(const HashOnly& val1, const HashOnly& val2) {
    return val1.getHash() != val2.getHash();
}

std::ostream& operator<<(std::ostream& os, const HashOnly& val);

#endif /* hashonlyvalue_hpp */
