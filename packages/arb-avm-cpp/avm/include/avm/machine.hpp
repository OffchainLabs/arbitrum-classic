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
    MachineState machine_state;

    friend std::ostream& operator<<(std::ostream&, const Machine&);
    void runOne();

   public:
    bool initializeMachine(const std::string& filename);
    void initializeMachine(const MachineState& initial_state);

    Assertion run(uint64_t stepCount,
                  uint256_t timeBoundStart,
                  uint256_t timeBoundEnd);

    Status currentStatus() { return machine_state.state; }
    BlockReason lastBlockReason() { return machine_state.blockReason; }
    uint256_t hash() const { return machine_state.hash(); }
    std::vector<unsigned char> marshalForProof() {
        return machine_state.marshalForProof();
    }

    uint256_t inboxHash() const { return ::hash(machine_state.inbox.messages); }

    void deliverMessages(Tuple messages);

    TuplePool& getPool() { return *machine_state.pool; }

    SaveResults checkpoint(CheckpointStorage& storage);
    bool restoreCheckpoint(const CheckpointStorage& storage,
                           const std::vector<unsigned char>& checkpoint_key);
    DeleteResults deleteCheckpoint(CheckpointStorage& storage);
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
