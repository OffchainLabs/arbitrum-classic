/*
 * Copyright 2019, Offchain Labs, Inc.
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
    std::vector<value> outMessages;
    std::vector<value> logs;
    bool didInboxInsn;
};

class Machine {
    friend std::ostream& operator<<(std::ostream&, const Machine&);
    BlockReason runOne();

   public:
    MachineState machine_state;

    Machine() = default;
    Machine(MachineState machine_state_)
        : machine_state(std::move(machine_state_)) {}
    Machine(const Code& code_,
            const value& static_val_,
            std::shared_ptr<TuplePool> pool_)
        : machine_state(code_, static_val_, std::move(pool_)) {}

    static std::pair<Machine, bool> loadFromFile(
        const std::string& contract_filename) {
        auto result = MachineState::loadFromFile(contract_filename);
        if (!result.second) {
            return std::make_pair(Machine{}, false);
        }
        return std::make_pair(Machine{std::move(result.first)}, true);
    }

    Assertion run(uint64_t stepCount,
                  const TimeBounds& timeBounds,
                  Tuple messages,
                  std::chrono::seconds wallLimit);

    Status currentStatus() { return machine_state.state; }
    uint256_t hash() const { return machine_state.hash(); }
    BlockReason isBlocked(uint256_t currentTime, bool newMessages) const {
        return machine_state.isBlocked(currentTime, newMessages);
    }
    std::vector<unsigned char> marshalForProof() {
        return machine_state.marshalForProof();
    }

    TuplePool& getPool() { return *machine_state.pool; }

    void marshal_value(const value& val, std::vector<unsigned char>& buf) {
        return ::marshal_value(val, buf, machine_state.code);
    }
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
