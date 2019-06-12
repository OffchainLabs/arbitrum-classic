//
//  opcodes.hpp
//  AVMtest
//
//  Created by Harry Kalodner on 4/2/19.
//

#ifndef opcodes_h
#define opcodes_h

#define CURRENT_AO_VERSION 1

enum states{EXTENSIVE, HALTED, ERROR};

enum class OpCode : uint8_t {
    ADD = 0x01,
    MUL,
    SUB,
    DIV,
    SDIV,
    MOD,
    SMOD,
    ADDMOD,
    MULMOD,
    EXP,
    
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
    SIGNEXTEND,

    HASH = 0x20,
    TYPE,
    
    POP = 0x30,
    SPUSH,//31
    RPUSH,//32
    RSET,//33
//    UNUSED, // place holder for old .ao opcodes
    JUMP,//34
    CJUMP,//35
    STACKEMPTY,//36
    PCPUSH,//37
    AUXPUSH,//38
    AUXPOP,//39
    AUXSTACKEMPTY,//3a
    NOP,//3b
    ERRPUCH,
    ERRSET,
    
    DUP0 = 0x40,
    DUP1,
    DUP2,
    SWAP1,
    SWAP2,
    
    TGET = 0x50,
    TSET,
    TLEN,
    
    BREAKPOINT = 0x60,
    LOG,

    SEND = 0x70,
    NBSEND,
    GETTIME,
    INBOX,
    HALT = 0x00,
    ERROR = 0x75,
    DEBUG =0x076
};

#endif /* opcodes_h */
