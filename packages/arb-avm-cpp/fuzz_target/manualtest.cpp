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

#include <fstream>
#include <iomanip>
#include <iostream>

int main(int argc, char** argv) {
    if (argc < 2) {
        std::cerr << "Usage: ./manual_test [inputs...]" << std::endl;
        return -1;
    }

    ProofTester tester(true);

    for (size_t i = 1; i < argc; i++) {
        std::cerr << "Testing " << argv[i] << std::endl;
        std::ifstream file(argv[i], std::ios::binary | std::ios::ate);
        std::streamsize size = file.tellg();
        file.seekg(0, std::ios::beg);

        std::vector<unsigned char> buffer(size);
        if (!file.read(reinterpret_cast<char*>(buffer.data()), size)) {
            std::cerr << "Failed to read input file" << std::endl;
            return -1;
        }
        auto machine = parseFuzzInput(buffer.data(), buffer.size());
        {
            std::cerr << "Static value: " << machine.machine_state.static_val
                      << std::endl;
            auto segment = machine.machine_state.code->loadCodeSegment(0);
            for (size_t i = 0; i < segment.op_count; i++) {
                auto code_point =
                    segment.segment->loadCodePoint(segment.op_count - i - 1);
                std::cerr << std::setw(3) << i << std::setw(0) << " "
                          << code_point.op << std::endl;
            }
        }
        tester.testMachine(machine);
        std::cerr << "Success" << std::endl;
    }

    return 0;
}
