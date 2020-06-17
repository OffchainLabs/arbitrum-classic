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

#ifndef machinestate_hpp
#define machinestate_hpp

#include <avm/machinestate/blockreason.hpp>
#include <avm/machinestate/datastack.hpp>
#include <avm/machinestate/status.hpp>
#include <avm_values/value.hpp>
#include <avm_values/vmValueParser.hpp>

#include <memory>
#include <vector>

struct TimeBounds {
    uint256_t lowerBoundBlock;
    uint256_t upperBoundBlock;
    uint256_t lowerBoundTimestamp;
    uint256_t upperBoundTimestamp;
};

struct AssertionContext {
    TimeBounds timeBounds;
    Tuple inbox;
    uint32_t numSteps;
    bool didInboxInsn;
    uint64_t numGas;
    std::vector<value> outMessage;
    std::vector<value> logs;

    AssertionContext() = default;

    explicit AssertionContext(const TimeBounds& tb, Tuple inbox)
        : timeBounds(tb),
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
    std::shared_ptr<const StaticVmValues> static_values;
    value registerVal;
    Datastack stack;
    Datastack auxstack;
    Status state = Status::Extensive;
    CodePointRef pc;
    CodePointRef errpc;
    AssertionContext context;

    static std::pair<MachineState, bool> loadFromFile(
        const std::string& contract_filename);

    MachineState()
        : pool(std::make_unique<TuplePool>()), pc(0, false), errpc(0, true) {}

    MachineState(std::shared_ptr<const StaticVmValues> static_values_,
                 std::shared_ptr<TuplePool> pool_)
        : pool(std::move(pool_)),
          static_values(std::move(static_values_)),
          pc(0, false),
          errpc(0, true) {}

    MachineState(std::shared_ptr<TuplePool> pool_,
                 std::shared_ptr<const StaticVmValues> static_values_,
                 value register_val_,
                 Datastack stack_,
                 Datastack auxstack_,
                 Status state_,
                 CodePointRef pc_,
                 CodePointRef errpc_)
        : pool(std::move(pool_)),
          static_values(std::move(static_values_)),
          registerVal(std::move(register_val_)),
          stack(std::move(stack_)),
          auxstack(std::move(auxstack_)),
          state(state_),
          pc(pc_),
          errpc(errpc_) {}

    uint256_t getMachineSize();
    std::vector<unsigned char> marshalForProof();
    BlockReason runOp(OpCode opcode);
    uint256_t hash() const;
    BlockReason isBlocked(uint256_t currentTime, bool newMessages) const;
};

#endif /* machinestate_hpp */
