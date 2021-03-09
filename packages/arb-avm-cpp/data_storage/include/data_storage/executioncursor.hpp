/*
 * Copyright 2020, Offchain Labs, Inc.
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

#ifndef data_storage_executioncursor_hpp
#define data_storage_executioncursor_hpp

#include <avm/machinestate/status.hpp>
#include <avm/machinethread.hpp>
#include <avm_values/bigint.hpp>
#include <avm_values/codepointstub.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/machine.hpp>
#include <utility>

class ExecutionCursor {
   public:
    uint256_t first_message_sequence_number;
    std::vector<InboxMessage> messages;
    std::vector<uint256_t> inbox_accumulators;
    size_t messages_to_skip{0};
    std::variant<MachineStateKeys, std::unique_ptr<Machine>> machine;

   public:
    ExecutionCursor(MachineStateKeys machine_)
        : first_message_sequence_number(
              machine_.output.fully_processed_messages),
          machine(std::move(machine_)) {}

    //    ExecutionCursor(std::unique_ptr<Machine>& machine,
    //                    std::vector<InboxMessage>& messages,
    //                    std::vector<uint256_t>& inbox_accumulators,
    //                    size_t messages_to_skip)
    //        : machine(std::move(machine)),
    //          first_message_sequence_number(machine->machine_state.output.fully_processed_messages),
    //          messages(std::move(messages)),
    //          inbox_accumulators(std::move(inbox_accumulators)),
    //          messages_to_skip(messages_to_skip) {}

    ~ExecutionCursor() = default;

    ExecutionCursor(const ExecutionCursor& rhs)
        : machine(std::unique_ptr<Machine>(nullptr)) {
        if (std::holds_alternative<std::unique_ptr<Machine>>(machine)) {
            machine = std::make_unique<Machine>(
                *std::get<std::unique_ptr<Machine>>(machine));
        } else {
            machine = std::get<MachineStateKeys>(machine);
        }

        first_message_sequence_number = rhs.first_message_sequence_number;
        messages = rhs.messages;
        inbox_accumulators = rhs.inbox_accumulators;
        messages_to_skip = rhs.messages_to_skip;
    }

    ExecutionCursor& operator=(const ExecutionCursor& rhs) {
        if (std::holds_alternative<std::unique_ptr<Machine>>(machine)) {
            machine = std::make_unique<Machine>(
                *std::get<std::unique_ptr<Machine>>(machine));
        } else {
            machine = std::get<MachineStateKeys>(machine);
        }

        first_message_sequence_number = rhs.first_message_sequence_number;
        messages = rhs.messages;
        inbox_accumulators = rhs.inbox_accumulators;
        messages_to_skip = rhs.messages_to_skip;

        return *this;
    }

    ExecutionCursor* clone();

    std::optional<uint256_t> machineHash() const {
        if (std::holds_alternative<std::unique_ptr<Machine>>(machine)) {
            return std::get<std::unique_ptr<Machine>>(machine)->hash();
        } else {
            return std::get<MachineStateKeys>(machine).machineHash();
        }
    }

    const MachineOutput& getOutput() const {
        if (std::holds_alternative<std::unique_ptr<Machine>>(machine)) {
            return std::get<std::unique_ptr<Machine>>(machine)
                ->machine_state.output;
        } else {
            return std::get<MachineStateKeys>(machine).output;
        }
    }

    const staged_variant& getStaged() const {
        if (std::holds_alternative<std::unique_ptr<Machine>>(machine)) {
            return std::get<std::unique_ptr<Machine>>(machine)
                ->machine_state.staged_message;
        } else {
            return std::get<MachineStateKeys>(machine).staged_message;
        }
    }

    staged_variant& getStaged() {
        if (std::holds_alternative<std::unique_ptr<Machine>>(machine)) {
            return std::get<std::unique_ptr<Machine>>(machine)
                ->machine_state.staged_message;
        } else {
            return std::get<MachineStateKeys>(machine).staged_message;
        }
    }

    std::optional<uint256_t> getInboxAcc() const {
        auto fully_processed_acc =
            getOutput().fully_processed_inbox_accumulator;
        auto& staged_message = getStaged();
        if (std::holds_alternative<InboxMessage>(staged_message)) {
            return hash_inbox(
                fully_processed_acc,
                std::get<InboxMessage>(staged_message).serialize());
        } else if (std::holds_alternative<std::monostate>(staged_message)) {
            return fully_processed_acc;
        } else {
            return std::nullopt;
        }
    }

    uint256_t getTotalMessagesRead() const {
        auto fully_processed_messages = getOutput().fully_processed_messages;
        auto& staged_message = getStaged();
        if (std::holds_alternative<std::monostate>(staged_message)) {
            return fully_processed_messages;
        } else {
            return fully_processed_messages + 1;
        }
    }
};

#endif /* data_storage_executioncursor_hpp */
