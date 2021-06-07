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

#include <avm/machine.hpp>

static std::unique_ptr<ProofTester> tester;

extern "C" int LLVMFuzzerInitialize(int* argc, char*** argv) {
    bool testing = true;
    bool failed_arg_parsing = false;
    if (*argc == 2) {
        std::string option = (*argv)[1];
        if (option == "--exercise") {
            testing = false;
        } else {
            failed_arg_parsing = true;
        }
    } else if (*argc > 2) {
        failed_arg_parsing = true;
    }
    if (failed_arg_parsing) {
        throw std::runtime_error("Usage: fuzz_target [--exercise]");
    }
    if (testing) {
        tester = std::make_unique<ProofTester>(true);
    }
    return 0;
}

extern "C" int LLVMFuzzerTestOneInput(uint8_t* buf, size_t len) {
    auto machine = parseFuzzInput(buf, len);
    if (tester) {
        tester->testMachine(machine);
    } else {
        MachineExecutionConfig config;
        while (machine.currentStatus() == Status::Extensive &&
               machine.machine_state.output.arb_gas_used < FUZZ_MAX_GAS &&
               machine.machine_state.output.total_steps < FUZZ_MAX_STEPS &&
               opcodeAllowed(
                   machine.machine_state.loadCurrentOperation().opcode)) {
            MachineExecutionConfig config;
            config.max_gas = machine.machine_state.output.arb_gas_used + 1;
            config.go_over_gas = true;
            machine.machine_state.context = AssertionContext(config);
            machine.run();
        }
    }
    return 0;
}
