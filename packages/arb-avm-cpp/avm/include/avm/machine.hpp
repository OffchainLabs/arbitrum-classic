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

#ifndef machine_hpp
#define machine_hpp

#include <memory>
#include <vector>

#include <avm/checkpointstorage.hpp>
#include <avm/exceptions.hpp>
#include <avm/machinestate/datastack.hpp>
#include <avm/machinestate/machinestate.hpp>
#include <avm/machinestate/tokenTracker.hpp>
#include <avm/value/value.hpp>

struct Assertion {
    uint64_t stepCount;
    std::vector<Message> outMessages;
    std::vector<value> logs;
};

class Machine {
    MachineState machine_state;

    friend std::ostream& operator<<(std::ostream&, const Machine&);
    void runOne();

   public:
    Machine() = default;
    Machine(const std::string& filename);
    bool deserialize(char* data) { return machine_state.deserialize(data); }

    Assertion run(uint64_t stepCount,
                  uint64_t timeBoundStart,
                  uint64_t timeBoundEnd);

    Status currentStatus() { return machine_state.state; }
    BlockReason lastBlockReason() { return machine_state.blockReason; }
    uint256_t hash() const { return machine_state.hash(); }
    std::vector<unsigned char> marshalForProof() {
        return machine_state.marshalForProof();
    }
    uint64_t pendingMessageCount() const {
        return machine_state.pendingMessageCount();
    }

    bool canSpend(const TokenType& tokType, const uint256_t& amount) const {
        return machine_state.balance.canSpend(tokType, amount);
    }
    uint256_t inboxHash() const { return ::hash(machine_state.inbox.messages); }

    void sendOnchainMessage(const Message& msg);
    void deliverOnchainMessages();
    void sendOffchainMessages(const std::vector<Message>& messages);

    TuplePool& getPool() { return *machine_state.pool; }

    SaveResults checkpoint(CheckpointStorage& storage);
    bool restoreCheckpoint(CheckpointStorage& storage,
                           const std::vector<unsigned char>& checkpoint_key);
    DeleteResults deleteCheckpoint(
        CheckpointStorage& storage,
        const std::vector<unsigned char>& checkpoint_key);
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
