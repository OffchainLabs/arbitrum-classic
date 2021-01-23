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
   private:
    std::unique_ptr<Machine> machine;
    uint256_t first_message_sequence_number;
    std::vector<Tuple> messages;
    std::vector<uint256_t> inbox_hashes;
    size_t messages_to_skip{0};
    nonstd::optional<uint256_t> min_next_block_height;

   public:
    ExecutionCursor(Checkpoint& checkpoint,
                    std::unique_ptr<Machine>& machine,
                    std::vector<Tuple>& messages,
                    std::vector<uint256_t>& inbox_hashes,
                    nonstd::optional<uint256_t>& min_next_block_height)
        : Checkpoint(checkpoint),
          machine(std::move(machine)),
          first_message_sequence_number(checkpoint.total_messages_read),
          messages(std::move(messages)),
          inbox_hashes(std::move(inbox_hashes)),
          min_next_block_height(std::move(min_next_block_height)) {}
    ~ExecutionCursor() = default;
    ExecutionCursor(const ExecutionCursor& rhs) : Checkpoint(rhs) {
        machine = std::make_unique<Machine>(*rhs.machine);
        first_message_sequence_number = rhs.first_message_sequence_number;
        messages = rhs.messages;
        inbox_hashes = rhs.inbox_hashes;
        messages_to_skip = rhs.messages_to_skip;
        min_next_block_height = rhs.min_next_block_height;
    }
    ExecutionCursor& operator=(const ExecutionCursor& rhs) {
        Checkpoint::operator=(rhs);
        machine = std::make_unique<Machine>(*rhs.machine);
        first_message_sequence_number = rhs.first_message_sequence_number;
        messages = rhs.messages;
        inbox_hashes = rhs.inbox_hashes;
        messages_to_skip = rhs.messages_to_skip;
        min_next_block_height = rhs.min_next_block_height;

        return *this;
    }

    ExecutionCursor* clone();
    uint256_t machineHash();

    std::unique_ptr<Machine> TakeMachine();

    bool Advance(uint256_t max_gas, bool go_over_gas);
};

#endif /* data_storage_executioncursor_hpp */
