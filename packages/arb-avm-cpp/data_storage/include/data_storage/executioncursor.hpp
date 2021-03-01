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
#include <data_storage/checkpoint.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresultfwd.hpp>
#include <data_storage/value/machine.hpp>
#include <utility>

class ExecutionCursor : public Checkpoint {
   public:
    std::unique_ptr<Machine> machine;
    uint256_t first_message_sequence_number;
    std::vector<InboxMessage> messages;
    std::vector<uint256_t> inbox_accumulators;
    size_t messages_to_skip{0};

   public:
    ExecutionCursor() = default;
    ExecutionCursor(Checkpoint& checkpoint,
                    std::unique_ptr<Machine>& machine,
                    std::vector<InboxMessage>& messages,
                    std::vector<uint256_t>& inbox_accumulators,
                    size_t messages_to_skip)
        : Checkpoint(checkpoint),
          machine(std::move(machine)),
          first_message_sequence_number(checkpoint.total_messages_read),
          messages(std::move(messages)),
          inbox_accumulators(std::move(inbox_accumulators)),
          messages_to_skip(messages_to_skip) {}
    ~ExecutionCursor() = default;
    ExecutionCursor(const ExecutionCursor& rhs) : Checkpoint(rhs) {
        machine = std::make_unique<Machine>(*rhs.machine);
        first_message_sequence_number = rhs.first_message_sequence_number;
        messages = rhs.messages;
        inbox_accumulators = rhs.inbox_accumulators;
        messages_to_skip = rhs.messages_to_skip;
    }
    ExecutionCursor& operator=(const ExecutionCursor& rhs) {
        Checkpoint::operator=(rhs);
        machine = std::make_unique<Machine>(*rhs.machine);
        first_message_sequence_number = rhs.first_message_sequence_number;
        messages = rhs.messages;
        inbox_accumulators = rhs.inbox_accumulators;
        messages_to_skip = rhs.messages_to_skip;

        return *this;
    }

    void resetExecutionCursor();
    void setCheckpoint(Checkpoint& checkpoint);
    ExecutionCursor* clone();
    uint256_t machineHash();

    std::unique_ptr<Machine> takeMachine();
};

#endif /* data_storage_executioncursor_hpp */
