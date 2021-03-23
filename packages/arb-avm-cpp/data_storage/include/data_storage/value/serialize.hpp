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

#ifndef serialize_hpp
#define serialize_hpp

#include <avm_values/value.hpp>

void serializeValue(const value& val, std::vector<unsigned char>& bytes);

struct ValueHasher {
    size_t operator()(value const& val) const noexcept {
        return (size_t)hash_value(val);
    }
};

using ValueCounter = std::unordered_map<value, uint32_t, ValueHasher>;

void getCodeSegmentDependencies(const CodeSegment& val,
                                ValueCounter& dependencies,
                                uint64_t start);
void getValueDependencies(const value& val, ValueCounter& dependencies);

#endif
