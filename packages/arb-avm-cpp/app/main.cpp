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

#include "bigint_utils.hpp"

#include <avm/machine.hpp>

#include <boost/algorithm/hex.hpp>

#include <sys/stat.h>
#include <fstream>
#include <iostream>
#include <string>
#include <thread>

int main(int argc, char* argv[]) {
    using namespace std::chrono_literals;
    if (argc < 3 || (std::string(argv[1]) != "--hexops" &&
                     std::string(argv[1]) != "--mexe")) {
        std::cout << "Usage: \n"
                     "avm_runner --hexops filename [--inbox filename]\n"
                     "avm_runner --mexe filename [--inbox filename]\n";
        return 1;
    }
    auto mode = std::string(argv[1]);
    std::string filename = argv[2];

    auto mach = [&]() {
        if (mode == "--hexops") {
            std::ifstream file(filename, std::ios::binary);
            if (!file.is_open()) {
                throw std::runtime_error("Couldn't open file");
            }
            std::vector<unsigned char> raw_ops(
                (std::istreambuf_iterator<char>(file)),
                std::istreambuf_iterator<char>());

            auto code = std::make_shared<Code>();
            auto stub = code->addSegment();
            // Code segments are built back to front so add operations in
            // reverse
            for (auto it = raw_ops.rbegin(); it != raw_ops.rend(); ++it) {
                stub = code->addOperation(stub.pc,
                                          Operation(static_cast<OpCode>(*it)));
            }
            return Machine(MachineState(std::move(code), Tuple(),
                                        std::make_shared<TuplePool>()));
        } else {
            return Machine::loadFromFile(filename);
        }
    }();

    Tuple inbox;
    if (argc == 5 && std::string(argv[3]) == "--inbox") {
        std::ifstream file(argv[4], std::ios::binary);
        if (!file.is_open()) {
            throw std::runtime_error("Couldn't open file");
        }
        std::vector<unsigned char> raw_inbox(
            (std::istreambuf_iterator<char>(file)),
            std::istreambuf_iterator<char>());
        auto data = reinterpret_cast<const char*>(raw_inbox.data());
        auto inboxVal = deserialize_value(data, mach.getPool());
        inbox = nonstd::get<Tuple>(inboxVal);
    }

    auto assertionBase = mach.run(100000000, Tuple(), std::chrono::seconds(0));

    auto assertion = mach.run(100000000, inbox, std::chrono::seconds(0));

    std::cout << "Ran " << assertion.stepCount << " ending in state "
              << static_cast<int>(mach.currentStatus()) << "\n";
    return 0;
}
