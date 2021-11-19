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

#include <cassert>

struct BigUnloadedValue {
    ValueTypes type;
    uint256_t hash{};
    uint256_t value_size{};
};

struct InlineUnloadedTuple {
    uint256_t hash{};
    uint64_t value_size{};
};

struct InlineUnloadedNonTuple {
    ValueTypes type;
    uint256_t hash{};
};

using unloadedValue = std::variant<InlineUnloadedTuple,
                                   InlineUnloadedNonTuple,
                                   std::shared_ptr<BigUnloadedValue>>;

struct UnloadedValueUnpacker {
    BigUnloadedValue operator()(const InlineUnloadedTuple& val) {
        return BigUnloadedValue{ValueTypes::TUPLE, val.hash, val.value_size};
    }
    BigUnloadedValue operator()(const InlineUnloadedNonTuple& val) {
        return BigUnloadedValue{val.type, val.hash, 1};
    }
    BigUnloadedValue operator()(const std::shared_ptr<BigUnloadedValue>& val) {
        return *val;
    }
};

class UnloadedValue {
   private:
    unloadedValue inner;

    UnloadedValue(unloadedValue inner_) : inner(inner_) {}

   public:
    UnloadedValue(ValueTypes ty, uint256_t hash, uint256_t size) {
        assert(size > 0);
        if (ty == ValueTypes::TUPLE) {
            uint64_t small_size(size);
            if (uint256_t(small_size) == size) {
                inner = InlineUnloadedTuple{hash, small_size};
            } else {
                inner = std::make_shared<BigUnloadedValue>(
                    BigUnloadedValue{ty, hash, size});
            }
        } else {
            assert(size == 1);
            inner = InlineUnloadedNonTuple{ty, hash};
        }
    }

    BigUnloadedValue unpack() const {
        return std::visit(UnloadedValueUnpacker{}, inner);
    }
};

inline uint256_t hash(const UnloadedValue& uv) {
    return uv.unpack().hash;
}

#endif /* unloadedvalue_hpp */
