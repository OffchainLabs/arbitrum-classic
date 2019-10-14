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

#include "avm/machinestate/messagestack.hpp"
#include <catch2/catch.hpp>
#include "avm/machinestate/machinestatesaver.hpp"

std::string current_path =
    "/Users/minhtruong/Dev/arbitrum/packages/arb-avm-cpp/build/tests/rocksDb";

void saveMessageStack(MessageStack stack) {
    TuplePool pool;
    CheckpointStorage storage(current_path);
    auto saver = MachineStateSaver(&storage, &pool);

    auto results = stack.checkpointState(saver);
    REQUIRE(results.msgs_tuple_results.status.ok());
    REQUIRE(results.msg_count_results.status.ok());
    REQUIRE(results.msgs_tuple_results.reference_count == 1);
    REQUIRE(results.msg_count_results.reference_count == 1);
}

void saveAndGetMessageStack(MessageStack stack,
                            uint256_t expected_tuple_hash,
                            uint64_t expected_count) {
    TuplePool pool;
    CheckpointStorage storage(current_path);
    auto saver = MachineStateSaver(&storage, &pool);

    auto results = stack.checkpointState(saver);
    auto tuple_res = saver.getTuple(results.msgs_tuple_results.storage_key);
    auto count_results = saver.getValue(results.msg_count_results.storage_key);

    REQUIRE(tuple_res.status.ok());
    REQUIRE(count_results.status.ok());

    auto count = nonstd::get<uint256_t>(count_results.val);
    auto new_stack = MessageStack(&pool, tuple_res.tuple, count);

    REQUIRE(new_stack.messages.calculateHash() == expected_tuple_hash);
    REQUIRE(new_stack.messageCount == expected_count);
    REQUIRE(tuple_res.reference_count == 1);
    REQUIRE(count_results.reference_count == 1);
}

void saveTwiceAndGetMessageStack(MessageStack stack,
                                 uint256_t expected_tuple_hash,
                                 uint64_t expected_count) {
    TuplePool pool;
    CheckpointStorage storage(current_path);
    auto saver = MachineStateSaver(&storage, &pool);

    auto results = stack.checkpointState(saver);
    auto results2 = stack.checkpointState(saver);
    auto tuple_res = saver.getTuple(results.msgs_tuple_results.storage_key);
    auto count_results = saver.getValue(results.msg_count_results.storage_key);

    REQUIRE(tuple_res.status.ok());
    REQUIRE(count_results.status.ok());

    auto count = nonstd::get<uint256_t>(count_results.val);
    auto new_stack = MessageStack(&pool, tuple_res.tuple, count);

    REQUIRE(new_stack.messages.calculateHash() == expected_tuple_hash);
    REQUIRE(new_stack.messageCount == expected_count);

    REQUIRE(tuple_res.reference_count == 2);
    REQUIRE(count_results.reference_count == 2);
}

TEST_CASE("save messagestack") {
    SECTION("empty stack") {
        TuplePool pool;
        auto stack = MessageStack(&pool);

        saveMessageStack(stack);
    }
    SECTION("stack with values") {
        TuplePool pool;
        auto stack = MessageStack(&pool);

        uint256_t val_data = 111;
        uint256_t destination = 2;
        uint256_t currency = 3;
        auto token_type = std::array<unsigned char, 21>();
        token_type[0] = 'a';
        auto msg = Message{val_data, destination, currency, token_type};

        stack.addMessage(msg);
        saveMessageStack(stack);
    }
}

TEST_CASE("save and get messagestack") {
    SECTION("empty stack") {
        TuplePool pool;
        auto stack = MessageStack(&pool);
        saveAndGetMessageStack(stack, Tuple().calculateHash(), 0);
    }
    SECTION("stack with values") {
        TuplePool pool;
        auto stack = MessageStack(&pool);

        uint256_t val_data = 111;
        uint256_t destination = 2;
        uint256_t currency = 3;
        auto token_type = std::array<unsigned char, 21>();
        token_type[0] = 'a';
        auto msg = Message{val_data, destination, currency, token_type};

        stack.addMessage(msg);
        saveAndGetMessageStack(stack, stack.messages.calculateHash(), 1);
    }
    SECTION("save stack twice, with values") {
        TuplePool pool;
        auto stack = MessageStack(&pool);

        uint256_t val_data = 111;
        uint256_t destination = 2;
        uint256_t currency = 3;
        auto token_type = std::array<unsigned char, 21>();
        token_type[0] = 'a';
        auto msg = Message{val_data, destination, currency, token_type};

        stack.addMessage(msg);
        saveTwiceAndGetMessageStack(stack, stack.messages.calculateHash(), 1);
    }
}
