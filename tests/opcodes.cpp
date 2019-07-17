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

MachineState runBinaryOp(uint256_t arg1, uint256_t arg2, OpCode op) {
    MachineState m;
    m.stack.push(arg2);
    m.stack.push(arg1);
    m.runOp(op);
    return m;
}

void testBinaryOp(uint256_t arg1, uint256_t arg2, uint256_t result, OpCode op) {
    MachineState m = runBinaryOp(arg1, arg2, op);
    value res = m.stack.pop();
    auto num = mpark::get_if<uint256_t>(&res);
    REQUIRE(num);
    REQUIRE(*num == result);
    REQUIRE(m.stack.stacksize() == 0);
}

TEST_CASE("SUB opcode is correct") {
    SECTION("Non overlow is correct") { testBinaryOp(4, 3, 1, OpCode::SUB); }

    SECTION("Overlow is correct") { testBinaryOp(3, 4, -1, OpCode::SUB); }
}

TEST_CASE("DIV opcode is correct") {
    SECTION("Non overlow is correct") { testBinaryOp(12, 3, 4, OpCode::DIV); }

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
}

TEST_CASE("SMOD opcode is correct") {
    SECTION("Positive mod positive") { testBinaryOp(8, 3, 2, OpCode::SMOD); }
    
    SECTION("Positive mod negative") { testBinaryOp(8, -3, 2, OpCode::SMOD); }

    SECTION("Negative mod positive") { testBinaryOp(-8, 3, -2, OpCode::SMOD); }

    SECTION("Negative mod negative") { testBinaryOp(-8, -3, -2, OpCode::SMOD); }
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
