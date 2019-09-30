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

#include "avm/value.hpp"

#include "avm/codepoint.hpp"
#include "avm/pool.hpp"
#include "avm/tuple.hpp"

#include "bigint_utils.hpp"
#include "util.hpp"

#include <ostream>

#define UINT256_SIZE 32
#define UINT64_SIZE 8

uint256_t deserialize_int(char*& bufptr) {
    uint256_t ret = from_big_endian(bufptr, bufptr + UINT256_SIZE);
    bufptr += UINT256_SIZE;
    return ret;
}

// make sure correct
uint64_t deserialize_int64(char*& bufptr) {
    uint64_t ret_value;
    memcpy(&ret_value, bufptr, UINT64_SIZE);
    return ret_value;
}

Operation deserializeOperation(char*& bufptr, TuplePool& pool) {
    uint8_t immediateCount;
    memcpy(&immediateCount, bufptr, sizeof(immediateCount));
    bufptr += sizeof(immediateCount);
    OpCode opcode;
    memcpy(&opcode, bufptr, sizeof(opcode));
    bufptr += sizeof(opcode);

    if (immediateCount == 1) {
        return {opcode, deserialize_value(bufptr, pool)};
    } else {
        return {opcode};
    }
}

CodePoint deserializeCodePoint(char*& bufptr, TuplePool& pool) {
    CodePoint ret;
    memcpy(&ret.pc, bufptr, sizeof(ret.pc));
    bufptr += sizeof(ret.pc);
    ret.pc = boost::endian::big_to_native(ret.pc);
    ret.op = deserializeOperation(bufptr, pool);
    ret.nextHash = from_big_endian(bufptr, bufptr + UINT256_SIZE);
    bufptr += UINT256_SIZE;

    return ret;
}

Tuple deserialize_tuple(char*& bufptr, int size, TuplePool& pool) {
    Tuple tup(&pool, size);
    for (int i = 0; i < size; i++) {
        tup.set_element(i, deserialize_value(bufptr, pool));
    }
    return tup;
}

void marshal_Tuple(const Tuple& val, std::vector<unsigned char>& buf) {
    val.marshal(buf);
}

void marshal_CodePoint(const CodePoint& val, std::vector<unsigned char>& buf) {
    val.marshal(buf);
}

void marshal_uint256_t(const uint256_t& val, std::vector<unsigned char>& buf) {
    buf.push_back(NUM);
    std::array<unsigned char, 32> tmpbuf;
    to_big_endian(val, tmpbuf.begin());
    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

// make sure correct
void marshal_uint64_t(const uint64_t& val, std::vector<unsigned char>& buf) {
    auto big_endian_val = boost::endian::native_to_big(val);
    std::array<unsigned char, 8> tmpbuf;
    memcpy(tmpbuf.data(), &big_endian_val, sizeof(big_endian_val));

    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

void marshal_value(const value& val, std::vector<unsigned char>& buf) {
    if (nonstd::holds_alternative<Tuple>(val))
        marshal_Tuple(nonstd::get<Tuple>(val), buf);
    else if (nonstd::holds_alternative<uint256_t>(val))
        marshal_uint256_t(nonstd::get<uint256_t>(val), buf);
    else if (nonstd::holds_alternative<CodePoint>(val))
        marshal_CodePoint(nonstd::get<CodePoint>(val), buf);
}

void marshalShallow(const value& val, std::vector<unsigned char>& buf) {
    return nonstd::visit([&](const auto& v) { return marshalShallow(v, buf); },
                         val);
}

void marshalShallow(const Tuple& val, std::vector<unsigned char>& buf) {
    buf.push_back(TUPLE + val.tuple_size());
    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        buf.push_back(HASH_ONLY);
        std::array<unsigned char, 32> tmpbuf;
        to_big_endian(::hash(val.get_element(i)), tmpbuf.begin());
        buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
    }
}

void marshalShallow(const CodePoint& val, std::vector<unsigned char>& buf) {
    buf.push_back(CODEPT);
    val.op.marshalShallow(buf);
    std::array<unsigned char, 32> hashVal;
    to_big_endian(val.nextHash, hashVal.begin());
    buf.insert(buf.end(), hashVal.begin(), hashVal.end());
}

void marshalShallow(const uint256_t& val, std::vector<unsigned char>& buf) {
    marshal_uint256_t(val, buf);
}

value deserialize_value(char*& bufptr, TuplePool& pool) {
    uint8_t valType;
    memcpy(&valType, bufptr, sizeof(valType));
    bufptr += sizeof(valType);
    switch (valType) {
        case NUM:
            return deserialize_int(bufptr);
        case CODEPT:
            return deserializeCodePoint(bufptr, pool);
        default:
            if (valType >= TUPLE && valType <= TUPLE + 8) {
                return deserialize_tuple(bufptr, valType - TUPLE, pool);
            } else {
                std::printf("in deserialize_value, unhandled type = %X\n",
                            valType);
                throw std::runtime_error("Tried to deserialize unhandled type");
            }
    }
}

int get_tuple_size(char*& bufptr) {
    uint8_t valType;
    memcpy(&valType, bufptr, sizeof(valType));

    return valType - TUPLE;
}

uint256_t hash(const value& value) {
    return nonstd::visit([](const auto& val) { return hash(val); }, value);
}

uint256_t& assumeInt(value& val) {
    auto aNum = nonstd::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

const uint256_t& assumeInt(const value& val) {
    auto aNum = nonstd::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

uint64_t assumeInt64(uint256_t& val) {
    if (val > std::numeric_limits<uint64_t>::max())
        throw int_out_of_bounds{};

    return static_cast<uint64_t>(val);
}

Tuple& assumeTuple(value& val) {
    auto tup = nonstd::get_if<Tuple>(&val);
    if (!tup) {
        throw bad_pop_type{};
    }
    return *tup;
}

struct ValuePrinter {
    std::ostream& os;

    std::ostream* operator()(const Tuple& val) const {
        os << val;
        return &os;
    }

    std::ostream* operator()(const uint256_t& val) const {
        os << val;
        return &os;
    }

    std::ostream* operator()(const CodePoint& val) const {
        //        std::printf("in CodePoint ostream operator\n");
        os << "CodePoint(" << val.pc << ", " << val.op << ", "
           << to_hex_str(val.nextHash) << ")";
        return &os;
    }
};

std::ostream& operator<<(std::ostream& os, const value& val) {
    return *nonstd::visit(ValuePrinter{os}, val);
}

std::string ToHashString(uint256_t hash_key) {
    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash_key, hash_key_vector);

    return std::string(hash_key_vector.begin(), hash_key_vector.end());
}

std::vector<unsigned char> GetHashKey(const value& val) {
    auto hash_key = hash(val);
    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash_key, hash_key_vector);

    return hash_key_vector;
}

std::vector<unsigned char> serializeForCheckpoint(const uint256_t& val) {
    std::vector<unsigned char> value_vector;
    auto type_code = (unsigned char)NUM;
    value_vector.push_back(type_code);

    std::vector<unsigned char> num_vector;
    marshal_uint256_t(val, num_vector);

    value_vector.insert(value_vector.end(), num_vector.begin(),
                        num_vector.end());

    return value_vector;
}

struct Serializer {
    SerializedValue operator()(const Tuple& val) const {
        std::vector<unsigned char> value_vector;

        auto type_code = (unsigned char)TUPLE;
        value_vector.push_back(type_code);

        auto hash_key = hash(val);
        std::vector<unsigned char> hash_key_vector;
        marshal_uint256_t(hash_key, hash_key_vector);

        value_vector.insert(value_vector.end(), hash_key_vector.begin(),
                            hash_key_vector.end());

        std::string str_value(value_vector.begin(), value_vector.end());

        SerializedValue serialized_value{TUPLE, str_value};

        return serialized_value;
    }

    SerializedValue operator()(const uint256_t& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = (unsigned char)NUM;
        value_vector.push_back(type_code);

        std::vector<unsigned char> num_vector;
        marshal_uint256_t(val, num_vector);

        value_vector.insert(value_vector.end(), num_vector.begin(),
                            num_vector.end());
        std::string str_value(value_vector.begin(), value_vector.end());

        SerializedValue serialized_value{NUM, str_value};

        return serialized_value;
    }

    SerializedValue operator()(const CodePoint& val) const {
        std::vector<unsigned char> value_vector;
        auto type_code = (unsigned char)CODEPT;
        value_vector.push_back(type_code);

        std::vector<unsigned char> pc_vector;
        marshal_uint64_t(val.pc, pc_vector);

        value_vector.insert(value_vector.end(), pc_vector.begin(),
                            pc_vector.end());

        std::string str_value(value_vector.begin(), value_vector.end());

        SerializedValue serialized_value{CODEPT, str_value};

        return serialized_value;
    }
};

SerializedValue SerializeValue(const value& val) {
    return nonstd::visit(Serializer{}, val);
}

// marshal_uint256_t instead?
// make sure correct
// std::vector<unsigned char> ConvertToCharVector(uint256_t value) {
//
//    unsigned char value_char_list[sizeof(value)];
//    std::memcpy(value_char_list, &value, sizeof(value));
//    std::vector<unsigned char> return_vector;
//
//    return_vector.insert(return_vector.end(), std::begin(value_char_list),
//    std::end(value_char_list));
//
//    return return_vector;
//}
