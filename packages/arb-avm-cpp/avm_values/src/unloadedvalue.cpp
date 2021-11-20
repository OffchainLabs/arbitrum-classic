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

#include <avm_values/unloadedvalue.hpp>

inline const HeapedUnloadedValueInfo& UnloadedValue::getHeaped() const {
    assert(isHeaped());
    return *impl.heaped_value.ptr;
}

UnloadedValue::UnloadedValue(BigUnloadedValue big)
    : impl{InlineUnloadedValue{}} {
    assert(big.value_size > 0);
    // Attempt to inline it
    if (big.type == ValueTypes::TUPLE) {
        if (big.value_size < (uint256_t(1) << 62) && big.value_size > 0) {
            impl.inline_value.value_size =
                uint64_t(big.value_size) | unloaded_value_fixed_bit;
            impl.inline_value.hash = big.hash;
            return;
        }
    }

    // We can't inline this; put it in a shared_ptr
    impl.heaped_value.uv_flag = unloaded_value_fixed_bit;
    impl.heaped_value.type = big.type;
    impl.heaped_value.ptr = std::make_shared<HeapedUnloadedValueInfo>(
        HeapedUnloadedValueInfo{big.hash, big.value_size});
}

UnloadedValue::~UnloadedValue() {
    if (isHeaped()) [[unlikely]] {
        impl.heaped_value.ptr.~shared_ptr();
    }
}

UnloadedValue::UnloadedValue(const UnloadedValue& other)
    : impl{InlineUnloadedValue{}} {
    if (other.isHeaped()) [[unlikely]] {
        impl.heaped_value = other.impl.heaped_value;
    } else {
        impl.inline_value = other.impl.inline_value;
    }
}

UnloadedValue& UnloadedValue::operator=(const UnloadedValue& other) {
    *this = other;
    return *this;
}

UnloadedValue::UnloadedValue(UnloadedValue&& other)
    : impl{InlineUnloadedValue{}} {
    if (other.isHeaped()) [[unlikely]] {
        impl.heaped_value = std::move(other.impl.heaped_value);
    } else {
        impl.inline_value = other.impl.inline_value;
    }
}

UnloadedValue& UnloadedValue::operator=(UnloadedValue&& other) {
    *this = std::move(other);
    return *this;
}

uint256_t UnloadedValue::hash() const {
    if (isHeaped()) [[unlikely]] {
        return getHeaped().hash;
    } else {
        return impl.inline_value.hash;
    }
}

uint256_t UnloadedValue::value_size() const {
    if (isHeaped()) [[unlikely]] {
        return getHeaped().value_size;
    } else {
        return impl.inline_value.value_size & ~unloaded_value_fixed_bit;
    }
}

ValueTypes UnloadedValue::type() const {
    if (isHeaped()) [[unlikely]] {
        return impl.heaped_value.type;
    } else {
        return ValueTypes::TUPLE;
    }
}
