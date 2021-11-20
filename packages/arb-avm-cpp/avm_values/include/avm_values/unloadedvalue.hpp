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

#include <avm_values/bigint.hpp>
#include <avm_values/valuetype.hpp>
#include <memory>

#include <cassert>

constexpr uint64_t unloaded_value_fixed_bit = uint64_t(1) << 62;

struct BigUnloadedValue {
    ValueTypes type;
    uint256_t hash{};
    uint256_t value_size{};
};

struct InlineUnloadedValue {
    uint64_t value_size;
    uint256_t hash;
};

struct HeapedUnloadedValueInfo {
    uint256_t hash;
    uint256_t value_size;
};

struct HeapedUnloadedValue {
    uint64_t uv_flag{};
    ValueTypes type;
    std::shared_ptr<HeapedUnloadedValueInfo> ptr;
};

class UnloadedValue {
   private:
    union UnloadedValueImpl {
        InlineUnloadedValue inline_value;
        HeapedUnloadedValue heaped_value;
        ~UnloadedValueImpl() {}
    };

    UnloadedValueImpl impl;

    inline const HeapedUnloadedValueInfo& getHeaped() const;

   public:
    UnloadedValue(BigUnloadedValue);

    // Rule of 5: provide a destructor, copy-constructor, copy-assignment
    // operator, move constructor, and move-assignment operator.
    // These all forward to the shared_ptr impls if necessary.
    ~UnloadedValue();
    UnloadedValue(const UnloadedValue&);
    UnloadedValue& operator=(const UnloadedValue&);
    UnloadedValue(UnloadedValue&&);
    UnloadedValue& operator=(UnloadedValue&&);

    inline bool isHeaped() const {
        return __builtin_expect(
            impl.heaped_value.uv_flag == unloaded_value_fixed_bit, 0);
    }

    // Provide methods to access fields
    uint256_t hash() const;
    uint256_t value_size() const;
    ValueTypes type() const;
};

inline uint256_t hash(const UnloadedValue& uv) {
    return uv.hash();
}

#endif /* unloadedvalue_hpp */
