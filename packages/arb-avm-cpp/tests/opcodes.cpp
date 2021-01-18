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

#include "ecops.hpp"

#include <avm/machine.hpp>
#include <avm/machinestate/machineoperation.hpp>

#include <secp256k1_recovery.h>
#include <ethash/keccak.hpp>

#define CATCH_CONFIG_ENABLE_BENCHMARKING 1
#include <catch2/catch.hpp>

#include <boost/algorithm/hex.hpp>

#include <iostream>

using namespace intx;

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

TEST_CASE("OPCODE: ADD opcode is correct") {
    SECTION("Non overlow is correct") {
        testBinaryOp(4_u256, 3_u256, 7_u256, OpCode::ADD);
    }

    SECTION("0+0 is correct") {
        testBinaryOp(0_u256, 0_u256, 0_u256, OpCode::ADD);
    }

    SECTION("Overlow is correct") {
        testBinaryOp(-1_u256, 4_u256, 3_u256, OpCode::ADD);
    }

    SECTION("-2+1 is correct") {
        testBinaryOp(-2_u256, 1_u256, -1_u256, OpCode::ADD);
    }
}

TEST_CASE("OPCODE: MUL opcode is correct") {
    SECTION("Non overlow is correct") {
        testBinaryOp(4_u256, 3_u256, 12_u256, OpCode::MUL);
    }

    SECTION("3*0 is correct") {
        testBinaryOp(3_u256, 0_u256, 0_u256, OpCode::MUL);
    }

    SECTION("-1*1 is correct") {
        testBinaryOp(-1_u256, 1_u256, -1_u256, OpCode::MUL);
    }

    SECTION("-2+1 is correct") {
        testBinaryOp(-2_u256, 1_u256, -2_u256, OpCode::MUL);
    }
}

TEST_CASE("OPCODE: SUB opcode is correct") {
    SECTION("Non overlow is correct") {
        testBinaryOp(4_u256, 3_u256, 1_u256, OpCode::SUB);
    }

    SECTION("Overlow is correct") {
        testBinaryOp(3_u256, 4_u256, -1_u256, OpCode::SUB);
    }
}

TEST_CASE("OPCODE: DIV opcode is correct") {
    SECTION("Non overlow is correct") {
        testBinaryOp(12_u256, 3_u256, 4_u256, OpCode::DIV);
    }

    SECTION("unsigned division is correct") {
        MachineState m = runBinaryOp(-6_u256, 2_u256, OpCode::DIV);
        value res = m.stack.pop();
        auto num = nonstd::get_if<uint256_t>(&res);
        REQUIRE(num);
        REQUIRE(*num != 3);
        REQUIRE(m.stack.stacksize() == 0);
    }

    SECTION("Divide by zero") {
        MachineState m = runBinaryOp(3_u256, 0_u256, OpCode::DIV);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("SDIV opcode is correct") {
    SECTION("Positive divided by positive") {
        testBinaryOp(12_u256, 3_u256, 4_u256, OpCode::SDIV);
    }
    SECTION("Positive divided by negative") {
        testBinaryOp(12_u256, -3_u256, -4_u256, OpCode::SDIV);
    }
    SECTION("Negative divided by positive") {
        testBinaryOp(-12_u256, 3_u256, -4_u256, OpCode::SDIV);
    }
    SECTION("Negative divided by negative") {
        testBinaryOp(-12_u256, -3_u256, 4_u256, OpCode::SDIV);
    }
    SECTION("Divide by zero") {
        MachineState m = runBinaryOp(3_u256, 0_u256, OpCode::SDIV);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("OPCODE: MOD opcode is correct") {
    SECTION("8 mod 3") { testBinaryOp(8_u256, 3_u256, 2_u256, OpCode::MOD); }

    SECTION("8 mod very large") {
        testBinaryOp(8_u256, -3_u256, 8_u256, OpCode::MOD);
    }

    SECTION("0 mod 3") { testBinaryOp(0_u256, 3_u256, 0_u256, OpCode::MOD); }

    SECTION("Mod by zero") {
        MachineState m = runBinaryOp(3_u256, 0_u256, OpCode::MOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("OPCODE: SMOD opcode is correct") {
    SECTION("Positive mod positive") {
        testBinaryOp(8_u256, 3_u256, 2_u256, OpCode::SMOD);
    }

    SECTION("Positive mod negative") {
        testBinaryOp(8_u256, -3_u256, 2_u256, OpCode::SMOD);
    }

    SECTION("Negative mod positive") {
        testBinaryOp(-8_u256, 3_u256, -2_u256, OpCode::SMOD);
    }

    SECTION("Negative mod negative") {
        testBinaryOp(-8_u256, -3_u256, -2_u256, OpCode::SMOD);
    }
    SECTION("Mod by zero") {
        MachineState m = runBinaryOp(3, 0, OpCode::SMOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("OPCODE: ADDMOD opcode is correct") {
    SECTION("(8+5)%3") {
        testTertiaryOp(8_u256, 5_u256, 3_u256, 1_u256, OpCode::ADDMOD);
    }

    SECTION("(-1+1)%7") {
        testTertiaryOp(-1_u256, 1_u256, 7_u256, 2_u256, OpCode::ADDMOD);
    }

    SECTION("(0+0)%7") {
        testTertiaryOp(0_u256, 0_u256, 7_u256, 0_u256, OpCode::ADDMOD);
    }

    SECTION("(3+3)%-4") {
        testTertiaryOp(3_u256, 3_u256, -4_u256, 6_u256, OpCode::ADDMOD);
    }

    SECTION("Mod by zero") {
        MachineState m = runTertiaryOp(8_u256, 3_u256, 0_u256, OpCode::ADDMOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("OPCODE: MULMOD opcode is correct") {
    SECTION("(8*2)%3") {
        testTertiaryOp(8_u256, 2_u256, 3_u256, 1_u256, OpCode::MULMOD);
    }

    SECTION("(-1*2)%7") {
        testTertiaryOp(-1_u256, 2_u256, 7_u256, 2_u256, OpCode::MULMOD);
    }

    SECTION("(0*0)%7") {
        testTertiaryOp(0_u256, 0_u256, 7_u256, 0_u256, OpCode::MULMOD);
    }

    SECTION("Mod by zero") {
        MachineState m = runTertiaryOp(8_u256, 3_u256, 0_u256, OpCode::MULMOD);
        REQUIRE(m.state == Status::Error);
    }
}

TEST_CASE("OPCODE: EXP opcode is correct") {
    SECTION("All positive") {
        testBinaryOp(3_u256, 2_u256, 9_u256, OpCode::EXP);
    }
    SECTION("wrap") { testBinaryOp(2_u256, 256_u256, 0_u256, OpCode::EXP); }
}

TEST_CASE("OPCODE: LT opcode is correct") {
    SECTION("less") { testBinaryOp(3, 9, 1, OpCode::LT); }
    SECTION("greater") { testBinaryOp(9, 3, 0, OpCode::LT); }
    SECTION("equal") { testBinaryOp(3, 3, 0, OpCode::LT); }
    SECTION("First neg, second pos") { testBinaryOp(-3, 9, 0, OpCode::LT); }
}

TEST_CASE("OPCODE: GT opcode is correct") {
    testBinaryOp(3, 9, 0, OpCode::GT);
    testBinaryOp(9, 3, 1, OpCode::GT);
    testBinaryOp(3, 3, 0, OpCode::GT);
    testBinaryOp(-3, 9, 1, OpCode::GT);

    BENCHMARK_ADVANCED("gt 100x")(Catch::Benchmark::Chronometer meter) {
        MachineState sample_machine;
        for (int i = 0; i < 101; i++) {
            sample_machine.stack.push(uint256_t{100});
        }
        std::vector<MachineState> machines(meter.runs());
        std::fill(machines.begin(), machines.end(), sample_machine);
        meter.measure([&machines](int i) {
            auto& mach = machines[i];
            for (int j = 0; j < 100; j++) {
                mach.runOp(OpCode::GT);
            }
            return mach;
        });
    };
}

TEST_CASE("OPCODE: SLT opcode is correct") {
    SECTION("All positive true") {
        testBinaryOp(7_u256, 3_u256, 0_u256, OpCode::SLT);
    }
    SECTION("All positive false") {
        testBinaryOp(3_u256, 7_u256, 1_u256, OpCode::SLT);
    }
    SECTION("All negative true") {
        testBinaryOp(-3_u256, -7_u256, 0_u256, OpCode::SLT);
    }
    SECTION("All negative false") {
        testBinaryOp(-7_u256, -3_u256, 1_u256, OpCode::SLT);
    }
    SECTION("First pos, second neg true") {
        testBinaryOp(3_u256, -7_u256, 0_u256, OpCode::SLT);
    }
    SECTION("First neg, second pos false") {
        testBinaryOp(-3_u256, 7_u256, 1_u256, OpCode::SLT);
    }
    SECTION("equal") { testBinaryOp(3_u256, 3_u256, 0_u256, OpCode::SLT); }
}

TEST_CASE("OPCODE: SGT opcode is correct") {
    SECTION("All positive true") {
        testBinaryOp(7_u256, 3_u256, 1_u256, OpCode::SGT);
    }
    SECTION("All positive false") {
        testBinaryOp(3_u256, 7_u256, 0_u256, OpCode::SGT);
    }
    SECTION("All negative true") {
        testBinaryOp(-3_u256, -7_u256, 1_u256, OpCode::SGT);
    }
    SECTION("All negative false") {
        testBinaryOp(-7_u256, -3_u256, 0_u256, OpCode::SGT);
    }
    SECTION("First pos, second neg true") {
        testBinaryOp(3_u256, -7_u256, 1_u256, OpCode::SGT);
    }
    SECTION("First neg, second pos false") {
        testBinaryOp(-7_u256, 3_u256, 0_u256, OpCode::SGT);
    }
    SECTION("equal") { testBinaryOp(3_u256, 3_u256, 0_u256, OpCode::SGT); }
}

TEST_CASE("OPCODE: EQ opcode is correct") {
    SECTION("Not equal") { testBinaryOp(7_u256, 3_u256, 0_u256, OpCode::EQ); }
    SECTION("equal") { testBinaryOp(3_u256, 3_u256, 1_u256, OpCode::EQ); }
    SECTION("matching tuples") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}});
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}});
        m.runOp(OpCode::EQ);
        value res = m.stack.pop();
        auto actual = nonstd::get_if<uint256_t>(&res);
        REQUIRE(actual);
        REQUIRE(*actual == 1);
        REQUIRE(m.stack.stacksize() == 0);
    }
    SECTION("different tuples") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}});
        m.stack.push(Tuple{uint256_t{1}, uint256_t{3}});
        m.runOp(OpCode::EQ);
        value res = m.stack.pop();
        auto actual = nonstd::get_if<uint256_t>(&res);
        REQUIRE(actual);
        REQUIRE(*actual == 0);
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: ISZERO opcode is correct") {
    SECTION("true") { testUnaryOp(0, 1, OpCode::ISZERO); }
    SECTION("false") { testUnaryOp(2, 0, OpCode::ISZERO); }
}

TEST_CASE("OPCODE: AND opcode is correct") {
    SECTION("3 and 9 = 1") { testBinaryOp(3, 9, 1, OpCode::BITWISE_AND); }
    SECTION("3 and 3 = 3") { testBinaryOp(3, 3, 3, OpCode::BITWISE_AND); }
}

TEST_CASE("OPCODE: OR opcode is correct") {
    SECTION("3 or 9 = 11") { testBinaryOp(3, 9, 11, OpCode::BITWISE_OR); }
    SECTION("3 or 3 = 3") { testBinaryOp(3, 3, 3, OpCode::BITWISE_OR); }
}

TEST_CASE("OPCODE: XOR opcode is correct") {
    SECTION("3 or 9 = 11") { testBinaryOp(3, 9, 10, OpCode::BITWISE_XOR); }
    SECTION("3 or 3 = 3") { testBinaryOp(3, 3, 0, OpCode::BITWISE_XOR); }
}

TEST_CASE("OPCODE: NOT opcode is correct") {
    SECTION("0") { testUnaryOp(0_u256, -1_u256, OpCode::BITWISE_NOT); }
    SECTION("3") { testUnaryOp(3_u256, -4_u256, OpCode::BITWISE_NOT); }
    SECTION("-4") { testUnaryOp(-4_u256, 3_u256, OpCode::BITWISE_NOT); }
}

TEST_CASE("OPCODE: BYTE opcode is correct") {
    SECTION("31st byte of 16 = 16") { testBinaryOp(31, 16, 16, OpCode::BYTE); }
    SECTION("3rd byte of 3 = 0") { testBinaryOp(3, 3, 0, OpCode::BYTE); }
}

TEST_CASE("OPCODE: SIGNEXTEND opcode is correct") {
    SECTION("test1") {
        testBinaryOp(0_u256, -1_u256, -1_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test2") {
        testBinaryOp(0_u256, 1_u256, 1_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test3") {
        testBinaryOp(0_u256, 127_u256, 127_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test4") {
        testBinaryOp(0_u256, 128_u256, -128_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test5") {
        testBinaryOp(0_u256, 254_u256, -2_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test6") {
        testBinaryOp(0_u256, 257_u256, 1_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test7") {
        testBinaryOp(1_u256, 65534_u256, -2_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test8") {
        testBinaryOp(1_u256, 65537_u256, 1_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test8") {
        testBinaryOp(2_u256, 65537_u256, 65537_u256, OpCode::SIGNEXTEND);
    }
    SECTION("test9") {
        testBinaryOp(50_u256, 254_u256, 254_u256, OpCode::SIGNEXTEND);
    }
}

TEST_CASE("OPCODE: HASH opcode is correct") {
    SECTION("10") {
        testUnaryOp(
            10,
            intx::from_string<uint256_t>("0xc65a7bb8d6351c1cf70c95a316cc6a92839"
                                         "c986682d98bc35f958f4883f9d2a8"),
            OpCode::HASH);
    }
}

TEST_CASE("OPCODE: TYPE opcode is correct") {
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
    SECTION("type codepoint stub") {
        MachineState m;
        m.stack.push(value{CodePointStub({0, 0}, 0)});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::TYPE);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(1)});
        REQUIRE(m.stack.stacksize() == 0);
    }
    SECTION("type tuple") {
        MachineState m;
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::TYPE);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: POP opcode is correct") {
    SECTION("pop") {
        MachineState m;
        m.stack.push(uint256_t{3});
        REQUIRE(m.stack.stacksize() == 1);
        m.runOp(OpCode::POP);
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: SPUSH opcode is correct") {
    SECTION("spush") {
        auto code = std::make_shared<Code>();
        code->addSegment();
        MachineState m{std::move(code), uint256_t{5}};
        m.runOp(OpCode::SPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: RPUSH opcode is correct") {
    SECTION("rpush") {
        MachineState m;
        m.registerVal = uint256_t(5);
        m.runOp(OpCode::RPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: RSET opcode is correct") {
    SECTION("rset") {
        MachineState m;
        m.stack.push(value{uint256_t(5)});
        m.runOp(OpCode::RSET);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.registerVal == value{uint256_t(5)});
    }
}

TEST_CASE("OPCODE: JUMP opcode is correct") {
    SECTION("jump") {
        MachineState m;
        CodePointRef cpr{0, 2};
        m.stack.push(value{CodePointStub(cpr, 73665)});
        m.runOp(OpCode::JUMP);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == cpr);
    }
}

TEST_CASE("OPCODE: CJUMP opcode is correct") {
    SECTION("cjump true") {
        MachineState m;
        CodePointRef cpr{0, 2};
        m.pc = {0, 3};
        m.stack.push(uint256_t{1});
        m.stack.push(value{CodePointStub(cpr, 73665)});
        m.runOp(OpCode::CJUMP);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == cpr);
    }
    SECTION("cjump false") {
        MachineState m;
        CodePointRef initial_pc{0, 3};
        m.pc = initial_pc;
        m.stack.push(uint256_t{0});
        m.stack.push(value{CodePointStub({0, 10}, 73665)});
        m.runOp(OpCode::CJUMP);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == initial_pc + 1);
    }
}

TEST_CASE("OPCODE: STACKEMPTY opcode is correct") {
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

TEST_CASE("OPCODE: PCPUSH opcode is correct") {
    SECTION("pcpush") {
        auto code = std::make_shared<Code>();
        auto stub = code->addSegment();
        code->addOperation(stub.pc, Operation(OpCode::ADD));
        MachineState m{std::move(code), uint256_t(5)};
        auto initial_stub = CodePointStub(m.pc, m.loadCurrentInstruction());
        m.runOp(OpCode::PCPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        REQUIRE(m.pc == stub.pc);
        value res = m.stack.pop();
        REQUIRE(res == value{initial_stub});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: AUXPUSH opcode is correct") {
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

TEST_CASE("OPCODE: AUXPOP opcode is correct") {
    SECTION("auxpop") {
        MachineState m;
        m.auxstack.push(value{uint256_t(5)});
        m.runOp(OpCode::AUXPOP);
        REQUIRE(m.auxstack.stacksize() == 0);
        REQUIRE(m.stack.stacksize() == 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("OPCODE: AUXSTACKEMPTY opcode is correct") {
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

MachineState createTestMachineState(OpCode op) {
    auto code = std::make_shared<Code>();
    auto stub = code->addSegment();
    stub = code->addOperation(stub.pc, {OpCode::HALT});
    code->addOperation(stub.pc, {op});
    return {std::move(code), Tuple()};
}

TEST_CASE("OPCODE: NOP opcode is correct") {
    SECTION("nop") {
        MachineState m = createTestMachineState(OpCode::NOP);
        auto start_pc = m.pc;
        m.runOne();
        REQUIRE(m.auxstack.stacksize() == 0);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == start_pc + 1);
    }
}

TEST_CASE("OPCODE: ERRPUSH opcode is correct") {
    SECTION("errpush") {
        auto code = std::make_shared<Code>();
        auto stub = code->addSegment();
        stub = code->addOperation(stub.pc, Operation(OpCode::ADD));
        MachineState m{std::move(code), uint256_t(5)};
        m.errpc = stub;
        m.runOp(OpCode::ERRPUSH);
        REQUIRE(m.stack.stacksize() == 1);
        REQUIRE(m.pc == CodePointRef{0, 0});
        value res = m.stack.pop();
        REQUIRE(res == value{stub});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: ERRSET opcode is correct") {
    SECTION("errset") {
        MachineState m = createTestMachineState(OpCode::ERRSET);
        auto start_pc = m.pc;
        auto new_err_stub = CodePointStub({0, 54}, 968967);
        m.stack.push(value{new_err_stub});
        m.runOne();
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.pc == start_pc + 1);
        REQUIRE(m.errpc == new_err_stub);
    }
}

TEST_CASE("OPCODE: DUP0 opcode is correct") {
    SECTION("dup") {
        MachineState m = createTestMachineState(OpCode::DUP0);
        auto start_pc = m.pc;
        m.stack.push(uint256_t{5});
        m.runOne();
        REQUIRE(m.stack.stacksize() == 2);
        REQUIRE(m.pc == start_pc + 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("OPCODE: DUP1 opcode is correct") {
    SECTION("dup") {
        MachineState m = createTestMachineState(OpCode::DUP1);
        auto start_pc = m.pc;
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOne();
        REQUIRE(m.stack.stacksize() == 3);
        REQUIRE(m.pc == start_pc + 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
    }
}

TEST_CASE("OPCODE: DUP2 opcode is correct") {
    SECTION("dup") {
        MachineState m = createTestMachineState(OpCode::DUP2);
        auto start_pc = m.pc;
        m.stack.push(uint256_t{7});
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::DUP2);
        REQUIRE(m.stack.stacksize() == 4);
        REQUIRE(m.pc == start_pc + 1);
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

TEST_CASE("OPCODE: SWAP1 opcode is correct") {
    SECTION("swap") {
        MachineState m = createTestMachineState(OpCode::SWAP1);
        auto start_pc = m.pc;
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::SWAP1);
        REQUIRE(m.stack.stacksize() == 2);
        REQUIRE(m.pc == start_pc + 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
    }
}

TEST_CASE("OPCODE: SWAP2 opcode is correct") {
    SECTION("swap") {
        MachineState m = createTestMachineState(OpCode::SWAP2);
        auto start_pc = m.pc;
        m.stack.push(uint256_t{7});
        m.stack.push(uint256_t{5});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::SWAP2);
        REQUIRE(m.stack.stacksize() == 3);
        REQUIRE(m.pc == start_pc + 1);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(7)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(5)});
        res = m.stack.pop();
        REQUIRE(res == value{uint256_t(3)});
    }
}

TEST_CASE("OPCODE: TGET opcode is correct") {
    SECTION("tget") {
        MachineState m;
        m.stack.push(
            Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7}, uint256_t{6}});
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::TGET);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(8)});
        REQUIRE(m.stack.stacksize() == 0);
    }

    SECTION("index out range") {
        MachineState m;
        m.stack.push(
            Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7}, uint256_t{6}});
        m.stack.push(uint256_t{5});
        try {
            m.runOp(OpCode::TGET);
        } catch (const bad_tuple_index& e) {
            m.state = Status::Error;
        }
        // should throw bad_tuple_index and leave two items on stack
        REQUIRE(m.state == Status::Error);
        REQUIRE(m.stack.stacksize() == 2);
    }
}

TEST_CASE("OPCODE: TSET opcode is correct") {
    SECTION("2 tup") {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.stack.push(Tuple{uint256_t{1}, uint256_t{2}});
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::TSET);
        value res = m.stack.pop();
        REQUIRE(res == value{Tuple{uint256_t{1}, uint256_t{3}}});
        REQUIRE(m.stack.stacksize() == 0);
    }

    SECTION("8 tup") {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.stack.push(Tuple{uint256_t{9}, uint256_t{9}, uint256_t{9},
                           uint256_t{9}, uint256_t{9}, uint256_t{9},
                           uint256_t{9}, uint256_t{9}});
        m.stack.push(uint256_t{7});
        m.runOp(OpCode::TSET);
        value res = m.stack.pop();
        REQUIRE(res == value{Tuple{uint256_t{9}, uint256_t{9}, uint256_t{9},
                                   uint256_t{9}, uint256_t{9}, uint256_t{9},
                                   uint256_t{9}, uint256_t{3}}});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: TLEN opcode is correct") {
    SECTION("tlen") {
        MachineState m;
        m.stack.push(
            Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7}, uint256_t{6}});
        m.runOp(OpCode::TLEN);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(4)});
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: XGET opcode is correct") {
    SECTION("correct") {
        MachineState m;
        m.auxstack.push(
            Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7}, uint256_t{6}});
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::XGET);
        value res = m.stack.pop();
        REQUIRE(res == value{uint256_t(8)});
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.auxstack.stacksize() == 1);
    }

    SECTION("index out range") {
        MachineState m;
        m.auxstack.push(
            Tuple{uint256_t{9}, uint256_t{8}, uint256_t{7}, uint256_t{6}});
        m.stack.push(uint256_t{5});

        CHECK_THROWS(m.runOp(OpCode::XGET));
        // should throw bad_tuple_index and leave two items on stack
        REQUIRE(m.stack.stacksize() == 1);
    }
}

TEST_CASE("OPCODE: XSET opcode is correct") {
    SECTION("2 tup") {
        MachineState m;
        m.auxstack.push(Tuple{uint256_t{1}, uint256_t{2}});
        m.stack.push(uint256_t{3});
        m.stack.push(uint256_t{1});
        m.runOp(OpCode::XSET);
        value res = m.auxstack.pop();
        REQUIRE(res == value{Tuple{uint256_t{1}, uint256_t{3}}});
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.auxstack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: BREAKPOINT opcode is correct") {
    SECTION("break") {
        MachineState m;
        auto blockReason = m.runOp(OpCode::BREAKPOINT);
        REQUIRE(m.state == Status::Extensive);
        REQUIRE(nonstd::get_if<BreakpointBlocked>(&blockReason) != nullptr);
        REQUIRE(m.stack.stacksize() == 0);
    }
}

TEST_CASE("OPCODE: LOG opcode is correct") {
    SECTION("log") {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::LOG);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.context.logs[0] == value{uint256_t(3)});
    }
}

TEST_CASE("OPCODE: SEND opcode is correct") {
    SECTION("send") {
        // TODO: fill in send test
        MachineState m;
        Buffer buf{};
        buf.set(0, 200);
        m.stack.push(std::move(buf));
        m.stack.push(uint256_t{1});

        m.runOp(OpCode::SEND);
        REQUIRE(m.stack.stacksize() == 0);
        REQUIRE(m.state == Status::Extensive);
    }
}

TEST_CASE("OPCODE: PUSHGAS opcode is correct") {
    MachineState m;
    m.arb_gas_remaining = 250;
    m.runOp(OpCode::PUSH_GAS);
    value res = m.stack.pop();
    REQUIRE(res == value{uint256_t(250)});
    REQUIRE(m.stack.stacksize() == 0);
    REQUIRE(m.auxstack.stacksize() == 0);
}

TEST_CASE("OPCODE: SET_GAS opcode is correct") {
    MachineState m;
    m.stack.push(uint256_t{100});
    m.runOp(OpCode::SET_GAS);
    REQUIRE(m.arb_gas_remaining == 100);
    REQUIRE(m.stack.stacksize() == 0);
    REQUIRE(m.auxstack.stacksize() == 0);
}

uint256_t& assumeInt(value& val) {
    auto aNum = nonstd::get_if<uint256_t>(&val);
    if (!aNum) {
        throw bad_pop_type{};
    }
    return *aNum;
}

TEST_CASE("OPCODE: ecrecover opcode is correct") {
    std::array<unsigned char, 32> msg;
    std::generate(msg.begin(), msg.end(), std::rand);
    std::array<unsigned char, 32> seckey;
    std::generate(seckey.begin(), seckey.end(), std::rand);

    auto ctx = secp256k1_context_create(SECP256K1_CONTEXT_SIGN |
                                        SECP256K1_CONTEXT_VERIFY);
    secp256k1_ecdsa_recoverable_signature sig;
    secp256k1_pubkey pubkey;
    REQUIRE(secp256k1_ec_pubkey_create(ctx, &pubkey, seckey.data()) == 1);
    std::array<unsigned char, 65> pubkey_raw;
    size_t output_length = 65;
    REQUIRE(secp256k1_ec_pubkey_serialize(ctx, pubkey_raw.data(),
                                          &output_length, &pubkey,
                                          SECP256K1_EC_UNCOMPRESSED));
    REQUIRE(output_length == 65);

    REQUIRE(secp256k1_ecdsa_sign_recoverable(
                ctx, &sig, msg.data(), seckey.data(), nullptr, nullptr) == 1);

    std::array<unsigned char, 64> sig_raw;
    int recovery_id;
    REQUIRE(secp256k1_ecdsa_recoverable_signature_serialize_compact(
                ctx, sig_raw.data(), &recovery_id, &sig) == 1);

    MachineState s;
    s.stack.push(intx::be::unsafe::load<uint256_t>(msg.begin()));
    s.stack.push(uint256_t{recovery_id});
    s.stack.push(intx::be::unsafe::load<uint256_t>(sig_raw.begin() + 32));
    s.stack.push(intx::be::unsafe::load<uint256_t>(sig_raw.begin()));
    s.runOp(OpCode::ECRECOVER);
    REQUIRE(s.stack[0] != value(0));
    auto hash_val = ethash::keccak256(pubkey_raw.begin() + 1, 64);
    std::fill(&hash_val.bytes[0], &hash_val.bytes[12], 0);
    auto correct_address = intx::be::load<uint256_t>(hash_val);
    auto calculated_address = assumeInt(s.stack[0]);
    REQUIRE(correct_address == calculated_address);

    BENCHMARK_ADVANCED("ecrecover")(Catch::Benchmark::Chronometer meter) {
        MachineState sample_machine;
        sample_machine.stack.push(
            intx::be::unsafe::load<uint256_t>(msg.begin()));
        sample_machine.stack.push(uint256_t{recovery_id});
        sample_machine.stack.push(
            intx::be::unsafe::load<uint256_t>(sig_raw.begin() + 32));
        sample_machine.stack.push(
            intx::be::unsafe::load<uint256_t>(sig_raw.begin()));

        std::vector<MachineState> machines(meter.runs());
        std::fill(machines.begin(), machines.end(), sample_machine);
        meter.measure([&machines](int i) {
            return machines[i].runOp(OpCode::ECRECOVER);
        });
    };
}

TEST_CASE("OPCODE: ECADD") {
    alt_bn128_pp::init_public_params();
    for (const auto& test_case : prepareECAddCases()) {
        MachineState mach;
        mach.stack.push(test_case.b.y);
        mach.stack.push(test_case.b.x);
        mach.stack.push(test_case.a.y);
        mach.stack.push(test_case.a.x);
        mach.runOp(OpCode::ECADD);
        REQUIRE(mach.state == Status::Extensive);
        REQUIRE(mach.stack[1] == value{test_case.res.y});
        REQUIRE(mach.stack[0] == value{test_case.res.x});
    }
}

TEST_CASE("OPCODE: ECMUL") {
    alt_bn128_pp::init_public_params();
    for (const auto& test_case : prepareECMulCases()) {
        MachineState mach;
        mach.stack.push(test_case.k);
        mach.stack.push(test_case.a.y);
        mach.stack.push(test_case.a.x);
        mach.runOp(OpCode::ECMUL);
        REQUIRE(mach.state == Status::Extensive);
        REQUIRE(mach.stack[1] == value{test_case.res.y});
        REQUIRE(mach.stack[0] == value{test_case.res.x});
    }
}

TEST_CASE("OPCODE: ECPAIRING") {
    alt_bn128_pp::init_public_params();
    auto cases = preparePairingCases();
    for (const auto& testCase : cases) {
        Tuple tup;
        for (const auto& point : testCase.points) {
            tup =
                Tuple(Tuple(point.first.x, point.first.y, point.second.x0,
                            point.second.x1, point.second.y0, point.second.y1),
                      tup);
        }
        MachineState mach;
        mach.stack.push(tup);
        mach.runOp(OpCode::ECPAIRING);
        REQUIRE(mach.state == Status::Extensive);
        REQUIRE(mach.stack[0] == value(testCase.valid ? 1 : 0));
    }
}

TEST_CASE("OPCODE: INBOX opcode is correct") {
    SECTION("inbox") {
        // TODO: fill in inbox test
    }
}

TEST_CASE("OPCODE: ERROR opcode is correct") {
    SECTION("error") {
        // TODO: fill in error test
    }
}

TEST_CASE("OPCODE: HALT opcode is correct") {
    SECTION("halt") {
        // TODO: fill in halt test
    }
}

TEST_CASE("OPCODE: KECCAKF opcode is correct") {
    auto code = std::make_shared<Code>();
    SECTION("Inverts correctly") {
        Tuple input_data(intx::from_string<uint256_t>(
                             "94370651106686220754648249265079798778273"
                             "932128194559331492955050019282050496"),
                         intx::from_string<uint256_t>(
                             "42512909751185556122923115391154208487752"
                             "310613213055089416300774052282720344"),
                         intx::from_string<uint256_t>(
                             "56208326812724912066026123588383649819390"
                             "601658448049319166841561743369815863"),
                         intx::from_string<uint256_t>(
                             "42512909751185556122923115391154208487752"
                             "310613213055089416300774052282720344"),
                         intx::from_string<uint256_t>(
                             "11318235288944921066599402722758875429096"
                             "9798016938687372921424809289618385856"),
                         intx::from_string<uint256_t>(
                             "81755589384323691266272576345129881657705"
                             "914621008081459572116739688988488432"),
                         uint256_t{6345636445});
        uint64_t state[25];
        machineoperation::internal::encodeKeccakState(input_data, state);
        auto ret = machineoperation::internal::decodeKeccakState(state);
        REQUIRE(ret == input_data);
    }

    SECTION("Hashes correctly") {
        auto stub = code->addSegment();
        stub = code->addOperation(stub.pc, Operation(OpCode::KECCAKF));
        code->addOperation(stub.pc, Operation(OpCode::KECCAKF));
        MachineState m{std::move(code), Tuple()};
        m.stack.push(
            Tuple(0_u256, 0_u256, 0_u256, 0_u256, 0_u256, 0_u256, 0_u256));
        m.runOne();
        auto ret = m.stack.pop();
        {
            REQUIRE(nonstd::holds_alternative<Tuple>(ret));
            auto ret_tup = nonstd::get<Tuple>(ret);
            REQUIRE(ret_tup.tuple_size() == 7);
            std::array<uint256_t, 7> parts;
            for (size_t i = 0; i < 7; ++i) {
                auto val = ret_tup.get_element(i);
                REQUIRE(nonstd::holds_alternative<uint256_t>(val));
                parts[i] = nonstd::get<uint256_t>(val);
            }

            uint256_t correct0 = hexToInt(
                "bd1547306f80494dd598261ea65aa9ee84d5ccf933c0478af1258f7940e1dd"
                "e7");
            uint256_t correct1 = hexToInt(
                "8c5bda0cd6192e7690fee5a0a44647c4ff97a42d7f8e6fd48b284e056253d0"
                "57");
            uint256_t correct2 = hexToInt(
                "a9a6e6260d712103eb5aa93f2317d63530935ab7d08ffc64ad30a6f71b1905"
                "9c");
            uint256_t correct3 = hexToInt(
                "05e5635a21d9ae6101f22f1a11a5569f43b831cd0347c82681a57c16dbcf55"
                "5f");
            uint256_t correct4 = hexToInt(
                "8c3ee88a1ccf32c8b87c5a554fd00ecb613670957bc4661164befef28cc970"
                "f2");
            uint256_t correct5 = hexToInt(
                "75f644e97f30a13b16f53526e70465c21841f924a2c509e4940c7922ae3a26"
                "14");
            uint256_t correct6 = hexToInt(
                "000000000000000000000000000000000000000000000000eaf1ff7b5ceca2"
                "49");

            REQUIRE(parts[0] == correct0);
            REQUIRE(parts[1] == correct1);
            REQUIRE(parts[2] == correct2);
            REQUIRE(parts[3] == correct3);
            REQUIRE(parts[4] == correct4);
            REQUIRE(parts[5] == correct5);
            REQUIRE(parts[6] == correct6);
        }

        m.stack.push(std::move(ret));
        m.runOne();
        ret = m.stack.pop();
        {
            REQUIRE(nonstd::holds_alternative<Tuple>(ret));
            auto ret_tup = nonstd::get<Tuple>(ret);
            REQUIRE(ret_tup.tuple_size() == 7);
            std::array<uint256_t, 7> parts;
            for (size_t i = 0; i < 7; ++i) {
                auto val = ret_tup.get_element(i);
                REQUIRE(nonstd::holds_alternative<uint256_t>(val));
                parts[i] = nonstd::get<uint256_t>(val);
            }
            uint256_t correct0 = hexToInt(
                "8a20d9b25569d094093d8d1270d76b6c6a332cd07057b56d2d5c954df96ecb"
                "3c");
            uint256_t correct1 = hexToInt(
                "faf4f247c3d810f785773dae1275af0df957b9a2da65fb384f9c4f99e5e7f1"
                "56");
            uint256_t correct2 = hexToInt(
                "deea66c4ba8f974f68ce61b6b9ce68a1e4fecc0fee98b4251f1b9ee6f79a87"
                "59");
            uint256_t correct3 = hexToInt(
                "fd5449a6bf1747437cf8a9f009831265e00654042719dbd933c43d836eafb1"
                "f5");
            uint256_t correct4 = hexToInt(
                "91a0226e649e42e9e3b8c8ee55b7b03c48ead5fc5d0be77497ddad33d8994b"
                "40");
            uint256_t correct5 = hexToInt(
                "609f4e62a44c10595b3402464e1c3db6202a9ec5faa3cce8900e3129e7badd"
                "7b");
            uint256_t correct6 = hexToInt(
                "00000000000000000000000000000000000000000000000020d06cd26a8fbf"
                "5c");

            REQUIRE(parts[0] == correct0);
            REQUIRE(parts[1] == correct1);
            REQUIRE(parts[2] == correct2);
            REQUIRE(parts[3] == correct3);
            REQUIRE(parts[4] == correct4);
            REQUIRE(parts[5] == correct5);
            REQUIRE(parts[6] == correct6);
        }
    }
}

TEST_CASE("OPCODE: SHA256F opcode is correct") {
    uint256_t initial_hash_state = hexToInt(
        "6a09e667bb67ae853c6ef372a54ff53a510e527f9b05688c1f83d9ab5be0cd19");
    struct Sha256Case {
        std::vector<unsigned char> raw_input_data;
        uint256_t input_first;
        uint256_t input_second;
        uint256_t output_digest;
    };
    std::vector<Sha256Case> cases = {
        {hexToVec(
             "00000000000000000000000000000000000000000000000000000000000000000"
             "000000000000000000000000000000000000000000000000000000000000000"),
         0_u256, 0_u256,
         hexToInt("da5698be17b9b46962335799779fbeca8ce5d491c0d26243bafef9ea1837"
                  "a9d8")},
        {hexToVec(
             "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
             "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
         hexToInt("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
                  "FFFF"),
         hexToInt("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
                  "FFFF"),
         hexToInt("ef0c748df4da50a8d6c43c013edc3ce76c9d9fa9a1458ade56eb86c0a644"
                  "92d2")},
        {hexToVec(
             "000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f2"
             "02122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d3e3f"),
         hexToInt("000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d"
                  "1e1f"),
         hexToInt("202122232425262728292a2b2c2d2e2f303132333435363738393a3b3c3d"
                  "3e3f"),
         hexToInt("fc99a2df88f42a7a7bb9d18033cdc6a20256755f9d5b9a5044a9cc315abe"
                  "84a7")},
        {hexToVec(
             "243f6a8885a308d313198a2e03707344a4093822299f31d0082efa98ec4e6c894"
             "52821e638d01377be5466cf34e90c6cc0ac29b7c97c50dd3f84d5b5b5470917"),
         hexToInt("243F6A8885A308D313198A2E03707344A4093822299F31D0082EFA98EC4E"
                  "6C89"),
         hexToInt("452821E638D01377BE5466CF34E90C6CC0AC29B7C97C50DD3f84D5B5b547"
                  "0917"),
         hexToInt("cf0ae4eb67d38ffeb94068984b22abde4e92bc548d14585e48dca8882d7b"
                  "09ce")},
    };

    SECTION("Block works correction") {
        int i = 0;
        for (const auto& test_case : cases) {
            DYNAMIC_SECTION(i) {
                std::array<uint8_t, 64> input_data;
                std::copy(test_case.raw_input_data.begin(),
                          test_case.raw_input_data.end(), input_data.begin());
                auto ret = machineoperation::internal::sha256_block(
                    initial_hash_state, input_data);
                REQUIRE(ret == test_case.output_digest);
            }
            ++i;
        }
    }

    SECTION("Hashes correctly") {
        auto code = std::make_shared<Code>();
        auto stub = code->addSegment();
        stub = code->addOperation(stub.pc, Operation(OpCode::SHA256F));
        MachineState m{std::move(code), Tuple()};

        int i = 0;
        for (const auto& test_case : cases) {
            DYNAMIC_SECTION(i) {
                m.stack.push(test_case.input_second);
                m.stack.push(test_case.input_first);
                m.stack.push(initial_hash_state);
                m.runOne();
                auto ret = m.stack.pop();
                REQUIRE(nonstd::holds_alternative<uint256_t>(ret));
                REQUIRE(ret.get<uint256_t>() == test_case.output_digest);
            }
            ++i;
        }
    }
}

TEST_CASE("OPCODE: Stack underflow") {
    for (uint8_t op = static_cast<uint8_t>(OpCode::ADD);
         op <= static_cast<uint8_t>(OpCode::ECRECOVER); ++op) {
        auto code = std::make_shared<Code>();
        auto stub = code->addSegment();
        code->addOperation(stub.pc, Operation(static_cast<OpCode>(op)));
        MachineState m{std::move(code), uint256_t{5}};
        m.runOne();
    }
}

TEST_CASE("OPCODE: Newbuffer opcode") {
    SECTION("Creates new buffer") {
        MachineState mach;
        mach.runOp(OpCode::NEW_BUFFER);
        REQUIRE(mach.stack[0] == value{Buffer()});
    }
}

TEST_CASE("OPCODE: getbuffer8 opcode") {
    SECTION("Reads from buffer") {
        MachineState mach;
        Buffer buf;
        buf = buf.set(123, 7);
        mach.stack.push(buf);
        mach.stack.push(uint256_t{123});
        mach.runOp(OpCode::GET_BUFFER8);
        REQUIRE(mach.stack[0] == value{uint256_t{7}});
    }
}

TEST_CASE("OPCODE: getbuffer64 opcode") {
    SECTION("Reads from buffer") {
        MachineState mach;
        Buffer buf;
        buf = buf.set(123, 7);
        mach.stack.push(buf);
        mach.stack.push(uint256_t{123});
        mach.runOp(OpCode::GET_BUFFER64);
        REQUIRE(mach.stack[0] == value{uint256_t{7L << 56L}});
    }
}

TEST_CASE("OPCODE: getbuffer256 opcode") {
    SECTION("Reads from buffer") {
        MachineState mach;
        Buffer buf;
        buf = buf.set(123, 7);
        mach.stack.push(buf);
        mach.stack.push(uint256_t{123});
        mach.runOp(OpCode::GET_BUFFER256);
        REQUIRE(mach.stack[0] == value{uint256_t{7L} << 248});
    }
}

TEST_CASE("OPCODE: setbuffer8 opcode") {
    SECTION("Writes to buffer") {
        MachineState mach;
        Buffer buf;
        buf = buf.set(123, 7);
        mach.stack.push(Buffer());
        mach.stack.push(uint256_t{7});
        mach.stack.push(uint256_t{123});
        mach.runOp(OpCode::SET_BUFFER8);
        REQUIRE(mach.stack[0] == value{buf});
    }
}

TEST_CASE("OPCODE: setbuffer64 opcode") {
    SECTION("Writes to buffer") {
        MachineState mach;
        Buffer buf;
        buf = buf.set(123, 9);
        buf = buf.set(123 + 1, 8);
        buf = buf.set(123 + 7, 7);
        mach.stack.push(Buffer());
        mach.stack.push(uint256_t{0x0908000000000007L});
        mach.stack.push(uint256_t{123});
        mach.runOp(OpCode::SET_BUFFER64);
        REQUIRE(mach.stack[0] == value{buf});
    }
}

TEST_CASE("OPCODE: setbuffer256 opcode") {
    SECTION("Writes to buffer") {
        MachineState mach;
        Buffer buf;
        buf = buf.set(123, 9);
        buf = buf.set(123 + 1, 8);
        buf = buf.set(123 + 7, 7);
        buf = buf.set(123 + 8, 9);
        buf = buf.set(123 + 1 + 8, 8);
        buf = buf.set(123 + 7 + 8, 7);
        buf = buf.set(123 + 16, 9);
        buf = buf.set(123 + 1 + 16, 8);
        buf = buf.set(123 + 7 + 16, 7);
        buf = buf.set(123 + 24, 9);
        buf = buf.set(123 + 1 + 24, 8);
        buf = buf.set(123 + 7 + 24, 7);
        mach.stack.push(Buffer());
        mach.stack.push(
            hexToInt("090800000000000709080000000000070908000000000007090800000"
                     "0000007"));
        mach.stack.push(uint256_t{123});
        mach.runOp(OpCode::SET_BUFFER256);
        REQUIRE(mach.stack[0] == value{buf});
    }
}
