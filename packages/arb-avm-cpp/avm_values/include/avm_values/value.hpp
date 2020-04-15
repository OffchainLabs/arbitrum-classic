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

std::ostream& operator<<(std::ostream& os, const value& val);
uint256_t hash(const value& value);
int get_tuple_size(char*& bufptr);

uint256_t deserializeUint256t(const char*& srccode);
Operation deserializeOperation(const char*& bufptr, TuplePool& pool);
CodePoint deserializeCodePoint(const char*& bufptr, TuplePool& pool);
Tuple deserializeTuple(const char*& bufptr, int size, TuplePool& pool);
value deserialize_value(const char*& srccode, TuplePool& pool);
void marshal_value(const value& val, std::vector<unsigned char>& buf);
void marshal_Tuple(const Tuple& val, std::vector<unsigned char>& buf);
void marshal_CodePoint(const CodePoint& val, std::vector<unsigned char>& buf);
void marshal_uint256_t(const uint256_t& val, std::vector<unsigned char>& buf);

void marshalShallow(const value& val, std::vector<unsigned char>& buf);
void marshalShallow(const Tuple& val, std::vector<unsigned char>& buf);
void marshalShallow(const CodePoint& val, std::vector<unsigned char>& buf);
void marshalShallow(const uint256_t& val, std::vector<unsigned char>& buf);

template <typename T>
static T shrink(uint256_t i) {
    return static_cast<T>(i & std::numeric_limits<T>::max());
}

std::vector<unsigned char> GetHashKey(const value& val);
int getSize(const value& val);

class HashOnly {
    uint256_t hash;
    uint256_t size;

   public:
    HashOnly(uint256_t _hash, int _size) {
        hash = _hash;
        size = _size;
    }
    void ToBuff(std::vector<unsigned char>& buf);
};

#endif /* value_hpp */
