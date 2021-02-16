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

#include <data_storage/inboxmessage.hpp>

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

template <typename T>
void convertInboxMessagesFromBytes(
    const std::vector<std::vector<unsigned char>>& bytes,
    T& output) {
    for (const auto& data : bytes) {
        auto message = extractInboxMessage(data);
        output.emplace_back(message.toTuple());
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

Assertion Machine::run(const MachineExecutionConfig& config) {
    if (!validMessages(config.inbox_messages)) {
        throw std::runtime_error("invalid message format");
    }

    machine_state.context = AssertionContext(config);

    bool has_gas_limit = config.max_gas != 0;
    BlockReason block_reason = NotBlocked{};
    while (true) {
        if (has_gas_limit) {
            if (!config.go_over_gas) {
                if (machine_state.nextGasCost() + machine_state.context.numGas >
                    config.max_gas) {
                    // Next step would go over gas limit
                    break;
                }
            } else if (machine_state.context.numGas >= config.max_gas) {
                // Last step reached or went over gas limit
                break;
            }
        }

        block_reason = machine_state.runOne();
        if (!nonstd::get_if<NotBlocked>(&block_reason)) {
            break;
        }
    }
    nonstd::optional<uint256_t> sideload_block_number;
    if (auto sideload_blocked =
            nonstd::get_if<SideloadBlocked>(&block_reason)) {
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
