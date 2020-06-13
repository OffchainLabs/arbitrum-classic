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
#include "utils.hpp"

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
    auto machine = static_cast<Machine*>(m);
    auto result =
        machine->checkpoint(*(static_cast<CheckpointStorage*>(storage)));

    return result.status.ok();
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
        marshal_uint256_t(val.timout, inboxDataVec);
        return CBlockReason{BLOCK_TYPE_INBOX, returnCharVector(inboxDataVec)};
    }
};

CBlockReason machineIsBlocked(CMachine* m,
                              void* currentTimeData,
                              int newMessages) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    auto currentTime = receiveUint256(currentTimeData);
    auto blockReason = mach->isBlocked(currentTime, newMessages != 0);
    return nonstd::visit(ReasonConverter{}, blockReason);
}

ByteSlice machineMarshallForProof(CMachine* m) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    std::vector<unsigned char> buffer;
    return returnCharVector(mach->marshalForProof());
}

RawAssertion machineExecuteAssertion(CMachine* m,
                                     uint64_t maxSteps,
                                     void* lowerBoundBlockData,
                                     void* upperBoundBlockData,
                                     void* lowerBoundTimestampData,
                                     void* upperBoundTimestampData,
                                     void* inbox,
                                     uint64_t wallLimit) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    auto lowerBoundBlock = receiveUint256(lowerBoundBlockData);
    auto upperBoundBlock = receiveUint256(upperBoundBlockData);
    auto lowerBoundTimestamp = receiveUint256(lowerBoundTimestampData);
    auto upperBoundTimestamp = receiveUint256(upperBoundTimestampData);

    auto inboxData = reinterpret_cast<const char*>(inbox);
    auto messages = deserialize_value(inboxData, mach->getPool());

    TimeBounds timeBounds{lowerBoundBlock, upperBoundBlock, lowerBoundTimestamp,
                          upperBoundTimestamp};

    Assertion assertion =
        mach->run(maxSteps, timeBounds, nonstd::get<Tuple>(std::move(messages)),
                  std::chrono::seconds{wallLimit});
    std::vector<unsigned char> outMsgData;
    for (const auto& outMsg : assertion.outMessages) {
        mach->marshal_value(outMsg, outMsgData);
    }
    std::vector<unsigned char> logData;
    for (const auto& log : assertion.logs) {
        mach->marshal_value(log, logData);
    }

    return {returnCharVector(outMsgData),
            static_cast<int>(assertion.outMessages.size()),
            returnCharVector(logData),
            static_cast<int>(assertion.logs.size()),
            assertion.stepCount,
            assertion.gasCount,
            assertion.didInboxInsn};
}
