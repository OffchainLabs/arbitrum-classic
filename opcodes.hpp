//
//  opcodes.hpp
//  AVMtest
//
//  Created by Harry Kalodner on 4/2/19.
//

#ifndef opcodes_h
#define opcodes_h

enum states{EXTENSIVE, HALTED, ERROR};

enum class OpCode : uint8_t {
    HALT = 0x00,
    ADD,
    MUL,
    SUB,
    DIV,
    SDIV,
    MOD,
    SMOD,
    ADDMOD,
    MULMOD,
    EXP,
    SIGNEXTEND,
    
    LT = 0x10,
    GT,
    SLT,
    SGT,
    EQ,
    ISZERO,
    BITWISE_AND,
    BITWISE_OR,
    BITWISE_XOR,
    BITWISE_NOT,
    BYTE,
    
    HASH = 0x20,
    
    POP = 0x30,
    SPUSH,
    RPUSH,
    RSET,
    INBOX,
    JUMP,
    CJUMP,
    STACKEMPTY,
    PCPUSH,
    AUXPUSH,
    AUXPOP,
    AUXSTACKEMPTY,
    NOP,
    
    DUP0 = 0x40,
    DUP1,
    DUP2,
    SWAP1,
    SWAP2,
    
    TGET = 0x50,
    TSET,
    TLEN,
    ISTUPLE,
    
    PRINTCHAR = 0x60,
    ADVISE,
    DEBUG,
    
    SEND = 0x70,
    INCATOMIC,
    DECATOMIC,
    GETTIME
};

#endif /* opcodes_h */
