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
#include <string>

void sleep() {
    std::chrono::seconds dura(2000);
    std::this_thread::sleep_for(dura);  // this makes this thread sleep for 2s
}

int main(int argc, char* argv[]) {
    using namespace std::chrono_literals;
    if (argc < 3) {
        std::cout << "Usage: \narbcore_runner dbpath arbospath\n";
        return 1;
    }
    auto dbpath = std::string(argv[1]);
    auto arbospath = std::string(argv[2]);
    std::cout << "Loading db\n";
    ArbStorage storage{dbpath};
    std::cout << "Initializing arbstorage\n";
    auto status = storage.initialize(arbospath);
    if (!status.ok()) {
        std::cerr << "Failed to get initialize storage" << status.ToString()
                  << std::endl;
        return -1;
    }
    auto core = storage.getArbCore();
    auto log_count = core->logInsertedCount();
    if (!log_count.status.ok()) {
        std::cerr << "Failed to get log count " << log_count.status.ToString()
                  << std::endl;
        return -1;
    }
    std::cout << "Starting machine thread\n";
    core->startThread();
    std::cout << "Log count: " << log_count.data << "\n";
    sleep();
    return 0;
}
