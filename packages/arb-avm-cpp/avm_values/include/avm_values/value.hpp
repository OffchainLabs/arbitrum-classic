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
#include <avm_values/opcodes.hpp>

#include <nonstd/variant.hpp>

enum ValueTypes { NUM, CODEPT, HASH_PRE_IMAGE, TUPLE };

class TuplePool;
class Tuple;
struct Operation;
struct CodePoint;
class HashPreImage;
class Code;
struct CodePointStub;
struct CodePointRef;

// Note: uint256_t is actually 48 bytes long
using value = nonstd::variant<Tuple, uint256_t, CodePointStub, HashPreImage>;

std::ostream& operator<<(std::ostream& os, const value& val);
uint256_t hash_value(const value& value);

CodePointRef deserializeCodePointRef(const char*& bufptr);
CodePointStub deserializeCodePointStub(const char*& bufptr);
uint256_t deserializeUint256t(const char*& srccode);
value deserialize_value(const char*& srccode, TuplePool& pool);

void marshal_uint256_t(const uint256_t& val, std::vector<unsigned char>& buf);

void marshal_value(const value& val, std::vector<unsigned char>& buf);

void marshalForProof(const value& val,
                     MarshalLevel marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code);

uint256_t getSize(const value& val);

#endif /* value_hpp */
