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

#include <avm/checkpoint/checkpointdeleter.hpp>
#include <avm/machine.hpp>
#include <avm/opcodes.hpp>
#include <bigint_utils.hpp>
#include <util.hpp>

std::ostream& operator<<(std::ostream& os, const MachineState& val) {
    os << "codePointHash " << to_hex_str(hash(val.code[val.pc])) << "\n";
    os << "stackHash " << to_hex_str(val.stack.hash()) << "\n";
    os << "auxStackHash " << to_hex_str(val.auxstack.hash()) << "\n";
    os << "registerHash " << to_hex_str(hash(val.registerVal)) << "\n";
    os << "staticHash " << to_hex_str(hash(val.staticVal)) << "\n";
    os << "errHandlerHash " << to_hex_str(hash(val.errpc)) << "\n";
    return os;
}

std::ostream& operator<<(std::ostream& os, const Machine& val) {
    os << val.machine_state;
    return os;
}

bool Machine::initializeMachine(const std::string& filename) {
    std::ifstream myfile;

    struct stat filestatus;
    stat(filename.c_str(), &filestatus);

    char* buf = (char*)malloc(filestatus.st_size);

    myfile.open(filename, std::ios::in);

    if (myfile.is_open()) {
        myfile.read((char*)buf, filestatus.st_size);
        myfile.close();

        return deserialize(buf);
    } else {
        return false;
    }
}

void Machine::sendOnchainMessage(const Message& msg) {
    machine_state.sendOnchainMessage(msg);
}

void Machine::deliverOnchainMessages() {
    machine_state.deliverOnchainMessages();
}

void Machine::sendOffchainMessages(const std::vector<Message>& messages) {
    machine_state.sendOffchainMessages(messages);
}

Assertion Machine::run(uint64_t stepCount,
                       uint64_t timeBoundStart,
                       uint64_t timeBoundEnd) {
    machine_state.context =
        AssertionContext{TimeBounds{{timeBoundStart, timeBoundEnd}}};
    machine_state.blockReason = NotBlocked{};
    while (machine_state.context.numSteps < stepCount) {
        runOne();
        if (!nonstd::get_if<NotBlocked>(&machine_state.blockReason)) {
            break;
        }
    }
    return {machine_state.context.numSteps,
            std::move(machine_state.context.outMessage),
            std::move(machine_state.context.logs)};
}

bool isErrorCodePoint(const CodePoint& cp) {
    return cp.nextHash == 0 && cp.op == Operation{static_cast<OpCode>(0)};
}

void Machine::runOne() {
    if (machine_state.state == Status::Error) {
        machine_state.blockReason = ErrorBlocked();
        return;
    }

    if (machine_state.state == Status::Halted) {
        machine_state.blockReason = HaltBlocked();
        return;
    }

    machine_state.context.numSteps++;

    auto& instruction = machine_state.code[machine_state.pc];

    auto startStackSize = machine_state.stack.stacksize();

    if (!isValidOpcode(instruction.op.opcode)) {
        machine_state.state = Status::Error;
    } else {
        if (instruction.op.immediate) {
            auto imm = *instruction.op.immediate;
            machine_state.stack.push(std::move(imm));
        }
        try {
            machine_state.blockReason =
                machine_state.runOp(instruction.op.opcode);
        } catch (const bad_pop_type& e) {
            machine_state.state = Status::Error;
        } catch (const bad_tuple_index& e) {
            machine_state.state = Status::Error;
        }
    }

    if (machine_state.state != Status::Error) {
        return;
    }

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

    return;
}

SaveResults Machine::checkpoint(CheckpointStorage& storage) {
    return machine_state.checkpointState(storage);
}

bool Machine::restoreCheckpoint(
    CheckpointStorage& storage,
    const std::vector<unsigned char>& checkpoint_key) {
    return machine_state.restoreCheckpoint(storage, checkpoint_key).status.ok();
}

DeleteResults Machine::deleteCheckpoint(CheckpointStorage& storage) {
    auto checkpoint_key = GetHashKey(hash());
    auto results = ::deleteCheckpoint(storage, checkpoint_key);
    return results;
}
