/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

#include <sys/stat.h>
#include <fstream>
#include <iostream>

#include <avm/machine.hpp>
#include <avm_values/opcodes.hpp>

std::ostream& operator<<(std::ostream& os, const Machine& val) {
    os << val.machine_state;
    return os;
}

namespace {
bool validMessages(const std::vector<Tuple>& messages) {
    for (const auto& msg : messages) {
        if (msg.tuple_size() < 2) {
            return false;
        }
        if (!nonstd::holds_alternative<uint256_t>(msg.get_element(1))) {
            return false;
        }
    }
    return true;
}
}  // namespace

Assertion Machine::executeMachine(
    uint64_t stepCount,
    std::chrono::seconds wallLimit,
    std::vector<Tuple> inbox_messages,
    Tuple sideload,
    bool blockingSideload,
    nonstd::optional<value> fake_inbox_peak_value) {
    if (!validMessages(inbox_messages)) {
        throw std::runtime_error("invalid message format");
    }

    machine_state.context =
        AssertionContext{std::move(inbox_messages), std::move(sideload),
                         blockingSideload, std::move(fake_inbox_peak_value)};

    bool has_time_limit = wallLimit.count() != 0;
    auto start_time = std::chrono::system_clock::now();
    while (machine_state.context.numSteps < stepCount) {
        auto blockReason = machine_state.runOne();
        if (!nonstd::get_if<NotBlocked>(&blockReason)) {
            break;
        }
        if (has_time_limit && machine_state.context.numSteps % 10000 == 0) {
            auto end_time = std::chrono::system_clock::now();
            auto run_time = end_time - start_time;
            if (run_time >= wallLimit) {
                break;
            }
        }
    }
    return {machine_state.context.numSteps, machine_state.context.numGas,
            machine_state.context.inbox_messages_consumed,
            std::move(machine_state.context.outMessage),
            std::move(machine_state.context.logs)};
}

Assertion Machine::run(uint64_t stepCount,
                       std::vector<Tuple> inbox_messages,
                       std::chrono::seconds wallLimit) {
    return executeMachine(stepCount, wallLimit, std::move(inbox_messages),
                          Tuple(), false, nonstd::nullopt);
}

Assertion Machine::runCallServer(uint64_t stepCount,
                                 std::vector<Tuple> inbox_messages,
                                 std::chrono::seconds wallLimit,
                                 value fake_inbox_peak_value) {
    return executeMachine(stepCount, wallLimit, std::move(inbox_messages),
                          Tuple(), false, std::move(fake_inbox_peak_value));
}

Assertion Machine::runSideloaded(uint64_t stepCount,
                                 std::vector<Tuple> inbox_messages,
                                 std::chrono::seconds wallLimit,
                                 Tuple sideload_value) {
    return executeMachine(stepCount, wallLimit, std::move(inbox_messages),
                          std::move(sideload_value), true, nonstd::nullopt);
}
