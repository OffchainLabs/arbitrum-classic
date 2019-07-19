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
    auto num = mpark::get_if<uint256_t>(&res);
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

void testBinaryOp(uint256_t arg1, uint256_t arg2, uint256_t expected, OpCode op) {
    MachineState m = runBinaryOp(arg1, arg2, op);
    value res = m.stack.pop();
    auto actual = mpark::get_if<uint256_t>(&res);
    REQUIRE(actual);
    REQUIRE(*actual == expected);
    REQUIRE(m.stack.stacksize() == 0);
}

MachineState runTertiaryOp(uint256_t arg1, uint256_t arg2, uint256_t arg3, OpCode op) {
    MachineState m;
    m.stack.push(arg3);
    m.stack.push(arg2);
    m.stack.push(arg1);
    m.runOp(op);
    return m;
}

void testTertiaryOp(uint256_t arg1, uint256_t arg2, uint256_t arg3, uint256_t result, OpCode op) {
    MachineState m = runTertiaryOp(arg1, arg2, arg3, op);
    value res = m.stack.pop();
    auto num = mpark::get_if<uint256_t>(&res);
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
        auto num = mpark::get_if<uint256_t>(&res);
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

TEST_CASE("LT opcode is correct") {
    SECTION("less") { testBinaryOp(3, 9, 1, OpCode::LT); }
    SECTION("greater") { testBinaryOp(9, 3, 0, OpCode::LT); }
    SECTION("equal") { testBinaryOp(3, 3, 0	, OpCode::LT); }
    SECTION("First neg, second pos") { testBinaryOp(-3, 9, 0, OpCode::LT); }
}

TEST_CASE("SLT opcode is correct") {
    SECTION("All positive") { testBinaryOp(7, 3, 0, OpCode::SLT); }
    SECTION("All negative") { testBinaryOp(-7, -3, 1, OpCode::SLT); }
    SECTION("First pos, second neg") { testBinaryOp(-7, 3, 1, OpCode::SLT); }
    SECTION("First neg, second pos") { testBinaryOp(7, -3, 0, OpCode::SLT); }
}

TEST_CASE("SGT opcode is correct") {
    SECTION("All positive") { testBinaryOp(7, 3, 1, OpCode::SGT); }
    SECTION("All negative") { testBinaryOp(-7, -3, 0, OpCode::SGT); }
    SECTION("First pos, second neg") { testBinaryOp(-7, 3, 0, OpCode::SGT); }
    SECTION("First neg, second pos") { testBinaryOp(7, -3, 1, OpCode::SGT); }
}

TEST_CASE("TSET opcode is correct") {
    SECTION("2 tup") {
        MachineState m;
        m.stack.push(3);
        m.stack.push(Tuple{1, 2, m.pool.get()});
        m.stack.push(1);
        m.runOp(OpCode::TSET);
        value res = m.stack.pop();
        REQUIRE(res == value{Tuple{1, 3, m.pool.get()}});
        REQUIRE(m.stack.stacksize() == 0);
    }

    SECTION("8 tup") {
        MachineState m;
        m.stack.push(3);
        m.stack.push(Tuple{9, 9, 9, 9, 9, 9, 9, 9, m.pool.get()});
        m.stack.push(7);
        m.runOp(OpCode::TSET);
        value res = m.stack.pop();
        REQUIRE(res == value{Tuple{9, 9, 9, 9, 9, 9, 9, 3, m.pool.get()}});
        REQUIRE(m.stack.stacksize() == 0);
    }
}
