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

#include <avm/bigint.hpp>
#include <avm/opcodes.hpp>

#include <nonstd/variant.hpp>

class bad_tuple_index : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "bad_tuple_index";
    }
};

enum types { NUM, CODEPT, TUPLE = 3 };

class TuplePool;
class Tuple;
struct Operation;
struct CodePoint;

// Note: uint256_t is actually 48 bytes long
using value = nonstd::variant<Tuple, uint256_t, CodePoint>;

std::ostream& operator<<(std::ostream& os, const value& val);

uint256_t hash(const value& value);

uint256_t deserialize_int(char*& srccode);
Operation deserializeOperation(char*& bufptr, TuplePool& pool);
CodePoint deserializeCodePoint(char*& bufptr, TuplePool& pool);
Tuple deserialize_tuple(char*& bufptr, int size, TuplePool& pool);
value deserialize_value(char*& srccode, TuplePool& pool);
void marshal_value(const value val, std::vector<unsigned char>& buf);
void marshal_Tuple(const Tuple& val, std::vector<unsigned char>& buf);
void marshal_CodePoint(const CodePoint& val, std::vector<unsigned char>& buf);
void marshal_uint256_t(const uint256_t& val, std::vector<unsigned char>& buf);

#endif /* value_hpp */
