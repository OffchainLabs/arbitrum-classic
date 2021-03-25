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

    machine_state.context = AssertionContext(std::move(config));

    auto inbox_opt = machine_state.output.fully_processed_inbox.inboxWithStaged(
        machine_state.staged_message);
    if (!inbox_opt.has_value()) {
        if (machine_state.context.inbox_messages.empty()) {
            reorg_check_data =
                ReorgState{machine_state.output.fully_processed_inbox,
                           machine_state.staged_message};
        } else {
            std::cerr << "Tried to run machine adding messages without "
                         "resolving staged message"
                      << std::endl;
            return false;
        }
    } else {
        for (const auto& message : machine_state.context.inbox_messages) {
            inbox_opt->addMessage(message);
        }
        if (machine_state.context.next_block_height) {
            reorg_check_data = {
                *inbox_opt,
                staged_variant{*machine_state.context.next_block_height}};
        } else {
            reorg_check_data = {*inbox_opt, staged_variant{std::monostate{}}};
        }
    }

    machine_status = MACHINE_RUNNING;

    machine_thread = std::make_unique<std::thread>(
        (std::reference_wrapper<MachineThread>(*this)));

    return true;
}

bool MachineThread::continueRunningMachine() {
    if (machine_status != MACHINE_NONE) {
        return false;
    }

    machine_state.context.resetForContinuedRun();

    machine_status = MACHINE_RUNNING;

    machine_thread = std::make_unique<std::thread>(
        (std::reference_wrapper<MachineThread>(*this)));

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

void MachineThread::operator()() {
    last_assertion = run();
    if (machine_status == MACHINE_RUNNING) {
        machine_status = MACHINE_SUCCESS;
    }
}
