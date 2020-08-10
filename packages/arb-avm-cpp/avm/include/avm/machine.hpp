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

#include <avm/machinestate/machinestate.hpp>
#include <avm_values/value.hpp>

#include <chrono>
#include <memory>
#include <vector>

struct Assertion {
    uint64_t stepCount;
    uint64_t gasCount;
    uint64_t inbox_messages_consumed;
    std::vector<value> outMessages;
    std::vector<value> logs;
};

class Machine {
    friend std::ostream& operator<<(std::ostream&, const Machine&);

    Assertion executeMachine(uint64_t stepCount,
                             std::chrono::seconds wallLimit,
                             std::vector<Tuple> inbox_messages,
                             Tuple sideload,
                             bool blockingSideload_,
                             nonstd::optional<value> fake_inbox_peek_value_);

   public:
    MachineState machine_state;

    Machine() = default;
    Machine(MachineState machine_state_)
        : machine_state(std::move(machine_state_)) {}
    Machine(std::shared_ptr<Code> code, value static_val)
        : machine_state(std::move(code), std::move(static_val)) {}

    static Machine loadFromFile(const std::string& executable_filename) {
        return {MachineState::loadFromFile(executable_filename)};
    }

    Assertion runSideloaded(uint64_t stepCount,
                            std::vector<Tuple> inbox_messages,
                            std::chrono::seconds wallLimit,
                            Tuple sideload);

    Assertion run(uint64_t stepCount,
                  std::vector<Tuple> inbox_messages,
                  std::chrono::seconds wallLimit);

    Assertion runCallServer(uint64_t stepCount,
                            std::vector<Tuple> inbox_messages,
                            std::chrono::seconds wallLimit,
                            value fake_inbox_peek_value);

    Status currentStatus() { return machine_state.state; }
    uint256_t hash() const { return machine_state.hash(); }
    BlockReason isBlocked(bool newMessages) const {
        return machine_state.isBlocked(newMessages);
    }
    std::vector<unsigned char> marshalForProof() {
        return machine_state.marshalForProof();
    }

    std::vector<unsigned char> marshalState() const {
        return machine_state.marshalState();
    }
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
