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
    uint256_t max_gas,
    bool go_over_gas,
    const std::vector<std::vector<unsigned char>>& inbox_data,
    uint256_t messages_to_skip,
    bool final_message_of_block) {
    std::vector<Tuple> inbox_messages;
    inbox_messages.reserve(inbox_messages.size());
    for (const auto& data : inbox_data) {
        inbox_messages.emplace_back(messageDataToTuple(data));
    }

    return run(max_gas, go_over_gas, inbox_messages, messages_to_skip,
               final_message_of_block);
}

Assertion Machine::run(uint256_t max_gas,
                       bool go_over_gas,
                       const std::vector<Tuple>& inbox_messages,
                       uint256_t messages_to_skip,
                       bool final_message_of_block) {
    if (!validMessages(inbox_messages)) {
        throw std::runtime_error("invalid message format");
    }

    nonstd::optional<uint256_t> min_next_block_height;
    if (final_message_of_block && !inbox_messages.empty()) {
        // Last message is the final message of a block, so need to
        // set min_next_block_height to the block after the last block
        auto block_num =
            inbox_messages[inbox_messages.size() - 1].get_element(1);
        if (!nonstd::holds_alternative<uint256_t>(block_num)) {
            throw std::runtime_error("Cannot get final block from tuple");
        }

        min_next_block_height = nonstd::get<uint256_t>(block_num) + 1;
    }

    machine_state.context = AssertionContext{
        inbox_messages, min_next_block_height, messages_to_skip};

    bool has_gas_limit = max_gas != 0;
    auto start_time = std::chrono::system_clock::now();
    while (true) {
        if (has_gas_limit) {
            if (!go_over_gas) {
                if (machine_state.nextGasCost() + machine_state.context.numGas >
                    max_gas) {
                    // Next step would go over gas limit
                    break;
                }
            } else if (machine_state.nextGasCost() >= max_gas) {
                // Last step reached or went over gas limit
                break;
            }
        }

        auto blockReason = machine_state.runOne();
        if (!nonstd::get_if<NotBlocked>(&blockReason)) {
            break;
        }
    }
    return {intx::narrow_cast<uint64_t>(machine_state.context.numSteps),
            intx::narrow_cast<uint64_t>(machine_state.context.numGas),
            machine_state.context.inbox_messages_consumed,
            std::move(machine_state.context.sends),
            std::move(machine_state.context.logs),
            std::move(machine_state.context.debug_prints)};
}
