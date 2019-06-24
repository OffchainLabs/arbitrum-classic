//
//  value.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/25/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef value_h
#define value_h

#include "bigint.hpp"
#include "opcodes.hpp"

#include <mpark/variant.hpp>

class bad_tuple_index : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "bad_tuple_index";
    }
};

enum types { NUM, CODEPT, TUPLE = 3 };

class TuplePool;
class Tuple;
struct CodePoint;

// Note: uint256_t is actually 48 bytes long
using value = mpark::variant<uint256_t, Tuple, CodePoint>;

std::ostream& operator<<(std::ostream& os, const value& val);
bool operator==(const CodePoint& val1, const CodePoint& val2);

uint256_t value_hash(const value& value);

struct CodePoint {
    uint64_t pc;
    OpCode op;
    uint256_t nexthash;

    CodePoint() {}
    CodePoint(uint64_t pc_) : pc(pc_) {}
};

uint256_t deserialize_int(char*& srccode);
CodePoint deserialize_codepoint(char*& srccode);
Tuple deserialize_tuple(char*& bufptr, int size, TuplePool& pool);
value deserialize_value(char*& srccode, TuplePool& pool);

uint256_t value_hash(const value& value);

#endif /* value_h */
