//
//  main.cpp
//  tests
//
//  Created by Harry Kalodner on 6/23/19.
//

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
        REQUIRE(m.state == ERROR);
    }
}

TEST_CASE("SDIV opcode is correct") {
    SECTION("Positive divided by negative") {
        testBinaryOp(12, -3, -4, OpCode::SDIV);
    }

    SECTION("Negative divided by negative") {
        testBinaryOp(-12, -3, 4, OpCode::SDIV);
    }
}

TEST_CASE("SMOD opcode is correct") {
    SECTION("Positive mod negative") {
        testBinaryOp(8, -3, 2, OpCode::SMOD);
    }

    SECTION("Negative mod positive") {
        testBinaryOp(-8, 3, -2, OpCode::SMOD);
    }

    SECTION("Negative mod negative") {
        testBinaryOp(-8, -3, -2, OpCode::SMOD);
    }
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
