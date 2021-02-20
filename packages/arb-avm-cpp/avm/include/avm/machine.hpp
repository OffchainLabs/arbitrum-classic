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
    uint64_t stepCount;
    uint64_t gasCount;
    uint64_t inbox_messages_consumed;
    std::vector<std::vector<uint8_t>> sends;
    std::vector<value> logs;
    std::vector<value> debugPrints;
    std::optional<uint256_t> sideloadBlockNumber;
};

class MachineExecutionConfig {
   public:
    uint256_t max_gas;
    bool go_over_gas;
    std::vector<InboxMessage> inbox_messages;
    uint256_t messages_to_skip;
    bool final_message_of_block;
    std::deque<InboxMessage> sideloads;
    bool stop_on_sideload;

    MachineExecutionConfig();

    void setInboxMessagesFromBytes(
        const std::vector<std::vector<unsigned char>>&);
    void setSideloadsFromBytes(const std::vector<std::vector<unsigned char>>&);
};

class Machine {
    friend std::ostream& operator<<(std::ostream&, const Machine&);

   public:
    MachineState machine_state;

    Machine() = default;
    explicit Machine(MachineState machine_state_)
        : machine_state(std::move(machine_state_)) {}
    Machine(std::shared_ptr<Code> code, value static_val)
        : machine_state(std::move(code), std::move(static_val)) {}

    static Machine loadFromFile(const std::string& executable_filename) {
        return Machine{MachineState::loadFromFile(executable_filename)};
    }

    Assertion run(const MachineExecutionConfig& config);

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
