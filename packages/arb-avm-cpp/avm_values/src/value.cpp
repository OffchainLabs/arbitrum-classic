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

namespace {
Tuple deserializeTuple(const char*& bufptr, int size, TuplePool& pool) {
    Tuple tup;
    if (size > 0) {
        tup = Tuple(&pool, size);
        for (int i = 0; i < size; i++) {
            tup.set_element(i, deserialize_value(bufptr, pool));
        }
    }

    return tup;
}

CodePoint deserializeCodePoint(const char*& bufptr, TuplePool& pool) {
    CodePoint ret;
    memcpy(&ret.pc, bufptr, sizeof(ret.pc));
    bufptr += sizeof(ret.pc);
    ret.pc = boost::endian::big_to_native(ret.pc);
    ret.op = deserializeOperation(bufptr, pool);
    ret.nextHash = from_big_endian(bufptr, bufptr + UINT256_SIZE);
    bufptr += UINT256_SIZE;

    return ret;
}
}  // namespace

uint256_t deserializeUint256t(const char*& bufptr) {
    uint256_t ret = from_big_endian(bufptr, bufptr + UINT256_SIZE);
    bufptr += UINT256_SIZE;
    return ret;
}

Operation deserializeOperation(const char*& bufptr, TuplePool& pool) {
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

value deserialize_value(const char*& bufptr, TuplePool& pool) {
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

void marshal_HashPreImage(const HashPreImage& val,
                          std::vector<unsigned char>& buf) {
    val.marshal(buf);
}

void marshal_uint256_t(const uint256_t& val, std::vector<unsigned char>& buf) {
    std::array<unsigned char, 32> tmpbuf;
    to_big_endian(val, tmpbuf.begin());
    buf.insert(buf.end(), tmpbuf.begin(), tmpbuf.end());
}

void marshal_value(const value& val, std::vector<unsigned char>& buf) {
    if (nonstd::holds_alternative<Tuple>(val)) {
        nonstd::get<Tuple>(val).marshal(buf);
    } else if (nonstd::holds_alternative<uint256_t>(val)) {
        buf.push_back(NUM);
        marshal_uint256_t(nonstd::get<uint256_t>(val), buf);
    } else if (nonstd::holds_alternative<CodePoint>(val)) {
        nonstd::get<CodePoint>(val).marshal(buf);
    } else if (nonstd::holds_alternative<HashPreImage>(val)) {
        marshal_HashPreImage(nonstd::get<HashPreImage>(val), buf);
    }
}

void marshalStub(const value& val, std::vector<unsigned char>& buf) {
    if (nonstd::holds_alternative<Tuple>(val)) {
        auto tup = nonstd::get<Tuple>(val);
        marshalForProof(tup.getHashPreImage(), buf);
    } else {
        marshalForProof(val, buf);
    }
}

namespace {
// see similar functionality in tuple.cloneShallow and tuple.marshal
void marshalForProof(const Tuple& val, std::vector<unsigned char>& buf) {
    buf.push_back(TUPLE + val.tuple_size());
    for (uint64_t i = 0; i < val.tuple_size(); i++) {
        auto itemval = val.get_element(i);
        marshalStub(itemval, buf);
    }
}

void marshalForProof(const CodePoint& val, std::vector<unsigned char>& buf) {
    buf.push_back(CODEPT);
    val.op.marshalForProof(buf, false);
    std::array<unsigned char, 32> hashVal;
    to_big_endian(val.nextHash, hashVal.begin());
    buf.insert(buf.end(), hashVal.begin(), hashVal.end());
}

void marshalForProof(const uint256_t& val, std::vector<unsigned char>& buf) {
    buf.push_back(NUM);
    marshal_uint256_t(val, buf);
}

void marshalForProof(const HashPreImage& val, std::vector<unsigned char>& buf) {
    buf.push_back(HASH_PRE_IMAGE);
    val.marshal(buf);
}
}  // namespace

void marshalForProof(const value& val, std::vector<unsigned char>& buf) {
    return nonstd::visit([&](const auto& v) { return marshalForProof(v, buf); },
                         val);
}

uint256_t hash(const value& value) {
    return nonstd::visit([](const auto& val) { return hash(val); }, value);
}

struct GetSize {
    uint256_t operator()(const HashPreImage& val) const {
        return val.getSize();
    }

    uint256_t operator()(const Tuple& val) const { return val.getSize(); }

    uint256_t operator()(const uint256_t& val) const { return 1; }

    uint256_t operator()(const CodePoint& val) const { return val.getSize(); }
};

uint256_t getSize(const value& val) {
    return nonstd::visit(GetSize{}, val);
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

    std::ostream* operator()(const HashPreImage& val) const {
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
