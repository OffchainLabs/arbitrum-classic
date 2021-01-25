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

#include <avm/machinethread.hpp>

#include <iostream>
#include <thread>

MachineThread::machine_status_enum MachineThread::status() {
    return machine_status;
}

void MachineThread::clearStatus() {
    abortThread();
    machine_status = MACHINE_NONE;
    machine_error_string.clear();
}

bool MachineThread::startThread(
    uint256_t max_gas,
    bool go_over_gas,
    const std::vector<std::vector<unsigned char>>& inbox_messages,
    uint256_t messages_to_skip,
    const nonstd::optional<uint256_t>& min_next_block_height) {
    abortThread();
    machine_status = MACHINE_RUNNING;

    machine_thread = std::make_unique<std::thread>(
        (std::reference_wrapper<MachineThread>(*this)), max_gas, go_over_gas,
        std::move(inbox_messages), messages_to_skip,
        std::move(min_next_block_height));

    return true;
}

void MachineThread::abortThread() {
    if (machine_thread) {
        machine_abort = true;
        machine_thread->join();
        machine_thread = nullptr;
    }
    machine_abort = false;
}

Assertion MachineThread::getAssertion() {
    return last_assertion;
}

std::string MachineThread::get_error_string() {
    return machine_error_string;
}

void MachineThread::clear_error_string() {
    machine_error_string.clear();
}

void MachineThread::operator()(
    const uint256_t max_gas,
    const bool go_over_gas,
    const std::vector<std::vector<unsigned char>>& inbox_messages,
    const uint256_t messages_to_skip,
    const nonstd::optional<uint256_t>& min_next_block_height) {
    if (machine_status != MACHINE_NONE) {
        machine_error_string =
            "Unexpected machine_status when trying to run machine";
        machine_status = MACHINE_ERROR;
        return;
    }
    last_assertion = run(max_gas, go_over_gas, inbox_messages, messages_to_skip,
                         min_next_block_height);
    machine_status = MACHINE_FINISHED;
}
