//
//  value.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/28/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#include <avm/value.hpp>

#include <avm/codepoint.hpp>
#include <avm/tuple.hpp>

#include <avm/pool.hpp>
#include <avm/util.hpp>

#include <iostream>

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

bool operator==(const CodePoint& val1, const CodePoint& val2) {
    if (val1.pc != val2.pc)
        return false;
    else
        return true;
}

uint256_t hash(const value& value) {
    return mpark::visit([](const auto& val) { return hash(val); }, value);
}

struct ValuePrinter {
    std::ostream& os;

    std::ostream& operator()(const Tuple& val) { return os << val; }

    std::ostream& operator()(const uint256_t& val) { return os << val; }

    std::ostream& operator()(const CodePoint& val) {
        //        std::printf("in CodePoint ostream operator\n");
        os << "CodePoint(" << val.pc << ", " << val.op << ", "
           << to_hex_str(val.nextHash) << ")";
        return os;
    }
};

std::ostream& operator<<(std::ostream& os, const value& val) {
    return mpark::visit(ValuePrinter{os}, val);
}
