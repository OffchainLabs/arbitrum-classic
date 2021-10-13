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

bool MachineThread::runMachine(MachineExecutionConfig config,
                               bool asynchronous) {
    if (machine_status != MACHINE_NONE) {
        return false;
    }

    machine_state.context = AssertionContext(std::move(config));

    reorg_check_data = machine_state.output.fully_processed_inbox;
    for (const auto& message : machine_state.context.inbox_messages) {
        reorg_check_data.addMessage(message);
    }

    machine_status = MACHINE_RUNNING;

    if (asynchronous) {
        machine_thread = std::make_unique<std::thread>(
            (std::reference_wrapper<MachineThread>(*this)));
    } else {
        this->operator()();
    }

    return true;
}

bool MachineThread::continueRunningMachine(bool asynchronous) {
    if (machine_status != MACHINE_NONE) {
        return false;
    }

    machine_state.context.resetForContinuedRun();

    machine_status = MACHINE_RUNNING;

    if (asynchronous) {
        machine_thread = std::make_unique<std::thread>(
            (std::reference_wrapper<MachineThread>(*this)));
    } else {
        this->operator()();
    }

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
    if (machine_thread != nullptr) {
        machine_thread->join();
        machine_thread = nullptr;
    }
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

void MachineThread::operator()() {
    last_assertion = run();
    if (machine_status == MACHINE_RUNNING) {
        machine_status = MACHINE_SUCCESS;
    }
}
