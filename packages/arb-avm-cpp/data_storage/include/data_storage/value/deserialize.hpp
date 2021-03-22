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

#ifndef deserialize_hpp
#define deserialize_hpp

#include <avm_values/codepoint.hpp>
#include <avm_values/value.hpp>

using SlotPointer = std::variant<value*, Buffer*, CodeSegment*>;

struct Slot {
    SlotPointer ptr;
    uint256_t hash;

    Slot(SlotPointer ptr_, uint256_t hash_) : ptr(ptr_), hash(hash_) {}
};

void deserializeValue(std::vector<unsigned char>::const_iterator& bytes,
                      value* result,
                      std::vector<Slot>& slots);

// Special cases with explicit types for other SlotPointer types
void deserializeValue(std::vector<unsigned char>::const_iterator& bytes,
                      Buffer* result,
                      std::vector<Slot>& slots);
void deserializeValue(std::vector<unsigned char>::const_iterator& bytes,
                      CodeSegment* result,
                      std::vector<Slot>& slots);

#endif
