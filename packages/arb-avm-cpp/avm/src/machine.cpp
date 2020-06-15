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

std::ostream& operator<<(std::ostream& os, const Machine& val) {
    os << val.machine_state;
    return os;
}

Assertion Machine::run(uint64_t stepCount,
                       const TimeBounds& timeBounds,
                       Tuple messages,
                       std::chrono::seconds wallLimit) {
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

BlockReason Machine::runOne() {
    if (machine_state.state == Status::Error) {
        return ErrorBlocked();
    }

    if (machine_state.state == Status::Halted) {
        return HaltBlocked();
    }

    auto& instruction = machine_state.static_values->code[machine_state.pc];

    // if opcode is invalid, increment step count and return error or
    // errorCodePoint
    if (!isValidOpcode(instruction.op.opcode)) {
        machine_state.state = Status::Error;
        machine_state.context.numSteps++;
        if (!machine_state.errpc.is_err) {
            machine_state.pc = machine_state.errpc;
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

        if (!machine_state.errpc.is_err) {
            machine_state.pc = machine_state.errpc;
            machine_state.state = Status::Extensive;
        }
        return blockReason;
    }
}
