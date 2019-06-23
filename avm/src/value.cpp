//
//  value.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/28/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#include <avm/value.hpp>

#include <avm/code.hpp>
#include <avm/pool.hpp>
#include <avm/util.hpp>

#include <iostream>

#define UINT256_SIZE 32

uint256_t deserialize_int(char*& bufptr) {
    uint256_t ret = from_big_endian(bufptr, bufptr + UINT256_SIZE);
    bufptr += UINT256_SIZE;
    return ret;
}

OpCode deserialize_code_point_opcode(char*& bufptr, TuplePool& pool) {
    uint8_t immediateCount;
    memcpy(&immediateCount, bufptr, sizeof(immediateCount));
    bufptr += sizeof(immediateCount);

    OpCode opcode;
    memcpy(&opcode, bufptr, sizeof(opcode));
    bufptr += sizeof(opcode);

    if (immediateCount == 1) {
        deserialize_value(bufptr, pool);
    }
    return opcode;
}

CodePoint deserialize_codepoint(char*& bufptr, TuplePool& pool) {
    CodePoint ret;
    memcpy(&ret.pc, bufptr, sizeof(ret.pc));
    bufptr += sizeof(ret.pc);
    ret.pc = __builtin_bswap64(ret.pc);
    ret.op = deserialize_code_point_opcode(bufptr, pool);
    memcpy(&ret.nexthash, bufptr, UINT256_SIZE);
    bufptr += UINT256_SIZE;

    return ret;
}

Tuple deserialize_tuple(char*& bufptr, int size, TuplePool& pool) {
    Tuple tup(size, &pool);
    for (int i = 0; i < size; i++) {
        tup.set_element(i, deserialize_value(bufptr, pool));
    }
    return tup;
}

value deserialize_value(char*& bufptr, TuplePool& pool) {
    uint8_t valType;
    memcpy(&valType, bufptr, sizeof(valType));
    bufptr += sizeof(valType);
    switch (valType) {
        case NUM:
            return deserialize_int(bufptr);
        case CODEPT:
            return deserialize_codepoint(bufptr, pool);
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

bool operator==(const CodePoint& val1, const CodePoint& val2) {
    if (val1.pc != val2.pc)
        return false;
    else
        return true;
}

std::vector<unsigned char> value_hash_raw(const value& value) {
    return mpark::visit([](const auto& val) { return value_hash_raw(val); },
                        value);
}

std::vector<unsigned char> value_hash_raw(const uint256_t& value) {
    std::vector<unsigned char> intData;
    intData.resize(32);
    to_big_endian(value, intData.begin());

    std::vector<unsigned char> hashData;
    hashData.resize(32);
    evm::Keccak_256(intData.data(), 32, hashData.data());
    return hashData;
}

std::vector<unsigned char> value_hash_raw(const Tuple& tup) {
    std::vector<unsigned char> tupData;
    tupData.resize(1 + tup.tuple_size() * 32);
    auto oit = tupData.begin();
    tupData[0] = TUPLE + tup.tuple_size();
    ++oit;
    for (int i = 0; i < tup.tuple_size(); i++) {
        auto valHash = value_hash_raw(tup.get_element(i));
        std::copy(valHash.begin(), valHash.end(), oit);
        oit += 32;
    }

    std::vector<unsigned char> hashData;
    hashData.resize(32);
    evm::Keccak_256(tupData.data(), 32, hashData.data());
    return hashData;
}

std::vector<unsigned char> value_hash_raw(const CodePoint&) {
    throw std::runtime_error("CodePoint hash not supported");
}

uint256_t value_hash(const value& value) {
    return mpark::visit(
        [](const auto& val) {
            auto raw = value_hash_raw(val);
            return from_big_endian(raw.begin(), raw.end());
        },
        value);
}

struct ValuePrinter {
    std::ostream& os;

    std::ostream& operator()(const Tuple& val) {
        os << "tuple="
           << " [";
        for (int i = 0; i < val.tuple_size(); i++) {
            std::cout << val.get_element(i)
                      << ((i < val.tuple_size() - 1) ? "," : "");
        }
        std::cout << "]";
        return os;
    }

    std::ostream& operator()(const uint256_t& val) {
        os << "num=" << val;
        return os;
    }

    std::ostream& operator()(const CodePoint& val) {
        //        std::printf("in CodePoint ostream operator\n");
        os << "codept pc=" << val.pc << " opcode=" << val.op;
        return os;
    }
};

std::ostream& operator<<(std::ostream& os, const value& val) {
    return mpark::visit(ValuePrinter{os}, val);
}
