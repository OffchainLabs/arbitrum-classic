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

#include <avm_values/codepoint.hpp>
#include <avm_values/pool.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

#include <avm_values/util.hpp>
#include <bigint_utils.hpp>

#include <ostream>

#define UINT256_SIZE 32

auto deserializeUint256t(const char*& bufptr) -> uint256_t {
    uint256_t ret = from_big_endian(bufptr, bufptr + UINT256_SIZE);
    bufptr += UINT256_SIZE;
    return ret;
}

auto deserializeOperation(const char*& bufptr, TuplePool& pool) -> Operation {
    uint8_t immediateCount;
    memcpy(&immediateCount, bufptr, sizeof(immediateCount));
    bufptr += sizeof(immediateCount);
    OpCode opcode;
    memcpy(&opcode, bufptr, sizeof(opcode));
    bufptr += sizeof(opcode);

    if (immediateCount == 1) {
        return {opcode, deserialize_value(bufptr, pool)};
    } else {
        return Operation{opcode};
    }
}

auto deserializeCodePoint(const char*& bufptr, TuplePool& pool) -> CodePoint {
    CodePoint ret;
    memcpy(&ret.pc, bufptr, sizeof(ret.pc));
    bufptr += sizeof(ret.pc);
    ret.pc = boost::endian::big_to_native(ret.pc);
    ret.op = deserializeOperation(bufptr, pool);
    ret.nextHash = from_big_endian(bufptr, bufptr + UINT256_SIZE);
    bufptr += UINT256_SIZE;

    return ret;
}

auto deserializeTuple(const char*& bufptr, int size, TuplePool& pool) -> Tuple {
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
    std::array<unsigned char, 32> tmpbuf;
    to_big_endian(val, tmpbuf.begin());
    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

void marshal_value(const value& val, std::vector<unsigned char>& buf) {
    if (nonstd::holds_alternative<Tuple>(val)) {
        marshal_Tuple(nonstd::get<Tuple>(val), buf);
    } else if (nonstd::holds_alternative<uint256_t>(val)) {
        buf.push_back(NUM);
        marshal_uint256_t(nonstd::get<uint256_t>(val), buf);
    } else if (nonstd::holds_alternative<CodePoint>(val)) {
        marshal_CodePoint(nonstd::get<CodePoint>(val), buf);
    }
}

void marshalShallow(const value& val, std::vector<unsigned char>& buf) {
    return nonstd::visit([&](const auto& v) { return marshalShallow(v, buf); },
                         val);
}

// see similar functionality in tuple.cloneShallow and tuple.marshal
void marshalShallow(const Tuple& val, std::vector<unsigned char>& buf) {
    buf.push_back(TUPLE + val.tuple_size());
    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto itemval = val.get_element(i);
        if (nonstd::holds_alternative<uint256_t>(itemval)) {
            marshalShallow(itemval, buf);
        } else {
            buf.push_back(HASH_ONLY);
            std::array<unsigned char, 32> tmpbuf;
            to_big_endian(::hash(val.get_element(i)), tmpbuf.begin());
            buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
        }
    }
}

void marshalShallow(const CodePoint& val, std::vector<unsigned char>& buf) {
    buf.push_back(CODEPT);
    val.op.marshalForProof(buf, false);
    std::array<unsigned char, 32> hashVal;
    to_big_endian(val.nextHash, hashVal.begin());
    buf.insert(buf.end(), hashVal.begin(), hashVal.end());
}

void marshalShallow(const uint256_t& val, std::vector<unsigned char>& buf) {
    buf.push_back(NUM);
    marshal_uint256_t(val, buf);
}

auto deserialize_value(const char*& bufptr, TuplePool& pool) -> value {
    uint8_t valType;
    memcpy(&valType, bufptr, sizeof(valType));
    bufptr += sizeof(valType);
    switch (valType) {
        case NUM:
            return deserializeUint256t(bufptr);
        case CODEPT:
            return deserializeCodePoint(bufptr, pool);
        default:
            if (valType >= TUPLE && valType <= TUPLE + 8) {
                return deserializeTuple(bufptr, valType - TUPLE, pool);
            } else {
                std::printf("in deserialize_value, unhandled type = %X\n",
                            valType);
                throw std::runtime_error("Tried to deserialize unhandled type");
            }
    }
}

auto get_tuple_size(const char*& bufptr) -> int {
    uint8_t valType;
    memcpy(&valType, bufptr, sizeof(valType));

    return valType - TUPLE;
}

auto hash(const value& value) -> uint256_t {
    return nonstd::visit([](const auto& val) { return hash(val); }, value);
}

struct ValuePrinter {
    std::ostream& os;

    auto operator()(const Tuple& val) const -> std::ostream* {
        os << val;
        return &os;
    }

    auto operator()(const uint256_t& val) const -> std::ostream* {
        os << val;
        return &os;
    }

    auto operator()(const CodePoint& val) const -> std::ostream* {
        //        std::printf("in CodePoint ostream operator\n");
        os << "CodePoint(" << val.pc << ", " << val.op << ", "
           << to_hex_str(val.nextHash) << ")";
        return &os;
    }
};

auto operator<<(std::ostream& os, const value& val) -> std::ostream& {
    return *nonstd::visit(ValuePrinter{os}, val);
}

auto GetHashKey(const value& val) -> std::vector<unsigned char> {
    auto hash_key = hash(val);
    std::vector<unsigned char> hash_key_vector;
    marshal_value(hash_key, hash_key_vector);

    return hash_key_vector;
}
