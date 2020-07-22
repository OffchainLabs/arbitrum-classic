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
    std::vector<value> outMessages;
    std::vector<value> logs;
    bool didInboxInsn;

    //    void updateAssertion(Assertion addition){
    //        stepCount += addition.stepCount;
    //        gasCount += addition.gasCount;
    //        didInboxInsn = addition.didInboxInsn;
    //        outMessages.insert(outMessages.end(),
    //        addition.outMessages.begin(), addition.outMessages.end());
    //        logs.insert(logs.end(), addition.logs.begin(),
    //        addition.logs.end());
    //    };
};

class Machine {
    friend std::ostream& operator<<(std::ostream&, const Machine&);

    Assertion executeMachine(uint64_t stepCount,
                             std::chrono::seconds wallLimit);

   public:
    MachineState machine_state;

    Machine() = default;
    Machine(MachineState machine_state_)
        : machine_state(std::move(machine_state_)) {}
    Machine(std::shared_ptr<Code> code,
            value static_val,
            std::shared_ptr<TuplePool> pool_)
        : machine_state(std::move(code),
                        std::move(static_val),
                        std::move(pool_)) {}

    static Machine loadFromFile(const std::string& executable_filename) {
        return {MachineState::loadFromFile(executable_filename)};
    }

    Assertion run(uint64_t stepCount,
                  Tuple messages,
                  std::chrono::seconds wallLimit,
                  Tuple sideload);

    Assertion runNormal(uint64_t stepCount,
                        Tuple messages,
                        std::chrono::seconds wallLimit);

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

    TuplePool& getPool() { return *machine_state.pool; }
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
