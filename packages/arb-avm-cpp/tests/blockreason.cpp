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

#include <avm/machinestate/blockreason.hpp>
#include <data_storage/checkpoint/checkpointutils.hpp>

#include <catch2/catch.hpp>

void serializeBlockReason(BlockReason& block_reason, BlockType expected_type) {
    auto serialized = serializeForCheckpoint(block_reason);
    auto type = serialized[0];
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
