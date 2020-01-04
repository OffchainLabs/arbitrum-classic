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

#ifndef machinestate_hpp
#define machinestate_hpp

#include <avm/machinestate/blockreason.hpp>
#include <avm/machinestate/datastack.hpp>
#include <avm/machinestate/messagestack.hpp>
#include <avm_values/value.hpp>

#include <memory>
#include <vector>

enum class Status { Extensive, Halted, Error };

typedef std::array<uint256_t, 2> TimeBounds;

class CheckpointStorage;

struct AssertionContext {
    uint32_t numSteps;
    uint64_t numGas;
    TimeBounds timeBounds;
    std::vector<value> outMessage;
    std::vector<value> logs;

    explicit AssertionContext(const TimeBounds& tb)
        : numSteps{0}, numGas{0}, timeBounds(tb) {}
};

struct MachineState {
    std::shared_ptr<TuplePool> pool;
    std::vector<CodePoint> code;
    value staticVal;
    value registerVal;
    Datastack stack;
    Datastack auxstack;
    Status state = Status::Extensive;
    uint64_t pc = 0;
    CodePoint errpc;
    MessageStack pendingInbox;
    AssertionContext context;
    MessageStack inbox;
    BlockReason blockReason;

    MachineState();
    MachineState(const std::vector<CodePoint>& code_,
                 const value& static_val_,
                 std::shared_ptr<TuplePool>& pool_);
    bool initialize_machinestate(const std::string& contract_filename);

    void readInbox(char* newInbox);
    std::vector<unsigned char> marshalForProof();
    uint64_t pendingMessageCount() const;
    void sendOnchainMessage(const Message& msg);
    void deliverOnchainMessages();
    void sendOffchainMessages(const std::vector<Message>& messages);
    BlockReason runOp(OpCode opcode);
    uint256_t hash() const;
    void setInbox(MessageStack ms);
    void setPendingInbox(MessageStack ms);
    SaveResults checkpointState(CheckpointStorage& storage);
    bool restoreCheckpoint(const CheckpointStorage& storage,
                           const std::vector<unsigned char>& checkpoint_key);
};

#endif /* machinestate_hpp */
