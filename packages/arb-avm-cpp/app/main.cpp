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

#include <avm/machine.hpp>
#include <data_storage/arbstorage.hpp>
#include <data_storage/storageresult.hpp>
#include <data_storage/value/machine.hpp>

#include <boost/filesystem.hpp>

#include <iostream>
#include <string>

constexpr auto temp_db_path = "arb_runner_temp_db";

struct DBDeleter {
    DBDeleter() { boost::filesystem::remove_all(temp_db_path); }
    ~DBDeleter() { boost::filesystem::remove_all(temp_db_path); }
};

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

    DBDeleter deleter;
    ArbStorage storage{temp_db_path};

    if (mode == "--hexops") {
        std::ifstream file(filename, std::ios::binary);
        if (!file.is_open()) {
            throw std::runtime_error("Couldn't open file");
        }
        std::vector<unsigned char> raw_ops(
            (std::istreambuf_iterator<char>(file)),
            std::istreambuf_iterator<char>());

        std::vector<Operation> ops;
        ops.reserve(raw_ops.size());
        for (auto op : raw_ops) {
            ops.emplace_back(static_cast<OpCode>(op));
        }

        auto code = std::make_shared<UnsafeCodeSegment>(0, ops);
        auto status =
            storage.initialize(LoadedExecutable{std::move(code), Tuple()});
        if (!status.ok()) {
            throw std::runtime_error("Error initializing storage");
        }
    } else {
        auto status = storage.initialize(filename);
        if (!status.ok()) {
            throw std::runtime_error("Error initializing storage");
        }
    }

    ValueCache value_cache{1, 0};
    auto mach = storage.getInitialMachine();

    std::vector<MachineMessage> inbox_messages;
    if (argc == 5) {
        if (std::string(argv[3]) == "--inbox") {
            std::ifstream file(argv[4], std::ios::binary);
            if (!file.is_open()) {
                throw std::runtime_error("Couldn't open file");
            }
            std::vector<unsigned char> raw_inbox(
                (std::istreambuf_iterator<char>(file)),
                std::istreambuf_iterator<char>());
            auto data = reinterpret_cast<const char*>(raw_inbox.data());
            auto inbox_val = std::get<Tuple>(deserialize_value(data));
            while (inbox_val != Tuple{}) {
                inbox_messages.emplace_back(
                    InboxMessage::fromTuple(
                        std::get<Tuple>(inbox_val.get_element(1))),
                    0);
                inbox_val =
                    std::get<Tuple>(std::move(inbox_val.get_element(0)));
            }
            std::reverse(inbox_messages.begin(), inbox_messages.end());
        } else if (std::string(argv[3]) == "--json-inbox") {
            std::ifstream file(argv[4], std::ios::binary);
            nlohmann::json j;
            file >> j;

            for (auto& val : j["inbox"]) {
                inbox_messages.emplace_back(
                    InboxMessage::fromTuple(
                        std::get<Tuple>(simple_value_from_json(val))),
                    0);
            }
        }
    }

    MachineExecutionConfig execConfig;
    execConfig.inbox_messages = inbox_messages;
    mach->machine_state.context = AssertionContext{execConfig};
    auto assertion = mach->run();

    std::cout << "Produced " << assertion.logs.size() << " logs\n";

    std::cout << "Ran " << assertion.step_count << " steps in "
              << assertion.gas_count << " gas ending in state "
              << static_cast<int>(mach->currentStatus()) << "\n";

    auto tx = storage.makeReadWriteTransaction();
    saveTestMachine(*tx, *mach);
    tx->commit();

    auto mach2 = storage.getMachine(mach->hash(), value_cache);
    execConfig.inbox_messages = std::vector<MachineMessage>();
    mach2->machine_state.context = AssertionContext{execConfig};
    mach2->run();
    return 0;
}
