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
        // Disable these opcodes for now
        if (opcode == OpCode::INBOX || opcode == OpCode::BREAKPOINT) {
            opcode = static_cast<OpCode>(0);
        }
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

constexpr uint256_t MAX_GAS = 10'000;
constexpr uint64_t MAX_STEPS = 100;

void testMachine(Machine machine) {
    while (machine.currentStatus() == Status::Extensive &&
           machine.machine_state.output.arb_gas_used < MAX_GAS &&
           machine.machine_state.output.total_steps < MAX_STEPS) {
        MachineExecutionConfig config;
        config.max_gas = machine.machine_state.output.arb_gas_used + 1;
        config.go_over_gas = true;
        machine.machine_state.context = AssertionContext(config);
        auto proof = machine.marshalForProof();
        auto assertion = machine.run();
        fuzz_require(assertion.stepCount == 1,
                     "Assertion produced wrong step count ",
                     assertion.stepCount);
        // TODO test assertion proof against Go checker
    }
}
