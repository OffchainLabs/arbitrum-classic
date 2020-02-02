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

#ifndef value_hpp
#define value_hpp

#include <avm_values/bigint.hpp>

#include <nonstd/variant.hpp>

enum ValueTypes { NUM, CODEPT, HASH_ONLY, TUPLE };

class TuplePool;
class Tuple;
struct Operation;
struct CodePoint;

// Note: uint256_t is actually 48 bytes long
using value = nonstd::variant<Tuple, uint256_t, CodePoint>;

auto operator<<(std::ostream& os, const value& val) -> std::ostream&;
auto hash(const value& value) -> uint256_t;
auto get_tuple_size(const char*& bufptr) -> int;

auto deserializeUint256t(const char*& srccode) -> uint256_t;
auto deserializeOperation(const char*& bufptr, TuplePool& pool) -> Operation;
auto deserializeCodePoint(const char*& bufptr, TuplePool& pool) -> CodePoint;
auto deserializeTuple(const char*& bufptr, int size, TuplePool& pool) -> Tuple;
auto deserialize_value(const char*& srccode, TuplePool& pool) -> value;
void marshal_value(const value& val, std::vector<unsigned char>& buf);
void marshal_Tuple(const Tuple& val, std::vector<unsigned char>& buf);
void marshal_CodePoint(const CodePoint& val, std::vector<unsigned char>& buf);
void marshal_uint256_t(const uint256_t& val, std::vector<unsigned char>& buf);

void marshalShallow(const value& val, std::vector<unsigned char>& buf);
void marshalShallow(const Tuple& val, std::vector<unsigned char>& buf);
void marshalShallow(const CodePoint& val, std::vector<unsigned char>& buf);
void marshalShallow(const uint256_t& val, std::vector<unsigned char>& buf);

template <typename T>
static auto shrink(const uint256_t& i) -> T {
    return static_cast<T>(i & std::numeric_limits<T>::max());
}

auto GetHashKey(const value& val) -> std::vector<unsigned char>;

#endif /* value_hpp */
