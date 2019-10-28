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

#ifndef blockreason_hpp
#define blockreason_hpp

#include <avm/machinestate/tokenTracker.hpp>

enum BlockType { Not, Halt, Error, Breakpoint, Inbox, Send };

struct NotBlocked {
    static constexpr BlockType type = Not;
};

struct HaltBlocked {
    static constexpr BlockType type = Halt;
};

struct ErrorBlocked {
    static constexpr BlockType type = Error;
};

struct BreakpointBlocked {
    static constexpr BlockType type = Breakpoint;
};

struct InboxBlocked {
    static constexpr BlockType type = Inbox;
    uint256_t inbox;
    InboxBlocked() {}

    InboxBlocked(uint256_t inbox_) { inbox = inbox_; }
};

struct SendBlocked {
    static constexpr BlockType type = Send;
    uint256_t currency;
    TokenType tokenType;
    SendBlocked() { tokenType = {}; }

    SendBlocked(uint256_t currency_, TokenType tokenType_) {
        currency = currency_;
        tokenType = tokenType_;
    }
};

extern std::unordered_map<BlockType, int> blockreason_type_length;

using BlockReason = nonstd::variant<NotBlocked,
                                    HaltBlocked,
                                    ErrorBlocked,
                                    BreakpointBlocked,
                                    InboxBlocked,
                                    SendBlocked>;

std::vector<unsigned char> serializeForCheckpoint(const BlockReason& val);
BlockReason deserializeBlockReason(std::vector<unsigned char>& data);

#endif /* blockreason_hpp */
