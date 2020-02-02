/*
 * Copyright 2019, Offchain Labs, Inc.
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

#ifndef exceptions_hpp
#define exceptions_hpp

#include <nonstd/variant.hpp>

class bad_tuple_index : public std::exception {
   public:
    auto what() const noexcept -> const char* override {
        return "bad_tuple_index";
    }
};

class bad_pop_type : public std::exception {
   public:
    auto what() const noexcept -> const char* override {
        return "bad_variant_access";
    }
};

class int_out_of_bounds : public std::exception {
   public:
    auto what() const noexcept -> const char* override {
        return "int_out_of_bounds";
    }
};

#endif /* exceptions_hpp */
