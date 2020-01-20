/*
 * Copyright 2019, Offchain Labs, Inc.
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
#include "bigint_utils.hpp"

#include <avm/machine.hpp>
#include <data_storage/checkpoint/checkpointstorage.hpp>

#include <sys/stat.h>
#include <fstream>
#include <iostream>

typedef struct {
    uint64_t stepCount;
} cassertion;

Machine* read_files(std::string filename) {
    auto machine = new Machine();
    auto sucess = machine->initializeMachine(filename);

    if (sucess) {
        return machine;
    } else {
        return nullptr;
    }
}

// cmachine_t *machine_create(char *data)
CMachine* machineCreate(const char* filename) {
    Machine* mach = read_files(filename);
    return static_cast<void*>(mach);
}

void machineDestroy(CMachine* m) {
    if (m == NULL)
        return;
    delete static_cast<Machine*>(m);
}

int checkpointMachine(CMachine* m, CCheckpointStorage* storage) {
    auto machine = *(static_cast<Machine*>(m));
    auto result =
        machine.checkpoint(*(static_cast<CheckpointStorage*>(storage)));

    return result.status.ok();
}

int restoreMachine(CMachine* m,
                   CCheckpointStorage* storage,
                   const void* machine_hash) {
    auto machine = *(static_cast<Machine*>(m));

    auto machine_hash_ptr = reinterpret_cast<const char*>(machine_hash);
    auto hash = deserializeUint256t(machine_hash_ptr);

    std::vector<unsigned char> hash_vector;
    marshal_uint256_t(hash, hash_vector);

    return machine.restoreCheckpoint(
        *(static_cast<CheckpointStorage*>(storage)), hash_vector);
}

void machineHash(CMachine* m, void* ret) {
    assert(m);
    uint256_t retHash = static_cast<Machine*>(m)->hash();
    std::array<unsigned char, 32> val;
    to_big_endian(retHash, val.begin());
    std::copy(val.begin(), val.end(), reinterpret_cast<char*>(ret));
}

void* machineClone(CMachine* m) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    Machine* cloneMach = new Machine(*mach);
    return static_cast<void*>(cloneMach);
}

void machinePrint(CMachine* m) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    std::cout << "Machine info\n" << *mach << std::endl;
}

CStatus machineCurrentStatus(CMachine* m) {
    Machine* mach = static_cast<Machine*>(m);
    switch (mach->currentStatus()) {
        case Status::Extensive:
            return STATUS_EXTENSIVE;
        case Status::Error:
            return STATUS_ERROR_STOP;
        case Status::Halted:
            return STATUS_HALT;
        default:
            throw std::runtime_error("Bad machine status type");
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

    CBlockReason operator()(const InboxBlocked& val) const {
        std::vector<unsigned char> inboxDataVec;
        marshal_value(val.timout, inboxDataVec);
        unsigned char* cInboxData = (unsigned char*)malloc(inboxDataVec.size());
        std::copy(inboxDataVec.begin(), inboxDataVec.end(), cInboxData);
        return CBlockReason{
            BLOCK_TYPE_INBOX,
            ByteSlice{cInboxData, static_cast<int>(inboxDataVec.size())}};
    }
};

CBlockReason machineIsBlocked(CMachine* m,
                              void* currentTimeData,
                              int newMessages) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    auto currentTimePtr = reinterpret_cast<const char*>(currentTimeData);
    auto currentTime = deserializeUint256t(currentTimePtr);
    auto blockReason = mach->isBlocked(currentTime, newMessages != 0);
    return nonstd::visit(ReasonConverter{}, blockReason);
}

ByteSlice machineMarshallForProof(CMachine* m) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    std::vector<unsigned char> buffer;
    auto proof = mach->marshalForProof();
    auto proofData = (unsigned char*)malloc(proof.size());
    std::copy(proof.begin(), proof.end(), proofData);
    auto voidData = reinterpret_cast<void*>(proofData);
    return {voidData, static_cast<int>(proof.size())};
}

RawAssertion machineExecuteAssertion(CMachine* m,
                                     uint64_t maxSteps,
                                     void* timeboundStartData,
                                     void* timeboundEndData,
                                     void* inbox) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    auto timeboundStartPtr = reinterpret_cast<const char*>(timeboundStartData);
    auto timeboundStart = deserializeUint256t(timeboundStartPtr);
    auto timeboundEndPtr = reinterpret_cast<const char*>(timeboundEndData);
    auto timeboundEnd = deserializeUint256t(timeboundEndPtr);

    auto inboxData = reinterpret_cast<const char*>(inbox);
    auto messages = deserialize_value(inboxData, mach->getPool());

    Assertion assertion = mach->run(maxSteps, timeboundStart, timeboundEnd,
                                    nonstd::get<Tuple>(std::move(messages)));
    std::vector<unsigned char> outMsgData;
    for (const auto& outMsg : assertion.outMessages) {
        marshal_value(outMsg, outMsgData);
    }
    std::vector<unsigned char> logData;
    for (const auto& log : assertion.logs) {
        marshal_value(log, logData);
    }

    unsigned char* cMessageData = (unsigned char*)malloc(outMsgData.size());
    std::copy(outMsgData.begin(), outMsgData.end(), cMessageData);

    unsigned char* cLogData = (unsigned char*)malloc(logData.size());
    std::copy(logData.begin(), logData.end(), cLogData);

    return {cMessageData,
            static_cast<int>(outMsgData.size()),
            static_cast<int>(assertion.outMessages.size()),
            cLogData,
            static_cast<int>(logData.size()),
            static_cast<int>(assertion.logs.size()),
            assertion.stepCount,
            assertion.gasCount,
            assertion.didInboxInsn};
}
