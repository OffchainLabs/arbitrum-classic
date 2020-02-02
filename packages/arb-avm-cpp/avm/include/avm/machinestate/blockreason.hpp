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

#include <avm_values/bigint.hpp>

#include <nonstd/variant.hpp>

#include <unordered_map>

enum BlockType { Not, Halt, Error, Breakpoint, Inbox };

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
    uint256_t timeout;

    InboxBlocked(uint256_t timeout_) : timeout(std::move(timeout_)) {}
};

using BlockReason = nonstd::variant<NotBlocked,
                                    HaltBlocked,
                                    ErrorBlocked,
                                    BreakpointBlocked,
                                    InboxBlocked>;

auto operator<<(std::ostream& os, const NotBlocked& val) -> std::ostream&;
auto operator<<(std::ostream& os, const HaltBlocked& val) -> std::ostream&;
auto operator<<(std::ostream& os, const ErrorBlocked& val) -> std::ostream&;
auto operator<<(std::ostream& os, const BreakpointBlocked& val)
    -> std::ostream&;
auto operator<<(std::ostream& os, const InboxBlocked& val) -> std::ostream&;
auto operator<<(std::ostream& os, const BlockReason& val) -> std::ostream&;

#endif /* blockreason_hpp */
