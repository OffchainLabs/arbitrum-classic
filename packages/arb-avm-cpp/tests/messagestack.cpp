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

void saveMessageStack(MachineStateSaver& saver,
                      MessageStack stack,
                      int expected_tuple_ref_count,
                      int expected_num_ref_count) {
    auto results = stack.checkpointState(saver);
    REQUIRE(results.msgs_tuple_results.status.ok());
    REQUIRE(results.msg_count_results.status.ok());
    REQUIRE(results.msgs_tuple_results.reference_count ==
            expected_tuple_ref_count);
    REQUIRE(results.msg_count_results.reference_count ==
            expected_num_ref_count);
}

void getSavedMessageStack(MachineStateSaver& saver,
                          std::vector<unsigned char> msgs_key,
                          std::vector<unsigned char> count_key,
                          uint256_t expected_tuple_hash,
                          uint64_t expected_count) {
    TuplePool pool;
    MessageStack new_stack(&pool);
    auto success = new_stack.initializeMessageStack(saver, msgs_key, count_key);

    REQUIRE(success == true);
    REQUIRE(new_stack.messages.calculateHash() == expected_tuple_hash);
    REQUIRE(new_stack.messageCount == expected_count);
}

TEST_CASE("save messagestack") {
    SECTION("empty stack") {
        TuplePool pool;
        CheckpointStorage storage(current_path);
        std::vector<CodePoint> code;
        auto saver = MachineStateSaver(&storage, &pool, code);
        auto stack = MessageStack(&pool);

        saveMessageStack(saver, stack, 1, 1);
    }
    SECTION("empty stack, twice") {
        TuplePool pool;
        CheckpointStorage storage(current_path);
        std::vector<CodePoint> code;
        auto saver = MachineStateSaver(&storage, &pool, code);
        auto stack = MessageStack(&pool);

        stack.checkpointState(saver);
        saveMessageStack(saver, stack, 2, 2);
    }
    SECTION("stack with values") {
        TuplePool pool;
        CheckpointStorage storage(current_path);
        std::vector<CodePoint> code;
        auto saver = MachineStateSaver(&storage, &pool, code);
        auto stack = MessageStack(&pool);

        uint256_t val_data = 111;
        uint256_t destination = 2;
        uint256_t currency = 3;
        auto token_type = std::array<unsigned char, 21>();
        token_type[0] = 'a';
        auto msg = Message{val_data, destination, currency, token_type};

        stack.addMessage(msg);
        saveMessageStack(saver, stack, 1, 1);
    }
    SECTION("stack with values, twice") {
        TuplePool pool;
        CheckpointStorage storage(current_path);
        std::vector<CodePoint> code;
        auto saver = MachineStateSaver(&storage, &pool, code);
        auto stack = MessageStack(&pool);

        uint256_t val_data = 111;
        uint256_t destination = 2;
        uint256_t currency = 3;
        auto token_type = std::array<unsigned char, 21>();
        token_type[0] = 'a';
        auto msg = Message{val_data, destination, currency, token_type};

        stack.addMessage(msg);
        stack.checkpointState(saver);

        saveMessageStack(saver, stack, 2, 2);
    }
}

TEST_CASE("Get saved messagestack") {
    SECTION("empty stack") {
        TuplePool pool;
        CheckpointStorage storage(current_path);
        std::vector<CodePoint> code;
        auto saver = MachineStateSaver(&storage, &pool, code);

        auto stack = MessageStack(&pool);
        auto results = stack.checkpointState(saver);

        getSavedMessageStack(saver, results.msgs_tuple_results.storage_key,
                             results.msg_count_results.storage_key,
                             stack.messages.calculateHash(), 0);
    }
    SECTION("stack with values") {
        TuplePool pool;
        CheckpointStorage storage(current_path);
        std::vector<CodePoint> code;
        auto saver = MachineStateSaver(&storage, &pool, code);

        auto stack = MessageStack(&pool);

        uint256_t val_data = 111;
        uint256_t destination = 2;
        uint256_t currency = 3;
        auto token_type = std::array<unsigned char, 21>();
        token_type[0] = 'a';
        auto msg = Message{val_data, destination, currency, token_type};

        stack.addMessage(msg);
        auto results = stack.checkpointState(saver);

        getSavedMessageStack(saver, results.msgs_tuple_results.storage_key,
                             results.msg_count_results.storage_key,
                             stack.messages.calculateHash(), 1);
    }
    SECTION("save stack twice, with values") {
        TuplePool pool;
        CheckpointStorage storage(current_path);
        std::vector<CodePoint> code;
        auto saver = MachineStateSaver(&storage, &pool, code);

        auto stack = MessageStack(&pool);

        uint256_t val_data = 111;
        uint256_t destination = 2;
        uint256_t currency = 3;
        auto token_type = std::array<unsigned char, 21>();
        token_type[0] = 'a';
        auto msg = Message{val_data, destination, currency, token_type};

        stack.addMessage(msg);
        auto results = stack.checkpointState(saver);
        auto results2 = stack.checkpointState(saver);

        getSavedMessageStack(saver, results.msgs_tuple_results.storage_key,
                             results.msg_count_results.storage_key,
                             stack.messages.calculateHash(), 1);
    }
}
