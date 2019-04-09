//
//  code.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef code_hpp
#define code_hpp

#include "value.hpp"
#include "opcodes.hpp"

#include <nonstd/optional.hpp>

//class codept{
//public:
//    int pc;
//    int opcode;
//    uint256_t nextCodeHash;
//};

class instr {
public:
    unsigned long long pc;
    OpCode opcode;
    uint256_t codeHash;
    nonstd::optional<value> immediate;
    //    codept label;
    
    instr(unsigned long long pc_, OpCode opcode_, uint256_t codeHash_, value && immediate_) :
    pc(pc_),
    opcode(opcode_),
    codeHash(codeHash_),
    immediate(immediate_) {}
    
    instr(unsigned long long pc_, OpCode opcode_, uint256_t codeHash_) :
    pc(pc_),
    opcode(opcode_),
    codeHash(codeHash_) {}
};

std::ostream& operator<<(std::ostream& os, const instr& instruction);

#endif /* code_hpp */
