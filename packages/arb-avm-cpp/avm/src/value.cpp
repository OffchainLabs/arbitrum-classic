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

uint256_t deserialize_int(char*& bufptr) {
    uint256_t ret = from_big_endian(bufptr, bufptr + UINT256_SIZE);
    bufptr += UINT256_SIZE;
    return ret;
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
    std::vector<unsigned char> tmpbuf;
    tmpbuf.resize(32);
    to_big_endian(val, tmpbuf.begin());
    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

void marshal_value(const value val, std::vector<unsigned char>& buf) {
    if (nonstd::holds_alternative<Tuple>(val))
        marshal_Tuple(nonstd::get<Tuple>(val), buf);
    else if (nonstd::holds_alternative<uint256_t>(val))
        marshal_uint256_t(nonstd::get<uint256_t>(val), buf);
    else if (nonstd::holds_alternative<CodePoint>(val))
        marshal_CodePoint(nonstd::get<CodePoint>(val), buf);
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

uint256_t hash(const value& value) {
    return nonstd::visit([](const auto& val) { return hash(val); }, value);
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
