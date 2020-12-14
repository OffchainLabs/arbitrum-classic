/*
 * Copyright 2020, Offchain Labs, Inc.
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

#include "ccheckpointedmachine.h"

#include "data_storage/checkpointedmachine.hpp"
#include "data_storage/value/machine.hpp"
#include "utils.hpp"

#include <iostream>

CheckpointedMachine* read_files(const std::string& executable_filename,
                                const std::string& db_path) {
    try {
        return new CheckpointedMachine(
            CheckpointedMachine::loadFromFile(executable_filename, db_path));
    } catch (const std::exception& e) {
        std::cerr << "Error loading machine " << executable_filename << ": "
                  << e.what() << "\n";
        return nullptr;
    }
}

CCheckpointedMachine* checkpointedMachineCreate(const char* executable_filename,
                                                const char* db_path) {
    CheckpointedMachine* mach = read_files(executable_filename, db_path);
    return static_cast<void*>(mach);
}

int initializeCheckpointedMachine(CCheckpointedMachine* cm,
                                  const char* executable_path) {
    auto cmach = static_cast<CheckpointedMachine*>(cm);
    try {
        cmach->initialize(executable_path);
        return true;
    } catch (const std::exception&) {
        return false;
    }
}

int checkpointedMachineInitialized(CCheckpointedMachine* cm) {
    auto cmach = static_cast<CheckpointedMachine*>(cm);
    try {
        cmach->initialized();
        return true;
    } catch (const std::exception&) {
        return false;
    }
}

int closeCheckpointedMachine(CCheckpointedMachine* cm) {
    auto cmach = static_cast<CheckpointedMachine*>(cm);
    try {
        return cmach->closeCheckpointedMachine();
    } catch (const std::exception&) {
        return false;
    }
}

void checkpointedMachineDestroy(CCheckpointedMachine* cm) {
    if (cm == nullptr) {
        return;
    }
    delete static_cast<CheckpointedMachine*>(cm);
}

RawAssertion checkpointedExecuteAssertion(CCheckpointedMachine* cm,
                                          uint64_t maxSteps,
                                          void* inbox_messages,
                                          uint64_t message_count,
                                          uint64_t wallLimit) {
    assert(cm);
    auto cmach = static_cast<CheckpointedMachine*>(cm);
    auto messages = getInboxMessages(inbox_messages, message_count);

    try {
        Assertion assertion = cmach->run(maxSteps, std::move(messages),
                                         std::chrono::seconds{wallLimit});
        return makeRawAssertion(assertion);
    } catch (const std::exception& e) {
        std::cerr << "Failed to make assertion " << e.what() << "\n";
        return makeEmptyAssertion();
    }
}

RawAssertion checkpointedExecuteCallServerAssertion(CCheckpointedMachine* cm,
                                                    uint64_t maxSteps,
                                                    void* inbox_messages,
                                                    uint64_t message_count,
                                                    void* fake_inbox_peek_value,
                                                    uint64_t wallLimit) {
    assert(cm);
    auto cmach = static_cast<CheckpointedMachine*>(cm);

    auto messages = getInboxMessages(inbox_messages, message_count);
    auto fake_inbox_peek_value_data =
        reinterpret_cast<const char*>(fake_inbox_peek_value);
    auto fake_inbox_peek = deserialize_value(fake_inbox_peek_value_data);

    try {
        Assertion assertion = cmach->runCallServer(
            maxSteps, std::move(messages), std::chrono::seconds{wallLimit},
            std::move(fake_inbox_peek));
        return makeRawAssertion(assertion);
    } catch (const std::exception& e) {
        std::cerr << "Failed to make assertion " << e.what() << "\n";
        return makeEmptyAssertion();
    }
}

RawAssertion checkpointedExecuteSideloadedAssertion(CCheckpointedMachine* cm,
                                                    uint64_t maxSteps,
                                                    void* inbox_messages,
                                                    uint64_t message_count,
                                                    void* sideload,
                                                    uint64_t wallLimit) {
    assert(cm);
    auto cmach = static_cast<CheckpointedMachine*>(cm);

    auto messages = getInboxMessages(inbox_messages, message_count);
    auto sideload_value = getTuple(sideload);

    try {
        Assertion assertion = cmach->runSideloaded(
            maxSteps, std::move(messages), std::chrono::seconds{wallLimit},
            std::move(sideload_value));
        return makeRawAssertion(assertion);
    } catch (const std::exception& e) {
        std::cerr << "Failed to make assertion " << e.what() << "\n";
        return makeEmptyAssertion();
    }
}
