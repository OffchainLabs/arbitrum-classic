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

class Slot {
    friend SlotMap;

   private:
    std::variant<Tuple, std::shared_ptr<Buffer>, CodeSegment> inner;

    Slot(std::variant<Tuple, std::shared_ptr<Buffer>, CodeSegment> inner_)
        : inner(inner_) {}

    static void fillInner(Tuple inner, value val);
    static void fillInner(std::shared_ptr<Buffer> inner, value val);
    static void fillInner(CodeSegment inner, value val);

   public:
    void fill(value);
};

class SlotMap {
   private:
    struct HashHasher {
        std::size_t operator()(const uint256_t& hash) const noexcept {
            return intx::narrow_cast<std::size_t>(hash);
        }
    };

    std::unordered_map<uint256_t, Slot, HashHasher> slots;

   public:
    SlotMap() = default;

    Tuple tupleSlot(uint256_t hash);
    std::shared_ptr<Buffer> bufferSlot(uint256_t hash);
    CodeSegment codeSegmentSlot(uint256_t hash);

    bool empty();
    std::pair<uint256_t, Slot> takeSlot();
};

// Deserialize a value from bytes, returning a list of "slots" that need filled
// in. Note that while the value will have pointers to the slots, it may not
// directly contain the slots (i.e. the slot pointer may not point to an offset
// of the value).
value deserializeValue(std::vector<unsigned char>::const_iterator& bytes,
                       SlotMap& slots);

#endif
