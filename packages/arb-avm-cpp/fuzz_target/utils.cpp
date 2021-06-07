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

#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <fstream>
#include <iostream>

value parseFuzzInputValue(const uint8_t*& buf,
                          const uint8_t* bufEnd,
                          const std::shared_ptr<Code>& code) {
    value output;
    std::vector<value*> slots;
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

ProofTester::ProofTester(bool debug) {
    queryPipePath = std::tmpnam(nullptr);
    if (mkfifo(queryPipePath.data(), 0644)) {
        throw std::runtime_error(std::string("Failed to mkfifo: ") +
                                 std::strerror(errno));
    }
    resultPipePath = std::tmpnam(nullptr);
    if (mkfifo(resultPipePath.data(), 0644)) {
        throw std::runtime_error(std::string("Failed to mkfifo result pipe: ") +
                                 std::strerror(errno));
    }
    if (fork() == 0) {
        execl("/bin/sh", "/bin/sh", "-c",
              "../../arb-node-core/proof_test_server $1 $2", "-s",
              queryPipePath.data(), resultPipePath.data(), nullptr);

        throw std::runtime_error(
            std::string("Error starting Go proof test server: ") +
            std::strerror(errno));
    }
    queryPipe.open(queryPipePath, std::ios::out | std::ios::binary);
    resultPipe.open(resultPipePath, std::ios::in | std::ios::binary);
    uint8_t initMessage[1];
    initMessage[0] = debug;
    queryPipe.write(reinterpret_cast<const char*>(initMessage),
                    sizeof(initMessage));
    queryPipe.flush();
    resultPipe.read(reinterpret_cast<char*>(initMessage), sizeof(initMessage));
    if (debug) {
        std::cerr << "Established contact with Go proof test server"
                  << std::endl;
    }
}

ProofTester::~ProofTester() {
    if (!resultPipePath.empty()) {
        remove(resultPipePath.data());
    }
    if (!queryPipePath.empty()) {
        remove(queryPipePath.data());
    }
}

void ProofTester::writeMachineState(const Machine& machine) {
    std::vector<unsigned char> data;
    marshal_uint256_t(machine.machine_state.output.arb_gas_used, data);
    marshal_uint256_t(machine.machine_state.getTotalMessagesRead(), data);
    marshal_uint256_t(machine.machine_state.hash(), data);
    marshal_uint256_t(machine.machine_state.output.send_acc, data);
    marshal_uint256_t(machine.machine_state.output.log_acc, data);
    queryPipe.write(reinterpret_cast<const char*>(data.data()), data.size());
}

void ProofTester::writeVarSizedBytes(const std::vector<unsigned char>& data) {
    std::vector<unsigned char> tmp;
    marshal_uint64_t(data.size(), tmp);
    queryPipe.write(reinterpret_cast<const char*>(tmp.data()), tmp.size());
    queryPipe.write(reinterpret_cast<const char*>(data.data()), data.size());
}

void ProofTester::writeProof(const OneStepProof& proof) {
    writeVarSizedBytes(proof.standard_proof);
    writeVarSizedBytes(proof.buffer_proof);
}

uint8_t ProofTester::readResult() {
    uint8_t buf[1];
    resultPipe.read(reinterpret_cast<char*>(buf), 1);
    return buf[0];
}

void ProofTester::testMachine(Machine machine) {
    Machine machine2 = machine;
    size_t buffered_results = 0;
    while (machine.currentStatus() == Status::Extensive &&
           machine.machine_state.output.arb_gas_used < FUZZ_MAX_GAS &&
           machine.machine_state.output.total_steps < FUZZ_MAX_STEPS &&
           opcodeAllowed(machine.machine_state.loadCurrentOperation().opcode)) {
        MachineExecutionConfig config;
        config.max_gas = machine.machine_state.output.arb_gas_used + 1;
        config.go_over_gas = true;
        machine.machine_state.context = AssertionContext(config);
        writeMachineState(machine);
        auto proof = machine.marshalForProof();
        writeProof(proof);
        auto assertion = machine.run();
        fuzz_require(assertion.stepCount == 1,
                     "Assertion produced wrong step count ",
                     assertion.stepCount);
        writeMachineState(machine);
        queryPipe.flush();
        if (buffered_results < 4) {
            buffered_results++;
        } else if (readResult() != 0) {
            throw std::runtime_error("Go proof tester errored");
        }
    }
    for (; buffered_results > 0; buffered_results--) {
        if (readResult() != 0) {
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
