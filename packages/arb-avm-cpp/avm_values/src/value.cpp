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

#include <avm_values/codepoint.hpp>
#include <avm_values/codepointstub.hpp>
#include <avm_values/pool.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

#include <avm_values/util.hpp>
#include <bigint_utils.hpp>

#include <ostream>

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
}  // namespace

CodePointRef deserializeCodePointRef(const char*& bufptr) {
    uint64_t pc;
    memcpy(&pc, bufptr, sizeof(pc));
    pc = boost::endian::big_to_native(pc);
    bufptr += sizeof(pc);
    bool is_err = static_cast<bool>(*bufptr);
    ++bufptr;
    return {pc, is_err};
}

CodePointStub deserializeCodePointStub(const char*& bufptr) {
    uint64_t pc;
    memcpy(&pc, bufptr, sizeof(pc));
    bufptr += sizeof(pc);
    pc = boost::endian::big_to_native(pc);
    auto hash_val = deserializeUint256t(bufptr);
    return {pc, hash_val};
}

uint256_t deserializeUint256t(const char*& bufptr) {
    uint256_t ret = from_big_endian(bufptr, bufptr + UINT256_SIZE);
    bufptr += UINT256_SIZE;
    return ret;
}

value deserialize_value(const char*& bufptr, TuplePool& pool) {
    uint8_t valType;
    memcpy(&valType, bufptr, sizeof(valType));
    bufptr += sizeof(valType);
    switch (valType) {
        case NUM:
            return deserializeUint256t(bufptr);
        case CODEPT: {
            return deserializeCodePointStub(bufptr);
        }
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
    } else if (nonstd::holds_alternative<CodePointStub>(val)) {
        buf.push_back(CODE_POINT_STUB);
        nonstd::get<CodePointStub>(val).marshal(buf);
    } else if (nonstd::holds_alternative<HashPreImage>(val)) {
        buf.push_back(HASH_PRE_IMAGE);
        nonstd::get<HashPreImage>(val).marshal(buf);
    }
}

namespace {
void marshalForProof(const HashPreImage& val,
                     MarshalLevel,
                     std::vector<unsigned char>& buf,
                     const Code&) {
    buf.push_back(HASH_PRE_IMAGE);
    val.marshal(buf);
}

void marshalForProof(const Tuple& val,
                     MarshalLevel marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code) {
    if (marshal_level == MarshalLevel::STUB) {
        marshalForProof(val.getHashPreImage(), marshal_level, buf, code);
    } else {
        buf.push_back(TUPLE + val.tuple_size());
        MarshalLevel nested_level = MarshalLevel::STUB;
        if (marshal_level == MarshalLevel::FULL) {
            nested_level = MarshalLevel::FULL;
        }
        for (uint64_t i = 0; i < val.tuple_size(); i++) {
            auto itemval = val.get_element(i);
            marshalForProof(itemval, nested_level, buf, code);
        }
    }
}

void marshalForProof(const CodePointStub& val,
                     MarshalLevel marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code) {
    auto& cp = code[val.pc];
    MarshalLevel nested_level = MarshalLevel::STUB;
    if (marshal_level == MarshalLevel::FULL) {
        nested_level = MarshalLevel::FULL;
    }
    buf.push_back(CODEPT);
    cp.op.marshalForProof(buf, nested_level, code);
    std::array<unsigned char, 32> hashVal;
    to_big_endian(cp.nextHash, hashVal.begin());
    buf.insert(buf.end(), hashVal.begin(), hashVal.end());
}

void marshalForProof(const uint256_t& val,
                     MarshalLevel,
                     std::vector<unsigned char>& buf,
                     const Code&) {
    buf.push_back(NUM);
    marshal_uint256_t(val, buf);
}
}  // namespace

void marshalForProof(const value& val,
                     MarshalLevel marshal_level,
                     std::vector<unsigned char>& buf,
                     const Code& code) {
    return nonstd::visit(
        [&](const auto& v) {
            return marshalForProof(v, marshal_level, buf, code);
        },
        val);
}

uint256_t hash_value(const value& value) {
    return nonstd::visit([](const auto& val) { return hash(val); }, value);
}

struct GetSize {
    uint256_t operator()(const HashPreImage& val) const {
        return val.getSize();
    }

    uint256_t operator()(const Tuple& val) const { return val.getSize(); }

    uint256_t operator()(const uint256_t&) const { return 1; }

    uint256_t operator()(const CodePointStub&) const { return 1; }
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

    std::ostream* operator()(const CodePointStub& val) const {
        //        std::printf("in CodePoint ostream operator\n");
        os << "CodePointStub(" << val.pc.pc << ")";
        return &os;
    }
};

std::ostream& operator<<(std::ostream& os, const value& val) {
    return *nonstd::visit(ValuePrinter{os}, val);
}
