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
#include <nonstd/optional.hpp>

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
using value = mpark::variant<Tuple, uint256_t, CodePoint>;

std::ostream& operator<<(std::ostream& os, const value& val);
bool operator==(const CodePoint& val1, const CodePoint& val2);

uint256_t hash(const value& value);
void warmHash(value& val);

uint256_t deserialize_int(char*& srccode);
Operation deserializeOperation(char*& bufptr, TuplePool& pool);
CodePoint deserializeCodePoint(char*& bufptr, TuplePool& pool);
Tuple deserialize_tuple(char*& bufptr, int size, TuplePool& pool);
value deserialize_value(char*& srccode, TuplePool& pool);

#endif /* value_h */
