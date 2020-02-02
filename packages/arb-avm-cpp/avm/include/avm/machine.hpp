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
    MachineState machine_state;

    friend auto operator<<(std::ostream&, const Machine&) -> std::ostream&;
    auto runOne() -> BlockReason;

   public:
    auto initializeMachine(const std::string& filename) -> bool;
    void initializeMachine(const MachineState& initial_state);

    auto run(uint64_t stepCount,
             const TimeBounds& timeBounds,
             Tuple messages,
             std::chrono::seconds wallLimit) -> Assertion;

    auto currentStatus() -> Status { return machine_state.state; }
    auto hash() const -> uint256_t { return machine_state.hash(); }
    auto isBlocked(const uint256_t& currentTime, bool newMessages) const
        -> BlockReason {
        return machine_state.isBlocked(currentTime, newMessages);
    }
    auto marshalForProof() -> std::vector<unsigned char> {
        return machine_state.marshalForProof();
    }

    auto getPool() -> TuplePool& { return *machine_state.pool; }

    auto checkpoint(CheckpointStorage& storage) -> SaveResults;
    auto restoreCheckpoint(const CheckpointStorage& storage,
                           const std::vector<unsigned char>& checkpoint_key)
        -> bool;
    auto deleteCheckpoint(CheckpointStorage& storage) -> DeleteResults;
};

auto operator<<(std::ostream& os, const MachineState& val) -> std::ostream&;
auto operator<<(std::ostream& os, const Machine& val) -> std::ostream&;

#endif /* machine_hpp */
