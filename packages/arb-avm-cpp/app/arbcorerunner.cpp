/*
 * Copyright 2021, Offchain Labs, Inc.
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

#include <data_storage/arbstorage.hpp>

#include <iostream>
#include <sstream>
#include <string>

void sleep() {
    std::chrono::hours dura{20};
    //    std::chrono::seconds dura(2000);
    std::this_thread::sleep_for(dura);  // this makes this thread sleep for 2s
}

constexpr bool clearDb = false;

bool exclusively_code_points(Tuple tup) {
    for (size_t i = 0; i < tup.tuple_size(); i++) {
        auto elem = tup.get_element(i);
        if (std::holds_alternative<CodePointStub>(elem)) {
            continue;
        }

        if (auto inner_tup = std::get_if<Tuple>(&elem)) {
            if (exclusively_code_points(*inner_tup)) {
                continue;
            }
        }

        return false;
    }
    return true;
}

std::string print_path(const std::vector<size_t>& path) {
    std::stringstream output;
    output << "register";
    for (auto part : path) {
        output << "[";
        output << part;
        output << "]";
    }
    return output.str();
}

int main(int argc, char* argv[]) {
    using namespace std::chrono_literals;
    if (argc < 4) {
        std::cerr << "Usage: \narbcore_runner dbpath arbos1path arbos2path\n";
        return 1;
    }
    auto dbpath = std::string(argv[1]);
    auto arbospath = std::string(argv[2]);
    auto arbos2Path = std::string(argv[3]);

    if (clearDb) {
        {
            ArbStorage storage{dbpath};
            {
                auto tx = storage.makeReadWriteTransaction();
                saveNextSegmentID(*tx, 0);
                auto s = tx->commit();
                if (!s.ok()) {
                    std::cerr << "Error overwriting segment: " << s.ToString()
                              << std::endl;
                    return -1;
                }
            }
        }
        {
            DataStorage storage{dbpath};
            auto s = storage.clearDBExceptInbox();
            if (!s.ok()) {
                std::cerr << "Error deleting columns: " << s.ToString()
                          << std::endl;
                return -1;
            }
            storage.closeDb();
        }
        {
            ArbStorage storage{dbpath};
            auto s = storage.initialize(arbospath);
            if (!s.ok()) {
                std::cerr << "Failed to get initialize storage" << s.ToString()
                          << std::endl;
                return -1;
            }
        }
    }

    std::cerr << "Loading db\n";
    ArbStorage storage{dbpath};
    std::cerr << "Initializing arbstorage\n";
    auto core = storage.getArbCore();
    ValueCache cache{0, 2};

    auto res = core->getMachineForSideload(152, cache);
    if (!res.status.ok()) {
        std::cerr << "Failed to get machine for sideload: "
                  << res.status.ToString() << std::endl;
        return -1;
    }
    auto arbos1Machine = std::move(res.data);
    std::cout << "Got arbos1Machine with "
              << arbos1Machine->machine_state.output.arb_gas_used << " gas used"
              << std::endl;

    auto arbos2Machine = Machine::loadFromFile(arbos2Path);
    auto messages = core->getMessages(
        0, arbos1Machine->machine_state.output.fully_processed_inbox.count);
    if (!messages.status.ok()) {
        std::cerr << "Failed to get messages: " << messages.status.ToString()
                  << std::endl;
        return -1;
    }
    for (auto& message : messages.data) {
        message.insert(message.begin(), 32, 0);
    }
    MachineExecutionConfig config;
    config.setInboxMessagesFromBytes(messages.data);
    arbos2Machine.machine_state.context = AssertionContext(std::move(config));
    arbos2Machine.run();
    std::cout << "Got arbos2Machine with "
              << arbos2Machine.machine_state.output.arb_gas_used << " gas used"
              << std::endl;

    auto left_stack = std::vector(1, arbos1Machine->machine_state.registerVal);
    auto right_stack = std::vector(1, arbos2Machine.machine_state.registerVal);
    auto path_stack = std::vector(1, std::vector<size_t>());

    constexpr uint256_t min_balance = 1e9;
    constexpr uint256_t allowed_balance_diff = 1e3;

    while (!left_stack.empty() && !right_stack.empty()) {
        auto left = left_stack.back();
        left_stack.pop_back();
        auto right = right_stack.back();
        right_stack.pop_back();
        auto path = path_stack.back();
        path_stack.pop_back();

        if (std::holds_alternative<CodePointStub>(left) &&
            std::holds_alternative<CodePointStub>(right)) {
            continue;
        }

        if (hash_value(left) == hash_value(right)) {
            continue;
        }

        if (auto left_int = std::get_if<uint256_t>(&left)) {
            if (auto right_int = std::get_if<uint256_t>(&right)) {
                if (*left_int >= min_balance && *right_int >= min_balance &&
                    (*left_int - *right_int <= allowed_balance_diff ||
                     *right_int - *left_int <= allowed_balance_diff)) {
                    continue;
                } else {
                    std::cout << "Disallowed int difference" << std::endl
                              << "Path:  " << print_path(path) << std::endl
                              << "Left:  " << *left_int << std::endl
                              << "Right: " << *right_int << std::endl;
                    continue;
                }
            }
        }

        if (auto left_tup = std::get_if<Tuple>(&left)) {
            if (auto right_tup = std::get_if<Tuple>(&right)) {
                auto left_size = left_tup->tuple_size();
                auto right_size = right_tup->tuple_size();
                if (left_size == right_size) {
                    for (size_t i = 0; i < left_size; i++) {
                        left_stack.push_back(left_tup->get_element(i));
                        right_stack.push_back(right_tup->get_element(i));
                        std::vector<size_t> new_path = path;
                        new_path.push_back(i);
                        path_stack.push_back(new_path);
                    }
                    continue;
                } else if (exclusively_code_points(*left_tup) &&
                           exclusively_code_points(*right_tup)) {
                    continue;
                } else {
                    if (left_tup->getSize() > 100 ||
                        right_tup->getSize() > 100) {
                        std::cout << "Differing tuple sizes" << std::endl
                                  << "Path:  " << print_path(path) << std::endl
                                  << "Left:  " << left_size << " "
                                  << left_tup->getSize() << std::endl
                                  << "Right: " << right_size << " "
                                  << right_tup->getSize() << std::endl;
                    } else {
                        std::cout << "Differing tuples" << std::endl
                                  << "Path:  " << print_path(path) << std::endl
                                  << "Left:  " << left << std::endl
                                  << "Right: " << right << std::endl;
                    }
                    continue;
                }
            }
        }

        std::cout << "Differing values" << std::endl
                  << "Path:  " << print_path(path) << std::endl
                  << "Left:  " << intx::to_string(hash_value(left), 16)
                  << std::endl
                  << "Right: " << intx::to_string(hash_value(right), 16)
                  << std::endl;
    }

    if (!left_stack.empty() || !right_stack.empty()) {
        std::cerr << "Internal error: unbalanced traversal" << std::endl;
        return -1;
    }

    return 0;
}
