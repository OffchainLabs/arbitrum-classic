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

#include "avm/machinestate/tokenTracker.hpp"
#include <catch2/catch.hpp>

void serializeBlockReason(BlockReason& block_reason, BlockType expected_type) {
    auto serialized = serializeForCheckpoint(block_reason);
    auto type = (BlockType)serialized[0];
    REQUIRE(type == expected_type);
    REQUIRE(serialized.size() == blockreason_type_length[expected_type]);
}

void deserializeInboxBlocked(std::vector<unsigned char> serialized,
                             uint256_t expected_inbox) {
    auto deserialized = deserializeBlockReason(serialized);

    auto inbox_block = nonstd::get<InboxBlocked>(deserialized);
    REQUIRE(inbox_block.type == Inbox);
    REQUIRE(inbox_block.inbox == expected_inbox);
}

void deserializeSendBlocked(std::vector<unsigned char> serialized,
                            uint256_t expected_currency,
                            TokenType expected_token_type) {
    auto deserialized = deserializeBlockReason(serialized);

    auto inbox_block = nonstd::get<SendBlocked>(deserialized);
    REQUIRE(inbox_block.type == Send);
    REQUIRE(inbox_block.currency == expected_currency);
    REQUIRE(inbox_block.tokenType == expected_token_type);
}

TEST_CASE("Serialize blockreason") {
    SECTION("NotBlocked") {
        BlockReason not_blocked = NotBlocked();
        serializeBlockReason(not_blocked, Not);
    }
    SECTION("HaltBlocked") {
        BlockReason halt_blocked = HaltBlocked();
        serializeBlockReason(halt_blocked, Halt);
    }
    SECTION("ErrorBlocked") {
        BlockReason error_blocked = ErrorBlocked();
        serializeBlockReason(error_blocked, Error);
    }
    SECTION("BreakpointBlocked") {
        BlockReason breakpoint_blocked = BreakpointBlocked();
        serializeBlockReason(breakpoint_blocked, Breakpoint);
    }
    SECTION("BreakpointBlocked") {
        BlockReason inbox_blocked = InboxBlocked();
        serializeBlockReason(inbox_blocked, Inbox);
    }
    SECTION("BreakpointBlocked") {
        BlockReason send_blocked = SendBlocked();
        serializeBlockReason(send_blocked, Send);
    }
}

TEST_CASE("deserialize inbox blocked") {
    SECTION("0 inbox") {
        auto inbox_blocked = InboxBlocked();
        auto serialized = serializeForCheckpoint(inbox_blocked);
        deserializeInboxBlocked(serialized, 0);
    }
    SECTION("inbox with value") {
        auto inbox_blocked = InboxBlocked(100);
        auto serialized = serializeForCheckpoint(inbox_blocked);
        deserializeInboxBlocked(serialized, 100);
    }
}

TEST_CASE("deserialize send blocked") {
    SECTION("default") {
        auto send_blocked = SendBlocked();
        auto serialized = serializeForCheckpoint(send_blocked);

        deserializeSendBlocked(serialized, 0, std::array<unsigned char, 21>());
    }
    SECTION("with values") {
        std::array<unsigned char, 21> token_type = {10};
        auto send_blocked = SendBlocked(999, token_type);
        auto serialized = serializeForCheckpoint(send_blocked);
        deserializeSendBlocked(serialized, 999, token_type);
    }
}

void serializedBalancaVals(BalanceTracker& balance_tracker,
                           int expected_length,
                           unsigned int expected_pairs) {
    auto serialized = balance_tracker.serializeBalanceValues();

    unsigned int balance_tracker_length;
    auto current_iter = serialized.begin();
    memcpy(&balance_tracker_length, &(*current_iter), sizeof(expected_pairs));

    REQUIRE(serialized.size() == expected_length);
    REQUIRE(balance_tracker_length == expected_pairs);
}

TEST_CASE("serialize balance tracker") {
    SECTION("default") {
        auto tracker = BalanceTracker();

        serializedBalancaVals(tracker, sizeof(unsigned int), 0);
    }
    SECTION("deafult pair") {
        std::array<unsigned char, 21> token_type = {};
        uint256_t amount = 0;
        auto tracker = BalanceTracker();
        tracker.add(token_type, amount);

        auto token_pair_length = TOKEN_VAL_LENGTH + TOKEN_TYPE_LENGTH;
        auto total_len = token_pair_length * 1 + sizeof(unsigned int);

        serializedBalancaVals(tracker, total_len, 1);
    }
    SECTION("pair with values") {
        std::array<unsigned char, 21> token_type = {1, 99};
        uint256_t amount = 11;
        auto tracker = BalanceTracker();
        tracker.add(token_type, amount);

        auto token_pair_length = TOKEN_VAL_LENGTH + TOKEN_TYPE_LENGTH;
        auto total_len = token_pair_length * 1 + sizeof(unsigned int);

        serializedBalancaVals(tracker, total_len, 1);
    }
    SECTION("multiple pair with values") {
        std::array<unsigned char, 21> token_type1 = {1, 99};
        uint256_t amount1 = 11;
        std::array<unsigned char, 21> token_type2 = {3, 9};
        uint256_t amount2 = 2;
        auto tracker = BalanceTracker();
        tracker.add(token_type1, amount1);
        tracker.add(token_type2, amount2);

        auto token_pair_length = TOKEN_VAL_LENGTH + TOKEN_TYPE_LENGTH;
        auto total_len = token_pair_length * 2 + sizeof(unsigned int);

        serializedBalancaVals(tracker, total_len, 2);
    }
}

void intializeTracker(
    std::vector<unsigned char> data,
    std::unordered_map<TokenType, uint256_t> expected_lookup) {
    BalanceTracker tracker(data);

    for (auto& pair : expected_lookup) {
        auto val = tracker.tokenValue(pair.first);
        REQUIRE(val == pair.second);
    }
}

TEST_CASE("deserialize and initialize tracker") {
    SECTION("default") {
        auto tracker = BalanceTracker();
        std::unordered_map<TokenType, uint256_t> expected_lookup;

        intializeTracker(tracker.serializeBalanceValues(), expected_lookup);
    }
    SECTION("deafult pair") {
        std::array<unsigned char, 21> token_type = {};
        uint256_t amount = 0;
        auto tracker = BalanceTracker();
        tracker.add(token_type, amount);
        std::unordered_map<TokenType, uint256_t> expected_lookup;

        expected_lookup[token_type] = amount;

        intializeTracker(tracker.serializeBalanceValues(), expected_lookup);
    }
    SECTION("pair with values") {
        std::array<unsigned char, 21> token_type = {1, 99};
        uint256_t amount = 11;
        auto tracker = BalanceTracker();
        tracker.add(token_type, amount);

        std::unordered_map<TokenType, uint256_t> expected_lookup;
        expected_lookup[token_type] = amount;

        intializeTracker(tracker.serializeBalanceValues(), expected_lookup);
    }
    SECTION("multiple pair with values") {
        std::array<unsigned char, 21> token_type1 = {1, 99};
        uint256_t amount1 = 11;
        std::array<unsigned char, 21> token_type2 = {3, 9};
        uint256_t amount2 = 2;
        auto tracker = BalanceTracker();
        tracker.add(token_type1, amount1);
        tracker.add(token_type2, amount2);

        std::unordered_map<TokenType, uint256_t> expected_lookup;
        expected_lookup[token_type1] = amount1;
        expected_lookup[token_type2] = amount2;

        intializeTracker(tracker.serializeBalanceValues(), expected_lookup);
    }
}
