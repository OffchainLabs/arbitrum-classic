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

#include "avm/machine.hpp"

#include <sys/stat.h>
#include <fstream>
#include <iostream>
#include "avm/checkpointutils.hpp"
#include "avm/machinestatedata.hpp"
#include "avm/opcodes.hpp"
#include "bigint_utils.hpp"
#include "util.hpp"

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
    os << val.m;
    return os;
}

Machine::Machine(std::string filename) {
    std::ifstream myfile;

    struct stat filestatus;
    stat(filename.c_str(), &filestatus);

    char* buf = (char*)malloc(filestatus.st_size);

    myfile.open(filename, std::ios::in);

    if (myfile.is_open()) {
        myfile.read((char*)buf, filestatus.st_size);
        myfile.close();

        auto success = deserialize(buf);

        if (!success) {
            // log?
        }
    } else {
        // log??
    }
}

void Machine::sendOnchainMessage(const Message& msg) {
    m.sendOnchainMessage(msg);
}

void Machine::deliverOnchainMessages() {
    m.deliverOnchainMessages();
}

void Machine::sendOffchainMessages(const std::vector<Message>& messages) {
    m.sendOffchainMessages(messages);
}

CheckpointData Machine::getCheckPointData() {
    auto pc_value = CodePoint();
    pc_value.pc = m.pc;

    CheckpointData cp_data = {m.staticVal, m.registerVal,  m.stack,
                              m.auxstack,  m.pendingInbox, m.inbox,
                              m.pc,        m.errpc,        m.balance,
                              m.state,     m.blockReason};

    return cp_data;
}

Assertion Machine::run(uint64_t stepCount,
                       uint64_t timeBoundStart,
                       uint64_t timeBoundEnd) {
    m.context = AssertionContext{TimeBounds{{timeBoundStart, timeBoundEnd}}};
    m.blockReason = NotBlocked{};
    while (m.context.numSteps < stepCount) {
        runOne();
        if (!nonstd::get_if<NotBlocked>(&m.blockReason)) {
            break;
        }
    }
    return {m.context.numSteps, std::move(m.context.outMessage),
            std::move(m.context.logs)};
}

bool isErrorCodePoint(const CodePoint& cp) {
    return cp.nextHash == 0 && cp.op == Operation{static_cast<OpCode>(0)};
}

void Machine::runOne() {
    if (m.state == Status::Error) {
        m.blockReason = ErrorBlocked();
        return;
    }

    if (m.state == Status::Halted) {
        m.blockReason = HaltBlocked();
        return;
    }

    m.context.numSteps++;

    auto& instruction = m.code[m.pc];

    auto startStackSize = m.stack.stacksize();

    if (!isValidOpcode(instruction.op.opcode)) {
        m.state = Status::Error;
    } else {
        if (instruction.op.immediate) {
            auto imm = *instruction.op.immediate;
            m.stack.push(std::move(imm));
        }

        try {
            m.blockReason = m.runOp(instruction.op.opcode);
        } catch (const bad_pop_type& e) {
            m.state = Status::Error;
        } catch (const bad_tuple_index& e) {
            m.state = Status::Error;
        }
    }

    if (m.state != Status::Error) {
        return;
    }

    // Clear stack to base for instruction
    auto stackItems = InstructionStackPops.at(instruction.op.opcode).size();
    while (m.stack.stacksize() > 0 &&
           startStackSize - m.stack.stacksize() < stackItems) {
        m.stack.popClear();
    }

    if (!isErrorCodePoint(m.errpc)) {
        m.pc = m.errpc.pc;
        m.state = Status::Extensive;
    }

    return;
}
