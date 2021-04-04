/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

#include "cmachine.h"
#include "utils.hpp"

#include <avm/machine.hpp>
#include <data_storage/arbstorage.hpp>
#include <data_storage/value/machine.hpp>

#include <iostream>
#include <sstream>

typedef struct {
    uint64_t stepCount;
} cassertion;

Machine* read_files(const std::string& filename) {
    try {
        return new Machine(Machine::loadFromFile(filename));
    } catch (const std::exception& e) {
        std::cerr << "Error loading machine " << filename << ": " << e.what()
                  << "\n";
        return nullptr;
    }
}

// cmachine_t *machine_create(char *data)
CMachine* machineCreate(const char* filename) {
    Machine* mach = read_files(filename);
    return static_cast<void*>(mach);
}

void machineDestroy(CMachine* m) {
    if (m == nullptr) {
        return;
    }
    delete static_cast<Machine*>(m);
}

int machineHash(CMachine* m, void* ret) {
    assert(m);
    auto optionalHash = static_cast<Machine*>(m)->hash();
    if (!optionalHash) {
        return 0;
    }
    std::array<unsigned char, 32> val{};
    to_big_endian(*optionalHash, val.begin());
    std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));

    return 1;
}

void* machineClone(CMachine* m) {
    assert(m);
    auto mach = static_cast<Machine*>(m);
    auto cloneMach = new Machine(*mach);
    return static_cast<void*>(cloneMach);
}

char* machineInfo(CMachine* m) {
    assert(m);
    auto mach = static_cast<Machine*>(m);
    std::stringstream ss;
    ss << *mach;
    return strdup(ss.str().c_str());
}

CStatus machineCurrentStatus(CMachine* m) {
    auto mach = static_cast<Machine*>(m);
    switch (mach->currentStatus()) {
        case Status::Extensive:
            return STATUS_EXTENSIVE;
        case Status::Error:
            return STATUS_ERROR_STOP;
        case Status::Halted:
            return STATUS_HALT;
        default:
            return STATE_UNKNOWN;
    }
}

struct ReasonConverter {
    CBlockReason operator()(const NotBlocked&) const {
        return CBlockReason{BLOCK_TYPE_NOT_BLOCKED, ByteSlice{nullptr, 0}};
    }

    CBlockReason operator()(const HaltBlocked&) const {
        return CBlockReason{BLOCK_TYPE_HALT, ByteSlice{nullptr, 0}};
    }

    CBlockReason operator()(const ErrorBlocked&) const {
        return CBlockReason{BLOCK_TYPE_ERROR, ByteSlice{nullptr, 0}};
    }

    CBlockReason operator()(const BreakpointBlocked&) const {
        return CBlockReason{BLOCK_TYPE_BREAKPOINT, ByteSlice{nullptr, 0}};
    }

    CBlockReason operator()(const InboxBlocked&) const {
        return CBlockReason{BLOCK_TYPE_INBOX, ByteSlice{nullptr, 0}};
    }

    CBlockReason operator()(const SideloadBlocked&) const {
        return CBlockReason{BLOCK_TYPE_SIDELOAD, ByteSlice{nullptr, 0}};
    }
};

CBlockReason machineIsBlocked(CMachine* m, int newMessages) {
    assert(m);
    auto mach = static_cast<Machine*>(m);
    auto blockReason = mach->isBlocked(newMessages != 0);
    return std::visit(ReasonConverter{}, blockReason);
}

COneStepProof machineMarshallForProof(CMachine* m) {
    assert(m);
    auto mach = static_cast<Machine*>(m);
    auto osp = mach->marshalForProof();
    auto standard = returnCharVector(osp.standard_proof);
    auto buf = returnCharVector(osp.buffer_proof);
    return {standard, buf};
}

ByteSlice machineMarshallState(CMachine* m) {
    assert(m);
    auto mach = static_cast<Machine*>(m);
    std::vector<unsigned char> buffer;
    return returnCharVector(mach->marshalState());
}

CMachineExecutionConfig* machineExecutionConfigCreate() {
    return new MachineExecutionConfig();
}

void machineExecutionConfigDestroy(CMachineExecutionConfig* m) {
    if (m == nullptr) {
        return;
    }
    delete static_cast<MachineExecutionConfig*>(m);
}

void* machineExecutionConfigClone(CMachineExecutionConfig* c) {
    assert(c);
    auto config = static_cast<MachineExecutionConfig*>(c);
    auto cloneConf = new MachineExecutionConfig(*config);
    return static_cast<void*>(cloneConf);
}

void machineExecutionConfigSetMaxGas(CMachineExecutionConfig* c,
                                     uint64_t max_gas,
                                     int go_over_gas) {
    assert(c);
    auto config = static_cast<MachineExecutionConfig*>(c);
    config->max_gas = max_gas;
    config->go_over_gas = go_over_gas;
}

void machineExecutionConfigSetInboxMessages(CMachineExecutionConfig* c,
                                            ByteSliceArray bytes) {
    assert(c);
    auto config = static_cast<MachineExecutionConfig*>(c);
    config->setInboxMessagesFromBytes(receiveByteSliceArray(bytes));
}

void machineExecutionConfigSetSideloads(CMachineExecutionConfig* c,
                                        ByteSliceArray bytes) {
    assert(c);
    auto config = static_cast<MachineExecutionConfig*>(c);
    config->setSideloadsFromBytes(receiveByteSliceArray(bytes));
}

void machineExecutionConfigSetStopOnSideload(CMachineExecutionConfig* c,
                                             int stop_on_sideload) {
    assert(c);
    auto config = static_cast<MachineExecutionConfig*>(c);
    config->stop_on_sideload = stop_on_sideload;
}

RawAssertion executeAssertion(CMachine* m,
                              const CMachineExecutionConfig* c,
                              void* before_send_acc_data,
                              void* before_log_acc_data) {
    assert(m);
    assert(c);
    auto mach = static_cast<Machine*>(m);
    auto config = static_cast<const MachineExecutionConfig*>(c);
    auto before_send_acc = receiveUint256(before_send_acc_data);
    auto before_log_acc = receiveUint256(before_log_acc_data);

    try {
        mach->machine_state.context = AssertionContext{*config};
        mach->machine_state.context.max_gas +=
            mach->machine_state.output.arb_gas_used;
        Assertion assertion = mach->run();
        return makeRawAssertion(assertion, before_send_acc, before_log_acc);
    } catch (const std::exception& e) {
        std::cerr << "Failed to make assertion " << e.what() << "\n";
        return makeEmptyAssertion();
    }
}
