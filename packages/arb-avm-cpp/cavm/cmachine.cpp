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

int checkpointMachine(CMachine* m, CArbStorage* s) {
    auto machine = static_cast<Machine*>(m);
    auto storage = static_cast<ArbStorage*>(s);
    auto transaction = storage->makeTransaction();
    auto result = saveMachine(*transaction, *machine);
    if (!result.status.ok()) {
        return false;
    }
    return transaction->commit().ok();
}

void machineHash(CMachine* m, void* ret) {
    assert(m);
    uint256_t retHash = static_cast<Machine*>(m)->hash();
    std::array<unsigned char, 32> val{};
    to_big_endian(retHash, val.begin());
    std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));
}

void* machineClone(CMachine* m) {
    assert(m);
    auto mach = static_cast<Machine*>(m);
    auto cloneMach = new Machine(*mach);
    return static_cast<void*>(cloneMach);
}

void machinePrint(CMachine* m) {
    assert(m);
    auto mach = static_cast<Machine*>(m);
    std::cout << "Machine info\n" << *mach << std::endl;
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
    return nonstd::visit(ReasonConverter{}, blockReason);
}

ByteSlice machineMarshallForProof(CMachine* m) {
    assert(m);
    auto mach = static_cast<Machine*>(m);
    std::vector<unsigned char> buffer;
    return returnCharVector(mach->marshalForProof());
}

ByteSlice machineMarshallBufferProof(CMachine* m) {
    assert(m);
    auto mach = static_cast<Machine*>(m);
    std::vector<unsigned char> buffer;
    return returnCharVector(mach->marshalBufferProof());
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

void machineExecutionConfigSetFinalMessageOfBlock(CMachineExecutionConfig* c,
                                                  int final_message_of_block) {
    assert(c);
    auto config = static_cast<MachineExecutionConfig*>(c);
    config->final_message_of_block = final_message_of_block;
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

RawAssertion executeAssertion(CMachine* m, const CMachineExecutionConfig* c) {
    assert(m);
    assert(c);
    auto mach = static_cast<Machine*>(m);
    auto config = static_cast<const MachineExecutionConfig*>(c);

    try {
        Assertion assertion = mach->run(*config);
        return makeRawAssertion(assertion);
    } catch (const std::exception& e) {
        std::cerr << "Failed to make assertion " << e.what() << "\n";
        return makeEmptyAssertion();
    }
}
