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

#ifndef valueloader_hpp
#define valueloader_hpp

#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

class AbstractValueLoader {
   public:
    virtual ~AbstractValueLoader() = default;

    // Throws an exception if the tuple cannot be loaded
    virtual value loadValue(const uint256_t& hash) = 0;

    virtual std::unique_ptr<AbstractValueLoader> clone() const = 0;
};

class ValueLoader : public AbstractValueLoader {
   protected:
    std::unique_ptr<AbstractValueLoader> impl;

   public:
    ValueLoader() : impl(nullptr) {}
    ValueLoader(std::unique_ptr<AbstractValueLoader> impl_)
        : impl(std::move(impl_)) {}
    ValueLoader(const ValueLoader& other) : impl(std::move(other.clone())) {}

    ValueLoader& operator=(const ValueLoader& other) {
        impl = std::move(other.clone());
        return *this;
    }

    value loadValue(const uint256_t& hash) override {
        if (!impl) {
            throw std::runtime_error("Value loader needed but not defined");
        }
        return impl->loadValue(hash);
    }

    std::unique_ptr<AbstractValueLoader> clone() const override {
        if (impl) {
            return impl->clone();
        } else {
            return nullptr;
        }
    }
};

#endif /* valueloader_hpp */
