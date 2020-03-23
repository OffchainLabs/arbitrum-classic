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

#include <sys/stat.h>
#include <fstream>
#include <iostream>

#include <avm/machine.hpp>
#include <avm_values/opcodes.hpp>
#include <avm_values/util.hpp>
#include <bigint_utils.hpp>
#include <data_storage/checkpoint/checkpointstorage.hpp>
#include <data_storage/checkpoint/machinestatedeleter.hpp>

std::ostream& operator<<(std::ostream& os, const MachineState& val) {
    os << "status " << static_cast<int>(val.state) << "\n";
    os << "codePointHash " << to_hex_str(hash(val.code->code[val.pc])) << "\n";
    os << "stackHash " << to_hex_str(val.stack.hash()) << "\n";
    os << "auxStackHash " << to_hex_str(val.auxstack.hash()) << "\n";
    os << "registerHash " << to_hex_str(hash(val.registerVal)) << "\n";
    os << "staticHash " << to_hex_str(hash(val.code->staticVal)) << "\n";
    os << "errHandlerHash " << to_hex_str(hash(val.errpc)) << "\n";
    return os;
}

std::ostream& operator<<(std::ostream& os, const Machine& val) {
    os << val.machine_state;
    return os;
}

bool Machine::initializeMachine(const std::string& filename) {
    return machine_state.initialize_machinestate(filename);
}

void Machine::initializeMachine(const MachineState& initial_state) {
    machine_state = initial_state;
}

Assertion Machine::run(uint64_t stepCount,
                       uint256_t timeBoundStart,
                       uint256_t timeBoundEnd,
                       Tuple messages,
                       std::chrono::seconds wallLimit) {
    bool has_time_limit = wallLimit.count() != 0;
    auto start_time = std::chrono::system_clock::now();
    machine_state.context = AssertionContext{
        TimeBounds{{timeBoundStart, timeBoundEnd}}, std::move(messages)};
    while (machine_state.context.numSteps < stepCount) {
        auto blockReason = runOne();
        if (!nonstd::get_if<NotBlocked>(&blockReason)) {
            break;
        }
        if (has_time_limit && machine_state.context.numSteps % 10000 == 0) {
            auto end_time = std::chrono::system_clock::now();
            auto run_time = end_time - start_time;
            if (run_time >= wallLimit) {
                break;
            }
        }
    }
    return {machine_state.context.numSteps, machine_state.context.numGas,
            std::move(machine_state.context.outMessage),
            std::move(machine_state.context.logs),
            machine_state.context.didInboxInsn};
}

SaveResults Machine::checkpoint(CheckpointStorage& storage) {
    return machine_state.checkpointState(storage);
}

bool Machine::restoreCheckpoint(
    const CheckpointStorage& storage,
    const std::vector<unsigned char>& checkpoint_key) {
    return machine_state.restoreCheckpoint(storage, checkpoint_key);
}

DeleteResults Machine::deleteCheckpoint(CheckpointStorage& storage) {
    std::vector<unsigned char> checkpoint_key;
    marshal_uint256_t(hash(), checkpoint_key);

    return ::deleteCheckpoint(storage, checkpoint_key);
}
