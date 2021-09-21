/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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
#include <avm_values/buffer.hpp>
#include <avm_values/opcodes.hpp>
#include <avm_values/unloadedvalue.hpp>
#include <avm_values/valuetype.hpp>

#include <variant>

class Tuple;
struct Operation;
struct CodePoint;
class HashPreImage;
class Code;
struct CodePointStub;
struct CodePointRef;
struct UnloadedValue;

using value = std::variant<Tuple,
                           uint256_t,
                           CodePointStub,
                           std::shared_ptr<HashPreImage>,
                           Buffer,
                           UnloadedValue>;

struct TuplePlaceholder {
    uint8_t values;
};
using DeserializedValue = std::variant<TuplePlaceholder, value>;

std::ostream& operator<<(std::ostream& os, const value& val);
uint256_t hash_value(const value& value);
bool values_equal(const value& a, const value& b);

uint64_t deserialize_uint64_t(const char*& bufptr);
CodePointRef deserializeCodePointRef(const char*& bufptr);
CodePointStub deserializeCodePointStub(const char*& bufptr);
uint256_t deserializeUint256t(const char*& srccode);
value deserialize_value(const char*& srccode);

void marshal_uint64_t(uint64_t val, std::vector<unsigned char>& buf);

void marshal_value(const value& val, std::vector<unsigned char>& buf);

void marshalForProof(const value& val,
                     size_t marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code);

uint256_t getSize(const value& val);

value assembleValueFromDeserialized(std::vector<DeserializedValue> values);

#endif /* value_hpp */
