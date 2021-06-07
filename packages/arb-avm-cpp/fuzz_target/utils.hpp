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

#pragma once

#include <avm/machine.hpp>

#include <fstream>
#include <iostream>
#include <sstream>

template <class... Args>
void fuzz_require(bool check, Args... message) {
    if (!check) {
        std::stringstream err_stream;
        err_stream << "fuzz_require failed: ";
        ((err_stream << std::forward<Args>(message)), ...);
        err_stream << std::endl;
        throw std::runtime_error(err_stream.str());
    }
}

Machine parseFuzzInput(const uint8_t* buf, size_t len);

bool opcodeAllowed(OpCode opcode);

constexpr uint256_t FUZZ_MAX_GAS = 10'000;
constexpr uint64_t FUZZ_MAX_STEPS = 100;

class ProofTester {
   protected:
    std::string queryPipePath;
    std::string resultPipePath;
    std::fstream queryPipe;
    std::fstream resultPipe;

    void writeMachineState(const Machine& machine);
    void writeVarSizedBytes(const std::vector<unsigned char>& data);
    void writeProof(const OneStepProof& proof);
    uint8_t readResult();

   public:
    ProofTester(bool debug);
    ~ProofTester();

    void testMachine(Machine machine);
};
