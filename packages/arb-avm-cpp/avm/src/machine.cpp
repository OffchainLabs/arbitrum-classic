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

auto operator<<(std::ostream& os, const MachineState& val) -> std::ostream& {
    os << "status " << static_cast<int>(val.state) << "\n";
    os << "codePointHash " << to_hex_str(hash(val.code[val.pc])) << "\n";
    os << "stackHash " << to_hex_str(val.stack.hash()) << "\n";
    os << "auxStackHash " << to_hex_str(val.auxstack.hash()) << "\n";
    os << "registerHash " << to_hex_str(hash(val.registerVal)) << "\n";
    os << "staticHash " << to_hex_str(hash(val.staticVal)) << "\n";
    os << "errHandlerHash " << to_hex_str(hash(val.errpc)) << "\n";
    return os;
}

auto operator<<(std::ostream& os, const Machine& val) -> std::ostream& {
    os << val.machine_state;
    return os;
}

auto Machine::initializeMachine(const std::string& filename) -> bool {
    return machine_state.initialize_machinestate(filename);
}

void Machine::initializeMachine(const MachineState& initial_state) {
    machine_state = initial_state;
}

auto Machine::run(uint64_t stepCount,
                  const TimeBounds& timeBounds,
                  Tuple messages,
                  std::chrono::seconds wallLimit) -> Assertion {
    bool has_time_limit = wallLimit.count() != 0;
    auto start_time = std::chrono::system_clock::now();
    machine_state.context = AssertionContext{timeBounds, std::move(messages)};
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

auto isErrorCodePoint(const CodePoint& cp) -> bool {
    return cp.nextHash == 0 && cp.op == Operation{static_cast<OpCode>(0)};
}

auto Machine::runOne() -> BlockReason {
    if (machine_state.state == Status::Error) {
        return ErrorBlocked();
    }

    if (machine_state.state == Status::Halted) {
        return HaltBlocked();
    }

    auto& instruction = machine_state.code[machine_state.pc];

    // if opcode is invalid, increment step count and return error or
    // errorCodePoint
    if (!isValidOpcode(instruction.op.opcode)) {
        machine_state.state = Status::Error;
        machine_state.context.numSteps++;
        if (!isErrorCodePoint(machine_state.errpc)) {
            machine_state.pc = machine_state.errpc.pc;
            machine_state.state = Status::Extensive;
        }
        return NotBlocked();
    } else {
        if (instruction.op.immediate) {
            auto imm = *instruction.op.immediate;
            machine_state.stack.push(std::move(imm));
        }
        // save stack size for stack cleanup in case of error
        uint64_t startStackSize = machine_state.stack.stacksize();
        BlockReason blockReason = NotBlocked();
        try {
            blockReason = machine_state.runOp(instruction.op.opcode);
        } catch (const bad_pop_type& e) {
            machine_state.state = Status::Error;
        } catch (const bad_tuple_index& e) {
            machine_state.state = Status::Error;
        }
        // if not blocked, increment step count and gas count
        if (nonstd::get_if<NotBlocked>(&blockReason)) {
            machine_state.context.numSteps++;
            machine_state.context.numGas +=
                InstructionArbGasCost.at(instruction.op.opcode);
        } else {
            if (instruction.op.immediate) {
                machine_state.stack.popClear();
            }
        }

        if (machine_state.state != Status::Error) {
            return blockReason;
        }
        // if state is Error, clean up stack
        // Clear stack to base for instruction
        auto stackItems = InstructionStackPops.at(instruction.op.opcode).size();
        while (machine_state.stack.stacksize() > 0 &&
               startStackSize - machine_state.stack.stacksize() < stackItems) {
            machine_state.stack.popClear();
        }

        if (!isErrorCodePoint(machine_state.errpc)) {
            machine_state.pc = machine_state.errpc.pc;
            machine_state.state = Status::Extensive;
        }
        return blockReason;
    }
}

auto Machine::checkpoint(CheckpointStorage& storage) -> SaveResults {
    return machine_state.checkpointState(storage);
}

auto Machine::restoreCheckpoint(
    const CheckpointStorage& storage,
    const std::vector<unsigned char>& checkpoint_key) -> bool {
    return machine_state.restoreCheckpoint(storage, checkpoint_key);
}

auto Machine::deleteCheckpoint(CheckpointStorage& storage) -> DeleteResults {
    std::vector<unsigned char> checkpoint_key;
    marshal_uint256_t(hash(), checkpoint_key);

    return ::deleteCheckpoint(storage, checkpoint_key);
}
