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

#include <iostream>

#include <avm/machine.hpp>
#include <avm_values/opcodes.hpp>

#include <data_storage/messageentry.hpp>

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

Assertion Machine::run(
    uint64_t gas_limit,
    bool hard_gas_limit,
    const std::vector<std::vector<unsigned char>>& inbox_data,
    const nonstd::optional<uint256_t>& final_block) {
    std::vector<Tuple> inbox_messages;
    inbox_messages.reserve(inbox_data.size());
    for (const auto& data : inbox_data) {
        inbox_messages.emplace_back(messageDataToTuple(data));
    }

    return run(gas_limit, hard_gas_limit, inbox_messages, final_block);
}

Assertion Machine::run(
    uint64_t gas_limit,
    bool hard_gas_limit,
    const std::vector<Tuple>& inbox_messages,
    const nonstd::optional<uint256_t>& min_next_block_height) {
    if (!validMessages(inbox_messages)) {
        throw std::runtime_error("invalid message format");
    }

    machine_state.context =
        AssertionContext{inbox_messages, min_next_block_height};

    bool has_gas_limit = gas_limit != 0;
    auto start_time = std::chrono::system_clock::now();
    while (true) {
        if (has_gas_limit) {
            if (hard_gas_limit) {
                if (machine_state.nextGasCost() + machine_state.context.numGas >
                    gas_limit) {
                    // Next step would go over gas limit
                    break;
                }
            } else if (machine_state.nextGasCost() >= gas_limit) {
                // Last step reached or went over gas limit
                break;
            }
        }

        auto blockReason = machine_state.runOne();
        if (!nonstd::get_if<NotBlocked>(&blockReason)) {
            break;
        }
    }
    return {machine_state.context.numSteps,
            machine_state.context.numGas,
            machine_state.context.inbox_messages_consumed,
            std::move(machine_state.context.sends),
            std::move(machine_state.context.logs),
            std::move(machine_state.context.debug_prints)};
}
