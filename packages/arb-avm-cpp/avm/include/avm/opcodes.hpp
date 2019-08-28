/*
 * Copyright 2019, Offchain Labs, Inc.
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

#ifndef opcodes_hpp
#define opcodes_hpp

#define CURRENT_AO_VERSION 1

#include <cstdint>
#include <map>
#include <string>

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

inline bool isValidOpcode(OpCode op) {
    return (op >= OpCode::ADD && op <= OpCode::EXP) ||
           (op >= OpCode::LT && op <= OpCode::SIGNEXTEND) ||
           (op >= OpCode::HASH && op <= OpCode::TYPE) ||
           (op >= OpCode::POP && op <= OpCode::ERRSET) ||
           (op >= OpCode::DUP0 && op <= OpCode::SWAP2) ||
           (op >= OpCode::TGET && op <= OpCode::TLEN) ||
           (op >= OpCode::BREAKPOINT && op <= OpCode::LOG) ||
           (op >= OpCode::SEND && op <= OpCode::ERROR);
}

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

const std::map<OpCode, std::vector<bool>> InstructionStackPops = {
    {static_cast<OpCode>(0), {}},
    {OpCode::ADD, {true, true}},
    {OpCode::MUL, {true, true}},
    {OpCode::SUB, {true, true}},
    {OpCode::DIV, {true, true}},
    {OpCode::SDIV, {true, true}},
    {OpCode::MOD, {true, true}},
    {OpCode::SMOD, {true, true}},
    {OpCode::ADDMOD, {true, true}},
    {OpCode::MULMOD, {true, true}},
    {OpCode::EXP, {true, true}},

    {OpCode::LT, {true, true}},
    {OpCode::GT, {true, true}},
    {OpCode::SLT, {true, true}},
    {OpCode::SGT, {true, true}},
    {OpCode::EQ, {false, false}},
    {OpCode::ISZERO, {true, true}},
    {OpCode::BITWISE_AND, {true, true}},
    {OpCode::BITWISE_OR, {true, true}},
    {OpCode::BITWISE_XOR, {true, true}},
    {OpCode::BITWISE_NOT, {true}},
    {OpCode::BYTE, {true, true}},
    {OpCode::SIGNEXTEND, {true, true}},

    {OpCode::HASH, {false}},
    {OpCode::TYPE, {true}},

    {OpCode::POP, {false}},
    {OpCode::SPUSH, {}},
    {OpCode::RPUSH, {}},
    {OpCode::RSET, {false}},
    {OpCode::JUMP, {false}},
    {OpCode::CJUMP, {false, true}},
    {OpCode::STACKEMPTY, {}},
    {OpCode::PCPUSH, {}},
    {OpCode::AUXPUSH, {false}},
    {OpCode::AUXPOP, {}},
    {OpCode::AUXSTACKEMPTY, {}},
    {OpCode::NOP, {}},
    {OpCode::ERRPUSH, {}},
    {OpCode::ERRSET, {true}},

    {OpCode::DUP0, {false}},
    {OpCode::DUP1, {false, false}},
    {OpCode::DUP2, {false, false, false}},
    {OpCode::SWAP1, {false, false}},
    {OpCode::SWAP2, {false, false, false}},

    {OpCode::TGET, {true, true}},
    {OpCode::TSET, {true, true, false}},
    {OpCode::TLEN, {true}},

    {OpCode::BREAKPOINT, {}},
    {OpCode::LOG, {false}},

    {OpCode::SEND, {true}},
    {OpCode::NBSEND, {true}},
    {OpCode::GETTIME, {}},
    {OpCode::INBOX, {false}},
    {OpCode::ERROR, {}},
    {OpCode::HALT, {}},
    {OpCode::DEBUG, {}}};

const std::map<OpCode, std::vector<bool>> InstructionAuxStackPops = {
    {static_cast<OpCode>(0), {}},
    {OpCode::ADD, {}},
    {OpCode::MUL, {}},
    {OpCode::SUB, {}},
    {OpCode::DIV, {}},
    {OpCode::SDIV, {}},
    {OpCode::MOD, {}},
    {OpCode::SMOD, {}},
    {OpCode::ADDMOD, {}},
    {OpCode::MULMOD, {}},
    {OpCode::EXP, {}},

    {OpCode::LT, {}},
    {OpCode::GT, {}},
    {OpCode::SLT, {}},
    {OpCode::SGT, {}},
    {OpCode::EQ, {}},
    {OpCode::ISZERO, {}},
    {OpCode::BITWISE_AND, {}},
    {OpCode::BITWISE_OR, {}},
    {OpCode::BITWISE_XOR, {}},
    {OpCode::BITWISE_NOT, {}},
    {OpCode::BYTE, {}},
    {OpCode::SIGNEXTEND, {}},

    {OpCode::HASH, {}},
    {OpCode::TYPE, {}},

    {OpCode::POP, {}},
    {OpCode::SPUSH, {}},
    {OpCode::RPUSH, {}},
    {OpCode::RSET, {}},
    {OpCode::JUMP, {}},
    {OpCode::CJUMP, {}},
    {OpCode::STACKEMPTY, {}},
    {OpCode::PCPUSH, {}},
    {OpCode::AUXPUSH, {}},
    {OpCode::AUXPOP, {false}},
    {OpCode::AUXSTACKEMPTY, {}},
    {OpCode::NOP, {}},
    {OpCode::ERRPUSH, {}},
    {OpCode::ERRSET, {}},

    {OpCode::DUP0, {}},
    {OpCode::DUP1, {}},
    {OpCode::DUP2, {}},
    {OpCode::SWAP1, {}},
    {OpCode::SWAP2, {}},

    {OpCode::TGET, {}},
    {OpCode::TSET, {}},
    {OpCode::TLEN, {}},

    {OpCode::BREAKPOINT, {}},
    {OpCode::LOG, {}},

    {OpCode::SEND, {}},
    {OpCode::NBSEND, {}},
    {OpCode::GETTIME, {}},
    {OpCode::INBOX, {}},
    {OpCode::ERROR, {}},
    {OpCode::HALT, {}},
    {OpCode::DEBUG, {}}};

#endif /* opcodes_hpp */
