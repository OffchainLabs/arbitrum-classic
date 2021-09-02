/*
 * Copyright 2021, Offchain Labs, Inc.
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

#ifndef unloadedvalue_hpp
#define unloadedvalue_hpp

#include <avm_values/valuetype.hpp>

struct UnloadedValue {
    ValueTypes type;
    uint256_t hash;
    uint256_t value_size;
};

inline bool operator==(const UnloadedValue& val1, const UnloadedValue& val2) {
    return val1.hash == val2.hash;
}

inline bool operator!=(const UnloadedValue& val1, const UnloadedValue& val2) {
    return val1.hash != val2.hash;
}

inline uint256_t hash(const UnloadedValue& uv) {
    return uv.hash;
}

#endif /* unloadedvalue_hpp */
