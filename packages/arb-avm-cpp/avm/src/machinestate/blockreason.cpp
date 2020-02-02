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

#include <avm_values/codepoint.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

constexpr BlockType InboxBlocked::type;

auto operator<<(std::ostream& os, const NotBlocked&) -> std::ostream& {
    return os << "NotBlocked";
}

auto operator<<(std::ostream& os, const HaltBlocked&) -> std::ostream& {
    return os << "HaltBlocked";
}

auto operator<<(std::ostream& os, const ErrorBlocked&) -> std::ostream& {
    return os << "ErrorBlocked";
}

auto operator<<(std::ostream& os, const BreakpointBlocked&) -> std::ostream& {
    return os << "BreakpointBlocked";
}

auto operator<<(std::ostream& os, const InboxBlocked& val) -> std::ostream& {
    return os << "InboxBlocked(" << val.timeout << ")";
}

auto operator<<(std::ostream& os, const BlockReason& val) -> std::ostream& {
    nonstd::visit([&](const auto& reason) { os << reason; }, val);
    return os;
}
