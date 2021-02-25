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

#include <avm/inboxmessage.hpp>
#include <avm/machine.hpp>
#include <avm_values/opcodes.hpp>

std::ostream& operator<<(std::ostream& os, const Machine& val) {
    os << val.machine_state;
    return os;
}

namespace {
template <typename T>
void convertInboxMessagesFromBytes(
    const std::vector<std::vector<unsigned char>>& bytes,
    T& output) {
    for (const auto& data : bytes) {
        auto message = extractInboxMessage(data);
        output.emplace_back(message);
    }
}
}  // namespace

MachineExecutionConfig::MachineExecutionConfig()
    : max_gas(0),
      go_over_gas(false),
      inbox_messages(),
      messages_to_skip(0),
      final_message_of_block(false),
      sideloads(),
      stop_on_sideload(false) {}

void MachineExecutionConfig::setInboxMessagesFromBytes(
    const std::vector<std::vector<unsigned char>>& bytes) {
    inbox_messages.clear();
    inbox_messages.reserve(bytes.size());
    convertInboxMessagesFromBytes(bytes, inbox_messages);
}

void MachineExecutionConfig::setSideloadsFromBytes(
    const std::vector<std::vector<unsigned char>>& bytes) {
    sideloads.clear();
    convertInboxMessagesFromBytes(bytes, sideloads);
}

Assertion Machine::run(MachineExecutionConfig config) {
    auto config_max_gas = config.max_gas;
    auto config_go_over_gas = config.go_over_gas;
    bool has_gas_limit = config.max_gas != 0;

    machine_state.context = AssertionContext(std::move(config));

    BlockReason block_reason = NotBlocked{};
    while (true) {
        if (has_gas_limit) {
            if (!config_go_over_gas) {
                if (machine_state.nextGasCost() + machine_state.context.numGas >
                    config_max_gas) {
                    // Next step would go over gas limit
                    break;
                }
            } else if (machine_state.context.numGas >= config_max_gas) {
                // Last step reached or went over gas limit
                break;
            }
        }

        block_reason = machine_state.runOne();
        if (!std::get_if<NotBlocked>(&block_reason)) {
            break;
        }
    }
    std::optional<uint256_t> sideload_block_number;
    if (auto sideload_blocked = std::get_if<SideloadBlocked>(&block_reason)) {
        sideload_block_number = sideload_blocked->block_number;
    }
    return {intx::narrow_cast<uint64_t>(machine_state.context.numSteps),
            intx::narrow_cast<uint64_t>(machine_state.context.numGas),
            machine_state.context.inbox_messages_consumed,
            std::move(machine_state.context.sends),
            std::move(machine_state.context.logs),
            std::move(machine_state.context.debug_prints),
            sideload_block_number};
}
