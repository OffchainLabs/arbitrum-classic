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
#include "utils.hpp"

#include <avm/machine.hpp>
#include <data_storage/checkpointstorage.hpp>
#include <data_storage/value/machine.hpp>

#include <sys/stat.h>
#include <fstream>
#include <iostream>

typedef struct {
    uint64_t stepCount;
} cassertion;

Machine* read_files(std::string filename) {
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

int checkpointMachine(CMachine* m, CCheckpointStorage* s) {
    auto machine = static_cast<Machine*>(m);
    auto storage = static_cast<CheckpointStorage*>(s);
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
    Machine* mach = static_cast<Machine*>(m);
    auto blockReason = mach->isBlocked(newMessages != 0);
    return nonstd::visit(ReasonConverter{}, blockReason);
}

ByteSlice machineMarshallForProof(CMachine* m) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    std::vector<unsigned char> buffer;
    return returnCharVector(mach->marshalForProof());
}

ByteSlice machineMarshallState(CMachine* m) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    std::vector<unsigned char> buffer;
    return returnCharVector(mach->marshalState());
}

RawAssertion makeRawAssertion(Assertion& assertion) {
    std::vector<unsigned char> outMsgData;
    for (const auto& outMsg : assertion.outMessages) {
        marshal_value(outMsg, outMsgData);
    }
    std::vector<unsigned char> logData;
    for (const auto& log : assertion.logs) {
        marshal_value(log, logData);
    }

    return {assertion.inbox_messages_consumed,
            returnCharVector(outMsgData),
            static_cast<int>(assertion.outMessages.size()),
            returnCharVector(logData),
            static_cast<int>(assertion.logs.size()),
            assertion.stepCount,
            assertion.gasCount};
}

RawAssertion makeEmptyAssertion() {
    return {0, returnCharVector(std::vector<char>{}),
            0, returnCharVector(std::vector<char>{}),
            0, 0,
            0};
}

Tuple getTuple(TuplePool& pool, void* data) {
    auto charData = reinterpret_cast<const char*>(data);
    return nonstd::get<Tuple>(deserialize_value(charData, pool));
}

std::vector<Tuple> getInboxMessages(TuplePool& pool,
                                    void* data,
                                    uint64_t message_count) {
    auto charData = reinterpret_cast<const char*>(data);
    std::vector<Tuple> messages;
    for (uint64_t i = 0; i < message_count; ++i) {
        messages.push_back(deserialize_value(charData, pool).get<Tuple>());
    }
    return messages;
}

RawAssertion executeAssertion(CMachine* m,
                              uint64_t maxSteps,
                              void* inbox_messages,
                              uint64_t message_count,
                              uint64_t wallLimit) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);
    auto messages =
        getInboxMessages(mach->getPool(), inbox_messages, message_count);

    try {
        Assertion assertion = mach->run(maxSteps, std::move(messages),
                                        std::chrono::seconds{wallLimit});
        return makeRawAssertion(assertion);
    } catch (const std::exception& e) {
        std::cerr << "Failed to make assertion " << e.what() << "\n";
        return makeEmptyAssertion();
    }
}

RawAssertion executeCallServerAssertion(CMachine* m,
                                        uint64_t maxSteps,
                                        void* inbox_messages,
                                        uint64_t message_count,
                                        void* fake_inbox_peak_value,
                                        uint64_t wallLimit) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);

    auto messages =
        getInboxMessages(mach->getPool(), inbox_messages, message_count);
    auto fake_inbox_peak_value_data =
        reinterpret_cast<const char*>(fake_inbox_peak_value);
    auto fake_inbox_peak =
        deserialize_value(fake_inbox_peak_value_data, mach->getPool());

    try {
        Assertion assertion = mach->runCallServer(
            maxSteps, std::move(messages), std::chrono::seconds{wallLimit},
            std::move(fake_inbox_peak));
        return makeRawAssertion(assertion);
    } catch (const std::exception& e) {
        std::cerr << "Failed to make assertion " << e.what() << "\n";
        return makeEmptyAssertion();
    }
}

RawAssertion executeSideloadedAssertion(CMachine* m,
                                        uint64_t maxSteps,
                                        void* inbox_messages,
                                        uint64_t message_count,
                                        void* sideload,
                                        uint64_t wallLimit) {
    assert(m);
    Machine* mach = static_cast<Machine*>(m);

    auto messages =
        getInboxMessages(mach->getPool(), inbox_messages, message_count);
    auto sideload_value = getTuple(mach->getPool(), sideload);

    try {
        Assertion assertion = mach->runSideloaded(
            maxSteps, std::move(messages), std::chrono::seconds{wallLimit},
            std::move(sideload_value));
        return makeRawAssertion(assertion);
    } catch (const std::exception& e) {
        std::cerr << "Failed to make assertion " << e.what() << "\n";
        return makeEmptyAssertion();
    }
}
