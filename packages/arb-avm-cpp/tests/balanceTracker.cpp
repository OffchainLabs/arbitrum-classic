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

#include <stdio.h>

void pushMessage(MachineState& m,
                 const value& data,
                 const uint256_t& dest,
                 const uint256_t& currency,
                 const TokenType& tokType,
                 TuplePool& pool) {
    m.stack.push(Tuple{data, dest, currency, fromTokenType(tokType), &pool});
}

void sendInboxMessage(MachineState& m,
                      value data,
                      uint256_t dest,
                      uint256_t currency,
                      TokenType tokType) {
    m.sendOnchainMessage({data, dest, currency, tokType});
}

MachineState setupMachine(value data,
                          uint256_t dest,
                          uint256_t currency,
                          TokenType tokType) {
    MachineState m;

    return m;
}

TEST_CASE("SEND") {
    SECTION("Successful send") {
        value data;
        TokenType token = {};
        token[0] = 15;
        token[20] = 0;
        uint256_t currency = 20;
        uint256_t destination = 25;
        MachineState m;
        m.state = Status::Extensive;
        sendInboxMessage(m, data, destination, currency, token);

        currency = 4;
        pushMessage(m, data, destination, currency, token, *(m.pool.get()));
        m.runOp(OpCode::SEND);
        uint256_t resNum = m.balance.tokenValue(token);
        REQUIRE(resNum == 16);
        REQUIRE(m.state == Status::Extensive);
    }

    SECTION("Not enough funds send") {
        value data;
        TokenType token = {};
        token[0] = 15;
        token[20] = 0;
        uint256_t currency = 20;
        uint256_t destination = 25;
        MachineState m;
        m.state = Status::Extensive;
        sendInboxMessage(m, data, destination, currency, token);

        currency = 25;
        pushMessage(m, data, destination, currency, token, *(m.pool.get()));
        auto blockReason = m.runOp(OpCode::SEND);
        uint256_t resNum = m.balance.tokenValue(token);
        REQUIRE(resNum == 20);
        REQUIRE(m.state == Status::Extensive);
        REQUIRE(nonstd::get_if<SendBlocked>(&blockReason));
    }

    SECTION("Successful NF send") {
        value data;
        TokenType token = {};
        token.fill(0);
        token[0] = 15;
        token[20] = 1;
        uint256_t currency = 1;
        uint256_t destination = 25;
        MachineState m;
        m.state = Status::Extensive;
        sendInboxMessage(m, data, destination, currency, token);

        pushMessage(m, data, destination, currency, token, *(m.pool.get()));

        REQUIRE(m.balance.hasNFT(token, currency));
        m.runOp(OpCode::SEND);
        REQUIRE(!m.balance.hasNFT(token, currency));
        REQUIRE(m.state == Status::Extensive);
    }

    SECTION("send unowned NF token") {
        value data;
        TokenType token = {};
        token[0] = 15;
        token[20] = 1;
        uint256_t currency = 1;
        uint256_t destination = 25;
        MachineState m;
        m.state = Status::Extensive;

        pushMessage(m, data, destination, currency, token, *(m.pool.get()));
        auto blockReason = m.runOp(OpCode::SEND);
        REQUIRE(!m.balance.hasNFT(token, currency));
        REQUIRE(m.state == Status::Extensive);
        REQUIRE(nonstd::get_if<SendBlocked>(&blockReason));
    }
}

TEST_CASE("NBSEND") {
    SECTION("Successful nbsend") {
        value data;
        TokenType token = {};
        token[0] = 15;
        token[20] = 0;
        uint256_t currency = 20;
        uint256_t destination = 25;
        MachineState m;
        m.state = Status::Extensive;
        sendInboxMessage(m, data, destination, currency, token);

        currency = 4;
        pushMessage(m, data, destination, currency, token, *(m.pool.get()));
        m.runOp(OpCode::NBSEND);
        uint256_t resNum = m.balance.tokenValue(token);
        REQUIRE(resNum == 16);
        value retval = m.stack.pop();
        auto ret = nonstd::get_if<uint256_t>(&retval);
        REQUIRE(*ret == 1);
        REQUIRE(m.state == Status::Extensive);
    }

    SECTION("Not enough funds nbsend") {
        value data;
        TokenType token = {};
        token[0] = 15;
        token[20] = 0;
        uint256_t currency = 20;
        uint256_t destination = 25;
        MachineState m;
        m.state = Status::Extensive;
        sendInboxMessage(m, data, destination, currency, token);

        currency = 25;
        pushMessage(m, data, destination, currency, token, *(m.pool.get()));
        m.runOp(OpCode::NBSEND);
        uint256_t resNum = m.balance.tokenValue(token);
        REQUIRE(resNum == 20);
        value retval = m.stack.pop();
        auto ret = nonstd::get_if<uint256_t>(&retval);
        REQUIRE(*ret == 0);
        REQUIRE(m.state == Status::Extensive);
    }
    SECTION("Successful NF nbsend") {
        value data;
        TokenType token = {};
        token.fill(0);
        token[0] = 15;
        token[20] = 1;
        uint256_t currency = 1;
        uint256_t destination = 25;
        MachineState m;
        m.state = Status::Extensive;
        sendInboxMessage(m, data, destination, currency, token);

        pushMessage(m, data, destination, currency, token, *(m.pool.get()));
        REQUIRE(m.balance.hasNFT(token, currency));
        m.runOp(OpCode::NBSEND);
        REQUIRE(!m.balance.hasNFT(token, currency));
        value retval = m.stack.pop();
        auto ret = nonstd::get_if<uint256_t>(&retval);
        REQUIRE(*ret == 1);
        REQUIRE(m.state == Status::Extensive);
    }

    SECTION("nbsend unowned NF token") {
        value data;
        TokenType token = {};
        token[0] = 15;
        token[20] = 1;
        uint256_t currency = 1;
        uint256_t destination = 25;
        MachineState m;
        m.state = Status::Extensive;

        pushMessage(m, data, destination, currency, token, *(m.pool.get()));
        REQUIRE(!m.balance.hasNFT(token, currency));
        m.runOp(OpCode::NBSEND);
        REQUIRE(!m.balance.hasNFT(token, currency));
        value retval = m.stack.pop();
        auto ret = nonstd::get_if<uint256_t>(&retval);
        REQUIRE(*ret == 0);
        REQUIRE(m.state == Status::Extensive);
    }
}
