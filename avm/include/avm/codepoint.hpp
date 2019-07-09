//
//  codepoint.hpp
//  AVMtest
//
//  Created by Harry Kalodner on 6/23/19.
//

#ifndef codepoint_h
#define codepoint_h

#include "value.hpp"

struct Operation {
    OpCode opcode;
    std::unique_ptr<value> immediate;

    Operation() = default;
    Operation(OpCode opcode_) : opcode(opcode_) {}
    Operation(OpCode opcode_, value val);

    Operation(const Operation&);
    Operation(Operation&&);
    Operation& operator=(const Operation&);
    Operation& operator=(Operation&&);
    ~Operation();
    void marshal(std::vector<unsigned char>& buf) const;
};

struct CodePoint {
    uint64_t pc;
    Operation op;
    uint256_t nextHash;

    CodePoint() {}
    CodePoint(uint64_t pc_, Operation op_, uint256_t nextHash_)
        : pc(pc_), op(op_), nextHash(nextHash_) {}
    void marshal(std::vector<unsigned char>& buf) const;
};

uint256_t hash(const CodePoint& cp);

std::ostream& operator<<(std::ostream& os, const Operation& val);

#endif /* codepoint_h */
