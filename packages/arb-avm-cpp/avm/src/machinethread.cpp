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

#include <iostream>

#include <avm/machinethread.hpp>

MachineThread::machine_status_enum MachineThread::status() {
    return machine_status;
}

void MachineThread::setStatus(machine_status_enum status) {
    machine_status = status;
}

void MachineThread::abortThread() {
    machine_abort = true;
}

bool MachineThread::setRunning() {
    if (machine_status == MACHINE_RUNNING) {
        return false;
    }

    machine_abort = false;
    machine_status = MACHINE_RUNNING;

    return true;
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
    uint64_t gas_limit,
    bool hard_gas_limit,
    const std::vector<std::vector<unsigned char>>& inbox_messages,
    const nonstd::optional<uint256_t>& final_block) {
    if (machine_status != MACHINE_NONE) {
        machine_error_string =
            "Unexpected machine_status when trying to run machine";
        machine_status = MACHINE_ERROR;
        return;
    }
    last_assertion =
        run(gas_limit, hard_gas_limit, inbox_messages, final_block);
    machine_status = MACHINE_FINISHED;
}
