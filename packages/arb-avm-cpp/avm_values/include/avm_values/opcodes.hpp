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

#include <array>
#include <cstdint>
#include <limits>
#include <string>
#include <unordered_map>
#include <vector>

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
    SHL,
    SHR,
    SAR,

    HASH = 0x20,
    TYPE,
    ETHHASH2,
    KECCAKF,
    SHA256F,

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
    XGET,
    XSET,

    BREAKPOINT = 0x60,
    LOG,

    SEND = 0x70,
    // INBOX_PEEK,
    INBOX = 0x72,
    ERROR,
    HALT,
    SET_GAS,
    PUSH_GAS,
    ERR_CODE_POINT,
    PUSH_INSN,
    PUSH_INSN_IMM,
    //    OPEN_INSN,
    SIDELOAD = 0x7B,

    ECRECOVER = 0x80,
    ECADD,
    ECMUL,
    ECPAIRING,

    DEBUG_PRINT = 0x90,

    NEW_BUFFER = 0xa0,
    GET_BUFFER8,
    GET_BUFFER64,
    GET_BUFFER256,
    SET_BUFFER8,
    SET_BUFFER64,
    SET_BUFFER256,
};

const std::unordered_map<OpCode, std::string> InstructionNames = {
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
    {OpCode::SIGNEXTEND, "signextend"},

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
    {OpCode::SHL, "shl"},
    {OpCode::SHR, "shr"},
    {OpCode::SAR, "sar"},

    {OpCode::HASH, "hash"},
    {OpCode::TYPE, "type"},
    {OpCode::ETHHASH2, "ethhash2"},
    {OpCode::KECCAKF, "keccakf"},
    {OpCode::SHA256F, "sha256f"},

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
    {OpCode::XGET, "xget"},
    {OpCode::XSET, "xset"},

    {OpCode::BREAKPOINT, "breakpoint"},
    {OpCode::LOG, "log"},

    {OpCode::SEND, "send"},
    {OpCode::INBOX, "inbox"},
    {OpCode::ERROR, "error"},
    {OpCode::HALT, "halt"},
    {OpCode::SET_GAS, "setgas"},
    {OpCode::PUSH_GAS, "pushgas"},
    {OpCode::ERR_CODE_POINT, "errcodepoint"},
    {OpCode::PUSH_INSN, "pushinsn"},
    {OpCode::PUSH_INSN_IMM, "pushinsnimm"},
    {OpCode::SIDELOAD, "sideload"},

    {OpCode::NEW_BUFFER, "newbuffer"},
    {OpCode::GET_BUFFER8, "getbuffer8"},
    {OpCode::GET_BUFFER64, "getbuffer64"},
    {OpCode::GET_BUFFER256, "getbuffer256"},
    {OpCode::SET_BUFFER8, "setbuffer8"},
    {OpCode::SET_BUFFER64, "setbuffer64"},
    {OpCode::SET_BUFFER256, "setbuffer256"},

    {OpCode::ECRECOVER, "ecrecover"},
    {OpCode::ECADD, "ecadd"},
    {OpCode::ECMUL, "ecmul"},
    {OpCode::ECPAIRING, "ecpairing"},
    {OpCode::DEBUG_PRINT, "debugprint"}};

// Each size_t is a tuple depth to marshal to
const std::unordered_map<OpCode, std::vector<size_t>> InstructionStackPops = {
    {static_cast<OpCode>(0), {}},
    {OpCode::ADD, {1, 1}},
    {OpCode::MUL, {1, 1}},
    {OpCode::SUB, {1, 1}},
    {OpCode::DIV, {1, 1}},
    {OpCode::SDIV, {1, 1}},
    {OpCode::MOD, {1, 1}},
    {OpCode::SMOD, {1, 1}},
    {OpCode::ADDMOD, {1, 1, 1}},
    {OpCode::MULMOD, {1, 1, 1}},
    {OpCode::EXP, {1, 1}},
    {OpCode::SIGNEXTEND, {1, 1}},

    {OpCode::LT, {1, 1}},
    {OpCode::GT, {1, 1}},
    {OpCode::SLT, {1, 1}},
    {OpCode::SGT, {1, 1}},
    {OpCode::EQ, {0, 0}},
    {OpCode::ISZERO, {1}},
    {OpCode::BITWISE_AND, {1, 1}},
    {OpCode::BITWISE_OR, {1, 1}},
    {OpCode::BITWISE_XOR, {1, 1}},
    {OpCode::BITWISE_NOT, {1}},
    {OpCode::BYTE, {1, 1}},
    {OpCode::SHL, {1, 1}},
    {OpCode::SHR, {1, 1}},
    {OpCode::SAR, {1, 1}},

    {OpCode::HASH, {0}},
    {OpCode::TYPE, {1}},
    {OpCode::ETHHASH2, {1, 1}},
    {OpCode::KECCAKF, {1}},
    {OpCode::SHA256F, {1, 1, 1}},

    {OpCode::POP, {0}},
    {OpCode::SPUSH, {}},
    {OpCode::RPUSH, {}},
    {OpCode::RSET, {0}},
    {OpCode::JUMP, {0}},
    {OpCode::CJUMP, {1, 1}},
    {OpCode::STACKEMPTY, {}},
    {OpCode::PCPUSH, {}},
    {OpCode::AUXPUSH, {0}},
    {OpCode::AUXPOP, {}},
    {OpCode::AUXSTACKEMPTY, {}},
    {OpCode::NOP, {}},
    {OpCode::ERRPUSH, {}},
    {OpCode::ERRSET, {1}},

    {OpCode::DUP0, {0}},
    {OpCode::DUP1, {0, 0}},
    {OpCode::DUP2, {0, 0, 0}},
    {OpCode::SWAP1, {0, 0}},
    {OpCode::SWAP2, {0, 0, 0}},

    {OpCode::TGET, {1, 1}},
    {OpCode::TSET, {1, 1, 0}},
    {OpCode::TLEN, {1}},
    {OpCode::XGET, {1}},
    {OpCode::XSET, {1, 0}},

    {OpCode::BREAKPOINT, {}},
    {OpCode::LOG, {0}},

    {OpCode::SEND, {1, 1}},
    {OpCode::INBOX, {}},
    {OpCode::ERROR, {}},
    {OpCode::HALT, {}},
    {OpCode::SET_GAS, {1}},
    {OpCode::PUSH_GAS, {}},
    {OpCode::ERR_CODE_POINT, {}},
    {OpCode::PUSH_INSN, {1, 1}},
    {OpCode::PUSH_INSN_IMM, {1, 0, 1}},
    {OpCode::SIDELOAD, {1}},
    {OpCode::DEBUG_PRINT, {0}},

    {OpCode::NEW_BUFFER, {}},
    {OpCode::GET_BUFFER8, {1, 1}},
    {OpCode::GET_BUFFER64, {1, 1}},
    {OpCode::GET_BUFFER256, {1, 1}},
    {OpCode::SET_BUFFER8, {1, 1, 1}},
    {OpCode::SET_BUFFER64, {1, 1, 1}},
    {OpCode::SET_BUFFER256, {1, 1, 1}},

    {OpCode::ECRECOVER, {1, 1, 1, 1}},
    {OpCode::ECADD, {1, 1, 1, 1}},
    {OpCode::ECMUL, {1, 1, 1}},
    {OpCode::ECPAIRING, {32}}};

const std::unordered_map<OpCode, std::vector<size_t>> InstructionAuxStackPops =
    {{static_cast<OpCode>(0), {}},
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
     {OpCode::SIGNEXTEND, {}},

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
     {OpCode::SHL, {}},
     {OpCode::SHR, {}},
     {OpCode::SAR, {}},

     {OpCode::HASH, {}},
     {OpCode::TYPE, {}},
     {OpCode::ETHHASH2, {}},
     {OpCode::KECCAKF, {}},
     {OpCode::SHA256F, {}},

     {OpCode::POP, {}},
     {OpCode::SPUSH, {}},
     {OpCode::RPUSH, {}},
     {OpCode::RSET, {}},
     {OpCode::JUMP, {}},
     {OpCode::CJUMP, {}},
     {OpCode::STACKEMPTY, {}},
     {OpCode::PCPUSH, {}},
     {OpCode::AUXPUSH, {}},
     {OpCode::AUXPOP, {0}},
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
     {OpCode::XGET, {1}},
     {OpCode::XSET, {1}},

     {OpCode::BREAKPOINT, {}},
     {OpCode::LOG, {}},

     {OpCode::SEND, {}},
     {OpCode::INBOX, {}},
     {OpCode::ERROR, {}},
     {OpCode::HALT, {}},
     {OpCode::SET_GAS, {}},
     {OpCode::PUSH_GAS, {}},
     {OpCode::ERR_CODE_POINT, {}},
     {OpCode::PUSH_INSN, {}},
     {OpCode::PUSH_INSN_IMM, {}},
     {OpCode::SIDELOAD, {}},
     {OpCode::DEBUG_PRINT, {}},

     {OpCode::NEW_BUFFER, {}},
     {OpCode::GET_BUFFER8, {}},
     {OpCode::GET_BUFFER64, {}},
     {OpCode::GET_BUFFER256, {}},
     {OpCode::SET_BUFFER8, {}},
     {OpCode::SET_BUFFER64, {}},
     {OpCode::SET_BUFFER256, {}},
     {OpCode::ECRECOVER, {}},
     {OpCode::ECADD, {}},
     {OpCode::ECMUL, {}},
     {OpCode::ECPAIRING, {}}};

const std::unordered_map<OpCode, uint64_t> InstructionArbGasCost = {
    {OpCode::ADD, 3},
    {OpCode::MUL, 3},
    {OpCode::SUB, 3},
    {OpCode::DIV, 4},
    {OpCode::SDIV, 7},
    {OpCode::MOD, 4},
    {OpCode::SMOD, 7},
    {OpCode::ADDMOD, 4},
    {OpCode::MULMOD, 4},
    {OpCode::EXP, 25},
    {OpCode::SIGNEXTEND, 7},

    {OpCode::LT, 2},
    {OpCode::GT, 2},
    {OpCode::SLT, 2},
    {OpCode::SGT, 2},
    {OpCode::EQ, 2},
    {OpCode::ISZERO, 1},
    {OpCode::BITWISE_AND, 2},
    {OpCode::BITWISE_OR, 2},
    {OpCode::BITWISE_XOR, 2},
    {OpCode::BITWISE_NOT, 1},
    {OpCode::BYTE, 4},
    {OpCode::SHL, 4},
    {OpCode::SHR, 4},
    {OpCode::SAR, 4},

    {OpCode::HASH, 7},
    {OpCode::TYPE, 3},
    {OpCode::ETHHASH2, 8},
    {OpCode::KECCAKF, 600},
    {OpCode::SHA256F, 250},

    {OpCode::POP, 1},
    {OpCode::SPUSH, 1},
    {OpCode::RPUSH, 1},
    {OpCode::RSET, 2},
    {OpCode::JUMP, 4},
    {OpCode::CJUMP, 4},
    {OpCode::STACKEMPTY, 2},
    {OpCode::PCPUSH, 1},
    {OpCode::AUXPUSH, 1},
    {OpCode::AUXPOP, 1},
    {OpCode::AUXSTACKEMPTY, 2},
    {OpCode::NOP, 1},
    {OpCode::ERRPUSH, 1},
    {OpCode::ERRSET, 1},

    {OpCode::DUP0, 1},
    {OpCode::DUP1, 1},
    {OpCode::DUP2, 1},
    {OpCode::SWAP1, 1},
    {OpCode::SWAP2, 1},

    {OpCode::TGET, 2},
    {OpCode::TSET, 40},
    {OpCode::TLEN, 2},
    {OpCode::XGET, 3},
    {OpCode::XSET, 41},

    {OpCode::BREAKPOINT, 100},
    {OpCode::LOG, 100},

    {OpCode::SEND, 100},
    {OpCode::INBOX, 40},
    {OpCode::ERROR, 5},
    {OpCode::HALT, 10},
    {OpCode::SET_GAS, 1},
    {OpCode::PUSH_GAS, 1},
    {OpCode::ERR_CODE_POINT, 25},
    {OpCode::PUSH_INSN, 25},
    {OpCode::PUSH_INSN_IMM, 25},
    {OpCode::SIDELOAD, 10},
    {OpCode::DEBUG_PRINT, 1},

    {OpCode::NEW_BUFFER, 1},
    {OpCode::GET_BUFFER8, 10},
    {OpCode::GET_BUFFER64, 10},
    {OpCode::GET_BUFFER256, 10},
    {OpCode::SET_BUFFER8, 100},
    {OpCode::SET_BUFFER64, 100},
    {OpCode::SET_BUFFER256, 100},

    {OpCode::ECRECOVER, 20000},
    {OpCode::ECADD, 3500},
    {OpCode::ECMUL, 82000},
    {OpCode::ECPAIRING, 1000}};

constexpr size_t MaxValidOpcode =
    static_cast<size_t>(std::numeric_limits<uint8_t>::max());

template <typename T>
using OpCodeArray = std::array<T, MaxValidOpcode>;

inline OpCodeArray<uint64_t> initializeGasCosts() {
    OpCodeArray<uint64_t> arr;
    for (size_t i = 0; i <= MaxValidOpcode; ++i) {
        auto it = InstructionArbGasCost.find(static_cast<OpCode>(i));
        if (it != InstructionArbGasCost.end()) {
            arr[i] = it->second;
        } else {
            arr[i] = 0;
        }
    }
    return arr;
}

inline const OpCodeArray<uint64_t>& instructionGasCosts() {
    static OpCodeArray<uint64_t> costs = initializeGasCosts();
    return costs;
}

inline OpCodeArray<bool> initializeValidity() {
    OpCodeArray<bool> arr;
    for (size_t i = 0; i <= MaxValidOpcode; ++i) {
        auto it = InstructionArbGasCost.find(static_cast<OpCode>(i));
        if (it != InstructionArbGasCost.end()) {
            arr[i] = true;
        } else {
            arr[i] = false;
        }
    }
    return arr;
}

inline const OpCodeArray<bool>& instructionValidity() {
    static OpCodeArray<bool> costs = initializeValidity();
    return costs;
}

#endif /* opcodes_hpp */
