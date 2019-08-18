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
                 value data,
                 uint256_t dest,
                 uint256_t currency,
                 TokenType tokType,
                 TuplePool& pool) {
    Tuple msgTup(&pool, 4);
    msgTup.set_element(0, data);
    msgTup.set_element(1, dest);
    msgTup.set_element(2, currency);
    msgTup.set_element(3, fromTokenType(tokType));
    m.stack.push(std::move(msgTup));
}

void sendInboxMessage(MachineState& m,
                      value data,
                      uint256_t dest,
                      uint256_t currency,
                      TokenType tokType) {
    Message msg;
    msg.data = data;
    msg.token = tokType;
    msg.currency = currency;
    msg.destination = dest;
    m.sendOnchainMessage(msg);
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
        m.runOp(OpCode::SEND);
        uint256_t resNum = m.balance.tokenValue(token);
        REQUIRE(resNum == 0);
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
        uint256_t resNum = m.balance.tokenValue(token);
        REQUIRE(resNum == 0);
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
        m.runOp(OpCode::NBSEND);
        uint256_t resNum = m.balance.tokenValue(token);
        REQUIRE(resNum == 0);
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
        m.runOp(OpCode::NBSEND);
        uint256_t resNum = m.balance.tokenValue(token);
        REQUIRE(resNum == 0);
        value retval = m.stack.pop();
        auto ret = nonstd::get_if<uint256_t>(&retval);
        REQUIRE(*ret == 0);
        REQUIRE(m.state == Status::Extensive);
    }
}

/*{
    value data;
    Message msg;
    msg.data = data;
    msg.token[0]=15;
    msg.token[20]=0;
    msg.currency = 20;
    msg.destination = 25;

    uint256_t tokval =fromTokenType(msg.token);
    TokenType tokenval;
    toTokenType(tokval, tokenval);
    std::cout<<tokval<<std::endl;
    std::cout<<"token value="<<tokval<<"
   tokenval[0]={"<<tokenval[0]<<"}"<<std::endl; m.addInboxMessage(msg);

    Tuple msgTup(4, m.pool.get());
    msgTup.set_element(0, msg.data);
    msgTup.set_element(1, fromTokenType(msg.token));
    msg.currency = 4;
    msgTup.set_element(2, msg.currency);
    msgTup.set_element(3, msg.destination);
    m.stack.push(std::move(msgTup));
    m.runOp(OpCode::SEND);
    uint256_t resNum = m.context.afterBalance.tokenValue(msg.token);
    if (resNum != 16){
        std::cout<<"ERROR - send failed - expected 16 received
   "<<resNum<<std::endl;
    }

    msg.currency = 18;
    msgTup.set_element(2, msg.currency);
    m.stack.push(std::move(msgTup));
    m.runOp(OpCode::SEND);
    if (m.state != HALTED){
        std::cout<<"ERROR - send failed - expected state HALTED received
   "<<m.state<<std::endl;
    }
    resNum = m.context.afterBalance.tokenValue(msg.token);
    if (resNum != 16){
        std::cout<<"ERROR - send failed - expected 16 received
   "<<resNum<<std::endl;
    }

    msg.currency = 4;
    msgTup.set_element(2, msg.currency);
    m.stack.push(std::move(msgTup));
    m.runOp(OpCode::NBSEND);
    resNum = m.context.afterBalance.tokenValue(msg.token);
    if (resNum != 12){
        std::cout<<"ERROR - send failed - expected 12 received
   "<<resNum<<std::endl;
    }
    value res = m.stack.pop();
    resNum = assumeInt(res);
    if (resNum != 1){
        std::cout<<"ERROR - sub failed - expected 1 received
   "<<resNum<<std::endl;
    }

    msg.currency = 14;
    msgTup.set_element(2, msg.currency);
    m.stack.push(std::move(msgTup));
    m.runOp(OpCode::NBSEND);
    resNum = m.context.afterBalance.tokenValue(msg.token);
    if (resNum != 12){
        std::cout<<"ERROR - send failed - expected 12 received
   "<<resNum<<std::endl;
    }
    res = m.stack.pop();
    resNum = assumeInt(res);
    if (resNum != 0){
        std::cout<<"ERROR - sub failed - expected 1 received
   "<<resNum<<std::endl;
    }
    }
*/
