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

#include <thread>

MachineThread::machine_status_enum MachineThread::status() {
    return machine_status;
}

bool MachineThread::runMachine(MachineExecutionConfig config) {
    if (machine_status != MACHINE_NONE) {
        return false;
    }

    machine_status = MACHINE_RUNNING;

    machine_thread = std::make_unique<std::thread>(
        (std::reference_wrapper<MachineThread>(*this)), config);

    return true;
}

void MachineThread::abortMachine() {
    if (machine_thread) {
        machine_abort = true;
        machine_thread->join();
        machine_thread = nullptr;
        machine_status = MACHINE_ABORTED;
    }
    machine_abort = false;
}

Assertion MachineThread::nextAssertion() {
    if (machine_status != MACHINE_SUCCESS) {
        return {};
    }
    machine_thread->join();
    machine_thread = nullptr;
    machine_status = MACHINE_NONE;
    return std::move(last_assertion);
}

std::string MachineThread::getErrorString() {
    return machine_error_string;
}

void MachineThread::clearError() {
    abortMachine();
    machine_status = MACHINE_NONE;
    machine_error_string.clear();
}

void MachineThread::operator()(MachineExecutionConfig config) {
    last_assertion = run(config);
    machine_status = MACHINE_SUCCESS;
}
