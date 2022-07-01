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

#include "utils.hpp"
#include "libproofchecker.h"

#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <fstream>
#include <iostream>

value parseFuzzInputValue(const uint8_t*& buf,
                          const uint8_t* bufEnd,
                          const std::shared_ptr<Code>& code) {
    value output;
    std::vector<Value*> slots;
    slots.push_back(&output);
    while (!slots.empty()) {
        auto slot = slots.back();
        slots.pop_back();
        if (buf >= bufEnd) {
            continue;
        }
        uint8_t ty = *buf;
        buf++;
        if (ty == NUM) {
            if (buf + 32 >= bufEnd) {
                continue;
            }
            auto ptr = reinterpret_cast<const char*>(buf);
            *slot = deserializeUint256t(ptr);
            buf += 32;
        } else if (ty == CODEPT) {
            if (buf + 8 >= bufEnd) {
                continue;
            }
            auto ptr = reinterpret_cast<const char*>(buf);
            auto pc = deserialize_uint64_t(ptr);
            buf += 8;
            if (!code->containsSegment(0) ||
                code->loadCodeSegment(0).op_count <= pc) {
                continue;
            }
            CodePointRef ref{0, pc};
            *slot = CodePointStub{ref, code->loadCodePoint(ref)};
        } else if (ty >= TUPLE && ty <= TUPLE + 8) {
            auto size = ty - TUPLE;
            auto tup = Tuple::createSizedTuple(size);
            *slot = tup;
            for (size_t i = size; i > 0; i--) {
                slots.push_back(&tup.get_element_mutable_unsafe(i - 1));
            }
        } else if (ty == BUFFER) {
            if (buf + 8 >= bufEnd) {
                continue;
            }
            auto ptr = reinterpret_cast<const char*>(buf);
            auto size = deserialize_uint64_t(ptr);
            buf += 8;
            if (size > (bufEnd - buf)) {
                size = bufEnd - buf;
            }
            *slot = Buffer::fromData(std::vector(buf, buf + size));
            buf += size;
        }
    }
    return output;
}

Machine parseFuzzInput(const uint8_t* buf, size_t len) {
    auto code = std::make_shared<Code>();
    auto bufEnd = buf + len;
    value staticVal = parseFuzzInputValue(buf, bufEnd, code);
    auto stub = code->addSegment();
    while (buf + 1 < bufEnd) {
        auto opcode = static_cast<OpCode>(*buf++);
        if (*buf == 0xff) {
            buf++;
            Operation op = {opcode};
            stub = code->addOperation(stub.pc, op);
        } else {
            auto immediate = parseFuzzInputValue(buf, bufEnd, code);
            Operation op = {opcode, immediate};
            stub = code->addOperation(stub.pc, op);
        }
    }
    return Machine(code, staticVal);
}

bool opcodeAllowed(OpCode opcode) {
    switch (opcode) {
        case OpCode::INBOX:
        case OpCode::BREAKPOINT:
            return false;
        default:
            return true;
    }
}

ProofTester::ProofTester(bool debug) {}

void writeMachineState(const Machine& machine,
                       std::vector<unsigned char>& output_data) {
    marshal_uint256_t(machine.machine_state.output.arb_gas_used, output_data);
    marshal_uint256_t(machine.machine_state.getTotalMessagesRead(),
                      output_data);
    marshal_uint256_t(machine.machine_state.hash(), output_data);
    marshal_uint256_t(machine.machine_state.output.send_acc, output_data);
    marshal_uint256_t(machine.machine_state.output.log_acc, output_data);
}

void writeVarSizedBytes(const std::vector<unsigned char>& data,
                        std::vector<unsigned char>& output_data) {
    marshal_uint64_t(data.size(), output_data);
    output_data.insert(output_data.end(), data.begin(), data.end());
}

void writeProof(const OneStepProof& proof,
                std::vector<unsigned char>& output_data) {
    writeVarSizedBytes(proof.standard_proof, output_data);
    writeVarSizedBytes(proof.buffer_proof, output_data);
}

void ProofTester::testMachine(Machine machine) {
    Machine machine2 = machine;
    while (machine.currentStatus() == Status::Extensive &&
           machine.machine_state.output.arb_gas_used < FUZZ_MAX_GAS &&
           machine.machine_state.output.total_steps < FUZZ_MAX_STEPS &&
           opcodeAllowed(machine.machine_state.loadCurrentOperation().opcode)) {
        MachineExecutionConfig config;
        config.max_gas = machine.machine_state.output.arb_gas_used + 1;
        config.go_over_gas = true;
        machine.machine_state.context = AssertionContext(config);

        std::vector<unsigned char> proof_data;
        writeMachineState(machine, proof_data);
        auto proof = machine.marshalForProof();
        writeProof(proof, proof_data);
        auto assertion = machine.run();
        fuzz_require(assertion.stepCount == 1,
                     "Assertion produced wrong step count ",
                     assertion.stepCount);
        writeMachineState(machine, proof_data);
        int res = CheckProof(0, proof_data.data(), proof_data.size());
        if (res != 0) {
            throw std::runtime_error("Go proof tester errored");
        }
    }
    if (machine.machine_state.output.arb_gas_used != 0) {
        MachineExecutionConfig config;
        config.max_gas = machine.machine_state.output.arb_gas_used;
        config.go_over_gas = true;
        machine2.machine_state.context = AssertionContext(config);
        machine2.run();
    }
    fuzz_require(
        machine.hash() == machine2.hash(),
        "Machine execution differed when re-executed",
        "\nFirst execution gas:  ", machine.machine_state.output.arb_gas_used,
        "\nSecond execution gas: ", machine2.machine_state.output.arb_gas_used);
}
