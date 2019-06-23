//
//  main.cpp
//  tests
//
//  Created by Harry Kalodner on 6/23/19.
//

#include <avm/machine.hpp>

#include <catch2/catch.hpp>

TEST_CASE( "SUB opcode is correct") {
    SECTION( "Non overlow is correct" ) {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.stack.push(uint256_t{4});
        m.runOp(OpCode::SUB);
        value res = m.stack.pop();
        auto num = mpark::get_if<uint256_t>(&res);
        REQUIRE(num);
        REQUIRE(*num == 1);
    }
    
    SECTION( "Overlow is correct" ) {
        MachineState m;
        m.stack.push(uint256_t{4});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::SUB);
        value res = m.stack.pop();
        auto num = mpark::get_if<uint256_t>(&res);
        REQUIRE(num);
        REQUIRE(*num == uint256_t{-1});
    }
}


TEST_CASE( "DIV opcode is correct") {
    SECTION( "Non overlow is correct" ) {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.stack.push(uint256_t{12});
        m.runOp(OpCode::DIV);
        value res = m.stack.pop();
        auto num = mpark::get_if<uint256_t>(&res);
        REQUIRE(num);
        REQUIRE(*num == 4);
    }
    
    SECTION( "Divide by zero" ) {
        MachineState m;
        m.stack.push(uint256_t{0});
        m.stack.push(uint256_t{3});
        m.runOp(OpCode::DIV);
        value res = m.stack.pop();
        REQUIRE(m.state == ERROR);
    }
}

TEST_CASE( "SDIV opcode is correct") {
    SECTION( "Positive divided by negative" ) {
        MachineState m;
        m.stack.push(uint256_t{-3});
        m.stack.push(uint256_t{12});
        m.runOp(OpCode::SDIV);
        value res = m.stack.pop();
        auto num = mpark::get_if<uint256_t>(&res);
        REQUIRE(num);
        REQUIRE(*num == uint256_t{-4});
    }
    
    SECTION( "Negative divided by negative" ) {
        MachineState m;
        m.stack.push(uint256_t{-3});
        m.stack.push(uint256_t{-12});
        m.runOp(OpCode::SDIV);
        value res = m.stack.pop();
        auto num = mpark::get_if<uint256_t>(&res);
        REQUIRE(num);
        REQUIRE(*num == uint256_t{4});
    }
}

TEST_CASE( "SMOD opcode is correct") {
    SECTION( "Positive divided by negative" ) {
        MachineState m;
        m.stack.push(uint256_t{3});
        m.stack.push(uint256_t{-8});
        m.runOp(OpCode::SMOD);
        value res = m.stack.pop();
        auto num = mpark::get_if<uint256_t>(&res);
        REQUIRE(num);
        REQUIRE(*num == uint256_t{-2});
    }
}
