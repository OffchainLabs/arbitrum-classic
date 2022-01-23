/*
 * Copyright 2019-2021, Offchain Labs, Inc.
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

#ifndef machine_hpp
#define machine_hpp

#include <avm/inboxmessage.hpp>
#include <avm/machinestate/machinestate.hpp>
#include <avm_values/value.hpp>

#include <rocksdb/slice.h>
#include <chrono>
#include <deque>
#include <memory>
#include <vector>

struct Assertion {
    uint64_t step_count;
    uint64_t gas_count;
    uint64_t inbox_messages_consumed;
    std::vector<MachineEmission<std::vector<uint8_t>>> sends;
    std::vector<MachineEmission<Value>> logs;
    std::vector<MachineEmission<Value>> debug_prints;
    std::optional<uint256_t> sideload_block_number;
};

class MachineExecutionConfig {
   public:
    uint256_t max_gas;
    bool go_over_gas;
    std::vector<MachineMessage> inbox_messages;
    std::deque<InboxMessage> sideloads;
    bool stop_on_sideload;
    bool stop_on_breakpoint;

    MachineExecutionConfig();

    void setInboxMessagesFromBytes(
        const std::vector<std::vector<unsigned char>>&);
    void setSideloadsFromBytes(const std::vector<std::vector<unsigned char>>&);
};

class Machine {
    friend std::ostream& operator<<(std::ostream&, const Machine&);

   private:
    std::atomic<bool> is_aborted{false};

   public:
    MachineState machine_state;

    Machine() = default;
    explicit Machine(MachineState machine_state_)
        : machine_state(std::move(machine_state_)) {}
    explicit Machine(const Machine& machine)
        : is_aborted(machine.is_aborted.load()),
          machine_state(machine.machine_state) {}
    explicit Machine(Machine&& machine) noexcept
        : is_aborted(machine.is_aborted.load()),
          machine_state(std::move(machine.machine_state)) {}
    Machine& operator=(const Machine& machine) {
        return *this = Machine(machine);
    }
    Machine& operator=(Machine&& machine) noexcept {
        is_aborted = machine.is_aborted.load();
        machine_state = std::move(machine.machine_state);
        return *this;
    }
    ~Machine() = default;

    static Machine loadFromFile(const std::string& executable_filename) {
        return Machine{MachineState::loadFromFile(executable_filename)};
    }

    void abort();

    Assertion run();

    Status currentStatus() const { return machine_state.state; }
    uint256_t hash() const { return machine_state.hash(); }
    BlockReason isBlocked(bool newMessages) const {
        return machine_state.isBlocked(newMessages);
    }
    OneStepProof marshalForProof() const {
        return machine_state.marshalForProof();
    }

    std::vector<unsigned char> marshalState() const {
        return machine_state.marshalState();
    }
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
