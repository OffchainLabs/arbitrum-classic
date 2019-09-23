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

#include <avm/datastack.hpp>
#include <avm/tokenTracker.hpp>
#include <avm/value.hpp>
#include <memory>
#include <vector>
#include "avm/checkpointutils.hpp"
#include "avm/exceptions.hpp"
#include "avm/machinestate.hpp"

struct Assertion {
    uint64_t stepCount;
    std::vector<Message> outMessages;
    std::vector<value> logs;
};

class Machine {
    MachineState m;

    friend std::ostream& operator<<(std::ostream&, const Machine&);
    void runOne();

   public:
    Machine() = default;
    Machine(std::string filename);
    bool deserialize(char* data) { return m.deserialize(data); }

    Assertion run(uint64_t stepCount,
                  uint64_t timeBoundStart,
                  uint64_t timeBoundEnd);

    Status currentStatus() { return m.state; }
    BlockReason lastBlockReason() { return m.blockReason; }
    uint256_t hash() const { return m.hash(); }
    std::vector<unsigned char> marshalForProof() { return m.marshalForProof(); }
    uint64_t pendingMessageCount() const { return m.pendingMessageCount(); }

    bool canSpend(const TokenType& tokType, const uint256_t& amount) const {
        return m.balance.canSpend(tokType, amount);
    }
    uint256_t inboxHash() const { return ::hash(m.inbox.messages); }

    void sendOnchainMessage(const Message& msg);
    void deliverOnchainMessages();
    void sendOffchainMessages(const std::vector<Message>& messages);

    TuplePool& getPool() { return *m.pool; }

    // should this be a tuple or some struct?
    CheckpointData getCheckPointTuple();
};

std::ostream& operator<<(std::ostream& os, const MachineState& val);
std::ostream& operator<<(std::ostream& os, const Machine& val);

#endif /* machine_hpp */
