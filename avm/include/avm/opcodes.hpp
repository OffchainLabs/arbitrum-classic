//
//  opcodes.hpp
//  AVMtest
//
//  Created by Harry Kalodner on 4/2/19.
//

#ifndef opcodes_h
#define opcodes_h

#define CURRENT_AO_VERSION 1

#include <cstdint>
#include <map>
#include <string>

enum states { EXTENSIVE, HALTED, ERROR, BLOCKED };

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
    SPUSH,          // 31
    RPUSH,          // 32
    RSET,           // 33
    JUMP,           // 34
    CJUMP,          // 35
    STACKEMPTY,     // 36
    PCPUSH,         // 37
    AUXPUSH,        // 38
    AUXPOP,         // 39
    AUXSTACKEMPTY,  // 3a
    NOP,            // 3b
    ERRPUSH,
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
    HALT,
    ERROR,
    DEBUG
};

const std::map<OpCode, std::string> InstructionNames = {
    {static_cast<OpCode>(0), "unhandled opcode"},
    {OpCode::ADD, "add"},
    {OpCode::MUL, "mul"},
    {OpCode::SUB, "sub"},
    {OpCode::DIV, "div"},
    {OpCode::SDIV, "sdiv"},
    {OpCode::MOD, "mod"},
    {OpCode::SMOD, "smod"},
    {OpCode::ADDMOD, "addmod"},
    {OpCode::MULMOD, "mulmod"},
    {OpCode::EXP, "exp"},

    {OpCode::LT, "lt"},
    {OpCode::GT, "gt"},
    {OpCode::SLT, "slt"},
    {OpCode::SGT, "sgt"},
    {OpCode::EQ, "eq"},
    {OpCode::ISZERO, "iszero"},
    {OpCode::BITWISE_AND, "and"},
    {OpCode::BITWISE_OR, "or"},
    {OpCode::BITWISE_XOR, "xor"},
    {OpCode::BITWISE_NOT, "not"},
    {OpCode::BYTE, "byte"},
    {OpCode::SIGNEXTEND, "signextend"},

    {OpCode::HASH, "hash"},
    {OpCode::TYPE, "type"},

    {OpCode::POP, "pop"},
    {OpCode::SPUSH, "spush"},
    {OpCode::RPUSH, "rpush"},
    {OpCode::RSET, "rset"},
    {OpCode::JUMP, "jump"},
    {OpCode::CJUMP, "cjump"},
    {OpCode::STACKEMPTY, "stackempty"},
    {OpCode::PCPUSH, "pcpush"},
    {OpCode::AUXPUSH, "auxpush"},
    {OpCode::AUXPOP, "auxpop"},
    {OpCode::AUXSTACKEMPTY, "auxstackempty"},
    {OpCode::NOP, "nop"},
    {OpCode::ERRPUSH, "errpush"},
    {OpCode::ERRSET, "errset"},

    {OpCode::DUP0, "dup0"},
    {OpCode::DUP1, "dup1"},
    {OpCode::DUP2, "dup2"},
    {OpCode::SWAP1, "swap1"},
    {OpCode::SWAP2, "swap2"},

    {OpCode::TGET, "tget"},
    {OpCode::TSET, "tset"},
    {OpCode::TLEN, "tlen"},

    {OpCode::BREAKPOINT, "breakpoint"},
    {OpCode::LOG, "log"},

    {OpCode::SEND, "send"},
    {OpCode::NBSEND, "nbsend"},
    {OpCode::GETTIME, "gettime"},
    {OpCode::INBOX, "inbox"},
    {OpCode::ERROR, "error"},
    {OpCode::HALT, "halt"},
    {OpCode::DEBUG, "debug"}};

#endif /* opcodes_h */
