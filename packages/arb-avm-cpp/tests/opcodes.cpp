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

#include "bigint_utils.hpp"

#include <avm/machine.hpp>

#include <catch2/catch.hpp>

MachineState runUnaryOp(uint256_t arg1, OpCode op) {
    MachineState m;
    m.stack.push(arg1);
    m.runOp(op);
    return m;
}

void testUnaryOp(uint256_t arg1, uint256_t result, OpCode op) {
    MachineState m = runUnaryOp(arg1, op);
    value res = m.stack.pop();
    auto num = nonstd::get_if<uint256_t>(&res);
    REQUIRE(num);
    REQUIRE(*num == result);
    REQUIRE(m.stack.stacksize() == 0);
}

MachineState runBinaryOp(uint256_t arg1, uint256_t arg2, OpCode op) {
    MachineState m;
    m.stack.push(arg2);
    m.stack.push(arg1);
    m.runOp(op);
    return m;
}

void testBinaryOp(uint256_t arg1,
                  uint256_t arg2,
                  uint256_t expected,
                  OpCode op) {
    MachineState m = runBinaryOp(arg1, arg2, op);
    value res = m.stack.pop();
    auto actual = nonstd::get_if<uint256_t>(&res);
    REQUIRE(actual);
    REQUIRE(*actual == expected);
    REQUIRE(m.stack.stacksize() == 0);
}

MachineState runTertiaryOp(uint256_t arg1,
                           uint256_t arg2,
                           uint256_t arg3,
                           OpCode op) {
    MachineState m;
    m.stack.push(arg3);
    m.stack.push(arg2);
    m.stack.push(arg1);
    m.runOp(op);
    return m;
}

void testTertiaryOp(uint256_t arg1,
                    uint256_t arg2,
                    uint256_t arg3,
                    uint256_t result,
                    OpCode op) {
    MachineState m = runTertiaryOp(arg1, arg2, arg3, op);
    value res = m.stack.pop();
    auto num = nonstd::get_if<uint256_t>(&res);
    REQUIRE(num);
    REQUIRE(*num == result);
    REQUIRE(m.stack.stacksize() == 0);
}

TEST_CASE("ADD opcode is correct") {
    SECTION("Non overlow is correct") { testBinaryOp(4, 3, 7, OpCode::ADD); }

    SECTION("0+0 is correct") { testBinaryOp(0, 0, 0, OpCode::ADD); }

    SECTION("Overlow is correct") { testBinaryOp(-1, 4, 3, OpCode::ADD); }

    SECTION("-2+1 is correct") { testBinaryOp(-2, 1, -1, OpCode::ADD); }
}

TEST_CASE("MUL opcode is correct") {
    SECTION("Non overlow is correct") { testBinaryOp(4, 3, 12, OpCode::MUL); }

    SECTION("3*0 is correct") { testBinaryOp(3, 0, 0, OpCode::MUL); }

    SECTION("-1*1 is correct") { testBinaryOp(-1, 1, -1, OpCode::MUL); }

    SECTION("-2+1 is correct") { testBinaryOp(-2, 1, -2, OpCode::MUL); }
}

TEST_CASE("SUB opcode is correct") {
    SECTION("Non overlow is correct") { testBinaryOp(4, 3, 1, OpCode::SUB); }

    SECTION("Overlow is correct") { testBinaryOp(3, 4, -1, OpCode::SUB); }
}

TEST_CASE("DIV opcode is correct") {
    SECTION("Non overlow is correct") { testBinaryOp(12, 3, 4, OpCode::DIV); }

    SECTION("unsigned division is correct") {
        MachineState m = runBinaryOp(-6, 2, OpCode::DIV);
        value res = m.stack.pop();
        auto num = nonstd::get_if<uint256_t>(&res);
        REQUIRE(num);
        REQUIRE(*num != 3);
        REQUIRE(m.stack.stacksize() == 0);
    }

    SECTION("Divide by zero") {
        MachineState m = runBinaryOp(3, 0, OpCode::DIV);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("SDIV opcode is correct") {
    SECTION("Positive divided by positive") {
        testBinaryOp(12, 3, 4, OpCode::SDIV);
    }
    SECTION("Positive divided by negative") {
        testBinaryOp(12, -3, -4, OpCode::SDIV);
    }
    SECTION("Negative divided by positive") {
        testBinaryOp(-12, 3, -4, OpCode::SDIV);
    }
    SECTION("Negative divided by negative") {
        testBinaryOp(-12, -3, 4, OpCode::SDIV);
    }
    SECTION("Divide by zero") {
        MachineState m = runBinaryOp(3, 0, OpCode::SDIV);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("MOD opcode is correct") {
    SECTION("8 mod 3") { testBinaryOp(8, 3, 2, OpCode::MOD); }

    SECTION("8 mod very large") { testBinaryOp(8, -3, 8, OpCode::MOD); }

    SECTION("0 mod 3") { testBinaryOp(0, 3, 0, OpCode::MOD); }

    SECTION("Mod by zero") {
        MachineState m = runBinaryOp(3, 0, OpCode::MOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("SMOD opcode is correct") {
    SECTION("Positive mod positive") { testBinaryOp(8, 3, 2, OpCode::SMOD); }

    SECTION("Positive mod negative") { testBinaryOp(8, -3, 2, OpCode::SMOD); }

    SECTION("Negative mod positive") { testBinaryOp(-8, 3, -2, OpCode::SMOD); }

    SECTION("Negative mod negative") { testBinaryOp(-8, -3, -2, OpCode::SMOD); }
    SECTION("Mod by zero") {
        MachineState m = runBinaryOp(3, 0, OpCode::SMOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("ADDMOD opcode is correct") {
    SECTION("(8+5)%3") { testTertiaryOp(8, 5, 3, 1, OpCode::ADDMOD); }

    SECTION("(-1+1)%7") { testTertiaryOp(-1, 1, 7, 2, OpCode::ADDMOD); }

    SECTION("(0+0)%7") { testTertiaryOp(0, 0, 7, 0, OpCode::ADDMOD); }

    SECTION("Mod by zero") {
        MachineState m = runTertiaryOp(8, 3, 0, OpCode::ADDMOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("MULMOD opcode is correct") {
    SECTION("(8*2)%3") { testTertiaryOp(8, 2, 3, 1, OpCode::MULMOD); }

    SECTION("(-1*2)%7") { testTertiaryOp(-1, 2, 7, 2, OpCode::MULMOD); }

    SECTION("(0*0)%7") { testTertiaryOp(0, 0, 7, 0, OpCode::MULMOD); }

    SECTION("Mod by zero") {
        MachineState m = runTertiaryOp(8, 3, 0, OpCode::MULMOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("EXP opcode is correct") {
    SECTION("All positive") { testBinaryOp(3, 2, 9, OpCode::EXP); }
    SECTION("wrap") { testBinaryOp(2, 256, 0, OpCode::EXP); }
}

TEST_CASE("LT opcode is correct") {
    SECTION("less") { testBinaryOp(3, 9, 1, OpCode::LT); }
    SECTION("greater") { testBinaryOp(9, 3, 0, OpCode::LT); }
    SECTION("equal") { testBinaryOp(3, 3, 0, OpCode::LT); }
    SECTION("First neg, second pos") { testBinaryOp(-3, 9, 0, OpCode::LT); }
}

TEST_CASE("GT opcode is correct") {
    SECTION("less") { testBinaryOp(3, 9, 0, OpCode::GT); }
    SECTION("greater") { testBinaryOp(9, 3, 1, OpCode::GT); }
    SECTION("equal") { testBinaryOp(3, 3, 0, OpCode::GT); }
    SECTION("First neg, second pos") { testBinaryOp(-3, 9, 1, OpCode::GT); }
}

TEST_CASE("SLT opcode is correct") {
    SECTION("All positive true") { testBinaryOp(7, 3, 0, OpCode::SLT); }
    SECTION("All positive false") { testBinaryOp(3, 7, 1, OpCode::SLT); }
    SECTION("All negative true") { testBinaryOp(-3, -7, 0, OpCode::SLT); }
    SECTION("All negative false") { testBinaryOp(-7, -3, 1, OpCode::SLT); }
    SECTION("First pos, second neg true") {
        testBinaryOp(3, -7, 0, OpCode::SLT);
    }
    SECTION("First neg, second pos false") {
        testBinaryOp(-3, 7, 1, OpCode::SLT);
    }
    SECTION("equal") { testBinaryOp(3, 3, 0, OpCode::SLT); }
}

TEST_CASE("SGT opcode is correct") {
    SECTION("All positive true") { testBinaryOp(7, 3, 1, OpCode::SGT); }
    SECTION("All positive false") { testBinaryOp(3, 7, 0, OpCode::SGT); }
    SECTION("All negative true") { testBinaryOp(-3, -7, 1, OpCode::SGT); }
    SECTION("All negative false") { testBinaryOp(-7, -3, 0, OpCode::SGT); }
    SECTION("First pos, second neg true") {
        testBinaryOp(3, -7, 1, OpCode::SGT);
    }
    SECTION("First neg, second pos false") {
        testBinaryOp(-7, 3, 0, OpCode::SGT);
    }
    SECTION("equal") { testBinaryOp(3, 3, 0, OpCode::SGT); }
}

TEST_CASE("EQ opcode is correct") {
    SECTION("Not equal") { testBinaryOp(7, 3, 0, OpCode::EQ); }
    SECTION("equal") { testBinaryOp(3, 3, 1, OpCode::EQ); }
    SECTION("matching tuples") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        m.runOp(OpCode::EQ);
        value res = m.stack.pop();
        auto actual = nonstd::get_if<uint256_t>(&res);
        REQUIRE(actual);
        REQUIRE(*actual == 1);
        REQUIRE(m.stack.stacksize() == 0);
    }
    SECTION("different tuples") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        m.stack.push(Tuple{uint256_t{1}, uint256_t{3}, m.pool.get()});
        m.runOp(OpCode::EQ);
        value res = m.stack.pop();
        auto actual = nonstd::get_if<uint256_t>(&res);
        REQUIRE(actual);
        REQUIRE(*actual == 0);
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("ISZERO opcode is correct") {
    SECTION("true") { testUnaryOp(0, 1, OpCode::ISZERO); }
    SECTION("false") { testUnaryOp(2, 0, OpCode::ISZERO); }
}

TEST_CASE("AND opcode is correct") {
    SECTION("3 and 9 = 1") { testBinaryOp(3, 9, 1, OpCode::BITWISE_AND); }
    SECTION("3 and 3 = 3") { testBinaryOp(3, 3, 3, OpCode::BITWISE_AND); }
}

TEST_CASE("OR opcode is correct") {
    SECTION("3 or 9 = 11") { testBinaryOp(3, 9, 11, OpCode::BITWISE_OR); }
    SECTION("3 or 3 = 3") { testBinaryOp(3, 3, 3, OpCode::BITWISE_OR); }
}

TEST_CASE("XOR opcode is correct") {
    SECTION("3 or 9 = 11") { testBinaryOp(3, 9, 10, OpCode::BITWISE_XOR); }
    SECTION("3 or 3 = 3") { testBinaryOp(3, 3, 0, OpCode::BITWISE_XOR); }
}

TEST_CASE("NOT opcode is correct") {
    SECTION("0") { testUnaryOp(0, -1, OpCode::BITWISE_NOT); }
    SECTION("3") { testUnaryOp(3, -4, OpCode::BITWISE_NOT); }
    SECTION("-4") { testUnaryOp(-4, 3, OpCode::BITWISE_NOT); }
}

TEST_CASE("BYTE opcode is correct") {
    SECTION("31st byte of 16 = 16") { testBinaryOp(16, 31, 16, OpCode::BYTE); }
    SECTION("3rd byte of 3 = 0") { testBinaryOp(3, 3, 0, OpCode::BYTE); }
}

TEST_CASE("SIGNEXTEND opcode is correct") {
    SECTION("test1") { testBinaryOp(-1, 0, -1, OpCode::SIGNEXTEND); }
    SECTION("test2") { testBinaryOp(1, 0, 1, OpCode::SIGNEXTEND); }
    SECTION("test3") { testBinaryOp(127, 0, 127, OpCode::SIGNEXTEND); }
    SECTION("test4") { testBinaryOp(128, 0, -128, OpCode::SIGNEXTEND); }
    SECTION("test5") { testBinaryOp(254, 0, -2, OpCode::SIGNEXTEND); }
    SECTION("test6") { testBinaryOp(257, 0, 1, OpCode::SIGNEXTEND); }
    SECTION("test7") { testBinaryOp(65534, 1, -2, OpCode::SIGNEXTEND); }
    SECTION("test8") { testBinaryOp(65537, 1, 1, OpCode::SIGNEXTEND); }
    SECTION("test8") { testBinaryOp(65537, 2, 65537, OpCode::SIGNEXTEND); }
}

TEST_CASE("HASH opcode is correct") {
    SECTION("10") {
        testUnaryOp(10,
                    from_hex_str("c65a7bb8d6351c1cf70c95a316cc6a92839c986682d98"
                                 "bc35f958f4883f9d2a8"),
                    OpCode::HASH);
    }
}

TEST_CASE("TYPE opcode is correct") {
    SECTION("type int") {
        MachineState m;
        m.stack.push(value{uint256_t(3)});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::TYPE);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(0)});
        REQUIRE(m.stack.stacksize() == 0);
    }
    SECTION("type codepoint") {
        MachineState m;
        m.stack.push(value{CodePoint()});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::TYPE);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(1)});
        REQUIRE(m.stack.stacksize() == 0);
    }
    SECTION("type tuple") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::TYPE);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("POP opcode is correct") {
    SECTION("pop") {
        MachineState m;
        m.stack.push(uint256_t{3});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::POP);
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("SPUSH opcode is correct") {
    SECTION("pop") {
        MachineState m;
        m.staticVal = uint256_t(5);
        m.runOp(OpCode::SPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("RPUSH opcode is correct") {
    SECTION("pop") {
        MachineState m;
        m.registerVal = uint256_t(5);
        m.runOp(OpCode::RPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("RSET opcode is correct") {
    SECTION("pop") {
        MachineState m;
        m.stack.push(value{uint256_t(5)});
        m.runOp(OpCode::RSET);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.registerVal == value{uint256_t(5)});
    }
}

TEST_CASE("JUMP opcode is correct") {
    SECTION("jump") {
        MachineState m;
        m.stack.push(value{CodePoint(2, OpCode::ADD, 0)});
        m.runOp(OpCode::JUMP);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == 2);
    }
}

TEST_CASE("CJUMP opcode is correct") {
    SECTION("cjump true") {
        MachineState m;
        m.pc = 3;
        m.stack.push(uint256_t{0});
        m.stack.push(value{CodePoint(2, OpCode::ADD, 0)});
        m.runOp(OpCode::CJUMP);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == 4);
    }
    SECTION("cjump false") {
        MachineState m;
        m.pc = 3;
        m.stack.push(uint256_t{1});
        m.stack.push(value{CodePoint(2, OpCode::ADD, 0)});
        m.runOp(OpCode::CJUMP);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == 2);
    }
}

TEST_CASE("STACKEMPTY opcode is correct") {
    SECTION("empty") {
        MachineState m;
        m.runOp(OpCode::STACKEMPTY);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(1)});
        REQUIRE(m.stack.stacksize() == 0);
    }
    SECTION("not empty") {
        MachineState m;
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::STACKEMPTY);
        REQUIRE(m.stack.stacksize() == 2);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(0)});
        REQUIRE(m.stack.stacksize() == 1);
    }
}

TEST_CASE("PCPUSH opcode is correct") {
    SECTION("pcpush") {
        MachineState m;
        m.code.push_back(CodePoint(0, OpCode::ADD, 0));
        m.runOp(OpCode::PCPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        REQUIRE(m.pc == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{CodePoint(0, OpCode::ADD, 0)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("AUXPUSH opcode is correct") {
    SECTION("auxpush") {
        MachineState m;
        m.stack.push(value{uint256_t(5)});
        m.runOp(OpCode::AUXPUSH);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.auxstack.stacksize() == 1);
        value res = m.auxstack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("AUXPOP opcode is correct") {
    SECTION("auxpush") {
        MachineState m;
        m.auxstack.push(value{uint256_t(5)});
        m.runOp(OpCode::AUXPOP);
        REQUIRE(m.auxstack.stacksize() == 0);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("AUXSTACKEMPTY opcode is correct") {
    SECTION("empty") {
        MachineState m;
        m.runOp(OpCode::AUXSTACKEMPTY);
        REQUIRE(m.auxstack.stacksize() == 0);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(1)});
    }
    SECTION("not empty") {
        MachineState m;
        m.auxstack.push(value{uint256_t(5)});
        m.runOp(OpCode::AUXSTACKEMPTY);
        REQUIRE(m.auxstack.stacksize() == 1);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(0)});
    }
}

TEST_CASE("NOP opcode is correct") {
    SECTION("nop") {
        MachineState m;
        m.runOp(OpCode::NOP);
        REQUIRE(m.auxstack.stacksize() == 0);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == 1);
    }
}

TEST_CASE("ERRPUSH opcode is correct") {
    SECTION("errpush") {
        MachineState m;
        m.errpc = CodePoint(0, OpCode::ADD, 0);
        m.runOp(OpCode::ERRPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        REQUIRE(m.pc == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{CodePoint(0, OpCode::ADD, 0)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("ERRSET opcode is correct") {
    SECTION("errpush") {
        MachineState m;
        m.stack.push(value{CodePoint(0, OpCode::ADD, 0)});
        m.runOp(OpCode::ERRSET);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == 1);
        REQUIRE(m.errpc == CodePoint(0, OpCode::ADD, 0));
    }
}

TEST_CASE("DUP0 opcode is correct") {
    SECTION("dup") {
        MachineState m;
        m.stack.push(uint256_t{5});
        m.runOp(OpCode::DUP0);
        REQUIRE(m.stack.stacksize() == 2);
        REQUIRE(m.pc == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("DUP1 opcode is correct") {
    SECTION("dup") {
        MachineState m;
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::DUP1);
        REQUIRE(m.stack.stacksize() == 3);
        REQUIRE(m.pc == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("DUP2 opcode is correct") {
    SECTION("dup") {
        MachineState m;
        m.stack.push(uint256_t{7});
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::DUP2);
        REQUIRE(m.stack.stacksize() == 4);
        REQUIRE(m.pc == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(7)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(7)});
    }
}

TEST_CASE("SWAP1 opcode is correct") {
    SECTION("swap") {
        MachineState m;
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::SWAP1);
        REQUIRE(m.stack.stacksize() == 2);
        REQUIRE(m.pc == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
    }
}

TEST_CASE("SWAP2 opcode is correct") {
    SECTION("dup") {
        MachineState m;
        m.stack.push(uint256_t{7});
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::SWAP2);
        REQUIRE(m.stack.stacksize() == 3);
        REQUIRE(m.pc == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(7)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
    }
}

TEST_CASE("TGET opcode is correct") {
    SECTION("tget") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7},
                           uint256_t{6}, m.pool.get()});
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::TGET);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(8)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("TSET opcode is correct") {
    SECTION("2 tup") {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}, m.pool.get()});
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::TSET);
        value res = m.stack.pop();
        REQUIRE(res == value{Tuple{uint256_t{1}, uint256_t{3}, m.pool.get()}});
        REQUIRE(m.stack.stacksize() == 0);
    }

    SECTION("8 tup") {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.stack.push(Tuple{uint256_t{9}, uint256_t{9}, uint256_t{9},
                           uint256_t{9}, uint256_t{9}, uint256_t{9},
                           uint256_t{9}, uint256_t{9}, m.pool.get()});
        m.stack.push(uint256_t{7});
        m.runOp(OpCode::TSET);
        value res = m.stack.pop();
        REQUIRE(res == value{Tuple{uint256_t{9}, uint256_t{9}, uint256_t{9},
                                   uint256_t{9}, uint256_t{9}, uint256_t{9},
                                   uint256_t{9}, uint256_t{3}, m.pool.get()}});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("TLEN opcode is correct") {
    SECTION("tlen") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7},
                           uint256_t{6}, m.pool.get()});
        m.runOp(OpCode::TLEN);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(4)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("BREAKPOINT opcode is correct") {
    SECTION("break") {
        MachineState m;
        auto blockReason = m.runOp(OpCode::BREAKPOINT);
        REQUIRE(m.state == Status::Extensive);
        REQUIRE(nonstd::get_if<BreakpointBlocked>(&blockReason));
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("LOG opcode is correct") {
    SECTION("log") {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::LOG);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.context.logs[0] == value{uint256_t(3)});
    }
}

TEST_CASE("SEND opcode is correct") {
    SECTION("send") {
        // TODO: fill in send test
    }
}

TEST_CASE("NBSEND opcode is correct") {
    SECTION("nbsend") {
        // TODO: fill in nbsend test
    }
}

TEST_CASE("GETTIME opcode is correct") {
    SECTION("time") {
        // TODO: fill in gettime test
    }
}

TEST_CASE("INBOX opcode is correct") {
    SECTION("inbox") {
        // TODO: fill in inbox test
    }
}

TEST_CASE("ERROR opcode is correct") {
    SECTION("error") {
        // TODO: fill in error test
    }
}

TEST_CASE("HALT opcode is correct") {
    SECTION("halt") {
        // TODO: fill in halt test
    }
}
