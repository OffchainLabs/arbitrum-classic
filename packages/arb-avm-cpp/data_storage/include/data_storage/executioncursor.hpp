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
    std::variant<MachineStateKeys, std::unique_ptr<Machine>> machine;

   public:
    explicit ExecutionCursor(MachineStateKeys machine_)
        : machine(std::move(machine_)) {}

    ~ExecutionCursor() = default;

    ExecutionCursor(const ExecutionCursor& rhs)
        : machine(std::unique_ptr<Machine>(nullptr)) {
        if (std::holds_alternative<std::unique_ptr<Machine>>(rhs.machine)) {
            machine = std::make_unique<Machine>(
                *std::get<std::unique_ptr<Machine>>(rhs.machine));
        } else {
            machine = std::get<MachineStateKeys>(rhs.machine);
        }
    }

    ExecutionCursor& operator=(const ExecutionCursor& rhs) {
        if (std::holds_alternative<std::unique_ptr<Machine>>(rhs.machine)) {
            machine = std::make_unique<Machine>(
                *std::get<std::unique_ptr<Machine>>(rhs.machine));
        } else {
            machine = std::get<MachineStateKeys>(rhs.machine);
        }
        return *this;
    }

    ExecutionCursor* clone();

    [[nodiscard]] std::optional<uint256_t> machineHash() const {
        if (std::holds_alternative<std::unique_ptr<Machine>>(machine)) {
            return std::get<std::unique_ptr<Machine>>(machine)->hash();
        } else {
            return std::get<MachineStateKeys>(machine).machineHash();
        }
    }

    [[nodiscard]] const MachineOutput& getOutput() const {
        if (std::holds_alternative<std::unique_ptr<Machine>>(machine)) {
            return std::get<std::unique_ptr<Machine>>(machine)
                ->machine_state.output;
        } else {
            return std::get<MachineStateKeys>(machine).output;
        }
    }

    [[nodiscard]] std::optional<uint256_t> getInboxAcc() const {
        return getOutput().fully_processed_inbox.accumulator;
    }

    [[nodiscard]] uint256_t getTotalMessagesRead() const {
        return getOutput().fully_processed_inbox.count;
    }
};

#endif /* data_storage_executioncursor_hpp */
