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

#ifndef machinestate_hpp
#define machinestate_hpp

#include <avm/machinestate/blockreason.hpp>
#include <avm/machinestate/datastack.hpp>
#include <avm_values/value.hpp>
#include <data_storage/storageresult.hpp>

#include <memory>
#include <utility>

#include <vector>

enum class Status { Extensive, Halted, Error };

using TimeBounds = std::array<uint256_t, 2>;

class CheckpointStorage;

struct AssertionContext {
    TimeBounds timeBounds;
    Tuple inbox;
    uint32_t numSteps;
    bool didInboxInsn;
    uint64_t numGas;
    std::vector<value> outMessage;
    std::vector<value> logs;

    explicit AssertionContext(TimeBounds tb, Tuple inbox)
        : timeBounds(std::move(tb)),
          inbox(std::move(inbox)),
          numSteps{0},
          didInboxInsn(false),
          numGas{0} {}

    void executedInbox() {
        didInboxInsn = true;
        inbox = Tuple();
    }
};

struct MachineState {
    std::shared_ptr<TuplePool> pool;
    std::vector<CodePoint> code;
    value staticVal;
    value registerVal;
    Datastack stack;
    Datastack auxstack;
    Status state = Status::Extensive;
    uint64_t pc = 0;
    CodePoint errpc;
    AssertionContext context;

    MachineState();
    MachineState(const std::vector<CodePoint>& code_,
                 const value& static_val_,
                 std::shared_ptr<TuplePool> pool_);
    auto initialize_machinestate(const std::string& contract_filename) -> bool;

    auto marshalForProof() -> std::vector<unsigned char>;
    auto runOp(OpCode opcode) -> BlockReason;
    auto hash() const -> uint256_t;
    auto isBlocked(const uint256_t& currentTime, bool newMessages) const
        -> BlockReason;
    auto checkpointState(CheckpointStorage& storage) const -> SaveResults;
    auto restoreCheckpoint(const CheckpointStorage& storage,
                           const std::vector<unsigned char>& checkpoint_key)
        -> bool;
};

#endif /* machinestate_hpp */
